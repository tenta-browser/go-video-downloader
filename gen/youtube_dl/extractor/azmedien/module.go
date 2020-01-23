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
 * azmedien/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/azmedien.py
 */

package azmedien

import (
	Ωjson "github.com/tenta-browser/go-video-downloader/gen/json"
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωkaltura "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/kaltura"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	AZMedienIE    λ.Object
	InfoExtractor λ.Object
	KalturaIE     λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		KalturaIE = Ωkaltura.KalturaIE
		AZMedienIE = λ.Cal(λ.TypeType, λ.StrLiteral("AZMedienIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				AZMedienIE__API_TEMPL    λ.Object
				AZMedienIE__PARTNER_ID   λ.Object
				AZMedienIE__VALID_URL    λ.Object
				AZMedienIE__real_extract λ.Object
			)
			AZMedienIE__VALID_URL = λ.StrLiteral("(?x)\n                    https?://\n                        (?:www\\.)?\n                        (?P<host>\n                            telezueri\\.ch|\n                            telebaern\\.tv|\n                            telem1\\.ch\n                        )/\n                        [^/]+/\n                        (?P<id>\n                            [^/]+-(?P<article_id>\\d+)\n                        )\n                        (?:\n                            \\#video=\n                            (?P<kaltura_id>\n                                [_0-9a-z]+\n                            )\n                        )?\n                    ")
			AZMedienIE__API_TEMPL = λ.StrLiteral("https://www.%s/api/pub/gql/%s/NewsArticleTeaser/cb9f2f81ed22e9b47f4ca64ea3cc5a5d13e88d1d")
			AZMedienIE__PARTNER_ID = λ.StrLiteral("1719221")
			AZMedienIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒarticle_id λ.Object
						ϒdisplay_id λ.Object
						ϒentry_id   λ.Object
						ϒhost       λ.Object
						ϒself       = λargs[0]
						ϒurl        = λargs[1]
						τmp0        λ.Object
					)
					τmp0 = λ.Calm(λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl), "groups")
					ϒhost = λ.GetItem(τmp0, λ.IntLiteral(0))
					ϒdisplay_id = λ.GetItem(τmp0, λ.IntLiteral(1))
					ϒarticle_id = λ.GetItem(τmp0, λ.IntLiteral(2))
					ϒentry_id = λ.GetItem(τmp0, λ.IntLiteral(3))
					if !λ.IsTrue(ϒentry_id) {
						ϒentry_id = λ.GetItem(λ.GetItem(λ.GetItem(λ.GetItem(λ.GetItem(λ.GetItem(λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
							λ.Mod(λ.GetAttr(ϒself, "_API_TEMPL", nil), λ.NewTuple(
								ϒhost,
								λ.GetItem(λ.Calm(ϒhost, "split", λ.StrLiteral(".")), λ.IntLiteral(0)),
							)),
							ϒdisplay_id,
						), λ.KWArgs{
							{Name: "query", Value: λ.DictLiteral(map[string]λ.Object{
								"variables": λ.Cal(Ωjson.ϒdumps, λ.DictLiteral(map[string]λ.Object{
									"contextId": λ.Add(λ.StrLiteral("NewsArticle:"), ϒarticle_id),
								})),
							})},
						}), λ.StrLiteral("data")), λ.StrLiteral("context")), λ.StrLiteral("mainAsset")), λ.StrLiteral("video")), λ.StrLiteral("kaltura")), λ.StrLiteral("kalturaId"))
					}
					return λ.Call(λ.GetAttr(ϒself, "url_result", nil), λ.NewArgs(λ.Mod(λ.StrLiteral("kaltura:%s:%s"), λ.NewTuple(
						λ.GetAttr(ϒself, "_PARTNER_ID", nil),
						ϒentry_id,
					))), λ.KWArgs{
						{Name: "ie", Value: λ.Calm(KalturaIE, "ie_key")},
						{Name: "video_id", Value: ϒentry_id},
					})
				})
			return λ.DictLiteral(map[string]λ.Object{
				"_API_TEMPL":    AZMedienIE__API_TEMPL,
				"_PARTNER_ID":   AZMedienIE__PARTNER_ID,
				"_VALID_URL":    AZMedienIE__VALID_URL,
				"_real_extract": AZMedienIE__real_extract,
			})
		}())
	})
}
