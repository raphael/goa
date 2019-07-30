package codegen

import (
	"bytes"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"text/template"

	"goa.design/goa/v3/codegen"
	"goa.design/goa/v3/codegen/service"
	"goa.design/goa/v3/expr"
)

// HTTPServices holds the data computed from the design needed to generate the
// transport code of the services.
var HTTPServices = make(ServicesData)

var (
	// pathInitTmpl is the template used to render path constructors code.
	pathInitTmpl = template.Must(template.New("path-init").Funcs(template.FuncMap{"goify": codegen.Goify}).Parse(pathInitT))
	// requestInitTmpl is the template used to render request constructors.
	requestInitTmpl = template.Must(template.New("request-init").Parse(requestInitT))
)

type (
	// ServicesData encapsulates the data computed from the design.
	ServicesData map[string]*ServiceData

	// ServiceData contains the data used to render the code related to a
	// single service.
	ServiceData struct {
		// Service contains the related service data.
		Service *service.Data
		// Endpoints describes the endpoint data for this service.
		Endpoints []*EndpointData
		// FileServers lists the file servers for this service.
		FileServers []*FileServerData
		// ServerStruct is the name of the HTTP server struct.
		ServerStruct string
		// MountPointStruct is the name of the mount point struct.
		MountPointStruct string
		// ServerInit is the name of the constructor of the server
		// struct.
		ServerInit string
		// MountServer is the name of the mount function.
		MountServer string
		// ServerService is the name of service function.
		ServerService string
		// ClientStruct is the name of the HTTP client struct.
		ClientStruct string
		// ServerBodyAttributeTypes is the list of user types used to
		// define the request, response and error response type
		// attributes in the server code.
		ServerBodyAttributeTypes []*TypeData
		// ClientBodyAttributeTypes is the list of user types used to
		// define the request, response and error response type
		// attributes in the client code.
		ClientBodyAttributeTypes []*TypeData
		// ServerTypeNames records the user type names used to define
		// the endpoint request and response bodies for server code.
		// The type name is used as the key and a bool as the value
		// which if true indicates that the type has been generated
		// in the server package.
		ServerTypeNames map[string]bool
		// ClientTypeNames records the user type names used to define
		// the endpoint request and response bodies for client code.
		// The type name is used as the key and a bool as the value
		// which if true indicates that the type has been generated
		// in the client package.
		ClientTypeNames map[string]bool
		// ServerTransformHelpers is the list of transform functions
		// required by the various server side constructors.
		ServerTransformHelpers []*codegen.TransformFunctionData
		// ClientTransformHelpers is the list of transform functions
		// required by the various client side constructors.
		ClientTransformHelpers []*codegen.TransformFunctionData
		// Scope initialized with all the server and client types.
		Scope *codegen.NameScope
	}

	// EndpointData contains the data used to render the code related to a
	// single service HTTP endpoint.
	EndpointData struct {
		// Method contains the related service method data.
		Method *service.MethodData
		// ServiceName is the name of the service exposing the endpoint.
		ServiceName string
		// ServiceVarName is the goified service name (first letter
		// lowercase).
		ServiceVarName string
		// ServicePkgName is the name of the service package.
		ServicePkgName string
		// Payload describes the method HTTP payload.
		Payload *PayloadData
		// Result describes the method HTTP result.
		Result *ResultData
		// Errors describes the method HTTP errors.
		Errors []*ErrorGroupData
		// Routes describes the possible routes for this endpoint.
		Routes []*RouteData
		// BasicScheme is the basic auth security scheme if any.
		BasicScheme *service.SchemeData
		// HeaderSchemes lists all the security requirement schemes that
		// apply to the method and are encoded in the request header.
		HeaderSchemes service.SchemesData
		// BodySchemes lists all the security requirement schemes that
		// apply to the method and are encoded in the request body.
		BodySchemes service.SchemesData
		// QuerySchemes lists all the security requirement schemes that
		// apply to the method and are encoded in the request query
		// string.
		QuerySchemes service.SchemesData

		// server

		// MountHandler is the name of the mount handler function.
		MountHandler string
		// HandlerInit is the name of the constructor function for the
		// http handler function.
		HandlerInit string
		// RequestDecoder is the name of the request decoder function.
		RequestDecoder string
		// ResponseEncoder is the name of the response encoder function.
		ResponseEncoder string
		// ErrorEncoder is the name of the error encoder function.
		ErrorEncoder string
		// MultipartRequestDecoder indicates the request decoder for
		// multipart content type.
		MultipartRequestDecoder *MultipartData
		// ServerStream holds the data to render the server struct which
		// implements the server stream interface.
		ServerStream *StreamData

		// client

		// ClientStruct is the name of the HTTP client struct.
		ClientStruct string
		// EndpointInit is the name of the constructor function for the
		// client endpoint.
		EndpointInit string
		// RequestInit is the request builder function.
		RequestInit *InitData
		// RequestEncoder is the name of the request encoder function.
		RequestEncoder string
		// ResponseDecoder is the name of the response decoder function.
		ResponseDecoder string
		// MultipartRequestEncoder indicates the request encoder for
		// multipart content type.
		MultipartRequestEncoder *MultipartData
		// ClientStream holds the data to render the client struct which
		// implements the client stream interface.
		ClientStream *StreamData
	}

	// FileServerData lists the data needed to generate file servers.
	FileServerData struct {
		// MountHandler is the name of the mount handler function.
		MountHandler string
		// RequestPaths is the set of HTTP paths to the server.
		RequestPaths []string
		// Root is the root server file path.
		FilePath string
		// Dir is true if the file server servers files under a
		// directory, false if it serves a single file.
		IsDir bool
		// PathParam is the name of the parameter used to capture the
		// path for file servers that serve files under a directory.
		PathParam string
	}

	// PayloadData contains the payload information required to generate the
	// transport decode (server) and encode (client) code.
	PayloadData struct {
		// Name is the name of the payload type.
		Name string
		// Ref is the fully qualified reference to the payload type.
		Ref string
		// Request contains the data for the corresponding HTTP request.
		Request *RequestData
		// DecoderReturnValue is a reference to the decoder return value
		// if there is no payload constructor (i.e. if Init is nil).
		DecoderReturnValue string
	}

	// ResultData contains the result information required to generate the
	// transport decode (client) and encode (server) code.
	ResultData struct {
		// Name is the name of the result type.
		Name string
		// Ref is the reference to the result type.
		Ref string
		// IsStruct is true if the result type is a user type defining
		// an object.
		IsStruct bool
		// Inits contains the data required to render the result
		// constructors if any.
		Inits []*InitData
		// Responses contains the data for the corresponding HTTP
		// responses.
		Responses []*ResponseData
		// View is the view used to render the result.
		View string
		// MustInit indicates if a variable holding the result type must be
		// initialized. It is used by server response encoder to initialize
		// the result variable only if there are multiple responses, or the
		// response has a body or a header.
		MustInit bool
	}

	// ErrorGroupData contains the error information required to generate
	// the transport decode (client) and encode (server) code for all errors
	// with responses using a given HTTP status code.
	ErrorGroupData struct {
		// StatusCode is the response HTTP status code.
		StatusCode string
		// Errors contains the information for each error.
		Errors []*ErrorData
	}

	// ErrorData contains the error information required to generate the
	// transport decode (client) and encode (server) code.
	ErrorData struct {
		// Name is the error name.
		Name string
		// Ref is a reference to the error type.
		Ref string
		// Response is the error response data.
		Response *ResponseData
	}

	// RequestData describes a request.
	RequestData struct {
		// PathParams describes the information about params that are
		// present in the request path.
		PathParams []*ParamData
		// QueryParams describes the information about the params that
		// are present in the request query string.
		QueryParams []*ParamData
		// Headers contains the HTTP request headers used to build the
		// method payload.
		Headers []*HeaderData
		// ServerBody describes the request body type used by server
		// code. The type is generated using pointers for all fields so
		// that it can be validated.
		ServerBody *TypeData
		// ClientBody describes the request body type used by client
		// code. The type does NOT use pointers for every fields since
		// no validation is required.
		ClientBody *TypeData
		// PayloadInit contains the data required to render the
		// payload constructor used by server code if any.
		PayloadInit *InitData
		// MustValidate is true if the request body or at least one
		// parameter or header requires validation.
		MustValidate bool
		// Multipart if true indicates the request is a multipart
		// request.
		Multipart bool
	}

	// ResponseData describes a response.
	ResponseData struct {
		// StatusCode is the return code of the response.
		StatusCode string
		// Description is the response description.
		Description string
		// Headers provides information about the headers in the
		// response.
		Headers []*HeaderData
		// ContentType contains the value of the response
		// "Content-Type" header.
		ContentType string
		// ErrorHeader contains the value of the response "goa-error"
		// header if any.
		ErrorHeader string
		// ServerBody is the type of the response body used by server
		// code, nil if body should be empty. The type does NOT use
		// pointers for all fields. If the method result is a result
		// type and the response data describes a success response, then
		// this field contains a type for every view in the result type.
		// The type name is suffixed with the name of the view (except
		// for "default" view where no suffix is added). A constructor
		// is also generated server side for each view to transform the
		// result type to the corresponding response body type. If
		// method result is not a result type or if the response
		// describes an error response, then this field contains at most
		// one item.
		ServerBody []*TypeData
		// ClientBody is the type of the response body used by client
		// code, nil if body should be empty. The type uses pointers for
		// all fields so they can be validated.
		ClientBody *TypeData
		// Init contains the data required to render the result or error
		// constructor if any.
		ResultInit *InitData
		// TagName is the name of the attribute used to test whether the
		// response is the one to use.
		TagName string
		// TagValue is the value the result attribute named by TagName
		// must have for this response to be used.
		TagValue string
		// TagPointer is true if the tag attribute is a pointer.
		TagPointer bool
		// MustValidate is true if at least one header requires validation.
		MustValidate bool
		// ResultAttr sets the response body from the specified result
		// type attribute. This field is set when the design uses
		// Body("name") syntax to set the response body and the result
		// type is an object.
		ResultAttr string
		// ViewedResult indicates whether the response body type is a
		// result type.
		ViewedResult *service.ViewedResultTypeData
	}

	// InitData contains the data required to render a constructor.
	InitData struct {
		// Name is the constructor function name.
		Name string
		// Description is the function description.
		Description string
		// ServerArgs is the list of constructor arguments for server
		// side code.
		ServerArgs []*InitArgData
		// ClientArgs is the list of constructor arguments for client
		// side code.
		ClientArgs []*InitArgData
		// CLIArgs is the list of arguments that should be initialized
		// from CLI flags. This is used for implicit attributes which
		// as the time of writing is only used for the basic auth
		// username and password.
		CLIArgs []*InitArgData
		// ReturnTypeName is the qualified (including the package name)
		// name of the payload, result or error type.
		ReturnTypeName string
		// ReturnTypeRef is the qualified (including the package name)
		// reference to the payload, result or error type.
		ReturnTypeRef string
		// ReturnTypeAttribute is the name of the attribute initialized
		// by this constructor when it only initializes one attribute
		// (i.e. body was defined with Body("name") syntax).
		ReturnTypeAttribute string
		// ReturnIsStruct is true if the return type is a struct.
		ReturnIsStruct bool
		// ReturnIsPrimitivePointer indicates whether the return type is
		// a primitive pointer.
		ReturnIsPrimitivePointer bool
		// ReturnTypePkg is the package where the return type is
		// present.
		ReturnTypePkg string
		// ServerCode is the code that builds the payload from the
		// request on the server when it contains user types.
		ServerCode string
		// ClientCode is the code that builds the payload or result type
		// from the request or response state on the client when it
		// contains user types.
		ClientCode string
	}

	// InitArgData represents a single constructor argument.
	InitArgData struct {
		// Name is the argument name.
		Name string
		// Description is the argument description.
		Description string
		// Reference to the argument, e.g. "&body".
		Ref string
		// FieldName is the name of the data structure field that should
		// be initialized with the argument if any.
		FieldName string
		// FieldPointer if true indicates that the data structure field is a
		// pointer.
		FieldPointer bool
		// FieldType is the type of the data structure field that should
		// be initialized with the argument if any.
		FieldType expr.DataType
		// TypeName is the argument type name.
		TypeName string
		// TypeRef is the argument type reference.
		TypeRef string
		// Type is the argument type. It is never an aliased user type.
		Type expr.DataType
		// Pointer is true if a pointer to the arg should be used.
		Pointer bool
		// Required is true if the arg is required to build the payload.
		Required bool
		// DefaultValue is the default value of the arg.
		DefaultValue interface{}
		// Validate contains the validation code for the argument
		// value if any.
		Validate string
		// Example is a example value
		Example interface{}
	}

	// RouteData describes a route.
	RouteData struct {
		// Verb is the HTTP method.
		Verb string
		// Path is the fullpath including wildcards.
		Path string
		// PathInit contains the information needed to render and call
		// the path constructor for the route.
		PathInit *InitData
	}

	// ParamData describes a HTTP request parameter.
	ParamData struct {
		// Name is the name of the mapping to the actual variable name.
		Name string
		// AttributeName is the name of the corresponding attribute.
		AttributeName string
		// Description is the parameter description
		Description string
		// FieldName is the name of the struct field that holds the
		// param value.
		FieldName string
		// FieldPointer if true indicates that the struct field that holds the
		// param value is a pointer.
		FieldPointer bool
		// FieldType is the type of the struct field.
		FieldType expr.DataType
		// VarName is the name of the Go variable used to read or
		// convert the param value.
		VarName string
		// ServiceField is true if there is a corresponding attribute in
		// the service types.
		ServiceField bool
		// Type is the datatype of the variable.
		Type expr.DataType
		// TypeName is the name of the type.
		TypeName string
		// TypeRef is the reference to the type.
		TypeRef string
		// Required is true if the param is required.
		Required bool
		// Pointer is true if and only the param variable is a pointer.
		Pointer bool
		// StringSlice is true if the param type is array of strings.
		StringSlice bool
		// Slice is true if the param type is an array.
		Slice bool
		// MapStringSlice is true if the param type is a map of string
		// slice.
		MapStringSlice bool
		// Map is true if the param type is a map.
		Map bool
		// Validate contains the validation code if any.
		Validate string
		// DefaultValue contains the default value if any.
		DefaultValue interface{}
		// Example is an example value.
		Example interface{}
		// MapQueryParams indicates that the query params must be mapped
		// to the entire payload (empty string) or a payload attribute
		// (attribute name).
		MapQueryParams *string
	}

	// HeaderData describes a HTTP request or response header.
	HeaderData struct {
		// Name is the name of the header key.
		Name string
		// AttributeName is the name of the corresponding attribute.
		AttributeName string
		// Description is the header description.
		Description string
		// CanonicalName is the canonical header key.
		CanonicalName string
		// FieldName is the name of the struct field that holds the
		// header value if any, empty string otherwise.
		FieldName string
		// FieldType is the type of the struct field.
		FieldType expr.DataType
		// FieldPointer if true indicates that the struct field that holds the
		// header value is a pointer.
		FieldPointer bool
		// VarName is the name of the Go variable used to read or
		// convert the header value.
		VarName string
		// TypeName is the name of the type.
		TypeName string
		// TypeRef is the reference to the type.
		TypeRef string
		// Required is true if the header is required.
		Required bool
		// Pointer is true if and only the param variable is a pointer.
		Pointer bool
		// StringSlice is true if the param type is array of strings.
		StringSlice bool
		// Slice is true if the param type is an array.
		Slice bool
		// Type describes the datatype of the variable value. Mainly
		// used for conversion.
		Type expr.DataType
		// Validate contains the validation code if any.
		Validate string
		// DefaultValue contains the default value if any.
		DefaultValue interface{}
		// Example is an example value.
		Example interface{}
	}

	// TypeData contains the data needed to render a type definition.
	TypeData struct {
		// Name is the type name.
		Name string
		// VarName is the Go type name.
		VarName string
		// Description is the type human description.
		Description string
		// Init contains the data needed to render and call the type
		// constructor if any.
		Init *InitData
		// Def is the type definition Go code.
		Def string
		// Ref is the reference to the type.
		Ref string
		// ValidateDef contains the validation code.
		ValidateDef string
		// ValidateRef contains the call to the validation code.
		ValidateRef string
		// Example is an example value for the type.
		Example interface{}
		// View is the view using which the type is rendered.
		View string
	}

	// MultipartData contains the data needed to render multipart
	// encoder/decoder.
	MultipartData struct {
		// FuncName is the name used to generate function type.
		FuncName string
		// InitName is the name of the constructor.
		InitName string
		// VarName is the name of the variable referring to the function.
		VarName string
		// ServiceName is the name of the service.
		ServiceName string
		// MethodName is the name of the method.
		MethodName string
		// Payload is the payload data required to generate
		// encoder/decoder.
		Payload *PayloadData
	}

	// StreamData contains the data needed to render struct type that
	// implements the server and client stream interfaces.
	StreamData struct {
		// VarName is the name of the struct.
		VarName string
		// Type is type of the stream (server or client).
		Type string
		// Interface is the fully qualified name of the interface that
		// the struct implements.
		Interface string
		// Endpoint is endpoint data that defines streaming
		// payload/result.
		Endpoint *EndpointData
		// Payload is the streaming payload type sent via the stream.
		Payload *TypeData
		// Response is the successful response data for the streaming
		// endpoint.
		Response *ResponseData
		// SendName is the name of the send function.
		SendName string
		// SendDesc is the description for the send function.
		SendDesc string
		// SendTypeName is the fully qualified type name sent through
		// the stream.
		SendTypeName string
		// SendTypeRef is the fully qualified type ref sent through the
		// stream.
		SendTypeRef string
		// RecvName is the name of the receive function.
		RecvName string
		// RecvDesc is the description for the recv function.
		RecvDesc string
		// RecvTypeName is the fully qualified type name received from
		// the stream.
		RecvTypeName string
		// RecvTypeRef is the fully qualified type ref received from the
		// stream.
		RecvTypeRef string
		// MustClose indicates whether to generate the Close() function
		// for the stream.
		MustClose bool
		// PkgName is the service package name.
		PkgName string
		// Kind is the kind of the stream (payload, result or
		// bidirectional).
		Kind expr.StreamKind
	}
)

