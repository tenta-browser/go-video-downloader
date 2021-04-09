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
 * tfo/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/tfo.py
 */

package tfo

import (
	Ωjson "github.com/tenta-browser/go-video-downloader/gen/json"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError λ.Object
	HEADRequest    λ.Object
	InfoExtractor  λ.Object
	TFOIE          λ.Object
	ϒclean_html    λ.Object
	ϒint_or_none   λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		HEADRequest = Ωutils.HEADRequest
		ExtractorError = Ωutils.ExtractorError
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒclean_html = Ωutils.ϒclean_html
		TFOIE = λ.Cal(λ.TypeType, λ.StrLiteral("TFOIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				TFOIE__GEO_COUNTRIES λ.Object
				TFOIE__VALID_URL     λ.Object
				TFOIE__real_extract  λ.Object
			)
			TFOIE__GEO_COUNTRIES = λ.NewList(λ.StrLiteral("CA"))
			TFOIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?tfo\\.org/(?:en|fr)/(?:[^/]+/){2}(?P<id>\\d+)")
			TFOIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒinfos      λ.Object
						ϒself       = λargs[0]
						ϒurl        = λargs[1]
						ϒvideo_data λ.Object
						ϒvideo_id   λ.Object
					)
					ϒvideo_id = λ.Calm(ϒself, "_match_id", ϒurl)
					λ.Calm(ϒself, "_request_webpage", λ.Cal(HEADRequest, λ.StrLiteral("http://www.tfo.org/")), ϒvideo_id)
					ϒinfos = λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
						λ.StrLiteral("http://www.tfo.org/api/web/video/get_infos"),
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "data", Value: λ.Calm(λ.Cal(Ωjson.ϒdumps, λ.DictLiteral(map[string]λ.Object{
							"product_id": ϒvideo_id,
						})), "encode")},
						{Name: "headers", Value: λ.DictLiteral(map[string]λ.Object{
							"X-tfo-session": λ.GetAttr(λ.GetItem(λ.Calm(ϒself, "_get_cookies", λ.StrLiteral("http://www.tfo.org/")), λ.StrLiteral("tfo-session")), "value", nil),
						})},
					})
					if λ.IsTrue(λ.Eq(λ.Calm(ϒinfos, "get", λ.StrLiteral("success")), λ.IntLiteral(0))) {
						if λ.IsTrue(λ.Eq(λ.Calm(ϒinfos, "get", λ.StrLiteral("code")), λ.StrLiteral("ErrGeoBlocked"))) {
							λ.Call(λ.GetAttr(ϒself, "raise_geo_restricted", nil), nil, λ.KWArgs{
								{Name: "countries", Value: λ.GetAttr(ϒself, "_GEO_COUNTRIES", nil)},
							})
						}
						panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.Mod(λ.StrLiteral("%s said: %s"), λ.NewTuple(
							λ.GetAttr(ϒself, "IE_NAME", nil),
							λ.Cal(ϒclean_html, λ.GetItem(ϒinfos, λ.StrLiteral("msg"))),
						))), λ.KWArgs{
							{Name: "expected", Value: λ.True},
						})))
					}
					ϒvideo_data = λ.GetItem(ϒinfos, λ.StrLiteral("data"))
					return λ.DictLiteral(map[string]λ.Object{
						"_type":          λ.StrLiteral("url_transparent"),
						"id":             ϒvideo_id,
						"url":            λ.Add(λ.StrLiteral("limelight:media:"), λ.GetItem(ϒvideo_data, λ.StrLiteral("llid"))),
						"title":          λ.GetItem(ϒvideo_data, λ.StrLiteral("title")),
						"description":    λ.Calm(ϒvideo_data, "get", λ.StrLiteral("description")),
						"series":         λ.Calm(ϒvideo_data, "get", λ.StrLiteral("collection")),
						"season_number":  λ.Cal(ϒint_or_none, λ.Calm(ϒvideo_data, "get", λ.StrLiteral("season"))),
						"episode_number": λ.Cal(ϒint_or_none, λ.Calm(ϒvideo_data, "get", λ.StrLiteral("episode"))),
						"duration":       λ.Cal(ϒint_or_none, λ.Calm(ϒvideo_data, "get", λ.StrLiteral("duration"))),
						"ie_key":         λ.StrLiteral("LimelightMedia"),
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_GEO_COUNTRIES": TFOIE__GEO_COUNTRIES,
				"_VALID_URL":     TFOIE__VALID_URL,
				"_real_extract":  TFOIE__real_extract,
			})
		}())
	})
}
