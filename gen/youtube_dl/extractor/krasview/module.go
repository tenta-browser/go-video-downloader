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
 * krasview/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/krasview.py
 */

package krasview

import (
	Ωjson "github.com/tenta-browser/go-video-downloader/gen/json"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor λ.Object
	KrasViewIE    λ.Object
	ϒint_or_none  λ.Object
	ϒjs_to_json   λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒjs_to_json = Ωutils.ϒjs_to_json
		KrasViewIE = λ.Cal(λ.TypeType, λ.NewStr("KrasViewIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				KrasViewIE__TEST         λ.Object
				KrasViewIE__VALID_URL    λ.Object
				KrasViewIE__real_extract λ.Object
			)
			KrasViewIE__VALID_URL = λ.NewStr("https?://krasview\\.ru/(?:video|embed)/(?P<id>\\d+)")
			KrasViewIE__TEST = λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("url"): λ.NewStr("http://krasview.ru/video/512228"),
				λ.NewStr("md5"): λ.NewStr("3b91003cf85fc5db277870c8ebd98eae"),
				λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("id"):          λ.NewStr("512228"),
					λ.NewStr("ext"):         λ.NewStr("mp4"),
					λ.NewStr("title"):       λ.NewStr("Снег, лёд, заносы"),
					λ.NewStr("description"): λ.NewStr("Снято в городе Нягань, в Ханты-Мансийском автономном округе."),
					λ.NewStr("duration"):    λ.NewInt(27),
					λ.NewStr("thumbnail"):   λ.NewStr("re:^https?://.*\\.jpg"),
				}),
				λ.NewStr("params"): λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("skip_download"): λ.NewStr("Not accessible from Travis CI server"),
				}),
			})
			KrasViewIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdescription λ.Object
						ϒduration    λ.Object
						ϒflashvars   λ.Object
						ϒheight      λ.Object
						ϒself        = λargs[0]
						ϒthumbnail   λ.Object
						ϒtitle       λ.Object
						ϒurl         = λargs[1]
						ϒvideo_id    λ.Object
						ϒvideo_url   λ.Object
						ϒwebpage     λ.Object
						ϒwidth       λ.Object
					)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒvideo_id)
					ϒflashvars = λ.Cal(Ωjson.ϒloads, λ.Cal(ϒjs_to_json, λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewStr("video_Init\\(({.+?})"), ϒwebpage, λ.NewStr("flashvars"))))
					ϒvideo_url = λ.GetItem(ϒflashvars, λ.NewStr("url"))
					ϒtitle = λ.Cal(λ.GetAttr(ϒself, "_og_search_title", nil), ϒwebpage)
					ϒdescription = λ.Call(λ.GetAttr(ϒself, "_og_search_description", nil), λ.NewArgs(ϒwebpage), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					ϒthumbnail = func() λ.Object {
						if λv := λ.Cal(λ.GetAttr(ϒflashvars, "get", nil), λ.NewStr("image")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Cal(λ.GetAttr(ϒself, "_og_search_thumbnail", nil), ϒwebpage)
						}
					}()
					ϒduration = λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒflashvars, "get", nil), λ.NewStr("duration")))
					ϒwidth = λ.Cal(ϒint_or_none, λ.Call(λ.GetAttr(ϒself, "_og_search_property", nil), λ.NewArgs(
						λ.NewStr("video:width"),
						ϒwebpage,
						λ.NewStr("video width"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					}))
					ϒheight = λ.Cal(ϒint_or_none, λ.Call(λ.GetAttr(ϒself, "_og_search_property", nil), λ.NewArgs(
						λ.NewStr("video:height"),
						ϒwebpage,
						λ.NewStr("video height"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					}))
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          ϒvideo_id,
						λ.NewStr("url"):         ϒvideo_url,
						λ.NewStr("title"):       ϒtitle,
						λ.NewStr("description"): ϒdescription,
						λ.NewStr("thumbnail"):   ϒthumbnail,
						λ.NewStr("duration"):    ϒduration,
						λ.NewStr("width"):       ϒwidth,
						λ.NewStr("height"):      ϒheight,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TEST"):         KrasViewIE__TEST,
				λ.NewStr("_VALID_URL"):    KrasViewIE__VALID_URL,
				λ.NewStr("_real_extract"): KrasViewIE__real_extract,
			})
		}())
	})
}
