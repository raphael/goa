package codegen

import (
	"testing"

	"goa.design/goa/v3/codegen/testdata"
	"goa.design/goa/v3/expr"
)

func TestGoTransformHelpers(t *testing.T) {
	root := RunDSL(t, testdata.TestTypesDSL)
	var (
		scope = NewNameScope()
		// types to test
		simple        = root.UserType("Simple")
		recursive     = root.UserType("Recursive")
		composite     = root.UserType("Composite")
		deep          = root.UserType("Deep")
		deepArray     = root.UserType("DeepArray")
		simpleAlias   = root.UserType("SimpleAlias")
		mapAlias      = root.UserType("NestedMapAlias")
		arrayMapAlias = root.UserType("ArrayMapAlias")
		collection    = root.UserType("ResultTypeCollection")
		// attribute contexts used in test cases
		defaultCtx = NewAttributeContext(false, false, true, "", scope)
	)
	tc := []struct {
		Name        string
		Type        expr.DataType
		HelperNames []string
	}{
		{"simple", simple, []string{}},
		{"recursive", recursive, []string{"transformRecursiveToRecursive"}},
		{"composite", composite, []string{"transformSimpleToSimple"}},
		{"deep", deep, []string{"transformCompositeToComposite", "transformSimpleToSimple"}},
		{"deep-array", deepArray, []string{"transformCompositeToComposite", "transformSimpleToSimple"}},
		{"simple-alias", simpleAlias, []string{}},
		{"nested-map-alias", mapAlias, []string{}},
		{"array-map-alias", arrayMapAlias, []string{}},
		{"result-type-collection", collection, []string{"transformResultTypeToResultType"}},
	}
	for _, c := range tc {
		t.Run(c.Name, func(t *testing.T) {
			if c.Type == nil {
				t.Fatal("source type not found in testdata")
			}
			_, funcs, err := GoTransform(&expr.AttributeExpr{Type: c.Type}, &expr.AttributeExpr{Type: c.Type}, "source", "target", defaultCtx, defaultCtx, "")
			if err != nil {
				t.Fatal(err)
			}
			if len(funcs) != len(c.HelperNames) {
				t.Errorf("invalid helpers count, got: %d, expected %d", len(funcs), len(c.HelperNames))
			} else {
				var diffs []string
				actual := make([]string, len(funcs))
				for i, f := range funcs {
					actual[i] = f.Name
					if c.HelperNames[i] != f.Name {
						diffs = append(diffs, f.Name)
					}
				}
				if len(diffs) > 0 {
					t.Errorf("invalid helper names, got: %v, expected: %v", actual, c.HelperNames)
				}
			}
		})
	}
}
