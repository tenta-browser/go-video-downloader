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
 * stitcher/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/stitcher.py
 */

package stitcher

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor  λ.Object
	StitcherIE     λ.Object
	ϒdetermine_ext λ.Object
	ϒint_or_none   λ.Object
	ϒjs_to_json    λ.Object
	ϒunescapeHTML  λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒdetermine_ext = Ωutils.ϒdetermine_ext
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒjs_to_json = Ωutils.ϒjs_to_json
		ϒunescapeHTML = Ωutils.ϒunescapeHTML
		StitcherIE = λ.Cal(λ.TypeType, λ.NewStr("StitcherIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				StitcherIE__TESTS        λ.Object
				StitcherIE__VALID_URL    λ.Object
				StitcherIE__real_extract λ.Object
			)
			StitcherIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?stitcher\\.com/podcast/(?:[^/]+/)+e/(?:(?P<display_id>[^/#?&]+?)-)?(?P<id>\\d+)(?:[/#?&]|$)")
			StitcherIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.stitcher.com/podcast/the-talking-machines/e/40789481?autoplay=true"),
					λ.NewStr("md5"): λ.NewStr("391dd4e021e6edeb7b8e68fbf2e9e940"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("40789481"),
						λ.NewStr("ext"):         λ.NewStr("mp3"),
						λ.NewStr("title"):       λ.NewStr("Machine Learning Mastery and Cancer Clusters"),
						λ.NewStr("description"): λ.NewStr("md5:55163197a44e915a14a1ac3a1de0f2d3"),
						λ.NewStr("duration"):    λ.NewInt(1604),
						λ.NewStr("thumbnail"):   λ.NewStr("re:^https?://.*\\.jpg"),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.stitcher.com/podcast/panoply/vulture-tv/e/the-rare-hourlong-comedy-plus-40846275?autoplay=true"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("40846275"),
						λ.NewStr("display_id"):  λ.NewStr("the-rare-hourlong-comedy-plus"),
						λ.NewStr("ext"):         λ.NewStr("mp3"),
						λ.NewStr("title"):       λ.NewStr("The CW's 'Crazy Ex-Girlfriend'"),
						λ.NewStr("description"): λ.NewStr("md5:04f1e2f98eb3f5cbb094cea0f9e19b17"),
						λ.NewStr("duration"):    λ.NewInt(2235),
						λ.NewStr("thumbnail"):   λ.NewStr("re:^https?://.*\\.jpg"),
					}),
					λ.NewStr("params"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("skip_download"): λ.True,
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://www.stitcher.com/podcast/marketplace-on-stitcher/e/40910226?autoplay=true"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://www.stitcher.com/podcast/panoply/getting-in/e/episode-2a-how-many-extracurriculars-should-i-have-40876278?autoplay=true"),
					λ.NewStr("only_matching"): λ.True,
				}),
			)
			StitcherIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒaudio_id    λ.Object
						ϒdescription λ.Object
						ϒdisplay_id  λ.Object
						ϒduration    λ.Object
						ϒepisode     λ.Object
						ϒformats     λ.Object
						ϒmobj        λ.Object
						ϒself        = λargs[0]
						ϒthumbnail   λ.Object
						ϒtitle       λ.Object
						ϒurl         = λargs[1]
						ϒwebpage     λ.Object
					)
					ϒmobj = λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl)
					ϒaudio_id = λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("id"))
					ϒdisplay_id = func() λ.Object {
						if λv := λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("display_id")); λ.IsTrue(λv) {
							return λv
						} else {
							return ϒaudio_id
						}
					}()
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒdisplay_id)
					ϒepisode = λ.GetItem(λ.GetItem(λ.Cal(λ.GetAttr(ϒself, "_parse_json", nil), λ.Cal(ϒjs_to_json, λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewStr("(?s)var\\s+stitcher(?:Config)?\\s*=\\s*({.+?});\\n"), ϒwebpage, λ.NewStr("episode config"))), ϒdisplay_id), λ.NewStr("config")), λ.NewStr("episode"))
					ϒtitle = λ.Cal(ϒunescapeHTML, λ.GetItem(ϒepisode, λ.NewStr("title")))
					ϒformats = λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
						nil,
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							return λ.NewGenerator(func(λgy λ.Yielder) λ.Object {
								var (
									ϒepisode_key λ.Object
									τmp0         λ.Object
									τmp1         λ.Object
								)
								τmp0 = λ.Cal(λ.BuiltinIter, λ.NewTuple(λ.NewStr("episodeURL")))
								for {
									if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
										break
									}
									ϒepisode_key = τmp1
									if λ.IsTrue(λ.Cal(λ.GetAttr(ϒepisode, "get", nil), ϒepisode_key)) {
										λgy.Yield(λ.NewDictWithTable(map[λ.Object]λ.Object{
											λ.NewStr("url"): λ.GetItem(ϒepisode, ϒepisode_key),
											λ.NewStr("ext"): func() λ.Object {
												if λv := λ.Cal(ϒdetermine_ext, λ.GetItem(ϒepisode, ϒepisode_key)); λ.IsTrue(λv) {
													return λv
												} else {
													return λ.NewStr("mp3")
												}
											}(),
											λ.NewStr("vcodec"): λ.NewStr("none"),
										}))
									}
								}
								return λ.None
							})
						})))
					ϒdescription = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr("Episode Info:\\s*</span>([^<]+)<"),
						ϒwebpage,
						λ.NewStr("description"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					})
					ϒduration = λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒepisode, "get", nil), λ.NewStr("duration")))
					ϒthumbnail = λ.Cal(λ.GetAttr(ϒepisode, "get", nil), λ.NewStr("episodeImage"))
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          ϒaudio_id,
						λ.NewStr("display_id"):  ϒdisplay_id,
						λ.NewStr("title"):       ϒtitle,
						λ.NewStr("description"): ϒdescription,
						λ.NewStr("duration"):    ϒduration,
						λ.NewStr("thumbnail"):   ϒthumbnail,
						λ.NewStr("formats"):     ϒformats,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TESTS"):        StitcherIE__TESTS,
				λ.NewStr("_VALID_URL"):    StitcherIE__VALID_URL,
				λ.NewStr("_real_extract"): StitcherIE__real_extract,
			})
		}())
	})
}