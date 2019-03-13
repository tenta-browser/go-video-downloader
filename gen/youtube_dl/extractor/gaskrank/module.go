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
 * gaskrank/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/gaskrank.py
 */

package gaskrank

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	GaskrankIE       λ.Object
	InfoExtractor    λ.Object
	ϒfloat_or_none   λ.Object
	ϒint_or_none     λ.Object
	ϒunified_strdate λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒfloat_or_none = Ωutils.ϒfloat_or_none
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒunified_strdate = Ωutils.ϒunified_strdate
		GaskrankIE = λ.Cal(λ.TypeType, λ.NewStr("GaskrankIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				GaskrankIE__TESTS        λ.Object
				GaskrankIE__VALID_URL    λ.Object
				GaskrankIE__real_extract λ.Object
			)
			GaskrankIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?gaskrank\\.tv/tv/(?P<categories>[^/]+)/(?P<id>[^/]+)\\.htm")
			GaskrankIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.gaskrank.tv/tv/motorrad-fun/strike-einparken-durch-anfaenger-crash-mit-groesserem-flurschaden.htm"),
					λ.NewStr("md5"): λ.NewStr("1ae88dbac97887d85ebd1157a95fc4f9"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):           λ.NewStr("201601/26955"),
						λ.NewStr("ext"):          λ.NewStr("mp4"),
						λ.NewStr("title"):        λ.NewStr("Strike! Einparken können nur Männer - Flurschaden hält sich in Grenzen *lol*"),
						λ.NewStr("thumbnail"):    λ.NewStr("re:^https?://.*\\.jpg$"),
						λ.NewStr("categories"):   λ.NewList(λ.NewStr("motorrad-fun")),
						λ.NewStr("display_id"):   λ.NewStr("strike-einparken-durch-anfaenger-crash-mit-groesserem-flurschaden"),
						λ.NewStr("uploader_id"):  λ.NewStr("Bikefun"),
						λ.NewStr("upload_date"):  λ.NewStr("20170110"),
						λ.NewStr("uploader_url"): λ.None,
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.gaskrank.tv/tv/racing/isle-of-man-tt-2011-michael-du-15920.htm"),
					λ.NewStr("md5"): λ.NewStr("c33ee32c711bc6c8224bfcbe62b23095"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):           λ.NewStr("201106/15920"),
						λ.NewStr("ext"):          λ.NewStr("mp4"),
						λ.NewStr("title"):        λ.NewStr("Isle of Man - Michael Dunlop vs Guy Martin - schwindelig kucken"),
						λ.NewStr("thumbnail"):    λ.NewStr("re:^https?://.*\\.jpg$"),
						λ.NewStr("categories"):   λ.NewList(λ.NewStr("racing")),
						λ.NewStr("display_id"):   λ.NewStr("isle-of-man-tt-2011-michael-du-15920"),
						λ.NewStr("uploader_id"):  λ.NewStr("IOM"),
						λ.NewStr("upload_date"):  λ.NewStr("20170523"),
						λ.NewStr("uploader_url"): λ.NewStr("www.iomtt.com"),
					}),
				}),
			)
			GaskrankIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒaverage_rating λ.Object
						ϒcategories     λ.Object
						ϒdisplay_id     λ.Object
						ϒentry          λ.Object
						ϒmobj           λ.Object
						ϒself           = λargs[0]
						ϒtags           λ.Object
						ϒtitle          λ.Object
						ϒupload_date    λ.Object
						ϒuploader_id    λ.Object
						ϒuploader_url   λ.Object
						ϒurl            = λargs[1]
						ϒvideo_id       λ.Object
						ϒview_count     λ.Object
						ϒwebpage        λ.Object
					)
					ϒdisplay_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒdisplay_id)
					ϒtitle = func() λ.Object {
						if λv := λ.Call(λ.GetAttr(ϒself, "_og_search_title", nil), λ.NewArgs(ϒwebpage), λ.KWArgs{
							{Name: "default", Value: λ.None},
						}); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Call(λ.GetAttr(ϒself, "_html_search_meta", nil), λ.NewArgs(
								λ.NewStr("title"),
								ϒwebpage,
							), λ.KWArgs{
								{Name: "fatal", Value: λ.True},
							})
						}
					}()
					ϒcategories = λ.NewList(λ.Cal(λ.GetAttr(λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl), "group", nil), λ.NewStr("categories")))
					ϒmobj = λ.Cal(Ωre.ϒsearch, λ.NewStr("Video von:\\s*(?P<uploader_id>[^|]*?)\\s*\\|\\s*vom:\\s*(?P<upload_date>[0-9][0-9]\\.[0-9][0-9]\\.[0-9][0-9][0-9][0-9])"), ϒwebpage)
					if λ.IsTrue(λ.NewBool(ϒmobj != λ.None)) {
						ϒuploader_id = λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒmobj, "groupdict", nil)), "get", nil), λ.NewStr("uploader_id"))
						ϒupload_date = λ.Cal(ϒunified_strdate, λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒmobj, "groupdict", nil)), "get", nil), λ.NewStr("upload_date")))
					}
					ϒuploader_url = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr("Homepage:\\s*<[^>]*>(?P<uploader_url>[^<]*)"),
						ϒwebpage,
						λ.NewStr("uploader_url"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					ϒtags = λ.Cal(Ωre.ϒfindall, λ.NewStr("/tv/tags/[^/]+/\"\\s*>(?P<tag>[^<]*?)<"), ϒwebpage)
					ϒview_count = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr("class\\s*=\\s*\"gkRight\"(?:[^>]*>\\s*<[^>]*)*icon-eye-open(?:[^>]*>\\s*<[^>]*)*>\\s*(?P<view_count>[0-9\\.]*)"),
						ϒwebpage,
						λ.NewStr("view_count"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					if λ.IsTrue(ϒview_count) {
						ϒview_count = λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒview_count, "replace", nil), λ.NewStr("."), λ.NewStr("")))
					}
					ϒaverage_rating = λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewStr("itemprop\\s*=\\s*\"ratingValue\"[^>]*>\\s*(?P<average_rating>[0-9,]+)"), ϒwebpage, λ.NewStr("average_rating"))
					if λ.IsTrue(ϒaverage_rating) {
						ϒaverage_rating = λ.Cal(ϒfloat_or_none, λ.Cal(λ.GetAttr(ϒaverage_rating, "replace", nil), λ.NewStr(","), λ.NewStr(".")))
					}
					ϒvideo_id = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr("https?://movies\\.gaskrank\\.tv/([^-]*?)(-[^\\.]*)?\\.mp4"),
						ϒwebpage,
						λ.NewStr("video id"),
					), λ.KWArgs{
						{Name: "default", Value: ϒdisplay_id},
					})
					ϒentry = λ.GetItem(λ.Cal(λ.GetAttr(ϒself, "_parse_html5_media_entries", nil), ϒurl, ϒwebpage, ϒvideo_id), λ.NewInt(0))
					λ.Cal(λ.GetAttr(ϒentry, "update", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):             ϒvideo_id,
						λ.NewStr("title"):          ϒtitle,
						λ.NewStr("categories"):     ϒcategories,
						λ.NewStr("display_id"):     ϒdisplay_id,
						λ.NewStr("uploader_id"):    ϒuploader_id,
						λ.NewStr("upload_date"):    ϒupload_date,
						λ.NewStr("uploader_url"):   ϒuploader_url,
						λ.NewStr("tags"):           ϒtags,
						λ.NewStr("view_count"):     ϒview_count,
						λ.NewStr("average_rating"): ϒaverage_rating,
					}))
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), λ.GetItem(ϒentry, λ.NewStr("formats")))
					return ϒentry
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TESTS"):        GaskrankIE__TESTS,
				λ.NewStr("_VALID_URL"):    GaskrankIE__VALID_URL,
				λ.NewStr("_real_extract"): GaskrankIE__real_extract,
			})
		}())
	})
}
