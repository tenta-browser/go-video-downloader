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
 * ina/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/ina.py
 */

package ina

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InaIE          λ.Object
	InfoExtractor  λ.Object
	ϒdetermine_ext λ.Object
	ϒint_or_none   λ.Object
	ϒstrip_or_none λ.Object
	ϒxpath_attr    λ.Object
	ϒxpath_text    λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒdetermine_ext = Ωutils.ϒdetermine_ext
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒstrip_or_none = Ωutils.ϒstrip_or_none
		ϒxpath_attr = Ωutils.ϒxpath_attr
		ϒxpath_text = Ωutils.ϒxpath_text
		InaIE = λ.Cal(λ.TypeType, λ.StrLiteral("InaIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				InaIE__VALID_URL    λ.Object
				InaIE__real_extract λ.Object
			)
			InaIE__VALID_URL = λ.StrLiteral("https?://(?:(?:www|m)\\.)?ina\\.fr/(?:video|audio)/(?P<id>[A-Z0-9_]+)")
			InaIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒcontent        λ.Object
						ϒext            λ.Object
						ϒformats        λ.Object
						ϒfurl           λ.Object
						ϒget_furl       λ.Object
						ϒh              λ.Object
						ϒinfo_doc       λ.Object
						ϒitem           λ.Object
						ϒmedia_ns_xpath λ.Object
						ϒq              λ.Object
						ϒq_url          λ.Object
						ϒself           = λargs[0]
						ϒthumbnail      λ.Object
						ϒthumbnail_url  λ.Object
						ϒthumbnails     λ.Object
						ϒtitle          λ.Object
						ϒurl            = λargs[1]
						ϒvideo_id       λ.Object
						ϒw              λ.Object
						τmp0            λ.Object
						τmp1            λ.Object
						τmp2            λ.Object
					)
					ϒvideo_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒinfo_doc = λ.Calm(ϒself, "_download_xml", λ.Mod(λ.StrLiteral("http://player.ina.fr/notices/%s.mrss"), ϒvideo_id), ϒvideo_id)
					ϒitem = λ.Calm(ϒinfo_doc, "find", λ.StrLiteral("channel/item"))
					ϒtitle = λ.Call(ϒxpath_text, λ.NewArgs(
						ϒitem,
						λ.StrLiteral("title"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.True},
					})
					ϒmedia_ns_xpath = λ.NewFunction("<lambda>",
						[]λ.Param{
							{Name: "x"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒx = λargs[0]
							)
							return λ.Calm(ϒself, "_xpath_ns", ϒx, λ.StrLiteral("http://search.yahoo.com/mrss/"))
						})
					ϒcontent = λ.Calm(ϒitem, "find", λ.Cal(ϒmedia_ns_xpath, λ.StrLiteral("content")))
					ϒget_furl = λ.NewFunction("<lambda>",
						[]λ.Param{
							{Name: "x"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒx = λargs[0]
							)
							return λ.Cal(ϒxpath_attr, ϒcontent, λ.Cal(ϒmedia_ns_xpath, ϒx), λ.StrLiteral("url"))
						})
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.NewTuple(
						λ.NewTuple(
							λ.StrLiteral("bq"),
							λ.IntLiteral(400),
							λ.IntLiteral(300),
						),
						λ.NewTuple(
							λ.StrLiteral("mq"),
							λ.IntLiteral(512),
							λ.IntLiteral(384),
						),
						λ.NewTuple(
							λ.StrLiteral("hq"),
							λ.IntLiteral(768),
							λ.IntLiteral(576),
						),
					))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						τmp2 = τmp1
						ϒq = λ.GetItem(τmp2, λ.IntLiteral(0))
						ϒw = λ.GetItem(τmp2, λ.IntLiteral(1))
						ϒh = λ.GetItem(τmp2, λ.IntLiteral(2))
						ϒq_url = λ.Cal(ϒget_furl, ϒq)
						if !λ.IsTrue(ϒq_url) {
							continue
						}
						λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
							"format_id": ϒq,
							"url":       ϒq_url,
							"width":     ϒw,
							"height":    ϒh,
						}))
					}
					if !λ.IsTrue(ϒformats) {
						ϒfurl = func() λ.Object {
							if λv := λ.Cal(ϒget_furl, λ.StrLiteral("player")); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.GetItem(λ.GetAttr(ϒcontent, "attrib", nil), λ.StrLiteral("url"))
							}
						}()
						ϒext = λ.Cal(ϒdetermine_ext, ϒfurl)
						ϒformats = λ.NewList(λ.DictLiteral(map[string]λ.Object{
							"url": ϒfurl,
							"vcodec": func() λ.Object {
								if λ.IsTrue(λ.Eq(ϒext, λ.StrLiteral("mp3"))) {
									return λ.StrLiteral("none")
								} else {
									return λ.None
								}
							}(),
							"ext": ϒext,
						}))
					}
					ϒthumbnails = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(ϒcontent, "findall", λ.Cal(ϒmedia_ns_xpath, λ.StrLiteral("thumbnail"))))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒthumbnail = τmp1
						ϒthumbnail_url = λ.Calm(ϒthumbnail, "get", λ.StrLiteral("url"))
						if !λ.IsTrue(ϒthumbnail_url) {
							continue
						}
						λ.Calm(ϒthumbnails, "append", λ.DictLiteral(map[string]λ.Object{
							"url":    ϒthumbnail_url,
							"height": λ.Cal(ϒint_or_none, λ.Calm(ϒthumbnail, "get", λ.StrLiteral("height"))),
							"width":  λ.Cal(ϒint_or_none, λ.Calm(ϒthumbnail, "get", λ.StrLiteral("width"))),
						}))
					}
					return λ.DictLiteral(map[string]λ.Object{
						"id":          ϒvideo_id,
						"formats":     ϒformats,
						"title":       ϒtitle,
						"description": λ.Cal(ϒstrip_or_none, λ.Cal(ϒxpath_text, ϒitem, λ.StrLiteral("description"))),
						"thumbnails":  ϒthumbnails,
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL":    InaIE__VALID_URL,
				"_real_extract": InaIE__real_extract,
			})
		}())
	})
}