// Get retrieves the transport data for the service with the given name
// computing it if needed. It returns nil if there is no service with the given
// name.
func (d ServicesData) Get(name string) *ServiceData {
	if data, ok := d[name]; ok {
		return data
	}
	service := expr.Root.API.HTTP.Service(name)
	if service == nil {
		return nil
	}
	d[name] = d.analyze(service)
	return d[name]
}

// Endpoint returns the service method transport data for the endpoint with the
// given name, nil if there isn't one.
func (svc *ServiceData) Endpoint(name string) *EndpointData {
	for _, e := range svc.Endpoints {
		if e.Method.Name == name {
			return e
		}
	}
	return nil
}

// analyze creates the data necessary to render the code of the given service.
// It records the user types needed by the service definition in userTypes.
func (d ServicesData) analyze(hs *expr.HTTPServiceExpr) *ServiceData {
	svc := service.Services.Get(hs.ServiceExpr.Name)
	scope := codegen.NewNameScope()
	scope.Unique("c") // 'c' is reserved as the client's receiver name.
	scope.Unique("v") // 'v' is reserved as the request builder payload argument name.
	rd := &ServiceData{
		Service:          svc,
		ServerStruct:     "Server",
		MountPointStruct: "MountPoint",
		ServerInit:       "New",
		MountServer:      "Mount",
		ServerService:    "Service",
		ClientStruct:     "Client",
		ServerTypeNames:  make(map[string]bool),
		ClientTypeNames:  make(map[string]bool),
		Scope:            scope,
	}

	for _, s := range hs.FileServers {
		paths := make([]string, len(s.RequestPaths))
		for i, p := range s.RequestPaths {
			idx := strings.LastIndex(p, "/{")
			if idx > 0 {
				paths[i] = p[:idx]
			} else {
				paths[i] = p
			}
		}
		var pp string
		if s.IsDir() {
			pp = expr.ExtractHTTPWildcards(s.RequestPaths[0])[0]
		}
		data := &FileServerData{
			MountHandler: fmt.Sprintf("Mount%s", codegen.Goify(s.FilePath, true)),
			RequestPaths: paths,
			FilePath:     s.FilePath,
			IsDir:        s.IsDir(),
			PathParam:    pp,
		}
		rd.FileServers = append(rd.FileServers, data)
	}

	for _, a := range hs.HTTPEndpoints {
		ep := svc.Method(a.MethodExpr.Name)

		var routes []*RouteData
		i := 0
		for _, r := range a.Routes {
			for _, rpath := range r.FullPaths() {
				params := expr.ExtractHTTPWildcards(rpath)
				var (
					init *InitData
				)
				{
					initArgs := make([]*InitArgData, len(params))
					pathParamsObj := expr.AsObject(a.PathParams().Type)
					suffix := ""
					if i > 0 {
						suffix = strconv.Itoa(i + 1)
					}
					i++
					name := fmt.Sprintf("%s%sPath%s", ep.VarName, svc.StructName, suffix)
					for j, arg := range params {
						att := pathParamsObj.Attribute(arg)
						pointer := a.Params.IsPrimitivePointer(arg, true)
						name := rd.Scope.Name(codegen.Goify(arg, false))
						var vcode string
						if att.Validation != nil {
							ctx := httpContext("", rd.Scope, true, false)
							vcode = codegen.RecursiveValidationCode(att, ctx, true, name)
						}
						initArgs[j] = &InitArgData{
							Name:        name,
							Description: att.Description,
							Ref:         name,
							FieldName:   codegen.Goify(arg, true),
							TypeName:    rd.Scope.GoTypeName(att),
							TypeRef:     rd.Scope.GoTypeRef(att),
							Pointer:     pointer,
							Required:    true,
							Example:     att.Example(expr.Root.API.Random()),
							Validate:    vcode,
						}
					}

					var buffer bytes.Buffer
					pf := expr.HTTPWildcardRegex.ReplaceAllString(rpath, "/%v")
					err := pathInitTmpl.Execute(&buffer, map[string]interface{}{
						"Args":       initArgs,
						"PathParams": pathParamsObj,
						"PathFormat": pf,
					})
					if err != nil {
						panic(err)
					}
					init = &InitData{
						Name:           name,
						Description:    fmt.Sprintf("%s returns the URL path to the %s service %s HTTP endpoint. ", name, svc.Name, ep.Name),
						ServerArgs:     initArgs,
						ClientArgs:     initArgs,
						ReturnTypeName: "string",
						ReturnTypeRef:  "string",
						ServerCode:     buffer.String(),
						ClientCode:     buffer.String(),
					}
				}

				routes = append(routes, &RouteData{
					Verb:     strings.ToUpper(r.Method),
					Path:     rpath,
					PathInit: init,
				})
			}
		}

		payload := buildPayloadData(a, rd)

		var (
			hsch  service.SchemesData
			bosch service.SchemesData
			qsch  service.SchemesData
			basch *service.SchemeData
		)
		{
			for _, req := range ep.Requirements {
				for _, s := range req.Schemes {
					switch s.Type {
					case "Basic":
						basch = s
					default:
						switch s.In {
						case "query":
							qsch = qsch.Append(s)
						case "header":
							hsch = hsch.Append(s)
						default:
							bosch = bosch.Append(s)
						}
					}
				}
			}
		}

		var requestEncoder string
		{
			if payload.Request.ClientBody != nil || len(payload.Request.Headers) > 0 || len(payload.Request.QueryParams) > 0 || basch != nil {
				requestEncoder = fmt.Sprintf("Encode%sRequest", ep.VarName)
			}
		}

		var requestInit *InitData
		{
			var (
				name       string
				args       []*InitArgData
				payloadRef string
			)
			{
				name = fmt.Sprintf("Build%sRequest", ep.VarName)
				s := codegen.NewNameScope()
				s.Unique("c") // 'c' is reserved as the client's receiver name.
				for _, ca := range routes[0].PathInit.ClientArgs {
					if ca.FieldName != "" {
						ca.Name = s.Unique(ca.Name)
						ca.Ref = ca.Name
						args = append(args, ca)
					}
				}
				if len(routes[0].PathInit.ClientArgs) > 0 && a.MethodExpr.Payload.Type != expr.Empty {
					payloadRef = svc.Scope.GoFullTypeRef(a.MethodExpr.Payload, svc.PkgName)
				}
			}
			data := map[string]interface{}{
				"PayloadRef":   payloadRef,
				"HasFields":    expr.IsObject(a.MethodExpr.Payload.Type),
				"ServiceName":  svc.Name,
				"EndpointName": ep.Name,
				"Args":         args,
				"PathInit":     routes[0].PathInit,
				"Verb":         routes[0].Verb,
				"IsStreaming":  a.MethodExpr.IsStreaming(),
			}
			var buf bytes.Buffer
			if err := requestInitTmpl.Execute(&buf, data); err != nil {
				panic(err) // bug
			}
			requestInit = &InitData{
				Name:        name,
				Description: fmt.Sprintf("%s instantiates a HTTP request object with method and path set to call the %q service %q endpoint", name, svc.Name, ep.Name),
				ClientCode:  buf.String(),
				ClientArgs: []*InitArgData{{
					Name:    "v",
					Ref:     "v",
					TypeRef: "interface{}",
				}},
			}
		}

		ad := &EndpointData{
			Method:          ep,
			ServiceName:     svc.Name,
			ServiceVarName:  svc.VarName,
			ServicePkgName:  svc.PkgName,
			Payload:         payload,
			Result:          buildResultData(a, rd),
			Errors:          buildErrorsData(a, rd),
			HeaderSchemes:   hsch,
			BodySchemes:     bosch,
			QuerySchemes:    qsch,
			BasicScheme:     basch,
			Routes:          routes,
			MountHandler:    fmt.Sprintf("Mount%sHandler", ep.VarName),
			HandlerInit:     fmt.Sprintf("New%sHandler", ep.VarName),
			RequestDecoder:  fmt.Sprintf("Decode%sRequest", ep.VarName),
			ResponseEncoder: fmt.Sprintf("Encode%sResponse", ep.VarName),
			ErrorEncoder:    fmt.Sprintf("Encode%sError", ep.VarName),
			ClientStruct:    "Client",
			EndpointInit:    ep.VarName,
			RequestInit:     requestInit,
			RequestEncoder:  requestEncoder,
			ResponseDecoder: fmt.Sprintf("Decode%sResponse", ep.VarName),
		}
		buildStreamData(ad, a, rd)

		if a.MultipartRequest {
			ad.MultipartRequestDecoder = &MultipartData{
				FuncName:    fmt.Sprintf("%s%sDecoderFunc", svc.StructName, ep.VarName),
				InitName:    fmt.Sprintf("New%s%sDecoder", svc.StructName, ep.VarName),
				VarName:     fmt.Sprintf("%s%sDecoderFn", svc.VarName, ep.VarName),
				ServiceName: svc.Name,
				MethodName:  ep.Name,
				Payload:     ad.Payload,
			}
			ad.MultipartRequestEncoder = &MultipartData{
				FuncName:    fmt.Sprintf("%s%sEncoderFunc", svc.StructName, ep.VarName),
				InitName:    fmt.Sprintf("New%s%sEncoder", svc.StructName, ep.VarName),
				VarName:     fmt.Sprintf("%s%sEncoderFn", svc.VarName, ep.VarName),
				ServiceName: svc.Name,
				MethodName:  ep.Name,
				Payload:     ad.Payload,
			}
		}

		rd.Endpoints = append(rd.Endpoints, ad)
	}

	for _, a := range hs.HTTPEndpoints {
		collectUserTypes(a.Body.Type, func(ut expr.UserType) {
			if d := attributeTypeData(ut, true, true, true, rd); d != nil {
				rd.ServerBodyAttributeTypes = append(rd.ServerBodyAttributeTypes, d)
			}
			if d := attributeTypeData(ut, true, false, false, rd); d != nil {
				rd.ClientBodyAttributeTypes = append(rd.ClientBodyAttributeTypes, d)
			}
		})

		if a.MethodExpr.StreamingPayload.Type != expr.Empty {
			collectUserTypes(a.StreamingBody.Type, func(ut expr.UserType) {
				if d := attributeTypeData(ut, true, true, true, rd); d != nil {
					rd.ServerBodyAttributeTypes = append(rd.ServerBodyAttributeTypes, d)
				}
				if d := attributeTypeData(ut, true, false, false, rd); d != nil {
					rd.ClientBodyAttributeTypes = append(rd.ClientBodyAttributeTypes, d)
				}
			})
		}

		if res := a.MethodExpr.Result; res != nil {
			for _, v := range a.Responses {
				collectUserTypes(v.Body.Type, func(ut expr.UserType) {
					// NOTE: ServerBodyAttributeTypes for response body types are
					// collected in buildResponseBodyType because we have to generate
					// body types for each view in a result type.
					if d := attributeTypeData(ut, false, true, false, rd); d != nil {
						rd.ClientBodyAttributeTypes = append(rd.ClientBodyAttributeTypes, d)
					}
				})
			}
		}

		for _, v := range a.HTTPErrors {
			collectUserTypes(v.Response.Body.Type, func(ut expr.UserType) {
				// NOTE: ServerBodyAttributeTypes for error response body types are
				// collected in buildResponseBodyType because we have to generate
				// body types for each view in a result type.
				if d := attributeTypeData(ut, false, true, false, rd); d != nil {
					rd.ClientBodyAttributeTypes = append(rd.ClientBodyAttributeTypes, d)
				}
			})
		}
	}

	return rd
}

