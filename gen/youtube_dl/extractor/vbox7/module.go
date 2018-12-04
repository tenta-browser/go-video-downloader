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
 * vbox7/module.go: transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/vbox7.py
 */

package vbox7

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError λ.Object
	InfoExtractor  λ.Object
	Vbox7IE        λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ExtractorError = Ωutils.ExtractorError
		Vbox7IE = λ.Cal(λ.TypeType, λ.NewStr("Vbox7IE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				Vbox7IE__GEO_COUNTRIES λ.Object
				Vbox7IE__TESTS         λ.Object
				Vbox7IE__VALID_URL     λ.Object
				Vbox7IE__real_extract  λ.Object
			)
			Vbox7IE__VALID_URL = λ.NewStr("(?x)\n                    https?://\n                        (?:[^/]+\\.)?vbox7\\.com/\n                        (?:\n                            play:|\n                            (?:\n                                emb/external\\.php|\n                                player/ext\\.swf\n                            )\\?.*?\\bvid=\n                        )\n                        (?P<id>[\\da-fA-F]+)\n                    ")
			Vbox7IE__GEO_COUNTRIES = λ.NewList(λ.NewStr("BG"))
			Vbox7IE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://vbox7.com/play:0946fff23c"),
					λ.NewStr("md5"): λ.NewStr("a60f9ab3a3a2f013ef9a967d5f7be5bf"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("0946fff23c"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Борисов: Притеснен съм за бъдещето на България"),
						λ.NewStr("description"): λ.NewStr("По думите му е опасно страната ни да бъде обявена за \"сигурна\""),
						λ.NewStr("thumbnail"):   λ.NewStr("re:^https?://.*\\.jpg$"),
						λ.NewStr("timestamp"):   λ.NewInt(1470982814),
						λ.NewStr("upload_date"): λ.NewStr("20160812"),
						λ.NewStr("uploader"):    λ.NewStr("zdraveibulgaria"),
					}),
					λ.NewStr("params"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("proxy"): λ.NewStr("127.0.0.1:8118"),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://vbox7.com/play:249bb972c2"),
					λ.NewStr("md5"): λ.NewStr("99f65c0c9ef9b682b97313e052734c3f"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):    λ.NewStr("249bb972c2"),
						λ.NewStr("ext"):   λ.NewStr("mp4"),
						λ.NewStr("title"): λ.NewStr("Смях! Чудо - чист за секунди - Скрита камера"),
					}),
					λ.NewStr("skip"): λ.NewStr("georestricted"),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://vbox7.com/emb/external.php?vid=a240d20f9c&autoplay=1"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://i49.vbox7.com/player/ext.swf?vid=0946fff23c&autoplay=1"),
					λ.NewStr("only_matching"): λ.True,
				}),
			)
			Vbox7IE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒinfo      λ.Object
						ϒresponse  λ.Object
						ϒself      = λargs[0]
						ϒtitle     λ.Object
						ϒuploader  λ.Object
						ϒurl       = λargs[1]
						ϒvideo     λ.Object
						ϒvideo_id  λ.Object
						ϒvideo_url λ.Object
						ϒwebpage   λ.Object
					)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒresponse = λ.Cal(λ.GetAttr(ϒself, "_download_json", nil), λ.Mod(λ.NewStr("https://www.vbox7.com/ajax/video/nextvideo.php?vid=%s"), ϒvideo_id), ϒvideo_id)
					if λ.IsTrue(λ.NewBool(λ.Contains(ϒresponse, λ.NewStr("error")))) {
						panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.Mod(λ.NewStr("%s said: %s"), λ.NewTuple(
							λ.GetAttr(ϒself, "IE_NAME", nil),
							λ.GetItem(ϒresponse, λ.NewStr("error")),
						))), λ.KWArgs{
							{Name: "expected", Value: λ.True},
						})))
					}
					ϒvideo = λ.GetItem(ϒresponse, λ.NewStr("options"))
					ϒtitle = λ.GetItem(ϒvideo, λ.NewStr("title"))
					ϒvideo_url = λ.GetItem(ϒvideo, λ.NewStr("src"))
					if λ.IsTrue(λ.NewBool(λ.Contains(ϒvideo_url, λ.NewStr("/na.mp4")))) {
						λ.Call(λ.GetAttr(ϒself, "raise_geo_restricted", nil), nil, λ.KWArgs{
							{Name: "countries", Value: λ.GetAttr(ϒself, "_GEO_COUNTRIES", nil)},
						})
					}
					ϒuploader = λ.Cal(λ.GetAttr(ϒvideo, "get", nil), λ.NewStr("uploader"))
					ϒwebpage = λ.Call(λ.GetAttr(ϒself, "_download_webpage", nil), λ.NewArgs(
						λ.Mod(λ.NewStr("http://vbox7.com/play:%s"), ϒvideo_id),
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "fatal", Value: λ.None},
					})
					ϒinfo = λ.NewDictWithTable(map[λ.Object]λ.Object{})
					if λ.IsTrue(ϒwebpage) {
						ϒinfo = λ.Call(λ.GetAttr(ϒself, "_search_json_ld", nil), λ.NewArgs(
							λ.Cal(λ.GetAttr(ϒwebpage, "replace", nil), λ.NewStr("\"/*@context\""), λ.NewStr("\"@context\"")),
							ϒvideo_id,
						), λ.KWArgs{
							{Name: "fatal", Value: λ.False},
						})
					}
					λ.Cal(λ.GetAttr(ϒinfo, "update", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):       ϒvideo_id,
						λ.NewStr("title"):    ϒtitle,
						λ.NewStr("url"):      ϒvideo_url,
						λ.NewStr("uploader"): ϒuploader,
						λ.NewStr("thumbnail"): λ.Cal(λ.GetAttr(ϒself, "_proto_relative_url", nil), func() λ.Object {
							if λv := λ.Cal(λ.GetAttr(ϒinfo, "get", nil), λ.NewStr("thumbnail")); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Cal(λ.GetAttr(ϒself, "_og_search_thumbnail", nil), ϒwebpage)
							}
						}(), λ.NewStr("http:")),
					}))
					return ϒinfo
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_GEO_COUNTRIES"): Vbox7IE__GEO_COUNTRIES,
				λ.NewStr("_TESTS"):         Vbox7IE__TESTS,
				λ.NewStr("_VALID_URL"):     Vbox7IE__VALID_URL,
				λ.NewStr("_real_extract"):  Vbox7IE__real_extract,
			})
		}())
	})
}