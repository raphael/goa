package codegen

import (
	"fmt"
	"path/filepath"

	"goa.design/goa/v3/codegen"
	"goa.design/goa/v3/codegen/service"
	"goa.design/goa/v3/expr"
)

// ClientFiles returns the client HTTP transport files.
func ClientFiles(genpkg string, root *expr.RootExpr) []*codegen.File {
	fw := make([]*codegen.File, 2*len(root.API.HTTP.Services))
	for i, r := range root.API.HTTP.Services {
		fw[i] = client(genpkg, r)
	}
	for i, r := range root.API.HTTP.Services {
		fw[i+len(root.API.HTTP.Services)] = clientEncodeDecode(genpkg, r)
	}
	return fw
}

// client returns the client HTTP transport file
func client(genpkg string, svc *expr.HTTPServiceExpr) *codegen.File {
	data := HTTPServices.Get(svc.Name())
	svcName := codegen.SnakeCase(data.Service.VarName)
	path := filepath.Join(codegen.Gendir, "http", svcName, "client", "client.go")
	title := fmt.Sprintf("%s client HTTP transport", svc.Name())
	sections := []*codegen.SectionTemplate{
		codegen.Header(title, "client", []*codegen.ImportSpec{
			{Path: "context"},
			{Path: "fmt"},
			{Path: "io"},
			{Path: "mime/multipart"},
			{Path: "net/http"},
			{Path: "strconv"},
			{Path: "strings"},
			{Path: "sync"},
			{Path: "time"},
			{Path: "github.com/gorilla/websocket"},
			codegen.GoaImport(""),
			codegen.GoaNamedImport("http", "goahttp"),
			{Path: genpkg + "/" + svcName, Name: data.Service.PkgName},
			{Path: genpkg + "/" + svcName + "/" + "views", Name: data.Service.ViewsPkg},
		}),
	}
	sections = append(sections, &codegen.SectionTemplate{
		Name:   "client-struct",
		Source: clientStructT,
		Data:   data,
		FuncMap: map[string]interface{}{
			"streamingEndpointExists": streamingEndpointExists,
		},
	})
	if streamingEndpointExists(data) {
		sections = append(sections, &codegen.SectionTemplate{
			Name:   "client-stream-conn-configurer-struct",
			Source: streamConnConfigurerStructT,
			Data:   data,
			FuncMap: map[string]interface{}{
				"isStreamingEndpoint": isStreamingEndpoint,
			},
		})
	}
	for _, e := range data.Endpoints {
		if e.ClientStream != nil {
			sections = append(sections, &codegen.SectionTemplate{
				Name:   "client-stream-struct-type",
				Source: streamStructTypeT,
				Data:   e.ClientStream,
			})
		}
	}

	for _, e := range data.Endpoints {
		if e.MultipartRequestEncoder != nil {
			sections = append(sections, &codegen.SectionTemplate{
				Name:   "multipart-request-encoder-type",
				Source: multipartRequestEncoderTypeT,
				Data:   e.MultipartRequestEncoder,
			})
		}
	}

	sections = append(sections, &codegen.SectionTemplate{
		Name:   "client-init",
		Source: clientInitT,
		Data:   data,
		FuncMap: map[string]interface{}{
			"streamingEndpointExists": streamingEndpointExists,
		},
	})

	if streamingEndpointExists(data) {
		sections = append(sections, &codegen.SectionTemplate{
			Name:   "client-stream-conn-configurer-struct-init",
			Source: streamConnConfigurerStructInitT,
			Data:   data,
			FuncMap: map[string]interface{}{
				"isStreamingEndpoint": isStreamingEndpoint,
			},
		})
	}

	for _, e := range data.Endpoints {
		sections = append(sections, &codegen.SectionTemplate{
			Name:   "client-endpoint-init",
			Source: endpointInitT,
			Data:   e,
		})
		if e.ClientStream != nil {
			if e.ClientStream.RecvTypeRef != "" {
				sections = append(sections, &codegen.SectionTemplate{
					Name:   "client-stream-recv",
					Source: streamRecvT,
					Data:   e.ClientStream,
					FuncMap: map[string]interface{}{
						"upgradeParams": upgradeParams,
					},
				})
			}
			switch e.ClientStream.Kind {
			case expr.ClientStreamKind, expr.BidirectionalStreamKind:
				sections = append(sections, &codegen.SectionTemplate{
					Name:   "client-stream-send",
					Source: streamSendT,
					Data:   e.ClientStream,
					FuncMap: map[string]interface{}{
						"upgradeParams":    upgradeParams,
						"viewedServerBody": viewedServerBody,
					},
				})
			}
			if e.ClientStream.MustClose {
				sections = append(sections, &codegen.SectionTemplate{
					Name:   "client-stream-close",
					Source: streamCloseT,
					Data:   e.ClientStream,
					FuncMap: map[string]interface{}{
						"upgradeParams": upgradeParams,
					},
				})
			}
			if e.Method.ViewedResult != nil && e.Method.ViewedResult.ViewName == "" {
				sections = append(sections, &codegen.SectionTemplate{
					Name:   "client-stream-set-view",
					Source: streamSetViewT,
					Data:   e.ClientStream,
				})
			}
		}
	}

	return &codegen.File{Path: path, SectionTemplates: sections}
}

