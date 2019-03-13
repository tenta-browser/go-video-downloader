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
		LibsynIE = λ.Cal(λ.TypeType, λ.NewStr("LibsynIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				LibsynIE__TESTS        λ.Object
				LibsynIE__VALID_URL    λ.Object
				LibsynIE__real_extract λ.Object
			)
			LibsynIE__VALID_URL = λ.NewStr("(?P<mainurl>https?://html5-player\\.libsyn\\.com/embed/episode/id/(?P<id>[0-9]+))")
			LibsynIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://html5-player.libsyn.com/embed/episode/id/6385796/"),
					λ.NewStr("md5"): λ.NewStr("2a55e75496c790cdeb058e7e6c087746"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("6385796"),
						λ.NewStr("ext"):         λ.NewStr("mp3"),
						λ.NewStr("title"):       λ.NewStr("Champion Minded - Developing a Growth Mindset"),
						λ.NewStr("upload_date"): λ.NewStr("20180320"),
						λ.NewStr("thumbnail"):   λ.NewStr("re:^https?://.*"),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("https://html5-player.libsyn.com/embed/episode/id/3727166/height/75/width/200/theme/standard/direction/no/autoplay/no/autonext/no/thumbnail/no/preload/no/no_addthis/no/"),
					λ.NewStr("md5"): λ.NewStr("6c5cb21acd622d754d3b1a92b582ce42"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("3727166"),
						λ.NewStr("ext"):         λ.NewStr("mp3"),
						λ.NewStr("title"):       λ.NewStr("Clients From Hell Podcast - How a Sex Toy Company Kickstarted my Freelance Career"),
						λ.NewStr("upload_date"): λ.NewStr("20150818"),
						λ.NewStr("thumbnail"):   λ.NewStr("re:^https?://.*"),
					}),
				}),
			)
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
					τmp0 = λ.Cal(λ.GetAttr(λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl), "groups", nil))
					ϒurl = λ.GetItem(τmp0, λ.NewInt(0))
					ϒvideo_id = λ.GetItem(τmp0, λ.NewInt(1))
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒvideo_id)
					ϒdata = λ.Cal(λ.GetAttr(ϒself, "_parse_json", nil), λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewStr("var\\s+playlistItem\\s*=\\s*({.+?});"), ϒwebpage, λ.NewStr("JSON data block")), ϒvideo_id)
					ϒepisode_title = func() λ.Object {
						if λv := λ.Cal(λ.GetAttr(ϒdata, "get", nil), λ.NewStr("item_title")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Cal(ϒget_element_by_class, λ.NewStr("episode-title"), ϒwebpage)
						}
					}()
					if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒepisode_title))) {
						λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewList(
							λ.NewStr("data-title=\"([^\"]+)\""),
							λ.NewStr("<title>(.+?)</title>"),
						), ϒwebpage, λ.NewStr("episode title"))
					}
					ϒepisode_title = λ.Cal(λ.GetAttr(ϒepisode_title, "strip", nil))
					ϒpodcast_title = λ.Cal(ϒstrip_or_none, λ.Cal(ϒclean_html, func() λ.Object {
						if λv := λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
							λ.NewStr("<h3>([^<]+)</h3>"),
							ϒwebpage,
							λ.NewStr("podcast title"),
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
						}); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Cal(ϒget_element_by_class, λ.NewStr("podcast-title"), ϒwebpage)
						}
					}()))
					ϒtitle = func() λ.Object {
						if λ.IsTrue(ϒpodcast_title) {
							return λ.Mod(λ.NewStr("%s - %s"), λ.NewTuple(
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
							λ.NewStr("media_url_libsyn"),
							λ.NewStr("libsyn"),
						),
						λ.NewTuple(
							λ.NewStr("media_url"),
							λ.NewStr("main"),
						),
						λ.NewTuple(
							λ.NewStr("download_link"),
							λ.NewStr("download"),
						),
					))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						τmp2 = τmp1
						ϒk = λ.GetItem(τmp2, λ.NewInt(0))
						ϒformat_id = λ.GetItem(τmp2, λ.NewInt(1))
						ϒf_url = λ.Cal(λ.GetAttr(ϒdata, "get", nil), ϒk)
						if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒf_url))) {
							continue
						}
						λ.Cal(λ.GetAttr(ϒformats, "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("url"):       ϒf_url,
							λ.NewStr("format_id"): ϒformat_id,
						}))
					}
					ϒdescription = λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
						λ.NewStr("<p\\s+id=\"info_text_body\">(.+?)</p>"),
						ϒwebpage,
						λ.NewStr("description"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					if λ.IsTrue(ϒdescription) {
						ϒdescription = λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒdescription, "replace", nil), λ.NewStr(" "), λ.NewStr(" ")), "strip", nil))
					}
					ϒrelease_date = λ.Cal(ϒunified_strdate, func() λ.Object {
						if λv := λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
							λ.NewStr("<div class=\"release_date\">Released: ([^<]+)<"),
							ϒwebpage,
							λ.NewStr("release date"),
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
						}); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Cal(λ.GetAttr(ϒdata, "get", nil), λ.NewStr("release_date"))
						}
					}())
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          ϒvideo_id,
						λ.NewStr("title"):       ϒtitle,
						λ.NewStr("description"): ϒdescription,
						λ.NewStr("thumbnail"):   λ.Cal(λ.GetAttr(ϒdata, "get", nil), λ.NewStr("thumbnail_url")),
						λ.NewStr("upload_date"): ϒrelease_date,
						λ.NewStr("duration"):    λ.Cal(ϒparse_duration, λ.Cal(λ.GetAttr(ϒdata, "get", nil), λ.NewStr("duration"))),
						λ.NewStr("formats"):     ϒformats,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TESTS"):        LibsynIE__TESTS,
				λ.NewStr("_VALID_URL"):    LibsynIE__VALID_URL,
				λ.NewStr("_real_extract"): LibsynIE__real_extract,
			})
		}())
	})
}
