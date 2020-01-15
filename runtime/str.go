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
 * str.go: Implementation of Python's string type
 */

package runtime

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Str ..
type Str interface {
	Object
	Value() string
}

type strStruct struct {
	Object
	value string
}

func (ss *strStruct) Value() string {
	return ss.value
}

var _ Str = (*strStruct)(nil)

func newStr(t Type, value string) Str {
	if t == StrType {
		return StrLiteral(value)
	}
	return &strStruct{newObject(t), value}
}

// NewStr ..
func NewStr(value string) Str {
	return newStr(StrType, value)
}

// StrType ..
var StrType = newBuiltinType("str")

func strNew(t Type, args Args, kwArgs KWArgs) Object {
	if t != StrType {
		s := strNew(StrType, args, kwArgs)
		return newStr(t, s.(Str).Value())
	}
	argCount := len(args)
	if argCount == 0 {
		return NewStr("")
	}
	if argCount != 1 {
		panic(RaiseType(NotImplementedErrorType, "Unicode decoding not implemented"))
	}
	o := args[0]
	// try __str__()
	if strFn := o.Type().Slots().Str; strFn != nil {
		return strFn(o)
	}
	// fallback to repr()
	return Repr(o)
}

func strHash(o Object) Object {
	s := o.(Str)
	h := newHash()
	h.Write([]byte(s.Value()))
	return NewInt(sumHash(h))
}

var strReprReplacer = strings.NewReplacer(
	"\\", `\\`,
	"'", `\'`,
	"\n", `\n`,
	"\r", `\r`,
	"\t", `\t`,
)

func strRepr(o Object) Object {
	s := o.(Str)
	return NewStr(fmt.Sprintf("'%s'", strReprReplacer.Replace(s.Value())))
}

func strStr(o Object) Object {
	s := o.(Str)
	return NewStr(s.Value())
}

func strInt(o Object) Object {
	s := o.(Str).Value()
	s = strings.TrimSpace(s)
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(RaiseType(ValueErrorType, fmt.Sprintf(
			"invalid literal for int() with base 10: '%s'", s)))
	}
	return NewInt(i)
}

func strFloat(o Object) Object {
	s := o.(Str).Value()
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(RaiseType(ValueErrorType, fmt.Sprintf(
			"could not convert string to float: '%s'", s)))
	}
	return NewFloat(f)
}

// copy of seqGetItem, using runes instead of Objects
func strGetItem(o Object, index Object) Object {
	t := o.Type()
	runes := []rune(o.(Str).Value())
	if index.IsInstance(IntType) {
		i := seqCheckedIndex(t, len(runes), index.(Int).Value())
		return NewStr(string(runes[i]))
	} else if index.IsInstance(SliceType) {
		start, stop, step, len := index.(Slice).CalcSlice(len(runes))
		var sliced []rune
		if step == 1 {
			sliced = runes[start:stop]
		} else {
			sliced = make([]rune, len)
			i := 0
			for j := start; j != stop; j += step {
				sliced[i] = runes[j]
				i++
			}
		}
		return NewStr(string(sliced))
	} else {
		panic(RaiseType(TypeErrorType, fmt.Sprintf(
			"%s indices must be integers or slices, not %s",
			t.Name(), index.Type().Name())))
	}
}

func strLen(o Object) Object {
	s := o.(Str)
	return NewInt(len([]rune(s.Value())))
}

func strCompare(op compareOp, a, b Object) Object {
	if !b.IsInstance(StrType) {
		return NotImplemented
	}
	i := a.(Str).Value()
	j := b.(Str).Value()
	var r bool
	switch op {
	case compareEq:
		r = (i == j)
	case compareNe:
		r = (i != j)
	case compareLt:
		r = (i < j)
	case compareGt:
		r = (i > j)
	case compareLe:
		r = (i <= j)
	case compareGe:
		r = (i >= j)
	}
	return NewBool(r)
}

func strContains(o Object, needle Object) Object {
	if !needle.IsInstance(StrType) {
		panic(RaiseType(TypeErrorType, fmt.Sprintf(
			"'in <string>' requires string as left operand, not %s",
			needle.Type().Name())))
	}
	s := o.(Str).Value()
	ns := needle.(Str).Value()
	return NewBool(strings.Contains(s, ns))
}

