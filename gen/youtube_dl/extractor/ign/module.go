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
 * ign/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/ign.py
 */

package ign

import (
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	HEADRequest                   λ.Object
	IGNArticleIE                  λ.Object
	IGNBaseIE                     λ.Object
	IGNIE                         λ.Object
	IGNVideoIE                    λ.Object
	InfoExtractor                 λ.Object
	ϒcompat_parse_qs              λ.Object
	ϒcompat_urllib_parse_urlparse λ.Object
	ϒdetermine_ext                λ.Object
	ϒint_or_none                  λ.Object
	ϒparse_iso8601                λ.Object
	ϒstrip_or_none                λ.Object
	ϒtry_get                      λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_parse_qs = Ωcompat.ϒcompat_parse_qs
		ϒcompat_urllib_parse_urlparse = Ωcompat.ϒcompat_urllib_parse_urlparse
		HEADRequest = Ωutils.HEADRequest
		ϒdetermine_ext = Ωutils.ϒdetermine_ext
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒparse_iso8601 = Ωutils.ϒparse_iso8601
		ϒstrip_or_none = Ωutils.ϒstrip_or_none
		ϒtry_get = Ωutils.ϒtry_get
		IGNBaseIE = λ.Cal(λ.TypeType, λ.StrLiteral("IGNBaseIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				IGNBaseIE__call_api λ.Object
			)
			IGNBaseIE__call_api = λ.NewFunction("_call_api",
				[]λ.Param{
					{Name: "self"},
					{Name: "slug"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒself = λargs[0]
						ϒslug = λargs[1]
					)
					return λ.Calm(ϒself, "_download_json", λ.Calm(λ.StrLiteral("http://apis.ign.com/{0}/v3/{0}s/slug/{1}"), "format", λ.GetAttr(ϒself, "_PAGE_TYPE", nil), ϒslug), ϒslug)
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_call_api": IGNBaseIE__call_api,
			})
		}())
		IGNIE = λ.Cal(λ.TypeType, λ.StrLiteral("IGNIE"), λ.NewTuple(IGNBaseIE), func() λ.Dict {
			var (
				IGNIE_IE_NAME       λ.Object
				IGNIE__PAGE_TYPE    λ.Object
				IGNIE__VALID_URL    λ.Object
				IGNIE__real_extract λ.Object
			)
			IGNIE__VALID_URL = λ.StrLiteral("https?://(?:.+?\\.ign|www\\.pcmag)\\.com/videos/(?:\\d{4}/\\d{2}/\\d{2}/)?(?P<id>[^/?&#]+)")
			IGNIE_IE_NAME = λ.StrLiteral("ign.com")
			IGNIE__PAGE_TYPE = λ.StrLiteral("video")
			IGNIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒasset         λ.Object
						ϒasset_url     λ.Object
						ϒdisplay_id    λ.Object
						ϒdisplay_name  λ.Object
						ϒf4m_url       λ.Object
						ϒformats       λ.Object
						ϒm3u8_url      λ.Object
						ϒmetadata      λ.Object
						ϒmezzanine_url λ.Object
						ϒrefs          λ.Object
						ϒself          = λargs[0]
						ϒtag           λ.Object
						ϒtags          λ.Object
						ϒthumbnail     λ.Object
						ϒthumbnail_url λ.Object
						ϒthumbnails    λ.Object
						ϒtitle         λ.Object
						ϒurl           = λargs[1]
						ϒvideo         λ.Object
						ϒvideo_id      λ.Object
						τmp0           λ.Object
						τmp1           λ.Object
					)
					ϒdisplay_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒvideo = λ.Calm(ϒself, "_call_api", ϒdisplay_id)
					ϒvideo_id = λ.GetItem(ϒvideo, λ.StrLiteral("videoId"))
					ϒmetadata = λ.GetItem(ϒvideo, λ.StrLiteral("metadata"))
					ϒtitle = func() λ.Object {
						if λv := λ.Calm(ϒmetadata, "get", λ.StrLiteral("longTitle")); λ.IsTrue(λv) {
							return λv
						} else if λv := λ.Calm(ϒmetadata, "get", λ.StrLiteral("title")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.GetItem(ϒmetadata, λ.StrLiteral("name"))
						}
					}()
					ϒformats = λ.NewList()
					ϒrefs = func() λ.Object {
						if λv := λ.Calm(ϒvideo, "get", λ.StrLiteral("refs")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.DictLiteral(map[λ.Object]λ.Object{})
						}
					}()
					ϒm3u8_url = λ.Calm(ϒrefs, "get", λ.StrLiteral("m3uUrl"))
					if λ.IsTrue(ϒm3u8_url) {
						λ.Calm(ϒformats, "extend", λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
							ϒm3u8_url,
							ϒvideo_id,
							λ.StrLiteral("mp4"),
							λ.StrLiteral("m3u8_native"),
						), λ.KWArgs{
							{Name: "m3u8_id", Value: λ.StrLiteral("hls")},
							{Name: "fatal", Value: λ.False},
						}))
					}
					ϒf4m_url = λ.Calm(ϒrefs, "get", λ.StrLiteral("f4mUrl"))
					if λ.IsTrue(ϒf4m_url) {
						λ.Calm(ϒformats, "extend", λ.Call(λ.GetAttr(ϒself, "_extract_f4m_formats", nil), λ.NewArgs(
							ϒf4m_url,
							ϒvideo_id,
						), λ.KWArgs{
							{Name: "f4m_id", Value: λ.StrLiteral("hds")},
							{Name: "fatal", Value: λ.False},
						}))
					}
					τmp0 = λ.Cal(λ.BuiltinIter, func() λ.Object {
						if λv := λ.Calm(ϒvideo, "get", λ.StrLiteral("assets")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.NewList()
						}
					}())
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒasset = τmp1
						ϒasset_url = λ.Calm(ϒasset, "get", λ.StrLiteral("url"))
						if !λ.IsTrue(ϒasset_url) {
							continue
						}
						λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
							"url":    ϒasset_url,
							"tbr":    λ.Cal(ϒint_or_none, λ.Calm(ϒasset, "get", λ.StrLiteral("bitrate")), λ.IntLiteral(1000)),
							"fps":    λ.Cal(ϒint_or_none, λ.Calm(ϒasset, "get", λ.StrLiteral("frame_rate"))),
							"height": λ.Cal(ϒint_or_none, λ.Calm(ϒasset, "get", λ.StrLiteral("height"))),
							"width":  λ.Cal(ϒint_or_none, λ.Calm(ϒasset, "get", λ.StrLiteral("width"))),
						}))
					}
					ϒmezzanine_url = λ.Cal(ϒtry_get, ϒvideo, λ.NewFunction("<lambda>",
						[]λ.Param{
							{Name: "x"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒx = λargs[0]
							)
							return λ.GetItem(λ.GetItem(ϒx, λ.StrLiteral("system")), λ.StrLiteral("mezzanineUrl"))
						}))
					if λ.IsTrue(ϒmezzanine_url) {
						λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
							"ext":        λ.Cal(ϒdetermine_ext, ϒmezzanine_url, λ.StrLiteral("mp4")),
							"format_id":  λ.StrLiteral("mezzanine"),
							"preference": λ.IntLiteral(1),
							"url":        ϒmezzanine_url,
						}))
					}
					λ.Calm(ϒself, "_sort_formats", ϒformats)
					ϒthumbnails = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, func() λ.Object {
						if λv := λ.Calm(ϒvideo, "get", λ.StrLiteral("thumbnails")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.NewList()
						}
					}())
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒthumbnail = τmp1
						ϒthumbnail_url = λ.Calm(ϒthumbnail, "get", λ.StrLiteral("url"))
						if !λ.IsTrue(ϒthumbnail_url) {
							continue
						}
						λ.Calm(ϒthumbnails, "append", λ.DictLiteral(map[string]λ.Object{
							"url": ϒthumbnail_url,
						}))
					}
					ϒtags = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, func() λ.Object {
						if λv := λ.Calm(ϒvideo, "get", λ.StrLiteral("tags")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.NewList()
						}
					}())
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒtag = τmp1
						ϒdisplay_name = λ.Calm(ϒtag, "get", λ.StrLiteral("displayName"))
						if !λ.IsTrue(ϒdisplay_name) {
							continue
						}
						λ.Calm(ϒtags, "append", ϒdisplay_name)
					}
					return λ.DictLiteral(map[string]λ.Object{
						"id":          ϒvideo_id,
						"title":       ϒtitle,
						"description": λ.Cal(ϒstrip_or_none, λ.Calm(ϒmetadata, "get", λ.StrLiteral("description"))),
						"timestamp":   λ.Cal(ϒparse_iso8601, λ.Calm(ϒmetadata, "get", λ.StrLiteral("publishDate"))),
						"duration":    λ.Cal(ϒint_or_none, λ.Calm(ϒmetadata, "get", λ.StrLiteral("duration"))),
						"display_id":  ϒdisplay_id,
						"thumbnails":  ϒthumbnails,
						"formats":     ϒformats,
						"tags":        ϒtags,
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"IE_NAME":       IGNIE_IE_NAME,
				"_PAGE_TYPE":    IGNIE__PAGE_TYPE,
				"_VALID_URL":    IGNIE__VALID_URL,
				"_real_extract": IGNIE__real_extract,
			})
		}())
		IGNVideoIE = λ.Cal(λ.TypeType, λ.StrLiteral("IGNVideoIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				IGNVideoIE__VALID_URL λ.Object
			)
			IGNVideoIE__VALID_URL = λ.StrLiteral("https?://.+?\\.ign\\.com/(?:[a-z]{2}/)?[^/]+/(?P<id>\\d+)/(?:video|trailer)/")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": IGNVideoIE__VALID_URL,
			})
		}())
		IGNArticleIE = λ.Cal(λ.TypeType, λ.StrLiteral("IGNArticleIE"), λ.NewTuple(IGNBaseIE), func() λ.Dict {
			var (
				IGNArticleIE__VALID_URL λ.Object
			)
			IGNArticleIE__VALID_URL = λ.StrLiteral("https?://.+?\\.ign\\.com/(?:articles(?:/\\d{4}/\\d{2}/\\d{2})?|(?:[a-z]{2}/)?feature/\\d+)/(?P<id>[^/?&#]+)")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": IGNArticleIE__VALID_URL,
			})
		}())
	})
}
