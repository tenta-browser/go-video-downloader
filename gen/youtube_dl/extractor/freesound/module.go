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
 * freesound/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/freesound.py
 */

package freesound

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	FreesoundIE           λ.Object
	InfoExtractor         λ.Object
	ϒfloat_or_none        λ.Object
	ϒget_element_by_class λ.Object
	ϒget_element_by_id    λ.Object
	ϒunified_strdate      λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒfloat_or_none = Ωutils.ϒfloat_or_none
		ϒget_element_by_class = Ωutils.ϒget_element_by_class
		ϒget_element_by_id = Ωutils.ϒget_element_by_id
		ϒunified_strdate = Ωutils.ϒunified_strdate
		FreesoundIE = λ.Cal(λ.TypeType, λ.StrLiteral("FreesoundIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				FreesoundIE__VALID_URL    λ.Object
				FreesoundIE__real_extract λ.Object
			)
			FreesoundIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?freesound\\.org/people/[^/]+/sounds/(?P<id>[^/]+)")
			FreesoundIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						LQ_FORMAT    λ.Object
						ϒaudio_id    λ.Object
						ϒaudio_url   λ.Object
						ϒaudio_urls  λ.Object
						ϒchannels    λ.Object
						ϒdescription λ.Object
						ϒduration    λ.Object
						ϒformats     λ.Object
						ϒself        = λargs[0]
						ϒtags        λ.Object
						ϒtags_str    λ.Object
						ϒtitle       λ.Object
						ϒupload_date λ.Object
						ϒuploader    λ.Object
						ϒurl         = λargs[1]
						ϒwebpage     λ.Object
					)
					ϒaudio_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒwebpage = λ.Calm(ϒself, "_download_webpage", ϒurl, ϒaudio_id)
					ϒaudio_url = λ.Calm(ϒself, "_og_search_property", λ.StrLiteral("audio"), ϒwebpage, λ.StrLiteral("song url"))
					ϒtitle = λ.Calm(ϒself, "_og_search_property", λ.StrLiteral("audio:title"), ϒwebpage, λ.StrLiteral("song title"))
					ϒdescription = λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("(?s)id=[\"\\']sound_description[\"\\'][^>]*>(.+?)</div>"),
						ϒwebpage,
						λ.StrLiteral("description"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					})
					ϒduration = λ.Call(ϒfloat_or_none, λ.NewArgs(λ.Cal(ϒget_element_by_class, λ.StrLiteral("duration"), ϒwebpage)), λ.KWArgs{
						{Name: "scale", Value: λ.IntLiteral(1000)},
					})
					ϒupload_date = λ.Cal(ϒunified_strdate, λ.Cal(ϒget_element_by_id, λ.StrLiteral("sound_date"), ϒwebpage))
					ϒuploader = λ.Call(λ.GetAttr(ϒself, "_og_search_property", nil), λ.NewArgs(
						λ.StrLiteral("audio:artist"),
						ϒwebpage,
						λ.StrLiteral("uploader"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					})
					ϒchannels = λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("Channels</dt><dd>(.+?)</dd>"),
						ϒwebpage,
						λ.StrLiteral("channels info"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					})
					ϒtags_str = λ.Cal(ϒget_element_by_class, λ.StrLiteral("tags"), ϒwebpage)
					ϒtags = func() λ.Object {
						if λ.IsTrue(ϒtags_str) {
							return λ.Cal(Ωre.ϒfindall, λ.StrLiteral("<a[^>]+>([^<]+)"), ϒtags_str)
						} else {
							return λ.None
						}
					}()
					ϒaudio_urls = λ.NewList(ϒaudio_url)
					LQ_FORMAT = λ.StrLiteral("-lq.mp3")
					if λ.Contains(ϒaudio_url, LQ_FORMAT) {
						λ.Calm(ϒaudio_urls, "append", λ.Calm(ϒaudio_url, "replace", LQ_FORMAT, λ.StrLiteral("-hq.mp3")))
					}
					ϒformats = λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
						nil,
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							return λ.NewGenerator(func(λgy λ.Yielder) λ.Object {
								var (
									ϒformat_url λ.Object
									ϒquality    λ.Object
									τmp0        λ.Object
									τmp1        λ.Object
									τmp2        λ.Object
								)
								τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.EnumerateIteratorType, ϒaudio_urls))
								for {
									if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
										break
									}
									τmp2 = λ.UnpackIterable(τmp1, 2)
									ϒquality = λ.GetItem(τmp2, λ.IntLiteral(0))
									ϒformat_url = λ.GetItem(τmp2, λ.IntLiteral(1))
									λgy.Yield(λ.DictLiteral(map[string]λ.Object{
										"url":         ϒformat_url,
										"format_note": ϒchannels,
										"quality":     ϒquality,
									}))
								}
								return λ.None
							})
						})))
					λ.Calm(ϒself, "_sort_formats", ϒformats)
					return λ.DictLiteral(map[string]λ.Object{
						"id":          ϒaudio_id,
						"title":       ϒtitle,
						"description": ϒdescription,
						"duration":    ϒduration,
						"uploader":    ϒuploader,
						"upload_date": ϒupload_date,
						"tags":        ϒtags,
						"formats":     ϒformats,
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL":    FreesoundIE__VALID_URL,
				"_real_extract": FreesoundIE__real_extract,
			})
		}())
	})
}
