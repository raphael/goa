package codegen

import (
	"bytes"
	"testing"

	"goa.design/goa/v3/codegen"
	"goa.design/goa/v3/expr"
	"goa.design/goa/v3/http/codegen/testdata"
)

func TestServerTypes(t *testing.T) {
	const genpkg = "gen"
	cases := []struct {
		Name string
		DSL  func()
		Code string
	}{
		{"mixed-payload-attrs", testdata.MixedPayloadInBodyDSL, MixedPayloadInBodyServerTypesFile},
		{"multiple-methods", testdata.MultipleMethodsDSL, MultipleMethodsServerTypesFile},
		{"payload-extend-validate", testdata.PayloadExtendedValidateDSL, PayloadExtendedValidateServerTypesFile},
		{"result-type-validate", testdata.ResultTypeValidateDSL, ResultTypeValidateServerTypesFile},
		{"with-result-collection", testdata.ResultWithResultCollectionDSL, ResultWithResultCollectionServerTypesFile},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			RunHTTPDSL(t, c.DSL)
			fs := serverType(genpkg, expr.Root.API.HTTP.Services[0], make(map[string]struct{}))
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

const MixedPayloadInBodyServerTypesFile = `// MethodARequestBody is the type of the "ServiceMixedPayloadInBody" service
// "MethodA" endpoint HTTP request body.
type MethodARequestBody struct {
	Any    interface{}          ` + "`" + `form:"any,omitempty" json:"any,omitempty" xml:"any,omitempty"` + "`" + `
	Array  []float32            ` + "`" + `form:"array,omitempty" json:"array,omitempty" xml:"array,omitempty"` + "`" + `
	Map    map[uint]interface{} ` + "`" + `form:"map,omitempty" json:"map,omitempty" xml:"map,omitempty"` + "`" + `
	Object *BPayloadRequestBody ` + "`" + `form:"object,omitempty" json:"object,omitempty" xml:"object,omitempty"` + "`" + `
	DupObj *BPayloadRequestBody ` + "`" + `form:"dup_obj,omitempty" json:"dup_obj,omitempty" xml:"dup_obj,omitempty"` + "`" + `
}

// BPayloadRequestBody is used to define fields on request body types.
type BPayloadRequestBody struct {
	Int   *int   ` + "`" + `form:"int,omitempty" json:"int,omitempty" xml:"int,omitempty"` + "`" + `
	Bytes []byte ` + "`" + `form:"bytes,omitempty" json:"bytes,omitempty" xml:"bytes,omitempty"` + "`" + `
}

// NewMethodAAPayload builds a ServiceMixedPayloadInBody service MethodA
// endpoint payload.
func NewMethodAAPayload(body *MethodARequestBody) *servicemixedpayloadinbody.APayload {
	v := &servicemixedpayloadinbody.APayload{
		Any: body.Any,
	}
	v.Array = make([]float32, len(body.Array))
	for i, val := range body.Array {
		v.Array[i] = val
	}
	if body.Map != nil {
		v.Map = make(map[uint]interface{}, len(body.Map))
		for key, val := range body.Map {
			tk := key
			tv := val
			v.Map[tk] = tv
		}
	}
	v.Object = unmarshalBPayloadRequestBodyToServicemixedpayloadinbodyBPayload(body.Object)
	if body.DupObj != nil {
		v.DupObj = unmarshalBPayloadRequestBodyToServicemixedpayloadinbodyBPayload(body.DupObj)
	}

	return v
}

// ValidateMethodARequestBody runs the validations defined on MethodARequestBody
func ValidateMethodARequestBody(body *MethodARequestBody) (err error) {
	if body.Array == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("array", "body"))
	}
	if body.Object == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("object", "body"))
	}
	if body.Object != nil {
		if err2 := ValidateBPayloadRequestBody(body.Object); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if body.DupObj != nil {
		if err2 := ValidateBPayloadRequestBody(body.DupObj); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateBPayloadRequestBody runs the validations defined on
// BPayloadRequestBody
func ValidateBPayloadRequestBody(body *BPayloadRequestBody) (err error) {
	if body.Int == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("int", "body"))
	}
	return
}
`

const MultipleMethodsServerTypesFile = `// MethodARequestBody is the type of the "ServiceMultipleMethods" service
// "MethodA" endpoint HTTP request body.
type MethodARequestBody struct {
	A *string ` + "`" + `form:"a,omitempty" json:"a,omitempty" xml:"a,omitempty"` + "`" + `
}

// MethodBRequestBody is the type of the "ServiceMultipleMethods" service
// "MethodB" endpoint HTTP request body.
type MethodBRequestBody struct {
	A *string              ` + "`" + `form:"a,omitempty" json:"a,omitempty" xml:"a,omitempty"` + "`" + `
	B *string              ` + "`" + `form:"b,omitempty" json:"b,omitempty" xml:"b,omitempty"` + "`" + `
	C *APayloadRequestBody ` + "`" + `form:"c,omitempty" json:"c,omitempty" xml:"c,omitempty"` + "`" + `
}

// APayloadRequestBody is used to define fields on request body types.
type APayloadRequestBody struct {
	A *string ` + "`" + `form:"a,omitempty" json:"a,omitempty" xml:"a,omitempty"` + "`" + `
}

// NewMethodAAPayload builds a ServiceMultipleMethods service MethodA endpoint
// payload.
func NewMethodAAPayload(body *MethodARequestBody) *servicemultiplemethods.APayload {
	v := &servicemultiplemethods.APayload{
		A: body.A,
	}

	return v
}

// NewMethodBPayloadType builds a ServiceMultipleMethods service MethodB
// endpoint payload.
func NewMethodBPayloadType(body *MethodBRequestBody) *servicemultiplemethods.PayloadType {
	v := &servicemultiplemethods.PayloadType{
		A: *body.A,
		B: body.B,
	}
	v.C = unmarshalAPayloadRequestBodyToServicemultiplemethodsAPayload(body.C)

	return v
}

// ValidateMethodARequestBody runs the validations defined on MethodARequestBody
func ValidateMethodARequestBody(body *MethodARequestBody) (err error) {
	if body.A != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("body.a", *body.A, "patterna"))
	}
	return
}

// ValidateMethodBRequestBody runs the validations defined on MethodBRequestBody
func ValidateMethodBRequestBody(body *MethodBRequestBody) (err error) {
	if body.A == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("a", "body"))
	}
	if body.C == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("c", "body"))
	}
	if body.A != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("body.a", *body.A, "patterna"))
	}
	if body.B != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("body.b", *body.B, "patternb"))
	}
	if body.C != nil {
		if err2 := ValidateAPayloadRequestBody(body.C); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
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

const PayloadExtendedValidateServerTypesFile = `// MethodQueryStringExtendedValidatePayloadRequestBody is the type of the
// "ServiceQueryStringExtendedValidatePayload" service
// "MethodQueryStringExtendedValidatePayload" endpoint HTTP request body.
type MethodQueryStringExtendedValidatePayloadRequestBody struct {
	Body *string ` + "`" + `form:"body,omitempty" json:"body,omitempty" xml:"body,omitempty"` + "`" + `
}

// NewMethodQueryStringExtendedValidatePayloadPayload builds a
// ServiceQueryStringExtendedValidatePayload service
// MethodQueryStringExtendedValidatePayload endpoint payload.
func NewMethodQueryStringExtendedValidatePayloadPayload(body *MethodQueryStringExtendedValidatePayloadRequestBody, q string, h int) *servicequerystringextendedvalidatepayload.MethodQueryStringExtendedValidatePayloadPayload {
	v := &servicequerystringextendedvalidatepayload.MethodQueryStringExtendedValidatePayloadPayload{
		Body: *body.Body,
	}
	v.Q = q
	v.H = h

	return v
}

// ValidateMethodQueryStringExtendedValidatePayloadRequestBody runs the
// validations defined on MethodQueryStringExtendedValidatePayloadRequestBody
func ValidateMethodQueryStringExtendedValidatePayloadRequestBody(body *MethodQueryStringExtendedValidatePayloadRequestBody) (err error) {
	if body.Body == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("body", "body"))
	}
	return
}
`

const ResultTypeValidateServerTypesFile = `// MethodResultTypeValidateResponseBody is the type of the
// "ServiceResultTypeValidate" service "MethodResultTypeValidate" endpoint HTTP
// response body.
type MethodResultTypeValidateResponseBody struct {
	A *string ` + "`" + `form:"a,omitempty" json:"a,omitempty" xml:"a,omitempty"` + "`" + `
}

// NewMethodResultTypeValidateResponseBody builds the HTTP response body from
// the result of the "MethodResultTypeValidate" endpoint of the
// "ServiceResultTypeValidate" service.
func NewMethodResultTypeValidateResponseBody(res *serviceresulttypevalidate.ResultType) *MethodResultTypeValidateResponseBody {
	body := &MethodResultTypeValidateResponseBody{
		A: res.A,
	}
	return body
}
`

const ResultWithResultCollectionServerTypesFile = `// MethodResultWithResultCollectionResponseBody is the type of the
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

// NewMethodResultWithResultCollectionResponseBody builds the HTTP response
// body from the result of the "MethodResultWithResultCollection" endpoint of
// the "ServiceResultWithResultCollection" service.
func NewMethodResultWithResultCollectionResponseBody(res *serviceresultwithresultcollection.MethodResultWithResultCollectionResult) *MethodResultWithResultCollectionResponseBody {
	body := &MethodResultWithResultCollectionResponseBody{}
	if res.A != nil {
		body.A = marshalServiceresultwithresultcollectionResulttypeToResulttypeResponseBody(res.A)
	}
	return body
}
`
