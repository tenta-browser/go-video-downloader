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
 * beatport/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/beatport.py
 */

package beatport

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	BeatportIE    λ.Object
	InfoExtractor λ.Object
	ϒcompat_str   λ.Object
	ϒint_or_none  λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_str = Ωcompat.ϒcompat_str
		ϒint_or_none = Ωutils.ϒint_or_none
		BeatportIE = λ.Cal(λ.TypeType, λ.StrLiteral("BeatportIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				BeatportIE__VALID_URL    λ.Object
				BeatportIE__real_extract λ.Object
			)
			BeatportIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.|pro\\.)?beatport\\.com/track/(?P<display_id>[^/]+)/(?P<id>[0-9]+)")
			BeatportIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdisplay_id λ.Object
						ϒext        λ.Object
						ϒfmt        λ.Object
						ϒformats    λ.Object
						ϒimage      λ.Object
						ϒimage_url  λ.Object
						ϒimages     λ.Object
						ϒinfo       λ.Object
						ϒmobj       λ.Object
						ϒname       λ.Object
						ϒplayables  λ.Object
						ϒself       = λargs[0]
						ϒtitle      λ.Object
						ϒtrack      λ.Object
						ϒtrack_id   λ.Object
						ϒurl        = λargs[1]
						ϒwebpage    λ.Object
						τmp0        λ.Object
						τmp1        λ.Object
						τmp2        λ.Object
					)
					ϒmobj = λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl)
					ϒtrack_id = λ.Calm(ϒmobj, "group", λ.StrLiteral("id"))
					ϒdisplay_id = λ.Calm(ϒmobj, "group", λ.StrLiteral("display_id"))
					ϒwebpage = λ.Calm(ϒself, "_download_webpage", ϒurl, ϒdisplay_id)
					ϒplayables = λ.Calm(ϒself, "_parse_json", λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("window\\.Playables\\s*=\\s*({.+?});"),
						ϒwebpage,
						λ.StrLiteral("playables info"),
					), λ.KWArgs{
						{Name: "flags", Value: Ωre.DOTALL},
					}), ϒtrack_id)
					ϒtrack = λ.Cal(λ.BuiltinNext, λ.Cal(λ.NewFunction("<generator>",
						nil,
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							return λ.NewGenerator(func(λgy λ.Yielder) λ.Object {
								var (
									ϒt   λ.Object
									τmp0 λ.Object
									τmp1 λ.Object
								)
								τmp0 = λ.Cal(λ.BuiltinIter, λ.GetItem(ϒplayables, λ.StrLiteral("tracks")))
								for {
									if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
										break
									}
									ϒt = τmp1
									if λ.IsTrue(λ.Eq(λ.GetItem(ϒt, λ.StrLiteral("id")), λ.Cal(λ.IntType, ϒtrack_id))) {
										λgy.Yield(ϒt)
									}
								}
								return λ.None
							})
						})))
					ϒtitle = λ.Add(λ.Add(λ.Calm(λ.StrLiteral(", "), "join", λ.Cal(λ.NewFunction("<generator>",
						nil,
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							return λ.NewGenerator(func(λgy λ.Yielder) λ.Object {
								var (
									ϒa   λ.Object
									τmp0 λ.Object
									τmp1 λ.Object
								)
								τmp0 = λ.Cal(λ.BuiltinIter, λ.GetItem(ϒtrack, λ.StrLiteral("artists")))
								for {
									if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
										break
									}
									ϒa = τmp1
									λgy.Yield(λ.GetItem(ϒa, λ.StrLiteral("name")))
								}
								return λ.None
							})
						}))), λ.StrLiteral(" - ")), λ.GetItem(ϒtrack, λ.StrLiteral("name")))
					if λ.IsTrue(λ.GetItem(ϒtrack, λ.StrLiteral("mix"))) {
						τmp0 = λ.IAdd(ϒtitle, λ.Add(λ.Add(λ.StrLiteral(" ("), λ.GetItem(ϒtrack, λ.StrLiteral("mix"))), λ.StrLiteral(")")))
						ϒtitle = τmp0
					}
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(λ.GetItem(ϒtrack, λ.StrLiteral("preview")), "items"))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						τmp2 = λ.UnpackIterable(τmp1, 2)
						ϒext = λ.GetItem(τmp2, λ.IntLiteral(0))
						ϒinfo = λ.GetItem(τmp2, λ.IntLiteral(1))
						if !λ.IsTrue(λ.GetItem(ϒinfo, λ.StrLiteral("url"))) {
							continue
						}
						ϒfmt = λ.DictLiteral(map[string]λ.Object{
							"url":       λ.GetItem(ϒinfo, λ.StrLiteral("url")),
							"ext":       ϒext,
							"format_id": ϒext,
							"vcodec":    λ.StrLiteral("none"),
						})
						if λ.IsTrue(λ.Eq(ϒext, λ.StrLiteral("mp3"))) {
							λ.SetItem(ϒfmt, λ.StrLiteral("preference"), λ.IntLiteral(0))
							λ.SetItem(ϒfmt, λ.StrLiteral("acodec"), λ.StrLiteral("mp3"))
							λ.SetItem(ϒfmt, λ.StrLiteral("abr"), λ.IntLiteral(96))
							λ.SetItem(ϒfmt, λ.StrLiteral("asr"), λ.IntLiteral(44100))
						} else {
							if λ.IsTrue(λ.Eq(ϒext, λ.StrLiteral("mp4"))) {
								λ.SetItem(ϒfmt, λ.StrLiteral("preference"), λ.IntLiteral(1))
								λ.SetItem(ϒfmt, λ.StrLiteral("acodec"), λ.StrLiteral("aac"))
								λ.SetItem(ϒfmt, λ.StrLiteral("abr"), λ.IntLiteral(96))
								λ.SetItem(ϒfmt, λ.StrLiteral("asr"), λ.IntLiteral(44100))
							}
						}
						λ.Calm(ϒformats, "append", ϒfmt)
					}
					λ.Calm(ϒself, "_sort_formats", ϒformats)
					ϒimages = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(λ.GetItem(ϒtrack, λ.StrLiteral("images")), "items"))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						τmp2 = λ.UnpackIterable(τmp1, 2)
						ϒname = λ.GetItem(τmp2, λ.IntLiteral(0))
						ϒinfo = λ.GetItem(τmp2, λ.IntLiteral(1))
						ϒimage_url = λ.Calm(ϒinfo, "get", λ.StrLiteral("url"))
						if λ.IsTrue(func() λ.Object {
							if λv := λ.Eq(ϒname, λ.StrLiteral("dynamic")); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.NewBool(!λ.IsTrue(ϒimage_url))
							}
						}()) {
							continue
						}
						ϒimage = λ.DictLiteral(map[string]λ.Object{
							"id":     ϒname,
							"url":    ϒimage_url,
							"height": λ.Cal(ϒint_or_none, λ.Calm(ϒinfo, "get", λ.StrLiteral("height"))),
							"width":  λ.Cal(ϒint_or_none, λ.Calm(ϒinfo, "get", λ.StrLiteral("width"))),
						})
						λ.Calm(ϒimages, "append", ϒimage)
					}
					return λ.DictLiteral(map[string]λ.Object{
						"id": func() λ.Object {
							if λv := λ.Cal(ϒcompat_str, λ.Calm(ϒtrack, "get", λ.StrLiteral("id"))); λ.IsTrue(λv) {
								return λv
							} else {
								return ϒtrack_id
							}
						}(),
						"display_id": func() λ.Object {
							if λv := λ.Calm(ϒtrack, "get", λ.StrLiteral("slug")); λ.IsTrue(λv) {
								return λv
							} else {
								return ϒdisplay_id
							}
						}(),
						"title":      ϒtitle,
						"formats":    ϒformats,
						"thumbnails": ϒimages,
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL":    BeatportIE__VALID_URL,
				"_real_extract": BeatportIE__real_extract,
			})
		}())
	})
}
