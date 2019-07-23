package codegen

import (
	"fmt"

	"goa.design/goa/v3/expr"
)

type (
	// Attributor defines the behavior of an attribute expression during code
	// generation.
	Attributor interface {
		Scoper
		// Name generates a valid name for the given attribute type.
		Name(att *expr.AttributeExpr, pkg string) string
		// Ref generates a valid reference to the given attribute type.
		Ref(att *expr.AttributeExpr, pkg string) string
		// Field generates a valid data structure field identifier for the given
		// attribute and field name. If firstUpper is true the field name's first
		// letter is capitalized.
		Field(att *expr.AttributeExpr, name string, firstUpper bool) string
	}

	// AttributeContext contains properties which impacts the code generating
	// behavior of an attribute.
	AttributeContext struct {
		// Pointer if true indicates that the attribute uses pointers to hold
		// primitive types even if they are required or has a default value.
		// It ignores UseDefault and IgnoreRequired properties.
		Pointer bool
		// IgnoreRequired if true indicates that the attribute uses non-pointers
		// to hold optional attributes (i.e. attributes that are not required).
		IgnoreRequired bool
		// UseDefault if true indicates that the attribute uses non-pointers for
		// primitive types if they have default value. If false, the attribute with
		// primitive types are non-pointers if they are required, otherwise they
		// are pointers.
		UseDefault bool
		// Pkg is the package name where the attribute type is found.
		Pkg string
		// Scope is the attribute scope.
		Scope Attributor
	}

	// AttributeScope contains the scope of an attribute. It implements the
	// Attributor interface.
	AttributeScope struct {
		// scope is the name scope for the attribute.
		scope *NameScope
	}

	// TransformAttrs are the attributes that help in the transformation.
	TransformAttrs struct {
		// SourceCtx and TargetCtx are the source and target attribute context.
		SourceCtx, TargetCtx *AttributeContext
		// Prefix is the transform function helper prefix.
		Prefix string
	}

	// TransformFunctionData describes a helper function used to transform
	// user types. These are necessary to prevent potential infinite
	// recursion when a type attribute is defined recursively. For example:
	//
	//     var Recursive = Type("Recursive", func() {
	//         Attribute("r", "Recursive")
	//     }
	//
	// Transforming this type requires generating an intermediary function:
	//
	//     func recursiveToRecursive(r *Recursive) *service.Recursive {
	//         var t service.Recursive
	//         if r.R != nil {
	//             t.R = recursiveToRecursive(r.R)
	//         }
	//    }
	//
	TransformFunctionData struct {
		Name          string
		ParamTypeRef  string
		ResultTypeRef string
		Code          string
	}
)

// NewAttributeContext initializes an attribute context.
func NewAttributeContext(pointer, reqIgnore, useDefault bool, pkg string, scope *NameScope) *AttributeContext {
	return &AttributeContext{
		Pointer:        pointer,
		IgnoreRequired: reqIgnore,
		UseDefault:     useDefault,
		Pkg:            pkg,
		Scope:          &AttributeScope{scope: scope},
	}
}

// IsCompatible returns an error if a and b are not both objects, both arrays,
// both maps or both the same primitive type. actx and bctx are used to build
// the error message if any.
func IsCompatible(a, b expr.DataType, actx, bctx string) error {
	switch {
	case expr.IsObject(a):
		if !expr.IsObject(b) {
			return fmt.Errorf("%s is an object but %s type is %s", actx, bctx, b.Name())
		}
	case expr.IsArray(a):
		if !expr.IsArray(b) {
			return fmt.Errorf("%s is an array but %s type is %s", actx, bctx, b.Name())
		}
	case expr.IsMap(a):
		if !expr.IsMap(b) {
			return fmt.Errorf("%s is a hash but %s type is %s", actx, bctx, b.Name())
		}
	default:
		if ut, ok := a.(expr.UserType); ok {
			return IsCompatible(ut.Attribute().Type, b, actx, bctx)
		} else if ut, ok := b.(expr.UserType); ok {
			return IsCompatible(a, ut.Attribute().Type, actx, bctx)
		}
		if a.Kind() != b.Kind() {
			return fmt.Errorf("%s is a %s but %s type is %s", actx, a.Name(), bctx, b.Name())
		}
	}
	return nil
}

