package codegen

import (
	"path"
	"path/filepath"

	"goa.design/goa/v3/codegen"
	"goa.design/goa/v3/codegen/cli"
	"goa.design/goa/v3/expr"
)

// ClientCLIFiles returns the CLI files to generate a command-line client that
// makes gRPC requests.
func ClientCLIFiles(genpkg string, root *expr.RootExpr) []*codegen.File {
	if len(root.API.GRPC.Services) == 0 {
		return nil
	}
	var (
		data []*cli.CommandData
		svcs []*expr.GRPCServiceExpr
	)
	{
		for _, svc := range root.API.GRPC.Services {
			if len(svc.GRPCEndpoints) == 0 {
				continue
			}
			sd := GRPCServices.Get(svc.Name())
			command := cli.BuildCommandData(sd.Service)
			for _, e := range sd.Endpoints {
				flags, buildFunction := buildFlags(sd, e)
				subcmd := cli.BuildSubcommandData(sd.Service.Name, e.Method, buildFunction, flags)
				command.Subcommands = append(command.Subcommands, subcmd)
			}
			command.Example = command.Subcommands[0].Example
			data = append(data, command)
			svcs = append(svcs, svc)
		}
	}
	var files []*codegen.File
	{
		for _, svr := range root.API.Servers {
			files = append(files, endpointParser(genpkg, root, svr, data))
		}
		for i, svc := range svcs {
			files = append(files, payloadBuilders(genpkg, svc, data[i]))
		}
	}
	return files
}

// endpointParser returns the file that implements the command line parser that
// builds the client endpoint and payload necessary to perform a request.
func endpointParser(genpkg string, root *expr.RootExpr, svr *expr.ServerExpr, data []*cli.CommandData) *codegen.File {
	pkg := codegen.SnakeCase(codegen.Goify(svr.Name, true))
	fpath := filepath.Join(codegen.Gendir, "grpc", "cli", pkg, "cli.go")
	title := svr.Name + " gRPC client CLI support package"
	specs := []*codegen.ImportSpec{
		{Path: "context"},
		{Path: "flag"},
		{Path: "fmt"},
		{Path: "os"},
		codegen.GoaImport(""),
		codegen.GoaNamedImport("grpc", "goagrpc"),
		{Path: "google.golang.org/grpc", Name: "grpc"},
	}
	for _, svc := range root.API.GRPC.Services {
		sd := GRPCServices.Get(svc.Name())
		if sd == nil {
			continue
		}
		svcName := codegen.SnakeCase(sd.Service.VarName)
		specs = append(specs, &codegen.ImportSpec{
			Path: path.Join(genpkg, "grpc", svcName, "client"),
			Name: sd.Service.PkgName + "c",
		})
		specs = append(specs, &codegen.ImportSpec{
			Path: path.Join(genpkg, "grpc", svcName, pbPkgName),
			Name: svcName + pbPkgName,
		})
	}

	sections := []*codegen.SectionTemplate{
		codegen.Header(title, "cli", specs),
		cli.UsageCommands(data),
		cli.UsageExamples(data),
		{
			Name:   "parse-endpoint",
			Source: parseEndpointT,
			Data: struct {
				FlagsCode string
				Commands  []*cli.CommandData
			}{
				cli.FlagsCode(data),
				data,
			},
		},
	}
	for _, cmd := range data {
		sections = append(sections, cli.CommandUsage(cmd))
	}
	return &codegen.File{Path: fpath, SectionTemplates: sections}
}

// payloadBuilders returns the file that contains the payload constructors that
// use flag values as arguments.
func payloadBuilders(genpkg string, svc *expr.GRPCServiceExpr, data *cli.CommandData) *codegen.File {
	sd := GRPCServices.Get(svc.Name())
	svcName := codegen.SnakeCase(sd.Service.VarName)
	fpath := filepath.Join(codegen.Gendir, "grpc", svcName, "client", "cli.go")
	title := svc.Name() + " gRPC client CLI support package"
	specs := []*codegen.ImportSpec{
		{Path: "encoding/json"},
		{Path: "fmt"},
		{Path: path.Join(genpkg, svcName), Name: sd.Service.PkgName},
		{Path: path.Join(genpkg, "grpc", svcName, pbPkgName), Name: sd.PkgName},
	}
	sections := []*codegen.SectionTemplate{
		codegen.Header(title, "client", specs),
	}
	for _, sub := range data.Subcommands {
		if sub.BuildFunction != nil {
			sections = append(sections, cli.PayloadBuilderSection(sub.BuildFunction))
		}
	}
	return &codegen.File{Path: fpath, SectionTemplates: sections}
}

func buildFlags(svc *ServiceData, e *EndpointData) ([]*cli.FlagData, *cli.BuildFunctionData) {
	var (
		flags         []*cli.FlagData
		buildFunction *cli.BuildFunctionData
	)
	{
		if e.Request != nil {
			flags, buildFunction = makeFlags(e, e.Request.CLIArgs)
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
		pinit     *cli.PayloadInitData
	)
	for i, arg := range args {
		pInitArgs[i] = &codegen.InitArgData{
			Name:      arg.Name,
			FieldName: arg.FieldName,
			FieldType: arg.FieldType,
			Type:      arg.Type,
		}

		f := cli.NewFlagData(e.ServiceName, e.Method.Name, arg.Name, arg.TypeName, arg.Description, arg.Required, arg.Example)
		flags[i] = f
		params[i] = f.FullName
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
	if e.Method.PayloadRef == "" {
		return flags, nil
	}
	if e.Request.ServerConvert != nil {
		pinit = &cli.PayloadInitData{
			Code:           e.Request.ServerConvert.Init.Code,
			ReturnIsStruct: e.Request.ServerConvert.Init.ReturnIsStruct,
			ReturnTypePkg:  e.Request.ServerConvert.Init.ReturnTypePkg,
			Args:           pInitArgs,
		}
	}

	return flags, &cli.BuildFunctionData{
		Name:         "Build" + e.Method.VarName + "Payload",
		ActualParams: params,
		FormalParams: params,
		ServiceName:  e.ServiceName,
		MethodName:   e.Method.Name,
		ResultType:   e.PayloadRef,
		Fields:       fdata,
		PayloadInit:  pinit,
		CheckErr:     check,
	}
}

const parseEndpointT = `// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(cc *grpc.ClientConn, opts ...grpc.CallOption) (goa.Endpoint, interface{}, error) {
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
			c := {{ .PkgName }}.NewClient(cc, opts...)
			switch epn {
		{{- $pkgName := .PkgName }}{{ range .Subcommands }}
			case "{{ .Name }}":
				endpoint = c.{{ .MethodVarName }}()
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
