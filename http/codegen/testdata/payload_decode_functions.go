package testdata

var PayloadQueryBoolDecodeCode = `// DecodeMethodQueryBoolRequest returns a decoder for requests sent to the
// ServiceQueryBool MethodQueryBool endpoint.
func DecodeMethodQueryBoolRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   *bool
			err error
		)
		{
			qRaw := r.URL.Query().Get("q")
			if qRaw != "" {
				v, err2 := strconv.ParseBool(qRaw)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("q", qRaw, "boolean"))
				}
				q = &v
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodQueryBoolPayload(q)

		return payload, nil
	}
}
`

var PayloadQueryBoolValidateDecodeCode = `// DecodeMethodQueryBoolValidateRequest returns a decoder for requests sent to
// the ServiceQueryBoolValidate MethodQueryBoolValidate endpoint.
func DecodeMethodQueryBoolValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   bool
			err error
		)
		{
			qRaw := r.URL.Query().Get("q")
			if qRaw == "" {
				err = goa.MergeErrors(err, goa.MissingFieldError("q", "query string"))
			}
			v, err2 := strconv.ParseBool(qRaw)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("q", qRaw, "boolean"))
			}
			q = v
		}
		if !(q == true) {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("q", q, []interface{}{true}))
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodQueryBoolValidatePayload(q)

		return payload, nil
	}
}
`

var PayloadQueryIntDecodeCode = `// DecodeMethodQueryIntRequest returns a decoder for requests sent to the
// ServiceQueryInt MethodQueryInt endpoint.
func DecodeMethodQueryIntRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   *int
			err error
		)
		{
			qRaw := r.URL.Query().Get("q")
			if qRaw != "" {
				v, err2 := strconv.ParseInt(qRaw, 10, strconv.IntSize)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("q", qRaw, "integer"))
				}
				pv := int(v)
				q = &pv
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodQueryIntPayload(q)

		return payload, nil
	}
}
`

var PayloadQueryIntValidateDecodeCode = `// DecodeMethodQueryIntValidateRequest returns a decoder for requests sent to
// the ServiceQueryIntValidate MethodQueryIntValidate endpoint.
func DecodeMethodQueryIntValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   int
			err error
		)
		{
			qRaw := r.URL.Query().Get("q")
			if qRaw == "" {
				err = goa.MergeErrors(err, goa.MissingFieldError("q", "query string"))
			}
			v, err2 := strconv.ParseInt(qRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("q", qRaw, "integer"))
			}
			q = int(v)
		}
		if q < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("q", q, 1, true))
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodQueryIntValidatePayload(q)

		return payload, nil
	}
}
`

var PayloadQueryInt32DecodeCode = `// DecodeMethodQueryInt32Request returns a decoder for requests sent to the
// ServiceQueryInt32 MethodQueryInt32 endpoint.
func DecodeMethodQueryInt32Request(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   *int32
			err error
		)
		{
			qRaw := r.URL.Query().Get("q")
			if qRaw != "" {
				v, err2 := strconv.ParseInt(qRaw, 10, 32)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("q", qRaw, "integer"))
				}
				pv := int32(v)
				q = &pv
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodQueryInt32Payload(q)

		return payload, nil
	}
}
`

var PayloadQueryInt32ValidateDecodeCode = `// DecodeMethodQueryInt32ValidateRequest returns a decoder for requests sent to
// the ServiceQueryInt32Validate MethodQueryInt32Validate endpoint.
func DecodeMethodQueryInt32ValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   int32
			err error
		)
		{
			qRaw := r.URL.Query().Get("q")
			if qRaw == "" {
				err = goa.MergeErrors(err, goa.MissingFieldError("q", "query string"))
			}
			v, err2 := strconv.ParseInt(qRaw, 10, 32)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("q", qRaw, "integer"))
			}
			q = int32(v)
		}
		if q < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("q", q, 1, true))
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodQueryInt32ValidatePayload(q)

		return payload, nil
	}
}
`

var PayloadQueryInt64DecodeCode = `// DecodeMethodQueryInt64Request returns a decoder for requests sent to the
// ServiceQueryInt64 MethodQueryInt64 endpoint.
func DecodeMethodQueryInt64Request(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   *int64
			err error
		)
		{
			qRaw := r.URL.Query().Get("q")
			if qRaw != "" {
				v, err2 := strconv.ParseInt(qRaw, 10, 64)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("q", qRaw, "integer"))
				}
				q = &v
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodQueryInt64Payload(q)

		return payload, nil
	}
}
`

var PayloadQueryInt64ValidateDecodeCode = `// DecodeMethodQueryInt64ValidateRequest returns a decoder for requests sent to
// the ServiceQueryInt64Validate MethodQueryInt64Validate endpoint.
func DecodeMethodQueryInt64ValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   int64
			err error
		)
		{
			qRaw := r.URL.Query().Get("q")
			if qRaw == "" {
				err = goa.MergeErrors(err, goa.MissingFieldError("q", "query string"))
			}
			v, err2 := strconv.ParseInt(qRaw, 10, 64)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("q", qRaw, "integer"))
			}
			q = v
		}
		if q < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("q", q, 1, true))
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodQueryInt64ValidatePayload(q)

		return payload, nil
	}
}
`

var PayloadQueryUIntDecodeCode = `// DecodeMethodQueryUIntRequest returns a decoder for requests sent to the
// ServiceQueryUInt MethodQueryUInt endpoint.
func DecodeMethodQueryUIntRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   *uint
			err error
		)
		{
			qRaw := r.URL.Query().Get("q")
			if qRaw != "" {
				v, err2 := strconv.ParseUint(qRaw, 10, strconv.IntSize)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("q", qRaw, "unsigned integer"))
				}
				pv := uint(v)
				q = &pv
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodQueryUIntPayload(q)

		return payload, nil
	}
}
`

var PayloadQueryUIntValidateDecodeCode = `// DecodeMethodQueryUIntValidateRequest returns a decoder for requests sent to
// the ServiceQueryUIntValidate MethodQueryUIntValidate endpoint.
func DecodeMethodQueryUIntValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   uint
			err error
		)
		{
			qRaw := r.URL.Query().Get("q")
			if qRaw == "" {
				err = goa.MergeErrors(err, goa.MissingFieldError("q", "query string"))
			}
			v, err2 := strconv.ParseUint(qRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("q", qRaw, "unsigned integer"))
			}
			q = uint(v)
		}
		if q < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("q", q, 1, true))
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodQueryUIntValidatePayload(q)

		return payload, nil
	}
}
`

var PayloadQueryUInt32DecodeCode = `// DecodeMethodQueryUInt32Request returns a decoder for requests sent to the
// ServiceQueryUInt32 MethodQueryUInt32 endpoint.
func DecodeMethodQueryUInt32Request(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   *uint32
			err error
		)
		{
			qRaw := r.URL.Query().Get("q")
			if qRaw != "" {
				v, err2 := strconv.ParseUint(qRaw, 10, 32)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("q", qRaw, "unsigned integer"))
				}
				pv := uint32(v)
				q = &pv
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodQueryUInt32Payload(q)

		return payload, nil
	}
}
`

var PayloadQueryUInt32ValidateDecodeCode = `// DecodeMethodQueryUInt32ValidateRequest returns a decoder for requests sent
// to the ServiceQueryUInt32Validate MethodQueryUInt32Validate endpoint.
func DecodeMethodQueryUInt32ValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   uint32
			err error
		)
		{
			qRaw := r.URL.Query().Get("q")
			if qRaw == "" {
				err = goa.MergeErrors(err, goa.MissingFieldError("q", "query string"))
			}
			v, err2 := strconv.ParseUint(qRaw, 10, 32)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("q", qRaw, "unsigned integer"))
			}
			q = uint32(v)
		}
		if q < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("q", q, 1, true))
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodQueryUInt32ValidatePayload(q)

		return payload, nil
	}
}
`

var PayloadQueryUInt64DecodeCode = `// DecodeMethodQueryUInt64Request returns a decoder for requests sent to the
// ServiceQueryUInt64 MethodQueryUInt64 endpoint.
func DecodeMethodQueryUInt64Request(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   *uint64
			err error
		)
		{
			qRaw := r.URL.Query().Get("q")
			if qRaw != "" {
				v, err2 := strconv.ParseUint(qRaw, 10, 64)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("q", qRaw, "unsigned integer"))
				}
				q = &v
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodQueryUInt64Payload(q)

		return payload, nil
	}
}
`

var PayloadQueryUInt64ValidateDecodeCode = `// DecodeMethodQueryUInt64ValidateRequest returns a decoder for requests sent
// to the ServiceQueryUInt64Validate MethodQueryUInt64Validate endpoint.
func DecodeMethodQueryUInt64ValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   uint64
			err error
		)
		{
			qRaw := r.URL.Query().Get("q")
			if qRaw == "" {
				err = goa.MergeErrors(err, goa.MissingFieldError("q", "query string"))
			}
			v, err2 := strconv.ParseUint(qRaw, 10, 64)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("q", qRaw, "unsigned integer"))
			}
			q = v
		}
		if q < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("q", q, 1, true))
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodQueryUInt64ValidatePayload(q)

		return payload, nil
	}
}
`

var PayloadQueryFloat32DecodeCode = `// DecodeMethodQueryFloat32Request returns a decoder for requests sent to the
// ServiceQueryFloat32 MethodQueryFloat32 endpoint.
func DecodeMethodQueryFloat32Request(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   *float32
			err error
		)
		{
			qRaw := r.URL.Query().Get("q")
			if qRaw != "" {
				v, err2 := strconv.ParseFloat(qRaw, 32)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("q", qRaw, "float"))
				}
				pv := float32(v)
				q = &pv
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodQueryFloat32Payload(q)

		return payload, nil
	}
}
`

var PayloadQueryFloat32ValidateDecodeCode = `// DecodeMethodQueryFloat32ValidateRequest returns a decoder for requests sent
// to the ServiceQueryFloat32Validate MethodQueryFloat32Validate endpoint.
func DecodeMethodQueryFloat32ValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   float32
			err error
		)
		{
			qRaw := r.URL.Query().Get("q")
			if qRaw == "" {
				err = goa.MergeErrors(err, goa.MissingFieldError("q", "query string"))
			}
			v, err2 := strconv.ParseFloat(qRaw, 32)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("q", qRaw, "float"))
			}
			q = float32(v)
		}
		if q < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("q", q, 1, true))
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodQueryFloat32ValidatePayload(q)

		return payload, nil
	}
}
`

var PayloadQueryFloat64DecodeCode = `// DecodeMethodQueryFloat64Request returns a decoder for requests sent to the
// ServiceQueryFloat64 MethodQueryFloat64 endpoint.
func DecodeMethodQueryFloat64Request(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   *float64
			err error
		)
		{
			qRaw := r.URL.Query().Get("q")
			if qRaw != "" {
				v, err2 := strconv.ParseFloat(qRaw, 64)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("q", qRaw, "float"))
				}
				q = &v
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodQueryFloat64Payload(q)

		return payload, nil
	}
}
`

var PayloadQueryFloat64ValidateDecodeCode = `// DecodeMethodQueryFloat64ValidateRequest returns a decoder for requests sent
// to the ServiceQueryFloat64Validate MethodQueryFloat64Validate endpoint.
func DecodeMethodQueryFloat64ValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   float64
			err error
		)
		{
			qRaw := r.URL.Query().Get("q")
			if qRaw == "" {
				err = goa.MergeErrors(err, goa.MissingFieldError("q", "query string"))
			}
			v, err2 := strconv.ParseFloat(qRaw, 64)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("q", qRaw, "float"))
			}
			q = v
		}
		if q < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("q", q, 1, true))
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodQueryFloat64ValidatePayload(q)

		return payload, nil
	}
}
`

var PayloadQueryStringDecodeCode = `// DecodeMethodQueryStringRequest returns a decoder for requests sent to the
// ServiceQueryString MethodQueryString endpoint.
func DecodeMethodQueryStringRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q *string
		)
		qRaw := r.URL.Query().Get("q")
		if qRaw != "" {
			q = &qRaw
		}
		payload := NewMethodQueryStringPayload(q)

		return payload, nil
	}
}
`

var PayloadQueryStringValidateDecodeCode = `// DecodeMethodQueryStringValidateRequest returns a decoder for requests sent
// to the ServiceQueryStringValidate MethodQueryStringValidate endpoint.
func DecodeMethodQueryStringValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   string
			err error
		)
		q = r.URL.Query().Get("q")
		if q == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("q", "query string"))
		}
		if !(q == "val") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("q", q, []interface{}{"val"}))
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodQueryStringValidatePayload(q)

		return payload, nil
	}
}
`

var PayloadQueryStringNotRequiredValidateDecodeCode = `// DecodeMethodQueryStringNotRequiredValidateRequest returns a decoder for
// requests sent to the ServiceQueryStringNotRequiredValidate
// MethodQueryStringNotRequiredValidate endpoint.
func DecodeMethodQueryStringNotRequiredValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   *string
			err error
		)
		qRaw := r.URL.Query().Get("q")
		if qRaw != "" {
			q = &qRaw
		}
		if q != nil {
			if !(*q == "val") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("q", *q, []interface{}{"val"}))
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodQueryStringNotRequiredValidatePayload(q)

		return payload, nil
	}
}
`

var PayloadQueryBytesDecodeCode = `// DecodeMethodQueryBytesRequest returns a decoder for requests sent to the
// ServiceQueryBytes MethodQueryBytes endpoint.
func DecodeMethodQueryBytesRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q []byte
		)
		{
			qRaw := r.URL.Query().Get("q")
			if qRaw != "" {
				q = []byte(qRaw)
			}
		}
		payload := NewMethodQueryBytesPayload(q)

		return payload, nil
	}
}
`

var PayloadQueryBytesValidateDecodeCode = `// DecodeMethodQueryBytesValidateRequest returns a decoder for requests sent to
// the ServiceQueryBytesValidate MethodQueryBytesValidate endpoint.
func DecodeMethodQueryBytesValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   []byte
			err error
		)
		{
			qRaw := r.URL.Query().Get("q")
			if qRaw == "" {
				err = goa.MergeErrors(err, goa.MissingFieldError("q", "query string"))
			}
			q = []byte(qRaw)
		}
		if len(q) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("q", q, len(q), 1, true))
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodQueryBytesValidatePayload(q)

		return payload, nil
	}
}
`

var PayloadQueryAnyDecodeCode = `// DecodeMethodQueryAnyRequest returns a decoder for requests sent to the
// ServiceQueryAny MethodQueryAny endpoint.
func DecodeMethodQueryAnyRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q interface{}
		)
		qRaw := r.URL.Query().Get("q")
		if qRaw != "" {
			q = qRaw
		}
		payload := NewMethodQueryAnyPayload(q)

		return payload, nil
	}
}
`

var PayloadQueryAnyValidateDecodeCode = `// DecodeMethodQueryAnyValidateRequest returns a decoder for requests sent to
// the ServiceQueryAnyValidate MethodQueryAnyValidate endpoint.
func DecodeMethodQueryAnyValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   interface{}
			err error
		)
		q = r.URL.Query().Get("q")
		if q == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("q", "query string"))
		}
		if !(q == "val" || q == 1) {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("q", q, []interface{}{"val", 1}))
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodQueryAnyValidatePayload(q)

		return payload, nil
	}
}
`