// makeHTTPType traverses the attribute recursively and performs these actions
//
// * removes aliased user type by replacing them with the underlying type
//
func makeHTTPType(att *expr.AttributeExpr, seen ...map[string]struct{}) {
	if att == nil {
		return
	}
	switch dt := att.Type.(type) {
	case expr.UserType:
		if _, ok := dt.(*expr.ResultTypeExpr); !ok && !expr.IsObject(dt) {
			// Aliased user type. Use the underlying aliased type instead of
			// generating new types in the client and server packages
			att.Type = dt.Attribute().Type
			if v := dt.Attribute().Validation; v != nil {
				if att.Validation == nil {
					att.Validation = v
				} else {
					att.Validation.Merge(v)
				}
			}
		}
		var s map[string]struct{}
		if len(seen) > 0 {
			s = seen[0]
		} else {
			s = make(map[string]struct{})
			seen = append(seen, s)
		}
		if _, ok := s[dt.ID()]; ok {
			return
		}
		s[dt.ID()] = struct{}{}
		makeHTTPType(dt.Attribute(), seen...)
	case *expr.Array:
		makeHTTPType(dt.ElemType, seen...)
	case *expr.Map:
		makeHTTPType(dt.KeyType, seen...)
		makeHTTPType(dt.ElemType, seen...)
	case *expr.Object:
		for _, nat := range *dt {
			makeHTTPType(nat.Attribute, seen...)
		}
	}
}

// buildPayloadData returns the data structure used to describe the endpoint
// payload including the HTTP request details. It also returns the user types
// used by the request body type recursively if any.
func buildPayloadData(e *expr.HTTPEndpointExpr, sd *ServiceData) *PayloadData {
	makeHTTPType(e.Body)
	var (
		payload    = e.MethodExpr.Payload
		svc        = sd.Service
		body       = e.Body.Type
		ep         = svc.Method(e.MethodExpr.Name)
		httpsvrctx = httpContext("", sd.Scope, true, true)
		httpclictx = httpContext("", sd.Scope, true, false)
		svcctx     = serviceContext(sd.Service.PkgName, sd.Service.Scope)

		request       *RequestData
		mapQueryParam *ParamData
	)
	{
		var (
			serverBodyData = buildRequestBodyType(e.Body, payload, e, true, sd)
			clientBodyData = buildRequestBodyType(e.Body, payload, e, false, sd)
			paramsData     = extractPathParams(e.PathParams(), payload, sd.Scope)
			queryData      = extractQueryParams(e.QueryParams(), payload, sd.Scope)
			headersData    = extractHeaders(e.Headers, payload, svcctx, sd.Scope)

			mustValidate bool
		)
		{
			if e.MapQueryParams != nil {
				var (
					fieldName string
					name      = "query"
					required  = true
					pAtt      = payload
				)
				if n := *e.MapQueryParams; n != "" {
					pAtt = expr.AsObject(payload.Type).Attribute(n)
					required = payload.IsRequired(n)
					name = n
					fieldName = codegen.Goify(name, true)
				}
				varn := codegen.Goify(name, false)
				mapQueryParam = &ParamData{
					Name:           name,
					VarName:        varn,
					FieldName:      fieldName,
					FieldType:      pAtt.Type,
					Required:       required,
					Type:           pAtt.Type,
					TypeName:       sd.Scope.GoTypeName(pAtt),
					TypeRef:        sd.Scope.GoTypeRef(pAtt),
					Map:            expr.AsMap(payload.Type) != nil,
					Validate:       codegen.RecursiveValidationCode(pAtt, httpsvrctx, required, varn),
					DefaultValue:   pAtt.DefaultValue,
					Example:        pAtt.Example(expr.Root.API.Random()),
					MapQueryParams: e.MapQueryParams,
				}
				queryData = append(queryData, mapQueryParam)
			}
			if serverBodyData != nil {
				sd.ServerTypeNames[serverBodyData.Name] = false
				sd.ClientTypeNames[serverBodyData.Name] = false
			}
			for _, p := range paramsData {
				if p.Validate != "" || needConversion(p.Type) {
					mustValidate = true
					break
				}
			}
			if !mustValidate {
				for _, q := range queryData {
					if q.Validate != "" || q.Required || needConversion(q.Type) {
						mustValidate = true
						break
					}
				}
			}
			if !mustValidate {
				for _, h := range headersData {
					if h.Validate != "" || h.Required || needConversion(h.Type) {
						mustValidate = true
						break
					}
				}
			}
		}
		request = &RequestData{
			PathParams:   paramsData,
			QueryParams:  queryData,
			Headers:      headersData,
			ServerBody:   serverBodyData,
			ClientBody:   clientBodyData,
			MustValidate: mustValidate,
			Multipart:    e.MultipartRequest,
		}
	}

	var init *InitData
	if needInit(payload.Type) {
		// generate constructor function to transform request body,
		// params, and headers into the method payload type
		var (
			name       string
			desc       string
			isObject   bool
			clientArgs []*InitArgData
			serverArgs []*InitArgData
		)
		n := codegen.Goify(ep.Name, true)
		p := codegen.Goify(ep.Payload, true)
		// Raw payload object has type name prefixed with endpoint name. No need to
		// prefix the type name again.
		if strings.HasPrefix(p, n) {
			p = svc.Scope.HashedUnique(payload.Type, p)
			name = fmt.Sprintf("New%s", p)
		} else {
			name = fmt.Sprintf("New%s%s", n, p)
		}
		desc = fmt.Sprintf("%s builds a %s service %s endpoint payload.",
			name, svc.Name, e.Name())
		isObject = expr.IsObject(payload.Type)
		if body != expr.Empty {
			var (
				svcode string
				cvcode string
			)
			if ut, ok := body.(expr.UserType); ok {
				if val := ut.Attribute().Validation; val != nil {
					svcode = codegen.RecursiveValidationCode(ut.Attribute(), httpsvrctx, true, "body")
					cvcode = codegen.RecursiveValidationCode(ut.Attribute(), httpclictx, true, "body")
				}
			}
			serverArgs = []*InitArgData{{
				Name:     "body",
				Ref:      sd.Scope.GoVar("body", body),
				TypeName: sd.Scope.GoTypeName(e.Body),
				TypeRef:  sd.Scope.GoTypeRef(e.Body),
				Type:     body,
				Required: true,
				Example:  e.Body.Example(expr.Root.API.Random()),
				Validate: svcode,
			}}
			clientArgs = []*InitArgData{{
				Name:     "body",
				Ref:      sd.Scope.GoVar("body", body),
				TypeName: sd.Scope.GoTypeName(e.Body),
				TypeRef:  sd.Scope.GoTypeRef(e.Body),
				Type:     body,
				Required: true,
				Example:  e.Body.Example(expr.Root.API.Random()),
				Validate: cvcode,
			}}
		}
		var args []*InitArgData
		for _, p := range request.PathParams {
			args = append(args, &InitArgData{
				Name:         p.VarName,
				Description:  p.Description,
				Ref:          p.VarName,
				FieldName:    p.FieldName,
				FieldPointer: p.FieldPointer,
				FieldType:    p.FieldType,
				TypeName:     p.TypeName,
				TypeRef:      p.TypeRef,
				Type:         p.Type,
				Pointer:      p.Pointer,
				Required:     p.Required,
				Validate:     p.Validate,
				Example:      p.Example,
			})
		}
		for _, p := range request.QueryParams {
			args = append(args, &InitArgData{
				Name:         p.VarName,
				Ref:          p.VarName,
				FieldName:    p.FieldName,
				FieldPointer: p.FieldPointer,
				FieldType:    p.FieldType,
				TypeName:     p.TypeName,
				TypeRef:      p.TypeRef,
				Type:         p.Type,
				Pointer:      p.Pointer,
				Required:     p.Required,
				DefaultValue: p.DefaultValue,
				Validate:     p.Validate,
				Example:      p.Example,
			})
		}
		for _, h := range request.Headers {
			args = append(args, &InitArgData{
				Name:         h.VarName,
				Ref:          h.VarName,
				FieldName:    h.FieldName,
				FieldPointer: h.FieldPointer,
				FieldType:    h.FieldType,
				TypeName:     h.TypeName,
				TypeRef:      h.TypeRef,
				Type:         h.Type,
				Pointer:      h.Pointer,
				Required:     h.Required,
				DefaultValue: h.DefaultValue,
				Validate:     h.Validate,
				Example:      h.Example,
			})
		}
		serverArgs = append(serverArgs, args...)
		clientArgs = append(clientArgs, args...)

		var (
			cliArgs []*InitArgData
		)
		for _, r := range ep.Requirements {
			done := false
			for _, sc := range r.Schemes {
				if sc.Type == "Basic" {
					uatt := e.MethodExpr.Payload.Find(sc.UsernameAttr)
					uref := svc.Scope.GoTypeRef(uatt)
					if sc.UsernamePointer {
						uref = "*" + uref
					}
					uarg := &InitArgData{
						Name:         sc.UsernameAttr,
						FieldName:    sc.UsernameField,
						FieldPointer: sc.UsernamePointer,
						FieldType:    uatt.Type,
						Description:  uatt.Description,
						Ref:          sc.UsernameAttr,
						Required:     sc.UsernameRequired,
						TypeName:     svc.Scope.GoTypeName(uatt),
						TypeRef:      uref,
						Type:         uatt.Type,
						Pointer:      sc.UsernamePointer,
						Validate:     codegen.RecursiveValidationCode(uatt, httpsvrctx, sc.UsernameRequired, sc.UsernameAttr),
						Example:      uatt.Example(expr.Root.API.Random()),
					}
					patt := e.MethodExpr.Payload.Find(sc.PasswordAttr)
					pref := svc.Scope.GoTypeRef(patt)
					if sc.PasswordPointer {
						pref = "*" + pref
					}
					parg := &InitArgData{
						Name:         sc.PasswordAttr,
						FieldName:    sc.PasswordField,
						FieldPointer: sc.PasswordPointer,
						FieldType:    patt.Type,
						Description:  patt.Description,
						Ref:          sc.PasswordAttr,
						Required:     sc.PasswordRequired,
						TypeName:     svc.Scope.GoTypeName(patt),
						TypeRef:      pref,
						Type:         patt.Type,
						Pointer:      sc.PasswordPointer,
						Validate:     codegen.RecursiveValidationCode(patt, httpsvrctx, sc.PasswordRequired, sc.PasswordAttr),
						Example:      patt.Example(expr.Root.API.Random()),
					}
					cliArgs = []*InitArgData{uarg, parg}
					done = true
					break
				}
			}
			if done {
				break
			}
		}

		var (
			serverCode string
			clientCode string
			err        error
			origin     string

			pAtt = payload
		)
		if body != expr.Empty {
			// If design uses Body("name") syntax then need to use payload
			// attribute to transform.
			if o, ok := e.Body.Meta["origin:attribute"]; ok {
				origin = o[0]
				pAtt = expr.AsObject(payload.Type).Attribute(origin)
			}

			var (
				helpers []*codegen.TransformFunctionData
			)
			serverCode, helpers, err = unmarshal(e.Body, pAtt, "body", "v", httpsvrctx, svcctx)
			if err == nil {
				sd.ServerTransformHelpers = codegen.AppendHelpers(sd.ServerTransformHelpers, helpers)
			}
			// The client code for building the method payload from a request
			// body is used by the CLI tool to build the payload given to the
			// client endpoint. It differs because the body type there does not
			// use pointers for all fields (no need to validate).
			clientCode, helpers, err = marshal(e.Body, pAtt, "body", "v", httpclictx, svcctx)
			if err == nil {
				sd.ClientTransformHelpers = codegen.AppendHelpers(sd.ClientTransformHelpers, helpers)
			}
		} else if expr.IsArray(payload.Type) || expr.IsMap(payload.Type) {
			if params := expr.AsObject(e.Params.Type); len(*params) > 0 {
				var helpers []*codegen.TransformFunctionData
				serverCode, helpers, err = unmarshal((*params)[0].Attribute, payload, codegen.Goify((*params)[0].Name, false), "v", httpsvrctx, svcctx)
				if err == nil {
					sd.ServerTransformHelpers = codegen.AppendHelpers(sd.ServerTransformHelpers, helpers)
				}
				clientCode, helpers, err = marshal((*params)[0].Attribute, payload, codegen.Goify((*params)[0].Name, false), "v", httpclictx, svcctx)
				if err == nil {
					sd.ClientTransformHelpers = codegen.AppendHelpers(sd.ClientTransformHelpers, helpers)
				}
			}
		}
		if err != nil {
			fmt.Println(err.Error()) // TBD validate DSL so errors are not possible
		}
		init = &InitData{
			Name:                name,
			Description:         desc,
			ServerArgs:          serverArgs,
			ClientArgs:          clientArgs,
			CLIArgs:             cliArgs,
			ReturnTypeName:      svc.Scope.GoFullTypeName(payload, svc.PkgName),
			ReturnTypeRef:       svc.Scope.GoFullTypeRef(payload, svc.PkgName),
			ReturnIsStruct:      isObject,
			ReturnTypeAttribute: codegen.Goify(origin, true),
			ReturnTypePkg:       svc.PkgName,
			ServerCode:          serverCode,
			ClientCode:          clientCode,
		}
	}
	request.PayloadInit = init

	var (
		returnValue string
		name        string
		ref         string
	)
	{
		if payload.Type != expr.Empty {
			name = svc.Scope.GoFullTypeName(payload, svc.PkgName)
			ref = svc.Scope.GoFullTypeRef(payload, svc.PkgName)
		}
		if init == nil {
			if o := expr.AsObject(e.Params.Type); o != nil && len(*o) > 0 {
				returnValue = codegen.Goify((*o)[0].Name, false)
			} else if o := expr.AsObject(e.Headers.Type); o != nil && len(*o) > 0 {
				returnValue = codegen.Goify((*o)[0].Name, false)
			} else if e.MapQueryParams != nil && *e.MapQueryParams == "" {
				returnValue = mapQueryParam.Name
			}
		}
	}

	return &PayloadData{
		Name:               name,
		Ref:                ref,
		Request:            request,
		DecoderReturnValue: returnValue,
	}
}

