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
 * box/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/box.py
 */

package box

import (
	Ωjson "github.com/tenta-browser/go-video-downloader/gen/json"
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	BoxIE             λ.Object
	InfoExtractor     λ.Object
	ϒdetermine_ext    λ.Object
	ϒparse_iso8601    λ.Object
	ϒupdate_url_query λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒdetermine_ext = Ωutils.ϒdetermine_ext
		ϒparse_iso8601 = Ωutils.ϒparse_iso8601
		ϒupdate_url_query = Ωutils.ϒupdate_url_query
		BoxIE = λ.Cal(λ.TypeType, λ.StrLiteral("BoxIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				BoxIE__VALID_URL    λ.Object
				BoxIE__real_extract λ.Object
			)
			BoxIE__VALID_URL = λ.StrLiteral("https?://(?:[^.]+\\.)?app\\.box\\.com/s/(?P<shared_name>[^/]+)/file/(?P<id>\\d+)")
			BoxIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒaccess_token               λ.Object
						ϒauthenticated_download_url λ.Object
						ϒcreator                    λ.Object
						ϒf                          λ.Object
						ϒfile_id                    λ.Object
						ϒformats                    λ.Object
						ϒquery                      λ.Object
						ϒrequest_token              λ.Object
						ϒself                       = λargs[0]
						ϒshared_link                λ.Object
						ϒshared_name                λ.Object
						ϒtitle                      λ.Object
						ϒurl                        = λargs[1]
						ϒwebpage                    λ.Object
						τmp0                        λ.Object
					)
					τmp0 = λ.UnpackIterable(λ.Calm(λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl), "groups"), 2)
					ϒshared_name = λ.GetItem(τmp0, λ.IntLiteral(0))
					ϒfile_id = λ.GetItem(τmp0, λ.IntLiteral(1))
					ϒwebpage = λ.Calm(ϒself, "_download_webpage", ϒurl, ϒfile_id)
					ϒrequest_token = λ.GetItem(λ.Calm(ϒself, "_parse_json", λ.Calm(ϒself, "_search_regex", λ.StrLiteral("Box\\.config\\s*=\\s*({.+?});"), ϒwebpage, λ.StrLiteral("Box config")), ϒfile_id), λ.StrLiteral("requestToken"))
					ϒaccess_token = λ.GetItem(λ.GetItem(λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
						λ.StrLiteral("https://app.box.com/app-api/enduserapp/elements/tokens"),
						ϒfile_id,
						λ.StrLiteral("Downloading token JSON metadata"),
					), λ.KWArgs{
						{Name: "data", Value: λ.Calm(λ.Cal(Ωjson.ϒdumps, λ.DictLiteral(map[string]λ.Object{
							"fileIDs": λ.NewList(ϒfile_id),
						})), "encode")},
						{Name: "headers", Value: λ.DictLiteral(map[string]λ.Object{
							"Content-Type":      λ.StrLiteral("application/json"),
							"X-Request-Token":   ϒrequest_token,
							"X-Box-EndUser-API": λ.Add(λ.StrLiteral("sharedName="), ϒshared_name),
						})},
					}), ϒfile_id), λ.StrLiteral("read"))
					ϒshared_link = λ.Add(λ.StrLiteral("https://app.box.com/s/"), ϒshared_name)
					ϒf = λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
						λ.Add(λ.StrLiteral("https://api.box.com/2.0/files/"), ϒfile_id),
						ϒfile_id,
						λ.StrLiteral("Downloading file JSON metadata"),
					), λ.KWArgs{
						{Name: "headers", Value: λ.DictLiteral(map[string]λ.Object{
							"Authorization": λ.Add(λ.StrLiteral("Bearer "), ϒaccess_token),
							"BoxApi":        λ.Add(λ.StrLiteral("shared_link="), ϒshared_link),
							"X-Rep-Hints":   λ.StrLiteral("[dash]"),
						})},
						{Name: "query", Value: λ.DictLiteral(map[string]string{
							"fields": "authenticated_download_url,created_at,created_by,description,extension,is_download_available,name,representations,size",
						})},
					})
					ϒtitle = λ.GetItem(ϒf, λ.StrLiteral("name"))
					ϒquery = λ.DictLiteral(map[string]λ.Object{
						"access_token": ϒaccess_token,
						"shared_link":  ϒshared_link,
					})
					ϒformats = λ.NewList()
					ϒauthenticated_download_url = λ.Calm(ϒf, "get", λ.StrLiteral("authenticated_download_url"))
					if λ.IsTrue(func() λ.Object {
						if λv := ϒauthenticated_download_url; !λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Calm(ϒf, "get", λ.StrLiteral("is_download_available"))
						}
					}()) {
						λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
							"ext": func() λ.Object {
								if λv := λ.Calm(ϒf, "get", λ.StrLiteral("extension")); λ.IsTrue(λv) {
									return λv
								} else {
									return λ.Cal(ϒdetermine_ext, ϒtitle)
								}
							}(),
							"filesize":  λ.Calm(ϒf, "get", λ.StrLiteral("size")),
							"format_id": λ.StrLiteral("download"),
							"url":       λ.Cal(ϒupdate_url_query, ϒauthenticated_download_url, ϒquery),
						}))
					}
					λ.Calm(ϒself, "_sort_formats", ϒformats)
					ϒcreator = func() λ.Object {
						if λv := λ.Calm(ϒf, "get", λ.StrLiteral("created_by")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.DictLiteral(map[λ.Object]λ.Object{})
						}
					}()
					return λ.DictLiteral(map[string]λ.Object{
						"id":      ϒfile_id,
						"title":   ϒtitle,
						"formats": ϒformats,
						"description": func() λ.Object {
							if λv := λ.Calm(ϒf, "get", λ.StrLiteral("description")); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.None
							}
						}(),
						"uploader":    λ.Calm(ϒcreator, "get", λ.StrLiteral("name")),
						"timestamp":   λ.Cal(ϒparse_iso8601, λ.Calm(ϒf, "get", λ.StrLiteral("created_at"))),
						"uploader_id": λ.Calm(ϒcreator, "get", λ.StrLiteral("id")),
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL":    BoxIE__VALID_URL,
				"_real_extract": BoxIE__real_extract,
			})
		}())
	})
}
