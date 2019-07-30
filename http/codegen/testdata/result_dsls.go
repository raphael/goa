package testdata

import (
	. "goa.design/goa/v3/dsl"
)

// The DSL function names follow the following pattern:
//
// Result(Header|Body)(Type)(Required|Default)?DSL
//
// Where Type is the type of the result or result.

var ResultHeaderBoolDSL = func() {
	Service("ServiceHeaderBool", func() {
		Method("MethodHeaderBool", func() {
			Result(func() {
				Attribute("h", Boolean)
			})
			HTTP(func() {
				GET("/")
				Response(StatusOK, func() {
					Header("h")
				})
			})
		})
	})
}

var ResultHeaderIntDSL = func() {
	Service("ServiceHeaderInt", func() {
		Method("MethodHeaderInt", func() {
			Result(func() {
				Attribute("h", Int)
			})
			HTTP(func() {
				GET("/")
				Response(StatusOK, func() {
					Header("h")
				})
			})
		})
	})
}

var ResultHeaderInt32DSL = func() {
	Service("ServiceHeaderInt32", func() {
		Method("MethodHeaderInt32", func() {
			Result(func() {
				Attribute("h", Int32)
			})
			HTTP(func() {
				GET("/")
				Response(StatusOK, func() {
					Header("h")
				})
			})
		})
	})
}

var ResultHeaderInt64DSL = func() {
	Service("ServiceHeaderInt64", func() {
		Method("MethodHeaderInt64", func() {
			Result(func() {
				Attribute("h", Int64)
			})
			HTTP(func() {
				GET("/")
				Response(StatusOK, func() {
					Header("h")
				})
			})
		})
	})
}

var ResultHeaderUIntDSL = func() {
	Service("ServiceHeaderUInt", func() {
		Method("MethodHeaderUInt", func() {
			Result(func() {
				Attribute("h", UInt)
			})
			HTTP(func() {
				GET("/")
				Response(StatusOK, func() {
					Header("h")
				})
			})
		})
	})
}

var ResultHeaderUInt32DSL = func() {
	Service("ServiceHeaderUInt32", func() {
		Method("MethodHeaderUInt32", func() {
			Result(func() {
				Attribute("h", UInt32)
			})
			HTTP(func() {
				GET("/")
				Response(StatusOK, func() {
					Header("h")
				})
			})
		})
	})
}

var ResultHeaderUInt64DSL = func() {
	Service("ServiceHeaderUInt64", func() {
		Method("MethodHeaderUInt64", func() {
			Result(func() {
				Attribute("h", UInt64)
			})
			HTTP(func() {
				GET("/")
				Response(StatusOK, func() {
					Header("h")
				})
			})
		})
	})
}

var ResultHeaderFloat32DSL = func() {
	Service("ServiceHeaderFloat32", func() {
		Method("MethodHeaderFloat32", func() {
			Result(func() {
				Attribute("h", Float32)
			})
			HTTP(func() {
				GET("/")
				Response(StatusOK, func() {
					Header("h")
				})
			})
		})
	})
}

var ResultHeaderFloat64DSL = func() {
	Service("ServiceHeaderFloat64", func() {
		Method("MethodHeaderFloat64", func() {
			Result(func() {
				Attribute("h", Float64)
			})
			HTTP(func() {
				GET("/")
				Response(StatusOK, func() {
					Header("h")
				})
			})
		})
	})
}

var ResultHeaderStringDSL = func() {
	Service("ServiceHeaderString", func() {
		Method("MethodHeaderString", func() {
			Result(func() {
				Attribute("h", String)
			})
			HTTP(func() {
				GET("/")
				Response(StatusOK, func() {
					Header("h")
				})
			})
		})
	})
}

var ResultHeaderBytesDSL = func() {
	Service("ServiceHeaderBytes", func() {
		Method("MethodHeaderBytes", func() {
			Result(func() {
				Attribute("h", Bytes)
			})
			HTTP(func() {
				GET("/")
				Response(StatusOK, func() {
					Header("h")
				})
			})
		})
	})
}

var ResultHeaderAnyDSL = func() {
	Service("ServiceHeaderAny", func() {
		Method("MethodHeaderAny", func() {
			Result(func() {
				Attribute("h", Any)
			})
			HTTP(func() {
				GET("/")
				Response(StatusOK, func() {
					Header("h")
				})
			})
		})
	})
}