// buildResultData builds the result data for the given service endpoint.
func buildResultData(e *expr.HTTPEndpointExpr, sd *ServiceData) *ResultData {
	var (
		svc    = sd.Service
		ep     = svc.Method(e.MethodExpr.Name)
		result = e.MethodExpr.Result

		name string
		ref  string
		view string
	)
	{
		view = "default"
		if result.Meta != nil {
			if v, ok := result.Meta["view"]; ok {
				view = v[0]
			}
		}
		if result.Type != expr.Empty {
			name = svc.Scope.GoFullTypeName(result, svc.PkgName)
			ref = svc.Scope.GoFullTypeRef(result, svc.PkgName)
		}
	}

	var (
		mustInit  bool
		responses []*ResponseData
	)
	{
		viewed := false
		if ep.ViewedResult != nil {
			result = expr.AsObject(ep.ViewedResult.Type).Attribute("projected")
			viewed = true
		}
		responses = buildResponses(e, result, viewed, sd)
		for _, r := range responses {
			// response has a body or headers or tag
			if len(r.ServerBody) > 0 || len(r.Headers) > 0 || r.TagName != "" {
				mustInit = true
			}
		}
	}
	return &ResultData{
		IsStruct:  expr.IsObject(result.Type),
		Name:      name,
		Ref:       ref,
		Responses: responses,
		View:      view,
		MustInit:  mustInit,
	}
}

// buildResponses builds the response data for all the responses in the
// endpoint expression. The response headers and body for each response
// are inferred from the method's result expression if not specified
// explicitly.
//
// viewed parameter indicates if the method result uses views.
func buildResponses(e *expr.HTTPEndpointExpr, result *expr.AttributeExpr, viewed bool, sd *ServiceData) []*ResponseData {
	var (
		responses []*ResponseData
		scope     *codegen.NameScope

		svc        = sd.Service
		md         = svc.Method(e.Name())
		httpclictx = httpContext("", sd.Scope, false, false)
		svcctx     = serviceContext(sd.Service.PkgName, sd.Service.Scope)
	)
	{
		scope = svc.Scope
		if viewed {
			scope = svc.ViewScope
			svcctx = viewContext(sd.Service.ViewsPkg, sd.Service.ViewScope)
		}
		notag := -1
		for i, resp := range e.Responses {
			makeHTTPType(resp.Body)
			if resp.Tag[0] == "" {
				if notag > -1 {
					continue // we don't want more than one response with no tag
				}
				notag = i
			}
			var (
				headersData    []*HeaderData
				serverBodyData []*TypeData
				clientBodyData *TypeData
				init           *InitData
				origin         string
				mustValidate   bool

				resAttr = result
			)
			{
				headersData = extractHeaders(resp.Headers, result, svcctx, scope)
				if resp.Body.Type != expr.Empty {
					// If design uses Body("name") syntax we need to use the
					// corresponding attribute in the result type for body
					// transformation.
					if o, ok := resp.Body.Meta["origin:attribute"]; ok {
						origin = o[0]
						resAttr = expr.AsObject(resAttr.Type).Attribute(origin)
					}
				}
				if viewed {
					vname := ""
					if origin != "" {
						// Response body is explicitly set to an attribute in the method
						// result type. No need to do any view-based projections server side.
						if sbd := buildResponseBodyType(resp.Body, result, e, true, &vname, sd); sbd != nil {
							serverBodyData = append(serverBodyData, sbd)
						}
					} else if v, ok := e.MethodExpr.Result.Meta["view"]; ok && len(v) > 0 {
						// Design explicitly sets the view to render the result.
						// We generate only one server body type which will be rendered
						// using the specified view.
						if sbd := buildResponseBodyType(resp.Body, result, e, true, &v[0], sd); sbd != nil {
							serverBodyData = append(serverBodyData, sbd)
						}
					} else {
						// If a method result uses views (i.e., a result type), we generate
						// one response body type per view defined in the result type. The
						// generated body type names are suffixed with the name of the view
						// (except for "default" view). Constructors are also generated to
						// create a view-specific body type from the method result. This
						// makes it possible for the server side to return only the
						// attributes defined in the view in the response (NOTE: a required
						// attribute in the result type may not be present in all its views)
						for _, view := range md.ViewedResult.Views {
							if sbd := buildResponseBodyType(resp.Body, result, e, true, &view.Name, sd); sbd != nil {
								serverBodyData = append(serverBodyData, sbd)
							}
						}
					}
					clientBodyData = buildResponseBodyType(resp.Body, result, e, false, &vname, sd)
				} else {
					if sbd := buildResponseBodyType(resp.Body, result, e, true, nil, sd); sbd != nil {
						serverBodyData = append(serverBodyData, sbd)
					}
					clientBodyData = buildResponseBodyType(resp.Body, result, e, false, nil, sd)
				}
				if clientBodyData != nil {
					sd.ClientTypeNames[clientBodyData.Name] = false
				}
				for _, h := range headersData {
					if h.Validate != "" || h.Required || needConversion(h.Type) {
						mustValidate = true
						break
					}
				}
				if needInit(result.Type) {
					// generate constructor function to transform response body
					// and headers into the method result type
					var (
						name       string
						desc       string
						code       string
						tname      string
						tref       string
						err        error
						pointer    bool
						clientArgs []*InitArgData
						helpers    []*codegen.TransformFunctionData
					)
					{
						tname = svc.Scope.GoFullTypeName(result, svc.PkgName)
						tref = svc.Scope.GoFullTypeRef(result, svc.PkgName)
						if viewed {
							tname = svc.ViewScope.GoFullTypeName(result, svc.ViewsPkg)
							tref = svc.ViewScope.GoFullTypeRef(result, svc.ViewsPkg)
						}
						status := codegen.Goify(http.StatusText(resp.StatusCode), true)
						n := codegen.Goify(md.Name, true)
						r := codegen.Goify(md.Result, true)
						// Raw result object has type name prefixed with endpoint name. No need to
						// prefix the type name again.
						if strings.HasPrefix(r, n) {
							r = scope.HashedUnique(result.Type, r)
							name = fmt.Sprintf("New%s%s", r, status)
						} else {
							name = fmt.Sprintf("New%s%s%s", n, r, status)
						}
						desc = fmt.Sprintf("%s builds a %q service %q endpoint result from a HTTP %q response.", name, svc.Name, e.Name(), status)
						if resp.Body.Type != expr.Empty {
							if origin != "" {
								pointer = result.IsPrimitivePointer(origin, true)
							}
							ref := "body"
							if expr.IsObject(resp.Body.Type) {
								ref = "&body"
								pointer = false
							}
							var vcode string
							if ut, ok := resp.Body.Type.(expr.UserType); ok {
								if val := ut.Attribute().Validation; val != nil {
									vcode = codegen.RecursiveValidationCode(ut.Attribute(), httpclictx, true, "body")
								}
							}
							clientArgs = []*InitArgData{{
								Name:     "body",
								Ref:      ref,
								TypeRef:  sd.Scope.GoTypeRef(resp.Body),
								Validate: vcode,
							}}
							// If the method result is a
							// * result type - we unmarshal the client response body to the
							//   corresponding type in the views package so that view-specific
							//   validation logic can be applied.
							// * user type - we unmarshal the client response body to the
							//   corresponding type in the service package after validating the
							//   response body. Here, the transformation code must rely that the
							//   required attributes are set in the response body (otherwise
							//   validation would fail).
							code, helpers, err = unmarshal(resp.Body, resAttr, "body", "v", httpclictx, svcctx)
							if err == nil {
								sd.ClientTransformHelpers = codegen.AppendHelpers(sd.ClientTransformHelpers, helpers)
							}
						} else if expr.IsArray(result.Type) || expr.IsMap(result.Type) {
							if params := expr.AsObject(e.QueryParams().Type); len(*params) > 0 {
								code, helpers, err = unmarshal((*params)[0].Attribute, result, codegen.Goify((*params)[0].Name, false), "v", httpclictx, svcctx)
								if err == nil {
									sd.ClientTransformHelpers = codegen.AppendHelpers(sd.ClientTransformHelpers, helpers)
								}
							}
						}
						if err != nil {
							fmt.Println(err.Error()) // TBD validate DSL so errors are not possible
						}
						for _, h := range headersData {
							clientArgs = append(clientArgs, &InitArgData{
								Name:         h.VarName,
								Ref:          h.VarName,
								FieldName:    h.FieldName,
								FieldPointer: h.FieldPointer,
								FieldType:    h.FieldType,
								Required:     h.Required,
								Pointer:      h.Pointer,
								TypeRef:      h.TypeRef,
								Type:         h.Type,
								Validate:     h.Validate,
								Example:      h.Example,
							})
						}
					}
					init = &InitData{
						Name:                     name,
						Description:              desc,
						ClientArgs:               clientArgs,
						ReturnTypeName:           tname,
						ReturnTypeRef:            tref,
						ReturnIsStruct:           expr.IsObject(result.Type),
						ReturnTypeAttribute:      codegen.Goify(origin, true),
						ReturnTypePkg:            svc.PkgName,
						ReturnIsPrimitivePointer: pointer,
						ClientCode:               code,
					}
				}

				var (
					tagName string
					tagVal  string
					tagPtr  bool
				)
				{
					if resp.Tag[0] != "" {
						tagName = codegen.Goify(resp.Tag[0], true)
						tagVal = resp.Tag[1]
						tagPtr = viewed || result.IsPrimitivePointer(resp.Tag[0], true)
					}
				}
				responses = append(responses, &ResponseData{
					StatusCode:   statusCodeToHTTPConst(resp.StatusCode),
					Description:  resp.Description,
					Headers:      headersData,
					ContentType:  resp.ContentType,
					ServerBody:   serverBodyData,
					ClientBody:   clientBodyData,
					ResultInit:   init,
					TagName:      tagName,
					TagValue:     tagVal,
					TagPointer:   tagPtr,
					MustValidate: mustValidate,
					ResultAttr:   codegen.Goify(origin, true),
					ViewedResult: md.ViewedResult,
				})
			}
		}
		count := len(responses)
		if notag >= 0 && notag < count-1 {
			// Make sure tagless response is last
			responses[notag], responses[count-1] = responses[count-1], responses[notag]
		}
	}
	return responses
}

