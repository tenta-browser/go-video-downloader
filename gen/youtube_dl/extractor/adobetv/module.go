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
 * adobetv/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/adobetv.py
 */

package adobetv

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	AdobeTVBaseIE         λ.Object
	AdobeTVChannelIE      λ.Object
	AdobeTVEmbedIE        λ.Object
	AdobeTVIE             λ.Object
	AdobeTVPlaylistBaseIE λ.Object
	AdobeTVShowIE         λ.Object
	AdobeTVVideoIE        λ.Object
	InfoExtractor         λ.Object
	ϒcompat_str           λ.Object
	ϒfloat_or_none        λ.Object
	ϒint_or_none          λ.Object
	ϒparse_duration       λ.Object
	ϒstr_or_none          λ.Object
	ϒstr_to_int           λ.Object
	ϒunified_strdate      λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_str = Ωcompat.ϒcompat_str
		ϒfloat_or_none = Ωutils.ϒfloat_or_none
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒparse_duration = Ωutils.ϒparse_duration
		ϒstr_or_none = Ωutils.ϒstr_or_none
		ϒstr_to_int = Ωutils.ϒstr_to_int
		ϒunified_strdate = Ωutils.ϒunified_strdate
		AdobeTVBaseIE = λ.Cal(λ.TypeType, λ.StrLiteral("AdobeTVBaseIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				AdobeTVBaseIE__parse_subtitles  λ.Object
				AdobeTVBaseIE__parse_video_data λ.Object
			)
			AdobeTVBaseIE__parse_subtitles = λ.NewFunction("_parse_subtitles",
				[]λ.Param{
					{Name: "self"},
					{Name: "video_data"},
					{Name: "url_key"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒlang        λ.Object
						ϒself        = λargs[0]
						ϒsubtitles   λ.Object
						ϒtranslation λ.Object
						ϒurl_key     = λargs[2]
						ϒvideo_data  = λargs[1]
						ϒvtt_path    λ.Object
						τmp0         λ.Object
						τmp1         λ.Object
					)
					_ = ϒself
					ϒsubtitles = λ.DictLiteral(map[λ.Object]λ.Object{})
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(ϒvideo_data, "get", λ.StrLiteral("translations"), λ.NewList()))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒtranslation = τmp1
						ϒvtt_path = λ.Calm(ϒtranslation, "get", ϒurl_key)
						if !λ.IsTrue(ϒvtt_path) {
							continue
						}
						ϒlang = func() λ.Object {
							if λv := λ.Calm(ϒtranslation, "get", λ.StrLiteral("language_w3c")); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Calm(λ.None, "long2short", λ.GetItem(ϒtranslation, λ.StrLiteral("language_medium")))
							}
						}()
						λ.Calm(λ.Calm(ϒsubtitles, "setdefault", ϒlang, λ.NewList()), "append", λ.DictLiteral(map[string]λ.Object{
							"ext": λ.StrLiteral("vtt"),
							"url": ϒvtt_path,
						}))
					}
					return ϒsubtitles
				})
			AdobeTVBaseIE__parse_video_data = λ.NewFunction("_parse_video_data",
				[]λ.Param{
					{Name: "self"},
					{Name: "video_data"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒf                 λ.Object
						ϒformats           λ.Object
						ϒmobj              λ.Object
						ϒoriginal_filename λ.Object
						ϒs3_extracted      λ.Object
						ϒself              = λargs[0]
						ϒsource            λ.Object
						ϒsource_url        λ.Object
						ϒtitle             λ.Object
						ϒvideo_data        = λargs[1]
						ϒvideo_id          λ.Object
						τmp0               λ.Object
						τmp1               λ.Object
					)
					ϒvideo_id = λ.Cal(ϒcompat_str, λ.GetItem(ϒvideo_data, λ.StrLiteral("id")))
					ϒtitle = λ.GetItem(ϒvideo_data, λ.StrLiteral("title"))
					ϒs3_extracted = λ.False
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(ϒvideo_data, "get", λ.StrLiteral("videos"), λ.NewList()))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒsource = τmp1
						ϒsource_url = λ.Calm(ϒsource, "get", λ.StrLiteral("url"))
						if !λ.IsTrue(ϒsource_url) {
							continue
						}
						ϒf = λ.DictLiteral(map[string]λ.Object{
							"format_id": λ.Calm(ϒsource, "get", λ.StrLiteral("quality_level")),
							"fps":       λ.Cal(ϒint_or_none, λ.Calm(ϒsource, "get", λ.StrLiteral("frame_rate"))),
							"height":    λ.Cal(ϒint_or_none, λ.Calm(ϒsource, "get", λ.StrLiteral("height"))),
							"tbr":       λ.Cal(ϒint_or_none, λ.Calm(ϒsource, "get", λ.StrLiteral("video_data_rate"))),
							"width":     λ.Cal(ϒint_or_none, λ.Calm(ϒsource, "get", λ.StrLiteral("width"))),
							"url":       ϒsource_url,
						})
						ϒoriginal_filename = λ.Calm(ϒsource, "get", λ.StrLiteral("original_filename"))
						if λ.IsTrue(ϒoriginal_filename) {
							if !λ.IsTrue(func() λ.Object {
								if λv := λ.Calm(ϒf, "get", λ.StrLiteral("height")); !λ.IsTrue(λv) {
									return λv
								} else {
									return λ.Calm(ϒf, "get", λ.StrLiteral("width"))
								}
							}()) {
								ϒmobj = λ.Cal(Ωre.ϒsearch, λ.StrLiteral("_(\\d+)x(\\d+)"), ϒoriginal_filename)
								if λ.IsTrue(ϒmobj) {
									λ.Calm(ϒf, "update", λ.DictLiteral(map[string]λ.Object{
										"height": λ.Cal(λ.IntType, λ.Calm(ϒmobj, "group", λ.IntLiteral(2))),
										"width":  λ.Cal(λ.IntType, λ.Calm(ϒmobj, "group", λ.IntLiteral(1))),
									}))
								}
							}
							if λ.IsTrue(func() λ.Object {
								if λv := λ.Calm(ϒoriginal_filename, "startswith", λ.StrLiteral("s3://")); !λ.IsTrue(λv) {
									return λv
								} else {
									return λ.NewBool(!λ.IsTrue(ϒs3_extracted))
								}
							}()) {
								λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
									"format_id":  λ.StrLiteral("original"),
									"preference": λ.IntLiteral(1),
									"url":        λ.Calm(ϒoriginal_filename, "replace", λ.StrLiteral("s3://"), λ.StrLiteral("https://s3.amazonaws.com/")),
								}))
								ϒs3_extracted = λ.True
							}
						}
						λ.Calm(ϒformats, "append", ϒf)
					}
					λ.Calm(ϒself, "_sort_formats", ϒformats)
					return λ.DictLiteral(map[string]λ.Object{
						"id":          ϒvideo_id,
						"title":       ϒtitle,
						"description": λ.Calm(ϒvideo_data, "get", λ.StrLiteral("description")),
						"thumbnail":   λ.Calm(ϒvideo_data, "get", λ.StrLiteral("thumbnail")),
						"upload_date": λ.Cal(ϒunified_strdate, λ.Calm(ϒvideo_data, "get", λ.StrLiteral("start_date"))),
						"duration":    λ.Cal(ϒparse_duration, λ.Calm(ϒvideo_data, "get", λ.StrLiteral("duration"))),
						"view_count":  λ.Cal(ϒstr_to_int, λ.Calm(ϒvideo_data, "get", λ.StrLiteral("playcount"))),
						"formats":     ϒformats,
						"subtitles":   λ.Calm(ϒself, "_parse_subtitles", ϒvideo_data, λ.StrLiteral("vtt")),
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_parse_subtitles":  AdobeTVBaseIE__parse_subtitles,
				"_parse_video_data": AdobeTVBaseIE__parse_video_data,
			})
		}())
		AdobeTVEmbedIE = λ.Cal(λ.TypeType, λ.StrLiteral("AdobeTVEmbedIE"), λ.NewTuple(AdobeTVBaseIE), func() λ.Dict {
			var (
				AdobeTVEmbedIE__VALID_URL λ.Object
			)
			AdobeTVEmbedIE__VALID_URL = λ.StrLiteral("https?://tv\\.adobe\\.com/embed/\\d+/(?P<id>\\d+)")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": AdobeTVEmbedIE__VALID_URL,
			})
		}())
		AdobeTVIE = λ.Cal(λ.TypeType, λ.StrLiteral("AdobeTVIE"), λ.NewTuple(AdobeTVBaseIE), func() λ.Dict {
			var (
				AdobeTVIE__VALID_URL λ.Object
			)
			AdobeTVIE__VALID_URL = λ.StrLiteral("https?://tv\\.adobe\\.com/(?:(?P<language>fr|de|es|jp)/)?watch/(?P<show_urlname>[^/]+)/(?P<id>[^/]+)")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": AdobeTVIE__VALID_URL,
			})
		}())
		AdobeTVPlaylistBaseIE = λ.Cal(λ.TypeType, λ.StrLiteral("AdobeTVPlaylistBaseIE"), λ.NewTuple(AdobeTVBaseIE), func() λ.Dict {

			return λ.ClassDictLiteral(map[λ.Object]λ.Object{})
		}())
		AdobeTVShowIE = λ.Cal(λ.TypeType, λ.StrLiteral("AdobeTVShowIE"), λ.NewTuple(AdobeTVPlaylistBaseIE), func() λ.Dict {
			var (
				AdobeTVShowIE__VALID_URL λ.Object
			)
			AdobeTVShowIE__VALID_URL = λ.StrLiteral("https?://tv\\.adobe\\.com/(?:(?P<language>fr|de|es|jp)/)?show/(?P<id>[^/]+)")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": AdobeTVShowIE__VALID_URL,
			})
		}())
		AdobeTVChannelIE = λ.Cal(λ.TypeType, λ.StrLiteral("AdobeTVChannelIE"), λ.NewTuple(AdobeTVPlaylistBaseIE), func() λ.Dict {
			var (
				AdobeTVChannelIE__VALID_URL λ.Object
			)
			AdobeTVChannelIE__VALID_URL = λ.StrLiteral("https?://tv\\.adobe\\.com/(?:(?P<language>fr|de|es|jp)/)?channel/(?P<id>[^/]+)(?:/(?P<category_urlname>[^/]+))?")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": AdobeTVChannelIE__VALID_URL,
			})
		}())
		AdobeTVVideoIE = λ.Cal(λ.TypeType, λ.StrLiteral("AdobeTVVideoIE"), λ.NewTuple(AdobeTVBaseIE), func() λ.Dict {
			var (
				AdobeTVVideoIE_IE_NAME       λ.Object
				AdobeTVVideoIE__VALID_URL    λ.Object
				AdobeTVVideoIE__real_extract λ.Object
			)
			AdobeTVVideoIE_IE_NAME = λ.StrLiteral("adobetv:video")
			AdobeTVVideoIE__VALID_URL = λ.StrLiteral("https?://video\\.tv\\.adobe\\.com/v/(?P<id>\\d+)")
			AdobeTVVideoIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒduration   λ.Object
						ϒformats    λ.Object
						ϒself       = λargs[0]
						ϒsource     λ.Object
						ϒsource_src λ.Object
						ϒsources    λ.Object
						ϒtitle      λ.Object
						ϒurl        = λargs[1]
						ϒvideo_data λ.Object
						ϒvideo_id   λ.Object
						ϒwebpage    λ.Object
						τmp0        λ.Object
						τmp1        λ.Object
					)
					ϒvideo_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒwebpage = λ.Calm(ϒself, "_download_webpage", ϒurl, ϒvideo_id)
					ϒvideo_data = λ.Calm(ϒself, "_parse_json", λ.Calm(ϒself, "_search_regex", λ.StrLiteral("var\\s+bridge\\s*=\\s*([^;]+);"), ϒwebpage, λ.StrLiteral("bridged data")), ϒvideo_id)
					ϒtitle = λ.GetItem(ϒvideo_data, λ.StrLiteral("title"))
					ϒformats = λ.NewList()
					ϒsources = func() λ.Object {
						if λv := λ.Calm(ϒvideo_data, "get", λ.StrLiteral("sources")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.NewList()
						}
					}()
					τmp0 = λ.Cal(λ.BuiltinIter, ϒsources)
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒsource = τmp1
						ϒsource_src = λ.Calm(ϒsource, "get", λ.StrLiteral("src"))
						if !λ.IsTrue(ϒsource_src) {
							continue
						}
						λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
							"filesize": λ.Call(ϒint_or_none, λ.NewArgs(func() λ.Object {
								if λv := λ.Calm(ϒsource, "get", λ.StrLiteral("kilobytes")); λ.IsTrue(λv) {
									return λv
								} else {
									return λ.None
								}
							}()), λ.KWArgs{
								{Name: "invscale", Value: λ.IntLiteral(1000)},
							}),
							"format_id": λ.Calm(λ.StrLiteral("-"), "join", λ.Cal(λ.FilterIteratorType, λ.None, λ.NewList(
								λ.Calm(ϒsource, "get", λ.StrLiteral("format")),
								λ.Calm(ϒsource, "get", λ.StrLiteral("label")),
							))),
							"height": λ.Cal(ϒint_or_none, func() λ.Object {
								if λv := λ.Calm(ϒsource, "get", λ.StrLiteral("height")); λ.IsTrue(λv) {
									return λv
								} else {
									return λ.None
								}
							}()),
							"tbr": λ.Cal(ϒint_or_none, func() λ.Object {
								if λv := λ.Calm(ϒsource, "get", λ.StrLiteral("bitrate")); λ.IsTrue(λv) {
									return λv
								} else {
									return λ.None
								}
							}()),
							"width": λ.Cal(ϒint_or_none, func() λ.Object {
								if λv := λ.Calm(ϒsource, "get", λ.StrLiteral("width")); λ.IsTrue(λv) {
									return λv
								} else {
									return λ.None
								}
							}()),
							"url": ϒsource_src,
						}))
					}
					λ.Calm(ϒself, "_sort_formats", ϒformats)
					ϒduration = λ.Cal(λ.BuiltinMax, λ.Cal(λ.FilterIteratorType, λ.None, λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
						nil,
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							return λ.NewGenerator(func(λgy λ.Yielder) λ.Object {
								var (
									ϒsource λ.Object
									τmp0    λ.Object
									τmp1    λ.Object
								)
								τmp0 = λ.Cal(λ.BuiltinIter, ϒsources)
								for {
									if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
										break
									}
									ϒsource = τmp1
									λgy.Yield(λ.Call(ϒfloat_or_none, λ.NewArgs(λ.Calm(ϒsource, "get", λ.StrLiteral("duration"))), λ.KWArgs{
										{Name: "scale", Value: λ.IntLiteral(1000)},
									}))
								}
								return λ.None
							})
						})))))
					return λ.DictLiteral(map[string]λ.Object{
						"id":          ϒvideo_id,
						"formats":     ϒformats,
						"title":       ϒtitle,
						"description": λ.Calm(ϒvideo_data, "get", λ.StrLiteral("description")),
						"thumbnail":   λ.Calm(λ.Calm(ϒvideo_data, "get", λ.StrLiteral("video"), λ.DictLiteral(map[λ.Object]λ.Object{})), "get", λ.StrLiteral("poster")),
						"duration":    ϒduration,
						"subtitles":   λ.Calm(ϒself, "_parse_subtitles", ϒvideo_data, λ.StrLiteral("vttPath")),
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"IE_NAME":       AdobeTVVideoIE_IE_NAME,
				"_VALID_URL":    AdobeTVVideoIE__VALID_URL,
				"_real_extract": AdobeTVVideoIE__real_extract,
			})
		}())
	})
}
