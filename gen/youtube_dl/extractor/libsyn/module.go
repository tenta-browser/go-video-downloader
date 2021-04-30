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
 * libsyn/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/libsyn.py
 */

package libsyn

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor         λ.Object
	LibsynIE              λ.Object
	ϒclean_html           λ.Object
	ϒget_element_by_class λ.Object
	ϒparse_duration       λ.Object
	ϒstrip_or_none        λ.Object
	ϒunified_strdate      λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒclean_html = Ωutils.ϒclean_html
		ϒget_element_by_class = Ωutils.ϒget_element_by_class
		ϒparse_duration = Ωutils.ϒparse_duration
		ϒstrip_or_none = Ωutils.ϒstrip_or_none
		ϒunified_strdate = Ωutils.ϒunified_strdate
		LibsynIE = λ.Cal(λ.TypeType, λ.StrLiteral("LibsynIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				LibsynIE__VALID_URL    λ.Object
				LibsynIE__real_extract λ.Object
			)
			LibsynIE__VALID_URL = λ.StrLiteral("(?P<mainurl>https?://html5-player\\.libsyn\\.com/embed/episode/id/(?P<id>[0-9]+))")
			LibsynIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdata          λ.Object
						ϒdescription   λ.Object
						ϒepisode_title λ.Object
						ϒf_url         λ.Object
						ϒformat_id     λ.Object
						ϒformats       λ.Object
						ϒk             λ.Object
						ϒpodcast_title λ.Object
						ϒrelease_date  λ.Object
						ϒself          = λargs[0]
						ϒtitle         λ.Object
						ϒurl           = λargs[1]
						ϒvideo_id      λ.Object
						ϒwebpage       λ.Object
						τmp0           λ.Object
						τmp1           λ.Object
						τmp2           λ.Object
					)
					τmp0 = λ.UnpackIterable(λ.Calm(λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl), "groups"), 2)
					ϒurl = λ.GetItem(τmp0, λ.IntLiteral(0))
					ϒvideo_id = λ.GetItem(τmp0, λ.IntLiteral(1))
					ϒwebpage = λ.Calm(ϒself, "_download_webpage", ϒurl, ϒvideo_id)
					ϒdata = λ.Calm(ϒself, "_parse_json", λ.Calm(ϒself, "_search_regex", λ.StrLiteral("var\\s+playlistItem\\s*=\\s*({.+?});"), ϒwebpage, λ.StrLiteral("JSON data block")), ϒvideo_id)
					ϒepisode_title = func() λ.Object {
						if λv := λ.Calm(ϒdata, "get", λ.StrLiteral("item_title")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Cal(ϒget_element_by_class, λ.StrLiteral("episode-title"), ϒwebpage)
						}
					}()
					if !λ.IsTrue(ϒepisode_title) {
						λ.Calm(ϒself, "_search_regex", λ.NewList(
							λ.StrLiteral("data-title=\"([^\"]+)\""),
							λ.StrLiteral("<title>(.+?)</title>"),
						), ϒwebpage, λ.StrLiteral("episode title"))
					}
					ϒepisode_title = λ.Calm(ϒepisode_title, "strip")
					ϒpodcast_title = λ.Cal(ϒstrip_or_none, λ.Cal(ϒclean_html, func() λ.Object {
						if λv := λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
							λ.StrLiteral("<h3>([^<]+)</h3>"),
							ϒwebpage,
							λ.StrLiteral("podcast title"),
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
						}); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Cal(ϒget_element_by_class, λ.StrLiteral("podcast-title"), ϒwebpage)
						}
					}()))
					ϒtitle = func() λ.Object {
						if λ.IsTrue(ϒpodcast_title) {
							return λ.Mod(λ.StrLiteral("%s - %s"), λ.NewTuple(
								ϒpodcast_title,
								ϒepisode_title,
							))
						} else {
							return ϒepisode_title
						}
					}()
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.NewTuple(
						λ.NewTuple(
							λ.StrLiteral("media_url_libsyn"),
							λ.StrLiteral("libsyn"),
						),
						λ.NewTuple(
							λ.StrLiteral("media_url"),
							λ.StrLiteral("main"),
						),
						λ.NewTuple(
							λ.StrLiteral("download_link"),
							λ.StrLiteral("download"),
						),
					))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						τmp2 = λ.UnpackIterable(τmp1, 2)
						ϒk = λ.GetItem(τmp2, λ.IntLiteral(0))
						ϒformat_id = λ.GetItem(τmp2, λ.IntLiteral(1))
						ϒf_url = λ.Calm(ϒdata, "get", ϒk)
						if !λ.IsTrue(ϒf_url) {
							continue
						}
						λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
							"url":       ϒf_url,
							"format_id": ϒformat_id,
						}))
					}
					ϒdescription = λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("<p\\s+id=\"info_text_body\">(.+?)</p>"),
						ϒwebpage,
						λ.StrLiteral("description"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					if λ.IsTrue(ϒdescription) {
						ϒdescription = λ.Calm(λ.Calm(ϒdescription, "replace", λ.StrLiteral(" "), λ.StrLiteral(" ")), "strip")
					}
					ϒrelease_date = λ.Cal(ϒunified_strdate, func() λ.Object {
						if λv := λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
							λ.StrLiteral("<div class=\"release_date\">Released: ([^<]+)<"),
							ϒwebpage,
							λ.StrLiteral("release date"),
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
						}); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Calm(ϒdata, "get", λ.StrLiteral("release_date"))
						}
					}())
					return λ.DictLiteral(map[string]λ.Object{
						"id":          ϒvideo_id,
						"title":       ϒtitle,
						"description": ϒdescription,
						"thumbnail":   λ.Calm(ϒdata, "get", λ.StrLiteral("thumbnail_url")),
						"upload_date": ϒrelease_date,
						"duration":    λ.Cal(ϒparse_duration, λ.Calm(ϒdata, "get", λ.StrLiteral("duration"))),
						"formats":     ϒformats,
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL":    LibsynIE__VALID_URL,
				"_real_extract": LibsynIE__real_extract,
			})
		}())
	})
}
