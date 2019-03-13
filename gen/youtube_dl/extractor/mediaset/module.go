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
 * mediaset/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/mediaset.py
 */

package mediaset

import (
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωtheplatform "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/theplatform"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError                λ.Object
	MediasetIE                    λ.Object
	ThePlatformBaseIE             λ.Object
	ϒcompat_parse_qs              λ.Object
	ϒcompat_str                   λ.Object
	ϒcompat_urllib_parse_urlparse λ.Object
	ϒint_or_none                  λ.Object
	ϒupdate_url_query             λ.Object
)

func init() {
	λ.InitModule(func() {
		ThePlatformBaseIE = Ωtheplatform.ThePlatformBaseIE
		ϒcompat_parse_qs = Ωcompat.ϒcompat_parse_qs
		ϒcompat_str = Ωcompat.ϒcompat_str
		ϒcompat_urllib_parse_urlparse = Ωcompat.ϒcompat_urllib_parse_urlparse
		ExtractorError = Ωutils.ExtractorError
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒupdate_url_query = Ωutils.ϒupdate_url_query
		MediasetIE = λ.Cal(λ.TypeType, λ.NewStr("MediasetIE"), λ.NewTuple(ThePlatformBaseIE), func() λ.Dict {
			var (
				MediasetIE__VALID_URL λ.Object
			)
			MediasetIE__VALID_URL = λ.NewStr("(?x)\n                    (?:\n                        mediaset:|\n                        https?://\n                            (?:(?:www|static3)\\.)?mediasetplay\\.mediaset\\.it/\n                            (?:\n                                (?:video|on-demand)/(?:[^/]+/)+[^/]+_|\n                                player/index\\.html\\?.*?\\bprogramGuid=\n                            )\n                    )(?P<id>[0-9A-Z]{16})\n                    ")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): MediasetIE__VALID_URL,
			})
		}())
	})
}
