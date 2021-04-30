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
 * ElementTree/module.go: transpiled from xml/etree/ElementTree.py
 */

package ElementTree

import (
	Ωencoding "github.com/tenta-browser/go-video-downloader/lib/encoding"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	Element            λ.Object
	ParseError         λ.Object
	XML                λ.Object
	XMLElementInfos    λ.Object
	XMLFind            λ.Object
	XMLFindAll         λ.Object
	XMLGetChildElement λ.Object
	XMLParse           λ.Object
)

func init() {
	λ.InitModule(func() {
		ParseError = Ωencoding.XMLParseErrorType
		XMLParse = Ωencoding.XMLParse
		XMLElementInfos = Ωencoding.XMLElementInfos
		XMLGetChildElement = Ωencoding.XMLGetChildElement
		XMLFind = Ωencoding.XMLFind
		XMLFindAll = Ωencoding.XMLFindAll
		XML = λ.NewFunction("XML",
			[]λ.Param{
				{Name: "text"},
			},
			0, false, false,
			func(λargs []λ.Object) λ.Object {
				var (
					ϒe    λ.Object
					ϒtext = λargs[0]
				)
				ϒe = λ.Cal(XMLParse, ϒtext)
				return func() λ.Object {
					if λ.IsTrue(ϒe) {
						return λ.Cal(Element, ϒe)
					} else {
						return λ.None
					}
				}()
			})
		Element = λ.Cal(λ.TypeType, λ.StrLiteral("Element"), λ.NewTuple(), func() λ.Dict {
			var (
				Element___getitem__ λ.Object
				Element___init__    λ.Object
				Element___len__     λ.Object
				Element_find        λ.Object
				Element_findall     λ.Object
				Element_findtext    λ.Object
				Element_get         λ.Object
			)
			Element___init__ = λ.NewFunction("__init__",
				[]λ.Param{
					{Name: "self"},
					{Name: "e"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒe    = λargs[1]
						ϒself = λargs[0]
						ϒtext λ.Object
						τmp0  λ.Object
					)
					λ.SetAttr(ϒself, "e", ϒe)
					τmp0 = λ.UnpackIterable(λ.Cal(XMLElementInfos, ϒe), 4)
					λ.SetAttr(ϒself, "tag", λ.GetItem(τmp0, λ.IntLiteral(0)))
					ϒtext = λ.GetItem(τmp0, λ.IntLiteral(1))
					λ.SetAttr(ϒself, "attrib", λ.GetItem(τmp0, λ.IntLiteral(2)))
					λ.SetAttr(ϒself, "childno", λ.GetItem(τmp0, λ.IntLiteral(3)))
					λ.SetAttr(ϒself, "text", func() λ.Object {
						if λ.IsTrue(ϒtext) {
							return ϒtext
						} else {
							return λ.None
						}
					}())
					return λ.None
				})
			Element___len__ = λ.NewFunction("__len__",
				[]λ.Param{
					{Name: "self"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒself = λargs[0]
					)
					return λ.GetAttr(ϒself, "childno", nil)
				})
			Element___getitem__ = λ.NewFunction("__getitem__",
				[]λ.Param{
					{Name: "self"},
					{Name: "key"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒkey  = λargs[1]
						ϒself = λargs[0]
					)
					return λ.Cal(Element, λ.Cal(XMLGetChildElement, λ.GetAttr(ϒself, "e", nil), ϒkey))
				})
			Element_get = λ.NewFunction("get",
				[]λ.Param{
					{Name: "self"},
					{Name: "key"},
					{Name: "default", Def: λ.None},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdefault = λargs[2]
						ϒkey     = λargs[1]
						ϒself    = λargs[0]
					)
					return λ.Calm(λ.GetAttr(ϒself, "attrib", nil), "get", ϒkey, ϒdefault)
				})
			Element_find = λ.NewFunction("find",
				[]λ.Param{
					{Name: "self"},
					{Name: "path"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒe    λ.Object
						ϒpath = λargs[1]
						ϒself = λargs[0]
					)
					ϒe = λ.Cal(XMLFind, λ.GetAttr(ϒself, "e", nil), ϒpath)
					return func() λ.Object {
						if λ.IsTrue(ϒe) {
							return λ.Cal(Element, ϒe)
						} else {
							return λ.None
						}
					}()
				})
			Element_findtext = λ.NewFunction("findtext",
				[]λ.Param{
					{Name: "self"},
					{Name: "path"},
					{Name: "default", Def: λ.None},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdefault = λargs[2]
						ϒel      λ.Object
						ϒpath    = λargs[1]
						ϒself    = λargs[0]
					)
					ϒel = λ.Calm(ϒself, "find", ϒpath)
					return func() λ.Object {
						if ϒel != λ.None {
							return λ.GetAttr(ϒel, "text", nil)
						} else {
							return ϒdefault
						}
					}()
				})
			Element_findall = λ.NewFunction("findall",
				[]λ.Param{
					{Name: "self"},
					{Name: "path"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒes   λ.Object
						ϒpath = λargs[1]
						ϒself = λargs[0]
					)
					ϒes = λ.Cal(XMLFindAll, λ.GetAttr(ϒself, "e", nil), ϒpath)
					return λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
						nil,
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							return λ.NewGenerator(func(λgy λ.Yielder) λ.Object {
								var (
									ϒe   λ.Object
									τmp0 λ.Object
									τmp1 λ.Object
								)
								τmp0 = λ.Cal(λ.BuiltinIter, ϒes)
								for {
									if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
										break
									}
									ϒe = τmp1
									λgy.Yield(λ.Cal(Element, ϒe))
								}
								return λ.None
							})
						})))
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"__getitem__": Element___getitem__,
				"__init__":    Element___init__,
				"__len__":     Element___len__,
				"find":        Element_find,
				"findall":     Element_findall,
				"findtext":    Element_findtext,
				"get":         Element_get,
			})
		}())
	})
}