// clientEncodeDecode returns the file containing the HTTP client encoding and
// decoding logic.
func clientEncodeDecode(genpkg string, svc *expr.HTTPServiceExpr) *codegen.File {
	data := HTTPServices.Get(svc.Name())
	svcName := codegen.SnakeCase(data.Service.VarName)
	path := filepath.Join(codegen.Gendir, "http", svcName, "client", "encode_decode.go")
	title := fmt.Sprintf("%s HTTP client encoders and decoders", svc.Name())
	sections := []*codegen.SectionTemplate{
		codegen.Header(title, "client", []*codegen.ImportSpec{
			{Path: "bytes"},
			{Path: "context"},
			{Path: "fmt"},
			{Path: "io"},
			{Path: "io/ioutil"},
			{Path: "mime/multipart"},
			{Path: "net/http"},
			{Path: "net/url"},
			{Path: "strconv"},
			{Path: "strings"},
			{Path: "unicode/utf8"},
			codegen.GoaImport(""),
			codegen.GoaNamedImport("http", "goahttp"),
			{Path: genpkg + "/" + svcName, Name: data.Service.PkgName},
			{Path: genpkg + "/" + svcName + "/" + "views", Name: data.Service.ViewsPkg},
		}),
	}

	for _, e := range data.Endpoints {
		sections = append(sections, &codegen.SectionTemplate{
			Name:   "request-builder",
			Source: requestBuilderT,
			Data:   e,
		})
		if e.RequestEncoder != "" {
			sections = append(sections, &codegen.SectionTemplate{
				Name:   "request-encoder",
				Source: requestEncoderT,
				FuncMap: map[string]interface{}{
					"typeConversionData": typeConversionData,
					"mapConversionData":  mapConversionData,
					"goTypeRef": func(dt expr.DataType) string {
						return service.Services.Get(svc.Name()).Scope.GoTypeRef(&expr.AttributeExpr{Type: dt})
					},
					"isBearer":    isBearer,
					"aliasedType": fieldType,
					"isAliased": func(dt expr.DataType) bool {
						_, ok := dt.(expr.UserType)
						return ok
					},
				},
				Data: e,
			})
		}
		if e.MultipartRequestEncoder != nil {
			sections = append(sections, &codegen.SectionTemplate{
				Name:   "multipart-request-encoder",
				Source: multipartRequestEncoderT,
				Data:   e.MultipartRequestEncoder,
			})
		}
		if e.Result != nil || len(e.Errors) > 0 {
			sections = append(sections, &codegen.SectionTemplate{
				Name:   "response-decoder",
				Source: responseDecoderT,
				Data:   e,
				FuncMap: map[string]interface{}{
					"goTypeRef": func(dt expr.DataType) string {
						return service.Services.Get(svc.Name()).Scope.GoTypeRef(&expr.AttributeExpr{Type: dt})
					},
				},
			})
		}
	}
	for _, h := range data.ClientTransformHelpers {
		sections = append(sections, &codegen.SectionTemplate{
			Name:   "client-transform-helper",
			Source: transformHelperT,
			Data:   h,
		})
	}

	return &codegen.File{Path: path, SectionTemplates: sections}
}

