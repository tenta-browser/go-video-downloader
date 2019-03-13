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
 * nova/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/nova.py
 */

package nova

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor    λ.Object
	NovaEmbedIE      λ.Object
	NovaIE           λ.Object
	ϒclean_html      λ.Object
	ϒint_or_none     λ.Object
	ϒjs_to_json      λ.Object
	ϒqualities       λ.Object
	ϒunified_strdate λ.Object
	ϒurl_or_none     λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒclean_html = Ωutils.ϒclean_html
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒjs_to_json = Ωutils.ϒjs_to_json
		ϒqualities = Ωutils.ϒqualities
		ϒunified_strdate = Ωutils.ϒunified_strdate
		ϒurl_or_none = Ωutils.ϒurl_or_none
		NovaEmbedIE = λ.Cal(λ.TypeType, λ.NewStr("NovaEmbedIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				NovaEmbedIE__TEST         λ.Object
				NovaEmbedIE__VALID_URL    λ.Object
				NovaEmbedIE__real_extract λ.Object
			)
			NovaEmbedIE__VALID_URL = λ.NewStr("https?://media\\.cms\\.nova\\.cz/embed/(?P<id>[^/?#&]+)")
			NovaEmbedIE__TEST = λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("url"): λ.NewStr("https://media.cms.nova.cz/embed/8o0n0r?autoplay=1"),
				λ.NewStr("md5"): λ.NewStr("b3834f6de5401baabf31ed57456463f7"),
				λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("id"):        λ.NewStr("8o0n0r"),
					λ.NewStr("ext"):       λ.NewStr("mp4"),
					λ.NewStr("title"):     λ.NewStr("2180. díl"),
					λ.NewStr("thumbnail"): λ.NewStr("re:^https?://.*\\.jpg"),
					λ.NewStr("duration"):  λ.NewInt(2578),
				}),
			})
			NovaEmbedIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						QUALITIES    λ.Object
						ϒbitrates    λ.Object
						ϒduration    λ.Object
						ϒf           λ.Object
						ϒf_id        λ.Object
						ϒformat_id   λ.Object
						ϒformat_list λ.Object
						ϒformat_url  λ.Object
						ϒformats     λ.Object
						ϒquality     λ.Object
						ϒquality_key λ.Object
						ϒself        = λargs[0]
						ϒthumbnail   λ.Object
						ϒtitle       λ.Object
						ϒurl         = λargs[1]
						ϒvideo_id    λ.Object
						ϒwebpage     λ.Object
						τmp0         λ.Object
						τmp1         λ.Object
						τmp2         λ.Object
						τmp3         λ.Object
						τmp4         λ.Object
						τmp5         λ.Object
						τmp6         λ.Object
					)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒvideo_id)
					ϒbitrates = λ.Call(λ.GetAttr(ϒself, "_parse_json", nil), λ.NewArgs(
						λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewStr("(?s)(?:src|bitrates)\\s*=\\s*({.+?})\\s*;"), ϒwebpage, λ.NewStr("formats")),
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "transform_source", Value: ϒjs_to_json},
					})
					QUALITIES = λ.NewTuple(
						λ.NewStr("lq"),
						λ.NewStr("mq"),
						λ.NewStr("hq"),
						λ.NewStr("hd"),
					)
					ϒquality_key = λ.Cal(ϒqualities, QUALITIES)
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(ϒbitrates, "items", nil)))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						τmp2 = τmp1
						ϒformat_id = λ.GetItem(τmp2, λ.NewInt(0))
						ϒformat_list = λ.GetItem(τmp2, λ.NewInt(1))
						if λ.IsTrue(λ.NewBool(!λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒformat_list, λ.ListType)))) {
							continue
						}
						τmp2 = λ.Cal(λ.BuiltinIter, ϒformat_list)
						for {
							if τmp3 = λ.NextDefault(τmp2, λ.AfterLast); τmp3 == λ.AfterLast {
								break
							}
							ϒformat_url = τmp3
							ϒformat_url = λ.Cal(ϒurl_or_none, ϒformat_url)
							if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒformat_url))) {
								continue
							}
							ϒf = λ.NewDictWithTable(map[λ.Object]λ.Object{
								λ.NewStr("url"): ϒformat_url,
							})
							ϒf_id = ϒformat_id
							τmp4 = λ.Cal(λ.BuiltinIter, QUALITIES)
							for {
								if τmp5 = λ.NextDefault(τmp4, λ.AfterLast); τmp5 == λ.AfterLast {
									break
								}
								ϒquality = τmp5
								if λ.IsTrue(λ.NewBool(λ.Contains(ϒformat_url, λ.Mod(λ.NewStr("%s.mp4"), ϒquality)))) {
									τmp6 = λ.IAdd(ϒf_id, λ.Mod(λ.NewStr("-%s"), ϒquality))
									ϒf_id = τmp6
									λ.Cal(λ.GetAttr(ϒf, "update", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
										λ.NewStr("quality"):     λ.Cal(ϒquality_key, ϒquality),
										λ.NewStr("format_note"): λ.Cal(λ.GetAttr(ϒquality, "upper", nil)),
									}))
									break
								}
							}
							λ.SetItem(ϒf, λ.NewStr("format_id"), ϒf_id)
							λ.Cal(λ.GetAttr(ϒformats, "append", nil), ϒf)
						}
					}
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					ϒtitle = func() λ.Object {
						if λv := λ.Call(λ.GetAttr(ϒself, "_og_search_title", nil), λ.NewArgs(ϒwebpage), λ.KWArgs{
							{Name: "default", Value: λ.None},
						}); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
								λ.NewTuple(
									λ.NewStr("<value>(?P<title>[^<]+)"),
									λ.NewStr("videoTitle\\s*:\\s*([\"\\'])(?P<value>(?:(?!\\1).)+)\\1"),
								),
								ϒwebpage,
								λ.NewStr("title"),
							), λ.KWArgs{
								{Name: "group", Value: λ.NewStr("value")},
							})
						}
					}()
					ϒthumbnail = func() λ.Object {
						if λv := λ.Call(λ.GetAttr(ϒself, "_og_search_thumbnail", nil), λ.NewArgs(ϒwebpage), λ.KWArgs{
							{Name: "default", Value: λ.None},
						}); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
								λ.NewStr("poster\\s*:\\s*([\"\\'])(?P<value>(?:(?!\\1).)+)\\1"),
								ϒwebpage,
								λ.NewStr("thumbnail"),
							), λ.KWArgs{
								{Name: "fatal", Value: λ.False},
								{Name: "group", Value: λ.NewStr("value")},
							})
						}
					}()
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
				λ.NewStr("_TEST"):         NovaEmbedIE__TEST,
				λ.NewStr("_VALID_URL"):    NovaEmbedIE__VALID_URL,
				λ.NewStr("_real_extract"): NovaEmbedIE__real_extract,
			})
		}())
		NovaIE = λ.Cal(λ.TypeType, λ.NewStr("NovaIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				NovaIE__VALID_URL λ.Object
			)
			NovaIE__VALID_URL = λ.NewStr("https?://(?:[^.]+\\.)?(?P<site>tv(?:noviny)?|tn|novaplus|vymena|fanda|krasna|doma|prask)\\.nova\\.cz/(?:[^/]+/)+(?P<id>[^/]+?)(?:\\.html|/|$)")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): NovaIE__VALID_URL,
			})
		}())
	})
}
