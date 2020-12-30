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
 * engadget/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/engadget.py
 */

package engadget

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	EngadgetIE    λ.Object
	InfoExtractor λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		EngadgetIE = λ.Cal(λ.TypeType, λ.StrLiteral("EngadgetIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				EngadgetIE__VALID_URL    λ.Object
				EngadgetIE__real_extract λ.Object
			)
			EngadgetIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?engadget\\.com/video/(?P<id>[^/?#]+)")
			EngadgetIE__real_extract = λ.NewFunction("_real_extract",
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
					ϒvideo_id = λ.Calm(ϒself, "_match_id", ϒurl)
					return λ.Calm(ϒself, "url_result", λ.Mod(λ.StrLiteral("aol-video:%s"), ϒvideo_id))
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL":    EngadgetIE__VALID_URL,
				"_real_extract": EngadgetIE__real_extract,
			})
		}())
	})
}