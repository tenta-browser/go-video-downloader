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
 * tva/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/tva.py
 */

package tva

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor  λ.Object
	QubIE          λ.Object
	TVAIE          λ.Object
	ϒfloat_or_none λ.Object
	ϒint_or_none   λ.Object
	ϒsmuggle_url   λ.Object
	ϒstrip_or_none λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒfloat_or_none = Ωutils.ϒfloat_or_none
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒsmuggle_url = Ωutils.ϒsmuggle_url
		ϒstrip_or_none = Ωutils.ϒstrip_or_none
		TVAIE = λ.Cal(λ.TypeType, λ.StrLiteral("TVAIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				TVAIE_BRIGHTCOVE_URL_TEMPLATE λ.Object
				TVAIE__VALID_URL              λ.Object
				TVAIE__real_extract           λ.Object
			)
			TVAIE__VALID_URL = λ.StrLiteral("https?://videos?\\.tva\\.ca/details/_(?P<id>\\d+)")
			TVAIE_BRIGHTCOVE_URL_TEMPLATE = λ.StrLiteral("http://players.brightcove.net/5481942443001/default_default/index.html?videoId=%s")
			TVAIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒself     = λargs[0]
						ϒurl      = λargs[1]
						ϒvideo_id λ.Object
					)
					ϒvideo_id = λ.Calm(ϒself, "_match_id", ϒurl)
					return λ.DictLiteral(map[string]λ.Object{
						"_type": λ.StrLiteral("url_transparent"),
						"id":    ϒvideo_id,
						"url": λ.Cal(ϒsmuggle_url, λ.Mod(λ.GetAttr(ϒself, "BRIGHTCOVE_URL_TEMPLATE", nil), ϒvideo_id), λ.DictLiteral(map[string]λ.Object{
							"geo_countries": λ.NewList(λ.StrLiteral("CA")),
						})),
						"ie_key": λ.StrLiteral("BrightcoveNew"),
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"BRIGHTCOVE_URL_TEMPLATE": TVAIE_BRIGHTCOVE_URL_TEMPLATE,
				"_VALID_URL":              TVAIE__VALID_URL,
				"_real_extract":           TVAIE__real_extract,
			})
		}())
		QubIE = λ.Cal(λ.TypeType, λ.StrLiteral("QubIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				QubIE__VALID_URL    λ.Object
				QubIE__real_extract λ.Object
			)
			QubIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?qub\\.ca/(?:[^/]+/)*[0-9a-z-]+-(?P<id>\\d+)")
			QubIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒentity    λ.Object
						ϒentity_id λ.Object
						ϒepisode   λ.Object
						ϒself      = λargs[0]
						ϒurl       = λargs[1]
						ϒvideo_id  λ.Object
					)
					ϒentity_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒentity = λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
						λ.StrLiteral("https://www.qub.ca/proxy/pfu/content-delivery-service/v1/entities"),
						ϒentity_id,
					), λ.KWArgs{
						{Name: "query", Value: λ.DictLiteral(map[string]λ.Object{
							"id": ϒentity_id,
						})},
					})
					ϒvideo_id = λ.GetItem(ϒentity, λ.StrLiteral("videoId"))
					ϒepisode = λ.Cal(ϒstrip_or_none, λ.Calm(ϒentity, "get", λ.StrLiteral("name")))
					return λ.DictLiteral(map[string]λ.Object{
						"_type":          λ.StrLiteral("url_transparent"),
						"id":             ϒvideo_id,
						"title":          ϒepisode,
						"url":            λ.Add(λ.StrLiteral("https://videos.tva.ca/details/_"), ϒvideo_id),
						"description":    λ.Calm(ϒentity, "get", λ.StrLiteral("longDescription")),
						"duration":       λ.Cal(ϒfloat_or_none, λ.Calm(ϒentity, "get", λ.StrLiteral("durationMillis")), λ.IntLiteral(1000)),
						"episode":        ϒepisode,
						"episode_number": λ.Cal(ϒint_or_none, λ.Calm(ϒentity, "get", λ.StrLiteral("episodeNumber"))),
						"ie_key":         λ.Calm(TVAIE, "ie_key"),
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL":    QubIE__VALID_URL,
				"_real_extract": QubIE__real_extract,
			})
		}())
	})
}
