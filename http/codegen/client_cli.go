package codegen

import (
	"fmt"
	"path/filepath"

	"goa.design/goa/v3/codegen"
	"goa.design/goa/v3/codegen/cli"
	"goa.design/goa/v3/expr"
)

// commandData wraps the common CommandData and adds HTTP-specific fields.
type commandData struct {
	*cli.CommandData
	// Subcommands is the list of endpoint commands.
	Subcommands []*subcommandData
	// NeedStream if true initializes the websocket dialer.
	NeedStream bool
}

// commandData wraps the common SubcommandData and adds HTTP-specific fields.
type subcommandData struct {
	*cli.SubcommandData
	// MultipartFuncName is the function name used to render a multipart request encoder.
	MultipartFuncName string
	// MultipartFuncName is the variable name used to render a multipart request encoder.
	MultipartVarName string
}

// ClientCLIFiles returns the client HTTP CLI support file.
func ClientCLIFiles(genpkg string, root *expr.RootExpr) []*codegen.File {
	if len(root.API.HTTP.Services) == 0 {
		return nil
	}
	var (
		data []*commandData
		svcs []*expr.HTTPServiceExpr
	)
	for _, svc := range root.API.HTTP.Services {
		sd := HTTPServices.Get(svc.Name())
		if len(sd.Endpoints) > 0 {
			command := &commandData{
				CommandData: cli.BuildCommandData(sd.Service),
				NeedStream:  streamingEndpointExists(sd),
			}

			for _, e := range sd.Endpoints {
				sub := buildSubcommandData(sd, e)
				command.Subcommands = append(command.Subcommands, sub)
				command.CommandData.Subcommands = append(command.CommandData.Subcommands, sub.SubcommandData)
			}

			command.Example = command.Subcommands[0].Example

			data = append(data, command)
			svcs = append(svcs, svc)
		}
	}
	var files []*codegen.File
	for _, svr := range root.API.Servers {
		files = append(files, endpointParser(genpkg, root, svr, data))
	}
	for i, svc := range svcs {
		files = append(files, payloadBuilders(genpkg, svc, data[i].CommandData))
	}
	return files
}

func buildSubcommandData(sd *ServiceData, e *EndpointData) *subcommandData {
	flags, buildFunction := buildFlags(sd, e)

	sub := &subcommandData{
		SubcommandData: cli.BuildSubcommandData(sd.Service.Name, e.Method, buildFunction, flags),
	}
	if e.MultipartRequestEncoder != nil {
		sub.MultipartVarName = e.MultipartRequestEncoder.VarName
		sub.MultipartFuncName = e.MultipartRequestEncoder.FuncName
	}
	return sub
}

// endpointParser returns the file that implements the command line parser that
// builds the client endpoint and payload necessary to perform a request.
func endpointParser(genpkg string, root *expr.RootExpr, svr *expr.ServerExpr, data []*commandData) *codegen.File {
	pkg := codegen.SnakeCase(codegen.Goify(svr.Name, true))
	path := filepath.Join(codegen.Gendir, "http", "cli", pkg, "cli.go")
	title := fmt.Sprintf("%s HTTP client CLI support package", svr.Name)
	specs := []*codegen.ImportSpec{
		{Path: "encoding/json"},
		{Path: "flag"},
		{Path: "fmt"},
		{Path: "net/http"},
		{Path: "os"},
		{Path: "strconv"},
		{Path: "unicode/utf8"},
		codegen.GoaImport(""),
		codegen.GoaNamedImport("http", "goahttp"),
	}
	for _, sv := range svr.Services {
		svc := root.Service(sv)
		sd := HTTPServices.Get(svc.Name)
		if sd == nil {
			continue
		}
		specs = append(specs, &codegen.ImportSpec{
			Path: genpkg + "/http/" + codegen.SnakeCase(sd.Service.VarName) + "/client",
			Name: sd.Service.PkgName + "c",
		})
	}

	cliData := make([]*cli.CommandData, len(data))
	for i, cmd := range data {
		cliData[i] = cmd.CommandData
	}

	sections := []*codegen.SectionTemplate{
		codegen.Header(title, "cli", specs),
		cli.UsageCommands(cliData),
		cli.UsageExamples(cliData),
		{
			Name:   "parse-endpoint",
			Source: parseEndpointT,
			Data: struct {
				FlagsCode string
				Commands  []*commandData
			}{
				cli.FlagsCode(cliData),
				data,
			},
			FuncMap: map[string]interface{}{
				"streamingCmdExists": streamingCmdExists,
			},
		},
	}
	for _, cmd := range cliData {
		sections = append(sections, cli.CommandUsage(cmd))
	}
	return &codegen.File{Path: path, SectionTemplates: sections}
}

// payloadBuilders returns the file that contains the payload constructors that
// use flag values as arguments.
func payloadBuilders(genpkg string, svc *expr.HTTPServiceExpr, data *cli.CommandData) *codegen.File {
	sd := HTTPServices.Get(svc.Name())
	path := filepath.Join(codegen.Gendir, "http", codegen.SnakeCase(sd.Service.VarName), "client", "cli.go")
	title := fmt.Sprintf("%s HTTP client CLI support package", svc.Name())
	specs := []*codegen.ImportSpec{
		{Path: "encoding/json"},
		{Path: "fmt"},
		{Path: "net/http"},
		{Path: "os"},
		{Path: "strconv"},
		{Path: "unicode/utf8"},
		codegen.GoaImport(""),
		codegen.GoaNamedImport("http", "goahttp"),
		{Path: genpkg + "/" + codegen.SnakeCase(sd.Service.VarName), Name: sd.Service.PkgName},
	}
	sections := []*codegen.SectionTemplate{
		codegen.Header(title, "client", specs),
	}
	for _, sub := range data.Subcommands {
		if sub.BuildFunction != nil {
			sections = append(sections, cli.PayloadBuilderSection(sub.BuildFunction))
		}
	}

	return &codegen.File{Path: path, SectionTemplates: sections}
}

