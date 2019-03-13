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
 * eagleplatform/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/eagleplatform.py
 */

package eagleplatform

import (
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	EaglePlatformIE   λ.Object
	ExtractorError    λ.Object
	InfoExtractor     λ.Object
	ϒcompat_HTTPError λ.Object
	ϒint_or_none      λ.Object
	ϒunsmuggle_url    λ.Object
	ϒurl_or_none      λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_HTTPError = Ωcompat.ϒcompat_HTTPError
		ExtractorError = Ωutils.ExtractorError
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒunsmuggle_url = Ωutils.ϒunsmuggle_url
		ϒurl_or_none = Ωutils.ϒurl_or_none
		EaglePlatformIE = λ.Cal(λ.TypeType, λ.NewStr("EaglePlatformIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				EaglePlatformIE__VALID_URL λ.Object
			)
			EaglePlatformIE__VALID_URL = λ.NewStr("(?x)\n                    (?:\n                        eagleplatform:(?P<custom_host>[^/]+):|\n                        https?://(?P<host>.+?\\.media\\.eagleplatform\\.com)/index/player\\?.*\\brecord_id=\n                    )\n                    (?P<id>\\d+)\n                ")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): EaglePlatformIE__VALID_URL,
			})
		}())
	})
}
