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
 * pokemon/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/pokemon.py
 */

package pokemon

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor       λ.Object
	PokemonIE           λ.Object
	ϒextract_attributes λ.Object
	ϒint_or_none        λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒextract_attributes = Ωutils.ϒextract_attributes
		ϒint_or_none = Ωutils.ϒint_or_none
		PokemonIE = λ.Cal(λ.TypeType, λ.StrLiteral("PokemonIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				PokemonIE__VALID_URL    λ.Object
				PokemonIE__real_extract λ.Object
			)
			PokemonIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?pokemon\\.com/[a-z]{2}(?:.*?play=(?P<id>[a-z0-9]{32})|/(?:[^/]+/)+(?P<display_id>[^/?#&]+))")
			PokemonIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdisplay_id λ.Object
						ϒself       = λargs[0]
						ϒtitle      λ.Object
						ϒurl        = λargs[1]
						ϒvideo_data λ.Object
						ϒvideo_id   λ.Object
						ϒwebpage    λ.Object
						τmp0        λ.Object
					)
					τmp0 = λ.UnpackIterable(λ.Calm(λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl), "groups"), 2)
					ϒvideo_id = λ.GetItem(τmp0, λ.IntLiteral(0))
					ϒdisplay_id = λ.GetItem(τmp0, λ.IntLiteral(1))
					ϒwebpage = λ.Calm(ϒself, "_download_webpage", ϒurl, func() λ.Object {
						if λv := ϒvideo_id; λ.IsTrue(λv) {
							return λv
						} else {
							return ϒdisplay_id
						}
					}())
					ϒvideo_data = λ.Cal(ϒextract_attributes, λ.Calm(ϒself, "_search_regex", λ.Mod(λ.StrLiteral("(<[^>]+data-video-id=\"%s\"[^>]*>)"), func() λ.Object {
						if λ.IsTrue(ϒvideo_id) {
							return ϒvideo_id
						} else {
							return λ.StrLiteral("[a-z0-9]{32}")
						}
					}()), ϒwebpage, λ.StrLiteral("video data element")))
					ϒvideo_id = λ.GetItem(ϒvideo_data, λ.StrLiteral("data-video-id"))
					ϒtitle = func() λ.Object {
						if λv := λ.Calm(ϒvideo_data, "get", λ.StrLiteral("data-video-title")); λ.IsTrue(λv) {
							return λv
						} else if λv := λ.Call(λ.GetAttr(ϒself, "_html_search_meta", nil), λ.NewArgs(
							λ.StrLiteral("pkm-title"),
							ϒwebpage,
							λ.StrLiteral(" title"),
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
						}); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Calm(ϒself, "_search_regex", λ.StrLiteral("<h1[^>]+\\bclass=[\"\\']us-title[^>]+>([^<]+)"), ϒwebpage, λ.StrLiteral("title"))
						}
					}()
					return λ.DictLiteral(map[string]λ.Object{
						"_type":          λ.StrLiteral("url_transparent"),
						"id":             ϒvideo_id,
						"url":            λ.Mod(λ.StrLiteral("limelight:media:%s"), ϒvideo_id),
						"title":          ϒtitle,
						"description":    λ.Calm(ϒvideo_data, "get", λ.StrLiteral("data-video-summary")),
						"thumbnail":      λ.Calm(ϒvideo_data, "get", λ.StrLiteral("data-video-poster")),
						"series":         λ.StrLiteral("Pokémon"),
						"season_number":  λ.Cal(ϒint_or_none, λ.Calm(ϒvideo_data, "get", λ.StrLiteral("data-video-season"))),
						"episode":        ϒtitle,
						"episode_number": λ.Cal(ϒint_or_none, λ.Calm(ϒvideo_data, "get", λ.StrLiteral("data-video-episode"))),
						"ie_key":         λ.StrLiteral("LimelightMedia"),
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL":    PokemonIE__VALID_URL,
				"_real_extract": PokemonIE__real_extract,
			})
		}())
	})
}
