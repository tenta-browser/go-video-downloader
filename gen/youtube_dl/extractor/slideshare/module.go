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
 * slideshare/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/slideshare.py
 */

package slideshare

import (
	Ωjson "github.com/tenta-browser/go-video-downloader/gen/json"
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωparse "github.com/tenta-browser/go-video-downloader/gen/urllib/parse"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError     λ.Object
	InfoExtractor      λ.Object
	SlideshareIE       λ.Object
	ϒget_element_by_id λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ExtractorError = Ωutils.ExtractorError
		ϒget_element_by_id = Ωutils.ϒget_element_by_id
		SlideshareIE = λ.Cal(λ.TypeType, λ.NewStr("SlideshareIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				SlideshareIE__TEST         λ.Object
				SlideshareIE__VALID_URL    λ.Object
				SlideshareIE__real_extract λ.Object
			)
			SlideshareIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?slideshare\\.net/[^/]+?/(?P<title>.+?)($|\\?)")
			SlideshareIE__TEST = λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("url"): λ.NewStr("http://www.slideshare.net/Dataversity/keynote-presentation-managing-scale-and-complexity"),
				λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("id"):          λ.NewStr("25665706"),
					λ.NewStr("ext"):         λ.NewStr("mp4"),
					λ.NewStr("title"):       λ.NewStr("Managing Scale and Complexity"),
					λ.NewStr("description"): λ.NewStr("This was a keynote presentation at the NoSQL Now! 2013 Conference & Expo (http://www.nosqlnow.com). This presentation was given by Adrian Cockcroft from Netflix."),
				}),
			})
			SlideshareIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒbucket         λ.Object
						ϒdescription    λ.Object
						ϒdoc            λ.Object
						ϒext            λ.Object
						ϒinfo           λ.Object
						ϒmobj           λ.Object
						ϒpage_title     λ.Object
						ϒself           = λargs[0]
						ϒslideshare_obj λ.Object
						ϒurl            = λargs[1]
						ϒvideo_url      λ.Object
						ϒwebpage        λ.Object
					)
					ϒmobj = λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl)
					ϒpage_title = λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("title"))
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒpage_title)
					ϒslideshare_obj = λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewStr("\\$\\.extend\\(.*?slideshare_object,\\s*(\\{.*?\\})\\);"), ϒwebpage, λ.NewStr("slideshare object"))
					ϒinfo = λ.Cal(Ωjson.ϒloads, ϒslideshare_obj)
					if λ.IsTrue(λ.Ne(λ.GetItem(λ.GetItem(ϒinfo, λ.NewStr("slideshow")), λ.NewStr("type")), λ.NewStr("video"))) {
						panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.Mod(λ.NewStr("Webpage type is \"%s\": only video extraction is supported for Slideshare"), λ.GetItem(λ.GetItem(ϒinfo, λ.NewStr("slideshow")), λ.NewStr("type")))), λ.KWArgs{
							{Name: "expected", Value: λ.True},
						})))
					}
					ϒdoc = λ.GetItem(ϒinfo, λ.NewStr("doc"))
					ϒbucket = λ.GetItem(λ.GetItem(ϒinfo, λ.NewStr("jsplayer")), λ.NewStr("video_bucket"))
					ϒext = λ.GetItem(λ.GetItem(ϒinfo, λ.NewStr("jsplayer")), λ.NewStr("video_extension"))
					ϒvideo_url = λ.Cal(Ωparse.ϒurljoin, ϒbucket, λ.Add(λ.Add(ϒdoc, λ.NewStr("-SD.")), ϒext))
					ϒdescription = func() λ.Object {
						if λv := λ.Cal(ϒget_element_by_id, λ.NewStr("slideshow-description-paragraph"), ϒwebpage); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
								λ.NewStr("(?s)<p[^>]+itemprop=\"description\"[^>]*>(.+?)</p>"),
								ϒwebpage,
								λ.NewStr("description"),
							), λ.KWArgs{
								{Name: "fatal", Value: λ.False},
							})
						}
					}()
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("_type"):     λ.NewStr("video"),
						λ.NewStr("id"):        λ.GetItem(λ.GetItem(ϒinfo, λ.NewStr("slideshow")), λ.NewStr("id")),
						λ.NewStr("title"):     λ.GetItem(λ.GetItem(ϒinfo, λ.NewStr("slideshow")), λ.NewStr("title")),
						λ.NewStr("ext"):       ϒext,
						λ.NewStr("url"):       ϒvideo_url,
						λ.NewStr("thumbnail"): λ.GetItem(λ.GetItem(ϒinfo, λ.NewStr("slideshow")), λ.NewStr("pin_image_url")),
						λ.NewStr("description"): func() λ.Object {
							if λ.IsTrue(ϒdescription) {
								return λ.Cal(λ.GetAttr(ϒdescription, "strip", nil))
							} else {
								return λ.None
							}
						}(),
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TEST"):         SlideshareIE__TEST,
				λ.NewStr("_VALID_URL"):    SlideshareIE__VALID_URL,
				λ.NewStr("_real_extract"): SlideshareIE__real_extract,
			})
		}())
	})
}
