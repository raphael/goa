package codegen

import (
	"bytes"
	"testing"

	"goa.design/goa/v3/codegen"
	"goa.design/goa/v3/expr"
	"goa.design/goa/v3/http/codegen/testdata"
)

func TestBodyTypeDecl(t *testing.T) {
	const genpkg = "gen"

	cases := []struct {
		Name string
		DSL  func()
		Code string
	}{
		{"body-user-inner", testdata.PayloadBodyUserInnerDSL, BodyUserInnerDeclCode},
		{"body-path-user-validate", testdata.PayloadBodyPathUserValidateDSL, BodyPathUserValidateDeclCode},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			RunHTTPDSL(t, c.DSL)
			fs := clientType(genpkg, expr.Root.API.HTTP.Services[0], make(map[string]struct{}))
			section := fs.SectionTemplates[1]
			code := codegen.SectionCode(t, section)
			if code != c.Code {
				t.Errorf("invalid code, got:\n%s\ngot vs. expected:\n%s", code, codegen.Diff(t, code, c.Code))
			}
		})
	}
}

func TestBodyTypeInit(t *testing.T) {
	const genpkg = "gen"
	cases := []struct {
		Name         string
		DSL          func()
		SectionIndex int
		Code         string
	}{
		{"body-user-inner", testdata.PayloadBodyUserInnerDSL, 3, BodyUserInnerInitCode},
		{"body-path-user-validate", testdata.PayloadBodyPathUserValidateDSL, 2, BodyPathUserValidateInitCode},
		{"body-primitive-array-user-validate", testdata.PayloadBodyPrimitiveArrayUserValidateDSL, 2, BodyPrimitiveArrayUserValidateInitCode},
		{"result-body-user", testdata.ResultBodyObjectHeaderDSL, 2, ResultBodyObjectHeaderInitCode},
		{"result-explicit-body-primitive", testdata.ExplicitBodyPrimitiveResultMultipleViewsDSL, 1, ExplicitBodyPrimitiveResultMultipleViewsInitCode},
		{"result-explicit-body-user-type", testdata.ExplicitBodyUserResultMultipleViewsDSL, 2, ExplicitBodyUserResultMultipleViewsInitCode},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			RunHTTPDSL(t, c.DSL)
			fs := clientType(genpkg, expr.Root.API.HTTP.Services[0], make(map[string]struct{}))
			section := fs.SectionTemplates[c.SectionIndex]
			code := codegen.SectionCode(t, section)
			if code != c.Code {
				t.Errorf("invalid code, got:\n%s\ngot vs. expected:\n%s", code, codegen.Diff(t, code, c.Code))
			}
		})
	}
}

func TestClientTypes(t *testing.T) {
	const genpkg = "gen"
	cases := []struct {
		Name string
		DSL  func()
		Code string
	}{
		{"mixed-payload-attrs", testdata.MixedPayloadInBodyDSL, MixedPayloadInBodyClientTypesFile},
		{"multiple-methods", testdata.MultipleMethodsDSL, MultipleMethodsClientTypesFile},
		{"payload-extend-validate", testdata.PayloadExtendedValidateDSL, PayloadExtendedValidateClientTypesFile},
		{"result-type-validate", testdata.ResultTypeValidateDSL, ResultTypeValidateClientTypesFile},
		{"with-result-collection", testdata.ResultWithResultCollectionDSL, WithResultCollectionClientTypesFile},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			RunHTTPDSL(t, c.DSL)
			fs := clientType(genpkg, expr.Root.API.HTTP.Services[0], make(map[string]struct{}))
			var buf bytes.Buffer
			for _, s := range fs.SectionTemplates[1:] {
				if err := s.Write(&buf); err != nil {
					t.Fatal(err)
				}
			}
			code := codegen.FormatTestCode(t, "package foo\n"+buf.String())
			if code != c.Code {
				t.Errorf("invalid code, got:\n%s\ngot vs. expected:\n%s", code, codegen.Diff(t, code, c.Code))
			}
		})
	}
}

const BodyUserInnerDeclCode = `// MethodBodyUserInnerRequestBody is the type of the "ServiceBodyUserInner"
// service "MethodBodyUserInner" endpoint HTTP request body.
type MethodBodyUserInnerRequestBody struct {
	Inner *InnerTypeRequestBody ` + "`" + `form:"inner,omitempty" json:"inner,omitempty" xml:"inner,omitempty"` + "`" + `
}
`

