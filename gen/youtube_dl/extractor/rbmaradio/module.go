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
 * rbmaradio/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/rbmaradio.py
 */

package rbmaradio

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor      λ.Object
	RBMARadioIE        λ.Object
	ϒclean_html        λ.Object
	ϒcompat_str        λ.Object
	ϒint_or_none       λ.Object
	ϒunified_timestamp λ.Object
	ϒupdate_url_query  λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_str = Ωcompat.ϒcompat_str
		ϒclean_html = Ωutils.ϒclean_html
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒunified_timestamp = Ωutils.ϒunified_timestamp
		ϒupdate_url_query = Ωutils.ϒupdate_url_query
		RBMARadioIE = λ.Cal(λ.TypeType, λ.NewStr("RBMARadioIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				RBMARadioIE__TEST         λ.Object
				RBMARadioIE__VALID_URL    λ.Object
				RBMARadioIE__real_extract λ.Object
			)
			RBMARadioIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?(?:rbmaradio|redbullradio)\\.com/shows/(?P<show_id>[^/]+)/episodes/(?P<id>[^/?#&]+)")
			RBMARadioIE__TEST = λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("url"): λ.NewStr("https://www.rbmaradio.com/shows/main-stage/episodes/ford-lopatin-live-at-primavera-sound-2011"),
				λ.NewStr("md5"): λ.NewStr("6bc6f9bcb18994b4c983bc3bf4384d95"),
				λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("id"):          λ.NewStr("ford-lopatin-live-at-primavera-sound-2011"),
					λ.NewStr("ext"):         λ.NewStr("mp3"),
					λ.NewStr("title"):       λ.NewStr("Main Stage - Ford & Lopatin at Primavera Sound"),
					λ.NewStr("description"): λ.NewStr("md5:d41d8cd98f00b204e9800998ecf8427e"),
					λ.NewStr("thumbnail"):   λ.NewStr("re:^https?://.*\\.jpg"),
					λ.NewStr("duration"):    λ.NewInt(2452),
					λ.NewStr("timestamp"):   λ.NewInt(1307103164),
					λ.NewStr("upload_date"): λ.NewStr("20110603"),
				}),
			})
			RBMARadioIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdescription λ.Object
						ϒduration    λ.Object
						ϒepisode     λ.Object
						ϒepisode_id  λ.Object
						ϒformats     λ.Object
						ϒmobj        λ.Object
						ϒself        = λargs[0]
						ϒshow_id     λ.Object
						ϒshow_title  λ.Object
						ϒthumbnail   λ.Object
						ϒtimestamp   λ.Object
						ϒtitle       λ.Object
						ϒurl         = λargs[1]
						ϒwebpage     λ.Object
					)
					ϒmobj = λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl)
					ϒshow_id = λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("show_id"))
					ϒepisode_id = λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("id"))
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒepisode_id)
					ϒepisode = λ.GetItem(λ.GetItem(λ.GetItem(λ.Cal(λ.GetAttr(ϒself, "_parse_json", nil), λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewStr("__INITIAL_STATE__\\s*=\\s*({.+?})\\s*</script>"), ϒwebpage, λ.NewStr("json data")), ϒepisode_id), λ.NewStr("episodes")), ϒshow_id), ϒepisode_id)
					ϒtitle = λ.GetItem(ϒepisode, λ.NewStr("title"))
					ϒshow_title = λ.Cal(λ.GetAttr(ϒepisode, "get", nil), λ.NewStr("showTitle"))
					if λ.IsTrue(ϒshow_title) {
						ϒtitle = λ.Mod(λ.NewStr("%s - %s"), λ.NewTuple(
							ϒshow_title,
							ϒtitle,
						))
					}
					ϒformats = λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
						nil,
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							return λ.NewGenerator(func(λgen λ.Generator) λ.Object {
								var (
									ϒabr λ.Object
									τmp0 λ.Object
									τmp1 λ.Object
								)
								τmp0 = λ.Cal(λ.BuiltinIter, λ.NewTuple(
									λ.NewInt(96),
									λ.NewInt(128),
									λ.NewInt(192),
									λ.NewInt(256),
								))
								for {
									if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
										break
									}
									ϒabr = τmp1
									λgen.Yield(λ.NewDictWithTable(map[λ.Object]λ.Object{
										λ.NewStr("url"): λ.Call(ϒupdate_url_query, λ.NewArgs(λ.GetItem(ϒepisode, λ.NewStr("audioURL"))), λ.KWArgs{
											{Name: "query", Value: λ.NewDictWithTable(map[λ.Object]λ.Object{
												λ.NewStr("cbr"): ϒabr,
											})},
										}),
										λ.NewStr("format_id"): λ.Cal(ϒcompat_str, ϒabr),
										λ.NewStr("abr"):       ϒabr,
										λ.NewStr("vcodec"):    λ.NewStr("none"),
									}))
								}
								return λ.None
							})
						})))
					λ.Cal(λ.GetAttr(ϒself, "_check_formats", nil), ϒformats, ϒepisode_id)
					ϒdescription = λ.Cal(ϒclean_html, λ.Cal(λ.GetAttr(ϒepisode, "get", nil), λ.NewStr("longTeaser")))
					ϒthumbnail = λ.Cal(λ.GetAttr(ϒself, "_proto_relative_url", nil), λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒepisode, "get", nil), λ.NewStr("imageURL"), λ.NewDictWithTable(map[λ.Object]λ.Object{})), "get", nil), λ.NewStr("landscape")))
					ϒduration = λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒepisode, "get", nil), λ.NewStr("duration")))
					ϒtimestamp = λ.Cal(ϒunified_timestamp, λ.Cal(λ.GetAttr(ϒepisode, "get", nil), λ.NewStr("publishedAt")))
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          ϒepisode_id,
						λ.NewStr("title"):       ϒtitle,
						λ.NewStr("description"): ϒdescription,
						λ.NewStr("thumbnail"):   ϒthumbnail,
						λ.NewStr("duration"):    ϒduration,
						λ.NewStr("timestamp"):   ϒtimestamp,
						λ.NewStr("formats"):     ϒformats,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TEST"):         RBMARadioIE__TEST,
				λ.NewStr("_VALID_URL"):    RBMARadioIE__VALID_URL,
				λ.NewStr("_real_extract"): RBMARadioIE__real_extract,
			})
		}())
	})
}
