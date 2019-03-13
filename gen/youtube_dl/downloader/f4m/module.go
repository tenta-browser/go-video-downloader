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
 * f4m/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/downloader/f4m.py
 */

package f4m

import (
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωfragment "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/downloader/fragment"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	FragmentFD                    λ.Object
	ϒ_add_ns                      λ.Object
	ϒcompat_b64decode             λ.Object
	ϒcompat_etree_fromstring      λ.Object
	ϒcompat_urllib_parse_urlparse λ.Object
	ϒfix_xml_ampersands           λ.Object
	ϒget_base_url                 λ.Object
	ϒremove_encrypted_media       λ.Object
	ϒxpath_text                   λ.Object
)

func init() {
	λ.InitModule(func() {
		FragmentFD = Ωfragment.FragmentFD
		ϒcompat_b64decode = Ωcompat.ϒcompat_b64decode
		ϒcompat_etree_fromstring = Ωcompat.ϒcompat_etree_fromstring
		ϒcompat_urllib_parse_urlparse = Ωcompat.ϒcompat_urllib_parse_urlparse
		ϒfix_xml_ampersands = Ωutils.ϒfix_xml_ampersands
		ϒxpath_text = Ωutils.ϒxpath_text
		ϒremove_encrypted_media = λ.NewFunction("remove_encrypted_media",
			[]λ.Param{
				{Name: "media"},
			},
			0, false, false,
			func(λargs []λ.Object) λ.Object {
				var (
					ϒmedia = λargs[0]
				)
				return λ.Cal(λ.ListType, λ.Cal(λ.FilterIteratorType, λ.NewFunction("<lambda>",
					[]λ.Param{
						{Name: "e"},
					},
					0, false, false,
					func(λargs []λ.Object) λ.Object {
						var (
							ϒe = λargs[0]
						)
						return func() λ.Object {
							if λv := λ.NewBool(!λ.Contains(λ.GetAttr(ϒe, "attrib", nil), λ.NewStr("drmAdditionalHeaderId"))); !λ.IsTrue(λv) {
								return λv
							} else {
								return λ.NewBool(!λ.Contains(λ.GetAttr(ϒe, "attrib", nil), λ.NewStr("drmAdditionalHeaderSetId")))
							}
						}()
					}), ϒmedia))
			})
		ϒ_add_ns = λ.NewFunction("_add_ns",
			[]λ.Param{
				{Name: "prop"},
				{Name: "ver", Def: λ.NewInt(1)},
			},
			0, false, false,
			func(λargs []λ.Object) λ.Object {
				var (
					ϒprop = λargs[0]
					ϒver  = λargs[1]
				)
				return λ.Mod(λ.NewStr("{http://ns.adobe.com/f4m/%d.0}%s"), λ.NewTuple(
					ϒver,
					ϒprop,
				))
			})
		ϒget_base_url = λ.NewFunction("get_base_url",
			[]λ.Param{
				{Name: "manifest"},
			},
			0, false, false,
			func(λargs []λ.Object) λ.Object {
				var (
					ϒbase_url λ.Object
					ϒmanifest = λargs[0]
				)
				ϒbase_url = λ.Call(ϒxpath_text, λ.NewArgs(
					ϒmanifest,
					λ.NewList(
						λ.Cal(ϒ_add_ns, λ.NewStr("baseURL")),
						λ.Cal(ϒ_add_ns, λ.NewStr("baseURL"), λ.NewInt(2)),
					),
					λ.NewStr("base URL"),
				), λ.KWArgs{
					{Name: "default", Value: λ.None},
				})
				if λ.IsTrue(ϒbase_url) {
					ϒbase_url = λ.Cal(λ.GetAttr(ϒbase_url, "strip", nil))
				}
				return ϒbase_url
			})
	})
}