var PayloadQueryArrayBoolDecodeCode = `// DecodeMethodQueryArrayBoolRequest returns a decoder for requests sent to the
// ServiceQueryArrayBool MethodQueryArrayBool endpoint.
func DecodeMethodQueryArrayBoolRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   []bool
			err error
		)
		{
			qRaw := r.URL.Query()["q"]
			if qRaw != nil {
				q = make([]bool, len(qRaw))
				for i, rv := range qRaw {
					v, err2 := strconv.ParseBool(rv)
					if err2 != nil {
						err = goa.MergeErrors(err, goa.InvalidFieldTypeError("q", qRaw, "array of booleans"))
					}
					q[i] = v
				}
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodQueryArrayBoolPayload(q)

		return payload, nil
	}
}
`

var PayloadQueryArrayBoolValidateDecodeCode = `// DecodeMethodQueryArrayBoolValidateRequest returns a decoder for requests
// sent to the ServiceQueryArrayBoolValidate MethodQueryArrayBoolValidate
// endpoint.
func DecodeMethodQueryArrayBoolValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   []bool
			err error
		)
		{
			qRaw := r.URL.Query()["q"]
			if qRaw == nil {
				return goa.MergeErrors(err, goa.MissingFieldError("q", "query string"))
			}
			q = make([]bool, len(qRaw))
			for i, rv := range qRaw {
				v, err2 := strconv.ParseBool(rv)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("q", qRaw, "array of booleans"))
				}
				q[i] = v
			}
		}
		if len(q) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("q", q, len(q), 1, true))
		}
		for _, e := range q {
			if !(e == true) {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("q[*]", e, []interface{}{true}))
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodQueryArrayBoolValidatePayload(q)

		return payload, nil
	}
}
`

var PayloadQueryArrayIntDecodeCode = `// DecodeMethodQueryArrayIntRequest returns a decoder for requests sent to the
// ServiceQueryArrayInt MethodQueryArrayInt endpoint.
func DecodeMethodQueryArrayIntRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   []int
			err error
		)
		{
			qRaw := r.URL.Query()["q"]
			if qRaw != nil {
				q = make([]int, len(qRaw))
				for i, rv := range qRaw {
					v, err2 := strconv.ParseInt(rv, 10, strconv.IntSize)
					if err2 != nil {
						err = goa.MergeErrors(err, goa.InvalidFieldTypeError("q", qRaw, "array of integers"))
					}
					q[i] = int(v)
				}
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodQueryArrayIntPayload(q)

		return payload, nil
	}
}
`

var PayloadQueryArrayIntValidateDecodeCode = `// DecodeMethodQueryArrayIntValidateRequest returns a decoder for requests sent
// to the ServiceQueryArrayIntValidate MethodQueryArrayIntValidate endpoint.
func DecodeMethodQueryArrayIntValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   []int
			err error
		)
		{
			qRaw := r.URL.Query()["q"]
			if qRaw == nil {
				return goa.MergeErrors(err, goa.MissingFieldError("q", "query string"))
			}
			q = make([]int, len(qRaw))
			for i, rv := range qRaw {
				v, err2 := strconv.ParseInt(rv, 10, strconv.IntSize)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("q", qRaw, "array of integers"))
				}
				q[i] = int(v)
			}
		}
		if len(q) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("q", q, len(q), 1, true))
		}
		for _, e := range q {
			if e < 1 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("q[*]", e, 1, true))
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodQueryArrayIntValidatePayload(q)

		return payload, nil
	}
}
`

var PayloadQueryArrayInt32DecodeCode = `// DecodeMethodQueryArrayInt32Request returns a decoder for requests sent to
// the ServiceQueryArrayInt32 MethodQueryArrayInt32 endpoint.
func DecodeMethodQueryArrayInt32Request(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   []int32
			err error
		)
		{
			qRaw := r.URL.Query()["q"]
			if qRaw != nil {
				q = make([]int32, len(qRaw))
				for i, rv := range qRaw {
					v, err2 := strconv.ParseInt(rv, 10, 32)
					if err2 != nil {
						err = goa.MergeErrors(err, goa.InvalidFieldTypeError("q", qRaw, "array of integers"))
					}
					q[i] = int32(v)
				}
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodQueryArrayInt32Payload(q)

		return payload, nil
	}
}
`

var PayloadQueryArrayInt32ValidateDecodeCode = `// DecodeMethodQueryArrayInt32ValidateRequest returns a decoder for requests
// sent to the ServiceQueryArrayInt32Validate MethodQueryArrayInt32Validate
// endpoint.
func DecodeMethodQueryArrayInt32ValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   []int32
			err error
		)
		{
			qRaw := r.URL.Query()["q"]
			if qRaw == nil {
				return goa.MergeErrors(err, goa.MissingFieldError("q", "query string"))
			}
			q = make([]int32, len(qRaw))
			for i, rv := range qRaw {
				v, err2 := strconv.ParseInt(rv, 10, 32)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("q", qRaw, "array of integers"))
				}
				q[i] = int32(v)
			}
		}
		if len(q) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("q", q, len(q), 1, true))
		}
		for _, e := range q {
			if e < 1 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("q[*]", e, 1, true))
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodQueryArrayInt32ValidatePayload(q)

		return payload, nil
	}
}
`

var PayloadQueryArrayInt64DecodeCode = `// DecodeMethodQueryArrayInt64Request returns a decoder for requests sent to
// the ServiceQueryArrayInt64 MethodQueryArrayInt64 endpoint.
func DecodeMethodQueryArrayInt64Request(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   []int64
			err error
		)
		{
			qRaw := r.URL.Query()["q"]
			if qRaw != nil {
				q = make([]int64, len(qRaw))
				for i, rv := range qRaw {
					v, err2 := strconv.ParseInt(rv, 10, 64)
					if err2 != nil {
						err = goa.MergeErrors(err, goa.InvalidFieldTypeError("q", qRaw, "array of integers"))
					}
					q[i] = v
				}
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodQueryArrayInt64Payload(q)

		return payload, nil
	}
}
`

var PayloadQueryArrayInt64ValidateDecodeCode = `// DecodeMethodQueryArrayInt64ValidateRequest returns a decoder for requests
// sent to the ServiceQueryArrayInt64Validate MethodQueryArrayInt64Validate
// endpoint.
func DecodeMethodQueryArrayInt64ValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   []int64
			err error
		)
		{
			qRaw := r.URL.Query()["q"]
			if qRaw == nil {
				return goa.MergeErrors(err, goa.MissingFieldError("q", "query string"))
			}
			q = make([]int64, len(qRaw))
			for i, rv := range qRaw {
				v, err2 := strconv.ParseInt(rv, 10, 64)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("q", qRaw, "array of integers"))
				}
				q[i] = v
			}
		}
		if len(q) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("q", q, len(q), 1, true))
		}
		for _, e := range q {
			if e < 1 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("q[*]", e, 1, true))
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodQueryArrayInt64ValidatePayload(q)

		return payload, nil
	}
}
`

var PayloadQueryArrayUIntDecodeCode = `// DecodeMethodQueryArrayUIntRequest returns a decoder for requests sent to the
// ServiceQueryArrayUInt MethodQueryArrayUInt endpoint.
func DecodeMethodQueryArrayUIntRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   []uint
			err error
		)
		{
			qRaw := r.URL.Query()["q"]
			if qRaw != nil {
				q = make([]uint, len(qRaw))
				for i, rv := range qRaw {
					v, err2 := strconv.ParseUint(rv, 10, strconv.IntSize)
					if err2 != nil {
						err = goa.MergeErrors(err, goa.InvalidFieldTypeError("q", qRaw, "array of unsigned integers"))
					}
					q[i] = uint(v)
				}
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodQueryArrayUIntPayload(q)

		return payload, nil
	}
}
`

var PayloadQueryArrayUIntValidateDecodeCode = `// DecodeMethodQueryArrayUIntValidateRequest returns a decoder for requests
// sent to the ServiceQueryArrayUIntValidate MethodQueryArrayUIntValidate
// endpoint.
func DecodeMethodQueryArrayUIntValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   []uint
			err error
		)
		{
			qRaw := r.URL.Query()["q"]
			if qRaw == nil {
				return goa.MergeErrors(err, goa.MissingFieldError("q", "query string"))
			}
			q = make([]uint, len(qRaw))
			for i, rv := range qRaw {
				v, err2 := strconv.ParseUint(rv, 10, strconv.IntSize)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("q", qRaw, "array of unsigned integers"))
				}
				q[i] = uint(v)
			}
		}
		if len(q) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("q", q, len(q), 1, true))
		}
		for _, e := range q {
			if e < 1 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("q[*]", e, 1, true))
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodQueryArrayUIntValidatePayload(q)

		return payload, nil
	}
}
`

var PayloadQueryArrayUInt32DecodeCode = `// DecodeMethodQueryArrayUInt32Request returns a decoder for requests sent to
// the ServiceQueryArrayUInt32 MethodQueryArrayUInt32 endpoint.
func DecodeMethodQueryArrayUInt32Request(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   []uint32
			err error
		)
		{
			qRaw := r.URL.Query()["q"]
			if qRaw != nil {
				q = make([]uint32, len(qRaw))
				for i, rv := range qRaw {
					v, err2 := strconv.ParseUint(rv, 10, 32)
					if err2 != nil {
						err = goa.MergeErrors(err, goa.InvalidFieldTypeError("q", qRaw, "array of unsigned integers"))
					}
					q[i] = uint32(v)
				}
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodQueryArrayUInt32Payload(q)

		return payload, nil
	}
}
`

var PayloadQueryArrayUInt32ValidateDecodeCode = `// DecodeMethodQueryArrayUInt32ValidateRequest returns a decoder for requests
// sent to the ServiceQueryArrayUInt32Validate MethodQueryArrayUInt32Validate
// endpoint.
func DecodeMethodQueryArrayUInt32ValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   []uint32
			err error
		)
		{
			qRaw := r.URL.Query()["q"]
			if qRaw == nil {
				return goa.MergeErrors(err, goa.MissingFieldError("q", "query string"))
			}
			q = make([]uint32, len(qRaw))
			for i, rv := range qRaw {
				v, err2 := strconv.ParseUint(rv, 10, 32)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("q", qRaw, "array of unsigned integers"))
				}
				q[i] = uint32(v)
			}
		}
		if len(q) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("q", q, len(q), 1, true))
		}
		for _, e := range q {
			if e < 1 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("q[*]", e, 1, true))
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodQueryArrayUInt32ValidatePayload(q)

		return payload, nil
	}
}
`

var PayloadQueryArrayUInt64DecodeCode = `// DecodeMethodQueryArrayUInt64Request returns a decoder for requests sent to
// the ServiceQueryArrayUInt64 MethodQueryArrayUInt64 endpoint.
func DecodeMethodQueryArrayUInt64Request(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   []uint64
			err error
		)
		{
			qRaw := r.URL.Query()["q"]
			if qRaw != nil {
				q = make([]uint64, len(qRaw))
				for i, rv := range qRaw {
					v, err2 := strconv.ParseUint(rv, 10, 64)
					if err2 != nil {
						err = goa.MergeErrors(err, goa.InvalidFieldTypeError("q", qRaw, "array of unsigned integers"))
					}
					q[i] = v
				}
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodQueryArrayUInt64Payload(q)

		return payload, nil
	}
}
`

var PayloadQueryArrayUInt64ValidateDecodeCode = `// DecodeMethodQueryArrayUInt64ValidateRequest returns a decoder for requests
// sent to the ServiceQueryArrayUInt64Validate MethodQueryArrayUInt64Validate
// endpoint.
func DecodeMethodQueryArrayUInt64ValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   []uint64
			err error
		)
		{
			qRaw := r.URL.Query()["q"]
			if qRaw == nil {
				return goa.MergeErrors(err, goa.MissingFieldError("q", "query string"))
			}
			q = make([]uint64, len(qRaw))
			for i, rv := range qRaw {
				v, err2 := strconv.ParseUint(rv, 10, 64)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("q", qRaw, "array of unsigned integers"))
				}
				q[i] = v
			}
		}
		if len(q) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("q", q, len(q), 1, true))
		}
		for _, e := range q {
			if e < 1 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("q[*]", e, 1, true))
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodQueryArrayUInt64ValidatePayload(q)

		return payload, nil
	}
}
`

var PayloadQueryArrayFloat32DecodeCode = `// DecodeMethodQueryArrayFloat32Request returns a decoder for requests sent to
// the ServiceQueryArrayFloat32 MethodQueryArrayFloat32 endpoint.
func DecodeMethodQueryArrayFloat32Request(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   []float32
			err error
		)
		{
			qRaw := r.URL.Query()["q"]
			if qRaw != nil {
				q = make([]float32, len(qRaw))
				for i, rv := range qRaw {
					v, err2 := strconv.ParseFloat(rv, 32)
					if err2 != nil {
						err = goa.MergeErrors(err, goa.InvalidFieldTypeError("q", qRaw, "array of floats"))
					}
					q[i] = float32(v)
				}
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodQueryArrayFloat32Payload(q)

		return payload, nil
	}
}
`

var PayloadQueryArrayFloat32ValidateDecodeCode = `// DecodeMethodQueryArrayFloat32ValidateRequest returns a decoder for requests
// sent to the ServiceQueryArrayFloat32Validate MethodQueryArrayFloat32Validate
// endpoint.
func DecodeMethodQueryArrayFloat32ValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   []float32
			err error
		)
		{
			qRaw := r.URL.Query()["q"]
			if qRaw == nil {
				return goa.MergeErrors(err, goa.MissingFieldError("q", "query string"))
			}
			q = make([]float32, len(qRaw))
			for i, rv := range qRaw {
				v, err2 := strconv.ParseFloat(rv, 32)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("q", qRaw, "array of floats"))
				}
				q[i] = float32(v)
			}
		}
		if len(q) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("q", q, len(q), 1, true))
		}
		for _, e := range q {
			if e < 1 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("q[*]", e, 1, true))
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodQueryArrayFloat32ValidatePayload(q)

		return payload, nil
	}
}
`

var PayloadQueryArrayFloat64DecodeCode = `// DecodeMethodQueryArrayFloat64Request returns a decoder for requests sent to
// the ServiceQueryArrayFloat64 MethodQueryArrayFloat64 endpoint.
func DecodeMethodQueryArrayFloat64Request(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   []float64
			err error
		)
		{
			qRaw := r.URL.Query()["q"]
			if qRaw != nil {
				q = make([]float64, len(qRaw))
				for i, rv := range qRaw {
					v, err2 := strconv.ParseFloat(rv, 64)
					if err2 != nil {
						err = goa.MergeErrors(err, goa.InvalidFieldTypeError("q", qRaw, "array of floats"))
					}
					q[i] = v
				}
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodQueryArrayFloat64Payload(q)

		return payload, nil
	}
}
`

