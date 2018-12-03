/**
 * Go Video Downloader
 *
 *    Copyright 2018 Tenta, LLC
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
 * re.go: Re-implementation of Python's regexp functions (re)
 */

package re

import (
	"fmt"
	"strconv"
	"strings"

	rnt "github.com/tenta-browser/go-video-downloader/runtime"

	"github.com/tenta-browser/go-pcre-matcher"
	"github.com/tenta-browser/go-pcre-matcher/replacer"
	"github.com/tenta-browser/go-video-downloader/lib"
)

// Python regexp flags
var (
	FlagIgnoreCase = rnt.NewInt(2)
	FlagMultiline  = rnt.NewInt(8)
	FlagDotAll     = rnt.NewInt(16)
	FlagUnicode    = rnt.NewInt(32)
	FlagVerbose    = rnt.NewInt(64)
)

// ErrorType is the error type raised by the regexp functions ..
var ErrorType = rnt.Cal(
	rnt.TypeType,
	rnt.NewStr("re.error"),
	rnt.NewTuple(rnt.ExceptionType),
	rnt.NewDict()).(rnt.Type)

// translatePyFlags converts a Python regexp flagset to our engine's flagset
func translatePyFlags(pyFlags int) (int, error) {
	flags := 0
	var flagMappers = map[rnt.Int](func() int){
		FlagIgnoreCase: matcher.ReEngine.FlagCaseInsensitive,
		FlagMultiline:  matcher.ReEngine.FlagMultiline,
		FlagDotAll:     matcher.ReEngine.FlagDotAll,
		FlagUnicode:    matcher.ReEngine.FlagUnicode,
		FlagVerbose:    matcher.ReEngine.FlagExtended,
	}
	for pyFlagI, flagF := range flagMappers {
		pyFlag := pyFlagI.Value()
		if pyFlags&pyFlag != 0 {
			flags |= flagF()
			pyFlags &^= pyFlag // turn off translated flag
		}
	}
	if pyFlags > 0 { // err on non-translated flags
		return 0, fmt.Errorf("Unknown flags: %d", pyFlags)
	}
	return flags, nil
}

// translatePyPattern removes Python-specific stuff from the pattern
// - escape {...} when it's not a valid quantifier
// - convert {,...} qualifiers to {0,...} since PCRE does not like the former
// - convert (?P=group) -> \k<group>
// - convert (?P<group>pattern) -> (?<group>pattern)
// - escape # in character classes, so that it doesn't get treated as comment in verbose mode
func translatePyPattern(pattern string) string {
	ln := len(pattern)
	buf := make([]byte, 0, ln)
	inCharClass := false
	for i := 0; i < ln; i++ {
		if pattern[i] == '\\' {
			i++
			if i == ln {
				panic(rnt.RaiseType(ErrorType, "Pattern ends with backslash"))
			} else {
				buf = append(buf, '\\', pattern[i])
			}
			continue
		} else if pattern[i] == '[' {
			inCharClass = true
		} else if pattern[i] == ']' {
			inCharClass = false
		} else if pattern[i] == '#' && inCharClass {
			buf = append(buf, "\\#"...)
			continue
		} else if !inCharClass {
			if pattern[i] == '{' {
				si := i
				i++
				buf2 := make([]byte, 0, 4)
				for ; i < ln && pattern[i] >= '0' && pattern[i] <= '9'; i++ {
					buf2 = append(buf2, pattern[i])
				}
				if i < ln && pattern[i] == ',' {
					buf2 = append(buf2, pattern[i])
					i++
				}
				for ; i < ln && pattern[i] >= '0' && pattern[i] <= '9'; i++ {
					buf2 = append(buf2, pattern[i])
				}
				if len(buf2) > 0 && i < ln && pattern[i] == '}' {
					buf = append(buf, '{')
					if buf2[0] == ',' {
						buf = append(buf, '0')
					}
					buf = append(buf, buf2...)
					buf = append(buf, '}')
				} else {
					i = si
					buf = append(buf, "\\{"...)
				}
				continue
			} else if pattern[i] == '}' {
				buf = append(buf, "\\}"...)
				continue
			} else if strings.HasPrefix(pattern[i:], "(?P<") {
				buf = append(buf, "(?<"...)
				i += len("(?P<") - 1
				continue
			} else if strings.HasPrefix(pattern[i:], "(?P=") {
				buf = append(buf, "\\k<"...)
				i += len("(?P=")
				for ; i < ln && pattern[i] != ')'; i++ {
					buf = append(buf, pattern[i])
				}
				buf = append(buf, '>')
				continue
			} else if strings.HasPrefix(pattern[i:], "(?") {
				buf2 := []byte("(?")
				handled := false
				for j := i + len(buf2); !handled && j < ln; j++ {
					switch ch := pattern[j]; ch {
					case 'a', 'i', 'L', 'm', 's', 'x':
						buf2 = append(buf2, ch)
					case 'u':
						// TODO we skip this flag for now, because PCRE does not support it
					case ')':
						buf2 = append(buf2, ch)
						if len(buf2) > 3 { // suppress (?)
							buf = append(buf, buf2...)
						}
						i = j
						handled = true
					default:
						// this was not a flag setting construct, break out
						j = ln
					}
				}
				if handled {
					continue
				}
			}
		}
		buf = append(buf, pattern[i])
	}
	return string(buf)
}

