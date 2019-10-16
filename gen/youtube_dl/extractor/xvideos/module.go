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
 * xvideos/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/xvideos.py
 */

package xvideos

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError               λ.Object
	InfoExtractor                λ.Object
	XVideosIE                    λ.Object
	ϒclean_html                  λ.Object
	ϒcompat_urllib_parse_unquote λ.Object
	ϒdetermine_ext               λ.Object
	ϒint_or_none                 λ.Object
	ϒparse_duration              λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_urllib_parse_unquote = Ωcompat.ϒcompat_urllib_parse_unquote
		ϒclean_html = Ωutils.ϒclean_html
		ϒdetermine_ext = Ωutils.ϒdetermine_ext
		ExtractorError = Ωutils.ExtractorError
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒparse_duration = Ωutils.ϒparse_duration
		XVideosIE = λ.Cal(λ.TypeType, λ.NewStr("XVideosIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				XVideosIE__TESTS        λ.Object
				XVideosIE__VALID_URL    λ.Object
				XVideosIE__real_extract λ.Object
			)
			XVideosIE__VALID_URL = λ.NewStr("(?x)\n                    https?://\n                        (?:\n                            (?:[^/]+\\.)?xvideos2?\\.com/video|\n                            (?:www\\.)?xvideos\\.es/video|\n                            flashservice\\.xvideos\\.com/embedframe/|\n                            static-hw\\.xvideos\\.com/swf/xv-player\\.swf\\?.*?\\bid_video=\n                        )\n                        (?P<id>[0-9]+)\n                    ")
			XVideosIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.xvideos.com/video4588838/biker_takes_his_girl"),
					λ.NewStr("md5"): λ.NewStr("14cea69fcb84db54293b1e971466c2e1"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):        λ.NewStr("4588838"),
						λ.NewStr("ext"):       λ.NewStr("mp4"),
						λ.NewStr("title"):     λ.NewStr("Biker Takes his Girl"),
						λ.NewStr("duration"):  λ.NewInt(108),
						λ.NewStr("age_limit"): λ.NewInt(18),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://flashservice.xvideos.com/embedframe/4588838"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://static-hw.xvideos.com/swf/xv-player.swf?id_video=4588838"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://xvideos.com/video4588838/biker_takes_his_girl"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://xvideos.com/video4588838/biker_takes_his_girl"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://xvideos.es/video4588838/biker_takes_his_girl"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://www.xvideos.es/video4588838/biker_takes_his_girl"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://xvideos.es/video4588838/biker_takes_his_girl"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://www.xvideos.es/video4588838/biker_takes_his_girl"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://fr.xvideos.com/video4588838/biker_takes_his_girl"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://fr.xvideos.com/video4588838/biker_takes_his_girl"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://it.xvideos.com/video4588838/biker_takes_his_girl"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://it.xvideos.com/video4588838/biker_takes_his_girl"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://de.xvideos.com/video4588838/biker_takes_his_girl"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://de.xvideos.com/video4588838/biker_takes_his_girl"),
					λ.NewStr("only_matching"): λ.True,
				}),
			)
			XVideosIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒduration      λ.Object
						ϒformat_id     λ.Object
						ϒformat_url    λ.Object
						ϒformats       λ.Object
						ϒkind          λ.Object
						ϒmobj          λ.Object
						ϒpreference    λ.Object
						ϒself          = λargs[0]
						ϒthumbnail     λ.Object
						ϒthumbnail_url λ.Object
						ϒthumbnails    λ.Object
						ϒtitle         λ.Object
						ϒurl           = λargs[1]
						ϒvideo_id      λ.Object
						ϒvideo_url     λ.Object
						ϒwebpage       λ.Object
						τmp0           λ.Object
						τmp1           λ.Object
						τmp2           λ.Object
					)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), λ.Mod(λ.NewStr("https://www.xvideos.com/video%s/"), ϒvideo_id), ϒvideo_id)
					ϒmobj = λ.Cal(Ωre.ϒsearch, λ.NewStr("<h1 class=\"inlineError\">(.+?)</h1>"), ϒwebpage)
					if λ.IsTrue(ϒmobj) {
						panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.Mod(λ.NewStr("%s said: %s"), λ.NewTuple(
							λ.GetAttr(ϒself, "IE_NAME", nil),
							λ.Cal(ϒclean_html, λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewInt(1))),
						))), λ.KWArgs{
							{Name: "expected", Value: λ.True},
						})))
					}
					ϒtitle = func() λ.Object {
						if λv := λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
							λ.NewTuple(
								λ.NewStr("<title>(?P<title>.+?)\\s+-\\s+XVID"),
								λ.NewStr("setVideoTitle\\s*\\(\\s*([\"\\'])(?P<title>(?:(?!\\1).)+)\\1"),
							),
							ϒwebpage,
							λ.NewStr("title"),
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
							{Name: "group", Value: λ.NewStr("title")},
						}); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Cal(λ.GetAttr(ϒself, "_og_search_title", nil), ϒwebpage)
						}
					}()
					ϒthumbnails = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.EnumerateIteratorType, λ.NewTuple(
						λ.NewStr(""),
						λ.NewStr("169"),
					)))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						τmp2 = τmp1
						ϒpreference = λ.GetItem(τmp2, λ.NewInt(0))
						ϒthumbnail = λ.GetItem(τmp2, λ.NewInt(1))
						ϒthumbnail_url = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
							λ.Mod(λ.NewStr("setThumbUrl%s\\(\\s*([\"\\'])(?P<thumbnail>(?:(?!\\1).)+)\\1"), ϒthumbnail),
							ϒwebpage,
							λ.NewStr("thumbnail"),
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
							{Name: "group", Value: λ.NewStr("thumbnail")},
						})
						if λ.IsTrue(ϒthumbnail_url) {
							λ.Cal(λ.GetAttr(ϒthumbnails, "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
								λ.NewStr("url"):        ϒthumbnail_url,
								λ.NewStr("preference"): ϒpreference,
							}))
						}
					}
					ϒduration = func() λ.Object {
						if λv := λ.Cal(ϒint_or_none, λ.Call(λ.GetAttr(ϒself, "_og_search_property", nil), λ.NewArgs(
							λ.NewStr("duration"),
							ϒwebpage,
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
						})); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Cal(ϒparse_duration, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
								λ.NewStr("<span[^>]+class=[\"\\']duration[\"\\'][^>]*>.*?(\\d[^<]+)"),
								ϒwebpage,
								λ.NewStr("duration"),
							), λ.KWArgs{
								{Name: "fatal", Value: λ.False},
							}))
						}
					}()
					ϒformats = λ.NewList()
					ϒvideo_url = λ.Cal(ϒcompat_urllib_parse_unquote, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr("flv_url=(.+?)&"),
						ϒwebpage,
						λ.NewStr("video URL"),
					), λ.KWArgs{
						{Name: "default", Value: λ.NewStr("")},
					}))
					if λ.IsTrue(ϒvideo_url) {
						λ.Cal(λ.GetAttr(ϒformats, "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("url"):       ϒvideo_url,
							λ.NewStr("format_id"): λ.NewStr("flv"),
						}))
					}
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(Ωre.ϒfindall, λ.NewStr("setVideo([^(]+)\\(([\"\\'])(http.+?)\\2\\)"), ϒwebpage))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						τmp2 = τmp1
						ϒkind = λ.GetItem(τmp2, λ.NewInt(0))
						_ = λ.GetItem(τmp2, λ.NewInt(1))
						ϒformat_url = λ.GetItem(τmp2, λ.NewInt(2))
						ϒformat_id = λ.Cal(λ.GetAttr(ϒkind, "lower", nil))
						if λ.IsTrue(λ.Eq(ϒformat_id, λ.NewStr("hls"))) {
							λ.Cal(λ.GetAttr(ϒformats, "extend", nil), λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
								ϒformat_url,
								ϒvideo_id,
								λ.NewStr("mp4"),
							), λ.KWArgs{
								{Name: "entry_protocol", Value: λ.NewStr("m3u8_native")},
								{Name: "m3u8_id", Value: λ.NewStr("hls")},
								{Name: "fatal", Value: λ.False},
							}))
						} else {
							if λ.IsTrue(λ.NewBool(λ.Contains(λ.NewTuple(
								λ.NewStr("urllow"),
								λ.NewStr("urlhigh"),
							), ϒformat_id))) {
								λ.Cal(λ.GetAttr(ϒformats, "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
									λ.NewStr("url"): ϒformat_url,
									λ.NewStr("format_id"): λ.Mod(λ.NewStr("%s-%s"), λ.NewTuple(
										λ.Cal(ϒdetermine_ext, ϒformat_url, λ.NewStr("mp4")),
										λ.GetItem(ϒformat_id, λ.NewSlice(λ.NewInt(3), λ.None, λ.None)),
									)),
									λ.NewStr("quality"): func() λ.Object {
										if λ.IsTrue(λ.Cal(λ.GetAttr(ϒformat_id, "endswith", nil), λ.NewStr("low"))) {
											return λ.Neg(λ.NewInt(2))
										} else {
											return λ.None
										}
									}(),
								}))
							}
						}
					}
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):         ϒvideo_id,
						λ.NewStr("formats"):    ϒformats,
						λ.NewStr("title"):      ϒtitle,
						λ.NewStr("duration"):   ϒduration,
						λ.NewStr("thumbnails"): ϒthumbnails,
						λ.NewStr("age_limit"):  λ.NewInt(18),
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TESTS"):        XVideosIE__TESTS,
				λ.NewStr("_VALID_URL"):    XVideosIE__VALID_URL,
				λ.NewStr("_real_extract"): XVideosIE__real_extract,
			})
		}())
	})
}