var PayloadQueryArrayFloat64ValidateDecodeCode = `// DecodeMethodQueryArrayFloat64ValidateRequest returns a decoder for requests
// sent to the ServiceQueryArrayFloat64Validate MethodQueryArrayFloat64Validate
// endpoint.
func DecodeMethodQueryArrayFloat64ValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   []float64
			err error
		)
		{
			qRaw := r.URL.Query()["q"]
			if qRaw == nil {
				return goa.MergeErrors(err, goa.MissingFieldError("q", "query string"))
			}
			q = make([]float64, len(qRaw))
			for i, rv := range qRaw {
				v, err2 := strconv.ParseFloat(rv, 64)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("q", qRaw, "array of floats"))
				}
				q[i] = v
			}
		}
		if len(q) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("q", q, len(q), 1, true))
		}
		for _, e := range q {
			if e < 1 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("q[*]", e, 1, true))
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodQueryArrayFloat64ValidatePayload(q)

		return payload, nil
	}
}
`

var PayloadQueryArrayStringDecodeCode = `// DecodeMethodQueryArrayStringRequest returns a decoder for requests sent to
// the ServiceQueryArrayString MethodQueryArrayString endpoint.
func DecodeMethodQueryArrayStringRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q []string
		)
		q = r.URL.Query()["q"]
		payload := NewMethodQueryArrayStringPayload(q)

		return payload, nil
	}
}
`

var PayloadQueryArrayStringValidateDecodeCode = `// DecodeMethodQueryArrayStringValidateRequest returns a decoder for requests
// sent to the ServiceQueryArrayStringValidate MethodQueryArrayStringValidate
// endpoint.
func DecodeMethodQueryArrayStringValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   []string
			err error
		)
		q = r.URL.Query()["q"]
		if q == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("q", "query string"))
		}
		if len(q) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("q", q, len(q), 1, true))
		}
		for _, e := range q {
			if !(e == "val") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("q[*]", e, []interface{}{"val"}))
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodQueryArrayStringValidatePayload(q)

		return payload, nil
	}
}
`

var PayloadQueryArrayBytesDecodeCode = `// DecodeMethodQueryArrayBytesRequest returns a decoder for requests sent to
// the ServiceQueryArrayBytes MethodQueryArrayBytes endpoint.
func DecodeMethodQueryArrayBytesRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q [][]byte
		)
		{
			qRaw := r.URL.Query()["q"]
			if qRaw != nil {
				q = make([][]byte, len(qRaw))
				for i, rv := range qRaw {
					q[i] = []byte(rv)
				}
			}
		}
		payload := NewMethodQueryArrayBytesPayload(q)

		return payload, nil
	}
}
`

var PayloadQueryArrayBytesValidateDecodeCode = `// DecodeMethodQueryArrayBytesValidateRequest returns a decoder for requests
// sent to the ServiceQueryArrayBytesValidate MethodQueryArrayBytesValidate
// endpoint.
func DecodeMethodQueryArrayBytesValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   [][]byte
			err error
		)
		{
			qRaw := r.URL.Query()["q"]
			if qRaw == nil {
				return goa.MergeErrors(err, goa.MissingFieldError("q", "query string"))
			}
			q = make([][]byte, len(qRaw))
			for i, rv := range qRaw {
				q[i] = []byte(rv)
			}
		}
		if len(q) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("q", q, len(q), 1, true))
		}
		for _, e := range q {
			if len(e) < 2 {
				err = goa.MergeErrors(err, goa.InvalidLengthError("q[*]", e, len(e), 2, true))
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodQueryArrayBytesValidatePayload(q)

		return payload, nil
	}
}
`

var PayloadQueryArrayAnyDecodeCode = `// DecodeMethodQueryArrayAnyRequest returns a decoder for requests sent to the
// ServiceQueryArrayAny MethodQueryArrayAny endpoint.
func DecodeMethodQueryArrayAnyRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q []interface{}
		)
		{
			qRaw := r.URL.Query()["q"]
			if qRaw != nil {
				q = make([]interface{}, len(qRaw))
				for i, rv := range qRaw {
					q[i] = rv
				}
			}
		}
		payload := NewMethodQueryArrayAnyPayload(q)

		return payload, nil
	}
}
`

var PayloadQueryArrayAnyValidateDecodeCode = `// DecodeMethodQueryArrayAnyValidateRequest returns a decoder for requests sent
// to the ServiceQueryArrayAnyValidate MethodQueryArrayAnyValidate endpoint.
func DecodeMethodQueryArrayAnyValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   []interface{}
			err error
		)
		{
			qRaw := r.URL.Query()["q"]
			if qRaw == nil {
				return goa.MergeErrors(err, goa.MissingFieldError("q", "query string"))
			}
			q = make([]interface{}, len(qRaw))
			for i, rv := range qRaw {
				q[i] = rv
			}
		}
		if len(q) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("q", q, len(q), 1, true))
		}
		for _, e := range q {
			if !(e == "val" || e == 1) {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("q[*]", e, []interface{}{"val", 1}))
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodQueryArrayAnyValidatePayload(q)

		return payload, nil
	}
}
`

var PayloadQueryMapStringStringDecodeCode = `// DecodeMethodQueryMapStringStringRequest returns a decoder for requests sent
// to the ServiceQueryMapStringString MethodQueryMapStringString endpoint.
func DecodeMethodQueryMapStringStringRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q map[string]string
		)
		{
			qRaw := r.URL.Query()
			if len(qRaw) != 0 {
				for keyRaw, valRaw := range qRaw {
					if strings.HasPrefix(keyRaw, "q[") {
						if q == nil {
							q = make(map[string]string)
						}
						var keya string
						{
							openIdx := strings.IndexRune(keyRaw, '[')
							closeIdx := strings.IndexRune(keyRaw, ']')
							keya = keyRaw[openIdx+1 : closeIdx]
						}
						q[keya] = valRaw[0]
					}
				}
			}
		}
		payload := NewMethodQueryMapStringStringPayload(q)

		return payload, nil
	}
}
`

var PayloadQueryMapStringStringValidateDecodeCode = `// DecodeMethodQueryMapStringStringValidateRequest returns a decoder for
// requests sent to the ServiceQueryMapStringStringValidate
// MethodQueryMapStringStringValidate endpoint.
func DecodeMethodQueryMapStringStringValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   map[string]string
			err error
		)
		{
			qRaw := r.URL.Query()
			if len(qRaw) == 0 {
				err = goa.MergeErrors(err, goa.MissingFieldError("q", "query string"))
			}
			for keyRaw, valRaw := range qRaw {
				if strings.HasPrefix(keyRaw, "q[") {
					if q == nil {
						q = make(map[string]string)
					}
					var keya string
					{
						openIdx := strings.IndexRune(keyRaw, '[')
						closeIdx := strings.IndexRune(keyRaw, ']')
						keya = keyRaw[openIdx+1 : closeIdx]
					}
					q[keya] = valRaw[0]
				}
			}
		}
		if len(q) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("q", q, len(q), 1, true))
		}
		for k, v := range q {
			if !(k == "key") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("q.key", k, []interface{}{"key"}))
			}
			if !(v == "val") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("q[key]", v, []interface{}{"val"}))
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodQueryMapStringStringValidatePayload(q)

		return payload, nil
	}
}
`

var PayloadQueryMapStringBoolDecodeCode = `// DecodeMethodQueryMapStringBoolRequest returns a decoder for requests sent to
// the ServiceQueryMapStringBool MethodQueryMapStringBool endpoint.
func DecodeMethodQueryMapStringBoolRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   map[string]bool
			err error
		)
		{
			qRaw := r.URL.Query()
			if len(qRaw) != 0 {
				for keyRaw, valRaw := range qRaw {
					if strings.HasPrefix(keyRaw, "q[") {
						if q == nil {
							q = make(map[string]bool)
						}
						var keya string
						{
							openIdx := strings.IndexRune(keyRaw, '[')
							closeIdx := strings.IndexRune(keyRaw, ']')
							keya = keyRaw[openIdx+1 : closeIdx]
						}
						var vala bool
						{
							valaRaw := valRaw[0]
							v, err2 := strconv.ParseBool(valaRaw)
							if err2 != nil {
								err = goa.MergeErrors(err, goa.InvalidFieldTypeError("vala", valaRaw, "boolean"))
							}
							vala = v
						}
						q[keya] = vala
					}
				}
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodQueryMapStringBoolPayload(q)

		return payload, nil
	}
}
`

var PayloadQueryMapStringBoolValidateDecodeCode = `// DecodeMethodQueryMapStringBoolValidateRequest returns a decoder for requests
// sent to the ServiceQueryMapStringBoolValidate
// MethodQueryMapStringBoolValidate endpoint.
func DecodeMethodQueryMapStringBoolValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   map[string]bool
			err error
		)
		{
			qRaw := r.URL.Query()
			if len(qRaw) == 0 {
				err = goa.MergeErrors(err, goa.MissingFieldError("q", "query string"))
			}
			for keyRaw, valRaw := range qRaw {
				if strings.HasPrefix(keyRaw, "q[") {
					if q == nil {
						q = make(map[string]bool)
					}
					var keya string
					{
						openIdx := strings.IndexRune(keyRaw, '[')
						closeIdx := strings.IndexRune(keyRaw, ']')
						keya = keyRaw[openIdx+1 : closeIdx]
					}
					var vala bool
					{
						valaRaw := valRaw[0]
						v, err2 := strconv.ParseBool(valaRaw)
						if err2 != nil {
							err = goa.MergeErrors(err, goa.InvalidFieldTypeError("vala", valaRaw, "boolean"))
						}
						vala = v
					}
					q[keya] = vala
				}
			}
		}
		if len(q) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("q", q, len(q), 1, true))
		}
		for k, v := range q {
			if !(k == "key") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("q.key", k, []interface{}{"key"}))
			}
			if !(v == true) {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("q[key]", v, []interface{}{true}))
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodQueryMapStringBoolValidatePayload(q)

		return payload, nil
	}
}
`

var PayloadQueryMapBoolStringDecodeCode = `// DecodeMethodQueryMapBoolStringRequest returns a decoder for requests sent to
// the ServiceQueryMapBoolString MethodQueryMapBoolString endpoint.
func DecodeMethodQueryMapBoolStringRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   map[bool]string
			err error
		)
		{
			qRaw := r.URL.Query()
			if len(qRaw) != 0 {
				for keyRaw, valRaw := range qRaw {
					if strings.HasPrefix(keyRaw, "q[") {
						if q == nil {
							q = make(map[bool]string)
						}
						var keya bool
						{
							openIdx := strings.IndexRune(keyRaw, '[')
							closeIdx := strings.IndexRune(keyRaw, ']')
							keyaRaw := keyRaw[openIdx+1 : closeIdx]
							v, err2 := strconv.ParseBool(keyaRaw)
							if err2 != nil {
								err = goa.MergeErrors(err, goa.InvalidFieldTypeError("keya", keyaRaw, "boolean"))
							}
							keya = v
						}
						q[keya] = valRaw[0]
					}
				}
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodQueryMapBoolStringPayload(q)

		return payload, nil
	}
}
`

var PayloadQueryMapBoolStringValidateDecodeCode = `// DecodeMethodQueryMapBoolStringValidateRequest returns a decoder for requests
// sent to the ServiceQueryMapBoolStringValidate
// MethodQueryMapBoolStringValidate endpoint.
func DecodeMethodQueryMapBoolStringValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   map[bool]string
			err error
		)
		{
			qRaw := r.URL.Query()
			if len(qRaw) == 0 {
				err = goa.MergeErrors(err, goa.MissingFieldError("q", "query string"))
			}
			for keyRaw, valRaw := range qRaw {
				if strings.HasPrefix(keyRaw, "q[") {
					if q == nil {
						q = make(map[bool]string)
					}
					var keya bool
					{
						openIdx := strings.IndexRune(keyRaw, '[')
						closeIdx := strings.IndexRune(keyRaw, ']')
						keyaRaw := keyRaw[openIdx+1 : closeIdx]
						v, err2 := strconv.ParseBool(keyaRaw)
						if err2 != nil {
							err = goa.MergeErrors(err, goa.InvalidFieldTypeError("keya", keyaRaw, "boolean"))
						}
						keya = v
					}
					q[keya] = valRaw[0]
				}
			}
		}
		if len(q) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("q", q, len(q), 1, true))
		}
		for k, v := range q {
			if !(k == true) {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("q.key", k, []interface{}{true}))
			}
			if !(v == "val") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("q[key]", v, []interface{}{"val"}))
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodQueryMapBoolStringValidatePayload(q)

		return payload, nil
	}
}
`

var PayloadQueryMapBoolBoolDecodeCode = `// DecodeMethodQueryMapBoolBoolRequest returns a decoder for requests sent to
// the ServiceQueryMapBoolBool MethodQueryMapBoolBool endpoint.
func DecodeMethodQueryMapBoolBoolRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   map[bool]bool
			err error
		)
		{
			qRaw := r.URL.Query()
			if len(qRaw) != 0 {
				for keyRaw, valRaw := range qRaw {
					if strings.HasPrefix(keyRaw, "q[") {
						if q == nil {
							q = make(map[bool]bool)
						}
						var keya bool
						{
							openIdx := strings.IndexRune(keyRaw, '[')
							closeIdx := strings.IndexRune(keyRaw, ']')
							keyaRaw := keyRaw[openIdx+1 : closeIdx]
							v, err2 := strconv.ParseBool(keyaRaw)
							if err2 != nil {
								err = goa.MergeErrors(err, goa.InvalidFieldTypeError("keya", keyaRaw, "boolean"))
							}
							keya = v
						}
						var vala bool
						{
							valaRaw := valRaw[0]
							v, err2 := strconv.ParseBool(valaRaw)
							if err2 != nil {
								err = goa.MergeErrors(err, goa.InvalidFieldTypeError("vala", valaRaw, "boolean"))
							}
							vala = v
						}
						q[keya] = vala
					}
				}
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodQueryMapBoolBoolPayload(q)

		return payload, nil
	}
}
`

var PayloadQueryMapBoolBoolValidateDecodeCode = `// DecodeMethodQueryMapBoolBoolValidateRequest returns a decoder for requests
// sent to the ServiceQueryMapBoolBoolValidate MethodQueryMapBoolBoolValidate
// endpoint.
func DecodeMethodQueryMapBoolBoolValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   map[bool]bool
			err error
		)
		{
			qRaw := r.URL.Query()
			if len(qRaw) == 0 {
				err = goa.MergeErrors(err, goa.MissingFieldError("q", "query string"))
			}
			for keyRaw, valRaw := range qRaw {
				if strings.HasPrefix(keyRaw, "q[") {
					if q == nil {
						q = make(map[bool]bool)
					}
					var keya bool
					{
						openIdx := strings.IndexRune(keyRaw, '[')
						closeIdx := strings.IndexRune(keyRaw, ']')
						keyaRaw := keyRaw[openIdx+1 : closeIdx]
						v, err2 := strconv.ParseBool(keyaRaw)
						if err2 != nil {
							err = goa.MergeErrors(err, goa.InvalidFieldTypeError("keya", keyaRaw, "boolean"))
						}
						keya = v
					}
					var vala bool
					{
						valaRaw := valRaw[0]
						v, err2 := strconv.ParseBool(valaRaw)
						if err2 != nil {
							err = goa.MergeErrors(err, goa.InvalidFieldTypeError("vala", valaRaw, "boolean"))
						}
						vala = v
					}
					q[keya] = vala
				}
			}
		}
		if len(q) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("q", q, len(q), 1, true))
		}
		for k, v := range q {
			if !(k == false) {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("q.key", k, []interface{}{false}))
			}
			if !(v == true) {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("q[key]", v, []interface{}{true}))
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodQueryMapBoolBoolValidatePayload(q)

		return payload, nil
	}
}
`