var ResultHeaderArrayBoolDSL = func() {
	Service("ServiceHeaderArrayBool", func() {
		Method("MethodHeaderArrayBool", func() {
			Result(func() {
				Attribute("h", ArrayOf(Boolean))
			})
			HTTP(func() {
				GET("/")
				Response(StatusOK, func() {
					Header("h")
				})
			})
		})
	})
}

var ResultHeaderArrayIntDSL = func() {
	Service("ServiceHeaderArrayInt", func() {
		Method("MethodHeaderArrayInt", func() {
			Result(func() {
				Attribute("h", ArrayOf(Int))
			})
			HTTP(func() {
				GET("/")
				Response(StatusOK, func() {
					Header("h")
				})
			})
		})
	})
}

var ResultHeaderArrayInt32DSL = func() {
	Service("ServiceHeaderArrayInt32", func() {
		Method("MethodHeaderArrayInt32", func() {
			Result(func() {
				Attribute("h", ArrayOf(Int32))
			})
			HTTP(func() {
				GET("/")
				Response(StatusOK, func() {
					Header("h")
				})
			})
		})
	})
}

var ResultHeaderArrayInt64DSL = func() {
	Service("ServiceHeaderArrayInt64", func() {
		Method("MethodHeaderArrayInt64", func() {
			Result(func() {
				Attribute("h", ArrayOf(Int64))
			})
			HTTP(func() {
				GET("/")
				Response(StatusOK, func() {
					Header("h")
				})
			})
		})
	})
}

var ResultHeaderArrayUIntDSL = func() {
	Service("ServiceHeaderArrayUInt", func() {
		Method("MethodHeaderArrayUInt", func() {
			Result(func() {
				Attribute("h", ArrayOf(UInt))
			})
			HTTP(func() {
				GET("/")
				Response(StatusOK, func() {
					Header("h")
				})
			})
		})
	})
}

var ResultHeaderArrayUInt32DSL = func() {
	Service("ServiceHeaderArrayUInt32", func() {
		Method("MethodHeaderArrayUInt32", func() {
			Result(func() {
				Attribute("h", ArrayOf(UInt32))
			})
			HTTP(func() {
				GET("/")
				Response(StatusOK, func() {
					Header("h")
				})
			})
		})
	})
}

var ResultHeaderArrayUInt64DSL = func() {
	Service("ServiceHeaderArrayUInt64", func() {
		Method("MethodHeaderArrayUInt64", func() {
			Result(func() {
				Attribute("h", ArrayOf(UInt64))
			})
			HTTP(func() {
				GET("/")
				Response(StatusOK, func() {
					Header("h")
				})
			})
		})
	})
}

var ResultHeaderArrayFloat32DSL = func() {
	Service("ServiceHeaderArrayFloat32", func() {
		Method("MethodHeaderArrayFloat32", func() {
			Result(func() {
				Attribute("h", ArrayOf(Float32))
			})
			HTTP(func() {
				GET("/")
				Response(StatusOK, func() {
					Header("h")
				})
			})
		})
	})
}

var ResultHeaderArrayFloat64DSL = func() {
	Service("ServiceHeaderArrayFloat64", func() {
		Method("MethodHeaderArrayFloat64", func() {
			Result(func() {
				Attribute("h", ArrayOf(Float64))
			})
			HTTP(func() {
				GET("/")
				Response(StatusOK, func() {
					Header("h")
				})
			})
		})
	})
}

var ResultHeaderArrayStringDSL = func() {
	Service("ServiceHeaderArrayString", func() {
		Method("MethodHeaderArrayString", func() {
			Result(func() {
				Attribute("h", ArrayOf(String))
			})
			HTTP(func() {
				GET("/")
				Response(StatusOK, func() {
					Header("h")
				})
			})
		})
	})
}

var ResultHeaderArrayBytesDSL = func() {
	Service("ServiceHeaderArrayBytes", func() {
		Method("MethodHeaderArrayBytes", func() {
			Result(func() {
				Attribute("h", ArrayOf(Bytes))
			})
			HTTP(func() {
				GET("/")
				Response(StatusOK, func() {
					Header("h")
				})
			})
		})
	})
}

