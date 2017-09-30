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
 * builtin-re.go: Re-implementation of Python's regexp functions (re)
 */

package runtime

import (
	"fmt"
	"strings"

	"github.com/tenta-browser/go-pcre-matcher"
)

// Re is dummy
const Re = 0

// Python regexp flags
const (
	ReFlagDotAll = 16
)

// translatePyFlags converts a Python regexp flagset to our engine's flagset
func translatePyFlags(pyFlags int) (int, error) {
	flags := 0
	var flagMappers = map[int](func() int){
		ReFlagDotAll: matcher.ReEngine.FlagDotAll,
	}
	for pyFlag, flagF := range flagMappers {
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
func translatePyPattern(pattern string) string {
	pattern = strings.Replace(pattern, "(?P<", "(?<", -1) // named group syntax: (?<name>pattern)
	return pattern
}

// translatePyReplacement converts Python-specific parts of the replacement string
func translatePyReplacement(repl string) string {
	ln := len(repl)
	buf := make([]byte, 0, ln)
	for i := 0; i < ln; i++ {
		if repl[i] == '\\' {
			i++
			if i == ln {
				panic(newExtractorError("Replacement string ends with backslash"))
			} else if repl[i] == '\\' {
				buf = append(buf, '\\')
			} else if repl[i] >= '0' && repl[i] <= '9' {
				buf = append(buf, '$')
				i--
			} else {
				panic(newExtractorError(fmt.Sprintf("Unknown char escaped: '%c'", repl[i])))
			}
		} else if repl[i] == '$' {
			buf = append(buf, "\\$"...)
		} else {
			buf = append(buf, repl[i])
		}
	}
	return string(buf)
}

func reMustCompile(pattern string, flags int) matcher.Regexp {
	flags, err := translatePyFlags(flags)
	if err != nil {
		panic(newExtractorError(err.Error()))
	}
	pattern = translatePyPattern(pattern)
	re, err := matcher.ReEngine.Compile(pattern, flags)
	if err != nil {
		panic(newExtractorError(err.Error()))
	}
	return re
}

// ReSearch implements python/re.search
func ReSearch(re int, pattern, str string, flags int) matcher.Match {
	return reMustCompile(pattern, flags).Search(str)
}

// ReMatch implements python/re.match
func ReMatch(re int, pattern, str string, flags int) matcher.Match {
	return ReSearch(re, fmt.Sprintf("^(?:%s)", pattern), str, flags)
}

// ReFindAllOne implements python/re.findall returning a list matches (groupcount <= 1)
func ReFindAllOne(re int, pattern, str string, flags int) []string {
	match := ReSearch(re, pattern, str, flags)
	res := []string{}
	if match != nil {
		groupCount := match.Groups()
		for {
			if groupCount <= 1 {
				res = append(res, match.GroupByIdx(groupCount))
			} else {
				panic(newExtractorError("Multiple groups found when zero or one were expected"))
			}
			if !match.Next() {
				break
			}
		}
	}
	return res
}

// ReFindAllMulti implements python/re.findall returning a list of group-tuples (groupcount >= 2)
func ReFindAllMulti(re int, pattern, str string, flags int) [][]string {
	match := ReSearch(re, pattern, str, flags)
	res := [][]string{}
	if match != nil {
		groupCount := match.Groups()
		for {
			if groupCount <= 1 {
				panic(newExtractorError("Zero or one group found when multiple were expected"))
			} else {
				groups := []string{}
				for i := 1; i <= groupCount; i++ {
					groups = append(groups, match.GroupByIdx(i))
				}
				res = append(res, groups)
			}
			if !match.Next() {
				break
			}
		}
	}
	return res
}

// ReSub implements python/re.sub
func ReSub(re int, pattern, repl, subject string, count, flags int) string {
	if count != 0 {
		panic(newExtractorError("Non-zero count at ReSub")) // TODO
	}
	return reMustCompile(pattern, flags).Replace(subject, translatePyReplacement(repl))
}

// ReMatchGroupNone implements python/match.group with 0 args
func ReMatchGroupNone(match matcher.Match) OptString {
	return ReMatchGroupOne(match, 0)
}

// ReMatchGroupOne implements python/match.group with 1 arg
func ReMatchGroupOne(match matcher.Match, group interface{}) OptString {
	switch grp := group.(type) {
	case int:
		if !match.GroupPresentByIdx(grp) {
			return OptString{}
		}
		return AsOptString(match.GroupByIdx(grp))
	case string:
		if !match.GroupPresentByName(grp) {
			return OptString{}
		}
		return AsOptString(match.GroupByName(grp))
	default:
		panic(fmt.Sprintf("Group has invalid type: %T", grp))
	}
}

// ReMatchGroupMulti implements python/match.group with >=2 args
func ReMatchGroupMulti(match matcher.Match, groups []interface{}) []OptString {
	values := make([]OptString, len(groups))
	for idx, group := range groups {
		values[idx] = ReMatchGroupOne(match, group)
	}
	return values
}

// ReMatchGroups implements python/match.groups
func ReMatchGroups(match matcher.Match, def OptString) []OptString {
	groups := make([]OptString, match.Groups())
	for idx := 1; idx <= len(groups); idx++ {
		if match.GroupPresentByIdx(idx) {
			groups[idx-1] = AsOptString(match.GroupByIdx(idx))
		} else {
			groups[idx-1] = def
		}
	}
	return groups
}