func strAdd(a, b Object) Object {
	if !b.IsInstance(StrType) {
		return NotImplemented
	}
	as := a.(Str).Value()
	bs := b.(Str).Value()
	return NewStr(as + bs)
}

func strMul(o, c Object) Object {
	if !c.IsInstance(IntType) {
		return NotImplemented
	}
	s := o.(Str).Value()
	cnt := c.(Int).Value()
	return NewStr(strings.Repeat(s, cnt))
}

var strInterpolationRegexp = regexp.MustCompile(
	`^%(\(([^)]+)\))?([#0 +-]?)(\*|[0-9]+)?(\.(\*|[0-9]+))?[hlL]?([diouxXeEfFgGcrs%])`)

func strMod(o, args Object) Object {
	format := o.(Str).Value()
	// If the format string contains mappings, args must be a dict,
	// otherwise it must be treated as a tuple of values, so if it's not a tuple already
	// it must be transformed into a single element tuple.
	var values Tuple
	var mappings Dict
	if args.IsInstance(TupleType) {
		values = args.(Tuple)
	} else {
		values = NewTuple(args)
		if args.IsInstance(DictType) {
			mappings = args.(Dict)
		}
	}
	const (
		idxAll = iota
		// Todo: mapping keys can contain balanced parentheses, these should be matched
		// manually before using the regexp
		_
		idxMappingKey
		idxFlags
		idxWidth
		idxPrecision
		_
		idxType
	)
	var buf bytes.Buffer
	valueIndex := 0
	index := strings.Index(format, "%")
	for index != -1 {
		buf.WriteString(format[:index])
		format = format[index:]
		matches := strInterpolationRegexp.FindStringSubmatch(format)
		if matches == nil {
			panic(RaiseType(ValueErrorType, "invalid format spec"))
		}
		mappingKey, fieldType := matches[idxMappingKey], matches[idxType]
		var value Object
		if mappingKey != "" {
			// Nb: mappings are checked even in case of "%%"
			if mappings == nil {
				panic(RaiseType(TypeErrorType, "format requires a mapping"))
			}
			value = mappings.GetItem(NewStr(mappingKey))
			if value == nil {
				panic(RaiseType(KeyErrorType, fmt.Sprintf("'%s'", mappingKey)))
			}
			valueIndex = 1
		} else if fieldType != "%" {
			if valueIndex >= len(values.Elems()) {
				panic(RaiseType(TypeErrorType, "not enough arguments for format string"))
			}
			value = values.Elems()[valueIndex]
			valueIndex++
		}
		fieldWidth := -1
		if matches[idxWidth] == "*" || matches[idxPrecision] != "" {
			panic(RaiseType(NotImplementedErrorType, "field width not yet supported"))
		}
		if matches[idxWidth] != "" {
			var err error
			fieldWidth, err = strconv.Atoi(matches[idxWidth])
			if err != nil {
				panic(RaiseType(TypeErrorType, fmt.Sprint(err)))
			}
		}
		flags := matches[idxFlags]
		if flags != "" && flags != "0" {
			panic(RaiseType(NotImplementedErrorType, "conversion flags not yet supported"))
		}
		var val string
		switch fieldType {
		case "r", "s":
			var s Str
			if fieldType == "r" {
				s = Repr(value)
			} else {
				s = ToStr(value)
			}
			val = s.Value()
			if fieldWidth > 0 {
				val = strModPad(val, fieldWidth, ' ')
			}
			buf.WriteString(val)
		case "f":
			if v, ok := floatCoerce(value); ok {
				val := strconv.FormatFloat(v, 'f', 6, 64)
				if fieldWidth > 0 {
					fillchar := byte(' ')
					if flags != "" {
						fillchar = flags[0]
					}
					val = strModPad(val, fieldWidth, fillchar)
				}
				buf.WriteString(val)
			} else {
				panic(RaiseType(TypeErrorType,
					fmt.Sprintf("float argument required, not %s", value.Type().Name())))
			}
		case "d", "x", "X", "o":
			i := ToInt(value)
			var base int
			switch fieldType {
			case "d":
				base = 10
			case "x", "X":
				base = 16
			case "o":
				base = 8
			}
			val = strconv.FormatInt(int64(i.(Int).Value()), base)
			if fieldType == "X" {
				val = strings.ToUpper(val)
			}
			if fieldWidth > 0 {
				fillchar := byte(' ')
				if flags != "" {
					fillchar = flags[0]
				}
				val = strModPad(val, fieldWidth, fillchar)
			}
			buf.WriteString(val)
		case "%":
			val = "%"
			if fieldWidth > 0 {
				val = strModPad(val, fieldWidth, ' ')
			}
			buf.WriteString(val)
		default:
			panic(RaiseType(NotImplementedErrorType,
				fmt.Sprintf("conversion type not yet supported: %s", fieldType)))
		}
		format = format[len(matches[idxAll]):]
		index = strings.Index(format, "%")
	}
	if mappings == nil && valueIndex < len(values.Elems()) {
		panic(RaiseType(TypeErrorType,
			"not all arguments converted during string formatting"))
	}
	buf.WriteString(format)
	return NewStr(buf.String())
}

