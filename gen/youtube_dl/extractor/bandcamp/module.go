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
 * bandcamp/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/bandcamp.py
 */

package bandcamp

import (
	Ωrandom "github.com/tenta-browser/go-video-downloader/gen/random"
	Ωtime "github.com/tenta-browser/go-video-downloader/gen/time"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	BandcampAlbumIE    λ.Object
	BandcampIE         λ.Object
	BandcampWeeklyIE   λ.Object
	ExtractorError     λ.Object
	InfoExtractor      λ.Object
	KNOWN_EXTENSIONS   λ.Object
	ϒcompat_str        λ.Object
	ϒfloat_or_none     λ.Object
	ϒint_or_none       λ.Object
	ϒparse_filesize    λ.Object
	ϒstr_or_none       λ.Object
	ϒtry_get           λ.Object
	ϒunified_strdate   λ.Object
	ϒunified_timestamp λ.Object
	ϒupdate_url_query  λ.Object
	ϒurl_or_none       λ.Object
	ϒurljoin           λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_str = Ωcompat.ϒcompat_str
		ExtractorError = Ωutils.ExtractorError
		ϒfloat_or_none = Ωutils.ϒfloat_or_none
		ϒint_or_none = Ωutils.ϒint_or_none
		KNOWN_EXTENSIONS = Ωutils.KNOWN_EXTENSIONS
		ϒparse_filesize = Ωutils.ϒparse_filesize
		ϒstr_or_none = Ωutils.ϒstr_or_none
		ϒtry_get = Ωutils.ϒtry_get
		ϒupdate_url_query = Ωutils.ϒupdate_url_query
		ϒunified_strdate = Ωutils.ϒunified_strdate
		ϒunified_timestamp = Ωutils.ϒunified_timestamp
		ϒurl_or_none = Ωutils.ϒurl_or_none
		ϒurljoin = Ωutils.ϒurljoin
		BandcampIE = λ.Cal(λ.TypeType, λ.StrLiteral("BandcampIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				BandcampIE__VALID_URL         λ.Object
				BandcampIE__extract_data_attr λ.Object
				BandcampIE__real_extract      λ.Object
			)
			BandcampIE__VALID_URL = λ.StrLiteral("https?://[^/]+\\.bandcamp\\.com/track/(?P<id>[^/?#&]+)")
			BandcampIE__extract_data_attr = λ.NewFunction("_extract_data_attr",
				[]λ.Param{
					{Name: "self"},
					{Name: "webpage"},
					{Name: "video_id"},
					{Name: "attr", Def: λ.StrLiteral("tralbum")},
					{Name: "fatal", Def: λ.True},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒattr     = λargs[3]
						ϒfatal    = λargs[4]
						ϒself     = λargs[0]
						ϒvideo_id = λargs[2]
						ϒwebpage  = λargs[1]
					)
					return λ.Call(λ.GetAttr(ϒself, "_parse_json", nil), λ.NewArgs(
						λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
							λ.Mod(λ.StrLiteral("data-%s=([\"\\'])({.+?})\\1"), ϒattr),
							ϒwebpage,
							λ.Add(ϒattr, λ.StrLiteral(" data")),
						), λ.KWArgs{
							{Name: "group", Value: λ.IntLiteral(2)},
						}),
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "fatal", Value: ϒfatal},
					})
				})
			BandcampIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒabr_str               λ.Object
						ϒartist                λ.Object
						ϒblob                  λ.Object
						ϒcurrent               λ.Object
						ϒdownload_formats      λ.Object
						ϒdownload_formats_list λ.Object
						ϒdownload_link         λ.Object
						ϒdownload_webpage      λ.Object
						ϒdownloads             λ.Object
						ϒduration              λ.Object
						ϒembed                 λ.Object
						ϒext                   λ.Object
						ϒf                     λ.Object
						ϒfile_                 λ.Object
						ϒformat_id             λ.Object
						ϒformat_url            λ.Object
						ϒformats               λ.Object
						ϒinfo                  λ.Object
						ϒname                  λ.Object
						ϒretry_url             λ.Object
						ϒself                  = λargs[0]
						ϒstat                  λ.Object
						ϒstat_url              λ.Object
						ϒthumbnail             λ.Object
						ϒtimestamp             λ.Object
						ϒtitle                 λ.Object
						ϒtrack                 λ.Object
						ϒtrack_id              λ.Object
						ϒtrack_info            λ.Object
						ϒtrack_number          λ.Object
						ϒtralbum               λ.Object
						ϒurl                   = λargs[1]
						ϒwebpage               λ.Object
						τmp0                   λ.Object
						τmp1                   λ.Object
						τmp2                   λ.Object
					)
					ϒtitle = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒwebpage = λ.Calm(ϒself, "_download_webpage", ϒurl, ϒtitle)
					ϒtralbum = λ.Calm(ϒself, "_extract_data_attr", ϒwebpage, ϒtitle)
					ϒthumbnail = λ.Calm(ϒself, "_og_search_thumbnail", ϒwebpage)
					ϒtrack_id = λ.None
					ϒtrack = λ.None
					ϒtrack_number = λ.None
					ϒduration = λ.None
					ϒformats = λ.NewList()
					ϒtrack_info = λ.Cal(ϒtry_get, ϒtralbum, λ.NewFunction("<lambda>",
						[]λ.Param{
							{Name: "x"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒx = λargs[0]
							)
							return λ.GetItem(λ.GetItem(ϒx, λ.StrLiteral("trackinfo")), λ.IntLiteral(0))
						}), λ.DictType)
					if λ.IsTrue(ϒtrack_info) {
						ϒfile_ = λ.Calm(ϒtrack_info, "get", λ.StrLiteral("file"))
						if λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒfile_, λ.DictType)) {
							τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(ϒfile_, "items"))
							for {
								if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
									break
								}
								τmp2 = τmp1
								ϒformat_id = λ.GetItem(τmp2, λ.IntLiteral(0))
								ϒformat_url = λ.GetItem(τmp2, λ.IntLiteral(1))
								if !λ.IsTrue(λ.Cal(ϒurl_or_none, ϒformat_url)) {
									continue
								}
								τmp2 = λ.Calm(ϒformat_id, "split", λ.StrLiteral("-"), λ.IntLiteral(1))
								ϒext = λ.GetItem(τmp2, λ.IntLiteral(0))
								ϒabr_str = λ.GetItem(τmp2, λ.IntLiteral(1))
								λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
									"format_id": ϒformat_id,
									"url":       λ.Calm(ϒself, "_proto_relative_url", ϒformat_url, λ.StrLiteral("http:")),
									"ext":       ϒext,
									"vcodec":    λ.StrLiteral("none"),
									"acodec":    ϒext,
									"abr":       λ.Cal(ϒint_or_none, ϒabr_str),
								}))
							}
						}
						ϒtrack = λ.Calm(ϒtrack_info, "get", λ.StrLiteral("title"))
						ϒtrack_id = λ.Cal(ϒstr_or_none, func() λ.Object {
							if λv := λ.Calm(ϒtrack_info, "get", λ.StrLiteral("track_id")); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Calm(ϒtrack_info, "get", λ.StrLiteral("id"))
							}
						}())
						ϒtrack_number = λ.Cal(ϒint_or_none, λ.Calm(ϒtrack_info, "get", λ.StrLiteral("track_num")))
						ϒduration = λ.Cal(ϒfloat_or_none, λ.Calm(ϒtrack_info, "get", λ.StrLiteral("duration")))
					}
					ϒembed = λ.Calm(ϒself, "_extract_data_attr", ϒwebpage, ϒtitle, λ.StrLiteral("embed"), λ.False)
					ϒcurrent = func() λ.Object {
						if λv := λ.Calm(ϒtralbum, "get", λ.StrLiteral("current")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.DictLiteral(map[λ.Object]λ.Object{})
						}
					}()
					ϒartist = func() λ.Object {
						if λv := λ.Calm(ϒembed, "get", λ.StrLiteral("artist")); λ.IsTrue(λv) {
							return λv
						} else if λv := λ.Calm(ϒcurrent, "get", λ.StrLiteral("artist")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Calm(ϒtralbum, "get", λ.StrLiteral("artist"))
						}
					}()
					ϒtimestamp = λ.Cal(ϒunified_timestamp, func() λ.Object {
						if λv := λ.Calm(ϒcurrent, "get", λ.StrLiteral("publish_date")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Calm(ϒtralbum, "get", λ.StrLiteral("album_publish_date"))
						}
					}())
					ϒdownload_link = λ.Calm(ϒtralbum, "get", λ.StrLiteral("freeDownloadPage"))
					if λ.IsTrue(ϒdownload_link) {
						ϒtrack_id = λ.Cal(ϒcompat_str, λ.GetItem(ϒtralbum, λ.StrLiteral("id")))
						ϒdownload_webpage = λ.Calm(ϒself, "_download_webpage", ϒdownload_link, ϒtrack_id, λ.StrLiteral("Downloading free downloads page"))
						ϒblob = λ.Calm(ϒself, "_extract_data_attr", ϒdownload_webpage, ϒtrack_id, λ.StrLiteral("blob"))
						ϒinfo = λ.Cal(ϒtry_get, ϒblob, λ.NewTuple(
							λ.NewFunction("<lambda>",
								[]λ.Param{
									{Name: "x"},
								},
								0, false, false,
								func(λargs []λ.Object) λ.Object {
									var (
										ϒx = λargs[0]
									)
									return λ.GetItem(λ.GetItem(ϒx, λ.StrLiteral("digital_items")), λ.IntLiteral(0))
								}),
							λ.NewFunction("<lambda>",
								[]λ.Param{
									{Name: "x"},
								},
								0, false, false,
								func(λargs []λ.Object) λ.Object {
									var (
										ϒx = λargs[0]
									)
									return λ.GetItem(λ.GetItem(ϒx, λ.StrLiteral("download_items")), λ.IntLiteral(0))
								}),
						), λ.DictType)
						if λ.IsTrue(ϒinfo) {
							ϒdownloads = λ.Calm(ϒinfo, "get", λ.StrLiteral("downloads"))
							if λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒdownloads, λ.DictType)) {
								if !λ.IsTrue(ϒtrack) {
									ϒtrack = λ.Calm(ϒinfo, "get", λ.StrLiteral("title"))
								}
								if !λ.IsTrue(ϒartist) {
									ϒartist = λ.Calm(ϒinfo, "get", λ.StrLiteral("artist"))
								}
								if !λ.IsTrue(ϒthumbnail) {
									ϒthumbnail = λ.Calm(ϒinfo, "get", λ.StrLiteral("thumb_url"))
								}
								ϒdownload_formats = λ.DictLiteral(map[λ.Object]λ.Object{})
								ϒdownload_formats_list = λ.Calm(ϒblob, "get", λ.StrLiteral("download_formats"))
								if λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒdownload_formats_list, λ.ListType)) {
									τmp0 = λ.Cal(λ.BuiltinIter, λ.GetItem(ϒblob, λ.StrLiteral("download_formats")))
									for {
										if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
											break
										}
										ϒf = τmp1
										τmp2 = λ.NewTuple(
											λ.Calm(ϒf, "get", λ.StrLiteral("name")),
											λ.Calm(ϒf, "get", λ.StrLiteral("file_extension")),
										)
										ϒname = λ.GetItem(τmp2, λ.IntLiteral(0))
										ϒext = λ.GetItem(τmp2, λ.IntLiteral(1))
										if λ.IsTrue(λ.Cal(λ.BuiltinAll, λ.Cal(λ.NewFunction("<generator>",
											nil,
											0, false, false,
											func(λargs []λ.Object) λ.Object {
												return λ.NewGenerator(func(λgy λ.Yielder) λ.Object {
													var (
														ϒx   λ.Object
														τmp0 λ.Object
														τmp1 λ.Object
													)
													τmp0 = λ.Cal(λ.BuiltinIter, λ.NewTuple(
														ϒname,
														ϒext,
													))
													for {
														if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
															break
														}
														ϒx = τmp1
														λgy.Yield(λ.Cal(λ.BuiltinIsInstance, ϒx, ϒcompat_str))
													}
													return λ.None
												})
											})))) {
											λ.SetItem(ϒdownload_formats, ϒname, λ.Calm(ϒext, "strip", λ.StrLiteral(".")))
										}
									}
								}
								τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(ϒdownloads, "items"))
								for {
									if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
										break
									}
									τmp2 = τmp1
									ϒformat_id = λ.GetItem(τmp2, λ.IntLiteral(0))
									ϒf = λ.GetItem(τmp2, λ.IntLiteral(1))
									ϒformat_url = λ.Calm(ϒf, "get", λ.StrLiteral("url"))
									if !λ.IsTrue(ϒformat_url) {
										continue
									}
									ϒstat_url = λ.Cal(ϒupdate_url_query, λ.Calm(ϒformat_url, "replace", λ.StrLiteral("/download/"), λ.StrLiteral("/statdownload/")), λ.DictLiteral(map[string]λ.Object{
										".rand": λ.Cal(λ.IntType, λ.Mul(λ.Mul(λ.Cal(Ωtime.ϒtime), λ.IntLiteral(1000)), λ.Cal(Ωrandom.ϒrandom))),
									}))
									ϒformat_id = func() λ.Object {
										if λv := λ.Calm(ϒf, "get", λ.StrLiteral("encoding_name")); λ.IsTrue(λv) {
											return λv
										} else {
											return ϒformat_id
										}
									}()
									ϒstat = λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
										ϒstat_url,
										ϒtrack_id,
										λ.Mod(λ.StrLiteral("Downloading %s JSON"), ϒformat_id),
									), λ.KWArgs{
										{Name: "transform_source", Value: λ.NewFunction("<lambda>",
											[]λ.Param{
												{Name: "s"},
											},
											0, false, false,
											func(λargs []λ.Object) λ.Object {
												var (
													ϒs = λargs[0]
												)
												return λ.GetItem(ϒs, λ.NewSlice(λ.Calm(ϒs, "index", λ.StrLiteral("{")), λ.Add(λ.Calm(ϒs, "rindex", λ.StrLiteral("}")), λ.IntLiteral(1)), λ.None))
											})},
										{Name: "fatal", Value: λ.False},
									})
									if !λ.IsTrue(ϒstat) {
										continue
									}
									ϒretry_url = λ.Cal(ϒurl_or_none, λ.Calm(ϒstat, "get", λ.StrLiteral("retry_url")))
									if !λ.IsTrue(ϒretry_url) {
										continue
									}
									λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
										"url":         λ.Calm(ϒself, "_proto_relative_url", ϒretry_url, λ.StrLiteral("http:")),
										"ext":         λ.Calm(ϒdownload_formats, "get", ϒformat_id),
										"format_id":   ϒformat_id,
										"format_note": λ.Calm(ϒf, "get", λ.StrLiteral("description")),
										"filesize":    λ.Cal(ϒparse_filesize, λ.Calm(ϒf, "get", λ.StrLiteral("size_mb"))),
										"vcodec":      λ.StrLiteral("none"),
									}))
								}
							}
						}
					}
					λ.Calm(ϒself, "_sort_formats", ϒformats)
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
					if !λ.IsTrue(ϒduration) {
						ϒduration = λ.Cal(ϒfloat_or_none, λ.Call(λ.GetAttr(ϒself, "_html_search_meta", nil), λ.NewArgs(
							λ.StrLiteral("duration"),
							ϒwebpage,
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
						}))
					}
					return λ.DictLiteral(map[string]λ.Object{
						"id":           ϒtrack_id,
						"title":        ϒtitle,
						"thumbnail":    ϒthumbnail,
						"uploader":     ϒartist,
						"timestamp":    ϒtimestamp,
						"release_date": λ.Cal(ϒunified_strdate, λ.Calm(ϒtralbum, "get", λ.StrLiteral("album_release_date"))),
						"duration":     ϒduration,
						"track":        ϒtrack,
						"track_number": ϒtrack_number,
						"track_id":     ϒtrack_id,
						"artist":       ϒartist,
						"album":        λ.Calm(ϒembed, "get", λ.StrLiteral("album_title")),
						"formats":      ϒformats,
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL":         BandcampIE__VALID_URL,
				"_extract_data_attr": BandcampIE__extract_data_attr,
				"_real_extract":      BandcampIE__real_extract,
			})
		}())
		BandcampAlbumIE = λ.Cal(λ.TypeType, λ.StrLiteral("BandcampAlbumIE"), λ.NewTuple(BandcampIE), func() λ.Dict {
			var (
				BandcampAlbumIE__VALID_URL λ.Object
				BandcampAlbumIE_suitable   λ.Object
			)
			BandcampAlbumIE__VALID_URL = λ.StrLiteral("https?://(?:(?P<subdomain>[^.]+)\\.)?bandcamp\\.com(?:/album/(?P<id>[^/?#&]+))?")
			BandcampAlbumIE_suitable = λ.NewFunction("suitable",
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
						if λ.IsTrue(func() λ.Object {
							if λv := λ.Calm(BandcampWeeklyIE, "suitable", ϒurl); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Calm(BandcampIE, "suitable", ϒurl)
							}
						}()) {
							return λ.False
						} else {
							return λ.Calm(λ.Cal(λ.SuperType, BandcampAlbumIE, ϒcls), "suitable", ϒurl)
						}
					}()
				})
			BandcampAlbumIE_suitable = λ.Cal(λ.ClassMethodType, BandcampAlbumIE_suitable)
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": BandcampAlbumIE__VALID_URL,
				"suitable":   BandcampAlbumIE_suitable,
			})
		}())
		BandcampWeeklyIE = λ.Cal(λ.TypeType, λ.StrLiteral("BandcampWeeklyIE"), λ.NewTuple(BandcampIE), func() λ.Dict {
			var (
				BandcampWeeklyIE_IE_NAME       λ.Object
				BandcampWeeklyIE__VALID_URL    λ.Object
				BandcampWeeklyIE__real_extract λ.Object
			)
			BandcampWeeklyIE_IE_NAME = λ.StrLiteral("Bandcamp:weekly")
			BandcampWeeklyIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?bandcamp\\.com/?\\?(?:.*?&)?show=(?P<id>\\d+)")
			BandcampWeeklyIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒblob       λ.Object
						ϒext        λ.Object
						ϒformat_id  λ.Object
						ϒformat_url λ.Object
						ϒformats    λ.Object
						ϒknown_ext  λ.Object
						ϒself       = λargs[0]
						ϒshow       λ.Object
						ϒshow_id    λ.Object
						ϒsubtitle   λ.Object
						ϒtitle      λ.Object
						ϒurl        = λargs[1]
						ϒwebpage    λ.Object
						τmp0        λ.Object
						τmp1        λ.Object
						τmp2        λ.Object
						τmp3        λ.Object
					)
					ϒshow_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒwebpage = λ.Calm(ϒself, "_download_webpage", ϒurl, ϒshow_id)
					ϒblob = λ.Calm(ϒself, "_extract_data_attr", ϒwebpage, ϒshow_id, λ.StrLiteral("blob"))
					ϒshow = λ.GetItem(λ.GetItem(ϒblob, λ.StrLiteral("bcw_data")), ϒshow_id)
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(λ.GetItem(ϒshow, λ.StrLiteral("audio_stream")), "items"))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						τmp2 = τmp1
						ϒformat_id = λ.GetItem(τmp2, λ.IntLiteral(0))
						ϒformat_url = λ.GetItem(τmp2, λ.IntLiteral(1))
						if !λ.IsTrue(λ.Cal(ϒurl_or_none, ϒformat_url)) {
							continue
						}
						τmp2 = λ.Cal(λ.BuiltinIter, KNOWN_EXTENSIONS)
						for {
							if τmp3 = λ.NextDefault(τmp2, λ.AfterLast); τmp3 == λ.AfterLast {
								break
							}
							ϒknown_ext = τmp3
							if λ.Contains(ϒformat_id, ϒknown_ext) {
								ϒext = ϒknown_ext
								break
							}
						}
						if τmp3 == λ.AfterLast {
							ϒext = λ.None
						}
						λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
							"format_id": ϒformat_id,
							"url":       ϒformat_url,
							"ext":       ϒext,
							"vcodec":    λ.StrLiteral("none"),
						}))
					}
					λ.Calm(ϒself, "_sort_formats", ϒformats)
					ϒtitle = func() λ.Object {
						if λv := λ.Calm(ϒshow, "get", λ.StrLiteral("audio_title")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.StrLiteral("Bandcamp Weekly")
						}
					}()
					ϒsubtitle = λ.Calm(ϒshow, "get", λ.StrLiteral("subtitle"))
					if λ.IsTrue(ϒsubtitle) {
						τmp0 = λ.IAdd(ϒtitle, λ.Mod(λ.StrLiteral(" - %s"), ϒsubtitle))
						ϒtitle = τmp0
					}
					return λ.DictLiteral(map[string]λ.Object{
						"id":    ϒshow_id,
						"title": ϒtitle,
						"description": func() λ.Object {
							if λv := λ.Calm(ϒshow, "get", λ.StrLiteral("desc")); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Calm(ϒshow, "get", λ.StrLiteral("short_desc"))
							}
						}(),
						"duration":     λ.Cal(ϒfloat_or_none, λ.Calm(ϒshow, "get", λ.StrLiteral("audio_duration"))),
						"is_live":      λ.False,
						"release_date": λ.Cal(ϒunified_strdate, λ.Calm(ϒshow, "get", λ.StrLiteral("published_date"))),
						"series":       λ.StrLiteral("Bandcamp Weekly"),
						"episode":      λ.Calm(ϒshow, "get", λ.StrLiteral("subtitle")),
						"episode_id":   ϒshow_id,
						"formats":      ϒformats,
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"IE_NAME":       BandcampWeeklyIE_IE_NAME,
				"_VALID_URL":    BandcampWeeklyIE__VALID_URL,
				"_real_extract": BandcampWeeklyIE__real_extract,
			})
		}())
	})
}