// translatePyReplacement converts Python-specific parts of the replacement string
func translatePyReplacement(repl string) string {
	ln := len(repl)
	buf := make([]byte, 0, ln)
	for i := 0; i < ln; i++ {
		switch repl[i] {
		case '\\':
			i++
			if i == ln {
				panic(rnt.RaiseType(ErrorType,
					"Replacement string ends with backslash"))
			} else if repl[i] == '\\' {
				buf = append(buf, '\\')
			} else if repl[i] >= '0' && repl[i] <= '9' {
				buf = append(buf, '$')
				i--
			} else if repl[i] == 'g' {
				// transform \g<group> -> ${groupName} or $groupIdx
				i++
				grp := []byte{}
				for ; i < ln; i++ {
					grp = append(grp, repl[i])
					if repl[i] == '>' {
						break
					}
				}
				if l := len(grp); l <= 2 || grp[0] != '<' || grp[l-1] != '>' {
					panic(rnt.RaiseType(ErrorType,
						fmt.Sprintf("Invalid group name: '%s'", grp)))
				} else {
					grp = grp[1 : l-1]
				}
				if _, err := strconv.Atoi(string(grp)); err == nil {
					buf = append(buf, '$')
					buf = append(buf, grp...)
				} else {
					buf = append(buf, "${"...)
					buf = append(buf, grp...)
					buf = append(buf, '}')
				}
			} else {
				panic(rnt.RaiseType(ErrorType,
					fmt.Sprintf("Unknown char escaped: '%c'", repl[i])))
			}
		case '$':
			buf = append(buf, "\\$"...)
		default:
			buf = append(buf, repl[i])
		}
	}
	return string(buf)
}

func groupNames(pattern string) []string {
	groups := []string{}
	ln := len(pattern)
	inCharClass := false
	for i := 0; i < ln; i++ {
		if pattern[i] == '\\' {
			i++
			if i == ln {
				panic(rnt.RaiseType(ErrorType, "Pattern ends with backslash"))
			}
		} else if pattern[i] == '[' {
			inCharClass = true
		} else if pattern[i] == ']' {
			inCharClass = false
		} else if !inCharClass {
			if strings.HasPrefix(pattern[i:], "(?P<") {
				i += len("(?P<")
				j := i
				for ; j < ln && pattern[j] != '>'; j++ {
				}
				if j == ln {
					panic(rnt.RaiseType(ErrorType, "Unterminated named group"))
				}
				groups = append(groups, string(pattern[i:j]))
				i = j + 1
			}
		}
	}
	return groups
}

// CompileInternal prepares and compiles a pattern (unwrapped)
func CompileInternal(pattern string, flags int) (matcher.Regexp, error) {
	flags, err := translatePyFlags(flags)
	if err != nil {
		return nil, err
	}
	pattern = translatePyPattern(pattern)
	re, err := matcher.ReEngine.Compile(pattern, flags)
	if err != nil {
		return nil, err
	}
	return re, nil
}

// Compile prepares and compiles a pattern
var Compile = rnt.NewSimpleFunction("compile", []string{"pattern", "flags"},
	func(args []rnt.Object) rnt.Object {
		pattern, _ := lib.StringlikeVal(args[0])
		flags := args[1].(rnt.Int).Value()
		flags, err := translatePyFlags(flags)
		if err != nil {
			panic(rnt.RaiseType(ErrorType, err.Error()))
		}
		pattern = translatePyPattern(pattern)
		re, err := matcher.ReEngine.Compile(pattern, flags)
		if err != nil {
			panic(rnt.RaiseType(ErrorType, err.Error()))
		}
		return rnt.NewNative(re)
	})

