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
 * tele13/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/tele13.py
 */

package tele13

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωyoutube "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/youtube"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor  λ.Object
	Tele13IE       λ.Object
	YoutubeIE      λ.Object
	ϒdetermine_ext λ.Object
	ϒjs_to_json    λ.Object
	ϒqualities     λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		YoutubeIE = Ωyoutube.YoutubeIE
		ϒjs_to_json = Ωutils.ϒjs_to_json
		ϒqualities = Ωutils.ϒqualities
		ϒdetermine_ext = Ωutils.ϒdetermine_ext
		Tele13IE = λ.Cal(λ.TypeType, λ.NewStr("Tele13IE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				Tele13IE__TESTS        λ.Object
				Tele13IE__VALID_URL    λ.Object
				Tele13IE__real_extract λ.Object
			)
			Tele13IE__VALID_URL = λ.NewStr("^https?://(?:www\\.)?t13\\.cl/videos(?:/[^/]+)+/(?P<id>[\\w-]+)")
			Tele13IE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.t13.cl/videos/actualidad/el-circulo-de-hierro-de-michelle-bachelet-en-su-regreso-a-la-moneda"),
					λ.NewStr("md5"): λ.NewStr("4cb1fa38adcad8fea88487a078831755"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):    λ.NewStr("el-circulo-de-hierro-de-michelle-bachelet-en-su-regreso-a-la-moneda"),
						λ.NewStr("ext"):   λ.NewStr("mp4"),
						λ.NewStr("title"): λ.NewStr("El círculo de hierro de Michelle Bachelet en su regreso a La Moneda"),
					}),
					λ.NewStr("params"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("skip_download"): λ.True,
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.t13.cl/videos/mundo/tendencias/video-captan-misteriosa-bola-fuego-cielos-bangkok"),
					λ.NewStr("md5"): λ.NewStr("867adf6a3b3fef932c68a71d70b70946"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("rOoKv2OMpOw"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Shooting star seen on 7-Sep-2015"),
						λ.NewStr("description"): λ.NewStr("md5:7292ff2a34b2f673da77da222ae77e1e"),
						λ.NewStr("uploader"):    λ.NewStr("Porjai Jaturongkhakun"),
						λ.NewStr("upload_date"): λ.NewStr("20150906"),
						λ.NewStr("uploader_id"): λ.NewStr("UCnLY_3ezwNcDSC_Wc6suZxw"),
					}),
					λ.NewStr("add_ie"): λ.NewList(λ.NewStr("Youtube")),
				}),
			)
			Tele13IE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdisplay_id λ.Object
						ϒext        λ.Object
						ϒf          λ.Object
						ϒformat_url λ.Object
						ϒformats    λ.Object
						ϒpreference λ.Object
						ϒself       = λargs[0]
						ϒsetup_js   λ.Object
						ϒsources    λ.Object
						ϒurl        = λargs[1]
						ϒurls       λ.Object
						ϒwebpage    λ.Object
						τmp0        λ.Object
						τmp1        λ.Object
					)
					ϒdisplay_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒdisplay_id)
					ϒsetup_js = λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewStr("(?s)jwplayer\\('player-vivo'\\).setup\\((\\{.*?\\})\\)"), ϒwebpage, λ.NewStr("setup code"))
					ϒsources = λ.Cal(λ.GetAttr(ϒself, "_parse_json", nil), λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewStr("sources\\s*:\\s*(\\[[^\\]]+\\])"), ϒsetup_js, λ.NewStr("sources")), ϒdisplay_id, ϒjs_to_json)
					ϒpreference = λ.Cal(ϒqualities, λ.NewList(
						λ.NewStr("Móvil"),
						λ.NewStr("SD"),
						λ.NewStr("HD"),
					))
					ϒformats = λ.NewList()
					ϒurls = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, ϒsources)
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒf = τmp1
						ϒformat_url = λ.GetItem(ϒf, λ.NewStr("file"))
						if λ.IsTrue(func() λ.Object {
							if λv := ϒformat_url; !λ.IsTrue(λv) {
								return λv
							} else {
								return λ.NewBool(!λ.Contains(ϒurls, ϒformat_url))
							}
						}()) {
							ϒext = λ.Cal(ϒdetermine_ext, ϒformat_url)
							if λ.IsTrue(λ.Eq(ϒext, λ.NewStr("m3u8"))) {
								λ.Cal(λ.GetAttr(ϒformats, "extend", nil), λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
									ϒformat_url,
									ϒdisplay_id,
									λ.NewStr("mp4"),
									λ.NewStr("m3u8_native"),
								), λ.KWArgs{
									{Name: "m3u8_id", Value: λ.NewStr("hls")},
									{Name: "fatal", Value: λ.False},
								}))
							} else {
								if λ.IsTrue(λ.Cal(λ.GetAttr(YoutubeIE, "suitable", nil), ϒformat_url)) {
									return λ.Cal(λ.GetAttr(ϒself, "url_result", nil), ϒformat_url, λ.NewStr("Youtube"))
								} else {
									λ.Cal(λ.GetAttr(ϒformats, "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
										λ.NewStr("url"):        ϒformat_url,
										λ.NewStr("format_id"):  λ.Cal(λ.GetAttr(ϒf, "get", nil), λ.NewStr("label")),
										λ.NewStr("preference"): λ.Cal(ϒpreference, λ.Cal(λ.GetAttr(ϒf, "get", nil), λ.NewStr("label"))),
										λ.NewStr("ext"):        ϒext,
									}))
								}
							}
							λ.Cal(λ.GetAttr(ϒurls, "append", nil), ϒformat_url)
						}
					}
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          ϒdisplay_id,
						λ.NewStr("title"):       λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewStr("title\\s*:\\s*\"([^\"]+)\""), ϒsetup_js, λ.NewStr("title")),
						λ.NewStr("description"): λ.Cal(λ.GetAttr(ϒself, "_html_search_meta", nil), λ.NewStr("description"), ϒwebpage, λ.NewStr("description")),
						λ.NewStr("thumbnail"): λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
							λ.NewStr("image\\s*:\\s*\"([^\"]+)\""),
							ϒsetup_js,
							λ.NewStr("thumbnail"),
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
						}),
						λ.NewStr("formats"): ϒformats,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TESTS"):        Tele13IE__TESTS,
				λ.NewStr("_VALID_URL"):    Tele13IE__VALID_URL,
				λ.NewStr("_real_extract"): Tele13IE__real_extract,
			})
		}())
	})
}
