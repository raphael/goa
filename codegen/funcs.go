package codegen

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"unicode"

	"goa.design/goa/v3/expr"
)

type (
	// InitArgData contains the data needed to render code to initlize struct
	// fields with the given arguments.
	InitArgData struct {
		// Name is the argument name.
		Name string
		// Pointer if true indicates that the argument is a pointer.
		Pointer bool
		// Type is the argument type.
		Type expr.DataType
		// FieldName is the name of the field in the struct initialized by the
		// argument.
		FieldName string
		// FieldPointer if true indicates that the field in the struct is a
		// pointer.
		FieldPointer bool
		// FieldType is the type of the field in the struct.
		FieldType expr.DataType
	}
)

// TemplateFuncs lists common template helper functions.
func TemplateFuncs() map[string]interface{} {
	return map[string]interface{}{
		"commandLine": CommandLine,
		"comment":     Comment,
	}
}

// CommandLine return the command used to run this process.
func CommandLine() string {
	cmdl := "$ goa"
	for _, arg := range os.Args {
		if strings.HasPrefix(arg, "--cmd=") {
			cmdl = arg[6:]
			break
		}
	}
	return cmdl
}

// Comment produces line comments by concatenating the given strings and
// producing 80 characters long lines starting with "//".
func Comment(elems ...string) string {
	var lines []string
	for _, e := range elems {
		lines = append(lines, strings.Split(e, "\n")...)
	}
	var trimmed = make([]string, len(lines))
	for i, l := range lines {
		trimmed[i] = strings.TrimLeft(l, " \t")
	}
	t := strings.Join(trimmed, "\n")

	return Indent(WrapText(t, 77), "// ")
}

// Indent inserts prefix at the beginning of each non-empty line of s. The
// end-of-line marker is NL.
func Indent(s, prefix string) string {
	var (
		res []byte
		b   = []byte(s)
		p   = []byte(prefix)
		bol = true
	)
	for _, c := range b {
		if bol && c != '\n' {
			res = append(res, p...)
		}
		res = append(res, c)
		bol = c == '\n'
	}
	return string(res)
}

// Casing exceptions
var toLower = map[string]string{"OAuth": "oauth"}

// CamelCase produces the CamelCase version of the given string. It removes any
// non letter and non digit character.
//
// If firstUpper is true the first letter of the string is capitalized else
// the first letter is in lowercase.
//
// If acronym is true and a part of the string is a common acronym
// then it keeps the part capitalized (firstUpper = true)
// (e.g. APIVersion) or lowercase (firstUpper = false) (e.g. apiVersion).
func CamelCase(name string, firstUpper bool, acronym bool) string {
	if name == "" {
		return ""
	}

	runes := []rune(name)
	// remove trailing invalid identifiers (makes code below simpler)
	runes = removeTrailingInvalid(runes)

	// all characters are invalid
	if len(runes) == 0 {
		return ""
	}

	w, i := 0, 0 // index of start of word, scan
	for i+1 <= len(runes) {
		eow := false // whether we hit the end of a word

		// remove leading invalid identifiers
		runes = removeInvalidAtIndex(i, runes)

		if i+1 == len(runes) {
			eow = true
		} else if !validIdentifier(runes[i]) {
			// get rid of it
			runes = append(runes[:i], runes[i+1:]...)
		} else if runes[i+1] == '_' {
			// underscore; shift the remainder forward over any run of underscores
			eow = true
			n := 1
			for i+n+1 < len(runes) && runes[i+n+1] == '_' {
				n++
			}
			copy(runes[i+1:], runes[i+n+1:])
			runes = runes[:len(runes)-n]
		} else if isLower(runes[i]) && !isLower(runes[i+1]) {
			// lower->non-lower
			eow = true
		}
		i++
		if !eow {
			continue
		}

		// [w,i] is a word.
		word := string(runes[w:i])
		// is it one of our initialisms?
		if u := strings.ToUpper(word); acronym && commonInitialisms[u] {
			if firstUpper {
				u = strings.ToUpper(u)
			} else if w == 0 {
				u = strings.ToLower(u)
			}

			// All the common initialisms are ASCII,
			// so we can replace the bytes exactly.
			copy(runes[w:], []rune(u))
		} else if w > 0 && strings.ToLower(word) == word {
			// already all lowercase, and not the first word, so uppercase the first character.
			runes[w] = unicode.ToUpper(runes[w])
		} else if w == 0 && strings.ToLower(word) == word && firstUpper {
			runes[w] = unicode.ToUpper(runes[w])
		}
		if w == 0 && !firstUpper {
			runes[w] = unicode.ToLower(runes[w])
		}
		//advance to next word
		w = i
	}

	return string(runes)
}

