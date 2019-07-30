package expr

import (
	"fmt"

	"goa.design/goa/v3/eval"
)

type (
	// AttributeExpr defines an object field with optional description,
	// default value and validations.
	AttributeExpr struct {
		// DSLFunc contains the DSL used to initialize the expression.
		eval.DSLFunc
		// Attribute type
		Type DataType
		// Base types if any
		Bases []DataType
		// Attribute reference types if any
		References []DataType
		// Optional description
		Description string
		// Docs points to external documentation
		Docs *DocsExpr
		// Optional validations
		Validation *ValidationExpr
		// Meta is a list of key/value pairs
		Meta MetaExpr
		// Optional member default value
		DefaultValue interface{}
		// ZeroValue sets the zero value for the attribute type
		ZeroValue interface{}
		// UserExample set in DSL or computed in Finalize
		UserExamples []*ExampleExpr
	}

	// ExampleExpr represents an example.
	ExampleExpr struct {
		// Summary is the example short summary.
		Summary string
		// Description is an optional long description.
		Description string
		// Value is the example value.
		Value interface{}
	}

	// Val is the type used to provide the value of examples for attributes that are
	// objects.
	Val map[string]interface{}

	// CompositeExpr defines a generic composite expression that contains an
	// attribute.  This makes it possible for plugins to use attributes in
	// their own data structures.
	CompositeExpr interface {
		// Attribute returns the composite expression embedded attribute.
		Attribute() *AttributeExpr
	}

	// ValidationExpr contains validation rules for an attribute.
	ValidationExpr struct {
		// Values represents an enum validation as described at
		// http://json-schema.org/latest/json-schema-validation.html#anchor76.
		Values []interface{}
		// Format represents a format validation as described at
		// http://json-schema.org/latest/json-schema-validation.html#anchor104.
		Format ValidationFormat
		// PatternValidationExpr represents a pattern validation as
		// described at
		// http://json-schema.org/latest/json-schema-validation.html#anchor33
		Pattern string
		// Minimum represents an minimum value validation as described
		// at
		// http://json-schema.org/latest/json-schema-validation.html#anchor21.
		Minimum *float64
		// Maximum represents a maximum value validation as described at
		// http://json-schema.org/latest/json-schema-validation.html#anchor17.
		Maximum *float64
		// MinLength represents an minimum length validation as
		// described at
		// http://json-schema.org/latest/json-schema-validation.html#anchor29.
		MinLength *int
		// MaxLength represents an maximum length validation as
		// described at
		// http://json-schema.org/latest/json-schema-validation.html#anchor26.
		MaxLength *int
		// Required list the required fields of object attributes as
		// described at
		// http://json-schema.org/latest/json-schema-validation.html#anchor61.
		Required []string
	}

	// ValidationFormat is the type used to enumerate the possible string
	// formats.
	ValidationFormat string
)

const (
	// FormatDate describes RFC3339 date values.
	FormatDate ValidationFormat = "date"

	// FormatDateTime describes RFC3339 date time values.
	FormatDateTime ValidationFormat = "date-time"

	// FormatUUID describes RFC4122 UUID values.
	FormatUUID = "uuid"

	// FormatEmail describes RFC5322 email addresses.
	FormatEmail = "email"

	// FormatHostname describes RFC1035 Internet hostnames.
	FormatHostname = "hostname"

	// FormatIPv4 describes RFC2373 IPv4 address values.
	FormatIPv4 = "ipv4"

	// FormatIPv6 describes RFC2373 IPv6 address values.
	FormatIPv6 = "ipv6"

	// FormatIP describes RFC2373 IPv4 or IPv6 address values.
	FormatIP = "ip"

	// FormatURI describes RFC3986 URI values.
	FormatURI = "uri"

	// FormatMAC describes IEEE 802 MAC-48, EUI-48 or EUI-64 MAC address values.
	FormatMAC = "mac"

	// FormatCIDR describes RFC4632 and RFC4291 CIDR notation IP address values.
	FormatCIDR = "cidr"

	// FormatRegexp describes regular expression syntax accepted by RE2.
	FormatRegexp = "regexp"

	// FormatJSON describes JSON text.
	FormatJSON = "json"

	// FormatRFC1123 describes RFC1123 date time values.
	FormatRFC1123 = "rfc1123"
)

// EvalName returns the name used by the DSL evaluation.
func (a *AttributeExpr) EvalName() string {
	return "attribute"
}

