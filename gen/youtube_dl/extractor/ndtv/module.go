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
 * ndtv/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/ndtv.py
 */

package ndtv

import (
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor                     λ.Object
	NDTVIE                            λ.Object
	ϒcompat_urllib_parse_unquote_plus λ.Object
	ϒparse_duration                   λ.Object
	ϒremove_end                       λ.Object
	ϒunified_strdate                  λ.Object
	ϒurljoin                          λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_urllib_parse_unquote_plus = Ωcompat.ϒcompat_urllib_parse_unquote_plus
		ϒparse_duration = Ωutils.ϒparse_duration
		ϒremove_end = Ωutils.ϒremove_end
		ϒunified_strdate = Ωutils.ϒunified_strdate
		ϒurljoin = Ωutils.ϒurljoin
		NDTVIE = λ.Cal(λ.TypeType, λ.NewStr("NDTVIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				NDTVIE__TESTS        λ.Object
				NDTVIE__VALID_URL    λ.Object
				NDTVIE__real_extract λ.Object
			)
			NDTVIE__VALID_URL = λ.NewStr("https?://(?:[^/]+\\.)?ndtv\\.com/(?:[^/]+/)*videos?/?(?:[^/]+/)*[^/?^&]+-(?P<id>\\d+)")
			NDTVIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("https://khabar.ndtv.com/video/show/prime-time/prime-time-ill-system-and-poor-education-468818"),
					λ.NewStr("md5"): λ.NewStr("78efcf3880ef3fd9b83d405ca94a38eb"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("468818"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("प्राइम टाइम: सिस्टम बीमार, स्कूल बदहाल"),
						λ.NewStr("description"): λ.NewStr("md5:f410512f1b49672e5695dea16ef2731d"),
						λ.NewStr("upload_date"): λ.NewStr("20170928"),
						λ.NewStr("duration"):    λ.NewInt(2218),
						λ.NewStr("thumbnail"):   λ.NewStr("re:https?://.*\\.jpg"),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://movies.ndtv.com/videos/cracker-free-diwali-wishes-from-karan-johar-kriti-sanon-other-stars-470304"),
					λ.NewStr("md5"): λ.NewStr("f1d709352305b44443515ac56b45aa46"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("470304"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Cracker-Free Diwali Wishes From Karan Johar, Kriti Sanon & Other Stars"),
						λ.NewStr("description"): λ.NewStr("md5:f115bba1adf2f6433fa7c1ade5feb465"),
						λ.NewStr("upload_date"): λ.NewStr("20171019"),
						λ.NewStr("duration"):    λ.NewInt(137),
						λ.NewStr("thumbnail"):   λ.NewStr("re:https?://.*\\.jpg"),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://www.ndtv.com/video/news/news/delhi-s-air-quality-status-report-after-diwali-is-very-poor-470372"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://auto.ndtv.com/videos/the-cnb-daily-october-13-2017-469935"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://sports.ndtv.com/cricket/videos/2nd-t20i-rock-thrown-at-australia-cricket-team-bus-after-win-over-india-469764"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://gadgets.ndtv.com/videos/uncharted-the-lost-legacy-review-465568"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://profit.ndtv.com/videos/news/video-indian-economy-on-very-solid-track-international-monetary-fund-chief-470040"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://food.ndtv.com/video-basil-seeds-coconut-porridge-419083"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://doctor.ndtv.com/videos/top-health-stories-of-the-week-467396"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://swirlster.ndtv.com/video/how-to-make-friends-at-work-469324"),
					λ.NewStr("only_matching"): λ.True,
				}),
			)
			NDTVIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdescription λ.Object
						ϒduration    λ.Object
						ϒfilename    λ.Object
						ϒself        = λargs[0]
						ϒtitle       λ.Object
						ϒupload_date λ.Object
						ϒurl         = λargs[1]
						ϒvideo_id    λ.Object
						ϒvideo_url   λ.Object
						ϒwebpage     λ.Object
					)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒvideo_id)
					ϒtitle = λ.Cal(ϒcompat_urllib_parse_unquote_plus, func() λ.Object {
						if λv := λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
							λ.NewStr("__title\\s*=\\s*'([^']+)'"),
							ϒwebpage,
							λ.NewStr("title"),
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
						}); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Cal(λ.GetAttr(ϒself, "_og_search_title", nil), ϒwebpage)
						}
					}())
					ϒfilename = λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewStr("(?:__)?filename\\s*[:=]\\s*'([^']+)'"), ϒwebpage, λ.NewStr("video filename"))
					ϒvideo_url = λ.Cal(ϒurljoin, λ.NewStr("https://ndtvod.bc-ssl.cdn.bitgravity.com/23372/ndtv/"), λ.Cal(λ.GetAttr(ϒfilename, "lstrip", nil), λ.NewStr("/")))
					ϒduration = λ.Cal(ϒparse_duration, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr("(?:__)?duration\\s*[:=]\\s*'([^']+)'"),
						ϒwebpage,
						λ.NewStr("duration"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					}))
					ϒupload_date = λ.Cal(ϒunified_strdate, func() λ.Object {
						if λv := λ.Call(λ.GetAttr(ϒself, "_html_search_meta", nil), λ.NewArgs(
							λ.NewStr("publish-date"),
							ϒwebpage,
							λ.NewStr("upload date"),
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
						}); λ.IsTrue(λv) {
							return λv
						} else if λv := λ.Call(λ.GetAttr(ϒself, "_html_search_meta", nil), λ.NewArgs(
							λ.NewStr("uploadDate"),
							ϒwebpage,
							λ.NewStr("upload date"),
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
						}); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
								λ.NewStr("datePublished\"\\s*:\\s*\"([^\"]+)\""),
								ϒwebpage,
								λ.NewStr("upload date"),
							), λ.KWArgs{
								{Name: "fatal", Value: λ.False},
							})
						}
					}())
					ϒdescription = λ.Cal(ϒremove_end, λ.Cal(λ.GetAttr(ϒself, "_og_search_description", nil), ϒwebpage), λ.NewStr(" (Read more)"))
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          ϒvideo_id,
						λ.NewStr("url"):         ϒvideo_url,
						λ.NewStr("title"):       ϒtitle,
						λ.NewStr("description"): ϒdescription,
						λ.NewStr("thumbnail"):   λ.Cal(λ.GetAttr(ϒself, "_og_search_thumbnail", nil), ϒwebpage),
						λ.NewStr("duration"):    ϒduration,
						λ.NewStr("upload_date"): ϒupload_date,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TESTS"):        NDTVIE__TESTS,
				λ.NewStr("_VALID_URL"):    NDTVIE__VALID_URL,
				λ.NewStr("_real_extract"): NDTVIE__real_extract,
			})
		}())
	})
}