// AppendHelpers takes care of only appending helper functions from newH that
// are not already in oldH.
func AppendHelpers(oldH, newH []*TransformFunctionData) []*TransformFunctionData {
	for _, h := range newH {
		found := false
		for _, h2 := range oldH {
			if h.Name == h2.Name {
				found = true
				break
			}
		}
		if !found {
			oldH = append(oldH, h)
		}
	}
	return oldH
}

// MapDepth returns the level of nested maps. For unnested maps, it returns 0.
func MapDepth(m *expr.Map) int {
	return mapDepth(m.ElemType.Type, 0)
}

func mapDepth(dt expr.DataType, depth int, seen ...map[string]struct{}) int {
	if mp := expr.AsMap(dt); mp != nil {
		depth++
		depth = mapDepth(mp.ElemType.Type, depth, seen...)
	} else if ar := expr.AsArray(dt); ar != nil {
		depth = mapDepth(ar.ElemType.Type, depth, seen...)
	} else if mo := expr.AsObject(dt); mo != nil {
		var s map[string]struct{}
		if len(seen) > 0 {
			s = seen[0]
		} else {
			s = make(map[string]struct{})
			seen = append(seen, s)
		}
		key := dt.Name()
		if u, ok := dt.(expr.UserType); ok {
			key = u.ID()
		}
		if _, ok := s[key]; ok {
			return depth
		}
		s[key] = struct{}{}
		var level int
		for _, nat := range *mo {
			// if object type has attributes of type map then find out the attribute that has
			// the deepest level of nested maps
			lvl := 0
			lvl = mapDepth(nat.Attribute.Type, lvl, seen...)
			if lvl > level {
				level = lvl
			}
		}
		depth += level
	}
	return depth
}

// IsPrimitivePointer returns true if the attribute with the given name is a
// primitive pointer in the given parent attribute.
func (a *AttributeContext) IsPrimitivePointer(name string, att *expr.AttributeExpr) bool {
	if at := att.Find(name); at != nil && at.Type == expr.Any || at.Type == expr.Bytes {
		return false
	}
	if a.Pointer {
		return true
	}
	if a.IgnoreRequired {
		return false
	}
	return att.IsPrimitivePointer(name, a.UseDefault)
}

// IsRequired returns true if the attribute with given name is a required
// attribute in the parent. If IgnoreRequired is set to true, IsRequired always
// returns false.
func (a *AttributeContext) IsRequired(name string, att *expr.AttributeExpr) bool {
	if a.IgnoreRequired {
		return false
	}
	return att.IsRequired(name)
}

// Dup creates a shallow copy of the AttributeContext.
func (a *AttributeContext) Dup() *AttributeContext {
	return &AttributeContext{
		Pointer:        a.Pointer,
		IgnoreRequired: a.IgnoreRequired,
		UseDefault:     a.UseDefault,
		Scope:          a.Scope,
	}
}

// Name returns the type name for the given attribute.
func (a *AttributeScope) Name(att *expr.AttributeExpr, pkg string) string {
	return a.scope.GoFullTypeName(att, pkg)
}

// Ref returns the type name for the given attribute.
func (a *AttributeScope) Ref(att *expr.AttributeExpr, pkg string) string {
	return a.scope.GoFullTypeRef(att, pkg)
}

// Field returns a valid Go struct field name.
func (a *AttributeScope) Field(att *expr.AttributeExpr, name string, firstUpper bool) string {
	return GoifyAtt(att, name, firstUpper)
}

// Scope returns the name scope.
func (a *AttributeScope) Scope() *NameScope {
	return a.scope
}