// typeConversionData produces the template data suitable for executing the
// "header_conversion" template.
func typeConversionData(dt, ft expr.DataType, varName string, target string) map[string]interface{} {
	ut, isut := ft.(expr.UserType)
	if isut {
		ft = ut.Attribute().Type
	}
	return map[string]interface{}{
		"Type":      dt,
		"FieldType": ft,
		"VarName":   varName,
		"Target":    target,
		"IsAliased": isut,
	}
}

func mapConversionData(dt, ft expr.DataType, varName, sourceVar, sourceField string, newVar bool) map[string]interface{} {
	ut, isut := ft.(expr.UserType)
	if isut {
		ft = ut.Attribute().Type
	}
	return map[string]interface{}{
		"Type":        dt,
		"FieldType":   ft,
		"VarName":     varName,
		"SourceVar":   sourceVar,
		"SourceField": sourceField,
		"NewVar":      newVar,
		"IsAliased":   isut,
	}
}

func fieldType(ft expr.DataType) expr.DataType {
	ut, isut := ft.(expr.UserType)
	if isut {
		return ut.Attribute().Type
	}
	return ft
}

// isBearer returns true if the security scheme uses a Bearer scheme.
func isBearer(schemes []*service.SchemeData) bool {
	for _, s := range schemes {
		if s.Name != "Authorization" {
			continue
		}
		if s.Type == "JWT" || s.Type == "OAuth2" {
			return true
		}
	}
	return false
}

// input: ServiceData
const clientStructT = `{{ printf "%s lists the %s service endpoint HTTP clients." .ClientStruct .Service.Name | comment }}
type {{ .ClientStruct }} struct {
	{{- range .Endpoints }}
	{{ printf "%s Doer is the HTTP client used to make requests to the %s endpoint." .Method.VarName .Method.Name | comment }}
	{{ .Method.VarName }}Doer goahttp.Doer
	{{ end }}
	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme     string
	host       string
	encoder    func(*http.Request) goahttp.Encoder
	decoder    func(*http.Response) goahttp.Decoder
	{{- if streamingEndpointExists . }}
	dialer goahttp.Dialer
	configurer *ConnConfigurer
	{{- end }}
}
`

// input: ServiceData
const clientInitT = `{{ printf "New%s instantiates HTTP clients for all the %s service servers." .ClientStruct .Service.Name | comment }}
func New{{ .ClientStruct }}(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
	{{- if streamingEndpointExists . }}
	dialer goahttp.Dialer,
	cfn *ConnConfigurer,
	{{- end }}
) *{{ .ClientStruct }} {
{{- if streamingEndpointExists . }}
	if cfn == nil {
		cfn = &ConnConfigurer{}
	}
{{- end }}
	return &{{ .ClientStruct }}{
		{{- range .Endpoints }}
		{{ .Method.VarName }}Doer: doer,
		{{- end }}
		RestoreResponseBody: restoreBody,
		scheme:            scheme,
		host:              host,
		decoder:           dec,
		encoder:           enc,
		{{- if streamingEndpointExists . }}
		dialer: dialer,
		configurer: cfn,
		{{- end }}
	}
}
`