var ResultHeaderArrayAnyDSL = func() {
	Service("ServiceHeaderArrayAny", func() {
		Method("MethodHeaderArrayAny", func() {
			Result(func() {
				Attribute("h", ArrayOf(Any))
			})
			HTTP(func() {
				GET("/")
				Response(StatusOK, func() {
					Header("h")
				})
			})
		})
	})
}

var ResultHeaderBoolDefaultDSL = func() {
	Service("ServiceHeaderBoolDefault", func() {
		Method("MethodHeaderBoolDefault", func() {
			Result(func() {
				Attribute("h", Boolean, func() {
					Default(true)
				})
			})
			HTTP(func() {
				GET("/")
				Response(StatusOK, func() {
					Header("h")
				})
			})
		})
	})
}

var ResultHeaderBoolRequiredDefaultDSL = func() {
	Service("ServiceHeaderBoolRequiredDefault", func() {
		Method("MethodHeaderBoolRequiredDefault", func() {
			Result(func() {
				Attribute("h", Boolean, func() {
					Default(true)
				})
				Required("h")
			})
			HTTP(func() {
				GET("/")
				Response(StatusOK, func() {
					Header("h")
				})
			})
		})
	})
}

var ResultHeaderStringDefaultDSL = func() {
	Service("ServiceHeaderStringDefault", func() {
		Method("MethodHeaderStringDefault", func() {
			Result(func() {
				Attribute("h", func() {
					Default("def")
				})
			})
			HTTP(func() {
				GET("/")
				Response(StatusOK, func() {
					Header("h")
				})
			})
		})
	})
}

var ResultHeaderStringRequiredDefaultDSL = func() {
	Service("ServiceHeaderStringRequiredDefault", func() {
		Method("MethodHeaderStringRequiredDefault", func() {
			Result(func() {
				Attribute("h", func() {
					Default("def")
				})
				Required("h")
			})
			HTTP(func() {
				GET("/")
				Response(StatusOK, func() {
					Header("h")
				})
			})
		})
	})
}

var ResultHeaderArrayBoolDefaultDSL = func() {
	Service("ServiceHeaderArrayBoolDefault", func() {
		Method("MethodHeaderArrayBoolDefault", func() {
			Result(func() {
				Attribute("h", ArrayOf(Boolean), func() {
					Default([]bool{true, false})
				})
			})
			HTTP(func() {
				GET("/")
				Response(StatusOK, func() {
					Header("h")
				})
			})
		})
	})
}

var ResultHeaderArrayBoolRequiredDefaultDSL = func() {
	Service("ServiceHeaderArrayBoolRequiredDefault", func() {
		Method("MethodHeaderArrayBoolRequiredDefault", func() {
			Result(func() {
				Attribute("h", ArrayOf(Boolean), func() {
					Default([]bool{true, false})
				})
				Required("h")
			})
			HTTP(func() {
				GET("/")
				Response(StatusOK, func() {
					Header("h")
				})
			})
		})
	})
}

var ResultHeaderArrayStringDefaultDSL = func() {
	Service("ServiceHeaderArrayStringDefault", func() {
		Method("MethodHeaderArrayStringDefault", func() {
			Result(func() {
				Attribute("h", ArrayOf(String), func() {
					Default([]string{"foo", "bar"})
				})
			})
			HTTP(func() {
				GET("/")
				Response(StatusOK, func() {
					Header("h")
				})
			})
		})
	})
}

var ResultHeaderArrayStringRequiredDefaultDSL = func() {
	Service("ServiceHeaderArrayStringRequiredDefault", func() {
		Method("MethodHeaderArrayStringRequiredDefault", func() {
			Result(func() {
				Attribute("h", ArrayOf(String), func() {
					Default([]string{"foo", "bar"})
				})
				Required("h")
			})
			HTTP(func() {
				GET("/")
				Response(StatusOK, func() {
					Header("h")
				})
			})
		})
	})
}

var ResultBodyStringDSL = func() {
	Service("ServiceBodyString", func() {
		Method("MethodBodyString", func() {
			Result(func() {
				Attribute("b", String)
			})
			HTTP(func() {
				POST("/")
			})
		})
	})
}

