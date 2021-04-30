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
 * zdf/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/zdf.py
 */

package zdf

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor      λ.Object
	NO_DEFAULT         λ.Object
	ZDFBaseIE          λ.Object
	ZDFChannelIE       λ.Object
	ZDFIE              λ.Object
	ϒcompat_str        λ.Object
	ϒdetermine_ext     λ.Object
	ϒfloat_or_none     λ.Object
	ϒint_or_none       λ.Object
	ϒmerge_dicts       λ.Object
	ϒparse_codecs      λ.Object
	ϒqualities         λ.Object
	ϒtry_get           λ.Object
	ϒunified_timestamp λ.Object
	ϒupdate_url_query  λ.Object
	ϒurl_or_none       λ.Object
	ϒurljoin           λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_str = Ωcompat.ϒcompat_str
		ϒdetermine_ext = Ωutils.ϒdetermine_ext
		ϒfloat_or_none = Ωutils.ϒfloat_or_none
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒmerge_dicts = Ωutils.ϒmerge_dicts
		NO_DEFAULT = Ωutils.NO_DEFAULT
		ϒparse_codecs = Ωutils.ϒparse_codecs
		ϒqualities = Ωutils.ϒqualities
		ϒtry_get = Ωutils.ϒtry_get
		ϒunified_timestamp = Ωutils.ϒunified_timestamp
		ϒupdate_url_query = Ωutils.ϒupdate_url_query
		ϒurl_or_none = Ωutils.ϒurl_or_none
		ϒurljoin = Ωutils.ϒurljoin
		ZDFBaseIE = λ.Cal(λ.TypeType, λ.StrLiteral("ZDFBaseIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				ZDFBaseIE__GEO_COUNTRIES     λ.Object
				ZDFBaseIE__QUALITIES         λ.Object
				ZDFBaseIE__call_api          λ.Object
				ZDFBaseIE__extract_format    λ.Object
				ZDFBaseIE__extract_player    λ.Object
				ZDFBaseIE__extract_ptmd      λ.Object
				ZDFBaseIE__extract_subtitles λ.Object
			)
			ZDFBaseIE__GEO_COUNTRIES = λ.NewList(λ.StrLiteral("DE"))
			ZDFBaseIE__QUALITIES = λ.NewTuple(
				λ.StrLiteral("auto"),
				λ.StrLiteral("low"),
				λ.StrLiteral("med"),
				λ.StrLiteral("high"),
				λ.StrLiteral("veryhigh"),
				λ.StrLiteral("hd"),
			)
			ZDFBaseIE__call_api = λ.NewFunction("_call_api",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
					{Name: "video_id"},
					{Name: "item"},
					{Name: "api_token", Def: λ.None},
					{Name: "referrer", Def: λ.None},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒapi_token = λargs[4]
						ϒheaders   λ.Object
						ϒitem      = λargs[3]
						ϒreferrer  = λargs[5]
						ϒself      = λargs[0]
						ϒurl       = λargs[1]
						ϒvideo_id  = λargs[2]
					)
					ϒheaders = λ.DictLiteral(map[λ.Object]λ.Object{})
					if λ.IsTrue(ϒapi_token) {
						λ.SetItem(ϒheaders, λ.StrLiteral("Api-Auth"), λ.Mod(λ.StrLiteral("Bearer %s"), ϒapi_token))
					}
					if λ.IsTrue(ϒreferrer) {
						λ.SetItem(ϒheaders, λ.StrLiteral("Referer"), ϒreferrer)
					}
					return λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
						ϒurl,
						ϒvideo_id,
						λ.Mod(λ.StrLiteral("Downloading JSON %s"), ϒitem),
					), λ.KWArgs{
						{Name: "headers", Value: ϒheaders},
					})
				})
			ZDFBaseIE__extract_subtitles = λ.NewFunction("_extract_subtitles",
				[]λ.Param{
					{Name: "src"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒcaption      λ.Object
						ϒlang         λ.Object
						ϒsrc          = λargs[0]
						ϒsubtitle_url λ.Object
						ϒsubtitles    λ.Object
						τmp0          λ.Object
						τmp1          λ.Object
					)
					ϒsubtitles = λ.DictLiteral(map[λ.Object]λ.Object{})
					τmp0 = λ.Cal(λ.BuiltinIter, func() λ.Object {
						if λv := λ.Cal(ϒtry_get, ϒsrc, λ.NewFunction("<lambda>",
							[]λ.Param{
								{Name: "x"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒx = λargs[0]
								)
								return λ.GetItem(ϒx, λ.StrLiteral("captions"))
							}), λ.ListType); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.NewList()
						}
					}())
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒcaption = τmp1
						ϒsubtitle_url = λ.Cal(ϒurl_or_none, λ.Calm(ϒcaption, "get", λ.StrLiteral("uri")))
						if λ.IsTrue(ϒsubtitle_url) {
							ϒlang = λ.Calm(ϒcaption, "get", λ.StrLiteral("language"), λ.StrLiteral("deu"))
							λ.Calm(λ.Calm(ϒsubtitles, "setdefault", ϒlang, λ.NewList()), "append", λ.DictLiteral(map[string]λ.Object{
								"url": ϒsubtitle_url,
							}))
						}
					}
					return ϒsubtitles
				})
			ZDFBaseIE__extract_subtitles = λ.Cal(λ.StaticMethodType, ZDFBaseIE__extract_subtitles)
			ZDFBaseIE__extract_format = λ.NewFunction("_extract_format",
				[]λ.Param{
					{Name: "self"},
					{Name: "video_id"},
					{Name: "formats"},
					{Name: "format_urls"},
					{Name: "meta"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒext         λ.Object
						ϒf           λ.Object
						ϒformat_id   λ.Object
						ϒformat_url  λ.Object
						ϒformat_urls = λargs[3]
						ϒformats     = λargs[2]
						ϒmeta        = λargs[4]
						ϒmime_type   λ.Object
						ϒp           λ.Object
						ϒself        = λargs[0]
						ϒvideo_id    = λargs[1]
						τmp0         λ.Object
						τmp1         λ.Object
					)
					ϒformat_url = λ.Cal(ϒurl_or_none, λ.Calm(ϒmeta, "get", λ.StrLiteral("url")))
					if !λ.IsTrue(ϒformat_url) {
						return λ.None
					}
					if λ.Contains(ϒformat_urls, ϒformat_url) {
						return λ.None
					}
					λ.Calm(ϒformat_urls, "add", ϒformat_url)
					ϒmime_type = λ.Calm(ϒmeta, "get", λ.StrLiteral("mimeType"))
					ϒext = λ.Cal(ϒdetermine_ext, ϒformat_url)
					if λ.IsTrue(func() λ.Object {
						if λv := λ.Eq(ϒmime_type, λ.StrLiteral("application/x-mpegURL")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Eq(ϒext, λ.StrLiteral("m3u8"))
						}
					}()) {
						λ.Calm(ϒformats, "extend", λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
							ϒformat_url,
							ϒvideo_id,
							λ.StrLiteral("mp4"),
						), λ.KWArgs{
							{Name: "m3u8_id", Value: λ.StrLiteral("hls")},
							{Name: "entry_protocol", Value: λ.StrLiteral("m3u8_native")},
							{Name: "fatal", Value: λ.False},
						}))
					} else {
						if λ.IsTrue(func() λ.Object {
							if λv := λ.Eq(ϒmime_type, λ.StrLiteral("application/f4m+xml")); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Eq(ϒext, λ.StrLiteral("f4m"))
							}
						}()) {
							λ.Calm(ϒformats, "extend", λ.Call(λ.GetAttr(ϒself, "_extract_f4m_formats", nil), λ.NewArgs(
								λ.Cal(ϒupdate_url_query, ϒformat_url, λ.DictLiteral(map[string]string{
									"hdcore": "3.7.0",
								})),
								ϒvideo_id,
							), λ.KWArgs{
								{Name: "f4m_id", Value: λ.StrLiteral("hds")},
								{Name: "fatal", Value: λ.False},
							}))
						} else {
							ϒf = λ.Cal(ϒparse_codecs, λ.Calm(ϒmeta, "get", λ.StrLiteral("mimeCodec")))
							ϒformat_id = λ.NewList(λ.StrLiteral("http"))
							τmp0 = λ.Cal(λ.BuiltinIter, λ.NewTuple(
								λ.Calm(ϒmeta, "get", λ.StrLiteral("type")),
								λ.Calm(ϒmeta, "get", λ.StrLiteral("quality")),
							))
							for {
								if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
									break
								}
								ϒp = τmp1
								if λ.IsTrue(func() λ.Object {
									if λv := ϒp; !λ.IsTrue(λv) {
										return λv
									} else {
										return λ.Cal(λ.BuiltinIsInstance, ϒp, ϒcompat_str)
									}
								}()) {
									λ.Calm(ϒformat_id, "append", ϒp)
								}
							}
							λ.Calm(ϒf, "update", λ.DictLiteral(map[string]λ.Object{
								"url":         ϒformat_url,
								"format_id":   λ.Calm(λ.StrLiteral("-"), "join", ϒformat_id),
								"format_note": λ.Calm(ϒmeta, "get", λ.StrLiteral("quality")),
								"language":    λ.Calm(ϒmeta, "get", λ.StrLiteral("language")),
								"quality":     λ.Cal(λ.Cal(ϒqualities, λ.GetAttr(ϒself, "_QUALITIES", nil)), λ.Calm(ϒmeta, "get", λ.StrLiteral("quality"))),
								"preference":  λ.Neg(λ.IntLiteral(10)),
							}))
							λ.Calm(ϒformats, "append", ϒf)
						}
					}
					return λ.None
				})
			ZDFBaseIE__extract_ptmd = λ.NewFunction("_extract_ptmd",
				[]λ.Param{
					{Name: "self"},
					{Name: "ptmd_url"},
					{Name: "video_id"},
					{Name: "api_token"},
					{Name: "referrer"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒapi_token   = λargs[3]
						ϒcontent_id  λ.Object
						ϒduration    λ.Object
						ϒf           λ.Object
						ϒf_qualities λ.Object
						ϒformats     λ.Object
						ϒformitaeten λ.Object
						ϒp           λ.Object
						ϒptmd        λ.Object
						ϒptmd_url    = λargs[1]
						ϒquality     λ.Object
						ϒreferrer    = λargs[4]
						ϒself        = λargs[0]
						ϒtrack       λ.Object
						ϒtrack_uris  λ.Object
						ϒtracks      λ.Object
						ϒvideo_id    = λargs[2]
						τmp0         λ.Object
						τmp1         λ.Object
						τmp2         λ.Object
						τmp3         λ.Object
						τmp4         λ.Object
						τmp5         λ.Object
						τmp6         λ.Object
						τmp7         λ.Object
					)
					ϒptmd = λ.Calm(ϒself, "_call_api", ϒptmd_url, ϒvideo_id, λ.StrLiteral("metadata"), ϒapi_token, ϒreferrer)
					ϒcontent_id = func() λ.Object {
						if λv := λ.Calm(ϒptmd, "get", λ.StrLiteral("basename")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.GetItem(λ.Calm(ϒptmd_url, "split", λ.StrLiteral("/")), λ.Neg(λ.IntLiteral(1)))
						}
					}()
					ϒformats = λ.NewList()
					ϒtrack_uris = λ.Cal(λ.SetType)
					τmp0 = λ.Cal(λ.BuiltinIter, λ.GetItem(ϒptmd, λ.StrLiteral("priorityList")))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒp = τmp1
						ϒformitaeten = λ.Calm(ϒp, "get", λ.StrLiteral("formitaeten"))
						if !λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒformitaeten, λ.ListType)) {
							continue
						}
						τmp2 = λ.Cal(λ.BuiltinIter, ϒformitaeten)
						for {
							if τmp3 = λ.NextDefault(τmp2, λ.AfterLast); τmp3 == λ.AfterLast {
								break
							}
							ϒf = τmp3
							ϒf_qualities = λ.Calm(ϒf, "get", λ.StrLiteral("qualities"))
							if !λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒf_qualities, λ.ListType)) {
								continue
							}
							τmp4 = λ.Cal(λ.BuiltinIter, ϒf_qualities)
							for {
								if τmp5 = λ.NextDefault(τmp4, λ.AfterLast); τmp5 == λ.AfterLast {
									break
								}
								ϒquality = τmp5
								ϒtracks = λ.Cal(ϒtry_get, ϒquality, λ.NewFunction("<lambda>",
									[]λ.Param{
										{Name: "x"},
									},
									0, false, false,
									func(λargs []λ.Object) λ.Object {
										var (
											ϒx = λargs[0]
										)
										return λ.GetItem(λ.GetItem(ϒx, λ.StrLiteral("audio")), λ.StrLiteral("tracks"))
									}), λ.ListType)
								if !λ.IsTrue(ϒtracks) {
									continue
								}
								τmp6 = λ.Cal(λ.BuiltinIter, ϒtracks)
								for {
									if τmp7 = λ.NextDefault(τmp6, λ.AfterLast); τmp7 == λ.AfterLast {
										break
									}
									ϒtrack = τmp7
									λ.Calm(ϒself, "_extract_format", ϒcontent_id, ϒformats, ϒtrack_uris, λ.DictLiteral(map[string]λ.Object{
										"url":      λ.Calm(ϒtrack, "get", λ.StrLiteral("uri")),
										"type":     λ.Calm(ϒf, "get", λ.StrLiteral("type")),
										"mimeType": λ.Calm(ϒf, "get", λ.StrLiteral("mimeType")),
										"quality":  λ.Calm(ϒquality, "get", λ.StrLiteral("quality")),
										"language": λ.Calm(ϒtrack, "get", λ.StrLiteral("language")),
									}))
								}
							}
						}
					}
					λ.Calm(ϒself, "_sort_formats", ϒformats)
					ϒduration = λ.Call(ϒfloat_or_none, λ.NewArgs(λ.Cal(ϒtry_get, ϒptmd, λ.NewFunction("<lambda>",
						[]λ.Param{
							{Name: "x"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒx = λargs[0]
							)
							return λ.GetItem(λ.GetItem(λ.GetItem(ϒx, λ.StrLiteral("attributes")), λ.StrLiteral("duration")), λ.StrLiteral("value"))
						}))), λ.KWArgs{
						{Name: "scale", Value: λ.IntLiteral(1000)},
					})
					return λ.DictLiteral(map[string]λ.Object{
						"extractor_key": λ.Calm(ZDFIE, "ie_key"),
						"id":            ϒcontent_id,
						"duration":      ϒduration,
						"formats":       ϒformats,
						"subtitles":     λ.Calm(ϒself, "_extract_subtitles", ϒptmd),
					})
				})
			ZDFBaseIE__extract_player = λ.NewFunction("_extract_player",
				[]λ.Param{
					{Name: "self"},
					{Name: "webpage"},
					{Name: "video_id"},
					{Name: "fatal", Def: λ.True},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒfatal    = λargs[3]
						ϒself     = λargs[0]
						ϒvideo_id = λargs[2]
						ϒwebpage  = λargs[1]
					)
					return λ.Calm(ϒself, "_parse_json", λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("(?s)data-zdfplayer-jsb=([\"\\'])(?P<json>{.+?})\\1"),
						ϒwebpage,
						λ.StrLiteral("player JSON"),
					), λ.KWArgs{
						{Name: "default", Value: func() λ.Object {
							if !λ.IsTrue(ϒfatal) {
								return λ.StrLiteral("{}")
							} else {
								return NO_DEFAULT
							}
						}()},
						{Name: "group", Value: λ.StrLiteral("json")},
					}), ϒvideo_id)
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_GEO_COUNTRIES":     ZDFBaseIE__GEO_COUNTRIES,
				"_QUALITIES":         ZDFBaseIE__QUALITIES,
				"_call_api":          ZDFBaseIE__call_api,
				"_extract_format":    ZDFBaseIE__extract_format,
				"_extract_player":    ZDFBaseIE__extract_player,
				"_extract_ptmd":      ZDFBaseIE__extract_ptmd,
				"_extract_subtitles": ZDFBaseIE__extract_subtitles,
			})
		}())
		ZDFIE = λ.Cal(λ.TypeType, λ.StrLiteral("ZDFIE"), λ.NewTuple(ZDFBaseIE), func() λ.Dict {
			var (
				ZDFIE__VALID_URL       λ.Object
				ZDFIE__extract_entry   λ.Object
				ZDFIE__extract_regular λ.Object
				ZDFIE__real_extract    λ.Object
			)
			ZDFIE__VALID_URL = λ.StrLiteral("https?://www\\.zdf\\.de/(?:[^/]+/)*(?P<id>[^/?#&]+)\\.html")
			ZDFIE__extract_entry = λ.NewFunction("_extract_entry",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
					{Name: "player"},
					{Name: "content"},
					{Name: "video_id"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒcontent    = λargs[3]
						ϒinfo       λ.Object
						ϒlayout_key λ.Object
						ϒlayout_url λ.Object
						ϒlayouts    λ.Object
						ϒmobj       λ.Object
						ϒplayer     = λargs[2]
						ϒptmd_path  λ.Object
						ϒself       = λargs[0]
						ϒt          λ.Object
						ϒthumbnail  λ.Object
						ϒthumbnails λ.Object
						ϒtitle      λ.Object
						ϒurl        = λargs[1]
						ϒvideo_id   = λargs[4]
						τmp0        λ.Object
						τmp1        λ.Object
						τmp2        λ.Object
					)
					ϒtitle = func() λ.Object {
						if λv := λ.Calm(ϒcontent, "get", λ.StrLiteral("title")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.GetItem(ϒcontent, λ.StrLiteral("teaserHeadline"))
						}
					}()
					ϒt = λ.GetItem(λ.GetItem(ϒcontent, λ.StrLiteral("mainVideoContent")), λ.StrLiteral("http://zdf.de/rels/target"))
					ϒptmd_path = λ.Calm(ϒt, "get", λ.StrLiteral("http://zdf.de/rels/streams/ptmd"))
					if !λ.IsTrue(ϒptmd_path) {
						ϒptmd_path = λ.Calm(λ.GetItem(ϒt, λ.StrLiteral("http://zdf.de/rels/streams/ptmd-template")), "replace", λ.StrLiteral("{playerId}"), λ.StrLiteral("ngplayer_2_4"))
					}
					ϒinfo = λ.Calm(ϒself, "_extract_ptmd", λ.Cal(ϒurljoin, ϒurl, ϒptmd_path), ϒvideo_id, λ.GetItem(ϒplayer, λ.StrLiteral("apiToken")), ϒurl)
					ϒthumbnails = λ.NewList()
					ϒlayouts = λ.Cal(ϒtry_get, ϒcontent, λ.NewFunction("<lambda>",
						[]λ.Param{
							{Name: "x"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒx = λargs[0]
							)
							return λ.GetItem(λ.GetItem(ϒx, λ.StrLiteral("teaserImageRef")), λ.StrLiteral("layouts"))
						}), λ.DictType)
					if λ.IsTrue(ϒlayouts) {
						τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(ϒlayouts, "items"))
						for {
							if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
								break
							}
							τmp2 = λ.UnpackIterable(τmp1, 2)
							ϒlayout_key = λ.GetItem(τmp2, λ.IntLiteral(0))
							ϒlayout_url = λ.GetItem(τmp2, λ.IntLiteral(1))
							ϒlayout_url = λ.Cal(ϒurl_or_none, ϒlayout_url)
							if !λ.IsTrue(ϒlayout_url) {
								continue
							}
							ϒthumbnail = λ.DictLiteral(map[string]λ.Object{
								"url":       ϒlayout_url,
								"format_id": ϒlayout_key,
							})
							ϒmobj = λ.Cal(Ωre.ϒsearch, λ.StrLiteral("(?P<width>\\d+)x(?P<height>\\d+)"), ϒlayout_key)
							if λ.IsTrue(ϒmobj) {
								λ.Calm(ϒthumbnail, "update", λ.DictLiteral(map[string]λ.Object{
									"width":  λ.Cal(λ.IntType, λ.Calm(ϒmobj, "group", λ.StrLiteral("width"))),
									"height": λ.Cal(λ.IntType, λ.Calm(ϒmobj, "group", λ.StrLiteral("height"))),
								}))
							}
							λ.Calm(ϒthumbnails, "append", ϒthumbnail)
						}
					}
					return λ.Cal(ϒmerge_dicts, ϒinfo, λ.DictLiteral(map[string]λ.Object{
						"title": ϒtitle,
						"description": func() λ.Object {
							if λv := λ.Calm(ϒcontent, "get", λ.StrLiteral("leadParagraph")); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Calm(ϒcontent, "get", λ.StrLiteral("teasertext"))
							}
						}(),
						"duration":   λ.Cal(ϒint_or_none, λ.Calm(ϒt, "get", λ.StrLiteral("duration"))),
						"timestamp":  λ.Cal(ϒunified_timestamp, λ.Calm(ϒcontent, "get", λ.StrLiteral("editorialDate"))),
						"thumbnails": ϒthumbnails,
					}))
				})
			ZDFIE__extract_regular = λ.NewFunction("_extract_regular",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
					{Name: "player"},
					{Name: "video_id"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒcontent  λ.Object
						ϒplayer   = λargs[2]
						ϒself     = λargs[0]
						ϒurl      = λargs[1]
						ϒvideo_id = λargs[3]
					)
					ϒcontent = λ.Calm(ϒself, "_call_api", λ.GetItem(ϒplayer, λ.StrLiteral("content")), ϒvideo_id, λ.StrLiteral("content"), λ.GetItem(ϒplayer, λ.StrLiteral("apiToken")), ϒurl)
					return λ.Calm(ϒself, "_extract_entry", λ.GetItem(ϒplayer, λ.StrLiteral("content")), ϒplayer, ϒcontent, ϒvideo_id)
				})
			ZDFIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒplayer   λ.Object
						ϒself     = λargs[0]
						ϒurl      = λargs[1]
						ϒvideo_id λ.Object
						ϒwebpage  λ.Object
					)
					ϒvideo_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒwebpage = λ.Call(λ.GetAttr(ϒself, "_download_webpage", nil), λ.NewArgs(
						ϒurl,
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					})
					if λ.IsTrue(ϒwebpage) {
						ϒplayer = λ.Call(λ.GetAttr(ϒself, "_extract_player", nil), λ.NewArgs(
							ϒwebpage,
							ϒurl,
						), λ.KWArgs{
							{Name: "fatal", Value: λ.False},
						})
						if λ.IsTrue(ϒplayer) {
							return λ.Calm(ϒself, "_extract_regular", ϒurl, ϒplayer, ϒvideo_id)
						}
					}
					return λ.Calm(ϒself, "_extract_mobile", ϒvideo_id)
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL":       ZDFIE__VALID_URL,
				"_extract_entry":   ZDFIE__extract_entry,
				"_extract_regular": ZDFIE__extract_regular,
				"_real_extract":    ZDFIE__real_extract,
			})
		}())
		ZDFChannelIE = λ.Cal(λ.TypeType, λ.StrLiteral("ZDFChannelIE"), λ.NewTuple(ZDFBaseIE), func() λ.Dict {
			var (
				ZDFChannelIE__VALID_URL λ.Object
				ZDFChannelIE_suitable   λ.Object
			)
			ZDFChannelIE__VALID_URL = λ.StrLiteral("https?://www\\.zdf\\.de/(?:[^/]+/)*(?P<id>[^/?#&]+)")
			ZDFChannelIE_suitable = λ.NewFunction("suitable",
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
						if λ.IsTrue(λ.Calm(ZDFIE, "suitable", ϒurl)) {
							return λ.False
						} else {
							return λ.Calm(λ.Cal(λ.SuperType, ZDFChannelIE, ϒcls), "suitable", ϒurl)
						}
					}()
				})
			ZDFChannelIE_suitable = λ.Cal(λ.ClassMethodType, ZDFChannelIE_suitable)
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": ZDFChannelIE__VALID_URL,
				"suitable":   ZDFChannelIE_suitable,
			})
		}())
	})
}