func strModPad(s string, width int, fillchar byte) string {
	l := len(s)
	if l > width {
		return s
	}
	buf := strings.Builder{}
	buf.Grow(width)
	// if the string starts with +/- and fillchar is '0' treat it like a number
	if l > 0 && fillchar == '0' && (s[0] == '+' || s[0] == '-') {
		buf.WriteByte(s[0])
		s = s[1:]
	}
	for i := 0; i < width-l; i++ {
		buf.WriteByte(fillchar)
	}
	buf.WriteString(s)
	return buf.String()
}

func strCapitalize(args Args, kwArgs KWArgs) Object {
	checkFunctionArgs("capitalize", args, kwArgs, StrType)
	s := args[0].(Str).Value()
	if s == "" {
		return args[0]
	}
	r, size := utf8.DecodeRuneInString(s)
	b := strings.Builder{}
	b.Grow(len(s))
	b.WriteRune(unicode.ToUpper(r))
	b.WriteString(strings.ToLower(s[size:]))
	return NewStr(b.String())
}

func strEncode(args Args, kwArgs KWArgs) Object {
	checkFunctionArgsMin("encode", args, kwArgs, 1, StrType, StrType, StrType)
	argc := len(args)
	encoding := "utf-8"
	if argc >= 2 {
		encoding = args[1].(Str).Value()
	}
	errors := NewStr(codingErrorsStrict)
	if argc >= 3 {
		errors = args[2].(Str)
	}
	bytes, _ := GetCodec(encoding).Encode(args[0], errors)
	return bytes
}

func strEndswith(args Args, kwArgs KWArgs) Object {
	return strStartsOrEndsWith("endswith", args, kwArgs, strings.HasSuffix)
}

func strFind(args Args, kwArgs KWArgs) Object {
	return strFindOrdIndex("find", args, kwArgs, func(s, sub string) int {
		return strings.Index(s, sub)
	})
}

func strFindOrdIndex(name string, args Args, kwArgs KWArgs,
	resFn func(string, string) int) Object {
	checkFunctionArgsMin(name, args, kwArgs, 2, StrType, StrType, IntType, IntType)
	s := args[0].(Str).Value()
	sub := args[1].(Str).Value()
	argc := len(args)
	l := len(s)
	start, end := 0, l
	if argc >= 3 {
		start = args[2].(Int).Value()
	}
	if argc >= 4 {
		end = args[3].(Int).Value()
	}
	ss, ssub := "", ":)"
	if start <= l {
		start = seqClampIndex(start, l)
		end = seqClampIndex(end, l)
		if start <= end {
			ss, ssub = s[start:end], sub
		}
	}
	index := resFn(ss, ssub)
	if index >= 0 {
		index += start
	}
	return NewInt(index)
}

var strFormatRegexp = regexp.MustCompile("^{" +
	"(|[0-9]+|[^}!:]+)" + // field name
	"(?:!(r|s))?" + // conversion
	"(?::([^}]+))?" + // format spec
	"}")