var PayloadQueryMapStringArrayStringDecodeCode = `// DecodeMethodQueryMapStringArrayStringRequest returns a decoder for requests
// sent to the ServiceQueryMapStringArrayString MethodQueryMapStringArrayString
// endpoint.
func DecodeMethodQueryMapStringArrayStringRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q map[string][]string
		)
		{
			qRaw := r.URL.Query()
			if len(qRaw) != 0 {
				for keyRaw, valRaw := range qRaw {
					if strings.HasPrefix(keyRaw, "q[") {
						if q == nil {
							q = make(map[string][]string)
						}
						var keya string
						{
							openIdx := strings.IndexRune(keyRaw, '[')
							closeIdx := strings.IndexRune(keyRaw, ']')
							keya = keyRaw[openIdx+1 : closeIdx]
						}
						q[keya] = valRaw
					}
				}
			}
		}
		payload := NewMethodQueryMapStringArrayStringPayload(q)

		return payload, nil
	}
}
`

var PayloadQueryMapStringArrayStringValidateDecodeCode = `// DecodeMethodQueryMapStringArrayStringValidateRequest returns a decoder for
// requests sent to the ServiceQueryMapStringArrayStringValidate
// MethodQueryMapStringArrayStringValidate endpoint.
func DecodeMethodQueryMapStringArrayStringValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   map[string][]string
			err error
		)
		{
			qRaw := r.URL.Query()
			if len(qRaw) == 0 {
				err = goa.MergeErrors(err, goa.MissingFieldError("q", "query string"))
			}
			for keyRaw, valRaw := range qRaw {
				if strings.HasPrefix(keyRaw, "q[") {
					if q == nil {
						q = make(map[string][]string)
					}
					var keya string
					{
						openIdx := strings.IndexRune(keyRaw, '[')
						closeIdx := strings.IndexRune(keyRaw, ']')
						keya = keyRaw[openIdx+1 : closeIdx]
					}
					q[keya] = valRaw
				}
			}
		}
		if len(q) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("q", q, len(q), 1, true))
		}
		for k, v := range q {
			if !(k == "key") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("q.key", k, []interface{}{"key"}))
			}
			if len(v) < 2 {
				err = goa.MergeErrors(err, goa.InvalidLengthError("q[key]", v, len(v), 2, true))
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodQueryMapStringArrayStringValidatePayload(q)

		return payload, nil
	}
}
`

var PayloadQueryMapStringArrayBoolDecodeCode = `// DecodeMethodQueryMapStringArrayBoolRequest returns a decoder for requests
// sent to the ServiceQueryMapStringArrayBool MethodQueryMapStringArrayBool
// endpoint.
func DecodeMethodQueryMapStringArrayBoolRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   map[string][]bool
			err error
		)
		{
			qRaw := r.URL.Query()
			if len(qRaw) != 0 {
				for keyRaw, valRaw := range qRaw {
					if strings.HasPrefix(keyRaw, "q[") {
						if q == nil {
							q = make(map[string][]bool)
						}
						var keya string
						{
							openIdx := strings.IndexRune(keyRaw, '[')
							closeIdx := strings.IndexRune(keyRaw, ']')
							keya = keyRaw[openIdx+1 : closeIdx]
						}
						var val []bool
						{
							val = make([]bool, len(valRaw))
							for i, rv := range valRaw {
								v, err2 := strconv.ParseBool(rv)
								if err2 != nil {
									err = goa.MergeErrors(err, goa.InvalidFieldTypeError("val", valRaw, "array of booleans"))
								}
								val[i] = v
							}
						}
						q[keya] = val
					}
				}
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodQueryMapStringArrayBoolPayload(q)

		return payload, nil
	}
}
`

var PayloadQueryMapStringArrayBoolValidateDecodeCode = `// DecodeMethodQueryMapStringArrayBoolValidateRequest returns a decoder for
// requests sent to the ServiceQueryMapStringArrayBoolValidate
// MethodQueryMapStringArrayBoolValidate endpoint.
func DecodeMethodQueryMapStringArrayBoolValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   map[string][]bool
			err error
		)
		{
			qRaw := r.URL.Query()
			if len(qRaw) == 0 {
				err = goa.MergeErrors(err, goa.MissingFieldError("q", "query string"))
			}
			for keyRaw, valRaw := range qRaw {
				if strings.HasPrefix(keyRaw, "q[") {
					if q == nil {
						q = make(map[string][]bool)
					}
					var keya string
					{
						openIdx := strings.IndexRune(keyRaw, '[')
						closeIdx := strings.IndexRune(keyRaw, ']')
						keya = keyRaw[openIdx+1 : closeIdx]
					}
					var val []bool
					{
						val = make([]bool, len(valRaw))
						for i, rv := range valRaw {
							v, err2 := strconv.ParseBool(rv)
							if err2 != nil {
								err = goa.MergeErrors(err, goa.InvalidFieldTypeError("val", valRaw, "array of booleans"))
							}
							val[i] = v
						}
					}
					q[keya] = val
				}
			}
		}
		if len(q) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("q", q, len(q), 1, true))
		}
		for k, v := range q {
			if !(k == "key") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("q.key", k, []interface{}{"key"}))
			}
			if len(v) < 2 {
				err = goa.MergeErrors(err, goa.InvalidLengthError("q[key]", v, len(v), 2, true))
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodQueryMapStringArrayBoolValidatePayload(q)

		return payload, nil
	}
}
`

var PayloadQueryMapBoolArrayStringDecodeCode = `// DecodeMethodQueryMapBoolArrayStringRequest returns a decoder for requests
// sent to the ServiceQueryMapBoolArrayString MethodQueryMapBoolArrayString
// endpoint.
func DecodeMethodQueryMapBoolArrayStringRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   map[bool][]string
			err error
		)
		{
			qRaw := r.URL.Query()
			if len(qRaw) != 0 {
				for keyRaw, valRaw := range qRaw {
					if strings.HasPrefix(keyRaw, "q[") {
						if q == nil {
							q = make(map[bool][]string)
						}
						var keya bool
						{
							openIdx := strings.IndexRune(keyRaw, '[')
							closeIdx := strings.IndexRune(keyRaw, ']')
							keyaRaw := keyRaw[openIdx+1 : closeIdx]
							v, err2 := strconv.ParseBool(keyaRaw)
							if err2 != nil {
								err = goa.MergeErrors(err, goa.InvalidFieldTypeError("keya", keyaRaw, "boolean"))
							}
							keya = v
						}
						q[keya] = valRaw
					}
				}
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodQueryMapBoolArrayStringPayload(q)

		return payload, nil
	}
}
`

var PayloadQueryMapBoolArrayStringValidateDecodeCode = `// DecodeMethodQueryMapBoolArrayStringValidateRequest returns a decoder for
// requests sent to the ServiceQueryMapBoolArrayStringValidate
// MethodQueryMapBoolArrayStringValidate endpoint.
func DecodeMethodQueryMapBoolArrayStringValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   map[bool][]string
			err error
		)
		{
			qRaw := r.URL.Query()
			if len(qRaw) != 0 {
				for keyRaw, valRaw := range qRaw {
					if strings.HasPrefix(keyRaw, "q[") {
						if q == nil {
							q = make(map[bool][]string)
						}
						var keya bool
						{
							openIdx := strings.IndexRune(keyRaw, '[')
							closeIdx := strings.IndexRune(keyRaw, ']')
							keyaRaw := keyRaw[openIdx+1 : closeIdx]
							v, err2 := strconv.ParseBool(keyaRaw)
							if err2 != nil {
								err = goa.MergeErrors(err, goa.InvalidFieldTypeError("keya", keyaRaw, "boolean"))
							}
							keya = v
						}
						q[keya] = valRaw
					}
				}
			}
		}
		if len(q) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("q", q, len(q), 1, true))
		}
		for k, v := range q {
			if !(k == true) {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("q.key", k, []interface{}{true}))
			}
			if len(v) < 2 {
				err = goa.MergeErrors(err, goa.InvalidLengthError("q[key]", v, len(v), 2, true))
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodQueryMapBoolArrayStringValidatePayload(q)

		return payload, nil
	}
}
`

var PayloadQueryMapBoolArrayBoolDecodeCode = `// DecodeMethodQueryMapBoolArrayBoolRequest returns a decoder for requests sent
// to the ServiceQueryMapBoolArrayBool MethodQueryMapBoolArrayBool endpoint.
func DecodeMethodQueryMapBoolArrayBoolRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   map[bool][]bool
			err error
		)
		{
			qRaw := r.URL.Query()
			if len(qRaw) != 0 {
				for keyRaw, valRaw := range qRaw {
					if strings.HasPrefix(keyRaw, "q[") {
						if q == nil {
							q = make(map[bool][]bool)
						}
						var keya bool
						{
							openIdx := strings.IndexRune(keyRaw, '[')
							closeIdx := strings.IndexRune(keyRaw, ']')
							keyaRaw := keyRaw[openIdx+1 : closeIdx]
							v, err2 := strconv.ParseBool(keyaRaw)
							if err2 != nil {
								err = goa.MergeErrors(err, goa.InvalidFieldTypeError("keya", keyaRaw, "boolean"))
							}
							keya = v
						}
						var val []bool
						{
							val = make([]bool, len(valRaw))
							for i, rv := range valRaw {
								v, err2 := strconv.ParseBool(rv)
								if err2 != nil {
									err = goa.MergeErrors(err, goa.InvalidFieldTypeError("val", valRaw, "array of booleans"))
								}
								val[i] = v
							}
						}
						q[keya] = val
					}
				}
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodQueryMapBoolArrayBoolPayload(q)

		return payload, nil
	}
}
`

var PayloadQueryMapBoolArrayBoolValidateDecodeCode = `// DecodeMethodQueryMapBoolArrayBoolValidateRequest returns a decoder for
// requests sent to the ServiceQueryMapBoolArrayBoolValidate
// MethodQueryMapBoolArrayBoolValidate endpoint.
func DecodeMethodQueryMapBoolArrayBoolValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   map[bool][]bool
			err error
		)
		{
			qRaw := r.URL.Query()
			if len(qRaw) == 0 {
				err = goa.MergeErrors(err, goa.MissingFieldError("q", "query string"))
			}
			for keyRaw, valRaw := range qRaw {
				if strings.HasPrefix(keyRaw, "q[") {
					if q == nil {
						q = make(map[bool][]bool)
					}
					var keya bool
					{
						openIdx := strings.IndexRune(keyRaw, '[')
						closeIdx := strings.IndexRune(keyRaw, ']')
						keyaRaw := keyRaw[openIdx+1 : closeIdx]
						v, err2 := strconv.ParseBool(keyaRaw)
						if err2 != nil {
							err = goa.MergeErrors(err, goa.InvalidFieldTypeError("keya", keyaRaw, "boolean"))
						}
						keya = v
					}
					var val []bool
					{
						val = make([]bool, len(valRaw))
						for i, rv := range valRaw {
							v, err2 := strconv.ParseBool(rv)
							if err2 != nil {
								err = goa.MergeErrors(err, goa.InvalidFieldTypeError("val", valRaw, "array of booleans"))
							}
							val[i] = v
						}
					}
					q[keya] = val
				}
			}
		}
		if len(q) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("q", q, len(q), 1, true))
		}
		for k, v := range q {
			if !(k == true) {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("q.key", k, []interface{}{true}))
			}
			if len(v) < 2 {
				err = goa.MergeErrors(err, goa.InvalidLengthError("q[key]", v, len(v), 2, true))
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodQueryMapBoolArrayBoolValidatePayload(q)

		return payload, nil
	}
}
`

var PayloadQueryPrimitiveStringValidateDecodeCode = `// DecodeMethodQueryPrimitiveStringValidateRequest returns a decoder for
// requests sent to the ServiceQueryPrimitiveStringValidate
// MethodQueryPrimitiveStringValidate endpoint.
func DecodeMethodQueryPrimitiveStringValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   string
			err error
		)
		q = r.URL.Query().Get("q")
		if q == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("q", "query string"))
		}
		if !(q == "val") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("q", q, []interface{}{"val"}))
		}
		if err != nil {
			return nil, err
		}
		payload := q

		return payload, nil
	}
}
`

var PayloadQueryPrimitiveBoolValidateDecodeCode = `// DecodeMethodQueryPrimitiveBoolValidateRequest returns a decoder for requests
// sent to the ServiceQueryPrimitiveBoolValidate
// MethodQueryPrimitiveBoolValidate endpoint.
func DecodeMethodQueryPrimitiveBoolValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   bool
			err error
		)
		{
			qRaw := r.URL.Query().Get("q")
			if qRaw == "" {
				err = goa.MergeErrors(err, goa.MissingFieldError("q", "query string"))
			}
			v, err2 := strconv.ParseBool(qRaw)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("q", qRaw, "boolean"))
			}
			q = v
		}
		if !(q == true) {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("q", q, []interface{}{true}))
		}
		if err != nil {
			return nil, err
		}
		payload := q

		return payload, nil
	}
}
`

var PayloadQueryPrimitiveArrayStringValidateDecodeCode = `// DecodeMethodQueryPrimitiveArrayStringValidateRequest returns a decoder for
// requests sent to the ServiceQueryPrimitiveArrayStringValidate
// MethodQueryPrimitiveArrayStringValidate endpoint.
func DecodeMethodQueryPrimitiveArrayStringValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   []string
			err error
		)
		q = r.URL.Query()["q"]
		if q == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("q", "query string"))
		}
		if len(q) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("q", q, len(q), 1, true))
		}
		for _, e := range q {
			if !(e == "val") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("q[*]", e, []interface{}{"val"}))
			}
		}
		if err != nil {
			return nil, err
		}
		payload := q

		return payload, nil
	}
}
`

var PayloadQueryPrimitiveArrayBoolValidateDecodeCode = `// DecodeMethodQueryPrimitiveArrayBoolValidateRequest returns a decoder for
// requests sent to the ServiceQueryPrimitiveArrayBoolValidate
// MethodQueryPrimitiveArrayBoolValidate endpoint.
func DecodeMethodQueryPrimitiveArrayBoolValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   []bool
			err error
		)
		{
			qRaw := r.URL.Query()["q"]
			if qRaw == nil {
				return goa.MergeErrors(err, goa.MissingFieldError("q", "query string"))
			}
			q = make([]bool, len(qRaw))
			for i, rv := range qRaw {
				v, err2 := strconv.ParseBool(rv)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("q", qRaw, "array of booleans"))
				}
				q[i] = v
			}
		}
		if len(q) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("q", q, len(q), 1, true))
		}
		for _, e := range q {
			if !(e == true) {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("q[*]", e, []interface{}{true}))
			}
		}
		if err != nil {
			return nil, err
		}
		payload := q

		return payload, nil
	}
}
`

