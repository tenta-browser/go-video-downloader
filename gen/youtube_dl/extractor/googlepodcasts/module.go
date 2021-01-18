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
 * googlepodcasts/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/googlepodcasts.py
 */

package googlepodcasts

import (
	Ωjson "github.com/tenta-browser/go-video-downloader/gen/json"
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	GooglePodcastsBaseIE λ.Object
	GooglePodcastsFeedIE λ.Object
	GooglePodcastsIE     λ.Object
	InfoExtractor        λ.Object
	ϒclean_podcast_url   λ.Object
	ϒint_or_none         λ.Object
	ϒtry_get             λ.Object
	ϒurlencode_postdata  λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒclean_podcast_url = Ωutils.ϒclean_podcast_url
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒtry_get = Ωutils.ϒtry_get
		ϒurlencode_postdata = Ωutils.ϒurlencode_postdata
		GooglePodcastsBaseIE = λ.Cal(λ.TypeType, λ.StrLiteral("GooglePodcastsBaseIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				GooglePodcastsBaseIE__VALID_URL_BASE  λ.Object
				GooglePodcastsBaseIE__batch_execute   λ.Object
				GooglePodcastsBaseIE__extract_episode λ.Object
			)
			GooglePodcastsBaseIE__VALID_URL_BASE = λ.StrLiteral("https?://podcasts\\.google\\.com/feed/")
			GooglePodcastsBaseIE__batch_execute = λ.NewFunction("_batch_execute",
				[]λ.Param{
					{Name: "self"},
					{Name: "func_id"},
					{Name: "video_id"},
					{Name: "params"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒfunc_id  = λargs[1]
						ϒparams   = λargs[3]
						ϒself     = λargs[0]
						ϒvideo_id = λargs[2]
					)
					return λ.Cal(Ωjson.ϒloads, λ.GetItem(λ.GetItem(λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
						λ.StrLiteral("https://podcasts.google.com/_/PodcastsUi/data/batchexecute"),
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "data", Value: λ.Cal(ϒurlencode_postdata, λ.DictLiteral(map[string]λ.Object{
							"f.req": λ.Cal(Ωjson.ϒdumps, λ.NewList(λ.NewList(λ.NewList(
								ϒfunc_id,
								λ.Cal(Ωjson.ϒdumps, ϒparams),
								λ.None,
								λ.StrLiteral("1"),
							)))),
						}))},
						{Name: "transform_source", Value: λ.NewFunction("<lambda>",
							[]λ.Param{
								{Name: "x"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒx = λargs[0]
								)
								return λ.Calm(ϒself, "_search_regex", λ.StrLiteral("(?s)(\\[.+\\])"), ϒx, λ.StrLiteral("data"))
							})},
					}), λ.IntLiteral(0)), λ.IntLiteral(2)))
				})
			GooglePodcastsBaseIE__extract_episode = λ.NewFunction("_extract_episode",
				[]λ.Param{
					{Name: "self"},
					{Name: "episode"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒepisode = λargs[1]
						ϒself    = λargs[0]
					)
					_ = ϒself
					return λ.DictLiteral(map[string]λ.Object{
						"id":          λ.GetItem(λ.GetItem(ϒepisode, λ.IntLiteral(4)), λ.IntLiteral(3)),
						"title":       λ.GetItem(ϒepisode, λ.IntLiteral(8)),
						"url":         λ.Cal(ϒclean_podcast_url, λ.GetItem(ϒepisode, λ.IntLiteral(13))),
						"thumbnail":   λ.GetItem(ϒepisode, λ.IntLiteral(2)),
						"description": λ.GetItem(ϒepisode, λ.IntLiteral(9)),
						"creator": λ.Cal(ϒtry_get, ϒepisode, λ.NewFunction("<lambda>",
							[]λ.Param{
								{Name: "x"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒx = λargs[0]
								)
								return λ.GetItem(ϒx, λ.IntLiteral(14))
							})),
						"timestamp": λ.Cal(ϒint_or_none, λ.GetItem(ϒepisode, λ.IntLiteral(11))),
						"duration":  λ.Cal(ϒint_or_none, λ.GetItem(ϒepisode, λ.IntLiteral(12))),
						"series":    λ.GetItem(ϒepisode, λ.IntLiteral(1)),
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL_BASE":  GooglePodcastsBaseIE__VALID_URL_BASE,
				"_batch_execute":   GooglePodcastsBaseIE__batch_execute,
				"_extract_episode": GooglePodcastsBaseIE__extract_episode,
			})
		}())
		GooglePodcastsIE = λ.Cal(λ.TypeType, λ.StrLiteral("GooglePodcastsIE"), λ.NewTuple(GooglePodcastsBaseIE), func() λ.Dict {
			var (
				GooglePodcastsIE_IE_NAME       λ.Object
				GooglePodcastsIE__VALID_URL    λ.Object
				GooglePodcastsIE__real_extract λ.Object
			)
			GooglePodcastsIE_IE_NAME = λ.StrLiteral("google:podcasts")
			GooglePodcastsIE__VALID_URL = λ.Add(λ.GetAttr(GooglePodcastsBaseIE, "_VALID_URL_BASE", nil), λ.StrLiteral("(?P<feed_url>[^/]+)/episode/(?P<id>[^/?&#]+)"))
			GooglePodcastsIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒb64_feed_url λ.Object
						ϒb64_guid     λ.Object
						ϒepisode      λ.Object
						ϒself         = λargs[0]
						ϒurl          = λargs[1]
						τmp0          λ.Object
					)
					τmp0 = λ.Calm(λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl), "groups")
					ϒb64_feed_url = λ.GetItem(τmp0, λ.IntLiteral(0))
					ϒb64_guid = λ.GetItem(τmp0, λ.IntLiteral(1))
					ϒepisode = λ.GetItem(λ.Calm(ϒself, "_batch_execute", λ.StrLiteral("oNjqVe"), ϒb64_guid, λ.NewList(
						ϒb64_feed_url,
						ϒb64_guid,
					)), λ.IntLiteral(1))
					return λ.Calm(ϒself, "_extract_episode", ϒepisode)
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"IE_NAME":       GooglePodcastsIE_IE_NAME,
				"_VALID_URL":    GooglePodcastsIE__VALID_URL,
				"_real_extract": GooglePodcastsIE__real_extract,
			})
		}())
		GooglePodcastsFeedIE = λ.Cal(λ.TypeType, λ.StrLiteral("GooglePodcastsFeedIE"), λ.NewTuple(GooglePodcastsBaseIE), func() λ.Dict {
			var (
				GooglePodcastsFeedIE__VALID_URL λ.Object
			)
			GooglePodcastsFeedIE__VALID_URL = λ.Add(λ.GetAttr(GooglePodcastsBaseIE, "_VALID_URL_BASE", nil), λ.StrLiteral("(?P<id>[^/?&#]+)/?(?:[?#&]|$)"))
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": GooglePodcastsFeedIE__VALID_URL,
			})
		}())
	})
}