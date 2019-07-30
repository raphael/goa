package testdata

import . "goa.design/goa/v3/dsl"

var ValidationTypesDSL = func() {
	var (
		IntegerT = Type("Integer", func() {
			Attribute("required_integer", Int, func() {
				Minimum(1)
			})
			Attribute("default_integer", Int32, func() {
				Enum(1, 5, 10, 100)
				Default(5)
			})
			Attribute("integer", Int64, func() {
				Maximum(100)
			})
			Required("required_integer")
		})

		FloatT = Type("Float", func() {
			Attribute("required_float", Float32, func() {
				Minimum(1.0)
			})
			Attribute("default_integer", Float32, func() {
				Enum(1.2, 5, 10, 100.8)
				Default(5.0)
			})
			Attribute("float64", Float64, func() {
				Maximum(100.1)
			})
			Required("required_float")
		})

		StringT = Type("String", func() {
			Attribute("required_string", String, func() {
				MinLength(1)
				MaxLength(10)
				Pattern("^[A-z].*[a-z]$")
			})
			Attribute("default_string", String, func() {
				Enum("foo", "bar")
				Default("foo")
			})
			Attribute("string", String, func() {
				Format(FormatDateTime)
			})
			Required("required_string")
		})

		_ = Type("UserType", func() {
			Attribute("required_integer", IntegerT)
			Attribute("default_string", StringT, func() {
				Default(struct{ RequiredString, DefaultString, String string }{RequiredString: "Atoz", DefaultString: "bar", String: "2018-12-18T13:22:53.108Z"})
			})
			Attribute("float", FloatT)
			Required("required_integer")
		})

		_ = Type("ArrayUserType", func() {
			Attribute("array", ArrayOf(FloatT))
		})

		_ = Type("Array", func() {
			Attribute("required_array", ArrayOf(Int), func() {
				MinLength(5)
			})
			Attribute("default_array", ArrayOf(String), func() {
				MaxLength(3)
				Default([]string{"foo", "bar"})
			})
			Attribute("array", ArrayOf(UInt64), func() {
				Elem(func() {
					Enum(0, 1, 1, 2, 3, 5)
				})
			})
			Required("required_array")
		})

		_ = Type("Map", func() {
			Attribute("required_map", MapOf(Int, String), func() {
				MinLength(5)
			})
			Attribute("default_map", MapOf(String, String), func() {
				MaxLength(3)
				Default(map[string]string{"foo": "bar"})
			})
			Attribute("map", MapOf(String, UInt64), func() {
				Key(func() {
					Pattern("^[A-Z]")
				})
				Elem(func() {
					Maximum(5)
				})
			})
			Required("required_map")
		})

		Result = ResultType("application/vnd.goa.result", func() {
			TypeName("Result")
			Attributes(func() {
				Attribute("required", Int, func() {
					Minimum(10)
				})
			})
		})

		_ = Type("Collection", CollectionOf(Result))

		_ = ResultType("application/vnd.goa.collection", func() {
			TypeName("TypeWithCollection")
			Attributes(func() {
				Attribute("collection", CollectionOf(Result))
			})
		})
	)
}
