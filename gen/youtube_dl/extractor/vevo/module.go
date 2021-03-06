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
 * vevo/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/vevo.py
 */

package vevo

import (
	Ωjson "github.com/tenta-browser/go-video-downloader/gen/json"
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError λ.Object
	InfoExtractor  λ.Object
	VevoBaseIE     λ.Object
	VevoIE         λ.Object
	VevoPlaylistIE λ.Object
	ϒcompat_str    λ.Object
	ϒint_or_none   λ.Object
	ϒparse_iso8601 λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_str = Ωcompat.ϒcompat_str
		ExtractorError = Ωutils.ExtractorError
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒparse_iso8601 = Ωutils.ϒparse_iso8601
		VevoBaseIE = λ.Cal(λ.TypeType, λ.StrLiteral("VevoBaseIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				VevoBaseIE__extract_json λ.Object
			)
			VevoBaseIE__extract_json = λ.NewFunction("_extract_json",
				[]λ.Param{
					{Name: "self"},
					{Name: "webpage"},
					{Name: "video_id"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒself     = λargs[0]
						ϒvideo_id = λargs[2]
						ϒwebpage  = λargs[1]
					)
					return λ.Calm(ϒself, "_parse_json", λ.Calm(ϒself, "_search_regex", λ.StrLiteral("window\\.__INITIAL_STORE__\\s*=\\s*({.+?});\\s*</script>"), ϒwebpage, λ.StrLiteral("initial store")), ϒvideo_id)
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_extract_json": VevoBaseIE__extract_json,
			})
		}())
		VevoIE = λ.Cal(λ.TypeType, λ.StrLiteral("VevoIE"), λ.NewTuple(VevoBaseIE), func() λ.Dict {
			var (
				VevoIE__VALID_URL      λ.Object
				VevoIE__VERSIONS       λ.Object
				VevoIE__call_api       λ.Object
				VevoIE__initialize_api λ.Object
				VevoIE__real_extract   λ.Object
			)
			VevoIE__VALID_URL = λ.StrLiteral("(?x)\n        (?:https?://(?:www\\.)?vevo\\.com/watch/(?!playlist|genre)(?:[^/]+/(?:[^/]+/)?)?|\n           https?://cache\\.vevo\\.com/m/html/embed\\.html\\?video=|\n           https?://videoplayer\\.vevo\\.com/embed/embedded\\?videoId=|\n           https?://embed\\.vevo\\.com/.*?[?&]isrc=|\n           vevo:)\n        (?P<id>[^&?#]+)")
			VevoIE__VERSIONS = λ.DictLiteral(map[int]string{
				0: "youtube",
				1: "level3",
				2: "akamai",
				3: "level3",
				4: "amazon",
			})
			VevoIE__initialize_api = λ.NewFunction("_initialize_api",
				[]λ.Param{
					{Name: "self"},
					{Name: "video_id"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒauth_info λ.Object
						ϒself      = λargs[0]
						ϒvideo_id  = λargs[1]
						ϒwebpage   λ.Object
					)
					ϒwebpage = λ.Call(λ.GetAttr(ϒself, "_download_webpage", nil), λ.NewArgs(
						λ.StrLiteral("https://accounts.vevo.com/token"),
						λ.None,
					), λ.KWArgs{
						{Name: "note", Value: λ.StrLiteral("Retrieving oauth token")},
						{Name: "errnote", Value: λ.StrLiteral("Unable to retrieve oauth token")},
						{Name: "data", Value: λ.Calm(λ.Cal(Ωjson.ϒdumps, λ.DictLiteral(map[string]string{
							"client_id":  "SPupX1tvqFEopQ1YS6SS",
							"grant_type": "urn:vevo:params:oauth:grant-type:anonymous",
						})), "encode", λ.StrLiteral("utf-8"))},
						{Name: "headers", Value: λ.DictLiteral(map[string]string{
							"Content-Type": "application/json",
						})},
					})
					if λ.IsTrue(λ.Cal(Ωre.ϒsearch, λ.StrLiteral("(?i)THIS PAGE IS CURRENTLY UNAVAILABLE IN YOUR REGION"), ϒwebpage)) {
						λ.Calm(ϒself, "raise_geo_restricted", λ.Mod(λ.StrLiteral("%s said: This page is currently unavailable in your region"), λ.GetAttr(ϒself, "IE_NAME", nil)))
					}
					ϒauth_info = λ.Calm(ϒself, "_parse_json", ϒwebpage, ϒvideo_id)
					λ.SetAttr(ϒself, "_api_url_template", λ.Add(λ.Add(λ.Calm(ϒself, "http_scheme"), λ.StrLiteral("//apiv2.vevo.com/%s?token=")), λ.GetItem(ϒauth_info, λ.StrLiteral("legacy_token"))))
					return λ.None
				})
			VevoIE__call_api = λ.NewFunction("_call_api",
				[]λ.Param{
					{Name: "self"},
					{Name: "path"},
				},
				0, true, true,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒargs          = λargs[2]
						ϒdata          λ.Object
						ϒerror_message λ.Object
						ϒerrors        λ.Object
						ϒkwargs        = λargs[3]
						ϒpath          = λargs[1]
						ϒself          = λargs[0]
						τmp0           λ.Object
						τmp1           λ.Object
					)
					_ = τmp0
					_ = τmp1
					τmp0, τmp1 = func() (λexit λ.Object, λret λ.Object) {
						defer λ.CatchMulti(
							nil,
							&λ.Catcher{ExtractorError, func(λex λ.BaseException) {
								var ϒe λ.Object = λex
								if λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, λ.GetAttr(ϒe, "cause", nil), λ.None)) {
									ϒerrors = λ.GetItem(λ.Calm(ϒself, "_parse_json", λ.Calm(λ.Calm(λ.GetAttr(ϒe, "cause", nil), "read"), "decode"), λ.None), λ.StrLiteral("errors"))
									ϒerror_message = λ.Calm(λ.StrLiteral(", "), "join", λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
										nil,
										0, false, false,
										func(λargs []λ.Object) λ.Object {
											return λ.NewGenerator(func(λgy λ.Yielder) λ.Object {
												var (
													ϒerror λ.Object
													τmp0   λ.Object
													τmp1   λ.Object
												)
												τmp0 = λ.Cal(λ.BuiltinIter, ϒerrors)
												for {
													if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
														break
													}
													ϒerror = τmp1
													λgy.Yield(λ.GetItem(ϒerror, λ.StrLiteral("message")))
												}
												return λ.None
											})
										}))))
									panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.Mod(λ.StrLiteral("%s said: %s"), λ.NewTuple(
										λ.GetAttr(ϒself, "IE_NAME", nil),
										ϒerror_message,
									))), λ.KWArgs{
										{Name: "expected", Value: λ.True},
									})))
								}
								panic(λ.Raise(λex))
							}},
						)
						ϒdata = λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(λ.Unpack(
							λ.Mod(λ.GetAttr(ϒself, "_api_url_template", nil), ϒpath),
							λ.AsStarred(ϒargs),
						)...), λ.KWArgs{
							{Name: "", Value: ϒkwargs},
						})
						return λ.BlockExitNormally, nil
					}()
					return ϒdata
				})
			VevoIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒage_limit       λ.Object
						ϒartist          λ.Object
						ϒartists         λ.Object
						ϒcurr_artist     λ.Object
						ϒfeatured_artist λ.Object
						ϒformats         λ.Object
						ϒgenre           λ.Object
						ϒgenres          λ.Object
						ϒis_explicit     λ.Object
						ϒjson_data       λ.Object
						ϒm               λ.Object
						ϒself            = λargs[0]
						ϒtitle           λ.Object
						ϒtrack           λ.Object
						ϒuploader        λ.Object
						ϒurl             = λargs[1]
						ϒversion         λ.Object
						ϒversion_url     λ.Object
						ϒvideo_id        λ.Object
						ϒvideo_info      λ.Object
						ϒvideo_version   λ.Object
						ϒvideo_versions  λ.Object
						ϒwebpage         λ.Object
						τmp0             λ.Object
						τmp1             λ.Object
						τmp2             λ.Object
					)
					ϒvideo_id = λ.Calm(ϒself, "_match_id", ϒurl)
					λ.Calm(ϒself, "_initialize_api", ϒvideo_id)
					ϒvideo_info = λ.Calm(ϒself, "_call_api", λ.Mod(λ.StrLiteral("video/%s"), ϒvideo_id), ϒvideo_id, λ.StrLiteral("Downloading api video info"), λ.StrLiteral("Failed to download video info"))
					ϒvideo_versions = λ.Call(λ.GetAttr(ϒself, "_call_api", nil), λ.NewArgs(
						λ.Mod(λ.StrLiteral("video/%s/streams"), ϒvideo_id),
						ϒvideo_id,
						λ.StrLiteral("Downloading video versions info"),
						λ.StrLiteral("Failed to download video versions info"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					})
					if !λ.IsTrue(ϒvideo_versions) {
						ϒwebpage = λ.Calm(ϒself, "_download_webpage", ϒurl, ϒvideo_id)
						ϒjson_data = λ.Calm(ϒself, "_extract_json", ϒwebpage, ϒvideo_id)
						if λ.Contains(λ.Calm(ϒjson_data, "get", λ.StrLiteral("default"), λ.DictLiteral(map[λ.Object]λ.Object{})), λ.StrLiteral("streams")) {
							ϒvideo_versions = λ.GetItem(λ.GetItem(λ.GetItem(λ.GetItem(ϒjson_data, λ.StrLiteral("default")), λ.StrLiteral("streams")), ϒvideo_id), λ.IntLiteral(0))
						} else {
							ϒvideo_versions = λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
								nil,
								0, false, false,
								func(λargs []λ.Object) λ.Object {
									return λ.NewGenerator(func(λgy λ.Yielder) λ.Object {
										var (
											ϒkey   λ.Object
											ϒvalue λ.Object
											τmp0   λ.Object
											τmp1   λ.Object
											τmp2   λ.Object
										)
										τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(λ.GetItem(λ.GetItem(ϒjson_data, λ.StrLiteral("apollo")), λ.StrLiteral("data")), "items"))
										for {
											if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
												break
											}
											τmp2 = λ.UnpackIterable(τmp1, 2)
											ϒkey = λ.GetItem(τmp2, λ.IntLiteral(0))
											ϒvalue = λ.GetItem(τmp2, λ.IntLiteral(1))
											if λ.IsTrue(λ.Calm(ϒkey, "startswith", λ.Mod(λ.StrLiteral("%s.streams"), ϒvideo_id))) {
												λgy.Yield(ϒvalue)
											}
										}
										return λ.None
									})
								})))
						}
					}
					ϒuploader = λ.None
					ϒartist = λ.None
					ϒfeatured_artist = λ.None
					ϒartists = λ.Calm(ϒvideo_info, "get", λ.StrLiteral("artists"))
					τmp0 = λ.Cal(λ.BuiltinIter, ϒartists)
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒcurr_artist = τmp1
						if λ.IsTrue(λ.Eq(λ.Calm(ϒcurr_artist, "get", λ.StrLiteral("role")), λ.StrLiteral("Featured"))) {
							ϒfeatured_artist = λ.GetItem(ϒcurr_artist, λ.StrLiteral("name"))
						} else {
							τmp2 = λ.GetItem(ϒcurr_artist, λ.StrLiteral("name"))
							ϒartist = τmp2
							ϒuploader = τmp2
						}
					}
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, ϒvideo_versions)
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒvideo_version = τmp1
						ϒversion = λ.Calm(λ.GetAttr(ϒself, "_VERSIONS", nil), "get", λ.Calm(ϒvideo_version, "get", λ.StrLiteral("version")), λ.StrLiteral("generic"))
						ϒversion_url = λ.Calm(ϒvideo_version, "get", λ.StrLiteral("url"))
						if !λ.IsTrue(ϒversion_url) {
							continue
						}
						if λ.Contains(ϒversion_url, λ.StrLiteral(".ism")) {
							continue
						} else {
							if λ.Contains(ϒversion_url, λ.StrLiteral(".mpd")) {
								λ.Calm(ϒformats, "extend", λ.Call(λ.GetAttr(ϒself, "_extract_mpd_formats", nil), λ.NewArgs(
									ϒversion_url,
									ϒvideo_id,
								), λ.KWArgs{
									{Name: "mpd_id", Value: λ.Mod(λ.StrLiteral("dash-%s"), ϒversion)},
									{Name: "note", Value: λ.Mod(λ.StrLiteral("Downloading %s MPD information"), ϒversion)},
									{Name: "errnote", Value: λ.Mod(λ.StrLiteral("Failed to download %s MPD information"), ϒversion)},
									{Name: "fatal", Value: λ.False},
								}))
							} else {
								if λ.Contains(ϒversion_url, λ.StrLiteral(".m3u8")) {
									λ.Calm(ϒformats, "extend", λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
										ϒversion_url,
										ϒvideo_id,
										λ.StrLiteral("mp4"),
										λ.StrLiteral("m3u8_native"),
									), λ.KWArgs{
										{Name: "m3u8_id", Value: λ.Mod(λ.StrLiteral("hls-%s"), ϒversion)},
										{Name: "note", Value: λ.Mod(λ.StrLiteral("Downloading %s m3u8 information"), ϒversion)},
										{Name: "errnote", Value: λ.Mod(λ.StrLiteral("Failed to download %s m3u8 information"), ϒversion)},
										{Name: "fatal", Value: λ.False},
									}))
								} else {
									ϒm = λ.Cal(Ωre.ϒsearch, λ.StrLiteral("(?xi)\n                    _(?P<width>[0-9]+)x(?P<height>[0-9]+)\n                    _(?P<vcodec>[a-z0-9]+)\n                    _(?P<vbr>[0-9]+)\n                    _(?P<acodec>[a-z0-9]+)\n                    _(?P<abr>[0-9]+)\n                    \\.(?P<ext>[a-z0-9]+)"), ϒversion_url)
									if !λ.IsTrue(ϒm) {
										continue
									}
									λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
										"url": ϒversion_url,
										"format_id": λ.Mod(λ.StrLiteral("http-%s-%s"), λ.NewTuple(
											ϒversion,
											λ.GetItem(ϒvideo_version, λ.StrLiteral("quality")),
										)),
										"vcodec": λ.Calm(ϒm, "group", λ.StrLiteral("vcodec")),
										"acodec": λ.Calm(ϒm, "group", λ.StrLiteral("acodec")),
										"vbr":    λ.Cal(λ.IntType, λ.Calm(ϒm, "group", λ.StrLiteral("vbr"))),
										"abr":    λ.Cal(λ.IntType, λ.Calm(ϒm, "group", λ.StrLiteral("abr"))),
										"ext":    λ.Calm(ϒm, "group", λ.StrLiteral("ext")),
										"width":  λ.Cal(λ.IntType, λ.Calm(ϒm, "group", λ.StrLiteral("width"))),
										"height": λ.Cal(λ.IntType, λ.Calm(ϒm, "group", λ.StrLiteral("height"))),
									}))
								}
							}
						}
					}
					λ.Calm(ϒself, "_sort_formats", ϒformats)
					ϒtrack = λ.GetItem(ϒvideo_info, λ.StrLiteral("title"))
					if λ.IsTrue(ϒfeatured_artist) {
						ϒartist = λ.Mod(λ.StrLiteral("%s ft. %s"), λ.NewTuple(
							ϒartist,
							ϒfeatured_artist,
						))
					}
					ϒtitle = func() λ.Object {
						if λ.IsTrue(ϒartist) {
							return λ.Mod(λ.StrLiteral("%s - %s"), λ.NewTuple(
								ϒartist,
								ϒtrack,
							))
						} else {
							return ϒtrack
						}
					}()
					ϒgenres = λ.Calm(ϒvideo_info, "get", λ.StrLiteral("genres"))
					ϒgenre = func() λ.Object {
						if λ.IsTrue(func() λ.Object {
							if λv := ϒgenres; !λ.IsTrue(λv) {
								return λv
							} else if λv := λ.Cal(λ.BuiltinIsInstance, ϒgenres, λ.ListType); !λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Cal(λ.BuiltinIsInstance, λ.GetItem(ϒgenres, λ.IntLiteral(0)), ϒcompat_str)
							}
						}()) {
							return λ.GetItem(ϒgenres, λ.IntLiteral(0))
						} else {
							return λ.None
						}
					}()
					ϒis_explicit = λ.Calm(ϒvideo_info, "get", λ.StrLiteral("isExplicit"))
					if ϒis_explicit == λ.True {
						ϒage_limit = λ.IntLiteral(18)
					} else {
						if ϒis_explicit == λ.False {
							ϒage_limit = λ.IntLiteral(0)
						} else {
							ϒage_limit = λ.None
						}
					}
					return λ.DictLiteral(map[string]λ.Object{
						"id":      ϒvideo_id,
						"title":   ϒtitle,
						"formats": ϒformats,
						"thumbnail": func() λ.Object {
							if λv := λ.Calm(ϒvideo_info, "get", λ.StrLiteral("imageUrl")); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Calm(ϒvideo_info, "get", λ.StrLiteral("thumbnailUrl"))
							}
						}(),
						"timestamp":  λ.Cal(ϒparse_iso8601, λ.Calm(ϒvideo_info, "get", λ.StrLiteral("releaseDate"))),
						"uploader":   ϒuploader,
						"duration":   λ.Cal(ϒint_or_none, λ.Calm(ϒvideo_info, "get", λ.StrLiteral("duration"))),
						"view_count": λ.Cal(ϒint_or_none, λ.Calm(λ.Calm(ϒvideo_info, "get", λ.StrLiteral("views"), λ.DictLiteral(map[λ.Object]λ.Object{})), "get", λ.StrLiteral("total"))),
						"age_limit":  ϒage_limit,
						"track":      ϒtrack,
						"artist":     ϒuploader,
						"genre":      ϒgenre,
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL":      VevoIE__VALID_URL,
				"_VERSIONS":       VevoIE__VERSIONS,
				"_call_api":       VevoIE__call_api,
				"_initialize_api": VevoIE__initialize_api,
				"_real_extract":   VevoIE__real_extract,
			})
		}())
		VevoPlaylistIE = λ.Cal(λ.TypeType, λ.StrLiteral("VevoPlaylistIE"), λ.NewTuple(VevoBaseIE), func() λ.Dict {
			var (
				VevoPlaylistIE__VALID_URL λ.Object
			)
			VevoPlaylistIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?vevo\\.com/watch/(?P<kind>playlist|genre)/(?P<id>[^/?#&]+)")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": VevoPlaylistIE__VALID_URL,
			})
		}())
	})
}
