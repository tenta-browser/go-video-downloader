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
 * tube8/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/tube8.py
 */

package tube8

import (
	Ωkeezmovies "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/keezmovies"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	KeezMoviesIE λ.Object
	Tube8IE      λ.Object
	ϒint_or_none λ.Object
	ϒstr_to_int  λ.Object
)

func init() {
	λ.InitModule(func() {
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒstr_to_int = Ωutils.ϒstr_to_int
		KeezMoviesIE = Ωkeezmovies.KeezMoviesIE
		Tube8IE = λ.Cal(λ.TypeType, λ.StrLiteral("Tube8IE"), λ.NewTuple(KeezMoviesIE), func() λ.Dict {
			var (
				Tube8IE__VALID_URL λ.Object
			)
			Tube8IE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?tube8\\.com/(?:[^/]+/)+(?P<display_id>[^/]+)/(?P<id>\\d+)")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": Tube8IE__VALID_URL,
			})
		}())
	})
}
