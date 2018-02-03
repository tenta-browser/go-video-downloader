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
	"strconv"
	"strings"
	"unicode"

	"github.com/tenta-browser/go-video-downloader/utils"
)

// StrSlice implements Python slicing for strings
func StrSlice(str string, low, high, step OptInt) string {
	rstr := []rune(str)
	rres := []rune{}
	lo, up, st := InitSliceArgs(len(str), low, high, step)
	if st > 0 {
		for i := lo; i < up; i += st {
			rres = append(rres, rstr[i])
		}
	} else {
		for i := up; i > lo; i += st {
			rres = append(rres, rstr[i])
		}
	}
	return string(rres)
}

// StrFormat2 implements Python 2's formatting (% operator)
func StrFormat2(format string, args ...interface{}) string {
	nargs := unboxStrFormatArgs(args)

	// make sure we don't pass funky formatting specs blindly to fmt.Sprintf
	// because it'll produce strange output; it's better to panic & correct
	argIdx := 0
	for _, fmtspec := range ReFindAllOne("%.*?[a-zA-Z%]", format, 0) {
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

// StrFormat2Named implements Python 2's formatting (% operator) with named parameters
func StrFormat2Named(format string, argsd SDict) string {
	args := []interface{}{}

	// match %% and %(name) tokens
	pattern := `%(?:%|\([^)]+\))`

	repl := func(m Match) string {
		s := m.GroupByIdx(0)
		if s == "%%" {
			return s
		}
		name := s[2 : len(s)-1] // s: %(name)
		arg, found := argsd[name]
		if !found {
			panic(newExtractorError("No value supplied for named parameter: " + name))
		}
		args = append(args, arg)
		return "%"
	}

	// convert named parameters to positional
	format = ReSubFunc(pattern, repl, format, 0, 0)
	return StrFormat2(format, args...)
}

// StrFormat3 implements python/str.format (Python 3 style formatting)
func StrFormat3(format string, args ...interface{}) string {
	nargs := unboxStrFormatArgs(args)

	// make sure we don't pass funky formatting specs blindly to fmt.Sprintf
	// because it'll produce strange output; it's better to panic & correct
	ln := len(format)
	buf := make([]byte, 0, ln)
	for i := 0; i < ln; i++ {
		if format[i] == '{' {
			i++
			if i == ln || format[i] != '{' {
				fmtspec := make([]byte, 0, 2)
				for ; i < ln && format[i] != '}'; i++ {
					fmtspec = append(fmtspec, format[i])
				}
				if i == ln {
					panic("Unterminated format spec")
				}
				if len(fmtspec) == 0 {
					buf = append(buf, "%v"...)
				} else if fmtidx, err := strconv.Atoi(string(fmtspec)); err == nil {
					buf = append(buf, "%["...)
					buf = append(buf, strconv.Itoa(fmtidx+1)...) // Python indexes from 0, Go from 1
					buf = append(buf, "]v"...)
				} else {
					panic("Unsupported formatter: " + string(fmtspec))
				}
				continue
			}
		} else if format[i] == '}' {
			i++
			if i == ln || format[i] != '}' {
				panic("Unstarted format spec")
			}
		}
		buf = append(buf, format[i])
	}

	format = string(buf)
	return fmt.Sprintf(format, nargs...)
}

func unboxStrFormatArgs(args []interface{}) []interface{} {
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
	return nargs
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
	if encoding == "ascii" {
		for _, r := range str {
			if r >= 128 {
				panic(newExtractorError(fmt.Sprintf("Invalid ascii char: %d", r)))
			}
		}
		return []byte(str)
	} else if encoding == "utf-8" {
		return []byte(str)
	}
	panic(newExtractorError("Unknown encoding: " + encoding))
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

// StrLower implements python/str.lower
func StrLower(str string) string {
	return strings.ToLower(str)
}

// StrUpper implements python/str.upper
func StrUpper(str string) string {
	return strings.ToUpper(str)
}

// StrStartswith implements python/str.startswith
func StrStartswith(str, prefix string, start, end OptInt) bool {
	return strings.HasPrefix(StrSlice(str, start, end, OptInt{}), prefix)
}

// StrEndswith implements python/str.endswith
func StrEndswith(str, suffix string, start, end OptInt) bool {
	return strings.HasSuffix(StrSlice(str, start, end, OptInt{}), suffix)
}

// StrJoin implements python/str.join
func StrJoin(str string, parts []string) string {
	return strings.Join(parts, str)
}

// BytesToStr implements python/bytes.decode
func BytesToStr(bytes []byte, encoding string) string {
	if encoding == "ascii" || encoding == "utf-8" {
		return string(bytes)
	}
	panic(newExtractorError("Unknown encoding: " + encoding))
}