const BodyPathUserValidateDeclCode = `// MethodUserBodyPathValidateRequestBody is the type of the
// "ServiceBodyPathUserValidate" service "MethodUserBodyPathValidate" endpoint
// HTTP request body.
type MethodUserBodyPathValidateRequestBody struct {
	A string ` + "`" + `form:"a" json:"a" xml:"a"` + "`" + `
}
`

const BodyPrimitiveArrayUserValidateInitCode = `// NewPayloadTypeRequestBody builds the HTTP request body from the payload of
// the "MethodBodyPrimitiveArrayUserValidate" endpoint of the
// "ServiceBodyPrimitiveArrayUserValidate" service.
func NewPayloadTypeRequestBody(p []*servicebodyprimitivearrayuservalidate.PayloadType) []*PayloadTypeRequestBody {
	body := make([]*PayloadTypeRequestBody, len(p))
	for i, val := range p {
		body[i] = marshalServicebodyprimitivearrayuservalidatePayloadTypeToPayloadTypeRequestBody(val)
	}
	return body
}
`

const BodyUserInnerInitCode = `// NewMethodBodyUserInnerRequestBody builds the HTTP request body from the
// payload of the "MethodBodyUserInner" endpoint of the "ServiceBodyUserInner"
// service.
func NewMethodBodyUserInnerRequestBody(p *servicebodyuserinner.PayloadType) *MethodBodyUserInnerRequestBody {
	body := &MethodBodyUserInnerRequestBody{}
	if p.Inner != nil {
		body.Inner = marshalServicebodyuserinnerInnerTypeToInnerTypeRequestBody(p.Inner)
	}
	return body
}
`

const BodyPathUserValidateInitCode = `// NewMethodUserBodyPathValidateRequestBody builds the HTTP request body from
// the payload of the "MethodUserBodyPathValidate" endpoint of the
// "ServiceBodyPathUserValidate" service.
func NewMethodUserBodyPathValidateRequestBody(p *servicebodypathuservalidate.PayloadType) *MethodUserBodyPathValidateRequestBody {
	body := &MethodUserBodyPathValidateRequestBody{
		A: p.A,
	}
	return body
}
`

const ResultBodyObjectHeaderInitCode = `// NewMethodBodyObjectHeaderResultOK builds a "ServiceBodyObjectHeader" service
// "MethodBodyObjectHeader" endpoint result from a HTTP "OK" response.
func NewMethodBodyObjectHeaderResultOK(body *MethodBodyObjectHeaderResponseBody, b *string) *servicebodyobjectheader.MethodBodyObjectHeaderResult {
	v := &servicebodyobjectheader.MethodBodyObjectHeaderResult{
		A: body.A,
	}
	v.B = b

	return v
}
`

const ExplicitBodyPrimitiveResultMultipleViewsInitCode = `// NewMethodExplicitBodyPrimitiveResultMultipleViewResulttypemultipleviewsOK
// builds a "ServiceExplicitBodyPrimitiveResultMultipleView" service
// "MethodExplicitBodyPrimitiveResultMultipleView" endpoint result from a HTTP
// "OK" response.
func NewMethodExplicitBodyPrimitiveResultMultipleViewResulttypemultipleviewsOK(body string, c *string) *serviceexplicitbodyprimitiveresultmultipleviewviews.ResulttypemultipleviewsView {
	v := body
	res := &serviceexplicitbodyprimitiveresultmultipleviewviews.ResulttypemultipleviewsView{
		A: &v,
	}
	res.C = c

	return res
}
`

const ExplicitBodyUserResultMultipleViewsInitCode = `// NewMethodExplicitBodyUserResultMultipleViewResulttypemultipleviewsOK builds
// a "ServiceExplicitBodyUserResultMultipleView" service
// "MethodExplicitBodyUserResultMultipleView" endpoint result from a HTTP "OK"
// response.
func NewMethodExplicitBodyUserResultMultipleViewResulttypemultipleviewsOK(body *UserType, c *string) *serviceexplicitbodyuserresultmultipleviewviews.ResulttypemultipleviewsView {
	v := &serviceexplicitbodyuserresultmultipleviewviews.UserTypeView{
		X: body.X,
		Y: body.Y,
	}
	res := &serviceexplicitbodyuserresultmultipleviewviews.ResulttypemultipleviewsView{
		A: v,
	}
	res.C = c

	return res
}
`