var ResultBodyObjectDSL = func() {
	Service("ServiceBodyObject", func() {
		Method("MethodBodyObject", func() {
			Result(func() {
				Attribute("b", String)
			})
			HTTP(func() {
				POST("/")
			})
		})
	})
}

var ResultBodyObjectHeaderDSL = func() {
	Service("ServiceBodyObjectHeader", func() {
		Method("MethodBodyObjectHeader", func() {
			Result(func() {
				Attribute("a", String)
				Attribute("b", String)
			})
			HTTP(func() {
				POST("/")
				Response(StatusOK, func() {
					Header("b:Authorization")
				})
			})
		})
	})
}

var ResultBodyUserDSL = func() {
	var ResultType = Type("ResultType", func() {
		Attribute("a", String)
	})
	Service("ServiceBodyUser", func() {
		Method("MethodBodyUser", func() {
			Result(ResultType)
			HTTP(func() {
				POST("/")
			})
		})
	})
}

var ResultTypeValidateDSL = func() {
	var ResultType = Type("ResultType", func() {
		Attribute("a", String, func() {
			MinLength(5)
		})
	})
	Service("ServiceResultTypeValidate", func() {
		Method("MethodResultTypeValidate", func() {
			Result(ResultType)
			HTTP(func() {
				POST("/")
			})
		})
	})
}

var ResultBodyMultipleViewsDSL = func() {
	var ResultType = ResultType("ResultTypeMultipleViews", func() {
		Attribute("a", String)
		Attribute("b", String)
		Attribute("c", String)
		View("default", func() {
			Attribute("a")
			Attribute("b")
			Attribute("c")
		})
		View("tiny", func() {
			Attribute("c")
		})
	})
	Service("ServiceBodyMultipleView", func() {
		Method("MethodBodyMultipleView", func() {
			Result(ResultType)
			HTTP(func() {
				POST("/")
				Response(StatusOK, func() {
					Header("c:Location")
				})
			})
		})
	})
}

var ResultBodyCollectionDSL = func() {
	var RT = ResultType("ResultTypeCollection", func() {
		Attributes(func() {
			Attribute("a", String)
			Attribute("b", String)
			Attribute("c", String)
		})
		View("default", func() {
			Attribute("a")
			Attribute("b")
			Attribute("c")
		})
		View("tiny", func() {
			Attribute("c")
		})
	})
	Service("ServiceBodyCollection", func() {
		Method("MethodBodyCollection", func() {
			Result(CollectionOf(RT))
			HTTP(func() {
				POST("/")
				Response(StatusOK)
			})
		})
	})
}

var ResultBodyCollectionExplicitViewDSL = func() {
	var RT = ResultType("ResultTypeCollection", func() {
		Attributes(func() {
			Attribute("a", String)
			Attribute("b", String)
			Attribute("c", String)
		})
		View("default", func() {
			Attribute("a")
			Attribute("b")
			Attribute("c")
		})
		View("tiny", func() {
			Attribute("c")
		})
	})
	Service("ServiceBodyCollectionExplicitView", func() {
		Method("MethodBodyCollectionExplicitView", func() {
			Result(CollectionOf(RT), func() {
				View("tiny")
			})
			HTTP(func() {
				POST("/")
				Response(StatusOK)
			})
		})
	})
}

var ResultWithResultCollectionDSL = func() {
	var RT = ResultType("RT", func() {
		Attributes(func() {
			Attribute("x", String, func() {
				MinLength(5)
			})
		})
	})
	var ResultType = ResultType("ResultType", func() {
		Attributes(func() {
			Attribute("x", CollectionOf(RT))
		})
	})
	Service("ServiceResultWithResultCollection", func() {
		Method("MethodResultWithResultCollection", func() {
			Result(func() {
				Attribute("a", ResultType)
			})
			HTTP(func() {
				POST("/")
				Response(StatusOK)
			})
		})
	})
}

