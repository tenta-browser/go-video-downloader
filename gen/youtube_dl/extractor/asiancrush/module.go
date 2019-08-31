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
 * asiancrush/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/asiancrush.py
 */

package asiancrush

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωkaltura "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/kaltura"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	AsianCrushIE         λ.Object
	AsianCrushPlaylistIE λ.Object
	InfoExtractor        λ.Object
	KalturaIE            λ.Object
	ϒextract_attributes  λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		KalturaIE = Ωkaltura.KalturaIE
		ϒextract_attributes = Ωutils.ϒextract_attributes
		AsianCrushIE = λ.Cal(λ.TypeType, λ.NewStr("AsianCrushIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				AsianCrushIE__TESTS          λ.Object
				AsianCrushIE__VALID_URL      λ.Object
				AsianCrushIE__VALID_URL_BASE λ.Object
				AsianCrushIE__real_extract   λ.Object
			)
			AsianCrushIE__VALID_URL_BASE = λ.NewStr("https?://(?:www\\.)?(?P<host>(?:(?:asiancrush|yuyutv|midnightpulp)\\.com|cocoro\\.tv))")
			AsianCrushIE__VALID_URL = λ.Mod(λ.NewStr("%s/video/(?:[^/]+/)?0+(?P<id>\\d+)v\\b"), AsianCrushIE__VALID_URL_BASE)
			AsianCrushIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("https://www.asiancrush.com/video/012869v/women-who-flirt/"),
					λ.NewStr("md5"): λ.NewStr("c3b740e48d0ba002a42c0b72857beae6"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("1_y4tmjm5r"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Women Who Flirt"),
						λ.NewStr("description"): λ.NewStr("md5:7e986615808bcfb11756eb503a751487"),
						λ.NewStr("timestamp"):   λ.NewInt(1496936429),
						λ.NewStr("upload_date"): λ.NewStr("20170608"),
						λ.NewStr("uploader_id"): λ.NewStr("craig@crifkin.com"),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://www.asiancrush.com/video/she-was-pretty/011886v-pretty-episode-3/"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://www.yuyutv.com/video/013886v/the-act-of-killing/"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://www.yuyutv.com/video/peep-show/013922v-warring-factions/"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://www.midnightpulp.com/video/010400v/drifters/"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://www.midnightpulp.com/video/mononoke/016378v-zashikiwarashi-part-1/"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://www.cocoro.tv/video/the-wonderful-wizard-of-oz/008878v-the-wonderful-wizard-of-oz-ep01/"),
					λ.NewStr("only_matching"): λ.True,
				}),
			)
			AsianCrushIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdescription λ.Object
						ϒentry_id    λ.Object
						ϒhost        λ.Object
						ϒkaltura_id  λ.Object
						ϒmobj        λ.Object
						ϒpartner_id  λ.Object
						ϒplayer      λ.Object
						ϒself        = λargs[0]
						ϒtitle       λ.Object
						ϒurl         = λargs[1]
						ϒvars        λ.Object
						ϒvideo_id    λ.Object
						ϒwebpage     λ.Object
						τmp0         λ.Object
					)
					ϒmobj = λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl)
					ϒhost = λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("host"))
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("id"))
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒvideo_id)
					τmp0 = λ.Mul(λ.NewList(λ.None), λ.NewInt(3))
					ϒentry_id = λ.GetItem(τmp0, λ.NewInt(0))
					ϒpartner_id = λ.GetItem(τmp0, λ.NewInt(1))
					ϒtitle = λ.GetItem(τmp0, λ.NewInt(2))
					ϒvars = λ.Call(λ.GetAttr(ϒself, "_parse_json", nil), λ.NewArgs(
						λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
							λ.NewStr("iEmbedVars\\s*=\\s*({.+?})"),
							ϒwebpage,
							λ.NewStr("embed vars"),
						), λ.KWArgs{
							{Name: "default", Value: λ.NewStr("{}")},
						}),
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					})
					if λ.IsTrue(ϒvars) {
						ϒentry_id = λ.Cal(λ.GetAttr(ϒvars, "get", nil), λ.NewStr("entry_id"))
						ϒpartner_id = λ.Cal(λ.GetAttr(ϒvars, "get", nil), λ.NewStr("partner_id"))
						ϒtitle = λ.Cal(λ.GetAttr(ϒvars, "get", nil), λ.NewStr("vid_label"))
					}
					if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒentry_id))) {
						ϒentry_id = λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewStr("\\bentry_id[\"\\']\\s*:\\s*[\"\\'](\\d+)"), ϒwebpage, λ.NewStr("entry id"))
					}
					ϒplayer = λ.Call(λ.GetAttr(ϒself, "_download_webpage", nil), λ.NewArgs(
						λ.Mod(λ.NewStr("https://api.%s/embeddedVideoPlayer"), ϒhost),
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "query", Value: λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("id"): ϒentry_id,
						})},
					})
					ϒkaltura_id = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr("entry_id[\"\\']\\s*:\\s*([\"\\'])(?P<id>(?:(?!\\1).)+)\\1"),
						ϒplayer,
						λ.NewStr("kaltura id"),
					), λ.KWArgs{
						{Name: "group", Value: λ.NewStr("id")},
					})
					if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒpartner_id))) {
						ϒpartner_id = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
							λ.NewStr("/p(?:artner_id)?/(\\d+)"),
							ϒplayer,
							λ.NewStr("partner id"),
						), λ.KWArgs{
							{Name: "default", Value: λ.NewStr("513551")},
						})
					}
					ϒdescription = λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
						λ.NewStr("(?s)<div[^>]+\\bclass=[\"\\']description[\"\\'][^>]*>(.+?)</div>"),
						ϒwebpage,
						λ.NewStr("description"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					})
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("_type"): λ.NewStr("url_transparent"),
						λ.NewStr("url"): λ.Mod(λ.NewStr("kaltura:%s:%s"), λ.NewTuple(
							ϒpartner_id,
							ϒkaltura_id,
						)),
						λ.NewStr("ie_key"):      λ.Cal(λ.GetAttr(KalturaIE, "ie_key", nil)),
						λ.NewStr("id"):          ϒvideo_id,
						λ.NewStr("title"):       ϒtitle,
						λ.NewStr("description"): ϒdescription,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TESTS"):          AsianCrushIE__TESTS,
				λ.NewStr("_VALID_URL"):      AsianCrushIE__VALID_URL,
				λ.NewStr("_VALID_URL_BASE"): AsianCrushIE__VALID_URL_BASE,
				λ.NewStr("_real_extract"):   AsianCrushIE__real_extract,
			})
		}())
		AsianCrushPlaylistIE = λ.Cal(λ.TypeType, λ.NewStr("AsianCrushPlaylistIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				AsianCrushPlaylistIE__VALID_URL λ.Object
			)
			AsianCrushPlaylistIE__VALID_URL = λ.Mod(λ.NewStr("%s/series/0+(?P<id>\\d+)s\\b"), λ.GetAttr(AsianCrushIE, "_VALID_URL_BASE", nil))
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): AsianCrushPlaylistIE__VALID_URL,
			})
		}())
	})
}