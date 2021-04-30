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
 * xvideos/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/xvideos.py
 */

package xvideos

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError               λ.Object
	InfoExtractor                λ.Object
	XVideosIE                    λ.Object
	ϒclean_html                  λ.Object
	ϒcompat_urllib_parse_unquote λ.Object
	ϒdetermine_ext               λ.Object
	ϒint_or_none                 λ.Object
	ϒparse_duration              λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_urllib_parse_unquote = Ωcompat.ϒcompat_urllib_parse_unquote
		ϒclean_html = Ωutils.ϒclean_html
		ϒdetermine_ext = Ωutils.ϒdetermine_ext
		ExtractorError = Ωutils.ExtractorError
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒparse_duration = Ωutils.ϒparse_duration
		XVideosIE = λ.Cal(λ.TypeType, λ.StrLiteral("XVideosIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				XVideosIE__VALID_URL    λ.Object
				XVideosIE__real_extract λ.Object
			)
			XVideosIE__VALID_URL = λ.StrLiteral("(?x)\n                    https?://\n                        (?:\n                            (?:[^/]+\\.)?xvideos2?\\.com/video|\n                            (?:www\\.)?xvideos\\.es/video|\n                            flashservice\\.xvideos\\.com/embedframe/|\n                            static-hw\\.xvideos\\.com/swf/xv-player\\.swf\\?.*?\\bid_video=\n                        )\n                        (?P<id>[0-9]+)\n                    ")
			XVideosIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒduration      λ.Object
						ϒformat_id     λ.Object
						ϒformat_url    λ.Object
						ϒformats       λ.Object
						ϒkind          λ.Object
						ϒmobj          λ.Object
						ϒpreference    λ.Object
						ϒself          = λargs[0]
						ϒthumbnail     λ.Object
						ϒthumbnail_url λ.Object
						ϒthumbnails    λ.Object
						ϒtitle         λ.Object
						ϒurl           = λargs[1]
						ϒvideo_id      λ.Object
						ϒvideo_url     λ.Object
						ϒwebpage       λ.Object
						τmp0           λ.Object
						τmp1           λ.Object
						τmp2           λ.Object
					)
					ϒvideo_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒwebpage = λ.Calm(ϒself, "_download_webpage", λ.Mod(λ.StrLiteral("https://www.xvideos.com/video%s/"), ϒvideo_id), ϒvideo_id)
					ϒmobj = λ.Cal(Ωre.ϒsearch, λ.StrLiteral("<h1 class=\"inlineError\">(.+?)</h1>"), ϒwebpage)
					if λ.IsTrue(ϒmobj) {
						panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.Mod(λ.StrLiteral("%s said: %s"), λ.NewTuple(
							λ.GetAttr(ϒself, "IE_NAME", nil),
							λ.Cal(ϒclean_html, λ.Calm(ϒmobj, "group", λ.IntLiteral(1))),
						))), λ.KWArgs{
							{Name: "expected", Value: λ.True},
						})))
					}
					ϒtitle = func() λ.Object {
						if λv := λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
							λ.NewTuple(
								λ.StrLiteral("<title>(?P<title>.+?)\\s+-\\s+XVID"),
								λ.StrLiteral("setVideoTitle\\s*\\(\\s*([\"\\'])(?P<title>(?:(?!\\1).)+)\\1"),
							),
							ϒwebpage,
							λ.StrLiteral("title"),
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
							{Name: "group", Value: λ.StrLiteral("title")},
						}); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Calm(ϒself, "_og_search_title", ϒwebpage)
						}
					}()
					ϒthumbnails = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.EnumerateIteratorType, λ.NewTuple(
						λ.StrLiteral(""),
						λ.StrLiteral("169"),
					)))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						τmp2 = λ.UnpackIterable(τmp1, 2)
						ϒpreference = λ.GetItem(τmp2, λ.IntLiteral(0))
						ϒthumbnail = λ.GetItem(τmp2, λ.IntLiteral(1))
						ϒthumbnail_url = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
							λ.Mod(λ.StrLiteral("setThumbUrl%s\\(\\s*([\"\\'])(?P<thumbnail>(?:(?!\\1).)+)\\1"), ϒthumbnail),
							ϒwebpage,
							λ.StrLiteral("thumbnail"),
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
							{Name: "group", Value: λ.StrLiteral("thumbnail")},
						})
						if λ.IsTrue(ϒthumbnail_url) {
							λ.Calm(ϒthumbnails, "append", λ.DictLiteral(map[string]λ.Object{
								"url":        ϒthumbnail_url,
								"preference": ϒpreference,
							}))
						}
					}
					ϒduration = func() λ.Object {
						if λv := λ.Cal(ϒint_or_none, λ.Call(λ.GetAttr(ϒself, "_og_search_property", nil), λ.NewArgs(
							λ.StrLiteral("duration"),
							ϒwebpage,
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
						})); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Cal(ϒparse_duration, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
								λ.StrLiteral("<span[^>]+class=[\"\\']duration[\"\\'][^>]*>.*?(\\d[^<]+)"),
								ϒwebpage,
								λ.StrLiteral("duration"),
							), λ.KWArgs{
								{Name: "fatal", Value: λ.False},
							}))
						}
					}()
					ϒformats = λ.NewList()
					ϒvideo_url = λ.Cal(ϒcompat_urllib_parse_unquote, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("flv_url=(.+?)&"),
						ϒwebpage,
						λ.StrLiteral("video URL"),
					), λ.KWArgs{
						{Name: "default", Value: λ.StrLiteral("")},
					}))
					if λ.IsTrue(ϒvideo_url) {
						λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
							"url":       ϒvideo_url,
							"format_id": λ.StrLiteral("flv"),
						}))
					}
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(Ωre.ϒfindall, λ.StrLiteral("setVideo([^(]+)\\(([\"\\'])(http.+?)\\2\\)"), ϒwebpage))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						τmp2 = λ.UnpackIterable(τmp1, 3)
						ϒkind = λ.GetItem(τmp2, λ.IntLiteral(0))
						_ = λ.GetItem(τmp2, λ.IntLiteral(1))
						ϒformat_url = λ.GetItem(τmp2, λ.IntLiteral(2))
						ϒformat_id = λ.Calm(ϒkind, "lower")
						if λ.IsTrue(λ.Eq(ϒformat_id, λ.StrLiteral("hls"))) {
							λ.Calm(ϒformats, "extend", λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
								ϒformat_url,
								ϒvideo_id,
								λ.StrLiteral("mp4"),
							), λ.KWArgs{
								{Name: "entry_protocol", Value: λ.StrLiteral("m3u8_native")},
								{Name: "m3u8_id", Value: λ.StrLiteral("hls")},
								{Name: "fatal", Value: λ.False},
							}))
						} else {
							if λ.Contains(λ.NewTuple(
								λ.StrLiteral("urllow"),
								λ.StrLiteral("urlhigh"),
							), ϒformat_id) {
								λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
									"url": ϒformat_url,
									"format_id": λ.Mod(λ.StrLiteral("%s-%s"), λ.NewTuple(
										λ.Cal(ϒdetermine_ext, ϒformat_url, λ.StrLiteral("mp4")),
										λ.GetItem(ϒformat_id, λ.NewSlice(λ.IntLiteral(3), λ.None, λ.None)),
									)),
									"quality": func() λ.Object {
										if λ.IsTrue(λ.Calm(ϒformat_id, "endswith", λ.StrLiteral("low"))) {
											return λ.Neg(λ.IntLiteral(2))
										} else {
											return λ.None
										}
									}(),
								}))
							}
						}
					}
					λ.Calm(ϒself, "_sort_formats", ϒformats)
					return λ.DictLiteral(map[string]λ.Object{
						"id":         ϒvideo_id,
						"formats":    ϒformats,
						"title":      ϒtitle,
						"duration":   ϒduration,
						"thumbnails": ϒthumbnails,
						"age_limit":  λ.IntLiteral(18),
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL":    XVideosIE__VALID_URL,
				"_real_extract": XVideosIE__real_extract,
			})
		}())
	})
}
