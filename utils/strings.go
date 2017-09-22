/**
 * Go Video Downloader
 *
 *    Copyright 2017 Tenta, LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * For any questions, please contact developer@tenta.io
 *
 * strings.go: String/text handling utilities
 */

package utils

// IndexRune returns the index of the first appearance of r in s.
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

// LastIndexRune returns the index of the last appearance of r in s.
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

// Ellipsize shortens s to maxlen length including ellipsis as a prefix.
func Ellipsize(s string, maxlen int, ellipsis string) string {
	l, el := len(s), len(ellipsis)
	if l <= maxlen {
		return s
	}
	return s[:(maxlen-el)] + ellipsis
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
