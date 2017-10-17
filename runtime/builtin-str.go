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
 * builtin-str.go: Re-implementation of various Python builtin string functions
 */

package runtime

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"

	"github.com/tenta-browser/go-video-downloader/utils"
)

// StrFormat implements Python 2's formatting (% operator)
func StrFormat(format string, args ...interface{}) string {
	// prepare args for formatting
	nargs := make([]interface{}, len(args))
	for argIdx, arg := range args {
		switch argVal := arg.(type) {
		case OptString:
			nargs[argIdx] = argVal.GetOrDef("")
		case OptInt:
			nargs[argIdx] = argVal.GetOrDef(0)
		default:
			nargs[argIdx] = arg
		}
	}

	// make sure we don't pass funky formatting specs blindly to fmt.Sprintf
	// because it'll produce strange output; it's better to panic & correct
	re := regexp.MustCompile("%.*?[a-zA-Z%]")
	argIdx := 0
	for _, fmtspec := range re.FindAllString(format, -1) {
		switch fmtspec {
		case "%%":
		case "%s":
			// Python is much more lax than Go when pairing format types to arg types
			// so we must make sure that only strings are passed for %s
			nargs[argIdx] = fmt.Sprintf("%v", nargs[argIdx])
			argIdx++
		case "%d":
			argIdx++
		default:
			panic("Unknown formatter spec: " + fmtspec)
		}
	}

	return fmt.Sprintf(format, nargs...)
}

// StrStrip implements python/str.strip
func StrStrip(str, chars string) string {
	if chars == "" {
		return strings.TrimSpace(str)
	}
	return strings.Trim(str, chars)
}

// StrLStrip implements python/str.lstrip
func StrLStrip(str, chars string) string {
	if chars == "" {
		return strings.TrimLeftFunc(str, unicode.IsSpace)
	}
	return strings.TrimLeft(str, chars)
}

// StrRStrip implements python/str.rstrip
func StrRStrip(str, chars string) string {
	if chars == "" {
		return strings.TrimRightFunc(str, unicode.IsSpace)
	}
	return strings.TrimRight(str, chars)
}

// StrSplit implements python/str.split
func StrSplit(str, sep string, max int) []string {
	max = max + 1
	if max == 0 {
		max = -1
	}
	if sep == "" {
		if max >= 0 {
			panic(newExtractorError("Unsupported limited word splitting"))
		}
		return strings.Fields(str)
	}
	return strings.SplitN(str, sep, max)
}

// StrRSplit implements python/str.rsplit
func StrRSplit(str, sep string, max int) []string {
	// TODO someday implement this properly
	res := StrSplit(utils.Reverse(str), utils.Reverse(sep), max)
	utils.ReverseStrSlice(res)
	for i, s := range res {
		res[i] = utils.Reverse(s)
	}
	return res
}

// StrContains implements python/(needle in haystack)
func StrContains(haystack, needle string) bool {
	return strings.Contains(haystack, needle)
}

// StrToBytes implements python/str.encode
func StrToBytes(str, encoding string) []byte {
	if encoding != "utf-8" {
		panic(newExtractorError("Unknown encoding: " + encoding))
	}
	return []byte(str)
}

// StrIsDigit implements python/str.isdigit
func StrIsDigit(str string) bool {
	if str == "" {
		return false
	}
	for _, r := range str {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

// StrFind implements python/str.find
func StrFind(str, sub string) int {
	return strings.Index(str, sub)
}

// StrReplace implements python/str.replace
func StrReplace(str, old, new string, count int) string {
	return strings.Replace(str, old, new, count)
}

// BytesToStr implements python/bytes.decode
func BytesToStr(bytes []byte, encoding string) string {
	if encoding != "utf-8" {
		panic(newExtractorError("Unknown encoding: " + encoding))
	}
	return string(bytes)
}
