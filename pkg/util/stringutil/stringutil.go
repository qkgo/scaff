package stringutil

import "strings"

func ContainsFold(a, b string) bool {
	return strings.Contains(strings.ToLower(a), strings.ToLower(b))
}

func EqualFold(a, b string) bool {
	return strings.EqualFold(a, b)
}

func CompareFold(a, b string) bool {
	return strings.EqualFold(a, b)
}

func SuffixFold(a, b string) bool {
	return strings.HasSuffix(strings.ToLower(a), strings.ToLower(b))
}

func PrefixFold(a, b string) bool {
	return strings.HasPrefix(strings.ToLower(a), strings.ToLower(b))
}