// input: EndpointData
const endpointInitT = `{{ printf "%s returns an endpoint that makes HTTP requests to the %s service %s server." .EndpointInit .ServiceName .Method.Name | comment }}
func (c *{{ .ClientStruct }}) {{ .EndpointInit }}({{ if .MultipartRequestEncoder }}{{ .MultipartRequestEncoder.VarName }} {{ .MultipartRequestEncoder.FuncName }}{{ end }}) goa.Endpoint {
	var (
		{{- if and .ClientStream .RequestEncoder }}
		encodeRequest  = {{ .RequestEncoder }}({{ if .MultipartRequestEncoder }}{{ .MultipartRequestEncoder.InitName }}({{ .MultipartRequestEncoder.VarName }}){{ else }}c.encoder{{ end }})
		{{- else }}
			{{- if .RequestEncoder }}
		encodeRequest  = {{ .RequestEncoder }}({{ if .MultipartRequestEncoder }}{{ .MultipartRequestEncoder.InitName }}({{ .MultipartRequestEncoder.VarName }}){{ else }}c.encoder{{ end }})
			{{- end }}
		{{- end }}
		decodeResponse = {{ .ResponseDecoder }}(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.{{ .RequestInit.Name }}(ctx, {{ range .RequestInit.ClientArgs }}{{ .Ref }}{{ end }})
		if err != nil {
			return nil, err
		}
	{{- if .RequestEncoder }}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
	{{- end }}

	{{- if .ClientStream }}
		var cancel context.CancelFunc
		{
			ctx, cancel = context.WithCancel(ctx)
		}
		conn, resp, err := c.dialer.DialContext(ctx, req.URL.String(), req.Header)
		if err != nil {
			if resp != nil {
				return decodeResponse(resp)
			}
			return nil, goahttp.ErrRequestError("{{ .ServiceName }}", "{{ .Method.Name }}", err)
		}
		if c.configurer.{{ .Method.VarName }}Fn != nil {
			conn = c.configurer.{{ .Method.VarName }}Fn(conn, cancel)
		}
	{{- if eq .ClientStream.SendName "" }}
		go func() {
			<-ctx.Done()
			conn.WriteControl(
				websocket.CloseMessage,
				websocket.FormatCloseMessage(websocket.CloseNormalClosure, "client closing connection"),
				time.Now().Add(time.Second),
			)
			conn.Close()
		}()
	{{- end }}
		stream := &{{ .ClientStream.VarName }}{conn: conn}
		{{- if .Method.ViewedResult }}
			{{- if not .Method.ViewedResult.ViewName }}
		view := resp.Header.Get("goa-view")
		stream.SetView(view)
			{{- end }}
		{{- end }}
		return stream, nil
	{{- else }}
		resp, err := c.{{ .Method.VarName }}Doer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("{{ .ServiceName }}", "{{ .Method.Name }}", err)
		}
		return decodeResponse(resp)
	{{- end }}
	}
}
`

// input: EndpointData
const requestBuilderT = `{{ comment .RequestInit.Description }}
func (c *{{ .ClientStruct }}) {{ .RequestInit.Name }}(ctx context.Context, {{ range .RequestInit.ClientArgs }}{{ .Name }} {{ .TypeRef }}{{ end }}) (*http.Request, error) {
	{{- .RequestInit.ClientCode }}
}
`

