package codegen

import (
	"strings"
	"testing"

	"goa.design/goa/v3/codegen"
	"goa.design/goa/v3/expr"
	"goa.design/goa/v3/http/codegen/testdata"
)

func TestDecode(t *testing.T) {
	cases := []struct {
		Name string
		DSL  func()
		Code string
	}{
		{"query-bool", testdata.PayloadQueryBoolDSL, testdata.PayloadQueryBoolDecodeCode},
		{"query-bool-validate", testdata.PayloadQueryBoolValidateDSL, testdata.PayloadQueryBoolValidateDecodeCode},
		{"query-int", testdata.PayloadQueryIntDSL, testdata.PayloadQueryIntDecodeCode},
		{"query-int-validate", testdata.PayloadQueryIntValidateDSL, testdata.PayloadQueryIntValidateDecodeCode},
		{"query-int32", testdata.PayloadQueryInt32DSL, testdata.PayloadQueryInt32DecodeCode},
		{"query-int32-validate", testdata.PayloadQueryInt32ValidateDSL, testdata.PayloadQueryInt32ValidateDecodeCode},
		{"query-int64", testdata.PayloadQueryInt64DSL, testdata.PayloadQueryInt64DecodeCode},
		{"query-int64-validate", testdata.PayloadQueryInt64ValidateDSL, testdata.PayloadQueryInt64ValidateDecodeCode},
		{"query-uint", testdata.PayloadQueryUIntDSL, testdata.PayloadQueryUIntDecodeCode},
		{"query-uint-validate", testdata.PayloadQueryUIntValidateDSL, testdata.PayloadQueryUIntValidateDecodeCode},
		{"query-uint32", testdata.PayloadQueryUInt32DSL, testdata.PayloadQueryUInt32DecodeCode},
		{"query-uint32-validate", testdata.PayloadQueryUInt32ValidateDSL, testdata.PayloadQueryUInt32ValidateDecodeCode},
		{"query-uint64", testdata.PayloadQueryUInt64DSL, testdata.PayloadQueryUInt64DecodeCode},
		{"query-uint64-validate", testdata.PayloadQueryUInt64ValidateDSL, testdata.PayloadQueryUInt64ValidateDecodeCode},
		{"query-float32", testdata.PayloadQueryFloat32DSL, testdata.PayloadQueryFloat32DecodeCode},
		{"query-float32-validate", testdata.PayloadQueryFloat32ValidateDSL, testdata.PayloadQueryFloat32ValidateDecodeCode},
		{"query-float64", testdata.PayloadQueryFloat64DSL, testdata.PayloadQueryFloat64DecodeCode},
		{"query-float64-validate", testdata.PayloadQueryFloat64ValidateDSL, testdata.PayloadQueryFloat64ValidateDecodeCode},
		{"query-string", testdata.PayloadQueryStringDSL, testdata.PayloadQueryStringDecodeCode},
		{"query-string-validate", testdata.PayloadQueryStringValidateDSL, testdata.PayloadQueryStringValidateDecodeCode},
		{"query-string-not-required-validate", testdata.PayloadQueryStringNotRequiredValidateDSL, testdata.PayloadQueryStringNotRequiredValidateDecodeCode},
		{"query-bytes", testdata.PayloadQueryBytesDSL, testdata.PayloadQueryBytesDecodeCode},
		{"query-bytes-validate", testdata.PayloadQueryBytesValidateDSL, testdata.PayloadQueryBytesValidateDecodeCode},
		{"query-any", testdata.PayloadQueryAnyDSL, testdata.PayloadQueryAnyDecodeCode},
		{"query-any-validate", testdata.PayloadQueryAnyValidateDSL, testdata.PayloadQueryAnyValidateDecodeCode},
		{"query-array-bool", testdata.PayloadQueryArrayBoolDSL, testdata.PayloadQueryArrayBoolDecodeCode},
		{"query-array-bool-validate", testdata.PayloadQueryArrayBoolValidateDSL, testdata.PayloadQueryArrayBoolValidateDecodeCode},
		{"query-array-int", testdata.PayloadQueryArrayIntDSL, testdata.PayloadQueryArrayIntDecodeCode},
		{"query-array-int-validate", testdata.PayloadQueryArrayIntValidateDSL, testdata.PayloadQueryArrayIntValidateDecodeCode},
		{"query-array-int32", testdata.PayloadQueryArrayInt32DSL, testdata.PayloadQueryArrayInt32DecodeCode},
		{"query-array-int32-validate", testdata.PayloadQueryArrayInt32ValidateDSL, testdata.PayloadQueryArrayInt32ValidateDecodeCode},
		{"query-array-int64", testdata.PayloadQueryArrayInt64DSL, testdata.PayloadQueryArrayInt64DecodeCode},
		{"query-array-int64-validate", testdata.PayloadQueryArrayInt64ValidateDSL, testdata.PayloadQueryArrayInt64ValidateDecodeCode},
		{"query-array-uint", testdata.PayloadQueryArrayUIntDSL, testdata.PayloadQueryArrayUIntDecodeCode},
		{"query-array-uint-validate", testdata.PayloadQueryArrayUIntValidateDSL, testdata.PayloadQueryArrayUIntValidateDecodeCode},
		{"query-array-uint32", testdata.PayloadQueryArrayUInt32DSL, testdata.PayloadQueryArrayUInt32DecodeCode},
		{"query-array-uint32-validate", testdata.PayloadQueryArrayUInt32ValidateDSL, testdata.PayloadQueryArrayUInt32ValidateDecodeCode},
		{"query-array-uint64", testdata.PayloadQueryArrayUInt64DSL, testdata.PayloadQueryArrayUInt64DecodeCode},
		{"query-array-uint64-validate", testdata.PayloadQueryArrayUInt64ValidateDSL, testdata.PayloadQueryArrayUInt64ValidateDecodeCode},
		{"query-array-float32", testdata.PayloadQueryArrayFloat32DSL, testdata.PayloadQueryArrayFloat32DecodeCode},
		{"query-array-float32-validate", testdata.PayloadQueryArrayFloat32ValidateDSL, testdata.PayloadQueryArrayFloat32ValidateDecodeCode},
		{"query-array-float64", testdata.PayloadQueryArrayFloat64DSL, testdata.PayloadQueryArrayFloat64DecodeCode},
		{"query-array-float64-validate", testdata.PayloadQueryArrayFloat64ValidateDSL, testdata.PayloadQueryArrayFloat64ValidateDecodeCode},
		{"query-array-string", testdata.PayloadQueryArrayStringDSL, testdata.PayloadQueryArrayStringDecodeCode},
		{"query-array-string-validate", testdata.PayloadQueryArrayStringValidateDSL, testdata.PayloadQueryArrayStringValidateDecodeCode},
		{"query-array-bytes", testdata.PayloadQueryArrayBytesDSL, testdata.PayloadQueryArrayBytesDecodeCode},
		{"query-array-bytes-validate", testdata.PayloadQueryArrayBytesValidateDSL, testdata.PayloadQueryArrayBytesValidateDecodeCode},
		{"query-array-any", testdata.PayloadQueryArrayAnyDSL, testdata.PayloadQueryArrayAnyDecodeCode},
		{"query-array-any-validate", testdata.PayloadQueryArrayAnyValidateDSL, testdata.PayloadQueryArrayAnyValidateDecodeCode},
		{"query-map-string-string", testdata.PayloadQueryMapStringStringDSL, testdata.PayloadQueryMapStringStringDecodeCode},
		{"query-map-string-string-validate", testdata.PayloadQueryMapStringStringValidateDSL, testdata.PayloadQueryMapStringStringValidateDecodeCode},
		{"query-map-string-bool", testdata.PayloadQueryMapStringBoolDSL, testdata.PayloadQueryMapStringBoolDecodeCode},
		{"query-map-string-bool-validate", testdata.PayloadQueryMapStringBoolValidateDSL, testdata.PayloadQueryMapStringBoolValidateDecodeCode},
		{"query-map-bool-string", testdata.PayloadQueryMapBoolStringDSL, testdata.PayloadQueryMapBoolStringDecodeCode},
		{"query-map-bool-string-validate", testdata.PayloadQueryMapBoolStringValidateDSL, testdata.PayloadQueryMapBoolStringValidateDecodeCode},
		{"query-map-bool-bool", testdata.PayloadQueryMapBoolBoolDSL, testdata.PayloadQueryMapBoolBoolDecodeCode},
		{"query-map-bool-bool-validate", testdata.PayloadQueryMapBoolBoolValidateDSL, testdata.PayloadQueryMapBoolBoolValidateDecodeCode},
		{"query-map-string-array-string", testdata.PayloadQueryMapStringArrayStringDSL, testdata.PayloadQueryMapStringArrayStringDecodeCode},
		{"query-map-string-array-string-validate", testdata.PayloadQueryMapStringArrayStringValidateDSL, testdata.PayloadQueryMapStringArrayStringValidateDecodeCode},
		{"query-map-string-array-bool", testdata.PayloadQueryMapStringArrayBoolDSL, testdata.PayloadQueryMapStringArrayBoolDecodeCode},
		{"query-map-string-array-bool-validate", testdata.PayloadQueryMapStringArrayBoolValidateDSL, testdata.PayloadQueryMapStringArrayBoolValidateDecodeCode},
		{"query-map-bool-array-string", testdata.PayloadQueryMapBoolArrayStringDSL, testdata.PayloadQueryMapBoolArrayStringDecodeCode},
		{"query-map-bool-array-string-validate", testdata.PayloadQueryMapBoolArrayStringValidateDSL, testdata.PayloadQueryMapBoolArrayStringValidateDecodeCode},
		{"query-map-bool-array-bool", testdata.PayloadQueryMapBoolArrayBoolDSL, testdata.PayloadQueryMapBoolArrayBoolDecodeCode},
		{"query-map-bool-array-bool-validate", testdata.PayloadQueryMapBoolArrayBoolValidateDSL, testdata.PayloadQueryMapBoolArrayBoolValidateDecodeCode},

		{"query-primitive-string-validate", testdata.PayloadQueryPrimitiveStringValidateDSL, testdata.PayloadQueryPrimitiveStringValidateDecodeCode},
		{"query-primitive-bool-validate", testdata.PayloadQueryPrimitiveBoolValidateDSL, testdata.PayloadQueryPrimitiveBoolValidateDecodeCode},
		{"query-primitive-array-string-validate", testdata.PayloadQueryPrimitiveArrayStringValidateDSL, testdata.PayloadQueryPrimitiveArrayStringValidateDecodeCode},
		{"query-primitive-array-bool-validate", testdata.PayloadQueryPrimitiveArrayBoolValidateDSL, testdata.PayloadQueryPrimitiveArrayBoolValidateDecodeCode},
		{"query-primitive-map-string-array-string-validate", testdata.PayloadQueryPrimitiveMapStringArrayStringValidateDSL, testdata.PayloadQueryPrimitiveMapStringArrayStringValidateDecodeCode},
		{"query-primitive-map-string-bool-validate", testdata.PayloadQueryPrimitiveMapStringBoolValidateDSL, testdata.PayloadQueryPrimitiveMapStringBoolValidateDecodeCode},
		{"query-primitive-map-bool-array-bool-validate", testdata.PayloadQueryPrimitiveMapBoolArrayBoolValidateDSL, testdata.PayloadQueryPrimitiveMapBoolArrayBoolValidateDecodeCode},
		{"query-map-string-map-int-string-validate", testdata.PayloadQueryMapStringMapIntStringValidateDSL, testdata.PayloadQueryMapStringMapIntStringValidateDecodeCode},
		{"query-map-int-map-string-array-int-validate", testdata.PayloadQueryMapIntMapStringArrayIntValidateDSL, testdata.PayloadQueryMapIntMapStringArrayIntValidateDecodeCode},

		{"query-string-mapped", testdata.PayloadQueryStringMappedDSL, testdata.PayloadQueryStringMappedDecodeCode},

		{"query-string-default", testdata.PayloadQueryStringDefaultDSL, testdata.PayloadQueryStringDefaultDecodeCode},
		{"query-string-default-validate", testdata.PayloadQueryStringDefaultValidateDSL, testdata.PayloadQueryStringDefaultValidateDecodeCode},
		{"query-primitive-string-default", testdata.PayloadQueryPrimitiveStringDefaultDSL, testdata.PayloadQueryPrimitiveStringDefaultDecodeCode},
		{"query-string-extended-payload", testdata.PayloadExtendedQueryStringDSL, testdata.PayloadExtendedQueryStringDecodeCode},

		{"path-string", testdata.PayloadPathStringDSL, testdata.PayloadPathStringDecodeCode},
		{"path-string-validate", testdata.PayloadPathStringValidateDSL, testdata.PayloadPathStringValidateDecodeCode},
		{"path-array-string", testdata.PayloadPathArrayStringDSL, testdata.PayloadPathArrayStringDecodeCode},
		{"path-array-string-validate", testdata.PayloadPathArrayStringValidateDSL, testdata.PayloadPathArrayStringValidateDecodeCode},

		{"path-primitive-string-validate", testdata.PayloadPathPrimitiveStringValidateDSL, testdata.PayloadPathPrimitiveStringValidateDecodeCode},
		{"path-primitive-bool-validate", testdata.PayloadPathPrimitiveBoolValidateDSL, testdata.PayloadPathPrimitiveBoolValidateDecodeCode},
		{"path-primitive-array-string-validate", testdata.PayloadPathPrimitiveArrayStringValidateDSL, testdata.PayloadPathPrimitiveArrayStringValidateDecodeCode},
		{"path-primitive-array-bool-validate", testdata.PayloadPathPrimitiveArrayBoolValidateDSL, testdata.PayloadPathPrimitiveArrayBoolValidateDecodeCode},

		{"header-string", testdata.PayloadHeaderStringDSL, testdata.PayloadHeaderStringDecodeCode},
		{"header-string-validate", testdata.PayloadHeaderStringValidateDSL, testdata.PayloadHeaderStringValidateDecodeCode},
		{"header-array-string", testdata.PayloadHeaderArrayStringDSL, testdata.PayloadHeaderArrayStringDecodeCode},
		{"header-array-string-validate", testdata.PayloadHeaderArrayStringValidateDSL, testdata.PayloadHeaderArrayStringValidateDecodeCode},

		{"header-primitive-string-validate", testdata.PayloadHeaderPrimitiveStringValidateDSL, testdata.PayloadHeaderPrimitiveStringValidateDecodeCode},
		{"header-primitive-bool-validate", testdata.PayloadHeaderPrimitiveBoolValidateDSL, testdata.PayloadHeaderPrimitiveBoolValidateDecodeCode},
		{"header-primitive-array-string-validate", testdata.PayloadHeaderPrimitiveArrayStringValidateDSL, testdata.PayloadHeaderPrimitiveArrayStringValidateDecodeCode},
		{"header-primitive-array-bool-validate", testdata.PayloadHeaderPrimitiveArrayBoolValidateDSL, testdata.PayloadHeaderPrimitiveArrayBoolValidateDecodeCode},

		{"header-string-default", testdata.PayloadHeaderStringDefaultDSL, testdata.PayloadHeaderStringDefaultDecodeCode},
		{"header-string-default-validate", testdata.PayloadHeaderStringDefaultValidateDSL, testdata.PayloadHeaderStringDefaultValidateDecodeCode},
		{"header-primitive-string-default", testdata.PayloadHeaderPrimitiveStringDefaultDSL, testdata.PayloadHeaderPrimitiveStringDefaultDecodeCode},

		{"body-string", testdata.PayloadBodyStringDSL, testdata.PayloadBodyStringDecodeCode},
		{"body-string-validate", testdata.PayloadBodyStringValidateDSL, testdata.PayloadBodyStringValidateDecodeCode},
		{"body-user", testdata.PayloadBodyUserDSL, testdata.PayloadBodyUserDecodeCode},
		{"body-user-required", testdata.PayloadBodyUserRequiredDSL, testdata.PayloadBodyUserRequiredDecodeCode},
		{"body-user-nested", testdata.PayloadBodyNestedUserDSL, testdata.PayloadBodyNestedUserDecodeCode},
		{"body-user-validate", testdata.PayloadBodyUserValidateDSL, testdata.PayloadBodyUserValidateDecodeCode},
		{"body-array-string", testdata.PayloadBodyArrayStringDSL, testdata.PayloadBodyArrayStringDecodeCode},
		{"body-array-string-validate", testdata.PayloadBodyArrayStringValidateDSL, testdata.PayloadBodyArrayStringValidateDecodeCode},
		{"body-array-user", testdata.PayloadBodyArrayUserDSL, testdata.PayloadBodyArrayUserDecodeCode},
		{"body-array-user-validate", testdata.PayloadBodyArrayUserValidateDSL, testdata.PayloadBodyArrayUserValidateDecodeCode},
		{"body-map-string", testdata.PayloadBodyMapStringDSL, testdata.PayloadBodyMapStringDecodeCode},
		{"body-map-string-validate", testdata.PayloadBodyMapStringValidateDSL, testdata.PayloadBodyMapStringValidateDecodeCode},
		{"body-map-user", testdata.PayloadBodyMapUserDSL, testdata.PayloadBodyMapUserDecodeCode},
		{"body-map-user-validate", testdata.PayloadBodyMapUserValidateDSL, testdata.PayloadBodyMapUserValidateDecodeCode},

		{"body-primitive-string-validate", testdata.PayloadBodyPrimitiveStringValidateDSL, testdata.PayloadBodyPrimitiveStringValidateDecodeCode},
		{"body-primitive-bool-validate", testdata.PayloadBodyPrimitiveBoolValidateDSL, testdata.PayloadBodyPrimitiveBoolValidateDecodeCode},
		{"body-primitive-array-string-validate", testdata.PayloadBodyPrimitiveArrayStringValidateDSL, testdata.PayloadBodyPrimitiveArrayStringValidateDecodeCode},
		{"body-primitive-array-bool-validate", testdata.PayloadBodyPrimitiveArrayBoolValidateDSL, testdata.PayloadBodyPrimitiveArrayBoolValidateDecodeCode},

		{"body-primitive-array-user-validate", testdata.PayloadBodyPrimitiveArrayUserValidateDSL, testdata.PayloadBodyPrimitiveArrayUserValidateDecodeCode},
		{"body-primitive-field-array-user", testdata.PayloadBodyPrimitiveFieldArrayUserDSL, testdata.PayloadBodyPrimitiveFieldArrayUserDecodeCode},
		{"body-extend-primitive-field-array-user", testdata.PayloadExtendBodyPrimitiveFieldArrayUserDSL, testdata.PayloadBodyPrimitiveFieldArrayUserDecodeCode},
		{"body-extend-primitive-field-string", testdata.PayloadExtendBodyPrimitiveFieldStringDSL, testdata.PayloadBodyPrimitiveFieldStringDecodeCode},
		{"body-primitive-field-array-user-validate", testdata.PayloadBodyPrimitiveFieldArrayUserValidateDSL, testdata.PayloadBodyPrimitiveFieldArrayUserValidateDecodeCode},

		{"body-query-object", testdata.PayloadBodyQueryObjectDSL, testdata.PayloadBodyQueryObjectDecodeCode},
		{"body-query-object-validate", testdata.PayloadBodyQueryObjectValidateDSL, testdata.PayloadBodyQueryObjectValidateDecodeCode},
		{"body-query-user", testdata.PayloadBodyQueryUserDSL, testdata.PayloadBodyQueryUserDecodeCode},
		{"body-query-user-validate", testdata.PayloadBodyQueryUserValidateDSL, testdata.PayloadBodyQueryUserValidateDecodeCode},

		{"body-path-object", testdata.PayloadBodyPathObjectDSL, testdata.PayloadBodyPathObjectDecodeCode},
		{"body-path-object-validate", testdata.PayloadBodyPathObjectValidateDSL, testdata.PayloadBodyPathObjectValidateDecodeCode},
		{"body-path-user", testdata.PayloadBodyPathUserDSL, testdata.PayloadBodyPathUserDecodeCode},
		{"body-path-user-validate", testdata.PayloadBodyPathUserValidateDSL, testdata.PayloadBodyPathUserValidateDecodeCode},

		{"body-query-path-object", testdata.PayloadBodyQueryPathObjectDSL, testdata.PayloadBodyQueryPathObjectDecodeCode},
		{"body-query-path-object-validate", testdata.PayloadBodyQueryPathObjectValidateDSL, testdata.PayloadBodyQueryPathObjectValidateDecodeCode},
		{"body-query-path-user", testdata.PayloadBodyQueryPathUserDSL, testdata.PayloadBodyQueryPathUserDecodeCode},
		{"body-query-path-user-validate", testdata.PayloadBodyQueryPathUserValidateDSL, testdata.PayloadBodyQueryPathUserValidateDecodeCode},

		{"map-query-primitive-primitive", testdata.PayloadMapQueryPrimitivePrimitiveDSL, testdata.PayloadMapQueryPrimitivePrimitiveDecodeCode},
		{"map-query-primitive-array", testdata.PayloadMapQueryPrimitiveArrayDSL, testdata.PayloadMapQueryPrimitiveArrayDecodeCode},
		{"map-query-object", testdata.PayloadMapQueryObjectDSL, testdata.PayloadMapQueryObjectDecodeCode},
		{"multipart-body-primitive", testdata.PayloadMultipartPrimitiveDSL, testdata.PayloadMultipartPrimitiveDecodeCode},
		{"multipart-body-user-type", testdata.PayloadMultipartUserTypeDSL, testdata.PayloadMultipartUserTypeDecodeCode},
		{"multipart-body-array-type", testdata.PayloadMultipartArrayTypeDSL, testdata.PayloadMultipartArrayTypeDecodeCode},
		{"multipart-body-map-type", testdata.PayloadMultipartMapTypeDSL, testdata.PayloadMultipartMapTypeDecodeCode},
		{"with-params-and-headers-dsl", testdata.WithParamsAndHeadersBlockDSL, testdata.WithParamsAndHeadersBlockDecodeCode},

		// aliases
		{"query-int-alias", testdata.QueryIntAliasDSL, testdata.QueryIntAliasDecodeCode},
		{"query-int-alias-validate", testdata.QueryIntAliasValidateDSL, testdata.QueryIntAliasValidateDecodeCode},
		{"query-array-alias", testdata.QueryArrayAliasDSL, testdata.QueryArrayAliasDecodeCode},
		{"query-array-alias-validate", testdata.QueryArrayAliasValidateDSL, testdata.QueryArrayAliasValidateDecodeCode},
		{"query-map-alias", testdata.QueryMapAliasDSL, testdata.QueryMapAliasDecodeCode},
		{"query-map-alias-validate", testdata.QueryMapAliasValidateDSL, testdata.QueryMapAliasValidateDecodeCode},
		{"query-array-nested-alias-validate", testdata.QueryArrayNestedAliasValidateDSL, testdata.QueryArrayNestedAliasValidateDecodeCode},
		{"header-int-alias", testdata.HeaderIntAliasDSL, testdata.HeaderIntAliasDecodeCode},
		{"path-int-alias", testdata.PathIntAliasDSL, testdata.PathIntAliasDecodeCode},
	}
	golden := makeGolden(t, "testdata/payload_decode_functions.go")
	if golden != nil {
		golden.WriteString("package testdata\n")
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			RunHTTPDSL(t, c.DSL)
			fs := ServerFiles("", expr.Root)
			if len(fs) != 2 {
				t.Fatalf("got %d files, expected two", len(fs))
			}
			sections := fs[1].SectionTemplates
			if len(sections) < 3 {
				t.Fatalf("got %d sections, expected at least 3", len(sections))
			}
			code := codegen.SectionCode(t, sections[2])
			if golden != nil {
				name := codegen.Goify(c.Name, true)
				name = strings.Replace(name, "Uint", "UInt", -1)
				code = "\nvar Payload" + name + "DecodeCode = `" + code + "`"
				golden.WriteString(code + "\n")
			} else if code != c.Code {
				t.Errorf("invalid code, got:\n%s\ngot vs. expected:\n%s", code, codegen.Diff(t, code, c.Code))
			}
		})
	}
}
