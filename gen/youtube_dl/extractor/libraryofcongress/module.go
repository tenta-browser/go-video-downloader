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
 * libraryofcongress/module.go: transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/libraryofcongress.py
 */

package libraryofcongress

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor       λ.Object
	LibraryOfCongressIE λ.Object
	ϒdetermine_ext      λ.Object
	ϒfloat_or_none      λ.Object
	ϒint_or_none        λ.Object
	ϒparse_filesize     λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒdetermine_ext = Ωutils.ϒdetermine_ext
		ϒfloat_or_none = Ωutils.ϒfloat_or_none
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒparse_filesize = Ωutils.ϒparse_filesize
		LibraryOfCongressIE = λ.Cal(λ.TypeType, λ.NewStr("LibraryOfCongressIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				LibraryOfCongressIE_IE_NAME       λ.Object
				LibraryOfCongressIE__TESTS        λ.Object
				LibraryOfCongressIE__VALID_URL    λ.Object
				LibraryOfCongressIE__real_extract λ.Object
			)
			LibraryOfCongressIE_IE_NAME = λ.NewStr("loc")
			LibraryOfCongressIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?loc\\.gov/(?:item/|today/cyberlc/feature_wdesc\\.php\\?.*\\brec=)(?P<id>[0-9a-z_.]+)")
			LibraryOfCongressIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://loc.gov/item/90716351/"),
					λ.NewStr("md5"): λ.NewStr("6ec0ae8f07f86731b1b2ff70f046210a"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):         λ.NewStr("90716351"),
						λ.NewStr("ext"):        λ.NewStr("mp4"),
						λ.NewStr("title"):      λ.NewStr("Pa's trip to Mars"),
						λ.NewStr("duration"):   λ.NewInt(0),
						λ.NewStr("view_count"): λ.IntType,
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("https://www.loc.gov/today/cyberlc/feature_wdesc.php?rec=5578"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):         λ.NewStr("5578"),
						λ.NewStr("ext"):        λ.NewStr("mp4"),
						λ.NewStr("title"):      λ.NewStr("Help! Preservation Training Needs Here, There & Everywhere"),
						λ.NewStr("duration"):   λ.NewInt(3765),
						λ.NewStr("view_count"): λ.IntType,
						λ.NewStr("subtitles"):  λ.NewStr("mincount:1"),
					}),
					λ.NewStr("params"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("skip_download"): λ.True,
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("https://www.loc.gov/item/78710669/"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):         λ.NewStr("78710669"),
						λ.NewStr("ext"):        λ.NewStr("mp4"),
						λ.NewStr("title"):      λ.NewStr("La vie et la passion de Jesus-Christ"),
						λ.NewStr("duration"):   λ.NewInt(0),
						λ.NewStr("view_count"): λ.IntType,
						λ.NewStr("formats"):    λ.NewStr("mincount:4"),
					}),
					λ.NewStr("params"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("skip_download"): λ.True,
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://www.loc.gov/item/ihas.200197114/"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://www.loc.gov/item/afc1981005_afs20503/"),
					λ.NewStr("only_matching"): λ.True,
				}),
			)
			LibraryOfCongressIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒcc_url        λ.Object
						ϒdata          λ.Object
						ϒderivative    λ.Object
						ϒdownload_url  λ.Object
						ϒdownload_urls λ.Object
						ϒduration      λ.Object
						ϒext           λ.Object
						ϒformat_id     λ.Object
						ϒformats       λ.Object
						ϒhttp_format   λ.Object
						ϒis_video      λ.Object
						ϒm             λ.Object
						ϒmedia_id      λ.Object
						ϒmedia_url     λ.Object
						ϒself          = λargs[0]
						ϒsubtitles     λ.Object
						ϒtitle         λ.Object
						ϒurl           = λargs[1]
						ϒvideo_id      λ.Object
						ϒview_count    λ.Object
						ϒwebpage       λ.Object
						τmp0           λ.Object
						τmp1           λ.Object
					)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒvideo_id)
					ϒmedia_id = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewTuple(
							λ.NewStr("id=([\"\\'])media-player-(?P<id>.+?)\\1"),
							λ.NewStr("<video[^>]+id=([\"\\'])uuid-(?P<id>.+?)\\1"),
							λ.NewStr("<video[^>]+data-uuid=([\"\\'])(?P<id>.+?)\\1"),
							λ.NewStr("mediaObjectId\\s*:\\s*([\"\\'])(?P<id>.+?)\\1"),
							λ.NewStr("data-tab=\"share-media-(?P<id>[0-9A-F]{32})\""),
						),
						ϒwebpage,
						λ.NewStr("media id"),
					), λ.KWArgs{
						{Name: "group", Value: λ.NewStr("id")},
					})
					ϒdata = λ.GetItem(λ.Cal(λ.GetAttr(ϒself, "_download_json", nil), λ.Mod(λ.NewStr("https://media.loc.gov/services/v1/media?id=%s&context=json"), ϒmedia_id), ϒmedia_id), λ.NewStr("mediaObject"))
					ϒderivative = λ.GetItem(λ.GetItem(ϒdata, λ.NewStr("derivatives")), λ.NewInt(0))
					ϒmedia_url = λ.GetItem(ϒderivative, λ.NewStr("derivativeUrl"))
					ϒtitle = func() λ.Object {
						if λv := λ.Cal(λ.GetAttr(ϒderivative, "get", nil), λ.NewStr("shortName")); λ.IsTrue(λv) {
							return λv
						} else if λv := λ.Cal(λ.GetAttr(ϒdata, "get", nil), λ.NewStr("shortName")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Cal(λ.GetAttr(ϒself, "_og_search_title", nil), ϒwebpage)
						}
					}()
					ϒmedia_url = λ.Cal(λ.GetAttr(ϒmedia_url, "replace", nil), λ.NewStr("rtmp"), λ.NewStr("https"))
					ϒis_video = λ.Eq(λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒdata, "get", nil), λ.NewStr("mediaType"), λ.NewStr("v")), "lower", nil)), λ.NewStr("v"))
					ϒext = λ.Cal(ϒdetermine_ext, ϒmedia_url)
					if λ.IsTrue(λ.NewBool(!λ.Contains(λ.NewTuple(
						λ.NewStr("mp4"),
						λ.NewStr("mp3"),
					), ϒext))) {
						τmp0 = λ.IAdd(ϒmedia_url, func() λ.Object {
							if λ.IsTrue(ϒis_video) {
								return λ.NewStr(".mp4")
							} else {
								return λ.NewStr(".mp3")
							}
						}())
						ϒmedia_url = τmp0
					}
					ϒformats = λ.NewList()
					if λ.IsTrue(λ.NewBool(λ.Contains(ϒmedia_url, λ.NewStr("/vod/mp4:")))) {
						λ.Cal(λ.GetAttr(ϒformats, "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("url"):       λ.Add(λ.Cal(λ.GetAttr(ϒmedia_url, "replace", nil), λ.NewStr("/vod/mp4:"), λ.NewStr("/hls-vod/media/")), λ.NewStr(".m3u8")),
							λ.NewStr("format_id"): λ.NewStr("hls"),
							λ.NewStr("ext"):       λ.NewStr("mp4"),
							λ.NewStr("protocol"):  λ.NewStr("m3u8_native"),
							λ.NewStr("quality"):   λ.NewInt(1),
						}))
					}
					ϒhttp_format = λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("url"):       λ.Cal(Ωre.ϒsub, λ.NewStr("(://[^/]+/)(?:[^/]+/)*(?:mp4|mp3):"), λ.NewStr("\\1"), ϒmedia_url),
						λ.NewStr("format_id"): λ.NewStr("http"),
						λ.NewStr("quality"):   λ.NewInt(1),
					})
					if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒis_video))) {
						λ.SetItem(ϒhttp_format, λ.NewStr("vcodec"), λ.NewStr("none"))
					}
					λ.Cal(λ.GetAttr(ϒformats, "append", nil), ϒhttp_format)
					ϒdownload_urls = λ.Cal(λ.SetType)
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(Ωre.ϒfinditer, λ.NewStr("<option[^>]+value=([\"\\'])(?P<url>.+?)\\1[^>]+data-file-download=[^>]+>\\s*(?P<id>.+?)(?:(?:&nbsp;|\\s+)\\((?P<size>.+?)\\))?\\s*<"), ϒwebpage))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒm = τmp1
						ϒformat_id = λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒm, "group", nil), λ.NewStr("id")), "lower", nil))
						if λ.IsTrue(λ.NewBool(λ.Contains(λ.NewTuple(
							λ.NewStr("gif"),
							λ.NewStr("jpeg"),
						), ϒformat_id))) {
							continue
						}
						ϒdownload_url = λ.Cal(λ.GetAttr(ϒm, "group", nil), λ.NewStr("url"))
						if λ.IsTrue(λ.NewBool(λ.Contains(ϒdownload_urls, ϒdownload_url))) {
							continue
						}
						λ.Cal(λ.GetAttr(ϒdownload_urls, "add", nil), ϒdownload_url)
						λ.Cal(λ.GetAttr(ϒformats, "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("url"):             ϒdownload_url,
							λ.NewStr("format_id"):       ϒformat_id,
							λ.NewStr("filesize_approx"): λ.Cal(ϒparse_filesize, λ.Cal(λ.GetAttr(ϒm, "group", nil), λ.NewStr("size"))),
						}))
					}
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					ϒduration = λ.Cal(ϒfloat_or_none, λ.Cal(λ.GetAttr(ϒdata, "get", nil), λ.NewStr("duration")))
					ϒview_count = λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒdata, "get", nil), λ.NewStr("viewCount")))
					ϒsubtitles = λ.NewDictWithTable(map[λ.Object]λ.Object{})
					ϒcc_url = λ.Cal(λ.GetAttr(ϒdata, "get", nil), λ.NewStr("ccUrl"))
					if λ.IsTrue(ϒcc_url) {
						λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒsubtitles, "setdefault", nil), λ.NewStr("en"), λ.NewList()), "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("url"): ϒcc_url,
							λ.NewStr("ext"): λ.NewStr("ttml"),
						}))
					}
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):    ϒvideo_id,
						λ.NewStr("title"): ϒtitle,
						λ.NewStr("thumbnail"): λ.Call(λ.GetAttr(ϒself, "_og_search_thumbnail", nil), λ.NewArgs(ϒwebpage), λ.KWArgs{
							{Name: "default", Value: λ.None},
						}),
						λ.NewStr("duration"):   ϒduration,
						λ.NewStr("view_count"): ϒview_count,
						λ.NewStr("formats"):    ϒformats,
						λ.NewStr("subtitles"):  ϒsubtitles,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("IE_NAME"):       LibraryOfCongressIE_IE_NAME,
				λ.NewStr("_TESTS"):        LibraryOfCongressIE__TESTS,
				λ.NewStr("_VALID_URL"):    LibraryOfCongressIE__VALID_URL,
				λ.NewStr("_real_extract"): LibraryOfCongressIE__real_extract,
			})
		}())
	})
}