var EmptyBodyResultMultipleViewsDSL = func() {
	var ResultType = ResultType("ResultTypeMultipleViews", func() {
		Attribute("a", String)
		Attribute("b", String)
		Attribute("c", String)
		View("default", func() {
			Attribute("a")
			Attribute("c")
		})
		View("tiny", func() {
			Attribute("c")
		})
	})
	Service("ServiceEmptyBodyResultMultipleView", func() {
		Method("MethodEmptyBodyResultMultipleView", func() {
			Result(ResultType)
			HTTP(func() {
				POST("/")
				Response(StatusOK, func() {
					Header("c:Location")
					Body(Empty)
				})
			})
		})
	})
}

var ExplicitBodyPrimitiveResultMultipleViewsDSL = func() {
	var ResultType = ResultType("ResultTypeMultipleViews", func() {
		Attribute("a", String, func() {
			MinLength(5)
		})
		Attribute("b", String)
		Attribute("c", String)
		View("default", func() {
			Attribute("a")
			Attribute("b")
			Attribute("c")
		})
		View("tiny", func() {
			Attribute("a")
			Attribute("c")
		})
	})
	Service("ServiceExplicitBodyPrimitiveResultMultipleView", func() {
		Method("MethodExplicitBodyPrimitiveResultMultipleView", func() {
			Result(ResultType)
			HTTP(func() {
				POST("/")
				Response(StatusOK, func() {
					Header("c:Location")
					Body("a")
				})
			})
		})
	})
}

var ExplicitBodyUserResultMultipleViewsDSL = func() {
	var UserType = Type("UserType", func() {
		Attribute("x", String)
		Attribute("y", Int)
	})
	var ResultType = ResultType("ResultTypeMultipleViews", func() {
		Attribute("a", UserType)
		Attribute("b", String)
		Attribute("c", String)
		View("default", func() {
			Attribute("a")
			Attribute("b")
			Attribute("c")
		})
		View("tiny", func() {
			Attribute("a")
			Attribute("c")
		})
	})
	Service("ServiceExplicitBodyUserResultMultipleView", func() {
		Method("MethodExplicitBodyUserResultMultipleView", func() {
			Result(ResultType)
			HTTP(func() {
				POST("/")
				Response(StatusOK, func() {
					Header("c:Location")
					Body("a")
				})
			})
		})
	})
}

var ExplicitBodyResultCollectionDSL = func() {
	var ResultType = ResultType("ResultType", func() {
		Attributes(func() {
			Attribute("x", String, func() {
				MinLength(5)
			})
		})
	})
	Service("ServiceExplicitBodyResultCollection", func() {
		Method("MethodExplicitBodyResultCollection", func() {
			Result(func() {
				Attribute("a", CollectionOf(ResultType))
			})
			HTTP(func() {
				POST("/")
				Response(StatusOK, func() {
					Body("a")
				})
			})
		})
	})
}

var ExplicitContentTypeResultDSL = func() {
	var ResultType = ResultType("ResultType", func() {
		ContentType("application/custom+json")
		Attribute("a", String)
		Attribute("b", String)
	})
	Service("ServiceExplicitContentTypeResult", func() {
		Method("MethodExplicitContentTypeResult", func() {
			Result(ResultType)
			HTTP(func() {
				POST("/")
			})
		})
	})
}

var ExplicitContentTypeResponseDSL = func() {
	var ResultType = ResultType("ResultType", func() {
		Attribute("a", String)
		Attribute("b", String)
	})
	Service("ServiceExplicitContentTypeResponse", func() {
		Method("MethodExplicitContentTypeResponse", func() {
			Result(ResultType)
			HTTP(func() {
				POST("/")
				Response(StatusOK, func() {
					ContentType("application/custom+json")
				})
			})
		})
	})
}

var ResultBodyArrayStringDSL = func() {
	Service("ServiceBodyArrayString", func() {
		Method("MethodBodyArrayString", func() {
			Result(func() {
				Attribute("b", ArrayOf(String))
			})
			HTTP(func() {
				POST("/")
			})
		})
	})
}

var ResultBodyArrayUserDSL = func() {
	var ResultType = Type("ResultType", func() {
		Attribute("a", String, func() {
			Pattern("apattern")
		})
	})
	Service("ServiceBodyArrayUser", func() {
		Method("MethodBodyArrayUser", func() {
			Result(func() {
				Attribute("b", ArrayOf(ResultType))
			})
			HTTP(func() {
				POST("/")
			})
		})
	})
}