func strFormat(args Args, kwArgs KWArgs) Object {
	checkFunctionVarArgs("format", args, nil, StrType)
	format := args[0].(Str).Value()
	argst := NewTuple(args[1:]...)
	argsd := NewDict()
	for _, kwArg := range kwArgs {
		argsd.SetItem(NewStr(kwArg.Name), kwArg.Value)
	}
	const (
		idxAll = iota
		idxField
		idxConversion
		idxFormatSpec
	)
	argIdx := 0
	buf := strings.Builder{}
	ln := len(format)
	for i := 0; i < ln; {
		ch := format[i]
		if i+1 < ln && (ch == '{' || ch == '}') && ch == format[i+1] {
			i++
		} else if ch == '{' {
			matches := strFormatRegexp.FindStringSubmatch(format[i:])
			if matches == nil {
				panic(RaiseType(ValueErrorType, "invalid format string"))
			}
			i += len(matches[idxAll])
			field := matches[idxField]
			conversion := matches[idxConversion]
			if formatSpec := matches[idxFormatSpec]; formatSpec != "" {
				panic(RaiseType(NotImplementedErrorType, "format specs not str.format"))
			}
			var val Object
			if field == "" {
				val = GetItem(argst, NewInt(argIdx))
				argIdx++
			} else if num, err := strconv.Atoi(field); err == nil {
				val = GetItem(argst, NewInt(num))
			} else {
				val = GetItem(argsd, NewStr(field))
			}
			switch conversion {
			case "s", "":
				buf.WriteString(ToStr(val).Value())
			case "r":
				buf.WriteString(Repr(val).Value())
			default:
				panic(RaiseType(ValueErrorType, fmt.Sprintf(
					"Unknown conversion specifier %s", conversion)))
			}
			continue
		}
		buf.WriteByte(ch)
		i++
	}
	return NewStr(buf.String())
}

func strIndex(args Args, kwArgs KWArgs) Object {
	return strFindOrdIndex("index", args, kwArgs, func(s, sub string) int {
		index := strings.Index(s, sub)
		if index < 0 {
			panic(RaiseType(ValueErrorType, "substring not found"))
		}
		return index
	})
}

func strIs(name string, args Args, kwArgs KWArgs, emptyRes bool, isFn func(rune) bool) Object {
	checkFunctionArgs(name, args, kwArgs, StrType)
	s := args[0].(Str).Value()
	if s == "" {
		return NewBool(emptyRes)
	}
	for _, r := range s {
		if !isFn(r) {
			return NewBool(false)
		}
	}
	return NewBool(true)
}

func strIsAlNum(args Args, kwArgs KWArgs) Object {
	return strIs("isalnum", args, kwArgs, false,
		func(r rune) bool { return unicode.In(r, unicode.Letter, unicode.Number) })
}

func strIsAlpha(args Args, kwArgs KWArgs) Object {
	return strIs("isalpha", args, kwArgs, false, unicode.IsLetter)
}

func strIsASCII(args Args, kwArgs KWArgs) Object {
	return strIs("isascii", args, kwArgs, true, func(r rune) bool { return r <= unicode.MaxASCII })
}

func strIsDecimal(args Args, kwArgs KWArgs) Object {
	// Docs: Formally a decimal character is a character in the Unicode General Category "Nd".
	return strIs("isdecimal", args, kwArgs, false, unicode.IsDigit)
}

func strIsDigit(args Args, kwArgs KWArgs) Object {
	// Docs: Formally, a digit is a character that has the property value
	// Numeric_Type=Digit or Numeric_Type=Decimal.
	return strIs("isdigit", args, kwArgs, false, unicode.IsNumber)
}

func strIsLowerOrUpper(name string, args Args, kwArgs KWArgs, isLower bool) Object {
	checkFunctionArgs(name, args, kwArgs, StrType)
	s := args[0].(Str).Value()
	cased := false
	for _, r := range s {
		upper := unicode.IsUpper(r)
		lower := unicode.IsLower(r)
		if upper || lower {
			cased = true
			if lower != isLower {
				return NewBool(false)
			}
		}
	}
	return NewBool(cased)
}

func strIsLower(args Args, kwArgs KWArgs) Object {
	return strIsLowerOrUpper("islower", args, kwArgs, true)
}

