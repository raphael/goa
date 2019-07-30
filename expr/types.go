package expr

import (
	"fmt"
	"reflect"

	"goa.design/goa/v3/eval"
)

type (
	// A Kind defines the conceptual type that a DataType represents.
	Kind uint

	// DataType is the common interface to all types.
	DataType interface {
		// Kind of data type, one of the Kind enum.
		Kind() Kind
		// Name returns the type name.
		Name() string
		// IsCompatible checks whether val has a Go type that is compatible with the data
		// type.
		IsCompatible(interface{}) bool
		// Example generates a pseudo-random value using the given random generator.
		Example(*Random) interface{}
		// Hash returns a unique hash value for the instance of the type.
		Hash() string
	}

	// Primitive is the type for null, boolean, integer, number, string, and time.
	Primitive Kind

	// Array is the type used to describe field arrays or repeated fields.
	Array struct {
		ElemType *AttributeExpr
	}

	// Map is the type used to describe maps of fields.
	Map struct {
		KeyType  *AttributeExpr
		ElemType *AttributeExpr
	}

	// NamedAttributeExpr describes object attributes together with their
	// names.
	NamedAttributeExpr struct {
		// Name of attribute
		Name string
		// Attribute
		Attribute *AttributeExpr
	}

	// Object is the type used to describe composite data structures.
	// Note: not a map because order matters.
	Object []*NamedAttributeExpr

	// UserType is the interface implemented by all user type
	// implementations. Plugins may leverage this interface to introduce
	// their own types.
	UserType interface {
		DataType
		eval.Expression
		// Finalizes the underlying type.
		eval.Finalizer
		// Provides the underlying type and validations.
		CompositeExpr
		// ID returns the identifier for the user type.
		ID() string
		// Rename changes the type name to the given value.
		Rename(string)
		// SetAttribute updates the underlying attribute.
		SetAttribute(*AttributeExpr)
		// Dup makes a shallow copy of the type and assigns its
		// attribute with att.
		Dup(att *AttributeExpr) UserType
		// Validate checks that the user type expression is consistent.
		Validate(ctx string, parent eval.Expression) *eval.ValidationErrors
	}

	// ArrayVal is the type used to set the default value for arrays.
	ArrayVal []interface{}

	// MapVal is the type used to set the default value for maps.
	MapVal map[interface{}]interface{}
)

const (
	// BooleanKind represents a boolean.
	BooleanKind Kind = iota + 1
	// IntKind represents a signed integer.
	IntKind
	// Int32Kind represents a signed 32-bit integer.
	Int32Kind
	// Int64Kind represents a signed 64-bit integer.
	Int64Kind
	// UIntKind represents an unsigned integer.
	UIntKind
	// UInt32Kind represents an unsigned 32-bit integer.
	UInt32Kind
	// UInt64Kind represents an unsigned 64-bit integer.
	UInt64Kind
	// Float32Kind represents a 32-bit floating number.
	Float32Kind
	// Float64Kind represents a 64-bit floating number.
	Float64Kind
	// StringKind represents a JSON string.
	StringKind
	// BytesKind represent a series of bytes (binary data).
	BytesKind
	// ArrayKind represents a JSON array.
	ArrayKind
	// ObjectKind represents a JSON object.
	ObjectKind
	// MapKind represents a JSON object where the keys are not known in
	// advance.
	MapKind
	// UserTypeKind represents a user type.
	UserTypeKind
	// ResultTypeKind represents a result type.
	ResultTypeKind
	// AnyKind represents an unknown type.
	AnyKind
)

const (
	// Boolean is the type for a JSON boolean.
	Boolean = Primitive(BooleanKind)

	// Int is the type for a signed integer.
	Int = Primitive(IntKind)

	// Int32 is the type for a signed 32-bit integer.
	Int32 = Primitive(Int32Kind)

	// Int64 is the type for a signed 64-bit integer.
	Int64 = Primitive(Int64Kind)

	// UInt is the type for an unsigned integer.
	UInt = Primitive(UIntKind)

	// UInt32 is the type for an unsigned 32-bit integer.
	UInt32 = Primitive(UInt32Kind)

	// UInt64 is the type for an unsigned 64-bit integer.
	UInt64 = Primitive(UInt64Kind)

	// Float32 is the type for a 32-bit floating number.
	Float32 = Primitive(Float32Kind)

	// Float64 is the type for a 64-bit floating number.
	Float64 = Primitive(Float64Kind)

	// String is the type for a JSON string.
	String = Primitive(StringKind)

	// Bytes is the type for binary data.
	Bytes = Primitive(BytesKind)

	// Any is the type for an arbitrary JSON value (interface{} in Go).
	Any = Primitive(AnyKind)
)

