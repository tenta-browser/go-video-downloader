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
 * request/module.go: transpiled from urllib/request.py
 */

package request

import (
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	HTTPCookieProcessor λ.Object
	HTTPHandler         λ.Object
	HTTPRedirectHandler λ.Object
	HTTPSHandler        λ.Object
	ProxyHandler        λ.Object
	Request             λ.Object
)

func init() {
	λ.InitModule(func() {
		Request = λ.Cal(λ.TypeType, λ.StrLiteral("Request"), λ.NewTuple(), func() λ.Dict {
			var (
				Request___init__     λ.Object
				Request_add_header   λ.Object
				Request_get_full_url λ.Object
				Request_get_header   λ.Object
				Request_get_method   λ.Object
			)
			Request___init__ = λ.NewFunction("__init__",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
					{Name: "data", Def: λ.None},
					{Name: "headers", Def: λ.None},
					{Name: "origin_req_host", Def: λ.None},
					{Name: "unverifiable", Def: λ.False},
					{Name: "method", Def: λ.None},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdata            = λargs[2]
						ϒheaders         = λargs[3]
						ϒkey             λ.Object
						ϒmethod          = λargs[6]
						ϒorigin_req_host = λargs[4]
						ϒself            = λargs[0]
						ϒunverifiable    = λargs[5]
						ϒurl             = λargs[1]
						ϒvalue           λ.Object
						τmp0             λ.Object
						τmp1             λ.Object
						τmp2             λ.Object
					)
					λ.SetAttr(ϒself, "url", ϒurl)
					λ.SetAttr(ϒself, "headers", λ.DictLiteral(map[λ.Object]λ.Object{}))
					λ.SetAttr(ϒself, "data", ϒdata)
					if λ.IsTrue(ϒheaders) {
						τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(ϒheaders, "items"))
						for {
							if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
								break
							}
							τmp2 = λ.UnpackIterable(τmp1, 2)
							ϒkey = λ.GetItem(τmp2, λ.IntLiteral(0))
							ϒvalue = λ.GetItem(τmp2, λ.IntLiteral(1))
							λ.Calm(ϒself, "add_header", ϒkey, ϒvalue)
						}
					}
					λ.SetAttr(ϒself, "origin_req_host", ϒorigin_req_host)
					λ.SetAttr(ϒself, "unverifiable", ϒunverifiable)
					λ.SetAttr(ϒself, "method", ϒmethod)
					return λ.None
				})
			Request_add_header = λ.NewFunction("add_header",
				[]λ.Param{
					{Name: "self"},
					{Name: "key"},
					{Name: "value"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒkey   = λargs[1]
						ϒself  = λargs[0]
						ϒvalue = λargs[2]
					)
					λ.SetItem(λ.GetAttr(ϒself, "headers", nil), λ.Calm(ϒkey, "capitalize"), ϒvalue)
					return λ.None
				})
			Request_get_header = λ.NewFunction("get_header",
				[]λ.Param{
					{Name: "self"},
					{Name: "header_name"},
					{Name: "default", Def: λ.None},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdefault     = λargs[2]
						ϒheader_name = λargs[1]
						ϒself        = λargs[0]
					)
					return λ.Calm(λ.GetAttr(ϒself, "headers", nil), "get", λ.Calm(ϒheader_name, "capitalize"), ϒdefault)
				})
			Request_get_full_url = λ.NewFunction("get_full_url",
				[]λ.Param{
					{Name: "self"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒself = λargs[0]
					)
					return λ.GetAttr(ϒself, "url", nil)
				})
			Request_get_method = λ.NewFunction("get_method",
				[]λ.Param{
					{Name: "self"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒself = λargs[0]
					)
					if λ.IsTrue(λ.GetAttr(ϒself, "method", nil)) {
						return λ.GetAttr(ϒself, "method", nil)
					}
					if λ.IsTrue(λ.GetAttr(ϒself, "data", nil)) {
						return λ.StrLiteral("POST")
					}
					return λ.StrLiteral("GET")
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"__init__":     Request___init__,
				"add_header":   Request_add_header,
				"get_full_url": Request_get_full_url,
				"get_header":   Request_get_header,
				"get_method":   Request_get_method,
			})
		}())
		HTTPHandler = λ.Cal(λ.TypeType, λ.StrLiteral("HTTPHandler"), λ.NewTuple(), func() λ.Dict {

			return λ.ClassDictLiteral(map[λ.Object]λ.Object{})
		}())
		HTTPSHandler = λ.Cal(λ.TypeType, λ.StrLiteral("HTTPSHandler"), λ.NewTuple(), func() λ.Dict {

			return λ.ClassDictLiteral(map[λ.Object]λ.Object{})
		}())
		HTTPRedirectHandler = λ.Cal(λ.TypeType, λ.StrLiteral("HTTPRedirectHandler"), λ.NewTuple(), func() λ.Dict {
			var (
				HTTPRedirectHandler_http_error_302 λ.Object
			)
			HTTPRedirectHandler_http_error_302 = λ.NewFunction("http_error_302",
				[]λ.Param{
					{Name: "self"},
				},
				0, true, true,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒargs   = λargs[1]
						ϒkwargs = λargs[2]
						ϒself   = λargs[0]
					)
					_ = ϒargs
					_ = ϒkwargs
					_ = ϒself
					panic(λ.Raise(λ.Cal(λ.NotImplementedErrorType, λ.StrLiteral("urllib.request.HTTPRedirectHandler.http_error_302"))))
					return λ.None
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"http_error_302": HTTPRedirectHandler_http_error_302,
			})
		}())
		HTTPCookieProcessor = λ.Cal(λ.TypeType, λ.StrLiteral("HTTPCookieProcessor"), λ.NewTuple(), func() λ.Dict {
			var (
				HTTPCookieProcessor_http_request λ.Object
			)
			HTTPCookieProcessor_http_request = λ.Cal(λ.TypeType, λ.StrLiteral("http_request"), λ.NewTuple(), func() λ.Dict {

				return λ.ClassDictLiteral(map[λ.Object]λ.Object{})
			}())
			return λ.ClassDictLiteral(map[string]λ.Object{
				"http_request": HTTPCookieProcessor_http_request,
			})
		}())
		ProxyHandler = λ.Cal(λ.TypeType, λ.StrLiteral("ProxyHandler"), λ.NewTuple(), func() λ.Dict {

			return λ.ClassDictLiteral(map[λ.Object]λ.Object{})
		}())
	})
}
