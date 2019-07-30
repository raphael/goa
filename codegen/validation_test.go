package codegen

import (
	"testing"

	"goa.design/goa/v3/codegen/testdata"
	"goa.design/goa/v3/expr"
)

func TestRecursiveValidationCode(t *testing.T) {
	root := RunDSL(t, testdata.ValidationTypesDSL)
	var (
		scope = NewNameScope()

		integerT = root.UserType("Integer")
		stringT  = root.UserType("String")
		floatT   = root.UserType("Float")
		userT    = root.UserType("UserType")
		arrayUT  = root.UserType("ArrayUserType")
		arrayT   = root.UserType("Array")
		mapT     = root.UserType("Map")
		rtT      = root.UserType("Result")
		rtcolT   = root.UserType("Collection")
		colT     = root.UserType("TypeWithCollection")
	)
	cases := []struct {
		Name       string
		Type       expr.UserType
		Required   bool
		Pointer    bool
		UseDefault bool
		Code       string
	}{
		{"integer-required", integerT, true, false, false, testdata.IntegerRequiredValidationCode},
		{"integer-pointer", integerT, false, true, false, testdata.IntegerPointerValidationCode},
		{"integer-use-default", integerT, false, false, true, testdata.IntegerUseDefaultValidationCode},
		{"float-required", floatT, true, false, false, testdata.FloatRequiredValidationCode},
		{"float-pointer", floatT, false, true, false, testdata.FloatPointerValidationCode},
		{"float-use-default", floatT, false, false, true, testdata.FloatUseDefaultValidationCode},
		{"string-required", stringT, true, false, false, testdata.StringRequiredValidationCode},
		{"string-pointer", stringT, false, true, false, testdata.StringPointerValidationCode},
		{"string-use-default", stringT, false, false, true, testdata.StringUseDefaultValidationCode},
		{"user-type-required", userT, true, false, false, testdata.UserTypeRequiredValidationCode},
		{"user-type-pointer", userT, false, true, false, testdata.UserTypePointerValidationCode},
		{"user-type-default", userT, false, false, true, testdata.UserTypeUseDefaultValidationCode},
		{"user-type-array-required", arrayUT, true, true, false, testdata.UserTypeArrayValidationCode},
		{"array-required", arrayT, true, false, false, testdata.ArrayRequiredValidationCode},
		{"array-pointer", arrayT, false, true, false, testdata.ArrayPointerValidationCode},
		{"array-use-default", arrayT, false, false, true, testdata.ArrayUseDefaultValidationCode},
		{"map-required", mapT, true, false, false, testdata.MapRequiredValidationCode},
		{"map-pointer", mapT, false, true, false, testdata.MapPointerValidationCode},
		{"map-use-default", mapT, false, false, true, testdata.MapUseDefaultValidationCode},
		{"result-type-pointer", rtT, false, true, false, testdata.ResultTypePointerValidationCode},
		{"collection-required", rtcolT, true, false, false, testdata.ResultCollectionPointerValidationCode},
		{"collection-pointer", rtcolT, false, true, false, testdata.ResultCollectionPointerValidationCode},
		{"type-with-collection-pointer", colT, false, true, false, testdata.TypeWithCollectionPointerValidationCode},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			ctx := NewAttributeContext(c.Pointer, false, c.UseDefault, "", scope)
			code := RecursiveValidationCode(&expr.AttributeExpr{Type: c.Type}, ctx, c.Required, "target")
			code = FormatTestCode(t, "package foo\nfunc Validate() (err error){\n"+code+"}")
			if code != c.Code {
				t.Errorf("invalid code, got:\n%s\ngot vs. expected:\n%s", code, Diff(t, code, c.Code))
			}
		})
	}
}