// validated keeps track of validated attributes to handle cyclical definitions.
var validated = make(map[*AttributeExpr]bool)

// TaggedAttribute returns the name of the child attribute of a with the given
// tag if a is an object.
func TaggedAttribute(a *AttributeExpr, tag string) string {
	obj := AsObject(a.Type)
	if obj == nil {
		return ""
	}
	for _, at := range *obj {
		if _, ok := at.Attribute.Meta[tag]; ok {
			return at.Name
		}
	}
	for _, b := range a.Bases {
		at := &AttributeExpr{Type: b}
		if ut, ok := b.(UserType); ok {
			at = ut.Attribute()
		}
		if n := TaggedAttribute(at, tag); n != "" {
			return n
		}
	}
	return ""
}

// Prepare initializes the Meta expression.
func (a *AttributeExpr) Prepare() {
	if a.Meta == nil {
		a.Meta = MetaExpr{}
	}
}

// Validate tests whether the attribute required fields exist.  Since attributes
// are unaware of their context, additional context information can be provided
// to be used in error messages.  The parent definition context is automatically
// added to error messages.
func (a *AttributeExpr) Validate(ctx string, parent eval.Expression) *eval.ValidationErrors {
	if validated[a] {
		return nil
	}
	validated[a] = true
	verr := new(eval.ValidationErrors)
	if a.Type == nil {
		verr.Add(parent, "attribute type is nil")
		return verr
	}
	if ctx != "" {
		ctx += " - "
	}
	verr.Merge(a.validateEnumDefault(ctx, parent))
	if o := AsObject(a.Type); o != nil {
		for _, n := range a.AllRequired() {
			if a.Find(n) == nil {
				verr.Add(parent, `%srequired field %q does not exist in type %s`, ctx, n, a.Type.Name())
			}
		}
		for _, nat := range *o {
			ctx = fmt.Sprintf("field %s", nat.Name)
			verr.Merge(nat.Attribute.Validate(ctx, parent))
		}
	} else {
		if ar := AsArray(a.Type); ar != nil {
			elemType := ar.ElemType
			verr.Merge(elemType.Validate(ctx, a))
		}
	}

	if views, ok := a.Meta["view"]; ok {
		rt, ok := a.Type.(*ResultTypeExpr)
		if !ok {
			verr.Add(parent, "%sdefines a view %v but type %s is not a result type", ctx, views, a.Type.Name())
		}
		if name := views[0]; name != "default" && rt != nil {
			found := false
			for _, v := range rt.Views {
				if v.Name == name {
					found = true
					break
				}
			}
			if !found {
				verr.Add(parent, "%stype %s does not define view %q", ctx, a.Type.Name(), name)
			}
		}
	}

	return verr
}

// Finalize merges base and reference type attributes and finalizes the Type
// attribute.
func (a *AttributeExpr) Finalize() {
	if ut, ok := a.Type.(UserType); ok {
		ut.Finalize()
	}
	if IsObject(a.Type) {
		for _, ref := range a.References {
			ru, ok := ref.(UserType)
			if !ok {
				continue
			}
			a.Inherit(ru.Attribute())
		}
		for _, base := range a.Bases {
			ru, ok := base.(UserType)
			if !ok {
				continue
			}
			a.Merge(ru.Attribute())
		}
	}
}

// Merge merges other's attributes into a overriding attributes of a with
// attributes of other with identical names.
//
// This only applies to attributes of type Object and Merge panics if the
// argument or the target is not of type Object.
func (a *AttributeExpr) Merge(other *AttributeExpr) {
	if other == nil {
		return
	}
	left := AsObject(a.Type)
	right := AsObject(other.Type)
	if left == nil || right == nil {
		panic("cannot merge non object attributes") // bug
	}
	if a.Type == Empty && len(*right) > 0 {
		a.Type = &Object{}
		left = AsObject(a.Type)
	}
	if other.Validation != nil {
		if a.Validation == nil {
			a.Validation = other.Validation.Dup()
		} else {
			a.Validation.Merge(other.Validation)
		}
	}
	for _, nat := range *right {
		left.Set(nat.Name, nat.Attribute)
	}
}