// Built-in composite types

// Empty represents empty values.
var Empty = &UserTypeExpr{
	TypeName: "Empty",
	AttributeExpr: &AttributeExpr{
		Description: "Empty represents empty values",
		Type:        &Object{},
	},
}

// Convenience methods

// AsObject returns the type underlying object if any, nil otherwise.
func AsObject(dt DataType) *Object {
	switch t := dt.(type) {
	case *UserTypeExpr:
		return AsObject(t.Type)
	case *ResultTypeExpr:
		return AsObject(t.Type)
	case *Object:
		return t
	default:
		return nil
	}
}

// AsArray returns the type underlying array if any, nil otherwise.
func AsArray(dt DataType) *Array {
	switch t := dt.(type) {
	case *UserTypeExpr:
		return AsArray(t.Type)
	case *ResultTypeExpr:
		return AsArray(t.Type)
	case *Array:
		return t
	default:
		return nil
	}
}

// AsMap returns the type underlying map if any, nil otherwise.
func AsMap(dt DataType) *Map {
	switch t := dt.(type) {
	case *UserTypeExpr:
		return AsMap(t.Type)
	case *ResultTypeExpr:
		return AsMap(t.Type)
	case *Map:
		return t
	default:
		return nil
	}
}

// IsObject returns true if the data type is an object.
func IsObject(dt DataType) bool { return AsObject(dt) != nil }

// IsArray returns true if the data type is an array.
func IsArray(dt DataType) bool { return AsArray(dt) != nil }

// IsMap returns true if the data type is a map.
func IsMap(dt DataType) bool { return AsMap(dt) != nil }

// IsPrimitive returns true if the data type is a primitive type.
func IsPrimitive(dt DataType) bool {
	switch t := dt.(type) {
	case Primitive:
		return true
	case *UserTypeExpr:
		return IsPrimitive(t.Type)
	case *ResultTypeExpr:
		return IsPrimitive(t.Type)
	default:
		return false
	}
}

// Equal compares the types recursively and returns true if they are equal. Two
// types are equal if:
//
//    - both types have the same kind
//    - array types have elements whose types are equal
//    - map types have keys and elements whose types are equal
//    - objects have the same attribute names and the attribute types are equal
//
// Note: calling Equal is not equivalent to evaluation dt.Hash() == dt2.Hash()
// as the former may return true for two user types with different names and
// thus with different hash values.
func Equal(dt, dt2 DataType) bool {
	bs := *equal(dt, dt2)
	for _, b := range bs {
		if !*b {
			return false
		}
	}
	return true
}

// Support recursive types by doing lazy evaluation.
func equal(dt, dt2 DataType, seen ...map[string]*[]*bool) *[]*bool {
	f := false
	fs := []*bool{&f}
	if dt.Kind() != dt2.Kind() {
		return &fs
	}
	var s map[string]*[]*bool
	if len(seen) > 0 {
		s = seen[0]
	} else {
		s = make(map[string]*[]*bool)
	}
	switch actual := dt.(type) {
	case *Array:
		return equal(actual.ElemType.Type, AsArray(dt2).ElemType.Type, s)
	case *Map:
		s1 := equal(actual.ElemType.Type, AsMap(dt2).ElemType.Type, s)
		s2 := equal(actual.KeyType.Type, AsMap(dt2).KeyType.Type, s)
		s3 := append(*s1, *s2...)
		return &s3
	case *Object:
		if len(*actual) != len(*AsObject(dt2)) {
			return &fs
		}
		var bs []*bool
		for _, nat := range *actual {
			obj := AsObject(dt2)
			at := obj.Attribute(nat.Name)
			if at == nil {
				return &fs
			}
			bs = append(bs, *equal(nat.Attribute.Type, at.Type, s)...)
		}
		return &bs
	case UserType:
		key := actual.Name() + "=" + dt2.Name()
		if v, ok := s[key]; ok {
			return v
		}
		var res []*bool
		pres := &res
		s[key] = pres
		if IsObject(actual) {
			*pres = *equal(AsObject(dt), AsObject(dt2), s)
		} else if IsMap(actual) {
			// Map aliased as UserType
			*pres = *equal(AsMap(dt), AsMap(dt2), s)
		} else if IsArray(actual) {
			// Array aliased as UserType or CollectionOf
			*pres = *equal(AsArray(dt), AsArray(dt2), s)
		} else {
			// Primitive type aliased as UserType
			*pres = *equal(dt, dt2, s)
		}
		return pres
	}

	t := true
	ts := []*bool{&t}
	return &ts
}