const MixedPayloadInBodyClientTypesFile = `// MethodARequestBody is the type of the "ServiceMixedPayloadInBody" service
// "MethodA" endpoint HTTP request body.
type MethodARequestBody struct {
	Any    interface{}          ` + "`" + `form:"any,omitempty" json:"any,omitempty" xml:"any,omitempty"` + "`" + `
	Array  []float32            ` + "`" + `form:"array" json:"array" xml:"array"` + "`" + `
	Map    map[uint]interface{} ` + "`" + `form:"map,omitempty" json:"map,omitempty" xml:"map,omitempty"` + "`" + `
	Object *BPayloadRequestBody ` + "`" + `form:"object" json:"object" xml:"object"` + "`" + `
	DupObj *BPayloadRequestBody ` + "`" + `form:"dup_obj,omitempty" json:"dup_obj,omitempty" xml:"dup_obj,omitempty"` + "`" + `
}

// BPayloadRequestBody is used to define fields on request body types.
type BPayloadRequestBody struct {
	Int   int    ` + "`" + `form:"int" json:"int" xml:"int"` + "`" + `
	Bytes []byte ` + "`" + `form:"bytes,omitempty" json:"bytes,omitempty" xml:"bytes,omitempty"` + "`" + `
}

// NewMethodARequestBody builds the HTTP request body from the payload of the
// "MethodA" endpoint of the "ServiceMixedPayloadInBody" service.
func NewMethodARequestBody(p *servicemixedpayloadinbody.APayload) *MethodARequestBody {
	body := &MethodARequestBody{
		Any: p.Any,
	}
	if p.Array != nil {
		body.Array = make([]float32, len(p.Array))
		for i, val := range p.Array {
			body.Array[i] = val
		}
	}
	if p.Map != nil {
		body.Map = make(map[uint]interface{}, len(p.Map))
		for key, val := range p.Map {
			tk := key
			tv := val
			body.Map[tk] = tv
		}
	}
	if p.Object != nil {
		body.Object = marshalServicemixedpayloadinbodyBPayloadToBPayloadRequestBody(p.Object)
	}
	if p.DupObj != nil {
		body.DupObj = marshalServicemixedpayloadinbodyBPayloadToBPayloadRequestBody(p.DupObj)
	}
	return body
}
`

const MultipleMethodsClientTypesFile = `// MethodARequestBody is the type of the "ServiceMultipleMethods" service
// "MethodA" endpoint HTTP request body.
type MethodARequestBody struct {
	A *string ` + "`" + `form:"a,omitempty" json:"a,omitempty" xml:"a,omitempty"` + "`" + `
}

// MethodBRequestBody is the type of the "ServiceMultipleMethods" service
// "MethodB" endpoint HTTP request body.
type MethodBRequestBody struct {
	A string               ` + "`" + `form:"a" json:"a" xml:"a"` + "`" + `
	B *string              ` + "`" + `form:"b,omitempty" json:"b,omitempty" xml:"b,omitempty"` + "`" + `
	C *APayloadRequestBody ` + "`" + `form:"c" json:"c" xml:"c"` + "`" + `
}

// APayloadRequestBody is used to define fields on request body types.
type APayloadRequestBody struct {
	A *string ` + "`" + `form:"a,omitempty" json:"a,omitempty" xml:"a,omitempty"` + "`" + `
}

// NewMethodARequestBody builds the HTTP request body from the payload of the
// "MethodA" endpoint of the "ServiceMultipleMethods" service.
func NewMethodARequestBody(p *servicemultiplemethods.APayload) *MethodARequestBody {
	body := &MethodARequestBody{
		A: p.A,
	}
	return body
}

// NewMethodBRequestBody builds the HTTP request body from the payload of the
// "MethodB" endpoint of the "ServiceMultipleMethods" service.
func NewMethodBRequestBody(p *servicemultiplemethods.PayloadType) *MethodBRequestBody {
	body := &MethodBRequestBody{
		A: p.A,
		B: p.B,
	}
	if p.C != nil {
		body.C = marshalServicemultiplemethodsAPayloadToAPayloadRequestBody(p.C)
	}
	return body
}

// ValidateAPayloadRequestBody runs the validations defined on
// APayloadRequestBody
func ValidateAPayloadRequestBody(body *APayloadRequestBody) (err error) {
	if body.A != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("body.a", *body.A, "patterna"))
	}
	return
}
`

const PayloadExtendedValidateClientTypesFile = `// MethodQueryStringExtendedValidatePayloadRequestBody is the type of the
// "ServiceQueryStringExtendedValidatePayload" service
// "MethodQueryStringExtendedValidatePayload" endpoint HTTP request body.
type MethodQueryStringExtendedValidatePayloadRequestBody struct {
	Body string ` + "`" + `form:"body" json:"body" xml:"body"` + "`" + `
}

// NewMethodQueryStringExtendedValidatePayloadRequestBody builds the HTTP
// request body from the payload of the
// "MethodQueryStringExtendedValidatePayload" endpoint of the
// "ServiceQueryStringExtendedValidatePayload" service.
func NewMethodQueryStringExtendedValidatePayloadRequestBody(p *servicequerystringextendedvalidatepayload.MethodQueryStringExtendedValidatePayloadPayload) *MethodQueryStringExtendedValidatePayloadRequestBody {
	body := &MethodQueryStringExtendedValidatePayloadRequestBody{
		Body: p.Body,
	}
	return body
}
`

