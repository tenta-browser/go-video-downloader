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
 * skyit/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/skyit.py
 */

package skyit

import (
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	CieloTVItIE                   λ.Object
	InfoExtractor                 λ.Object
	SkyItAcademyIE                λ.Object
	SkyItArteIE                   λ.Object
	SkyItIE                       λ.Object
	SkyItPlayerIE                 λ.Object
	SkyItVideoIE                  λ.Object
	SkyItVideoLiveIE              λ.Object
	TV8ItIE                       λ.Object
	ϒcompat_parse_qs              λ.Object
	ϒcompat_str                   λ.Object
	ϒcompat_urllib_parse_urlparse λ.Object
	ϒdict_get                     λ.Object
	ϒint_or_none                  λ.Object
	ϒparse_duration               λ.Object
	ϒunified_timestamp            λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_str = Ωcompat.ϒcompat_str
		ϒcompat_parse_qs = Ωcompat.ϒcompat_parse_qs
		ϒcompat_urllib_parse_urlparse = Ωcompat.ϒcompat_urllib_parse_urlparse
		ϒdict_get = Ωutils.ϒdict_get
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒparse_duration = Ωutils.ϒparse_duration
		ϒunified_timestamp = Ωutils.ϒunified_timestamp
		SkyItPlayerIE = λ.Cal(λ.TypeType, λ.StrLiteral("SkyItPlayerIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				SkyItPlayerIE_IE_NAME            λ.Object
				SkyItPlayerIE__DOMAIN            λ.Object
				SkyItPlayerIE__GEO_BYPASS        λ.Object
				SkyItPlayerIE__PLAYER_TMPL       λ.Object
				SkyItPlayerIE__TOKEN_MAP         λ.Object
				SkyItPlayerIE__VALID_URL         λ.Object
				SkyItPlayerIE__parse_video       λ.Object
				SkyItPlayerIE__player_url_result λ.Object
				SkyItPlayerIE__real_extract      λ.Object
			)
			SkyItPlayerIE_IE_NAME = λ.StrLiteral("player.sky.it")
			SkyItPlayerIE__VALID_URL = λ.StrLiteral("https?://player\\.sky\\.it/player/(?:external|social)\\.html\\?.*?\\bid=(?P<id>\\d+)")
			SkyItPlayerIE__GEO_BYPASS = λ.False
			SkyItPlayerIE__DOMAIN = λ.StrLiteral("sky")
			SkyItPlayerIE__PLAYER_TMPL = λ.StrLiteral("https://player.sky.it/player/external.html?id=%s&domain=%s")
			SkyItPlayerIE__TOKEN_MAP = λ.DictLiteral(map[string]string{
				"cielo":           "Hh9O7M8ks5yi6nSROL7bKYz933rdf3GhwZlTLMgvy4Q",
				"hotclub":         "kW020K2jq2lk2eKRJD2vWEg832ncx2EivZlTLQput2C",
				"mtv8":            "A5Nn9GGb326CI7vP5e27d7E4PIaQjota",
				"salesforce":      "C6D585FD1615272C98DE38235F38BD86",
				"sitocommerciale": "VJwfFuSGnLKnd9Phe9y96WkXgYDCguPMJ2dLhGMb2RE",
				"sky":             "F96WlOd8yoFmLQgiqv6fNQRvHZcsWk5jDaYnDvhbiJk",
				"skyacademy":      "A6LAn7EkO2Q26FRy0IAMBekX6jzDXYL3",
				"skyarte":         "LWk29hfiU39NNdq87ePeRach3nzTSV20o0lTv2001Cd",
				"theupfront":      "PRSGmDMsg6QMGc04Obpoy7Vsbn7i2Whp",
			})
			SkyItPlayerIE__player_url_result = λ.NewFunction("_player_url_result",
				[]λ.Param{
					{Name: "self"},
					{Name: "video_id"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒself     = λargs[0]
						ϒvideo_id = λargs[1]
					)
					return λ.Calm(ϒself, "url_result", λ.Mod(λ.GetAttr(ϒself, "_PLAYER_TMPL", nil), λ.NewTuple(
						ϒvideo_id,
						λ.GetAttr(ϒself, "_DOMAIN", nil),
					)), λ.Calm(SkyItPlayerIE, "ie_key"), ϒvideo_id)
				})
			SkyItPlayerIE__parse_video = λ.NewFunction("_parse_video",
				[]λ.Param{
					{Name: "self"},
					{Name: "video"},
					{Name: "video_id"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒformats  λ.Object
						ϒhls_url  λ.Object
						ϒis_live  λ.Object
						ϒself     = λargs[0]
						ϒtitle    λ.Object
						ϒvideo    = λargs[1]
						ϒvideo_id = λargs[2]
					)
					ϒtitle = λ.GetItem(ϒvideo, λ.StrLiteral("title"))
					ϒis_live = λ.Eq(λ.Calm(ϒvideo, "get", λ.StrLiteral("type")), λ.StrLiteral("live"))
					ϒhls_url = λ.Calm(ϒvideo, "get", λ.Add(func() λ.Object {
						if λ.IsTrue(ϒis_live) {
							return λ.StrLiteral("streaming")
						} else {
							return λ.StrLiteral("hls")
						}
					}(), λ.StrLiteral("_url")))
					if λ.IsTrue(func() λ.Object {
						if λv := λ.NewBool(!λ.IsTrue(ϒhls_url)); !λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Calm(ϒvideo, "get", func() λ.Object {
								if λ.IsTrue(ϒis_live) {
									return λ.StrLiteral("geoblock")
								} else {
									return λ.StrLiteral("geob")
								}
							}())
						}
					}()) {
						λ.Call(λ.GetAttr(ϒself, "raise_geo_restricted", nil), nil, λ.KWArgs{
							{Name: "countries", Value: λ.NewList(λ.StrLiteral("IT"))},
						})
					}
					if λ.IsTrue(ϒis_live) {
						ϒformats = λ.Calm(ϒself, "_extract_m3u8_formats", ϒhls_url, ϒvideo_id, λ.StrLiteral("mp4"))
					} else {
						ϒformats = λ.Calm(ϒself, "_extract_akamai_formats", ϒhls_url, ϒvideo_id, λ.DictLiteral(map[string]string{
							"http": "videoplatform.sky.it",
						}))
					}
					λ.Calm(ϒself, "_sort_formats", ϒformats)
					return λ.DictLiteral(map[string]λ.Object{
						"id": ϒvideo_id,
						"title": func() λ.Object {
							if λ.IsTrue(ϒis_live) {
								return λ.Calm(ϒself, "_live_title", ϒtitle)
							} else {
								return ϒtitle
							}
						}(),
						"formats": ϒformats,
						"thumbnail": λ.Cal(ϒdict_get, ϒvideo, λ.NewTuple(
							λ.StrLiteral("video_still"),
							λ.StrLiteral("video_still_medium"),
							λ.StrLiteral("thumb"),
						)),
						"description": func() λ.Object {
							if λv := λ.Calm(ϒvideo, "get", λ.StrLiteral("short_desc")); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.None
							}
						}(),
						"timestamp": λ.Cal(ϒunified_timestamp, λ.Calm(ϒvideo, "get", λ.StrLiteral("create_date"))),
						"duration": func() λ.Object {
							if λv := λ.Cal(ϒint_or_none, λ.Calm(ϒvideo, "get", λ.StrLiteral("duration_sec"))); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Cal(ϒparse_duration, λ.Calm(ϒvideo, "get", λ.StrLiteral("duration")))
							}
						}(),
						"is_live": ϒis_live,
					})
				})
			SkyItPlayerIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdomain   λ.Object
						ϒself     = λargs[0]
						ϒtoken    λ.Object
						ϒurl      = λargs[1]
						ϒvideo    λ.Object
						ϒvideo_id λ.Object
					)
					ϒvideo_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒdomain = λ.GetItem(λ.Calm(λ.Cal(ϒcompat_parse_qs, λ.GetAttr(λ.Cal(ϒcompat_urllib_parse_urlparse, ϒurl), "query", nil)), "get", λ.StrLiteral("domain"), λ.NewList(λ.None)), λ.IntLiteral(0))
					ϒtoken = λ.Cal(ϒdict_get, λ.GetAttr(ϒself, "_TOKEN_MAP", nil), λ.NewTuple(
						ϒdomain,
						λ.StrLiteral("sky"),
					))
					ϒvideo = λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
						λ.StrLiteral("https://apid.sky.it/vdp/v1/getVideoData"),
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "query", Value: λ.DictLiteral(map[string]λ.Object{
							"caller": λ.StrLiteral("sky"),
							"id":     ϒvideo_id,
							"token":  ϒtoken,
						})},
						{Name: "headers", Value: λ.Calm(ϒself, "geo_verification_headers")},
					})
					return λ.Calm(ϒself, "_parse_video", ϒvideo, ϒvideo_id)
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"IE_NAME":            SkyItPlayerIE_IE_NAME,
				"_DOMAIN":            SkyItPlayerIE__DOMAIN,
				"_GEO_BYPASS":        SkyItPlayerIE__GEO_BYPASS,
				"_PLAYER_TMPL":       SkyItPlayerIE__PLAYER_TMPL,
				"_TOKEN_MAP":         SkyItPlayerIE__TOKEN_MAP,
				"_VALID_URL":         SkyItPlayerIE__VALID_URL,
				"_parse_video":       SkyItPlayerIE__parse_video,
				"_player_url_result": SkyItPlayerIE__player_url_result,
				"_real_extract":      SkyItPlayerIE__real_extract,
			})
		}())
		SkyItVideoIE = λ.Cal(λ.TypeType, λ.StrLiteral("SkyItVideoIE"), λ.NewTuple(SkyItPlayerIE), func() λ.Dict {
			var (
				SkyItVideoIE_IE_NAME       λ.Object
				SkyItVideoIE__VALID_URL    λ.Object
				SkyItVideoIE__real_extract λ.Object
			)
			SkyItVideoIE_IE_NAME = λ.StrLiteral("video.sky.it")
			SkyItVideoIE__VALID_URL = λ.StrLiteral("https?://(?:masterchef|video|xfactor)\\.sky\\.it(?:/[^/]+)*/video/[0-9a-z-]+-(?P<id>\\d+)")
			SkyItVideoIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒself     = λargs[0]
						ϒurl      = λargs[1]
						ϒvideo_id λ.Object
					)
					ϒvideo_id = λ.Calm(ϒself, "_match_id", ϒurl)
					return λ.Calm(ϒself, "_player_url_result", ϒvideo_id)
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"IE_NAME":       SkyItVideoIE_IE_NAME,
				"_VALID_URL":    SkyItVideoIE__VALID_URL,
				"_real_extract": SkyItVideoIE__real_extract,
			})
		}())
		SkyItVideoLiveIE = λ.Cal(λ.TypeType, λ.StrLiteral("SkyItVideoLiveIE"), λ.NewTuple(SkyItPlayerIE), func() λ.Dict {
			var (
				SkyItVideoLiveIE__VALID_URL λ.Object
			)
			SkyItVideoLiveIE__VALID_URL = λ.StrLiteral("https?://video\\.sky\\.it/diretta/(?P<id>[^/?&#]+)")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": SkyItVideoLiveIE__VALID_URL,
			})
		}())
		SkyItIE = λ.Cal(λ.TypeType, λ.StrLiteral("SkyItIE"), λ.NewTuple(SkyItPlayerIE), func() λ.Dict {
			var (
				SkyItIE_IE_NAME         λ.Object
				SkyItIE__VALID_URL      λ.Object
				SkyItIE__VIDEO_ID_REGEX λ.Object
				SkyItIE__real_extract   λ.Object
			)
			SkyItIE_IE_NAME = λ.StrLiteral("sky.it")
			SkyItIE__VALID_URL = λ.StrLiteral("https?://(?:sport|tg24)\\.sky\\.it(?:/[^/]+)*/\\d{4}/\\d{2}/\\d{2}/(?P<id>[^/?&#]+)")
			SkyItIE__VIDEO_ID_REGEX = λ.StrLiteral("data-videoid=\"(\\d+)\"")
			SkyItIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdisplay_id λ.Object
						ϒself       = λargs[0]
						ϒurl        = λargs[1]
						ϒvideo_id   λ.Object
						ϒwebpage    λ.Object
					)
					ϒdisplay_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒwebpage = λ.Calm(ϒself, "_download_webpage", ϒurl, ϒdisplay_id)
					ϒvideo_id = λ.Calm(ϒself, "_search_regex", λ.GetAttr(ϒself, "_VIDEO_ID_REGEX", nil), ϒwebpage, λ.StrLiteral("video id"))
					return λ.Calm(ϒself, "_player_url_result", ϒvideo_id)
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"IE_NAME":         SkyItIE_IE_NAME,
				"_VALID_URL":      SkyItIE__VALID_URL,
				"_VIDEO_ID_REGEX": SkyItIE__VIDEO_ID_REGEX,
				"_real_extract":   SkyItIE__real_extract,
			})
		}())
		SkyItAcademyIE = λ.Cal(λ.TypeType, λ.StrLiteral("SkyItAcademyIE"), λ.NewTuple(SkyItIE), func() λ.Dict {
			var (
				SkyItAcademyIE__VALID_URL λ.Object
			)
			SkyItAcademyIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?skyacademy\\.it(?:/[^/]+)*/\\d{4}/\\d{2}/\\d{2}/(?P<id>[^/?&#]+)")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": SkyItAcademyIE__VALID_URL,
			})
		}())
		SkyItArteIE = λ.Cal(λ.TypeType, λ.StrLiteral("SkyItArteIE"), λ.NewTuple(SkyItIE), func() λ.Dict {
			var (
				SkyItArteIE_IE_NAME         λ.Object
				SkyItArteIE__DOMAIN         λ.Object
				SkyItArteIE__VALID_URL      λ.Object
				SkyItArteIE__VIDEO_ID_REGEX λ.Object
			)
			SkyItArteIE_IE_NAME = λ.StrLiteral("arte.sky.it")
			SkyItArteIE__VALID_URL = λ.StrLiteral("https?://arte\\.sky\\.it/video/(?P<id>[^/?&#]+)")
			SkyItArteIE__DOMAIN = λ.StrLiteral("skyarte")
			SkyItArteIE__VIDEO_ID_REGEX = λ.StrLiteral("(?s)<iframe[^>]+src=\"(?:https:)?//player\\.sky\\.it/player/external\\.html\\?[^\"]*\\bid=(\\d+)")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"IE_NAME":         SkyItArteIE_IE_NAME,
				"_DOMAIN":         SkyItArteIE__DOMAIN,
				"_VALID_URL":      SkyItArteIE__VALID_URL,
				"_VIDEO_ID_REGEX": SkyItArteIE__VIDEO_ID_REGEX,
			})
		}())
		CieloTVItIE = λ.Cal(λ.TypeType, λ.StrLiteral("CieloTVItIE"), λ.NewTuple(SkyItIE), func() λ.Dict {
			var (
				CieloTVItIE_IE_NAME         λ.Object
				CieloTVItIE__DOMAIN         λ.Object
				CieloTVItIE__VALID_URL      λ.Object
				CieloTVItIE__VIDEO_ID_REGEX λ.Object
			)
			CieloTVItIE_IE_NAME = λ.StrLiteral("cielotv.it")
			CieloTVItIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?cielotv\\.it/video/(?P<id>[^.]+)\\.html")
			CieloTVItIE__DOMAIN = λ.StrLiteral("cielo")
			CieloTVItIE__VIDEO_ID_REGEX = λ.StrLiteral("videoId\\s*=\\s*\"(\\d+)\"")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"IE_NAME":         CieloTVItIE_IE_NAME,
				"_DOMAIN":         CieloTVItIE__DOMAIN,
				"_VALID_URL":      CieloTVItIE__VALID_URL,
				"_VIDEO_ID_REGEX": CieloTVItIE__VIDEO_ID_REGEX,
			})
		}())
		TV8ItIE = λ.Cal(λ.TypeType, λ.StrLiteral("TV8ItIE"), λ.NewTuple(SkyItVideoIE), func() λ.Dict {
			var (
				TV8ItIE_IE_NAME    λ.Object
				TV8ItIE__DOMAIN    λ.Object
				TV8ItIE__VALID_URL λ.Object
			)
			TV8ItIE_IE_NAME = λ.StrLiteral("tv8.it")
			TV8ItIE__VALID_URL = λ.StrLiteral("https?://tv8\\.it/showvideo/(?P<id>\\d+)")
			TV8ItIE__DOMAIN = λ.StrLiteral("mtv8")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"IE_NAME":    TV8ItIE_IE_NAME,
				"_DOMAIN":    TV8ItIE__DOMAIN,
				"_VALID_URL": TV8ItIE__VALID_URL,
			})
		}())
	})
}
