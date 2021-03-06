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
 * condenast/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/condenast.py
 */

package condenast

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	CondeNastIE                   λ.Object
	InfoExtractor                 λ.Object
	ϒcompat_urllib_parse_urlparse λ.Object
	ϒdetermine_ext                λ.Object
	ϒextract_attributes           λ.Object
	ϒint_or_none                  λ.Object
	ϒjs_to_json                   λ.Object
	ϒmimetype2ext                 λ.Object
	ϒparse_iso8601                λ.Object
	ϒstrip_or_none                λ.Object
	ϒtry_get                      λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_urllib_parse_urlparse = Ωcompat.ϒcompat_urllib_parse_urlparse
		ϒdetermine_ext = Ωutils.ϒdetermine_ext
		ϒextract_attributes = Ωutils.ϒextract_attributes
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒjs_to_json = Ωutils.ϒjs_to_json
		ϒmimetype2ext = Ωutils.ϒmimetype2ext
		ϒparse_iso8601 = Ωutils.ϒparse_iso8601
		ϒstrip_or_none = Ωutils.ϒstrip_or_none
		ϒtry_get = Ωutils.ϒtry_get
		CondeNastIE = λ.Cal(λ.TypeType, λ.StrLiteral("CondeNastIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				CondeNastIE__SITES         λ.Object
				CondeNastIE__VALID_URL     λ.Object
				CondeNastIE__extract_video λ.Object
				CondeNastIE__real_extract  λ.Object
			)
			CondeNastIE__SITES = λ.DictLiteral(map[string]string{
				"allure":              "Allure",
				"architecturaldigest": "Architectural Digest",
				"arstechnica":         "Ars Technica",
				"bonappetit":          "Bon Appétit",
				"brides":              "Brides",
				"cnevids":             "Condé Nast",
				"cntraveler":          "Condé Nast Traveler",
				"details":             "Details",
				"epicurious":          "Epicurious",
				"glamour":             "Glamour",
				"golfdigest":          "Golf Digest",
				"gq":                  "GQ",
				"newyorker":           "The New Yorker",
				"self":                "SELF",
				"teenvogue":           "Teen Vogue",
				"vanityfair":          "Vanity Fair",
				"vogue":               "Vogue",
				"wired":               "WIRED",
				"wmagazine":           "W Magazine",
			})
			CondeNastIE__VALID_URL = λ.Mod(λ.StrLiteral("(?x)https?://(?:video|www|player(?:-backend)?)\\.(?:%s)\\.com/\n        (?:\n            (?:\n                embed(?:js)?|\n                (?:script|inline)/video\n            )/(?P<id>[0-9a-f]{24})(?:/(?P<player_id>[0-9a-f]{24}))?(?:.+?\\btarget=(?P<target>[^&]+))?|\n            (?P<type>watch|series|video)/(?P<display_id>[^/?#]+)\n        )"), λ.Calm(λ.StrLiteral("|"), "join", λ.Calm(CondeNastIE__SITES, "keys")))
			CondeNastIE__extract_video = λ.NewFunction("_extract_video",
				[]λ.Param{
					{Name: "self"},
					{Name: "params"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒcaption     λ.Object
						ϒcaption_url λ.Object
						ϒext         λ.Object
						ϒfdata       λ.Object
						ϒformats     λ.Object
						ϒinfo_page   λ.Object
						ϒparams      = λargs[1]
						ϒquality     λ.Object
						ϒquery       λ.Object
						ϒself        = λargs[0]
						ϒsrc         λ.Object
						ϒsubtitles   λ.Object
						ϒt           λ.Object
						ϒtitle       λ.Object
						ϒvideo_id    λ.Object
						ϒvideo_info  λ.Object
						τmp0         λ.Object
						τmp1         λ.Object
						τmp2         λ.Object
					)
					ϒvideo_id = λ.GetItem(ϒparams, λ.StrLiteral("videoId"))
					ϒvideo_info = λ.None
					ϒquery = λ.Calm(ϒparams, "copy")
					λ.SetItem(ϒquery, λ.StrLiteral("embedType"), λ.StrLiteral("inline"))
					ϒinfo_page = λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
						λ.StrLiteral("http://player.cnevids.com/embed-api.json"),
						ϒvideo_id,
						λ.StrLiteral("Downloading embed info"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
						{Name: "query", Value: ϒquery},
					})
					if !λ.IsTrue(ϒinfo_page) {
						if λ.IsTrue(λ.Calm(ϒparams, "get", λ.StrLiteral("playerId"))) {
							ϒinfo_page = λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
								λ.StrLiteral("http://player.cnevids.com/player/video.js"),
								ϒvideo_id,
								λ.StrLiteral("Downloading video info"),
							), λ.KWArgs{
								{Name: "fatal", Value: λ.False},
								{Name: "query", Value: ϒparams},
							})
						}
					}
					if λ.IsTrue(ϒinfo_page) {
						ϒvideo_info = λ.Calm(ϒinfo_page, "get", λ.StrLiteral("video"))
					}
					if !λ.IsTrue(ϒvideo_info) {
						ϒinfo_page = λ.Call(λ.GetAttr(ϒself, "_download_webpage", nil), λ.NewArgs(
							λ.StrLiteral("http://player.cnevids.com/player/loader.js"),
							ϒvideo_id,
							λ.StrLiteral("Downloading loader info"),
						), λ.KWArgs{
							{Name: "query", Value: ϒparams},
						})
					}
					if !λ.IsTrue(ϒvideo_info) {
						ϒinfo_page = λ.Call(λ.GetAttr(ϒself, "_download_webpage", nil), λ.NewArgs(
							λ.Mod(λ.StrLiteral("https://player.cnevids.com/inline/video/%s.js"), ϒvideo_id),
							ϒvideo_id,
							λ.StrLiteral("Downloading inline info"),
						), λ.KWArgs{
							{Name: "query", Value: λ.DictLiteral(map[string]λ.Object{
								"target": λ.Calm(ϒparams, "get", λ.StrLiteral("target"), λ.StrLiteral("embedplayer")),
							})},
						})
					}
					if !λ.IsTrue(ϒvideo_info) {
						ϒvideo_info = λ.GetItem(λ.Call(λ.GetAttr(ϒself, "_parse_json", nil), λ.NewArgs(
							λ.Calm(ϒself, "_search_regex", λ.StrLiteral("(?s)var\\s+config\\s*=\\s*({.+?});"), ϒinfo_page, λ.StrLiteral("config")),
							ϒvideo_id,
						), λ.KWArgs{
							{Name: "transform_source", Value: ϒjs_to_json},
						}), λ.StrLiteral("video"))
					}
					ϒtitle = λ.GetItem(ϒvideo_info, λ.StrLiteral("title"))
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.GetItem(ϒvideo_info, λ.StrLiteral("sources")))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒfdata = τmp1
						ϒsrc = λ.Calm(ϒfdata, "get", λ.StrLiteral("src"))
						if !λ.IsTrue(ϒsrc) {
							continue
						}
						ϒext = func() λ.Object {
							if λv := λ.Cal(ϒmimetype2ext, λ.Calm(ϒfdata, "get", λ.StrLiteral("type"))); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Cal(ϒdetermine_ext, ϒsrc)
							}
						}()
						if λ.IsTrue(λ.Eq(ϒext, λ.StrLiteral("m3u8"))) {
							λ.Calm(ϒformats, "extend", λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
								ϒsrc,
								ϒvideo_id,
								λ.StrLiteral("mp4"),
							), λ.KWArgs{
								{Name: "entry_protocol", Value: λ.StrLiteral("m3u8_native")},
								{Name: "m3u8_id", Value: λ.StrLiteral("hls")},
								{Name: "fatal", Value: λ.False},
							}))
							continue
						}
						ϒquality = λ.Calm(ϒfdata, "get", λ.StrLiteral("quality"))
						λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
							"format_id": λ.Add(ϒext, func() λ.Object {
								if λ.IsTrue(ϒquality) {
									return λ.Mod(λ.StrLiteral("-%s"), ϒquality)
								} else {
									return λ.StrLiteral("")
								}
							}()),
							"url": ϒsrc,
							"ext": ϒext,
							"quality": func() λ.Object {
								if λ.IsTrue(λ.Eq(ϒquality, λ.StrLiteral("high"))) {
									return λ.IntLiteral(1)
								} else {
									return λ.IntLiteral(0)
								}
							}(),
						}))
					}
					λ.Calm(ϒself, "_sort_formats", ϒformats)
					ϒsubtitles = λ.DictLiteral(map[λ.Object]λ.Object{})
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(λ.Calm(ϒvideo_info, "get", λ.StrLiteral("captions"), λ.DictLiteral(map[λ.Object]λ.Object{})), "items"))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						τmp2 = λ.UnpackIterable(τmp1, 2)
						ϒt = λ.GetItem(τmp2, λ.IntLiteral(0))
						ϒcaption = λ.GetItem(τmp2, λ.IntLiteral(1))
						ϒcaption_url = λ.Calm(ϒcaption, "get", λ.StrLiteral("src"))
						if !λ.IsTrue(func() λ.Object {
							if λv := λ.NewBool(λ.Contains(λ.NewTuple(
								λ.StrLiteral("vtt"),
								λ.StrLiteral("srt"),
								λ.StrLiteral("tml"),
							), ϒt)); !λ.IsTrue(λv) {
								return λv
							} else {
								return ϒcaption_url
							}
						}()) {
							continue
						}
						λ.Calm(λ.Calm(ϒsubtitles, "setdefault", λ.StrLiteral("en"), λ.NewList()), "append", λ.DictLiteral(map[string]λ.Object{
							"url": ϒcaption_url,
						}))
					}
					return λ.DictLiteral(map[string]λ.Object{
						"id":         ϒvideo_id,
						"formats":    ϒformats,
						"title":      ϒtitle,
						"thumbnail":  λ.Calm(ϒvideo_info, "get", λ.StrLiteral("poster_frame")),
						"uploader":   λ.Calm(ϒvideo_info, "get", λ.StrLiteral("brand")),
						"duration":   λ.Cal(ϒint_or_none, λ.Calm(ϒvideo_info, "get", λ.StrLiteral("duration"))),
						"tags":       λ.Calm(ϒvideo_info, "get", λ.StrLiteral("tags")),
						"series":     λ.Calm(ϒvideo_info, "get", λ.StrLiteral("series_title")),
						"season":     λ.Calm(ϒvideo_info, "get", λ.StrLiteral("season_title")),
						"timestamp":  λ.Cal(ϒparse_iso8601, λ.Calm(ϒvideo_info, "get", λ.StrLiteral("premiere_date"))),
						"categories": λ.Calm(ϒvideo_info, "get", λ.StrLiteral("categories")),
						"subtitles":  ϒsubtitles,
					})
				})
			CondeNastIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdisplay_id λ.Object
						ϒinfo       λ.Object
						ϒparams     λ.Object
						ϒplayer_id  λ.Object
						ϒself       = λargs[0]
						ϒtarget     λ.Object
						ϒurl        = λargs[1]
						ϒurl_type   λ.Object
						ϒvideo      λ.Object
						ϒvideo_id   λ.Object
						ϒwebpage    λ.Object
						τmp0        λ.Object
					)
					τmp0 = λ.UnpackIterable(λ.Calm(λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl), "groups"), 5)
					ϒvideo_id = λ.GetItem(τmp0, λ.IntLiteral(0))
					ϒplayer_id = λ.GetItem(τmp0, λ.IntLiteral(1))
					ϒtarget = λ.GetItem(τmp0, λ.IntLiteral(2))
					ϒurl_type = λ.GetItem(τmp0, λ.IntLiteral(3))
					ϒdisplay_id = λ.GetItem(τmp0, λ.IntLiteral(4))
					if λ.IsTrue(ϒvideo_id) {
						return λ.Calm(ϒself, "_extract_video", λ.DictLiteral(map[string]λ.Object{
							"videoId":  ϒvideo_id,
							"playerId": ϒplayer_id,
							"target":   ϒtarget,
						}))
					}
					ϒwebpage = λ.Calm(ϒself, "_download_webpage", ϒurl, ϒdisplay_id)
					if λ.IsTrue(λ.Eq(ϒurl_type, λ.StrLiteral("series"))) {
						return λ.Calm(ϒself, "_extract_series", ϒurl, ϒwebpage)
					} else {
						ϒvideo = λ.Cal(ϒtry_get, λ.Calm(ϒself, "_parse_json", λ.Calm(ϒself, "_search_regex", λ.StrLiteral("__PRELOADED_STATE__\\s*=\\s*({.+?});"), ϒwebpage, λ.StrLiteral("preload state"), λ.StrLiteral("{}")), ϒdisplay_id), λ.NewFunction("<lambda>",
							[]λ.Param{
								{Name: "x"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒx = λargs[0]
								)
								return λ.GetItem(λ.GetItem(ϒx, λ.StrLiteral("transformed")), λ.StrLiteral("video"))
							}))
						if λ.IsTrue(ϒvideo) {
							ϒparams = λ.DictLiteral(map[string]λ.Object{
								"videoId": λ.GetItem(ϒvideo, λ.StrLiteral("id")),
							})
							ϒinfo = λ.DictLiteral(map[string]λ.Object{
								"description": λ.Cal(ϒstrip_or_none, λ.Calm(ϒvideo, "get", λ.StrLiteral("description"))),
							})
						} else {
							ϒparams = λ.Calm(ϒself, "_extract_video_params", ϒwebpage, ϒdisplay_id)
							ϒinfo = λ.Call(λ.GetAttr(ϒself, "_search_json_ld", nil), λ.NewArgs(
								ϒwebpage,
								ϒdisplay_id,
							), λ.KWArgs{
								{Name: "fatal", Value: λ.False},
							})
						}
						λ.Calm(ϒinfo, "update", λ.Calm(ϒself, "_extract_video", ϒparams))
						return ϒinfo
					}
					return λ.None
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_SITES":         CondeNastIE__SITES,
				"_VALID_URL":     CondeNastIE__VALID_URL,
				"_extract_video": CondeNastIE__extract_video,
				"_real_extract":  CondeNastIE__real_extract,
			})
		}())
	})
}
