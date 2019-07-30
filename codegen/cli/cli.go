// Package cli contains helpers used by transport-specific command-line client
// generators for parsing the command-line flags to identify the service and
// the method to make a request along with the request payload to be sent.
package cli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"goa.design/goa/v3/codegen"
	"goa.design/goa/v3/codegen/service"
	"goa.design/goa/v3/expr"
)

type (
	// CommandData contains the data needed to render a command.
	CommandData struct {
		// Name of command e.g. "cellar-storage"
		Name string
		// VarName is the name of the command variable e.g.
		// "cellarStorage"
		VarName string
		// Description is the help text.
		Description string
		// Subcommands is the list of endpoint commands.
		Subcommands []*SubcommandData
		// Example is a valid command invocation, starting with the
		// command name.
		Example string
		// PkgName is the service HTTP client package import name,
		// e.g. "storagec".
		PkgName string
	}

	// SubcommandData contains the data needed to render a sub-command.
	SubcommandData struct {
		// Name is the sub-command name e.g. "add"
		Name string
		// FullName is the sub-command full name e.g. "storageAdd"
		FullName string
		// Description is the help text.
		Description string
		// Flags is the list of flags supported by the subcommand.
		Flags []*FlagData
		// MethodVarName is the endpoint method name, e.g. "Add"
		MethodVarName string
		// BuildFunction contains the data to generate a payload builder function
		// if any. Exclusive with Conversion.
		BuildFunction *BuildFunctionData
		// Conversion contains the flag value to payload conversion function if
		// any. Exclusive with BuildFunction.
		Conversion string
		// Example is a valid command invocation, starting with the command name.
		Example string
	}

	// FlagData contains the data needed to render a command-line flag.
	FlagData struct {
		// Name is the name of the flag, e.g. "list-vintage"
		Name string
		// VarName is the name of the flag variable, e.g. "listVintage"
		VarName string
		// Type is the type of the flag, e.g. INT
		Type string
		// FullName is the flag full name e.g. "storageAddVintage"
		FullName string
		// Description is the flag help text.
		Description string
		// Required is true if the flag is required.
		Required bool
		// Example returns a JSON serialized example value.
		Example string
	}

	// BuildFunctionData contains the data needed to generate a constructor
	// function that builds a service method payload type from the command-line
	// flags.
	BuildFunctionData struct {
		// Name is the build payload function name.
		Name string
		// Description describes the payload function.
		Description string
		// ActualParams is the list of passed build function parameters.
		ActualParams []string
		// FormalParams is the list of build function formal parameter
		// names.
		FormalParams []string
		// ServiceName is the name of the service.
		ServiceName string
		// MethodName is the name of the method.
		MethodName string
		// ResultType is the fully qualified payload type name.
		ResultType string
		// Fields describes the payload fields.
		Fields []*FieldData
		// PayloadInit contains the data needed to render the function
		// body.
		PayloadInit *PayloadInitData
		// CheckErr is true if the payload initialization code requires an
		// "err error" variable that must be checked.
		CheckErr bool
	}

	// FieldData contains the data needed to generate the code that initializes a
	// field in the method payload type.
	FieldData struct {
		// Name is the field name, e.g. "Vintage"
		Name string
		// VarName is the name of the local variable holding the field
		// value, e.g. "vintage"
		VarName string
		// TypeRef is the reference to the type.
		TypeRef string
		// Init is the code initializing the variable.
		Init string
	}

	// PayloadInitData contains the data needed to generate a constructor
	// function that initializes a service method payload type from the
	// command-ling arguments.
	PayloadInitData struct {
		// Code is the payload initialization code.
		Code string
		// ReturnTypeAttribute if non-empty returns an attribute in the payload
		// type that describes the shape of the method payload.
		ReturnTypeAttribute string
		// ReturnIsStruct if true indicates that the method payload is an object.
		ReturnIsStruct bool
		// ReturnTypeName is the fully-qualified name of the payload.
		ReturnTypeName string
		// ReturnTypePkg is the package name where the payload is present.
		ReturnTypePkg string
		// Args is the list of arguments for the constructor.
		Args []*codegen.InitArgData
	}
)

// BuildCommandData builds the data needed by CLI code generators to render the
// parsing of the service command.
func BuildCommandData(data *service.Data) *CommandData {
	description := data.Description
	if description == "" {
		description = fmt.Sprintf("Make requests to the %q service", data.Name)
	}
	return &CommandData{
		Name:        codegen.KebabCase(data.Name),
		VarName:     codegen.Goify(data.Name, false),
		Description: description,
		PkgName:     data.PkgName + "c",
	}
}