// input: EndpointData
const requestEncoderT = `{{ printf "%s returns an encoder for requests sent to the %s %s server." .RequestEncoder .ServiceName .Method.Name | comment }}
func {{ .RequestEncoder }}(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.({{ .Payload.Ref }})
		if !ok {
			return goahttp.ErrInvalidType("{{ .ServiceName }}", "{{ .Method.Name }}", "{{ .Payload.Ref }}", v)
		}
	{{- range .Payload.Request.Headers }}
		{{- if .FieldName }}
			{{- if .FieldPointer }}
		if p.{{ .FieldName }} != nil {
			{{- end }}
			{{- if (and (eq .Name "Authorization") (isBearer $.HeaderSchemes)) }}
		if !strings.Contains({{ if .FieldPointer }}*{{ end }}p.{{ .FieldName }}, " ") {
			req.Header.Set({{ printf "%q" .Name }}, "Bearer "+{{ if .FieldPointer }}*{{ end }}p.{{ .FieldName }})
		} else {
			{{- end }}
			req.Header.Set({{ printf "%q" .Name }}, {{ if .FieldPointer }}*{{ end }}p.{{ .FieldName }})
			{{- if (and (eq .Name "Authorization") (isBearer $.HeaderSchemes)) }}
		}
			{{- end }}
			{{- if .FieldPointer }}
		}
			{{- end }}
		{{- end }}
	{{- end }}
	{{- if or .Payload.Request.QueryParams }}
		values := req.URL.Query()
	{{- end }}
	{{- range .Payload.Request.QueryParams }}
		{{- if .MapQueryParams }}
		for key, value := range p{{ if .FieldName }}.{{ .FieldName }}{{ end }} {
			{{ template "type_conversion" (typeConversionData .Type.KeyType.Type (aliasedType .FieldType).KeyType.Type "keyStr" "key") }}
			{{- if eq .Type.ElemType.Type.Name "array" }}
			for _, val := range value {
				{{ template "type_conversion" (typeConversionData .Type.ElemType.Type.ElemType.Type (aliasedType (aliasedType .FieldType).ElemType.Type).ElemType.Type "valStr" "val") }}
				values.Add(keyStr, valStr)
			}
			{{- else }}
			{{ template "type_conversion" (typeConversionData .Type.ElemType.Type (aliasedType .FieldType).ElemType.Type "valueStr" "value") }}
			values.Add(keyStr, valueStr)
			{{- end }}
    }
		{{- else if .StringSlice }}
			for _, value := range p{{ if .FieldName }}.{{ .FieldName }}{{ end }} {
				values.Add("{{ .Name }}", value)
			}
		{{- else if .Slice }}
			for _, value := range p{{ if .FieldName }}.{{ .FieldName }}{{ end }} {
				{{ template "type_conversion" (typeConversionData .Type.ElemType.Type (aliasedType .FieldType).ElemType.Type "valueStr" "value") }}
				values.Add("{{ .Name }}", valueStr)
			}
		{{- else if .Map }}
			{{- template "map_conversion" (mapConversionData .Type .FieldType .Name "p" .FieldName true) }}
		{{- else if .FieldName }}
			{{- if .FieldPointer }}
		if p.{{ .FieldName }} != nil {
			{{- end }}
		values.Add("{{ .Name }}",
			{{- if eq .Type.Name "bytes" }} string(
			{{- else if not (eq .Type.Name "string") }} fmt.Sprintf("%v",
			{{- end }}
			{{- if .FieldPointer }}*{{ end }}p.{{ .FieldName }}
			{{- if or (eq .Type.Name "bytes") (not (eq .Type.Name "string")) }})
			{{- end }})
			{{- if .FieldPointer }}
		}
			{{- end }}
		{{- else }}
			{{- if eq .Type.Name "string" }}
				values.Add("{{ .Name }}", p)
			{{- else }}
				{{ template "type_conversion" (typeConversionData .Type .FieldType "pStr" "p") }}
				values.Add("{{ .Name }}", pStr)
			{{- end }}
		{{- end }}
	{{- end }}
	{{- if .Payload.Request.QueryParams }}
		req.URL.RawQuery = values.Encode()
	{{- end }}
	{{- if .MultipartRequestEncoder }}
		if err := encoder(req).Encode(p); err != nil {
			return goahttp.ErrEncodingError("{{ .ServiceName }}", "{{ .Method.Name }}", err)
		}
	{{- else if .Payload.Request.ClientBody }}
		{{- if .Payload.Request.ClientBody.Init }}
		body := {{ .Payload.Request.ClientBody.Init.Name }}({{ range .Payload.Request.ClientBody.Init.ClientArgs }}{{ if .FieldPointer }}&{{ end }}{{ .Name }}, {{ end }})
		{{- else }}
		body := p
		{{- end }}
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("{{ .ServiceName }}", "{{ .Method.Name }}", err)
		}
	{{- end }}
	{{- if .BasicScheme }}{{ with .BasicScheme }}
		{{- if not .UsernameRequired }}
		if p.{{ .UsernameField }} != nil {
		{{- end }}
		{{- if not .PasswordRequired }}
		if p.{{ .PasswordField }} != nil {
		{{- end }}
		req.SetBasicAuth({{ if .UsernamePointer }}*{{ end }}p.{{ .UsernameField }}, {{ if .PasswordPointer }}*{{ end }}p.{{ .PasswordField }})
		{{- if not .UsernameRequired }}
		}
		{{- end }}
		{{- if not .PasswordRequired }}
		}
		{{- end }}
	{{- end }}{{ end }}
		return nil
	}
}

{{- define "map_conversion" }}
  for k{{ if not (eq .Type.KeyType.Type.Name "string") }}Raw{{ end }}, value := range {{ .SourceVar }}{{ if .SourceField }}.{{ .SourceField }}{{ end }} {
		{{- if not (eq .Type.KeyType.Type.Name "string") }}
			{{ template "type_conversion" (typeConversionData .Type.KeyType.Type .FieldType.KeyType.Type "k" "kRaw") }}
		{{- end }}
		key {{ if .NewVar }}:={{ else }}={{ end }} fmt.Sprintf("{{ .VarName }}[%s]", {{ if not .NewVar }}key, {{ end }}k)
		{{- if eq .Type.ElemType.Type.Name "string" }}
			values.Add(key, {{ if (isAliased .FieldType.ElemType.Type) }}string({{ end }}value{{ if (isAliased .FieldType.ElemType.Type) }}){{ end }})
		{{- else if eq .Type.ElemType.Type.Name "map" }}
			{{- template "map_conversion" (mapConversionData .Type.ElemType.Type .FieldType.ElemType.Type "%s" "value" "" false) }}
		{{- else if eq .Type.ElemType.Type.Name "array" }}
			{{- if and (eq .Type.ElemType.Type.ElemType.Type.Name "string") (not (isAliased .FieldType.ElemType.Type.ElemType.Type)) }}
				values[key] = value
			{{- else }}
				for _, val := range value {
					{{ template "type_conversion" (typeConversionData .Type.ElemType.Type.ElemType.Type (aliasedType .FieldType.ElemType.Type).ElemType.Type "valStr" "val") }}
					values.Add(key, valStr)
				}
			{{- end }}
		{{- else }}
			{{ template "type_conversion" (typeConversionData .Type.ElemType.Type .FieldType.ElemType.Type "valueStr" "value") }}
			values.Add(key, valueStr)
		{{- end }}
	}
{{- end }}

{{- define "type_conversion" }}
  {{- if eq .Type.Name "boolean" -}}
    {{ .VarName }} := strconv.FormatBool({{ if .IsAliased }}bool({{ end }}{{ .Target }}{{ if .IsAliased }}){{ end }})
  {{- else if eq .Type.Name "int" -}}
    {{ .VarName }} := strconv.Itoa({{ if .IsAliased }}int({{ end }}{{ .Target }}{{ if .IsAliased }}){{ end }})
  {{- else if eq .Type.Name "int32" -}}
    {{ .VarName }} := strconv.FormatInt(int64({{ .Target }}), 10)
  {{- else if eq .Type.Name "int64" -}}
    {{ .VarName }} := strconv.FormatInt({{ if .IsAliased }}int64({{ end }}{{ .Target }}{{ if .IsAliased }}){{ end }}, 10)
  {{- else if eq .Type.Name "uint" -}}
    {{ .VarName }} := strconv.FormatUint(uint64({{ .Target }}), 10)
  {{- else if eq .Type.Name "uint32" -}}
    {{ .VarName }} := strconv.FormatUint(uint64({{ .Target }}), 10)
  {{- else if eq .Type.Name "uint64" -}}
    {{ .VarName }} := strconv.FormatUint({{ if .IsAliased }}uint64({{ end }}{{ .Target }}{{ if .IsAliased }}){{ end }}, 10)
  {{- else if eq .Type.Name "float32" -}}
    {{ .VarName }} := strconv.FormatFloat(float64({{ .Target }}), 'f', -1, 32)
  {{- else if eq .Type.Name "float64" -}}
    {{ .VarName }} := strconv.FormatFloat({{ if .IsAliased }}float64({{ end }}{{ .Target }}{{ if .IsAliased }}){{ end }}, 'f', -1, 64)
	{{- else if eq .Type.Name "string" -}}
    {{ .VarName }} := {{ if .IsAliased }}string({{ end }}{{ .Target }}{{ if .IsAliased }}){{ end }}
  {{- else if eq .Type.Name "bytes" -}}
    {{ .VarName }} := string({{ .Target }})
  {{- else if eq .Type.Name "any" -}}
    {{ .VarName }} := fmt.Sprintf("%v", {{ .Target }})
  {{- else }}
    // unsupported type {{ .Type.Name }} for field {{ .FieldName }}
  {{- end }}
{{- end }}
`

