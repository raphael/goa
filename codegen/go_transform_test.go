package codegen

import (
	"testing"

	"goa.design/goa/v3/codegen/testdata"
	"goa.design/goa/v3/expr"
)

func TestGoTransform(t *testing.T) {
	root := RunDSL(t, testdata.TestTypesDSL)
	var (
		scope = NewNameScope()

		// types to test
		simple   = root.UserType("Simple")
		super    = root.UserType("Super")
		required = root.UserType("Required")
		defaultT = root.UserType("Default")

		simpleMap   = root.UserType("SimpleMap")
		requiredMap = root.UserType("RequiredMap")
		defaultMap  = root.UserType("DefaultMap")
		nestedMap   = root.UserType("NestedMap")
		typeMap     = root.UserType("TypeMap")
		arrayMap    = root.UserType("ArrayMap")

		simpleArray   = root.UserType("SimpleArray")
		requiredArray = root.UserType("RequiredArray")
		defaultArray  = root.UserType("DefaultArray")
		nestedArray   = root.UserType("NestedArray")
		typeArray     = root.UserType("TypeArray")
		mapArray      = root.UserType("MapArray")

		recursive      = root.UserType("Recursive")
		recursiveArray = root.UserType("RecursiveArray")
		recursiveMap   = root.UserType("RecursiveMap")
		composite      = root.UserType("Composite")
		customField    = root.UserType("CompositeWithCustomField")

		resultType = root.UserType("ResultType")
		rtCol      = root.UserType("ResultTypeCollection")

		simpleAlias    = root.UserType("SimpleAlias")
		nestedMapAlias = root.UserType("NestedMapAlias")
		arrayMapAlias  = root.UserType("ArrayMapAlias")
		stringAlias    = root.UserType("StringAlias")

		// primitive tyes
		stringT = expr.String

		// attribute contexts used in test cases
		defaultCtx    = NewAttributeContext(false, false, true, "", scope)
		defaultCtxPkg = NewAttributeContext(false, false, true, "mypkg", scope)
		pointerCtx    = NewAttributeContext(true, false, false, "", scope)
	)
	tc := map[string][]struct {
		Name      string
		Source    expr.DataType
		Target    expr.DataType
		SourceCtx *AttributeContext
		TargetCtx *AttributeContext
		Code      string
	}{
		// source and target type use default
		"source-target-type-use-default": {
			{"simple-to-simple", simple, simple, defaultCtx, defaultCtx, srcTgtUseDefaultSimpleToSimpleCode},
			{"simple-to-required", simple, required, defaultCtx, defaultCtx, srcTgtUseDefaultSimpleToRequiredCode},
			{"required-to-simple", required, simple, defaultCtx, defaultCtx, srcTgtUseDefaultRequiredToSimpleCode},
			{"simple-to-super", simple, super, defaultCtx, defaultCtx, srcTgtUseDefaultSimpleToSuperCode},
			{"super-to-simple", super, simple, defaultCtx, defaultCtx, srcTgtUseDefaultSuperToSimpleCode},
			{"simple-to-default", simple, defaultT, defaultCtx, defaultCtx, srcTgtUseDefaultSimpleToDefaultCode},
			{"default-to-simple", defaultT, simple, defaultCtx, defaultCtx, srcTgtUseDefaultDefaultToSimpleCode},

			// maps
			{"map-to-map", simpleMap, simpleMap, defaultCtx, defaultCtx, srcTgtUseDefaultMapToMapCode},
			{"map-to-required-map", simpleMap, requiredMap, defaultCtx, defaultCtx, srcTgtUseDefaultMapToRequiredMapCode},
			{"required-map-to-map", requiredMap, simpleMap, defaultCtx, defaultCtx, srcTgtUseDefaultRequiredMapToMapCode},
			{"map-to-default-map", simpleMap, defaultMap, defaultCtx, defaultCtx, srcTgtUseDefaultMapToDefaultMapCode},
			{"default-map-to-map", defaultMap, simpleMap, defaultCtx, defaultCtx, srcTgtUseDefaultDefaultMapToMapCode},
			{"required-map-to-default-map", requiredMap, defaultMap, defaultCtx, defaultCtx, srcTgtUseDefaultRequiredMapToDefaultMapCode},
			{"default-map-to-required-map", defaultMap, requiredMap, defaultCtx, defaultCtx, srcTgtUseDefaultDefaultMapToRequiredMapCode},
			{"nested-map-to-nested-map", nestedMap, nestedMap, defaultCtx, defaultCtx, srcTgtUseDefaultNestedMapToNestedMapCode},
			{"type-map-to-type-map", typeMap, typeMap, defaultCtx, defaultCtx, srcTgtUseDefaultTypeMapToTypeMapCode},
			{"array-map-to-array-map", arrayMap, arrayMap, defaultCtx, defaultCtx, srcTgtUseDefaultArrayMapToArrayMapCode},

			// arrays
			{"array-to-array", simpleArray, simpleArray, defaultCtx, defaultCtx, srcTgtUseDefaultArrayToArrayCode},
			{"array-to-required-array", simpleArray, requiredArray, defaultCtx, defaultCtx, srcTgtUseDefaultArrayToRequiredArrayCode},
			{"required-array-to-array", requiredArray, simpleArray, defaultCtx, defaultCtx, srcTgtUseDefaultRequiredArrayToArrayCode},
			{"array-to-default-array", simpleArray, defaultArray, defaultCtx, defaultCtx, srcTgtUseDefaultArrayToDefaultArrayCode},
			{"default-array-to-array", defaultArray, simpleArray, defaultCtx, defaultCtx, srcTgtUseDefaultDefaultArrayToArrayCode},
			{"required-array-to-default-array", requiredArray, defaultArray, defaultCtx, defaultCtx, srcTgtUseDefaultRequiredArrayToDefaultArrayCode},
			{"default-array-to-required-array", defaultArray, requiredArray, defaultCtx, defaultCtx, srcTgtUseDefaultDefaultArrayToRequiredArrayCode},
			{"nested-array-to-nested-array", nestedArray, nestedArray, defaultCtx, defaultCtx, srcTgtUseDefaultNestedArrayToNestedArrayCode},
			{"type-array-to-type-array", typeArray, typeArray, defaultCtx, defaultCtx, srcTgtUseDefaultTypeArrayToTypeArrayCode},
			{"map-array-to-map-array", mapArray, mapArray, defaultCtx, defaultCtx, srcTgtUseDefaultMapArrayToMapArrayCode},

			// others
			{"recursive-to-recursive", recursive, recursive, defaultCtx, defaultCtx, srcTgtUseDefaultRecursiveToRecursiveCode},
			{"recursive-array-to-recursive-array", recursiveArray, recursiveArray, defaultCtx, defaultCtx, srcTgtUseDefaultRecursiveArrayToRecursiveArrayCode},
			{"recursive-map-to-recursive-map", recursiveMap, recursiveMap, defaultCtx, defaultCtx, srcTgtUseDefaultRecursiveMapToRecursiveMapCode},
			{"composite-to-custom-field", composite, customField, defaultCtx, defaultCtx, srcTgtUseDefaultCompositeToCustomFieldCode},
			{"custom-field-to-composite", customField, composite, defaultCtx, defaultCtx, srcTgtUseDefaultCustomFieldToCompositeCode},
			{"composite-to-custom-field-pkg", composite, customField, defaultCtx, defaultCtxPkg, srcTgtUseDefaultCompositeToCustomFieldPkgCode},
			{"result-type-to-result-type", resultType, resultType, defaultCtx, defaultCtx, srcTgtUseDefaultResultTypeToResultTypeCode},
			{"result-type-collection-to-result-type-collection", rtCol, rtCol, defaultCtx, defaultCtx, srcTgtUseDefaultRTColToRTColCode},

			// alias
			{"simple-alias-to-simple", simpleAlias, simple, defaultCtx, defaultCtx, srcTgtUseDefaultSimpleAliasToSimpleCode},
			{"simple-to-simple-alias", simple, simpleAlias, defaultCtx, defaultCtx, srcTgtUseDefaultSimpleToSimpleAliasCode},
			{"nested-map-alias-to-nested-map", nestedMapAlias, nestedMap, defaultCtx, defaultCtx, srcTgtUseDefaultNestedMapAliasToNestedMapCode},
			{"nested-map-to-nested-map-alias", nestedMap, nestedMapAlias, defaultCtx, defaultCtx, srcTgtUseDefaultNestedMapToNestedMapAliasCode},
			{"array-map-alias-to-array-map", arrayMapAlias, arrayMap, defaultCtx, defaultCtx, srcTgtUseDefaultArrayMapAliasToArrayMapCode},
			{"array-map-to-array-map-alias", arrayMap, arrayMapAlias, defaultCtx, defaultCtx, srcTgtUseDefaultArrayMapToArrayMapAliasCode},
			{"string-to-string-alias", stringT, stringAlias, defaultCtx, defaultCtx, srcTgtUseDefaultStringToStringAliasCode},
			{"string-alias-to-string", stringAlias, stringT, defaultCtx, defaultCtx, srcTgtUseDefaultStringAliasToStringCode},
			{"string-alias-to-string-alias", stringAlias, stringAlias, defaultCtx, defaultCtx, srcTgtUseDefaultStringAliasToStringAliasCode},
		},

		// source type uses pointers for all fields, target type uses default
		"source-type-all-ptrs-target-type-uses-default": {
			{"simple-to-simple", simple, simple, pointerCtx, defaultCtx, srcAllPtrsTgtUseDefaultSimpleToSimpleCode},
			{"simple-to-required", simple, required, pointerCtx, defaultCtx, srcAllPtrsTgtUseDefaultSimpleToRequiredCode},
			{"required-to-simple", required, simple, pointerCtx, defaultCtx, srcAllPtrsTgtUseDefaultRequiredToSimpleCode},
			{"simple-to-super", simple, super, pointerCtx, defaultCtx, srcAllPtrsTgtUseDefaultSimpleToSuperCode},
			{"super-to-simple", super, simple, pointerCtx, defaultCtx, srcAllPtrsTgtUseDefaultSuperToSimpleCode},
			{"simple-to-default", simple, defaultT, pointerCtx, defaultCtx, srcAllPtrsTgtUseDefaultSimpleToDefaultCode},
			{"default-to-simple", defaultT, simple, pointerCtx, defaultCtx, srcAllPtrsTgtUseDefaultDefaultToSimpleCode},

			// maps
			{"required-map-to-map", requiredMap, simpleMap, pointerCtx, defaultCtx, srcAllPtrsTgtUseDefaultRequiredMapToMapCode},
			{"default-map-to-map", defaultMap, simpleMap, pointerCtx, defaultCtx, srcAllPtrsTgtUseDefaultDefaultMapToMapCode},
			{"required-map-to-default-map", requiredMap, defaultMap, pointerCtx, defaultCtx, srcAllPtrsTgtUseDefaultRequiredMapToDefaultMapCode},
			{"default-map-to-required-map", defaultMap, requiredMap, pointerCtx, defaultCtx, srcAllPtrsTgtUseDefaultDefaultMapToRequiredMapCode},

			// arrays
			{"default-array-to-array", defaultArray, simpleArray, pointerCtx, defaultCtx, srcAllPtrsTgtUseDefaultDefaultArrayToArrayCode},
			{"required-array-to-default-array", requiredArray, defaultArray, pointerCtx, defaultCtx, srcAllPtrsTgtUseDefaultRequiredArrayToDefaultArrayCode},
			{"default-array-to-required-array", defaultArray, requiredArray, pointerCtx, defaultCtx, srcAllPtrsTgtUseDefaultDefaultArrayToRequiredArrayCode},

			// others
			{"custom-field-to-composite", customField, composite, pointerCtx, defaultCtx, srcAllPtrsTgtUseDefaultCustomFieldToCompositeCode},

			// alias
			{"simple-alias-to-simple", simpleAlias, simple, pointerCtx, defaultCtx, srcAllPtrsTgtUseDefaultSimpleAliasToSimpleCode},
			{"simple-to-simple-alias", simple, simpleAlias, pointerCtx, defaultCtx, srcAllPtrsTgtUseDefaultSimpleToSimpleAliasCode},
		},

		// source type uses default, target type uses pointers for all fields
		"source-type-uses-default-target-type-all-ptrs": {
			{"simple-to-simple", simple, simple, defaultCtx, pointerCtx, srcUseDefaultTgtAllPtrsSimpleToSimpleCode},
			{"simple-to-required", simple, required, defaultCtx, pointerCtx, srcUseDefaultTgtAllPtrsSimpleToRequiredCode},
			{"required-to-simple", required, simple, defaultCtx, pointerCtx, srcUseDefaultTgtAllPtrsRequiredToSimpleCode},
			{"simple-to-default", simple, defaultT, defaultCtx, pointerCtx, srcUseDefaultTgtAllPtrsSimpleToDefaultCode},
			{"default-to-simple", defaultT, simple, defaultCtx, pointerCtx, srcUseDefaultTgtAllPtrsDefaultToSimpleCode},

			// maps
			{"map-to-default-map", simpleMap, defaultMap, defaultCtx, pointerCtx, srcUseDefaultTgtAllPtrsMapToDefaultMapCode},

			// arrays
			{"array-to-default-array", simpleArray, defaultArray, defaultCtx, pointerCtx, srcUseDefaultTgtAllPtrsArrayToDefaultArrayCode},

			// others
			{"recursive-to-recursive", recursive, recursive, defaultCtx, pointerCtx, srcUseDefaultTgtAllPtrsRecursiveToRecursiveCode},
			{"composite-to-custom-field", composite, customField, defaultCtx, pointerCtx, srcUseDefaultTgtAllPtrsCompositeToCustomFieldCode},
		},
	}
	for name, cases := range tc {
		t.Run(name, func(t *testing.T) {
			for _, c := range cases {
				t.Run(c.Name, func(t *testing.T) {
					if c.Source == nil {
						t.Fatal("source type not found in testdata")
					}
					if c.Target == nil {
						t.Fatal("target type not found in testdata")
					}
					code, _, err := GoTransform(&expr.AttributeExpr{Type: c.Source}, &expr.AttributeExpr{Type: c.Target}, "source", "target", c.SourceCtx, c.TargetCtx, "")
					if err != nil {
						t.Fatal(err)
					}
					code = FormatTestCode(t, "package foo\nfunc transform(){\n"+code+"}")
					if code != c.Code {
						t.Errorf("invalid code, got:\n%s\ngot vs. expected:\n%s", code, Diff(t, code, c.Code))
					}
				})
			}
		})
	}
}