// Inherit merges the properties of existing target type attributes with the
// argument's. The algorithm is recursive so that child attributes are also
// merged.
func (a *AttributeExpr) Inherit(parent *AttributeExpr) {
	if !a.shouldInherit(parent) {
		return
	}
	pobj := AsObject(parent.Type)
	if a.Type == Empty && len(*pobj) > 0 {
		a.Type = &Object{}
	}
	a.inheritValidations(parent)
	a.inheritRecursive(parent, make(map[*AttributeExpr]struct{}))
}

// AllRequired returns the list of all required fields from the underlying
// object. This method recurses if the type is itself an attribute (i.e. a
// UserType, this happens with the Reference DSL for example).
func (a *AttributeExpr) AllRequired() []string {
	if u, ok := a.Type.(UserType); ok {
		return u.Attribute().AllRequired()
	}
	if a.Validation != nil {
		return a.Validation.Required
	}
	return nil
}

// IsRequired returns true if the given string matches the name of a required
// attribute, false otherwise. This method only applies to attributes of type
// Object.
func (a *AttributeExpr) IsRequired(attName string) bool {
	for _, name := range a.AllRequired() {
		if name == attName {
			return true
		}
	}
	return false
}

// IsRequiredNoDefault returns true if the given string matches the name of a
// required attribute and the attribute has no default value, false otherwise.
// This method only applies to attributes of type Object.
func (a *AttributeExpr) IsRequiredNoDefault(attName string) bool {
	for _, name := range a.AllRequired() {
		if name == attName {
			return a.GetDefault(name) == nil
		}
	}
	return false
}

// IsPrimitivePointer returns true if the field generated for the given
// attribute should be a pointer to a primitive type. The receiver attribute must
// be an object.
//
// If useDefault is true and the attribute has a default value then
// IsPrimitivePointer returns false. This makes it possible to differentiate
// between request types where attributes with default values should not be
// generated using a pointer value and response types where they should.
//
//    DefaultValue UseDefault Pointer (assuming all other conditions are true)
//    Yes          True       False
//    Yes          False      True
//    No           True       True
//    No           False      True
//
func (a *AttributeExpr) IsPrimitivePointer(attName string, useDefault bool) bool {
	o := AsObject(a.Type)
	if o == nil {
		panic("checking pointer field on non-object") // bug
	}
	att := o.Attribute(attName)
	if att == nil {
		return false
	}
	if IsPrimitive(att.Type) {
		return att.Type.Kind() != BytesKind && att.Type.Kind() != AnyKind &&
			!a.IsRequired(attName) && (!a.HasDefaultValue(attName) || !useDefault)
	}
	return false
}

// HasTag returns true if the attribute is an object that has an attribute with
// the given tag.
func (a *AttributeExpr) HasTag(tag string) bool {
	if a == nil {
		return false
	}
	obj := AsObject(a.Type)
	if obj == nil {
		return false
	}
	for _, at := range *obj {
		if _, ok := at.Attribute.Meta[tag]; ok {
			return true
		}
	}
	return false
}

// FieldTag returns the field tag if the attribute is a field.
func (a *AttributeExpr) FieldTag() (tag string, found bool) {
	if a == nil {
		return
	}
	return a.Meta.Last("rpc:tag")
}

// HasDefaultValue returns true if the attribute with the given name has a
// default value.
func (a *AttributeExpr) HasDefaultValue(attName string) bool {
	return a.GetDefault(attName) != nil
}

// GetDefault gets the default value for the child attribute with the given
// name. It returns nil if the child attribute with the given name does not
// exist or if the child attribute does not have a default value.
func (a *AttributeExpr) GetDefault(attName string) interface{} {
	if o := AsObject(a.Type); o != nil {
		att := o.Attribute(attName)
		if att.DefaultValue != nil {
			return att.DefaultValue
		}
		if ut, ok := att.Type.(UserType); ok && !IsObject(ut) {
			return ut.Attribute().DefaultValue
		}
	}
	return nil
}

// SetDefault sets the default for the attribute. It also converts HashVal
// and ArrayVal to map and slice respectively.
func (a *AttributeExpr) SetDefault(def interface{}) {
	switch actual := def.(type) {
	case MapVal:
		a.DefaultValue = actual.ToMap()
	case ArrayVal:
		a.DefaultValue = actual.ToSlice()
	default:
		a.DefaultValue = actual
	}
}

