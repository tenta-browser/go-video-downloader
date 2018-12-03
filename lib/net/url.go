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
 * url.go: Methods for parsing, constructing, combining, encoding, decoding URLs
 */

package http

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/tenta-browser/go-video-downloader/lib"
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

// URLJoin implements approximately to python/urllib.parse.urljoin
var URLJoin = rnt.NewSimpleFunction("urljoin", []string{"base", "url"},
	func(args []rnt.Object) rnt.Object {
		base := args[0].(rnt.Str).Value()
		urlv := args[1].(rnt.Str).Value()
		baseURL, err := url.Parse(base)
		if err != nil {
			return args[1]
		}
		resultURL, err := baseURL.Parse(urlv)
		if err != nil {
			return args[1]
		}
		return rnt.NewStr(resultURL.String())
	})

// URLParse implements python/urllib.parse.urlparse
var URLParse = rnt.NewSimpleFunction("urlparse", []string{"urlstring"},
	func(args []rnt.Object) rnt.Object {
		urlv := args[0].(rnt.Str).Value()
		parsedURL, err := url.Parse(urlv)
		if err != nil {
			panic(rnt.RaiseType(rnt.ValueErrorType, fmt.Sprintf(
				"failed to parse url: '%s'", urlv)))
		}
		escPath := parsedURL.EscapedPath()
		parameters := ""
		if pos := strings.LastIndex(escPath, ";"); pos >= 0 {
			parameters = escPath[pos+1:]
			escPath = escPath[:pos]
		}
		return rnt.NewTuple(
			rnt.NewStr(parsedURL.Scheme),
			rnt.NewStr(parsedURL.Host),
			rnt.NewStr(escPath),
			rnt.NewStr(parameters),
			rnt.NewStr(parsedURL.RawQuery),
			rnt.NewStr(parsedURL.Fragment),
		)
	})

// URLUnparse implements python/urllib.parse.urlunparse
var URLUnparse = rnt.NewSimpleFunction("urlunparse", []string{"parts"},
	func(args []rnt.Object) rnt.Object {
		parts := rnt.Cal(rnt.ListType, args[0]).(rnt.List).Elems()
		if len(parts) != 6 {
			panic(rnt.RaiseType(rnt.ValueErrorType, "expected 6 values to unpack"))
		}
		// 'scheme', 'netloc', 'path', 'params','query', 'fragment'
		u := new(url.URL)
		u.Scheme = parts[0].(rnt.Str).Value()
		u.Host = parts[1].(rnt.Str).Value()
		u.RawPath = parts[2].(rnt.Str).Value()
		if params := parts[3].(rnt.Str).Value(); params != "" {
			u.RawPath += ";" + params
		}
		path, err := url.PathUnescape(u.RawPath)
		if err != nil {
			panic(rnt.RaiseType(rnt.ValueErrorType, err.Error()))
		}
		u.Path = path
		u.RawQuery = parts[4].(rnt.Str).Value()
		u.Fragment = parts[5].(rnt.Str).Value()
		return rnt.NewStr(u.String())
	})

// Unquote implements python/urllib.parse.unquote
var Unquote = rnt.NewSimpleFunction("unquote", []string{"string"},
	func(args []rnt.Object) rnt.Object {
		s := args[0].(rnt.Str).Value()
		res, err := url.PathUnescape(s)
		if err != nil {
			panic(rnt.RaiseType(lib.ExtractorErrorType, err.Error()))
		}
		return rnt.NewStr(res)
	})

// UnquotePlus implements python/urllib.parse.unquote_plus
var UnquotePlus = rnt.NewSimpleFunction("unquote_plus", []string{"string"},
	func(args []rnt.Object) rnt.Object {
		s := args[0].(rnt.Str).Value()
		res, err := url.QueryUnescape(s)
		if err != nil {
			panic(rnt.RaiseType(lib.ExtractorErrorType, err.Error()))
		}
		return rnt.NewStr(res)
	})

// ParseQS implements python/urllib.parse.parse_qs
var ParseQS = rnt.NewFunction("parse_qs",
	[]rnt.Param{
		{Name: "qs"},
		{Name: "keep_blank_values", Def: rnt.False},
	}, 0, false, false,
	func(args []rnt.Object) rnt.Object {
		qs, st := lib.StringlikeVal(args[0])
		keepBlanks := rnt.IsTrue(args[1])
		values, err := url.ParseQuery(qs)
		if err != nil {
			panic(rnt.RaiseType(rnt.ValueErrorType, err.Error()))
		}
		resDict := rnt.NewDict()
		for name, vals := range values {
			resVals := []rnt.Object{}
			for _, val := range vals {
				if val != "" || keepBlanks {
					resVals = append(resVals, lib.NewStringlike(val, st))
				}
			}
			resDict.SetItem(rnt.NewStr(name), rnt.NewList(resVals...))
		}
		return resDict
	})

// URLEncode implements python/urllib.parse.urlencode
var URLEncode = rnt.NewFunction("urlencode",
	[]rnt.Param{
		{Name: "query"},
		{Name: "doseq", Def: rnt.False},
	}, 0, false, false,
	func(args []rnt.Object) rnt.Object {
		toString := func(o rnt.Object) string {
			if o.IsInstance(rnt.BytesType) {
				return string(o.(rnt.Bytes).Value())
			}
			return rnt.ToStr(o).Value()
		}
		doSeq := rnt.IsTrue(args[1])
		if args[0].IsInstance(rnt.DictType) {
			args[0] = rnt.Cal(rnt.GetAttr(args[0], "items", nil))
		}
		pairs := rnt.Cal(rnt.ListType, args[0]).(rnt.List).Elems()
		resValues := make(url.Values)
		for _, pair := range pairs {
			name := toString(rnt.GetItem(pair, rnt.NewInt(0)))
			val := rnt.GetItem(pair, rnt.NewInt(1))
			listVal := func() rnt.Object {
				defer rnt.Catch(rnt.TypeErrorType, nil, nil)
				return rnt.Cal(rnt.ListType, val)
			}()
			if listVal != nil && doSeq {
				for _, val := range listVal.(rnt.List).Elems() {
					resValues.Add(name, toString(val))
				}
			} else {
				resValues.Add(name, toString(val))
			}
		}
		return rnt.NewStr(resValues.Encode())
	})
