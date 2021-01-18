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
 * acast/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/acast.py
 */

package acast

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ACastBaseIE        λ.Object
	ACastChannelIE     λ.Object
	ACastIE            λ.Object
	InfoExtractor      λ.Object
	ϒclean_html        λ.Object
	ϒclean_podcast_url λ.Object
	ϒint_or_none       λ.Object
	ϒparse_iso8601     λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒclean_html = Ωutils.ϒclean_html
		ϒclean_podcast_url = Ωutils.ϒclean_podcast_url
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒparse_iso8601 = Ωutils.ϒparse_iso8601
		ACastBaseIE = λ.Cal(λ.TypeType, λ.StrLiteral("ACastBaseIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				ACastBaseIE__call_api          λ.Object
				ACastBaseIE__extract_episode   λ.Object
				ACastBaseIE__extract_show_info λ.Object
			)
			ACastBaseIE__extract_episode = λ.NewFunction("_extract_episode",
				[]λ.Param{
					{Name: "self"},
					{Name: "episode"},
					{Name: "show_info"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒepisode   = λargs[1]
						ϒinfo      λ.Object
						ϒself      = λargs[0]
						ϒshow_info = λargs[2]
						ϒtitle     λ.Object
					)
					_ = ϒself
					ϒtitle = λ.GetItem(ϒepisode, λ.StrLiteral("title"))
					ϒinfo = λ.DictLiteral(map[string]λ.Object{
						"id":         λ.GetItem(ϒepisode, λ.StrLiteral("id")),
						"display_id": λ.Calm(ϒepisode, "get", λ.StrLiteral("episodeUrl")),
						"url":        λ.Cal(ϒclean_podcast_url, λ.GetItem(ϒepisode, λ.StrLiteral("url"))),
						"title":      ϒtitle,
						"description": λ.Cal(ϒclean_html, func() λ.Object {
							if λv := λ.Calm(ϒepisode, "get", λ.StrLiteral("description")); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Calm(ϒepisode, "get", λ.StrLiteral("summary"))
							}
						}()),
						"thumbnail":      λ.Calm(ϒepisode, "get", λ.StrLiteral("image")),
						"timestamp":      λ.Cal(ϒparse_iso8601, λ.Calm(ϒepisode, "get", λ.StrLiteral("publishDate"))),
						"duration":       λ.Cal(ϒint_or_none, λ.Calm(ϒepisode, "get", λ.StrLiteral("duration"))),
						"filesize":       λ.Cal(ϒint_or_none, λ.Calm(ϒepisode, "get", λ.StrLiteral("contentLength"))),
						"season_number":  λ.Cal(ϒint_or_none, λ.Calm(ϒepisode, "get", λ.StrLiteral("season"))),
						"episode":        ϒtitle,
						"episode_number": λ.Cal(ϒint_or_none, λ.Calm(ϒepisode, "get", λ.StrLiteral("episode"))),
					})
					λ.Calm(ϒinfo, "update", ϒshow_info)
					return ϒinfo
				})
			ACastBaseIE__extract_show_info = λ.NewFunction("_extract_show_info",
				[]λ.Param{
					{Name: "self"},
					{Name: "show"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒself = λargs[0]
						ϒshow = λargs[1]
					)
					_ = ϒself
					return λ.DictLiteral(map[string]λ.Object{
						"creator": λ.Calm(ϒshow, "get", λ.StrLiteral("author")),
						"series":  λ.Calm(ϒshow, "get", λ.StrLiteral("title")),
					})
				})
			ACastBaseIE__call_api = λ.NewFunction("_call_api",
				[]λ.Param{
					{Name: "self"},
					{Name: "path"},
					{Name: "video_id"},
					{Name: "query", Def: λ.None},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒpath     = λargs[1]
						ϒquery    = λargs[3]
						ϒself     = λargs[0]
						ϒvideo_id = λargs[2]
					)
					return λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
						λ.Add(λ.StrLiteral("https://feeder.acast.com/api/v1/shows/"), ϒpath),
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "query", Value: ϒquery},
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_call_api":          ACastBaseIE__call_api,
				"_extract_episode":   ACastBaseIE__extract_episode,
				"_extract_show_info": ACastBaseIE__extract_show_info,
			})
		}())
		ACastIE = λ.Cal(λ.TypeType, λ.StrLiteral("ACastIE"), λ.NewTuple(ACastBaseIE), func() λ.Dict {
			var (
				ACastIE_IE_NAME       λ.Object
				ACastIE__VALID_URL    λ.Object
				ACastIE__real_extract λ.Object
			)
			ACastIE_IE_NAME = λ.StrLiteral("acast")
			ACastIE__VALID_URL = λ.StrLiteral("(?x)\n                    https?://\n                        (?:\n                            (?:(?:embed|www)\\.)?acast\\.com/|\n                            play\\.acast\\.com/s/\n                        )\n                        (?P<channel>[^/]+)/(?P<id>[^/#?]+)\n                    ")
			ACastIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒchannel    λ.Object
						ϒdisplay_id λ.Object
						ϒepisode    λ.Object
						ϒself       = λargs[0]
						ϒurl        = λargs[1]
						τmp0        λ.Object
					)
					τmp0 = λ.Calm(λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl), "groups")
					ϒchannel = λ.GetItem(τmp0, λ.IntLiteral(0))
					ϒdisplay_id = λ.GetItem(τmp0, λ.IntLiteral(1))
					ϒepisode = λ.Calm(ϒself, "_call_api", λ.Mod(λ.StrLiteral("%s/episodes/%s"), λ.NewTuple(
						ϒchannel,
						ϒdisplay_id,
					)), ϒdisplay_id, λ.DictLiteral(map[string]string{
						"showInfo": "true",
					}))
					return λ.Calm(ϒself, "_extract_episode", ϒepisode, λ.Calm(ϒself, "_extract_show_info", func() λ.Object {
						if λv := λ.Calm(ϒepisode, "get", λ.StrLiteral("show")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.DictLiteral(map[λ.Object]λ.Object{})
						}
					}()))
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"IE_NAME":       ACastIE_IE_NAME,
				"_VALID_URL":    ACastIE__VALID_URL,
				"_real_extract": ACastIE__real_extract,
			})
		}())
		ACastChannelIE = λ.Cal(λ.TypeType, λ.StrLiteral("ACastChannelIE"), λ.NewTuple(ACastBaseIE), func() λ.Dict {
			var (
				ACastChannelIE__VALID_URL λ.Object
				ACastChannelIE_suitable   λ.Object
			)
			ACastChannelIE__VALID_URL = λ.StrLiteral("(?x)\n                    https?://\n                        (?:\n                            (?:www\\.)?acast\\.com/|\n                            play\\.acast\\.com/s/\n                        )\n                        (?P<id>[^/#?]+)\n                    ")
			ACastChannelIE_suitable = λ.NewFunction("suitable",
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
						if λ.IsTrue(λ.Calm(ACastIE, "suitable", ϒurl)) {
							return λ.False
						} else {
							return λ.Calm(λ.Cal(λ.SuperType, ACastChannelIE, ϒcls), "suitable", ϒurl)
						}
					}()
				})
			ACastChannelIE_suitable = λ.Cal(λ.ClassMethodType, ACastChannelIE_suitable)
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": ACastChannelIE__VALID_URL,
				"suitable":   ACastChannelIE_suitable,
			})
		}())
	})
}