func strIsNumeric(args Args, kwArgs KWArgs) Object {
	// Docs: Formally, numeric characters are those with the property value
	// Numeric_Type=Digit, Numeric_Type=Decimal or Numeric_Type=Numeric.
	// TODO this definition is too narrow, ideally we should have:
	//  一二三四五  isdecimal=False  isdigit=False  isnumeric=True
	//  12345      isdecimal=True   isdigit=True   isnumeric=True
	//  2²         isdecimal=False  isdigit=True   isnumeric=True
	return strIs("isnumeric", args, kwArgs, false, unicode.IsNumber)
}

func strIsPrintable(args Args, kwArgs KWArgs) Object {
	return strIs("isprintable", args, kwArgs, true, unicode.IsPrint)
}

func strIsSpace(args Args, kwArgs KWArgs) Object {
	return strIs("isspace", args, kwArgs, false, unicode.IsSpace)
}

func strIsUpper(args Args, kwArgs KWArgs) Object {
	return strIsLowerOrUpper("isupper", args, kwArgs, false)
}

func strJoin(args Args, kwArgs KWArgs) Object {
	checkFunctionArgs("join", args, kwArgs, StrType, ObjectType)
	sep := args[0].(Str).Value()
	i := 0
	b := strings.Builder{}
	seqForEach(args[1], func(item Object) {
		if !item.IsInstance(StrType) {
			panic(RaiseType(TypeErrorType, fmt.Sprintf(
				"sequence item %d: expected str instance, %s found",
				i, item.Type().Name())))
		}
		if i != 0 {
			b.WriteString(sep)
		}
		b.WriteString(item.(Str).Value())
		i++
	})
	return NewStr(b.String())
}

func strLower(args Args, kwArgs KWArgs) Object {
	checkFunctionArgs("lower", args, kwArgs, StrType)
	s := args[0].(Str)
	return NewStr(strings.ToLower(s.Value()))
}

func strPartition(args Args, kwArgs KWArgs) Object {
	checkFunctionArgs("partition", args, kwArgs, StrType, StrType)
	sep := args[1].(Str).Value()
	if sep == "" {
		panic(RaiseType(ValueErrorType, "empty separator"))
	}
	s := args[0].(Str).Value()
	pos := strings.Index(s, sep)
	if pos < 0 {
		emptyStr := NewStr("")
		return NewTuple(args[0], emptyStr, emptyStr)
	}
	start := NewStr(s[0:pos])
	end := NewStr(s[pos+len(sep):])
	return NewTuple(start, args[1], end)
}

func strReplace(args Args, kwArgs KWArgs) Object {
	checkFunctionArgsMin("replace", args, kwArgs, 3, StrType,
		StrType, StrType, IntType)
	s := args[0].(Str).Value()
	old := args[1].(Str).Value()
	new := args[2].(Str).Value()
	n := -1
	if len(args) >= 4 {
		n = args[3].(Int).Value()
	}
	if n == 0 {
		return args[0] // no subs
	}
	if old != "" {
		return NewStr(strings.Replace(s, old, new, n))
	}
	// when old is empty, new has to be inserted between characters
	b := strings.Builder{}
	i := 1
	b.WriteString(new)
	for _, r := range s {
		b.WriteRune(r)
		if n < 0 || i < n {
			b.WriteString(new)
			i++
		}
	}
	return NewStr(b.String())
}

func strRFind(args Args, kwArgs KWArgs) Object {
	return strFindOrdIndex("rfind", args, kwArgs, func(s, sub string) int {
		return strings.LastIndex(s, sub)
	})
}

func strRIndex(args Args, kwArgs KWArgs) Object {
	return strFindOrdIndex("rindex", args, kwArgs, func(s, sub string) int {
		index := strings.LastIndex(s, sub)
		if index < 0 {
			panic(RaiseType(ValueErrorType, "substring not found"))
		}
		return index
	})
}

func strRPartition(args Args, kwArgs KWArgs) Object {
	checkFunctionArgs("rpartition", args, kwArgs, StrType, StrType)
	sep := args[1].(Str).Value()
	if sep == "" {
		panic(RaiseType(ValueErrorType, "empty separator"))
	}
	s := args[0].(Str).Value()
	pos := strings.LastIndex(s, sep)
	if pos < 0 {
		emptyStr := NewStr("")
		return NewTuple(emptyStr, emptyStr, args[0])
	}
	start := NewStr(s[0:pos])
	end := NewStr(s[pos+len(sep):])
	return NewTuple(start, args[1], end)
}