// buildErrorsData builds the error data for all the error responses in the
// endpoint expression. The response headers and body for each response
// are inferred from the method's error expression if not specified
// explicitly.
func buildErrorsData(e *expr.HTTPEndpointExpr, sd *ServiceData) []*ErrorGroupData {
	var (
		svc        = sd.Service
		httpclictx = httpContext("", sd.Scope, false, false)
		svcctx     = serviceContext(sd.Service.PkgName, sd.Service.Scope)
	)

	data := make(map[string][]*ErrorData)
	for _, v := range e.HTTPErrors {
		var (
			init *InitData
			body = v.Response.Body.Type
		)
		if needInit(v.ErrorExpr.Type) {
			var (
				name     string
				desc     string
				isObject bool
				args     []*InitArgData
			)
			{
				ep := svc.Method(e.MethodExpr.Name)
				name = fmt.Sprintf("New%s%s", codegen.Goify(ep.Name, true), codegen.Goify(v.ErrorExpr.Name, true))
				desc = fmt.Sprintf("%s builds a %s service %s endpoint %s error.",
					name, svc.Name, e.Name(), v.ErrorExpr.Name)
				if body != expr.Empty {
					isObject = expr.IsObject(body)
					ref := "body"
					if isObject {
						ref = "&body"
					}
					args = []*InitArgData{{
						Name:    "body",
						Ref:     ref,
						TypeRef: sd.Scope.GoTypeRef(v.Response.Body),
					}}
				}
				for _, h := range extractHeaders(v.Response.Headers, v.ErrorExpr.AttributeExpr, svcctx, sd.Scope) {
					args = append(args, &InitArgData{
						Name:         h.VarName,
						Ref:          h.VarName,
						FieldName:    h.FieldName,
						FieldPointer: h.FieldPointer,
						FieldType:    h.FieldType,
						TypeRef:      h.TypeRef,
						Type:         h.Type,
						Validate:     h.Validate,
						Example:      h.Example,
					})
				}
			}

			var (
				code   string
				origin string
				err    error
			)
			{
				if body != expr.Empty {
					eAtt := v.ErrorExpr.AttributeExpr
					// If design uses Body("name") syntax then need to use payload
					// attribute to transform.
					if o, ok := v.Response.Body.Meta["origin:attribute"]; ok {
						origin = o[0]
						eAtt = expr.AsObject(v.ErrorExpr.Type).Attribute(origin)
					}

					var helpers []*codegen.TransformFunctionData
					code, helpers, err = unmarshal(v.Response.Body, eAtt, "body", "v", httpclictx, svcctx)
					if err == nil {
						sd.ClientTransformHelpers = codegen.AppendHelpers(sd.ClientTransformHelpers, helpers)
					}
				} else if expr.IsArray(v.ErrorExpr.Type) || expr.IsMap(v.ErrorExpr.Type) {
					if params := expr.AsObject(e.QueryParams().Type); len(*params) > 0 {
						var helpers []*codegen.TransformFunctionData
						code, helpers, err = unmarshal((*params)[0].Attribute, v.ErrorExpr.AttributeExpr, codegen.Goify((*params)[0].Name, false), "v", httpclictx, svcctx)
						if err == nil {
							sd.ClientTransformHelpers = codegen.AppendHelpers(sd.ClientTransformHelpers, helpers)
						}
					}
				}
				if err != nil {
					fmt.Println(err.Error()) // TBD validate DSL so errors are not possible
				}
			}

			init = &InitData{
				Name:                name,
				Description:         desc,
				ClientArgs:          args,
				ReturnTypeName:      svc.Scope.GoFullTypeName(v.ErrorExpr.AttributeExpr, svc.PkgName),
				ReturnTypeRef:       svc.Scope.GoFullTypeRef(v.ErrorExpr.AttributeExpr, svc.PkgName),
				ReturnIsStruct:      isObject,
				ReturnTypeAttribute: codegen.Goify(origin, true),
				ReturnTypePkg:       svc.PkgName,
				ClientCode:          code,
			}
		}

		var (
			responseData *ResponseData
		)
		{
			var (
				serverBodyData []*TypeData
				clientBodyData *TypeData
			)
			{
				if sbd := buildResponseBodyType(v.Response.Body, v.ErrorExpr.AttributeExpr, e, true, nil, sd); sbd != nil {
					serverBodyData = append(serverBodyData, sbd)
				}
				clientBodyData = buildResponseBodyType(v.Response.Body, v.ErrorExpr.AttributeExpr, e, false, nil, sd)
				if clientBodyData != nil {
					sd.ClientTypeNames[clientBodyData.Name] = false
					clientBodyData.Description = fmt.Sprintf("%s is the type of the %q service %q endpoint HTTP response body for the %q error.",
						clientBodyData.VarName, svc.Name, e.Name(), v.Name)
					serverBodyData[0].Description = fmt.Sprintf("%s is the type of the %q service %q endpoint HTTP response body for the %q error.",
						serverBodyData[0].VarName, svc.Name, e.Name(), v.Name)
				}
			}

			headers := extractHeaders(v.Response.Headers, v.ErrorExpr.AttributeExpr, svcctx, sd.Scope)
			var mustValidate bool
			{
				for _, h := range headers {
					if h.Validate != "" || h.Required || needConversion(h.Type) {
						mustValidate = true
						break
					}
				}
			}
			responseData = &ResponseData{
				StatusCode:   statusCodeToHTTPConst(v.Response.StatusCode),
				Headers:      headers,
				ErrorHeader:  v.Name,
				ServerBody:   serverBodyData,
				ClientBody:   clientBodyData,
				ResultInit:   init,
				MustValidate: mustValidate,
			}
		}

		ref := svc.Scope.GoFullTypeRef(v.ErrorExpr.AttributeExpr, svc.PkgName)
		data[ref] = append(data[ref], &ErrorData{
			Name:     v.Name,
			Response: responseData,
			Ref:      ref,
		})
	}
	keys := make([]string, len(data))
	i := 0
	for k := range data {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	var vals []*ErrorGroupData
	for _, k := range keys {
		es := data[k]
		for _, e := range es {
			found := false
			for _, eg := range vals {
				if eg.StatusCode == e.Response.StatusCode {
					eg.Errors = append(eg.Errors, e)
					found = true
					break
				}
			}
			if !found {
				vals = append(vals,
					&ErrorGroupData{
						StatusCode: e.Response.StatusCode,
						Errors:     []*ErrorData{e},
					})
			}
		}
	}
	return vals
}

func buildStreamData(ed *EndpointData, e *expr.HTTPEndpointExpr, sd *ServiceData) {
	if !e.MethodExpr.IsStreaming() {
		return
	}
	var (
		svrSendTypeName string
		svrSendTypeRef  string
		svrRecvTypeName string
		svrRecvTypeRef  string
		svrSendDesc     string
		svrRecvDesc     string
		svrPayload      *TypeData
		cliSendDesc     string
		cliRecvDesc     string
		cliPayload      *TypeData

		md     = ed.Method
		svc    = sd.Service
		svcctx = serviceContext(sd.Service.PkgName, sd.Service.Scope)
	)
	{
		svrSendTypeName = ed.Result.Name
		svrSendTypeRef = ed.Result.Ref
		svrSendDesc = fmt.Sprintf("%s streams instances of %q to the %q endpoint websocket connection.", md.ServerStream.SendName, svrSendTypeName, md.Name)
		cliRecvDesc = fmt.Sprintf("%s reads instances of %q from the %q endpoint websocket connection.", md.ClientStream.RecvName, svrSendTypeName, md.Name)
		if e.MethodExpr.Stream == expr.ClientStreamKind || e.MethodExpr.Stream == expr.BidirectionalStreamKind {
			svrRecvTypeName = sd.Scope.GoFullTypeName(e.MethodExpr.StreamingPayload, svc.PkgName)
			svrRecvTypeRef = sd.Scope.GoFullTypeRef(e.MethodExpr.StreamingPayload, svc.PkgName)
			svrPayload = buildRequestBodyType(e.StreamingBody, e.MethodExpr.StreamingPayload, e, true, sd)
			if needInit(e.MethodExpr.StreamingPayload.Type) {
				makeHTTPType(e.StreamingBody)
				body := e.StreamingBody.Type
				// generate constructor function to transform request body,
				// into the method streaming payload type
				var (
					name       string
					desc       string
					isObject   bool
					serverArgs []*InitArgData
					serverCode string
					err        error
				)
				{
					n := codegen.Goify(e.MethodExpr.Name, true)
					p := codegen.Goify(svrPayload.Name, true)
					// Raw payload object has type name prefixed with endpoint name. No need to
					// prefix the type name again.
					if strings.HasPrefix(p, n) {
						name = fmt.Sprintf("New%s", p)
					} else {
						name = fmt.Sprintf("New%s%s", n, p)
					}
					desc = fmt.Sprintf("%s builds a %s service %s endpoint payload.", name, svc.Name, e.MethodExpr.Name)
					isObject = expr.IsObject(e.MethodExpr.StreamingPayload.Type)
					if body != expr.Empty {
						var (
							ref    string
							svcode string
						)
						{
							ref = "body"
							if expr.IsObject(body) {
								ref = "&body"
							}
							if ut, ok := body.(expr.UserType); ok {
								if val := ut.Attribute().Validation; val != nil {
									httpctx := httpContext("", sd.Scope, true, true)
									svcode = codegen.RecursiveValidationCode(ut.Attribute(), httpctx, true, "body")
								}
							}
						}
						serverArgs = []*InitArgData{{
							Name:     "body",
							Ref:      ref,
							TypeName: sd.Scope.GoTypeName(e.StreamingBody),
							TypeRef:  sd.Scope.GoTypeRef(e.StreamingBody),
							Type:     e.StreamingBody.Type,
							Required: true,
							Example:  e.Body.Example(expr.Root.API.Random()),
							Validate: svcode,
						}}
					}
					if body != expr.Empty {
						var helpers []*codegen.TransformFunctionData
						httpctx := httpContext("", sd.Scope, true, true)
						serverCode, helpers, err = marshal(e.StreamingBody, e.MethodExpr.StreamingPayload, "body", "v", httpctx, svcctx)
						if err == nil {
							sd.ServerTransformHelpers = codegen.AppendHelpers(sd.ServerTransformHelpers, helpers)
						}
					}
					if err != nil {
						fmt.Println(err.Error()) // TBD validate DSL so errors are not possible
					}
				}
				svrPayload.Init = &InitData{
					Name:           name,
					Description:    desc,
					ServerArgs:     serverArgs,
					ReturnTypeName: svc.Scope.GoFullTypeName(e.MethodExpr.StreamingPayload, svc.PkgName),
					ReturnTypeRef:  svc.Scope.GoFullTypeRef(e.MethodExpr.StreamingPayload, svc.PkgName),
					ReturnIsStruct: isObject,
					ReturnTypePkg:  svc.PkgName,
					ServerCode:     serverCode,
				}
			}
			cliPayload = buildRequestBodyType(e.StreamingBody, e.MethodExpr.StreamingPayload, e, false, sd)
			if cliPayload != nil {
				sd.ClientTypeNames[cliPayload.Name] = false
				sd.ServerTypeNames[cliPayload.Name] = false
			}
			if e.MethodExpr.Stream == expr.ClientStreamKind {
				svrSendDesc = fmt.Sprintf("%s streams instances of %q to the %q endpoint websocket connection and closes the connection.", md.ServerStream.SendName, svrSendTypeName, md.Name)
				cliRecvDesc = fmt.Sprintf("%s stops sending messages to the %q endpoint websocket connection and reads instances of %q from the connection.", md.ClientStream.RecvName, md.Name, svrSendTypeName)
			}
			svrRecvDesc = fmt.Sprintf("%s reads instances of %q from the %q endpoint websocket connection.", md.ServerStream.RecvName, svrRecvTypeName, md.Name)
			cliSendDesc = fmt.Sprintf("%s streams instances of %q to the %q endpoint websocket connection.", md.ClientStream.SendName, svrRecvTypeName, md.Name)
		}
	}
	ed.ServerStream = &StreamData{
		VarName:      md.ServerStream.VarName,
		Interface:    fmt.Sprintf("%s.%s", svc.PkgName, md.ServerStream.Interface),
		Endpoint:     ed,
		Payload:      svrPayload,
		Response:     ed.Result.Responses[0],
		PkgName:      svc.PkgName,
		Type:         "server",
		Kind:         md.ServerStream.Kind,
		SendName:     md.ServerStream.SendName,
		SendDesc:     svrSendDesc,
		SendTypeName: svrSendTypeName,
		SendTypeRef:  svrSendTypeRef,
		RecvName:     md.ServerStream.RecvName,
		RecvDesc:     svrRecvDesc,
		RecvTypeName: svrRecvTypeName,
		RecvTypeRef:  svrRecvTypeRef,
		MustClose:    md.ServerStream.MustClose,
	}
	ed.ClientStream = &StreamData{
		VarName:      md.ClientStream.VarName,
		Interface:    fmt.Sprintf("%s.%s", svc.PkgName, md.ClientStream.Interface),
		Endpoint:     ed,
		Payload:      cliPayload,
		Response:     ed.Result.Responses[0],
		PkgName:      svc.PkgName,
		Type:         "client",
		Kind:         md.ClientStream.Kind,
		SendName:     md.ClientStream.SendName,
		SendDesc:     cliSendDesc,
		SendTypeName: svrRecvTypeName,
		SendTypeRef:  svrRecvTypeRef,
		RecvName:     md.ClientStream.RecvName,
		RecvDesc:     cliRecvDesc,
		RecvTypeName: svrSendTypeName,
		RecvTypeRef:  svrSendTypeRef,
		MustClose:    md.ClientStream.MustClose,
	}
}

// buildRequestBodyType builds the TypeData for a request body. The data makes
// it possible to generate a function on the client side that creates the body
// from the service method payload.
//
// body is the HTTP request body
//
// att is the payload attribute
//
// e is the HTTP endpoint expression
//
// svr is true if the function is generated for server side code.
//
// sd is the service data
//
func buildRequestBodyType(body, att *expr.AttributeExpr, e *expr.HTTPEndpointExpr, svr bool, sd *ServiceData) *TypeData {
	if body.Type == expr.Empty {
		return nil
	}
	var (
		name        string
		varname     string
		desc        string
		def         string
		ref         string
		validateDef string
		validateRef string

		svc     = sd.Service
		httpctx = httpContext("", sd.Scope, true, svr)
		svcctx  = serviceContext(sd.Service.PkgName, sd.Service.Scope)
	)
	{
		name = body.Type.Name()
		ref = sd.Scope.GoTypeRef(body)
		if ut, ok := body.Type.(expr.UserType); ok {
			varname = codegen.Goify(ut.Name(), true)
			def = goTypeDef(sd.Scope, ut.Attribute(), svr, !svr)
			desc = fmt.Sprintf("%s is the type of the %q service %q endpoint HTTP request body.",
				varname, svc.Name, e.Name())
			if svr {
				// generate validation code for unmarshaled type (server-side).
				validateDef = codegen.RecursiveValidationCode(ut.Attribute(), httpctx, true, "body")
				if validateDef != "" {
					validateRef = fmt.Sprintf("err = Validate%s(&body)", varname)
				}
			}
		} else {
			varname = sd.Scope.GoTypeRef(body)
			validateRef = codegen.RecursiveValidationCode(body, httpctx, true, "body")
			desc = body.Description
		}
	}
	var init *InitData
	{
		if !svr && att.Type != expr.Empty && needInit(body.Type) {
			var (
				name    string
				desc    string
				code    string
				origin  string
				err     error
				helpers []*codegen.TransformFunctionData

				sourceVar = "p"
				svc       = sd.Service
			)
			{
				name = fmt.Sprintf("New%s", codegen.Goify(sd.Scope.GoTypeName(body), true))
				desc = fmt.Sprintf("%s builds the HTTP request body from the payload of the %q endpoint of the %q service.",
					name, e.Name(), svc.Name)
				src := sourceVar
				srcAtt := att
				// If design uses Body("name") syntax then need to use payload attribute
				// to transform.
				if o, ok := body.Meta["origin:attribute"]; ok {
					srcObj := expr.AsObject(att.Type)
					origin = o[0]
					srcAtt = srcObj.Attribute(origin)
					src += "." + codegen.Goify(origin, true)
				}
				code, helpers, err = marshal(srcAtt, body, src, "body", svcctx, httpctx)
				if err != nil {
					fmt.Println(err.Error()) // TBD validate DSL so errors are not possible
				}
				sd.ClientTransformHelpers = codegen.AppendHelpers(sd.ClientTransformHelpers, helpers)
			}
			arg := InitArgData{
				Name:     sourceVar,
				Ref:      sourceVar,
				TypeRef:  svc.Scope.GoFullTypeRef(att, svc.PkgName),
				Type:     att.Type,
				Validate: validateDef,
				Example:  att.Example(expr.Root.API.Random()),
			}
			init = &InitData{
				Name:                name,
				Description:         desc,
				ReturnTypeRef:       sd.Scope.GoTypeRef(body),
				ReturnTypeAttribute: codegen.Goify(origin, true),
				ClientCode:          code,
				ClientArgs:          []*InitArgData{&arg},
			}
		}
	}
	return &TypeData{
		Name:        name,
		VarName:     varname,
		Description: desc,
		Def:         def,
		Ref:         ref,
		Init:        init,
		ValidateDef: validateDef,
		ValidateRef: validateRef,
		Example:     body.Example(expr.Root.API.Random()),
	}
}

// buildResponseBodyType builds the TypeData for a response body. The data
// makes it possible to generate a function that creates the server response
// body from the service method result/projected result or error.
//
// body is the response (success or error) HTTP body.
//
// att is the result/projected attribute.
//
// svr is true if the function is generated for server side code
//
// view is the view name to add as a suffix to the type name.
//
func buildResponseBodyType(body, att *expr.AttributeExpr, e *expr.HTTPEndpointExpr, svr bool, view *string, sd *ServiceData) *TypeData {
	if body.Type == expr.Empty {
		return nil
	}
	var (
		name        string
		varname     string
		desc        string
		def         string
		ref         string
		validateDef string
		validateRef string
		viewName    string
		mustInit    bool

		svc     = sd.Service
		httpctx = httpContext("", sd.Scope, false, svr)
		svcctx  = serviceContext(sd.Service.PkgName, sd.Service.Scope)
	)
	{
		// For server code, we project the response body type if the type is a result
		// type and generate a type for each view in the result type. This makes it
		// possible to return only the attributes in the view in the server response.
		if svr && view != nil && *view != "" {
			viewName = *view
			body = expr.DupAtt(body)
			if rt, ok := body.Type.(*expr.ResultTypeExpr); ok {
				var err error
				rt, err = expr.Project(rt, *view)
				if err != nil {
					panic(err)
				}
				body.Type = rt
				sd.ServerTypeNames[rt.Name()] = false
			}
		}

		name = body.Type.Name()
		ref = sd.Scope.GoTypeRef(body)
		mustInit = att.Type != expr.Empty && needInit(body.Type)

		if ut, ok := body.Type.(expr.UserType); ok {
			// response body is a user type.
			varname = codegen.Goify(ut.Name(), true)
			def = goTypeDef(sd.Scope, ut.Attribute(), !svr, svr)
			desc = fmt.Sprintf("%s is the type of the %q service %q endpoint HTTP response body.",
				varname, svc.Name, e.Name())
			if !svr && view == nil {
				// generate validation code for unmarshaled type (client-side).
				validateDef = codegen.RecursiveValidationCode(body, httpctx, true, "body")
				if validateDef != "" {
					target := "&body"
					if expr.IsArray(ut) {
						// result type collection
						target = "body"
					}
					validateRef = fmt.Sprintf("err = Validate%s(%s)", varname, target)
				}
			}
		} else if !expr.IsPrimitive(body.Type) && mustInit {
			// response body is an array or map type.
			name = codegen.Goify(e.Name(), true) + "ResponseBody"
			varname = name
			desc = fmt.Sprintf("%s is the type of the %q service %q endpoint HTTP response body.",
				varname, svc.Name, e.Name())
			def = goTypeDef(sd.Scope, body, !svr, svr)
			validateRef = codegen.RecursiveValidationCode(body, httpctx, true, "body")
		} else {
			// response body is a primitive type. They are used as non-pointers when
			// encoding/decoding responses.
			httpctx = httpContext("", sd.Scope, false, true)
			validateRef = codegen.RecursiveValidationCode(body, httpctx, true, "body")
			varname = sd.Scope.GoTypeRef(body)
			desc = body.Description
		}
	}
	if svr {
		sd.ServerTypeNames[name] = false
		// We collect the server body types need to generate a response body type
		// here because the response body type would be different from the actual
		// type in the HTTPResponseExpr since we projected the body type above.
		// For client side, we don't have to generate a separate body type per
		// view. Hence the client types are collected in "analyze" function.
		collectUserTypes(body.Type, func(ut expr.UserType) {
			if d := attributeTypeData(ut, false, false, true, sd); d != nil {
				sd.ServerBodyAttributeTypes = append(sd.ServerBodyAttributeTypes, d)
			}
		})

	}

	var init *InitData
	{
		if svr && mustInit {
			var (
				name    string
				desc    string
				rtref   string
				code    string
				origin  string
				err     error
				helpers []*codegen.TransformFunctionData

				sourceVar = "res"
				svc       = sd.Service
			)
			{
				var rtname string
				if _, ok := body.Type.(expr.UserType); !ok && !expr.IsPrimitive(body.Type) {
					rtname = codegen.Goify(e.Name(), true) + "ResponseBody"
					rtref = rtname
				} else {
					rtname = codegen.Goify(sd.Scope.GoTypeName(body), true)
					rtref = sd.Scope.GoTypeRef(body)
				}
				name = fmt.Sprintf("New%s", rtname)
				desc = fmt.Sprintf("%s builds the HTTP response body from the result of the %q endpoint of the %q service.",
					name, e.Name(), svc.Name)
				if view != nil {
					svcctx = viewContext(sd.Service.ViewsPkg, sd.Service.ViewScope)
				}
				src := sourceVar
				srcAtt := att
				// If design uses Body("name") syntax then need to use result attribute
				// to transform.
				if o, ok := body.Meta["origin:attribute"]; ok {
					srcObj := expr.AsObject(att.Type)
					origin = o[0]
					srcAtt = srcObj.Attribute(origin)
					src += "." + codegen.Goify(origin, true)
				}
				code, helpers, err = marshal(srcAtt, body, src, "body", svcctx, httpctx)
				if err != nil {
					fmt.Println(err.Error()) // TBD validate DSL so errors are not possible
				}
				sd.ServerTransformHelpers = codegen.AppendHelpers(sd.ServerTransformHelpers, helpers)
			}
			ref := sourceVar
			if view != nil {
				ref += ".Projected"
			}
			tref := svc.Scope.GoFullTypeRef(att, svc.PkgName)
			if view != nil {
				tref = svc.ViewScope.GoFullTypeRef(att, svc.ViewsPkg)
			}
			arg := InitArgData{
				Name:     sourceVar,
				Ref:      ref,
				TypeRef:  tref,
				Type:     att.Type,
				Validate: validateDef,
				Example:  att.Example(expr.Root.API.Random()),
			}
			init = &InitData{
				Name:                name,
				Description:         desc,
				ReturnTypeRef:       rtref,
				ReturnTypeAttribute: codegen.Goify(origin, true),
				ServerCode:          code,
				ServerArgs:          []*InitArgData{&arg},
			}
		}
	}
	return &TypeData{
		Name:        name,
		VarName:     varname,
		Description: desc,
		Def:         def,
		Ref:         ref,
		Init:        init,
		ValidateDef: validateDef,
		ValidateRef: validateRef,
		Example:     body.Example(expr.Root.API.Random()),
		View:        viewName,
	}
}

func extractPathParams(a *expr.MappedAttributeExpr, service *expr.AttributeExpr, scope *codegen.NameScope) []*ParamData {
	var params []*ParamData
	codegen.WalkMappedAttr(a, func(name, elem string, _ bool, c *expr.AttributeExpr) error {
		makeHTTPType(c)
		var (
			varn = scope.Name(codegen.Goify(name, false))
			arr  = expr.AsArray(c.Type)
			ctx  = serviceContext("", scope)
			ft   = service.Type

			fptr bool
		)
		fieldName := codegen.Goify(name, true)
		if !expr.IsObject(service.Type) {
			fieldName = ""
		} else {
			fptr = service.IsPrimitivePointer(name, true)
			ft = service.Find(name).Type
		}
		params = append(params, &ParamData{
			Name:           elem,
			AttributeName:  name,
			Description:    c.Description,
			FieldName:      fieldName,
			FieldPointer:   fptr,
			FieldType:      ft,
			VarName:        varn,
			Required:       true,
			Type:           c.Type,
			TypeName:       scope.GoTypeName(c),
			TypeRef:        scope.GoTypeRef(c),
			Pointer:        false,
			Slice:          arr != nil,
			StringSlice:    arr != nil && arr.ElemType.Type.Kind() == expr.StringKind,
			Map:            false,
			MapStringSlice: false,
			Validate:       codegen.RecursiveValidationCode(c, ctx, true, varn),
			DefaultValue:   c.DefaultValue,
			Example:        c.Example(expr.Root.API.Random()),
		})
		return nil
	})

	return params
}

func extractQueryParams(a *expr.MappedAttributeExpr, service *expr.AttributeExpr, scope *codegen.NameScope) []*ParamData {
	var params []*ParamData
	codegen.WalkMappedAttr(a, func(name, elem string, required bool, c *expr.AttributeExpr) error {
		makeHTTPType(c)
		var (
			varn    = scope.Name(codegen.Goify(name, false))
			arr     = expr.AsArray(c.Type)
			mp      = expr.AsMap(c.Type)
			typeRef = scope.GoTypeRef(c)
			ctx     = serviceContext("", scope)
			ft      = service.Type

			pointer bool
			fptr    bool
		)
		if pointer = a.IsPrimitivePointer(name, true); pointer {
			typeRef = "*" + typeRef
		}
		fieldName := codegen.Goify(name, true)
		if !expr.IsObject(service.Type) {
			fieldName = ""
		} else {
			fptr = service.IsPrimitivePointer(name, true)
			ft = service.Find(name).Type
		}
		params = append(params, &ParamData{
			Name:          elem,
			AttributeName: name,
			Description:   c.Description,
			FieldName:     fieldName,
			FieldPointer:  fptr,
			FieldType:     ft,
			VarName:       varn,
			Required:      required,
			Type:          c.Type,
			TypeName:      scope.GoTypeName(c),
			TypeRef:       typeRef,
			Pointer:       pointer,
			Slice:         arr != nil,
			StringSlice:   arr != nil && arr.ElemType.Type.Kind() == expr.StringKind,
			Map:           mp != nil,
			MapStringSlice: mp != nil &&
				mp.KeyType.Type.Kind() == expr.StringKind &&
				mp.ElemType.Type.Kind() == expr.ArrayKind &&
				expr.AsArray(mp.ElemType.Type).ElemType.Type.Kind() == expr.StringKind,
			Validate:     codegen.RecursiveValidationCode(c, ctx, required, varn),
			DefaultValue: c.DefaultValue,
			Example:      c.Example(expr.Root.API.Random()),
		})
		return nil
	})

	return params
}

func extractHeaders(a *expr.MappedAttributeExpr, svcAtt *expr.AttributeExpr, svcCtx *codegen.AttributeContext, scope *codegen.NameScope) []*HeaderData {
	var headers []*HeaderData
	codegen.WalkMappedAttr(a, func(name, elem string, required bool, _ *expr.AttributeExpr) error {
		var hattr *expr.AttributeExpr
		{
			if hattr = svcAtt.Find(name); hattr == nil {
				hattr = svcAtt
			}
			makeHTTPType(hattr)
		}
		var (
			varn    = scope.Name(codegen.Goify(name, false))
			arr     = expr.AsArray(hattr.Type)
			typeRef = scope.GoTypeRef(hattr)
			ft      = svcAtt.Type

			fieldName string
			pointer   bool
			fptr      bool
		)
		{
			pointer = a.IsPrimitivePointer(name, true)
			if expr.IsObject(svcAtt.Type) {
				fieldName = codegen.Goify(name, true)
				fptr = svcCtx.IsPrimitivePointer(name, svcAtt)
				ft = svcAtt.Find(name).Type
			}
			if pointer {
				typeRef = "*" + typeRef
			}
		}
		headers = append(headers, &HeaderData{
			Name:          elem,
			AttributeName: name,
			Description:   hattr.Description,
			CanonicalName: http.CanonicalHeaderKey(elem),
			FieldName:     fieldName,
			FieldPointer:  fptr,
			FieldType:     ft,
			VarName:       varn,
			TypeName:      scope.GoTypeName(hattr),
			TypeRef:       typeRef,
			Required:      required,
			Pointer:       pointer,
			Slice:         arr != nil,
			StringSlice:   arr != nil && arr.ElemType.Type.Kind() == expr.StringKind,
			Type:          hattr.Type,
			Validate:      codegen.RecursiveValidationCode(hattr, svcCtx, required, varn),
			DefaultValue:  hattr.DefaultValue,
			Example:       hattr.Example(expr.Root.API.Random()),
		})
		return nil
	})
	return headers
}

// collectUserTypes traverses the given data type recursively and calls back the
// given function for each attribute using a user type.
func collectUserTypes(dt expr.DataType, cb func(expr.UserType), seen ...map[string]struct{}) {
	if dt == expr.Empty {
		return
	}
	var s map[string]struct{}
	if len(seen) > 0 {
		s = seen[0]
	} else {
		s = make(map[string]struct{})
	}
	switch actual := dt.(type) {
	case *expr.Object:
		for _, nat := range *actual {
			collectUserTypes(nat.Attribute.Type, cb, seen...)
		}
	case *expr.Array:
		collectUserTypes(actual.ElemType.Type, cb, seen...)
	case *expr.Map:
		collectUserTypes(actual.KeyType.Type, cb, seen...)
		collectUserTypes(actual.ElemType.Type, cb, seen...)
	case expr.UserType:
		if _, ok := s[actual.ID()]; ok {
			return
		}
		s[actual.ID()] = struct{}{}
		cb(actual)
		collectUserTypes(actual.Attribute().Type, cb, s)
	}
}

func attributeTypeData(ut expr.UserType, req, ptr, server bool, rd *ServiceData) *TypeData {
	if ut == expr.Empty {
		return nil
	}
	seen := rd.ServerTypeNames
	if !server {
		seen = rd.ClientTypeNames
	}
	if _, ok := seen[ut.Name()]; ok {
		return nil
	}
	seen[ut.Name()] = false

	var (
		name        string
		desc        string
		validate    string
		validateRef string

		att  = &expr.AttributeExpr{Type: ut}
		hctx = httpContext("", rd.Scope, req, server)
	)
	{
		name = rd.Scope.GoTypeName(att)
		ctx := "request"
		if !req {
			ctx = "response"
		}
		desc = name + " is used to define fields on " + ctx + " body types."
		if req || !req && !server {
			// generate validations for responses client-side and for
			// requests server-side and CLI
			validate = codegen.RecursiveValidationCode(ut.Attribute(), hctx, true, "body")
		}
		if validate != "" {
			validateRef = fmt.Sprintf("err = Validate%s(v)", name)
		}
	}
	return &TypeData{
		Name:        ut.Name(),
		VarName:     name,
		Description: desc,
		Def:         goTypeDef(rd.Scope, ut.Attribute(), ptr, hctx.UseDefault),
		Ref:         rd.Scope.GoTypeRef(att),
		ValidateDef: validate,
		ValidateRef: validateRef,
		Example:     att.Example(expr.Root.API.Random()),
	}
}

// httpContext returns a context for attributes of types used to marshal and
// unmarshal HTTP requests and responses.
//
// pkg is the package name where the body type exists
//
// scope is the named scope
//
// request if true indicates that the type is a request type, else response
// type
//
// svr if true indicates that the type is a server type, else client type
//
func httpContext(pkg string, scope *codegen.NameScope, request, svr bool) *codegen.AttributeContext {
	marshal := !request && svr || request && !svr
	return codegen.NewAttributeContext(!marshal, false, marshal, pkg, scope)
}

// serviceContext returns an attribute context for service types.
func serviceContext(pkg string, scope *codegen.NameScope) *codegen.AttributeContext {
	return codegen.NewAttributeContext(false, false, true, pkg, scope)
}

// viewContext returns an attribute context for projected types.
func viewContext(pkg string, scope *codegen.NameScope) *codegen.AttributeContext {
	return codegen.NewAttributeContext(true, false, true, pkg, scope)
}

// unmarshal initializes a data structure defined by target type from a data
// structure defined by source type. The attributes in the source data
// structure are pointers and the attributes in the target data structure that
// have default values are non-pointers. Fields in target type are initialized
// with their default values (if any).
//
// source, target are the attributes used in the transformation
//
// sourceVar, targetVar are the variable names for source and target used in
// the transformation code
//
// sourceCtx, targetCtx are the source and target attribute contexts
//
func unmarshal(source, target *expr.AttributeExpr, sourceVar, targetVar string, sourceCtx, targetCtx *codegen.AttributeContext) (string, []*codegen.TransformFunctionData, error) {
	return codegen.GoTransform(source, target, sourceVar, targetVar, sourceCtx, targetCtx, "unmarshal")
}

// marshal initializes a data structure defined by target type from a data
// structure defined by source type. The fields in the source and target
// data structure use non-pointers for attributes with default values.
//
// source, target are the attributes used in the transformation
//
// sourceVar, targetVar are the variable names for source and target used in
// the transformation code
//
// sourceCtx, targetCtx are the source and target attribute contexts
//
func marshal(source, target *expr.AttributeExpr, sourceVar, targetVar string, sourceCtx, targetCtx *codegen.AttributeContext) (string, []*codegen.TransformFunctionData, error) {
	return codegen.GoTransform(source, target, sourceVar, targetVar, sourceCtx, targetCtx, "marshal")
}

// needConversion returns true if the type needs to be converted from a string.
func needConversion(dt expr.DataType) bool {
	if dt == expr.Empty {
		return false
	}
	switch actual := dt.(type) {
	case expr.Primitive:
		if actual.Kind() == expr.StringKind ||
			actual.Kind() == expr.AnyKind ||
			actual.Kind() == expr.BytesKind {
			return false
		}
		return true
	case *expr.Array:
		return needConversion(actual.ElemType.Type)
	case *expr.Map:
		return needConversion(actual.KeyType.Type) ||
			needConversion(actual.ElemType.Type)
	default:
		return true
	}
}

// needInit returns true if and only if the given type is or makes use of user
// types.
func needInit(dt expr.DataType) bool {
	if dt == expr.Empty {
		return false
	}
	switch actual := dt.(type) {
	case expr.Primitive:
		return false
	case *expr.Array:
		return needInit(actual.ElemType.Type)
	case *expr.Map:
		return needInit(actual.KeyType.Type) ||
			needInit(actual.ElemType.Type)
	case *expr.Object:
		for _, nat := range *actual {
			if needInit(nat.Attribute.Type) {
				return true
			}
		}
		return false
	case expr.UserType:
		return true
	default:
		panic(fmt.Sprintf("unknown data type %T", actual)) // bug
	}
}

// upgradeParams returns the data required to render the websocket_upgrade
// template.
func upgradeParams(e *EndpointData, fn string) map[string]interface{} {
	return map[string]interface{}{
		"ViewedResult": e.Method.ViewedResult,
		"Function":     fn,
	}
}

// needStream returns true if at least one method in the defined services
// uses stream for sending payload/result.
func needStream(data []*ServiceData) bool {
	for _, svc := range data {
		if streamingEndpointExists(svc) {
			return true
		}
	}
	return false
}

// streamingEndpointExists returns true if at least one of the endpoints in
// the service defines a streaming payload or result.
func streamingEndpointExists(sd *ServiceData) bool {
	for _, e := range sd.Endpoints {
		if isStreamingEndpoint(e) {
			return true
		}
	}
	return false
}

// isStreamingEndpoint returns true if the endpoint defines a streaming payload
// or result.
func isStreamingEndpoint(ed *EndpointData) bool {
	return ed.ServerStream != nil || ed.ClientStream != nil
}

const (
	// pathInitT is the template used to render the code of path constructors.
	pathInitT = `
{{- if .Args }}
	{{- range $i, $arg := .Args }}
		{{- $typ := (index $.PathParams $i).Attribute.Type }}
		{{- if eq $typ.Name "array" }}
	{{ .Name }}Slice := make([]string, len({{ .Name }}))
	for i, v := range {{ .Name }} {
		{{ .Name }}Slice[i] = {{ template "slice_conversion" $typ.ElemType.Type.Name }}
	}
		{{- end }}
	{{- end }}
	return fmt.Sprintf("{{ .PathFormat }}", {{ range $i, $arg := .Args }}
	{{- if eq (index $.PathParams $i).Attribute.Type.Name "array" }}strings.Join({{ .Name }}Slice, ", ")
	{{- else }}{{ .Name }}
	{{- end }}, {{ end }})
{{- else }}
	return "{{ .PathFormat }}"
{{- end }}

{{- define "slice_conversion" }}
	{{- if eq . "string" }} url.QueryEscape(v)
	{{- else if eq . "int" "int32" }} strconv.FormatInt(int64(v), 10)
	{{- else if eq . "int64" }} strconv.FormatInt(v, 10)
	{{- else if eq . "uint" "uint32" }} strconv.FormatUint(uint64(v), 10)
	{{- else if eq . "uint64" }} strconv.FormatUint(v, 10)
	{{- else if eq . "float32" }} strconv.FormatFloat(float64(v), 'f', -1, 32)
	{{- else if eq . "float64" }} strconv.FormatFloat(v, 'f', -1, 64)
	{{- else if eq . "boolean" }} strconv.FormatBool(v)
	{{- else if eq . "bytes" }} url.QueryEscape(string(v))
	{{- else }} url.QueryEscape(fmt.Sprintf("%v", v))
	{{- end }}
{{- end }}`

	// requestInitT is the template used to render the code of HTTP
	// request constructors.
	requestInitT = `
{{- if .Args }}
	var (
	{{- range .Args }}
	{{ .Name }} {{ .TypeRef }}
	{{- end }}
	)
{{- end }}
{{- if and .PayloadRef .Args }}
	{
		p, ok := v.({{ .PayloadRef }})
		if !ok {
			return nil, goahttp.ErrInvalidType("{{ .ServiceName }}", "{{ .EndpointName }}", "{{ .PayloadRef }}", v)
		}
	{{- range .Args }}
		{{- if .Pointer }}
		if p{{ if $.HasFields }}.{{ .FieldName }}{{ end }} != nil {
		{{- end }}
			{{ .Name }} = {{ if .Pointer }}*{{ end }}p{{ if $.HasFields }}.{{ .FieldName }}{{ end }}
		{{- if .Pointer }}
		}
		{{- end }}
	{{- end }}
	}
{{- end }}
	{{- if .IsStreaming }}
		scheme := c.scheme
		switch c.scheme {
		case "http":
			scheme = "ws"
		case "https":
			scheme = "wss"
		}
	{{- end }}
	u := &url.URL{Scheme: {{ if .IsStreaming }}scheme{{ else }}c.scheme{{ end }}, Host: c.host, Path: {{ .PathInit.Name }}({{ range .Args }}{{ .Ref }}, {{ end }})}
	req, err := http.NewRequest("{{ .Verb }}", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("{{ .ServiceName }}", "{{ .EndpointName }}", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil`

	// streamStructTypeT renders the server and client struct types that
	// implements the client and server stream interfaces. The data to render
	// input: StreamData
	streamStructTypeT = `{{ printf "%s implements the %s interface." .VarName .Interface | comment }}
type {{ .VarName }} struct {
{{- if eq .Type "server" }}
	once sync.Once
	{{ comment "upgrader is the websocket connection upgrader." }}
	upgrader goahttp.Upgrader
	{{ comment "connConfigFn is the websocket connection configurer." }}
	connConfigFn goahttp.ConnConfigureFunc
	{{ comment "cancel is the context cancellation function which cancels the request context when invoked." }}
	cancel context.CancelFunc
	{{ comment "w is the HTTP response writer used in upgrading the connection." }}
	w http.ResponseWriter
	{{ comment "r is the HTTP request." }}
	r *http.Request
{{- end }}
	{{ comment "conn is the underlying websocket connection." }}
	conn *websocket.Conn
	{{- if .Endpoint.Method.ViewedResult }}
		{{- if not .Endpoint.Method.ViewedResult.ViewName }}
	{{ printf "view is the view to render %s result type before sending to the websocket connection." .SendTypeName | comment }}
	view string
		{{- end }}
	{{- end }}
}
`

	// streamConnConfigurerStructT generates the struct type that holds the
	// websocket connection configurers for all the streaming endpoints in the
	// service.
	// input: ServiceData
	streamConnConfigurerStructT = `{{ printf "ConnConfigurer holds the websocket connection configurer functions for the streaming endpoints in %q service." .Service.Name | comment }}
type ConnConfigurer struct {
{{- range .Endpoints }}
	{{- if isStreamingEndpoint . }}
		{{ .Method.VarName }}Fn goahttp.ConnConfigureFunc
	{{- end }}
{{- end }}
}
`

	// streamConnConfigurerStructInitT generates the constructor function to
	// initialize the websocket connection configurer struct.
	// input: ServiceData
	streamConnConfigurerStructInitT = `{{ printf "NewConnConfigurer initializes the websocket connection configurer function with fn for all the streaming endpoints in %q service." .Service.Name | comment }}
func NewConnConfigurer(fn goahttp.ConnConfigureFunc) *ConnConfigurer {
	return &ConnConfigurer{
{{- range .Endpoints }}
	{{- if isStreamingEndpoint . }}
		{{ .Method.VarName}}Fn: fn,
	{{- end }}
{{- end }}
	}
}
`

	// streamSendT renders the function implementing the Send method in
	// stream interface.
	// input: StreamData
	streamSendT = `{{ comment .SendDesc }}
func (s *{{ .VarName }}) {{ .SendName }}(v {{ .SendTypeRef }}) error {
{{- if eq .Type "server" }}
	{{- if eq .SendName "Send" }}
		var err error
		{{- template "websocket_upgrade" (upgradeParams .Endpoint .SendName) }}
	{{- else }} {{/* SendAndClose */}}
		defer s.conn.Close()
	{{- end }}
	{{- if .Endpoint.Method.ViewedResult }}
		{{- if .Endpoint.Method.ViewedResult.ViewName }}
			res := {{ .PkgName }}.{{ .Endpoint.Method.ViewedResult.Init.Name }}(v, {{ printf "%q" .Endpoint.Method.ViewedResult.ViewName }})
		{{- else }}
			res := {{ .PkgName }}.{{ .Endpoint.Method.ViewedResult.Init.Name }}(v, s.view)
		{{- end }}
	{{- else }}
	res := v
	{{- end }}
	{{- $servBodyLen := len .Response.ServerBody }}
	{{- if gt $servBodyLen 0 }}
		{{- if (index .Response.ServerBody 0).Init }}
			{{- if .Endpoint.Method.ViewedResult }}
				{{- if .Endpoint.Method.ViewedResult.ViewName }}
					{{- $vsb := (viewedServerBody $.Response.ServerBody .Endpoint.Method.ViewedResult.ViewName) }}
					body := {{ $vsb.Init.Name }}({{ range $vsb.Init.ServerArgs }}{{ .Ref }}, {{ end }})
				{{- else }}
					var body interface{}
					switch s.view {
					{{- range .Endpoint.Method.ViewedResult.Views }}
						case {{ printf "%q" .Name }}{{ if eq .Name "default" }}, ""{{ end }}:
						{{- $vsb := (viewedServerBody $.Response.ServerBody .Name) }}
							body = {{ $vsb.Init.Name }}({{ range $vsb.Init.ServerArgs }}{{ .Ref }}, {{ end }})
						{{- end }}
					}
				{{- end }}
			{{- else }}
				body := {{ (index .Response.ServerBody 0).Init.Name }}({{ range (index .Response.ServerBody 0).Init.ServerArgs }}{{ .Ref }}, {{ end }})
			{{- end }}
			return s.conn.WriteJSON(body)
		{{- else }}
			return s.conn.WriteJSON(res)
		{{- end }}
	{{- else }}
		return s.conn.WriteJSON(res)
	{{- end }}
{{- else }}
	{{- if .Payload.Init }}
		body := {{ .Payload.Init.Name }}(v)
		return s.conn.WriteJSON(body)
	{{- else }}
		return s.conn.WriteJSON(v)
	{{- end }}
{{- end }}
}
` + upgradeT

	// streamRecvT renders the function implementing the Recv method in
	// stream interface.
	// input: StreamData
	streamRecvT = `{{ comment .RecvDesc }}
func (s *{{ .VarName }}) {{ .RecvName }}() ({{ .RecvTypeRef }}, error) {
	var (
		rv {{ .RecvTypeRef }}
	{{- if eq .Type "server" }}
		msg *{{ .Payload.Ref }}
	{{- else }}
		body {{ .Response.ClientBody.VarName }}
	{{- end }}
		err error
	)
{{- if eq .Type "server" }}
	{{- template "websocket_upgrade" (upgradeParams .Endpoint .RecvName) }}
	if err = s.conn.ReadJSON(&msg); err != nil {
		return rv, err
	}
	if msg == nil {
		return rv, io.EOF
	}
	body := *msg
	{{- if .Payload.ValidateRef }}
		{{ .Payload.ValidateRef }}
		if err != nil {
			return rv, err
		}
	{{- end }}
	{{- if .Payload.Init }}
		return {{ .Payload.Init.Name }}(body), nil
	{{- else }}
		return body, nil
	{{- end }}
{{- else }} {{/* client side code */}}
	{{- if eq .RecvName "CloseAndRecv" }}
		defer s.conn.Close()
		{{ comment "Send a nil payload to the server implying end of message" }}
		if err = s.conn.WriteJSON(nil); err != nil {
			return rv, err
		}
	{{- end }}
	err = s.conn.ReadJSON(&body)
	if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
		{{- if not .MustClose }}
			s.conn.Close()
		{{- end }}
		return rv, io.EOF
	}
	if err != nil {
		return rv, err
	}
	{{- if and .Response.ClientBody.ValidateRef (not .Endpoint.Method.ViewedResult) }}
	{{ .Response.ClientBody.ValidateRef }}
	if err != nil {
		return rv, err
	}
	{{- end }}
	{{- if .Response.ResultInit }}
		res := {{ .Response.ResultInit.Name }}({{ range .Response.ResultInit.ClientArgs }}{{ .Ref }},{{ end }})
		{{- if .Endpoint.Method.ViewedResult }}{{ with .Endpoint.Method.ViewedResult }}
			vres := {{ if not .IsCollection }}&{{ end }}{{ .ViewsPkg }}.{{ .VarName }}{res, {{ if .ViewName }}{{ printf "%q" .ViewName }}{{ else }}s.view{{ end }} }
			if err := {{ .ViewsPkg }}.Validate{{ $.Endpoint.Method.Result }}(vres); err != nil {
				return rv, goahttp.ErrValidationError("{{ $.Endpoint.ServiceName }}", "{{ $.Endpoint.Method.Name }}", err)
			}
			return {{ $.PkgName }}.{{ .ResultInit.Name }}(vres){{ end }}, nil
		{{- else }}
			return res, nil
		{{- end }}
	{{- else }}
		return body, nil
	{{- end }}
{{- end }}
}
` + upgradeT

	// upgradeT renders the code to upgrade the HTTP connection to a gorilla
	// websocket connection.
	upgradeT = `{{- define "websocket_upgrade" }}
	{{ printf "Upgrade the HTTP connection to a websocket connection only once. Connection upgrade is done here so that authorization logic in the endpoint is executed before calling the actual service method which may call %s()." .Function | comment }}
	s.once.Do(func() {
	{{- if and .ViewedResult (eq .Function "Send") }}
		{{- if not .ViewedResult.ViewName }}
			respHdr := make(http.Header)
			respHdr.Add("goa-view", s.view)
		{{- end }}
	{{- end }}
		var conn *websocket.Conn
		{{- if eq .Function "Send" }}
			{{- if .ViewedResult }}
				{{- if not .ViewedResult.ViewName }}
					conn, err = s.upgrader.Upgrade(s.w, s.r, respHdr)
				{{- else }}
					conn, err = s.upgrader.Upgrade(s.w, s.r, nil)
				{{- end }}
			{{- else }}
				conn, err = s.upgrader.Upgrade(s.w, s.r, nil)
			{{- end }}
		{{- else }}
			conn, err = s.upgrader.Upgrade(s.w, s.r, nil)
		{{- end }}
		if err != nil {
			return
		}
		if s.connConfigFn != nil {
			conn = s.connConfigFn(conn, s.cancel)
		}
		s.conn = conn
	})
	if err != nil {
		return {{ if eq .Function "Recv" }}rv, {{ end }}err
	}
{{- end }}
`

	// streamCloseT renders the function implementing the Close method in
	// stream interface.
	// input: StreamData
	streamCloseT = `{{ printf "Close closes the %q endpoint websocket connection." .Endpoint.Method.Name | comment }}
func (s *{{ .VarName }}) Close() error {
	var err error
{{- if eq .Type "server" }}
	if s.conn == nil {
		return nil
	}
	if err = s.conn.WriteControl(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, "server closing connection"),
		time.Now().Add(time.Second),
	); err != nil {
		return err
	}
{{- else }} {{/* client side code */}}
	{{ comment "Send a nil payload to the server implying client closing connection." }}
  if err = s.conn.WriteJSON(nil); err != nil {
    return err
  }
{{- end }}
	return s.conn.Close()
}
` + upgradeT

	// streamSetViewT renders the function implementing the SetView method in
	// server stream interface.
	// input: StreamData
	streamSetViewT = `{{ printf "SetView sets the view to render the %s type before sending to the %q endpoint websocket connection." .SendTypeName .Endpoint.Method.Name | comment }}
func (s *{{ .VarName }}) SetView(view string) {
	s.view = view
}
`
)
