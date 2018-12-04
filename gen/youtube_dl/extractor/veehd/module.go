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
 * veehd/module.go: transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/veehd.py
 */

package veehd

import (
	Ωjson "github.com/tenta-browser/go-video-downloader/gen/json"
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωparse "github.com/tenta-browser/go-video-downloader/gen/urllib/parse"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError               λ.Object
	InfoExtractor                λ.Object
	VeeHDIE                      λ.Object
	ϒclean_html                  λ.Object
	ϒcompat_urllib_parse_unquote λ.Object
	ϒget_element_by_id           λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_urllib_parse_unquote = Ωcompat.ϒcompat_urllib_parse_unquote
		ExtractorError = Ωutils.ExtractorError
		ϒclean_html = Ωutils.ϒclean_html
		ϒget_element_by_id = Ωutils.ϒget_element_by_id
		VeeHDIE = λ.Cal(λ.TypeType, λ.NewStr("VeeHDIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				VeeHDIE__TESTS        λ.Object
				VeeHDIE__VALID_URL    λ.Object
				VeeHDIE__real_extract λ.Object
			)
			VeeHDIE__VALID_URL = λ.NewStr("https?://veehd\\.com/video/(?P<id>\\d+)")
			VeeHDIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://veehd.com/video/4639434_Solar-Sinter"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("4639434"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Solar Sinter"),
						λ.NewStr("uploader_id"): λ.NewStr("VideoEyes"),
						λ.NewStr("description"): λ.NewStr("md5:46a840e8692ddbaffb5f81d9885cb457"),
					}),
					λ.NewStr("skip"): λ.NewStr("Video deleted"),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://veehd.com/video/4905758_Elysian-Fields-Channeling"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("4905758"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Elysian Fields - Channeling"),
						λ.NewStr("description"): λ.NewStr("md5:360e4e95fdab58aefbea0f2a19e5604b"),
						λ.NewStr("uploader_id"): λ.NewStr("spotted"),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://veehd.com/video/2046729_2012-2009-DivX-Trailer"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("2046729"),
						λ.NewStr("ext"):         λ.NewStr("avi"),
						λ.NewStr("title"):       λ.NewStr("2012 (2009) DivX Trailer"),
						λ.NewStr("description"): λ.NewStr("md5:75435ee95255e6a9838ac6f6f3a2396b"),
						λ.NewStr("uploader_id"): λ.NewStr("Movie_Trailers"),
					}),
				}),
			)
			VeeHDIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒconfig      λ.Object
						ϒconfig_json λ.Object
						ϒdescription λ.Object
						ϒiframe_page λ.Object
						ϒiframe_src  λ.Object
						ϒiframe_url  λ.Object
						ϒplayer_page λ.Object
						ϒplayer_path λ.Object
						ϒplayer_url  λ.Object
						ϒself        = λargs[0]
						ϒthumbnail   λ.Object
						ϒtitle       λ.Object
						ϒuploader_id λ.Object
						ϒurl         = λargs[1]
						ϒvideo_id    λ.Object
						ϒvideo_url   λ.Object
						ϒwebpage     λ.Object
					)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒvideo_id, λ.NewStr("Requesting webpage"))
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒvideo_id)
					if λ.IsTrue(λ.NewBool(λ.Contains(ϒwebpage, λ.NewStr("This video has been removed<")))) {
						panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.Mod(λ.NewStr("Video %s has been removed"), ϒvideo_id)), λ.KWArgs{
							{Name: "expected", Value: λ.True},
						})))
					}
					ϒplayer_path = λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewStr("\\$\\(\"#playeriframe\"\\).attr\\({src : \"(.+?)\""), ϒwebpage, λ.NewStr("player path"))
					ϒplayer_url = λ.Cal(Ωparse.ϒurljoin, ϒurl, ϒplayer_path)
					λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒplayer_url, ϒvideo_id, λ.NewStr("Requesting player page"))
					ϒplayer_page = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒplayer_url, ϒvideo_id, λ.NewStr("Downloading player page"))
					ϒvideo_url = λ.None
					ϒconfig_json = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr("value=\\'config=({.+?})\\'"),
						ϒplayer_page,
						λ.NewStr("config json"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					if λ.IsTrue(ϒconfig_json) {
						ϒconfig = λ.Cal(Ωjson.ϒloads, ϒconfig_json)
						ϒvideo_url = λ.Cal(ϒcompat_urllib_parse_unquote, λ.GetItem(λ.GetItem(ϒconfig, λ.NewStr("clip")), λ.NewStr("url")))
					}
					if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒvideo_url))) {
						ϒvideo_url = λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
							λ.NewStr("<embed[^>]+type=\"video/divx\"[^>]+src=\"([^\"]+)\""),
							ϒplayer_page,
							λ.NewStr("video url"),
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
						})
					}
					if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒvideo_url))) {
						ϒiframe_src = λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewStr("<iframe[^>]+src=\"/?([^\"]+)\""), ϒplayer_page, λ.NewStr("iframe url"))
						ϒiframe_url = λ.Mod(λ.NewStr("http://veehd.com/%s"), ϒiframe_src)
						λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒiframe_url, ϒvideo_id, λ.NewStr("Requesting iframe page"))
						ϒiframe_page = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒiframe_url, ϒvideo_id, λ.NewStr("Downloading iframe page"))
						ϒvideo_url = λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewStr("file\\s*:\\s*'([^']+)'"), ϒiframe_page, λ.NewStr("video url"))
					}
					ϒtitle = λ.Cal(ϒclean_html, λ.GetItem(λ.Cal(λ.GetAttr(λ.Cal(ϒget_element_by_id, λ.NewStr("videoName"), ϒwebpage), "rpartition", nil), λ.NewStr("|")), λ.NewInt(0)))
					ϒuploader_id = λ.Cal(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewStr("<a href=\"/profile/\\d+\">(.+?)</a>"), ϒwebpage, λ.NewStr("uploader"))
					ϒthumbnail = λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewStr("<img id=\"veehdpreview\" src=\"(.+?)\""), ϒwebpage, λ.NewStr("thumbnail"))
					ϒdescription = λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
						λ.NewStr("<td class=\"infodropdown\".*?<div>(.*?)<ul"),
						ϒwebpage,
						λ.NewStr("description"),
					), λ.KWArgs{
						{Name: "flags", Value: Ωre.DOTALL},
					})
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("_type"):       λ.NewStr("video"),
						λ.NewStr("id"):          ϒvideo_id,
						λ.NewStr("title"):       ϒtitle,
						λ.NewStr("url"):         ϒvideo_url,
						λ.NewStr("uploader_id"): ϒuploader_id,
						λ.NewStr("thumbnail"):   ϒthumbnail,
						λ.NewStr("description"): ϒdescription,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TESTS"):        VeeHDIE__TESTS,
				λ.NewStr("_VALID_URL"):    VeeHDIE__VALID_URL,
				λ.NewStr("_real_extract"): VeeHDIE__real_extract,
			})
		}())
	})
}