// input: EndpointData
const responseDecoderT = `{{ printf "%s returns a decoder for responses returned by the %s %s endpoint. restoreBody controls whether the response body should be restored after having been read." .ResponseDecoder .ServiceName .Method.Name | comment }}
{{- if .Errors }}
{{ printf "%s may return the following errors:" .ResponseDecoder | comment }}
	{{- range $gerr := .Errors }}
	{{- range $errors := .Errors }}
//	- {{ printf "%q" .Name }} (type {{ .Ref }}): {{ .Response.StatusCode }}{{ if .Response.Description }}, {{ .Response.Description }}{{ end }}
	{{- end }}
	{{- end }}
//	- error: internal error
{{- end }}
func {{ .ResponseDecoder }}(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
	{{- range .Result.Responses }}
		case {{ .StatusCode }}:
` + singleResponseT + `
		{{- if .ResultInit }}
			{{- if .ViewedResult }}
			p := {{ .ResultInit.Name }}({{ range .ResultInit.ClientArgs }}{{ .Ref }},{{ end }})
			{{- if .TagName }}
				tmp := {{ printf "%q" .TagValue }}
				p.{{ .TagName }} = &tmp
			{{- end }}
				{{- if $.Method.ViewedResult.ViewName }}
			view := {{ printf "%q" $.Method.ViewedResult.ViewName }}
				{{- else }}
			view := resp.Header.Get("goa-view")
				{{- end }}
			vres := {{ if not $.Method.ViewedResult.IsCollection }}&{{ end }}{{ $.Method.ViewedResult.ViewsPkg}}.{{ $.Method.ViewedResult.VarName }}{p, view}
			{{- if .ClientBody }}
				if err = {{ $.Method.ViewedResult.ViewsPkg}}.Validate{{ $.Method.Result }}(vres); err != nil {
					return nil, goahttp.ErrValidationError("{{ $.ServiceName }}", "{{ $.Method.Name }}", err)
				}
			{{- end }}
			res := {{ $.ServicePkgName }}.{{ $.Method.ViewedResult.ResultInit.Name }}(vres)
			{{- else }}
			res := {{ .ResultInit.Name }}({{ range .ResultInit.ClientArgs }}{{ .Ref }},{{ end }})
			{{- end }}
			{{- if and .TagName (not .ViewedResult) }}
				{{- if .TagPointer }}
					tmp := {{ printf "%q" .TagValue }}
					res.{{ .TagName }} = &tmp
				{{- else }}
					res.{{ .TagName }} = {{ printf "%q" .TagValue }}
				{{- end }}
			{{- end }}
			return res, nil
		{{- else if .ClientBody }}
			return body, nil
		{{- else }}
			return nil, nil
		{{- end }}
	{{- end }}
	{{- range .Errors }}
		case {{ .StatusCode }}:
		{{- if gt (len .Errors) 1 }}
		en := resp.Header.Get("goa-error")
		switch en {
			{{- range .Errors }}
		case {{ printf "%q" .Name }}:
				{{- with .Response }}
` + singleResponseT + `
					{{- if .ResultInit }}
			return nil, {{ .ResultInit.Name }}({{ range .ResultInit.ClientArgs }}{{ .Ref }},{{ end }})
					{{- else if .ClientBody }}
			return nil, body
					{{- else }}
			return nil, nil
					{{- end }}
				{{- end }}
			{{- end }}
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse({{ printf "%q" $.ServiceName }}, {{ printf "%q" $.Method.Name }}, resp.StatusCode, string(body))
		}
		{{- else }}
			{{- with (index .Errors 0).Response }}
` + singleResponseT + `
				{{- if .ResultInit }}
			return nil, {{ .ResultInit.Name }}({{ range .ResultInit.ClientArgs }}{{ .Ref }},{{ end }})
				{{- else if .ClientBody }}
			return nil, body
				{{- else }}
			return nil, nil
				{{- end }}
			{{- end }}
		{{- end }}
	{{- end }}
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse({{ printf "%q" .ServiceName }}, {{ printf "%q" .Method.Name }}, resp.StatusCode, string(body))
		}
	}
}
` + typeConversionT

