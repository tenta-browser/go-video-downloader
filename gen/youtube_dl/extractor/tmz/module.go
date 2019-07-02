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
 * tmz/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/tmz.py
 */

package tmz

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor λ.Object
	TMZArticleIE  λ.Object
	TMZIE         λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		TMZIE = λ.Cal(λ.TypeType, λ.NewStr("TMZIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				TMZIE__TESTS        λ.Object
				TMZIE__VALID_URL    λ.Object
				TMZIE__real_extract λ.Object
			)
			TMZIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?tmz\\.com/videos/(?P<id>[^/?#]+)")
			TMZIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.tmz.com/videos/0_okj015ty/"),
					λ.NewStr("md5"): λ.NewStr("4d22a51ef205b6c06395d8394f72d560"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("0_okj015ty"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Kim Kardashian's Boobs Unlock a Mystery!"),
						λ.NewStr("description"): λ.NewStr("Did Kim Kardasain try to one-up Khloe by one-upping Kylie???  Or is she just showing off her amazing boobs?"),
						λ.NewStr("timestamp"):   λ.NewInt(1394747163),
						λ.NewStr("uploader_id"): λ.NewStr("batchUser"),
						λ.NewStr("upload_date"): λ.NewStr("20140313"),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://www.tmz.com/videos/0-cegprt2p/"),
					λ.NewStr("only_matching"): λ.True,
				}),
			)
			TMZIE__real_extract = λ.NewFunction("_real_extract",
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
					ϒvideo_id = λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl), "replace", nil), λ.NewStr("-"), λ.NewStr("_"))
					return λ.Cal(λ.GetAttr(ϒself, "url_result", nil), λ.Mod(λ.NewStr("kaltura:591531:%s"), ϒvideo_id), λ.NewStr("Kaltura"), ϒvideo_id)
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TESTS"):        TMZIE__TESTS,
				λ.NewStr("_VALID_URL"):    TMZIE__VALID_URL,
				λ.NewStr("_real_extract"): TMZIE__real_extract,
			})
		}())
		TMZArticleIE = λ.Cal(λ.TypeType, λ.NewStr("TMZArticleIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				TMZArticleIE__VALID_URL λ.Object
			)
			TMZArticleIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?tmz\\.com/\\d{4}/\\d{2}/\\d{2}/(?P<id>[^/]+)/?")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): TMZArticleIE__VALID_URL,
			})
		}())
	})
}
