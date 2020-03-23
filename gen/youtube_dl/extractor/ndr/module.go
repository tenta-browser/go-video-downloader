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
 * ndr/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/ndr.py
 */

package ndr

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor  λ.Object
	NDRBaseIE      λ.Object
	NDREmbedBaseIE λ.Object
	NDREmbedIE     λ.Object
	NDRIE          λ.Object
	NJoyEmbedIE    λ.Object
	NJoyIE         λ.Object
	ϒdetermine_ext λ.Object
	ϒint_or_none   λ.Object
	ϒmerge_dicts   λ.Object
	ϒparse_iso8601 λ.Object
	ϒqualities     λ.Object
	ϒtry_get       λ.Object
	ϒurljoin       λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒdetermine_ext = Ωutils.ϒdetermine_ext
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒmerge_dicts = Ωutils.ϒmerge_dicts
		ϒparse_iso8601 = Ωutils.ϒparse_iso8601
		ϒqualities = Ωutils.ϒqualities
		ϒtry_get = Ωutils.ϒtry_get
		ϒurljoin = Ωutils.ϒurljoin
		NDRBaseIE = λ.Cal(λ.TypeType, λ.StrLiteral("NDRBaseIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				NDRBaseIE__real_extract λ.Object
			)
			NDRBaseIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdisplay_id λ.Object
						ϒmobj       λ.Object
						ϒself       = λargs[0]
						ϒurl        = λargs[1]
						ϒwebpage    λ.Object
					)
					ϒmobj = λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl)
					ϒdisplay_id = λ.Cal(λ.BuiltinNext, λ.Cal(λ.NewFunction("<generator>",
						nil,
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							return λ.NewGenerator(func(λgy λ.Yielder) λ.Object {
								var (
									ϒgroup λ.Object
									τmp0   λ.Object
									τmp1   λ.Object
								)
								τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(ϒmobj, "groups"))
								for {
									if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
										break
									}
									ϒgroup = τmp1
									if λ.IsTrue(ϒgroup) {
										λgy.Yield(ϒgroup)
									}
								}
								return λ.None
							})
						})))
					ϒwebpage = λ.Calm(ϒself, "_download_webpage", ϒurl, ϒdisplay_id)
					return λ.Calm(ϒself, "_extract_embed", ϒwebpage, ϒdisplay_id)
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_real_extract": NDRBaseIE__real_extract,
			})
		}())
		NDRIE = λ.Cal(λ.TypeType, λ.StrLiteral("NDRIE"), λ.NewTuple(NDRBaseIE), func() λ.Dict {
			var (
				NDRIE_IE_NAME        λ.Object
				NDRIE__VALID_URL     λ.Object
				NDRIE__extract_embed λ.Object
			)
			NDRIE_IE_NAME = λ.StrLiteral("ndr")
			NDRIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?ndr\\.de/(?:[^/]+/)*(?P<id>[^/?#]+),[\\da-z]+\\.html")
			NDRIE__extract_embed = λ.NewFunction("_extract_embed",
				[]λ.Param{
					{Name: "self"},
					{Name: "webpage"},
					{Name: "display_id"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdescription λ.Object
						ϒdisplay_id  = λargs[2]
						ϒembed_url   λ.Object
						ϒinfo        λ.Object
						ϒself        = λargs[0]
						ϒtimestamp   λ.Object
						ϒwebpage     = λargs[1]
					)
					ϒembed_url = func() λ.Object {
						if λv := λ.Call(λ.GetAttr(ϒself, "_html_search_meta", nil), λ.NewArgs(
							λ.StrLiteral("embedURL"),
							ϒwebpage,
							λ.StrLiteral("embed URL"),
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
						}); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
								λ.StrLiteral("\\bembedUrl[\"\\']\\s*:\\s*([\"\\'])(?P<url>(?:(?!\\1).)+)\\1"),
								ϒwebpage,
								λ.StrLiteral("embed URL"),
							), λ.KWArgs{
								{Name: "group", Value: λ.StrLiteral("url")},
							})
						}
					}()
					ϒdescription = func() λ.Object {
						if λv := λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
							λ.StrLiteral("<p[^>]+itemprop=\"description\">([^<]+)</p>"),
							ϒwebpage,
							λ.StrLiteral("description"),
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
						}); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Calm(ϒself, "_og_search_description", ϒwebpage)
						}
					}()
					ϒtimestamp = λ.Cal(ϒparse_iso8601, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("<span[^>]+itemprop=\"(?:datePublished|uploadDate)\"[^>]+content=\"([^\"]+)\""),
						ϒwebpage,
						λ.StrLiteral("upload date"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					}))
					ϒinfo = λ.Call(λ.GetAttr(ϒself, "_search_json_ld", nil), λ.NewArgs(
						ϒwebpage,
						ϒdisplay_id,
					), λ.KWArgs{
						{Name: "default", Value: λ.DictLiteral(map[λ.Object]λ.Object{})},
					})
					return λ.Cal(ϒmerge_dicts, λ.DictLiteral(map[string]λ.Object{
						"_type":       λ.StrLiteral("url_transparent"),
						"url":         ϒembed_url,
						"display_id":  ϒdisplay_id,
						"description": ϒdescription,
						"timestamp":   ϒtimestamp,
					}), ϒinfo)
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"IE_NAME":        NDRIE_IE_NAME,
				"_VALID_URL":     NDRIE__VALID_URL,
				"_extract_embed": NDRIE__extract_embed,
			})
		}())
		NJoyIE = λ.Cal(λ.TypeType, λ.StrLiteral("NJoyIE"), λ.NewTuple(NDRBaseIE), func() λ.Dict {
			var (
				NJoyIE__VALID_URL λ.Object
			)
			NJoyIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?n-joy\\.de/(?:[^/]+/)*(?:(?P<display_id>[^/?#]+),)?(?P<id>[\\da-z]+)\\.html")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": NJoyIE__VALID_URL,
			})
		}())
		NDREmbedBaseIE = λ.Cal(λ.TypeType, λ.StrLiteral("NDREmbedBaseIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				NDREmbedBaseIE__VALID_URL    λ.Object
				NDREmbedBaseIE__real_extract λ.Object
			)
			NDREmbedBaseIE__VALID_URL = λ.StrLiteral("(?:ndr:(?P<id_s>[\\da-z]+)|https?://www\\.ndr\\.de/(?P<id>[\\da-z]+)-ppjson\\.json)")
			NDREmbedBaseIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒconfig        λ.Object
						ϒduration      λ.Object
						ϒext           λ.Object
						ϒf             λ.Object
						ϒff            λ.Object
						ϒformat_id     λ.Object
						ϒformats       λ.Object
						ϒlive          λ.Object
						ϒmobj          λ.Object
						ϒplaylist      λ.Object
						ϒposter        λ.Object
						ϒppjson        λ.Object
						ϒquality       λ.Object
						ϒquality_key   λ.Object
						ϒself          = λargs[0]
						ϒsrc           λ.Object
						ϒthumbnail     λ.Object
						ϒthumbnail_id  λ.Object
						ϒthumbnail_url λ.Object
						ϒthumbnails    λ.Object
						ϒtitle         λ.Object
						ϒtype_         λ.Object
						ϒupload_date   λ.Object
						ϒuploader      λ.Object
						ϒurl           = λargs[1]
						ϒvideo_id      λ.Object
						τmp0           λ.Object
						τmp1           λ.Object
						τmp2           λ.Object
					)
					ϒmobj = λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl)
					ϒvideo_id = func() λ.Object {
						if λv := λ.Calm(ϒmobj, "group", λ.StrLiteral("id")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Calm(ϒmobj, "group", λ.StrLiteral("id_s"))
						}
					}()
					ϒppjson = λ.Calm(ϒself, "_download_json", λ.Mod(λ.StrLiteral("http://www.ndr.de/%s-ppjson.json"), ϒvideo_id), ϒvideo_id)
					ϒplaylist = λ.GetItem(ϒppjson, λ.StrLiteral("playlist"))
					ϒformats = λ.NewList()
					ϒquality_key = λ.Cal(ϒqualities, λ.NewTuple(
						λ.StrLiteral("xs"),
						λ.StrLiteral("s"),
						λ.StrLiteral("m"),
						λ.StrLiteral("l"),
						λ.StrLiteral("xl"),
					))
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(ϒplaylist, "items"))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						τmp2 = τmp1
						ϒformat_id = λ.GetItem(τmp2, λ.IntLiteral(0))
						ϒf = λ.GetItem(τmp2, λ.IntLiteral(1))
						ϒsrc = λ.Calm(ϒf, "get", λ.StrLiteral("src"))
						if !λ.IsTrue(ϒsrc) {
							continue
						}
						ϒext = λ.Cal(ϒdetermine_ext, ϒsrc, λ.None)
						if λ.IsTrue(λ.Eq(ϒext, λ.StrLiteral("f4m"))) {
							λ.Calm(ϒformats, "extend", λ.Call(λ.GetAttr(ϒself, "_extract_f4m_formats", nil), λ.NewArgs(
								λ.Add(ϒsrc, λ.StrLiteral("?hdcore=3.7.0&plugin=aasp-3.7.0.39.44")),
								ϒvideo_id,
							), λ.KWArgs{
								{Name: "f4m_id", Value: λ.StrLiteral("hds")},
								{Name: "fatal", Value: λ.False},
							}))
						} else {
							if λ.IsTrue(λ.Eq(ϒext, λ.StrLiteral("m3u8"))) {
								λ.Calm(ϒformats, "extend", λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
									ϒsrc,
									ϒvideo_id,
									λ.StrLiteral("mp4"),
								), λ.KWArgs{
									{Name: "m3u8_id", Value: λ.StrLiteral("hls")},
									{Name: "entry_protocol", Value: λ.StrLiteral("m3u8_native")},
									{Name: "fatal", Value: λ.False},
								}))
							} else {
								ϒquality = λ.Calm(ϒf, "get", λ.StrLiteral("quality"))
								ϒff = λ.DictLiteral(map[string]λ.Object{
									"url": ϒsrc,
									"format_id": func() λ.Object {
										if λv := ϒquality; λ.IsTrue(λv) {
											return λv
										} else {
											return ϒformat_id
										}
									}(),
									"quality": λ.Cal(ϒquality_key, ϒquality),
								})
								ϒtype_ = λ.Calm(ϒf, "get", λ.StrLiteral("type"))
								if λ.IsTrue(func() λ.Object {
									if λv := ϒtype_; !λ.IsTrue(λv) {
										return λv
									} else {
										return λ.Eq(λ.GetItem(λ.Calm(ϒtype_, "split", λ.StrLiteral("/")), λ.IntLiteral(0)), λ.StrLiteral("audio"))
									}
								}()) {
									λ.SetItem(ϒff, λ.StrLiteral("vcodec"), λ.StrLiteral("none"))
									λ.SetItem(ϒff, λ.StrLiteral("ext"), func() λ.Object {
										if λv := ϒext; λ.IsTrue(λv) {
											return λv
										} else {
											return λ.StrLiteral("mp3")
										}
									}())
								}
								λ.Calm(ϒformats, "append", ϒff)
							}
						}
					}
					λ.Calm(ϒself, "_sort_formats", ϒformats)
					ϒconfig = λ.GetItem(ϒplaylist, λ.StrLiteral("config"))
					ϒlive = λ.NewBool(λ.Contains(λ.NewList(
						λ.StrLiteral("httpVideoLive"),
						λ.StrLiteral("httpAudioLive"),
					), λ.Calm(λ.Calm(ϒplaylist, "get", λ.StrLiteral("config"), λ.DictLiteral(map[λ.Object]λ.Object{})), "get", λ.StrLiteral("streamType"))))
					ϒtitle = λ.GetItem(ϒconfig, λ.StrLiteral("title"))
					if λ.IsTrue(ϒlive) {
						ϒtitle = λ.Calm(ϒself, "_live_title", ϒtitle)
					}
					ϒuploader = λ.Calm(λ.Calm(ϒppjson, "get", λ.StrLiteral("config"), λ.DictLiteral(map[λ.Object]λ.Object{})), "get", λ.StrLiteral("branding"))
					ϒupload_date = λ.Calm(λ.Calm(ϒppjson, "get", λ.StrLiteral("config"), λ.DictLiteral(map[λ.Object]λ.Object{})), "get", λ.StrLiteral("publicationDate"))
					ϒduration = λ.Cal(ϒint_or_none, λ.Calm(ϒconfig, "get", λ.StrLiteral("duration")))
					ϒthumbnails = λ.NewList()
					ϒposter = func() λ.Object {
						if λv := λ.Cal(ϒtry_get, ϒconfig, λ.NewFunction("<lambda>",
							[]λ.Param{
								{Name: "x"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒx = λargs[0]
								)
								return λ.GetItem(ϒx, λ.StrLiteral("poster"))
							}), λ.DictType); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.DictLiteral(map[λ.Object]λ.Object{})
						}
					}()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(ϒposter, "items"))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						τmp2 = τmp1
						ϒthumbnail_id = λ.GetItem(τmp2, λ.IntLiteral(0))
						ϒthumbnail = λ.GetItem(τmp2, λ.IntLiteral(1))
						ϒthumbnail_url = λ.Cal(ϒurljoin, ϒurl, λ.Calm(ϒthumbnail, "get", λ.StrLiteral("src")))
						if !λ.IsTrue(ϒthumbnail_url) {
							continue
						}
						λ.Calm(ϒthumbnails, "append", λ.DictLiteral(map[string]λ.Object{
							"id": func() λ.Object {
								if λv := λ.Calm(ϒthumbnail, "get", λ.StrLiteral("quality")); λ.IsTrue(λv) {
									return λv
								} else {
									return ϒthumbnail_id
								}
							}(),
							"url":        ϒthumbnail_url,
							"preference": λ.Cal(ϒquality_key, λ.Calm(ϒthumbnail, "get", λ.StrLiteral("quality"))),
						}))
					}
					return λ.DictLiteral(map[string]λ.Object{
						"id":      ϒvideo_id,
						"title":   ϒtitle,
						"is_live": ϒlive,
						"uploader": func() λ.Object {
							if λ.IsTrue(λ.Ne(ϒuploader, λ.StrLiteral("-"))) {
								return ϒuploader
							} else {
								return λ.None
							}
						}(),
						"upload_date": func() λ.Object {
							if λ.IsTrue(ϒupload_date) {
								return λ.GetItem(ϒupload_date, λ.NewSlice(λ.IntLiteral(0), λ.IntLiteral(8), λ.None))
							} else {
								return λ.None
							}
						}(),
						"duration":   ϒduration,
						"thumbnails": ϒthumbnails,
						"formats":    ϒformats,
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL":    NDREmbedBaseIE__VALID_URL,
				"_real_extract": NDREmbedBaseIE__real_extract,
			})
		}())
		NDREmbedIE = λ.Cal(λ.TypeType, λ.StrLiteral("NDREmbedIE"), λ.NewTuple(NDREmbedBaseIE), func() λ.Dict {
			var (
				NDREmbedIE_IE_NAME    λ.Object
				NDREmbedIE__VALID_URL λ.Object
			)
			NDREmbedIE_IE_NAME = λ.StrLiteral("ndr:embed")
			NDREmbedIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?ndr\\.de/(?:[^/]+/)*(?P<id>[\\da-z]+)-(?:player|externalPlayer)\\.html")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"IE_NAME":    NDREmbedIE_IE_NAME,
				"_VALID_URL": NDREmbedIE__VALID_URL,
			})
		}())
		NJoyEmbedIE = λ.Cal(λ.TypeType, λ.StrLiteral("NJoyEmbedIE"), λ.NewTuple(NDREmbedBaseIE), func() λ.Dict {
			var (
				NJoyEmbedIE_IE_NAME    λ.Object
				NJoyEmbedIE__VALID_URL λ.Object
			)
			NJoyEmbedIE_IE_NAME = λ.StrLiteral("njoy:embed")
			NJoyEmbedIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?n-joy\\.de/(?:[^/]+/)*(?P<id>[\\da-z]+)-(?:player|externalPlayer)_[^/]+\\.html")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"IE_NAME":    NJoyEmbedIE_IE_NAME,
				"_VALID_URL": NJoyEmbedIE__VALID_URL,
			})
		}())
	})
}
