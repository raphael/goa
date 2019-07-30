package testdata

import (
	. "goa.design/goa/v3/dsl"
)

// The DSL function names follow the following pattern:
//
// (Payload|Result)(Query|Path|Body)+(Type)(Validate)?DSL
//
// Where Type is the type of the payload or result.

var PayloadQueryBoolDSL = func() {
	Service("ServiceQueryBool", func() {
		Method("MethodQueryBool", func() {
			Payload(func() {
				Attribute("q", Boolean)
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryBoolValidateDSL = func() {
	Service("ServiceQueryBoolValidate", func() {
		Method("MethodQueryBoolValidate", func() {
			Payload(func() {
				Attribute("q", Boolean, func() {
					Enum(true)
				})
				Required("q")
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryIntDSL = func() {
	Service("ServiceQueryInt", func() {
		Method("MethodQueryInt", func() {
			Payload(func() {
				Attribute("q", Int)
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryIntValidateDSL = func() {
	Service("ServiceQueryIntValidate", func() {
		Method("MethodQueryIntValidate", func() {
			Payload(func() {
				Attribute("q", Int, func() {
					Minimum(1)
				})
				Required("q")
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryInt32DSL = func() {
	Service("ServiceQueryInt32", func() {
		Method("MethodQueryInt32", func() {
			Payload(func() {
				Attribute("q", Int32)
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryInt32ValidateDSL = func() {
	Service("ServiceQueryInt32Validate", func() {
		Method("MethodQueryInt32Validate", func() {
			Payload(func() {
				Attribute("q", Int32, func() {
					Minimum(1)
				})
				Required("q")
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryInt64DSL = func() {
	Service("ServiceQueryInt64", func() {
		Method("MethodQueryInt64", func() {
			Payload(func() {
				Attribute("q", Int64)
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryInt64ValidateDSL = func() {
	Service("ServiceQueryInt64Validate", func() {
		Method("MethodQueryInt64Validate", func() {
			Payload(func() {
				Attribute("q", Int64, func() {
					Minimum(1)
				})
				Required("q")
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryUIntDSL = func() {
	Service("ServiceQueryUInt", func() {
		Method("MethodQueryUInt", func() {
			Payload(func() {
				Attribute("q", UInt)
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryUIntValidateDSL = func() {
	Service("ServiceQueryUIntValidate", func() {
		Method("MethodQueryUIntValidate", func() {
			Payload(func() {
				Attribute("q", UInt, func() {
					Minimum(1)
				})
				Required("q")
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryUInt32DSL = func() {
	Service("ServiceQueryUInt32", func() {
		Method("MethodQueryUInt32", func() {
			Payload(func() {
				Attribute("q", UInt32)
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryUInt32ValidateDSL = func() {
	Service("ServiceQueryUInt32Validate", func() {
		Method("MethodQueryUInt32Validate", func() {
			Payload(func() {
				Attribute("q", UInt32, func() {
					Minimum(1)
				})
				Required("q")
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryUInt64DSL = func() {
	Service("ServiceQueryUInt64", func() {
		Method("MethodQueryUInt64", func() {
			Payload(func() {
				Attribute("q", UInt64)
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryUInt64ValidateDSL = func() {
	Service("ServiceQueryUInt64Validate", func() {
		Method("MethodQueryUInt64Validate", func() {
			Payload(func() {
				Attribute("q", UInt64, func() {
					Minimum(1)
				})
				Required("q")
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryFloat32DSL = func() {
	Service("ServiceQueryFloat32", func() {
		Method("MethodQueryFloat32", func() {
			Payload(func() {
				Attribute("q", Float32)
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryFloat32ValidateDSL = func() {
	Service("ServiceQueryFloat32Validate", func() {
		Method("MethodQueryFloat32Validate", func() {
			Payload(func() {
				Attribute("q", Float32, func() {
					Minimum(1)
				})
				Required("q")
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryFloat64DSL = func() {
	Service("ServiceQueryFloat64", func() {
		Method("MethodQueryFloat64", func() {
			Payload(func() {
				Attribute("q", Float64)
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryFloat64ValidateDSL = func() {
	Service("ServiceQueryFloat64Validate", func() {
		Method("MethodQueryFloat64Validate", func() {
			Payload(func() {
				Attribute("q", Float64, func() {
					Minimum(1)
				})
				Required("q")
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryStringDSL = func() {
	Service("ServiceQueryString", func() {
		Method("MethodQueryString", func() {
			Payload(func() {
				Attribute("q", String)
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryStringValidateDSL = func() {
	Service("ServiceQueryStringValidate", func() {
		Method("MethodQueryStringValidate", func() {
			Payload(func() {
				Attribute("q", String, func() {
					Enum("val")
				})
				Required("q")
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryStringNotRequiredValidateDSL = func() {
	Service("ServiceQueryStringNotRequiredValidate", func() {
		Method("MethodQueryStringNotRequiredValidate", func() {
			Payload(func() {
				Attribute("q", String, func() {
					Enum("val")
				})
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryBytesDSL = func() {
	Service("ServiceQueryBytes", func() {
		Method("MethodQueryBytes", func() {
			Payload(func() {
				Attribute("q", Bytes)
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryBytesValidateDSL = func() {
	Service("ServiceQueryBytesValidate", func() {
		Method("MethodQueryBytesValidate", func() {
			Payload(func() {
				Attribute("q", Bytes, func() {
					MinLength(1)
				})
				Required("q")
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryAnyDSL = func() {
	Service("ServiceQueryAny", func() {
		Method("MethodQueryAny", func() {
			Payload(func() {
				Attribute("q", Any)
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryAnyValidateDSL = func() {
	Service("ServiceQueryAnyValidate", func() {
		Method("MethodQueryAnyValidate", func() {
			Payload(func() {
				Attribute("q", Any, func() {
					Enum("val", 1)
				})
				Required("q")
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryArrayBoolDSL = func() {
	Service("ServiceQueryArrayBool", func() {
		Method("MethodQueryArrayBool", func() {
			Payload(func() {
				Attribute("q", ArrayOf(Boolean))
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryArrayBoolValidateDSL = func() {
	Service("ServiceQueryArrayBoolValidate", func() {
		Method("MethodQueryArrayBoolValidate", func() {
			Payload(func() {
				Attribute("q", ArrayOf(Boolean), func() {
					MinLength(1)
					Elem(func() {
						Enum(true)
					})
				})
				Required("q")
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryArrayIntDSL = func() {
	Service("ServiceQueryArrayInt", func() {
		Method("MethodQueryArrayInt", func() {
			Payload(func() {
				Attribute("q", ArrayOf(Int))
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryArrayIntValidateDSL = func() {
	Service("ServiceQueryArrayIntValidate", func() {
		Method("MethodQueryArrayIntValidate", func() {
			Payload(func() {
				Attribute("q", ArrayOf(Int), func() {
					MinLength(1)
					Elem(func() {
						Minimum(1)
					})
				})
				Required("q")
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryArrayInt32DSL = func() {
	Service("ServiceQueryArrayInt32", func() {
		Method("MethodQueryArrayInt32", func() {
			Payload(func() {
				Attribute("q", ArrayOf(Int32))
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryArrayInt32ValidateDSL = func() {
	Service("ServiceQueryArrayInt32Validate", func() {
		Method("MethodQueryArrayInt32Validate", func() {
			Payload(func() {
				Attribute("q", ArrayOf(Int32), func() {
					MinLength(1)
					Elem(func() {
						Minimum(1)
					})
				})
				Required("q")
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryArrayInt64DSL = func() {
	Service("ServiceQueryArrayInt64", func() {
		Method("MethodQueryArrayInt64", func() {
			Payload(func() {
				Attribute("q", ArrayOf(Int64))
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryArrayInt64ValidateDSL = func() {
	Service("ServiceQueryArrayInt64Validate", func() {
		Method("MethodQueryArrayInt64Validate", func() {
			Payload(func() {
				Attribute("q", ArrayOf(Int64), func() {
					MinLength(1)
					Elem(func() {
						Minimum(1)
					})
				})
				Required("q")
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryArrayUIntDSL = func() {
	Service("ServiceQueryArrayUInt", func() {
		Method("MethodQueryArrayUInt", func() {
			Payload(func() {
				Attribute("q", ArrayOf(UInt))
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryArrayUIntValidateDSL = func() {
	Service("ServiceQueryArrayUIntValidate", func() {
		Method("MethodQueryArrayUIntValidate", func() {
			Payload(func() {
				Attribute("q", ArrayOf(UInt), func() {
					MinLength(1)
					Elem(func() {
						Minimum(1)
					})
				})
				Required("q")
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryArrayUInt32DSL = func() {
	Service("ServiceQueryArrayUInt32", func() {
		Method("MethodQueryArrayUInt32", func() {
			Payload(func() {
				Attribute("q", ArrayOf(UInt32))
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryArrayUInt32ValidateDSL = func() {
	Service("ServiceQueryArrayUInt32Validate", func() {
		Method("MethodQueryArrayUInt32Validate", func() {
			Payload(func() {
				Attribute("q", ArrayOf(UInt32), func() {
					MinLength(1)
					Elem(func() {
						Minimum(1)
					})
				})
				Required("q")
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryArrayUInt64DSL = func() {
	Service("ServiceQueryArrayUInt64", func() {
		Method("MethodQueryArrayUInt64", func() {
			Payload(func() {
				Attribute("q", ArrayOf(UInt64))
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryArrayUInt64ValidateDSL = func() {
	Service("ServiceQueryArrayUInt64Validate", func() {
		Method("MethodQueryArrayUInt64Validate", func() {
			Payload(func() {
				Attribute("q", ArrayOf(UInt64), func() {
					MinLength(1)
					Elem(func() {
						Minimum(1)
					})
				})
				Required("q")
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryArrayFloat32DSL = func() {
	Service("ServiceQueryArrayFloat32", func() {
		Method("MethodQueryArrayFloat32", func() {
			Payload(func() {
				Attribute("q", ArrayOf(Float32))
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryArrayFloat32ValidateDSL = func() {
	Service("ServiceQueryArrayFloat32Validate", func() {
		Method("MethodQueryArrayFloat32Validate", func() {
			Payload(func() {
				Attribute("q", ArrayOf(Float32), func() {
					MinLength(1)
					Elem(func() {
						Minimum(1)
					})
				})
				Required("q")
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryArrayFloat64DSL = func() {
	Service("ServiceQueryArrayFloat64", func() {
		Method("MethodQueryArrayFloat64", func() {
			Payload(func() {
				Attribute("q", ArrayOf(Float64))
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryArrayFloat64ValidateDSL = func() {
	Service("ServiceQueryArrayFloat64Validate", func() {
		Method("MethodQueryArrayFloat64Validate", func() {
			Payload(func() {
				Attribute("q", ArrayOf(Float64), func() {
					MinLength(1)
					Elem(func() {
						Minimum(1)
					})
				})
				Required("q")
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryArrayStringDSL = func() {
	Service("ServiceQueryArrayString", func() {
		Method("MethodQueryArrayString", func() {
			Payload(func() {
				Attribute("q", ArrayOf(String))
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryArrayStringValidateDSL = func() {
	Service("ServiceQueryArrayStringValidate", func() {
		Method("MethodQueryArrayStringValidate", func() {
			Payload(func() {
				Attribute("q", ArrayOf(String), func() {
					MinLength(1)
					Elem(func() {
						Enum("val")
					})
				})
				Required("q")
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryArrayBytesDSL = func() {
	Service("ServiceQueryArrayBytes", func() {
		Method("MethodQueryArrayBytes", func() {
			Payload(func() {
				Attribute("q", ArrayOf(Bytes))
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryArrayBytesValidateDSL = func() {
	Service("ServiceQueryArrayBytesValidate", func() {
		Method("MethodQueryArrayBytesValidate", func() {
			Payload(func() {
				Attribute("q", ArrayOf(Bytes), func() {
					MinLength(1)
					Elem(func() {
						MinLength(2)
					})
				})
				Required("q")
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryArrayAnyDSL = func() {
	Service("ServiceQueryArrayAny", func() {
		Method("MethodQueryArrayAny", func() {
			Payload(func() {
				Attribute("q", ArrayOf(Any))
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryArrayAnyValidateDSL = func() {
	Service("ServiceQueryArrayAnyValidate", func() {
		Method("MethodQueryArrayAnyValidate", func() {
			Payload(func() {
				Attribute("q", ArrayOf(Any), func() {
					MinLength(1)
					Elem(func() {
						Enum("val", 1)
					})
				})
				Required("q")
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryMapStringStringDSL = func() {
	Service("ServiceQueryMapStringString", func() {
		Method("MethodQueryMapStringString", func() {
			Payload(func() {
				Attribute("q", MapOf(String, String))
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryMapStringStringValidateDSL = func() {
	Service("ServiceQueryMapStringStringValidate", func() {
		Method("MethodQueryMapStringStringValidate", func() {
			Payload(func() {
				Attribute("q", MapOf(String, String), func() {
					MinLength(1)
					Key(func() {
						Enum("key")
					})
					Elem(func() {
						Enum("val")
					})
				})
				Required("q")
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryMapStringBoolDSL = func() {
	Service("ServiceQueryMapStringBool", func() {
		Method("MethodQueryMapStringBool", func() {
			Payload(func() {
				Attribute("q", MapOf(String, Boolean))
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryMapStringBoolValidateDSL = func() {
	Service("ServiceQueryMapStringBoolValidate", func() {
		Method("MethodQueryMapStringBoolValidate", func() {
			Payload(func() {
				Attribute("q", MapOf(String, Boolean), func() {
					MinLength(1)
					Key(func() {
						Enum("key")
					})
					Elem(func() {
						Enum(true)
					})
				})
				Required("q")
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryMapBoolStringDSL = func() {
	Service("ServiceQueryMapBoolString", func() {
		Method("MethodQueryMapBoolString", func() {
			Payload(func() {
				Attribute("q", MapOf(Boolean, String))
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryMapBoolStringValidateDSL = func() {
	Service("ServiceQueryMapBoolStringValidate", func() {
		Method("MethodQueryMapBoolStringValidate", func() {
			Payload(func() {
				Attribute("q", MapOf(Boolean, String), func() {
					MinLength(1)
					Key(func() {
						Enum(true)
					})
					Elem(func() {
						Enum("val")
					})
				})
				Required("q")
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryMapBoolBoolDSL = func() {
	Service("ServiceQueryMapBoolBool", func() {
		Method("MethodQueryMapBoolBool", func() {
			Payload(func() {
				Attribute("q", MapOf(Boolean, Boolean))
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryMapBoolBoolValidateDSL = func() {
	Service("ServiceQueryMapBoolBoolValidate", func() {
		Method("MethodQueryMapBoolBoolValidate", func() {
			Payload(func() {
				Attribute("q", MapOf(Boolean, Boolean), func() {
					MinLength(1)
					Key(func() {
						Enum(false)
					})
					Elem(func() {
						Enum(true)
					})
				})
				Required("q")
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryMapStringArrayStringDSL = func() {
	Service("ServiceQueryMapStringArrayString", func() {
		Method("MethodQueryMapStringArrayString", func() {
			Payload(func() {
				Attribute("q", MapOf(String, ArrayOf(String)))
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryMapStringArrayStringValidateDSL = func() {
	Service("ServiceQueryMapStringArrayStringValidate", func() {
		Method("MethodQueryMapStringArrayStringValidate", func() {
			Payload(func() {
				Attribute("q", MapOf(String, ArrayOf(String)), func() {
					MinLength(1)
					Key(func() {
						Enum("key")
					})
					Elem(func() {
						MinLength(2)
					})
				})
				Required("q")
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryMapStringArrayBoolDSL = func() {
	Service("ServiceQueryMapStringArrayBool", func() {
		Method("MethodQueryMapStringArrayBool", func() {
			Payload(func() {
				Attribute("q", MapOf(String, ArrayOf(Boolean)))
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryMapStringArrayBoolValidateDSL = func() {
	Service("ServiceQueryMapStringArrayBoolValidate", func() {
		Method("MethodQueryMapStringArrayBoolValidate", func() {
			Payload(func() {
				Attribute("q", MapOf(String, ArrayOf(Boolean)), func() {
					MinLength(1)
					Key(func() {
						Enum("key")
					})
					Elem(func() {
						MinLength(2)
					})
				})
				Required("q")
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryMapBoolArrayBoolDSL = func() {
	Service("ServiceQueryMapBoolArrayBool", func() {
		Method("MethodQueryMapBoolArrayBool", func() {
			Payload(func() {
				Attribute("q", MapOf(Boolean, ArrayOf(Boolean)))
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryMapBoolArrayBoolValidateDSL = func() {
	Service("ServiceQueryMapBoolArrayBoolValidate", func() {
		Method("MethodQueryMapBoolArrayBoolValidate", func() {
			Payload(func() {
				Attribute("q", MapOf(Boolean, ArrayOf(Boolean)), func() {
					MinLength(1)
					Key(func() {
						Enum(true)
					})
					Elem(func() {
						MinLength(2)
					})
				})
				Required("q")
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryMapBoolArrayStringDSL = func() {
	Service("ServiceQueryMapBoolArrayString", func() {
		Method("MethodQueryMapBoolArrayString", func() {
			Payload(func() {
				Attribute("q", MapOf(Boolean, ArrayOf(String)))
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryMapBoolArrayStringValidateDSL = func() {
	Service("ServiceQueryMapBoolArrayStringValidate", func() {
		Method("MethodQueryMapBoolArrayStringValidate", func() {
			Payload(func() {
				Attribute("q", MapOf(Boolean, ArrayOf(String)), func() {
					MinLength(1)
					Key(func() {
						Enum(true)
					})
					Elem(func() {
						MinLength(2)
					})
				})
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryPrimitiveStringValidateDSL = func() {
	Service("ServiceQueryPrimitiveStringValidate", func() {
		Method("MethodQueryPrimitiveStringValidate", func() {
			Payload(String, func() {
				Enum("val")
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryPrimitiveBoolValidateDSL = func() {
	Service("ServiceQueryPrimitiveBoolValidate", func() {
		Method("MethodQueryPrimitiveBoolValidate", func() {
			Payload(Boolean, func() {
				Enum(true)
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryPrimitiveArrayStringValidateDSL = func() {
	Service("ServiceQueryPrimitiveArrayStringValidate", func() {
		Method("MethodQueryPrimitiveArrayStringValidate", func() {
			Payload(ArrayOf(String), func() {
				MinLength(1)
				Elem(func() {
					Enum("val")
				})
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryPrimitiveArrayBoolValidateDSL = func() {
	Service("ServiceQueryPrimitiveArrayBoolValidate", func() {
		Method("MethodQueryPrimitiveArrayBoolValidate", func() {
			Payload(ArrayOf(Boolean), func() {
				MinLength(1)
				Elem(func() {
					Enum(true)
				})
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryPrimitiveMapStringArrayStringValidateDSL = func() {
	Service("ServiceQueryPrimitiveMapStringArrayStringValidate", func() {
		Method("MethodQueryPrimitiveMapStringArrayStringValidate", func() {
			Payload(MapOf(String, ArrayOf(String)), func() {
				MinLength(1)
				Key(func() {
					Pattern("key")
				})
				Elem(func() {
					MinLength(2)
					Elem(func() {
						Pattern("val")
					})
				})
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryPrimitiveMapStringBoolValidateDSL = func() {
	Service("ServiceQueryPrimitiveMapStringBoolValidate", func() {
		Method("MethodQueryPrimitiveMapStringBoolValidate", func() {
			Payload(MapOf(String, Boolean), func() {
				MinLength(1)
				Key(func() {
					Pattern("key")
				})
				Elem(func() {
					Enum(true)
				})
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryPrimitiveMapBoolArrayBoolValidateDSL = func() {
	Service("ServiceQueryPrimitiveMapBoolArrayBoolValidate", func() {
		Method("MethodQueryPrimitiveMapBoolArrayBoolValidate", func() {
			Payload(MapOf(Boolean, ArrayOf(Boolean)), func() {
				MinLength(1)
				Key(func() {
					Enum(true)
				})
				Elem(func() {
					MinLength(2)
					Elem(func() {
						Enum(false)
					})
				})
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryMapStringMapIntStringValidateDSL = func() {
	Service("ServiceQueryMapStringMapIntStringValidate", func() {
		Method("MethodQueryMapStringMapIntStringValidate", func() {
			Payload(MapOf(String, MapOf(Int, String)), func() {
				Key(func() {
					Enum("foo")
				})
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryMapIntMapStringArrayIntValidateDSL = func() {
	Service("ServiceQueryMapIntMapStringArrayIntValidate", func() {
		Method("MethodQueryMapIntMapStringArrayIntValidate", func() {
			Payload(MapOf(Int, MapOf(String, ArrayOf(Int))), func() {
				Key(func() {
					Enum(1, 2, 3)
				})
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryStringMappedDSL = func() {
	Service("ServiceQueryStringMapped", func() {
		Method("MethodQueryStringMapped", func() {
			Payload(func() {
				Attribute("query")
			})
			HTTP(func() {
				GET("/")
				Param("query:q")
			})
		})
	})
}

var PayloadQueryStringDefaultDSL = func() {
	Service("ServiceQueryStringDefault", func() {
		Method("MethodQueryStringDefault", func() {
			Payload(func() {
				Attribute("q", func() {
					Default("def")
				})
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryStringDefaultValidateDSL = func() {
	Service("ServiceQueryStringDefaultValidate", func() {
		Method("MethodQueryStringDefaultValidate", func() {
			Payload(func() {
				Attribute("q", func() {
					Default("def")
					Enum("def")
				})
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadQueryPrimitiveStringDefaultDSL = func() {
	Service("ServiceQueryPrimitiveStringDefault", func() {
		Method("MethodQueryPrimitiveStringDefault", func() {
			Payload(String, func() {
				Default("def")
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadExtendedQueryStringDSL = func() {
	var UT = Type("UserType", func() {
		Attribute("q", String)
	})
	Service("ServiceQueryStringExtendedPayload", func() {
		Method("MethodQueryStringExtendedPayload", func() {
			Payload(func() {
				Extend(UT)
			})
			HTTP(func() {
				GET("/")
				Param("q")
			})
		})
	})
}

var PayloadExtendedValidateDSL = func() {
	var UT = Type("UserType", func() {
		Attribute("q", String)
		Attribute("h", Int)
		Attribute("body", String)
		Required("h")
	})
	Service("ServiceQueryStringExtendedValidatePayload", func() {
		Method("MethodQueryStringExtendedValidatePayload", func() {
			Payload(func() {
				Extend(UT)
				Required("q", "body")
			})
			HTTP(func() {
				GET("/")
				Param("q")
				Header("h:Location")
			})
		})
	})
}

var PayloadPathStringDSL = func() {
	Service("ServicePathString", func() {
		Method("MethodPathString", func() {
			Payload(func() {
				Attribute("p", String)
			})
			HTTP(func() {
				GET("/{p}")
			})
		})
	})
}

var PayloadPathStringValidateDSL = func() {
	Service("ServicePathStringValidate", func() {
		Method("MethodPathStringValidate", func() {
			Payload(func() {
				Attribute("p", String, func() {
					Enum("val")
				})
				Required("p")
			})
			HTTP(func() {
				GET("/{p}")
			})
		})
	})
}

var PayloadPathStringDefaultDSL = func() {
	Service("ServicePathStringDefault", func() {
		Method("MethodPathStringDefault", func() {
			Payload(func() {
				Attribute("p", String, func() {
					Default("def")
				})
			})
			HTTP(func() {
				GET("/{p}")
			})
		})
	})
}

var PayloadPathArrayStringDSL = func() {
	Service("ServicePathArrayString", func() {
		Method("MethodPathArrayString", func() {
			Payload(func() {
				Attribute("p", ArrayOf(String))
			})
			HTTP(func() {
				GET("/{p}")
			})
		})
	})
}

var PayloadPathArrayStringValidateDSL = func() {
	Service("ServicePathArrayStringValidate", func() {
		Method("MethodPathArrayStringValidate", func() {
			Payload(func() {
				Attribute("p", ArrayOf(String), func() {
					Enum([]string{"val"})
				})
			})
			HTTP(func() {
				GET("/{p}")
			})
		})
	})
}

var PayloadPathPrimitiveStringValidateDSL = func() {
	Service("ServicePathPrimitiveStringValidate", func() {
		Method("MethodPathPrimitiveStringValidate", func() {
			Payload(String, func() {
				Enum("val")
			})
			HTTP(func() {
				GET("/{p}")
			})
		})
	})
}

var PayloadPathPrimitiveBoolValidateDSL = func() {
	Service("ServicePathPrimitiveBoolValidate", func() {
		Method("MethodPathPrimitiveBoolValidate", func() {
			Payload(Boolean, func() {
				Enum(true)
			})
			HTTP(func() {
				GET("/{p}")
			})
		})
	})
}

var PayloadPathPrimitiveArrayStringValidateDSL = func() {
	Service("ServicePathPrimitiveArrayStringValidate", func() {
		Method("MethodPathPrimitiveArrayStringValidate", func() {
			Payload(ArrayOf(String), func() {
				MinLength(1)
				Elem(func() {
					Enum("val")
				})
			})
			HTTP(func() {
				GET("/{p}")
			})
		})
	})
}

var PayloadPathPrimitiveArrayBoolValidateDSL = func() {
	Service("ServicePathPrimitiveArrayBoolValidate", func() {
		Method("MethodPathPrimitiveArrayBoolValidate", func() {
			Payload(ArrayOf(Boolean), func() {
				MinLength(1)
				Elem(func() {
					Enum(true)
				})
			})
			HTTP(func() {
				GET("/{p}")
			})
		})
	})
}

var PayloadHeaderStringDSL = func() {
	Service("ServiceHeaderString", func() {
		Method("MethodHeaderString", func() {
			Payload(func() {
				Attribute("h", String)
			})
			HTTP(func() {
				GET("/")
				Header("h")
			})
		})
	})
}

var PayloadHeaderStringValidateDSL = func() {
	Service("ServiceHeaderStringValidate", func() {
		Method("MethodHeaderStringValidate", func() {
			Payload(func() {
				Attribute("h", String, func() {
					Pattern("header")
				})
			})
			HTTP(func() {
				GET("/")
				Header("h")
			})
		})
	})
}

var PayloadHeaderArrayStringDSL = func() {
	Service("ServiceHeaderArrayString", func() {
		Method("MethodHeaderArrayString", func() {
			Payload(func() {
				Attribute("h", ArrayOf(String))
			})
			HTTP(func() {
				GET("/")
				Header("h")
			})
		})
	})
}

var PayloadHeaderArrayStringValidateDSL = func() {
	Service("ServiceHeaderArrayStringValidate", func() {
		Method("MethodHeaderArrayStringValidate", func() {
			Payload(func() {
				Attribute("h", ArrayOf(String, func() {
					Enum("val")
				}))
			})
			HTTP(func() {
				GET("/")
				Header("h")
			})
		})
	})
}

var PayloadHeaderPrimitiveStringValidateDSL = func() {
	Service("ServiceHeaderPrimitiveStringValidate", func() {
		Method("MethodHeaderPrimitiveStringValidate", func() {
			Payload(String, func() {
				Enum("val")
			})
			HTTP(func() {
				GET("/")
				Header("h")
			})
		})
	})
}

var PayloadHeaderPrimitiveBoolValidateDSL = func() {
	Service("ServiceHeaderPrimitiveBoolValidate", func() {
		Method("MethodHeaderPrimitiveBoolValidate", func() {
			Payload(Boolean, func() {
				Enum(true)
			})
			HTTP(func() {
				GET("/")
				Header("h")
			})
		})
	})
}

var PayloadHeaderPrimitiveArrayStringValidateDSL = func() {
	Service("ServiceHeaderPrimitiveArrayStringValidate", func() {
		Method("MethodHeaderPrimitiveArrayStringValidate", func() {
			Payload(ArrayOf(String), func() {
				MinLength(1)
				Elem(func() {
					Pattern("val")
				})
			})
			HTTP(func() {
				GET("/")
				Header("h")
			})
		})
	})
}

var PayloadHeaderPrimitiveArrayBoolValidateDSL = func() {
	Service("ServiceHeaderPrimitiveArrayBoolValidate", func() {
		Method("MethodHeaderPrimitiveArrayBoolValidate", func() {
			Payload(ArrayOf(Boolean), func() {
				MinLength(1)
				Elem(func() {
					Enum(true)
				})
			})
			HTTP(func() {
				GET("/")
				Header("h")
			})
		})
	})
}

var PayloadHeaderStringDefaultDSL = func() {
	Service("ServiceHeaderStringDefault", func() {
		Method("MethodHeaderStringDefault", func() {
			Payload(func() {
				Attribute("h", String, func() {
					Default("def")
				})
			})
			HTTP(func() {
				GET("/")
				Header("h")
			})
		})
	})
}

var PayloadHeaderStringDefaultValidateDSL = func() {
	Service("ServiceHeaderStringDefaultValidate", func() {
		Method("MethodHeaderStringDefaultValidate", func() {
			Payload(func() {
				Attribute("h", String, func() {
					Default("def")
					Enum("def")
				})
			})
			HTTP(func() {
				GET("/")
				Header("h")
			})
		})
	})
}

var PayloadHeaderPrimitiveStringDefaultDSL = func() {
	Service("ServiceHeaderPrimitiveStringDefault", func() {
		Method("MethodHeaderPrimitiveStringDefault", func() {
			Payload(String, func() {
				Default("def")
			})
			HTTP(func() {
				GET("")
				Header("h")
			})
		})
	})
}

var PayloadBodyStringDSL = func() {
	Service("ServiceBodyString", func() {
		Method("MethodBodyString", func() {
			Payload(func() {
				Attribute("b", String)
			})
			HTTP(func() {
				POST("/")
			})
		})
	})
}

var PayloadBodyStringValidateDSL = func() {
	Service("ServiceBodyStringValidate", func() {
		Method("MethodBodyStringValidate", func() {
			Payload(func() {
				Attribute("b", String, func() {
					Pattern("pattern")
				})
				Required("b")
			})
			HTTP(func() {
				POST("/")
			})
		})
	})
}

var PayloadBodyUserDSL = func() {
	var PayloadType = Type("PayloadType", func() {
		Attribute("a", String)
	})
	Service("ServiceBodyUser", func() {
		Method("MethodBodyUser", func() {
			Payload(PayloadType)
			HTTP(func() {
				POST("/")
			})
		})
	})
}

var PayloadBodyUserRequiredDSL = func() {
	var PayloadType = Type("PayloadType", func() {
		Attribute("a", String)
		Attribute("b", String)
		Required("a")
	})
	Service("ServiceBodyUser", func() {
		Method("MethodBodyUser", func() {
			Payload(PayloadType)
			HTTP(func() {
				POST("/")
			})
		})
	})
}

var PayloadBodyNestedUserDSL = func() {
	var NestedType = Type("NestedType", func() {
		Attribute("a", String)
		Attribute("b", String)
		Required("a")
	})
	var PayloadType = Type("PayloadType", func() {
		Attribute("data", NestedType)
	})
	Service("ServiceBodyUser", func() {
		Method("MethodBodyUser", func() {
			Payload(PayloadType)
			HTTP(func() {
				POST("/")
				Body("data")
			})
		})
	})
}

var PayloadBodyUserValidateDSL = func() {
	var PayloadType = Type("PayloadType", func() {
		Attribute("a", String, func() {
			Pattern("apattern")
		})
	})
	Service("ServiceBodyUserValidate", func() {
		Method("MethodBodyUserValidate", func() {
			Payload(PayloadType)
			HTTP(func() {
				POST("/")
			})
		})
	})
}

var PayloadBodyArrayStringDSL = func() {
	Service("ServiceBodyArrayString", func() {
		Method("MethodBodyArrayString", func() {
			Payload(func() {
				Attribute("b", ArrayOf(String))
			})
			HTTP(func() {
				POST("/")
			})
		})
	})
}

var PayloadBodyArrayStringValidateDSL = func() {
	Service("ServiceBodyArrayStringValidate", func() {
		Method("MethodBodyArrayStringValidate", func() {
			Payload(func() {
				Attribute("b", ArrayOf(String), func() {
					MinLength(2)
					Elem(func() {
						MinLength(3)
					})
				})
			})
			HTTP(func() {
				POST("/")
			})
		})
	})
}

var PayloadBodyArrayUserDSL = func() {
	var PayloadType = Type("PayloadType", func() {
		Attribute("a", String, func() {
			Pattern("apattern")
		})
	})
	Service("ServiceBodyArrayUser", func() {
		Method("MethodBodyArrayUser", func() {
			Payload(func() {
				Attribute("b", ArrayOf(PayloadType))
			})
			HTTP(func() {
				POST("/")
			})
		})
	})
}

var PayloadBodyArrayUserValidateDSL = func() {
	var PayloadType = Type("PayloadType", func() {
		Attribute("a", String, func() {
			Pattern("apattern")
		})
	})
	Service("ServiceBodyArrayUserValidate", func() {
		Method("MethodBodyArrayUserValidate", func() {
			Payload(func() {
				Attribute("b", ArrayOf(PayloadType), func() {
					MinLength(2)
				})
			})
			HTTP(func() {
				POST("/")
			})
		})
	})
}

var PayloadBodyMapStringDSL = func() {
	Service("ServiceBodyMapString", func() {
		Method("MethodBodyMapString", func() {
			Payload(func() {
				Attribute("b", MapOf(String, String))
			})
			HTTP(func() {
				POST("/")
			})
		})
	})
}

var PayloadBodyMapStringValidateDSL = func() {
	Service("ServiceBodyMapStringValidate", func() {
		Method("MethodBodyMapStringValidate", func() {
			Payload(func() {
				Attribute("b", MapOf(String, String), func() {
					Elem(func() {
						MinLength(2)
					})
				})
			})
			HTTP(func() {
				POST("/")
			})
		})
	})
}

var PayloadBodyMapUserDSL = func() {
	var PayloadType = Type("PayloadType", func() {
		Attribute("a", String, func() {
			Pattern("apattern")
		})
	})
	Service("ServiceBodyMapUser", func() {
		Method("MethodBodyMapUser", func() {
			Payload(func() {
				Attribute("b", MapOf(String, PayloadType))
			})
			HTTP(func() {
				POST("/")
			})
		})
	})
}

var PayloadBodyMapUserValidateDSL = func() {
	var PayloadType = Type("PayloadType", func() {
		Attribute("a", String, func() {
			Pattern("apattern")
		})
	})
	Service("ServiceBodyMapUserValidate", func() {
		Method("MethodBodyMapUserValidate", func() {
			Payload(func() {
				Attribute("b", MapOf(String, PayloadType), func() {
					Key(func() {
						MinLength(2)
					})
				})
			})
			HTTP(func() {
				POST("/")
			})
		})
	})
}

var PayloadBodyPrimitiveStringValidateDSL = func() {
	Service("ServiceBodyPrimitiveStringValidate", func() {
		Method("MethodBodyPrimitiveStringValidate", func() {
			Payload(String, func() {
				Enum("val")
			})
			HTTP(func() {
				POST("/")
			})
		})
	})
}

var PayloadBodyPrimitiveBoolValidateDSL = func() {
	Service("ServiceBodyPrimitiveBoolValidate", func() {
		Method("MethodBodyPrimitiveBoolValidate", func() {
			Payload(Boolean, func() {
				Enum(true)
			})
			HTTP(func() {
				POST("/")
			})
		})
	})
}

var PayloadBodyPrimitiveArrayStringValidateDSL = func() {
	Service("ServiceBodyPrimitiveArrayStringValidate", func() {
		Method("MethodBodyPrimitiveArrayStringValidate", func() {
			Payload(ArrayOf(String), func() {
				MinLength(1)
				Elem(func() {
					Enum("val")
				})
			})
			HTTP(func() {
				POST("/")
			})
		})
	})
}

var PayloadBodyPrimitiveArrayBoolValidateDSL = func() {
	Service("ServiceBodyPrimitiveArrayBoolValidate", func() {
		Method("MethodBodyPrimitiveArrayBoolValidate", func() {
			Payload(ArrayOf(Boolean), func() {
				MinLength(1)
				Elem(func() {
					Enum(true)
				})
			})
			HTTP(func() {
				POST("/")
			})
		})
	})
}

var PayloadBodyPrimitiveArrayUserValidateDSL = func() {
	var PayloadType = Type("PayloadType", func() {
		Attribute("a", String, func() {
			Pattern("pattern")
		})
		Required("a")
	})
	Service("ServiceBodyPrimitiveArrayUserValidate", func() {
		Method("MethodBodyPrimitiveArrayUserValidate", func() {
			Payload(ArrayOf(PayloadType), func() {
				MinLength(1)
			})
			HTTP(func() {
				POST("/")
			})
		})
	})
}

var PayloadBodyPrimitiveFieldEmptyDSL = func() {
	var PayloadType = Type("PayloadType", func() {
		Attribute("a", ArrayOf(String))
	})
	Service("ServiceBodyPrimitiveArrayUser", func() {
		Method("MethodBodyPrimitiveArrayUser", func() {
			Payload(PayloadType)
			HTTP(func() {
				POST("/")
				Param("a")
				Body(Empty)
			})
		})
	})
}

var PayloadBodyPrimitiveFieldArrayUserDSL = func() {
	var PayloadType = Type("PayloadType", func() {
		Attribute("a", ArrayOf(String))
	})
	Service("ServiceBodyPrimitiveArrayUser", func() {
		Method("MethodBodyPrimitiveArrayUser", func() {
			Payload(PayloadType)
			HTTP(func() {
				POST("/")
				Body("a")
			})
		})
	})
}

var PayloadExtendBodyPrimitiveFieldArrayUserDSL = func() {
	var Ext = Type("Ext", func() {
		Attribute("b", String)
	})
	var PayloadType = Type("PayloadType", func() {
		Extend(Ext)
		Attribute("a", ArrayOf(String))
	})
	Service("ServiceBodyPrimitiveArrayUser", func() {
		Method("MethodBodyPrimitiveArrayUser", func() {
			Payload(PayloadType)
			HTTP(func() {
				POST("/")
				Body("a")
			})
		})
	})
}

var PayloadExtendBodyPrimitiveFieldStringDSL = func() {
	var Ext = Type("Ext", func() {
		Attribute("b", String)
	})
	var PayloadType = Type("PayloadType", func() {
		Extend(Ext)
		Attribute("a", ArrayOf(String))
	})
	Service("ServiceBodyPrimitiveArrayUser", func() {
		Method("MethodBodyPrimitiveArrayUser", func() {
			Payload(PayloadType)
			HTTP(func() {
				POST("/")
				Body("b")
			})
		})
	})
}

var PayloadBodyPrimitiveFieldArrayUserValidateDSL = func() {
	var PayloadType = Type("PayloadType", func() {
		Attribute("a", ArrayOf(String), func() {
			MinLength(1)
			Elem(func() {
				Pattern("pattern")
			})
		})
		Required("a")
	})
	Service("ServiceBodyPrimitiveArrayUserValidate", func() {
		Method("MethodBodyPrimitiveArrayUserValidate", func() {
			Payload(PayloadType)
			HTTP(func() {
				POST("/")
				Body("a")
			})
		})
	})
}

var PayloadBodyQueryObjectDSL = func() {
	Service("ServiceBodyQueryObject", func() {
		Method("MethodBodyQueryObject", func() {
			Payload(func() {
				Attribute("a", String)
				Attribute("b", String)
			})
			HTTP(func() {
				POST("/")
				Param("b")
			})
		})
	})
}

var PayloadBodyQueryObjectValidateDSL = func() {
	Service("ServiceBodyQueryObjectValidate", func() {
		Method("MethodBodyQueryObjectValidate", func() {
			Payload(func() {
				Attribute("a", String, func() {
					Pattern("patterna")
				})
				Attribute("b", String, func() {
					Pattern("patternb")
				})
				Required("a", "b")
			})
			HTTP(func() {
				POST("/")
				Param("b")
			})
		})
	})
}

var PayloadBodyQueryUserDSL = func() {
	var PayloadType = Type("PayloadType", func() {
		Attribute("a", String)
		Attribute("b", String)
	})
	Service("ServiceBodyQueryUser", func() {
		Method("MethodBodyQueryUser", func() {
			Payload(PayloadType)
			HTTP(func() {
				POST("/")
				Param("b")
			})
		})
	})
}

var PayloadBodyQueryUserValidateDSL = func() {
	var PayloadType = Type("PayloadType", func() {
		Attribute("a", String, func() {
			Pattern("patterna")
		})
		Attribute("b", String, func() {
			Pattern("patternb")
		})
		Required("a", "b")
	})
	Service("ServiceBodyQueryUserValidate", func() {
		Method("MethodBodyQueryUserValidate", func() {
			Payload(PayloadType)
			HTTP(func() {
				POST("/")
				Param("b")
			})
		})
	})
}

var PayloadBodyPathObjectDSL = func() {
	Service("ServiceBodyPathObject", func() {
		Method("MethodBodyPathObject", func() {
			Payload(func() {
				Attribute("a", String)
				Attribute("b", String)
			})
			HTTP(func() {
				POST("/{b}")
			})
		})
	})
}

var PayloadBodyPathObjectValidateDSL = func() {
	Service("ServiceBodyPathObjectValidate", func() {
		Method("MethodBodyPathObjectValidate", func() {
			Payload(func() {
				Attribute("a", String, func() {
					Pattern("patterna")
				})
				Attribute("b", String, func() {
					Pattern("patternb")
				})
				Required("a", "b")
			})
			HTTP(func() {
				POST("/{b}")
			})
		})
	})
}

var PayloadBodyPathUserDSL = func() {
	var PayloadType = Type("PayloadType", func() {
		Attribute("a", String)
		Attribute("b", String)
	})
	Service("ServiceBodyPathUser", func() {
		Method("MethodBodyPathUser", func() {
			Payload(PayloadType)
			HTTP(func() {
				POST("/{b}")
			})
		})
	})
}

var PayloadBodyPathUserValidateDSL = func() {
	var PayloadType = Type("PayloadType", func() {
		Attribute("a", String, func() {
			Pattern("patterna")
		})
		Attribute("b", String, func() {
			Pattern("patternb")
		})
		Required("a", "b")
	})
	Service("ServiceBodyPathUserValidate", func() {
		Method("MethodUserBodyPathValidate", func() {
			Payload(PayloadType)
			HTTP(func() {
				POST("/{b}")
			})
		})
	})
}

var PayloadBodyQueryPathObjectDSL = func() {
	Service("ServiceBodyQueryPathObject", func() {
		Method("MethodBodyQueryPathObject", func() {
			Payload(func() {
				Attribute("a", String)
				Attribute("b", String)
				Attribute("c", String)
			})
			HTTP(func() {
				POST("/{c}")
				Param("b")
			})
		})
	})
}

var PayloadBodyQueryPathObjectValidateDSL = func() {
	Service("ServiceBodyQueryPathObjectValidate", func() {
		Method("MethodBodyQueryPathObjectValidate", func() {
			Payload(func() {
				Attribute("a", String, func() {
					Pattern("patterna")
				})
				Attribute("b", String, func() {
					Pattern("patternb")
				})
				Attribute("c", String, func() {
					Pattern("patternc")
				})
				Required("a", "b", "c")
			})
			HTTP(func() {
				POST("/{c}")
				Param("b")
			})
		})
	})
}

var PayloadBodyQueryPathUserDSL = func() {
	var PayloadType = Type("PayloadType", func() {
		Attribute("a", String)
		Attribute("b", String)
		Attribute("c", String)
	})
	Service("ServiceBodyQueryPathUser", func() {
		Method("MethodBodyQueryPathUser", func() {
			Payload(PayloadType)
			HTTP(func() {
				POST("/{c}")
				Param("b")
			})
		})
	})
}

var PayloadBodyQueryPathUserValidateDSL = func() {
	var PayloadType = Type("PayloadType", func() {
		Attribute("a", String, func() {
			Pattern("patterna")
		})
		Attribute("b", String, func() {
			Pattern("patternb")
		})
		Attribute("c", String, func() {
			Pattern("patternc")
		})
		Required("a", "b", "c")
	})
	Service("ServiceBodyQueryPathUserValidate", func() {
		Method("MethodBodyQueryPathUserValidate", func() {
			Payload(PayloadType)
			HTTP(func() {
				POST("/{c}")
				Param("b")
			})
		})
	})
}

var PayloadBodyUserInnerDSL = func() {
	var InnerType = Type("InnerType", func() {
		Attribute("a", String, func() {
			Pattern("patterna")
		})
		Attribute("b", String, func() {
			Pattern("patternb")
		})
		Required("a")
	})
	var PayloadType = Type("PayloadType", func() {
		Attribute("inner", InnerType)
	})
	Service("ServiceBodyUserInner", func() {
		Method("MethodBodyUserInner", func() {
			Payload(PayloadType)
			HTTP(func() {
				POST("/")
			})
		})
	})
}

var PayloadBodyUserInnerDefaultDSL = func() {
	var InnerType = Type("InnerType", func() {
		Attribute("a", String, func() {
			Default("defaulta")
			Pattern("patterna")
		})
		Attribute("b", String, func() {
			Default("defaultb")
			Pattern("patternb")
		})
		Required("a")
	})
	var PayloadType = Type("PayloadType", func() {
		Attribute("inner", InnerType)
	})
	Service("ServiceBodyUserInnerDefault", func() {
		Method("MethodBodyUserInnerDefault", func() {
			Payload(PayloadType)
			HTTP(func() {
				POST("/")
			})
		})
	})
}

var PayloadBodyInlineArrayUserDSL = func() {
	var ElemType = Type("ElemType", func() {
		Attribute("a", String, func() {
			Pattern("patterna")
		})
		Attribute("b", String, func() {
			Pattern("patternb")
		})
		Required("a")
	})
	Service("ServiceBodyInlineArrayUser", func() {
		Method("MethodBodyInlineArrayUser", func() {
			Payload(ArrayOf(ElemType))
			HTTP(func() {
				POST("/")
			})
		})
	})
}

var PayloadBodyInlineMapUserDSL = func() {
	var KeyType = Type("KeyType", func() {
		Attribute("a", String, func() {
			Pattern("patterna")
		})
		Attribute("b", String, func() {
			Pattern("patternb")
		})
		Required("a")
	})
	var ElemType = Type("ElemType", func() {
		Attribute("a", String, func() {
			Pattern("patterna")
		})
		Attribute("b", String, func() {
			Pattern("patternb")
		})
		Required("a")
	})
	Service("ServiceBodyInlineMapUser", func() {
		Method("MethodBodyInlineMapUser", func() {
			Payload(MapOf(KeyType, ElemType))
			HTTP(func() {
				POST("/")
			})
		})
	})
}

var PayloadBodyInlineRecursiveUserDSL = func() {
	var PayloadType = Type("PayloadType", func() {
		Attribute("a", String, func() {
			Pattern("patterna")
		})
		Attribute("b", String, func() {
			Pattern("patternb")
		})
		Attribute("c", "PayloadType")
		Required("a", "c")
	})

	Service("ServiceBodyInlineRecursiveUser", func() {
		Method("MethodBodyInlineRecursiveUser", func() {
			Payload(PayloadType)
			HTTP(func() {
				POST("/{a}")
				Param("b")
			})
		})
	})
}

var PayloadMapQueryPrimitivePrimitiveDSL = func() {
	Service("ServiceMapQueryPrimitivePrimitive", func() {
		Method("MapQueryPrimitivePrimitive", func() {
			Payload(MapOf(String, String))
			HTTP(func() {
				POST("/")
				MapParams()
			})
		})
	})
}

var PayloadMapQueryPrimitiveArrayDSL = func() {
	Service("ServiceMapQueryPrimitiveArray", func() {
		Method("MapQueryPrimitiveArray", func() {
			Payload(MapOf(String, ArrayOf(UInt)))
			HTTP(func() {
				POST("/")
				MapParams()
			})
		})
	})
}

var PayloadMapQueryObjectDSL = func() {
	var PayloadType = Type("PayloadType", func() {
		Attribute("a", String, func() {
			Pattern("patterna")
		})
		Attribute("b", String, func() {
			Pattern("patternb")
		})
		Attribute("c", MapOf(Int, ArrayOf(String)))
		Required("a", "c")
	})

	Service("ServiceMapQueryObject", func() {
		Method("MethodMapQueryObject", func() {
			Payload(PayloadType)
			HTTP(func() {
				POST("/{a}")
				MapParams("c")
			})
		})
	})
}

var PayloadMultipartPrimitiveDSL = func() {
	Service("ServiceMultipartPrimitive", func() {
		Method("MethodMultipartPrimitive", func() {
			Payload(String)
			HTTP(func() {
				POST("/")
				MultipartRequest()
			})
		})
	})
}

var PayloadMultipartUserTypeDSL = func() {
	Service("ServiceMultipartUserType", func() {
		Method("MethodMultipartUserType", func() {
			Payload(func() {
				Attribute("b", String, func() {
					Pattern("patternb")
				})
				Attribute("c", MapOf(Int, ArrayOf(String)))
				Required("b", "c")
			})
			HTTP(func() {
				POST("/")
				MultipartRequest()
			})
		})
	})
}

var PayloadMultipartArrayTypeDSL = func() {
	var PayloadType = Type("PayloadType", func() {
		Attribute("a", String, func() {
			Pattern("patterna")
		})
		Attribute("b", String, func() {
			Pattern("patternb")
		})
		Attribute("c", MapOf(Int, ArrayOf(String)))
		Required("a", "c")
	})
	Service("ServiceMultipartArrayType", func() {
		Method("MethodMultipartArrayType", func() {
			Payload(ArrayOf(PayloadType))
			HTTP(func() {
				POST("/")
				MultipartRequest()
			})
		})
	})
}

var PayloadMultipartMapTypeDSL = func() {
	Service("ServiceMultipartMapType", func() {
		Method("MethodMultipartMapType", func() {
			Payload(MapOf(String, Int))
			HTTP(func() {
				POST("/")
				MultipartRequest()
			})
		})
	})
}

var PayloadMultipartWithParamDSL = func() {
	var PayloadType = Type("PayloadType", func() {
		Attribute("a", String, func() {
			Pattern("patterna")
		})
		Attribute("b", String, func() {
			Pattern("patternb")
		})
		Attribute("c", MapOf(Int, ArrayOf(String)))
		Required("a", "c")
	})
	Service("ServiceMultipartWithParam", func() {
		Method("MethodMultipartWithParam", func() {
			Payload(PayloadType)
			Result(String)
			HTTP(func() {
				POST("/")
				Param("c")
				MultipartRequest()
			})
		})
	})
}

var PayloadMultipartWithParamsAndHeadersDSL = func() {
	var PayloadType = Type("PayloadType", func() {
		Attribute("a", String, func() {
			Pattern("patterna")
		})
		Attribute("b", String, func() {
			Pattern("patternb")
		})
		Attribute("c", MapOf(Int, ArrayOf(String)))
		Required("a", "c")
	})
	Service("ServiceMultipartWithParamsAndHeaders", func() {
		Method("MethodMultipartWithParamsAndHeaders", func() {
			Payload(PayloadType)
			Result(String)
			HTTP(func() {
				POST("/{a}")
				Param("c")
				Header("b:Authorization", String)
				MultipartRequest()
			})
		})
	})
}

var MultipleMethodsDSL = func() {
	var APayload = Type("APayload", func() {
		Attribute("a", String, func() {
			Pattern("patterna")
		})
	})
	var PayloadType = Type("PayloadType", func() {
		Attribute("a", String, func() {
			Pattern("patterna")
		})
		Attribute("b", String, func() {
			Pattern("patternb")
		})
		Attribute("c", APayload)
		Required("a", "c")
	})
	Service("ServiceMultipleMethods", func() {
		Method("MethodA", func() {
			Payload(APayload)
			HTTP(func() {
				POST("/")
				Body(APayload)
			})
		})
		Method("MethodB", func() {
			Payload(PayloadType)
			HTTP(func() {
				PUT("/")
			})
		})
	})
}

var MixedPayloadInBodyDSL = func() {
	var BPayload = Type("BPayload", func() {
		Attribute("int", Int)
		Attribute("bytes", Bytes)
		Required("int")
	})
	var APayload = Type("APayload", func() {
		Attribute("any", Any)
		Attribute("array", ArrayOf(Float32))
		Attribute("map", MapOf(UInt, Any))
		Attribute("object", BPayload)
		Attribute("dup_obj", BPayload)
		Required("array", "object")
	})
	Service("ServiceMixedPayloadInBody", func() {
		Method("MethodA", func() {
			Payload(APayload)
			HTTP(func() {
				POST("/")
			})
		})
	})
}

var WithParamsAndHeadersBlockDSL = func() {
	Service("ServiceWithParamsAndHeadersBlock", func() {
		Method("MethodA", func() {
			Payload(func() {
				Attribute("required", String)
				Attribute("optional", Int)
				Attribute("optional_but_required_param", Float32)
				Attribute("optional_but_required_header", Float32)
				Attribute("path", UInt)
				Attribute("body", String)
				Required("required")
			})
			HTTP(func() {
				POST("/{path}")
				Params(func() {
					Param("optional", Int)
					Param("optional_but_required_param", Float32)
					Required("optional_but_required_param")
				})
				Headers(func() {
					Header("required", String)
					Header("optional_but_required_header", Float32)
					Required("optional_but_required_header")
				})
			})
		})
	})
}

var QueryIntAliasDSL = func() {
	var IntAlias = Type("IntAlias", Int)
	var Int32Alias = Type("Int32Alias", Int32)
	var Int64Alias = Type("Int64Alias", Int64)
	Service("ServiceQueryIntAlias", func() {
		Method("MethodA", func() {
			Payload(func() {
				Attribute("int", IntAlias)
				Attribute("int32", Int32Alias)
				Attribute("int64", Int64Alias)
			})
			HTTP(func() {
				POST("/")
				Params(func() {
					Param("int")
					Param("int32")
					Param("int64")
				})
			})
		})
	})
}

var HeaderIntAliasDSL = func() {
	var IntAlias = Type("IntAlias", Int)
	var Int32Alias = Type("Int32Alias", Int32)
	var Int64Alias = Type("Int64Alias", Int64)
	Service("ServiceHeaderIntAlias", func() {
		Method("MethodA", func() {
			Payload(func() {
				Attribute("int", IntAlias)
				Attribute("int32", Int32Alias)
				Attribute("int64", Int64Alias)
			})
			HTTP(func() {
				POST("/")
				Headers(func() {
					Header("int")
					Header("int32")
					Header("int64")
				})
			})
		})
	})
}

var PathIntAliasDSL = func() {
	var IntAlias = Type("IntAlias", Int)
	var Int32Alias = Type("Int32Alias", Int32)
	var Int64Alias = Type("Int64Alias", Int64)
	Service("ServicePathIntAlias", func() {
		Method("MethodA", func() {
			Payload(func() {
				Attribute("int", IntAlias)
				Attribute("int32", Int32Alias)
				Attribute("int64", Int64Alias)
			})
			HTTP(func() {
				POST("/{int}/{int32}/{int64}")
			})
		})
	})
}

var QueryIntAliasValidateDSL = func() {
	var IntAlias = Type("IntAlias", Int, func() {
		Minimum(10)
	})
	var Int32Alias = Type("Int32Alias", Int32, func() {
		Maximum(100)
	})
	var Int64Alias = Type("Int64Alias", Int64, func() {
		Minimum(0)
	})
	Service("ServiceQueryIntAliasValidate", func() {
		Method("MethodA", func() {
			Payload(func() {
				Attribute("int", IntAlias)
				Attribute("int32", Int32Alias)
				Attribute("int64", Int64Alias)
			})
			HTTP(func() {
				POST("/")
				Params(func() {
					Param("int")
					Param("int32")
					Param("int64")
				})
			})
		})
	})
}

var QueryArrayAliasDSL = func() {
	var ArrayAlias = Type("ArrayAlias", ArrayOf(UInt))
	Service("ServiceQueryArrayAlias", func() {
		Method("MethodA", func() {
			Payload(func() {
				Attribute("array", ArrayAlias)
			})
			HTTP(func() {
				POST("/")
				Params(func() {
					Param("array")
				})
			})
		})
	})
}

var QueryArrayAliasValidateDSL = func() {
	var ArrayAlias = Type("ArrayAlias", ArrayOf(UInt), func() {
		MinLength(3)
		Elem(func() {
			Minimum(10)
		})
	})
	Service("ServiceQueryArrayAliasValidate", func() {
		Method("MethodA", func() {
			Payload(func() {
				Attribute("array", ArrayAlias)
			})
			HTTP(func() {
				POST("/")
				Params(func() {
					Param("array")
				})
			})
		})
	})
}

var QueryMapAliasDSL = func() {
	var MapAlias = Type("MapAlias", MapOf(Float32, Boolean))
	Service("ServiceQueryMapAlias", func() {
		Method("MethodA", func() {
			Payload(func() {
				Attribute("map", MapAlias)
			})
			HTTP(func() {
				POST("/")
				Params(func() {
					Param("map")
				})
			})
		})
	})
}

var QueryMapAliasValidateDSL = func() {
	var MapAlias = Type("MapAlias", MapOf(Float32, Boolean), func() {
		MinLength(5)
	})
	Service("ServiceQueryMapAliasValidate", func() {
		Method("MethodA", func() {
			Payload(func() {
				Attribute("map", MapAlias)
			})
			HTTP(func() {
				POST("/")
				Params(func() {
					Param("map")
				})
			})
		})
	})
}

var QueryArrayNestedAliasValidateDSL = func() {
	var Float64Alias = Type("Float64Alias", Float64, func() {
		Minimum(10)
	})
	var ArrayAlias = Type("ArrayAlias", ArrayOf(Float64Alias))
	Service("ServiceQueryArrayAliasValidate", func() {
		Method("MethodA", func() {
			Payload(func() {
				Attribute("array", ArrayAlias)
			})
			HTTP(func() {
				POST("/")
				Params(func() {
					Param("array")
				})
			})
		})
	})
}
