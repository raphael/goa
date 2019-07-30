package testdata

const (
	IntegerRequiredValidationCode = `func Validate() (err error) {
	if target.RequiredInteger < 1 {
		err = goa.MergeErrors(err, goa.InvalidRangeError("target.required_integer", target.RequiredInteger, 1, true))
	}
	if target.DefaultInteger != nil {
		if !(*target.DefaultInteger == 1 || *target.DefaultInteger == 5 || *target.DefaultInteger == 10 || *target.DefaultInteger == 100) {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("target.default_integer", *target.DefaultInteger, []interface{}{1, 5, 10, 100}))
		}
	}
	if target.Integer != nil {
		if *target.Integer > 100 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("target.integer", *target.Integer, 100, false))
		}
	}
}
`

	IntegerPointerValidationCode = `func Validate() (err error) {
	if target.RequiredInteger == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("required_integer", "target"))
	}
	if target.RequiredInteger != nil {
		if *target.RequiredInteger < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("target.required_integer", *target.RequiredInteger, 1, true))
		}
	}
	if target.DefaultInteger != nil {
		if !(*target.DefaultInteger == 1 || *target.DefaultInteger == 5 || *target.DefaultInteger == 10 || *target.DefaultInteger == 100) {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("target.default_integer", *target.DefaultInteger, []interface{}{1, 5, 10, 100}))
		}
	}
	if target.Integer != nil {
		if *target.Integer > 100 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("target.integer", *target.Integer, 100, false))
		}
	}
}
`

	IntegerUseDefaultValidationCode = `func Validate() (err error) {
	if target.RequiredInteger < 1 {
		err = goa.MergeErrors(err, goa.InvalidRangeError("target.required_integer", target.RequiredInteger, 1, true))
	}
	if !(target.DefaultInteger == 1 || target.DefaultInteger == 5 || target.DefaultInteger == 10 || target.DefaultInteger == 100) {
		err = goa.MergeErrors(err, goa.InvalidEnumValueError("target.default_integer", target.DefaultInteger, []interface{}{1, 5, 10, 100}))
	}
	if target.Integer != nil {
		if *target.Integer > 100 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("target.integer", *target.Integer, 100, false))
		}
	}
}
`

	FloatRequiredValidationCode = `func Validate() (err error) {
	if target.RequiredFloat < 1 {
		err = goa.MergeErrors(err, goa.InvalidRangeError("target.required_float", target.RequiredFloat, 1, true))
	}
	if target.DefaultInteger != nil {
		if !(*target.DefaultInteger == 1.2 || *target.DefaultInteger == 5 || *target.DefaultInteger == 10 || *target.DefaultInteger == 100.8) {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("target.default_integer", *target.DefaultInteger, []interface{}{1.2, 5, 10, 100.8}))
		}
	}
	if target.Float64 != nil {
		if *target.Float64 > 100.1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("target.float64", *target.Float64, 100.1, false))
		}
	}
}
`

	FloatPointerValidationCode = `func Validate() (err error) {
	if target.RequiredFloat == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("required_float", "target"))
	}
	if target.RequiredFloat != nil {
		if *target.RequiredFloat < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("target.required_float", *target.RequiredFloat, 1, true))
		}
	}
	if target.DefaultInteger != nil {
		if !(*target.DefaultInteger == 1.2 || *target.DefaultInteger == 5 || *target.DefaultInteger == 10 || *target.DefaultInteger == 100.8) {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("target.default_integer", *target.DefaultInteger, []interface{}{1.2, 5, 10, 100.8}))
		}
	}
	if target.Float64 != nil {
		if *target.Float64 > 100.1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("target.float64", *target.Float64, 100.1, false))
		}
	}
}
`

	FloatUseDefaultValidationCode = `func Validate() (err error) {
	if target.RequiredFloat < 1 {
		err = goa.MergeErrors(err, goa.InvalidRangeError("target.required_float", target.RequiredFloat, 1, true))
	}
	if !(target.DefaultInteger == 1.2 || target.DefaultInteger == 5 || target.DefaultInteger == 10 || target.DefaultInteger == 100.8) {
		err = goa.MergeErrors(err, goa.InvalidEnumValueError("target.default_integer", target.DefaultInteger, []interface{}{1.2, 5, 10, 100.8}))
	}
	if target.Float64 != nil {
		if *target.Float64 > 100.1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("target.float64", *target.Float64, 100.1, false))
		}
	}
}
`

	StringRequiredValidationCode = `func Validate() (err error) {
	err = goa.MergeErrors(err, goa.ValidatePattern("target.required_string", target.RequiredString, "^[A-z].*[a-z]$"))
	if utf8.RuneCountInString(target.RequiredString) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("target.required_string", target.RequiredString, utf8.RuneCountInString(target.RequiredString), 1, true))
	}
	if utf8.RuneCountInString(target.RequiredString) > 10 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("target.required_string", target.RequiredString, utf8.RuneCountInString(target.RequiredString), 10, false))
	}
	if target.DefaultString != nil {
		if !(*target.DefaultString == "foo" || *target.DefaultString == "bar") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("target.default_string", *target.DefaultString, []interface{}{"foo", "bar"}))
		}
	}
	if target.String != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("target.string", *target.String, goa.FormatDateTime))
	}
}
`

	StringPointerValidationCode = `func Validate() (err error) {
	if target.RequiredString == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("required_string", "target"))
	}
	if target.RequiredString != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("target.required_string", *target.RequiredString, "^[A-z].*[a-z]$"))
	}
	if target.RequiredString != nil {
		if utf8.RuneCountInString(*target.RequiredString) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("target.required_string", *target.RequiredString, utf8.RuneCountInString(*target.RequiredString), 1, true))
		}
	}
	if target.RequiredString != nil {
		if utf8.RuneCountInString(*target.RequiredString) > 10 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("target.required_string", *target.RequiredString, utf8.RuneCountInString(*target.RequiredString), 10, false))
		}
	}
	if target.DefaultString != nil {
		if !(*target.DefaultString == "foo" || *target.DefaultString == "bar") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("target.default_string", *target.DefaultString, []interface{}{"foo", "bar"}))
		}
	}
	if target.String != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("target.string", *target.String, goa.FormatDateTime))
	}
}
`

	StringUseDefaultValidationCode = `func Validate() (err error) {
	err = goa.MergeErrors(err, goa.ValidatePattern("target.required_string", target.RequiredString, "^[A-z].*[a-z]$"))
	if utf8.RuneCountInString(target.RequiredString) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("target.required_string", target.RequiredString, utf8.RuneCountInString(target.RequiredString), 1, true))
	}
	if utf8.RuneCountInString(target.RequiredString) > 10 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("target.required_string", target.RequiredString, utf8.RuneCountInString(target.RequiredString), 10, false))
	}
	if !(target.DefaultString == "foo" || target.DefaultString == "bar") {
		err = goa.MergeErrors(err, goa.InvalidEnumValueError("target.default_string", target.DefaultString, []interface{}{"foo", "bar"}))
	}
	if target.String != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("target.string", *target.String, goa.FormatDateTime))
	}
}
`

	UserTypeRequiredValidationCode = `func Validate() (err error) {
	if target.RequiredInteger == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("required_integer", "target"))
	}
	if target.RequiredInteger != nil {
		if err2 := ValidateInteger(target.RequiredInteger); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if target.DefaultString != nil {
		if err2 := ValidateString(target.DefaultString); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if target.Float != nil {
		if err2 := ValidateFloat(target.Float); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
}
`

	UserTypePointerValidationCode = `func Validate() (err error) {
	if target.RequiredInteger == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("required_integer", "target"))
	}
	if target.RequiredInteger != nil {
		if err2 := ValidateInteger(target.RequiredInteger); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if target.DefaultString != nil {
		if err2 := ValidateString(target.DefaultString); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if target.Float != nil {
		if err2 := ValidateFloat(target.Float); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
}
`
	UserTypeUseDefaultValidationCode = `func Validate() (err error) {
	if target.RequiredInteger == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("required_integer", "target"))
	}
	if target.RequiredInteger != nil {
		if err2 := ValidateInteger(target.RequiredInteger); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if target.DefaultString != nil {
		if err2 := ValidateString(target.DefaultString); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if target.Float != nil {
		if err2 := ValidateFloat(target.Float); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
}
`

	UserTypeArrayValidationCode = `func Validate() (err error) {
	for _, e := range target.Array {
		if e != nil {
			if err2 := ValidateFloat(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
}
`

	ArrayRequiredValidationCode = `func Validate() (err error) {
	if target.RequiredArray == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("required_array", "target"))
	}
	if len(target.RequiredArray) < 5 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("target.required_array", target.RequiredArray, len(target.RequiredArray), 5, true))
	}
	if len(target.DefaultArray) > 3 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("target.default_array", target.DefaultArray, len(target.DefaultArray), 3, false))
	}
	for _, e := range target.Array {
		if !(e == 0 || e == 1 || e == 1 || e == 2 || e == 3 || e == 5) {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("target.array[*]", e, []interface{}{0, 1, 1, 2, 3, 5}))
		}
	}
}
`

	ArrayPointerValidationCode = `func Validate() (err error) {
	if target.RequiredArray == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("required_array", "target"))
	}
	if len(target.RequiredArray) < 5 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("target.required_array", target.RequiredArray, len(target.RequiredArray), 5, true))
	}
	if len(target.DefaultArray) > 3 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("target.default_array", target.DefaultArray, len(target.DefaultArray), 3, false))
	}
	for _, e := range target.Array {
		if !(e == 0 || e == 1 || e == 1 || e == 2 || e == 3 || e == 5) {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("target.array[*]", e, []interface{}{0, 1, 1, 2, 3, 5}))
		}
	}
}
`

	ArrayUseDefaultValidationCode = `func Validate() (err error) {
	if target.RequiredArray == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("required_array", "target"))
	}
	if len(target.RequiredArray) < 5 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("target.required_array", target.RequiredArray, len(target.RequiredArray), 5, true))
	}
	if len(target.DefaultArray) > 3 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("target.default_array", target.DefaultArray, len(target.DefaultArray), 3, false))
	}
	for _, e := range target.Array {
		if !(e == 0 || e == 1 || e == 1 || e == 2 || e == 3 || e == 5) {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("target.array[*]", e, []interface{}{0, 1, 1, 2, 3, 5}))
		}
	}
}
`

	MapRequiredValidationCode = `func Validate() (err error) {
	if target.RequiredMap == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("required_map", "target"))
	}
	if len(target.RequiredMap) < 5 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("target.required_map", target.RequiredMap, len(target.RequiredMap), 5, true))
	}
	if len(target.DefaultMap) > 3 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("target.default_map", target.DefaultMap, len(target.DefaultMap), 3, false))
	}
	for k, v := range target.Map {
		err = goa.MergeErrors(err, goa.ValidatePattern("target.map.key", k, "^[A-Z]"))
		if v > 5 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("target.map[key]", v, 5, false))
		}
	}
}
`

	MapPointerValidationCode = `func Validate() (err error) {
	if target.RequiredMap == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("required_map", "target"))
	}
	if len(target.RequiredMap) < 5 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("target.required_map", target.RequiredMap, len(target.RequiredMap), 5, true))
	}
	if len(target.DefaultMap) > 3 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("target.default_map", target.DefaultMap, len(target.DefaultMap), 3, false))
	}
	for k, v := range target.Map {
		err = goa.MergeErrors(err, goa.ValidatePattern("target.map.key", k, "^[A-Z]"))
		if v > 5 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("target.map[key]", v, 5, false))
		}
	}
}
`

	MapUseDefaultValidationCode = `func Validate() (err error) {
	if target.RequiredMap == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("required_map", "target"))
	}
	if len(target.RequiredMap) < 5 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("target.required_map", target.RequiredMap, len(target.RequiredMap), 5, true))
	}
	if len(target.DefaultMap) > 3 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("target.default_map", target.DefaultMap, len(target.DefaultMap), 3, false))
	}
	for k, v := range target.Map {
		err = goa.MergeErrors(err, goa.ValidatePattern("target.map.key", k, "^[A-Z]"))
		if v > 5 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("target.map[key]", v, 5, false))
		}
	}
}
`

	ResultTypePointerValidationCode = `func Validate() (err error) {
	if target.Required != nil {
		if *target.Required < 10 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("target.required", *target.Required, 10, true))
		}
	}
}
`

	ResultCollectionPointerValidationCode = `func Validate() (err error) {
	for _, e := range target {
		if e != nil {
			if err2 := ValidateResult(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
}
`

	TypeWithCollectionPointerValidationCode = `func Validate() (err error) {
	if err2 := ValidateResultCollection(target.Collection); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
}
`
)
