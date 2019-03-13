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
 * xxxymovies/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/xxxymovies.py
 */

package xxxymovies

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor   λ.Object
	XXXYMoviesIE    λ.Object
	ϒint_or_none    λ.Object
	ϒparse_duration λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒparse_duration = Ωutils.ϒparse_duration
		ϒint_or_none = Ωutils.ϒint_or_none
		XXXYMoviesIE = λ.Cal(λ.TypeType, λ.NewStr("XXXYMoviesIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				XXXYMoviesIE__TEST         λ.Object
				XXXYMoviesIE__VALID_URL    λ.Object
				XXXYMoviesIE__real_extract λ.Object
			)
			XXXYMoviesIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?xxxymovies\\.com/videos/(?P<id>\\d+)/(?P<display_id>[^/]+)")
			XXXYMoviesIE__TEST = λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("url"): λ.NewStr("http://xxxymovies.com/videos/138669/ecstatic-orgasm-sofcore/"),
				λ.NewStr("md5"): λ.NewStr("810b1bdbbffff89dd13bdb369fe7be4b"),
				λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("id"):            λ.NewStr("138669"),
					λ.NewStr("display_id"):    λ.NewStr("ecstatic-orgasm-sofcore"),
					λ.NewStr("ext"):           λ.NewStr("mp4"),
					λ.NewStr("title"):         λ.NewStr("Ecstatic Orgasm Sofcore"),
					λ.NewStr("duration"):      λ.NewInt(931),
					λ.NewStr("categories"):    λ.ListType,
					λ.NewStr("view_count"):    λ.IntType,
					λ.NewStr("like_count"):    λ.IntType,
					λ.NewStr("dislike_count"): λ.IntType,
					λ.NewStr("age_limit"):     λ.NewInt(18),
				}),
			})
			XXXYMoviesIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒage_limit     λ.Object
						ϒcategories    λ.Object
						ϒdislike_count λ.Object
						ϒdisplay_id    λ.Object
						ϒduration      λ.Object
						ϒlike_count    λ.Object
						ϒmobj          λ.Object
						ϒself          = λargs[0]
						ϒthumbnail     λ.Object
						ϒtitle         λ.Object
						ϒurl           = λargs[1]
						ϒvideo_id      λ.Object
						ϒvideo_url     λ.Object
						ϒview_count    λ.Object
						ϒwebpage       λ.Object
					)
					ϒmobj = λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("id"))
					ϒdisplay_id = λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("display_id"))
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒdisplay_id)
					ϒvideo_url = λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewStr("video_url\\s*:\\s*'([^']+)'"), ϒwebpage, λ.NewStr("video URL"))
					ϒtitle = λ.Cal(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewList(
						λ.NewStr("<div[^>]+\\bclass=\"block_header\"[^>]*>\\s*<h1>([^<]+)<"),
						λ.NewStr("<title>(.*?)\\s*-\\s*(?:XXXYMovies\\.com|XXX\\s+Movies)</title>"),
					), ϒwebpage, λ.NewStr("title"))
					ϒthumbnail = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr("preview_url\\s*:\\s*'([^']+)'"),
						ϒwebpage,
						λ.NewStr("thumbnail"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					})
					ϒcategories = λ.Cal(λ.GetAttr(λ.Call(λ.GetAttr(ϒself, "_html_search_meta", nil), λ.NewArgs(
						λ.NewStr("keywords"),
						ϒwebpage,
						λ.NewStr("categories"),
					), λ.KWArgs{
						{Name: "default", Value: λ.NewStr("")},
					}), "split", nil), λ.NewStr(","))
					ϒduration = λ.Cal(ϒparse_duration, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr("<span>Duration:</span>\\s*(\\d+:\\d+)"),
						ϒwebpage,
						λ.NewStr("duration"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					}))
					ϒview_count = λ.Cal(ϒint_or_none, λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
						λ.NewStr("<div class=\"video_views\">\\s*(\\d+)"),
						ϒwebpage,
						λ.NewStr("view count"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					}))
					ϒlike_count = λ.Cal(ϒint_or_none, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr(">\\s*Likes? <b>\\((\\d+)\\)"),
						ϒwebpage,
						λ.NewStr("like count"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					}))
					ϒdislike_count = λ.Cal(ϒint_or_none, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr(">\\s*Dislike <b>\\((\\d+)\\)</b>"),
						ϒwebpage,
						λ.NewStr("dislike count"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					}))
					ϒage_limit = λ.Cal(λ.GetAttr(ϒself, "_rta_search", nil), ϒwebpage)
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):            ϒvideo_id,
						λ.NewStr("display_id"):    ϒdisplay_id,
						λ.NewStr("url"):           ϒvideo_url,
						λ.NewStr("title"):         ϒtitle,
						λ.NewStr("thumbnail"):     ϒthumbnail,
						λ.NewStr("categories"):    ϒcategories,
						λ.NewStr("duration"):      ϒduration,
						λ.NewStr("view_count"):    ϒview_count,
						λ.NewStr("like_count"):    ϒlike_count,
						λ.NewStr("dislike_count"): ϒdislike_count,
						λ.NewStr("age_limit"):     ϒage_limit,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TEST"):         XXXYMoviesIE__TEST,
				λ.NewStr("_VALID_URL"):    XXXYMoviesIE__VALID_URL,
				λ.NewStr("_real_extract"): XXXYMoviesIE__real_extract,
			})
		}())
	})
}
