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
 * arnes/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/arnes.py
 */

package arnes

import (
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ArnesIE                       λ.Object
	InfoExtractor                 λ.Object
	ϒcompat_parse_qs              λ.Object
	ϒcompat_urllib_parse_urlparse λ.Object
	ϒfloat_or_none                λ.Object
	ϒint_or_none                  λ.Object
	ϒparse_iso8601                λ.Object
	ϒremove_start                 λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_parse_qs = Ωcompat.ϒcompat_parse_qs
		ϒcompat_urllib_parse_urlparse = Ωcompat.ϒcompat_urllib_parse_urlparse
		ϒfloat_or_none = Ωutils.ϒfloat_or_none
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒparse_iso8601 = Ωutils.ϒparse_iso8601
		ϒremove_start = Ωutils.ϒremove_start
		ArnesIE = λ.Cal(λ.TypeType, λ.StrLiteral("ArnesIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				ArnesIE_IE_NAME       λ.Object
				ArnesIE__BASE_URL     λ.Object
				ArnesIE__VALID_URL    λ.Object
				ArnesIE__real_extract λ.Object
			)
			ArnesIE_IE_NAME = λ.StrLiteral("video.arnes.si")
			ArnesIE__VALID_URL = λ.StrLiteral("https?://video\\.arnes\\.si/(?:[a-z]{2}/)?(?:watch|embed|api/(?:asset|public/video))/(?P<id>[0-9a-zA-Z]{12})")
			ArnesIE__BASE_URL = λ.StrLiteral("https://video.arnes.si")
			ArnesIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒchannel    λ.Object
						ϒchannel_id λ.Object
						ϒformats    λ.Object
						ϒmedia      λ.Object
						ϒmedia_url  λ.Object
						ϒself       = λargs[0]
						ϒthumbnail  λ.Object
						ϒtitle      λ.Object
						ϒurl        = λargs[1]
						ϒvideo      λ.Object
						ϒvideo_id   λ.Object
						τmp0        λ.Object
						τmp1        λ.Object
					)
					ϒvideo_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒvideo = λ.GetItem(λ.Calm(ϒself, "_download_json", λ.Add(λ.Add(λ.GetAttr(ϒself, "_BASE_URL", nil), λ.StrLiteral("/api/public/video/")), ϒvideo_id), ϒvideo_id), λ.StrLiteral("data"))
					ϒtitle = λ.GetItem(ϒvideo, λ.StrLiteral("title"))
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, func() λ.Object {
						if λv := λ.Calm(ϒvideo, "get", λ.StrLiteral("media")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.NewList()
						}
					}())
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒmedia = τmp1
						ϒmedia_url = λ.Calm(ϒmedia, "get", λ.StrLiteral("url"))
						if !λ.IsTrue(ϒmedia_url) {
							continue
						}
						λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
							"url":         λ.Add(λ.GetAttr(ϒself, "_BASE_URL", nil), ϒmedia_url),
							"format_id":   λ.Cal(ϒremove_start, λ.Calm(ϒmedia, "get", λ.StrLiteral("format")), λ.StrLiteral("FORMAT_")),
							"format_note": λ.Calm(ϒmedia, "get", λ.StrLiteral("formatTranslation")),
							"width":       λ.Cal(ϒint_or_none, λ.Calm(ϒmedia, "get", λ.StrLiteral("width"))),
							"height":      λ.Cal(ϒint_or_none, λ.Calm(ϒmedia, "get", λ.StrLiteral("height"))),
						}))
					}
					λ.Calm(ϒself, "_sort_formats", ϒformats)
					ϒchannel = func() λ.Object {
						if λv := λ.Calm(ϒvideo, "get", λ.StrLiteral("channel")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.DictLiteral(map[λ.Object]λ.Object{})
						}
					}()
					ϒchannel_id = λ.Calm(ϒchannel, "get", λ.StrLiteral("url"))
					ϒthumbnail = λ.Calm(ϒvideo, "get", λ.StrLiteral("thumbnailUrl"))
					return λ.DictLiteral(map[string]λ.Object{
						"id":          ϒvideo_id,
						"title":       ϒtitle,
						"formats":     ϒformats,
						"thumbnail":   λ.Add(λ.GetAttr(ϒself, "_BASE_URL", nil), ϒthumbnail),
						"description": λ.Calm(ϒvideo, "get", λ.StrLiteral("description")),
						"license":     λ.Calm(ϒvideo, "get", λ.StrLiteral("license")),
						"creator":     λ.Calm(ϒvideo, "get", λ.StrLiteral("author")),
						"timestamp":   λ.Cal(ϒparse_iso8601, λ.Calm(ϒvideo, "get", λ.StrLiteral("creationTime"))),
						"channel":     λ.Calm(ϒchannel, "get", λ.StrLiteral("name")),
						"channel_id":  ϒchannel_id,
						"channel_url": func() λ.Object {
							if λ.IsTrue(ϒchannel_id) {
								return λ.Add(λ.Add(λ.GetAttr(ϒself, "_BASE_URL", nil), λ.StrLiteral("/?channel=")), ϒchannel_id)
							} else {
								return λ.None
							}
						}(),
						"duration":   λ.Cal(ϒfloat_or_none, λ.Calm(ϒvideo, "get", λ.StrLiteral("duration")), λ.IntLiteral(1000)),
						"view_count": λ.Cal(ϒint_or_none, λ.Calm(ϒvideo, "get", λ.StrLiteral("views"))),
						"tags":       λ.Calm(ϒvideo, "get", λ.StrLiteral("hashtags")),
						"start_time": λ.Cal(ϒint_or_none, λ.GetItem(λ.Calm(λ.Cal(ϒcompat_parse_qs, λ.GetAttr(λ.Cal(ϒcompat_urllib_parse_urlparse, ϒurl), "query", nil)), "get", λ.StrLiteral("t"), λ.NewList(λ.None)), λ.IntLiteral(0))),
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"IE_NAME":       ArnesIE_IE_NAME,
				"_BASE_URL":     ArnesIE__BASE_URL,
				"_VALID_URL":    ArnesIE__VALID_URL,
				"_real_extract": ArnesIE__real_extract,
			})
		}())
	})
}