// input: ResponseData
const singleResponseT = ` {{- if .ClientBody }}
			var (
				body {{ .ClientBody.VarName }}
				err error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("{{ $.ServiceName }}", "{{ $.Method.Name }}", err)
			}
		{{- if .ClientBody.ValidateRef }}
			{{ .ClientBody.ValidateRef }}
			if err != nil {
				return nil, goahttp.ErrValidationError("{{ $.ServiceName }}", "{{ $.Method.Name }}", err)
			}
		{{- end }}
	{{- end }}

	{{- if .Headers }}
			var (
		{{- range .Headers }}
				{{ .VarName }} {{ .TypeRef }}
		{{- end }}
		{{- if not .ClientBody }}
			{{- if .MustValidate }}
				err error
			{{- end }}
		{{- end }}
			)
		{{- range .Headers }}

		{{- if (or (eq .Type.Name "string") (eq .Type.Name "any")) }}
			{{ .VarName }}Raw := resp.Header.Get("{{ .CanonicalName }}")
			{{- if .Required }}
				if {{ .VarName }}Raw == "" {
					err = goa.MergeErrors(err, goa.MissingFieldError("{{ .Name }}", "header"))
				}
				{{ .VarName }} = {{ if and (eq .Type.Name "string") .Pointer }}&{{ end }}{{ .VarName }}Raw
			{{- else }}
				if {{ .VarName }}Raw != "" {
					{{ .VarName }} = {{ if and (eq .Type.Name "string") .Pointer }}&{{ end }}{{ .VarName }}Raw
				}
				{{- if .DefaultValue }} else {
					{{ .VarName }} = {{ if eq .Type.Name "string" }}{{ printf "%q" .DefaultValue }}{{ else }}{{ printf "%#v" .DefaultValue }}{{ end }}
				}
				{{- end }}
			{{- end }}

		{{- else if .StringSlice }}
			{{ .VarName }} = resp.Header["{{ .CanonicalName }}"]
			{{ if .Required }}
			if {{ .VarName }} == nil {
				err = goa.MergeErrors(err, goa.MissingFieldError("{{ .Name }}", "header"))
			}
			{{- else if .DefaultValue }}
			if {{ .VarName }} == nil {
				{{ .VarName }} = {{ printf "%#v" .DefaultValue }}
			}
			{{- end }}

		{{- else if .Slice }}
		{
			{{ .VarName }}Raw := resp.Header["{{ .CanonicalName }}"]
				{{ if .Required }} if {{ .VarName }}Raw == nil {
				return nil, goahttp.ErrValidationError("{{ $.ServiceName }}", "{{ $.Method.Name }}", goa.MissingFieldError("{{ .Name }}", "header"))
			}
			{{- else if .DefaultValue }}
			if {{ .VarName }}Raw == nil {
				{{ .VarName }} = {{ printf "%#v" .DefaultValue }}
			}
			{{- end }}

			{{- if .DefaultValue }}else {
			{{- else if not .Required }}
			if {{ .VarName }}Raw != nil {
			{{- end }}
				{{- template "slice_conversion" . }}
			{{- if or .DefaultValue (not .Required) }}
			}
			{{- end }}
		}

		{{- else }}{{/* not string, not any and not slice */}}
		{
			{{ .VarName }}Raw := resp.Header.Get("{{ .CanonicalName }}")
			{{- if .Required }}
			if {{ .VarName }}Raw == "" {
				return nil, goahttp.ErrValidationError("{{ $.ServiceName }}", "{{ $.Method.Name }}", goa.MissingFieldError("{{ .Name }}", "header"))
			}
			{{- else if .DefaultValue }}
			if {{ .VarName }}Raw == "" {
				{{ .VarName }} = {{ printf "%#v" .DefaultValue }}
			}
			{{- end }}

			{{- if .DefaultValue }}else {
				{{- else if not .Required }}
			if {{ .VarName }}Raw != "" {
			{{- end }}
				{{- template "type_conversion" . }}
			{{- if or .DefaultValue (not .Required) }}
			}
			{{- end }}
		}
		{{- end }}
		{{- if .Validate }}
			{{ .Validate }}
		{{- end }}
		{{- end }}{{/* range .Headers */}}
	{{- end }}

	{{- if .MustValidate }}
			if err != nil {
				return nil, goahttp.ErrValidationError("{{ $.ServiceName }}", "{{ $.Method.Name }}", err)
			}
	{{- end }}
`

// input: multipartData
const multipartRequestEncoderTypeT = `{{ printf "%s is the type to encode multipart request for the %q service %q endpoint." .FuncName .ServiceName .MethodName | comment }}
type {{ .FuncName }} func(*multipart.Writer, {{ .Payload.Ref }}) error
`

// input: multipartData
const multipartRequestEncoderT = `{{ printf "%s returns an encoder to encode the multipart request for the %q service %q endpoint." .InitName .ServiceName .MethodName | comment }}
func {{ .InitName }}(encoderFn {{ .FuncName }}) func(r *http.Request) goahttp.Encoder {
	return func(r *http.Request) goahttp.Encoder {
		body := &bytes.Buffer{}
		mw := multipart.NewWriter(body)
		return goahttp.EncodingFunc(func(v interface{}) error {
			p := v.({{ .Payload.Ref }})
			if err := encoderFn(mw, p); err != nil {
				return err
			}
			r.Body = ioutil.NopCloser(body)
			r.Header.Set("Content-Type", mw.FormDataContentType())
			return mw.Close()
		})
	}
}
`
