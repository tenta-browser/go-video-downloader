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
 * phoenix/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/phoenix.py
 */

package phoenix

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωyoutube "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/youtube"
	Ωzdf "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/zdf"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	PhoenixIE          λ.Object
	YoutubeIE          λ.Object
	ZDFBaseIE          λ.Object
	ϒcompat_str        λ.Object
	ϒint_or_none       λ.Object
	ϒmerge_dicts       λ.Object
	ϒtry_get           λ.Object
	ϒunified_timestamp λ.Object
	ϒurljoin           λ.Object
)

func init() {
	λ.InitModule(func() {
		YoutubeIE = Ωyoutube.YoutubeIE
		ZDFBaseIE = Ωzdf.ZDFBaseIE
		ϒcompat_str = Ωcompat.ϒcompat_str
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒmerge_dicts = Ωutils.ϒmerge_dicts
		ϒtry_get = Ωutils.ϒtry_get
		ϒunified_timestamp = Ωutils.ϒunified_timestamp
		ϒurljoin = Ωutils.ϒurljoin
		PhoenixIE = λ.Cal(λ.TypeType, λ.StrLiteral("PhoenixIE"), λ.NewTuple(ZDFBaseIE), func() λ.Dict {
			var (
				PhoenixIE_IE_NAME       λ.Object
				PhoenixIE__VALID_URL    λ.Object
				PhoenixIE__real_extract λ.Object
			)
			PhoenixIE_IE_NAME = λ.StrLiteral("phoenix.de")
			PhoenixIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?phoenix\\.de/(?:[^/]+/)*[^/?#&]*-a-(?P<id>\\d+)\\.html")
			PhoenixIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒarticle       λ.Object
						ϒarticle_id    λ.Object
						ϒcontent_id    λ.Object
						ϒdetails       λ.Object
						ϒduration      λ.Object
						ϒepisode       λ.Object
						ϒinfo          λ.Object
						ϒm             λ.Object
						ϒself          = λargs[0]
						ϒseries        λ.Object
						ϒteaser_images λ.Object
						ϒthumbnail     λ.Object
						ϒthumbnail_key λ.Object
						ϒthumbnail_url λ.Object
						ϒthumbnails    λ.Object
						ϒtimestamp     λ.Object
						ϒtitle         λ.Object
						ϒurl           = λargs[1]
						ϒvideo         λ.Object
						ϒvideo_id      λ.Object
						τmp0           λ.Object
						τmp1           λ.Object
						τmp2           λ.Object
					)
					ϒarticle_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒarticle = λ.Calm(ϒself, "_download_json", λ.Mod(λ.StrLiteral("https://www.phoenix.de/response/id/%s"), ϒarticle_id), ϒarticle_id, λ.StrLiteral("Downloading article JSON"))
					ϒvideo = λ.GetItem(λ.GetItem(ϒarticle, λ.StrLiteral("absaetze")), λ.IntLiteral(0))
					ϒtitle = func() λ.Object {
						if λv := λ.Calm(ϒvideo, "get", λ.StrLiteral("titel")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Calm(ϒarticle, "get", λ.StrLiteral("subtitel"))
						}
					}()
					if λ.IsTrue(λ.Eq(λ.Calm(ϒvideo, "get", λ.StrLiteral("typ")), λ.StrLiteral("video-youtube"))) {
						ϒvideo_id = λ.GetItem(ϒvideo, λ.StrLiteral("id"))
						return λ.Call(λ.GetAttr(ϒself, "url_result", nil), λ.NewArgs(ϒvideo_id), λ.KWArgs{
							{Name: "ie", Value: λ.Calm(YoutubeIE, "ie_key")},
							{Name: "video_id", Value: ϒvideo_id},
							{Name: "video_title", Value: ϒtitle},
						})
					}
					ϒvideo_id = λ.Cal(ϒcompat_str, func() λ.Object {
						if λv := λ.Calm(ϒvideo, "get", λ.StrLiteral("basename")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Calm(ϒvideo, "get", λ.StrLiteral("content"))
						}
					}())
					ϒdetails = λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
						λ.StrLiteral("https://www.phoenix.de/php/mediaplayer/data/beitrags_details.php"),
						ϒvideo_id,
						λ.StrLiteral("Downloading details JSON"),
					), λ.KWArgs{
						{Name: "query", Value: λ.DictLiteral(map[string]λ.Object{
							"ak":      λ.StrLiteral("web"),
							"ptmd":    λ.StrLiteral("true"),
							"id":      ϒvideo_id,
							"profile": λ.StrLiteral("player2"),
						})},
					})
					ϒtitle = func() λ.Object {
						if λv := ϒtitle; λ.IsTrue(λv) {
							return λv
						} else {
							return λ.GetItem(ϒdetails, λ.StrLiteral("title"))
						}
					}()
					ϒcontent_id = λ.GetItem(λ.GetItem(λ.GetItem(λ.GetItem(ϒdetails, λ.StrLiteral("tracking")), λ.StrLiteral("nielsen")), λ.StrLiteral("content")), λ.StrLiteral("assetid"))
					ϒinfo = λ.Calm(ϒself, "_extract_ptmd", λ.Mod(λ.StrLiteral("https://tmd.phoenix.de/tmd/2/ngplayer_2_3/vod/ptmd/phoenix/%s"), ϒcontent_id), ϒcontent_id, λ.None, ϒurl)
					ϒduration = λ.Cal(ϒint_or_none, λ.Cal(ϒtry_get, ϒdetails, λ.NewFunction("<lambda>",
						[]λ.Param{
							{Name: "x"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒx = λargs[0]
							)
							return λ.GetItem(λ.GetItem(λ.GetItem(λ.GetItem(ϒx, λ.StrLiteral("tracking")), λ.StrLiteral("nielsen")), λ.StrLiteral("content")), λ.StrLiteral("length"))
						})))
					ϒtimestamp = λ.Cal(ϒunified_timestamp, λ.Calm(ϒdetails, "get", λ.StrLiteral("editorialDate")))
					ϒseries = λ.Cal(ϒtry_get, ϒdetails, λ.NewFunction("<lambda>",
						[]λ.Param{
							{Name: "x"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒx = λargs[0]
							)
							return λ.GetItem(λ.GetItem(λ.GetItem(λ.GetItem(ϒx, λ.StrLiteral("tracking")), λ.StrLiteral("nielsen")), λ.StrLiteral("content")), λ.StrLiteral("program"))
						}), ϒcompat_str)
					ϒepisode = func() λ.Object {
						if λ.IsTrue(λ.Eq(λ.Calm(ϒdetails, "get", λ.StrLiteral("contentType")), λ.StrLiteral("episode"))) {
							return ϒtitle
						} else {
							return λ.None
						}
					}()
					ϒthumbnails = λ.NewList()
					ϒteaser_images = func() λ.Object {
						if λv := λ.Cal(ϒtry_get, ϒdetails, λ.NewFunction("<lambda>",
							[]λ.Param{
								{Name: "x"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒx = λargs[0]
								)
								return λ.GetItem(λ.GetItem(ϒx, λ.StrLiteral("teaserImageRef")), λ.StrLiteral("layouts"))
							}), λ.DictType); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.DictLiteral(map[λ.Object]λ.Object{})
						}
					}()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(ϒteaser_images, "items"))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						τmp2 = λ.UnpackIterable(τmp1, 2)
						ϒthumbnail_key = λ.GetItem(τmp2, λ.IntLiteral(0))
						ϒthumbnail_url = λ.GetItem(τmp2, λ.IntLiteral(1))
						ϒthumbnail_url = λ.Cal(ϒurljoin, ϒurl, ϒthumbnail_url)
						if !λ.IsTrue(ϒthumbnail_url) {
							continue
						}
						ϒthumbnail = λ.DictLiteral(map[string]λ.Object{
							"url": ϒthumbnail_url,
						})
						ϒm = λ.Cal(Ωre.ϒmatch, λ.StrLiteral("^([0-9]+)x([0-9]+)$"), ϒthumbnail_key)
						if λ.IsTrue(ϒm) {
							λ.SetItem(ϒthumbnail, λ.StrLiteral("width"), λ.Cal(λ.IntType, λ.Calm(ϒm, "group", λ.IntLiteral(1))))
							λ.SetItem(ϒthumbnail, λ.StrLiteral("height"), λ.Cal(λ.IntType, λ.Calm(ϒm, "group", λ.IntLiteral(2))))
						}
						λ.Calm(ϒthumbnails, "append", ϒthumbnail)
					}
					return λ.Cal(ϒmerge_dicts, ϒinfo, λ.DictLiteral(map[string]λ.Object{
						"id":          ϒcontent_id,
						"title":       ϒtitle,
						"description": λ.Calm(ϒdetails, "get", λ.StrLiteral("leadParagraph")),
						"duration":    ϒduration,
						"thumbnails":  ϒthumbnails,
						"timestamp":   ϒtimestamp,
						"uploader":    λ.Calm(ϒdetails, "get", λ.StrLiteral("tvService")),
						"series":      ϒseries,
						"episode":     ϒepisode,
					}))
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"IE_NAME":       PhoenixIE_IE_NAME,
				"_VALID_URL":    PhoenixIE__VALID_URL,
				"_real_extract": PhoenixIE__real_extract,
			})
		}())
	})
}
