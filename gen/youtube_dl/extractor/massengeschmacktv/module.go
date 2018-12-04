// Code generated by transpiler. DO NOT EDIT.

/**
 * Go Video Downloader
 *
 *    Copyright 2018 Tenta, LLC
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
 * massengeschmacktv/module.go: transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/massengeschmacktv.py
 */

package massengeschmacktv

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor       λ.Object
	MassengeschmackTVIE λ.Object
	ϒclean_html         λ.Object
	ϒdetermine_ext      λ.Object
	ϒint_or_none        λ.Object
	ϒjs_to_json         λ.Object
	ϒmimetype2ext       λ.Object
	ϒparse_filesize     λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒclean_html = Ωutils.ϒclean_html
		ϒdetermine_ext = Ωutils.ϒdetermine_ext
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒjs_to_json = Ωutils.ϒjs_to_json
		ϒmimetype2ext = Ωutils.ϒmimetype2ext
		ϒparse_filesize = Ωutils.ϒparse_filesize
		MassengeschmackTVIE = λ.Cal(λ.TypeType, λ.NewStr("MassengeschmackTVIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				MassengeschmackTVIE_IE_NAME       λ.Object
				MassengeschmackTVIE__TEST         λ.Object
				MassengeschmackTVIE__VALID_URL    λ.Object
				MassengeschmackTVIE__real_extract λ.Object
			)
			MassengeschmackTVIE_IE_NAME = λ.NewStr("massengeschmack.tv")
			MassengeschmackTVIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?massengeschmack\\.tv/play/(?P<id>[^?&#]+)")
			MassengeschmackTVIE__TEST = λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("url"): λ.NewStr("https://massengeschmack.tv/play/fktv202"),
				λ.NewStr("md5"): λ.NewStr("a9e054db9c2b5a08f0a0527cc201e8d3"),
				λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("id"):    λ.NewStr("fktv202"),
					λ.NewStr("ext"):   λ.NewStr("mp4"),
					λ.NewStr("title"): λ.NewStr("Fernsehkritik-TV - Folge 202"),
				}),
			})
			MassengeschmackTVIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdurl      λ.Object
						ϒepisode   λ.Object
						ϒext       λ.Object
						ϒfilesize  λ.Object
						ϒformat_id λ.Object
						ϒformats   λ.Object
						ϒfurl      λ.Object
						ϒheight    λ.Object
						ϒself      = λargs[0]
						ϒsource    λ.Object
						ϒsources   λ.Object
						ϒthumbnail λ.Object
						ϒtitle     λ.Object
						ϒurl       = λargs[1]
						ϒwebpage   λ.Object
						ϒwidth     λ.Object
						τmp0       λ.Object
						τmp1       λ.Object
						τmp2       λ.Object
					)
					ϒepisode = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒepisode)
					ϒtitle = λ.Cal(ϒclean_html, λ.Cal(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewStr("<h3>([^<]+)</h3>"), ϒwebpage, λ.NewStr("title")))
					ϒthumbnail = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr("POSTER\\s*=\\s*\"([^\"]+)"),
						ϒwebpage,
						λ.NewStr("thumbnail"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					})
					ϒsources = λ.Cal(λ.GetAttr(ϒself, "_parse_json", nil), λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewStr("(?s)MEDIA\\s*=\\s*(\\[.+?\\]);"), ϒwebpage, λ.NewStr("media")), ϒepisode, ϒjs_to_json)
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, ϒsources)
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒsource = τmp1
						ϒfurl = λ.Cal(λ.GetAttr(ϒsource, "get", nil), λ.NewStr("src"))
						if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒfurl))) {
							continue
						}
						ϒfurl = λ.Cal(λ.GetAttr(ϒself, "_proto_relative_url", nil), ϒfurl)
						ϒext = func() λ.Object {
							if λv := λ.Cal(ϒdetermine_ext, ϒfurl); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Cal(ϒmimetype2ext, λ.Cal(λ.GetAttr(ϒsource, "get", nil), λ.NewStr("type")))
							}
						}()
						if λ.IsTrue(λ.Eq(ϒext, λ.NewStr("m3u8"))) {
							λ.Cal(λ.GetAttr(ϒformats, "extend", nil), λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
								ϒfurl,
								ϒepisode,
								λ.NewStr("mp4"),
								λ.NewStr("m3u8_native"),
							), λ.KWArgs{
								{Name: "m3u8_id", Value: λ.NewStr("hls")},
								{Name: "fatal", Value: λ.False},
							}))
						} else {
							λ.Cal(λ.GetAttr(ϒformats, "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
								λ.NewStr("url"):       ϒfurl,
								λ.NewStr("format_id"): λ.Cal(ϒdetermine_ext, ϒfurl),
							}))
						}
					}
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(Ωre.ϒfindall, λ.NewStr("(?x)\n                                   <a[^>]+?href=\"(?P<url>(?:https:)?//[^\"]+)\".*?\n                                   <strong>(?P<format_id>.+?)</strong>.*?\n                                   <small>(?:(?P<width>\\d+)x(?P<height>\\d+))?\\s+?\\((?P<filesize>[\\d,]+\\s*[GM]iB)\\)</small>\n                                "), ϒwebpage))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						τmp2 = τmp1
						ϒdurl = λ.GetItem(τmp2, λ.NewInt(0))
						ϒformat_id = λ.GetItem(τmp2, λ.NewInt(1))
						ϒwidth = λ.GetItem(τmp2, λ.NewInt(2))
						ϒheight = λ.GetItem(τmp2, λ.NewInt(3))
						ϒfilesize = λ.GetItem(τmp2, λ.NewInt(4))
						λ.Cal(λ.GetAttr(ϒformats, "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("url"):       ϒdurl,
							λ.NewStr("format_id"): ϒformat_id,
							λ.NewStr("width"):     λ.Cal(ϒint_or_none, ϒwidth),
							λ.NewStr("height"):    λ.Cal(ϒint_or_none, ϒheight),
							λ.NewStr("filesize"):  λ.Cal(ϒparse_filesize, ϒfilesize),
							λ.NewStr("vcodec"): func() λ.Object {
								if λ.IsTrue(λ.Cal(λ.GetAttr(ϒformat_id, "startswith", nil), λ.NewStr("Audio"))) {
									return λ.NewStr("none")
								} else {
									return λ.None
								}
							}(),
						}))
					}
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats, λ.NewTuple(
						λ.NewStr("width"),
						λ.NewStr("height"),
						λ.NewStr("filesize"),
						λ.NewStr("tbr"),
					))
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):        ϒepisode,
						λ.NewStr("title"):     ϒtitle,
						λ.NewStr("formats"):   ϒformats,
						λ.NewStr("thumbnail"): ϒthumbnail,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("IE_NAME"):       MassengeschmackTVIE_IE_NAME,
				λ.NewStr("_TEST"):         MassengeschmackTVIE__TEST,
				λ.NewStr("_VALID_URL"):    MassengeschmackTVIE__VALID_URL,
				λ.NewStr("_real_extract"): MassengeschmackTVIE__real_extract,
			})
		}())
	})
}