// BuildSubcommandData builds the data needed by CLI code generators to render
// the CLI parsing of the service sub-command.
func BuildSubcommandData(svcName string, m *service.MethodData, buildFunction *BuildFunctionData, flags []*FlagData) *SubcommandData {
	var (
		name        string
		fullName    string
		description string

		conversion string
	)
	{
		en := m.Name
		name = codegen.KebabCase(en)
		fullName = goifyTerms(svcName, en)
		description = m.Description
		if description == "" {
			description = fmt.Sprintf("Make request to the %q endpoint", m.Name)
		}

		if buildFunction == nil && len(flags) > 0 {
			// No build function, just convert the arg to the body type
			var convPre, convSuff string
			target := "data"
			if flagType(m.Payload) == "JSON" {
				target = "val"
				convPre = fmt.Sprintf("var val %s\n", m.Payload)
				convSuff = "\ndata = val"
			}
			conv, check := conversionCode(
				"*"+flags[0].FullName+"Flag",
				target,
				m.Payload,
				false,
			)
			conversion = convPre + conv + convSuff
			if check {
				conversion = "var err error\n" + conversion
				conversion += "\nif err != nil {\n"
				if flagType(m.Payload) == "JSON" {
					conversion += fmt.Sprintf(`return nil, nil, fmt.Errorf("invalid JSON for %s, example of valid JSON:\n%%s", %q)`,
						flags[0].FullName+"Flag", flags[0].Example)
				} else {
					conversion += fmt.Sprintf(`return nil, nil, fmt.Errorf("invalid value for %s, must be %s")`,
						flags[0].FullName+"Flag", flags[0].Type)
				}
				conversion += "\n}"
			}
		}
	}
	sub := &SubcommandData{
		Name:          name,
		FullName:      fullName,
		Description:   description,
		Flags:         flags,
		MethodVarName: m.VarName,
		BuildFunction: buildFunction,
		Conversion:    conversion,
	}
	generateExample(sub, svcName)

	return sub
}

// UsageCommands builds a section template that generates a help text showing
// the list of allowed commands and sub-commands.
func UsageCommands(data []*CommandData) *codegen.SectionTemplate {
	usages := make([]string, len(data))
	for i, cmd := range data {
		subs := make([]string, len(cmd.Subcommands))
		for i, s := range cmd.Subcommands {
			subs[i] = s.Name
		}
		var lp, rp string
		if len(subs) > 1 {
			lp = "("
			rp = ")"
		}
		usages[i] = fmt.Sprintf("%s %s%s%s", cmd.Name, lp, strings.Join(subs, "|"), rp)
	}

	return &codegen.SectionTemplate{Source: usageT, Data: usages}
}

// UsageExamples builds a section template that generates a help text showing
// a valid invocation of the CLI tool.
func UsageExamples(data []*CommandData) *codegen.SectionTemplate {
	var examples []string
	for i, cmd := range data {
		if i < 5 {
			examples = append(examples, cmd.Example)
		}
	}

	return &codegen.SectionTemplate{Source: exampleT, Data: examples}
}

// FlagsCode returns a string containing the code that parses the command-line
// flags to infer the command (service), sub-command (method), and the
// arguments (method payload) invoked by the tool. It panics if any error
// occurs during the generation of flag parsing code.
func FlagsCode(data []*CommandData) string {
	section := codegen.SectionTemplate{
		Name:    "parse-endpoint-flags",
		Source:  parseFlagsT,
		Data:    data,
		FuncMap: map[string]interface{}{"printDescription": printDescription},
	}
	var flagsCode bytes.Buffer
	err := section.Write(&flagsCode)
	if err != nil {
		panic(err)
	}

	return flagsCode.String()
}

// CommandUsage builds the section templates that can be used to generate the
// endpoint command usage code.
func CommandUsage(data *CommandData) *codegen.SectionTemplate {
	return &codegen.SectionTemplate{
		Name:    "cli-command-usage",
		Source:  commandUsageT,
		Data:    data,
		FuncMap: map[string]interface{}{"printDescription": printDescription},
	}
}

// PayloadBuilderSection builds the section template that can be used to
// generate the payload builder code.
func PayloadBuilderSection(buildFunction *BuildFunctionData) *codegen.SectionTemplate {
	return &codegen.SectionTemplate{
		Name:   "cli-build-payload",
		Source: buildPayloadT,
		Data:   buildFunction,
		FuncMap: map[string]interface{}{
			"fieldCode": fieldCode,
		},
	}
}

