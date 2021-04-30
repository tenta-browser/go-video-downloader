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
 * googledrive/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/googledrive.py
 */

package googledrive

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError        λ.Object
	GoogleDriveIE         λ.Object
	InfoExtractor         λ.Object
	ϒcompat_parse_qs      λ.Object
	ϒdetermine_ext        λ.Object
	ϒget_element_by_class λ.Object
	ϒint_or_none          λ.Object
	ϒlowercase_escape     λ.Object
	ϒtry_get              λ.Object
	ϒupdate_url_query     λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_parse_qs = Ωcompat.ϒcompat_parse_qs
		ϒdetermine_ext = Ωutils.ϒdetermine_ext
		ExtractorError = Ωutils.ExtractorError
		ϒget_element_by_class = Ωutils.ϒget_element_by_class
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒlowercase_escape = Ωutils.ϒlowercase_escape
		ϒtry_get = Ωutils.ϒtry_get
		ϒupdate_url_query = Ωutils.ϒupdate_url_query
		GoogleDriveIE = λ.Cal(λ.TypeType, λ.StrLiteral("GoogleDriveIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				GoogleDriveIE__FORMATS_EXT  λ.Object
				GoogleDriveIE__VALID_URL    λ.Object
				GoogleDriveIE__real_extract λ.Object
			)
			GoogleDriveIE__VALID_URL = λ.StrLiteral("(?x)\n                        https?://\n                            (?:\n                                (?:docs|drive)\\.google\\.com/\n                                (?:\n                                    (?:uc|open)\\?.*?id=|\n                                    file/d/\n                                )|\n                                video\\.google\\.com/get_player\\?.*?docid=\n                            )\n                            (?P<id>[a-zA-Z0-9_-]{28,})\n                    ")
			GoogleDriveIE__FORMATS_EXT = λ.DictLiteral(map[string]string{
				"5":  "flv",
				"6":  "flv",
				"13": "3gp",
				"17": "3gp",
				"18": "mp4",
				"22": "mp4",
				"34": "flv",
				"35": "flv",
				"36": "3gp",
				"37": "mp4",
				"38": "mp4",
				"43": "webm",
				"44": "webm",
				"45": "webm",
				"46": "webm",
				"59": "mp4",
			})
			GoogleDriveIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒadd_source_format    λ.Object
						ϒconfirm              λ.Object
						ϒconfirmation_webpage λ.Object
						ϒconfirmed_source_url λ.Object
						ϒf                    λ.Object
						ϒfmt                  λ.Object
						ϒfmt_list             λ.Object
						ϒfmt_stream           λ.Object
						ϒfmt_stream_map       λ.Object
						ϒfmt_stream_split     λ.Object
						ϒformat_id            λ.Object
						ϒformat_url           λ.Object
						ϒformats              λ.Object
						ϒget_value            λ.Object
						ϒhl                   λ.Object
						ϒmobj                 λ.Object
						ϒreason               λ.Object
						ϒrequest_source_file  λ.Object
						ϒresolution           λ.Object
						ϒresolutions          λ.Object
						ϒself                 = λargs[0]
						ϒsource_url           λ.Object
						ϒsubtitles_id         λ.Object
						ϒtitle                λ.Object
						ϒttsurl               λ.Object
						ϒurl                  = λargs[1]
						ϒurlh                 λ.Object
						ϒvideo_id             λ.Object
						ϒvideo_info           λ.Object
						τmp0                  λ.Object
						τmp1                  λ.Object
						τmp2                  λ.Object
					)
					ϒvideo_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒvideo_info = λ.Cal(ϒcompat_parse_qs, λ.Call(λ.GetAttr(ϒself, "_download_webpage", nil), λ.NewArgs(
						λ.StrLiteral("https://drive.google.com/get_video_info"),
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "query", Value: λ.DictLiteral(map[string]λ.Object{
							"docid": ϒvideo_id,
						})},
					}))
					ϒget_value = λ.NewFunction("get_value",
						[]λ.Param{
							{Name: "key"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒkey = λargs[0]
							)
							return λ.Cal(ϒtry_get, ϒvideo_info, λ.NewFunction("<lambda>",
								[]λ.Param{
									{Name: "x"},
								},
								0, false, false,
								func(λargs []λ.Object) λ.Object {
									var (
										ϒx = λargs[0]
									)
									return λ.GetItem(λ.GetItem(ϒx, ϒkey), λ.IntLiteral(0))
								}))
						})
					ϒreason = λ.Cal(ϒget_value, λ.StrLiteral("reason"))
					ϒtitle = λ.Cal(ϒget_value, λ.StrLiteral("title"))
					if λ.IsTrue(func() λ.Object {
						if λv := λ.NewBool(!λ.IsTrue(ϒtitle)); !λ.IsTrue(λv) {
							return λv
						} else {
							return ϒreason
						}
					}()) {
						panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(ϒreason), λ.KWArgs{
							{Name: "expected", Value: λ.True},
						})))
					}
					ϒformats = λ.NewList()
					ϒfmt_stream_map = λ.Calm(func() λ.Object {
						if λv := λ.Cal(ϒget_value, λ.StrLiteral("fmt_stream_map")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.StrLiteral("")
						}
					}(), "split", λ.StrLiteral(","))
					ϒfmt_list = λ.Calm(func() λ.Object {
						if λv := λ.Cal(ϒget_value, λ.StrLiteral("fmt_list")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.StrLiteral("")
						}
					}(), "split", λ.StrLiteral(","))
					if λ.IsTrue(func() λ.Object {
						if λv := ϒfmt_stream_map; !λ.IsTrue(λv) {
							return λv
						} else {
							return ϒfmt_list
						}
					}()) {
						ϒresolutions = λ.DictLiteral(map[λ.Object]λ.Object{})
						τmp0 = λ.Cal(λ.BuiltinIter, ϒfmt_list)
						for {
							if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
								break
							}
							ϒfmt = τmp1
							ϒmobj = λ.Cal(Ωre.ϒsearch, λ.StrLiteral("^(?P<format_id>\\d+)/(?P<width>\\d+)[xX](?P<height>\\d+)"), ϒfmt)
							if λ.IsTrue(ϒmobj) {
								λ.SetItem(ϒresolutions, λ.Calm(ϒmobj, "group", λ.StrLiteral("format_id")), λ.NewTuple(
									λ.Cal(λ.IntType, λ.Calm(ϒmobj, "group", λ.StrLiteral("width"))),
									λ.Cal(λ.IntType, λ.Calm(ϒmobj, "group", λ.StrLiteral("height"))),
								))
							}
						}
						τmp0 = λ.Cal(λ.BuiltinIter, ϒfmt_stream_map)
						for {
							if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
								break
							}
							ϒfmt_stream = τmp1
							ϒfmt_stream_split = λ.Calm(ϒfmt_stream, "split", λ.StrLiteral("|"))
							if λ.IsTrue(λ.Lt(λ.Cal(λ.BuiltinLen, ϒfmt_stream_split), λ.IntLiteral(2))) {
								continue
							}
							τmp2 = λ.UnpackIterable(λ.GetItem(ϒfmt_stream_split, λ.NewSlice(λ.None, λ.IntLiteral(2), λ.None)), 2)
							ϒformat_id = λ.GetItem(τmp2, λ.IntLiteral(0))
							ϒformat_url = λ.GetItem(τmp2, λ.IntLiteral(1))
							ϒf = λ.DictLiteral(map[string]λ.Object{
								"url":       λ.Cal(ϒlowercase_escape, ϒformat_url),
								"format_id": ϒformat_id,
								"ext":       λ.GetItem(λ.GetAttr(ϒself, "_FORMATS_EXT", nil), ϒformat_id),
							})
							ϒresolution = λ.Calm(ϒresolutions, "get", ϒformat_id)
							if λ.IsTrue(ϒresolution) {
								λ.Calm(ϒf, "update", λ.DictLiteral(map[string]λ.Object{
									"width":  λ.GetItem(ϒresolution, λ.IntLiteral(0)),
									"height": λ.GetItem(ϒresolution, λ.IntLiteral(1)),
								}))
							}
							λ.Calm(ϒformats, "append", ϒf)
						}
					}
					ϒsource_url = λ.Cal(ϒupdate_url_query, λ.StrLiteral("https://drive.google.com/uc"), λ.DictLiteral(map[string]λ.Object{
						"id":     ϒvideo_id,
						"export": λ.StrLiteral("download"),
					}))
					ϒrequest_source_file = λ.NewFunction("request_source_file",
						[]λ.Param{
							{Name: "source_url"},
							{Name: "kind"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒkind       = λargs[1]
								ϒsource_url = λargs[0]
							)
							return λ.Call(λ.GetAttr(ϒself, "_request_webpage", nil), λ.NewArgs(
								ϒsource_url,
								ϒvideo_id,
							), λ.KWArgs{
								{Name: "note", Value: λ.Mod(λ.StrLiteral("Requesting %s file"), ϒkind)},
								{Name: "errnote", Value: λ.Mod(λ.StrLiteral("Unable to request %s file"), ϒkind)},
								{Name: "fatal", Value: λ.False},
							})
						})
					ϒurlh = λ.Cal(ϒrequest_source_file, ϒsource_url, λ.StrLiteral("source"))
					if λ.IsTrue(ϒurlh) {
						ϒadd_source_format = λ.NewFunction("add_source_format",
							[]λ.Param{
								{Name: "urlh"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒurlh = λargs[0]
								)
								λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
									"url":       λ.Calm(ϒurlh, "geturl"),
									"ext":       λ.Calm(λ.Cal(ϒdetermine_ext, ϒtitle, λ.StrLiteral("mp4")), "lower"),
									"format_id": λ.StrLiteral("source"),
									"quality":   λ.IntLiteral(1),
								}))
								return λ.None
							})
						if λ.IsTrue(λ.Calm(λ.GetAttr(ϒurlh, "headers", nil), "get", λ.StrLiteral("Content-Disposition"))) {
							λ.Cal(ϒadd_source_format, ϒurlh)
						} else {
							ϒconfirmation_webpage = λ.Call(λ.GetAttr(ϒself, "_webpage_read_content", nil), λ.NewArgs(
								ϒurlh,
								ϒurl,
								ϒvideo_id,
							), λ.KWArgs{
								{Name: "note", Value: λ.StrLiteral("Downloading confirmation page")},
								{Name: "errnote", Value: λ.StrLiteral("Unable to confirm download")},
								{Name: "fatal", Value: λ.False},
							})
							if λ.IsTrue(ϒconfirmation_webpage) {
								ϒconfirm = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
									λ.StrLiteral("confirm=([^&\"\\']+)"),
									ϒconfirmation_webpage,
									λ.StrLiteral("confirmation code"),
								), λ.KWArgs{
									{Name: "default", Value: λ.None},
								})
								if λ.IsTrue(ϒconfirm) {
									ϒconfirmed_source_url = λ.Cal(ϒupdate_url_query, ϒsource_url, λ.DictLiteral(map[string]λ.Object{
										"confirm": ϒconfirm,
									}))
									ϒurlh = λ.Cal(ϒrequest_source_file, ϒconfirmed_source_url, λ.StrLiteral("confirmed source"))
									if λ.IsTrue(func() λ.Object {
										if λv := ϒurlh; !λ.IsTrue(λv) {
											return λv
										} else {
											return λ.Calm(λ.GetAttr(ϒurlh, "headers", nil), "get", λ.StrLiteral("Content-Disposition"))
										}
									}()) {
										λ.Cal(ϒadd_source_format, ϒurlh)
									}
								} else {
									λ.Calm(ϒself, "report_warning", func() λ.Object {
										if λv := λ.Cal(ϒget_element_by_class, λ.StrLiteral("uc-error-subcaption"), ϒconfirmation_webpage); λ.IsTrue(λv) {
											return λv
										} else if λv := λ.Cal(ϒget_element_by_class, λ.StrLiteral("uc-error-caption"), ϒconfirmation_webpage); λ.IsTrue(λv) {
											return λv
										} else {
											return λ.StrLiteral("unable to extract confirmation code")
										}
									}())
								}
							}
						}
					}
					if λ.IsTrue(func() λ.Object {
						if λv := λ.NewBool(!λ.IsTrue(ϒformats)); !λ.IsTrue(λv) {
							return λv
						} else {
							return ϒreason
						}
					}()) {
						panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(ϒreason), λ.KWArgs{
							{Name: "expected", Value: λ.True},
						})))
					}
					λ.Calm(ϒself, "_sort_formats", ϒformats)
					ϒhl = λ.Cal(ϒget_value, λ.StrLiteral("hl"))
					ϒsubtitles_id = λ.None
					ϒttsurl = λ.Cal(ϒget_value, λ.StrLiteral("ttsurl"))
					if λ.IsTrue(ϒttsurl) {
						ϒsubtitles_id = λ.GetItem(λ.Calm(λ.Calm(λ.Calm(ϒttsurl, "encode", λ.StrLiteral("utf-8")), "decode", λ.StrLiteral("unicode_escape")), "split", λ.StrLiteral("=")), λ.Neg(λ.IntLiteral(1)))
					}
					return λ.DictLiteral(map[string]λ.Object{
						"id":                 ϒvideo_id,
						"title":              ϒtitle,
						"thumbnail":          λ.Add(λ.StrLiteral("https://drive.google.com/thumbnail?id="), ϒvideo_id),
						"duration":           λ.Cal(ϒint_or_none, λ.Cal(ϒget_value, λ.StrLiteral("length_seconds"))),
						"formats":            ϒformats,
						"subtitles":          λ.Calm(ϒself, "extract_subtitles", ϒvideo_id, ϒsubtitles_id, ϒhl),
						"automatic_captions": λ.Calm(ϒself, "extract_automatic_captions", ϒvideo_id, ϒsubtitles_id, ϒhl),
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_FORMATS_EXT":  GoogleDriveIE__FORMATS_EXT,
				"_VALID_URL":    GoogleDriveIE__VALID_URL,
				"_real_extract": GoogleDriveIE__real_extract,
			})
		}())
	})
}
