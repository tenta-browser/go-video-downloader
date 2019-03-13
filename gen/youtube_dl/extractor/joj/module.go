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
 * joj/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/joj.py
 */

package joj

import (
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor λ.Object
	JojIE         λ.Object
	ϒcompat_str   λ.Object
	ϒint_or_none  λ.Object
	ϒjs_to_json   λ.Object
	ϒtry_get      λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_str = Ωcompat.ϒcompat_str
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒjs_to_json = Ωutils.ϒjs_to_json
		ϒtry_get = Ωutils.ϒtry_get
		JojIE = λ.Cal(λ.TypeType, λ.NewStr("JojIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				JojIE__TESTS        λ.Object
				JojIE__VALID_URL    λ.Object
				JojIE__real_extract λ.Object
			)
			JojIE__VALID_URL = λ.NewStr("(?x)\n                    (?:\n                        joj:|\n                        https?://media\\.joj\\.sk/embed/\n                    )\n                    (?P<id>[^/?#^]+)\n                ")
			JojIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("https://media.joj.sk/embed/a388ec4c-6019-4a4a-9312-b1bee194e932"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):        λ.NewStr("a388ec4c-6019-4a4a-9312-b1bee194e932"),
						λ.NewStr("ext"):       λ.NewStr("mp4"),
						λ.NewStr("title"):     λ.NewStr("NOVÉ BÝVANIE"),
						λ.NewStr("thumbnail"): λ.NewStr("re:^https?://.*\\.jpg$"),
						λ.NewStr("duration"):  λ.NewInt(3118),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://media.joj.sk/embed/9i1cxv"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("joj:a388ec4c-6019-4a4a-9312-b1bee194e932"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("joj:9i1cxv"),
					λ.NewStr("only_matching"): λ.True,
				}),
			)
			JojIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒbitrates   λ.Object
						ϒduration   λ.Object
						ϒfile_el    λ.Object
						ϒformat_id  λ.Object
						ϒformat_url λ.Object
						ϒformats    λ.Object
						ϒheight     λ.Object
						ϒpath       λ.Object
						ϒplaylist   λ.Object
						ϒself       = λargs[0]
						ϒthumbnail  λ.Object
						ϒtitle      λ.Object
						ϒurl        = λargs[1]
						ϒvideo_id   λ.Object
						ϒwebpage    λ.Object
						τmp0        λ.Object
						τmp1        λ.Object
					)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), λ.Mod(λ.NewStr("https://media.joj.sk/embed/%s"), ϒvideo_id), ϒvideo_id)
					ϒtitle = func() λ.Object {
						if λv := λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
							λ.NewTuple(
								λ.NewStr("videoTitle\\s*:\\s*([\"\\'])(?P<title>(?:(?!\\1).)+)\\1"),
								λ.NewStr("<title>(?P<title>[^<]+)"),
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
					ϒbitrates = λ.Call(λ.GetAttr(ϒself, "_parse_json", nil), λ.NewArgs(
						λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
							λ.NewStr("(?s)(?:src|bitrates)\\s*=\\s*({.+?});"),
							ϒwebpage,
							λ.NewStr("bitrates"),
						), λ.KWArgs{
							{Name: "default", Value: λ.NewStr("{}")},
						}),
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "transform_source", Value: ϒjs_to_json},
						{Name: "fatal", Value: λ.False},
					})
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, func() λ.Object {
						if λv := λ.Cal(ϒtry_get, ϒbitrates, λ.NewFunction("<lambda>",
							[]λ.Param{
								{Name: "x"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒx = λargs[0]
								)
								return λ.GetItem(ϒx, λ.NewStr("mp4"))
							}), λ.ListType); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.NewList()
						}
					}())
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒformat_url = τmp1
						if λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒformat_url, ϒcompat_str)) {
							ϒheight = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
								λ.NewStr("(\\d+)[pP]\\."),
								ϒformat_url,
								λ.NewStr("height"),
							), λ.KWArgs{
								{Name: "default", Value: λ.None},
							})
							λ.Cal(λ.GetAttr(ϒformats, "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
								λ.NewStr("url"): ϒformat_url,
								λ.NewStr("format_id"): func() λ.Object {
									if λ.IsTrue(ϒheight) {
										return λ.Mod(λ.NewStr("%sp"), ϒheight)
									} else {
										return λ.None
									}
								}(),
								λ.NewStr("height"): λ.Cal(λ.IntType, ϒheight),
							}))
						}
					}
					if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒformats))) {
						ϒplaylist = λ.Cal(λ.GetAttr(ϒself, "_download_xml", nil), λ.Mod(λ.NewStr("https://media.joj.sk/services/Video.php?clip=%s"), ϒvideo_id), ϒvideo_id)
						τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(ϒplaylist, "findall", nil), λ.NewStr("./files/file")))
						for {
							if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
								break
							}
							ϒfile_el = τmp1
							ϒpath = λ.Cal(λ.GetAttr(ϒfile_el, "get", nil), λ.NewStr("path"))
							if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒpath))) {
								continue
							}
							ϒformat_id = func() λ.Object {
								if λv := λ.Cal(λ.GetAttr(ϒfile_el, "get", nil), λ.NewStr("id")); λ.IsTrue(λv) {
									return λv
								} else {
									return λ.Cal(λ.GetAttr(ϒfile_el, "get", nil), λ.NewStr("label"))
								}
							}()
							λ.Cal(λ.GetAttr(ϒformats, "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
								λ.NewStr("url"):       λ.Mod(λ.NewStr("http://n16.joj.sk/storage/%s"), λ.Cal(λ.GetAttr(ϒpath, "replace", nil), λ.NewStr("dat/"), λ.NewStr(""), λ.NewInt(1))),
								λ.NewStr("format_id"): ϒformat_id,
								λ.NewStr("height"): λ.Cal(ϒint_or_none, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
									λ.NewStr("(\\d+)[pP]"),
									func() λ.Object {
										if λv := ϒformat_id; λ.IsTrue(λv) {
											return λv
										} else {
											return ϒpath
										}
									}(),
									λ.NewStr("height"),
								), λ.KWArgs{
									{Name: "default", Value: λ.None},
								})),
							}))
						}
					}
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					ϒthumbnail = λ.Cal(λ.GetAttr(ϒself, "_og_search_thumbnail", nil), ϒwebpage)
					ϒduration = λ.Cal(ϒint_or_none, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr("videoDuration\\s*:\\s*(\\d+)"),
						ϒwebpage,
						λ.NewStr("duration"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					}))
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):        ϒvideo_id,
						λ.NewStr("title"):     ϒtitle,
						λ.NewStr("thumbnail"): ϒthumbnail,
						λ.NewStr("duration"):  ϒduration,
						λ.NewStr("formats"):   ϒformats,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TESTS"):        JojIE__TESTS,
				λ.NewStr("_VALID_URL"):    JojIE__VALID_URL,
				λ.NewStr("_real_extract"): JojIE__real_extract,
			})
		}())
	})
}
