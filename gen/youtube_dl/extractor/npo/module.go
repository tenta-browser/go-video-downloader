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
 * npo/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/npo.py
 */

package npo

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	AndereTijdenIE      λ.Object
	ExtractorError      λ.Object
	HetKlokhuisIE       λ.Object
	InfoExtractor       λ.Object
	NPOBaseIE           λ.Object
	NPODataMidEmbedIE   λ.Object
	NPOIE               λ.Object
	NPOLiveIE           λ.Object
	NPOPlaylistBaseIE   λ.Object
	NPORadioFragmentIE  λ.Object
	NPORadioIE          λ.Object
	SchoolTVIE          λ.Object
	VPROIE              λ.Object
	WNLIE               λ.Object
	ϒcompat_HTTPError   λ.Object
	ϒcompat_str         λ.Object
	ϒdetermine_ext      λ.Object
	ϒfix_xml_ampersands λ.Object
	ϒint_or_none        λ.Object
	ϒmerge_dicts        λ.Object
	ϒorderedSet         λ.Object
	ϒparse_duration     λ.Object
	ϒqualities          λ.Object
	ϒstr_or_none        λ.Object
	ϒstrip_jsonp        λ.Object
	ϒunified_strdate    λ.Object
	ϒunified_timestamp  λ.Object
	ϒurl_or_none        λ.Object
	ϒurlencode_postdata λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_HTTPError = Ωcompat.ϒcompat_HTTPError
		ϒcompat_str = Ωcompat.ϒcompat_str
		ϒdetermine_ext = Ωutils.ϒdetermine_ext
		ExtractorError = Ωutils.ExtractorError
		ϒfix_xml_ampersands = Ωutils.ϒfix_xml_ampersands
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒmerge_dicts = Ωutils.ϒmerge_dicts
		ϒorderedSet = Ωutils.ϒorderedSet
		ϒparse_duration = Ωutils.ϒparse_duration
		ϒqualities = Ωutils.ϒqualities
		ϒstr_or_none = Ωutils.ϒstr_or_none
		ϒstrip_jsonp = Ωutils.ϒstrip_jsonp
		ϒunified_strdate = Ωutils.ϒunified_strdate
		ϒunified_timestamp = Ωutils.ϒunified_timestamp
		ϒurl_or_none = Ωutils.ϒurl_or_none
		ϒurlencode_postdata = Ωutils.ϒurlencode_postdata
		NPOBaseIE = λ.Cal(λ.TypeType, λ.StrLiteral("NPOBaseIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {

			return λ.ClassDictLiteral(map[λ.Object]λ.Object{})
		}())
		NPOIE = λ.Cal(λ.TypeType, λ.StrLiteral("NPOIE"), λ.NewTuple(NPOBaseIE), func() λ.Dict {
			var (
				NPOIE_IE_NAME       λ.Object
				NPOIE__VALID_URL    λ.Object
				NPOIE__get_info     λ.Object
				NPOIE__get_old_info λ.Object
				NPOIE__real_extract λ.Object
				NPOIE_suitable      λ.Object
			)
			NPOIE_IE_NAME = λ.StrLiteral("npo")
			NPOIE__VALID_URL = λ.StrLiteral("(?x)\n                    (?:\n                        npo:|\n                        https?://\n                            (?:www\\.)?\n                            (?:\n                                npo\\.nl/(?:[^/]+/)*|\n                                (?:ntr|npostart)\\.nl/(?:[^/]+/){2,}|\n                                omroepwnl\\.nl/video/fragment/[^/]+__|\n                                (?:zapp|npo3)\\.nl/(?:[^/]+/){2,}\n                            )\n                        )\n                        (?P<id>[^/?#]+)\n                ")
			NPOIE_suitable = λ.NewFunction("suitable",
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
						if λ.IsTrue(λ.Cal(λ.BuiltinAny, λ.Cal(λ.NewFunction("<generator>",
							nil,
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								return λ.NewGenerator(func(λgy λ.Yielder) λ.Object {
									var (
										ϒie  λ.Object
										τmp0 λ.Object
										τmp1 λ.Object
									)
									τmp0 = λ.Cal(λ.BuiltinIter, λ.NewTuple(
										NPOLiveIE,
										NPORadioIE,
										NPORadioFragmentIE,
									))
									for {
										if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
											break
										}
										ϒie = τmp1
										λgy.Yield(λ.Calm(ϒie, "suitable", ϒurl))
									}
									return λ.None
								})
							})))) {
							return λ.False
						} else {
							return λ.Calm(λ.Cal(λ.SuperType, NPOIE, ϒcls), "suitable", ϒurl)
						}
					}()
				})
			NPOIE_suitable = λ.Cal(λ.ClassMethodType, NPOIE_suitable)
			NPOIE__real_extract = λ.NewFunction("_real_extract",
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
					return func() λ.Object {
						if λv := λ.Calm(ϒself, "_get_info", ϒurl, ϒvideo_id); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Calm(ϒself, "_get_old_info", ϒvideo_id)
						}
					}()
				})
			NPOIE__get_info = λ.NewFunction("_get_info",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
					{Name: "video_id"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒcc             λ.Object
						ϒcc_url         λ.Object
						ϒdrm            λ.Object
						ϒembed_url      λ.Object
						ϒformat_urls    λ.Object
						ϒformats        λ.Object
						ϒinfo           λ.Object
						ϒlang           λ.Object
						ϒplayer         λ.Object
						ϒplayer_token   λ.Object
						ϒprofile        λ.Object
						ϒself           = λargs[0]
						ϒstream         λ.Object
						ϒstream_ext     λ.Object
						ϒstream_type    λ.Object
						ϒstream_url     λ.Object
						ϒstreams        λ.Object
						ϒsubtitles      λ.Object
						ϒsubtitles_list λ.Object
						ϒtitle          λ.Object
						ϒtoken          λ.Object
						ϒurl            = λargs[1]
						ϒvideo          λ.Object
						ϒvideo_id       = λargs[2]
						ϒwebpage        λ.Object
						τmp0            λ.Object
						τmp1            λ.Object
					)
					ϒtoken = λ.GetItem(λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
						λ.StrLiteral("https://www.npostart.nl/api/token"),
						ϒvideo_id,
						λ.StrLiteral("Downloading token"),
					), λ.KWArgs{
						{Name: "headers", Value: λ.DictLiteral(map[string]λ.Object{
							"Referer":          ϒurl,
							"X-Requested-With": λ.StrLiteral("XMLHttpRequest"),
						})},
					}), λ.StrLiteral("token"))
					ϒplayer = λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
						λ.Mod(λ.StrLiteral("https://www.npostart.nl/player/%s"), ϒvideo_id),
						ϒvideo_id,
						λ.StrLiteral("Downloading player JSON"),
					), λ.KWArgs{
						{Name: "data", Value: λ.Cal(ϒurlencode_postdata, λ.DictLiteral(map[string]λ.Object{
							"autoplay":     λ.IntLiteral(0),
							"share":        λ.IntLiteral(1),
							"pageUrl":      ϒurl,
							"hasAdConsent": λ.IntLiteral(0),
							"_token":       ϒtoken,
						}))},
					})
					ϒplayer_token = λ.GetItem(ϒplayer, λ.StrLiteral("token"))
					ϒdrm = λ.False
					ϒformat_urls = λ.Cal(λ.SetType)
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.NewTuple(
						λ.StrLiteral("hls"),
						λ.StrLiteral("dash-widevine"),
						λ.StrLiteral("dash-playready"),
						λ.StrLiteral("smooth"),
					))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒprofile = τmp1
						ϒstreams = λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
							λ.Mod(λ.StrLiteral("https://start-player.npo.nl/video/%s/streams"), ϒvideo_id),
							ϒvideo_id,
							λ.Mod(λ.StrLiteral("Downloading %s profile JSON"), ϒprofile),
						), λ.KWArgs{
							{Name: "fatal", Value: λ.False},
							{Name: "query", Value: λ.DictLiteral(map[string]λ.Object{
								"profile":    ϒprofile,
								"quality":    λ.StrLiteral("npo"),
								"tokenId":    ϒplayer_token,
								"streamType": λ.StrLiteral("broadcast"),
							})},
						})
						if !λ.IsTrue(ϒstreams) {
							continue
						}
						ϒstream = λ.Calm(ϒstreams, "get", λ.StrLiteral("stream"))
						if !λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒstream, λ.DictType)) {
							continue
						}
						ϒstream_url = λ.Cal(ϒurl_or_none, λ.Calm(ϒstream, "get", λ.StrLiteral("src")))
						if λ.IsTrue(func() λ.Object {
							if λv := λ.NewBool(!λ.IsTrue(ϒstream_url)); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.NewBool(λ.Contains(ϒformat_urls, ϒstream_url))
							}
						}()) {
							continue
						}
						λ.Calm(ϒformat_urls, "add", ϒstream_url)
						if λ.IsTrue(func() λ.Object {
							if λv := λ.NewBool(λ.Calm(ϒstream, "get", λ.StrLiteral("protection")) != λ.None); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.NewBool(λ.Calm(ϒstream, "get", λ.StrLiteral("keySystemOptions")) != λ.None)
							}
						}()) {
							ϒdrm = λ.True
							continue
						}
						ϒstream_type = λ.Calm(ϒstream, "get", λ.StrLiteral("type"))
						ϒstream_ext = λ.Cal(ϒdetermine_ext, ϒstream_url)
						if λ.IsTrue(func() λ.Object {
							if λv := λ.Eq(ϒstream_type, λ.StrLiteral("application/dash+xml")); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Eq(ϒstream_ext, λ.StrLiteral("mpd"))
							}
						}()) {
							λ.Calm(ϒformats, "extend", λ.Call(λ.GetAttr(ϒself, "_extract_mpd_formats", nil), λ.NewArgs(
								ϒstream_url,
								ϒvideo_id,
							), λ.KWArgs{
								{Name: "mpd_id", Value: λ.StrLiteral("dash")},
								{Name: "fatal", Value: λ.False},
							}))
						} else {
							if λ.IsTrue(func() λ.Object {
								if λv := λ.Eq(ϒstream_type, λ.StrLiteral("application/vnd.apple.mpegurl")); λ.IsTrue(λv) {
									return λv
								} else {
									return λ.Eq(ϒstream_ext, λ.StrLiteral("m3u8"))
								}
							}()) {
								λ.Calm(ϒformats, "extend", λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
									ϒstream_url,
									ϒvideo_id,
								), λ.KWArgs{
									{Name: "ext", Value: λ.StrLiteral("mp4")},
									{Name: "entry_protocol", Value: λ.StrLiteral("m3u8_native")},
									{Name: "m3u8_id", Value: λ.StrLiteral("hls")},
									{Name: "fatal", Value: λ.False},
								}))
							} else {
								if λ.IsTrue(λ.Cal(Ωre.ϒsearch, λ.StrLiteral("\\.isml?/Manifest"), ϒstream_url)) {
									λ.Calm(ϒformats, "extend", λ.Call(λ.GetAttr(ϒself, "_extract_ism_formats", nil), λ.NewArgs(
										ϒstream_url,
										ϒvideo_id,
									), λ.KWArgs{
										{Name: "ism_id", Value: λ.StrLiteral("mss")},
										{Name: "fatal", Value: λ.False},
									}))
								} else {
									λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
										"url": ϒstream_url,
									}))
								}
							}
						}
					}
					if !λ.IsTrue(ϒformats) {
						if λ.IsTrue(ϒdrm) {
							panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.StrLiteral("This video is DRM protected.")), λ.KWArgs{
								{Name: "expected", Value: λ.True},
							})))
						}
						return λ.None
					}
					λ.Calm(ϒself, "_sort_formats", ϒformats)
					ϒinfo = λ.DictLiteral(map[string]λ.Object{
						"id":      ϒvideo_id,
						"title":   ϒvideo_id,
						"formats": ϒformats,
					})
					ϒembed_url = λ.Cal(ϒurl_or_none, λ.Calm(ϒplayer, "get", λ.StrLiteral("embedUrl")))
					if λ.IsTrue(ϒembed_url) {
						ϒwebpage = λ.Call(λ.GetAttr(ϒself, "_download_webpage", nil), λ.NewArgs(
							ϒembed_url,
							ϒvideo_id,
							λ.StrLiteral("Downloading embed page"),
						), λ.KWArgs{
							{Name: "fatal", Value: λ.False},
						})
						if λ.IsTrue(ϒwebpage) {
							ϒvideo = λ.Calm(ϒself, "_parse_json", λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
								λ.StrLiteral("\\bvideo\\s*=\\s*({.+?})\\s*;"),
								ϒwebpage,
								λ.StrLiteral("video"),
							), λ.KWArgs{
								{Name: "default", Value: λ.StrLiteral("{}")},
							}), ϒvideo_id)
							if λ.IsTrue(ϒvideo) {
								ϒtitle = λ.Calm(ϒvideo, "get", λ.StrLiteral("episodeTitle"))
								ϒsubtitles = λ.DictLiteral(map[λ.Object]λ.Object{})
								ϒsubtitles_list = λ.Calm(ϒvideo, "get", λ.StrLiteral("subtitles"))
								if λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒsubtitles_list, λ.ListType)) {
									τmp0 = λ.Cal(λ.BuiltinIter, ϒsubtitles_list)
									for {
										if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
											break
										}
										ϒcc = τmp1
										ϒcc_url = λ.Cal(ϒurl_or_none, λ.Calm(ϒcc, "get", λ.StrLiteral("src")))
										if !λ.IsTrue(ϒcc_url) {
											continue
										}
										ϒlang = func() λ.Object {
											if λv := λ.Cal(ϒstr_or_none, λ.Calm(ϒcc, "get", λ.StrLiteral("language"))); λ.IsTrue(λv) {
												return λv
											} else {
												return λ.StrLiteral("nl")
											}
										}()
										λ.Calm(λ.Calm(ϒsubtitles, "setdefault", ϒlang, λ.NewList()), "append", λ.DictLiteral(map[string]λ.Object{
											"url": ϒcc_url,
										}))
									}
								}
								return λ.Cal(ϒmerge_dicts, λ.DictLiteral(map[string]λ.Object{
									"title":       ϒtitle,
									"description": λ.Calm(ϒvideo, "get", λ.StrLiteral("description")),
									"thumbnail": λ.Cal(ϒurl_or_none, func() λ.Object {
										if λv := λ.Calm(ϒvideo, "get", λ.StrLiteral("still_image_url")); λ.IsTrue(λv) {
											return λv
										} else {
											return λ.Calm(ϒvideo, "get", λ.StrLiteral("orig_image_url"))
										}
									}()),
									"duration":       λ.Cal(ϒint_or_none, λ.Calm(ϒvideo, "get", λ.StrLiteral("duration"))),
									"timestamp":      λ.Cal(ϒunified_timestamp, λ.Calm(ϒvideo, "get", λ.StrLiteral("broadcastDate"))),
									"creator":        λ.Calm(ϒvideo, "get", λ.StrLiteral("channel")),
									"series":         λ.Calm(ϒvideo, "get", λ.StrLiteral("title")),
									"episode":        ϒtitle,
									"episode_number": λ.Cal(ϒint_or_none, λ.Calm(ϒvideo, "get", λ.StrLiteral("episodeNumber"))),
									"subtitles":      ϒsubtitles,
								}), ϒinfo)
							}
						}
					}
					return ϒinfo
				})
			NPOIE__get_old_info = λ.NewFunction("_get_old_info",
				[]λ.Param{
					{Name: "self"},
					{Name: "video_id"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						QUALITY_FORMATS         λ.Object
						QUALITY_LABELS          λ.Object
						ϒadd_format_url         λ.Object
						ϒasx                    λ.Object
						ϒerror                  λ.Object
						ϒf4m_format             λ.Object
						ϒf4m_formats            λ.Object
						ϒformat_id              λ.Object
						ϒformats                λ.Object
						ϒis_legal_url           λ.Object
						ϒis_live                λ.Object
						ϒitem                   λ.Object
						ϒitem_label             λ.Object
						ϒitem_url               λ.Object
						ϒitems                  λ.Object
						ϒmetadata               λ.Object
						ϒnum                    λ.Object
						ϒquality_from_format_id λ.Object
						ϒquality_from_label     λ.Object
						ϒref                    λ.Object
						ϒself                   = λargs[0]
						ϒstream                 λ.Object
						ϒstream_info            λ.Object
						ϒstream_type            λ.Object
						ϒstream_url             λ.Object
						ϒsub_title              λ.Object
						ϒsubtitles              λ.Object
						ϒtitle                  λ.Object
						ϒtoken                  λ.Object
						ϒurls                   λ.Object
						ϒvideo_id               = λargs[1]
						ϒvideo_url              λ.Object
						τmp0                    λ.Object
						τmp1                    λ.Object
						τmp2                    λ.Object
						τmp3                    λ.Object
					)
					ϒmetadata = λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
						λ.Mod(λ.StrLiteral("http://e.omroep.nl/metadata/%s"), ϒvideo_id),
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "transform_source", Value: ϒstrip_jsonp},
					})
					ϒerror = λ.Calm(ϒmetadata, "get", λ.StrLiteral("error"))
					if λ.IsTrue(ϒerror) {
						panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(ϒerror), λ.KWArgs{
							{Name: "expected", Value: λ.True},
						})))
					}
					ϒvideo_id = func() λ.Object {
						if λv := λ.Calm(ϒmetadata, "get", λ.StrLiteral("prid")); λ.IsTrue(λv) {
							return λv
						} else {
							return ϒvideo_id
						}
					}()
					ϒtitle = λ.GetItem(ϒmetadata, λ.StrLiteral("titel"))
					ϒsub_title = λ.Calm(ϒmetadata, "get", λ.StrLiteral("aflevering_titel"))
					if λ.IsTrue(func() λ.Object {
						if λv := ϒsub_title; !λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Ne(ϒsub_title, ϒtitle)
						}
					}()) {
						τmp0 = λ.IAdd(ϒtitle, λ.Mod(λ.StrLiteral(": %s"), ϒsub_title))
						ϒtitle = τmp0
					}
					ϒtoken = λ.Calm(ϒself, "_get_token", ϒvideo_id)
					ϒformats = λ.NewList()
					ϒurls = λ.Cal(λ.SetType)
					ϒis_legal_url = λ.NewFunction("is_legal_url",
						[]λ.Param{
							{Name: "format_url"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒformat_url = λargs[0]
							)
							return func() λ.Object {
								if λv := ϒformat_url; !λ.IsTrue(λv) {
									return λv
								} else if λv := λ.NewBool(!λ.Contains(ϒurls, ϒformat_url)); !λ.IsTrue(λv) {
									return λv
								} else {
									return λ.Cal(Ωre.ϒmatch, λ.StrLiteral("^(?:https?:)?//"), ϒformat_url)
								}
							}()
						})
					QUALITY_LABELS = λ.NewTuple(
						λ.StrLiteral("Laag"),
						λ.StrLiteral("Normaal"),
						λ.StrLiteral("Hoog"),
					)
					QUALITY_FORMATS = λ.NewTuple(
						λ.StrLiteral("adaptive"),
						λ.StrLiteral("wmv_sb"),
						λ.StrLiteral("h264_sb"),
						λ.StrLiteral("wmv_bb"),
						λ.StrLiteral("h264_bb"),
						λ.StrLiteral("wvc1_std"),
						λ.StrLiteral("h264_std"),
					)
					ϒquality_from_label = λ.Cal(ϒqualities, QUALITY_LABELS)
					ϒquality_from_format_id = λ.Cal(ϒqualities, QUALITY_FORMATS)
					ϒitems = λ.GetItem(λ.GetItem(λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
						λ.Mod(λ.StrLiteral("http://ida.omroep.nl/app.php/%s"), ϒvideo_id),
						ϒvideo_id,
						λ.StrLiteral("Downloading formats JSON"),
					), λ.KWArgs{
						{Name: "query", Value: λ.DictLiteral(map[string]λ.Object{
							"adaptive": λ.StrLiteral("yes"),
							"token":    ϒtoken,
						})},
					}), λ.StrLiteral("items")), λ.IntLiteral(0))
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.EnumerateIteratorType, ϒitems))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						τmp2 = τmp1
						ϒnum = λ.GetItem(τmp2, λ.IntLiteral(0))
						ϒitem = λ.GetItem(τmp2, λ.IntLiteral(1))
						ϒitem_url = λ.Calm(ϒitem, "get", λ.StrLiteral("url"))
						if !λ.IsTrue(λ.Cal(ϒis_legal_url, ϒitem_url)) {
							continue
						}
						λ.Calm(ϒurls, "add", ϒitem_url)
						ϒformat_id = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
							λ.StrLiteral("video/ida/([^/]+)"),
							ϒitem_url,
							λ.StrLiteral("format id"),
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
						})
						ϒitem_label = λ.Calm(ϒitem, "get", λ.StrLiteral("label"))
						ϒadd_format_url = λ.NewFunction("add_format_url",
							[]λ.Param{
								{Name: "format_url"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒf_id       λ.Object
									ϒformat_url = λargs[0]
									ϒheight     λ.Object
									ϒquality    λ.Object
									ϒwidth      λ.Object
									τmp0        λ.Object
								)
								ϒwidth = λ.Cal(ϒint_or_none, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
									λ.StrLiteral("(\\d+)[xX]\\d+"),
									ϒformat_url,
									λ.StrLiteral("width"),
								), λ.KWArgs{
									{Name: "default", Value: λ.None},
								}))
								ϒheight = λ.Cal(ϒint_or_none, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
									λ.StrLiteral("\\d+[xX](\\d+)"),
									ϒformat_url,
									λ.StrLiteral("height"),
								), λ.KWArgs{
									{Name: "default", Value: λ.None},
								}))
								if λ.Contains(QUALITY_LABELS, ϒitem_label) {
									ϒquality = λ.Cal(ϒquality_from_label, ϒitem_label)
									ϒf_id = ϒitem_label
								} else {
									if λ.Contains(QUALITY_FORMATS, ϒitem_label) {
										ϒquality = λ.Cal(ϒquality_from_format_id, ϒformat_id)
										ϒf_id = ϒformat_id
									} else {
										τmp0 = λ.Mul(λ.NewList(λ.None), λ.IntLiteral(2))
										ϒquality = λ.GetItem(τmp0, λ.IntLiteral(0))
										ϒf_id = λ.GetItem(τmp0, λ.IntLiteral(1))
									}
								}
								λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
									"url":       ϒformat_url,
									"format_id": ϒf_id,
									"width":     ϒwidth,
									"height":    ϒheight,
									"quality":   ϒquality,
								}))
								return λ.None
							})
						if λ.Contains(λ.NewTuple(
							λ.StrLiteral("url"),
							λ.StrLiteral("audio"),
						), λ.Calm(ϒitem, "get", λ.StrLiteral("contentType"))) {
							λ.Cal(ϒadd_format_url, ϒitem_url)
							continue
						}
						τmp2, τmp3 = func() (λexit λ.Object, λret λ.Object) {
							defer λ.CatchMulti(
								nil,
								&λ.Catcher{ExtractorError, func(λex λ.BaseException) {
									var ϒee λ.Object = λex
									if λ.IsTrue(func() λ.Object {
										if λv := λ.Cal(λ.BuiltinIsInstance, λ.GetAttr(ϒee, "cause", nil), ϒcompat_HTTPError); !λ.IsTrue(λv) {
											return λv
										} else {
											return λ.Eq(λ.GetAttr(λ.GetAttr(ϒee, "cause", nil), "code", nil), λ.IntLiteral(404))
										}
									}()) {
										ϒerror = λ.Calm(func() λ.Object {
											if λv := λ.Call(λ.GetAttr(ϒself, "_parse_json", nil), λ.NewArgs(
												λ.Calm(λ.Calm(λ.GetAttr(ϒee, "cause", nil), "read"), "decode"),
												ϒvideo_id,
											), λ.KWArgs{
												{Name: "fatal", Value: λ.False},
											}); λ.IsTrue(λv) {
												return λv
											} else {
												return λ.DictLiteral(map[λ.Object]λ.Object{})
											}
										}(), "get", λ.StrLiteral("errorstring"))
										if λ.IsTrue(ϒerror) {
											panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(ϒerror), λ.KWArgs{
												{Name: "expected", Value: λ.True},
											})))
										}
									}
									panic(λ.Raise(λex))
								}},
							)
							ϒstream_info = λ.Calm(ϒself, "_download_json", λ.Add(ϒitem_url, λ.StrLiteral("&type=json")), ϒvideo_id, func() λ.Object {
								if λv := λ.Mod(λ.StrLiteral("Downloading %s stream JSON"), ϒitem_label); λ.IsTrue(λv) {
									return λv
								} else if λv := λ.Calm(ϒitem, "get", λ.StrLiteral("format")); λ.IsTrue(λv) {
									return λv
								} else if λv := ϒformat_id; λ.IsTrue(λv) {
									return λv
								} else {
									return ϒnum
								}
							}())
							return λ.BlockExitNormally, nil
						}()
						if λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒstream_info, ϒcompat_str)) {
							if !λ.IsTrue(λ.Calm(ϒstream_info, "startswith", λ.StrLiteral("http"))) {
								continue
							}
							ϒvideo_url = ϒstream_info
						} else {
							ϒvideo_url = λ.Calm(ϒstream_info, "get", λ.StrLiteral("url"))
						}
						if λ.IsTrue(func() λ.Object {
							if λv := λ.NewBool(!λ.IsTrue(ϒvideo_url)); λ.IsTrue(λv) {
								return λv
							} else if λv := λ.NewBool(λ.Contains(ϒvideo_url, λ.StrLiteral("vodnotavailable."))); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.NewBool(λ.Contains(ϒurls, ϒvideo_url))
							}
						}()) {
							continue
						}
						λ.Calm(ϒurls, "add", ϒvideo_url)
						if λ.IsTrue(λ.Eq(λ.Cal(ϒdetermine_ext, ϒvideo_url), λ.StrLiteral("m3u8"))) {
							λ.Calm(ϒformats, "extend", λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
								ϒvideo_url,
								ϒvideo_id,
							), λ.KWArgs{
								{Name: "ext", Value: λ.StrLiteral("mp4")},
								{Name: "entry_protocol", Value: λ.StrLiteral("m3u8_native")},
								{Name: "m3u8_id", Value: λ.StrLiteral("hls")},
								{Name: "fatal", Value: λ.False},
							}))
						} else {
							λ.Cal(ϒadd_format_url, ϒvideo_url)
						}
					}
					ϒis_live = λ.Eq(λ.Calm(ϒmetadata, "get", λ.StrLiteral("medium")), λ.StrLiteral("live"))
					if !λ.IsTrue(ϒis_live) {
						τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.EnumerateIteratorType, λ.Calm(ϒmetadata, "get", λ.StrLiteral("streams"), λ.NewList())))
						for {
							if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
								break
							}
							τmp3 = τmp1
							ϒnum = λ.GetItem(τmp3, λ.IntLiteral(0))
							ϒstream = λ.GetItem(τmp3, λ.IntLiteral(1))
							ϒstream_url = λ.Calm(ϒstream, "get", λ.StrLiteral("url"))
							if !λ.IsTrue(λ.Cal(ϒis_legal_url, ϒstream_url)) {
								continue
							}
							λ.Calm(ϒurls, "add", ϒstream_url)
							ϒstream_type = λ.Calm(λ.Calm(ϒstream, "get", λ.StrLiteral("type"), λ.StrLiteral("")), "lower")
							if λ.Contains(λ.NewList(
								λ.StrLiteral("ss"),
								λ.StrLiteral("ms"),
							), ϒstream_type) {
								continue
							}
							if λ.IsTrue(λ.Eq(ϒstream_type, λ.StrLiteral("hds"))) {
								ϒf4m_formats = λ.Call(λ.GetAttr(ϒself, "_extract_f4m_formats", nil), λ.NewArgs(
									ϒstream_url,
									ϒvideo_id,
								), λ.KWArgs{
									{Name: "fatal", Value: λ.False},
								})
								τmp3 = λ.Cal(λ.BuiltinIter, ϒf4m_formats)
								for {
									if τmp2 = λ.NextDefault(τmp3, λ.AfterLast); τmp2 == λ.AfterLast {
										break
									}
									ϒf4m_format = τmp2
									λ.SetItem(ϒf4m_format, λ.StrLiteral("preference"), λ.Neg(λ.IntLiteral(1)))
								}
								λ.Calm(ϒformats, "extend", ϒf4m_formats)
							} else {
								if λ.IsTrue(λ.Eq(ϒstream_type, λ.StrLiteral("hls"))) {
									λ.Calm(ϒformats, "extend", λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
										ϒstream_url,
										ϒvideo_id,
									), λ.KWArgs{
										{Name: "ext", Value: λ.StrLiteral("mp4")},
										{Name: "fatal", Value: λ.False},
									}))
								} else {
									if λ.Contains(ϒstream_url, λ.StrLiteral(".asf")) {
										ϒasx = λ.Call(λ.GetAttr(ϒself, "_download_xml", nil), λ.NewArgs(
											ϒstream_url,
											ϒvideo_id,
											λ.Mod(λ.StrLiteral("Downloading stream %d ASX playlist"), ϒnum),
										), λ.KWArgs{
											{Name: "transform_source", Value: ϒfix_xml_ampersands},
											{Name: "fatal", Value: λ.False},
										})
										if !λ.IsTrue(ϒasx) {
											continue
										}
										ϒref = λ.Calm(ϒasx, "find", λ.StrLiteral("./ENTRY/Ref"))
										if ϒref == λ.None {
											continue
										}
										ϒvideo_url = λ.Calm(ϒref, "get", λ.StrLiteral("href"))
										if λ.IsTrue(func() λ.Object {
											if λv := λ.NewBool(!λ.IsTrue(ϒvideo_url)); λ.IsTrue(λv) {
												return λv
											} else {
												return λ.NewBool(λ.Contains(ϒurls, ϒvideo_url))
											}
										}()) {
											continue
										}
										λ.Calm(ϒurls, "add", ϒvideo_url)
										λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
											"url":        ϒvideo_url,
											"ext":        λ.Calm(ϒstream, "get", λ.StrLiteral("formaat"), λ.StrLiteral("asf")),
											"quality":    λ.Calm(ϒstream, "get", λ.StrLiteral("kwaliteit")),
											"preference": λ.Neg(λ.IntLiteral(10)),
										}))
									} else {
										λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
											"url":     ϒstream_url,
											"quality": λ.Calm(ϒstream, "get", λ.StrLiteral("kwaliteit")),
										}))
									}
								}
							}
						}
					}
					λ.Calm(ϒself, "_sort_formats", ϒformats)
					ϒsubtitles = λ.DictLiteral(map[λ.Object]λ.Object{})
					if λ.IsTrue(λ.Eq(λ.Calm(ϒmetadata, "get", λ.StrLiteral("tt888")), λ.StrLiteral("ja"))) {
						λ.SetItem(ϒsubtitles, λ.StrLiteral("nl"), λ.NewList(λ.DictLiteral(map[string]λ.Object{
							"ext": λ.StrLiteral("vtt"),
							"url": λ.Mod(λ.StrLiteral("http://tt888.omroep.nl/tt888/%s"), ϒvideo_id),
						})))
					}
					return λ.DictLiteral(map[string]λ.Object{
						"id": ϒvideo_id,
						"title": func() λ.Object {
							if λ.IsTrue(ϒis_live) {
								return λ.Calm(ϒself, "_live_title", ϒtitle)
							} else {
								return ϒtitle
							}
						}(),
						"description": λ.Calm(ϒmetadata, "get", λ.StrLiteral("info")),
						"thumbnail": λ.GetItem(λ.GetItem(λ.Calm(ϒmetadata, "get", λ.StrLiteral("images"), λ.NewList(λ.DictLiteral(map[string]λ.Object{
							"url": λ.None,
						}))), λ.Neg(λ.IntLiteral(1))), λ.StrLiteral("url")),
						"upload_date": λ.Cal(ϒunified_strdate, λ.Calm(ϒmetadata, "get", λ.StrLiteral("gidsdatum"))),
						"duration":    λ.Cal(ϒparse_duration, λ.Calm(ϒmetadata, "get", λ.StrLiteral("tijdsduur"))),
						"formats":     ϒformats,
						"subtitles":   ϒsubtitles,
						"is_live":     ϒis_live,
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"IE_NAME":       NPOIE_IE_NAME,
				"_VALID_URL":    NPOIE__VALID_URL,
				"_get_info":     NPOIE__get_info,
				"_get_old_info": NPOIE__get_old_info,
				"_real_extract": NPOIE__real_extract,
				"suitable":      NPOIE_suitable,
			})
		}())
		NPOLiveIE = λ.Cal(λ.TypeType, λ.StrLiteral("NPOLiveIE"), λ.NewTuple(NPOBaseIE), func() λ.Dict {
			var (
				NPOLiveIE__VALID_URL λ.Object
			)
			NPOLiveIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?npo(?:start)?\\.nl/live(?:/(?P<id>[^/?#&]+))?")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": NPOLiveIE__VALID_URL,
			})
		}())
		NPORadioIE = λ.Cal(λ.TypeType, λ.StrLiteral("NPORadioIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				NPORadioIE__VALID_URL λ.Object
				NPORadioIE_suitable   λ.Object
			)
			NPORadioIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?npo\\.nl/radio/(?P<id>[^/]+)")
			NPORadioIE_suitable = λ.NewFunction("suitable",
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
						if λ.IsTrue(λ.Calm(NPORadioFragmentIE, "suitable", ϒurl)) {
							return λ.False
						} else {
							return λ.Calm(λ.Cal(λ.SuperType, NPORadioIE, ϒcls), "suitable", ϒurl)
						}
					}()
				})
			NPORadioIE_suitable = λ.Cal(λ.ClassMethodType, NPORadioIE_suitable)
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": NPORadioIE__VALID_URL,
				"suitable":   NPORadioIE_suitable,
			})
		}())
		NPORadioFragmentIE = λ.Cal(λ.TypeType, λ.StrLiteral("NPORadioFragmentIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				NPORadioFragmentIE__VALID_URL λ.Object
			)
			NPORadioFragmentIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?npo\\.nl/radio/[^/]+/fragment/(?P<id>\\d+)")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": NPORadioFragmentIE__VALID_URL,
			})
		}())
		NPODataMidEmbedIE = λ.Cal(λ.TypeType, λ.StrLiteral("NPODataMidEmbedIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {

			return λ.ClassDictLiteral(map[λ.Object]λ.Object{})
		}())
		SchoolTVIE = λ.Cal(λ.TypeType, λ.StrLiteral("SchoolTVIE"), λ.NewTuple(NPODataMidEmbedIE), func() λ.Dict {
			var (
				SchoolTVIE__VALID_URL λ.Object
			)
			SchoolTVIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?schooltv\\.nl/video/(?P<id>[^/?#&]+)")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": SchoolTVIE__VALID_URL,
			})
		}())
		HetKlokhuisIE = λ.Cal(λ.TypeType, λ.StrLiteral("HetKlokhuisIE"), λ.NewTuple(NPODataMidEmbedIE), func() λ.Dict {
			var (
				HetKlokhuisIE__VALID_URL λ.Object
			)
			HetKlokhuisIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?hetklokhuis\\.nl/[^/]+/\\d+/(?P<id>[^/?#&]+)")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": HetKlokhuisIE__VALID_URL,
			})
		}())
		NPOPlaylistBaseIE = λ.Cal(λ.TypeType, λ.StrLiteral("NPOPlaylistBaseIE"), λ.NewTuple(NPOIE), func() λ.Dict {

			return λ.ClassDictLiteral(map[λ.Object]λ.Object{})
		}())
		VPROIE = λ.Cal(λ.TypeType, λ.StrLiteral("VPROIE"), λ.NewTuple(NPOPlaylistBaseIE), func() λ.Dict {
			var (
				VPROIE__VALID_URL λ.Object
			)
			VPROIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?(?:(?:tegenlicht\\.)?vpro|2doc)\\.nl/(?:[^/]+/)*(?P<id>[^/]+)\\.html")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": VPROIE__VALID_URL,
			})
		}())
		WNLIE = λ.Cal(λ.TypeType, λ.StrLiteral("WNLIE"), λ.NewTuple(NPOPlaylistBaseIE), func() λ.Dict {
			var (
				WNLIE__VALID_URL λ.Object
			)
			WNLIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?omroepwnl\\.nl/video/detail/(?P<id>[^/]+)__\\d+")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": WNLIE__VALID_URL,
			})
		}())
		AndereTijdenIE = λ.Cal(λ.TypeType, λ.StrLiteral("AndereTijdenIE"), λ.NewTuple(NPOPlaylistBaseIE), func() λ.Dict {
			var (
				AndereTijdenIE__VALID_URL λ.Object
			)
			AndereTijdenIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?anderetijden\\.nl/programma/(?:[^/]+/)+(?P<id>[^/?#&]+)")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": AndereTijdenIE__VALID_URL,
			})
		}())
	})
}
