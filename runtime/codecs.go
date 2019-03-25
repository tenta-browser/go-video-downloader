/**
 * Go Video Downloader
 *
 *    Copyright 2019 Tenta, LLC
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
 * codecs.go: Implementation of Python's codecs;
 *            only for bytes<->string unicode conversions
 */

package runtime

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

const (
	codingErrorsStrict  = "strict"
	codingErrorsIgnore  = "ignore"
	codingErrorsReplace = "replace"
)

const (
	codingErrEncode             = "'%s' codec can't encode character %v in position %d"
	codingErrDecode             = "'%s' codec can't decode byte 0x%02x in position %d"
	codingErrUnkownErrorHandler = "unknown error handler name '%s'"
)

// Codec ..
type Codec interface {
	Name() string
	Encode(obj Object, errors Str) (Object, Int)
	Decode(obj Object, errors Str) (Object, Int)
}

var (
	utf8Codec          = utf8CodecStruct{}
	asciiCodec         = asciiCodecStruct{}
	unicodeEscapeCodec = unicodeEscapeCodecStruct{}
)

var codecs = map[string]Codec{
	utf8Codec.Name():          utf8Codec,
	asciiCodec.Name():         asciiCodec,
	unicodeEscapeCodec.Name(): unicodeEscapeCodec,
}

// GetCodec ..
func GetCodec(name string) Codec {
	name = strings.ToLower(name)
	codec, ok := codecs[name]
	if !ok {
		panic(RaiseType(LookupErrorType, fmt.Sprintf("unknown encoding: %s", name)))
	}
	return codec
}

// ---- utf-8 ----

type utf8CodecStruct struct{}

func (c utf8CodecStruct) Name() string {
	return "utf-8"
}

func (c utf8CodecStruct) Encode(obj Object, errors Str) (Object, Int) {
	return encodeUtf8orASCII(obj, errors, c.Name())
}

func encodeUtf8orASCII(obj Object, errors Str, cname string) (Object, Int) {
	s := obj.(Str).Value()
	var isValid func(rune) bool
	var repl rune
	if cname == utf8Codec.Name() {
		isValid = utf8.ValidRune
		repl = unicode.ReplacementChar
	} else {
		isValid = func(r rune) bool {
			return r < 128
		}
		repl = '?'
	}
	b := bytes.Buffer{}
	cnt := 0
	for i, r := range s {
		switch errors := errors.Value(); true {
		case isValid(r):
			b.WriteRune(r)
		case errors == codingErrorsIgnore:
			// Do nothing
		case errors == codingErrorsReplace:
			b.WriteRune(repl)
		case errors == codingErrorsStrict:
			panic(RaiseType(UnicodeEncodeErrorType, fmt.Sprintf(
				codingErrEncode, cname, r, i)))
		default:
			panic(RaiseType(LookupErrorType, fmt.Sprintf(
				codingErrUnkownErrorHandler, errors)))
		}
		cnt++
	}
	return NewBytes(b.Bytes()...), NewInt(cnt)
}

func (c utf8CodecStruct) Decode(obj Object, errors Str) (Object, Int) {
	bytes := obj.(Bytes).Value()
	s := string(bytes)
	sb := strings.Builder{}
	for pos, r := range s {
		switch errors := errors.Value(); true {
		case r != utf8.RuneError:
			sb.WriteRune(r)
		case errors == codingErrorsIgnore:
			// Do nothing
		case errors == codingErrorsReplace:
			sb.WriteRune(unicode.ReplacementChar)
		case errors == codingErrorsStrict:
			panic(RaiseType(UnicodeDecodeErrorType, fmt.Sprintf(
				codingErrDecode, c.Name(), int(s[pos]), pos)))
		default:
			panic(RaiseType(LookupErrorType, fmt.Sprintf(
				codingErrUnkownErrorHandler, errors)))
		}
	}
	return NewStr(sb.String()), NewInt(len(bytes))
}

// ---- ascii ----

type asciiCodecStruct struct{}

func (c asciiCodecStruct) Name() string {
	return "ascii"
}

func (c asciiCodecStruct) Encode(obj Object, errors Str) (Object, Int) {
	return encodeUtf8orASCII(obj, errors, c.Name())
}

func (c asciiCodecStruct) Decode(obj Object, errors Str) (Object, Int) {
	bytes := obj.(Bytes).Value()
	sb := strings.Builder{}
	for pos, b := range bytes {
		switch errors := errors.Value(); true {
		case b < 128:
			sb.WriteByte(b)
		case errors == codingErrorsIgnore:
			// Do nothing
		case errors == codingErrorsReplace:
			sb.WriteRune(unicode.ReplacementChar)
		case errors == codingErrorsStrict:
			panic(RaiseType(UnicodeDecodeErrorType, fmt.Sprintf(
				codingErrDecode, c.Name(), b, pos)))
		default:
			panic(RaiseType(LookupErrorType, fmt.Sprintf(
				codingErrUnkownErrorHandler, errors)))
		}
	}
	return NewStr(sb.String()), NewInt(len(bytes))
}

// ---- unicode_escape ----

type unicodeEscapeCodecStruct struct{}

func (c unicodeEscapeCodecStruct) Name() string {
	return "unicode_escape"
}

func (c unicodeEscapeCodecStruct) Encode(obj Object, errors Str) (Object, Int) {
	s := obj.(Str).Value()
	b := bytes.Buffer{}
	cnt := 0
	for i, r := range s {
		switch errors := errors.Value(); true {
		case utf8.ValidRune(r):
			// TODO horrifically inefficient
			rs := strconv.QuoteRuneToASCII(r)
			b.WriteString(rs[1 : len(rs)-1])
		case errors == codingErrorsIgnore:
			// Do nothing
		case errors == codingErrorsReplace:
			b.WriteString(strconv.QuoteRuneToASCII(unicode.ReplacementChar))
		case errors == codingErrorsStrict:
			panic(RaiseType(UnicodeEncodeErrorType, fmt.Sprintf(
				codingErrEncode, c.Name(), r, i)))
		default:
			panic(RaiseType(LookupErrorType, fmt.Sprintf(
				codingErrUnkownErrorHandler, errors)))
		}
		cnt++
	}
	return NewBytes(b.Bytes()...), NewInt(cnt)
}

func (c unicodeEscapeCodecStruct) Decode(obj Object, errors Str) (Object, Int) {
	var s string
	if obj.IsInstance(StrType) {
		s = obj.(Str).Value()
	} else {
		s = string(obj.(Bytes).Value())
	}
	ls := len(s)
	sb := strings.Builder{}
	pos := 0
	for s != "" {
		var r rune
		var err error
		r, _, s, err = strconv.UnquoteChar(s, '\'')
		switch errors := errors.Value(); true {
		case err == nil:
			sb.WriteRune(r)
		case errors == codingErrorsIgnore:
			// Do nothing
		case errors == codingErrorsReplace:
			sb.WriteRune(unicode.ReplacementChar)
		case errors == codingErrorsStrict:
			panic(RaiseType(UnicodeDecodeErrorType, fmt.Sprintf(
				codingErrDecode, c.Name(), int(s[pos]), pos)))
		default:
			panic(RaiseType(LookupErrorType, fmt.Sprintf(
				codingErrUnkownErrorHandler, errors)))
		}
		pos = ls - len(s)
	}
	return NewStr(sb.String()), NewInt(ls)
}
