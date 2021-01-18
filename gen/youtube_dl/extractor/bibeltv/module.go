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
 * bibeltv/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/bibeltv.py
 */

package bibeltv

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	BibelTVIE     λ.Object
	InfoExtractor λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		BibelTVIE = λ.Cal(λ.TypeType, λ.StrLiteral("BibelTVIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				BibelTVIE_BRIGHTCOVE_URL_TEMPLATE λ.Object
				BibelTVIE__VALID_URL              λ.Object
				BibelTVIE__real_extract           λ.Object
			)
			BibelTVIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?bibeltv\\.de/mediathek/videos/(?:crn/)?(?P<id>\\d+)")
			BibelTVIE_BRIGHTCOVE_URL_TEMPLATE = λ.StrLiteral("http://players.brightcove.net/5840105145001/default_default/index.html?videoId=ref:%s")
			BibelTVIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒcrn_id λ.Object
						ϒself   = λargs[0]
						ϒurl    = λargs[1]
					)
					ϒcrn_id = λ.Calm(ϒself, "_match_id", ϒurl)
					return λ.Calm(ϒself, "url_result", λ.Mod(λ.GetAttr(ϒself, "BRIGHTCOVE_URL_TEMPLATE", nil), ϒcrn_id), λ.StrLiteral("BrightcoveNew"))
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"BRIGHTCOVE_URL_TEMPLATE": BibelTVIE_BRIGHTCOVE_URL_TEMPLATE,
				"_VALID_URL":              BibelTVIE__VALID_URL,
				"_real_extract":           BibelTVIE__real_extract,
			})
		}())
	})
}