// Find finds a child attribute with the given name in the attribute and
// its bases and references. If the parent attribute is not an object, it
// returns nil.
func (a *AttributeExpr) Find(name string) *AttributeExpr {
	findAttrFn := func(typ DataType) *AttributeExpr {
		switch t := typ.(type) {
		case UserType:
			return t.Attribute().Find(name)
		case *Object:
			if att := t.Attribute(name); att != nil {
				return att
			}
		}
		return nil
	}

	if att := findAttrFn(a.Type); att != nil {
		return att
	}
	for _, b := range a.Bases {
		if att := findAttrFn(b); att != nil {
			return att
		}
	}
	return nil
}

// Delete removes an attribute with the given name. It does nothing if the
// attribute expression is not an object type.
func (a *AttributeExpr) Delete(name string) {
	switch t := a.Type.(type) {
	case UserType:
		t.Attribute().Delete(name)
	case *Object:
		AsObject(t).Delete(name)
		if a.Validation != nil {
			a.Validation.RemoveRequired(name)
		}
		for _, ex := range a.UserExamples {
			if m, ok := ex.Value.(map[string]interface{}); ok {
				delete(m, name)
			}
		}
	}
}

// Debug dumps the attribute to STDOUT in a goa developer friendly way.
func (a *AttributeExpr) Debug(prefix string) { a.debug(prefix, make(map[*AttributeExpr]int), 0) }
func (a *AttributeExpr) debug(prefix string, seen map[*AttributeExpr]int, indent int) {
	for i := 0; i < indent; i++ {
		prefix = "  " + prefix
	}
	if c, ok := seen[a]; ok && c > 1 {
		fmt.Printf("%s: ...\n", prefix)
		return
	}
	seen[a]++
	fmt.Printf("%s: %q\n", prefix, a.Type.Name())
	if o := AsObject(a.Type); o != nil {
		for _, att := range *o {
			att.Attribute.debug(" - "+att.Name, seen, indent+1)
		}
	}
	if a := AsArray(a.Type); a != nil {
		a.ElemType.debug(" Elem", seen, indent+2)
	}
	if m := AsMap(a.Type); m != nil {
		m.KeyType.debug(" Key", seen, indent+2)
		m.ElemType.debug(" Elem", seen, indent+2)
	}
}

// validateEnumDefault makes sure that the attribute default value is one of the
// enum values.
func (a *AttributeExpr) validateEnumDefault(ctx string, parent eval.Expression) *eval.ValidationErrors {
	//TODO: We only do the default value and enum check just for primitive types.
	if _, ok := a.Type.(Primitive); !ok {
		return nil
	}
	verr := new(eval.ValidationErrors)
	if a.DefaultValue != nil && a.Validation != nil && a.Validation.Values != nil {
		var found bool
		for _, e := range a.Validation.Values {
			if e == a.DefaultValue {
				found = true
				break
			}
		}
		if !found {
			verr.Add(
				parent,
				"%sdefault value %#v is not one of the accepted values: %#v",
				ctx,
				a.DefaultValue,
				a.Validation.Values,
			)
		}
	}
	return verr
}

func (a *AttributeExpr) inheritRecursive(parent *AttributeExpr, seen map[*AttributeExpr]struct{}) {
	if !a.shouldInherit(parent) {
		return
	}
	for _, nat := range *AsObject(a.Type) {
		if patt := AsObject(parent.Type).Attribute(nat.Name); patt != nil {
			att := nat.Attribute
			if att.Description == "" {
				att.Description = patt.Description
			}
			att.inheritValidations(patt)
			if att.DefaultValue == nil {
				att.DefaultValue = patt.DefaultValue
			}
			if att.Type == nil {
				att.Type = patt.Type
			} else if att.shouldInherit(patt) {
				if _, ok := seen[att]; ok {
					continue
				}
				seen[att] = struct{}{}
				for _, nat := range *AsObject(att.Type) {
					child := nat.Attribute
					parent := AsObject(patt.Type).Attribute(nat.Name)
					child.inheritValidations(parent)
					child.inheritRecursive(parent, seen)
				}
			}
		}
	}
}

func (a *AttributeExpr) inheritValidations(parent *AttributeExpr) {
	if parent.Validation == nil {
		return
	}
	if a.Validation == nil {
		a.Validation = &ValidationExpr{}
	}
	a.Validation.AddRequired(parent.Validation.Required...)
}

func (a *AttributeExpr) shouldInherit(parent *AttributeExpr) bool {
	return a != nil && AsObject(a.Type) != nil &&
		parent != nil && AsObject(parent.Type) != nil
}

