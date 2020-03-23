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
 * xtube/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/xtube.py
 */

package xtube

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor      λ.Object
	XTubeIE            λ.Object
	XTubeUserIE        λ.Object
	ϒint_or_none       λ.Object
	ϒjs_to_json        λ.Object
	ϒorderedSet        λ.Object
	ϒparse_duration    λ.Object
	ϒsanitized_Request λ.Object
	ϒstr_to_int        λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒjs_to_json = Ωutils.ϒjs_to_json
		ϒorderedSet = Ωutils.ϒorderedSet
		ϒparse_duration = Ωutils.ϒparse_duration
		ϒsanitized_Request = Ωutils.ϒsanitized_Request
		ϒstr_to_int = Ωutils.ϒstr_to_int
		XTubeIE = λ.Cal(λ.TypeType, λ.StrLiteral("XTubeIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				XTubeIE__VALID_URL    λ.Object
				XTubeIE__real_extract λ.Object
			)
			XTubeIE__VALID_URL = λ.StrLiteral("(?x)\n                        (?:\n                            xtube:|\n                            https?://(?:www\\.)?xtube\\.com/(?:watch\\.php\\?.*\\bv=|video-watch/(?:embedded/)?(?P<display_id>[^/]+)-)\n                        )\n                        (?P<id>[^/?&#]+)\n                    ")
			XTubeIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒcomment_count λ.Object
						ϒconfig        λ.Object
						ϒdescription   λ.Object
						ϒdisplay_id    λ.Object
						ϒduration      λ.Object
						ϒformat_id     λ.Object
						ϒformat_url    λ.Object
						ϒformats       λ.Object
						ϒmobj          λ.Object
						ϒself          = λargs[0]
						ϒsources       λ.Object
						ϒthumbnail     λ.Object
						ϒtitle         λ.Object
						ϒuploader      λ.Object
						ϒurl           = λargs[1]
						ϒurl_pattern   λ.Object
						ϒvideo_id      λ.Object
						ϒview_count    λ.Object
						ϒwebpage       λ.Object
						τmp0           λ.Object
						τmp1           λ.Object
						τmp2           λ.Object
					)
					ϒmobj = λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl)
					ϒvideo_id = λ.Calm(ϒmobj, "group", λ.StrLiteral("id"))
					ϒdisplay_id = λ.Calm(ϒmobj, "group", λ.StrLiteral("display_id"))
					if !λ.IsTrue(ϒdisplay_id) {
						ϒdisplay_id = ϒvideo_id
					}
					if λ.IsTrue(func() λ.Object {
						if λv := λ.Calm(ϒvideo_id, "isdigit"); !λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Lt(λ.Cal(λ.BuiltinLen, ϒvideo_id), λ.IntLiteral(11))
						}
					}()) {
						ϒurl_pattern = λ.StrLiteral("http://www.xtube.com/video-watch/-%s")
					} else {
						ϒurl_pattern = λ.StrLiteral("http://www.xtube.com/watch.php?v=%s")
					}
					ϒwebpage = λ.Call(λ.GetAttr(ϒself, "_download_webpage", nil), λ.NewArgs(
						λ.Mod(ϒurl_pattern, ϒvideo_id),
						ϒdisplay_id,
					), λ.KWArgs{
						{Name: "headers", Value: λ.DictLiteral(map[string]string{
							"Cookie": "age_verified=1; cookiesAccepted=1",
						})},
					})
					τmp0 = λ.Mul(λ.NewList(λ.None), λ.IntLiteral(3))
					ϒtitle = λ.GetItem(τmp0, λ.IntLiteral(0))
					ϒthumbnail = λ.GetItem(τmp0, λ.IntLiteral(1))
					ϒduration = λ.GetItem(τmp0, λ.IntLiteral(2))
					ϒconfig = λ.Call(λ.GetAttr(ϒself, "_parse_json", nil), λ.NewArgs(
						λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
							λ.StrLiteral("playerConf\\s*=\\s*({.+?})\\s*,\\s*\\n"),
							ϒwebpage,
							λ.StrLiteral("config"),
						), λ.KWArgs{
							{Name: "default", Value: λ.StrLiteral("{}")},
						}),
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "transform_source", Value: ϒjs_to_json},
						{Name: "fatal", Value: λ.False},
					})
					if λ.IsTrue(ϒconfig) {
						ϒconfig = λ.Calm(ϒconfig, "get", λ.StrLiteral("mainRoll"))
						if λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒconfig, λ.DictType)) {
							ϒtitle = λ.Calm(ϒconfig, "get", λ.StrLiteral("title"))
							ϒthumbnail = λ.Calm(ϒconfig, "get", λ.StrLiteral("poster"))
							ϒduration = λ.Cal(ϒint_or_none, λ.Calm(ϒconfig, "get", λ.StrLiteral("duration")))
							ϒsources = func() λ.Object {
								if λv := λ.Calm(ϒconfig, "get", λ.StrLiteral("sources")); λ.IsTrue(λv) {
									return λv
								} else {
									return λ.Calm(ϒconfig, "get", λ.StrLiteral("format"))
								}
							}()
						}
					}
					if !λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒsources, λ.DictType)) {
						ϒsources = λ.Call(λ.GetAttr(ϒself, "_parse_json", nil), λ.NewArgs(
							λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
								λ.StrLiteral("([\"\\'])?sources\\1?\\s*:\\s*(?P<sources>{.+?}),"),
								ϒwebpage,
								λ.StrLiteral("sources"),
							), λ.KWArgs{
								{Name: "group", Value: λ.StrLiteral("sources")},
							}),
							ϒvideo_id,
						), λ.KWArgs{
							{Name: "transform_source", Value: ϒjs_to_json},
						})
					}
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(ϒsources, "items"))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						τmp2 = τmp1
						ϒformat_id = λ.GetItem(τmp2, λ.IntLiteral(0))
						ϒformat_url = λ.GetItem(τmp2, λ.IntLiteral(1))
						λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
							"url":       ϒformat_url,
							"format_id": ϒformat_id,
							"height":    λ.Cal(ϒint_or_none, ϒformat_id),
						}))
					}
					λ.Calm(ϒself, "_remove_duplicate_formats", ϒformats)
					λ.Calm(ϒself, "_sort_formats", ϒformats)
					if !λ.IsTrue(ϒtitle) {
						ϒtitle = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
							λ.NewTuple(
								λ.StrLiteral("<h1>\\s*(?P<title>[^<]+?)\\s*</h1>"),
								λ.StrLiteral("videoTitle\\s*:\\s*([\"\\'])(?P<title>.+?)\\1"),
							),
							ϒwebpage,
							λ.StrLiteral("title"),
						), λ.KWArgs{
							{Name: "group", Value: λ.StrLiteral("title")},
						})
					}
					ϒdescription = func() λ.Object {
						if λv := λ.Call(λ.GetAttr(ϒself, "_og_search_description", nil), λ.NewArgs(ϒwebpage), λ.KWArgs{
							{Name: "default", Value: λ.None},
						}); λ.IsTrue(λv) {
							return λv
						} else if λv := λ.Call(λ.GetAttr(ϒself, "_html_search_meta", nil), λ.NewArgs(
							λ.StrLiteral("twitter:description"),
							ϒwebpage,
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
						}); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
								λ.StrLiteral("</h1>\\s*<p>([^<]+)"),
								ϒwebpage,
								λ.StrLiteral("description"),
							), λ.KWArgs{
								{Name: "fatal", Value: λ.False},
							})
						}
					}()
					ϒuploader = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewTuple(
							λ.StrLiteral("<input[^>]+name=\"contentOwnerId\"[^>]+value=\"([^\"]+)\""),
							λ.StrLiteral("<span[^>]+class=\"nickname\"[^>]*>([^<]+)"),
						),
						ϒwebpage,
						λ.StrLiteral("uploader"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					})
					if !λ.IsTrue(ϒduration) {
						ϒduration = λ.Cal(ϒparse_duration, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
							λ.StrLiteral("<dt>Runtime:?</dt>\\s*<dd>([^<]+)</dd>"),
							ϒwebpage,
							λ.StrLiteral("duration"),
						), λ.KWArgs{
							{Name: "fatal", Value: λ.False},
						}))
					}
					ϒview_count = λ.Cal(ϒstr_to_int, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewTuple(
							λ.StrLiteral("[\"\\']viewsCount[\"\\'][^>]*>(\\d+)\\s+views"),
							λ.StrLiteral("<dt>Views:?</dt>\\s*<dd>([\\d,\\.]+)</dd>"),
						),
						ϒwebpage,
						λ.StrLiteral("view count"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					}))
					ϒcomment_count = λ.Cal(ϒstr_to_int, λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
						λ.StrLiteral(">Comments? \\(([\\d,\\.]+)\\)<"),
						ϒwebpage,
						λ.StrLiteral("comment count"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					}))
					return λ.DictLiteral(map[string]λ.Object{
						"id":            ϒvideo_id,
						"display_id":    ϒdisplay_id,
						"title":         ϒtitle,
						"description":   ϒdescription,
						"thumbnail":     ϒthumbnail,
						"uploader":      ϒuploader,
						"duration":      ϒduration,
						"view_count":    ϒview_count,
						"comment_count": ϒcomment_count,
						"age_limit":     λ.IntLiteral(18),
						"formats":       ϒformats,
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL":    XTubeIE__VALID_URL,
				"_real_extract": XTubeIE__real_extract,
			})
		}())
		XTubeUserIE = λ.Cal(λ.TypeType, λ.StrLiteral("XTubeUserIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				XTubeUserIE__VALID_URL λ.Object
			)
			XTubeUserIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?xtube\\.com/profile/(?P<id>[^/]+-\\d+)")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": XTubeUserIE__VALID_URL,
			})
		}())
	})
}
