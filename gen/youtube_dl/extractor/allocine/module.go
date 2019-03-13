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
 * allocine/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/allocine.py
 */

package allocine

import (
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	AllocineIE         λ.Object
	InfoExtractor      λ.Object
	ϒcompat_str        λ.Object
	ϒint_or_none       λ.Object
	ϒqualities         λ.Object
	ϒremove_end        λ.Object
	ϒtry_get           λ.Object
	ϒunified_timestamp λ.Object
	ϒurl_basename      λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_str = Ωcompat.ϒcompat_str
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒqualities = Ωutils.ϒqualities
		ϒremove_end = Ωutils.ϒremove_end
		ϒtry_get = Ωutils.ϒtry_get
		ϒunified_timestamp = Ωutils.ϒunified_timestamp
		ϒurl_basename = Ωutils.ϒurl_basename
		AllocineIE = λ.Cal(λ.TypeType, λ.NewStr("AllocineIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				AllocineIE__TESTS        λ.Object
				AllocineIE__VALID_URL    λ.Object
				AllocineIE__real_extract λ.Object
			)
			AllocineIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?allocine\\.fr/(?:article|video|film)/(?:fichearticle_gen_carticle=|player_gen_cmedia=|fichefilm_gen_cfilm=|video-)(?P<id>[0-9]+)(?:\\.html)?")
			AllocineIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.allocine.fr/article/fichearticle_gen_carticle=18635087.html"),
					λ.NewStr("md5"): λ.NewStr("0c9fcf59a841f65635fa300ac43d8269"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("19546517"),
						λ.NewStr("display_id"):  λ.NewStr("18635087"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Astérix - Le Domaine des Dieux Teaser VF"),
						λ.NewStr("description"): λ.NewStr("md5:4a754271d9c6f16c72629a8a993ee884"),
						λ.NewStr("thumbnail"):   λ.NewStr("re:http://.*\\.jpg"),
						λ.NewStr("duration"):    λ.NewInt(39),
						λ.NewStr("timestamp"):   λ.NewInt(1404273600),
						λ.NewStr("upload_date"): λ.NewStr("20140702"),
						λ.NewStr("view_count"):  λ.IntType,
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.allocine.fr/video/player_gen_cmedia=19540403&cfilm=222257.html"),
					λ.NewStr("md5"): λ.NewStr("d0cdce5d2b9522ce279fdfec07ff16e0"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("19540403"),
						λ.NewStr("display_id"):  λ.NewStr("19540403"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Planes 2 Bande-annonce VF"),
						λ.NewStr("description"): λ.NewStr("Regardez la bande annonce du film Planes 2 (Planes 2 Bande-annonce VF). Planes 2, un film de Roberts Gannaway"),
						λ.NewStr("thumbnail"):   λ.NewStr("re:http://.*\\.jpg"),
						λ.NewStr("duration"):    λ.NewInt(69),
						λ.NewStr("timestamp"):   λ.NewInt(1385659800),
						λ.NewStr("upload_date"): λ.NewStr("20131128"),
						λ.NewStr("view_count"):  λ.IntType,
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.allocine.fr/video/player_gen_cmedia=19544709&cfilm=181290.html"),
					λ.NewStr("md5"): λ.NewStr("101250fb127ef9ca3d73186ff22a47ce"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("19544709"),
						λ.NewStr("display_id"):  λ.NewStr("19544709"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Dragons 2 - Bande annonce finale VF"),
						λ.NewStr("description"): λ.NewStr("md5:6cdd2d7c2687d4c6aafe80a35e17267a"),
						λ.NewStr("thumbnail"):   λ.NewStr("re:http://.*\\.jpg"),
						λ.NewStr("duration"):    λ.NewInt(144),
						λ.NewStr("timestamp"):   λ.NewInt(1397589900),
						λ.NewStr("upload_date"): λ.NewStr("20140415"),
						λ.NewStr("view_count"):  λ.IntType,
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.allocine.fr/video/video-19550147/"),
					λ.NewStr("md5"): λ.NewStr("3566c0668c0235e2d224fd8edb389f67"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("19550147"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Faux Raccord N°123 - Les gaffes de Cliffhanger"),
						λ.NewStr("description"): λ.NewStr("md5:bc734b83ffa2d8a12188d9eb48bb6354"),
						λ.NewStr("thumbnail"):   λ.NewStr("re:http://.*\\.jpg"),
					}),
				}),
			)
			AllocineIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdisplay_id λ.Object
						ϒduration   λ.Object
						ϒformat_id  λ.Object
						ϒformats    λ.Object
						ϒkey        λ.Object
						ϒmedia_data λ.Object
						ϒmodel      λ.Object
						ϒmodel_data λ.Object
						ϒquality    λ.Object
						ϒself       = λargs[0]
						ϒtimestamp  λ.Object
						ϒtitle      λ.Object
						ϒurl        = λargs[1]
						ϒvalue      λ.Object
						ϒvideo      λ.Object
						ϒvideo_id   λ.Object
						ϒvideo_url  λ.Object
						ϒview_count λ.Object
						ϒwebpage    λ.Object
						τmp0        λ.Object
						τmp1        λ.Object
						τmp2        λ.Object
					)
					ϒdisplay_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒdisplay_id)
					ϒformats = λ.NewList()
					ϒquality = λ.Cal(ϒqualities, λ.NewList(
						λ.NewStr("ld"),
						λ.NewStr("md"),
						λ.NewStr("hd"),
					))
					ϒmodel = λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
						λ.NewStr("data-model=\"([^\"]+)\""),
						ϒwebpage,
						λ.NewStr("data model"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					if λ.IsTrue(ϒmodel) {
						ϒmodel_data = λ.Cal(λ.GetAttr(ϒself, "_parse_json", nil), ϒmodel, ϒdisplay_id)
						ϒvideo = λ.GetItem(λ.GetItem(ϒmodel_data, λ.NewStr("videos")), λ.NewInt(0))
						ϒtitle = λ.GetItem(ϒvideo, λ.NewStr("title"))
						τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(λ.GetItem(ϒvideo, λ.NewStr("sources")), "values", nil)))
						for {
							if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
								break
							}
							ϒvideo_url = τmp1
							τmp2 = λ.GetItem(λ.Cal(λ.GetAttr(λ.Cal(ϒurl_basename, ϒvideo_url), "split", nil), λ.NewStr("_")), λ.NewSlice(λ.None, λ.NewInt(2), λ.None))
							ϒvideo_id = λ.GetItem(τmp2, λ.NewInt(0))
							ϒformat_id = λ.GetItem(τmp2, λ.NewInt(1))
							λ.Cal(λ.GetAttr(ϒformats, "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
								λ.NewStr("format_id"): ϒformat_id,
								λ.NewStr("quality"):   λ.Cal(ϒquality, ϒformat_id),
								λ.NewStr("url"):       ϒvideo_url,
							}))
						}
						ϒduration = λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒvideo, "get", nil), λ.NewStr("duration")))
						ϒview_count = λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒvideo, "get", nil), λ.NewStr("view_count")))
						ϒtimestamp = λ.Cal(ϒunified_timestamp, λ.Cal(ϒtry_get, ϒvideo, λ.NewFunction("<lambda>",
							[]λ.Param{
								{Name: "x"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒx = λargs[0]
								)
								return λ.GetItem(λ.GetItem(ϒx, λ.NewStr("added_at")), λ.NewStr("date"))
							}), ϒcompat_str))
					} else {
						ϒvideo_id = ϒdisplay_id
						ϒmedia_data = λ.Cal(λ.GetAttr(ϒself, "_download_json", nil), λ.Mod(λ.NewStr("http://www.allocine.fr/ws/AcVisiondataV5.ashx?media=%s"), ϒvideo_id), ϒdisplay_id)
						ϒtitle = λ.Cal(ϒremove_end, λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewStr("(?s)<title>(.+?)</title>"), ϒwebpage, λ.NewStr("title")), "strip", nil)), λ.NewStr(" - AlloCiné"))
						τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(λ.GetItem(ϒmedia_data, λ.NewStr("video")), "items", nil)))
						for {
							if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
								break
							}
							τmp2 = τmp1
							ϒkey = λ.GetItem(τmp2, λ.NewInt(0))
							ϒvalue = λ.GetItem(τmp2, λ.NewInt(1))
							if λ.IsTrue(λ.NewBool(!λ.IsTrue(λ.Cal(λ.GetAttr(ϒkey, "endswith", nil), λ.NewStr("Path"))))) {
								continue
							}
							ϒformat_id = λ.GetItem(ϒkey, λ.NewSlice(λ.None, λ.Neg(λ.Cal(λ.BuiltinLen, λ.NewStr("Path"))), λ.None))
							λ.Cal(λ.GetAttr(ϒformats, "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
								λ.NewStr("format_id"): ϒformat_id,
								λ.NewStr("quality"):   λ.Cal(ϒquality, ϒformat_id),
								λ.NewStr("url"):       ϒvalue,
							}))
						}
						τmp0 = λ.Mul(λ.NewList(λ.None), λ.NewInt(3))
						ϒduration = λ.GetItem(τmp0, λ.NewInt(0))
						ϒview_count = λ.GetItem(τmp0, λ.NewInt(1))
						ϒtimestamp = λ.GetItem(τmp0, λ.NewInt(2))
					}
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          ϒvideo_id,
						λ.NewStr("display_id"):  ϒdisplay_id,
						λ.NewStr("title"):       ϒtitle,
						λ.NewStr("description"): λ.Cal(λ.GetAttr(ϒself, "_og_search_description", nil), ϒwebpage),
						λ.NewStr("thumbnail"):   λ.Cal(λ.GetAttr(ϒself, "_og_search_thumbnail", nil), ϒwebpage),
						λ.NewStr("duration"):    ϒduration,
						λ.NewStr("timestamp"):   ϒtimestamp,
						λ.NewStr("view_count"):  ϒview_count,
						λ.NewStr("formats"):     ϒformats,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TESTS"):        AllocineIE__TESTS,
				λ.NewStr("_VALID_URL"):    AllocineIE__VALID_URL,
				λ.NewStr("_real_extract"): AllocineIE__real_extract,
			})
		}())
	})
}
