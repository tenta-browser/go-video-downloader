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
	pattern = strings.Replace(pattern, "(?P<", "(?<", -1) // named group syntax: (?<name>patt)
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
func ReSearch(re int, pattern, str string, flags int) matcher.Matcher {
	return reMustCompile(pattern, flags).Search(str)
}

// ReMatch implements python/re.match
func ReMatch(re int, pattern, str string, flags int) matcher.Matcher {
	return ReSearch(re, fmt.Sprintf("^(?:%s)", pattern), str, flags)
}

// ReFindAllOne implements python/re.findall returning a list matches (groupcount <= 1)
func ReFindAllOne(re int, pattern, str string, flags int) []string {
	matcher := ReSearch(re, pattern, str, flags)
	res := []string{}
	if matcher != nil {
		groupCount := matcher.Groups()
		for {
			if groupCount <= 1 {
				res = append(res, matcher.GroupByIdx(groupCount))
			} else {
				panic(newExtractorError("Multiple groups found when zero or one were expected"))
			}
			if !matcher.Next() {
				break
			}
		}
	}
	return res
}

// ReFindAllMulti implements python/re.findall returning a list of group-tuples (groupcount >= 2)
func ReFindAllMulti(re int, pattern, str string, flags int) [][]string {
	matcher := ReSearch(re, pattern, str, flags)
	res := [][]string{}
	if matcher != nil {
		groupCount := matcher.Groups()
		for {
			if groupCount <= 1 {
				panic(newExtractorError("Zero or one group fround when multiple were expected"))
			} else {
				grps := []string{}
				for i := 1; i <= groupCount; i++ {
					grps = append(grps, matcher.GroupByIdx(i))
				}
				res = append(res, grps)
			}
			if !matcher.Next() {
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
func ReMatchGroupNone(matcher matcher.Matcher) OptString {
	return ReMatchGroupOne(matcher, 0)
}

// ReMatchGroupOne implements python/match.group with 1 arg
func ReMatchGroupOne(matcher matcher.Matcher, group interface{}) OptString {
	switch grp := group.(type) {
	case int:
		if !matcher.GroupPresentByIdx(grp) {
			return OptString{}
		}
		return AsOptString(matcher.GroupByIdx(grp))
	case string:
		if !matcher.GroupPresentByName(grp) {
			return OptString{}
		}
		return AsOptString(matcher.GroupByName(grp))
	default:
		panic(fmt.Sprintf("Group has invalid type: %T", grp))
	}
}

// ReMatchGroupMulti implements python/match.group with >=2 args
func ReMatchGroupMulti(matcher matcher.Matcher, groups []interface{}) []OptString {
	values := make([]OptString, len(groups))
	for idx, group := range groups {
		values[idx] = ReMatchGroupOne(matcher, group)
	}
	return values
}
