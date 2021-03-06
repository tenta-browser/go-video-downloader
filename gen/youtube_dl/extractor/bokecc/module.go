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
 * bokecc/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/bokecc.py
 */

package bokecc

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	BokeCCBaseIE     λ.Object
	BokeCCIE         λ.Object
	ExtractorError   λ.Object
	InfoExtractor    λ.Object
	ϒcompat_parse_qs λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_parse_qs = Ωcompat.ϒcompat_parse_qs
		ExtractorError = Ωutils.ExtractorError
		BokeCCBaseIE = λ.Cal(λ.TypeType, λ.StrLiteral("BokeCCBaseIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				BokeCCBaseIE__extract_bokecc_formats λ.Object
			)
			BokeCCBaseIE__extract_bokecc_formats = λ.NewFunction("_extract_bokecc_formats",
				[]λ.Param{
					{Name: "self"},
					{Name: "webpage"},
					{Name: "video_id"},
					{Name: "format_id", Def: λ.None},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒformat_id         = λargs[3]
						ϒformats           λ.Object
						ϒinfo_xml          λ.Object
						ϒplayer_params     λ.Object
						ϒplayer_params_str λ.Object
						ϒself              = λargs[0]
						ϒvideo_id          = λargs[2]
						ϒwebpage           = λargs[1]
					)
					ϒplayer_params_str = λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("<(?:script|embed)[^>]+src=(?P<q>[\"\\'])(?:https?:)?//p\\.bokecc\\.com/(?:player|flash/player\\.swf)\\?(?P<query>.+?)(?P=q)"),
						ϒwebpage,
						λ.StrLiteral("player params"),
					), λ.KWArgs{
						{Name: "group", Value: λ.StrLiteral("query")},
					})
					ϒplayer_params = λ.Cal(ϒcompat_parse_qs, ϒplayer_params_str)
					ϒinfo_xml = λ.Calm(ϒself, "_download_xml", λ.Mod(λ.StrLiteral("http://p.bokecc.com/servlet/playinfo?uid=%s&vid=%s&m=1"), λ.NewTuple(
						λ.GetItem(λ.GetItem(ϒplayer_params, λ.StrLiteral("siteid")), λ.IntLiteral(0)),
						λ.GetItem(λ.GetItem(ϒplayer_params, λ.StrLiteral("vid")), λ.IntLiteral(0)),
					)), ϒvideo_id)
					ϒformats = λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
						nil,
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							return λ.NewGenerator(func(λgy λ.Yielder) λ.Object {
								var (
									ϒquality λ.Object
									τmp0     λ.Object
									τmp1     λ.Object
								)
								τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(ϒinfo_xml, "findall", λ.StrLiteral("./video/quality")))
								for {
									if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
										break
									}
									ϒquality = τmp1
									λgy.Yield(λ.DictLiteral(map[string]λ.Object{
										"format_id":  ϒformat_id,
										"url":        λ.GetItem(λ.GetAttr(λ.Calm(ϒquality, "find", λ.StrLiteral("./copy")), "attrib", nil), λ.StrLiteral("playurl")),
										"preference": λ.Cal(λ.IntType, λ.GetItem(λ.GetAttr(ϒquality, "attrib", nil), λ.StrLiteral("value"))),
									}))
								}
								return λ.None
							})
						})))
					λ.Calm(ϒself, "_sort_formats", ϒformats)
					return ϒformats
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_extract_bokecc_formats": BokeCCBaseIE__extract_bokecc_formats,
			})
		}())
		BokeCCIE = λ.Cal(λ.TypeType, λ.StrLiteral("BokeCCIE"), λ.NewTuple(BokeCCBaseIE), func() λ.Dict {
			var (
				BokeCCIE__VALID_URL    λ.Object
				BokeCCIE__real_extract λ.Object
			)
			BokeCCIE__VALID_URL = λ.StrLiteral("https?://union\\.bokecc\\.com/playvideo\\.bo\\?(?P<query>.*)")
			BokeCCIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒqs       λ.Object
						ϒself     = λargs[0]
						ϒurl      = λargs[1]
						ϒvideo_id λ.Object
						ϒwebpage  λ.Object
					)
					ϒqs = λ.Cal(ϒcompat_parse_qs, λ.Calm(λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl), "group", λ.StrLiteral("query")))
					if λ.IsTrue(func() λ.Object {
						if λv := λ.NewBool(!λ.IsTrue(λ.Calm(ϒqs, "get", λ.StrLiteral("vid")))); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.NewBool(!λ.IsTrue(λ.Calm(ϒqs, "get", λ.StrLiteral("uid"))))
						}
					}()) {
						panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.StrLiteral("Invalid URL")), λ.KWArgs{
							{Name: "expected", Value: λ.True},
						})))
					}
					ϒvideo_id = λ.Mod(λ.StrLiteral("%s_%s"), λ.NewTuple(
						λ.GetItem(λ.GetItem(ϒqs, λ.StrLiteral("uid")), λ.IntLiteral(0)),
						λ.GetItem(λ.GetItem(ϒqs, λ.StrLiteral("vid")), λ.IntLiteral(0)),
					))
					ϒwebpage = λ.Calm(ϒself, "_download_webpage", ϒurl, ϒvideo_id)
					return λ.DictLiteral(map[string]λ.Object{
						"id":      ϒvideo_id,
						"title":   λ.StrLiteral("BokeCC Video"),
						"formats": λ.Calm(ϒself, "_extract_bokecc_formats", ϒwebpage, ϒvideo_id),
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL":    BokeCCIE__VALID_URL,
				"_real_extract": BokeCCIE__real_extract,
			})
		}())
	})
}
