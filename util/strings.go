package util

import "unicode"

// CapitalizeFirst capitalizes the first letter of a string. This is safe to use with unicode characters
func CapitalizeFirst(s string) string {
	a := []rune(s)
	a[0] = unicode.ToUpper(a[0])
	return string(a)
}
