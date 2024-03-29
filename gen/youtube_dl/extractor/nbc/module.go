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
 * nbc/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/nbc.py
 */

package nbc

import (
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωadobepass "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/adobepass"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωtheplatform "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/theplatform"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	AdobePassIE                  λ.Object
	InfoExtractor                λ.Object
	NBCIE                        λ.Object
	NBCNewsIE                    λ.Object
	NBCOlympicsIE                λ.Object
	NBCOlympicsStreamIE          λ.Object
	NBCSportsIE                  λ.Object
	NBCSportsStreamIE            λ.Object
	NBCSportsVPlayerIE           λ.Object
	ThePlatformIE                λ.Object
	ϒcompat_urllib_parse_unquote λ.Object
	ϒint_or_none                 λ.Object
	ϒparse_duration              λ.Object
	ϒsmuggle_url                 λ.Object
	ϒtry_get                     λ.Object
	ϒunified_timestamp           λ.Object
	ϒupdate_url_query            λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ThePlatformIE = Ωtheplatform.ThePlatformIE
		AdobePassIE = Ωadobepass.AdobePassIE
		ϒcompat_urllib_parse_unquote = Ωcompat.ϒcompat_urllib_parse_unquote
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒparse_duration = Ωutils.ϒparse_duration
		ϒsmuggle_url = Ωutils.ϒsmuggle_url
		ϒtry_get = Ωutils.ϒtry_get
		ϒunified_timestamp = Ωutils.ϒunified_timestamp
		ϒupdate_url_query = Ωutils.ϒupdate_url_query
		NBCIE = λ.Cal(λ.TypeType, λ.StrLiteral("NBCIE"), λ.NewTuple(AdobePassIE), func() λ.Dict {
			var (
				NBCIE__VALID_URL λ.Object
			)
			NBCIE__VALID_URL = λ.StrLiteral("https?(?P<permalink>://(?:www\\.)?nbc\\.com/(?:classic-tv/)?[^/]+/video/[^/]+/(?P<id>n?\\d+))")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": NBCIE__VALID_URL,
			})
		}())
		NBCSportsVPlayerIE = λ.Cal(λ.TypeType, λ.StrLiteral("NBCSportsVPlayerIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				NBCSportsVPlayerIE__VALID_URL      λ.Object
				NBCSportsVPlayerIE__VALID_URL_BASE λ.Object
			)
			NBCSportsVPlayerIE__VALID_URL_BASE = λ.StrLiteral("https?://(?:vplayer\\.nbcsports\\.com|(?:www\\.)?nbcsports\\.com/vplayer)/")
			NBCSportsVPlayerIE__VALID_URL = λ.Add(NBCSportsVPlayerIE__VALID_URL_BASE, λ.StrLiteral("(?:[^/]+/)+(?P<id>[0-9a-zA-Z_]+)"))
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL":      NBCSportsVPlayerIE__VALID_URL,
				"_VALID_URL_BASE": NBCSportsVPlayerIE__VALID_URL_BASE,
			})
		}())
		NBCSportsIE = λ.Cal(λ.TypeType, λ.StrLiteral("NBCSportsIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				NBCSportsIE__VALID_URL λ.Object
			)
			NBCSportsIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?nbcsports\\.com//?(?!vplayer/)(?:[^/]+/)+(?P<id>[0-9a-z-]+)")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": NBCSportsIE__VALID_URL,
			})
		}())
		NBCSportsStreamIE = λ.Cal(λ.TypeType, λ.StrLiteral("NBCSportsStreamIE"), λ.NewTuple(AdobePassIE), func() λ.Dict {
			var (
				NBCSportsStreamIE__VALID_URL λ.Object
			)
			NBCSportsStreamIE__VALID_URL = λ.StrLiteral("https?://stream\\.nbcsports\\.com/.+?\\bpid=(?P<id>\\d+)")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": NBCSportsStreamIE__VALID_URL,
			})
		}())
		NBCNewsIE = λ.Cal(λ.TypeType, λ.StrLiteral("NBCNewsIE"), λ.NewTuple(ThePlatformIE), func() λ.Dict {
			var (
				NBCNewsIE__VALID_URL    λ.Object
				NBCNewsIE__real_extract λ.Object
			)
			NBCNewsIE__VALID_URL = λ.StrLiteral("(?x)https?://(?:www\\.)?(?:nbcnews|today|msnbc)\\.com/([^/]+/)*(?:.*-)?(?P<id>[^/?]+)")
			NBCNewsIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒcc_url            λ.Object
						ϒclosed_captioning λ.Object
						ϒdata              λ.Object
						ϒformat_id         λ.Object
						ϒformats           λ.Object
						ϒpublic_url        λ.Object
						ϒself              = λargs[0]
						ϒsubtitles         λ.Object
						ϒtbr               λ.Object
						ϒtitle             λ.Object
						ϒurl               = λargs[1]
						ϒva                λ.Object
						ϒvideo_data        λ.Object
						ϒvideo_id          λ.Object
						ϒwebpage           λ.Object
						τmp0               λ.Object
						τmp1               λ.Object
						τmp2               λ.Object
					)
					ϒvideo_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒwebpage = λ.Calm(ϒself, "_download_webpage", ϒurl, ϒvideo_id)
					ϒdata = λ.GetItem(λ.GetItem(λ.Calm(ϒself, "_parse_json", λ.Calm(ϒself, "_search_regex", λ.StrLiteral("<script[^>]+id=\"__NEXT_DATA__\"[^>]*>({.+?})</script>"), ϒwebpage, λ.StrLiteral("bootstrap json")), ϒvideo_id), λ.StrLiteral("props")), λ.StrLiteral("initialState"))
					ϒvideo_data = λ.Cal(ϒtry_get, ϒdata, λ.NewFunction("<lambda>",
						[]λ.Param{
							{Name: "x"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒx = λargs[0]
							)
							return λ.GetItem(λ.GetItem(ϒx, λ.StrLiteral("video")), λ.StrLiteral("current"))
						}), λ.DictType)
					if !λ.IsTrue(ϒvideo_data) {
						ϒvideo_data = λ.GetItem(λ.GetItem(λ.GetItem(λ.GetItem(λ.GetItem(ϒdata, λ.StrLiteral("article")), λ.StrLiteral("content")), λ.IntLiteral(0)), λ.StrLiteral("primaryMedia")), λ.StrLiteral("video"))
					}
					ϒtitle = λ.GetItem(λ.GetItem(ϒvideo_data, λ.StrLiteral("headline")), λ.StrLiteral("primary"))
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(ϒvideo_data, "get", λ.StrLiteral("videoAssets"), λ.NewList()))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒva = τmp1
						ϒpublic_url = λ.Calm(ϒva, "get", λ.StrLiteral("publicUrl"))
						if !λ.IsTrue(ϒpublic_url) {
							continue
						}
						if λ.Contains(ϒpublic_url, λ.StrLiteral("://link.theplatform.com/")) {
							ϒpublic_url = λ.Cal(ϒupdate_url_query, ϒpublic_url, λ.DictLiteral(map[string]string{
								"format": "redirect",
							}))
						}
						ϒformat_id = λ.Calm(ϒva, "get", λ.StrLiteral("format"))
						if λ.IsTrue(λ.Eq(ϒformat_id, λ.StrLiteral("M3U"))) {
							λ.Calm(ϒformats, "extend", λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
								ϒpublic_url,
								ϒvideo_id,
								λ.StrLiteral("mp4"),
								λ.StrLiteral("m3u8_native"),
							), λ.KWArgs{
								{Name: "m3u8_id", Value: ϒformat_id},
								{Name: "fatal", Value: λ.False},
							}))
							continue
						}
						ϒtbr = λ.Cal(ϒint_or_none, λ.Calm(ϒva, "get", λ.StrLiteral("bitrate")), λ.IntLiteral(1000))
						if λ.IsTrue(ϒtbr) {
							τmp2 = λ.IAdd(ϒformat_id, λ.Mod(λ.StrLiteral("-%d"), ϒtbr))
							ϒformat_id = τmp2
						}
						λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
							"format_id": ϒformat_id,
							"url":       ϒpublic_url,
							"width":     λ.Cal(ϒint_or_none, λ.Calm(ϒva, "get", λ.StrLiteral("width"))),
							"height":    λ.Cal(ϒint_or_none, λ.Calm(ϒva, "get", λ.StrLiteral("height"))),
							"tbr":       ϒtbr,
							"ext":       λ.StrLiteral("mp4"),
						}))
					}
					λ.Calm(ϒself, "_sort_formats", ϒformats)
					ϒsubtitles = λ.DictLiteral(map[λ.Object]λ.Object{})
					ϒclosed_captioning = λ.Calm(ϒvideo_data, "get", λ.StrLiteral("closedCaptioning"))
					if λ.IsTrue(ϒclosed_captioning) {
						τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(ϒclosed_captioning, "values"))
						for {
							if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
								break
							}
							ϒcc_url = τmp1
							if !λ.IsTrue(ϒcc_url) {
								continue
							}
							λ.Calm(λ.Calm(ϒsubtitles, "setdefault", λ.StrLiteral("en"), λ.NewList()), "append", λ.DictLiteral(map[string]λ.Object{
								"url": ϒcc_url,
							}))
						}
					}
					return λ.DictLiteral(map[string]λ.Object{
						"id":    ϒvideo_id,
						"title": ϒtitle,
						"description": λ.Cal(ϒtry_get, ϒvideo_data, λ.NewFunction("<lambda>",
							[]λ.Param{
								{Name: "x"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒx = λargs[0]
								)
								return λ.GetItem(λ.GetItem(ϒx, λ.StrLiteral("description")), λ.StrLiteral("primary"))
							})),
						"thumbnail": λ.Cal(ϒtry_get, ϒvideo_data, λ.NewFunction("<lambda>",
							[]λ.Param{
								{Name: "x"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒx = λargs[0]
								)
								return λ.GetItem(λ.GetItem(λ.GetItem(ϒx, λ.StrLiteral("primaryImage")), λ.StrLiteral("url")), λ.StrLiteral("primary"))
							})),
						"duration":  λ.Cal(ϒparse_duration, λ.Calm(ϒvideo_data, "get", λ.StrLiteral("duration"))),
						"timestamp": λ.Cal(ϒunified_timestamp, λ.Calm(ϒvideo_data, "get", λ.StrLiteral("datePublished"))),
						"formats":   ϒformats,
						"subtitles": ϒsubtitles,
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL":    NBCNewsIE__VALID_URL,
				"_real_extract": NBCNewsIE__real_extract,
			})
		}())
		NBCOlympicsIE = λ.Cal(λ.TypeType, λ.StrLiteral("NBCOlympicsIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				NBCOlympicsIE__VALID_URL λ.Object
			)
			NBCOlympicsIE__VALID_URL = λ.StrLiteral("https?://www\\.nbcolympics\\.com/video/(?P<id>[a-z-]+)")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": NBCOlympicsIE__VALID_URL,
			})
		}())
		NBCOlympicsStreamIE = λ.Cal(λ.TypeType, λ.StrLiteral("NBCOlympicsStreamIE"), λ.NewTuple(AdobePassIE), func() λ.Dict {
			var (
				NBCOlympicsStreamIE__VALID_URL λ.Object
			)
			NBCOlympicsStreamIE__VALID_URL = λ.StrLiteral("https?://stream\\.nbcolympics\\.com/(?P<id>[0-9a-z-]+)")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": NBCOlympicsStreamIE__VALID_URL,
			})
		}())
	})
}
