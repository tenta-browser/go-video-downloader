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
 * pearvideo/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/pearvideo.py
 */

package pearvideo

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor      λ.Object
	PearVideoIE        λ.Object
	ϒqualities         λ.Object
	ϒunified_timestamp λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒqualities = Ωutils.ϒqualities
		ϒunified_timestamp = Ωutils.ϒunified_timestamp
		PearVideoIE = λ.Cal(λ.TypeType, λ.NewStr("PearVideoIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				PearVideoIE__TEST         λ.Object
				PearVideoIE__VALID_URL    λ.Object
				PearVideoIE__real_extract λ.Object
			)
			PearVideoIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?pearvideo\\.com/video_(?P<id>\\d+)")
			PearVideoIE__TEST = λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("url"): λ.NewStr("http://www.pearvideo.com/video_1076290"),
				λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("id"):          λ.NewStr("1076290"),
					λ.NewStr("ext"):         λ.NewStr("mp4"),
					λ.NewStr("title"):       λ.NewStr("小浣熊在主人家玻璃上滚石头：没砸"),
					λ.NewStr("description"): λ.NewStr("md5:01d576b747de71be0ee85eb7cac25f9d"),
					λ.NewStr("timestamp"):   λ.NewInt(1494275280),
					λ.NewStr("upload_date"): λ.NewStr("20170508"),
				}),
			})
			PearVideoIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdescription λ.Object
						ϒformats     λ.Object
						ϒquality     λ.Object
						ϒself        = λargs[0]
						ϒtimestamp   λ.Object
						ϒtitle       λ.Object
						ϒurl         = λargs[1]
						ϒvideo_id    λ.Object
						ϒwebpage     λ.Object
					)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒvideo_id)
					ϒquality = λ.Cal(ϒqualities, λ.NewTuple(
						λ.NewStr("ldflv"),
						λ.NewStr("ld"),
						λ.NewStr("sdflv"),
						λ.NewStr("sd"),
						λ.NewStr("hdflv"),
						λ.NewStr("hd"),
						λ.NewStr("src"),
					))
					ϒformats = λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
						nil,
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							return λ.NewGenerator(func(λgen λ.Generator) λ.Object {
								var (
									ϒmobj λ.Object
									τmp0  λ.Object
									τmp1  λ.Object
								)
								τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(Ωre.ϒfinditer, λ.NewStr("(?P<id>[a-zA-Z]+)Url\\s*=\\s*([\"\\'])(?P<url>(?:https?:)?//.+?)\\2"), ϒwebpage))
								for {
									if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
										break
									}
									ϒmobj = τmp1
									λgen.Yield(λ.NewDictWithTable(map[λ.Object]λ.Object{
										λ.NewStr("url"):       λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("url")),
										λ.NewStr("format_id"): λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("id")),
										λ.NewStr("quality"):   λ.Cal(ϒquality, λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("id"))),
									}))
								}
								return λ.None
							})
						})))
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					ϒtitle = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewTuple(
							λ.NewStr("<h1[^>]+\\bclass=([\"\\'])video-tt\\1[^>]*>(?P<value>[^<]+)"),
							λ.NewStr("<[^>]+\\bdata-title=([\"\\'])(?P<value>(?:(?!\\1).)+)\\1"),
						),
						ϒwebpage,
						λ.NewStr("title"),
					), λ.KWArgs{
						{Name: "group", Value: λ.NewStr("value")},
					})
					ϒdescription = func() λ.Object {
						if λv := λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
							λ.NewTuple(
								λ.NewStr("<div[^>]+\\bclass=([\"\\'])summary\\1[^>]*>(?P<value>[^<]+)"),
								λ.NewStr("<[^>]+\\bdata-summary=([\"\\'])(?P<value>(?:(?!\\1).)+)\\1"),
							),
							ϒwebpage,
							λ.NewStr("description"),
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
							{Name: "group", Value: λ.NewStr("value")},
						}); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Cal(λ.GetAttr(ϒself, "_html_search_meta", nil), λ.NewStr("Description"), ϒwebpage)
						}
					}()
					ϒtimestamp = λ.Cal(ϒunified_timestamp, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr("<div[^>]+\\bclass=[\"\\']date[\"\\'][^>]*>([^<]+)"),
						ϒwebpage,
						λ.NewStr("timestamp"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					}))
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          ϒvideo_id,
						λ.NewStr("title"):       ϒtitle,
						λ.NewStr("description"): ϒdescription,
						λ.NewStr("timestamp"):   ϒtimestamp,
						λ.NewStr("formats"):     ϒformats,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TEST"):         PearVideoIE__TEST,
				λ.NewStr("_VALID_URL"):    PearVideoIE__VALID_URL,
				λ.NewStr("_real_extract"): PearVideoIE__real_extract,
			})
		}())
	})
}