func buildFlags(svc *ServiceData, e *EndpointData) ([]*cli.FlagData, *cli.BuildFunctionData) {
	var (
		flags         []*cli.FlagData
		buildFunction *cli.BuildFunctionData
	)

	svcn := svc.Service.Name
	en := e.Method.Name
	if e.Payload != nil {
		if e.Payload.Request.PayloadInit != nil {
			args := e.Payload.Request.PayloadInit.ClientArgs
			args = append(args, e.Payload.Request.PayloadInit.CLIArgs...)
			flags, buildFunction = makeFlags(e, args)
		} else if e.Payload.Ref != "" {
			flags = append(flags, cli.NewFlagData(svcn, en, "p", e.Method.PayloadRef, e.Method.PayloadDesc, true, e.Method.PayloadEx))
		}
	}

	return flags, buildFunction
}

func makeFlags(e *EndpointData, args []*InitArgData) ([]*cli.FlagData, *cli.BuildFunctionData) {
	var (
		fdata     []*cli.FieldData
		flags     = make([]*cli.FlagData, len(args))
		params    = make([]string, len(args))
		pInitArgs = make([]*codegen.InitArgData, len(args))
		check     bool
	)
	for i, arg := range args {
		pInitArgs[i] = &codegen.InitArgData{
			Name:         arg.Name,
			Pointer:      arg.Pointer,
			FieldName:    arg.FieldName,
			FieldPointer: arg.FieldPointer,
			FieldType:    arg.FieldType,
			Type:         arg.Type,
		}

		f := cli.NewFlagData(e.ServiceName, e.Method.Name, arg.Name, arg.TypeName, arg.Description, arg.Required, arg.Example)
		flags[i] = f
		params[i] = f.FullName
		if arg.FieldName == "" && arg.Name != "body" {
			continue
		}
		code, chek := cli.FieldLoadCode(f, arg.Name, arg.TypeName, arg.Validate, arg.DefaultValue)
		check = check || chek
		tn := arg.TypeRef
		if f.Type == "JSON" {
			// We need to declare the variable without
			// a pointer to be able to unmarshal the JSON
			// using its address.
			tn = arg.TypeName
		}
		fdata = append(fdata, &cli.FieldData{
			Name:    arg.Name,
			VarName: arg.Name,
			TypeRef: tn,
			Init:    code,
		})
	}

	pInit := cli.PayloadInitData{
		Code:                e.Payload.Request.PayloadInit.ClientCode,
		ReturnTypeAttribute: e.Payload.Request.PayloadInit.ReturnTypeAttribute,
		ReturnIsStruct:      e.Payload.Request.PayloadInit.ReturnIsStruct,
		ReturnTypeName:      e.Payload.Request.PayloadInit.ReturnTypeName,
		ReturnTypePkg:       e.Payload.Request.PayloadInit.ReturnTypePkg,
		Args:                pInitArgs,
	}

	return flags, &cli.BuildFunctionData{
		Name:         "Build" + e.Method.VarName + "Payload",
		ActualParams: params,
		FormalParams: params,
		ServiceName:  e.ServiceName,
		MethodName:   e.Method.Name,
		ResultType:   e.Payload.Ref,
		Fields:       fdata,
		PayloadInit:  &pInit,
		CheckErr:     check,
	}
}

// streamingCmdExists returns true if at least one command in the list of commands
// uses stream for sending payload/result.
func streamingCmdExists(data []*commandData) bool {
	for _, c := range data {
		if c.NeedStream {
			return true
		}
	}
	return false
}

const parseEndpointT = `// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
	{{- if streamingCmdExists .Commands }}
	dialer goahttp.Dialer,
		{{- range .Commands }}
			{{- if .NeedStream }}
				{{ .VarName }}Configurer *{{ .PkgName }}.ConnConfigurer,
			{{- end }}
		{{- end }}
	{{- end }}
	{{- range $c := .Commands }}
	{{- range .Subcommands }}
		{{- if .MultipartVarName }}
	{{ .MultipartVarName }} {{ $c.PkgName }}.{{ .MultipartFuncName }},
		{{- end }}
	{{- end }}
	{{- end }}
) (goa.Endpoint, interface{}, error) {
	{{ .FlagsCode }}
    var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
	{{- range .Commands }}
		case "{{ .Name }}":
			c := {{ .PkgName }}.NewClient(scheme, host, doer, enc, dec, restore{{ if .NeedStream }}, dialer, {{ .VarName }}Configurer{{ end }})
			switch epn {
		{{- $pkgName := .PkgName }}{{ range .Subcommands }}
			case "{{ .Name }}":
				endpoint = c.{{ .MethodVarName }}({{ if .MultipartVarName }}{{ .MultipartVarName }}{{ end }})
			{{- if .BuildFunction }}
				data, err = {{ $pkgName}}.{{ .BuildFunction.Name }}({{ range .BuildFunction.ActualParams }}*{{ . }}Flag, {{ end }})
			{{- else if .Conversion }}
				{{ .Conversion }}
			{{- else }}
				data = nil
			{{- end }}
		{{- end }}
			}
	{{- end }}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}
`
