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
 * radiojavan/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/radiojavan.py
 */

package radiojavan

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor       λ.Object
	RadioJavanIE        λ.Object
	ϒparse_resolution   λ.Object
	ϒstr_to_int         λ.Object
	ϒunified_strdate    λ.Object
	ϒurlencode_postdata λ.Object
	ϒurljoin            λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒparse_resolution = Ωutils.ϒparse_resolution
		ϒstr_to_int = Ωutils.ϒstr_to_int
		ϒunified_strdate = Ωutils.ϒunified_strdate
		ϒurlencode_postdata = Ωutils.ϒurlencode_postdata
		ϒurljoin = Ωutils.ϒurljoin
		RadioJavanIE = λ.Cal(λ.TypeType, λ.StrLiteral("RadioJavanIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				RadioJavanIE__VALID_URL    λ.Object
				RadioJavanIE__real_extract λ.Object
			)
			RadioJavanIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?radiojavan\\.com/videos/video/(?P<id>[^/]+)/?")
			RadioJavanIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdislike_count λ.Object
						ϒdownload_host λ.Object
						ϒf             λ.Object
						ϒformat_id     λ.Object
						ϒformats       λ.Object
						ϒlike_count    λ.Object
						ϒself          = λargs[0]
						ϒthumbnail     λ.Object
						ϒtitle         λ.Object
						ϒupload_date   λ.Object
						ϒurl           = λargs[1]
						ϒvideo_id      λ.Object
						ϒvideo_path    λ.Object
						ϒview_count    λ.Object
						ϒwebpage       λ.Object
						τmp0           λ.Object
						τmp1           λ.Object
						τmp2           λ.Object
					)
					ϒvideo_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒdownload_host = λ.Calm(λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
						λ.StrLiteral("https://www.radiojavan.com/videos/video_host"),
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "data", Value: λ.Cal(ϒurlencode_postdata, λ.DictLiteral(map[string]λ.Object{
							"id": ϒvideo_id,
						}))},
						{Name: "headers", Value: λ.DictLiteral(map[string]λ.Object{
							"Content-Type": λ.StrLiteral("application/x-www-form-urlencoded"),
							"Referer":      ϒurl,
						})},
					}), "get", λ.StrLiteral("host"), λ.StrLiteral("https://host1.rjmusicmedia.com"))
					ϒwebpage = λ.Calm(ϒself, "_download_webpage", ϒurl, ϒvideo_id)
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(Ωre.ϒfindall, λ.StrLiteral("RJ\\.video(?P<format_id>\\d+[pPkK])\\s*=\\s*([\"\\'])(?P<url>(?:(?!\\2).)+)\\2"), ϒwebpage))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						τmp2 = τmp1
						ϒformat_id = λ.GetItem(τmp2, λ.IntLiteral(0))
						_ = λ.GetItem(τmp2, λ.IntLiteral(1))
						ϒvideo_path = λ.GetItem(τmp2, λ.IntLiteral(2))
						ϒf = λ.Cal(ϒparse_resolution, ϒformat_id)
						λ.Calm(ϒf, "update", λ.DictLiteral(map[string]λ.Object{
							"url":       λ.Cal(ϒurljoin, ϒdownload_host, ϒvideo_path),
							"format_id": ϒformat_id,
						}))
						λ.Calm(ϒformats, "append", ϒf)
					}
					λ.Calm(ϒself, "_sort_formats", ϒformats)
					ϒtitle = λ.Calm(ϒself, "_og_search_title", ϒwebpage)
					ϒthumbnail = λ.Calm(ϒself, "_og_search_thumbnail", ϒwebpage)
					ϒupload_date = λ.Cal(ϒunified_strdate, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("class=\"date_added\">Date added: ([^<]+)<"),
						ϒwebpage,
						λ.StrLiteral("upload date"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					}))
					ϒview_count = λ.Cal(ϒstr_to_int, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("class=\"views\">Plays: ([\\d,]+)"),
						ϒwebpage,
						λ.StrLiteral("view count"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					}))
					ϒlike_count = λ.Cal(ϒstr_to_int, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("class=\"rating\">([\\d,]+) likes"),
						ϒwebpage,
						λ.StrLiteral("like count"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					}))
					ϒdislike_count = λ.Cal(ϒstr_to_int, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("class=\"rating\">([\\d,]+) dislikes"),
						ϒwebpage,
						λ.StrLiteral("dislike count"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					}))
					return λ.DictLiteral(map[string]λ.Object{
						"id":            ϒvideo_id,
						"title":         ϒtitle,
						"thumbnail":     ϒthumbnail,
						"upload_date":   ϒupload_date,
						"view_count":    ϒview_count,
						"like_count":    ϒlike_count,
						"dislike_count": ϒdislike_count,
						"formats":       ϒformats,
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL":    RadioJavanIE__VALID_URL,
				"_real_extract": RadioJavanIE__real_extract,
			})
		}())
	})
}
