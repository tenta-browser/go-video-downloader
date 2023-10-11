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
 * palcomp3/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/palcomp3.py
 */

package palcomp3

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor    λ.Object
	PalcoMP3ArtistIE λ.Object
	PalcoMP3BaseIE   λ.Object
	PalcoMP3IE       λ.Object
	PalcoMP3VideoIE  λ.Object
	ϒcompat_str      λ.Object
	ϒint_or_none     λ.Object
	ϒstr_or_none     λ.Object
	ϒtry_get         λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_str = Ωcompat.ϒcompat_str
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒstr_or_none = Ωutils.ϒstr_or_none
		ϒtry_get = Ωutils.ϒtry_get
		PalcoMP3BaseIE = λ.Cal(λ.TypeType, λ.StrLiteral("PalcoMP3BaseIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				PalcoMP3BaseIE__ARTIST_FIELDS_TMPL λ.Object
				PalcoMP3BaseIE__GQL_QUERY_TMPL     λ.Object
				PalcoMP3BaseIE__MUSIC_FIELDS       λ.Object
				PalcoMP3BaseIE__call_api           λ.Object
				PalcoMP3BaseIE__parse_music        λ.Object
				PalcoMP3BaseIE__real_extract       λ.Object
				PalcoMP3BaseIE__real_initialize    λ.Object
			)
			PalcoMP3BaseIE__GQL_QUERY_TMPL = λ.StrLiteral("{\n  artist(slug: \"%s\") {\n    %s\n  }\n}")
			PalcoMP3BaseIE__ARTIST_FIELDS_TMPL = λ.StrLiteral("music(slug: \"%%s\") {\n      %s\n    }")
			PalcoMP3BaseIE__MUSIC_FIELDS = λ.StrLiteral("duration\n      hls\n      mp3File\n      musicID\n      plays\n      title")
			PalcoMP3BaseIE__call_api = λ.NewFunction("_call_api",
				[]λ.Param{
					{Name: "self"},
					{Name: "artist_slug"},
					{Name: "artist_fields"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒartist_fields = λargs[2]
						ϒartist_slug   = λargs[1]
						ϒself          = λargs[0]
					)
					return λ.GetItem(λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
						λ.StrLiteral("https://www.palcomp3.com.br/graphql/"),
						ϒartist_slug,
					), λ.KWArgs{
						{Name: "query", Value: λ.DictLiteral(map[string]λ.Object{
							"query": λ.Mod(λ.GetAttr(ϒself, "_GQL_QUERY_TMPL", nil), λ.NewTuple(
								ϒartist_slug,
								ϒartist_fields,
							)),
						})},
					}), λ.StrLiteral("data"))
				})
			PalcoMP3BaseIE__parse_music = λ.NewFunction("_parse_music",
				[]λ.Param{
					{Name: "self"},
					{Name: "music"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒformats  λ.Object
						ϒhls_url  λ.Object
						ϒmp3_file λ.Object
						ϒmusic    = λargs[1]
						ϒmusic_id λ.Object
						ϒself     = λargs[0]
						ϒtitle    λ.Object
					)
					_ = ϒself
					ϒmusic_id = λ.Cal(ϒcompat_str, λ.GetItem(ϒmusic, λ.StrLiteral("musicID")))
					ϒtitle = λ.GetItem(ϒmusic, λ.StrLiteral("title"))
					ϒformats = λ.NewList()
					ϒhls_url = λ.Calm(ϒmusic, "get", λ.StrLiteral("hls"))
					if λ.IsTrue(ϒhls_url) {
						λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
							"url":      ϒhls_url,
							"protocol": λ.StrLiteral("m3u8_native"),
							"ext":      λ.StrLiteral("mp4"),
						}))
					}
					ϒmp3_file = λ.Calm(ϒmusic, "get", λ.StrLiteral("mp3File"))
					if λ.IsTrue(ϒmp3_file) {
						λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
							"url": ϒmp3_file,
						}))
					}
					return λ.DictLiteral(map[string]λ.Object{
						"id":         ϒmusic_id,
						"title":      ϒtitle,
						"formats":    ϒformats,
						"duration":   λ.Cal(ϒint_or_none, λ.Calm(ϒmusic, "get", λ.StrLiteral("duration"))),
						"view_count": λ.Cal(ϒint_or_none, λ.Calm(ϒmusic, "get", λ.StrLiteral("plays"))),
					})
				})
			PalcoMP3BaseIE__real_initialize = λ.NewFunction("_real_initialize",
				[]λ.Param{
					{Name: "self"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒself = λargs[0]
					)
					λ.SetAttr(ϒself, "_ARTIST_FIELDS_TMPL", λ.Mod(λ.GetAttr(ϒself, "_ARTIST_FIELDS_TMPL", nil), λ.GetAttr(ϒself, "_MUSIC_FIELDS", nil)))
					return λ.None
				})
			PalcoMP3BaseIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒartist_fields λ.Object
						ϒartist_slug   λ.Object
						ϒmusic         λ.Object
						ϒmusic_slug    λ.Object
						ϒself          = λargs[0]
						ϒurl           = λargs[1]
						τmp0           λ.Object
					)
					τmp0 = λ.UnpackIterable(λ.Calm(λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl), "groups"), 2)
					ϒartist_slug = λ.GetItem(τmp0, λ.IntLiteral(0))
					ϒmusic_slug = λ.GetItem(τmp0, λ.IntLiteral(1))
					ϒartist_fields = λ.Mod(λ.GetAttr(ϒself, "_ARTIST_FIELDS_TMPL", nil), ϒmusic_slug)
					ϒmusic = λ.GetItem(λ.GetItem(λ.Calm(ϒself, "_call_api", ϒartist_slug, ϒartist_fields), λ.StrLiteral("artist")), λ.StrLiteral("music"))
					return λ.Calm(ϒself, "_parse_music", ϒmusic)
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_ARTIST_FIELDS_TMPL": PalcoMP3BaseIE__ARTIST_FIELDS_TMPL,
				"_GQL_QUERY_TMPL":     PalcoMP3BaseIE__GQL_QUERY_TMPL,
				"_MUSIC_FIELDS":       PalcoMP3BaseIE__MUSIC_FIELDS,
				"_call_api":           PalcoMP3BaseIE__call_api,
				"_parse_music":        PalcoMP3BaseIE__parse_music,
				"_real_extract":       PalcoMP3BaseIE__real_extract,
				"_real_initialize":    PalcoMP3BaseIE__real_initialize,
			})
		}())
		PalcoMP3IE = λ.Cal(λ.TypeType, λ.StrLiteral("PalcoMP3IE"), λ.NewTuple(PalcoMP3BaseIE), func() λ.Dict {
			var (
				PalcoMP3IE_IE_NAME    λ.Object
				PalcoMP3IE__VALID_URL λ.Object
				PalcoMP3IE_suitable   λ.Object
			)
			PalcoMP3IE_IE_NAME = λ.StrLiteral("PalcoMP3:song")
			PalcoMP3IE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?palcomp3\\.com(?:\\.br)?/(?P<artist>[^/]+)/(?P<id>[^/?&#]+)")
			PalcoMP3IE_suitable = λ.NewFunction("suitable",
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
						if λ.IsTrue(λ.Calm(PalcoMP3VideoIE, "suitable", ϒurl)) {
							return λ.False
						} else {
							return λ.Calm(λ.Cal(λ.SuperType, PalcoMP3IE, ϒcls), "suitable", ϒurl)
						}
					}()
				})
			PalcoMP3IE_suitable = λ.Cal(λ.ClassMethodType, PalcoMP3IE_suitable)
			return λ.ClassDictLiteral(map[string]λ.Object{
				"IE_NAME":    PalcoMP3IE_IE_NAME,
				"_VALID_URL": PalcoMP3IE__VALID_URL,
				"suitable":   PalcoMP3IE_suitable,
			})
		}())
		PalcoMP3ArtistIE = λ.Cal(λ.TypeType, λ.StrLiteral("PalcoMP3ArtistIE"), λ.NewTuple(PalcoMP3BaseIE), func() λ.Dict {
			var (
				PalcoMP3ArtistIE__VALID_URL λ.Object
				PalcoMP3ArtistIE_suitable   λ.Object
			)
			PalcoMP3ArtistIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?palcomp3\\.com(?:\\.br)?/(?P<id>[^/?&#]+)")
			PalcoMP3ArtistIE_suitable = λ.NewFunction("suitable",
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
						if λ.IsTrue(λ.Cal(Ωre.ϒmatch, λ.GetAttr(PalcoMP3IE, "_VALID_URL", nil), ϒurl)) {
							return λ.False
						} else {
							return λ.Calm(λ.Cal(λ.SuperType, PalcoMP3ArtistIE, ϒcls), "suitable", ϒurl)
						}
					}()
				})
			PalcoMP3ArtistIE_suitable = λ.Cal(λ.ClassMethodType, PalcoMP3ArtistIE_suitable)
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": PalcoMP3ArtistIE__VALID_URL,
				"suitable":   PalcoMP3ArtistIE_suitable,
			})
		}())
		PalcoMP3VideoIE = λ.Cal(λ.TypeType, λ.StrLiteral("PalcoMP3VideoIE"), λ.NewTuple(PalcoMP3BaseIE), func() λ.Dict {
			var (
				PalcoMP3VideoIE_IE_NAME       λ.Object
				PalcoMP3VideoIE__MUSIC_FIELDS λ.Object
				PalcoMP3VideoIE__VALID_URL    λ.Object
				PalcoMP3VideoIE__parse_music  λ.Object
			)
			PalcoMP3VideoIE_IE_NAME = λ.StrLiteral("PalcoMP3:video")
			PalcoMP3VideoIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?palcomp3\\.com(?:\\.br)?/(?P<artist>[^/]+)/(?P<id>[^/?&#]+)/?#clipe")
			PalcoMP3VideoIE__MUSIC_FIELDS = λ.StrLiteral("youtubeID")
			PalcoMP3VideoIE__parse_music = λ.NewFunction("_parse_music",
				[]λ.Param{
					{Name: "self"},
					{Name: "music"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒmusic      = λargs[1]
						ϒself       = λargs[0]
						ϒyoutube_id λ.Object
					)
					ϒyoutube_id = λ.GetItem(ϒmusic, λ.StrLiteral("youtubeID"))
					return λ.Calm(ϒself, "url_result", ϒyoutube_id, λ.StrLiteral("Youtube"), ϒyoutube_id)
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"IE_NAME":       PalcoMP3VideoIE_IE_NAME,
				"_MUSIC_FIELDS": PalcoMP3VideoIE__MUSIC_FIELDS,
				"_VALID_URL":    PalcoMP3VideoIE__VALID_URL,
				"_parse_music":  PalcoMP3VideoIE__parse_music,
			})
		}())
	})
}