var PayloadQueryPrimitiveMapStringArrayStringValidateDecodeCode = `// DecodeMethodQueryPrimitiveMapStringArrayStringValidateRequest returns a
// decoder for requests sent to the
// ServiceQueryPrimitiveMapStringArrayStringValidate
// MethodQueryPrimitiveMapStringArrayStringValidate endpoint.
func DecodeMethodQueryPrimitiveMapStringArrayStringValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   map[string][]string
			err error
		)
		{
			qRaw := r.URL.Query()
			if len(qRaw) == 0 {
				err = goa.MergeErrors(err, goa.MissingFieldError("q", "query string"))
			}
			for keyRaw, valRaw := range qRaw {
				if strings.HasPrefix(keyRaw, "q[") {
					if q == nil {
						q = make(map[string][]string)
					}
					var keya string
					{
						openIdx := strings.IndexRune(keyRaw, '[')
						closeIdx := strings.IndexRune(keyRaw, ']')
						keya = keyRaw[openIdx+1 : closeIdx]
					}
					q[keya] = valRaw
				}
			}
		}
		if len(q) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("q", q, len(q), 1, true))
		}
		for k, v := range q {
			err = goa.MergeErrors(err, goa.ValidatePattern("q.key", k, "key"))
			if len(v) < 2 {
				err = goa.MergeErrors(err, goa.InvalidLengthError("q[key]", v, len(v), 2, true))
			}
			for _, e := range v {
				err = goa.MergeErrors(err, goa.ValidatePattern("q[key][*]", e, "val"))
			}
		}
		if err != nil {
			return nil, err
		}
		payload := q

		return payload, nil
	}
}
`

var PayloadQueryPrimitiveMapStringBoolValidateDecodeCode = `// DecodeMethodQueryPrimitiveMapStringBoolValidateRequest returns a decoder for
// requests sent to the ServiceQueryPrimitiveMapStringBoolValidate
// MethodQueryPrimitiveMapStringBoolValidate endpoint.
func DecodeMethodQueryPrimitiveMapStringBoolValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   map[string]bool
			err error
		)
		{
			qRaw := r.URL.Query()
			if len(qRaw) == 0 {
				err = goa.MergeErrors(err, goa.MissingFieldError("q", "query string"))
			}
			for keyRaw, valRaw := range qRaw {
				if strings.HasPrefix(keyRaw, "q[") {
					if q == nil {
						q = make(map[string]bool)
					}
					var keya string
					{
						openIdx := strings.IndexRune(keyRaw, '[')
						closeIdx := strings.IndexRune(keyRaw, ']')
						keya = keyRaw[openIdx+1 : closeIdx]
					}
					var vala bool
					{
						valaRaw := valRaw[0]
						v, err2 := strconv.ParseBool(valaRaw)
						if err2 != nil {
							err = goa.MergeErrors(err, goa.InvalidFieldTypeError("vala", valaRaw, "boolean"))
						}
						vala = v
					}
					q[keya] = vala
				}
			}
		}
		if len(q) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("q", q, len(q), 1, true))
		}
		for k, v := range q {
			err = goa.MergeErrors(err, goa.ValidatePattern("q.key", k, "key"))
			if !(v == true) {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("q[key]", v, []interface{}{true}))
			}
		}
		if err != nil {
			return nil, err
		}
		payload := q

		return payload, nil
	}
}
`

var PayloadQueryPrimitiveMapBoolArrayBoolValidateDecodeCode = `// DecodeMethodQueryPrimitiveMapBoolArrayBoolValidateRequest returns a decoder
// for requests sent to the ServiceQueryPrimitiveMapBoolArrayBoolValidate
// MethodQueryPrimitiveMapBoolArrayBoolValidate endpoint.
func DecodeMethodQueryPrimitiveMapBoolArrayBoolValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   map[bool][]bool
			err error
		)
		{
			qRaw := r.URL.Query()
			if len(qRaw) == 0 {
				err = goa.MergeErrors(err, goa.MissingFieldError("q", "query string"))
			}
			for keyRaw, valRaw := range qRaw {
				if strings.HasPrefix(keyRaw, "q[") {
					if q == nil {
						q = make(map[bool][]bool)
					}
					var keya bool
					{
						openIdx := strings.IndexRune(keyRaw, '[')
						closeIdx := strings.IndexRune(keyRaw, ']')
						keyaRaw := keyRaw[openIdx+1 : closeIdx]
						v, err2 := strconv.ParseBool(keyaRaw)
						if err2 != nil {
							err = goa.MergeErrors(err, goa.InvalidFieldTypeError("keya", keyaRaw, "boolean"))
						}
						keya = v
					}
					var val []bool
					{
						val = make([]bool, len(valRaw))
						for i, rv := range valRaw {
							v, err2 := strconv.ParseBool(rv)
							if err2 != nil {
								err = goa.MergeErrors(err, goa.InvalidFieldTypeError("val", valRaw, "array of booleans"))
							}
							val[i] = v
						}
					}
					q[keya] = val
				}
			}
		}
		if len(q) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("q", q, len(q), 1, true))
		}
		for k, v := range q {
			if !(k == true) {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("q.key", k, []interface{}{true}))
			}
			if len(v) < 2 {
				err = goa.MergeErrors(err, goa.InvalidLengthError("q[key]", v, len(v), 2, true))
			}
			for _, e := range v {
				if !(e == false) {
					err = goa.MergeErrors(err, goa.InvalidEnumValueError("q[key][*]", e, []interface{}{false}))
				}
			}
		}
		if err != nil {
			return nil, err
		}
		payload := q

		return payload, nil
	}
}
`

var PayloadQueryMapStringMapIntStringValidateDecodeCode = `// DecodeMethodQueryMapStringMapIntStringValidateRequest returns a decoder for
// requests sent to the ServiceQueryMapStringMapIntStringValidate
// MethodQueryMapStringMapIntStringValidate endpoint.
func DecodeMethodQueryMapStringMapIntStringValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   map[string]map[int]string
			err error
		)
		{
			qRaw := r.URL.Query()
			if len(qRaw) == 0 {
				err = goa.MergeErrors(err, goa.MissingFieldError("q", "query string"))
			}
			for keyRaw, valRaw := range qRaw {
				if strings.HasPrefix(keyRaw, "q[") {
					if q == nil {
						q = make(map[string]map[int]string)
					}
					var keya string
					{
						openIdx := strings.IndexRune(keyRaw, '[')
						closeIdx := strings.IndexRune(keyRaw, ']')
						keya = keyRaw[openIdx+1 : closeIdx]
						keyRaw = keyRaw[closeIdx+1:]
					}
					if q[keya] == nil {
						q[keya] = make(map[int]string)
					}
					var keyb int
					{
						openIdx := strings.IndexRune(keyRaw, '[')
						closeIdx := strings.IndexRune(keyRaw, ']')
						keybRaw := keyRaw[openIdx+1 : closeIdx]
						v, err2 := strconv.ParseInt(keybRaw, 10, strconv.IntSize)
						if err2 != nil {
							err = goa.MergeErrors(err, goa.InvalidFieldTypeError("keyb", keybRaw, "integer"))
						}
						keyb = int(v)
					}
					q[keya][keyb] = valRaw[0]
				}
			}
		}
		for k, _ := range q {
			if !(k == "foo") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("q.key", k, []interface{}{"foo"}))
			}
		}
		if err != nil {
			return nil, err
		}
		payload := q

		return payload, nil
	}
}
`

var PayloadQueryMapIntMapStringArrayIntValidateDecodeCode = `// DecodeMethodQueryMapIntMapStringArrayIntValidateRequest returns a decoder
// for requests sent to the ServiceQueryMapIntMapStringArrayIntValidate
// MethodQueryMapIntMapStringArrayIntValidate endpoint.
func DecodeMethodQueryMapIntMapStringArrayIntValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   map[int]map[string][]int
			err error
		)
		{
			qRaw := r.URL.Query()
			if len(qRaw) == 0 {
				err = goa.MergeErrors(err, goa.MissingFieldError("q", "query string"))
			}
			for keyRaw, valRaw := range qRaw {
				if strings.HasPrefix(keyRaw, "q[") {
					if q == nil {
						q = make(map[int]map[string][]int)
					}
					var keya int
					{
						openIdx := strings.IndexRune(keyRaw, '[')
						closeIdx := strings.IndexRune(keyRaw, ']')
						keyaRaw := keyRaw[openIdx+1 : closeIdx]
						v, err2 := strconv.ParseInt(keyaRaw, 10, strconv.IntSize)
						if err2 != nil {
							err = goa.MergeErrors(err, goa.InvalidFieldTypeError("keya", keyaRaw, "integer"))
						}
						keya = int(v)
						keyRaw = keyRaw[closeIdx+1:]
					}
					if q[keya] == nil {
						q[keya] = make(map[string][]int)
					}
					var keyb string
					{
						openIdx := strings.IndexRune(keyRaw, '[')
						closeIdx := strings.IndexRune(keyRaw, ']')
						keyb = keyRaw[openIdx+1 : closeIdx]
					}
					var val []int
					{
						val = make([]int, len(valRaw))
						for i, rv := range valRaw {
							v, err2 := strconv.ParseInt(rv, 10, strconv.IntSize)
							if err2 != nil {
								err = goa.MergeErrors(err, goa.InvalidFieldTypeError("val", valRaw, "array of integers"))
							}
							val[i] = int(v)
						}
					}
					q[keya][keyb] = val
				}
			}
		}
		for k, _ := range q {
			if !(k == 1 || k == 2 || k == 3) {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("q.key", k, []interface{}{1, 2, 3}))
			}
		}
		if err != nil {
			return nil, err
		}
		payload := q

		return payload, nil
	}
}
`

var PayloadQueryStringMappedDecodeCode = `// DecodeMethodQueryStringMappedRequest returns a decoder for requests sent to
// the ServiceQueryStringMapped MethodQueryStringMapped endpoint.
func DecodeMethodQueryStringMappedRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			query *string
		)
		queryRaw := r.URL.Query().Get("q")
		if queryRaw != "" {
			query = &queryRaw
		}
		payload := NewMethodQueryStringMappedPayload(query)

		return payload, nil
	}
}
`

var PayloadQueryStringDefaultDecodeCode = `// DecodeMethodQueryStringDefaultRequest returns a decoder for requests sent to
// the ServiceQueryStringDefault MethodQueryStringDefault endpoint.
func DecodeMethodQueryStringDefaultRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q string
		)
		qRaw := r.URL.Query().Get("q")
		if qRaw != "" {
			q = qRaw
		} else {
			q = "def"
		}
		payload := NewMethodQueryStringDefaultPayload(q)

		return payload, nil
	}
}
`

var PayloadQueryStringDefaultValidateDecodeCode = `// DecodeMethodQueryStringDefaultValidateRequest returns a decoder for requests
// sent to the ServiceQueryStringDefaultValidate
// MethodQueryStringDefaultValidate endpoint.
func DecodeMethodQueryStringDefaultValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   string
			err error
		)
		qRaw := r.URL.Query().Get("q")
		if qRaw != "" {
			q = qRaw
		} else {
			q = "def"
		}
		if !(q == "def") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("q", q, []interface{}{"def"}))
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodQueryStringDefaultValidatePayload(q)

		return payload, nil
	}
}
`

var PayloadQueryPrimitiveStringDefaultDecodeCode = `// DecodeMethodQueryPrimitiveStringDefaultRequest returns a decoder for
// requests sent to the ServiceQueryPrimitiveStringDefault
// MethodQueryPrimitiveStringDefault endpoint.
func DecodeMethodQueryPrimitiveStringDefaultRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q   string
			err error
		)
		q = r.URL.Query().Get("q")
		if q == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("q", "query string"))
		}
		if err != nil {
			return nil, err
		}
		payload := q

		return payload, nil
	}
}
`

var PayloadExtendedQueryStringDecodeCode = `// DecodeMethodQueryStringExtendedPayloadRequest returns a decoder for requests
// sent to the ServiceQueryStringExtendedPayload
// MethodQueryStringExtendedPayload endpoint.
func DecodeMethodQueryStringExtendedPayloadRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			q *string
		)
		qRaw := r.URL.Query().Get("q")
		if qRaw != "" {
			q = &qRaw
		}
		payload := NewMethodQueryStringExtendedPayloadPayload(q)

		return payload, nil
	}
}
`

var PayloadPathStringDecodeCode = `// DecodeMethodPathStringRequest returns a decoder for requests sent to the
// ServicePathString MethodPathString endpoint.
func DecodeMethodPathStringRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			p string

			params = mux.Vars(r)
		)
		p = params["p"]
		payload := NewMethodPathStringPayload(p)

		return payload, nil
	}
}
`

var PayloadPathStringValidateDecodeCode = `// DecodeMethodPathStringValidateRequest returns a decoder for requests sent to
// the ServicePathStringValidate MethodPathStringValidate endpoint.
func DecodeMethodPathStringValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			p   string
			err error

			params = mux.Vars(r)
		)
		p = params["p"]
		if !(p == "val") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("p", p, []interface{}{"val"}))
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodPathStringValidatePayload(p)

		return payload, nil
	}
}
`

var PayloadPathArrayStringDecodeCode = `// DecodeMethodPathArrayStringRequest returns a decoder for requests sent to
// the ServicePathArrayString MethodPathArrayString endpoint.
func DecodeMethodPathArrayStringRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			p []string

			params = mux.Vars(r)
		)
		{
			pRaw := params["p"]
			pRawSlice := strings.Split(pRaw, ",")
			p = make([]string, len(pRawSlice))
			for i, rv := range pRawSlice {
				p[i] = rv
			}
		}
		payload := NewMethodPathArrayStringPayload(p)

		return payload, nil
	}
}
`

var PayloadPathArrayStringValidateDecodeCode = `// DecodeMethodPathArrayStringValidateRequest returns a decoder for requests
// sent to the ServicePathArrayStringValidate MethodPathArrayStringValidate
// endpoint.
func DecodeMethodPathArrayStringValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			p   []string
			err error

			params = mux.Vars(r)
		)
		{
			pRaw := params["p"]
			pRawSlice := strings.Split(pRaw, ",")
			p = make([]string, len(pRawSlice))
			for i, rv := range pRawSlice {
				p[i] = rv
			}
		}
		if !(p == []string{"val"}) {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("p", p, []interface{}{[]string{"val"}}))
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodPathArrayStringValidatePayload(p)

		return payload, nil
	}
}
`

var PayloadPathPrimitiveStringValidateDecodeCode = `// DecodeMethodPathPrimitiveStringValidateRequest returns a decoder for
// requests sent to the ServicePathPrimitiveStringValidate
// MethodPathPrimitiveStringValidate endpoint.
func DecodeMethodPathPrimitiveStringValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			p   string
			err error

			params = mux.Vars(r)
		)
		p = params["p"]
		if !(p == "val") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("p", p, []interface{}{"val"}))
		}
		if err != nil {
			return nil, err
		}
		payload := p

		return payload, nil
	}
}
`