// DataType implementation

// Kind implements DataKind.
func (p Primitive) Kind() Kind { return Kind(p) }

// Name returns the type name appropriate for logging.
func (p Primitive) Name() string {
	switch p {
	case Boolean:
		return "boolean"
	case Int:
		return "int"
	case Int32:
		return "int32"
	case Int64:
		return "int64"
	case UInt:
		return "uint"
	case UInt32:
		return "uint32"
	case UInt64:
		return "uint64"
	case Float32:
		return "float32"
	case Float64:
		return "float64"
	case String:
		return "string"
	case Bytes:
		return "bytes"
	case Any:
		return "any"
	default:
		panic("unknown primitive type") // bug
	}
}

// IsCompatible returns true if val is compatible with p.
func (p Primitive) IsCompatible(val interface{}) bool {
	if p == Any {
		return true
	}
	switch val.(type) {
	case bool:
		return p == Boolean
	case int, int8, int16, int32, uint, uint8, uint16, uint32:
		return p == Int || p == Int32 || p == Int64 ||
			p == UInt || p == UInt32 || p == UInt64 ||
			p == Float32 || p == Float64
	case int64, uint64:
		return p == Int64 || p == UInt64 || p == Float32 || p == Float64
	case float32, float64:
		return p == Float32 || p == Float64
	case string:
		return p == String || p == Bytes
	case []byte:
		return p == Bytes
	}
	return false
}

// Example generates a pseudo-random primitive value using the given random
// generator.
func (p Primitive) Example(r *Random) interface{} {
	switch p {
	case Boolean:
		return r.Bool()
	case Int:
		return r.Int()
	case Int32:
		return r.Int32()
	case Int64:
		return r.Int64()
	case UInt:
		return r.UInt()
	case UInt32:
		return r.UInt32()
	case UInt64:
		return r.UInt64()
	case Float32:
		return r.Float32()
	case Float64:
		return r.Float64()
	case String, Any:
		return r.String()
	case Bytes:
		return []byte(r.String())
	default:
		panic("unknown primitive type") // bug
	}
}

// Hash returns a unique hash value for p.
func (p Primitive) Hash() string {
	return p.Name()
}

// Kind implements DataKind.
func (a *Array) Kind() Kind { return ArrayKind }

// Name returns the type name.
func (a *Array) Name() string {
	return "array"
}

// Hash returns a unique hash value for a.
func (a *Array) Hash() string {
	return "_array_+" + a.ElemType.Type.Hash()
}

// IsCompatible returns true if val is compatible with p.
func (a *Array) IsCompatible(val interface{}) bool {
	k := reflect.TypeOf(val).Kind()
	if k != reflect.Array && k != reflect.Slice {
		return false
	}
	v := reflect.ValueOf(val)
	for i := 0; i < v.Len(); i++ {
		compat := (a.ElemType.Type != nil) && a.ElemType.Type.IsCompatible(v.Index(i).Interface())
		if !compat {
			return false
		}
	}
	return true
}

// Example generates a pseudo-random array value using the given random
// generator.
func (a *Array) Example(r *Random) interface{} {
	count := r.Int()%3 + 2
	res := make([]interface{}, count)
	for i := 0; i < count; i++ {
		res[i] = a.ElemType.Example(r)
		if res[i] == nil {
			// Handle the case of recursive data structures
			res[i] = make(map[string]interface{})
		}
	}
	return a.MakeSlice(res)
}

