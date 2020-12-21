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
 * tv5unis/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/tv5unis.py
 */

package tv5unis

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor    λ.Object
	TV5UnisBaseIE    λ.Object
	TV5UnisIE        λ.Object
	TV5UnisVideoIE   λ.Object
	ϒint_or_none     λ.Object
	ϒparse_age_limit λ.Object
	ϒsmuggle_url     λ.Object
	ϒtry_get         λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒparse_age_limit = Ωutils.ϒparse_age_limit
		ϒsmuggle_url = Ωutils.ϒsmuggle_url
		ϒtry_get = Ωutils.ϒtry_get
		TV5UnisBaseIE = λ.Cal(λ.TypeType, λ.StrLiteral("TV5UnisBaseIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				TV5UnisBaseIE__GEO_COUNTRIES λ.Object
				TV5UnisBaseIE__real_extract  λ.Object
			)
			TV5UnisBaseIE__GEO_COUNTRIES = λ.NewList(λ.StrLiteral("CA"))
			TV5UnisBaseIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒgroups   λ.Object
						ϒmedia_id λ.Object
						ϒproduct  λ.Object
						ϒself     = λargs[0]
						ϒurl      = λargs[1]
					)
					ϒgroups = λ.Calm(λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl), "groups")
					ϒproduct = λ.GetItem(λ.GetItem(λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
						λ.StrLiteral("https://api.tv5unis.ca/graphql"),
						λ.GetItem(ϒgroups, λ.IntLiteral(0)),
					), λ.KWArgs{
						{Name: "query", Value: λ.DictLiteral(map[string]λ.Object{
							"query": λ.Mod(λ.StrLiteral("{\n  %s(%s) {\n    collection {\n      title\n    }\n    episodeNumber\n    rating {\n      name\n    }\n    seasonNumber\n    tags\n    title\n    videoElement {\n      ... on Video {\n        mediaId\n      }\n    }\n  }\n}"), λ.NewTuple(
								λ.GetAttr(ϒself, "_GQL_QUERY_NAME", nil),
								λ.Calm(ϒself, "_gql_args", ϒgroups),
							)),
						})},
					}), λ.StrLiteral("data")), λ.GetAttr(ϒself, "_GQL_QUERY_NAME", nil))
					ϒmedia_id = λ.GetItem(λ.GetItem(ϒproduct, λ.StrLiteral("videoElement")), λ.StrLiteral("mediaId"))
					return λ.DictLiteral(map[string]λ.Object{
						"_type": λ.StrLiteral("url_transparent"),
						"id":    ϒmedia_id,
						"title": λ.Calm(ϒproduct, "get", λ.StrLiteral("title")),
						"url": λ.Cal(ϒsmuggle_url, λ.Add(λ.StrLiteral("limelight:media:"), ϒmedia_id), λ.DictLiteral(map[string]λ.Object{
							"geo_countries": λ.GetAttr(ϒself, "_GEO_COUNTRIES", nil),
						})),
						"age_limit": λ.Cal(ϒparse_age_limit, λ.Cal(ϒtry_get, ϒproduct, λ.NewFunction("<lambda>",
							[]λ.Param{
								{Name: "x"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒx = λargs[0]
								)
								return λ.GetItem(λ.GetItem(ϒx, λ.StrLiteral("rating")), λ.StrLiteral("name"))
							}))),
						"tags": λ.Calm(ϒproduct, "get", λ.StrLiteral("tags")),
						"series": λ.Cal(ϒtry_get, ϒproduct, λ.NewFunction("<lambda>",
							[]λ.Param{
								{Name: "x"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒx = λargs[0]
								)
								return λ.GetItem(λ.GetItem(ϒx, λ.StrLiteral("collection")), λ.StrLiteral("title"))
							})),
						"season_number":  λ.Cal(ϒint_or_none, λ.Calm(ϒproduct, "get", λ.StrLiteral("seasonNumber"))),
						"episode_number": λ.Cal(ϒint_or_none, λ.Calm(ϒproduct, "get", λ.StrLiteral("episodeNumber"))),
						"ie_key":         λ.StrLiteral("LimelightMedia"),
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_GEO_COUNTRIES": TV5UnisBaseIE__GEO_COUNTRIES,
				"_real_extract":  TV5UnisBaseIE__real_extract,
			})
		}())
		TV5UnisVideoIE = λ.Cal(λ.TypeType, λ.StrLiteral("TV5UnisVideoIE"), λ.NewTuple(TV5UnisBaseIE), func() λ.Dict {
			var (
				TV5UnisVideoIE_IE_NAME         λ.Object
				TV5UnisVideoIE__GQL_QUERY_NAME λ.Object
				TV5UnisVideoIE__VALID_URL      λ.Object
				TV5UnisVideoIE__gql_args       λ.Object
			)
			TV5UnisVideoIE_IE_NAME = λ.StrLiteral("tv5unis:video")
			TV5UnisVideoIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?tv5unis\\.ca/videos/[^/]+/(?P<id>\\d+)")
			TV5UnisVideoIE__GQL_QUERY_NAME = λ.StrLiteral("productById")
			TV5UnisVideoIE__gql_args = λ.NewFunction("_gql_args",
				[]λ.Param{
					{Name: "groups"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒgroups = λargs[0]
					)
					return λ.Mod(λ.StrLiteral("id: %s"), ϒgroups)
				})
			TV5UnisVideoIE__gql_args = λ.Cal(λ.StaticMethodType, TV5UnisVideoIE__gql_args)
			return λ.ClassDictLiteral(map[string]λ.Object{
				"IE_NAME":         TV5UnisVideoIE_IE_NAME,
				"_GQL_QUERY_NAME": TV5UnisVideoIE__GQL_QUERY_NAME,
				"_VALID_URL":      TV5UnisVideoIE__VALID_URL,
				"_gql_args":       TV5UnisVideoIE__gql_args,
			})
		}())
		TV5UnisIE = λ.Cal(λ.TypeType, λ.StrLiteral("TV5UnisIE"), λ.NewTuple(TV5UnisBaseIE), func() λ.Dict {
			var (
				TV5UnisIE_IE_NAME         λ.Object
				TV5UnisIE__GQL_QUERY_NAME λ.Object
				TV5UnisIE__VALID_URL      λ.Object
				TV5UnisIE__gql_args       λ.Object
			)
			TV5UnisIE_IE_NAME = λ.StrLiteral("tv5unis")
			TV5UnisIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?tv5unis\\.ca/videos/(?P<id>[^/]+)(?:/saisons/(?P<season_number>\\d+)/episodes/(?P<episode_number>\\d+))?/?(?:[?#&]|$)")
			TV5UnisIE__GQL_QUERY_NAME = λ.StrLiteral("productByRootProductSlug")
			TV5UnisIE__gql_args = λ.NewFunction("_gql_args",
				[]λ.Param{
					{Name: "groups"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒargs   λ.Object
						ϒgroups = λargs[0]
						τmp0    λ.Object
					)
					ϒargs = λ.Mod(λ.StrLiteral("rootProductSlug: \"%s\""), λ.GetItem(ϒgroups, λ.IntLiteral(0)))
					if λ.IsTrue(λ.GetItem(ϒgroups, λ.IntLiteral(1))) {
						τmp0 = λ.IAdd(ϒargs, λ.Mod(λ.StrLiteral(", seasonNumber: %s, episodeNumber: %s"), λ.GetItem(ϒgroups, λ.NewSlice(λ.IntLiteral(1), λ.None, λ.None))))
						ϒargs = τmp0
					}
					return ϒargs
				})
			TV5UnisIE__gql_args = λ.Cal(λ.StaticMethodType, TV5UnisIE__gql_args)
			return λ.ClassDictLiteral(map[string]λ.Object{
				"IE_NAME":         TV5UnisIE_IE_NAME,
				"_GQL_QUERY_NAME": TV5UnisIE__GQL_QUERY_NAME,
				"_VALID_URL":      TV5UnisIE__VALID_URL,
				"_gql_args":       TV5UnisIE__gql_args,
			})
		}())
	})
}
