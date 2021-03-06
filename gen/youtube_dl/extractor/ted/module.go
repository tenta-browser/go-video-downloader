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
 * ted/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/ted.py
 */

package ted

import (
	Ωjson "github.com/tenta-browser/go-video-downloader/gen/json"
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωparse "github.com/tenta-browser/go-video-downloader/gen/urllib/parse"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor       λ.Object
	TEDIE               λ.Object
	ϒcompat_str         λ.Object
	ϒextract_attributes λ.Object
	ϒfloat_or_none      λ.Object
	ϒint_or_none        λ.Object
	ϒtry_get            λ.Object
	ϒurl_or_none        λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_str = Ωcompat.ϒcompat_str
		ϒextract_attributes = Ωutils.ϒextract_attributes
		ϒfloat_or_none = Ωutils.ϒfloat_or_none
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒtry_get = Ωutils.ϒtry_get
		ϒurl_or_none = Ωutils.ϒurl_or_none
		TEDIE = λ.Cal(λ.TypeType, λ.StrLiteral("TEDIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				TEDIE_IE_NAME               λ.Object
				TEDIE__NATIVE_FORMATS       λ.Object
				TEDIE__VALID_URL            λ.Object
				TEDIE__extract_info         λ.Object
				TEDIE__get_subtitles        λ.Object
				TEDIE__playlist_videos_info λ.Object
				TEDIE__real_extract         λ.Object
				TEDIE__talk_info            λ.Object
			)
			TEDIE_IE_NAME = λ.StrLiteral("ted")
			TEDIE__VALID_URL = λ.StrLiteral("(?x)\n        (?P<proto>https?://)\n        (?P<type>www|embed(?:-ssl)?)(?P<urlmain>\\.ted\\.com/\n        (\n            (?P<type_playlist>playlists(?:/(?P<playlist_id>\\d+))?) # We have a playlist\n            |\n            ((?P<type_talk>talks)) # We have a simple talk\n            |\n            (?P<type_watch>watch)/[^/]+/[^/]+\n        )\n        (/lang/(.*?))? # The url may contain the language\n        /(?P<name>[\\w-]+) # Here goes the name and then \".html\"\n        .*)$\n        ")
			TEDIE__NATIVE_FORMATS = λ.DictLiteral(map[string]λ.Object{
				"low": λ.DictLiteral(map[string]int{
					"width":  320,
					"height": 180,
				}),
				"medium": λ.DictLiteral(map[string]int{
					"width":  512,
					"height": 288,
				}),
				"high": λ.DictLiteral(map[string]int{
					"width":  854,
					"height": 480,
				}),
			})
			TEDIE__extract_info = λ.NewFunction("_extract_info",
				[]λ.Param{
					{Name: "self"},
					{Name: "webpage"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒinfo_json λ.Object
						ϒself      = λargs[0]
						ϒwebpage   = λargs[1]
					)
					ϒinfo_json = λ.Calm(ϒself, "_search_regex", λ.StrLiteral("(?s)q\\(\\s*\"\\w+.init\"\\s*,\\s*({.+?})\\)\\s*</script>"), ϒwebpage, λ.StrLiteral("info json"))
					return λ.Cal(Ωjson.ϒloads, ϒinfo_json)
				})
			TEDIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdesktop_url λ.Object
						ϒm           λ.Object
						ϒname        λ.Object
						ϒself        = λargs[0]
						ϒurl         = λargs[1]
					)
					ϒm = λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl, Ωre.VERBOSE)
					if λ.IsTrue(λ.Calm(λ.Calm(ϒm, "group", λ.StrLiteral("type")), "startswith", λ.StrLiteral("embed"))) {
						ϒdesktop_url = λ.Add(λ.Add(λ.Calm(ϒm, "group", λ.StrLiteral("proto")), λ.StrLiteral("www")), λ.Calm(ϒm, "group", λ.StrLiteral("urlmain")))
						return λ.Calm(ϒself, "url_result", ϒdesktop_url, λ.StrLiteral("TED"))
					}
					ϒname = λ.Calm(ϒm, "group", λ.StrLiteral("name"))
					if λ.IsTrue(λ.Calm(ϒm, "group", λ.StrLiteral("type_talk"))) {
						return λ.Calm(ϒself, "_talk_info", ϒurl, ϒname)
					} else {
						if λ.IsTrue(λ.Calm(ϒm, "group", λ.StrLiteral("type_watch"))) {
							return λ.Calm(ϒself, "_watch_info", ϒurl, ϒname)
						} else {
							return λ.Calm(ϒself, "_playlist_videos_info", ϒurl, ϒname)
						}
					}
					return λ.None
				})
			TEDIE__playlist_videos_info = λ.NewFunction("_playlist_videos_info",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
					{Name: "name"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒattrs            λ.Object
						ϒentry            λ.Object
						ϒentry_url        λ.Object
						ϒfinal_url        λ.Object
						ϒname             = λargs[2]
						ϒplaylist_entries λ.Object
						ϒplaylist_id      λ.Object
						ϒself             = λargs[0]
						ϒurl              = λargs[1]
						ϒwebpage          λ.Object
						τmp0              λ.Object
						τmp1              λ.Object
					)
					ϒwebpage = λ.Calm(ϒself, "_download_webpage", ϒurl, ϒname, λ.StrLiteral("Downloading playlist webpage"))
					ϒplaylist_entries = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(Ωre.ϒfindall, λ.StrLiteral("(?s)<[^>]+data-ga-context=[\"\\']playlist[\"\\'][^>]*>"), ϒwebpage))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒentry = τmp1
						ϒattrs = λ.Cal(ϒextract_attributes, ϒentry)
						ϒentry_url = λ.Cal(Ωparse.ϒurljoin, ϒurl, λ.GetItem(ϒattrs, λ.StrLiteral("href")))
						λ.Calm(ϒplaylist_entries, "append", λ.Calm(ϒself, "url_result", ϒentry_url, λ.Calm(ϒself, "ie_key")))
					}
					ϒfinal_url = λ.Call(λ.GetAttr(ϒself, "_og_search_url", nil), λ.NewArgs(ϒwebpage), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					})
					ϒplaylist_id = func() λ.Object {
						if λ.IsTrue(ϒfinal_url) {
							return λ.Calm(λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒfinal_url), "group", λ.StrLiteral("playlist_id"))
						} else {
							return λ.None
						}
					}()
					return λ.Call(λ.GetAttr(ϒself, "playlist_result", nil), λ.NewArgs(ϒplaylist_entries), λ.KWArgs{
						{Name: "playlist_id", Value: ϒplaylist_id},
						{Name: "playlist_title", Value: λ.Call(λ.GetAttr(ϒself, "_og_search_title", nil), λ.NewArgs(ϒwebpage), λ.KWArgs{
							{Name: "fatal", Value: λ.False},
						})},
						{Name: "playlist_description", Value: λ.Calm(ϒself, "_og_search_description", ϒwebpage)},
					})
				})
			TEDIE__talk_info = λ.NewFunction("_talk_info",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
					{Name: "video_name"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒaudio_download      λ.Object
						ϒbitrate             λ.Object
						ϒbitrate_url         λ.Object
						ϒdata                λ.Object
						ϒdownloads           λ.Object
						ϒext_url             λ.Object
						ϒexternal            λ.Object
						ϒf                   λ.Object
						ϒfinfo               λ.Object
						ϒformat_id           λ.Object
						ϒformats             λ.Object
						ϒh264_url            λ.Object
						ϒhttp_url            λ.Object
						ϒinfo                λ.Object
						ϒlang                λ.Object
						ϒm3u8_format         λ.Object
						ϒm3u8_formats        λ.Object
						ϒnative_downloads    λ.Object
						ϒplayer_talk         λ.Object
						ϒq                   λ.Object
						ϒq_url               λ.Object
						ϒresource            λ.Object
						ϒresources           λ.Object
						ϒresources_          λ.Object
						ϒself                = λargs[0]
						ϒservice             λ.Object
						ϒstream_url          λ.Object
						ϒstreamer            λ.Object
						ϒsubtitled_download  λ.Object
						ϒsubtitled_downloads λ.Object
						ϒtalk_info           λ.Object
						ϒtitle               λ.Object
						ϒurl                 = λargs[1]
						ϒvideo_id            λ.Object
						ϒvideo_name          = λargs[2]
						ϒwebpage             λ.Object
						τmp0                 λ.Object
						τmp1                 λ.Object
						τmp2                 λ.Object
						τmp3                 λ.Object
					)
					ϒwebpage = λ.Calm(ϒself, "_download_webpage", ϒurl, ϒvideo_name)
					ϒinfo = λ.Calm(ϒself, "_extract_info", ϒwebpage)
					ϒdata = func() λ.Object {
						if λv := λ.Cal(ϒtry_get, ϒinfo, λ.NewFunction("<lambda>",
							[]λ.Param{
								{Name: "x"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒx = λargs[0]
								)
								return λ.GetItem(ϒx, λ.StrLiteral("__INITIAL_DATA__"))
							}), λ.DictType); λ.IsTrue(λv) {
							return λv
						} else {
							return ϒinfo
						}
					}()
					ϒtalk_info = λ.GetItem(λ.GetItem(ϒdata, λ.StrLiteral("talks")), λ.IntLiteral(0))
					ϒtitle = λ.Calm(λ.GetItem(ϒtalk_info, λ.StrLiteral("title")), "strip")
					ϒdownloads = func() λ.Object {
						if λv := λ.Calm(ϒtalk_info, "get", λ.StrLiteral("downloads")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.DictLiteral(map[λ.Object]λ.Object{})
						}
					}()
					ϒnative_downloads = func() λ.Object {
						if λv := λ.Calm(ϒdownloads, "get", λ.StrLiteral("nativeDownloads")); λ.IsTrue(λv) {
							return λv
						} else if λv := λ.Calm(ϒtalk_info, "get", λ.StrLiteral("nativeDownloads")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.DictLiteral(map[λ.Object]λ.Object{})
						}
					}()
					ϒformats = λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
						nil,
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							return λ.NewGenerator(func(λgy λ.Yielder) λ.Object {
								var (
									ϒformat_id  λ.Object
									ϒformat_url λ.Object
									τmp0        λ.Object
									τmp1        λ.Object
									τmp2        λ.Object
								)
								τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(ϒnative_downloads, "items"))
								for {
									if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
										break
									}
									τmp2 = λ.UnpackIterable(τmp1, 2)
									ϒformat_id = λ.GetItem(τmp2, λ.IntLiteral(0))
									ϒformat_url = λ.GetItem(τmp2, λ.IntLiteral(1))
									if ϒformat_url != λ.None {
										λgy.Yield(λ.DictLiteral(map[string]λ.Object{
											"url":       ϒformat_url,
											"format_id": ϒformat_id,
										}))
									}
								}
								return λ.None
							})
						})))
					ϒsubtitled_downloads = func() λ.Object {
						if λv := λ.Calm(ϒdownloads, "get", λ.StrLiteral("subtitledDownloads")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.DictLiteral(map[λ.Object]λ.Object{})
						}
					}()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(ϒsubtitled_downloads, "items"))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						τmp2 = λ.UnpackIterable(τmp1, 2)
						ϒlang = λ.GetItem(τmp2, λ.IntLiteral(0))
						ϒsubtitled_download = λ.GetItem(τmp2, λ.IntLiteral(1))
						τmp2 = λ.Cal(λ.BuiltinIter, λ.GetAttr(ϒself, "_NATIVE_FORMATS", nil))
						for {
							if τmp3 = λ.NextDefault(τmp2, λ.AfterLast); τmp3 == λ.AfterLast {
								break
							}
							ϒq = τmp3
							ϒq_url = λ.Calm(ϒsubtitled_download, "get", ϒq)
							if !λ.IsTrue(ϒq_url) {
								continue
							}
							λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
								"url": ϒq_url,
								"format_id": λ.Mod(λ.StrLiteral("%s-%s"), λ.NewTuple(
									ϒq,
									ϒlang,
								)),
								"language": ϒlang,
							}))
						}
					}
					if λ.IsTrue(ϒformats) {
						τmp0 = λ.Cal(λ.BuiltinIter, ϒformats)
						for {
							if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
								break
							}
							ϒf = τmp1
							ϒfinfo = λ.Calm(λ.GetAttr(ϒself, "_NATIVE_FORMATS", nil), "get", λ.GetItem(λ.Calm(λ.GetItem(ϒf, λ.StrLiteral("format_id")), "split", λ.StrLiteral("-")), λ.IntLiteral(0)))
							if λ.IsTrue(ϒfinfo) {
								λ.Calm(ϒf, "update", ϒfinfo)
							}
						}
					}
					ϒplayer_talk = λ.GetItem(λ.GetItem(ϒtalk_info, λ.StrLiteral("player_talks")), λ.IntLiteral(0))
					ϒresources_ = func() λ.Object {
						if λv := λ.Calm(ϒplayer_talk, "get", λ.StrLiteral("resources")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Calm(ϒtalk_info, "get", λ.StrLiteral("resources"))
						}
					}()
					ϒhttp_url = λ.None
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(ϒresources_, "items"))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						τmp2 = λ.UnpackIterable(τmp1, 2)
						ϒformat_id = λ.GetItem(τmp2, λ.IntLiteral(0))
						ϒresources = λ.GetItem(τmp2, λ.IntLiteral(1))
						if λ.IsTrue(λ.Eq(ϒformat_id, λ.StrLiteral("hls"))) {
							if !λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒresources, λ.DictType)) {
								continue
							}
							ϒstream_url = λ.Cal(ϒurl_or_none, λ.Calm(ϒresources, "get", λ.StrLiteral("stream")))
							if !λ.IsTrue(ϒstream_url) {
								continue
							}
							λ.Calm(ϒformats, "extend", λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
								ϒstream_url,
								ϒvideo_name,
								λ.StrLiteral("mp4"),
							), λ.KWArgs{
								{Name: "m3u8_id", Value: ϒformat_id},
								{Name: "fatal", Value: λ.False},
							}))
						} else {
							if !λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒresources, λ.ListType)) {
								continue
							}
							if λ.IsTrue(λ.Eq(ϒformat_id, λ.StrLiteral("h264"))) {
								τmp2 = λ.Cal(λ.BuiltinIter, ϒresources)
								for {
									if τmp3 = λ.NextDefault(τmp2, λ.AfterLast); τmp3 == λ.AfterLast {
										break
									}
									ϒresource = τmp3
									ϒh264_url = λ.Calm(ϒresource, "get", λ.StrLiteral("file"))
									if !λ.IsTrue(ϒh264_url) {
										continue
									}
									ϒbitrate = λ.Cal(ϒint_or_none, λ.Calm(ϒresource, "get", λ.StrLiteral("bitrate")))
									λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
										"url": ϒh264_url,
										"format_id": λ.Mod(λ.StrLiteral("%s-%sk"), λ.NewTuple(
											ϒformat_id,
											ϒbitrate,
										)),
										"tbr": ϒbitrate,
									}))
									if λ.IsTrue(λ.Cal(Ωre.ϒsearch, λ.StrLiteral("\\d+k"), ϒh264_url)) {
										ϒhttp_url = ϒh264_url
									}
								}
							} else {
								if λ.IsTrue(λ.Eq(ϒformat_id, λ.StrLiteral("rtmp"))) {
									ϒstreamer = λ.Calm(ϒtalk_info, "get", λ.StrLiteral("streamer"))
									if !λ.IsTrue(ϒstreamer) {
										continue
									}
									τmp2 = λ.Cal(λ.BuiltinIter, ϒresources)
									for {
										if τmp3 = λ.NextDefault(τmp2, λ.AfterLast); τmp3 == λ.AfterLast {
											break
										}
										ϒresource = τmp3
										λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
											"format_id": λ.Mod(λ.StrLiteral("%s-%s"), λ.NewTuple(
												ϒformat_id,
												λ.Calm(ϒresource, "get", λ.StrLiteral("name")),
											)),
											"url":       ϒstreamer,
											"play_path": λ.GetItem(ϒresource, λ.StrLiteral("file")),
											"ext":       λ.StrLiteral("flv"),
											"width":     λ.Cal(ϒint_or_none, λ.Calm(ϒresource, "get", λ.StrLiteral("width"))),
											"height":    λ.Cal(ϒint_or_none, λ.Calm(ϒresource, "get", λ.StrLiteral("height"))),
											"tbr":       λ.Cal(ϒint_or_none, λ.Calm(ϒresource, "get", λ.StrLiteral("bitrate"))),
										}))
									}
								}
							}
						}
					}
					ϒm3u8_formats = λ.Cal(λ.ListType, λ.Cal(λ.FilterIteratorType, λ.NewFunction("<lambda>",
						[]λ.Param{
							{Name: "f"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒf = λargs[0]
							)
							return func() λ.Object {
								if λv := λ.Eq(λ.Calm(ϒf, "get", λ.StrLiteral("protocol")), λ.StrLiteral("m3u8")); !λ.IsTrue(λv) {
									return λv
								} else {
									return λ.Ne(λ.Calm(ϒf, "get", λ.StrLiteral("vcodec")), λ.StrLiteral("none"))
								}
							}()
						}), ϒformats))
					if λ.IsTrue(ϒhttp_url) {
						τmp0 = λ.Cal(λ.BuiltinIter, ϒm3u8_formats)
						for {
							if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
								break
							}
							ϒm3u8_format = τmp1
							ϒbitrate = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
								λ.StrLiteral("(\\d+k)"),
								λ.GetItem(ϒm3u8_format, λ.StrLiteral("url")),
								λ.StrLiteral("bitrate"),
							), λ.KWArgs{
								{Name: "default", Value: λ.None},
							})
							if !λ.IsTrue(ϒbitrate) {
								continue
							}
							ϒbitrate_url = λ.Cal(Ωre.ϒsub, λ.StrLiteral("\\d+k"), ϒbitrate, ϒhttp_url)
							if !λ.IsTrue(λ.Calm(ϒself, "_is_valid_url", ϒbitrate_url, ϒvideo_name, λ.Mod(λ.StrLiteral("%s bitrate"), ϒbitrate))) {
								continue
							}
							ϒf = λ.Calm(ϒm3u8_format, "copy")
							λ.Calm(ϒf, "update", λ.DictLiteral(map[string]λ.Object{
								"url":       ϒbitrate_url,
								"format_id": λ.Calm(λ.GetItem(ϒm3u8_format, λ.StrLiteral("format_id")), "replace", λ.StrLiteral("hls"), λ.StrLiteral("http")),
								"protocol":  λ.StrLiteral("http"),
							}))
							if λ.IsTrue(λ.Eq(λ.Calm(ϒf, "get", λ.StrLiteral("acodec")), λ.StrLiteral("none"))) {
								λ.DelItem(ϒf, λ.StrLiteral("acodec"))
							}
							λ.Calm(ϒformats, "append", ϒf)
						}
					}
					ϒaudio_download = λ.Calm(ϒtalk_info, "get", λ.StrLiteral("audioDownload"))
					if λ.IsTrue(ϒaudio_download) {
						λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
							"url":       ϒaudio_download,
							"format_id": λ.StrLiteral("audio"),
							"vcodec":    λ.StrLiteral("none"),
						}))
					}
					if !λ.IsTrue(ϒformats) {
						ϒexternal = λ.Calm(ϒplayer_talk, "get", λ.StrLiteral("external"))
						if λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒexternal, λ.DictType)) {
							ϒservice = λ.Calm(ϒexternal, "get", λ.StrLiteral("service"))
							if λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒservice, ϒcompat_str)) {
								ϒext_url = λ.None
								if λ.IsTrue(λ.Eq(λ.Calm(ϒservice, "lower"), λ.StrLiteral("youtube"))) {
									ϒext_url = λ.Calm(ϒexternal, "get", λ.StrLiteral("code"))
								}
								return λ.Calm(ϒself, "url_result", func() λ.Object {
									if λv := ϒext_url; λ.IsTrue(λv) {
										return λv
									} else {
										return λ.GetItem(ϒexternal, λ.StrLiteral("uri"))
									}
								}())
							}
						}
					}
					λ.Calm(ϒself, "_sort_formats", ϒformats)
					ϒvideo_id = λ.Cal(ϒcompat_str, λ.GetItem(ϒtalk_info, λ.StrLiteral("id")))
					return λ.DictLiteral(map[string]λ.Object{
						"id":    ϒvideo_id,
						"title": ϒtitle,
						"uploader": func() λ.Object {
							if λv := λ.Calm(ϒplayer_talk, "get", λ.StrLiteral("speaker")); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Calm(ϒtalk_info, "get", λ.StrLiteral("speaker"))
							}
						}(),
						"thumbnail": func() λ.Object {
							if λv := λ.Calm(ϒplayer_talk, "get", λ.StrLiteral("thumb")); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Calm(ϒtalk_info, "get", λ.StrLiteral("thumb"))
							}
						}(),
						"description": λ.Calm(ϒself, "_og_search_description", ϒwebpage),
						"subtitles":   λ.Calm(ϒself, "_get_subtitles", ϒvideo_id, ϒtalk_info),
						"formats":     ϒformats,
						"duration":    λ.Cal(ϒfloat_or_none, λ.Calm(ϒtalk_info, "get", λ.StrLiteral("duration"))),
						"view_count":  λ.Cal(ϒint_or_none, λ.Calm(ϒdata, "get", λ.StrLiteral("viewed_count"))),
						"comment_count": λ.Cal(ϒint_or_none, λ.Cal(ϒtry_get, ϒdata, λ.NewFunction("<lambda>",
							[]λ.Param{
								{Name: "x"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒx = λargs[0]
								)
								return λ.GetItem(λ.GetItem(ϒx, λ.StrLiteral("comments")), λ.StrLiteral("count"))
							}))),
						"tags": λ.Cal(ϒtry_get, ϒtalk_info, λ.NewFunction("<lambda>",
							[]λ.Param{
								{Name: "x"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒx = λargs[0]
								)
								return λ.GetItem(ϒx, λ.StrLiteral("tags"))
							}), λ.ListType),
					})
				})
			TEDIE__get_subtitles = λ.NewFunction("_get_subtitles",
				[]λ.Param{
					{Name: "self"},
					{Name: "video_id"},
					{Name: "talk_info"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒlang_code     λ.Object
						ϒlanguage      λ.Object
						ϒself          = λargs[0]
						ϒsub_lang_list λ.Object
						ϒtalk_info     = λargs[2]
						ϒvideo_id      = λargs[1]
						τmp0           λ.Object
						τmp1           λ.Object
					)
					_ = ϒself
					ϒsub_lang_list = λ.DictLiteral(map[λ.Object]λ.Object{})
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(ϒtry_get, ϒtalk_info, λ.NewTuple(
						λ.NewFunction("<lambda>",
							[]λ.Param{
								{Name: "x"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒx = λargs[0]
								)
								return λ.GetItem(λ.GetItem(ϒx, λ.StrLiteral("downloads")), λ.StrLiteral("languages"))
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
								return λ.GetItem(ϒx, λ.StrLiteral("languages"))
							}),
					), λ.ListType))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒlanguage = τmp1
						ϒlang_code = func() λ.Object {
							if λv := λ.Calm(ϒlanguage, "get", λ.StrLiteral("languageCode")); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Calm(ϒlanguage, "get", λ.StrLiteral("ianaCode"))
							}
						}()
						if !λ.IsTrue(ϒlang_code) {
							continue
						}
						λ.SetItem(ϒsub_lang_list, ϒlang_code, λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
							nil,
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								return λ.NewGenerator(func(λgy λ.Yielder) λ.Object {
									var (
										ϒext λ.Object
										τmp0 λ.Object
										τmp1 λ.Object
									)
									τmp0 = λ.Cal(λ.BuiltinIter, λ.NewList(
										λ.StrLiteral("ted"),
										λ.StrLiteral("srt"),
									))
									for {
										if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
											break
										}
										ϒext = τmp1
										λgy.Yield(λ.DictLiteral(map[string]λ.Object{
											"url": λ.Mod(λ.StrLiteral("http://www.ted.com/talks/subtitles/id/%s/lang/%s/format/%s"), λ.NewTuple(
												ϒvideo_id,
												ϒlang_code,
												ϒext,
											)),
											"ext": ϒext,
										}))
									}
									return λ.None
								})
							}))))
					}
					return ϒsub_lang_list
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"IE_NAME":               TEDIE_IE_NAME,
				"_NATIVE_FORMATS":       TEDIE__NATIVE_FORMATS,
				"_VALID_URL":            TEDIE__VALID_URL,
				"_extract_info":         TEDIE__extract_info,
				"_get_subtitles":        TEDIE__get_subtitles,
				"_playlist_videos_info": TEDIE__playlist_videos_info,
				"_real_extract":         TEDIE__real_extract,
				"_talk_info":            TEDIE__talk_info,
			})
		}())
	})
}
