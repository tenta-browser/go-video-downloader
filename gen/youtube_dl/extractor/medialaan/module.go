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
 * medialaan/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/medialaan.py
 */

package medialaan

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor       λ.Object
	MedialaanIE         λ.Object
	ϒextract_attributes λ.Object
	ϒint_or_none        λ.Object
	ϒmimetype2ext       λ.Object
	ϒparse_iso8601      λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒextract_attributes = Ωutils.ϒextract_attributes
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒmimetype2ext = Ωutils.ϒmimetype2ext
		ϒparse_iso8601 = Ωutils.ϒparse_iso8601
		MedialaanIE = λ.Cal(λ.TypeType, λ.StrLiteral("MedialaanIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				MedialaanIE__VALID_URL    λ.Object
				MedialaanIE__real_extract λ.Object
			)
			MedialaanIE__VALID_URL = λ.StrLiteral("(?x)\n                    https?://\n                        (?:\n                            (?:embed\\.)?mychannels.video/embed/|\n                            embed\\.mychannels\\.video/(?:s(?:dk|cript)/)?production/|\n                            (?:www\\.)?(?:\n                                (?:\n                                    7sur7|\n                                    demorgen|\n                                    hln|\n                                    joe|\n                                    qmusic\n                                )\\.be|\n                                (?:\n                                    [abe]d|\n                                    bndestem|\n                                    destentor|\n                                    gelderlander|\n                                    pzc|\n                                    tubantia|\n                                    volkskrant\n                                )\\.nl\n                            )/video/(?:[^/]+/)*[^/?&#]+~p\n                        )\n                        (?P<id>\\d+)\n                    ")
			MedialaanIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒext           λ.Object
						ϒformats       λ.Object
						ϒproduction    λ.Object
						ϒproduction_id λ.Object
						ϒself          = λargs[0]
						ϒsource        λ.Object
						ϒsrc           λ.Object
						ϒtitle         λ.Object
						ϒurl           = λargs[1]
						τmp0           λ.Object
						τmp1           λ.Object
					)
					ϒproduction_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒproduction = λ.GetItem(λ.GetItem(λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
						λ.Add(λ.StrLiteral("https://embed.mychannels.video/sdk/production/"), ϒproduction_id),
						ϒproduction_id,
					), λ.KWArgs{
						{Name: "query", Value: λ.DictLiteral(map[string]string{
							"options": "UUUU_default",
						})},
					}), λ.StrLiteral("productions")), λ.IntLiteral(0))
					ϒtitle = λ.GetItem(ϒproduction, λ.StrLiteral("title"))
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, func() λ.Object {
						if λv := λ.Calm(ϒproduction, "get", λ.StrLiteral("sources")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.NewList()
						}
					}())
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒsource = τmp1
						ϒsrc = λ.Calm(ϒsource, "get", λ.StrLiteral("src"))
						if !λ.IsTrue(ϒsrc) {
							continue
						}
						ϒext = λ.Cal(ϒmimetype2ext, λ.Calm(ϒsource, "get", λ.StrLiteral("type")))
						if λ.IsTrue(λ.Eq(ϒext, λ.StrLiteral("m3u8"))) {
							λ.Calm(ϒformats, "extend", λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
								ϒsrc,
								ϒproduction_id,
								λ.StrLiteral("mp4"),
								λ.StrLiteral("m3u8_native"),
							), λ.KWArgs{
								{Name: "m3u8_id", Value: λ.StrLiteral("hls")},
								{Name: "fatal", Value: λ.False},
							}))
						} else {
							λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
								"ext": ϒext,
								"url": ϒsrc,
							}))
						}
					}
					λ.Calm(ϒself, "_sort_formats", ϒformats)
					return λ.DictLiteral(map[string]λ.Object{
						"id":        ϒproduction_id,
						"title":     ϒtitle,
						"formats":   ϒformats,
						"thumbnail": λ.Calm(ϒproduction, "get", λ.StrLiteral("posterUrl")),
						"timestamp": λ.Cal(ϒparse_iso8601, λ.Calm(ϒproduction, "get", λ.StrLiteral("publicationDate")), λ.StrLiteral(" ")),
						"duration": func() λ.Object {
							if λv := λ.Cal(ϒint_or_none, λ.Calm(ϒproduction, "get", λ.StrLiteral("duration"))); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.None
							}
						}(),
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL":    MedialaanIE__VALID_URL,
				"_real_extract": MedialaanIE__real_extract,
			})
		}())
	})
}