const ResultTypeValidateClientTypesFile = `// MethodResultTypeValidateResponseBody is the type of the
// "ServiceResultTypeValidate" service "MethodResultTypeValidate" endpoint HTTP
// response body.
type MethodResultTypeValidateResponseBody struct {
	A *string ` + "`" + `form:"a,omitempty" json:"a,omitempty" xml:"a,omitempty"` + "`" + `
}

// NewMethodResultTypeValidateResultTypeNoContent builds a
// "ServiceResultTypeValidate" service "MethodResultTypeValidate" endpoint
// result from a HTTP "NoContent" response.
func NewMethodResultTypeValidateResultTypeNoContent(body *MethodResultTypeValidateResponseBody) *serviceresulttypevalidate.ResultType {
	v := &serviceresulttypevalidate.ResultType{
		A: body.A,
	}

	return v
}

// ValidateMethodResultTypeValidateResponseBody runs the validations defined on
// MethodResultTypeValidateResponseBody
func ValidateMethodResultTypeValidateResponseBody(body *MethodResultTypeValidateResponseBody) (err error) {
	if body.A != nil {
		if utf8.RuneCountInString(*body.A) < 5 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.a", *body.A, utf8.RuneCountInString(*body.A), 5, true))
		}
	}
	return
}
`

const WithResultCollectionClientTypesFile = `// MethodResultWithResultCollectionResponseBody is the type of the
// "ServiceResultWithResultCollection" service
// "MethodResultWithResultCollection" endpoint HTTP response body.
type MethodResultWithResultCollectionResponseBody struct {
	A *ResulttypeResponseBody ` + "`" + `form:"a,omitempty" json:"a,omitempty" xml:"a,omitempty"` + "`" + `
}

// ResulttypeResponseBody is used to define fields on response body types.
type ResulttypeResponseBody struct {
	X RtCollectionResponseBody ` + "`" + `form:"x,omitempty" json:"x,omitempty" xml:"x,omitempty"` + "`" + `
}

// RtCollectionResponseBody is used to define fields on response body types.
type RtCollectionResponseBody []*RtResponseBody

// RtResponseBody is used to define fields on response body types.
type RtResponseBody struct {
	X *string ` + "`" + `form:"x,omitempty" json:"x,omitempty" xml:"x,omitempty"` + "`" + `
}

// NewMethodResultWithResultCollectionResultOK builds a
// "ServiceResultWithResultCollection" service
// "MethodResultWithResultCollection" endpoint result from a HTTP "OK" response.
func NewMethodResultWithResultCollectionResultOK(body *MethodResultWithResultCollectionResponseBody) *serviceresultwithresultcollection.MethodResultWithResultCollectionResult {
	v := &serviceresultwithresultcollection.MethodResultWithResultCollectionResult{}
	if body.A != nil {
		v.A = unmarshalResulttypeResponseBodyToServiceresultwithresultcollectionResulttype(body.A)
	}

	return v
}

// ValidateMethodResultWithResultCollectionResponseBody runs the validations
// defined on MethodResultWithResultCollectionResponseBody
func ValidateMethodResultWithResultCollectionResponseBody(body *MethodResultWithResultCollectionResponseBody) (err error) {
	if body.A != nil {
		if err2 := ValidateResulttypeResponseBody(body.A); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateResulttypeResponseBody runs the validations defined on
// ResulttypeResponseBody
func ValidateResulttypeResponseBody(body *ResulttypeResponseBody) (err error) {
	if err2 := ValidateRtCollectionResponseBody(body.X); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	return
}

// ValidateRtCollectionResponseBody runs the validations defined on
// RtCollectionResponseBody
func ValidateRtCollectionResponseBody(body RtCollectionResponseBody) (err error) {
	for _, e := range body {
		if e != nil {
			if err2 := ValidateRtResponseBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateRtResponseBody runs the validations defined on RtResponseBody
func ValidateRtResponseBody(body *RtResponseBody) (err error) {
	if body.X != nil {
		if utf8.RuneCountInString(*body.X) < 5 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.x", *body.X, utf8.RuneCountInString(*body.X), 5, true))
		}
	}
	return
}
`
