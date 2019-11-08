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
 * operator/module.go: transpiled from operator.py
 */

package operator

import (
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ϒadd        λ.Object
	ϒand_       λ.Object
	ϒconcat     λ.Object
	ϒitemgetter λ.Object
	ϒlshift     λ.Object
	ϒmod        λ.Object
	ϒmul        λ.Object
	ϒor_        λ.Object
	ϒrshift     λ.Object
	ϒsub        λ.Object
	ϒtruediv    λ.Object
	ϒxor        λ.Object
)

func init() {
	λ.InitModule(func() {
		ϒadd = λ.NewFunction("add",
			[]λ.Param{
				{Name: "a"},
				{Name: "b"},
			},
			0, false, false,
			func(λargs []λ.Object) λ.Object {
				var (
					ϒa = λargs[0]
					ϒb = λargs[1]
				)
				return λ.Add(ϒa, ϒb)
			})
		ϒsub = λ.NewFunction("sub",
			[]λ.Param{
				{Name: "a"},
				{Name: "b"},
			},
			0, false, false,
			func(λargs []λ.Object) λ.Object {
				var (
					ϒa = λargs[0]
					ϒb = λargs[1]
				)
				return λ.Sub(ϒa, ϒb)
			})
		ϒmul = λ.NewFunction("mul",
			[]λ.Param{
				{Name: "a"},
				{Name: "b"},
			},
			0, false, false,
			func(λargs []λ.Object) λ.Object {
				var (
					ϒa = λargs[0]
					ϒb = λargs[1]
				)
				return λ.Mul(ϒa, ϒb)
			})
		ϒtruediv = λ.NewFunction("truediv",
			[]λ.Param{
				{Name: "a"},
				{Name: "b"},
			},
			0, false, false,
			func(λargs []λ.Object) λ.Object {
				var (
					ϒa = λargs[0]
					ϒb = λargs[1]
				)
				return λ.TrueDiv(ϒa, ϒb)
			})
		ϒmod = λ.NewFunction("mod",
			[]λ.Param{
				{Name: "a"},
				{Name: "b"},
			},
			0, false, false,
			func(λargs []λ.Object) λ.Object {
				var (
					ϒa = λargs[0]
					ϒb = λargs[1]
				)
				return λ.Mod(ϒa, ϒb)
			})
		ϒor_ = λ.NewFunction("or_",
			[]λ.Param{
				{Name: "a"},
				{Name: "b"},
			},
			0, false, false,
			func(λargs []λ.Object) λ.Object {
				var (
					ϒa = λargs[0]
					ϒb = λargs[1]
				)
				return λ.Or(ϒa, ϒb)
			})
		ϒxor = λ.NewFunction("xor",
			[]λ.Param{
				{Name: "a"},
				{Name: "b"},
			},
			0, false, false,
			func(λargs []λ.Object) λ.Object {
				var (
					ϒa = λargs[0]
					ϒb = λargs[1]
				)
				return λ.Xor(ϒa, ϒb)
			})
		ϒand_ = λ.NewFunction("and_",
			[]λ.Param{
				{Name: "a"},
				{Name: "b"},
			},
			0, false, false,
			func(λargs []λ.Object) λ.Object {
				var (
					ϒa = λargs[0]
					ϒb = λargs[1]
				)
				return λ.And(ϒa, ϒb)
			})
		ϒlshift = λ.NewFunction("lshift",
			[]λ.Param{
				{Name: "a"},
				{Name: "b"},
			},
			0, false, false,
			func(λargs []λ.Object) λ.Object {
				var (
					ϒa = λargs[0]
					ϒb = λargs[1]
				)
				return λ.LShift(ϒa, ϒb)
			})
		ϒrshift = λ.NewFunction("rshift",
			[]λ.Param{
				{Name: "a"},
				{Name: "b"},
			},
			0, false, false,
			func(λargs []λ.Object) λ.Object {
				var (
					ϒa = λargs[0]
					ϒb = λargs[1]
				)
				return λ.RShift(ϒa, ϒb)
			})
		ϒconcat = λ.NewFunction("concat",
			[]λ.Param{
				{Name: "a"},
				{Name: "b"},
			},
			0, false, false,
			func(λargs []λ.Object) λ.Object {
				var (
					ϒa   = λargs[0]
					ϒb   = λargs[1]
					ϒmsg λ.Object
				)
				λ.NewStr("Same as a + b, for a and b sequences.")
				if λ.IsTrue(λ.NewBool(!λ.IsTrue(λ.Cal(λ.BuiltinHasAttr, ϒa, λ.NewStr("__getitem__"))))) {
					ϒmsg = λ.Mod(λ.NewStr("'%s' object can't be concatenated"), λ.GetAttr(λ.Cal(λ.TypeType, ϒa), "__name__", nil))
					panic(λ.Raise(λ.Cal(λ.TypeErrorType, ϒmsg)))
				}
				return λ.Add(ϒa, ϒb)
			})
		ϒitemgetter = λ.Cal(λ.TypeType, λ.NewStr("itemgetter"), λ.NewTuple(), func() λ.Dict {
			var (
				ϒitemgetter___init__ λ.Object
			)
			ϒitemgetter___init__ = λ.NewFunction("__init__",
				[]λ.Param{
					{Name: "self"},
					{Name: "item"},
				},
				0, true, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒfunc  λ.Object
						ϒitem  = λargs[1]
						ϒitems = λargs[2]
						ϒself  = λargs[0]
						τmp0   λ.Object
					)
					if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒitems))) {
						λ.SetAttr(ϒself, "_items", λ.NewTuple(ϒitem))
						ϒfunc = λ.NewFunction("func",
							[]λ.Param{
								{Name: "obj"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒobj = λargs[0]
								)
								return λ.GetItem(ϒobj, ϒitem)
							})
						λ.SetAttr(ϒself, "_call", ϒfunc)
					} else {
						τmp0 = λ.Add(λ.NewTuple(ϒitem), ϒitems)
						λ.SetAttr(ϒself, "_items", τmp0)
						ϒitems = τmp0
						ϒfunc = λ.NewFunction("func",
							[]λ.Param{
								{Name: "obj"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒobj = λargs[0]
								)
								return λ.Cal(λ.TupleType, λ.Cal(λ.NewFunction("<generator>",
									nil,
									0, false, false,
									func(λargs []λ.Object) λ.Object {
										return λ.NewGenerator(func(λgy λ.Yielder) λ.Object {
											var (
												ϒi   λ.Object
												τmp0 λ.Object
												τmp1 λ.Object
											)
											τmp0 = λ.Cal(λ.BuiltinIter, ϒitems)
											for {
												if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
													break
												}
												ϒi = τmp1
												λgy.Yield(λ.GetItem(ϒobj, ϒi))
											}
											return λ.None
										})
									})))
							})
						λ.SetAttr(ϒself, "_call", ϒfunc)
					}
					return λ.None
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("__init__"): ϒitemgetter___init__,
			})
		}())
	})
}
