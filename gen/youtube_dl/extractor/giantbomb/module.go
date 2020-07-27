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
 * giantbomb/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/giantbomb.py
 */

package giantbomb

import (
	Ωjson "github.com/tenta-browser/go-video-downloader/gen/json"
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	GiantBombIE    λ.Object
	InfoExtractor  λ.Object
	ϒdetermine_ext λ.Object
	ϒint_or_none   λ.Object
	ϒqualities     λ.Object
	ϒunescapeHTML  λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒdetermine_ext = Ωutils.ϒdetermine_ext
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒqualities = Ωutils.ϒqualities
		ϒunescapeHTML = Ωutils.ϒunescapeHTML
		GiantBombIE = λ.Cal(λ.TypeType, λ.StrLiteral("GiantBombIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				GiantBombIE__VALID_URL    λ.Object
				GiantBombIE__real_extract λ.Object
			)
			GiantBombIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?giantbomb\\.com/(?:videos|shows)/(?P<display_id>[^/]+)/(?P<id>\\d+-\\d+)")
			GiantBombIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdescription λ.Object
						ϒdisplay_id  λ.Object
						ϒduration    λ.Object
						ϒext         λ.Object
						ϒf4m_formats λ.Object
						ϒformat_id   λ.Object
						ϒformats     λ.Object
						ϒmobj        λ.Object
						ϒquality     λ.Object
						ϒself        = λargs[0]
						ϒthumbnail   λ.Object
						ϒtitle       λ.Object
						ϒurl         = λargs[1]
						ϒvideo       λ.Object
						ϒvideo_id    λ.Object
						ϒvideo_url   λ.Object
						ϒwebpage     λ.Object
						ϒyoutube_id  λ.Object
						τmp0         λ.Object
						τmp1         λ.Object
						τmp2         λ.Object
					)
					ϒmobj = λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl)
					ϒvideo_id = λ.Calm(ϒmobj, "group", λ.StrLiteral("id"))
					ϒdisplay_id = λ.Calm(ϒmobj, "group", λ.StrLiteral("display_id"))
					ϒwebpage = λ.Calm(ϒself, "_download_webpage", ϒurl, ϒdisplay_id)
					ϒtitle = λ.Calm(ϒself, "_og_search_title", ϒwebpage)
					ϒdescription = λ.Calm(ϒself, "_og_search_description", ϒwebpage)
					ϒthumbnail = λ.Calm(ϒself, "_og_search_thumbnail", ϒwebpage)
					ϒvideo = λ.Cal(Ωjson.ϒloads, λ.Cal(ϒunescapeHTML, λ.Calm(ϒself, "_search_regex", λ.StrLiteral("data-video=\"([^\"]+)\""), ϒwebpage, λ.StrLiteral("data-video"))))
					ϒduration = λ.Cal(ϒint_or_none, λ.Calm(ϒvideo, "get", λ.StrLiteral("lengthSeconds")))
					ϒquality = λ.Cal(ϒqualities, λ.NewList(
						λ.StrLiteral("f4m_low"),
						λ.StrLiteral("progressive_low"),
						λ.StrLiteral("f4m_high"),
						λ.StrLiteral("progressive_high"),
						λ.StrLiteral("f4m_hd"),
						λ.StrLiteral("progressive_hd"),
					))
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(λ.GetItem(ϒvideo, λ.StrLiteral("videoStreams")), "items"))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						τmp2 = τmp1
						ϒformat_id = λ.GetItem(τmp2, λ.IntLiteral(0))
						ϒvideo_url = λ.GetItem(τmp2, λ.IntLiteral(1))
						if λ.IsTrue(λ.Eq(ϒformat_id, λ.StrLiteral("f4m_stream"))) {
							continue
						}
						ϒext = λ.Cal(ϒdetermine_ext, ϒvideo_url)
						if λ.IsTrue(λ.Eq(ϒext, λ.StrLiteral("f4m"))) {
							ϒf4m_formats = λ.Calm(ϒself, "_extract_f4m_formats", λ.Add(ϒvideo_url, λ.StrLiteral("?hdcore=3.3.1")), ϒdisplay_id)
							if λ.IsTrue(ϒf4m_formats) {
								λ.SetItem(λ.GetItem(ϒf4m_formats, λ.IntLiteral(0)), λ.StrLiteral("quality"), λ.Cal(ϒquality, ϒformat_id))
								λ.Calm(ϒformats, "extend", ϒf4m_formats)
							}
						} else {
							if λ.IsTrue(λ.Eq(ϒext, λ.StrLiteral("m3u8"))) {
								λ.Calm(ϒformats, "extend", λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
									ϒvideo_url,
									ϒdisplay_id,
								), λ.KWArgs{
									{Name: "ext", Value: λ.StrLiteral("mp4")},
									{Name: "entry_protocol", Value: λ.StrLiteral("m3u8_native")},
									{Name: "m3u8_id", Value: λ.StrLiteral("hls")},
									{Name: "fatal", Value: λ.False},
								}))
							} else {
								λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
									"url":       ϒvideo_url,
									"format_id": ϒformat_id,
									"quality":   λ.Cal(ϒquality, ϒformat_id),
								}))
							}
						}
					}
					if !λ.IsTrue(ϒformats) {
						ϒyoutube_id = λ.Calm(ϒvideo, "get", λ.StrLiteral("youtubeID"))
						if λ.IsTrue(ϒyoutube_id) {
							return λ.Calm(ϒself, "url_result", ϒyoutube_id, λ.StrLiteral("Youtube"))
						}
					}
					λ.Calm(ϒself, "_sort_formats", ϒformats)
					return λ.DictLiteral(map[string]λ.Object{
						"id":          ϒvideo_id,
						"display_id":  ϒdisplay_id,
						"title":       ϒtitle,
						"description": ϒdescription,
						"thumbnail":   ϒthumbnail,
						"duration":    ϒduration,
						"formats":     ϒformats,
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL":    GiantBombIE__VALID_URL,
				"_real_extract": GiantBombIE__real_extract,
			})
		}())
	})
}