// NewFlagData creates a new FlagData from the given argument attributes.
//
// svcn is the service name
// en is the endpoint name
// name is the flag name
// typeName is the flag type
// description is the flag description
// required determines if the flag is required
// example is an example value for the flag
//
func NewFlagData(svcn, en, name, typeName, description string, required bool, example interface{}) *FlagData {
	ex := jsonExample(example)
	fn := goifyTerms(svcn, en, name)
	return &FlagData{
		Name:        codegen.KebabCase(name),
		VarName:     codegen.Goify(name, false),
		Type:        flagType(typeName),
		FullName:    fn,
		Description: description,
		Required:    required,
		Example:     ex,
	}
}

// FieldLoadCode returns the code used in the build payload function that
// initializes one of the payload object fields. It returns the initialization
// code and a boolean indicating whether the code requires an "err" variable.
func FieldLoadCode(f *FlagData, argName, argTypeName, validate string, defaultValue interface{}) (string, bool) {
	var (
		code    string
		check   bool
		startIf string
		endIf   string
	)
	{
		if !f.Required {
			startIf = fmt.Sprintf("if %s != \"\" {\n", f.FullName)
			endIf = "\n}"
		}
		if argTypeName == codegen.GoNativeTypeName(expr.String) {
			ref := "&"
			if f.Required || defaultValue != nil {
				ref = ""
			}
			code = argName + " = " + ref + f.FullName
		} else {
			ex := f.Example
			code, check = conversionCode(f.FullName, argName, argTypeName, !f.Required && defaultValue == nil)
			if check {
				code += "\nif err != nil {\n"
				if flagType(argTypeName) == "JSON" {
					code += fmt.Sprintf(`return nil, fmt.Errorf("invalid JSON for %s, example of valid JSON:\n%%s", %q)`,
						argName, ex)
				} else {
					code += fmt.Sprintf(`return nil, fmt.Errorf("invalid value for %s, must be %s")`,
						argName, f.Type)
				}
				code += "\n}"
			}
			if validate != "" {
				code += "\n" + validate + "\n" + "if err != nil {\n\treturn nil, err\n}"
			}
		}
	}
	return fmt.Sprintf("%s%s%s", startIf, code, endIf), check
}

// flagType calculates the type of a flag
func flagType(tname string) string {
	switch tname {
	case boolN, intN, int32N, int64N, uintN, uint32N, uint64N, float32N, float64N, stringN:
		return strings.ToUpper(tname)
	case bytesN:
		return "STRING"
	default: // Any, Array, Map, Object, User
		return "JSON"
	}
}

// jsonExample generates a json example
func jsonExample(v interface{}) string {
	// In JSON, keys must be a string. But goa allows map keys to be anything.
	r := reflect.ValueOf(v)
	if r.Kind() == reflect.Map {
		keys := r.MapKeys()
		if keys[0].Kind() != reflect.String {
			a := make(map[string]interface{}, len(keys))
			var kstr string
			for _, k := range keys {
				switch t := k.Interface().(type) {
				case bool:
					kstr = strconv.FormatBool(t)
				case int32:
					kstr = strconv.FormatInt(int64(t), 10)
				case int64:
					kstr = strconv.FormatInt(t, 10)
				case int:
					kstr = strconv.Itoa(t)
				case float32:
					kstr = strconv.FormatFloat(float64(t), 'f', -1, 32)
				case float64:
					kstr = strconv.FormatFloat(t, 'f', -1, 64)
				default:
					kstr = k.String()
				}
				a[kstr] = r.MapIndex(k).Interface()
			}
			v = a
		}
	}
	b, err := json.MarshalIndent(v, "   ", "   ")
	ex := "?"
	if err == nil {
		ex = string(b)
	}
	if strings.Contains(ex, "\n") {
		ex = "'" + strings.Replace(ex, "'", "\\'", -1) + "'"
	}
	return ex
}

