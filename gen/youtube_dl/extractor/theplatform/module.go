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
 * theplatform/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/theplatform.py
 */

package theplatform

import (
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωadobepass "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/adobepass"
	Ωonce "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/once"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	AdobePassIE                   λ.Object
	ExtractorError                λ.Object
	OnceIE                        λ.Object
	ThePlatformBaseIE             λ.Object
	ThePlatformFeedIE             λ.Object
	ThePlatformIE                 λ.Object
	ϒcompat_parse_qs              λ.Object
	ϒcompat_urllib_parse_urlparse λ.Object
	ϒdetermine_ext                λ.Object
	ϒfind_xpath_attr              λ.Object
	ϒfloat_or_none                λ.Object
	ϒint_or_none                  λ.Object
	ϒmimetype2ext                 λ.Object
	ϒsanitized_Request            λ.Object
	ϒunsmuggle_url                λ.Object
	ϒupdate_url_query             λ.Object
)

func init() {
	λ.InitModule(func() {
		OnceIE = Ωonce.OnceIE
		AdobePassIE = Ωadobepass.AdobePassIE
		ϒcompat_parse_qs = Ωcompat.ϒcompat_parse_qs
		ϒcompat_urllib_parse_urlparse = Ωcompat.ϒcompat_urllib_parse_urlparse
		ϒdetermine_ext = Ωutils.ϒdetermine_ext
		ExtractorError = Ωutils.ExtractorError
		ϒfloat_or_none = Ωutils.ϒfloat_or_none
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒsanitized_Request = Ωutils.ϒsanitized_Request
		ϒunsmuggle_url = Ωutils.ϒunsmuggle_url
		ϒupdate_url_query = Ωutils.ϒupdate_url_query
		ϒmimetype2ext = Ωutils.ϒmimetype2ext
		ϒfind_xpath_attr = Ωutils.ϒfind_xpath_attr
		ThePlatformBaseIE = λ.Cal(λ.TypeType, λ.NewStr("ThePlatformBaseIE"), λ.NewTuple(OnceIE), func() λ.Dict {

			return λ.NewDictWithTable(map[λ.Object]λ.Object{})
		}())
		ThePlatformIE = λ.Cal(λ.TypeType, λ.NewStr("ThePlatformIE"), λ.NewTuple(
			ThePlatformBaseIE,
			AdobePassIE,
		), func() λ.Dict {
			var (
				ThePlatformIE__VALID_URL λ.Object
			)
			ThePlatformIE__VALID_URL = λ.NewStr("(?x)\n        (?:https?://(?:link|player)\\.theplatform\\.com/[sp]/(?P<provider_id>[^/]+)/\n           (?:(?:(?:[^/]+/)+select/)?(?P<media>media/(?:guid/\\d+/)?)?|(?P<config>(?:[^/\\?]+/(?:swf|config)|onsite)/select/))?\n         |theplatform:)(?P<id>[^/\\?&]+)")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): ThePlatformIE__VALID_URL,
			})
		}())
		ThePlatformFeedIE = λ.Cal(λ.TypeType, λ.NewStr("ThePlatformFeedIE"), λ.NewTuple(ThePlatformBaseIE), func() λ.Dict {
			var (
				ThePlatformFeedIE__VALID_URL λ.Object
			)
			ThePlatformFeedIE__VALID_URL = λ.NewStr("https?://feed\\.theplatform\\.com/f/(?P<provider_id>[^/]+)/(?P<feed_id>[^?/]+)\\?(?:[^&]+&)*(?P<filter>by(?:Gui|I)d=(?P<id>[^&]+))")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): ThePlatformFeedIE__VALID_URL,
			})
		}())
	})
}
