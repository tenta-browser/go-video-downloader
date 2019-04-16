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
 * jwplatform/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/jwplatform.py
 */

package jwplatform

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor λ.Object
	JWPlatformIE  λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		JWPlatformIE = λ.Cal(λ.TypeType, λ.NewStr("JWPlatformIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				JWPlatformIE__TESTS        λ.Object
				JWPlatformIE__VALID_URL    λ.Object
				JWPlatformIE__real_extract λ.Object
			)
			JWPlatformIE__VALID_URL = λ.NewStr("(?:https?://(?:content\\.jwplatform|cdn\\.jwplayer)\\.com/(?:(?:feed|player|thumb|preview|video)s|jw6|v2/media)/|jwplatform:)(?P<id>[a-zA-Z0-9]{8})")
			JWPlatformIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://content.jwplatform.com/players/nPripu9l-ALJ3XQCI.js"),
					λ.NewStr("md5"): λ.NewStr("fa8899fa601eb7c83a64e9d568bdf325"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("nPripu9l"),
						λ.NewStr("ext"):         λ.NewStr("mov"),
						λ.NewStr("title"):       λ.NewStr("Big Buck Bunny Trailer"),
						λ.NewStr("description"): λ.NewStr("Big Buck Bunny is a short animated film by the Blender Institute. It is made using free and open source software."),
						λ.NewStr("upload_date"): λ.NewStr("20081127"),
						λ.NewStr("timestamp"):   λ.NewInt(1227796140),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://cdn.jwplayer.com/players/nPripu9l-ALJ3XQCI.js"),
					λ.NewStr("only_matching"): λ.True,
				}),
			)
			JWPlatformIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒjson_data λ.Object
						ϒself      = λargs[0]
						ϒurl       = λargs[1]
						ϒvideo_id  λ.Object
					)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒjson_data = λ.Cal(λ.GetAttr(ϒself, "_download_json", nil), λ.Add(λ.NewStr("https://cdn.jwplayer.com/v2/media/"), ϒvideo_id), ϒvideo_id)
					return λ.Cal(λ.GetAttr(ϒself, "_parse_jwplayer_data", nil), ϒjson_data, ϒvideo_id)
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TESTS"):        JWPlatformIE__TESTS,
				λ.NewStr("_VALID_URL"):    JWPlatformIE__VALID_URL,
				λ.NewStr("_real_extract"): JWPlatformIE__real_extract,
			})
		}())
	})
}