// MakeSlice examines the key type from the Array and create a slice with
// builtin type if possible. The idea is to avoid generating []interface{} and
// produce more precise types.
func (a *Array) MakeSlice(s []interface{}) interface{} {
	slice := reflect.MakeSlice(toReflectType(a), 0, len(s))
	for _, item := range s {
		slice = reflect.Append(slice, reflect.ValueOf(item))
	}
	return slice.Interface()
}

// ToSlice converts an ArrayVal into a slice.
func (a ArrayVal) ToSlice() []interface{} {
	arr := make([]interface{}, len(a))
	for i, elem := range a {
		switch actual := elem.(type) {
		case ArrayVal:
			arr[i] = actual.ToSlice()
		case MapVal:
			arr[i] = actual.ToMap()
		default:
			arr[i] = actual
		}
	}
	return arr
}

// Attribute returns the attribute with the given name if any, nil otherwise.
func (o *Object) Attribute(name string) *AttributeExpr {
	for _, nat := range *o {
		if nat.Name == name {
			return nat.Attribute
		}
	}
	return nil
}

// Set replaces the object named attribute n if any - creates a new object by
// appending to the slice of named attributes otherwise. The resulting object is
// returned in both cases.
func (o *Object) Set(n string, att *AttributeExpr) {
	for _, nat := range *o {
		if nat.Name == n {
			nat.Attribute = att
			return
		}
	}
	*o = append(*o, &NamedAttributeExpr{n, att})
}

// Delete creates a new object with the same named attributes as o but without
// the named attribute n if any.
func (o *Object) Delete(n string) {
	index := -1
	for i, nat := range *o {
		if nat.Name == n {
			index = i
			break
		}
	}
	if index == -1 {
		return
	}
	*o = append((*o)[:index], (*o)[index+1:]...)
}

// Rename changes the name of the named attribute n to m. Rename does nothing if
// o does not have an attribute named n.
func (o *Object) Rename(n, m string) {
	for _, nat := range *o {
		if nat.Name == n {
			nat.Name = m
			return
		}
	}
}

// Kind implements DataKind.
func (o *Object) Kind() Kind { return ObjectKind }

// Name returns the type name.
func (o *Object) Name() string { return "object" }

// Hash returns a unique hash value for o.
func (o *Object) Hash() string {
	h := "_object_"
	for _, nat := range *o {
		h += "+" + nat.Name + "/" + nat.Attribute.Type.Hash()
	}
	return h
}

// Merge creates a new object consisting of the named attributes of o appended
// with duplicates of the named attributes of other. Named attributes of o that
// have an identical name to named attributes of other get overridden.
func (o *Object) Merge(other *Object) *Object {
	res := o
	for _, nat := range *other {
		res.Set(nat.Name, DupAtt(nat.Attribute))
	}
	return res
}

// IsCompatible returns true if o describes the (Go) type of val.
func (o *Object) IsCompatible(val interface{}) bool {
	k := reflect.TypeOf(val).Kind()
	return k == reflect.Map || k == reflect.Struct
}

// Example returns a random value of the object.
func (o *Object) Example(r *Random) interface{} {
	res := make(map[string]interface{})
	for _, nat := range *o {
		if v := nat.Attribute.Example(r); v != nil {
			res[nat.Name] = v
		}
	}
	return res
}

// Kind implements DataKind.
func (m *Map) Kind() Kind { return MapKind }

// Name returns the type name.
func (m *Map) Name() string { return "map" }

// Hash returns a unique hash value for m.
func (m *Map) Hash() string {
	return "_map_+" + m.KeyType.Type.Hash() + ":" + m.ElemType.Type.Hash()
}

// IsCompatible returns true if o describes the (Go) type of val.
func (m *Map) IsCompatible(val interface{}) bool {
	k := reflect.TypeOf(val).Kind()
	if k != reflect.Map {
		return false
	}
	v := reflect.ValueOf(val)
	for _, key := range v.MapKeys() {
		keyCompat := m.KeyType.Type == nil || m.KeyType.Type.IsCompatible(key.Interface())
		elemCompat := m.ElemType.Type == nil || m.ElemType.Type.IsCompatible(v.MapIndex(key).Interface())
		if !keyCompat || !elemCompat {
			return false
		}
	}
	return true
}

