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
 * viqeo/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/viqeo.py
 */

package viqeo

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor λ.Object
	ViqeoIE       λ.Object
	ϒint_or_none  λ.Object
	ϒstr_or_none  λ.Object
	ϒurl_or_none  λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒstr_or_none = Ωutils.ϒstr_or_none
		ϒurl_or_none = Ωutils.ϒurl_or_none
		ViqeoIE = λ.Cal(λ.TypeType, λ.StrLiteral("ViqeoIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				ViqeoIE__VALID_URL λ.Object
			)
			ViqeoIE__VALID_URL = λ.StrLiteral("(?x)\n                        (?:\n                            viqeo:|\n                            https?://cdn\\.viqeo\\.tv/embed/*\\?.*?\\bvid=|\n                            https?://api\\.viqeo\\.tv/v\\d+/data/startup?.*?\\bvideo(?:%5B%5D|\\[\\])=\n                        )\n                        (?P<id>[\\da-f]+)\n                    ")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": ViqeoIE__VALID_URL,
			})
		}())
	})
}
