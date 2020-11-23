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
 * soundcloud/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/soundcloud.py
 */

package soundcloud

import (
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError                λ.Object
	InfoExtractor                 λ.Object
	KNOWN_EXTENSIONS              λ.Object
	SearchInfoExtractor           λ.Object
	SoundcloudEmbedIE             λ.Object
	SoundcloudIE                  λ.Object
	SoundcloudPagedPlaylistBaseIE λ.Object
	SoundcloudPlaylistBaseIE      λ.Object
	SoundcloudPlaylistIE          λ.Object
	SoundcloudSearchIE            λ.Object
	SoundcloudSetIE               λ.Object
	SoundcloudTrackStationIE      λ.Object
	SoundcloudUserIE              λ.Object
	ϒcompat_HTTPError             λ.Object
	ϒcompat_kwargs                λ.Object
	ϒcompat_str                   λ.Object
	ϒerror_to_compat_str          λ.Object
	ϒfloat_or_none                λ.Object
	ϒint_or_none                  λ.Object
	ϒmimetype2ext                 λ.Object
	ϒstr_or_none                  λ.Object
	ϒtry_get                      λ.Object
	ϒunified_timestamp            λ.Object
	ϒupdate_url_query             λ.Object
	ϒurl_or_none                  λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		SearchInfoExtractor = Ωcommon.SearchInfoExtractor
		ϒcompat_HTTPError = Ωcompat.ϒcompat_HTTPError
		ϒcompat_kwargs = Ωcompat.ϒcompat_kwargs
		ϒcompat_str = Ωcompat.ϒcompat_str
		ϒerror_to_compat_str = Ωutils.ϒerror_to_compat_str
		ExtractorError = Ωutils.ExtractorError
		ϒfloat_or_none = Ωutils.ϒfloat_or_none
		ϒint_or_none = Ωutils.ϒint_or_none
		KNOWN_EXTENSIONS = Ωutils.KNOWN_EXTENSIONS
		ϒmimetype2ext = Ωutils.ϒmimetype2ext
		ϒstr_or_none = Ωutils.ϒstr_or_none
		ϒtry_get = Ωutils.ϒtry_get
		ϒunified_timestamp = Ωutils.ϒunified_timestamp
		ϒupdate_url_query = Ωutils.ϒupdate_url_query
		ϒurl_or_none = Ωutils.ϒurl_or_none
		SoundcloudEmbedIE = λ.Cal(λ.TypeType, λ.StrLiteral("SoundcloudEmbedIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				SoundcloudEmbedIE__VALID_URL λ.Object
			)
			SoundcloudEmbedIE__VALID_URL = λ.StrLiteral("https?://(?:w|player|p)\\.soundcloud\\.com/player/?.*?\\burl=(?P<id>.+)")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": SoundcloudEmbedIE__VALID_URL,
			})
		}())
		SoundcloudIE = λ.Cal(λ.TypeType, λ.StrLiteral("SoundcloudIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				SoundcloudIE__VALID_URL λ.Object
			)
			SoundcloudIE__VALID_URL = λ.StrLiteral("(?x)^(?:https?://)?\n                    (?:(?:(?:www\\.|m\\.)?soundcloud\\.com/\n                            (?!stations/track)\n                            (?P<uploader>[\\w\\d-]+)/\n                            (?!(?:tracks|albums|sets(?:/.+?)?|reposts|likes|spotlight)/?(?:$|[?#]))\n                            (?P<title>[\\w\\d-]+)/?\n                            (?P<token>[^?]+?)?(?:[?].*)?$)\n                       |(?:api(?:-v2)?\\.soundcloud\\.com/tracks/(?P<track_id>\\d+)\n                          (?:/?\\?secret_token=(?P<secret_token>[^&]+))?)\n                    )\n                    ")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": SoundcloudIE__VALID_URL,
			})
		}())
		SoundcloudPlaylistBaseIE = λ.Cal(λ.TypeType, λ.StrLiteral("SoundcloudPlaylistBaseIE"), λ.NewTuple(SoundcloudIE), func() λ.Dict {

			return λ.ClassDictLiteral(map[λ.Object]λ.Object{})
		}())
		SoundcloudSetIE = λ.Cal(λ.TypeType, λ.StrLiteral("SoundcloudSetIE"), λ.NewTuple(SoundcloudPlaylistBaseIE), func() λ.Dict {
			var (
				SoundcloudSetIE__VALID_URL λ.Object
			)
			SoundcloudSetIE__VALID_URL = λ.StrLiteral("https?://(?:(?:www|m)\\.)?soundcloud\\.com/(?P<uploader>[\\w\\d-]+)/sets/(?P<slug_title>[\\w\\d-]+)(?:/(?P<token>[^?/]+))?")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": SoundcloudSetIE__VALID_URL,
			})
		}())
		SoundcloudPagedPlaylistBaseIE = λ.Cal(λ.TypeType, λ.StrLiteral("SoundcloudPagedPlaylistBaseIE"), λ.NewTuple(SoundcloudIE), func() λ.Dict {

			return λ.ClassDictLiteral(map[λ.Object]λ.Object{})
		}())
		SoundcloudUserIE = λ.Cal(λ.TypeType, λ.StrLiteral("SoundcloudUserIE"), λ.NewTuple(SoundcloudPagedPlaylistBaseIE), func() λ.Dict {
			var (
				SoundcloudUserIE__VALID_URL λ.Object
			)
			SoundcloudUserIE__VALID_URL = λ.StrLiteral("(?x)\n                        https?://\n                            (?:(?:www|m)\\.)?soundcloud\\.com/\n                            (?P<user>[^/]+)\n                            (?:/\n                                (?P<rsrc>tracks|albums|sets|reposts|likes|spotlight)\n                            )?\n                            /?(?:[?#].*)?$\n                    ")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": SoundcloudUserIE__VALID_URL,
			})
		}())
		SoundcloudTrackStationIE = λ.Cal(λ.TypeType, λ.StrLiteral("SoundcloudTrackStationIE"), λ.NewTuple(SoundcloudPagedPlaylistBaseIE), func() λ.Dict {
			var (
				SoundcloudTrackStationIE__VALID_URL λ.Object
			)
			SoundcloudTrackStationIE__VALID_URL = λ.StrLiteral("https?://(?:(?:www|m)\\.)?soundcloud\\.com/stations/track/[^/]+/(?P<id>[^/?#&]+)")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": SoundcloudTrackStationIE__VALID_URL,
			})
		}())
		SoundcloudPlaylistIE = λ.Cal(λ.TypeType, λ.StrLiteral("SoundcloudPlaylistIE"), λ.NewTuple(SoundcloudPlaylistBaseIE), func() λ.Dict {
			var (
				SoundcloudPlaylistIE__VALID_URL λ.Object
			)
			SoundcloudPlaylistIE__VALID_URL = λ.StrLiteral("https?://api(?:-v2)?\\.soundcloud\\.com/playlists/(?P<id>[0-9]+)(?:/?\\?secret_token=(?P<token>[^&]+?))?$")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": SoundcloudPlaylistIE__VALID_URL,
			})
		}())
		SoundcloudSearchIE = λ.Cal(λ.TypeType, λ.StrLiteral("SoundcloudSearchIE"), λ.NewTuple(
			SearchInfoExtractor,
			SoundcloudIE,
		), func() λ.Dict {
			var (
				SoundcloudSearchIE__SEARCH_KEY λ.Object
			)
			SoundcloudSearchIE__SEARCH_KEY = λ.StrLiteral("scsearch")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_SEARCH_KEY": SoundcloudSearchIE__SEARCH_KEY,
			})
		}())
	})
}
