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
 * mofosex/module.go: transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/mofosex.py
 */

package mofosex

import (
	Ωkeezmovies "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/keezmovies"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	KeezMoviesIE     λ.Object
	MofosexIE        λ.Object
	ϒint_or_none     λ.Object
	ϒstr_to_int      λ.Object
	ϒunified_strdate λ.Object
)

func init() {
	λ.InitModule(func() {
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒstr_to_int = Ωutils.ϒstr_to_int
		ϒunified_strdate = Ωutils.ϒunified_strdate
		KeezMoviesIE = Ωkeezmovies.KeezMoviesIE
		MofosexIE = λ.Cal(λ.TypeType, λ.NewStr("MofosexIE"), λ.NewTuple(KeezMoviesIE), func() λ.Dict {
			var (
				MofosexIE__TESTS        λ.Object
				MofosexIE__VALID_URL    λ.Object
				MofosexIE__real_extract λ.Object
			)
			MofosexIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?mofosex\\.com/videos/(?P<id>\\d+)/(?P<display_id>[^/?#&.]+)\\.html")
			MofosexIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.mofosex.com/videos/318131/amateur-teen-playing-and-masturbating-318131.html"),
					λ.NewStr("md5"): λ.NewStr("558fcdafbb63a87c019218d6e49daf8a"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):            λ.NewStr("318131"),
						λ.NewStr("display_id"):    λ.NewStr("amateur-teen-playing-and-masturbating-318131"),
						λ.NewStr("ext"):           λ.NewStr("mp4"),
						λ.NewStr("title"):         λ.NewStr("amateur teen playing and masturbating"),
						λ.NewStr("thumbnail"):     λ.NewStr("re:^https?://.*\\.jpg$"),
						λ.NewStr("upload_date"):   λ.NewStr("20121114"),
						λ.NewStr("view_count"):    λ.IntType,
						λ.NewStr("like_count"):    λ.IntType,
						λ.NewStr("dislike_count"): λ.IntType,
						λ.NewStr("age_limit"):     λ.NewInt(18),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://www.mofosex.com/videos/5018/japanese-teen-music-video.html"),
					λ.NewStr("only_matching"): λ.True,
				}),
			)
			MofosexIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdislike_count λ.Object
						ϒinfo          λ.Object
						ϒlike_count    λ.Object
						ϒself          = λargs[0]
						ϒupload_date   λ.Object
						ϒurl           = λargs[1]
						ϒview_count    λ.Object
						ϒwebpage       λ.Object
						τmp0           λ.Object
					)
					τmp0 = λ.Cal(λ.GetAttr(ϒself, "_extract_info", nil), ϒurl)
					ϒwebpage = λ.GetItem(τmp0, λ.NewInt(0))
					ϒinfo = λ.GetItem(τmp0, λ.NewInt(1))
					ϒview_count = λ.Cal(ϒstr_to_int, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr("VIEWS:</span>\\s*([\\d,.]+)"),
						ϒwebpage,
						λ.NewStr("view count"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					}))
					ϒlike_count = λ.Cal(ϒint_or_none, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr("id=[\"\\']amountLikes[\"\\'][^>]*>(\\d+)"),
						ϒwebpage,
						λ.NewStr("like count"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					}))
					ϒdislike_count = λ.Cal(ϒint_or_none, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr("id=[\"\\']amountDislikes[\"\\'][^>]*>(\\d+)"),
						ϒwebpage,
						λ.NewStr("like count"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					}))
					ϒupload_date = λ.Cal(ϒunified_strdate, λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
						λ.NewStr("Added:</span>([^<]+)"),
						ϒwebpage,
						λ.NewStr("upload date"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					}))
					λ.Cal(λ.GetAttr(ϒinfo, "update", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("view_count"):    ϒview_count,
						λ.NewStr("like_count"):    ϒlike_count,
						λ.NewStr("dislike_count"): ϒdislike_count,
						λ.NewStr("upload_date"):   ϒupload_date,
						λ.NewStr("thumbnail"):     λ.Cal(λ.GetAttr(ϒself, "_og_search_thumbnail", nil), ϒwebpage),
					}))
					return ϒinfo
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TESTS"):        MofosexIE__TESTS,
				λ.NewStr("_VALID_URL"):    MofosexIE__VALID_URL,
				λ.NewStr("_real_extract"): MofosexIE__real_extract,
			})
		}())
	})
}