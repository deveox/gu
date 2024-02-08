package stringer

import (
	"strings"
	"unicode"
	"unsafe"
)

// ToCamelCase converts a title, snake or sentence case to a camel case string.
//
// Example: MyBigName -> myBigName
func ToCamelCase(s string) string {
	// complete := false
	n := strings.Builder{}
	for idx := 0; idx < len(s); idx++ {
		if idx == 0 {
			n.WriteRune(unicode.ToLower(rune(s[idx])))
			continue
		}
		switch s[idx] {
		case ' ', '_':
			if idx < len(s)-1 {
				switch s[idx+1] {
				case ' ', '_':
					continue
				default:
					n.WriteRune(unicode.ToUpper(rune(s[idx+1])))
					idx++
				}
			}
			continue
		default:
			if unicode.IsUpper(rune(s[idx])) && unicode.IsUpper(rune(s[idx-1])) && (idx+1 == len(s) || unicode.IsUpper(rune(s[idx+1]))) {
				n.WriteRune(unicode.ToLower(rune(s[idx])))
				continue
			}
			n.WriteByte(s[idx])
		}
	}
	return n.String()
}

// ToSnakeCase converts a title, camel or sentence case a snake case string.
func ToSnakeCase(s string) string {
	n := strings.Builder{}
	for idx, v := range s {
		switch v {
		case ' ':
			if idx != 0 && s[idx-1] != ' ' {
				n.WriteRune('_')
			}
			continue
		default:
			if unicode.IsUpper(v) {
				switch idx {
				case 0, len(s) - 1:
				default:
					switch s[idx-1] {
					case ' ', '_':
					default:
						if unicode.IsLower(rune(s[idx+1])) || unicode.IsLower(rune(s[idx-1])) {
							n.WriteRune('_')
						}
					}
				}
				v = unicode.ToLower(v)
			}
		}
		n.WriteRune(v)
	}
	return n.String()
}

func FromBytes(b []byte) string {
	// Ignore if your IDE shows an error here; it's a false positive.
	p := unsafe.SliceData(b)
	return unsafe.String(p, len(b))
}