const (
	srcTgtUseDefaultSimpleToSimpleCode = `func transform() {
	target := &Simple{
		RequiredString: source.RequiredString,
		DefaultBool:    source.DefaultBool,
		Integer:        source.Integer,
	}
}
`

	srcTgtUseDefaultSimpleToRequiredCode = `func transform() {
	target := &Required{
		RequiredString: source.RequiredString,
		DefaultBool:    source.DefaultBool,
	}
	if source.Integer != nil {
		target.Integer = *source.Integer
	}
}
`

	srcTgtUseDefaultRequiredToSimpleCode = `func transform() {
	target := &Simple{
		RequiredString: source.RequiredString,
		DefaultBool:    source.DefaultBool,
		Integer:        &source.Integer,
	}
}
`

	srcTgtUseDefaultSimpleToSuperCode = `func transform() {
	target := &Super{
		RequiredString: source.RequiredString,
		DefaultBool:    source.DefaultBool,
		Integer:        source.Integer,
	}
}
`

	srcTgtUseDefaultSuperToSimpleCode = `func transform() {
	target := &Simple{
		RequiredString: source.RequiredString,
		DefaultBool:    source.DefaultBool,
		Integer:        source.Integer,
	}
}
`

	srcTgtUseDefaultSimpleToDefaultCode = `func transform() {
	target := &Default{
		RequiredString: source.RequiredString,
		DefaultBool:    source.DefaultBool,
	}
	if source.Integer != nil {
		target.Integer = *source.Integer
	}
	if source.Integer == nil {
		target.Integer = 1
	}
}
`

	srcTgtUseDefaultDefaultToSimpleCode = `func transform() {
	target := &Simple{
		RequiredString: source.RequiredString,
		DefaultBool:    source.DefaultBool,
		Integer:        &source.Integer,
	}
}
`

	srcTgtUseDefaultMapToMapCode = `func transform() {
	target := &SimpleMap{}
	if source.Simple != nil {
		target.Simple = make(map[string]int, len(source.Simple))
		for key, val := range source.Simple {
			tk := key
			tv := val
			target.Simple[tk] = tv
		}
	}
}
`

	srcTgtUseDefaultMapToRequiredMapCode = `func transform() {
	target := &RequiredMap{}
	if source.Simple != nil {
		target.Simple = make(map[string]int, len(source.Simple))
		for key, val := range source.Simple {
			tk := key
			tv := val
			target.Simple[tk] = tv
		}
	}
}
`

	srcTgtUseDefaultRequiredMapToMapCode = `func transform() {
	target := &SimpleMap{}
	if source.Simple != nil {
		target.Simple = make(map[string]int, len(source.Simple))
		for key, val := range source.Simple {
			tk := key
			tv := val
			target.Simple[tk] = tv
		}
	}
}
`

	srcTgtUseDefaultMapToDefaultMapCode = `func transform() {
	target := &DefaultMap{}
	if source.Simple != nil {
		target.Simple = make(map[string]int, len(source.Simple))
		for key, val := range source.Simple {
			tk := key
			tv := val
			target.Simple[tk] = tv
		}
	}
	if source.Simple == nil {
		target.Simple = map[string]int{"foo": 1}
	}
}
`

	srcTgtUseDefaultDefaultMapToMapCode = `func transform() {
	target := &SimpleMap{}
	if source.Simple != nil {
		target.Simple = make(map[string]int, len(source.Simple))
		for key, val := range source.Simple {
			tk := key
			tv := val
			target.Simple[tk] = tv
		}
	}
}
`

	srcTgtUseDefaultRequiredMapToDefaultMapCode = `func transform() {
	target := &DefaultMap{}
	if source.Simple != nil {
		target.Simple = make(map[string]int, len(source.Simple))
		for key, val := range source.Simple {
			tk := key
			tv := val
			target.Simple[tk] = tv
		}
	}
}
`

	srcTgtUseDefaultDefaultMapToRequiredMapCode = `func transform() {
	target := &RequiredMap{}
	if source.Simple != nil {
		target.Simple = make(map[string]int, len(source.Simple))
		for key, val := range source.Simple {
			tk := key
			tv := val
			target.Simple[tk] = tv
		}
	}
}
`

	srcTgtUseDefaultNestedMapToNestedMapCode = `func transform() {
	target := &NestedMap{}
	if source.NestedMap != nil {
		target.NestedMap = make(map[float64]map[int]map[float64]uint64, len(source.NestedMap))
		for key, val := range source.NestedMap {
			tk := key
			tvc := make(map[int]map[float64]uint64, len(val))
			for key, val := range val {
				tk := key
				tvb := make(map[float64]uint64, len(val))
				for key, val := range val {
					tk := key
					tv := val
					tvb[tk] = tv
				}
				tvc[tk] = tvb
			}
			target.NestedMap[tk] = tvc
		}
	}
}
`

	srcTgtUseDefaultTypeMapToTypeMapCode = `func transform() {
	target := &TypeMap{}
	if source.TypeMap != nil {
		target.TypeMap = make(map[string]*SimpleMap, len(source.TypeMap))
		for key, val := range source.TypeMap {
			tk := key
			target.TypeMap[tk] = transformSimpleMapToSimpleMap(val)
		}
	}
}
`

	srcTgtUseDefaultArrayMapToArrayMapCode = `func transform() {
	target := &ArrayMap{}
	if source.ArrayMap != nil {
		target.ArrayMap = make(map[uint32][]float32, len(source.ArrayMap))
		for key, val := range source.ArrayMap {
			tk := key
			tv := make([]float32, len(val))
			for i, val := range val {
				tv[i] = val
			}
			target.ArrayMap[tk] = tv
		}
	}
}
`

	srcTgtUseDefaultArrayToArrayCode = `func transform() {
	target := &SimpleArray{}
	if source.StringArray != nil {
		target.StringArray = make([]string, len(source.StringArray))
		for i, val := range source.StringArray {
			target.StringArray[i] = val
		}
	}
}
`

	srcTgtUseDefaultArrayToRequiredArrayCode = `func transform() {
	target := &RequiredArray{}
	if source.StringArray != nil {
		target.StringArray = make([]string, len(source.StringArray))
		for i, val := range source.StringArray {
			target.StringArray[i] = val
		}
	}
}
`

	srcTgtUseDefaultRequiredArrayToArrayCode = `func transform() {
	target := &SimpleArray{}
	if source.StringArray != nil {
		target.StringArray = make([]string, len(source.StringArray))
		for i, val := range source.StringArray {
			target.StringArray[i] = val
		}
	}
}
`

	srcTgtUseDefaultArrayToDefaultArrayCode = `func transform() {
	target := &DefaultArray{}
	if source.StringArray != nil {
		target.StringArray = make([]string, len(source.StringArray))
		for i, val := range source.StringArray {
			target.StringArray[i] = val
		}
	}
	if source.StringArray == nil {
		target.StringArray = []string{"foo", "bar"}
	}
}
`

	srcTgtUseDefaultDefaultArrayToArrayCode = `func transform() {
	target := &SimpleArray{}
	if source.StringArray != nil {
		target.StringArray = make([]string, len(source.StringArray))
		for i, val := range source.StringArray {
			target.StringArray[i] = val
		}
	}
}
`

	srcTgtUseDefaultRequiredArrayToDefaultArrayCode = `func transform() {
	target := &DefaultArray{}
	if source.StringArray != nil {
		target.StringArray = make([]string, len(source.StringArray))
		for i, val := range source.StringArray {
			target.StringArray[i] = val
		}
	}
}
`

	srcTgtUseDefaultDefaultArrayToRequiredArrayCode = `func transform() {
	target := &RequiredArray{}
	if source.StringArray != nil {
		target.StringArray = make([]string, len(source.StringArray))
		for i, val := range source.StringArray {
			target.StringArray[i] = val
		}
	}
}
`

	srcTgtUseDefaultNestedArrayToNestedArrayCode = `func transform() {
	target := &NestedArray{}
	if source.NestedArray != nil {
		target.NestedArray = make([][][]float64, len(source.NestedArray))
		for i, val := range source.NestedArray {
			target.NestedArray[i] = make([][]float64, len(val))
			for j, val := range val {
				target.NestedArray[i][j] = make([]float64, len(val))
				for k, val := range val {
					target.NestedArray[i][j][k] = val
				}
			}
		}
	}
}
`

	srcTgtUseDefaultTypeArrayToTypeArrayCode = `func transform() {
	target := &TypeArray{}
	if source.TypeArray != nil {
		target.TypeArray = make([]*SimpleArray, len(source.TypeArray))
		for i, val := range source.TypeArray {
			target.TypeArray[i] = transformSimpleArrayToSimpleArray(val)
		}
	}
}
`

	srcTgtUseDefaultMapArrayToMapArrayCode = `func transform() {
	target := &MapArray{}
	if source.MapArray != nil {
		target.MapArray = make([]map[int]string, len(source.MapArray))
		for i, val := range source.MapArray {
			target.MapArray[i] = make(map[int]string, len(val))
			for key, val := range val {
				tk := key
				tv := val
				target.MapArray[i][tk] = tv
			}
		}
	}
}
`

	srcTgtUseDefaultRecursiveToRecursiveCode = `func transform() {
	target := &Recursive{
		RequiredString: source.RequiredString,
	}
	if source.Recursive != nil {
		target.Recursive = transformRecursiveToRecursive(source.Recursive)
	}
}
`

	srcTgtUseDefaultRecursiveArrayToRecursiveArrayCode = `func transform() {
	target := &RecursiveArray{
		RequiredString: source.RequiredString,
	}
	if source.Recursive != nil {
		target.Recursive = make([]*RecursiveArray, len(source.Recursive))
		for i, val := range source.Recursive {
			target.Recursive[i] = transformRecursiveArrayToRecursiveArray(val)
		}
	}
}
`

	srcTgtUseDefaultRecursiveMapToRecursiveMapCode = `func transform() {
	target := &RecursiveMap{
		RequiredString: source.RequiredString,
	}
	if source.Recursive != nil {
		target.Recursive = make(map[string]*RecursiveMap, len(source.Recursive))
		for key, val := range source.Recursive {
			tk := key
			target.Recursive[tk] = transformRecursiveMapToRecursiveMap(val)
		}
	}
}
`

	srcTgtUseDefaultCompositeToCustomFieldCode = `func transform() {
	target := &CompositeWithCustomField{}
	if source.RequiredString != nil {
		target.MyString = *source.RequiredString
	}
	if source.DefaultInt != nil {
		target.MyInt = *source.DefaultInt
	}
	if source.DefaultInt == nil {
		target.MyInt = 100
	}
	if source.Type != nil {
		target.MyType = transformSimpleToSimple(source.Type)
	}
	if source.Map != nil {
		target.MyMap = make(map[int]string, len(source.Map))
		for key, val := range source.Map {
			tk := key
			tv := val
			target.MyMap[tk] = tv
		}
	}
	if source.Array != nil {
		target.MyArray = make([]string, len(source.Array))
		for i, val := range source.Array {
			target.MyArray[i] = val
		}
	}
}
`

	srcTgtUseDefaultCustomFieldToCompositeCode = `func transform() {
	target := &Composite{
		RequiredString: &source.MyString,
		DefaultInt:     &source.MyInt,
	}
	if source.MyType != nil {
		target.Type = transformSimpleToSimple(source.MyType)
	}
	if source.MyMap != nil {
		target.Map = make(map[int]string, len(source.MyMap))
		for key, val := range source.MyMap {
			tk := key
			tv := val
			target.Map[tk] = tv
		}
	}
	if source.MyArray != nil {
		target.Array = make([]string, len(source.MyArray))
		for i, val := range source.MyArray {
			target.Array[i] = val
		}
	}
}
`

	srcTgtUseDefaultCompositeToCustomFieldPkgCode = `func transform() {
	target := &mypkg.CompositeWithCustomField{}
	if source.RequiredString != nil {
		target.MyString = *source.RequiredString
	}
	if source.DefaultInt != nil {
		target.MyInt = *source.DefaultInt
	}
	if source.DefaultInt == nil {
		target.MyInt = 100
	}
	if source.Type != nil {
		target.MyType = transformSimpleToMypkgSimple(source.Type)
	}
	if source.Map != nil {
		target.MyMap = make(map[int]string, len(source.Map))
		for key, val := range source.Map {
			tk := key
			tv := val
			target.MyMap[tk] = tv
		}
	}
	if source.Array != nil {
		target.MyArray = make([]string, len(source.Array))
		for i, val := range source.Array {
			target.MyArray[i] = val
		}
	}
}
`

	srcTgtUseDefaultResultTypeToResultTypeCode = `func transform() {
	target := &ResultType{
		Int: source.Int,
	}
	if source.Map != nil {
		target.Map = make(map[int]string, len(source.Map))
		for key, val := range source.Map {
			tk := key
			tv := val
			target.Map[tk] = tv
		}
	}
}
`

	srcTgtUseDefaultRTColToRTColCode = `func transform() {
	target := &ResultTypeCollection{}
	if source.Collection != nil {
		target.Collection = make([]*ResultType, len(source.Collection))
		for i, val := range source.Collection {
			target.Collection[i] = transformResultTypeToResultType(val)
		}
	}
}
`

	srcTgtUseDefaultSimpleAliasToSimpleCode = `func transform() {
	target := &Simple{
		RequiredString: string(source.RequiredString),
		DefaultBool:    bool(source.DefaultBool),
	}
	if source.Integer != nil {
		integer := int(*source.Integer)
		target.Integer = &integer
	}
}
`

	srcTgtUseDefaultSimpleToSimpleAliasCode = `func transform() {
	target := &SimpleAlias{
		RequiredString: StringAlias(source.RequiredString),
		DefaultBool:    BoolAlias(source.DefaultBool),
	}
	if source.Integer != nil {
		integer := IntAlias(*source.Integer)
		target.Integer = &integer
	}
}
`

	srcTgtUseDefaultNestedMapAliasToNestedMapCode = `func transform() {
	target := &NestedMap{}
	if source.NestedMap != nil {
		target.NestedMap = make(map[float64]map[int]map[float64]uint64, len(source.NestedMap))
		for key, val := range source.NestedMap {
			tk := float64(key)
			tvc := make(map[int]map[float64]uint64, len(val))
			for key, val := range val {
				tk := int(key)
				tvb := make(map[float64]uint64, len(val))
				for key, val := range val {
					tk := float64(key)
					tv := val
					tvb[tk] = tv
				}
				tvc[tk] = tvb
			}
			target.NestedMap[tk] = tvc
		}
	}
}
`

	srcTgtUseDefaultNestedMapToNestedMapAliasCode = `func transform() {
	target := &NestedMapAlias{}
	if source.NestedMap != nil {
		target.NestedMap = make(map[Float64Alias]map[IntAlias]map[Float64Alias]uint64, len(source.NestedMap))
		for key, val := range source.NestedMap {
			tk := Float64Alias(key)
			tvc := make(map[IntAlias]map[Float64Alias]uint64, len(val))
			for key, val := range val {
				tk := IntAlias(key)
				tvb := make(map[Float64Alias]uint64, len(val))
				for key, val := range val {
					tk := Float64Alias(key)
					tv := val
					tvb[tk] = tv
				}
				tvc[tk] = tvb
			}
			target.NestedMap[tk] = tvc
		}
	}
}
`

	srcTgtUseDefaultArrayMapAliasToArrayMapCode = `func transform() {
	target := &ArrayMap{}
	if source.ArrayMap != nil {
		target.ArrayMap = make(map[uint32][]float32, len(source.ArrayMap))
		for key, val := range source.ArrayMap {
			tk := key
			tv := make([]float32, len(val))
			for i, val := range val {
				tv[i] = float32(val)
			}
			target.ArrayMap[tk] = tv
		}
	}
}
`

	srcTgtUseDefaultArrayMapToArrayMapAliasCode = `func transform() {
	target := &ArrayMapAlias{}
	if source.ArrayMap != nil {
		target.ArrayMap = make(map[uint32]Float32ArrayAlias, len(source.ArrayMap))
		for key, val := range source.ArrayMap {
			tk := key
			tv := make([]Float32Alias, len(val))
			for i, val := range val {
				tv[i] = Float32Alias(val)
			}
			target.ArrayMap[tk] = tv
		}
	}
}
`

	srcTgtUseDefaultStringToStringAliasCode = `func transform() {
	target := StringAlias(source)
}
`

	srcTgtUseDefaultStringAliasToStringCode = `func transform() {
	target := string(source)
}
`

	srcTgtUseDefaultStringAliasToStringAliasCode = `func transform() {
	target := source
}
`

	srcAllPtrsTgtUseDefaultSimpleToSimpleCode = `func transform() {
	target := &Simple{
		RequiredString: *source.RequiredString,
		Integer:        source.Integer,
	}
	if source.DefaultBool != nil {
		target.DefaultBool = *source.DefaultBool
	}
	if source.DefaultBool == nil {
		target.DefaultBool = true
	}
}
`

	srcAllPtrsTgtUseDefaultSimpleToRequiredCode = `func transform() {
	target := &Required{
		RequiredString: *source.RequiredString,
	}
	if source.DefaultBool != nil {
		target.DefaultBool = *source.DefaultBool
	}
	if source.Integer != nil {
		target.Integer = *source.Integer
	}
	if source.DefaultBool == nil {
		target.DefaultBool = true
	}
}
`

	srcAllPtrsTgtUseDefaultRequiredToSimpleCode = `func transform() {
	target := &Simple{
		RequiredString: *source.RequiredString,
		DefaultBool:    *source.DefaultBool,
		Integer:        source.Integer,
	}
}
`

	srcAllPtrsTgtUseDefaultSimpleToSuperCode = `func transform() {
	target := &Super{
		RequiredString: *source.RequiredString,
		Integer:        source.Integer,
	}
	if source.DefaultBool != nil {
		target.DefaultBool = *source.DefaultBool
	}
	if source.DefaultBool == nil {
		target.DefaultBool = true
	}
}
`

	srcAllPtrsTgtUseDefaultSuperToSimpleCode = `func transform() {
	target := &Simple{
		RequiredString: *source.RequiredString,
		Integer:        source.Integer,
	}
	if source.DefaultBool != nil {
		target.DefaultBool = *source.DefaultBool
	}
	if source.DefaultBool == nil {
		target.DefaultBool = true
	}
}
`

	srcAllPtrsTgtUseDefaultSimpleToDefaultCode = `func transform() {
	target := &Default{
		RequiredString: *source.RequiredString,
	}
	if source.DefaultBool != nil {
		target.DefaultBool = *source.DefaultBool
	}
	if source.Integer != nil {
		target.Integer = *source.Integer
	}
	if source.DefaultBool == nil {
		target.DefaultBool = true
	}
	if source.Integer == nil {
		target.Integer = 1
	}
}
`

	srcAllPtrsTgtUseDefaultDefaultToSimpleCode = `func transform() {
	target := &Simple{
		Integer: source.Integer,
	}
	if source.RequiredString != nil {
		target.RequiredString = *source.RequiredString
	}
	if source.DefaultBool != nil {
		target.DefaultBool = *source.DefaultBool
	}
	if source.DefaultBool == nil {
		target.DefaultBool = true
	}
}
`

	srcAllPtrsTgtUseDefaultRequiredMapToMapCode = `func transform() {
	target := &SimpleMap{}
	target.Simple = make(map[string]int, len(source.Simple))
	for key, val := range source.Simple {
		tk := key
		tv := val
		target.Simple[tk] = tv
	}
}
`

	srcAllPtrsTgtUseDefaultDefaultMapToMapCode = `func transform() {
	target := &SimpleMap{}
	if source.Simple != nil {
		target.Simple = make(map[string]int, len(source.Simple))
		for key, val := range source.Simple {
			tk := key
			tv := val
			target.Simple[tk] = tv
		}
	}
}
`

	srcAllPtrsTgtUseDefaultRequiredMapToDefaultMapCode = `func transform() {
	target := &DefaultMap{}
	target.Simple = make(map[string]int, len(source.Simple))
	for key, val := range source.Simple {
		tk := key
		tv := val
		target.Simple[tk] = tv
	}
}
`

	srcAllPtrsTgtUseDefaultDefaultMapToRequiredMapCode = `func transform() {
	target := &RequiredMap{}
	if source.Simple != nil {
		target.Simple = make(map[string]int, len(source.Simple))
		for key, val := range source.Simple {
			tk := key
			tv := val
			target.Simple[tk] = tv
		}
	}
}
`

	srcAllPtrsTgtUseDefaultDefaultArrayToArrayCode = `func transform() {
	target := &SimpleArray{}
	if source.StringArray != nil {
		target.StringArray = make([]string, len(source.StringArray))
		for i, val := range source.StringArray {
			target.StringArray[i] = val
		}
	}
}
`

	srcAllPtrsTgtUseDefaultRequiredArrayToDefaultArrayCode = `func transform() {
	target := &DefaultArray{}
	target.StringArray = make([]string, len(source.StringArray))
	for i, val := range source.StringArray {
		target.StringArray[i] = val
	}
}
`

	srcAllPtrsTgtUseDefaultDefaultArrayToRequiredArrayCode = `func transform() {
	target := &RequiredArray{}
	if source.StringArray != nil {
		target.StringArray = make([]string, len(source.StringArray))
		for i, val := range source.StringArray {
			target.StringArray[i] = val
		}
	}
}
`

	srcAllPtrsTgtUseDefaultCustomFieldToCompositeCode = `func transform() {
	target := &Composite{
		RequiredString: source.MyString,
		DefaultInt:     source.MyInt,
	}
	target.Type = transformSimpleToSimple(source.MyType)
	target.Map = make(map[int]string, len(source.MyMap))
	for key, val := range source.MyMap {
		tk := key
		tv := val
		target.Map[tk] = tv
	}
	target.Array = make([]string, len(source.MyArray))
	for i, val := range source.MyArray {
		target.Array[i] = val
	}
}
`

	srcAllPtrsTgtUseDefaultSimpleAliasToSimpleCode = `func transform() {
	target := &Simple{
		RequiredString: string(*source.RequiredString),
	}
	if source.DefaultBool != nil {
		target.DefaultBool = bool(*source.DefaultBool)
	}
	if source.Integer != nil {
		integer := int(*source.Integer)
		target.Integer = &integer
	}
	if source.DefaultBool == nil {
		target.DefaultBool = true
	}
}
`

	srcAllPtrsTgtUseDefaultSimpleToSimpleAliasCode = `func transform() {
	target := &SimpleAlias{
		RequiredString: StringAlias(*source.RequiredString),
	}
	if source.DefaultBool != nil {
		target.DefaultBool = BoolAlias(*source.DefaultBool)
	}
	if source.Integer != nil {
		integer := IntAlias(*source.Integer)
		target.Integer = &integer
	}
	if source.DefaultBool == nil {
		target.DefaultBool = true
	}
}
`

	srcUseDefaultTgtAllPtrsSimpleToSimpleCode = `func transform() {
	target := &Simple{
		RequiredString: &source.RequiredString,
		DefaultBool:    &source.DefaultBool,
		Integer:        source.Integer,
	}
}
`

	srcUseDefaultTgtAllPtrsSimpleToRequiredCode = `func transform() {
	target := &Required{
		RequiredString: &source.RequiredString,
		DefaultBool:    &source.DefaultBool,
		Integer:        source.Integer,
	}
}
`

	srcUseDefaultTgtAllPtrsRequiredToSimpleCode = `func transform() {
	target := &Simple{
		RequiredString: &source.RequiredString,
		DefaultBool:    &source.DefaultBool,
		Integer:        &source.Integer,
	}
}
`

	srcUseDefaultTgtAllPtrsSimpleToDefaultCode = `func transform() {
	target := &Default{
		RequiredString: &source.RequiredString,
		DefaultBool:    &source.DefaultBool,
		Integer:        source.Integer,
	}
}
`

	srcUseDefaultTgtAllPtrsDefaultToSimpleCode = `func transform() {
	target := &Simple{
		RequiredString: &source.RequiredString,
		DefaultBool:    &source.DefaultBool,
		Integer:        &source.Integer,
	}
}
`

	srcUseDefaultTgtAllPtrsMapToDefaultMapCode = `func transform() {
	target := &DefaultMap{}
	if source.Simple != nil {
		target.Simple = make(map[string]int, len(source.Simple))
		for key, val := range source.Simple {
			tk := key
			tv := val
			target.Simple[tk] = tv
		}
	}
}
`

	srcUseDefaultTgtAllPtrsArrayToDefaultArrayCode = `func transform() {
	target := &DefaultArray{}
	if source.StringArray != nil {
		target.StringArray = make([]string, len(source.StringArray))
		for i, val := range source.StringArray {
			target.StringArray[i] = val
		}
	}
}
`

	srcUseDefaultTgtAllPtrsRecursiveToRecursiveCode = `func transform() {
	target := &Recursive{
		RequiredString: &source.RequiredString,
	}
	if source.Recursive != nil {
		target.Recursive = transformRecursiveToRecursive(source.Recursive)
	}
}
`

	srcUseDefaultTgtAllPtrsCompositeToCustomFieldCode = `func transform() {
	target := &CompositeWithCustomField{
		MyString: source.RequiredString,
		MyInt:    source.DefaultInt,
	}
	if source.Type != nil {
		target.MyType = transformSimpleToSimple(source.Type)
	}
	if source.Map != nil {
		target.MyMap = make(map[int]string, len(source.Map))
		for key, val := range source.Map {
			tk := key
			tv := val
			target.MyMap[tk] = tv
		}
	}
	if source.Array != nil {
		target.MyArray = make([]string, len(source.Array))
		for i, val := range source.Array {
			target.MyArray[i] = val
		}
	}
}
`
)
