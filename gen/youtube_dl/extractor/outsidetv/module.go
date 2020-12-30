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
 * outsidetv/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/outsidetv.py
 */

package outsidetv

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor λ.Object
	OutsideTVIE   λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		OutsideTVIE = λ.Cal(λ.TypeType, λ.StrLiteral("OutsideTVIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				OutsideTVIE__VALID_URL    λ.Object
				OutsideTVIE__real_extract λ.Object
			)
			OutsideTVIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?outsidetv\\.com/(?:[^/]+/)*?play/[a-zA-Z0-9]{8}/\\d+/\\d+/(?P<id>[a-zA-Z0-9]{8})")
			OutsideTVIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒjw_media_id λ.Object
						ϒself        = λargs[0]
						ϒurl         = λargs[1]
					)
					ϒjw_media_id = λ.Calm(ϒself, "_match_id", ϒurl)
					return λ.Calm(ϒself, "url_result", λ.Add(λ.StrLiteral("jwplatform:"), ϒjw_media_id), λ.StrLiteral("JWPlatform"), ϒjw_media_id)
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL":    OutsideTVIE__VALID_URL,
				"_real_extract": OutsideTVIE__real_extract,
			})
		}())
	})
}