var PayloadPathPrimitiveBoolValidateDecodeCode = `// DecodeMethodPathPrimitiveBoolValidateRequest returns a decoder for requests
// sent to the ServicePathPrimitiveBoolValidate MethodPathPrimitiveBoolValidate
// endpoint.
func DecodeMethodPathPrimitiveBoolValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			p   bool
			err error

			params = mux.Vars(r)
		)
		{
			pRaw := params["p"]
			v, err2 := strconv.ParseBool(pRaw)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("p", pRaw, "boolean"))
			}
			p = v
		}
		if !(p == true) {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("p", p, []interface{}{true}))
		}
		if err != nil {
			return nil, err
		}
		payload := p

		return payload, nil
	}
}
`

var PayloadPathPrimitiveArrayStringValidateDecodeCode = `// DecodeMethodPathPrimitiveArrayStringValidateRequest returns a decoder for
// requests sent to the ServicePathPrimitiveArrayStringValidate
// MethodPathPrimitiveArrayStringValidate endpoint.
func DecodeMethodPathPrimitiveArrayStringValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			p   []string
			err error

			params = mux.Vars(r)
		)
		{
			pRaw := params["p"]
			pRawSlice := strings.Split(pRaw, ",")
			p = make([]string, len(pRawSlice))
			for i, rv := range pRawSlice {
				p[i] = rv
			}
		}
		if len(p) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("p", p, len(p), 1, true))
		}
		for _, e := range p {
			if !(e == "val") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("p[*]", e, []interface{}{"val"}))
			}
		}
		if err != nil {
			return nil, err
		}
		payload := p

		return payload, nil
	}
}
`

var PayloadPathPrimitiveArrayBoolValidateDecodeCode = `// DecodeMethodPathPrimitiveArrayBoolValidateRequest returns a decoder for
// requests sent to the ServicePathPrimitiveArrayBoolValidate
// MethodPathPrimitiveArrayBoolValidate endpoint.
func DecodeMethodPathPrimitiveArrayBoolValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			p   []bool
			err error

			params = mux.Vars(r)
		)
		{
			pRaw := params["p"]
			pRawSlice := strings.Split(pRaw, ",")
			p = make([]bool, len(pRawSlice))
			for i, rv := range pRawSlice {
				v, err2 := strconv.ParseBool(rv)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("p", pRaw, "array of booleans"))
				}
				p[i] = v
			}
		}
		if len(p) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("p", p, len(p), 1, true))
		}
		for _, e := range p {
			if !(e == true) {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("p[*]", e, []interface{}{true}))
			}
		}
		if err != nil {
			return nil, err
		}
		payload := p

		return payload, nil
	}
}
`

var PayloadHeaderStringDecodeCode = `// DecodeMethodHeaderStringRequest returns a decoder for requests sent to the
// ServiceHeaderString MethodHeaderString endpoint.
func DecodeMethodHeaderStringRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			h *string
		)
		hRaw := r.Header.Get("h")
		if hRaw != "" {
			h = &hRaw
		}
		payload := NewMethodHeaderStringPayload(h)

		return payload, nil
	}
}
`

var PayloadHeaderStringValidateDecodeCode = `// DecodeMethodHeaderStringValidateRequest returns a decoder for requests sent
// to the ServiceHeaderStringValidate MethodHeaderStringValidate endpoint.
func DecodeMethodHeaderStringValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			h   *string
			err error
		)
		hRaw := r.Header.Get("h")
		if hRaw != "" {
			h = &hRaw
		}
		if h != nil {
			err = goa.MergeErrors(err, goa.ValidatePattern("h", *h, "header"))
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodHeaderStringValidatePayload(h)

		return payload, nil
	}
}
`

var PayloadHeaderArrayStringDecodeCode = `// DecodeMethodHeaderArrayStringRequest returns a decoder for requests sent to
// the ServiceHeaderArrayString MethodHeaderArrayString endpoint.
func DecodeMethodHeaderArrayStringRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			h []string
		)
		h = r.Header["H"]
		payload := NewMethodHeaderArrayStringPayload(h)

		return payload, nil
	}
}
`

var PayloadHeaderArrayStringValidateDecodeCode = `// DecodeMethodHeaderArrayStringValidateRequest returns a decoder for requests
// sent to the ServiceHeaderArrayStringValidate MethodHeaderArrayStringValidate
// endpoint.
func DecodeMethodHeaderArrayStringValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			h   []string
			err error
		)
		h = r.Header["H"]
		for _, e := range h {
			if !(e == "val") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("h[*]", e, []interface{}{"val"}))
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodHeaderArrayStringValidatePayload(h)

		return payload, nil
	}
}
`

var PayloadHeaderPrimitiveStringValidateDecodeCode = `// DecodeMethodHeaderPrimitiveStringValidateRequest returns a decoder for
// requests sent to the ServiceHeaderPrimitiveStringValidate
// MethodHeaderPrimitiveStringValidate endpoint.
func DecodeMethodHeaderPrimitiveStringValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			h   string
			err error
		)
		h = r.Header.Get("h")
		if h == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("h", "header"))
		}
		if !(h == "val") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("h", h, []interface{}{"val"}))
		}
		if err != nil {
			return nil, err
		}
		payload := h

		return payload, nil
	}
}
`

var PayloadHeaderPrimitiveBoolValidateDecodeCode = `// DecodeMethodHeaderPrimitiveBoolValidateRequest returns a decoder for
// requests sent to the ServiceHeaderPrimitiveBoolValidate
// MethodHeaderPrimitiveBoolValidate endpoint.
func DecodeMethodHeaderPrimitiveBoolValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			h   bool
			err error
		)
		{
			hRaw := r.Header.Get("h")
			if hRaw == "" {
				err = goa.MergeErrors(err, goa.MissingFieldError("h", "header"))
			}
			v, err2 := strconv.ParseBool(hRaw)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("h", hRaw, "boolean"))
			}
			h = v
		}
		if !(h == true) {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("h", h, []interface{}{true}))
		}
		if err != nil {
			return nil, err
		}
		payload := h

		return payload, nil
	}
}
`

var PayloadHeaderPrimitiveArrayStringValidateDecodeCode = `// DecodeMethodHeaderPrimitiveArrayStringValidateRequest returns a decoder for
// requests sent to the ServiceHeaderPrimitiveArrayStringValidate
// MethodHeaderPrimitiveArrayStringValidate endpoint.
func DecodeMethodHeaderPrimitiveArrayStringValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			h   []string
			err error
		)
		h = r.Header["H"]
		if h == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("h", "header"))
		}
		if len(h) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("h", h, len(h), 1, true))
		}
		for _, e := range h {
			err = goa.MergeErrors(err, goa.ValidatePattern("h[*]", e, "val"))
		}
		if err != nil {
			return nil, err
		}
		payload := h

		return payload, nil
	}
}
`

var PayloadHeaderPrimitiveArrayBoolValidateDecodeCode = `// DecodeMethodHeaderPrimitiveArrayBoolValidateRequest returns a decoder for
// requests sent to the ServiceHeaderPrimitiveArrayBoolValidate
// MethodHeaderPrimitiveArrayBoolValidate endpoint.
func DecodeMethodHeaderPrimitiveArrayBoolValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			h   []bool
			err error
		)
		{
			hRaw := r.Header["H"]
			if hRaw == nil {
				err = goa.MergeErrors(err, goa.MissingFieldError("h", "header"))
			}
			h = make([]bool, len(hRaw))
			for i, rv := range hRaw {
				v, err2 := strconv.ParseBool(rv)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("h", hRaw, "array of booleans"))
				}
				h[i] = v
			}
		}
		if len(h) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("h", h, len(h), 1, true))
		}
		for _, e := range h {
			if !(e == true) {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("h[*]", e, []interface{}{true}))
			}
		}
		if err != nil {
			return nil, err
		}
		payload := h

		return payload, nil
	}
}
`

var PayloadHeaderStringDefaultDecodeCode = `// DecodeMethodHeaderStringDefaultRequest returns a decoder for requests sent
// to the ServiceHeaderStringDefault MethodHeaderStringDefault endpoint.
func DecodeMethodHeaderStringDefaultRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			h string
		)
		hRaw := r.Header.Get("h")
		if hRaw != "" {
			h = hRaw
		} else {
			h = "def"
		}
		payload := NewMethodHeaderStringDefaultPayload(h)

		return payload, nil
	}
}
`

var PayloadHeaderStringDefaultValidateDecodeCode = `// DecodeMethodHeaderStringDefaultValidateRequest returns a decoder for
// requests sent to the ServiceHeaderStringDefaultValidate
// MethodHeaderStringDefaultValidate endpoint.
func DecodeMethodHeaderStringDefaultValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			h   string
			err error
		)
		hRaw := r.Header.Get("h")
		if hRaw != "" {
			h = hRaw
		} else {
			h = "def"
		}
		if !(h == "def") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("h", h, []interface{}{"def"}))
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodHeaderStringDefaultValidatePayload(h)

		return payload, nil
	}
}
`

var PayloadHeaderPrimitiveStringDefaultDecodeCode = `// DecodeMethodHeaderPrimitiveStringDefaultRequest returns a decoder for
// requests sent to the ServiceHeaderPrimitiveStringDefault
// MethodHeaderPrimitiveStringDefault endpoint.
func DecodeMethodHeaderPrimitiveStringDefaultRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			h   string
			err error
		)
		h = r.Header.Get("h")
		if h == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("h", "header"))
		}
		if err != nil {
			return nil, err
		}
		payload := h

		return payload, nil
	}
}
`

var PayloadBodyStringDecodeCode = `// DecodeMethodBodyStringRequest returns a decoder for requests sent to the
// ServiceBodyString MethodBodyString endpoint.
func DecodeMethodBodyStringRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body MethodBodyStringRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		payload := NewMethodBodyStringPayload(&body)

		return payload, nil
	}
}
`

var PayloadBodyStringValidateDecodeCode = `// DecodeMethodBodyStringValidateRequest returns a decoder for requests sent to
// the ServiceBodyStringValidate MethodBodyStringValidate endpoint.
func DecodeMethodBodyStringValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body MethodBodyStringValidateRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateMethodBodyStringValidateRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewMethodBodyStringValidatePayload(&body)

		return payload, nil
	}
}
`

var PayloadBodyUserDecodeCode = `// DecodeMethodBodyUserRequest returns a decoder for requests sent to the
// ServiceBodyUser MethodBodyUser endpoint.
func DecodeMethodBodyUserRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body MethodBodyUserRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		payload := NewMethodBodyUserPayloadType(&body)

		return payload, nil
	}
}
`

var PayloadBodyUserRequiredDecodeCode = `// DecodeMethodBodyUserRequest returns a decoder for requests sent to the
// ServiceBodyUser MethodBodyUser endpoint.
func DecodeMethodBodyUserRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body MethodBodyUserRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateMethodBodyUserRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewMethodBodyUserPayloadType(&body)

		return payload, nil
	}
}
`

var PayloadBodyNestedUserDecodeCode = `// DecodeMethodBodyUserRequest returns a decoder for requests sent to the
// ServiceBodyUser MethodBodyUser endpoint.
func DecodeMethodBodyUserRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body MethodBodyUserRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateMethodBodyUserRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewMethodBodyUserPayloadType(&body)

		return payload, nil
	}
}
`

var PayloadBodyUserValidateDecodeCode = `// DecodeMethodBodyUserValidateRequest returns a decoder for requests sent to
// the ServiceBodyUserValidate MethodBodyUserValidate endpoint.
func DecodeMethodBodyUserValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body MethodBodyUserValidateRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateMethodBodyUserValidateRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewMethodBodyUserValidatePayloadType(&body)

		return payload, nil
	}
}
`

var PayloadBodyArrayStringDecodeCode = `// DecodeMethodBodyArrayStringRequest returns a decoder for requests sent to
// the ServiceBodyArrayString MethodBodyArrayString endpoint.
func DecodeMethodBodyArrayStringRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body MethodBodyArrayStringRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		payload := NewMethodBodyArrayStringPayload(&body)

		return payload, nil
	}
}
`

var PayloadBodyArrayStringValidateDecodeCode = `// DecodeMethodBodyArrayStringValidateRequest returns a decoder for requests
// sent to the ServiceBodyArrayStringValidate MethodBodyArrayStringValidate
// endpoint.
func DecodeMethodBodyArrayStringValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body MethodBodyArrayStringValidateRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateMethodBodyArrayStringValidateRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewMethodBodyArrayStringValidatePayload(&body)

		return payload, nil
	}
}
`

var PayloadBodyArrayUserDecodeCode = `// DecodeMethodBodyArrayUserRequest returns a decoder for requests sent to the
// ServiceBodyArrayUser MethodBodyArrayUser endpoint.
func DecodeMethodBodyArrayUserRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body MethodBodyArrayUserRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateMethodBodyArrayUserRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewMethodBodyArrayUserPayload(&body)

		return payload, nil
	}
}
`

var PayloadBodyArrayUserValidateDecodeCode = `// DecodeMethodBodyArrayUserValidateRequest returns a decoder for requests sent
// to the ServiceBodyArrayUserValidate MethodBodyArrayUserValidate endpoint.
func DecodeMethodBodyArrayUserValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body MethodBodyArrayUserValidateRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateMethodBodyArrayUserValidateRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewMethodBodyArrayUserValidatePayload(&body)

		return payload, nil
	}
}
`

var PayloadBodyMapStringDecodeCode = `// DecodeMethodBodyMapStringRequest returns a decoder for requests sent to the
// ServiceBodyMapString MethodBodyMapString endpoint.
func DecodeMethodBodyMapStringRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body MethodBodyMapStringRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		payload := NewMethodBodyMapStringPayload(&body)

		return payload, nil
	}
}
`

var PayloadBodyMapStringValidateDecodeCode = `// DecodeMethodBodyMapStringValidateRequest returns a decoder for requests sent
// to the ServiceBodyMapStringValidate MethodBodyMapStringValidate endpoint.
func DecodeMethodBodyMapStringValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body MethodBodyMapStringValidateRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateMethodBodyMapStringValidateRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewMethodBodyMapStringValidatePayload(&body)

		return payload, nil
	}
}
`

var PayloadBodyMapUserDecodeCode = `// DecodeMethodBodyMapUserRequest returns a decoder for requests sent to the
// ServiceBodyMapUser MethodBodyMapUser endpoint.
func DecodeMethodBodyMapUserRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body MethodBodyMapUserRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateMethodBodyMapUserRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewMethodBodyMapUserPayload(&body)

		return payload, nil
	}
}
`

var PayloadBodyMapUserValidateDecodeCode = `// DecodeMethodBodyMapUserValidateRequest returns a decoder for requests sent
// to the ServiceBodyMapUserValidate MethodBodyMapUserValidate endpoint.
func DecodeMethodBodyMapUserValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body MethodBodyMapUserValidateRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateMethodBodyMapUserValidateRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewMethodBodyMapUserValidatePayload(&body)

		return payload, nil
	}
}
`

var PayloadBodyPrimitiveStringValidateDecodeCode = `// DecodeMethodBodyPrimitiveStringValidateRequest returns a decoder for
// requests sent to the ServiceBodyPrimitiveStringValidate
// MethodBodyPrimitiveStringValidate endpoint.
func DecodeMethodBodyPrimitiveStringValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body string
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		if body != nil {
			if !(*body == "val") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("body", *body, []interface{}{"val"}))
			}
		}
		if err != nil {
			return nil, err
		}
		payload := body

		return payload, nil
	}
}
`

