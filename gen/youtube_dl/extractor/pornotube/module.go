// Code generated by transpiler. DO NOT EDIT.

/**
 * Go Video Downloader
 *
 *    Copyright 2018 Tenta, LLC
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
 * pornotube/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/pornotube.py
 */

package pornotube

import (
	Ωjson "github.com/tenta-browser/go-video-downloader/gen/json"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor λ.Object
	PornotubeIE   λ.Object
	ϒint_or_none  λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒint_or_none = Ωutils.ϒint_or_none
		PornotubeIE = λ.Cal(λ.TypeType, λ.NewStr("PornotubeIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				PornotubeIE__TEST         λ.Object
				PornotubeIE__VALID_URL    λ.Object
				PornotubeIE__real_extract λ.Object
			)
			PornotubeIE__VALID_URL = λ.NewStr("https?://(?:\\w+\\.)?pornotube\\.com/(?:[^?#]*?)/video/(?P<id>[0-9]+)")
			PornotubeIE__TEST = λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("url"): λ.NewStr("http://www.pornotube.com/orientation/straight/video/4964/title/weird-hot-and-wet-science"),
				λ.NewStr("md5"): λ.NewStr("60fc5a4f0d93a97968fc7999d98260c9"),
				λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("id"):          λ.NewStr("4964"),
					λ.NewStr("ext"):         λ.NewStr("mp4"),
					λ.NewStr("upload_date"): λ.NewStr("20141203"),
					λ.NewStr("title"):       λ.NewStr("Weird Hot and Wet Science"),
					λ.NewStr("description"): λ.NewStr("md5:a8304bef7ef06cb4ab476ca6029b01b0"),
					λ.NewStr("categories"): λ.NewList(
						λ.NewStr("Adult Humor"),
						λ.NewStr("Blondes"),
					),
					λ.NewStr("uploader"):  λ.NewStr("Alpha Blue Archives"),
					λ.NewStr("thumbnail"): λ.NewStr("re:^https?://.*\\.jpg$"),
					λ.NewStr("timestamp"): λ.NewInt(1417582800),
					λ.NewStr("age_limit"): λ.NewInt(18),
				}),
			})
			PornotubeIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						FIELDS                λ.Object
						ϒcategories           λ.Object
						ϒduration             λ.Object
						ϒend                  λ.Object
						ϒinfo                 λ.Object
						ϒmovie_id             λ.Object
						ϒprimary_image_number λ.Object
						ϒself                 = λargs[0]
						ϒstart                λ.Object
						ϒthumbnail            λ.Object
						ϒtimestamp            λ.Object
						ϒtitle                λ.Object
						ϒtoken                λ.Object
						ϒuploader             λ.Object
						ϒurl                  = λargs[1]
						ϒvideo_id             λ.Object
						ϒvideo_url            λ.Object
					)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒtoken = λ.GetItem(λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
						λ.NewStr("https://api.aebn.net/auth/v2/origins/authenticate"),
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "note", Value: λ.NewStr("Downloading token")},
						{Name: "data", Value: λ.Cal(λ.GetAttr(λ.Cal(Ωjson.ϒdumps, λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("credentials"): λ.NewStr("Clip Application"),
						})), "encode", nil), λ.NewStr("utf-8"))},
						{Name: "headers", Value: λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("Content-Type"): λ.NewStr("application/json"),
							λ.NewStr("Origin"):       λ.NewStr("http://www.pornotube.com"),
						})},
					}), λ.NewStr("tokenKey"))
					ϒvideo_url = λ.GetItem(λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
						λ.Mod(λ.NewStr("https://api.aebn.net/delivery/v1/clips/%s/MP4"), ϒvideo_id),
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "note", Value: λ.NewStr("Downloading delivery information")},
						{Name: "headers", Value: λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("Authorization"): ϒtoken,
						})},
					}), λ.NewStr("mediaUrl"))
					FIELDS = λ.NewTuple(
						λ.NewStr("title"),
						λ.NewStr("description"),
						λ.NewStr("startSecond"),
						λ.NewStr("endSecond"),
						λ.NewStr("publishDate"),
						λ.NewStr("studios{name}"),
						λ.NewStr("categories{name}"),
						λ.NewStr("movieId"),
						λ.NewStr("primaryImageNumber"),
					)
					ϒinfo = λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
						λ.Mod(λ.NewStr("https://api.aebn.net/content/v2/clips/%s?fields=%s"), λ.NewTuple(
							ϒvideo_id,
							λ.Cal(λ.GetAttr(λ.NewStr(","), "join", nil), FIELDS),
						)),
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "note", Value: λ.NewStr("Downloading metadata")},
						{Name: "headers", Value: λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("Authorization"): ϒtoken,
						})},
					})
					if λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒinfo, λ.ListType)) {
						ϒinfo = λ.GetItem(ϒinfo, λ.NewInt(0))
					}
					ϒtitle = λ.GetItem(ϒinfo, λ.NewStr("title"))
					ϒtimestamp = λ.Call(ϒint_or_none, λ.NewArgs(λ.Cal(λ.GetAttr(ϒinfo, "get", nil), λ.NewStr("publishDate"))), λ.KWArgs{
						{Name: "scale", Value: λ.NewInt(1000)},
					})
					ϒuploader = λ.Cal(λ.GetAttr(λ.GetItem(λ.Cal(λ.GetAttr(ϒinfo, "get", nil), λ.NewStr("studios"), λ.NewList(λ.NewDictWithTable(map[λ.Object]λ.Object{}))), λ.NewInt(0)), "get", nil), λ.NewStr("name"))
					ϒmovie_id = λ.Cal(λ.GetAttr(ϒinfo, "get", nil), λ.NewStr("movieId"))
					ϒprimary_image_number = λ.Cal(λ.GetAttr(ϒinfo, "get", nil), λ.NewStr("primaryImageNumber"))
					ϒthumbnail = λ.None
					if λ.IsTrue(func() λ.Object {
						if λv := ϒmovie_id; !λ.IsTrue(λv) {
							return λv
						} else {
							return ϒprimary_image_number
						}
					}()) {
						ϒthumbnail = λ.Mod(λ.NewStr("http://pic.aebn.net/dis/t/%s/%s_%08d.jpg"), λ.NewTuple(
							ϒmovie_id,
							ϒmovie_id,
							ϒprimary_image_number,
						))
					}
					ϒstart = λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒinfo, "get", nil), λ.NewStr("startSecond")))
					ϒend = λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒinfo, "get", nil), λ.NewStr("endSecond")))
					ϒduration = func() λ.Object {
						if λ.IsTrue(func() λ.Object {
							if λv := ϒstart; !λ.IsTrue(λv) {
								return λv
							} else {
								return ϒend
							}
						}()) {
							return λ.Sub(ϒend, ϒstart)
						} else {
							return λ.None
						}
					}()
					ϒcategories = λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
						nil,
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							return λ.NewGenerator(func(λgen λ.Generator) λ.Object {
								var (
									ϒc   λ.Object
									τmp0 λ.Object
									τmp1 λ.Object
								)
								τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(ϒinfo, "get", nil), λ.NewStr("categories"), λ.NewList()))
								for {
									if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
										break
									}
									ϒc = τmp1
									if λ.IsTrue(λ.Cal(λ.GetAttr(ϒc, "get", nil), λ.NewStr("name"))) {
										λgen.Yield(λ.GetItem(ϒc, λ.NewStr("name")))
									}
								}
								return λ.None
							})
						})))
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          ϒvideo_id,
						λ.NewStr("url"):         ϒvideo_url,
						λ.NewStr("title"):       ϒtitle,
						λ.NewStr("description"): λ.Cal(λ.GetAttr(ϒinfo, "get", nil), λ.NewStr("description")),
						λ.NewStr("duration"):    ϒduration,
						λ.NewStr("timestamp"):   ϒtimestamp,
						λ.NewStr("uploader"):    ϒuploader,
						λ.NewStr("thumbnail"):   ϒthumbnail,
						λ.NewStr("categories"):  ϒcategories,
						λ.NewStr("age_limit"):   λ.NewInt(18),
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TEST"):         PornotubeIE__TEST,
				λ.NewStr("_VALID_URL"):    PornotubeIE__VALID_URL,
				λ.NewStr("_real_extract"): PornotubeIE__real_extract,
			})
		}())
	})
}
