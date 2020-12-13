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
 * sportbox/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/sportbox.py
 */

package sportbox

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor  λ.Object
	SportBoxIE     λ.Object
	ϒdetermine_ext λ.Object
	ϒint_or_none   λ.Object
	ϒjs_to_json    λ.Object
	ϒmerge_dicts   λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒdetermine_ext = Ωutils.ϒdetermine_ext
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒjs_to_json = Ωutils.ϒjs_to_json
		ϒmerge_dicts = Ωutils.ϒmerge_dicts
		SportBoxIE = λ.Cal(λ.TypeType, λ.StrLiteral("SportBoxIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				SportBoxIE__VALID_URL λ.Object
			)
			SportBoxIE__VALID_URL = λ.StrLiteral("https?://(?:news\\.sportbox|matchtv)\\.ru/vdl/player(?:/[^/]+/|\\?.*?\\bn?id=)(?P<id>\\d+)")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": SportBoxIE__VALID_URL,
			})
		}())
	})
}