// EvalName returns the name used by the DSL evaluation.
func (a *ExampleExpr) EvalName() string {
	return `example "` + a.Summary + `"`
}

// Merge merges other into v.
func (v *ValidationExpr) Merge(other *ValidationExpr) {
	if v.Values == nil {
		v.Values = other.Values
	}
	if v.Format == "" {
		v.Format = other.Format
	}
	if v.Pattern == "" {
		v.Pattern = other.Pattern
	}
	if v.Minimum == nil || (other.Minimum != nil && *v.Minimum > *other.Minimum) {
		v.Minimum = other.Minimum
	}
	if v.Maximum == nil || (other.Maximum != nil && *v.Maximum < *other.Maximum) {
		v.Maximum = other.Maximum
	}
	if v.MinLength == nil || (other.MinLength != nil && *v.MinLength > *other.MinLength) {
		v.MinLength = other.MinLength
	}
	if v.MaxLength == nil || (other.MaxLength != nil && *v.MaxLength < *other.MaxLength) {
		v.MaxLength = other.MaxLength
	}
	v.AddRequired(other.Required...)
}

// AddRequired merges the required fields into v.
func (v *ValidationExpr) AddRequired(required ...string) {
	for _, r := range required {
		found := false
		for _, rr := range v.Required {
			if r == rr {
				found = true
				break
			}
		}
		if !found {
			v.Required = append(v.Required, r)
		}
	}
}

// RemoveRequired removes the given field from the list of required fields.
func (v *ValidationExpr) RemoveRequired(required string) {
	for i, r := range v.Required {
		if required == r {
			v.Required = append(v.Required[:i], v.Required[i+1:]...)
			break
		}
	}
}

// HasRequiredOnly returns true if the validation only has the Required field
// with a non-zero value.
func (v *ValidationExpr) HasRequiredOnly() bool {
	if len(v.Values) > 0 {
		return false
	}
	if v.Format != "" || v.Pattern != "" {
		return false
	}
	if (v.Minimum != nil) || (v.Maximum != nil) || (v.MinLength != nil) || (v.MaxLength != nil) {
		return false
	}
	return true
}

// Dup makes a shallow dup of the validation.
func (v *ValidationExpr) Dup() *ValidationExpr {
	var req []string
	if len(v.Required) > 0 {
		req = make([]string, len(v.Required))
		copy(req, v.Required)
	}
	return &ValidationExpr{
		Values:    v.Values,
		Format:    v.Format,
		Pattern:   v.Pattern,
		Minimum:   v.Minimum,
		Maximum:   v.Maximum,
		MinLength: v.MinLength,
		MaxLength: v.MaxLength,
		Required:  req,
	}
}

// IsSupportedValidationFormat checks if the validation format is supported by goa.
func (a *AttributeExpr) IsSupportedValidationFormat(vf ValidationFormat) bool {
	switch vf {
	case FormatDate:
		return true
	case FormatDateTime:
		return true
	case FormatUUID:
		return true
	case FormatEmail:
		return true
	case FormatHostname:
		return true
	case FormatIPv4:
		return true
	case FormatIPv6:
		return true
	case FormatIP:
		return true
	case FormatURI:
		return true
	case FormatMAC:
		return true
	case FormatCIDR:
		return true
	case FormatRegexp:
		return true
	case FormatJSON:
		return true
	case FormatRFC1123:
		return true
	}
	return false
}

// walkAttribute iterates over the given attribute, its bases and references
// (if any). It calls the given function giving each attribute as it iterates.
// It stops if the given attribute is not an object type or there is no more
// attribute to iterate over or if the iterator function returned an error. It
// is generally used in implementing the Validator interface since attribute
// bases and references are only merged during Finalize. It is not a recursive
// implementation.
func walkAttribute(att *AttributeExpr, it func(name string, a *AttributeExpr) error) error {
	switch dt := att.Type.(type) {
	case UserType:
		if err := walkAttribute(dt.Attribute(), it); err != nil {
			return err
		}
	case *Object:
		for _, nat := range *dt {
			if err := it(nat.Name, nat.Attribute); err != nil {
				return err
			}
		}
	}
	for _, b := range att.Bases {
		if err := walkAttribute(&AttributeExpr{Type: b}, it); err != nil {
			return err
		}
	}
	for _, r := range att.References {
		if err := walkAttribute(&AttributeExpr{Type: r}, it); err != nil {
			return err
		}
	}
	return nil
}
