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
 * bokecc/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/bokecc.py
 */

package bokecc

import (
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
		BokeCCBaseIE = λ.Cal(λ.TypeType, λ.NewStr("BokeCCBaseIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
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
					ϒplayer_params_str = λ.Cal(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewStr("<(?:script|embed)[^>]+src=\"http://p\\.bokecc\\.com/player\\?([^\"]+)"), ϒwebpage, λ.NewStr("player params"))
					ϒplayer_params = λ.Cal(ϒcompat_parse_qs, ϒplayer_params_str)
					ϒinfo_xml = λ.Cal(λ.GetAttr(ϒself, "_download_xml", nil), λ.Mod(λ.NewStr("http://p.bokecc.com/servlet/playinfo?uid=%s&vid=%s&m=1"), λ.NewTuple(
						λ.GetItem(λ.GetItem(ϒplayer_params, λ.NewStr("siteid")), λ.NewInt(0)),
						λ.GetItem(λ.GetItem(ϒplayer_params, λ.NewStr("vid")), λ.NewInt(0)),
					)), ϒvideo_id)
					ϒformats = λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
						nil,
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							return λ.NewGenerator(func(λgen λ.Generator) λ.Object {
								var (
									ϒquality λ.Object
									τmp0     λ.Object
									τmp1     λ.Object
								)
								τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(ϒinfo_xml, "findall", nil), λ.NewStr("./video/quality")))
								for {
									if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
										break
									}
									ϒquality = τmp1
									λgen.Yield(λ.NewDictWithTable(map[λ.Object]λ.Object{
										λ.NewStr("format_id"):  ϒformat_id,
										λ.NewStr("url"):        λ.GetItem(λ.GetAttr(λ.Cal(λ.GetAttr(ϒquality, "find", nil), λ.NewStr("./copy")), "attrib", nil), λ.NewStr("playurl")),
										λ.NewStr("preference"): λ.Cal(λ.IntType, λ.GetItem(λ.GetAttr(ϒquality, "attrib", nil), λ.NewStr("value"))),
									}))
								}
								return λ.None
							})
						})))
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					return ϒformats
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_extract_bokecc_formats"): BokeCCBaseIE__extract_bokecc_formats,
			})
		}())
		BokeCCIE = λ.Cal(λ.TypeType, λ.NewStr("BokeCCIE"), λ.NewTuple(BokeCCBaseIE), func() λ.Dict {
			var (
				BokeCCIE__VALID_URL λ.Object
			)
			BokeCCIE__VALID_URL = λ.NewStr("https?://union\\.bokecc\\.com/playvideo\\.bo\\?(?P<query>.*)")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): BokeCCIE__VALID_URL,
			})
		}())
	})
}