var ResultBodyPrimitiveStringDSL = func() {
	Service("ServiceBodyPrimitiveString", func() {
		Method("MethodBodyPrimitiveString", func() {
			Result(String, func() {
				Enum("val")
			})
			HTTP(func() {
				POST("/")
			})
		})
	})
}

var ResultBodyPrimitiveBoolDSL = func() {
	Service("ServiceBodyPrimitiveBool", func() {
		Method("MethodBodyPrimitiveBool", func() {
			Result(Boolean, func() {
				Enum(true)
			})
			HTTP(func() {
				POST("/")
			})
		})
	})
}

var ResultBodyPrimitiveArrayStringDSL = func() {
	Service("ServiceBodyPrimitiveArrayString", func() {
		Method("MethodBodyPrimitiveArrayString", func() {
			Result(ArrayOf(String), func() {
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

var ResultBodyPrimitiveArrayBoolDSL = func() {
	Service("ServiceBodyPrimitiveArrayBool", func() {
		Method("MethodBodyPrimitiveArrayBool", func() {
			Result(ArrayOf(Boolean), func() {
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

var ResultBodyPrimitiveArrayUserDSL = func() {
	var ResultType = Type("ResultType", func() {
		Attribute("a", String, func() {
			Pattern("apattern")
		})
	})
	Service("ServiceBodyPrimitiveArrayUser", func() {
		Method("MethodBodyPrimitiveArrayUser", func() {
			Result(ArrayOf(ResultType))
			HTTP(func() {
				POST("/")
			})
		})
	})
}

var ResultBodyHeaderObjectDSL = func() {
	Service("ServiceBodyHeaderObject", func() {
		Method("MethodBodyHeaderObject", func() {
			Result(func() {
				Attribute("a", String)
				Attribute("b", String)
			})
			HTTP(func() {
				POST("/")
				Response(StatusOK, func() {
					Header("b")
				})
			})
		})
	})
}

var ResultBodyHeaderUserDSL = func() {
	var ResultType = Type("ResultType", func() {
		Attribute("a", String)
		Attribute("b", String)
	})
	Service("ServiceBodyHeaderUser", func() {
		Method("MethodBodyHeaderUser", func() {
			Result(ResultType)
			HTTP(func() {
				POST("/")
				Response(StatusOK, func() {
					Header("b")
				})
			})
		})
	})
}

var ResultTagStringDSL = func() {
	Service("ServiceTagString", func() {
		Method("MethodTagString", func() {
			Result(func() {
				Attribute("h", String)
			})
			HTTP(func() {
				GET("/")
				Response(StatusAccepted, func() {
					Header("h")
					Tag("h", "value")
				})
				Response(StatusOK)
			})
		})
	})
}

var ResultTagStringRequiredDSL = func() {
	Service("ServiceTagStringRequired", func() {
		Method("MethodTagStringRequired", func() {
			Result(func() {
				Attribute("h", String)
				Required("h")
			})
			HTTP(func() {
				GET("/")
				Response(StatusAccepted, func() {
					Header("h")
					Tag("h", "value")
				})
				Response(StatusOK)
			})
		})
	})
}

var ResultMultipleViewsTagDSL = func() {
	var ResultType = ResultType("ResultTypeMultipleViews", func() {
		Attribute("a", String)
		Attribute("b", String)
		Attribute("c", String)
		View("default", func() {
			Attribute("a")
			Attribute("c")
		})
		View("tiny", func() {
			Attribute("c")
		})
	})
	Service("ServiceTagMultipleViews", func() {
		Method("MethodTagMultipleViews", func() {
			Result(ResultType)
			HTTP(func() {
				GET("/")
				Response(StatusAccepted, func() {
					Header("c")
					Tag("b", "value")
				})
				Response(StatusOK)
			})
		})
	})
}

var EmptyServerResponseDSL = func() {
	Service("ServiceEmptyServerResponse", func() {
		Method("MethodEmptyServerResponse", func() {
			Result(func() {
				Attribute("h", String)
				Required("h")
			})
			HTTP(func() {
				GET("/")
				Response(StatusOK, func() {
					Body(Empty)
				})
			})
		})
	})
}

var EmptyServerResponseWithTagsDSL = func() {
	Service("ServiceEmptyServerResponseWithTags", func() {
		Method("MethodEmptyServerResponseWithTags", func() {
			Result(func() {
				Attribute("h", String)
				Required("h")
			})
			HTTP(func() {
				GET("/")
				Response(StatusNoContent, func() {
					Body(Empty)
				})
				Response(StatusNotModified, func() {
					Tag("h", "true")
					Body(Empty)
				})
			})
		})
	})
}

var ResultHeaderStringArrayDSL = func() {
	Service("ServiceHeaderStringArrayResponse", func() {
		Method("MethodA", func() {
			Result(func() {
				Attribute("array", ArrayOf(String))
			})
			HTTP(func() {
				GET("/")
				Response(StatusOK, func() {
					Header("array")
				})
			})
		})
	})
}

var ResultHeaderStringArrayValidateDSL = func() {
	Service("ServiceHeaderStringArrayValidateResponse", func() {
		Method("MethodA", func() {
			Result(func() {
				Attribute("array", ArrayOf(String), func() {
					MinLength(5)
				})
			})
			HTTP(func() {
				GET("/")
				Response(StatusOK, func() {
					Header("array")
				})
			})
		})
	})
}

var ResultHeaderArrayDSL = func() {
	Service("ServiceHeaderArrayResponse", func() {
		Method("MethodA", func() {
			Result(func() {
				Attribute("array", ArrayOf(UInt))
			})
			HTTP(func() {
				GET("/")
				Response(StatusOK, func() {
					Header("array")
				})
			})
		})
	})
}

var ResultHeaderArrayValidateDSL = func() {
	Service("ServiceHeaderArrayValidateResponse", func() {
		Method("MethodA", func() {
			Result(func() {
				Attribute("array", ArrayOf(Int), func() {
					Elem(func() {
						Minimum(5)
					})
				})
			})
			HTTP(func() {
				GET("/")
				Response(StatusOK, func() {
					Header("array")
				})
			})
		})
	})
}

var WithHeadersBlockDSL = func() {
	Service("ServiceWithHeadersBlock", func() {
		Method("MethodA", func() {
			Result(func() {
				Attribute("required", Int)
				Attribute("optional", Float32)
				Attribute("optional_but_required", UInt)
				Required("required")
			})
			HTTP(func() {
				GET("/")
				Response(StatusOK, func() {
					Headers(func() {
						Header("required:X-Request-ID")
						Header("optional:Authorization")
						Header("optional_but_required:Location")
						Required("optional_but_required")
					})
				})
			})
		})
	})
}

var WithHeadersBlockViewedResultDSL = func() {
	var AResult = ResultType("application/vnd.goa.aresult", func() {
		TypeName("AResult")
		Attributes(func() {
			Attribute("required", Int)
			Attribute("optional", Float32)
			Attribute("optional_but_required", UInt)
			Required("required")
		})
		View("tiny", func() {
			Attribute("required")
			Attribute("optional")
			Attribute("optional_but_required")
		})
	})
	Service("ServiceWithHeadersBlockViewedResult", func() {
		Method("MethodA", func() {
			Result(AResult)
			HTTP(func() {
				GET("/")
				Response(StatusOK, func() {
					Headers(func() {
						Header("required:X-Request-ID")
						Header("optional:Authorization")
						Header("optional_but_required:Location")
						Required("optional_but_required")
					})
				})
			})
		})
	})
}

var ValidateErrorResponseTypeDSL = func() {
	var AResult = ResultType("application/vnd.goa.aresult", func() {
		TypeName("AResult")
		Attributes(func() {
			Attribute("required", Int)
			Required("required")
		})
	})
	var AError = Type("AError", func() {
		Attribute("error", String)
		Attribute("num_occur", Int, func() {
			Minimum(1)
		})
		Required("error")
	})
	Service("ValidateErrorResponseType", func() {
		Method("MethodA", func() {
			Result(AResult)
			Error("some_error", AError)
			HTTP(func() {
				GET("/")
				Response(StatusOK, func() {
					Headers(func() {
						Header("required:X-Request-ID")
					})
				})
				Response("some_error", StatusBadRequest, func() {
					Header("error:X-Application-Error")
					Header("num_occur:X-Occur")
				})
			})
		})
	})
}
