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
 * mtv/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/mtv.py
 */

package mtv

import (
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError           λ.Object
	InfoExtractor            λ.Object
	MTVDEIE                  λ.Object
	MTVIE                    λ.Object
	MTVJapanIE               λ.Object
	MTVServicesEmbeddedIE    λ.Object
	MTVServicesInfoExtractor λ.Object
	MTVVideoIE               λ.Object
	RegexNotFoundError       λ.Object
	ϒcompat_str              λ.Object
	ϒcompat_xpath            λ.Object
	ϒfind_xpath_attr         λ.Object
	ϒfix_xml_ampersands      λ.Object
	ϒfloat_or_none           λ.Object
	ϒsanitized_Request       λ.Object
	ϒstrip_or_none           λ.Object
	ϒtry_get                 λ.Object
	ϒunescapeHTML            λ.Object
	ϒupdate_url_query        λ.Object
	ϒurl_basename            λ.Object
	ϒxpath_text              λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_str = Ωcompat.ϒcompat_str
		ϒcompat_xpath = Ωcompat.ϒcompat_xpath
		ExtractorError = Ωutils.ExtractorError
		ϒfind_xpath_attr = Ωutils.ϒfind_xpath_attr
		ϒfix_xml_ampersands = Ωutils.ϒfix_xml_ampersands
		ϒfloat_or_none = Ωutils.ϒfloat_or_none
		RegexNotFoundError = Ωutils.RegexNotFoundError
		ϒsanitized_Request = Ωutils.ϒsanitized_Request
		ϒstrip_or_none = Ωutils.ϒstrip_or_none
		ϒtry_get = Ωutils.ϒtry_get
		ϒunescapeHTML = Ωutils.ϒunescapeHTML
		ϒupdate_url_query = Ωutils.ϒupdate_url_query
		ϒurl_basename = Ωutils.ϒurl_basename
		ϒxpath_text = Ωutils.ϒxpath_text
		MTVServicesInfoExtractor = λ.Cal(λ.TypeType, λ.StrLiteral("MTVServicesInfoExtractor"), λ.NewTuple(InfoExtractor), func() λ.Dict {

			return λ.ClassDictLiteral(map[λ.Object]λ.Object{})
		}())
		MTVServicesEmbeddedIE = λ.Cal(λ.TypeType, λ.StrLiteral("MTVServicesEmbeddedIE"), λ.NewTuple(MTVServicesInfoExtractor), func() λ.Dict {
			var (
				MTVServicesEmbeddedIE__VALID_URL λ.Object
			)
			MTVServicesEmbeddedIE__VALID_URL = λ.StrLiteral("https?://media\\.mtvnservices\\.com/embed/(?P<mgid>.+?)(\\?|/|$)")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": MTVServicesEmbeddedIE__VALID_URL,
			})
		}())
		MTVIE = λ.Cal(λ.TypeType, λ.StrLiteral("MTVIE"), λ.NewTuple(MTVServicesInfoExtractor), func() λ.Dict {
			var (
				MTVIE__VALID_URL λ.Object
			)
			MTVIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?mtv\\.com/(?:video-clips|(?:full-)?episodes)/(?P<id>[^/?#.]+)")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": MTVIE__VALID_URL,
			})
		}())
		MTVJapanIE = λ.Cal(λ.TypeType, λ.StrLiteral("MTVJapanIE"), λ.NewTuple(MTVServicesInfoExtractor), func() λ.Dict {
			var (
				MTVJapanIE__VALID_URL λ.Object
			)
			MTVJapanIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?mtvjapan\\.com/videos/(?P<id>[0-9a-z]+)")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": MTVJapanIE__VALID_URL,
			})
		}())
		MTVVideoIE = λ.Cal(λ.TypeType, λ.StrLiteral("MTVVideoIE"), λ.NewTuple(MTVServicesInfoExtractor), func() λ.Dict {
			var (
				MTVVideoIE__VALID_URL λ.Object
			)
			MTVVideoIE__VALID_URL = λ.StrLiteral("(?x)^https?://\n        (?:(?:www\\.)?mtv\\.com/videos/.+?/(?P<videoid>[0-9]+)/[^/]+$|\n           m\\.mtv\\.com/videos/video\\.rbml\\?.*?id=(?P<mgid>[^&]+))")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": MTVVideoIE__VALID_URL,
			})
		}())
		MTVDEIE = λ.Cal(λ.TypeType, λ.StrLiteral("MTVDEIE"), λ.NewTuple(MTVServicesInfoExtractor), func() λ.Dict {
			var (
				MTVDEIE__VALID_URL λ.Object
			)
			MTVDEIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?mtv\\.de/(?:musik/videoclips|folgen|news)/(?P<id>[0-9a-z]+)")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": MTVDEIE__VALID_URL,
			})
		}())
	})
}
