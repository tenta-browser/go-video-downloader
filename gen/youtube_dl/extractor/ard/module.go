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
 * ard/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/ard.py
 */

package ard

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ARDBetaMediathekIE       λ.Object
	ARDIE                    λ.Object
	ARDMediathekBaseIE       λ.Object
	ARDMediathekIE           λ.Object
	ExtractorError           λ.Object
	InfoExtractor            λ.Object
	ϒcompat_etree_fromstring λ.Object
	ϒdetermine_ext           λ.Object
	ϒint_or_none             λ.Object
	ϒparse_duration          λ.Object
	ϒqualities               λ.Object
	ϒstr_or_none             λ.Object
	ϒtry_get                 λ.Object
	ϒunified_strdate         λ.Object
	ϒunified_timestamp       λ.Object
	ϒupdate_url_query        λ.Object
	ϒurl_or_none             λ.Object
	ϒxpath_text              λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒdetermine_ext = Ωutils.ϒdetermine_ext
		ExtractorError = Ωutils.ExtractorError
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒparse_duration = Ωutils.ϒparse_duration
		ϒqualities = Ωutils.ϒqualities
		ϒstr_or_none = Ωutils.ϒstr_or_none
		ϒtry_get = Ωutils.ϒtry_get
		ϒunified_strdate = Ωutils.ϒunified_strdate
		ϒunified_timestamp = Ωutils.ϒunified_timestamp
		ϒupdate_url_query = Ωutils.ϒupdate_url_query
		ϒurl_or_none = Ωutils.ϒurl_or_none
		ϒxpath_text = Ωutils.ϒxpath_text
		ϒcompat_etree_fromstring = Ωcompat.ϒcompat_etree_fromstring
		ARDMediathekBaseIE = λ.Cal(λ.TypeType, λ.StrLiteral("ARDMediathekBaseIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				ARDMediathekBaseIE__GEO_COUNTRIES      λ.Object
				ARDMediathekBaseIE__extract_formats    λ.Object
				ARDMediathekBaseIE__extract_media_info λ.Object
				ARDMediathekBaseIE__parse_media_info   λ.Object
			)
			ARDMediathekBaseIE__GEO_COUNTRIES = λ.NewList(λ.StrLiteral("DE"))
			ARDMediathekBaseIE__extract_media_info = λ.NewFunction("_extract_media_info",
				[]λ.Param{
					{Name: "self"},
					{Name: "media_info_url"},
					{Name: "webpage"},
					{Name: "video_id"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒmedia_info     λ.Object
						ϒmedia_info_url = λargs[1]
						ϒself           = λargs[0]
						ϒvideo_id       = λargs[3]
						ϒwebpage        = λargs[2]
					)
					ϒmedia_info = λ.Calm(ϒself, "_download_json", ϒmedia_info_url, ϒvideo_id, λ.StrLiteral("Downloading media JSON"))
					return λ.Calm(ϒself, "_parse_media_info", ϒmedia_info, ϒvideo_id, λ.NewBool(λ.Contains(ϒwebpage, λ.StrLiteral("\"fsk\""))))
				})
			ARDMediathekBaseIE__parse_media_info = λ.NewFunction("_parse_media_info",
				[]λ.Param{
					{Name: "self"},
					{Name: "media_info"},
					{Name: "video_id"},
					{Name: "fsk"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒformats      λ.Object
						ϒfsk          = λargs[3]
						ϒmedia_info   = λargs[1]
						ϒself         = λargs[0]
						ϒsubtitle_url λ.Object
						ϒsubtitles    λ.Object
						ϒvideo_id     = λargs[2]
					)
					ϒformats = λ.Calm(ϒself, "_extract_formats", ϒmedia_info, ϒvideo_id)
					if !λ.IsTrue(ϒformats) {
						if λ.IsTrue(ϒfsk) {
							panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.StrLiteral("This video is only available after 20:00")), λ.KWArgs{
								{Name: "expected", Value: λ.True},
							})))
						} else {
							if λ.IsTrue(λ.Calm(ϒmedia_info, "get", λ.StrLiteral("_geoblocked"))) {
								λ.Call(λ.GetAttr(ϒself, "raise_geo_restricted", nil), λ.NewArgs(λ.StrLiteral("This video is not available due to geoblocking")), λ.KWArgs{
									{Name: "countries", Value: λ.GetAttr(ϒself, "_GEO_COUNTRIES", nil)},
								})
							}
						}
					}
					λ.Calm(ϒself, "_sort_formats", ϒformats)
					ϒsubtitles = λ.DictLiteral(map[λ.Object]λ.Object{})
					ϒsubtitle_url = λ.Calm(ϒmedia_info, "get", λ.StrLiteral("_subtitleUrl"))
					if λ.IsTrue(ϒsubtitle_url) {
						λ.SetItem(ϒsubtitles, λ.StrLiteral("de"), λ.NewList(λ.DictLiteral(map[string]λ.Object{
							"ext": λ.StrLiteral("ttml"),
							"url": ϒsubtitle_url,
						})))
					}
					return λ.DictLiteral(map[string]λ.Object{
						"id":        ϒvideo_id,
						"duration":  λ.Cal(ϒint_or_none, λ.Calm(ϒmedia_info, "get", λ.StrLiteral("_duration"))),
						"thumbnail": λ.Calm(ϒmedia_info, "get", λ.StrLiteral("_previewImage")),
						"is_live":   λ.NewBool(λ.Calm(ϒmedia_info, "get", λ.StrLiteral("_isLive")) == λ.True),
						"formats":   ϒformats,
						"subtitles": ϒsubtitles,
					})
				})
			ARDMediathekBaseIE__extract_formats = λ.NewFunction("_extract_formats",
				[]λ.Param{
					{Name: "self"},
					{Name: "media_info"},
					{Name: "video_id"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒext         λ.Object
						ϒf           λ.Object
						ϒformats     λ.Object
						ϒm           λ.Object
						ϒmedia       λ.Object
						ϒmedia_array λ.Object
						ϒmedia_info  = λargs[1]
						ϒnum         λ.Object
						ϒquality     λ.Object
						ϒself        = λargs[0]
						ϒserver      λ.Object
						ϒstream      λ.Object
						ϒstream_url  λ.Object
						ϒstream_urls λ.Object
						ϒtype_       λ.Object
						ϒvideo_id    = λargs[2]
						τmp0         λ.Object
						τmp1         λ.Object
						τmp2         λ.Object
						τmp3         λ.Object
						τmp4         λ.Object
						τmp5         λ.Object
					)
					ϒtype_ = λ.Calm(ϒmedia_info, "get", λ.StrLiteral("_type"))
					ϒmedia_array = λ.Calm(ϒmedia_info, "get", λ.StrLiteral("_mediaArray"), λ.NewList())
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.EnumerateIteratorType, ϒmedia_array))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						τmp2 = τmp1
						ϒnum = λ.GetItem(τmp2, λ.IntLiteral(0))
						ϒmedia = λ.GetItem(τmp2, λ.IntLiteral(1))
						τmp2 = λ.Cal(λ.BuiltinIter, λ.Calm(ϒmedia, "get", λ.StrLiteral("_mediaStreamArray"), λ.NewList()))
						for {
							if τmp3 = λ.NextDefault(τmp2, λ.AfterLast); τmp3 == λ.AfterLast {
								break
							}
							ϒstream = τmp3
							ϒstream_urls = λ.Calm(ϒstream, "get", λ.StrLiteral("_stream"))
							if !λ.IsTrue(ϒstream_urls) {
								continue
							}
							if !λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒstream_urls, λ.ListType)) {
								ϒstream_urls = λ.NewList(ϒstream_urls)
							}
							ϒquality = λ.Calm(ϒstream, "get", λ.StrLiteral("_quality"))
							ϒserver = λ.Calm(ϒstream, "get", λ.StrLiteral("_server"))
							τmp4 = λ.Cal(λ.BuiltinIter, ϒstream_urls)
							for {
								if τmp5 = λ.NextDefault(τmp4, λ.AfterLast); τmp5 == λ.AfterLast {
									break
								}
								ϒstream_url = τmp5
								if !λ.IsTrue(λ.Cal(ϒurl_or_none, ϒstream_url)) {
									continue
								}
								ϒext = λ.Cal(ϒdetermine_ext, ϒstream_url)
								if λ.IsTrue(func() λ.Object {
									if λv := λ.Ne(ϒquality, λ.StrLiteral("auto")); !λ.IsTrue(λv) {
										return λv
									} else {
										return λ.NewBool(λ.Contains(λ.NewTuple(
											λ.StrLiteral("f4m"),
											λ.StrLiteral("m3u8"),
										), ϒext))
									}
								}()) {
									continue
								}
								if λ.IsTrue(λ.Eq(ϒext, λ.StrLiteral("f4m"))) {
									λ.Calm(ϒformats, "extend", λ.Call(λ.GetAttr(ϒself, "_extract_f4m_formats", nil), λ.NewArgs(
										λ.Cal(ϒupdate_url_query, ϒstream_url, λ.DictLiteral(map[string]string{
											"hdcore": "3.1.1",
											"plugin": "aasp-3.1.1.69.124",
										})),
										ϒvideo_id,
									), λ.KWArgs{
										{Name: "f4m_id", Value: λ.StrLiteral("hds")},
										{Name: "fatal", Value: λ.False},
									}))
								} else {
									if λ.IsTrue(λ.Eq(ϒext, λ.StrLiteral("m3u8"))) {
										λ.Calm(ϒformats, "extend", λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
											ϒstream_url,
											ϒvideo_id,
											λ.StrLiteral("mp4"),
											λ.StrLiteral("m3u8_native"),
										), λ.KWArgs{
											{Name: "m3u8_id", Value: λ.StrLiteral("hls")},
											{Name: "fatal", Value: λ.False},
										}))
									} else {
										if λ.IsTrue(func() λ.Object {
											if λv := ϒserver; !λ.IsTrue(λv) {
												return λv
											} else {
												return λ.Calm(ϒserver, "startswith", λ.StrLiteral("rtmp"))
											}
										}()) {
											ϒf = λ.DictLiteral(map[string]λ.Object{
												"url":       ϒserver,
												"play_path": ϒstream_url,
												"format_id": λ.Mod(λ.StrLiteral("a%s-rtmp-%s"), λ.NewTuple(
													ϒnum,
													ϒquality,
												)),
											})
										} else {
											ϒf = λ.DictLiteral(map[string]λ.Object{
												"url": ϒstream_url,
												"format_id": λ.Mod(λ.StrLiteral("a%s-%s-%s"), λ.NewTuple(
													ϒnum,
													ϒext,
													ϒquality,
												)),
											})
										}
										ϒm = λ.Cal(Ωre.ϒsearch, λ.StrLiteral("_(?P<width>\\d+)x(?P<height>\\d+)\\.mp4$"), ϒstream_url)
										if λ.IsTrue(ϒm) {
											λ.Calm(ϒf, "update", λ.DictLiteral(map[string]λ.Object{
												"width":  λ.Cal(λ.IntType, λ.Calm(ϒm, "group", λ.StrLiteral("width"))),
												"height": λ.Cal(λ.IntType, λ.Calm(ϒm, "group", λ.StrLiteral("height"))),
											}))
										}
										if λ.IsTrue(λ.Eq(ϒtype_, λ.StrLiteral("audio"))) {
											λ.SetItem(ϒf, λ.StrLiteral("vcodec"), λ.StrLiteral("none"))
										}
										λ.Calm(ϒformats, "append", ϒf)
									}
								}
							}
						}
					}
					return ϒformats
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_GEO_COUNTRIES":      ARDMediathekBaseIE__GEO_COUNTRIES,
				"_extract_formats":    ARDMediathekBaseIE__extract_formats,
				"_extract_media_info": ARDMediathekBaseIE__extract_media_info,
				"_parse_media_info":   ARDMediathekBaseIE__parse_media_info,
			})
		}())
		ARDMediathekIE = λ.Cal(λ.TypeType, λ.StrLiteral("ARDMediathekIE"), λ.NewTuple(ARDMediathekBaseIE), func() λ.Dict {
			var (
				ARDMediathekIE_IE_NAME       λ.Object
				ARDMediathekIE__VALID_URL    λ.Object
				ARDMediathekIE__real_extract λ.Object
				ARDMediathekIE_suitable      λ.Object
			)
			ARDMediathekIE_IE_NAME = λ.StrLiteral("ARD:mediathek")
			ARDMediathekIE__VALID_URL = λ.StrLiteral("^https?://(?:(?:(?:www|classic)\\.)?ardmediathek\\.de|mediathek\\.(?:daserste|rbb-online)\\.de|one\\.ard\\.de)/(?:.*/)(?P<video_id>[0-9]+|[^0-9][^/\\?]+)[^/\\?]*(?:\\?.*)?")
			ARDMediathekIE_suitable = λ.NewFunction("suitable",
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
						if λ.IsTrue(λ.Calm(ARDBetaMediathekIE, "suitable", ϒurl)) {
							return λ.False
						} else {
							return λ.Calm(λ.Cal(λ.SuperType, ARDMediathekIE, ϒcls), "suitable", ϒurl)
						}
					}()
				})
			ARDMediathekIE_suitable = λ.Cal(λ.ClassMethodType, ARDMediathekIE_suitable)
			ARDMediathekIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ERRORS         λ.Object
						QUALITIES      λ.Object
						ϒdescription   λ.Object
						ϒdoc           λ.Object
						ϒdocument_id   λ.Object
						ϒfid           λ.Object
						ϒfid_m         λ.Object
						ϒformats       λ.Object
						ϒfurl          λ.Object
						ϒinfo          λ.Object
						ϒm             λ.Object
						ϒmedia_streams λ.Object
						ϒmessage       λ.Object
						ϒnumid         λ.Object
						ϒpattern       λ.Object
						ϒself          = λargs[0]
						ϒthumbnail     λ.Object
						ϒtitle         λ.Object
						ϒurl           = λargs[1]
						ϒvideo_id      λ.Object
						ϒwebpage       λ.Object
						τmp0           λ.Object
						τmp1           λ.Object
						τmp2           λ.Object
					)
					ϒm = λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl)
					ϒdocument_id = λ.None
					ϒnumid = λ.Cal(Ωre.ϒsearch, λ.StrLiteral("documentId=([0-9]+)"), ϒurl)
					if λ.IsTrue(ϒnumid) {
						τmp0 = λ.Calm(ϒnumid, "group", λ.IntLiteral(1))
						ϒdocument_id = τmp0
						ϒvideo_id = τmp0
					} else {
						ϒvideo_id = λ.Calm(ϒm, "group", λ.StrLiteral("video_id"))
					}
					ϒwebpage = λ.Calm(ϒself, "_download_webpage", ϒurl, ϒvideo_id)
					ERRORS = λ.NewTuple(
						λ.NewTuple(
							λ.StrLiteral(">Leider liegt eine Störung vor."),
							λ.StrLiteral("Video %s is unavailable"),
						),
						λ.NewTuple(
							λ.StrLiteral(">Der gewünschte Beitrag ist nicht mehr verfügbar.<"),
							λ.StrLiteral("Video %s is no longer available"),
						),
					)
					τmp0 = λ.Cal(λ.BuiltinIter, ERRORS)
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						τmp2 = τmp1
						ϒpattern = λ.GetItem(τmp2, λ.IntLiteral(0))
						ϒmessage = λ.GetItem(τmp2, λ.IntLiteral(1))
						if λ.Contains(ϒwebpage, ϒpattern) {
							panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.Mod(ϒmessage, ϒvideo_id)), λ.KWArgs{
								{Name: "expected", Value: λ.True},
							})))
						}
					}
					if λ.IsTrue(λ.Cal(Ωre.ϒsearch, λ.StrLiteral("[\\?&]rss($|[=&])"), ϒurl)) {
						ϒdoc = λ.Cal(ϒcompat_etree_fromstring, λ.Calm(ϒwebpage, "encode", λ.StrLiteral("utf-8")))
						if λ.IsTrue(λ.Eq(λ.GetAttr(ϒdoc, "tag", nil), λ.StrLiteral("rss"))) {
							return λ.Calm(λ.Cal(λ.None), "_extract_rss", ϒurl, ϒvideo_id, ϒdoc)
						}
					}
					ϒtitle = λ.Calm(ϒself, "_html_search_regex", λ.NewList(
						λ.StrLiteral("<h1(?:\\s+class=\"boxTopHeadline\")?>(.*?)</h1>"),
						λ.StrLiteral("<meta name=\"dcterms\\.title\" content=\"(.*?)\"/>"),
						λ.StrLiteral("<h4 class=\"headline\">(.*?)</h4>"),
						λ.StrLiteral("<title[^>]*>(.*?)</title>"),
					), ϒwebpage, λ.StrLiteral("title"))
					ϒdescription = λ.Call(λ.GetAttr(ϒself, "_html_search_meta", nil), λ.NewArgs(
						λ.StrLiteral("dcterms.abstract"),
						ϒwebpage,
						λ.StrLiteral("description"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					if ϒdescription == λ.None {
						ϒdescription = λ.Call(λ.GetAttr(ϒself, "_html_search_meta", nil), λ.NewArgs(
							λ.StrLiteral("description"),
							ϒwebpage,
							λ.StrLiteral("meta description"),
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
						})
					}
					if ϒdescription == λ.None {
						ϒdescription = λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
							λ.StrLiteral("<p\\s+class=\"teasertext\">(.+?)</p>"),
							ϒwebpage,
							λ.StrLiteral("teaser text"),
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
						})
					}
					ϒthumbnail = λ.Call(λ.GetAttr(ϒself, "_og_search_thumbnail", nil), λ.NewArgs(ϒwebpage), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					ϒmedia_streams = λ.Cal(Ωre.ϒfindall, λ.StrLiteral("(?x)\n            mediaCollection\\.addMediaStream\\([0-9]+,\\s*[0-9]+,\\s*\"[^\"]*\",\\s*\n            \"([^\"]+)\""), ϒwebpage)
					if λ.IsTrue(ϒmedia_streams) {
						QUALITIES = λ.Cal(ϒqualities, λ.NewList(
							λ.StrLiteral("lo"),
							λ.StrLiteral("hi"),
							λ.StrLiteral("hq"),
						))
						ϒformats = λ.NewList()
						τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.SetType, ϒmedia_streams))
						for {
							if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
								break
							}
							ϒfurl = τmp1
							if λ.IsTrue(λ.Calm(ϒfurl, "endswith", λ.StrLiteral(".f4m"))) {
								ϒfid = λ.StrLiteral("f4m")
							} else {
								ϒfid_m = λ.Cal(Ωre.ϒmatch, λ.StrLiteral(".*\\.([^.]+)\\.[^.]+$"), ϒfurl)
								ϒfid = func() λ.Object {
									if λ.IsTrue(ϒfid_m) {
										return λ.Calm(ϒfid_m, "group", λ.IntLiteral(1))
									} else {
										return λ.None
									}
								}()
							}
							λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
								"quality":   λ.Cal(QUALITIES, ϒfid),
								"format_id": ϒfid,
								"url":       ϒfurl,
							}))
						}
						λ.Calm(ϒself, "_sort_formats", ϒformats)
						ϒinfo = λ.DictLiteral(map[string]λ.Object{
							"formats": ϒformats,
						})
					} else {
						if !λ.IsTrue(ϒdocument_id) {
							ϒvideo_id = λ.Calm(ϒself, "_search_regex", λ.StrLiteral("/play/(?:config|media)/(\\d+)"), ϒwebpage, λ.StrLiteral("media id"))
						}
						ϒinfo = λ.Calm(ϒself, "_extract_media_info", λ.Mod(λ.StrLiteral("http://www.ardmediathek.de/play/media/%s"), ϒvideo_id), ϒwebpage, ϒvideo_id)
					}
					λ.Calm(ϒinfo, "update", λ.DictLiteral(map[string]λ.Object{
						"id": ϒvideo_id,
						"title": func() λ.Object {
							if λ.IsTrue(λ.Calm(ϒinfo, "get", λ.StrLiteral("is_live"))) {
								return λ.Calm(ϒself, "_live_title", ϒtitle)
							} else {
								return ϒtitle
							}
						}(),
						"description": ϒdescription,
						"thumbnail":   ϒthumbnail,
					}))
					return ϒinfo
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"IE_NAME":       ARDMediathekIE_IE_NAME,
				"_VALID_URL":    ARDMediathekIE__VALID_URL,
				"_real_extract": ARDMediathekIE__real_extract,
				"suitable":      ARDMediathekIE_suitable,
			})
		}())
		ARDIE = λ.Cal(λ.TypeType, λ.StrLiteral("ARDIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				ARDIE__VALID_URL λ.Object
			)
			ARDIE__VALID_URL = λ.StrLiteral("(?P<mainurl>https?://(www\\.)?daserste\\.de/[^?#]+/videos(?:extern)?/(?P<display_id>[^/?#]+)-(?P<id>[0-9]+))\\.html")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": ARDIE__VALID_URL,
			})
		}())
		ARDBetaMediathekIE = λ.Cal(λ.TypeType, λ.StrLiteral("ARDBetaMediathekIE"), λ.NewTuple(ARDMediathekBaseIE), func() λ.Dict {
			var (
				ARDBetaMediathekIE__VALID_URL λ.Object
			)
			ARDBetaMediathekIE__VALID_URL = λ.StrLiteral("https://(?:(?:beta|www)\\.)?ardmediathek\\.de/(?P<client>[^/]+)/(?:player|live|video)/(?P<display_id>(?:[^/]+/)*)(?P<video_id>[a-zA-Z0-9]+)")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": ARDBetaMediathekIE__VALID_URL,
			})
		}())
	})
}
