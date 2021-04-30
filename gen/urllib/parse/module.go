// Code generated by transpiler. DO NOT EDIT.

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
 * parse/module.go: transpiled from urllib/parse.py
 */

package parse

import (
	Ωnet "github.com/tenta-browser/go-video-downloader/lib/net"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ParseResult   λ.Object
	URLParse      λ.Object
	ϒparse_qs     λ.Object
	ϒunquote      λ.Object
	ϒunquote_plus λ.Object
	ϒurlencode    λ.Object
	ϒurljoin      λ.Object
	ϒurlparse     λ.Object
	ϒurlunparse   λ.Object
)

func init() {
	λ.InitModule(func() {
		ϒurljoin = Ωnet.URLJoin
		URLParse = Ωnet.URLParse
		ϒurlunparse = Ωnet.URLUnparse
		ϒunquote = Ωnet.Unquote
		ϒunquote_plus = Ωnet.UnquotePlus
		ϒparse_qs = Ωnet.ParseQS
		ϒurlencode = Ωnet.URLEncode
		ParseResult = λ.Cal(λ.TypeType, λ.StrLiteral("ParseResult"), λ.NewTuple(λ.TupleType), func() λ.Dict {
			var (
				ParseResult__replace λ.Object
				ParseResult_fragment λ.Object
				ParseResult_path     λ.Object
				ParseResult_query    λ.Object
				ParseResult_scheme   λ.Object
			)
			ParseResult_scheme = λ.NewFunction("scheme",
				[]λ.Param{
					{Name: "self"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒself = λargs[0]
					)
					return λ.GetItem(ϒself, λ.IntLiteral(0))
				})
			ParseResult_scheme = λ.Cal(λ.PropertyType, ParseResult_scheme)
			ParseResult_path = λ.NewFunction("path",
				[]λ.Param{
					{Name: "self"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒself = λargs[0]
					)
					return λ.GetItem(ϒself, λ.IntLiteral(2))
				})
			ParseResult_path = λ.Cal(λ.PropertyType, ParseResult_path)
			ParseResult_query = λ.NewFunction("query",
				[]λ.Param{
					{Name: "self"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒself = λargs[0]
					)
					return λ.GetItem(ϒself, λ.IntLiteral(4))
				})
			ParseResult_query = λ.Cal(λ.PropertyType, ParseResult_query)
			ParseResult_fragment = λ.NewFunction("fragment",
				[]λ.Param{
					{Name: "self"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒself = λargs[0]
					)
					return λ.GetItem(ϒself, λ.IntLiteral(5))
				})
			ParseResult_fragment = λ.Cal(λ.PropertyType, ParseResult_fragment)
			ParseResult__replace = λ.NewFunction("_replace",
				[]λ.Param{
					{Name: "self"},
				},
				0, false, true,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒi      λ.Object
						ϒkwargs = λargs[1]
						ϒname   λ.Object
						ϒnames  λ.Object
						ϒself   = λargs[0]
						ϒvals   λ.Object
						τmp0    λ.Object
						τmp1    λ.Object
						τmp2    λ.Object
					)
					ϒnames = λ.NewTuple(
						λ.StrLiteral("scheme"),
						λ.StrLiteral("netloc"),
						λ.StrLiteral("path"),
						λ.StrLiteral("params"),
						λ.StrLiteral("query"),
						λ.StrLiteral("fragment"),
					)
					ϒvals = λ.Cal(λ.ListType, ϒself)
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.EnumerateIteratorType, ϒnames))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						τmp2 = λ.UnpackIterable(τmp1, 2)
						ϒi = λ.GetItem(τmp2, λ.IntLiteral(0))
						ϒname = λ.GetItem(τmp2, λ.IntLiteral(1))
						if λ.Contains(ϒkwargs, ϒname) {
							λ.SetItem(ϒvals, ϒi, λ.GetItem(ϒkwargs, ϒname))
						}
					}
					return λ.Cal(ParseResult, ϒvals)
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_replace": ParseResult__replace,
				"fragment": ParseResult_fragment,
				"path":     ParseResult_path,
				"query":    ParseResult_query,
				"scheme":   ParseResult_scheme,
			})
		}())
		ϒurlparse = λ.NewFunction("urlparse",
			[]λ.Param{
				{Name: "urlstring"},
				{Name: "scheme", Def: λ.StrLiteral("")},
			},
			0, false, false,
			func(λargs []λ.Object) λ.Object {
				var (
					ϒ_scheme   λ.Object
					ϒfragment  λ.Object
					ϒnetloc    λ.Object
					ϒparams    λ.Object
					ϒpath      λ.Object
					ϒquery     λ.Object
					ϒscheme    = λargs[1]
					ϒurlstring = λargs[0]
					τmp0       λ.Object
				)
				τmp0 = λ.UnpackIterable(λ.Cal(URLParse, ϒurlstring), 6)
				ϒ_scheme = λ.GetItem(τmp0, λ.IntLiteral(0))
				ϒnetloc = λ.GetItem(τmp0, λ.IntLiteral(1))
				ϒpath = λ.GetItem(τmp0, λ.IntLiteral(2))
				ϒparams = λ.GetItem(τmp0, λ.IntLiteral(3))
				ϒquery = λ.GetItem(τmp0, λ.IntLiteral(4))
				ϒfragment = λ.GetItem(τmp0, λ.IntLiteral(5))
				return λ.Cal(ParseResult, λ.NewTuple(
					func() λ.Object {
						if λv := ϒ_scheme; λ.IsTrue(λv) {
							return λv
						} else {
							return ϒscheme
						}
					}(),
					ϒnetloc,
					ϒpath,
					ϒparams,
					ϒquery,
					ϒfragment,
				))
			})
	})
}