// GroupNames returns the names of named groups in the pattern
var GroupNames = rnt.NewSimpleFunction("groupNames", []string{"pattern"},
	func(args []rnt.Object) rnt.Object {
		pattern, _ := lib.StringlikeVal(args[0])
		names := groupNames(pattern)
		ret := make([]rnt.Object, len(names))
		for i, name := range names {
			ret[i] = rnt.NewStr(name)
		}
		return rnt.NewList(ret...)
	})

// Search returns the first match.
// When anchorStart is true it anchors it to the beginning of str.
var Search = rnt.NewFunction("search",
	[]rnt.Param{
		{Name: "regexp"},
		{Name: "str"},
		{Name: "anchorStart", Def: rnt.False},
	}, 0, false, false,
	func(args []rnt.Object) rnt.Object {
		regexp := rnt.UnwrapNative(args[0]).(matcher.Regexp)
		str, _ := lib.StringlikeVal(args[1])
		anchs := rnt.IsTrue(args[2])
		match := regexp.Search(str)
		if match == nil {
			return rnt.None
		}
		if anchs && !strings.HasPrefix(str, match.GroupByIdx(0)) {
			// TODO this is a hack, because we cannot currently anchor compiled regexps
			return rnt.None
		}
		return rnt.NewNative(match)
	})

// Group returns the value of the group identified by idx (str or int) from match.
// If the group did not participate in the match it returns def.
var Group = rnt.NewSimpleFunction("group", []string{"match", "st", "idx", "def"},
	func(args []rnt.Object) rnt.Object {
		match := rnt.UnwrapNative(args[0]).(matcher.Match)
		st := rnt.IsTrue(args[1])
		idx := args[2]
		def := args[3]
		var val string
		if idx.IsInstance(rnt.IntType) {
			idx := idx.(rnt.Int).Value()
			if !match.GroupPresentByIdx(idx) {
				return def
			}
			val = match.GroupByIdx(idx)
		} else if idx.IsInstance(rnt.StrType) {
			idx := idx.(rnt.Str).Value()
			if !match.GroupPresentByName(idx) {
				return def
			}
			val = match.GroupByName(idx)
		} else {
			panic(rnt.RaiseType(rnt.TypeErrorType,
				fmt.Sprintf("expected index to be int or str, not %s", idx.Type().Name())))
		}
		return lib.NewStringlike(val, st)
	})

// Groups returns the number of groups in the match.
var Groups = rnt.NewSimpleFunction("groups", []string{"match"},
	func(args []rnt.Object) rnt.Object {
		match := rnt.UnwrapNative(args[0]).(matcher.Match)
		return rnt.NewInt(match.Groups())
	})

// Next looks for the next match and returns true if it's found.
var Next = rnt.NewSimpleFunction("next", []string{"match"},
	func(args []rnt.Object) rnt.Object {
		match := rnt.UnwrapNative(args[0]).(matcher.Match)
		return rnt.NewBool(match.Next())
	})

// Replace replaces every occurrence of regexp in subject with repl
// and returns the new string.
var Replace = rnt.NewSimpleFunction("replace", []string{"regexp", "repl", "subject"},
	func(args []rnt.Object) rnt.Object {
		regexp := rnt.UnwrapNative(args[0]).(matcher.Regexp)
		repl, st := lib.StringlikeVal(args[1])
		subject, _ := lib.StringlikeVal(args[2])
		return lib.NewStringlike(regexp.Replace(subject, translatePyReplacement(repl)), st)
	})

// ReplaceFunc replaces every occurrence of regexp in subject with
// the result returned by repl invoked for each match in turn,
// and returns the new string.
var ReplaceFunc = rnt.NewSimpleFunction("replaceFunc",
	[]string{"regexp", "repl", "subject"},
	func(args []rnt.Object) rnt.Object {
		regexp := rnt.UnwrapNative(args[0]).(matcher.Regexp)
		repl := args[1]
		subject, st := lib.StringlikeVal(args[2])
		var replErr interface{}
		repler := replacer.NewReplacer(func(m matcher.Match) string {
			defer func() {
				// Panicking into the regexp engine crashes it,
				// so recover it here, and re-panic it after the regexp
				// engine returns.
				replErr = recover()
			}()
			mrepl, _ := lib.StringlikeVal(rnt.Cal(repl, rnt.NewNative(m)))
			return mrepl
		})
		res := lib.NewStringlike(regexp.ReplaceFunc(subject, repler), st)
		if replErr != nil {
			panic(replErr)
		}
		return res
	})

// Escape escapes special characters in pattern.
var Escape = rnt.NewSimpleFunction("escape", []string{"pattern"},
	func(args []rnt.Object) rnt.Object {
		pattern, st := lib.StringlikeVal(args[0])
		return lib.NewStringlike(matcher.ReEngine.Quote(pattern), st)
	})
