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
 * fivemin/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/fivemin.py
 */

package fivemin

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	FiveMinIE     λ.Object
	InfoExtractor λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		FiveMinIE = λ.Cal(λ.TypeType, λ.NewStr("FiveMinIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				FiveMinIE_IE_NAME       λ.Object
				FiveMinIE__TESTS        λ.Object
				FiveMinIE__VALID_URL    λ.Object
				FiveMinIE__real_extract λ.Object
			)
			FiveMinIE_IE_NAME = λ.NewStr("5min")
			FiveMinIE__VALID_URL = λ.NewStr("(?:5min:|https?://(?:[^/]*?5min\\.com/|delivery\\.vidible\\.tv/aol)(?:(?:Scripts/PlayerSeed\\.js|playerseed/?)?\\?.*?playList=)?)(?P<id>\\d+)")
			FiveMinIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://pshared.5min.com/Scripts/PlayerSeed.js?sid=281&width=560&height=345&playList=518013791"),
					λ.NewStr("md5"): λ.NewStr("4f7b0b79bf1a470e5004f7112385941d"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("518013791"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("iPad Mini with Retina Display Review"),
						λ.NewStr("description"): λ.NewStr("iPad mini with Retina Display review"),
						λ.NewStr("duration"):    λ.NewInt(177),
						λ.NewStr("uploader"):    λ.NewStr("engadget"),
						λ.NewStr("upload_date"): λ.NewStr("20131115"),
						λ.NewStr("timestamp"):   λ.NewInt(1384515288),
					}),
					λ.NewStr("params"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("skip_download"): λ.True,
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("5min:518086247"),
					λ.NewStr("md5"): λ.NewStr("e539a9dd682c288ef5a498898009f69e"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):       λ.NewStr("518086247"),
						λ.NewStr("ext"):      λ.NewStr("mp4"),
						λ.NewStr("title"):    λ.NewStr("How to Make a Next-Level Fruit Salad"),
						λ.NewStr("duration"): λ.NewInt(184),
					}),
					λ.NewStr("skip"): λ.NewStr("no longer available"),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://embed.5min.com/518726732/"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://delivery.vidible.tv/aol?playList=518013791"),
					λ.NewStr("only_matching"): λ.True,
				}),
			)
			FiveMinIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒself     = λargs[0]
						ϒurl      = λargs[1]
						ϒvideo_id λ.Object
					)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					return λ.Cal(λ.GetAttr(ϒself, "url_result", nil), λ.Mod(λ.NewStr("aol-video:%s"), ϒvideo_id))
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("IE_NAME"):       FiveMinIE_IE_NAME,
				λ.NewStr("_TESTS"):        FiveMinIE__TESTS,
				λ.NewStr("_VALID_URL"):    FiveMinIE__VALID_URL,
				λ.NewStr("_real_extract"): FiveMinIE__real_extract,
			})
		}())
	})
}
