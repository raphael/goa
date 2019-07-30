package testdata

import . "goa.design/goa/v3/dsl"

var TestTypesDSL = func() {
	var (
		Simple = Type("Simple", func() {
			Attribute("required_string", String)
			Attribute("default_bool", Boolean, func() {
				Default(true)
			})
			Attribute("integer", Int)
			Required("required_string")
		})

		_ = Type("Required", func() {
			Extend(Simple)
			Required("required_string", "default_bool", "integer")
		})

		_ = Type("Super", func() {
			Extend(Simple)
			Attribute("ignored_attr", Float32)
		})

		_ = Type("Default", func() {
			Attribute("required_string", String, func() {
				Default("foo")
			})
			Attribute("default_bool", Boolean, func() {
				Default(true)
			})
			Attribute("integer", Int, func() {
				Default(1)
			})
			Required("integer")
		})

		SimpleMap = Type("SimpleMap", func() {
			Attribute("simple", MapOf(String, Int))
		})

		_ = Type("RequiredMap", func() {
			Extend(SimpleMap)
			Required("simple")
		})

		_ = Type("DefaultMap", func() {
			Attribute("simple", MapOf(String, Int), func() {
				Default(map[string]int{"foo": 1})
			})
		})

		_ = Type("NestedMap", func() {
			Attribute("nested_map", MapOf(Float64, MapOf(Int, MapOf(Float64, UInt64))))
		})

		_ = Type("TypeMap", func() {
			Attribute("type_map", MapOf(String, SimpleMap))
		})

		_ = Type("ArrayMap", func() {
			Attribute("array_map", MapOf(UInt32, ArrayOf(Float32)))
		})

		SimpleArray = Type("SimpleArray", func() {
			Attribute("string_array", ArrayOf(String))
		})

		_ = Type("RequiredArray", func() {
			Extend(SimpleArray)
			Required("string_array")
		})

		_ = Type("DefaultArray", func() {
			Attribute("string_array", ArrayOf(String), func() {
				Default([]string{"foo", "bar"})
			})
		})

		_ = Type("NestedArray", func() {
			Attribute("nested_array", ArrayOf(ArrayOf(ArrayOf(Float64))))
		})

		_ = Type("TypeArray", func() {
			Attribute("type_array", ArrayOf(SimpleArray))
		})

		_ = Type("MapArray", func() {
			Attribute("map_array", ArrayOf(MapOf(Int, String)))
		})

		Composite = Type("Composite", func() {
			Attribute("required_string", String)
			Attribute("default_int", Int)
			Attribute("type", Simple)
			Attribute("map", MapOf(Int, String))
			Attribute("array", ArrayOf(String))
		})

		_ = Type("Deep", func() {
			Attribute("string", String)
			Attribute("inner", Composite)
		})

		_ = Type("DeepArray", func() {
			Attribute("string", String)
			Attribute("inner", ArrayOf(Composite))
		})

		_ = Type("CompositeWithCustomField", func() {
			Attribute("required_string", String, func() {
				Meta("struct:field:name", "my_string")
			})
			Attribute("default_int", Int, func() {
				Meta("struct:field:name", "my_int")
				Default(100)
			})
			Attribute("type", Simple, func() {
				Meta("struct:field:name", "my_type")
			})
			Attribute("map", MapOf(Int, String), func() {
				Meta("struct:field:name", "my_map")
			})
			Attribute("array", ArrayOf(String), func() {
				Meta("struct:field:name", "my_array")
			})
			Required("required_string", "type", "map", "array")
		})

		_ = Type("Recursive", func() {
			Attribute("required_string", String)
			Attribute("recursive", "Recursive")
			Required("required_string")
		})

		_ = Type("RecursiveArray", func() {
			Attribute("required_string", String)
			Attribute("recursive", ArrayOf("RecursiveArray"))
			Required("required_string")
		})

		_ = Type("RecursiveMap", func() {
			Attribute("required_string", String)
			Attribute("recursive", MapOf(String, "RecursiveMap"))
			Required("required_string")
		})

		RT = ResultType("application/vnd.goa.example", func() {
			TypeName("ResultType")
			Attributes(func() {
				Attribute("int", Int)
				Attribute("map", MapOf(Int, String))
			})
		})

		_ = ResultType("application/vnd.goa.collection", func() {
			TypeName("ResultTypeCollection")
			Attributes(func() {
				Attribute("collection", CollectionOf(RT))
			})
		})

		_ = Type("Optional", func() {
			Attribute("int", Int)
			Attribute("uint", UInt)
			Attribute("float", Float32)
			Attribute("string", String)
			Attribute("bytes", Bytes)
			Attribute("any", Any)
			Attribute("array", ArrayOf(String))
			Attribute("map", MapOf(Int, String))
			Attribute("user_type", "Optional")
		})

		_ = Type("WithDefaults", func() {
			Attribute("int", Int, func() {
				Default(100)
			})
			Attribute("required_int", Int, func() {
				Default(99)
			})
			Attribute("string", String, func() {
				Default("foo")
			})
			Attribute("required_string", String, func() {
				Default("bar")
			})
			Attribute("bytes", Bytes, func() {
				Default([]byte("foobar"))
			})
			Attribute("required_bytes", Bytes, func() {
				Default([]byte("foobar_required"))
			})
			Attribute("any", Any, func() {
				Default("something")
			})
			Attribute("required_any", Any, func() {
				Default("anything")
			})
			Attribute("array", ArrayOf(String), func() {
				Default([]string{"foo", "bar"})
			})
			Attribute("required_array", ArrayOf(String), func() {
				Default([]string{"bar", "foo"})
			})
			Attribute("map", MapOf(Int, String), func() {
				Default(map[int]string{1: "foo"})
			})
			Attribute("required_map", MapOf(Int, String), func() {
				Default(map[int]string{2: "bar"})
			})
			Required("required_int", "required_string", "required_bytes", "required_any", "required_array", "required_map")
		})

		StringAlias = Type("StringAlias", String)
		BoolAlias   = Type("BoolAlias", Boolean, func() {
			Default(true)
		})
		IntAlias          = Type("IntAlias", Int)
		Float32Alias      = Type("Float32Alias", Float32)
		Float64Alias      = Type("Float64Alias", Float64)
		Float32ArrayAlias = Type("Float32ArrayAlias", ArrayOf(Float32Alias))
		NestedMapAlias    = Type("MapAlias", MapOf(Float64Alias, MapOf(IntAlias, MapOf(Float64Alias, UInt64))))
		ArrayMapAlias     = Type("MapWithArrayAlias", MapOf(UInt32, Float32ArrayAlias))

		_ = Type("SimpleAlias", func() {
			Attribute("required_string", StringAlias)
			Attribute("default_bool", BoolAlias)
			Attribute("integer", IntAlias)
			Required("required_string")
		})

		_ = Type("NestedMapAlias", func() {
			Attribute("nested_map", NestedMapAlias)
		})

		_ = Type("ArrayMapAlias", func() {
			Attribute("array_map", ArrayMapAlias)
		})
	)
}
