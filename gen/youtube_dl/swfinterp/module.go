// Code generated by transpiler. DO NOT EDIT.

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
 * swfinterp/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/swfinterp.py
 */

package swfinterp

import (
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ByteArrayClass  λ.Object
	ExtractorError  λ.Object
	StringClass     λ.Object
	TimerClass      λ.Object
	TimerEventClass λ.Object
	ϒ_AVMClass      λ.Object
	ϒ_ScopeDict     λ.Object
	ϒ_Undefined     λ.Object
	ϒ_read_int      λ.Object
	ϒcompat_str     λ.Object
)

func init() {
	λ.InitModule(func() {
		ϒcompat_str = Ωcompat.ϒcompat_str
		ExtractorError = Ωutils.ExtractorError
		ϒ_ScopeDict = λ.Cal(λ.TypeType, λ.NewStr("_ScopeDict"), λ.NewTuple(λ.DictType), func() λ.Dict {
			var (
				ϒ_ScopeDict___init__ λ.Object
			)
			ϒ_ScopeDict___init__ = λ.NewFunction("__init__",
				[]λ.Param{
					{Name: "self"},
					{Name: "avm_class"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒavm_class = λargs[1]
						ϒself      = λargs[0]
					)
					λ.Cal(λ.GetAttr(λ.Cal(λ.SuperType, ϒ_ScopeDict, ϒself), "__init__", nil))
					λ.SetAttr(ϒself, "avm_class", ϒavm_class)
					return λ.None
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("__init__"): ϒ_ScopeDict___init__,
			})
		}())
		ϒ_AVMClass = λ.Cal(λ.TypeType, λ.NewStr("_AVMClass"), λ.NewTuple(λ.ObjectType), func() λ.Dict {
			var (
				ϒ_AVMClass___init__ λ.Object
			)
			ϒ_AVMClass___init__ = λ.NewFunction("__init__",
				[]λ.Param{
					{Name: "self"},
					{Name: "name_idx"},
					{Name: "name"},
					{Name: "static_properties", Def: λ.None},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒname              = λargs[2]
						ϒname_idx          = λargs[1]
						ϒself              = λargs[0]
						ϒstatic_properties = λargs[3]
					)
					λ.SetAttr(ϒself, "name_idx", ϒname_idx)
					λ.SetAttr(ϒself, "name", ϒname)
					λ.SetAttr(ϒself, "method_names", λ.NewDictWithTable(map[λ.Object]λ.Object{}))
					λ.SetAttr(ϒself, "method_idxs", λ.NewDictWithTable(map[λ.Object]λ.Object{}))
					λ.SetAttr(ϒself, "methods", λ.NewDictWithTable(map[λ.Object]λ.Object{}))
					λ.SetAttr(ϒself, "method_pyfunctions", λ.NewDictWithTable(map[λ.Object]λ.Object{}))
					λ.SetAttr(ϒself, "static_properties", func() λ.Object {
						if λ.IsTrue(ϒstatic_properties) {
							return ϒstatic_properties
						} else {
							return λ.NewDictWithTable(map[λ.Object]λ.Object{})
						}
					}())
					λ.SetAttr(ϒself, "variables", λ.Cal(ϒ_ScopeDict, ϒself))
					λ.SetAttr(ϒself, "constants", λ.NewDictWithTable(map[λ.Object]λ.Object{}))
					return λ.None
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("__init__"): ϒ_AVMClass___init__,
			})
		}())
		// unknown_binary_op:BitOr
		ϒ_read_int = λ.None
		StringClass = λ.Cal(ϒ_AVMClass, λ.NewStr("(no name idx)"), λ.NewStr("String"))
		ByteArrayClass = λ.Cal(ϒ_AVMClass, λ.NewStr("(no name idx)"), λ.NewStr("ByteArray"))
		TimerClass = λ.Cal(ϒ_AVMClass, λ.NewStr("(no name idx)"), λ.NewStr("Timer"))
		TimerEventClass = λ.Cal(ϒ_AVMClass, λ.NewStr("(no name idx)"), λ.NewStr("TimerEvent"), λ.NewDictWithTable(map[λ.Object]λ.Object{
			λ.NewStr("TIMER"): λ.NewStr("timer"),
		}))
		ϒ_Undefined = λ.Cal(λ.TypeType, λ.NewStr("_Undefined"), λ.NewTuple(λ.ObjectType), func() λ.Dict {
			var (
				ϒ_Undefined___bool__ λ.Object
				ϒ_Undefined___str__  λ.Object
			)
			ϒ_Undefined___bool__ = λ.NewFunction("__bool__",
				[]λ.Param{
					{Name: "self"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒself = λargs[0]
					)
					_ = ϒself
					return λ.False
				})
			ϒ_Undefined___str__ = λ.NewFunction("__str__",
				[]λ.Param{
					{Name: "self"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒself = λargs[0]
					)
					_ = ϒself
					return λ.NewStr("undefined")
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("__bool__"): ϒ_Undefined___bool__,
				λ.NewStr("__str__"):  ϒ_Undefined___str__,
			})
		}())
	})
}