var (
	boolN    = codegen.GoNativeTypeName(expr.Boolean)
	intN     = codegen.GoNativeTypeName(expr.Int)
	int32N   = codegen.GoNativeTypeName(expr.Int32)
	int64N   = codegen.GoNativeTypeName(expr.Int64)
	uintN    = codegen.GoNativeTypeName(expr.UInt)
	uint32N  = codegen.GoNativeTypeName(expr.UInt32)
	uint64N  = codegen.GoNativeTypeName(expr.UInt64)
	float32N = codegen.GoNativeTypeName(expr.Float32)
	float64N = codegen.GoNativeTypeName(expr.Float64)
	stringN  = codegen.GoNativeTypeName(expr.String)
	bytesN   = codegen.GoNativeTypeName(expr.Bytes)
)

// conversionCode produces the code that converts the string stored in the
// variable "from" to the value stored in the variable "to" of type typeName.
func conversionCode(from, to, typeName string, pointer bool) (string, bool) {
	var (
		parse    string
		cast     string
		checkErr bool
	)
	target := to
	needCast := typeName != stringN && typeName != bytesN && flagType(typeName) != "JSON"
	decl := ""
	if needCast && pointer {
		target = "val"
		decl = ":"
	}
	switch typeName {
	case boolN:
		if pointer {
			parse = fmt.Sprintf("var %s bool\n", target)
		}
		parse += fmt.Sprintf("%s, err = strconv.ParseBool(%s)", target, from)
		checkErr = true
	case intN:
		parse = fmt.Sprintf("var v int64\nv, err = strconv.ParseInt(%s, 10, 64)", from)
		cast = fmt.Sprintf("%s %s= int(v)", target, decl)
		checkErr = true
	case int32N:
		parse = fmt.Sprintf("var v int64\nv, err = strconv.ParseInt(%s, 10, 32)", from)
		cast = fmt.Sprintf("%s %s= int32(v)", target, decl)
		checkErr = true
	case int64N:
		parse = fmt.Sprintf("%s, err %s= strconv.ParseInt(%s, 10, 64)", target, decl, from)
		checkErr = true
	case uintN:
		parse = fmt.Sprintf("var v uint64\nv, err = strconv.ParseUint(%s, 10, 64)", from)
		cast = fmt.Sprintf("%s %s= uint(v)", target, decl)
		checkErr = true
	case uint32N:
		parse = fmt.Sprintf("var v uint64\nv, err = strconv.ParseUint(%s, 10, 32)", from)
		cast = fmt.Sprintf("%s %s= uint32(v)", target, decl)
		checkErr = true
	case uint64N:
		parse = fmt.Sprintf("%s, err %s= strconv.ParseUint(%s, 10, 64)", target, decl, from)
		checkErr = true
	case float32N:
		parse = fmt.Sprintf("var v float64\nv, err = strconv.ParseFloat(%s, 32)", from)
		cast = fmt.Sprintf("%s %s= float32(v)", target, decl)
		checkErr = true
	case float64N:
		parse = fmt.Sprintf("%s, err %s= strconv.ParseFloat(%s, 64)", target, decl, from)
		checkErr = true
	case stringN:
		parse = fmt.Sprintf("%s %s= %s", target, decl, from)
	case bytesN:
		parse = fmt.Sprintf("%s %s= []byte(%s)", target, decl, from)
	default:
		parse = fmt.Sprintf("err = json.Unmarshal([]byte(%s), &%s)", from, target)
		checkErr = true
	}
	if !needCast {
		return parse, checkErr
	}
	if cast != "" {
		parse = parse + "\n" + cast
	}
	if to != target {
		ref := ""
		if pointer {
			ref = "&"
		}
		parse = parse + fmt.Sprintf("\n%s = %s%s", to, ref, target)
	}
	return parse, checkErr
}

// goifyTerms makes valid go identifiers out of the supplied terms
func goifyTerms(terms ...string) string {
	res := codegen.Goify(terms[0], false)
	if len(terms) == 1 {
		return res
	}
	for _, t := range terms[1:] {
		res += codegen.Goify(t, true)
	}
	return res
}

func printDescription(desc string) string {
	res := strings.Replace(desc, "`", "`+\"`\"+`", -1)
	res = strings.Replace(res, "\n", "\n\t", -1)
	return res
}

func generateExample(sub *SubcommandData, svc string) {
	ex := codegen.KebabCase(svc) + " " + codegen.KebabCase(sub.Name)
	for _, f := range sub.Flags {
		ex += " --" + f.Name + " " + f.Example
	}
	sub.Example = ex
}

