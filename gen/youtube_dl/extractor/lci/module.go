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
 * lci/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/lci.py
 */

package lci

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor λ.Object
	LCIIE         λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		LCIIE = λ.Cal(λ.TypeType, λ.NewStr("LCIIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				LCIIE__TEST         λ.Object
				LCIIE__VALID_URL    λ.Object
				LCIIE__real_extract λ.Object
			)
			LCIIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?lci\\.fr/[^/]+/[\\w-]+-(?P<id>\\d+)\\.html")
			LCIIE__TEST = λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("url"): λ.NewStr("http://www.lci.fr/international/etats-unis-a-j-62-hillary-clinton-reste-sans-voix-2001679.html"),
				λ.NewStr("md5"): λ.NewStr("2fdb2538b884d4d695f9bd2bde137e6c"),
				λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("id"):          λ.NewStr("13244802"),
					λ.NewStr("ext"):         λ.NewStr("mp4"),
					λ.NewStr("title"):       λ.NewStr("Hillary Clinton et sa quinte de toux, en plein meeting"),
					λ.NewStr("description"): λ.NewStr("md5:a4363e3a960860132f8124b62f4a01c9"),
				}),
			})
			LCIIE__real_extract = λ.NewFunction("_real_extract",
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
						ϒwat_id   λ.Object
						ϒwebpage  λ.Object
					)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒvideo_id)
					ϒwat_id = λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewTuple(
						λ.NewStr("data-watid=[\\'\"](\\d+)"),
						λ.NewStr("idwat[\"\\']?\\s*:\\s*[\"\\']?(\\d+)"),
					), ϒwebpage, λ.NewStr("wat id"))
					return λ.Cal(λ.GetAttr(ϒself, "url_result", nil), λ.Add(λ.NewStr("wat:"), ϒwat_id), λ.NewStr("Wat"), ϒwat_id)
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TEST"):         LCIIE__TEST,
				λ.NewStr("_VALID_URL"):    LCIIE__VALID_URL,
				λ.NewStr("_real_extract"): LCIIE__real_extract,
			})
		}())
	})
}
