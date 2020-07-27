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
 * myspass/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/myspass.py
 */

package myspass

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor   λ.Object
	MySpassIE       λ.Object
	ϒcompat_str     λ.Object
	ϒint_or_none    λ.Object
	ϒparse_duration λ.Object
	ϒxpath_text     λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_str = Ωcompat.ϒcompat_str
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒparse_duration = Ωutils.ϒparse_duration
		ϒxpath_text = Ωutils.ϒxpath_text
		MySpassIE = λ.Cal(λ.TypeType, λ.StrLiteral("MySpassIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				MySpassIE__VALID_URL    λ.Object
				MySpassIE__real_extract λ.Object
			)
			MySpassIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?myspass\\.de/([^/]+/)*(?P<id>\\d+)")
			MySpassIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒgroup        λ.Object
						ϒgroup_int    λ.Object
						ϒmetadata     λ.Object
						ϒself         = λargs[0]
						ϒtitle        λ.Object
						ϒurl          = λargs[1]
						ϒvideo_id     λ.Object
						ϒvideo_id_int λ.Object
						ϒvideo_url    λ.Object
						τmp0          λ.Object
						τmp1          λ.Object
					)
					ϒvideo_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒmetadata = λ.Calm(ϒself, "_download_xml", λ.Add(λ.StrLiteral("http://www.myspass.de/myspass/includes/apps/video/getvideometadataxml.php?id="), ϒvideo_id), ϒvideo_id)
					ϒtitle = λ.Call(ϒxpath_text, λ.NewArgs(
						ϒmetadata,
						λ.StrLiteral("title"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.True},
					})
					ϒvideo_url = λ.Cal(ϒxpath_text, ϒmetadata, λ.StrLiteral("url_flv"), λ.StrLiteral("download url"), λ.True)
					ϒvideo_id_int = λ.Cal(λ.IntType, ϒvideo_id)
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(λ.Cal(Ωre.ϒsearch, λ.StrLiteral("/myspass2009/\\d+/(\\d+)/(\\d+)/(\\d+)/"), ϒvideo_url), "groups"))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒgroup = τmp1
						ϒgroup_int = λ.Cal(λ.IntType, ϒgroup)
						if λ.IsTrue(λ.Gt(ϒgroup_int, ϒvideo_id_int)) {
							ϒvideo_url = λ.Calm(ϒvideo_url, "replace", ϒgroup, λ.Cal(ϒcompat_str, λ.FloorDiv(ϒgroup_int, ϒvideo_id_int)))
						}
					}
					return λ.DictLiteral(map[string]λ.Object{
						"id":             ϒvideo_id,
						"url":            ϒvideo_url,
						"title":          ϒtitle,
						"thumbnail":      λ.Cal(ϒxpath_text, ϒmetadata, λ.StrLiteral("imagePreview")),
						"description":    λ.Cal(ϒxpath_text, ϒmetadata, λ.StrLiteral("description")),
						"duration":       λ.Cal(ϒparse_duration, λ.Cal(ϒxpath_text, ϒmetadata, λ.StrLiteral("duration"))),
						"series":         λ.Cal(ϒxpath_text, ϒmetadata, λ.StrLiteral("format")),
						"season_number":  λ.Cal(ϒint_or_none, λ.Cal(ϒxpath_text, ϒmetadata, λ.StrLiteral("season"))),
						"season_id":      λ.Cal(ϒxpath_text, ϒmetadata, λ.StrLiteral("season_id")),
						"episode":        ϒtitle,
						"episode_number": λ.Cal(ϒint_or_none, λ.Cal(ϒxpath_text, ϒmetadata, λ.StrLiteral("episode"))),
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL":    MySpassIE__VALID_URL,
				"_real_extract": MySpassIE__real_extract,
			})
		}())
	})
}