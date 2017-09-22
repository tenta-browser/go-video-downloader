package utils

// IndexRune returns the index of the first appearence of r in s.
func IndexRune(s string, r rune) int {
	sr := ([]rune)(s)
	l := len(sr)
	for i := 0; i < l; i++ {
		if sr[i] == r {
			return i
		}
	}
	return -1
}

// LastIndexRune returns the index of the last appearence of r in s.
func LastIndexRune(s string, r rune) int {
	sr := ([]rune)(s)
	l := len(sr)
	for i := l - 1; i >= 0; i-- {
		if sr[i] == r {
			return i
		}
	}
	return -1
}

// Ellipsize shortens s to maxl length including ellipsis as a prefix.
func Ellipsize(s string, maxl int, ellipsis string) string {
	l, el := len(s), len(ellipsis)
	if l <= maxl {
		return s
	}
	return s[:(maxl-el)] + ellipsis
}

// Reverse reverses runewise a string.
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// ReverseStrSlice reverses a string slice (duh :))
func ReverseStrSlice(l []string) {
	ln := len(l)
	for i := ln/2 - 1; i >= 0; i-- {
		ni := ln - 1 - i
		l[i], l[ni] = l[ni], l[i]
	}
}