var PayloadBodyPrimitiveBoolValidateDecodeCode = `// DecodeMethodBodyPrimitiveBoolValidateRequest returns a decoder for requests
// sent to the ServiceBodyPrimitiveBoolValidate MethodBodyPrimitiveBoolValidate
// endpoint.
func DecodeMethodBodyPrimitiveBoolValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body bool
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		if body != nil {
			if !(*body == true) {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("body", *body, []interface{}{true}))
			}
		}
		if err != nil {
			return nil, err
		}
		payload := body

		return payload, nil
	}
}
`

var PayloadBodyPrimitiveArrayStringValidateDecodeCode = `// DecodeMethodBodyPrimitiveArrayStringValidateRequest returns a decoder for
// requests sent to the ServiceBodyPrimitiveArrayStringValidate
// MethodBodyPrimitiveArrayStringValidate endpoint.
func DecodeMethodBodyPrimitiveArrayStringValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body []string
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		if len(body) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body", body, len(body), 1, true))
		}
		for _, e := range body {
			if !(e == "val") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("body[*]", e, []interface{}{"val"}))
			}
		}
		if err != nil {
			return nil, err
		}
		payload := body

		return payload, nil
	}
}
`

var PayloadBodyPrimitiveArrayBoolValidateDecodeCode = `// DecodeMethodBodyPrimitiveArrayBoolValidateRequest returns a decoder for
// requests sent to the ServiceBodyPrimitiveArrayBoolValidate
// MethodBodyPrimitiveArrayBoolValidate endpoint.
func DecodeMethodBodyPrimitiveArrayBoolValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body []bool
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		if len(body) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body", body, len(body), 1, true))
		}
		for _, e := range body {
			if !(e == true) {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("body[*]", e, []interface{}{true}))
			}
		}
		if err != nil {
			return nil, err
		}
		payload := body

		return payload, nil
	}
}
`

var PayloadBodyPrimitiveArrayUserValidateDecodeCode = `// DecodeMethodBodyPrimitiveArrayUserValidateRequest returns a decoder for
// requests sent to the ServiceBodyPrimitiveArrayUserValidate
// MethodBodyPrimitiveArrayUserValidate endpoint.
func DecodeMethodBodyPrimitiveArrayUserValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body []*PayloadTypeRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		if len(body) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body", body, len(body), 1, true))
		}
		for _, e := range body {
			if e != nil {
				if err2 := ValidatePayloadTypeRequestBody(e); err2 != nil {
					err = goa.MergeErrors(err, err2)
				}
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodBodyPrimitiveArrayUserValidatePayloadType(body)

		return payload, nil
	}
}
`

var PayloadBodyPrimitiveFieldArrayUserDecodeCode = `// DecodeMethodBodyPrimitiveArrayUserRequest returns a decoder for requests
// sent to the ServiceBodyPrimitiveArrayUser MethodBodyPrimitiveArrayUser
// endpoint.
func DecodeMethodBodyPrimitiveArrayUserRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body []string
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		payload := NewMethodBodyPrimitiveArrayUserPayloadType(body)

		return payload, nil
	}
}
`

var PayloadBodyPrimitiveFieldStringDecodeCode = `// DecodeMethodBodyPrimitiveArrayUserRequest returns a decoder for requests
// sent to the ServiceBodyPrimitiveArrayUser MethodBodyPrimitiveArrayUser
// endpoint.
func DecodeMethodBodyPrimitiveArrayUserRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body string
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		payload := NewMethodBodyPrimitiveArrayUserPayloadType(body)

		return payload, nil
	}
}
`

var PayloadBodyPrimitiveFieldArrayUserValidateDecodeCode = `// DecodeMethodBodyPrimitiveArrayUserValidateRequest returns a decoder for
// requests sent to the ServiceBodyPrimitiveArrayUserValidate
// MethodBodyPrimitiveArrayUserValidate endpoint.
func DecodeMethodBodyPrimitiveArrayUserValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body []string
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		if len(body) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body", body, len(body), 1, true))
		}
		for _, e := range body {
			err = goa.MergeErrors(err, goa.ValidatePattern("body[*]", e, "pattern"))
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodBodyPrimitiveArrayUserValidatePayloadType(body)

		return payload, nil
	}
}
`

var PayloadBodyQueryObjectDecodeCode = `// DecodeMethodBodyQueryObjectRequest returns a decoder for requests sent to
// the ServiceBodyQueryObject MethodBodyQueryObject endpoint.
func DecodeMethodBodyQueryObjectRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body MethodBodyQueryObjectRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}

		var (
			b *string
		)
		bRaw := r.URL.Query().Get("b")
		if bRaw != "" {
			b = &bRaw
		}
		payload := NewMethodBodyQueryObjectPayload(&body, b)

		return payload, nil
	}
}
`

var PayloadBodyQueryObjectValidateDecodeCode = `// DecodeMethodBodyQueryObjectValidateRequest returns a decoder for requests
// sent to the ServiceBodyQueryObjectValidate MethodBodyQueryObjectValidate
// endpoint.
func DecodeMethodBodyQueryObjectValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body MethodBodyQueryObjectValidateRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateMethodBodyQueryObjectValidateRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			b string
		)
		b = r.URL.Query().Get("b")
		if b == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("b", "query string"))
		}
		err = goa.MergeErrors(err, goa.ValidatePattern("b", b, "patternb"))
		if err != nil {
			return nil, err
		}
		payload := NewMethodBodyQueryObjectValidatePayload(&body, b)

		return payload, nil
	}
}
`

var PayloadBodyQueryUserDecodeCode = `// DecodeMethodBodyQueryUserRequest returns a decoder for requests sent to the
// ServiceBodyQueryUser MethodBodyQueryUser endpoint.
func DecodeMethodBodyQueryUserRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body MethodBodyQueryUserRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}

		var (
			b *string
		)
		bRaw := r.URL.Query().Get("b")
		if bRaw != "" {
			b = &bRaw
		}
		payload := NewMethodBodyQueryUserPayloadType(&body, b)

		return payload, nil
	}
}
`

var PayloadBodyQueryUserValidateDecodeCode = `// DecodeMethodBodyQueryUserValidateRequest returns a decoder for requests sent
// to the ServiceBodyQueryUserValidate MethodBodyQueryUserValidate endpoint.
func DecodeMethodBodyQueryUserValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body MethodBodyQueryUserValidateRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateMethodBodyQueryUserValidateRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			b string
		)
		b = r.URL.Query().Get("b")
		if b == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("b", "query string"))
		}
		err = goa.MergeErrors(err, goa.ValidatePattern("b", b, "patternb"))
		if err != nil {
			return nil, err
		}
		payload := NewMethodBodyQueryUserValidatePayloadType(&body, b)

		return payload, nil
	}
}
`

var PayloadBodyPathObjectDecodeCode = `// DecodeMethodBodyPathObjectRequest returns a decoder for requests sent to the
// ServiceBodyPathObject MethodBodyPathObject endpoint.
func DecodeMethodBodyPathObjectRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body MethodBodyPathObjectRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}

		var (
			b string

			params = mux.Vars(r)
		)
		b = params["b"]
		payload := NewMethodBodyPathObjectPayload(&body, b)

		return payload, nil
	}
}
`

var PayloadBodyPathObjectValidateDecodeCode = `// DecodeMethodBodyPathObjectValidateRequest returns a decoder for requests
// sent to the ServiceBodyPathObjectValidate MethodBodyPathObjectValidate
// endpoint.
func DecodeMethodBodyPathObjectValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body MethodBodyPathObjectValidateRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateMethodBodyPathObjectValidateRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			b string

			params = mux.Vars(r)
		)
		b = params["b"]
		err = goa.MergeErrors(err, goa.ValidatePattern("b", b, "patternb"))
		if err != nil {
			return nil, err
		}
		payload := NewMethodBodyPathObjectValidatePayload(&body, b)

		return payload, nil
	}
}
`

var PayloadBodyPathUserDecodeCode = `// DecodeMethodBodyPathUserRequest returns a decoder for requests sent to the
// ServiceBodyPathUser MethodBodyPathUser endpoint.
func DecodeMethodBodyPathUserRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body MethodBodyPathUserRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}

		var (
			b string

			params = mux.Vars(r)
		)
		b = params["b"]
		payload := NewMethodBodyPathUserPayloadType(&body, b)

		return payload, nil
	}
}
`

var PayloadBodyPathUserValidateDecodeCode = `// DecodeMethodUserBodyPathValidateRequest returns a decoder for requests sent
// to the ServiceBodyPathUserValidate MethodUserBodyPathValidate endpoint.
func DecodeMethodUserBodyPathValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body MethodUserBodyPathValidateRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateMethodUserBodyPathValidateRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			b string

			params = mux.Vars(r)
		)
		b = params["b"]
		err = goa.MergeErrors(err, goa.ValidatePattern("b", b, "patternb"))
		if err != nil {
			return nil, err
		}
		payload := NewMethodUserBodyPathValidatePayloadType(&body, b)

		return payload, nil
	}
}
`

var PayloadBodyQueryPathObjectDecodeCode = `// DecodeMethodBodyQueryPathObjectRequest returns a decoder for requests sent
// to the ServiceBodyQueryPathObject MethodBodyQueryPathObject endpoint.
func DecodeMethodBodyQueryPathObjectRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body MethodBodyQueryPathObjectRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}

		var (
			c2 string
			b  *string

			params = mux.Vars(r)
		)
		c2 = params["c"]
		bRaw := r.URL.Query().Get("b")
		if bRaw != "" {
			b = &bRaw
		}
		payload := NewMethodBodyQueryPathObjectPayload(&body, c2, b)

		return payload, nil
	}
}
`

var PayloadBodyQueryPathObjectValidateDecodeCode = `// DecodeMethodBodyQueryPathObjectValidateRequest returns a decoder for
// requests sent to the ServiceBodyQueryPathObjectValidate
// MethodBodyQueryPathObjectValidate endpoint.
func DecodeMethodBodyQueryPathObjectValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body MethodBodyQueryPathObjectValidateRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateMethodBodyQueryPathObjectValidateRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			c2 string
			b  string

			params = mux.Vars(r)
		)
		c2 = params["c"]
		err = goa.MergeErrors(err, goa.ValidatePattern("c2", c2, "patternc"))
		b = r.URL.Query().Get("b")
		if b == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("b", "query string"))
		}
		err = goa.MergeErrors(err, goa.ValidatePattern("b", b, "patternb"))
		if err != nil {
			return nil, err
		}
		payload := NewMethodBodyQueryPathObjectValidatePayload(&body, c2, b)

		return payload, nil
	}
}
`

var PayloadBodyQueryPathUserDecodeCode = `// DecodeMethodBodyQueryPathUserRequest returns a decoder for requests sent to
// the ServiceBodyQueryPathUser MethodBodyQueryPathUser endpoint.
func DecodeMethodBodyQueryPathUserRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body MethodBodyQueryPathUserRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}

		var (
			c2 string
			b  *string

			params = mux.Vars(r)
		)
		c2 = params["c"]
		bRaw := r.URL.Query().Get("b")
		if bRaw != "" {
			b = &bRaw
		}
		payload := NewMethodBodyQueryPathUserPayloadType(&body, c2, b)

		return payload, nil
	}
}
`

var PayloadBodyQueryPathUserValidateDecodeCode = `// DecodeMethodBodyQueryPathUserValidateRequest returns a decoder for requests
// sent to the ServiceBodyQueryPathUserValidate MethodBodyQueryPathUserValidate
// endpoint.
func DecodeMethodBodyQueryPathUserValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body MethodBodyQueryPathUserValidateRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateMethodBodyQueryPathUserValidateRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			c2 string
			b  string

			params = mux.Vars(r)
		)
		c2 = params["c"]
		err = goa.MergeErrors(err, goa.ValidatePattern("c2", c2, "patternc"))
		b = r.URL.Query().Get("b")
		if b == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("b", "query string"))
		}
		err = goa.MergeErrors(err, goa.ValidatePattern("b", b, "patternb"))
		if err != nil {
			return nil, err
		}
		payload := NewMethodBodyQueryPathUserValidatePayloadType(&body, c2, b)

		return payload, nil
	}
}
`

var PayloadMapQueryPrimitivePrimitiveDecodeCode = `// DecodeMapQueryPrimitivePrimitiveRequest returns a decoder for requests sent
// to the ServiceMapQueryPrimitivePrimitive MapQueryPrimitivePrimitive endpoint.
func DecodeMapQueryPrimitivePrimitiveRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			query map[string]string
			err   error
		)
		{
			queryRaw := r.URL.Query()
			if len(queryRaw) == 0 {
				err = goa.MergeErrors(err, goa.MissingFieldError("query", "query string"))
			}
			for keyRaw, valRaw := range queryRaw {
				if strings.HasPrefix(keyRaw, "query[") {
					if query == nil {
						query = make(map[string]string)
					}
					var keya string
					{
						openIdx := strings.IndexRune(keyRaw, '[')
						closeIdx := strings.IndexRune(keyRaw, ']')
						keya = keyRaw[openIdx+1 : closeIdx]
					}
					query[keya] = valRaw[0]
				}
			}
		}
		if err != nil {
			return nil, err
		}
		payload := query

		return payload, nil
	}
}
`

var PayloadMapQueryPrimitiveArrayDecodeCode = `// DecodeMapQueryPrimitiveArrayRequest returns a decoder for requests sent to
// the ServiceMapQueryPrimitiveArray MapQueryPrimitiveArray endpoint.
func DecodeMapQueryPrimitiveArrayRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			query map[string][]uint
			err   error
		)
		{
			queryRaw := r.URL.Query()
			if len(queryRaw) == 0 {
				err = goa.MergeErrors(err, goa.MissingFieldError("query", "query string"))
			}
			for keyRaw, valRaw := range queryRaw {
				if strings.HasPrefix(keyRaw, "query[") {
					if query == nil {
						query = make(map[string][]uint)
					}
					var keya string
					{
						openIdx := strings.IndexRune(keyRaw, '[')
						closeIdx := strings.IndexRune(keyRaw, ']')
						keya = keyRaw[openIdx+1 : closeIdx]
					}
					var val []uint
					{
						val = make([]uint, len(valRaw))
						for i, rv := range valRaw {
							v, err2 := strconv.ParseUint(rv, 10, strconv.IntSize)
							if err2 != nil {
								err = goa.MergeErrors(err, goa.InvalidFieldTypeError("val", valRaw, "array of unsigned integers"))
							}
							val[i] = uint(v)
						}
					}
					query[keya] = val
				}
			}
		}
		if err != nil {
			return nil, err
		}
		payload := query

		return payload, nil
	}
}
`

