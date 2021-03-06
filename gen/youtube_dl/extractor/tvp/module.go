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
 * tvp/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/tvp.py
 */

package tvp

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError            λ.Object
	InfoExtractor             λ.Object
	TVPEmbedIE                λ.Object
	TVPIE                     λ.Object
	TVPWebsiteIE              λ.Object
	ϒclean_html               λ.Object
	ϒdetermine_ext            λ.Object
	ϒget_element_by_attribute λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒclean_html = Ωutils.ϒclean_html
		ϒdetermine_ext = Ωutils.ϒdetermine_ext
		ExtractorError = Ωutils.ExtractorError
		ϒget_element_by_attribute = Ωutils.ϒget_element_by_attribute
		TVPIE = λ.Cal(λ.TypeType, λ.StrLiteral("TVPIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				TVPIE_IE_NAME       λ.Object
				TVPIE__VALID_URL    λ.Object
				TVPIE__real_extract λ.Object
			)
			TVPIE_IE_NAME = λ.StrLiteral("tvp")
			TVPIE__VALID_URL = λ.StrLiteral("https?://[^/]+\\.tvp\\.(?:pl|info)/(?:video/(?:[^,\\s]*,)*|(?:(?!\\d+/)[^/]+/)*)(?P<id>\\d+)")
			TVPIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒpage_id  λ.Object
						ϒself     = λargs[0]
						ϒurl      = λargs[1]
						ϒvideo_id λ.Object
						ϒwebpage  λ.Object
					)
					ϒpage_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒwebpage = λ.Calm(ϒself, "_download_webpage", ϒurl, ϒpage_id)
					ϒvideo_id = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewList(
							λ.StrLiteral("<iframe[^>]+src=\"[^\"]*?object_id=(\\d+)"),
							λ.StrLiteral("object_id\\s*:\\s*'(\\d+)'"),
							λ.StrLiteral("data-video-id=\"(\\d+)\""),
						),
						ϒwebpage,
						λ.StrLiteral("video id"),
					), λ.KWArgs{
						{Name: "default", Value: ϒpage_id},
					})
					return λ.DictLiteral(map[string]λ.Object{
						"_type": λ.StrLiteral("url_transparent"),
						"url":   λ.Add(λ.StrLiteral("tvp:"), ϒvideo_id),
						"description": func() λ.Object {
							if λv := λ.Call(λ.GetAttr(ϒself, "_og_search_description", nil), λ.NewArgs(ϒwebpage), λ.KWArgs{
								{Name: "default", Value: λ.None},
							}); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Call(λ.GetAttr(ϒself, "_html_search_meta", nil), λ.NewArgs(
									λ.StrLiteral("description"),
									ϒwebpage,
								), λ.KWArgs{
									{Name: "default", Value: λ.None},
								})
							}
						}(),
						"thumbnail": λ.Call(λ.GetAttr(ϒself, "_og_search_thumbnail", nil), λ.NewArgs(ϒwebpage), λ.KWArgs{
							{Name: "default", Value: λ.None},
						}),
						"ie_key": λ.StrLiteral("TVPEmbed"),
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"IE_NAME":       TVPIE_IE_NAME,
				"_VALID_URL":    TVPIE__VALID_URL,
				"_real_extract": TVPIE__real_extract,
			})
		}())
		TVPEmbedIE = λ.Cal(λ.TypeType, λ.StrLiteral("TVPEmbedIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				TVPEmbedIE_IE_NAME       λ.Object
				TVPEmbedIE__VALID_URL    λ.Object
				TVPEmbedIE__real_extract λ.Object
			)
			TVPEmbedIE_IE_NAME = λ.StrLiteral("tvp:embed")
			TVPEmbedIE__VALID_URL = λ.StrLiteral("(?:tvp:|https?://[^/]+\\.tvp\\.(?:pl|info)/sess/tvplayer\\.php\\?.*?object_id=)(?P<id>\\d+)")
			TVPEmbedIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒerror          λ.Object
						ϒf              λ.Object
						ϒformats        λ.Object
						ϒhttp_url       λ.Object
						ϒi              λ.Object
						ϒm3u8_format    λ.Object
						ϒm3u8_formats   λ.Object
						ϒself           = λargs[0]
						ϒseries_title   λ.Object
						ϒthumbnail      λ.Object
						ϒtitle          λ.Object
						ϒurl            = λargs[1]
						ϒvideo_id       λ.Object
						ϒvideo_url      λ.Object
						ϒvideo_url_base λ.Object
						ϒwebpage        λ.Object
						τmp0            λ.Object
						τmp1            λ.Object
						τmp2            λ.Object
					)
					ϒvideo_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒwebpage = λ.Calm(ϒself, "_download_webpage", λ.Mod(λ.StrLiteral("http://www.tvp.pl/sess/tvplayer.php?object_id=%s"), ϒvideo_id), ϒvideo_id)
					ϒerror = func() λ.Object {
						if λv := λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
							λ.StrLiteral("(?s)<p[^>]+\\bclass=[\"\\']notAvailable__text[\"\\'][^>]*>(.+?)</p>"),
							ϒwebpage,
							λ.StrLiteral("error"),
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
						}); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Cal(ϒclean_html, λ.Cal(ϒget_element_by_attribute, λ.StrLiteral("class"), λ.StrLiteral("msg error"), ϒwebpage))
						}
					}()
					if λ.IsTrue(ϒerror) {
						panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.Mod(λ.StrLiteral("%s said: %s"), λ.NewTuple(
							λ.GetAttr(ϒself, "IE_NAME", nil),
							λ.Cal(ϒclean_html, ϒerror),
						))), λ.KWArgs{
							{Name: "expected", Value: λ.True},
						})))
					}
					ϒtitle = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("name\\s*:\\s*([\\'\"])Title\\1\\s*,\\s*value\\s*:\\s*\\1(?P<title>.+?)\\1"),
						ϒwebpage,
						λ.StrLiteral("title"),
					), λ.KWArgs{
						{Name: "group", Value: λ.StrLiteral("title")},
					})
					ϒseries_title = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("name\\s*:\\s*([\\'\"])SeriesTitle\\1\\s*,\\s*value\\s*:\\s*\\1(?P<series>.+?)\\1"),
						ϒwebpage,
						λ.StrLiteral("series"),
					), λ.KWArgs{
						{Name: "group", Value: λ.StrLiteral("series")},
						{Name: "default", Value: λ.None},
					})
					if λ.IsTrue(ϒseries_title) {
						ϒtitle = λ.Mod(λ.StrLiteral("%s, %s"), λ.NewTuple(
							ϒseries_title,
							ϒtitle,
						))
					}
					ϒthumbnail = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("poster\\s*:\\s*'([^']+)'"),
						ϒwebpage,
						λ.StrLiteral("thumbnail"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					ϒvideo_url = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("0:{src:([\\'\"])(?P<url>.*?)\\1"),
						ϒwebpage,
						λ.StrLiteral("formats"),
					), λ.KWArgs{
						{Name: "group", Value: λ.StrLiteral("url")},
						{Name: "default", Value: λ.None},
					})
					if λ.IsTrue(func() λ.Object {
						if λv := λ.NewBool(!λ.IsTrue(ϒvideo_url)); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.NewBool(λ.Contains(ϒvideo_url, λ.StrLiteral("material_niedostepny.mp4")))
						}
					}()) {
						ϒvideo_url = λ.GetItem(λ.Calm(ϒself, "_download_json", λ.Mod(λ.StrLiteral("http://www.tvp.pl/pub/stat/videofileinfo?video_id=%s"), ϒvideo_id), ϒvideo_id), λ.StrLiteral("video_url"))
					}
					ϒformats = λ.NewList()
					ϒvideo_url_base = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("(https?://.+?/video)(?:\\.(?:ism|f4m|m3u8)|-\\d+\\.mp4)"),
						ϒvideo_url,
						λ.StrLiteral("video base url"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					if λ.IsTrue(ϒvideo_url_base) {
						λ.Calm(ϒformats, "extend", λ.Call(λ.GetAttr(ϒself, "_extract_ism_formats", nil), λ.NewArgs(
							λ.Add(ϒvideo_url_base, λ.StrLiteral(".ism/Manifest")),
							ϒvideo_id,
							λ.StrLiteral("mss"),
						), λ.KWArgs{
							{Name: "fatal", Value: λ.False},
						}))
						λ.Calm(ϒformats, "extend", λ.Call(λ.GetAttr(ϒself, "_extract_f4m_formats", nil), λ.NewArgs(
							λ.Add(ϒvideo_url_base, λ.StrLiteral(".ism/video.f4m")),
							ϒvideo_id,
						), λ.KWArgs{
							{Name: "f4m_id", Value: λ.StrLiteral("hds")},
							{Name: "fatal", Value: λ.False},
						}))
						ϒm3u8_formats = λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
							λ.Add(ϒvideo_url_base, λ.StrLiteral(".ism/video.m3u8")),
							ϒvideo_id,
							λ.StrLiteral("mp4"),
							λ.StrLiteral("m3u8_native"),
						), λ.KWArgs{
							{Name: "m3u8_id", Value: λ.StrLiteral("hls")},
							{Name: "fatal", Value: λ.False},
						})
						λ.Calm(ϒself, "_sort_formats", ϒm3u8_formats)
						ϒm3u8_formats = λ.Cal(λ.ListType, λ.Cal(λ.FilterIteratorType, λ.NewFunction("<lambda>",
							[]λ.Param{
								{Name: "f"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒf = λargs[0]
								)
								return λ.Ne(λ.Calm(ϒf, "get", λ.StrLiteral("vcodec")), λ.StrLiteral("none"))
							}), ϒm3u8_formats))
						λ.Calm(ϒformats, "extend", ϒm3u8_formats)
						τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.EnumerateIteratorType, ϒm3u8_formats, λ.IntLiteral(2)))
						for {
							if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
								break
							}
							τmp2 = λ.UnpackIterable(τmp1, 2)
							ϒi = λ.GetItem(τmp2, λ.IntLiteral(0))
							ϒm3u8_format = λ.GetItem(τmp2, λ.IntLiteral(1))
							ϒhttp_url = λ.Mod(λ.StrLiteral("%s-%d.mp4"), λ.NewTuple(
								ϒvideo_url_base,
								ϒi,
							))
							if λ.IsTrue(λ.Calm(ϒself, "_is_valid_url", ϒhttp_url, ϒvideo_id)) {
								ϒf = λ.Calm(ϒm3u8_format, "copy")
								λ.Calm(ϒf, "update", λ.DictLiteral(map[string]λ.Object{
									"url":       ϒhttp_url,
									"format_id": λ.Calm(λ.GetItem(ϒf, λ.StrLiteral("format_id")), "replace", λ.StrLiteral("hls"), λ.StrLiteral("http")),
									"protocol":  λ.StrLiteral("http"),
								}))
								λ.Calm(ϒformats, "append", ϒf)
							}
						}
					} else {
						ϒformats = λ.NewList(λ.DictLiteral(map[string]λ.Object{
							"format_id": λ.StrLiteral("direct"),
							"url":       ϒvideo_url,
							"ext":       λ.Cal(ϒdetermine_ext, ϒvideo_url, λ.StrLiteral("mp4")),
						}))
					}
					λ.Calm(ϒself, "_sort_formats", ϒformats)
					return λ.DictLiteral(map[string]λ.Object{
						"id":        ϒvideo_id,
						"title":     ϒtitle,
						"thumbnail": ϒthumbnail,
						"formats":   ϒformats,
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"IE_NAME":       TVPEmbedIE_IE_NAME,
				"_VALID_URL":    TVPEmbedIE__VALID_URL,
				"_real_extract": TVPEmbedIE__real_extract,
			})
		}())
		TVPWebsiteIE = λ.Cal(λ.TypeType, λ.StrLiteral("TVPWebsiteIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				TVPWebsiteIE__VALID_URL λ.Object
			)
			TVPWebsiteIE__VALID_URL = λ.StrLiteral("https?://vod\\.tvp\\.pl/website/(?P<display_id>[^,]+),(?P<id>\\d+)")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": TVPWebsiteIE__VALID_URL,
			})
		}())
	})
}
