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
 * ir90tv/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/ir90tv.py
 */

package ir90tv

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor λ.Object
	Ir90TvIE      λ.Object
	ϒremove_start λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒremove_start = Ωutils.ϒremove_start
		Ir90TvIE = λ.Cal(λ.TypeType, λ.NewStr("Ir90TvIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				Ir90TvIE__TESTS        λ.Object
				Ir90TvIE__VALID_URL    λ.Object
				Ir90TvIE__real_extract λ.Object
			)
			Ir90TvIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?90tv\\.ir/video/(?P<id>[0-9]+)/.*")
			Ir90TvIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://90tv.ir/video/95719/%D8%B4%D8%A7%DB%8C%D8%B9%D8%A7%D8%AA-%D9%86%D9%82%D9%84-%D9%88-%D8%A7%D9%86%D8%AA%D9%82%D8%A7%D9%84%D8%A7%D8%AA-%D9%85%D9%87%D9%85-%D9%81%D9%88%D8%AA%D8%A8%D8%A7%D9%84-%D8%A7%D8%B1%D9%88%D9%BE%D8%A7-940218"),
					λ.NewStr("md5"): λ.NewStr("411dbd94891381960cb9e13daa47a869"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):        λ.NewStr("95719"),
						λ.NewStr("ext"):       λ.NewStr("mp4"),
						λ.NewStr("title"):     λ.NewStr("شایعات نقل و انتقالات مهم فوتبال اروپا 94/02/18"),
						λ.NewStr("thumbnail"): λ.NewStr("re:^https?://.*\\.jpg$"),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://www.90tv.ir/video/95719/%D8%B4%D8%A7%DB%8C%D8%B9%D8%A7%D8%AA-%D9%86%D9%82%D9%84-%D9%88-%D8%A7%D9%86%D8%AA%D9%82%D8%A7%D9%84%D8%A7%D8%AA-%D9%85%D9%87%D9%85-%D9%81%D9%88%D8%AA%D8%A8%D8%A7%D9%84-%D8%A7%D8%B1%D9%88%D9%BE%D8%A7-940218"),
					λ.NewStr("only_matching"): λ.True,
				}),
			)
			Ir90TvIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒself      = λargs[0]
						ϒthumbnail λ.Object
						ϒtitle     λ.Object
						ϒurl       = λargs[1]
						ϒvideo_id  λ.Object
						ϒvideo_url λ.Object
						ϒwebpage   λ.Object
					)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒvideo_id)
					ϒtitle = λ.Cal(ϒremove_start, λ.Cal(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewStr("<title>([^<]+)</title>"), ϒwebpage, λ.NewStr("title")), λ.NewStr("90tv.ir :: "))
					ϒvideo_url = λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewStr("<source[^>]+src=\"([^\"]+)\""), ϒwebpage, λ.NewStr("video url"))
					ϒthumbnail = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr("poster=\"([^\"]+)\""),
						ϒwebpage,
						λ.NewStr("thumbnail url"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					})
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("url"):       ϒvideo_url,
						λ.NewStr("id"):        ϒvideo_id,
						λ.NewStr("title"):     ϒtitle,
						λ.NewStr("video_url"): ϒvideo_url,
						λ.NewStr("thumbnail"): ϒthumbnail,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TESTS"):        Ir90TvIE__TESTS,
				λ.NewStr("_VALID_URL"):    Ir90TvIE__VALID_URL,
				λ.NewStr("_real_extract"): Ir90TvIE__real_extract,
			})
		}())
	})
}
