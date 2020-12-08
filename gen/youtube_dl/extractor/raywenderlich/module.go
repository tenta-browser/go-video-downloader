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
 * raywenderlich/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/raywenderlich.py
 */

package raywenderlich

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωvimeo "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/vimeo"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError        λ.Object
	InfoExtractor         λ.Object
	RayWenderlichCourseIE λ.Object
	RayWenderlichIE       λ.Object
	VimeoIE               λ.Object
	ϒcompat_str           λ.Object
	ϒint_or_none          λ.Object
	ϒmerge_dicts          λ.Object
	ϒtry_get              λ.Object
	ϒunescapeHTML         λ.Object
	ϒunified_timestamp    λ.Object
	ϒurljoin              λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		VimeoIE = Ωvimeo.VimeoIE
		ϒcompat_str = Ωcompat.ϒcompat_str
		ExtractorError = Ωutils.ExtractorError
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒmerge_dicts = Ωutils.ϒmerge_dicts
		ϒtry_get = Ωutils.ϒtry_get
		ϒunescapeHTML = Ωutils.ϒunescapeHTML
		ϒunified_timestamp = Ωutils.ϒunified_timestamp
		ϒurljoin = Ωutils.ϒurljoin
		RayWenderlichIE = λ.Cal(λ.TypeType, λ.StrLiteral("RayWenderlichIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				RayWenderlichIE__VALID_URL        λ.Object
				RayWenderlichIE__extract_video_id λ.Object
				RayWenderlichIE__real_extract     λ.Object
			)
			RayWenderlichIE__VALID_URL = λ.StrLiteral("(?x)\n                    https?://\n                        (?:\n                            videos\\.raywenderlich\\.com/courses|\n                            (?:www\\.)?raywenderlich\\.com\n                        )/\n                        (?P<course_id>[^/]+)/lessons/(?P<id>\\d+)\n                    ")
			RayWenderlichIE__extract_video_id = λ.NewFunction("_extract_video_id",
				[]λ.Param{
					{Name: "data"},
					{Name: "lesson_id"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒcontent   λ.Object
						ϒcontents  λ.Object
						ϒdata      = λargs[0]
						ϒgroup     λ.Object
						ϒgroups    λ.Object
						ϒlesson_id = λargs[1]
						ϒordinal   λ.Object
						ϒvideo_id  λ.Object
						τmp0       λ.Object
						τmp1       λ.Object
						τmp2       λ.Object
						τmp3       λ.Object
					)
					if !λ.IsTrue(ϒdata) {
						return λ.None
					}
					ϒgroups = func() λ.Object {
						if λv := λ.Cal(ϒtry_get, ϒdata, λ.NewFunction("<lambda>",
							[]λ.Param{
								{Name: "x"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒx = λargs[0]
								)
								return λ.GetItem(ϒx, λ.StrLiteral("groups"))
							}), λ.ListType); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.NewList()
						}
					}()
					if !λ.IsTrue(ϒgroups) {
						return λ.None
					}
					τmp0 = λ.Cal(λ.BuiltinIter, ϒgroups)
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒgroup = τmp1
						if !λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒgroup, λ.DictType)) {
							continue
						}
						ϒcontents = func() λ.Object {
							if λv := λ.Cal(ϒtry_get, ϒdata, λ.NewFunction("<lambda>",
								[]λ.Param{
									{Name: "x"},
								},
								0, false, false,
								func(λargs []λ.Object) λ.Object {
									var (
										ϒx = λargs[0]
									)
									return λ.GetItem(ϒx, λ.StrLiteral("contents"))
								}), λ.ListType); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.NewList()
							}
						}()
						τmp2 = λ.Cal(λ.BuiltinIter, ϒcontents)
						for {
							if τmp3 = λ.NextDefault(τmp2, λ.AfterLast); τmp3 == λ.AfterLast {
								break
							}
							ϒcontent = τmp3
							if !λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒcontent, λ.DictType)) {
								continue
							}
							ϒordinal = λ.Cal(ϒint_or_none, λ.Calm(ϒcontent, "get", λ.StrLiteral("ordinal")))
							if λ.IsTrue(λ.Ne(ϒordinal, ϒlesson_id)) {
								continue
							}
							ϒvideo_id = λ.Calm(ϒcontent, "get", λ.StrLiteral("identifier"))
							if λ.IsTrue(ϒvideo_id) {
								return λ.Cal(ϒcompat_str, ϒvideo_id)
							}
						}
					}
					return λ.None
				})
			RayWenderlichIE__extract_video_id = λ.Cal(λ.StaticMethodType, RayWenderlichIE__extract_video_id)
			RayWenderlichIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒcourse_id  λ.Object
						ϒcsrf_token λ.Object
						ϒdata       λ.Object
						ϒdisplay_id λ.Object
						ϒheaders    λ.Object
						ϒinfo       λ.Object
						ϒlesson_id  λ.Object
						ϒmobj       λ.Object
						ϒself       = λargs[0]
						ϒthumbnail  λ.Object
						ϒurl        = λargs[1]
						ϒvideo      λ.Object
						ϒvideo_id   λ.Object
						ϒvimeo_id   λ.Object
						ϒwebpage    λ.Object
						τmp0        λ.Object
					)
					ϒmobj = λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl)
					τmp0 = λ.Calm(ϒmobj, "group", λ.StrLiteral("course_id"), λ.StrLiteral("id"))
					ϒcourse_id = λ.GetItem(τmp0, λ.IntLiteral(0))
					ϒlesson_id = λ.GetItem(τmp0, λ.IntLiteral(1))
					ϒdisplay_id = λ.Mod(λ.StrLiteral("%s/%s"), λ.NewTuple(
						ϒcourse_id,
						ϒlesson_id,
					))
					ϒwebpage = λ.Calm(ϒself, "_download_webpage", ϒurl, ϒdisplay_id)
					ϒthumbnail = func() λ.Object {
						if λv := λ.Call(λ.GetAttr(ϒself, "_og_search_thumbnail", nil), λ.NewArgs(ϒwebpage), λ.KWArgs{
							{Name: "default", Value: λ.None},
						}); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Calm(ϒself, "_html_search_meta", λ.StrLiteral("twitter:image"), ϒwebpage, λ.StrLiteral("thumbnail"))
						}
					}()
					if λ.Contains(ϒwebpage, λ.StrLiteral(">Subscribe to unlock")) {
						panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.StrLiteral("This content is only available for subscribers")), λ.KWArgs{
							{Name: "expected", Value: λ.True},
						})))
					}
					ϒinfo = λ.DictLiteral(map[string]λ.Object{
						"thumbnail": ϒthumbnail,
					})
					ϒvimeo_id = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("data-vimeo-id=[\"\\'](\\d+)"),
						ϒwebpage,
						λ.StrLiteral("vimeo id"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					if !λ.IsTrue(ϒvimeo_id) {
						ϒdata = λ.Call(λ.GetAttr(ϒself, "_parse_json", nil), λ.NewArgs(
							λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
								λ.StrLiteral("data-collection=([\"\\'])(?P<data>{.+?})\\1"),
								ϒwebpage,
								λ.StrLiteral("data collection"),
							), λ.KWArgs{
								{Name: "default", Value: λ.StrLiteral("{}")},
								{Name: "group", Value: λ.StrLiteral("data")},
							}),
							ϒdisplay_id,
						), λ.KWArgs{
							{Name: "transform_source", Value: ϒunescapeHTML},
							{Name: "fatal", Value: λ.False},
						})
						ϒvideo_id = func() λ.Object {
							if λv := λ.Calm(ϒself, "_extract_video_id", ϒdata, ϒlesson_id); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Calm(ϒself, "_search_regex", λ.StrLiteral("/videos/(\\d+)/"), ϒthumbnail, λ.StrLiteral("video id"))
							}
						}()
						ϒheaders = λ.DictLiteral(map[string]λ.Object{
							"Referer":          ϒurl,
							"X-Requested-With": λ.StrLiteral("XMLHttpRequest"),
						})
						ϒcsrf_token = λ.Call(λ.GetAttr(ϒself, "_html_search_meta", nil), λ.NewArgs(
							λ.StrLiteral("csrf-token"),
							ϒwebpage,
							λ.StrLiteral("csrf token"),
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
						})
						if λ.IsTrue(ϒcsrf_token) {
							λ.SetItem(ϒheaders, λ.StrLiteral("X-CSRF-Token"), ϒcsrf_token)
						}
						ϒvideo = λ.GetItem(λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
							λ.Mod(λ.StrLiteral("https://videos.raywenderlich.com/api/v1/videos/%s.json"), ϒvideo_id),
							ϒdisplay_id,
						), λ.KWArgs{
							{Name: "headers", Value: ϒheaders},
						}), λ.StrLiteral("video"))
						ϒvimeo_id = λ.GetItem(λ.GetItem(λ.GetItem(ϒvideo, λ.StrLiteral("clips")), λ.IntLiteral(0)), λ.StrLiteral("provider_id"))
						λ.Calm(ϒinfo, "update", λ.DictLiteral(map[string]λ.Object{
							"_type": λ.StrLiteral("url_transparent"),
							"title": λ.Calm(ϒvideo, "get", λ.StrLiteral("name")),
							"description": func() λ.Object {
								if λv := λ.Calm(ϒvideo, "get", λ.StrLiteral("description")); λ.IsTrue(λv) {
									return λv
								} else {
									return λ.Calm(ϒvideo, "get", λ.StrLiteral("meta_description"))
								}
							}(),
							"duration":  λ.Cal(ϒint_or_none, λ.Calm(ϒvideo, "get", λ.StrLiteral("duration"))),
							"timestamp": λ.Cal(ϒunified_timestamp, λ.Calm(ϒvideo, "get", λ.StrLiteral("created_at"))),
						}))
					}
					return λ.Cal(ϒmerge_dicts, ϒinfo, λ.Call(λ.GetAttr(ϒself, "url_result", nil), λ.NewArgs(λ.Calm(VimeoIE, "_smuggle_referrer", λ.Mod(λ.StrLiteral("https://player.vimeo.com/video/%s"), ϒvimeo_id), ϒurl)), λ.KWArgs{
						{Name: "ie", Value: λ.Calm(VimeoIE, "ie_key")},
						{Name: "video_id", Value: ϒvimeo_id},
					}))
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL":        RayWenderlichIE__VALID_URL,
				"_extract_video_id": RayWenderlichIE__extract_video_id,
				"_real_extract":     RayWenderlichIE__real_extract,
			})
		}())
		RayWenderlichCourseIE = λ.Cal(λ.TypeType, λ.StrLiteral("RayWenderlichCourseIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				RayWenderlichCourseIE__VALID_URL λ.Object
				RayWenderlichCourseIE_suitable   λ.Object
			)
			RayWenderlichCourseIE__VALID_URL = λ.StrLiteral("(?x)\n                    https?://\n                        (?:\n                            videos\\.raywenderlich\\.com/courses|\n                            (?:www\\.)?raywenderlich\\.com\n                        )/\n                        (?P<id>[^/]+)\n                    ")
			RayWenderlichCourseIE_suitable = λ.NewFunction("suitable",
				[]λ.Param{
					{Name: "cls"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒcls = λargs[0]
						ϒurl = λargs[1]
					)
					return func() λ.Object {
						if λ.IsTrue(λ.Calm(RayWenderlichIE, "suitable", ϒurl)) {
							return λ.False
						} else {
							return λ.Calm(λ.Cal(λ.SuperType, RayWenderlichCourseIE, ϒcls), "suitable", ϒurl)
						}
					}()
				})
			RayWenderlichCourseIE_suitable = λ.Cal(λ.ClassMethodType, RayWenderlichCourseIE_suitable)
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": RayWenderlichCourseIE__VALID_URL,
				"suitable":   RayWenderlichCourseIE_suitable,
			})
		}())
	})
}