// fieldCode generates code to initialize the data structures fields
// from the given args. It is used only in templates.
func fieldCode(init *PayloadInitData) string {
	varn := "res"
	if init.ReturnTypeAttribute == "" {
		varn = "v"
	}
	// We can ignore the transform helpers as there won't be any generated
	// because the args cannot be user types.
	c, _, err := codegen.InitStructFields(init.Args, init.ReturnTypeName, varn, "", init.ReturnTypePkg, init.Code == "")
	if err != nil {
		panic(err) //bug
	}
	return c
}

// input: []string
const usageT = `// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return ` + "`" + `{{ range . }}{{ . }}
{{ end }}` + "`" + `
}
`

// input: []string
const exampleT = `// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return {{ range . }}os.Args[0] + ` + "`" + ` {{ . }}` + "`" + ` + "\n" +
	{{ end }}""
}
`

// input: []commandData
const parseFlagsT = `var (
		{{- range . }}
		{{ .VarName }}Flags = flag.NewFlagSet("{{ .Name }}", flag.ContinueOnError)
		{{ range .Subcommands }}
		{{ .FullName }}Flags = flag.NewFlagSet("{{ .Name }}", flag.ExitOnError)
		{{- $sub := . }}
		{{- range .Flags }}
		{{ .FullName }}Flag = {{ $sub.FullName }}Flags.String("{{ .Name }}", "{{ if .Required }}REQUIRED{{ end }}", {{ printf "%q" .Description }})
		{{- end }}
		{{ end }}
		{{- end }}
	)
	{{ range . -}}
	{{ $cmd := . -}}
	{{ .VarName }}Flags.Usage = {{ .VarName }}Usage
	{{ range .Subcommands -}}
	{{ .FullName }}Flags.Usage = {{ .FullName }}Usage
	{{ end }}
	{{ end }}
	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
	{{- range . }}
		case "{{ .Name }}":
			svcf = {{ .VarName }}Flags
	{{- end }}
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
	{{- range . }}
		case "{{ .Name }}":
			switch epn {
		{{- range .Subcommands }}
			case "{{ .Name }}":
				epf = {{ .FullName }}Flags
		{{ end }}
			}
	{{ end }}
		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}
`

// input: commandData
const commandUsageT = `{{ printf "%sUsage displays the usage of the %s command and its subcommands." .Name .Name | comment }}
func {{ .VarName }}Usage() {
	fmt.Fprintf(os.Stderr, ` + "`" + `{{ printDescription .Description }}
Usage:
    %s [globalflags] {{ .Name }} COMMAND [flags]

COMMAND:
    {{- range .Subcommands }}
    {{ .Name }}: {{ printDescription .Description }}
    {{- end }}

Additional help:
    %s {{ .Name }} COMMAND --help
` + "`" + `, os.Args[0], os.Args[0])
}

{{- range .Subcommands }}
func {{ .FullName }}Usage() {
	fmt.Fprintf(os.Stderr, ` + "`" + `%s [flags] {{ $.Name }} {{ .Name }}{{range .Flags }} -{{ .Name }} {{ .Type }}{{ end }}

{{ printDescription .Description}}
	{{- range .Flags }}
    -{{ .Name }} {{ .Type }}: {{ .Description }}
	{{- end }}

Example:
    ` + "`+os.Args[0]+" + "`" + ` {{ .Example }}
` + "`" + `, os.Args[0])
}
{{ end }}
`

// input: buildFunctionData
const buildPayloadT = `{{ printf "%s builds the payload for the %s %s endpoint from CLI flags." .Name .ServiceName .MethodName | comment }}
func {{ .Name }}({{ range .FormalParams }}{{ . }} string, {{ end }}) ({{ .ResultType }}, error) {
{{- if .CheckErr }}
	var err error
{{- end }}
{{- range .Fields }}
	{{- if .VarName }}
		var {{ .VarName }} {{ .TypeRef }}
		{
			{{ .Init }}
		}
	{{- end }}
{{- end }}
{{- with .PayloadInit }}
	{{- if .Code }}
		{{ .Code }}
		{{- if .ReturnTypeAttribute }}
			res := &{{ .ReturnTypeName }}{
				{{ .ReturnTypeAttribute }}: v,
			}
		{{- end }}
	{{- end }}
	{{- if .ReturnIsStruct }}
		{{- if not .Code }}
		{{ if .ReturnTypeAttribute }}res{{ else }}v{{ end }} := &{{ .ReturnTypeName }}{}
		{{- end }}
		{{ fieldCode . }}
	{{- end }}
	return {{ if .ReturnTypeAttribute }}res{{ else }}v{{ end }}, nil
{{- end }}
}
`
