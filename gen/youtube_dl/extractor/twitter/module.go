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
 * twitter/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/twitter.py
 */

package twitter

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωperiscope "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/periscope"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError                λ.Object
	InfoExtractor                 λ.Object
	PeriscopeBaseIE               λ.Object
	PeriscopeIE                   λ.Object
	TwitterAmplifyIE              λ.Object
	TwitterBaseIE                 λ.Object
	TwitterBroadcastIE            λ.Object
	TwitterCardIE                 λ.Object
	TwitterIE                     λ.Object
	ϒcompat_HTTPError             λ.Object
	ϒcompat_parse_qs              λ.Object
	ϒcompat_urllib_parse_unquote  λ.Object
	ϒcompat_urllib_parse_urlparse λ.Object
	ϒdict_get                     λ.Object
	ϒfloat_or_none                λ.Object
	ϒint_or_none                  λ.Object
	ϒstrip_or_none                λ.Object
	ϒtry_get                      λ.Object
	ϒunified_timestamp            λ.Object
	ϒupdate_url_query             λ.Object
	ϒxpath_text                   λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_HTTPError = Ωcompat.ϒcompat_HTTPError
		ϒcompat_parse_qs = Ωcompat.ϒcompat_parse_qs
		ϒcompat_urllib_parse_unquote = Ωcompat.ϒcompat_urllib_parse_unquote
		ϒcompat_urllib_parse_urlparse = Ωcompat.ϒcompat_urllib_parse_urlparse
		ϒdict_get = Ωutils.ϒdict_get
		ExtractorError = Ωutils.ExtractorError
		ϒfloat_or_none = Ωutils.ϒfloat_or_none
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒtry_get = Ωutils.ϒtry_get
		ϒstrip_or_none = Ωutils.ϒstrip_or_none
		ϒunified_timestamp = Ωutils.ϒunified_timestamp
		ϒupdate_url_query = Ωutils.ϒupdate_url_query
		ϒxpath_text = Ωutils.ϒxpath_text
		PeriscopeBaseIE = Ωperiscope.PeriscopeBaseIE
		PeriscopeIE = Ωperiscope.PeriscopeIE
		TwitterBaseIE = λ.Cal(λ.TypeType, λ.StrLiteral("TwitterBaseIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				TwitterBaseIE__BASE_REGEX                     λ.Object
				TwitterBaseIE__extract_formats_from_vmap_url  λ.Object
				TwitterBaseIE__extract_variant_formats        λ.Object
				TwitterBaseIE__search_dimensions_in_video_url λ.Object
			)
			TwitterBaseIE__BASE_REGEX = λ.StrLiteral("https?://(?:(?:www|m(?:obile)?)\\.)?twitter\\.com/")
			TwitterBaseIE__extract_variant_formats = λ.NewFunction("_extract_variant_formats",
				[]λ.Param{
					{Name: "self"},
					{Name: "variant"},
					{Name: "video_id"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒf           λ.Object
						ϒself        = λargs[0]
						ϒtbr         λ.Object
						ϒvariant     = λargs[1]
						ϒvariant_url λ.Object
						ϒvideo_id    = λargs[2]
					)
					ϒvariant_url = λ.Calm(ϒvariant, "get", λ.StrLiteral("url"))
					if !λ.IsTrue(ϒvariant_url) {
						return λ.NewList()
					} else {
						if λ.Contains(ϒvariant_url, λ.StrLiteral(".m3u8")) {
							return λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
								ϒvariant_url,
								ϒvideo_id,
								λ.StrLiteral("mp4"),
								λ.StrLiteral("m3u8_native"),
							), λ.KWArgs{
								{Name: "m3u8_id", Value: λ.StrLiteral("hls")},
								{Name: "fatal", Value: λ.False},
							})
						} else {
							ϒtbr = func() λ.Object {
								if λv := λ.Cal(ϒint_or_none, λ.Cal(ϒdict_get, ϒvariant, λ.NewTuple(
									λ.StrLiteral("bitrate"),
									λ.StrLiteral("bit_rate"),
								)), λ.IntLiteral(1000)); λ.IsTrue(λv) {
									return λv
								} else {
									return λ.None
								}
							}()
							ϒf = λ.DictLiteral(map[string]λ.Object{
								"url": ϒvariant_url,
								"format_id": λ.Add(λ.StrLiteral("http"), func() λ.Object {
									if λ.IsTrue(ϒtbr) {
										return λ.Mod(λ.StrLiteral("-%d"), ϒtbr)
									} else {
										return λ.StrLiteral("")
									}
								}()),
								"tbr": ϒtbr,
							})
							λ.Calm(ϒself, "_search_dimensions_in_video_url", ϒf, ϒvariant_url)
							return λ.NewList(ϒf)
						}
					}
					return λ.None
				})
			TwitterBaseIE__extract_formats_from_vmap_url = λ.NewFunction("_extract_formats_from_vmap_url",
				[]λ.Param{
					{Name: "self"},
					{Name: "vmap_url"},
					{Name: "video_id"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒformats       λ.Object
						ϒself          = λargs[0]
						ϒurls          λ.Object
						ϒvideo_id      = λargs[2]
						ϒvideo_url     λ.Object
						ϒvideo_variant λ.Object
						ϒvmap_data     λ.Object
						ϒvmap_url      = λargs[1]
						τmp0           λ.Object
						τmp1           λ.Object
					)
					ϒvmap_data = λ.Calm(ϒself, "_download_xml", ϒvmap_url, ϒvideo_id)
					ϒformats = λ.NewList()
					ϒurls = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(ϒvmap_data, "findall", λ.StrLiteral(".//{http://twitter.com/schema/videoVMapV2.xsd}videoVariant")))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒvideo_variant = τmp1
						λ.SetItem(λ.GetAttr(ϒvideo_variant, "attrib", nil), λ.StrLiteral("url"), λ.Cal(ϒcompat_urllib_parse_unquote, λ.GetItem(λ.GetAttr(ϒvideo_variant, "attrib", nil), λ.StrLiteral("url"))))
						λ.Calm(ϒurls, "append", λ.GetItem(λ.GetAttr(ϒvideo_variant, "attrib", nil), λ.StrLiteral("url")))
						λ.Calm(ϒformats, "extend", λ.Calm(ϒself, "_extract_variant_formats", λ.GetAttr(ϒvideo_variant, "attrib", nil), ϒvideo_id))
					}
					ϒvideo_url = λ.Cal(ϒstrip_or_none, λ.Cal(ϒxpath_text, ϒvmap_data, λ.StrLiteral(".//MediaFile")))
					if !λ.Contains(ϒurls, ϒvideo_url) {
						λ.Calm(ϒformats, "extend", λ.Calm(ϒself, "_extract_variant_formats", λ.DictLiteral(map[string]λ.Object{
							"url": ϒvideo_url,
						}), ϒvideo_id))
					}
					return ϒformats
				})
			TwitterBaseIE__search_dimensions_in_video_url = λ.NewFunction("_search_dimensions_in_video_url",
				[]λ.Param{
					{Name: "a_format"},
					{Name: "video_url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒa_format  = λargs[0]
						ϒm         λ.Object
						ϒvideo_url = λargs[1]
					)
					ϒm = λ.Cal(Ωre.ϒsearch, λ.StrLiteral("/(?P<width>\\d+)x(?P<height>\\d+)/"), ϒvideo_url)
					if λ.IsTrue(ϒm) {
						λ.Calm(ϒa_format, "update", λ.DictLiteral(map[string]λ.Object{
							"width":  λ.Cal(λ.IntType, λ.Calm(ϒm, "group", λ.StrLiteral("width"))),
							"height": λ.Cal(λ.IntType, λ.Calm(ϒm, "group", λ.StrLiteral("height"))),
						}))
					}
					return λ.None
				})
			TwitterBaseIE__search_dimensions_in_video_url = λ.Cal(λ.StaticMethodType, TwitterBaseIE__search_dimensions_in_video_url)
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_BASE_REGEX":                     TwitterBaseIE__BASE_REGEX,
				"_extract_formats_from_vmap_url":  TwitterBaseIE__extract_formats_from_vmap_url,
				"_extract_variant_formats":        TwitterBaseIE__extract_variant_formats,
				"_search_dimensions_in_video_url": TwitterBaseIE__search_dimensions_in_video_url,
			})
		}())
		TwitterCardIE = λ.Cal(λ.TypeType, λ.StrLiteral("TwitterCardIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				TwitterCardIE__VALID_URL λ.Object
			)
			TwitterCardIE__VALID_URL = λ.Add(λ.GetAttr(TwitterBaseIE, "_BASE_REGEX", nil), λ.StrLiteral("i/(?:cards/tfw/v1|videos(?:/tweet)?)/(?P<id>\\d+)"))
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": TwitterCardIE__VALID_URL,
			})
		}())
		TwitterIE = λ.Cal(λ.TypeType, λ.StrLiteral("TwitterIE"), λ.NewTuple(TwitterBaseIE), func() λ.Dict {
			var (
				TwitterIE__VALID_URL λ.Object
			)
			TwitterIE__VALID_URL = λ.Add(λ.GetAttr(TwitterBaseIE, "_BASE_REGEX", nil), λ.StrLiteral("(?:(?:i/web|[^/]+)/status|statuses)/(?P<id>\\d+)"))
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": TwitterIE__VALID_URL,
			})
		}())
		TwitterAmplifyIE = λ.Cal(λ.TypeType, λ.StrLiteral("TwitterAmplifyIE"), λ.NewTuple(TwitterBaseIE), func() λ.Dict {
			var (
				TwitterAmplifyIE_IE_NAME       λ.Object
				TwitterAmplifyIE__VALID_URL    λ.Object
				TwitterAmplifyIE__real_extract λ.Object
			)
			TwitterAmplifyIE_IE_NAME = λ.StrLiteral("twitter:amplify")
			TwitterAmplifyIE__VALID_URL = λ.StrLiteral("https?://amp\\.twimg\\.com/v/(?P<id>[0-9a-f\\-]{36})")
			TwitterAmplifyIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒ_find_dimension λ.Object
						ϒformats         λ.Object
						ϒself            = λargs[0]
						ϒthumbnail       λ.Object
						ϒthumbnail_h     λ.Object
						ϒthumbnail_w     λ.Object
						ϒthumbnails      λ.Object
						ϒurl             = λargs[1]
						ϒvideo_h         λ.Object
						ϒvideo_id        λ.Object
						ϒvideo_w         λ.Object
						ϒvmap_url        λ.Object
						ϒwebpage         λ.Object
						τmp0             λ.Object
					)
					ϒvideo_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒwebpage = λ.Calm(ϒself, "_download_webpage", ϒurl, ϒvideo_id)
					ϒvmap_url = λ.Calm(ϒself, "_html_search_meta", λ.StrLiteral("twitter:amplify:vmap"), ϒwebpage, λ.StrLiteral("vmap url"))
					ϒformats = λ.Calm(ϒself, "_extract_formats_from_vmap_url", ϒvmap_url, ϒvideo_id)
					ϒthumbnails = λ.NewList()
					ϒthumbnail = λ.Call(λ.GetAttr(ϒself, "_html_search_meta", nil), λ.NewArgs(
						λ.StrLiteral("twitter:image:src"),
						ϒwebpage,
						λ.StrLiteral("thumbnail"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					})
					ϒ_find_dimension = λ.NewFunction("_find_dimension",
						[]λ.Param{
							{Name: "target"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒh      λ.Object
								ϒtarget = λargs[0]
								ϒw      λ.Object
							)
							ϒw = λ.Cal(ϒint_or_none, λ.Call(λ.GetAttr(ϒself, "_html_search_meta", nil), λ.NewArgs(
								λ.Mod(λ.StrLiteral("twitter:%s:width"), ϒtarget),
								ϒwebpage,
							), λ.KWArgs{
								{Name: "fatal", Value: λ.False},
							}))
							ϒh = λ.Cal(ϒint_or_none, λ.Call(λ.GetAttr(ϒself, "_html_search_meta", nil), λ.NewArgs(
								λ.Mod(λ.StrLiteral("twitter:%s:height"), ϒtarget),
								ϒwebpage,
							), λ.KWArgs{
								{Name: "fatal", Value: λ.False},
							}))
							return λ.NewTuple(
								ϒw,
								ϒh,
							)
						})
					if λ.IsTrue(ϒthumbnail) {
						τmp0 = λ.Cal(ϒ_find_dimension, λ.StrLiteral("image"))
						ϒthumbnail_w = λ.GetItem(τmp0, λ.IntLiteral(0))
						ϒthumbnail_h = λ.GetItem(τmp0, λ.IntLiteral(1))
						λ.Calm(ϒthumbnails, "append", λ.DictLiteral(map[string]λ.Object{
							"url":    ϒthumbnail,
							"width":  ϒthumbnail_w,
							"height": ϒthumbnail_h,
						}))
					}
					τmp0 = λ.Cal(ϒ_find_dimension, λ.StrLiteral("player"))
					ϒvideo_w = λ.GetItem(τmp0, λ.IntLiteral(0))
					ϒvideo_h = λ.GetItem(τmp0, λ.IntLiteral(1))
					λ.Calm(λ.GetItem(ϒformats, λ.IntLiteral(0)), "update", λ.DictLiteral(map[string]λ.Object{
						"width":  ϒvideo_w,
						"height": ϒvideo_h,
					}))
					return λ.DictLiteral(map[string]λ.Object{
						"id":         ϒvideo_id,
						"title":      λ.StrLiteral("Twitter Video"),
						"formats":    ϒformats,
						"thumbnails": ϒthumbnails,
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"IE_NAME":       TwitterAmplifyIE_IE_NAME,
				"_VALID_URL":    TwitterAmplifyIE__VALID_URL,
				"_real_extract": TwitterAmplifyIE__real_extract,
			})
		}())
		TwitterBroadcastIE = λ.Cal(λ.TypeType, λ.StrLiteral("TwitterBroadcastIE"), λ.NewTuple(
			TwitterBaseIE,
			PeriscopeBaseIE,
		), func() λ.Dict {
			var (
				TwitterBroadcastIE__VALID_URL λ.Object
			)
			TwitterBroadcastIE__VALID_URL = λ.Add(λ.GetAttr(TwitterBaseIE, "_BASE_REGEX", nil), λ.StrLiteral("i/broadcasts/(?P<id>[0-9a-zA-Z]{13})"))
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": TwitterBroadcastIE__VALID_URL,
			})
		}())
	})
}