// Example returns a random hash value.
func (m *Map) Example(r *Random) interface{} {
	if IsObject(m.KeyType.Type) || IsArray(m.KeyType.Type) || IsMap(m.KeyType.Type) {
		// not much we can do for non hashable Go types
		return nil
	}
	count := r.Int()%3 + 1
	pair := map[interface{}]interface{}{}
	for i := 0; i < count; i++ {
		k := m.KeyType.Example(r)
		v := m.ElemType.Example(r)
		if k != nil && v != nil {
			pair[k] = v
		}
	}
	return m.MakeMap(pair)
}

// MakeMap examines the key type from a Map and create a map with builtin type
// if possible. The idea is to avoid generating map[interface{}]interface{},
// which cannot be handled by json.Marshal.
func (m *Map) MakeMap(raw map[interface{}]interface{}) interface{} {
	ma := reflect.MakeMap(toReflectType(m))
	for key, value := range raw {
		ma.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(value))
	}
	return ma.Interface()
}

// ToMap converts a MapVal to a map.
func (m MapVal) ToMap() map[interface{}]interface{} {
	mp := make(map[interface{}]interface{}, len(m))
	for k, v := range m {
		switch actual := v.(type) {
		case ArrayVal:
			mp[k] = actual.ToSlice()
		case MapVal:
			mp[k] = actual.ToMap()
		default:
			mp[k] = actual
		}
	}
	return mp
}

// QualifiedTypeName returns the qualified type name for the given data type.
// The qualified type name includes the name of the type of the elements of
// array or map types. This is useful in reporting types in error messages,
// examples of qualified type names:
//
//     "array<string>"
//     "map<string, string>"
//     "map<string, array<int32>>"
//
func QualifiedTypeName(t DataType) string {
	switch t.Kind() {
	case ArrayKind:
		a := t.(*Array)
		return fmt.Sprintf("%s<%s>",
			t.Name(),
			QualifiedTypeName(a.ElemType.Type),
		)
	case MapKind:
		h := t.(*Map)
		return fmt.Sprintf("%s<%s, %s>",
			t.Name(),
			QualifiedTypeName(h.KeyType.Type),
			QualifiedTypeName(h.ElemType.Type),
		)
	}
	return t.Name()
}

// toReflectType converts the DataType to reflect.Type.
func toReflectType(dtype DataType) reflect.Type {
	switch dtype.Kind() {
	case BooleanKind:
		return reflect.TypeOf(true)
	case IntKind:
		return reflect.TypeOf(int(0))
	case Int32Kind:
		return reflect.TypeOf(int32(0))
	case Int64Kind:
		return reflect.TypeOf(int64(0))
	case UIntKind:
		return reflect.TypeOf(uint(0))
	case UInt32Kind:
		return reflect.TypeOf(uint32(0))
	case UInt64Kind:
		return reflect.TypeOf(uint64(0))
	case Float32Kind:
		return reflect.TypeOf(float32(0))
	case Float64Kind:
		return reflect.TypeOf(float64(0))
	case StringKind:
		return reflect.TypeOf("")
	case BytesKind:
		return reflect.TypeOf([]byte{})
	case ObjectKind:
		return reflect.TypeOf(map[string]interface{}{})
	case UserTypeKind:
		return toReflectType(dtype.(*UserTypeExpr).Attribute().Type)
	case ResultTypeKind:
		return toReflectType(dtype.(*ResultTypeExpr).Attribute().Type)
	case ArrayKind:
		return reflect.SliceOf(toReflectType(dtype.(*Array).ElemType.Type))
	case MapKind:
		m := dtype.(*Map)
		// avoid complication: not allow object as the map key
		var ktype reflect.Type
		if m.KeyType.Type.Kind() != ObjectKind {
			ktype = toReflectType(m.KeyType.Type)
		} else {
			ktype = reflect.TypeOf([]interface{}{}).Elem()
		}
		return reflect.MapOf(ktype, toReflectType(m.ElemType.Type))
	default:
		return reflect.TypeOf([]interface{}{}).Elem()
	}
}