var strSplitWhiteSpaceSep = regexp.MustCompile(`\s+`)

func strSplit(args Args, kwArgs KWArgs) Object {
	checkFunctionArgsMin("split", args, kwArgs, 1, StrType, ObjectType, IntType)
	argc := len(args)
	s := args[0].(Str).Value()
	sep := ""
	max := -1
	if argc >= 2 && args[1] != None {
		sep = args[1].(Str).Value()
		if sep == "" {
			panic(RaiseType(ValueErrorType, "empty separator"))
		}
	}
	if argc >= 3 {
		max = args[2].(Int).Value()
	}
	max++
	if max == 0 {
		max = -1
	}
	var parts []string
	if sep == "" {
		s = strings.TrimSpace(s)
		parts = strSplitWhiteSpaceSep.Split(s, max)
		if len(parts) == 1 && parts[0] == "" {
			parts = parts[:0]
		}
	} else {
		parts = strings.SplitN(s, sep, max)
	}
	elems := make([]Object, len(parts))
	for i, part := range parts {
		elems[i] = NewStr(part)
	}
	return NewList(elems...)
}

func strSplitlines(args Args, kwArgs KWArgs) Object {
	checkFunctionArgsMin("splitlines", args, kwArgs, 1, StrType, IntType)
	s := args[0].(Str).Value()
	keepEnds := false
	if len(args) >= 2 {
		keepEnds = IsTrue(args[1])
	}
	n := len(s)
	start, end := 0, 0
	lines := []Object{}
	for start < n {
		eol := 0
		for end = start; end < n; end++ {
			c := s[end]
			// TODO Include \u2028 and \u2029 too.
			if c == '\n' || c == '\v' || c == '\f' ||
				c == '\x1c' || c == '\x1d' || c == '\x1e' || c == '\x85' {
				eol = end + 1
				break
			}
			if c == '\r' {
				eol = end + 1
				if eol < n && s[eol] == '\n' {
					eol++
				}
				break
			}
		}
		if end >= n {
			eol = end
		}
		line := ""
		if keepEnds {
			line = s[start:eol]
		} else {
			line = s[start:end]
		}
		lines = append(lines, NewStr(line))
		start = eol
	}
	return NewList(lines...)
}

func strStartswith(args Args, kwArgs KWArgs) Object {
	return strStartsOrEndsWith("startswith", args, kwArgs, strings.HasPrefix)
}

func strStartsOrEndsWith(name string, args Args, kwArgs KWArgs,
	matchFn func(string, string) bool) Object {
	checkFunctionArgsMin(name, args, kwArgs, 2, StrType, ObjectType, IntType, IntType)
	s := args[0].(Str).Value()
	var matches []string
	switch marg := args[1]; {
	case marg.IsInstance(TupleType):
		t := marg.(Tuple)
		matches = make([]string, len(t.Elems()))
		for i, o := range t.Elems() {
			if !o.IsInstance(StrType) {
				panic(RaiseType(TypeErrorType, fmt.Sprintf(
					"tuple for %s must only contain str, not %s",
					name, o.Type().Name())))
			}
			matches[i] = o.(Str).Value()
		}
	case marg.IsInstance(StrType):
		matches = []string{marg.(Str).Value()}
	default:
		panic(RaiseType(TypeErrorType, fmt.Sprintf(
			"%s first arg must be str or a tuple of str, not %s",
			name, marg.Type().Name())))
	}
	argc := len(args)
	l := len(s)
	start, end := 0, l
	if argc >= 3 {
		start = args[2].(Int).Value()
	}
	if argc >= 4 {
		end = args[3].(Int).Value()
	}
	start = seqClampIndex(start, l)
	end = seqClampIndex(end, l)
	if start > end {
		return False
	}
	s = s[start:end]
	for _, match := range matches {
		if matchFn(s, match) {
			return True
		}
	}
	return False
}

func strStrip(args Args, kwArgs KWArgs) Object {
	return strStripper("strip", strings.Trim, strings.TrimFunc, args, kwArgs)
}