// SnakeCase produces the snake_case version of the given CamelCase string.
// News    => news
// OldNews => old_news
// CNNNews => cnn_news
func SnakeCase(name string) string {
	// Special handling for single "words" starting with multiple upper case letters
	for u, l := range toLower {
		name = strings.Replace(name, u, l, -1)
	}

	// Special handling for dashes to convert them into underscores
	name = strings.Replace(name, "-", "_", -1)

	var b bytes.Buffer
	ln := len(name)
	if ln == 0 {
		return ""
	}
	n := rune(name[0])
	b.WriteRune(unicode.ToLower(n))
	lastLower, isLower, lastUnder, isUnder := false, true, false, false
	for i := 1; i < ln; i++ {
		r := rune(name[i])
		isLower = unicode.IsLower(r) && unicode.IsLetter(r) || unicode.IsDigit(r)
		isUnder = r == '_'
		if !isLower && !isUnder {
			if lastLower && !lastUnder {
				b.WriteRune('_')
			} else if ln > i+1 {
				rn := rune(name[i+1])
				if unicode.IsLower(rn) && rn != '_' && !lastUnder {
					b.WriteRune('_')
				}
			}
		}
		b.WriteRune(unicode.ToLower(r))
		lastLower = isLower
		lastUnder = isUnder
	}
	return b.String()
}

// KebabCase produces the kebab-case version of the given CamelCase string.
func KebabCase(name string) string {
	name = SnakeCase(name)
	ln := len(name)
	if name[ln-1] == '_' {
		name = name[:ln-1]
	}
	return strings.Replace(name, "_", "-", -1)
}

// WrapText produces lines with text capped at maxChars
// it will keep words intact and respects newlines.
func WrapText(text string, maxChars int) string {
	res := ""
	lines := strings.Split(text, "\n")
	for _, v := range lines {
		runes := []rune(strings.TrimSpace(v))
		for l := len(runes); l >= 0; l = len(runes) {
			if maxChars >= l {
				res = res + string(runes) + "\n"
				break
			}

			i := runeSpacePosRev(runes[:maxChars])
			if i == 0 {
				i = runeSpacePos(runes)
			}

			res = res + string(runes[:i]) + "\n"
			if l == i {
				break
			}
			runes = runes[i+1:]
		}
	}
	return res[:len(res)-1]
}

