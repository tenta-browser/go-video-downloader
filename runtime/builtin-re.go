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
	"github.com/tenta-browser/go-pcre-matcher/replacer"
)

// Python regexp flags
const (
	ReFlagDotAll = 16
)

type (
	// Regexp is the regexp type exposed to extractors
	Regexp = matcher.Regexp
	// Match is the match type exposed to extractors
	Match = matcher.Match
)

// MatchIterator is an iterator over all matches of a search
type MatchIterator interface {
	Next() (Match, bool)
}

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
				panic(newExtractorError("Pattern ends with backslash"))
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
				panic(newExtractorError("Replacement string ends with backslash"))
			} else if repl[i] == '\\' {
				buf = append(buf, '\\')
			} else if repl[i] >= '0' && repl[i] <= '9' {
				buf = append(buf, '$')
				i--
			} else {
				panic(newExtractorError(fmt.Sprintf("Unknown char escaped: '%c'", repl[i])))
			}
		case '$':
			buf = append(buf, "\\$"...)
		default:
			buf = append(buf, repl[i])
		}
	}
	return string(buf)
}

// ReMustCompile prepares and compiles a pattern
func ReMustCompile(pattern string, flags int) Regexp {
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
func ReSearch(pattern, str string, flags int) Match {
	return ReMustCompile(pattern, flags).Search(str)
}

// ReMatch implements python/re.match
func ReMatch(pattern, str string, flags int) Match {
	return ReSearch(fmt.Sprintf("^(?:%s)", pattern), str, flags)
}

// ReFindAllOne implements python/re.findall returning a list matches (groupcount <= 1)
func ReFindAllOne(pattern, str string, flags int) []string {
	match := ReSearch(pattern, str, flags)
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
func ReFindAllMulti(pattern, str string, flags int) [][]string {
	match := ReSearch(pattern, str, flags)
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

type matchIteratorProto struct {
	match Match
	dirty bool
}

func (mip *matchIteratorProto) Next() (Match, bool) {
	if !mip.dirty {
		mip.dirty = true
	} else if mip.match != nil {
		// if this is not the first invocation, pull the next match
		if !mip.match.Next() {
			mip.match = nil
		}
	}
	return mip.match, mip.match != nil
}

// ReFindIter implements python/re.finditer
func ReFindIter(pattern, str string, flags int) MatchIterator {
	match := ReSearch(pattern, str, flags)
	return &matchIteratorProto{match: match}
}

// ReSub implements python/re.sub (for string repl)
func ReSub(pattern, repl, subject string, count, flags int) string {
	if count != 0 {
		panic(newExtractorError("Non-zero count at ReSub")) // TODO
	}
	return ReMustCompile(pattern, flags).Replace(subject, translatePyReplacement(repl))
}

// ReSubFunc implements python/re.sub (for func repl)
func ReSubFunc(pattern string, repl func(Match) string, subject string, count, flags int) string {
	if count != 0 {
		panic(newExtractorError("Non-zero count at ReSubFunc")) // TODO
	}
	return ReMustCompile(pattern, flags).ReplaceFunc(subject, replacer.NewReplacer(repl))
}

// ReMatchGroupNone implements python/match.group with 0 args
func ReMatchGroupNone(match Match) OptString {
	return ReMatchGroupOne(match, 0)
}

// ReMatchGroupOne implements python/match.group with 1 arg
func ReMatchGroupOne(match Match, group interface{}) OptString {
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

// ReMatchGroupTwoResult is a tuple returned by the ReMatchGroupTwo function
// consisting of the optional values of the requested two groups
type ReMatchGroupTwoResult = struct {
	Φ0 OptString
	Φ1 OptString
}

// ReMatchGroupTwo implements python/match.group with 2 args
func ReMatchGroupTwo(match Match, group1, group2 interface{}) ReMatchGroupTwoResult {
	values := ReMatchGroupMulti(match, []interface{}{group1, group2})
	return ReMatchGroupTwoResult{values[0], values[1]}
}

// ReMatchGroupMulti implements python/match.group with >=2 args
func ReMatchGroupMulti(match Match, groups []interface{}) []OptString {
	values := make([]OptString, len(groups))
	for idx, group := range groups {
		values[idx] = ReMatchGroupOne(match, group)
	}
	return values
}

// ReMatchGroups implements python/match.groups
func ReMatchGroups(match Match, def OptString) []OptString {
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