var PayloadMapQueryObjectDecodeCode = `// DecodeMethodMapQueryObjectRequest returns a decoder for requests sent to the
// ServiceMapQueryObject MethodMapQueryObject endpoint.
func DecodeMethodMapQueryObjectRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body MethodMapQueryObjectRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateMethodMapQueryObjectRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			a string
			c map[int][]string

			params = mux.Vars(r)
		)
		a = params["a"]
		err = goa.MergeErrors(err, goa.ValidatePattern("a", a, "patterna"))
		{
			cRaw := r.URL.Query()
			if len(cRaw) == 0 {
				err = goa.MergeErrors(err, goa.MissingFieldError("c", "query string"))
			}
			for keyRaw, valRaw := range cRaw {
				if strings.HasPrefix(keyRaw, "c[") {
					if c == nil {
						c = make(map[int][]string)
					}
					var keya int
					{
						openIdx := strings.IndexRune(keyRaw, '[')
						closeIdx := strings.IndexRune(keyRaw, ']')
						keyaRaw := keyRaw[openIdx+1 : closeIdx]
						v, err2 := strconv.ParseInt(keyaRaw, 10, strconv.IntSize)
						if err2 != nil {
							err = goa.MergeErrors(err, goa.InvalidFieldTypeError("keya", keyaRaw, "integer"))
						}
						keya = int(v)
					}
					c[keya] = valRaw
				}
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodMapQueryObjectPayloadType(&body, a, c)

		return payload, nil
	}
}
`

var PayloadMultipartPrimitiveDecodeCode = `// DecodeMethodMultipartPrimitiveRequest returns a decoder for requests sent to
// the ServiceMultipartPrimitive MethodMultipartPrimitive endpoint.
func DecodeMethodMultipartPrimitiveRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var payload string
		if err := decoder(r).Decode(&payload); err != nil {
			return nil, goa.DecodePayloadError(err.Error())
		}

		return payload, nil
	}
}
`

var PayloadMultipartUserTypeDecodeCode = `// DecodeMethodMultipartUserTypeRequest returns a decoder for requests sent to
// the ServiceMultipartUserType MethodMultipartUserType endpoint.
func DecodeMethodMultipartUserTypeRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var payload *servicemultipartusertype.MethodMultipartUserTypePayload
		if err := decoder(r).Decode(&payload); err != nil {
			return nil, goa.DecodePayloadError(err.Error())
		}

		return payload, nil
	}
}
`

var PayloadMultipartArrayTypeDecodeCode = `// DecodeMethodMultipartArrayTypeRequest returns a decoder for requests sent to
// the ServiceMultipartArrayType MethodMultipartArrayType endpoint.
func DecodeMethodMultipartArrayTypeRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var payload []*servicemultipartarraytype.PayloadType
		if err := decoder(r).Decode(&payload); err != nil {
			return nil, goa.DecodePayloadError(err.Error())
		}

		return payload, nil
	}
}
`

var PayloadMultipartMapTypeDecodeCode = `// DecodeMethodMultipartMapTypeRequest returns a decoder for requests sent to
// the ServiceMultipartMapType MethodMultipartMapType endpoint.
func DecodeMethodMultipartMapTypeRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var payload map[string]int
		if err := decoder(r).Decode(&payload); err != nil {
			return nil, goa.DecodePayloadError(err.Error())
		}

		return payload, nil
	}
}
`

var WithParamsAndHeadersBlockDecodeCode = `// DecodeMethodARequest returns a decoder for requests sent to the
// ServiceWithParamsAndHeadersBlock MethodA endpoint.
func DecodeMethodARequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body MethodARequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}

		var (
			path                      uint
			optional                  *int
			optionalButRequiredParam  float32
			required                  string
			optionalButRequiredHeader float32

			params = mux.Vars(r)
		)
		{
			pathRaw := params["path"]
			v, err2 := strconv.ParseUint(pathRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("path", pathRaw, "unsigned integer"))
			}
			path = uint(v)
		}
		{
			optionalRaw := r.URL.Query().Get("optional")
			if optionalRaw != "" {
				v, err2 := strconv.ParseInt(optionalRaw, 10, strconv.IntSize)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("optional", optionalRaw, "integer"))
				}
				pv := int(v)
				optional = &pv
			}
		}
		{
			optionalButRequiredParamRaw := r.URL.Query().Get("optional_but_required_param")
			if optionalButRequiredParamRaw == "" {
				err = goa.MergeErrors(err, goa.MissingFieldError("optional_but_required_param", "query string"))
			}
			v, err2 := strconv.ParseFloat(optionalButRequiredParamRaw, 32)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("optionalButRequiredParam", optionalButRequiredParamRaw, "float"))
			}
			optionalButRequiredParam = float32(v)
		}
		required = r.Header.Get("required")
		if required == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("required", "header"))
		}
		{
			optionalButRequiredHeaderRaw := r.Header.Get("optional_but_required_header")
			if optionalButRequiredHeaderRaw == "" {
				err = goa.MergeErrors(err, goa.MissingFieldError("optional_but_required_header", "header"))
			}
			v, err2 := strconv.ParseFloat(optionalButRequiredHeaderRaw, 32)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("optionalButRequiredHeader", optionalButRequiredHeaderRaw, "float"))
			}
			optionalButRequiredHeader = float32(v)
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodAPayload(&body, path, optional, optionalButRequiredParam, required, optionalButRequiredHeader)

		return payload, nil
	}
}
`

var QueryIntAliasDecodeCode = `// DecodeMethodARequest returns a decoder for requests sent to the
// ServiceQueryIntAlias MethodA endpoint.
func DecodeMethodARequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			int_   *int
			int32_ *int32
			int64_ *int64
			err    error
		)
		{
			int_Raw := r.URL.Query().Get("int")
			if int_Raw != "" {
				v, err2 := strconv.ParseInt(int_Raw, 10, strconv.IntSize)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("int_", int_Raw, "integer"))
				}
				pv := int(v)
				int_ = &pv
			}
		}
		{
			int32_Raw := r.URL.Query().Get("int32")
			if int32_Raw != "" {
				v, err2 := strconv.ParseInt(int32_Raw, 10, 32)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("int32_", int32_Raw, "integer"))
				}
				pv := int32(v)
				int32_ = &pv
			}
		}
		{
			int64_Raw := r.URL.Query().Get("int64")
			if int64_Raw != "" {
				v, err2 := strconv.ParseInt(int64_Raw, 10, 64)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("int64_", int64_Raw, "integer"))
				}
				int64_ = &v
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodAPayload(int_, int32_, int64_)

		return payload, nil
	}
}
`

var QueryIntAliasValidateDecodeCode = `// DecodeMethodARequest returns a decoder for requests sent to the
// ServiceQueryIntAliasValidate MethodA endpoint.
func DecodeMethodARequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			int_   *int
			int32_ *int32
			int64_ *int64
			err    error
		)
		{
			int_Raw := r.URL.Query().Get("int")
			if int_Raw != "" {
				v, err2 := strconv.ParseInt(int_Raw, 10, strconv.IntSize)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("int_", int_Raw, "integer"))
				}
				pv := int(v)
				int_ = &pv
			}
		}
		if int_ != nil {
			if *int_ < 10 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("int_", *int_, 10, true))
			}
		}
		{
			int32_Raw := r.URL.Query().Get("int32")
			if int32_Raw != "" {
				v, err2 := strconv.ParseInt(int32_Raw, 10, 32)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("int32_", int32_Raw, "integer"))
				}
				pv := int32(v)
				int32_ = &pv
			}
		}
		if int32_ != nil {
			if *int32_ > 100 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("int32_", *int32_, 100, false))
			}
		}
		{
			int64_Raw := r.URL.Query().Get("int64")
			if int64_Raw != "" {
				v, err2 := strconv.ParseInt(int64_Raw, 10, 64)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("int64_", int64_Raw, "integer"))
				}
				int64_ = &v
			}
		}
		if int64_ != nil {
			if *int64_ < 0 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("int64_", *int64_, 0, true))
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodAPayload(int_, int32_, int64_)

		return payload, nil
	}
}
`

var QueryArrayAliasDecodeCode = `// DecodeMethodARequest returns a decoder for requests sent to the
// ServiceQueryArrayAlias MethodA endpoint.
func DecodeMethodARequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			array []uint
			err   error
		)
		{
			arrayRaw := r.URL.Query()["array"]
			if arrayRaw != nil {
				array = make([]uint, len(arrayRaw))
				for i, rv := range arrayRaw {
					v, err2 := strconv.ParseUint(rv, 10, strconv.IntSize)
					if err2 != nil {
						err = goa.MergeErrors(err, goa.InvalidFieldTypeError("array", arrayRaw, "array of unsigned integers"))
					}
					array[i] = uint(v)
				}
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodAPayload(array)

		return payload, nil
	}
}
`

var QueryArrayAliasValidateDecodeCode = `// DecodeMethodARequest returns a decoder for requests sent to the
// ServiceQueryArrayAliasValidate MethodA endpoint.
func DecodeMethodARequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			array []uint
			err   error
		)
		{
			arrayRaw := r.URL.Query()["array"]
			if arrayRaw != nil {
				array = make([]uint, len(arrayRaw))
				for i, rv := range arrayRaw {
					v, err2 := strconv.ParseUint(rv, 10, strconv.IntSize)
					if err2 != nil {
						err = goa.MergeErrors(err, goa.InvalidFieldTypeError("array", arrayRaw, "array of unsigned integers"))
					}
					array[i] = uint(v)
				}
			}
		}
		if len(array) < 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("array", array, len(array), 3, true))
		}
		for _, e := range array {
			if e < 10 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("array[*]", e, 10, true))
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodAPayload(array)

		return payload, nil
	}
}
`
var QueryMapAliasDecodeCode = `// DecodeMethodARequest returns a decoder for requests sent to the
// ServiceQueryMapAlias MethodA endpoint.
func DecodeMethodARequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			map_ map[float32]bool
			err  error
		)
		{
			map_Raw := r.URL.Query()
			if len(map_Raw) != 0 {
				for keyRaw, valRaw := range map_Raw {
					if strings.HasPrefix(keyRaw, "map[") {
						if map_ == nil {
							map_ = make(map[float32]bool)
						}
						var keya float32
						{
							openIdx := strings.IndexRune(keyRaw, '[')
							closeIdx := strings.IndexRune(keyRaw, ']')
							keyaRaw := keyRaw[openIdx+1 : closeIdx]
							v, err2 := strconv.ParseFloat(keyaRaw, 32)
							if err2 != nil {
								err = goa.MergeErrors(err, goa.InvalidFieldTypeError("keya", keyaRaw, "float"))
							}
							keya = float32(v)
						}
						var vala bool
						{
							valaRaw := valRaw[0]
							v, err2 := strconv.ParseBool(valaRaw)
							if err2 != nil {
								err = goa.MergeErrors(err, goa.InvalidFieldTypeError("vala", valaRaw, "boolean"))
							}
							vala = v
						}
						map_[keya] = vala
					}
				}
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodAPayload(map_)

		return payload, nil
	}
}
`

var QueryMapAliasValidateDecodeCode = `// DecodeMethodARequest returns a decoder for requests sent to the
// ServiceQueryMapAliasValidate MethodA endpoint.
func DecodeMethodARequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			map_ map[float32]bool
			err  error
		)
		{
			map_Raw := r.URL.Query()
			if len(map_Raw) != 0 {
				for keyRaw, valRaw := range map_Raw {
					if strings.HasPrefix(keyRaw, "map[") {
						if map_ == nil {
							map_ = make(map[float32]bool)
						}
						var keya float32
						{
							openIdx := strings.IndexRune(keyRaw, '[')
							closeIdx := strings.IndexRune(keyRaw, ']')
							keyaRaw := keyRaw[openIdx+1 : closeIdx]
							v, err2 := strconv.ParseFloat(keyaRaw, 32)
							if err2 != nil {
								err = goa.MergeErrors(err, goa.InvalidFieldTypeError("keya", keyaRaw, "float"))
							}
							keya = float32(v)
						}
						var vala bool
						{
							valaRaw := valRaw[0]
							v, err2 := strconv.ParseBool(valaRaw)
							if err2 != nil {
								err = goa.MergeErrors(err, goa.InvalidFieldTypeError("vala", valaRaw, "boolean"))
							}
							vala = v
						}
						map_[keya] = vala
					}
				}
			}
		}
		if len(map_) < 5 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("map_", map_, len(map_), 5, true))
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodAPayload(map_)

		return payload, nil
	}
}
`

var QueryArrayNestedAliasValidateDecodeCode = `// DecodeMethodARequest returns a decoder for requests sent to the
// ServiceQueryArrayAliasValidate MethodA endpoint.
func DecodeMethodARequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			array []float64
			err   error
		)
		{
			arrayRaw := r.URL.Query()["array"]
			if arrayRaw != nil {
				array = make([]float64, len(arrayRaw))
				for i, rv := range arrayRaw {
					v, err2 := strconv.ParseFloat(rv, 64)
					if err2 != nil {
						err = goa.MergeErrors(err, goa.InvalidFieldTypeError("array", arrayRaw, "array of floats"))
					}
					array[i] = v
				}
			}
		}
		for _, e := range array {
			if e < 10 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("array[*]", e, 10, true))
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodAPayload(array)

		return payload, nil
	}
}
`

var HeaderIntAliasDecodeCode = `// DecodeMethodARequest returns a decoder for requests sent to the
// ServiceHeaderIntAlias MethodA endpoint.
func DecodeMethodARequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			int_   *int
			int32_ *int32
			int64_ *int64
			err    error
		)
		{
			int_Raw := r.Header.Get("int")
			if int_Raw != "" {
				v, err2 := strconv.ParseInt(int_Raw, 10, strconv.IntSize)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("int_", int_Raw, "integer"))
				}
				pv := int(v)
				int_ = &pv
			}
		}
		{
			int32_Raw := r.Header.Get("int32")
			if int32_Raw != "" {
				v, err2 := strconv.ParseInt(int32_Raw, 10, 32)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("int32_", int32_Raw, "integer"))
				}
				pv := int32(v)
				int32_ = &pv
			}
		}
		{
			int64_Raw := r.Header.Get("int64")
			if int64_Raw != "" {
				v, err2 := strconv.ParseInt(int64_Raw, 10, 64)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("int64_", int64_Raw, "integer"))
				}
				int64_ = &v
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodAPayload(int_, int32_, int64_)

		return payload, nil
	}
}
`

var PathIntAliasDecodeCode = `// DecodeMethodARequest returns a decoder for requests sent to the
// ServicePathIntAlias MethodA endpoint.
func DecodeMethodARequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			int_   int
			int32_ int32
			int64_ int64
			err    error

			params = mux.Vars(r)
		)
		{
			int_Raw := params["int"]
			v, err2 := strconv.ParseInt(int_Raw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("int_", int_Raw, "integer"))
			}
			int_ = int(v)
		}
		{
			int32_Raw := params["int32"]
			v, err2 := strconv.ParseInt(int32_Raw, 10, 32)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("int32_", int32_Raw, "integer"))
			}
			int32_ = int32(v)
		}
		{
			int64_Raw := params["int64"]
			v, err2 := strconv.ParseInt(int64_Raw, 10, 64)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("int64_", int64_Raw, "integer"))
			}
			int64_ = v
		}
		if err != nil {
			return nil, err
		}
		payload := NewMethodAPayload(int_, int32_, int64_)

		return payload, nil
	}
}
`