func strLStrip(args Args, kwArgs KWArgs) Object {
	return strStripper("lstrip", strings.TrimLeft, strings.TrimLeftFunc, args, kwArgs)
}

func strRStrip(args Args, kwArgs KWArgs) Object {
	return strStripper("rstrip", strings.TrimRight, strings.TrimRightFunc, args, kwArgs)
}

func strStripper(name string, trim func(string, string) string,
	trimFunc func(string, func(rune) bool) string, args Args, kwArgs KWArgs) Object {

	checkFunctionArgsMin(name, args, kwArgs, 1, StrType, ObjectType)
	s := args[0].(Str).Value()
	chars := ""
	if len(args) >= 2 && args[1] != None {
		chars = args[1].(Str).Value()
	}
	if chars == "" {
		return NewStr(trimFunc(s, unicode.IsSpace))
	}
	return NewStr(trim(s, chars))
}

func strUpper(args Args, kwArgs KWArgs) Object {
	checkFunctionArgs("lower", args, kwArgs, StrType)
	s := args[0].(Str)
	return NewStr(strings.ToUpper(s.Value()))
}

func initStrType(slots *typeSlots, dict map[string]Object) {
	dict["capitalize"] = newBuiltinFunction("capitalize", strCapitalize)
	dict["encode"] = newBuiltinFunction("encode", strEncode)
	dict["endswith"] = newBuiltinFunction("endswith", strEndswith)
	dict["find"] = newBuiltinFunction("find", strFind)
	dict["format"] = newBuiltinFunction("format", strFormat)
	dict["index"] = newBuiltinFunction("index", strIndex)
	dict["isalnum"] = newBuiltinFunction("isalnum", strIsAlNum)
	dict["isalpha"] = newBuiltinFunction("isalpha", strIsAlpha)
	dict["isascii"] = newBuiltinFunction("isascii", strIsASCII)
	dict["isdecimal"] = newBuiltinFunction("isdecimal", strIsDecimal)
	dict["isdigit"] = newBuiltinFunction("isdigit", strIsDigit)
	dict["islower"] = newBuiltinFunction("islower", strIsLower)
	dict["isnumeric"] = newBuiltinFunction("isnumeric", strIsNumeric)
	dict["isprintable"] = newBuiltinFunction("isprintable", strIsPrintable)
	dict["isupper"] = newBuiltinFunction("isupper", strIsUpper)
	dict["isspace"] = newBuiltinFunction("isspace", strIsSpace)
	dict["join"] = newBuiltinFunction("join", strJoin)
	dict["lower"] = newBuiltinFunction("lower", strLower)
	dict["lstrip"] = newBuiltinFunction("lstrip", strLStrip)
	dict["partition"] = newBuiltinFunction("partition", strPartition)
	dict["replace"] = newBuiltinFunction("replace", strReplace)
	dict["rfind"] = newBuiltinFunction("rfind", strRFind)
	dict["rindex"] = newBuiltinFunction("rindex", strRIndex)
	dict["rpartition"] = newBuiltinFunction("rpartition", strRPartition)
	dict["rstrip"] = newBuiltinFunction("rstrip", strRStrip)
	dict["split"] = newBuiltinFunction("split", strSplit)
	dict["splitlines"] = newBuiltinFunction("splitlines", strSplitlines)
	dict["startswith"] = newBuiltinFunction("startswith", strStartswith)
	dict["strip"] = newBuiltinFunction("strip", strStrip)
	dict["upper"] = newBuiltinFunction("upper", strUpper)
	slots.New = strNew
	slots.Hash = strHash
	slots.Repr = strRepr
	slots.Str = strStr
	slots.Int = strInt
	slots.Float = strFloat
	slots.GetItem = strGetItem
	slots.Len = strLen
	slots.Eq = makeCompareFn(strCompare, compareEq)
	slots.Ne = makeCompareFn(strCompare, compareNe)
	slots.Lt = makeCompareFn(strCompare, compareLt)
	slots.Gt = makeCompareFn(strCompare, compareGt)
	slots.Le = makeCompareFn(strCompare, compareLe)
	slots.Ge = makeCompareFn(strCompare, compareGe)
	slots.Add = strAdd
	slots.Mul = strMul
	slots.Mod = strMod
	slots.Contains = strContains
}
