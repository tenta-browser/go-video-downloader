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
 * amp/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/amp.py
 */

package amp

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	AMPIE              λ.Object
	ExtractorError     λ.Object
	InfoExtractor      λ.Object
	ϒdetermine_ext     λ.Object
	ϒint_or_none       λ.Object
	ϒmimetype2ext      λ.Object
	ϒparse_iso8601     λ.Object
	ϒunified_timestamp λ.Object
	ϒurl_or_none       λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒdetermine_ext = Ωutils.ϒdetermine_ext
		ExtractorError = Ωutils.ExtractorError
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒmimetype2ext = Ωutils.ϒmimetype2ext
		ϒparse_iso8601 = Ωutils.ϒparse_iso8601
		ϒunified_timestamp = Ωutils.ϒunified_timestamp
		ϒurl_or_none = Ωutils.ϒurl_or_none
		AMPIE = λ.Cal(λ.TypeType, λ.StrLiteral("AMPIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				AMPIE__extract_feed_info λ.Object
			)
			AMPIE__extract_feed_info = λ.NewFunction("_extract_feed_info",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒext             λ.Object
						ϒfeed            λ.Object
						ϒformats         λ.Object
						ϒget_media_node  λ.Object
						ϒitem            λ.Object
						ϒmedia           λ.Object
						ϒmedia_content   λ.Object
						ϒmedia_data      λ.Object
						ϒmedia_subtitle  λ.Object
						ϒmedia_thumbnail λ.Object
						ϒmedia_url       λ.Object
						ϒself            = λargs[0]
						ϒsubtitle        λ.Object
						ϒsubtitle_data   λ.Object
						ϒsubtitle_href   λ.Object
						ϒsubtitles       λ.Object
						ϒthumbnail       λ.Object
						ϒthumbnail_data  λ.Object
						ϒthumbnail_url   λ.Object
						ϒthumbnails      λ.Object
						ϒtimestamp       λ.Object
						ϒurl             = λargs[1]
						ϒvideo_id        λ.Object
						τmp0             λ.Object
						τmp1             λ.Object
					)
					ϒfeed = λ.Calm(ϒself, "_download_json", ϒurl, λ.None, λ.StrLiteral("Downloading Akamai AMP feed"), λ.StrLiteral("Unable to download Akamai AMP feed"))
					ϒitem = λ.Calm(λ.Calm(ϒfeed, "get", λ.StrLiteral("channel"), λ.DictLiteral(map[λ.Object]λ.Object{})), "get", λ.StrLiteral("item"))
					if !λ.IsTrue(ϒitem) {
						panic(λ.Raise(λ.Cal(ExtractorError, λ.Mod(λ.StrLiteral("%s said: %s"), λ.NewTuple(
							λ.GetAttr(ϒself, "IE_NAME", nil),
							λ.GetItem(ϒfeed, λ.StrLiteral("error")),
						)))))
					}
					ϒvideo_id = λ.GetItem(ϒitem, λ.StrLiteral("guid"))
					ϒget_media_node = λ.NewFunction("get_media_node",
						[]λ.Param{
							{Name: "name"},
							{Name: "default", Def: λ.None},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒdefault     = λargs[1]
								ϒmedia_group λ.Object
								ϒmedia_name  λ.Object
								ϒname        = λargs[0]
							)
							ϒmedia_name = λ.Mod(λ.StrLiteral("media-%s"), ϒname)
							ϒmedia_group = func() λ.Object {
								if λv := λ.Calm(ϒitem, "get", λ.StrLiteral("media-group")); λ.IsTrue(λv) {
									return λv
								} else {
									return ϒitem
								}
							}()
							return func() λ.Object {
								if λv := λ.Calm(ϒmedia_group, "get", ϒmedia_name); λ.IsTrue(λv) {
									return λv
								} else if λv := λ.Calm(ϒitem, "get", ϒmedia_name); λ.IsTrue(λv) {
									return λv
								} else {
									return λ.Calm(ϒitem, "get", ϒname, ϒdefault)
								}
							}()
						})
					ϒthumbnails = λ.NewList()
					ϒmedia_thumbnail = λ.Cal(ϒget_media_node, λ.StrLiteral("thumbnail"))
					if λ.IsTrue(ϒmedia_thumbnail) {
						if λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒmedia_thumbnail, λ.DictType)) {
							ϒmedia_thumbnail = λ.NewList(ϒmedia_thumbnail)
						}
						τmp0 = λ.Cal(λ.BuiltinIter, ϒmedia_thumbnail)
						for {
							if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
								break
							}
							ϒthumbnail_data = τmp1
							ϒthumbnail = λ.Calm(ϒthumbnail_data, "get", λ.StrLiteral("@attributes"), λ.DictLiteral(map[λ.Object]λ.Object{}))
							ϒthumbnail_url = λ.Cal(ϒurl_or_none, λ.Calm(ϒthumbnail, "get", λ.StrLiteral("url")))
							if !λ.IsTrue(ϒthumbnail_url) {
								continue
							}
							λ.Calm(ϒthumbnails, "append", λ.DictLiteral(map[string]λ.Object{
								"url":    λ.Calm(ϒself, "_proto_relative_url", ϒthumbnail_url, λ.StrLiteral("http:")),
								"width":  λ.Cal(ϒint_or_none, λ.Calm(ϒthumbnail, "get", λ.StrLiteral("width"))),
								"height": λ.Cal(ϒint_or_none, λ.Calm(ϒthumbnail, "get", λ.StrLiteral("height"))),
							}))
						}
					}
					ϒsubtitles = λ.DictLiteral(map[λ.Object]λ.Object{})
					ϒmedia_subtitle = λ.Cal(ϒget_media_node, λ.StrLiteral("subTitle"))
					if λ.IsTrue(ϒmedia_subtitle) {
						if λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒmedia_subtitle, λ.DictType)) {
							ϒmedia_subtitle = λ.NewList(ϒmedia_subtitle)
						}
						τmp0 = λ.Cal(λ.BuiltinIter, ϒmedia_subtitle)
						for {
							if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
								break
							}
							ϒsubtitle_data = τmp1
							ϒsubtitle = λ.Calm(ϒsubtitle_data, "get", λ.StrLiteral("@attributes"), λ.DictLiteral(map[λ.Object]λ.Object{}))
							ϒsubtitle_href = λ.Cal(ϒurl_or_none, λ.Calm(ϒsubtitle, "get", λ.StrLiteral("href")))
							if !λ.IsTrue(ϒsubtitle_href) {
								continue
							}
							λ.Calm(λ.Calm(ϒsubtitles, "setdefault", func() λ.Object {
								if λv := λ.Calm(ϒsubtitle, "get", λ.StrLiteral("lang")); λ.IsTrue(λv) {
									return λv
								} else {
									return λ.StrLiteral("en")
								}
							}(), λ.NewList()), "append", λ.DictLiteral(map[string]λ.Object{
								"url": ϒsubtitle_href,
								"ext": func() λ.Object {
									if λv := λ.Cal(ϒmimetype2ext, λ.Calm(ϒsubtitle, "get", λ.StrLiteral("type"))); λ.IsTrue(λv) {
										return λv
									} else {
										return λ.Cal(ϒdetermine_ext, ϒsubtitle_href)
									}
								}(),
							}))
						}
					}
					ϒformats = λ.NewList()
					ϒmedia_content = λ.Cal(ϒget_media_node, λ.StrLiteral("content"))
					if λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒmedia_content, λ.DictType)) {
						ϒmedia_content = λ.NewList(ϒmedia_content)
					}
					τmp0 = λ.Cal(λ.BuiltinIter, ϒmedia_content)
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒmedia_data = τmp1
						ϒmedia = λ.Calm(ϒmedia_data, "get", λ.StrLiteral("@attributes"), λ.DictLiteral(map[λ.Object]λ.Object{}))
						ϒmedia_url = λ.Cal(ϒurl_or_none, λ.Calm(ϒmedia, "get", λ.StrLiteral("url")))
						if !λ.IsTrue(ϒmedia_url) {
							continue
						}
						ϒext = func() λ.Object {
							if λv := λ.Cal(ϒmimetype2ext, λ.Calm(ϒmedia, "get", λ.StrLiteral("type"))); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Cal(ϒdetermine_ext, ϒmedia_url)
							}
						}()
						if λ.IsTrue(λ.Eq(ϒext, λ.StrLiteral("f4m"))) {
							λ.Calm(ϒformats, "extend", λ.Call(λ.GetAttr(ϒself, "_extract_f4m_formats", nil), λ.NewArgs(
								λ.Add(ϒmedia_url, λ.StrLiteral("?hdcore=3.4.0&plugin=aasp-3.4.0.132.124")),
								ϒvideo_id,
							), λ.KWArgs{
								{Name: "f4m_id", Value: λ.StrLiteral("hds")},
								{Name: "fatal", Value: λ.False},
							}))
						} else {
							if λ.IsTrue(λ.Eq(ϒext, λ.StrLiteral("m3u8"))) {
								λ.Calm(ϒformats, "extend", λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
									ϒmedia_url,
									ϒvideo_id,
									λ.StrLiteral("mp4"),
								), λ.KWArgs{
									{Name: "m3u8_id", Value: λ.StrLiteral("hls")},
									{Name: "fatal", Value: λ.False},
								}))
							} else {
								λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
									"format_id": λ.Calm(λ.Calm(λ.Calm(ϒmedia_data, "get", λ.StrLiteral("media-category"), λ.DictLiteral(map[λ.Object]λ.Object{})), "get", λ.StrLiteral("@attributes"), λ.DictLiteral(map[λ.Object]λ.Object{})), "get", λ.StrLiteral("label")),
									"url":       ϒmedia_url,
									"tbr":       λ.Cal(ϒint_or_none, λ.Calm(ϒmedia, "get", λ.StrLiteral("bitrate"))),
									"filesize":  λ.Cal(ϒint_or_none, λ.Calm(ϒmedia, "get", λ.StrLiteral("fileSize"))),
									"ext":       ϒext,
								}))
							}
						}
					}
					λ.Calm(ϒself, "_sort_formats", ϒformats)
					ϒtimestamp = func() λ.Object {
						if λv := λ.Cal(ϒunified_timestamp, λ.Calm(ϒitem, "get", λ.StrLiteral("pubDate")), λ.StrLiteral(" ")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Cal(ϒparse_iso8601, λ.Calm(ϒitem, "get", λ.StrLiteral("dc-date")))
						}
					}()
					return λ.DictLiteral(map[string]λ.Object{
						"id":          ϒvideo_id,
						"title":       λ.Cal(ϒget_media_node, λ.StrLiteral("title")),
						"description": λ.Cal(ϒget_media_node, λ.StrLiteral("description")),
						"thumbnails":  ϒthumbnails,
						"timestamp":   ϒtimestamp,
						"duration":    λ.Cal(ϒint_or_none, λ.Calm(λ.Calm(λ.GetItem(ϒmedia_content, λ.IntLiteral(0)), "get", λ.StrLiteral("@attributes"), λ.DictLiteral(map[λ.Object]λ.Object{})), "get", λ.StrLiteral("duration"))),
						"subtitles":   ϒsubtitles,
						"formats":     ϒformats,
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_extract_feed_info": AMPIE__extract_feed_info,
			})
		}())
	})
}
