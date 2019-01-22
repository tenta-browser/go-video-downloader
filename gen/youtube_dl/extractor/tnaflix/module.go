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
 * tnaflix/module.go: transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/tnaflix.py
 */

package tnaflix

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	EMPFlixIE             λ.Object
	InfoExtractor         λ.Object
	MovieFapIE            λ.Object
	TNAEMPFlixBaseIE      λ.Object
	TNAFlixIE             λ.Object
	TNAFlixNetworkBaseIE  λ.Object
	TNAFlixNetworkEmbedIE λ.Object
	ϒcompat_str           λ.Object
	ϒfix_xml_ampersands   λ.Object
	ϒfloat_or_none        λ.Object
	ϒint_or_none          λ.Object
	ϒparse_duration       λ.Object
	ϒstr_to_int           λ.Object
	ϒunescapeHTML         λ.Object
	ϒxpath_text           λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_str = Ωcompat.ϒcompat_str
		ϒfix_xml_ampersands = Ωutils.ϒfix_xml_ampersands
		ϒfloat_or_none = Ωutils.ϒfloat_or_none
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒparse_duration = Ωutils.ϒparse_duration
		ϒstr_to_int = Ωutils.ϒstr_to_int
		ϒunescapeHTML = Ωutils.ϒunescapeHTML
		ϒxpath_text = Ωutils.ϒxpath_text
		TNAFlixNetworkBaseIE = λ.Cal(λ.TypeType, λ.NewStr("TNAFlixNetworkBaseIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				TNAFlixNetworkBaseIE__AVERAGE_RATING_REGEX λ.Object
				TNAFlixNetworkBaseIE__CATEGORIES_REGEX     λ.Object
				TNAFlixNetworkBaseIE__COMMENT_COUNT_REGEX  λ.Object
				TNAFlixNetworkBaseIE__CONFIG_REGEX         λ.Object
				TNAFlixNetworkBaseIE__DESCRIPTION_REGEX    λ.Object
				TNAFlixNetworkBaseIE__HOST                 λ.Object
				TNAFlixNetworkBaseIE__TITLE_REGEX          λ.Object
				TNAFlixNetworkBaseIE__UPLOADER_REGEX       λ.Object
				TNAFlixNetworkBaseIE__VIEW_COUNT_REGEX     λ.Object
				TNAFlixNetworkBaseIE__VKEY_SUFFIX          λ.Object
				TNAFlixNetworkBaseIE__extract_thumbnails   λ.Object
				TNAFlixNetworkBaseIE__real_extract         λ.Object
			)
			TNAFlixNetworkBaseIE__CONFIG_REGEX = λ.NewList(
				λ.NewStr("flashvars\\.config\\s*=\\s*escape\\(\"(?P<url>[^\"]+)\""),
				λ.NewStr("<input[^>]+name=\"config\\d?\" value=\"(?P<url>[^\"]+)\""),
				λ.NewStr("config\\s*=\\s*([\"\\'])(?P<url>(?:https?:)?//(?:(?!\\1).)+)\\1"),
			)
			TNAFlixNetworkBaseIE__HOST = λ.NewStr("tna")
			TNAFlixNetworkBaseIE__VKEY_SUFFIX = λ.NewStr("")
			TNAFlixNetworkBaseIE__TITLE_REGEX = λ.NewStr("<input[^>]+name=\"title\" value=\"([^\"]+)\"")
			TNAFlixNetworkBaseIE__DESCRIPTION_REGEX = λ.NewStr("<input[^>]+name=\"description\" value=\"([^\"]+)\"")
			TNAFlixNetworkBaseIE__UPLOADER_REGEX = λ.NewStr("<input[^>]+name=\"username\" value=\"([^\"]+)\"")
			TNAFlixNetworkBaseIE__VIEW_COUNT_REGEX = λ.None
			TNAFlixNetworkBaseIE__COMMENT_COUNT_REGEX = λ.None
			TNAFlixNetworkBaseIE__AVERAGE_RATING_REGEX = λ.None
			TNAFlixNetworkBaseIE__CATEGORIES_REGEX = λ.NewStr("<li[^>]*>\\s*<span[^>]+class=\"infoTitle\"[^>]*>Categories:</span>\\s*<span[^>]+class=\"listView\"[^>]*>(.+?)</span>\\s*</li>")
			TNAFlixNetworkBaseIE__extract_thumbnails = λ.NewFunction("_extract_thumbnails",
				[]λ.Param{
					{Name: "self"},
					{Name: "flix_xml"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒfirst      λ.Object
						ϒfirst_el   λ.Object
						ϒfirst_text λ.Object
						ϒflix_xml   = λargs[1]
						ϒget_child  λ.Object
						ϒheight     λ.Object
						ϒlast       λ.Object
						ϒlast_el    λ.Object
						ϒlast_text  λ.Object
						ϒpattern_el λ.Object
						ϒself       = λargs[0]
						ϒtimeline   λ.Object
						ϒwidth      λ.Object
					)
					ϒget_child = λ.NewFunction("get_child",
						[]λ.Param{
							{Name: "elem"},
							{Name: "names"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒchild λ.Object
								ϒelem  = λargs[0]
								ϒname  λ.Object
								ϒnames = λargs[1]
								τmp0   λ.Object
								τmp1   λ.Object
							)
							τmp0 = λ.Cal(λ.BuiltinIter, ϒnames)
							for {
								if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
									break
								}
								ϒname = τmp1
								ϒchild = λ.Cal(λ.GetAttr(ϒelem, "find", nil), ϒname)
								if λ.IsTrue(λ.NewBool(ϒchild != λ.None)) {
									return ϒchild
								}
							}
							return λ.None
						})
					ϒtimeline = λ.Cal(ϒget_child, ϒflix_xml, λ.NewList(
						λ.NewStr("timeline"),
						λ.NewStr("rolloverBarImage"),
					))
					if λ.IsTrue(λ.NewBool(ϒtimeline == λ.None)) {
						return λ.None
					}
					ϒpattern_el = λ.Cal(ϒget_child, ϒtimeline, λ.NewList(
						λ.NewStr("imagePattern"),
						λ.NewStr("pattern"),
					))
					if λ.IsTrue(func() λ.Object {
						if λv := λ.NewBool(ϒpattern_el == λ.None); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.NewBool(!λ.IsTrue(λ.GetAttr(ϒpattern_el, "text", nil)))
						}
					}()) {
						return λ.None
					}
					ϒfirst_el = λ.Cal(ϒget_child, ϒtimeline, λ.NewList(
						λ.NewStr("imageFirst"),
						λ.NewStr("first"),
					))
					ϒlast_el = λ.Cal(ϒget_child, ϒtimeline, λ.NewList(
						λ.NewStr("imageLast"),
						λ.NewStr("last"),
					))
					if λ.IsTrue(func() λ.Object {
						if λv := λ.NewBool(ϒfirst_el == λ.None); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.NewBool(ϒlast_el == λ.None)
						}
					}()) {
						return λ.None
					}
					ϒfirst_text = λ.GetAttr(ϒfirst_el, "text", nil)
					ϒlast_text = λ.GetAttr(ϒlast_el, "text", nil)
					if λ.IsTrue(func() λ.Object {
						if λv := λ.NewBool(!λ.IsTrue(λ.Cal(λ.GetAttr(ϒfirst_text, "isdigit", nil)))); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.NewBool(!λ.IsTrue(λ.Cal(λ.GetAttr(ϒlast_text, "isdigit", nil))))
						}
					}()) {
						return λ.None
					}
					ϒfirst = λ.Cal(λ.IntType, ϒfirst_text)
					ϒlast = λ.Cal(λ.IntType, ϒlast_text)
					if λ.IsTrue(λ.Gt(ϒfirst, ϒlast)) {
						return λ.None
					}
					ϒwidth = λ.Cal(ϒint_or_none, λ.Cal(ϒxpath_text, ϒtimeline, λ.NewStr("./imageWidth"), λ.NewStr("thumbnail width")))
					ϒheight = λ.Cal(ϒint_or_none, λ.Cal(ϒxpath_text, ϒtimeline, λ.NewStr("./imageHeight"), λ.NewStr("thumbnail height")))
					return λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
						nil,
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							return λ.NewGenerator(func(λgen λ.Generator) λ.Object {
								var (
									ϒi   λ.Object
									τmp0 λ.Object
									τmp1 λ.Object
								)
								τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.RangeType, ϒfirst, λ.Add(ϒlast, λ.NewInt(1))))
								for {
									if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
										break
									}
									ϒi = τmp1
									λgen.Yield(λ.NewDictWithTable(map[λ.Object]λ.Object{
										λ.NewStr("url"):    λ.Cal(λ.GetAttr(ϒself, "_proto_relative_url", nil), λ.Cal(λ.GetAttr(λ.GetAttr(ϒpattern_el, "text", nil), "replace", nil), λ.NewStr("#"), λ.Cal(ϒcompat_str, ϒi)), λ.NewStr("http:")),
										λ.NewStr("width"):  ϒwidth,
										λ.NewStr("height"): ϒheight,
									}))
								}
								return λ.None
							})
						})))
				})
			TNAFlixNetworkBaseIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒage_limit         λ.Object
						ϒaverage_rating    λ.Object
						ϒcategories        λ.Object
						ϒcategories_str    λ.Object
						ϒcfg_url           λ.Object
						ϒcfg_xml           λ.Object
						ϒcomment_count     λ.Object
						ϒdescription       λ.Object
						ϒdisplay_id        λ.Object
						ϒdisplay_id_key    λ.Object
						ϒduration          λ.Object
						ϒextract_field     λ.Object
						ϒextract_video_url λ.Object
						ϒformat_id         λ.Object
						ϒformats           λ.Object
						ϒheight            λ.Object
						ϒinputs            λ.Object
						ϒitem              λ.Object
						ϒmobj              λ.Object
						ϒres               λ.Object
						ϒself              = λargs[0]
						ϒthumbnail         λ.Object
						ϒthumbnails        λ.Object
						ϒtitle             λ.Object
						ϒuploader          λ.Object
						ϒurl               = λargs[1]
						ϒvideo_id          λ.Object
						ϒvideo_link        λ.Object
						ϒview_count        λ.Object
						ϒwebpage           λ.Object
						τmp0               λ.Object
						τmp1               λ.Object
					)
					ϒmobj = λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("id"))
					τmp0 = λ.Cal(λ.BuiltinIter, λ.NewTuple(
						λ.NewStr("display_id"),
						λ.NewStr("display_id_2"),
					))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒdisplay_id_key = τmp1
						if λ.IsTrue(λ.NewBool(λ.Contains(λ.Cal(λ.GetAttr(ϒmobj, "groupdict", nil)), ϒdisplay_id_key))) {
							ϒdisplay_id = λ.Cal(λ.GetAttr(ϒmobj, "group", nil), ϒdisplay_id_key)
							if λ.IsTrue(ϒdisplay_id) {
								break
							}
						}
					}
					if τmp1 == λ.AfterLast {
						ϒdisplay_id = ϒvideo_id
					}
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒdisplay_id)
					ϒcfg_url = λ.Cal(λ.GetAttr(ϒself, "_proto_relative_url", nil), λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
						λ.GetAttr(ϒself, "_CONFIG_REGEX", nil),
						ϒwebpage,
						λ.NewStr("flashvars.config"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
						{Name: "group", Value: λ.NewStr("url")},
					}), λ.NewStr("http:"))
					if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒcfg_url))) {
						ϒinputs = λ.Cal(λ.GetAttr(ϒself, "_hidden_inputs", nil), ϒwebpage)
						ϒcfg_url = λ.Mod(λ.NewStr("https://cdn-fck.%sflix.com/%sflix/%s%s.fid?key=%s&VID=%s&premium=1&vip=1&alpha"), λ.NewTuple(
							λ.GetAttr(ϒself, "_HOST", nil),
							λ.GetAttr(ϒself, "_HOST", nil),
							λ.GetItem(ϒinputs, λ.NewStr("vkey")),
							λ.GetAttr(ϒself, "_VKEY_SUFFIX", nil),
							λ.GetItem(ϒinputs, λ.NewStr("nkey")),
							ϒvideo_id,
						))
					}
					ϒcfg_xml = λ.Call(λ.GetAttr(ϒself, "_download_xml", nil), λ.NewArgs(
						ϒcfg_url,
						ϒdisplay_id,
						λ.NewStr("Downloading metadata"),
					), λ.KWArgs{
						{Name: "transform_source", Value: ϒfix_xml_ampersands},
						{Name: "headers", Value: λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("Referer"): ϒurl,
						})},
					})
					ϒformats = λ.NewList()
					ϒextract_video_url = λ.NewFunction("extract_video_url",
						[]λ.Param{
							{Name: "vl"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒvl = λargs[0]
							)
							return λ.Cal(ϒunescapeHTML, λ.GetAttr(ϒvl, "text", nil))
						})
					ϒvideo_link = λ.Cal(λ.GetAttr(ϒcfg_xml, "find", nil), λ.NewStr("./videoLink"))
					if λ.IsTrue(λ.NewBool(ϒvideo_link != λ.None)) {
						λ.Cal(λ.GetAttr(ϒformats, "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("url"): λ.Cal(ϒextract_video_url, ϒvideo_link),
							λ.NewStr("ext"): λ.Call(ϒxpath_text, λ.NewArgs(
								ϒcfg_xml,
								λ.NewStr("./videoConfig/type"),
								λ.NewStr("type"),
							), λ.KWArgs{
								{Name: "default", Value: λ.NewStr("flv")},
							}),
						}))
					}
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(ϒcfg_xml, "findall", nil), λ.NewStr("./quality/item")))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒitem = τmp1
						ϒvideo_link = λ.Cal(λ.GetAttr(ϒitem, "find", nil), λ.NewStr("./videoLink"))
						if λ.IsTrue(λ.NewBool(ϒvideo_link == λ.None)) {
							continue
						}
						ϒres = λ.Cal(λ.GetAttr(ϒitem, "find", nil), λ.NewStr("res"))
						ϒformat_id = func() λ.Object {
							if λ.IsTrue(λ.NewBool(ϒres == λ.None)) {
								return λ.None
							} else {
								return λ.GetAttr(ϒres, "text", nil)
							}
						}()
						ϒheight = λ.Cal(ϒint_or_none, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
							λ.NewStr("^(\\d+)[pP]"),
							ϒformat_id,
							λ.NewStr("height"),
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
						}))
						λ.Cal(λ.GetAttr(ϒformats, "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("url"):       λ.Cal(λ.GetAttr(ϒself, "_proto_relative_url", nil), λ.Cal(ϒextract_video_url, ϒvideo_link), λ.NewStr("http:")),
							λ.NewStr("format_id"): ϒformat_id,
							λ.NewStr("height"):    ϒheight,
						}))
					}
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					ϒthumbnail = λ.Cal(λ.GetAttr(ϒself, "_proto_relative_url", nil), λ.Cal(ϒxpath_text, ϒcfg_xml, λ.NewStr("./startThumb"), λ.NewStr("thumbnail")), λ.NewStr("http:"))
					ϒthumbnails = λ.Cal(λ.GetAttr(ϒself, "_extract_thumbnails", nil), ϒcfg_xml)
					ϒtitle = λ.None
					if λ.IsTrue(λ.GetAttr(ϒself, "_TITLE_REGEX", nil)) {
						ϒtitle = λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
							λ.GetAttr(ϒself, "_TITLE_REGEX", nil),
							ϒwebpage,
							λ.NewStr("title"),
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
						})
					}
					if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒtitle))) {
						ϒtitle = λ.Cal(λ.GetAttr(ϒself, "_og_search_title", nil), ϒwebpage)
					}
					ϒage_limit = func() λ.Object {
						if λv := λ.Cal(λ.GetAttr(ϒself, "_rta_search", nil), ϒwebpage); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.NewInt(18)
						}
					}()
					ϒduration = λ.Cal(ϒparse_duration, λ.Call(λ.GetAttr(ϒself, "_html_search_meta", nil), λ.NewArgs(
						λ.NewStr("duration"),
						ϒwebpage,
						λ.NewStr("duration"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					}))
					ϒextract_field = λ.NewFunction("extract_field",
						[]λ.Param{
							{Name: "pattern"},
							{Name: "name"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒname    = λargs[1]
								ϒpattern = λargs[0]
							)
							return func() λ.Object {
								if λ.IsTrue(ϒpattern) {
									return λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
										ϒpattern,
										ϒwebpage,
										ϒname,
									), λ.KWArgs{
										{Name: "default", Value: λ.None},
									})
								} else {
									return λ.None
								}
							}()
						})
					ϒdescription = λ.Cal(ϒextract_field, λ.GetAttr(ϒself, "_DESCRIPTION_REGEX", nil), λ.NewStr("description"))
					ϒuploader = λ.Cal(ϒextract_field, λ.GetAttr(ϒself, "_UPLOADER_REGEX", nil), λ.NewStr("uploader"))
					ϒview_count = λ.Cal(ϒstr_to_int, λ.Cal(ϒextract_field, λ.GetAttr(ϒself, "_VIEW_COUNT_REGEX", nil), λ.NewStr("view count")))
					ϒcomment_count = λ.Cal(ϒstr_to_int, λ.Cal(ϒextract_field, λ.GetAttr(ϒself, "_COMMENT_COUNT_REGEX", nil), λ.NewStr("comment count")))
					ϒaverage_rating = λ.Cal(ϒfloat_or_none, λ.Cal(ϒextract_field, λ.GetAttr(ϒself, "_AVERAGE_RATING_REGEX", nil), λ.NewStr("average rating")))
					ϒcategories_str = λ.Cal(ϒextract_field, λ.GetAttr(ϒself, "_CATEGORIES_REGEX", nil), λ.NewStr("categories"))
					ϒcategories = func() λ.Object {
						if λ.IsTrue(λ.NewBool(ϒcategories_str != λ.None)) {
							return λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
								nil,
								0, false, false,
								func(λargs []λ.Object) λ.Object {
									return λ.NewGenerator(func(λgen λ.Generator) λ.Object {
										var (
											ϒc   λ.Object
											τmp0 λ.Object
											τmp1 λ.Object
										)
										τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(ϒcategories_str, "split", nil), λ.NewStr(",")))
										for {
											if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
												break
											}
											ϒc = τmp1
											λgen.Yield(λ.Cal(λ.GetAttr(ϒc, "strip", nil)))
										}
										return λ.None
									})
								})))
						} else {
							return λ.NewList()
						}
					}()
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):             ϒvideo_id,
						λ.NewStr("display_id"):     ϒdisplay_id,
						λ.NewStr("title"):          ϒtitle,
						λ.NewStr("description"):    ϒdescription,
						λ.NewStr("thumbnail"):      ϒthumbnail,
						λ.NewStr("thumbnails"):     ϒthumbnails,
						λ.NewStr("duration"):       ϒduration,
						λ.NewStr("age_limit"):      ϒage_limit,
						λ.NewStr("uploader"):       ϒuploader,
						λ.NewStr("view_count"):     ϒview_count,
						λ.NewStr("comment_count"):  ϒcomment_count,
						λ.NewStr("average_rating"): ϒaverage_rating,
						λ.NewStr("categories"):     ϒcategories,
						λ.NewStr("formats"):        ϒformats,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_AVERAGE_RATING_REGEX"): TNAFlixNetworkBaseIE__AVERAGE_RATING_REGEX,
				λ.NewStr("_CATEGORIES_REGEX"):     TNAFlixNetworkBaseIE__CATEGORIES_REGEX,
				λ.NewStr("_COMMENT_COUNT_REGEX"):  TNAFlixNetworkBaseIE__COMMENT_COUNT_REGEX,
				λ.NewStr("_CONFIG_REGEX"):         TNAFlixNetworkBaseIE__CONFIG_REGEX,
				λ.NewStr("_DESCRIPTION_REGEX"):    TNAFlixNetworkBaseIE__DESCRIPTION_REGEX,
				λ.NewStr("_HOST"):                 TNAFlixNetworkBaseIE__HOST,
				λ.NewStr("_TITLE_REGEX"):          TNAFlixNetworkBaseIE__TITLE_REGEX,
				λ.NewStr("_UPLOADER_REGEX"):       TNAFlixNetworkBaseIE__UPLOADER_REGEX,
				λ.NewStr("_VIEW_COUNT_REGEX"):     TNAFlixNetworkBaseIE__VIEW_COUNT_REGEX,
				λ.NewStr("_VKEY_SUFFIX"):          TNAFlixNetworkBaseIE__VKEY_SUFFIX,
				λ.NewStr("_extract_thumbnails"):   TNAFlixNetworkBaseIE__extract_thumbnails,
				λ.NewStr("_real_extract"):         TNAFlixNetworkBaseIE__real_extract,
			})
		}())
		TNAFlixNetworkEmbedIE = λ.Cal(λ.TypeType, λ.NewStr("TNAFlixNetworkEmbedIE"), λ.NewTuple(TNAFlixNetworkBaseIE), func() λ.Dict {
			var (
				TNAFlixNetworkEmbedIE__TESTS       λ.Object
				TNAFlixNetworkEmbedIE__TITLE_REGEX λ.Object
				TNAFlixNetworkEmbedIE__VALID_URL   λ.Object
			)
			TNAFlixNetworkEmbedIE__VALID_URL = λ.NewStr("https?://player\\.(?:tna|emp)flix\\.com/video/(?P<id>\\d+)")
			TNAFlixNetworkEmbedIE__TITLE_REGEX = λ.NewStr("<title>([^<]+)</title>")
			TNAFlixNetworkEmbedIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("https://player.tnaflix.com/video/6538"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):         λ.NewStr("6538"),
						λ.NewStr("display_id"): λ.NewStr("6538"),
						λ.NewStr("ext"):        λ.NewStr("mp4"),
						λ.NewStr("title"):      λ.NewStr("Educational xxx video"),
						λ.NewStr("thumbnail"):  λ.NewStr("re:https?://.*\\.jpg$"),
						λ.NewStr("age_limit"):  λ.NewInt(18),
					}),
					λ.NewStr("params"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("skip_download"): λ.True,
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://player.empflix.com/video/33051"),
					λ.NewStr("only_matching"): λ.True,
				}),
			)
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TESTS"):       TNAFlixNetworkEmbedIE__TESTS,
				λ.NewStr("_TITLE_REGEX"): TNAFlixNetworkEmbedIE__TITLE_REGEX,
				λ.NewStr("_VALID_URL"):   TNAFlixNetworkEmbedIE__VALID_URL,
			})
		}())
		TNAEMPFlixBaseIE = λ.Cal(λ.TypeType, λ.NewStr("TNAEMPFlixBaseIE"), λ.NewTuple(TNAFlixNetworkBaseIE), func() λ.Dict {
			var (
				TNAEMPFlixBaseIE__CATEGORIES_REGEX  λ.Object
				TNAEMPFlixBaseIE__DESCRIPTION_REGEX λ.Object
				TNAEMPFlixBaseIE__UPLOADER_REGEX    λ.Object
			)
			TNAEMPFlixBaseIE__DESCRIPTION_REGEX = λ.NewStr("(?s)>Description:</[^>]+>(.+?)<")
			TNAEMPFlixBaseIE__UPLOADER_REGEX = λ.NewStr("<span>by\\s*<a[^>]+\\bhref=[\"\\']/profile/[^>]+>([^<]+)<")
			TNAEMPFlixBaseIE__CATEGORIES_REGEX = λ.NewStr("(?s)<span[^>]*>Categories:</span>(.+?)</div>")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_CATEGORIES_REGEX"):  TNAEMPFlixBaseIE__CATEGORIES_REGEX,
				λ.NewStr("_DESCRIPTION_REGEX"): TNAEMPFlixBaseIE__DESCRIPTION_REGEX,
				λ.NewStr("_UPLOADER_REGEX"):    TNAEMPFlixBaseIE__UPLOADER_REGEX,
			})
		}())
		TNAFlixIE = λ.Cal(λ.TypeType, λ.NewStr("TNAFlixIE"), λ.NewTuple(TNAEMPFlixBaseIE), func() λ.Dict {
			var (
				TNAFlixIE__TESTS       λ.Object
				TNAFlixIE__TITLE_REGEX λ.Object
				TNAFlixIE__VALID_URL   λ.Object
			)
			TNAFlixIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?tnaflix\\.com/[^/]+/(?P<display_id>[^/]+)/video(?P<id>\\d+)")
			TNAFlixIE__TITLE_REGEX = λ.NewStr("<title>(.+?) - (?:TNAFlix Porn Videos|TNAFlix\\.com)</title>")
			TNAFlixIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.tnaflix.com/porn-stars/Carmella-Decesare-striptease/video553878"),
					λ.NewStr("md5"): λ.NewStr("7e569419fe6d69543d01e6be22f5f7c4"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):         λ.NewStr("553878"),
						λ.NewStr("display_id"): λ.NewStr("Carmella-Decesare-striptease"),
						λ.NewStr("ext"):        λ.NewStr("mp4"),
						λ.NewStr("title"):      λ.NewStr("Carmella Decesare - striptease"),
						λ.NewStr("thumbnail"):  λ.NewStr("re:https?://.*\\.jpg$"),
						λ.NewStr("duration"):   λ.NewInt(91),
						λ.NewStr("age_limit"):  λ.NewInt(18),
						λ.NewStr("categories"): λ.NewList(λ.NewStr("Porn Stars")),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("https://www.tnaflix.com/teen-porn/Educational-xxx-video/video6538"),
					λ.NewStr("md5"): λ.NewStr("0f5d4d490dbfd117b8607054248a07c0"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("6538"),
						λ.NewStr("display_id"):  λ.NewStr("Educational-xxx-video"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Educational xxx video"),
						λ.NewStr("description"): λ.NewStr("md5:b4fab8f88a8621c8fabd361a173fe5b8"),
						λ.NewStr("thumbnail"):   λ.NewStr("re:https?://.*\\.jpg$"),
						λ.NewStr("duration"):    λ.NewInt(164),
						λ.NewStr("age_limit"):   λ.NewInt(18),
						λ.NewStr("uploader"):    λ.NewStr("bobwhite39"),
						λ.NewStr("categories"):  λ.ListType,
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://www.tnaflix.com/amateur-porn/bunzHD-Ms.Donk/video358632"),
					λ.NewStr("only_matching"): λ.True,
				}),
			)
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TESTS"):       TNAFlixIE__TESTS,
				λ.NewStr("_TITLE_REGEX"): TNAFlixIE__TITLE_REGEX,
				λ.NewStr("_VALID_URL"):   TNAFlixIE__VALID_URL,
			})
		}())
		EMPFlixIE = λ.Cal(λ.TypeType, λ.NewStr("EMPFlixIE"), λ.NewTuple(TNAEMPFlixBaseIE), func() λ.Dict {
			var (
				EMPFlixIE__HOST        λ.Object
				EMPFlixIE__TESTS       λ.Object
				EMPFlixIE__VALID_URL   λ.Object
				EMPFlixIE__VKEY_SUFFIX λ.Object
			)
			EMPFlixIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?empflix\\.com/(?:videos/(?P<display_id>.+?)-|[^/]+/(?P<display_id_2>[^/]+)/video)(?P<id>[0-9]+)")
			EMPFlixIE__HOST = λ.NewStr("emp")
			EMPFlixIE__VKEY_SUFFIX = λ.NewStr("-1")
			EMPFlixIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.empflix.com/videos/Amateur-Finger-Fuck-33051.html"),
					λ.NewStr("md5"): λ.NewStr("bc30d48b91a7179448a0bda465114676"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("33051"),
						λ.NewStr("display_id"):  λ.NewStr("Amateur-Finger-Fuck"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Amateur Finger Fuck"),
						λ.NewStr("description"): λ.NewStr("Amateur solo finger fucking."),
						λ.NewStr("thumbnail"):   λ.NewStr("re:https?://.*\\.jpg$"),
						λ.NewStr("duration"):    λ.NewInt(83),
						λ.NewStr("age_limit"):   λ.NewInt(18),
						λ.NewStr("uploader"):    λ.NewStr("cwbike"),
						λ.NewStr("categories"): λ.NewList(
							λ.NewStr("Amateur"),
							λ.NewStr("Anal"),
							λ.NewStr("Fisting"),
							λ.NewStr("Home made"),
							λ.NewStr("Solo"),
						),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://www.empflix.com/videos/[AROMA][ARMD-718]-Aoi-Yoshino-Sawa-25826.html"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://www.empflix.com/amateur-porn/Amateur-Finger-Fuck/video33051"),
					λ.NewStr("only_matching"): λ.True,
				}),
			)
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_HOST"):        EMPFlixIE__HOST,
				λ.NewStr("_TESTS"):       EMPFlixIE__TESTS,
				λ.NewStr("_VALID_URL"):   EMPFlixIE__VALID_URL,
				λ.NewStr("_VKEY_SUFFIX"): EMPFlixIE__VKEY_SUFFIX,
			})
		}())
		MovieFapIE = λ.Cal(λ.TypeType, λ.NewStr("MovieFapIE"), λ.NewTuple(TNAFlixNetworkBaseIE), func() λ.Dict {
			var (
				MovieFapIE__AVERAGE_RATING_REGEX λ.Object
				MovieFapIE__CATEGORIES_REGEX     λ.Object
				MovieFapIE__COMMENT_COUNT_REGEX  λ.Object
				MovieFapIE__TESTS                λ.Object
				MovieFapIE__VALID_URL            λ.Object
				MovieFapIE__VIEW_COUNT_REGEX     λ.Object
			)
			MovieFapIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?moviefap\\.com/videos/(?P<id>[0-9a-f]+)/(?P<display_id>[^/]+)\\.html")
			MovieFapIE__VIEW_COUNT_REGEX = λ.NewStr("<br>Views\\s*<strong>([\\d,.]+)</strong>")
			MovieFapIE__COMMENT_COUNT_REGEX = λ.NewStr("<span[^>]+id=\"comCount\"[^>]*>([\\d,.]+)</span>")
			MovieFapIE__AVERAGE_RATING_REGEX = λ.NewStr("Current Rating\\s*<br>\\s*<strong>([\\d.]+)</strong>")
			MovieFapIE__CATEGORIES_REGEX = λ.NewStr("(?s)<div[^>]+id=\"vid_info\"[^>]*>\\s*<div[^>]*>.+?</div>(.*?)<br>")
			MovieFapIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.moviefap.com/videos/be9867c9416c19f54a4a/experienced-milf-amazing-handjob.html"),
					λ.NewStr("md5"): λ.NewStr("26624b4e2523051b550067d547615906"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):             λ.NewStr("be9867c9416c19f54a4a"),
						λ.NewStr("display_id"):     λ.NewStr("experienced-milf-amazing-handjob"),
						λ.NewStr("ext"):            λ.NewStr("mp4"),
						λ.NewStr("title"):          λ.NewStr("Experienced MILF Amazing Handjob"),
						λ.NewStr("description"):    λ.NewStr("Experienced MILF giving an Amazing Handjob"),
						λ.NewStr("thumbnail"):      λ.NewStr("re:https?://.*\\.jpg$"),
						λ.NewStr("age_limit"):      λ.NewInt(18),
						λ.NewStr("uploader"):       λ.NewStr("darvinfred06"),
						λ.NewStr("view_count"):     λ.IntType,
						λ.NewStr("comment_count"):  λ.IntType,
						λ.NewStr("average_rating"): λ.FloatType,
						λ.NewStr("categories"): λ.NewList(
							λ.NewStr("Amateur"),
							λ.NewStr("Masturbation"),
							λ.NewStr("Mature"),
							λ.NewStr("Flashing"),
						),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.moviefap.com/videos/e5da0d3edce5404418f5/jeune-couple-russe.html"),
					λ.NewStr("md5"): λ.NewStr("fa56683e291fc80635907168a743c9ad"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):             λ.NewStr("e5da0d3edce5404418f5"),
						λ.NewStr("display_id"):     λ.NewStr("jeune-couple-russe"),
						λ.NewStr("ext"):            λ.NewStr("flv"),
						λ.NewStr("title"):          λ.NewStr("Jeune Couple Russe"),
						λ.NewStr("description"):    λ.NewStr("Amateur"),
						λ.NewStr("thumbnail"):      λ.NewStr("re:https?://.*\\.jpg$"),
						λ.NewStr("age_limit"):      λ.NewInt(18),
						λ.NewStr("uploader"):       λ.NewStr("whiskeyjar"),
						λ.NewStr("view_count"):     λ.IntType,
						λ.NewStr("comment_count"):  λ.IntType,
						λ.NewStr("average_rating"): λ.FloatType,
						λ.NewStr("categories"): λ.NewList(
							λ.NewStr("Amateur"),
							λ.NewStr("Teen"),
						),
					}),
				}),
			)
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_AVERAGE_RATING_REGEX"): MovieFapIE__AVERAGE_RATING_REGEX,
				λ.NewStr("_CATEGORIES_REGEX"):     MovieFapIE__CATEGORIES_REGEX,
				λ.NewStr("_COMMENT_COUNT_REGEX"):  MovieFapIE__COMMENT_COUNT_REGEX,
				λ.NewStr("_TESTS"):                MovieFapIE__TESTS,
				λ.NewStr("_VALID_URL"):            MovieFapIE__VALID_URL,
				λ.NewStr("_VIEW_COUNT_REGEX"):     MovieFapIE__VIEW_COUNT_REGEX,
			})
		}())
	})
}
