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
 * screencast/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/screencast.py
 */

package screencast

import (
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError   λ.Object
	InfoExtractor    λ.Object
	ScreencastIE     λ.Object
	ϒcompat_parse_qs λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_parse_qs = Ωcompat.ϒcompat_parse_qs
		ExtractorError = Ωutils.ExtractorError
		ScreencastIE = λ.Cal(λ.TypeType, λ.NewStr("ScreencastIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				ScreencastIE__TESTS        λ.Object
				ScreencastIE__VALID_URL    λ.Object
				ScreencastIE__real_extract λ.Object
			)
			ScreencastIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?screencast\\.com/t/(?P<id>[a-zA-Z0-9]+)")
			ScreencastIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.screencast.com/t/3ZEjQXlT"),
					λ.NewStr("md5"): λ.NewStr("917df1c13798a3e96211dd1561fded83"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("3ZEjQXlT"),
						λ.NewStr("ext"):         λ.NewStr("m4v"),
						λ.NewStr("title"):       λ.NewStr("Color Measurement with Ocean Optics Spectrometers"),
						λ.NewStr("description"): λ.NewStr("md5:240369cde69d8bed61349a199c5fb153"),
						λ.NewStr("thumbnail"):   λ.NewStr("re:^https?://.*\\.(?:gif|jpg)$"),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.screencast.com/t/V2uXehPJa1ZI"),
					λ.NewStr("md5"): λ.NewStr("e8e4b375a7660a9e7e35c33973410d34"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("V2uXehPJa1ZI"),
						λ.NewStr("ext"):         λ.NewStr("mov"),
						λ.NewStr("title"):       λ.NewStr("The Amadeus Spectrometer"),
						λ.NewStr("description"): λ.NewStr("re:^In this video, our friends at.*To learn more about Amadeus, visit"),
						λ.NewStr("thumbnail"):   λ.NewStr("re:^https?://.*\\.(?:gif|jpg)$"),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.screencast.com/t/aAB3iowa"),
					λ.NewStr("md5"): λ.NewStr("dedb2734ed00c9755761ccaee88527cd"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("aAB3iowa"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Google Earth Export"),
						λ.NewStr("description"): λ.NewStr("Provides a demo of a CommunityViz export to Google Earth, one of the 3D viewing options."),
						λ.NewStr("thumbnail"):   λ.NewStr("re:^https?://.*\\.(?:gif|jpg)$"),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.screencast.com/t/X3ddTrYh"),
					λ.NewStr("md5"): λ.NewStr("669ee55ff9c51988b4ebc0877cc8b159"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("X3ddTrYh"),
						λ.NewStr("ext"):         λ.NewStr("wmv"),
						λ.NewStr("title"):       λ.NewStr("Toolkit 6 User Group Webinar (2014-03-04) - Default Judgment and First Impression"),
						λ.NewStr("description"): λ.NewStr("md5:7b9f393bc92af02326a5c5889639eab0"),
						λ.NewStr("thumbnail"):   λ.NewStr("re:^https?://.*\\.(?:gif|jpg)$"),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://screencast.com/t/aAB3iowa"),
					λ.NewStr("only_matching"): λ.True,
				}),
			)
			ScreencastIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdescription   λ.Object
						ϒflash_vars    λ.Object
						ϒflash_vars_s  λ.Object
						ϒself          = λargs[0]
						ϒthumbnail     λ.Object
						ϒtitle         λ.Object
						ϒurl           = λargs[1]
						ϒvideo_id      λ.Object
						ϒvideo_meta    λ.Object
						ϒvideo_url     λ.Object
						ϒvideo_url_raw λ.Object
						ϒwebpage       λ.Object
					)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒvideo_id)
					ϒvideo_url = λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
						λ.NewStr("<embed name=\"Video\".*?src=\"([^\"]+)\""),
						ϒwebpage,
						λ.NewStr("QuickTime embed"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					if λ.IsTrue(λ.NewBool(ϒvideo_url == λ.None)) {
						ϒflash_vars_s = λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
							λ.NewStr("<param name=\"flashVars\" value=\"([^\"]+)\""),
							ϒwebpage,
							λ.NewStr("flash vars"),
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
						})
						if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒflash_vars_s))) {
							ϒflash_vars_s = λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
								λ.NewStr("<param name=\"initParams\" value=\"([^\"]+)\""),
								ϒwebpage,
								λ.NewStr("flash vars"),
							), λ.KWArgs{
								{Name: "default", Value: λ.None},
							})
							if λ.IsTrue(ϒflash_vars_s) {
								ϒflash_vars_s = λ.Cal(λ.GetAttr(ϒflash_vars_s, "replace", nil), λ.NewStr(","), λ.NewStr("&"))
							}
						}
						if λ.IsTrue(ϒflash_vars_s) {
							ϒflash_vars = λ.Cal(ϒcompat_parse_qs, ϒflash_vars_s)
							ϒvideo_url_raw = λ.Cal(λ.None, λ.GetItem(λ.GetItem(ϒflash_vars, λ.NewStr("content")), λ.NewInt(0)))
							ϒvideo_url = λ.Cal(λ.GetAttr(ϒvideo_url_raw, "replace", nil), λ.NewStr("http%3A"), λ.NewStr("http:"))
						}
					}
					if λ.IsTrue(λ.NewBool(ϒvideo_url == λ.None)) {
						ϒvideo_meta = λ.Call(λ.GetAttr(ϒself, "_html_search_meta", nil), λ.NewArgs(
							λ.NewStr("og:video"),
							ϒwebpage,
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
						})
						if λ.IsTrue(ϒvideo_meta) {
							ϒvideo_url = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
								λ.NewStr("src=(.*?)(?:$|&)"),
								ϒvideo_meta,
								λ.NewStr("meta tag video URL"),
							), λ.KWArgs{
								{Name: "default", Value: λ.None},
							})
						}
					}
					if λ.IsTrue(λ.NewBool(ϒvideo_url == λ.None)) {
						ϒvideo_url = λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
							λ.NewStr("MediaContentUrl[\"\\']\\s*:([\"\\'])(?P<url>(?:(?!\\1).)+)\\1"),
							ϒwebpage,
							λ.NewStr("video url"),
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
							{Name: "group", Value: λ.NewStr("url")},
						})
					}
					if λ.IsTrue(λ.NewBool(ϒvideo_url == λ.None)) {
						ϒvideo_url = λ.Call(λ.GetAttr(ϒself, "_html_search_meta", nil), λ.NewArgs(
							λ.NewStr("og:video"),
							ϒwebpage,
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
						})
					}
					if λ.IsTrue(λ.NewBool(ϒvideo_url == λ.None)) {
						panic(λ.Raise(λ.Cal(ExtractorError, λ.NewStr("Cannot find video"))))
					}
					ϒtitle = λ.Call(λ.GetAttr(ϒself, "_og_search_title", nil), λ.NewArgs(ϒwebpage), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					if λ.IsTrue(λ.NewBool(ϒtitle == λ.None)) {
						ϒtitle = λ.Cal(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewList(
							λ.NewStr("<b>Title:</b> ([^<]+)</div>"),
							λ.NewStr("class=\"tabSeperator\">></span><span class=\"tabText\">(.+?)<"),
							λ.NewStr("<title>([^<]+)</title>"),
						), ϒwebpage, λ.NewStr("title"))
					}
					ϒthumbnail = λ.Cal(λ.GetAttr(ϒself, "_og_search_thumbnail", nil), ϒwebpage)
					ϒdescription = λ.Call(λ.GetAttr(ϒself, "_og_search_description", nil), λ.NewArgs(ϒwebpage), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					if λ.IsTrue(λ.NewBool(ϒdescription == λ.None)) {
						ϒdescription = λ.Cal(λ.GetAttr(ϒself, "_html_search_meta", nil), λ.NewStr("description"), ϒwebpage)
					}
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          ϒvideo_id,
						λ.NewStr("url"):         ϒvideo_url,
						λ.NewStr("title"):       ϒtitle,
						λ.NewStr("description"): ϒdescription,
						λ.NewStr("thumbnail"):   ϒthumbnail,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TESTS"):        ScreencastIE__TESTS,
				λ.NewStr("_VALID_URL"):    ScreencastIE__VALID_URL,
				λ.NewStr("_real_extract"): ScreencastIE__real_extract,
			})
		}())
	})
}
