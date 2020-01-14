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
		ϒparse_iso8601 = Ωutils.ϒparse_iso8601
		ϒqualities = Ωutils.ϒqualities
		ϒtry_get = Ωutils.ϒtry_get
		ϒurljoin = Ωutils.ϒurljoin
		NDRBaseIE = λ.Cal(λ.TypeType, λ.NewStr("NDRBaseIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
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
								τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(ϒmobj, "groups", nil)))
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
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒdisplay_id)
					return λ.Cal(λ.GetAttr(ϒself, "_extract_embed", nil), ϒwebpage, ϒdisplay_id)
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_real_extract"): NDRBaseIE__real_extract,
			})
		}())
		NDRIE = λ.Cal(λ.TypeType, λ.NewStr("NDRIE"), λ.NewTuple(NDRBaseIE), func() λ.Dict {
			var (
				NDRIE_IE_NAME        λ.Object
				NDRIE__VALID_URL     λ.Object
				NDRIE__extract_embed λ.Object
			)
			NDRIE_IE_NAME = λ.NewStr("ndr")
			NDRIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?ndr\\.de/(?:[^/]+/)*(?P<id>[^/?#]+),[\\da-z]+\\.html")
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
						ϒself        = λargs[0]
						ϒtimestamp   λ.Object
						ϒwebpage     = λargs[1]
					)
					ϒembed_url = λ.Call(λ.GetAttr(ϒself, "_html_search_meta", nil), λ.NewArgs(
						λ.NewStr("embedURL"),
						ϒwebpage,
						λ.NewStr("embed URL"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.True},
					})
					ϒdescription = func() λ.Object {
						if λv := λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
							λ.NewStr("<p[^>]+itemprop=\"description\">([^<]+)</p>"),
							ϒwebpage,
							λ.NewStr("description"),
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
						}); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Cal(λ.GetAttr(ϒself, "_og_search_description", nil), ϒwebpage)
						}
					}()
					ϒtimestamp = λ.Cal(ϒparse_iso8601, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr("<span[^>]+itemprop=\"(?:datePublished|uploadDate)\"[^>]+content=\"([^\"]+)\""),
						ϒwebpage,
						λ.NewStr("upload date"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					}))
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("_type"):       λ.NewStr("url_transparent"),
						λ.NewStr("url"):         ϒembed_url,
						λ.NewStr("display_id"):  ϒdisplay_id,
						λ.NewStr("description"): ϒdescription,
						λ.NewStr("timestamp"):   ϒtimestamp,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("IE_NAME"):        NDRIE_IE_NAME,
				λ.NewStr("_VALID_URL"):     NDRIE__VALID_URL,
				λ.NewStr("_extract_embed"): NDRIE__extract_embed,
			})
		}())
		NJoyIE = λ.Cal(λ.TypeType, λ.NewStr("NJoyIE"), λ.NewTuple(NDRBaseIE), func() λ.Dict {
			var (
				NJoyIE__VALID_URL λ.Object
			)
			NJoyIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?n-joy\\.de/(?:[^/]+/)*(?:(?P<display_id>[^/?#]+),)?(?P<id>[\\da-z]+)\\.html")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): NJoyIE__VALID_URL,
			})
		}())
		NDREmbedBaseIE = λ.Cal(λ.TypeType, λ.NewStr("NDREmbedBaseIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				NDREmbedBaseIE__VALID_URL    λ.Object
				NDREmbedBaseIE__real_extract λ.Object
			)
			NDREmbedBaseIE__VALID_URL = λ.NewStr("(?:ndr:(?P<id_s>[\\da-z]+)|https?://www\\.ndr\\.de/(?P<id>[\\da-z]+)-ppjson\\.json)")
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
						if λv := λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("id")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("id_s"))
						}
					}()
					ϒppjson = λ.Cal(λ.GetAttr(ϒself, "_download_json", nil), λ.Mod(λ.NewStr("http://www.ndr.de/%s-ppjson.json"), ϒvideo_id), ϒvideo_id)
					ϒplaylist = λ.GetItem(ϒppjson, λ.NewStr("playlist"))
					ϒformats = λ.NewList()
					ϒquality_key = λ.Cal(ϒqualities, λ.NewTuple(
						λ.NewStr("xs"),
						λ.NewStr("s"),
						λ.NewStr("m"),
						λ.NewStr("l"),
						λ.NewStr("xl"),
					))
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(ϒplaylist, "items", nil)))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						τmp2 = τmp1
						ϒformat_id = λ.GetItem(τmp2, λ.NewInt(0))
						ϒf = λ.GetItem(τmp2, λ.NewInt(1))
						ϒsrc = λ.Cal(λ.GetAttr(ϒf, "get", nil), λ.NewStr("src"))
						if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒsrc))) {
							continue
						}
						ϒext = λ.Cal(ϒdetermine_ext, ϒsrc, λ.None)
						if λ.IsTrue(λ.Eq(ϒext, λ.NewStr("f4m"))) {
							λ.Cal(λ.GetAttr(ϒformats, "extend", nil), λ.Call(λ.GetAttr(ϒself, "_extract_f4m_formats", nil), λ.NewArgs(
								λ.Add(ϒsrc, λ.NewStr("?hdcore=3.7.0&plugin=aasp-3.7.0.39.44")),
								ϒvideo_id,
							), λ.KWArgs{
								{Name: "f4m_id", Value: λ.NewStr("hds")},
								{Name: "fatal", Value: λ.False},
							}))
						} else {
							if λ.IsTrue(λ.Eq(ϒext, λ.NewStr("m3u8"))) {
								λ.Cal(λ.GetAttr(ϒformats, "extend", nil), λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
									ϒsrc,
									ϒvideo_id,
									λ.NewStr("mp4"),
								), λ.KWArgs{
									{Name: "m3u8_id", Value: λ.NewStr("hls")},
									{Name: "entry_protocol", Value: λ.NewStr("m3u8_native")},
									{Name: "fatal", Value: λ.False},
								}))
							} else {
								ϒquality = λ.Cal(λ.GetAttr(ϒf, "get", nil), λ.NewStr("quality"))
								ϒff = λ.NewDictWithTable(map[λ.Object]λ.Object{
									λ.NewStr("url"): ϒsrc,
									λ.NewStr("format_id"): func() λ.Object {
										if λv := ϒquality; λ.IsTrue(λv) {
											return λv
										} else {
											return ϒformat_id
										}
									}(),
									λ.NewStr("quality"): λ.Cal(ϒquality_key, ϒquality),
								})
								ϒtype_ = λ.Cal(λ.GetAttr(ϒf, "get", nil), λ.NewStr("type"))
								if λ.IsTrue(func() λ.Object {
									if λv := ϒtype_; !λ.IsTrue(λv) {
										return λv
									} else {
										return λ.Eq(λ.GetItem(λ.Cal(λ.GetAttr(ϒtype_, "split", nil), λ.NewStr("/")), λ.NewInt(0)), λ.NewStr("audio"))
									}
								}()) {
									λ.SetItem(ϒff, λ.NewStr("vcodec"), λ.NewStr("none"))
									λ.SetItem(ϒff, λ.NewStr("ext"), func() λ.Object {
										if λv := ϒext; λ.IsTrue(λv) {
											return λv
										} else {
											return λ.NewStr("mp3")
										}
									}())
								}
								λ.Cal(λ.GetAttr(ϒformats, "append", nil), ϒff)
							}
						}
					}
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					ϒconfig = λ.GetItem(ϒplaylist, λ.NewStr("config"))
					ϒlive = λ.NewBool(λ.Contains(λ.NewList(
						λ.NewStr("httpVideoLive"),
						λ.NewStr("httpAudioLive"),
					), λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒplaylist, "get", nil), λ.NewStr("config"), λ.NewDictWithTable(map[λ.Object]λ.Object{})), "get", nil), λ.NewStr("streamType"))))
					ϒtitle = λ.GetItem(ϒconfig, λ.NewStr("title"))
					if λ.IsTrue(ϒlive) {
						ϒtitle = λ.Cal(λ.GetAttr(ϒself, "_live_title", nil), ϒtitle)
					}
					ϒuploader = λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒppjson, "get", nil), λ.NewStr("config"), λ.NewDictWithTable(map[λ.Object]λ.Object{})), "get", nil), λ.NewStr("branding"))
					ϒupload_date = λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒppjson, "get", nil), λ.NewStr("config"), λ.NewDictWithTable(map[λ.Object]λ.Object{})), "get", nil), λ.NewStr("publicationDate"))
					ϒduration = λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒconfig, "get", nil), λ.NewStr("duration")))
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
								return λ.GetItem(ϒx, λ.NewStr("poster"))
							}), λ.DictType); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.NewDictWithTable(map[λ.Object]λ.Object{})
						}
					}()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(ϒposter, "items", nil)))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						τmp2 = τmp1
						ϒthumbnail_id = λ.GetItem(τmp2, λ.NewInt(0))
						ϒthumbnail = λ.GetItem(τmp2, λ.NewInt(1))
						ϒthumbnail_url = λ.Cal(ϒurljoin, ϒurl, λ.Cal(λ.GetAttr(ϒthumbnail, "get", nil), λ.NewStr("src")))
						if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒthumbnail_url))) {
							continue
						}
						λ.Cal(λ.GetAttr(ϒthumbnails, "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("id"): func() λ.Object {
								if λv := λ.Cal(λ.GetAttr(ϒthumbnail, "get", nil), λ.NewStr("quality")); λ.IsTrue(λv) {
									return λv
								} else {
									return ϒthumbnail_id
								}
							}(),
							λ.NewStr("url"):        ϒthumbnail_url,
							λ.NewStr("preference"): λ.Cal(ϒquality_key, λ.Cal(λ.GetAttr(ϒthumbnail, "get", nil), λ.NewStr("quality"))),
						}))
					}
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):      ϒvideo_id,
						λ.NewStr("title"):   ϒtitle,
						λ.NewStr("is_live"): ϒlive,
						λ.NewStr("uploader"): func() λ.Object {
							if λ.IsTrue(λ.Ne(ϒuploader, λ.NewStr("-"))) {
								return ϒuploader
							} else {
								return λ.None
							}
						}(),
						λ.NewStr("upload_date"): func() λ.Object {
							if λ.IsTrue(ϒupload_date) {
								return λ.GetItem(ϒupload_date, λ.NewSlice(λ.NewInt(0), λ.NewInt(8), λ.None))
							} else {
								return λ.None
							}
						}(),
						λ.NewStr("duration"):   ϒduration,
						λ.NewStr("thumbnails"): ϒthumbnails,
						λ.NewStr("formats"):    ϒformats,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"):    NDREmbedBaseIE__VALID_URL,
				λ.NewStr("_real_extract"): NDREmbedBaseIE__real_extract,
			})
		}())
		NDREmbedIE = λ.Cal(λ.TypeType, λ.NewStr("NDREmbedIE"), λ.NewTuple(NDREmbedBaseIE), func() λ.Dict {
			var (
				NDREmbedIE_IE_NAME    λ.Object
				NDREmbedIE__VALID_URL λ.Object
			)
			NDREmbedIE_IE_NAME = λ.NewStr("ndr:embed")
			NDREmbedIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?ndr\\.de/(?:[^/]+/)*(?P<id>[\\da-z]+)-(?:player|externalPlayer)\\.html")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("IE_NAME"):    NDREmbedIE_IE_NAME,
				λ.NewStr("_VALID_URL"): NDREmbedIE__VALID_URL,
			})
		}())
		NJoyEmbedIE = λ.Cal(λ.TypeType, λ.NewStr("NJoyEmbedIE"), λ.NewTuple(NDREmbedBaseIE), func() λ.Dict {
			var (
				NJoyEmbedIE_IE_NAME    λ.Object
				NJoyEmbedIE__VALID_URL λ.Object
			)
			NJoyEmbedIE_IE_NAME = λ.NewStr("njoy:embed")
			NJoyEmbedIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?n-joy\\.de/(?:[^/]+/)*(?P<id>[\\da-z]+)-(?:player|externalPlayer)_[^/]+\\.html")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("IE_NAME"):    NJoyEmbedIE_IE_NAME,
				λ.NewStr("_VALID_URL"): NJoyEmbedIE__VALID_URL,
			})
		}())
	})
}