// InitStructFields produces Go code to initialize a struct and its fields from
// the given init arguments.
func InitStructFields(args []*InitArgData, name, targetVar, sourcePkg, targetPkg string, mustInit bool) (string, []*TransformFunctionData, error) {
	scope := NewNameScope()
	scope.Unique(targetVar)

	var (
		code    string
		helpers []*TransformFunctionData
	)
	for _, arg := range args {
		switch {
		case arg.FieldName == "":
			// do nothing
		case expr.Equal(arg.Type, arg.FieldType):
			// arg type and struct field type are the same. No need to call transform
			// to initialize the field
			deref := ""
			if !arg.Pointer && arg.FieldPointer && expr.IsPrimitive(arg.FieldType) {
				deref = "&"
			}
			code += fmt.Sprintf("%s.%s = %s%s\n", targetVar, arg.FieldName, deref, arg.Name)
		case expr.IsPrimitive(arg.FieldType):
			// aliased primitive type
			t := scope.GoFullTypeRef(&expr.AttributeExpr{Type: arg.FieldType}, targetPkg)
			cast := fmt.Sprintf("%s(%s)", t, arg.Name)
			if arg.Pointer {
				cast = fmt.Sprintf("%s(*%s)", t, arg.Name)
			}
			if arg.FieldPointer {
				code += fmt.Sprintf("tmp%s := %s\n%s.%s = &tmp%s\n", arg.Name, cast, targetVar, arg.FieldName, arg.Name)
			} else {
				code += fmt.Sprintf("%s.%s = %s\n", targetVar, arg.FieldName, cast)
			}
		default:
			srcctx := NewAttributeContext(arg.Pointer, false, true, sourcePkg, scope)
			tgtctx := NewAttributeContext(arg.FieldPointer, false, true, targetPkg, scope)
			c, h, err := GoTransformToVar(
				&expr.AttributeExpr{Type: arg.Type}, &expr.AttributeExpr{Type: arg.FieldType},
				arg.Name, fmt.Sprintf("%s.%s", targetVar, arg.FieldName), srcctx, tgtctx, "", false)
			if err != nil {
				return "", helpers, err
			}
			code += c + "\n"
			helpers = AppendHelpers(helpers, h)
		}
	}
	return code, helpers, nil
}

func runeSpacePosRev(r []rune) int {
	for i := len(r) - 1; i > 0; i-- {
		if unicode.IsSpace(r[i]) {
			return i
		}
	}
	return 0
}

func runeSpacePos(r []rune) int {
	for i := 0; i < len(r); i++ {
		if unicode.IsSpace(r[i]) {
			return i
		}
	}
	return len(r)
}

// isLower returns true if the character is considered a lower case character
// when transforming word into CamelCase.
func isLower(r rune) bool {
	return unicode.IsDigit(r) || unicode.IsLower(r)
}

// validIdentifier returns true if the rune is a letter or number
func validIdentifier(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}

// removeTrailingInvalid removes trailing invalid identifiers from runes.
func removeTrailingInvalid(runes []rune) []rune {
	valid := len(runes) - 1
	for ; valid >= 0 && !validIdentifier(runes[valid]); valid-- {
	}

	return runes[0 : valid+1]
}

// removeInvalidAtIndex removes consecutive invalid identifiers from runes starting at index i.
func removeInvalidAtIndex(i int, runes []rune) []rune {
	valid := i
	for ; valid < len(runes) && !validIdentifier(runes[valid]); valid++ {
	}

	return append(runes[:i], runes[valid:]...)
}

var (
	// common words who need to keep their
	commonInitialisms = map[string]bool{
		"API":   true,
		"ASCII": true,
		"CPU":   true,
		"CSS":   true,
		"DNS":   true,
		"EOF":   true,
		"GUID":  true,
		"HTML":  true,
		"HTTP":  true,
		"HTTPS": true,
		"ID":    true,
		"IP":    true,
		"JMES":  true,
		"JSON":  true,
		"JWT":   true,
		"LHS":   true,
		"OK":    true,
		"QPS":   true,
		"RAM":   true,
		"RHS":   true,
		"RPC":   true,
		"SLA":   true,
		"SMTP":  true,
		"SQL":   true,
		"SSH":   true,
		"TCP":   true,
		"TLS":   true,
		"TTL":   true,
		"UDP":   true,
		"UI":    true,
		"UID":   true,
		"UUID":  true,
		"URI":   true,
		"URL":   true,
		"UTF8":  true,
		"VM":    true,
		"XML":   true,
		"XSRF":  true,
		"XSS":   true,
	}
)
