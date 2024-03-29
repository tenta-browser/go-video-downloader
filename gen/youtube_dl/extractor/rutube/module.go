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
 * rutube/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/rutube.py
 */

package rutube

import (
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor                 λ.Object
	RutubeBaseIE                  λ.Object
	RutubeChannelIE               λ.Object
	RutubeEmbedIE                 λ.Object
	RutubeIE                      λ.Object
	RutubeMovieIE                 λ.Object
	RutubePersonIE                λ.Object
	RutubePlaylistBaseIE          λ.Object
	RutubePlaylistIE              λ.Object
	ϒcompat_parse_qs              λ.Object
	ϒcompat_str                   λ.Object
	ϒcompat_urllib_parse_urlparse λ.Object
	ϒdetermine_ext                λ.Object
	ϒint_or_none                  λ.Object
	ϒtry_get                      λ.Object
	ϒunified_timestamp            λ.Object
	ϒurl_or_none                  λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_str = Ωcompat.ϒcompat_str
		ϒcompat_parse_qs = Ωcompat.ϒcompat_parse_qs
		ϒcompat_urllib_parse_urlparse = Ωcompat.ϒcompat_urllib_parse_urlparse
		ϒdetermine_ext = Ωutils.ϒdetermine_ext
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒtry_get = Ωutils.ϒtry_get
		ϒunified_timestamp = Ωutils.ϒunified_timestamp
		ϒurl_or_none = Ωutils.ϒurl_or_none
		RutubeBaseIE = λ.Cal(λ.TypeType, λ.StrLiteral("RutubeBaseIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {

			return λ.ClassDictLiteral(map[λ.Object]λ.Object{})
		}())
		RutubeIE = λ.Cal(λ.TypeType, λ.StrLiteral("RutubeIE"), λ.NewTuple(RutubeBaseIE), func() λ.Dict {
			var (
				RutubeIE__VALID_URL λ.Object
				RutubeIE_suitable   λ.Object
			)
			RutubeIE__VALID_URL = λ.StrLiteral("https?://rutube\\.ru/(?:video|(?:play/)?embed)/(?P<id>[\\da-z]{32})")
			RutubeIE_suitable = λ.NewFunction("suitable",
				[]λ.Param{
					{Name: "cls"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒcls = λargs[0]
						ϒurl = λargs[1]
					)
					return func() λ.Object {
						if λ.IsTrue(λ.Calm(RutubePlaylistIE, "suitable", ϒurl)) {
							return λ.False
						} else {
							return λ.Calm(λ.Cal(λ.SuperType, RutubeIE, ϒcls), "suitable", ϒurl)
						}
					}()
				})
			RutubeIE_suitable = λ.Cal(λ.ClassMethodType, RutubeIE_suitable)
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": RutubeIE__VALID_URL,
				"suitable":   RutubeIE_suitable,
			})
		}())
		RutubeEmbedIE = λ.Cal(λ.TypeType, λ.StrLiteral("RutubeEmbedIE"), λ.NewTuple(RutubeBaseIE), func() λ.Dict {
			var (
				RutubeEmbedIE__VALID_URL λ.Object
			)
			RutubeEmbedIE__VALID_URL = λ.StrLiteral("https?://rutube\\.ru/(?:video|play)/embed/(?P<id>[0-9]+)")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": RutubeEmbedIE__VALID_URL,
			})
		}())
		RutubePlaylistBaseIE = λ.Cal(λ.TypeType, λ.StrLiteral("RutubePlaylistBaseIE"), λ.NewTuple(RutubeBaseIE), func() λ.Dict {

			return λ.ClassDictLiteral(map[λ.Object]λ.Object{})
		}())
		RutubeChannelIE = λ.Cal(λ.TypeType, λ.StrLiteral("RutubeChannelIE"), λ.NewTuple(RutubePlaylistBaseIE), func() λ.Dict {
			var (
				RutubeChannelIE__VALID_URL λ.Object
			)
			RutubeChannelIE__VALID_URL = λ.StrLiteral("https?://rutube\\.ru/tags/video/(?P<id>\\d+)")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": RutubeChannelIE__VALID_URL,
			})
		}())
		RutubeMovieIE = λ.Cal(λ.TypeType, λ.StrLiteral("RutubeMovieIE"), λ.NewTuple(RutubePlaylistBaseIE), func() λ.Dict {
			var (
				RutubeMovieIE__VALID_URL λ.Object
			)
			RutubeMovieIE__VALID_URL = λ.StrLiteral("https?://rutube\\.ru/metainfo/tv/(?P<id>\\d+)")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": RutubeMovieIE__VALID_URL,
			})
		}())
		RutubePersonIE = λ.Cal(λ.TypeType, λ.StrLiteral("RutubePersonIE"), λ.NewTuple(RutubePlaylistBaseIE), func() λ.Dict {
			var (
				RutubePersonIE__VALID_URL λ.Object
			)
			RutubePersonIE__VALID_URL = λ.StrLiteral("https?://rutube\\.ru/video/person/(?P<id>\\d+)")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": RutubePersonIE__VALID_URL,
			})
		}())
		RutubePlaylistIE = λ.Cal(λ.TypeType, λ.StrLiteral("RutubePlaylistIE"), λ.NewTuple(RutubePlaylistBaseIE), func() λ.Dict {
			var (
				RutubePlaylistIE__VALID_URL λ.Object
				RutubePlaylistIE_suitable   λ.Object
			)
			RutubePlaylistIE__VALID_URL = λ.StrLiteral("https?://rutube\\.ru/(?:video|(?:play/)?embed)/[\\da-z]{32}/\\?.*?\\bpl_id=(?P<id>\\d+)")
			RutubePlaylistIE_suitable = λ.NewFunction("suitable",
				[]λ.Param{
					{Name: "cls"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒcls    = λargs[0]
						ϒparams λ.Object
						ϒurl    = λargs[1]
					)
					if !λ.IsTrue(λ.Calm(λ.Cal(λ.SuperType, RutubePlaylistIE, ϒcls), "suitable", ϒurl)) {
						return λ.False
					}
					ϒparams = λ.Cal(ϒcompat_parse_qs, λ.GetAttr(λ.Cal(ϒcompat_urllib_parse_urlparse, ϒurl), "query", nil))
					return func() λ.Object {
						if λv := λ.GetItem(λ.Calm(ϒparams, "get", λ.StrLiteral("pl_type"), λ.NewList(λ.None)), λ.IntLiteral(0)); !λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Cal(ϒint_or_none, λ.GetItem(λ.Calm(ϒparams, "get", λ.StrLiteral("pl_id"), λ.NewList(λ.None)), λ.IntLiteral(0)))
						}
					}()
				})
			RutubePlaylistIE_suitable = λ.Cal(λ.ClassMethodType, RutubePlaylistIE_suitable)
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": RutubePlaylistIE__VALID_URL,
				"suitable":   RutubePlaylistIE_suitable,
			})
		}())
	})
}
