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
 * sapo/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/sapo.py
 */

package sapo

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor    λ.Object
	SapoIE           λ.Object
	ϒparse_duration  λ.Object
	ϒunified_strdate λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒparse_duration = Ωutils.ϒparse_duration
		ϒunified_strdate = Ωutils.ϒunified_strdate
		SapoIE = λ.Cal(λ.TypeType, λ.NewStr("SapoIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				SapoIE__TESTS        λ.Object
				SapoIE__VALID_URL    λ.Object
				SapoIE__real_extract λ.Object
			)
			SapoIE__VALID_URL = λ.NewStr("https?://(?:(?:v2|www)\\.)?videos\\.sapo\\.(?:pt|cv|ao|mz|tl)/(?P<id>[\\da-zA-Z]{20})")
			SapoIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):  λ.NewStr("http://videos.sapo.pt/UBz95kOtiWYUMTA5Ghfi"),
					λ.NewStr("md5"):  λ.NewStr("79ee523f6ecb9233ac25075dee0eda83"),
					λ.NewStr("note"): λ.NewStr("SD video"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("UBz95kOtiWYUMTA5Ghfi"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Benfica - Marcas na Hitória"),
						λ.NewStr("description"): λ.NewStr("md5:c9082000a128c3fd57bf0299e1367f22"),
						λ.NewStr("duration"):    λ.NewInt(264),
						λ.NewStr("uploader"):    λ.NewStr("tiago_1988"),
						λ.NewStr("upload_date"): λ.NewStr("20080229"),
						λ.NewStr("categories"): λ.NewList(
							λ.NewStr("benfica"),
							λ.NewStr("cabral"),
							λ.NewStr("desporto"),
							λ.NewStr("futebol"),
							λ.NewStr("geovanni"),
							λ.NewStr("hooijdonk"),
							λ.NewStr("joao"),
							λ.NewStr("karel"),
							λ.NewStr("lisboa"),
							λ.NewStr("miccoli"),
						),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):  λ.NewStr("http://videos.sapo.pt/IyusNAZ791ZdoCY5H5IF"),
					λ.NewStr("md5"):  λ.NewStr("90a2f283cfb49193fe06e861613a72aa"),
					λ.NewStr("note"): λ.NewStr("HD video"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("IyusNAZ791ZdoCY5H5IF"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Codebits VII - Report"),
						λ.NewStr("description"): λ.NewStr("md5:6448d6fd81ce86feac05321f354dbdc8"),
						λ.NewStr("duration"):    λ.NewInt(144),
						λ.NewStr("uploader"):    λ.NewStr("codebits"),
						λ.NewStr("upload_date"): λ.NewStr("20140427"),
						λ.NewStr("categories"): λ.NewList(
							λ.NewStr("codebits"),
							λ.NewStr("codebits2014"),
						),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):  λ.NewStr("http://v2.videos.sapo.pt/yLqjzPtbTimsn2wWBKHz"),
					λ.NewStr("md5"):  λ.NewStr("e5aa7cc0bdc6db9b33df1a48e49a15ac"),
					λ.NewStr("note"): λ.NewStr("v2 video"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("yLqjzPtbTimsn2wWBKHz"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Hipnose Condicionativa 4"),
						λ.NewStr("description"): λ.NewStr("md5:ef0481abf8fb4ae6f525088a6dadbc40"),
						λ.NewStr("duration"):    λ.NewInt(692),
						λ.NewStr("uploader"):    λ.NewStr("sapozen"),
						λ.NewStr("upload_date"): λ.NewStr("20090609"),
						λ.NewStr("categories"): λ.NewList(
							λ.NewStr("condicionativa"),
							λ.NewStr("heloisa"),
							λ.NewStr("hipnose"),
							λ.NewStr("miranda"),
							λ.NewStr("sapo"),
							λ.NewStr("zen"),
						),
					}),
				}),
			)
			SapoIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒage_limit     λ.Object
						ϒcategories    λ.Object
						ϒcomment_count λ.Object
						ϒdescription   λ.Object
						ϒduration      λ.Object
						ϒformats       λ.Object
						ϒitem          λ.Object
						ϒmobj          λ.Object
						ϒself          = λargs[0]
						ϒtags          λ.Object
						ϒthumbnail     λ.Object
						ϒtitle         λ.Object
						ϒupload_date   λ.Object
						ϒuploader      λ.Object
						ϒurl           = λargs[1]
						ϒvideo_id      λ.Object
						ϒvideo_size    λ.Object
						ϒvideo_url     λ.Object
						ϒview_count    λ.Object
					)
					ϒmobj = λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("id"))
					ϒitem = λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒself, "_download_xml", nil), λ.Mod(λ.NewStr("http://rd3.videos.sapo.pt/%s/rss2"), ϒvideo_id), ϒvideo_id), "find", nil), λ.NewStr("./channel/item"))
					ϒtitle = λ.GetAttr(λ.Cal(λ.GetAttr(ϒitem, "find", nil), λ.NewStr("./title")), "text", nil)
					ϒdescription = λ.GetAttr(λ.Cal(λ.GetAttr(ϒitem, "find", nil), λ.NewStr("./{http://videos.sapo.pt/mrss/}synopse")), "text", nil)
					ϒthumbnail = λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒitem, "find", nil), λ.NewStr("./{http://search.yahoo.com/mrss/}content")), "get", nil), λ.NewStr("url"))
					ϒduration = λ.Cal(ϒparse_duration, λ.GetAttr(λ.Cal(λ.GetAttr(ϒitem, "find", nil), λ.NewStr("./{http://videos.sapo.pt/mrss/}time")), "text", nil))
					ϒuploader = λ.GetAttr(λ.Cal(λ.GetAttr(ϒitem, "find", nil), λ.NewStr("./{http://videos.sapo.pt/mrss/}author")), "text", nil)
					ϒupload_date = λ.Cal(ϒunified_strdate, λ.GetAttr(λ.Cal(λ.GetAttr(ϒitem, "find", nil), λ.NewStr("./pubDate")), "text", nil))
					ϒview_count = λ.Cal(λ.IntType, λ.GetAttr(λ.Cal(λ.GetAttr(ϒitem, "find", nil), λ.NewStr("./{http://videos.sapo.pt/mrss/}views")), "text", nil))
					ϒcomment_count = λ.Cal(λ.IntType, λ.GetAttr(λ.Cal(λ.GetAttr(ϒitem, "find", nil), λ.NewStr("./{http://videos.sapo.pt/mrss/}comment_count")), "text", nil))
					ϒtags = λ.GetAttr(λ.Cal(λ.GetAttr(ϒitem, "find", nil), λ.NewStr("./{http://videos.sapo.pt/mrss/}tags")), "text", nil)
					ϒcategories = func() λ.Object {
						if λ.IsTrue(ϒtags) {
							return λ.Cal(λ.GetAttr(ϒtags, "split", nil))
						} else {
							return λ.NewList()
						}
					}()
					ϒage_limit = func() λ.Object {
						if λ.IsTrue(λ.Eq(λ.GetAttr(λ.Cal(λ.GetAttr(ϒitem, "find", nil), λ.NewStr("./{http://videos.sapo.pt/mrss/}m18")), "text", nil), λ.NewStr("true"))) {
							return λ.NewInt(18)
						} else {
							return λ.NewInt(0)
						}
					}()
					ϒvideo_url = λ.GetAttr(λ.Cal(λ.GetAttr(ϒitem, "find", nil), λ.NewStr("./{http://videos.sapo.pt/mrss/}videoFile")), "text", nil)
					ϒvideo_size = λ.Cal(λ.GetAttr(λ.GetAttr(λ.Cal(λ.GetAttr(ϒitem, "find", nil), λ.NewStr("./{http://videos.sapo.pt/mrss/}videoSize")), "text", nil), "split", nil), λ.NewStr("x"))
					ϒformats = λ.NewList(λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("url"):       ϒvideo_url,
						λ.NewStr("ext"):       λ.NewStr("mp4"),
						λ.NewStr("format_id"): λ.NewStr("sd"),
						λ.NewStr("width"):     λ.Cal(λ.IntType, λ.GetItem(ϒvideo_size, λ.NewInt(0))),
						λ.NewStr("height"):    λ.Cal(λ.IntType, λ.GetItem(ϒvideo_size, λ.NewInt(1))),
					}))
					if λ.IsTrue(λ.Eq(λ.GetAttr(λ.Cal(λ.GetAttr(ϒitem, "find", nil), λ.NewStr("./{http://videos.sapo.pt/mrss/}HD")), "text", nil), λ.NewStr("true"))) {
						λ.Cal(λ.GetAttr(ϒformats, "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("url"):       λ.Cal(Ωre.ϒsub, λ.NewStr("/mov/1$"), λ.NewStr("/mov/39"), ϒvideo_url),
							λ.NewStr("ext"):       λ.NewStr("mp4"),
							λ.NewStr("format_id"): λ.NewStr("hd"),
							λ.NewStr("width"):     λ.NewInt(1280),
							λ.NewStr("height"):    λ.NewInt(720),
						}))
					}
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):            ϒvideo_id,
						λ.NewStr("title"):         ϒtitle,
						λ.NewStr("description"):   ϒdescription,
						λ.NewStr("thumbnail"):     ϒthumbnail,
						λ.NewStr("duration"):      ϒduration,
						λ.NewStr("uploader"):      ϒuploader,
						λ.NewStr("upload_date"):   ϒupload_date,
						λ.NewStr("view_count"):    ϒview_count,
						λ.NewStr("comment_count"): ϒcomment_count,
						λ.NewStr("categories"):    ϒcategories,
						λ.NewStr("age_limit"):     ϒage_limit,
						λ.NewStr("formats"):       ϒformats,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TESTS"):        SapoIE__TESTS,
				λ.NewStr("_VALID_URL"):    SapoIE__VALID_URL,
				λ.NewStr("_real_extract"): SapoIE__real_extract,
			})
		}())
	})
}
