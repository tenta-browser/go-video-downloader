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
 * pornhub/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/pornhub.py
 */

package pornhub

import (
	Ωfunctools "github.com/tenta-browser/go-video-downloader/gen/functools"
	Ωoperator "github.com/tenta-browser/go-video-downloader/gen/operator"
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωrequest "github.com/tenta-browser/go-video-downloader/gen/urllib/request"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωopenload "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/openload"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError             λ.Object
	InfoExtractor              λ.Object
	NO_DEFAULT                 λ.Object
	PhantomJSwrapper           λ.Object
	PornHubBaseIE              λ.Object
	PornHubIE                  λ.Object
	PornHubPagedPlaylistBaseIE λ.Object
	PornHubPagedVideoListIE    λ.Object
	PornHubPlaylistBaseIE      λ.Object
	PornHubUserIE              λ.Object
	PornHubUserVideosUploadIE  λ.Object
	ϒcompat_HTTPError          λ.Object
	ϒcompat_str                λ.Object
	ϒdetermine_ext             λ.Object
	ϒint_or_none               λ.Object
	ϒmerge_dicts               λ.Object
	ϒremove_quotes             λ.Object
	ϒstr_to_int                λ.Object
	ϒupdate_url_query          λ.Object
	ϒurl_or_none               λ.Object
	ϒurlencode_postdata        λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_HTTPError = Ωcompat.ϒcompat_HTTPError
		ϒcompat_str = Ωcompat.ϒcompat_str
		PhantomJSwrapper = Ωopenload.PhantomJSwrapper
		ϒdetermine_ext = Ωutils.ϒdetermine_ext
		ExtractorError = Ωutils.ExtractorError
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒmerge_dicts = Ωutils.ϒmerge_dicts
		NO_DEFAULT = Ωutils.NO_DEFAULT
		ϒremove_quotes = Ωutils.ϒremove_quotes
		ϒstr_to_int = Ωutils.ϒstr_to_int
		ϒupdate_url_query = Ωutils.ϒupdate_url_query
		ϒurlencode_postdata = Ωutils.ϒurlencode_postdata
		ϒurl_or_none = Ωutils.ϒurl_or_none
		PornHubBaseIE = λ.Cal(λ.TypeType, λ.StrLiteral("PornHubBaseIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				PornHubBaseIE__download_webpage_handle λ.Object
				PornHubBaseIE__login                   λ.Object
				PornHubBaseIE__real_initialize         λ.Object
			)
			PornHubBaseIE__download_webpage_handle = λ.NewFunction("_download_webpage_handle",
				[]λ.Param{
					{Name: "self"},
				},
				0, true, true,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒargs           = λargs[1]
						ϒdl             λ.Object
						ϒkwargs         = λargs[2]
						ϒphantom        λ.Object
						ϒret            λ.Object
						ϒself           = λargs[0]
						ϒurl            λ.Object
						ϒurl_or_request λ.Object
						ϒurlh           λ.Object
						ϒwebpage        λ.Object
						τmp0            λ.Object
					)
					ϒdl = λ.NewFunction("dl",
						nil,
						0, true, true,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒargs   = λargs[0]
								ϒkwargs = λargs[1]
							)
							return λ.Call(λ.GetAttr(λ.Cal(λ.SuperType, PornHubBaseIE, ϒself), "_download_webpage_handle", nil), λ.NewArgs(λ.Unpack(λ.AsStarred(ϒargs))...), λ.KWArgs{
								{Name: "", Value: ϒkwargs},
							})
						})
					ϒret = λ.Call(ϒdl, λ.NewArgs(λ.Unpack(λ.AsStarred(ϒargs))...), λ.KWArgs{
						{Name: "", Value: ϒkwargs},
					})
					if !λ.IsTrue(ϒret) {
						return ϒret
					}
					τmp0 = ϒret
					ϒwebpage = λ.GetItem(τmp0, λ.IntLiteral(0))
					ϒurlh = λ.GetItem(τmp0, λ.IntLiteral(1))
					if λ.IsTrue(λ.Cal(λ.BuiltinAny, λ.Cal(λ.NewFunction("<generator>",
						nil,
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							return λ.NewGenerator(func(λgy λ.Yielder) λ.Object {
								var (
									ϒp   λ.Object
									τmp0 λ.Object
									τmp1 λ.Object
								)
								τmp0 = λ.Cal(λ.BuiltinIter, λ.NewTuple(
									λ.StrLiteral("<body\\b[^>]+\\bonload=[\"\\']go\\(\\)"),
									λ.StrLiteral("document\\.cookie\\s*=\\s*[\"\\']RNKEY="),
									λ.StrLiteral("document\\.location\\.reload\\(true\\)"),
								))
								for {
									if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
										break
									}
									ϒp = τmp1
									λgy.Yield(λ.Cal(Ωre.ϒsearch, ϒp, ϒwebpage))
								}
								return λ.None
							})
						})))) {
						ϒurl_or_request = λ.GetItem(ϒargs, λ.IntLiteral(0))
						ϒurl = func() λ.Object {
							if λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒurl_or_request, Ωrequest.Request)) {
								return λ.Calm(ϒurl_or_request, "get_full_url")
							} else {
								return ϒurl_or_request
							}
						}()
						ϒphantom = λ.Call(PhantomJSwrapper, λ.NewArgs(ϒself), λ.KWArgs{
							{Name: "required_version", Value: λ.StrLiteral("2.0")},
						})
						λ.Call(λ.GetAttr(ϒphantom, "get", nil), λ.NewArgs(ϒurl), λ.KWArgs{
							{Name: "html", Value: ϒwebpage},
						})
						τmp0 = λ.Call(ϒdl, λ.NewArgs(λ.Unpack(λ.AsStarred(ϒargs))...), λ.KWArgs{
							{Name: "", Value: ϒkwargs},
						})
						ϒwebpage = λ.GetItem(τmp0, λ.IntLiteral(0))
						ϒurlh = λ.GetItem(τmp0, λ.IntLiteral(1))
					}
					return λ.NewTuple(
						ϒwebpage,
						ϒurlh,
					)
				})
			PornHubBaseIE__real_initialize = λ.NewFunction("_real_initialize",
				[]λ.Param{
					{Name: "self"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒself = λargs[0]
					)
					λ.SetAttr(ϒself, "_logged_in", λ.False)
					return λ.None
				})
			PornHubBaseIE__login = λ.NewFunction("_login",
				[]λ.Param{
					{Name: "self"},
					{Name: "host"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒhost       = λargs[1]
						ϒis_logged  λ.Object
						ϒlogin_form λ.Object
						ϒlogin_page λ.Object
						ϒlogin_url  λ.Object
						ϒmessage    λ.Object
						ϒpassword   λ.Object
						ϒresponse   λ.Object
						ϒself       = λargs[0]
						ϒsite       λ.Object
						ϒusername   λ.Object
						τmp0        λ.Object
					)
					if λ.IsTrue(λ.GetAttr(ϒself, "_logged_in", nil)) {
						return λ.None
					}
					ϒsite = λ.GetItem(λ.Calm(ϒhost, "split", λ.StrLiteral(".")), λ.IntLiteral(0))
					τmp0 = λ.Call(λ.GetAttr(ϒself, "_get_login_info", nil), nil, λ.KWArgs{
						{Name: "netrc_machine", Value: ϒsite},
					})
					ϒusername = λ.GetItem(τmp0, λ.IntLiteral(0))
					ϒpassword = λ.GetItem(τmp0, λ.IntLiteral(1))
					if ϒusername == λ.None {
						return λ.None
					}
					ϒlogin_url = λ.Mod(λ.StrLiteral("https://www.%s/%slogin"), λ.NewTuple(
						ϒhost,
						func() λ.Object {
							if λ.Contains(ϒhost, λ.StrLiteral("premium")) {
								return λ.StrLiteral("premium/")
							} else {
								return λ.StrLiteral("")
							}
						}(),
					))
					ϒlogin_page = λ.Calm(ϒself, "_download_webpage", ϒlogin_url, λ.None, λ.Mod(λ.StrLiteral("Downloading %s login page"), ϒsite))
					ϒis_logged = λ.NewFunction("is_logged",
						[]λ.Param{
							{Name: "webpage"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒwebpage = λargs[0]
							)
							return λ.Cal(λ.BuiltinAny, λ.Cal(λ.NewFunction("<generator>",
								nil,
								0, false, false,
								func(λargs []λ.Object) λ.Object {
									return λ.NewGenerator(func(λgy λ.Yielder) λ.Object {
										var (
											ϒp   λ.Object
											τmp0 λ.Object
											τmp1 λ.Object
										)
										τmp0 = λ.Cal(λ.BuiltinIter, λ.NewTuple(
											λ.StrLiteral("class=[\"\\']signOut"),
											λ.StrLiteral(">Sign\\s+[Oo]ut\\s*<"),
										))
										for {
											if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
												break
											}
											ϒp = τmp1
											λgy.Yield(λ.Cal(Ωre.ϒsearch, ϒp, ϒwebpage))
										}
										return λ.None
									})
								})))
						})
					if λ.IsTrue(λ.Cal(ϒis_logged, ϒlogin_page)) {
						λ.SetAttr(ϒself, "_logged_in", λ.True)
						return λ.None
					}
					ϒlogin_form = λ.Calm(ϒself, "_hidden_inputs", ϒlogin_page)
					λ.Calm(ϒlogin_form, "update", λ.DictLiteral(map[string]λ.Object{
						"username": ϒusername,
						"password": ϒpassword,
					}))
					ϒresponse = λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
						λ.Mod(λ.StrLiteral("https://www.%s/front/authenticate"), ϒhost),
						λ.None,
						λ.Mod(λ.StrLiteral("Logging in to %s"), ϒsite),
					), λ.KWArgs{
						{Name: "data", Value: λ.Cal(ϒurlencode_postdata, ϒlogin_form)},
						{Name: "headers", Value: λ.DictLiteral(map[string]λ.Object{
							"Content-Type":     λ.StrLiteral("application/x-www-form-urlencoded; charset=UTF-8"),
							"Referer":          ϒlogin_url,
							"X-Requested-With": λ.StrLiteral("XMLHttpRequest"),
						})},
					})
					if λ.IsTrue(λ.Eq(λ.Calm(ϒresponse, "get", λ.StrLiteral("success")), λ.StrLiteral("1"))) {
						λ.SetAttr(ϒself, "_logged_in", λ.True)
						return λ.None
					}
					ϒmessage = λ.Calm(ϒresponse, "get", λ.StrLiteral("message"))
					if ϒmessage != λ.None {
						panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.Mod(λ.StrLiteral("Unable to login: %s"), ϒmessage)), λ.KWArgs{
							{Name: "expected", Value: λ.True},
						})))
					}
					panic(λ.Raise(λ.Cal(ExtractorError, λ.StrLiteral("Unable to log in"))))
					return λ.None
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_download_webpage_handle": PornHubBaseIE__download_webpage_handle,
				"_login":                   PornHubBaseIE__login,
				"_real_initialize":         PornHubBaseIE__real_initialize,
			})
		}())
		PornHubIE = λ.Cal(λ.TypeType, λ.StrLiteral("PornHubIE"), λ.NewTuple(PornHubBaseIE), func() λ.Dict {
			var (
				PornHubIE__VALID_URL     λ.Object
				PornHubIE__extract_count λ.Object
				PornHubIE__real_extract  λ.Object
			)
			PornHubIE__VALID_URL = λ.StrLiteral("(?x)\n                    https?://\n                        (?:\n                            (?:[^/]+\\.)?(?P<host>pornhub(?:premium)?\\.(?:com|net|org))/(?:(?:view_video\\.php|video/show)\\?viewkey=|embed/)|\n                            (?:www\\.)?thumbzilla\\.com/video/\n                        )\n                        (?P<id>[\\da-z]+)\n                    ")
			PornHubIE__extract_count = λ.NewFunction("_extract_count",
				[]λ.Param{
					{Name: "self"},
					{Name: "pattern"},
					{Name: "webpage"},
					{Name: "name"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒname    = λargs[3]
						ϒpattern = λargs[1]
						ϒself    = λargs[0]
						ϒwebpage = λargs[2]
					)
					return λ.Cal(ϒstr_to_int, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						ϒpattern,
						ϒwebpage,
						λ.Mod(λ.StrLiteral("%s count"), ϒname),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					}))
				})
			PornHubIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						FORMAT_PREFIXES      λ.Object
						ϒadd_video_url       λ.Object
						ϒcomment_count       λ.Object
						ϒdefinition          λ.Object
						ϒdislike_count       λ.Object
						ϒdl_webpage          λ.Object
						ϒduration            λ.Object
						ϒerror_msg           λ.Object
						ϒext                 λ.Object
						ϒextract_js_vars     λ.Object
						ϒextract_list        λ.Object
						ϒextract_vote_count  λ.Object
						ϒflashvars           λ.Object
						ϒformat_url          λ.Object
						ϒformats             λ.Object
						ϒheight              λ.Object
						ϒhost                λ.Object
						ϒinfo                λ.Object
						ϒjs_vars             λ.Object
						ϒkey                 λ.Object
						ϒlike_count          λ.Object
						ϒmedia_definitions   λ.Object
						ϒmobj                λ.Object
						ϒparse_quality_items λ.Object
						ϒself                = λargs[0]
						ϒsubtitle_url        λ.Object
						ϒsubtitles           λ.Object
						ϒtbr                 λ.Object
						ϒthumbnail           λ.Object
						ϒtitle               λ.Object
						ϒupload_date         λ.Object
						ϒurl                 = λargs[1]
						ϒvideo_id            λ.Object
						ϒvideo_uploader      λ.Object
						ϒvideo_url           λ.Object
						ϒvideo_urls          λ.Object
						ϒvideo_urls_set      λ.Object
						ϒview_count          λ.Object
						ϒwebpage             λ.Object
						τmp0                 λ.Object
						τmp1                 λ.Object
						τmp2                 λ.Object
					)
					ϒmobj = λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl)
					ϒhost = func() λ.Object {
						if λv := λ.Calm(ϒmobj, "group", λ.StrLiteral("host")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.StrLiteral("pornhub.com")
						}
					}()
					ϒvideo_id = λ.Calm(ϒmobj, "group", λ.StrLiteral("id"))
					λ.Calm(ϒself, "_login", ϒhost)
					λ.Calm(ϒself, "_set_cookie", ϒhost, λ.StrLiteral("age_verified"), λ.StrLiteral("1"))
					ϒdl_webpage = λ.NewFunction("dl_webpage",
						[]λ.Param{
							{Name: "platform"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒplatform = λargs[0]
							)
							λ.Calm(ϒself, "_set_cookie", ϒhost, λ.StrLiteral("platform"), ϒplatform)
							return λ.Calm(ϒself, "_download_webpage", λ.Mod(λ.StrLiteral("https://www.%s/view_video.php?viewkey=%s"), λ.NewTuple(
								ϒhost,
								ϒvideo_id,
							)), ϒvideo_id, λ.Mod(λ.StrLiteral("Downloading %s webpage"), ϒplatform))
						})
					ϒwebpage = λ.Cal(ϒdl_webpage, λ.StrLiteral("pc"))
					ϒerror_msg = λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("(?s)<div[^>]+class=([\"\\'])(?:(?!\\1).)*\\b(?:removed|userMessageSection)\\b(?:(?!\\1).)*\\1[^>]*>(?P<error>.+?)</div>"),
						ϒwebpage,
						λ.StrLiteral("error message"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
						{Name: "group", Value: λ.StrLiteral("error")},
					})
					if λ.IsTrue(ϒerror_msg) {
						ϒerror_msg = λ.Cal(Ωre.ϒsub, λ.StrLiteral("\\s+"), λ.StrLiteral(" "), ϒerror_msg)
						panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.Mod(λ.StrLiteral("PornHub said: %s"), ϒerror_msg)), λ.KWArgs{
							{Name: "expected", Value: λ.True},
							{Name: "video_id", Value: ϒvideo_id},
						})))
					}
					ϒtitle = func() λ.Object {
						if λv := λ.Call(λ.GetAttr(ϒself, "_html_search_meta", nil), λ.NewArgs(
							λ.StrLiteral("twitter:title"),
							ϒwebpage,
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
						}); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
								λ.NewTuple(
									λ.StrLiteral("(?s)<h1[^>]+class=[\"\\']title[\"\\'][^>]*>(?P<title>.+?)</h1>"),
									λ.StrLiteral("<div[^>]+data-video-title=([\"\\'])(?P<title>(?:(?!\\1).)+)\\1"),
									λ.StrLiteral("shareTitle[\"\\']\\s*[=:]\\s*([\"\\'])(?P<title>(?:(?!\\1).)+)\\1"),
								),
								ϒwebpage,
								λ.StrLiteral("title"),
							), λ.KWArgs{
								{Name: "group", Value: λ.StrLiteral("title")},
							})
						}
					}()
					ϒvideo_urls = λ.NewList()
					ϒvideo_urls_set = λ.Cal(λ.SetType)
					ϒsubtitles = λ.DictLiteral(map[λ.Object]λ.Object{})
					ϒflashvars = λ.Calm(ϒself, "_parse_json", λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("var\\s+flashvars_\\d+\\s*=\\s*({.+?});"),
						ϒwebpage,
						λ.StrLiteral("flashvars"),
					), λ.KWArgs{
						{Name: "default", Value: λ.StrLiteral("{}")},
					}), ϒvideo_id)
					if λ.IsTrue(ϒflashvars) {
						ϒsubtitle_url = λ.Cal(ϒurl_or_none, λ.Calm(ϒflashvars, "get", λ.StrLiteral("closedCaptionsFile")))
						if λ.IsTrue(ϒsubtitle_url) {
							λ.Calm(λ.Calm(ϒsubtitles, "setdefault", λ.StrLiteral("en"), λ.NewList()), "append", λ.DictLiteral(map[string]λ.Object{
								"url": ϒsubtitle_url,
								"ext": λ.StrLiteral("srt"),
							}))
						}
						ϒthumbnail = λ.Calm(ϒflashvars, "get", λ.StrLiteral("image_url"))
						ϒduration = λ.Cal(ϒint_or_none, λ.Calm(ϒflashvars, "get", λ.StrLiteral("video_duration")))
						ϒmedia_definitions = λ.Calm(ϒflashvars, "get", λ.StrLiteral("mediaDefinitions"))
						if λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒmedia_definitions, λ.ListType)) {
							τmp0 = λ.Cal(λ.BuiltinIter, ϒmedia_definitions)
							for {
								if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
									break
								}
								ϒdefinition = τmp1
								if !λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒdefinition, λ.DictType)) {
									continue
								}
								ϒvideo_url = λ.Calm(ϒdefinition, "get", λ.StrLiteral("videoUrl"))
								if λ.IsTrue(func() λ.Object {
									if λv := λ.NewBool(!λ.IsTrue(ϒvideo_url)); λ.IsTrue(λv) {
										return λv
									} else {
										return λ.NewBool(!λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒvideo_url, ϒcompat_str)))
									}
								}()) {
									continue
								}
								if λ.Contains(ϒvideo_urls_set, ϒvideo_url) {
									continue
								}
								λ.Calm(ϒvideo_urls_set, "add", ϒvideo_url)
								λ.Calm(ϒvideo_urls, "append", λ.NewTuple(
									ϒvideo_url,
									λ.Cal(ϒint_or_none, λ.Calm(ϒdefinition, "get", λ.StrLiteral("quality"))),
								))
							}
						}
					} else {
						τmp0 = λ.Mul(λ.NewList(λ.None), λ.IntLiteral(2))
						ϒthumbnail = λ.GetItem(τmp0, λ.IntLiteral(0))
						ϒduration = λ.GetItem(τmp0, λ.IntLiteral(1))
					}
					ϒextract_js_vars = λ.NewFunction("extract_js_vars",
						[]λ.Param{
							{Name: "webpage"},
							{Name: "pattern"},
							{Name: "default", Def: NO_DEFAULT},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒassignments    λ.Object
								ϒassn           λ.Object
								ϒdefault        = λargs[2]
								ϒjs_vars        λ.Object
								ϒparse_js_value λ.Object
								ϒpattern        = λargs[1]
								ϒvalue          λ.Object
								ϒvname          λ.Object
								ϒwebpage        = λargs[0]
								τmp0            λ.Object
								τmp1            λ.Object
								τmp2            λ.Object
							)
							ϒassignments = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
								ϒpattern,
								ϒwebpage,
								λ.StrLiteral("encoded url"),
							), λ.KWArgs{
								{Name: "default", Value: ϒdefault},
							})
							if !λ.IsTrue(ϒassignments) {
								return λ.DictLiteral(map[λ.Object]λ.Object{})
							}
							ϒassignments = λ.Calm(ϒassignments, "split", λ.StrLiteral(";"))
							ϒjs_vars = λ.DictLiteral(map[λ.Object]λ.Object{})
							ϒparse_js_value = λ.NewFunction("parse_js_value",
								[]λ.Param{
									{Name: "inp"},
								},
								0, false, false,
								func(λargs []λ.Object) λ.Object {
									var (
										ϒinp  = λargs[0]
										ϒinps λ.Object
									)
									ϒinp = λ.Cal(Ωre.ϒsub, λ.StrLiteral("/\\*(?:(?!\\*/).)*?\\*/"), λ.StrLiteral(""), ϒinp)
									if λ.Contains(ϒinp, λ.StrLiteral("+")) {
										ϒinps = λ.Calm(ϒinp, "split", λ.StrLiteral("+"))
										return λ.Cal(Ωfunctools.ϒreduce, Ωoperator.ϒconcat, λ.Cal(λ.MapIteratorType, ϒparse_js_value, ϒinps))
									}
									ϒinp = λ.Calm(ϒinp, "strip")
									if λ.Contains(ϒjs_vars, ϒinp) {
										return λ.GetItem(ϒjs_vars, ϒinp)
									}
									return λ.Cal(ϒremove_quotes, ϒinp)
								})
							τmp0 = λ.Cal(λ.BuiltinIter, ϒassignments)
							for {
								if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
									break
								}
								ϒassn = τmp1
								ϒassn = λ.Calm(ϒassn, "strip")
								if !λ.IsTrue(ϒassn) {
									continue
								}
								ϒassn = λ.Cal(Ωre.ϒsub, λ.StrLiteral("var\\s+"), λ.StrLiteral(""), ϒassn)
								τmp2 = λ.Calm(ϒassn, "split", λ.StrLiteral("="), λ.IntLiteral(1))
								ϒvname = λ.GetItem(τmp2, λ.IntLiteral(0))
								ϒvalue = λ.GetItem(τmp2, λ.IntLiteral(1))
								λ.SetItem(ϒjs_vars, ϒvname, λ.Cal(ϒparse_js_value, ϒvalue))
							}
							return ϒjs_vars
						})
					ϒadd_video_url = λ.NewFunction("add_video_url",
						[]λ.Param{
							{Name: "video_url"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒv_url     λ.Object
								ϒvideo_url = λargs[0]
							)
							ϒv_url = λ.Cal(ϒurl_or_none, ϒvideo_url)
							if !λ.IsTrue(ϒv_url) {
								return λ.None
							}
							if λ.Contains(ϒvideo_urls_set, ϒv_url) {
								return λ.None
							}
							λ.Calm(ϒvideo_urls, "append", λ.NewTuple(
								ϒv_url,
								λ.None,
							))
							λ.Calm(ϒvideo_urls_set, "add", ϒv_url)
							return λ.None
						})
					ϒparse_quality_items = λ.NewFunction("parse_quality_items",
						[]λ.Param{
							{Name: "quality_items"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒitem          λ.Object
								ϒq_items       λ.Object
								ϒquality_items = λargs[0]
								τmp0           λ.Object
								τmp1           λ.Object
							)
							ϒq_items = λ.Call(λ.GetAttr(ϒself, "_parse_json", nil), λ.NewArgs(
								ϒquality_items,
								ϒvideo_id,
							), λ.KWArgs{
								{Name: "fatal", Value: λ.False},
							})
							if !λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒq_items, λ.ListType)) {
								return λ.None
							}
							τmp0 = λ.Cal(λ.BuiltinIter, ϒq_items)
							for {
								if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
									break
								}
								ϒitem = τmp1
								if λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒitem, λ.DictType)) {
									λ.Cal(ϒadd_video_url, λ.Calm(ϒitem, "get", λ.StrLiteral("url")))
								}
							}
							return λ.None
						})
					if !λ.IsTrue(ϒvideo_urls) {
						FORMAT_PREFIXES = λ.NewTuple(
							λ.StrLiteral("media"),
							λ.StrLiteral("quality"),
							λ.StrLiteral("qualityItems"),
						)
						ϒjs_vars = λ.Call(ϒextract_js_vars, λ.NewArgs(
							ϒwebpage,
							λ.Mod(λ.StrLiteral("(var\\s+(?:%s)_.+)"), λ.Calm(λ.StrLiteral("|"), "join", FORMAT_PREFIXES)),
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
						})
						if λ.IsTrue(ϒjs_vars) {
							τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(ϒjs_vars, "items"))
							for {
								if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
									break
								}
								τmp2 = τmp1
								ϒkey = λ.GetItem(τmp2, λ.IntLiteral(0))
								ϒformat_url = λ.GetItem(τmp2, λ.IntLiteral(1))
								if λ.IsTrue(λ.Calm(ϒkey, "startswith", λ.GetItem(FORMAT_PREFIXES, λ.Neg(λ.IntLiteral(1))))) {
									λ.Cal(ϒparse_quality_items, ϒformat_url)
								} else {
									if λ.IsTrue(λ.Cal(λ.BuiltinAny, λ.Cal(λ.NewFunction("<generator>",
										nil,
										0, false, false,
										func(λargs []λ.Object) λ.Object {
											return λ.NewGenerator(func(λgy λ.Yielder) λ.Object {
												var (
													ϒp   λ.Object
													τmp0 λ.Object
													τmp1 λ.Object
												)
												τmp0 = λ.Cal(λ.BuiltinIter, λ.GetItem(FORMAT_PREFIXES, λ.NewSlice(λ.None, λ.IntLiteral(2), λ.None)))
												for {
													if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
														break
													}
													ϒp = τmp1
													λgy.Yield(λ.Calm(ϒkey, "startswith", ϒp))
												}
												return λ.None
											})
										})))) {
										λ.Cal(ϒadd_video_url, ϒformat_url)
									}
								}
							}
						}
						if λ.IsTrue(func() λ.Object {
							if λv := λ.NewBool(!λ.IsTrue(ϒvideo_urls)); !λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Cal(Ωre.ϒsearch, λ.StrLiteral("<[^>]+\\bid=[\"\\']lockedPlayer"), ϒwebpage)
							}
						}()) {
							panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.Mod(λ.StrLiteral("Video %s is locked"), ϒvideo_id)), λ.KWArgs{
								{Name: "expected", Value: λ.True},
							})))
						}
					}
					if !λ.IsTrue(ϒvideo_urls) {
						ϒjs_vars = λ.Cal(ϒextract_js_vars, λ.Cal(ϒdl_webpage, λ.StrLiteral("tv")), λ.StrLiteral("(var.+?mediastring.+?)</script>"))
						λ.Cal(ϒadd_video_url, λ.GetItem(ϒjs_vars, λ.StrLiteral("mediastring")))
					}
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(Ωre.ϒfinditer, λ.StrLiteral("<a[^>]+\\bclass=[\"\\']downloadBtn\\b[^>]+\\bhref=([\"\\'])(?P<url>(?:(?!\\1).)+)\\1"), ϒwebpage))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒmobj = τmp1
						ϒvideo_url = λ.Calm(ϒmobj, "group", λ.StrLiteral("url"))
						if !λ.Contains(ϒvideo_urls_set, ϒvideo_url) {
							λ.Calm(ϒvideo_urls, "append", λ.NewTuple(
								ϒvideo_url,
								λ.None,
							))
							λ.Calm(ϒvideo_urls_set, "add", ϒvideo_url)
						}
					}
					ϒupload_date = λ.None
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, ϒvideo_urls)
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						τmp2 = τmp1
						ϒvideo_url = λ.GetItem(τmp2, λ.IntLiteral(0))
						ϒheight = λ.GetItem(τmp2, λ.IntLiteral(1))
						if !λ.IsTrue(ϒupload_date) {
							ϒupload_date = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
								λ.StrLiteral("/(\\d{6}/\\d{2})/"),
								ϒvideo_url,
								λ.StrLiteral("upload data"),
							), λ.KWArgs{
								{Name: "default", Value: λ.None},
							})
							if λ.IsTrue(ϒupload_date) {
								ϒupload_date = λ.Calm(ϒupload_date, "replace", λ.StrLiteral("/"), λ.StrLiteral(""))
							}
						}
						ϒext = λ.Cal(ϒdetermine_ext, ϒvideo_url)
						if λ.IsTrue(λ.Eq(ϒext, λ.StrLiteral("mpd"))) {
							λ.Calm(ϒformats, "extend", λ.Call(λ.GetAttr(ϒself, "_extract_mpd_formats", nil), λ.NewArgs(
								ϒvideo_url,
								ϒvideo_id,
							), λ.KWArgs{
								{Name: "mpd_id", Value: λ.StrLiteral("dash")},
								{Name: "fatal", Value: λ.False},
							}))
							continue
						} else {
							if λ.IsTrue(λ.Eq(ϒext, λ.StrLiteral("m3u8"))) {
								λ.Calm(ϒformats, "extend", λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
									ϒvideo_url,
									ϒvideo_id,
									λ.StrLiteral("mp4"),
								), λ.KWArgs{
									{Name: "entry_protocol", Value: λ.StrLiteral("m3u8_native")},
									{Name: "m3u8_id", Value: λ.StrLiteral("hls")},
									{Name: "fatal", Value: λ.False},
								}))
								continue
							}
						}
						ϒtbr = λ.None
						ϒmobj = λ.Cal(Ωre.ϒsearch, λ.StrLiteral("(?P<height>\\d+)[pP]?_(?P<tbr>\\d+)[kK]"), ϒvideo_url)
						if λ.IsTrue(ϒmobj) {
							if !λ.IsTrue(ϒheight) {
								ϒheight = λ.Cal(λ.IntType, λ.Calm(ϒmobj, "group", λ.StrLiteral("height")))
							}
							ϒtbr = λ.Cal(λ.IntType, λ.Calm(ϒmobj, "group", λ.StrLiteral("tbr")))
						}
						λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
							"url": ϒvideo_url,
							"format_id": func() λ.Object {
								if λ.IsTrue(ϒheight) {
									return λ.Mod(λ.StrLiteral("%dp"), ϒheight)
								} else {
									return λ.None
								}
							}(),
							"height": ϒheight,
							"tbr":    ϒtbr,
						}))
					}
					λ.Calm(ϒself, "_sort_formats", ϒformats)
					ϒvideo_uploader = λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("(?s)From:&nbsp;.+?<(?:a\\b[^>]+\\bhref=[\"\\']/(?:(?:user|channel)s|model|pornstar)/|span\\b[^>]+\\bclass=[\"\\']username)[^>]+>(.+?)<"),
						ϒwebpage,
						λ.StrLiteral("uploader"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					ϒextract_vote_count = λ.NewFunction("extract_vote_count",
						[]λ.Param{
							{Name: "kind"},
							{Name: "name"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒkind = λargs[0]
								ϒname = λargs[1]
							)
							return λ.Calm(ϒself, "_extract_count", λ.NewTuple(
								λ.Mod(λ.StrLiteral("<span[^>]+\\bclass=\"votes%s\"[^>]*>([\\d,\\.]+)</span>"), ϒkind),
								λ.Mod(λ.StrLiteral("<span[^>]+\\bclass=[\"\\']votes%s[\"\\'][^>]*\\bdata-rating=[\"\\'](\\d+)"), ϒkind),
							), ϒwebpage, ϒname)
						})
					ϒview_count = λ.Calm(ϒself, "_extract_count", λ.StrLiteral("<span class=\"count\">([\\d,\\.]+)</span> [Vv]iews"), ϒwebpage, λ.StrLiteral("view"))
					ϒlike_count = λ.Cal(ϒextract_vote_count, λ.StrLiteral("Up"), λ.StrLiteral("like"))
					ϒdislike_count = λ.Cal(ϒextract_vote_count, λ.StrLiteral("Down"), λ.StrLiteral("dislike"))
					ϒcomment_count = λ.Calm(ϒself, "_extract_count", λ.StrLiteral("All Comments\\s*<span>\\(([\\d,.]+)\\)"), ϒwebpage, λ.StrLiteral("comment"))
					ϒextract_list = λ.NewFunction("extract_list",
						[]λ.Param{
							{Name: "meta_key"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒdiv      λ.Object
								ϒmeta_key = λargs[0]
							)
							ϒdiv = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
								λ.Mod(λ.StrLiteral("(?s)<div[^>]+\\bclass=[\"\\'].*?\\b%sWrapper[^>]*>(.+?)</div>"), ϒmeta_key),
								ϒwebpage,
								ϒmeta_key,
							), λ.KWArgs{
								{Name: "default", Value: λ.None},
							})
							if λ.IsTrue(ϒdiv) {
								return λ.Cal(Ωre.ϒfindall, λ.StrLiteral("<a[^>]+\\bhref=[^>]+>([^<]+)"), ϒdiv)
							}
							return λ.None
						})
					ϒinfo = λ.Call(λ.GetAttr(ϒself, "_search_json_ld", nil), λ.NewArgs(
						ϒwebpage,
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "default", Value: λ.DictLiteral(map[λ.Object]λ.Object{})},
					})
					λ.SetItem(ϒinfo, λ.StrLiteral("description"), λ.None)
					return λ.Cal(ϒmerge_dicts, λ.DictLiteral(map[string]λ.Object{
						"id":            ϒvideo_id,
						"uploader":      ϒvideo_uploader,
						"upload_date":   ϒupload_date,
						"title":         ϒtitle,
						"thumbnail":     ϒthumbnail,
						"duration":      ϒduration,
						"view_count":    ϒview_count,
						"like_count":    ϒlike_count,
						"dislike_count": ϒdislike_count,
						"comment_count": ϒcomment_count,
						"formats":       ϒformats,
						"age_limit":     λ.IntLiteral(18),
						"tags":          λ.Cal(ϒextract_list, λ.StrLiteral("tags")),
						"categories":    λ.Cal(ϒextract_list, λ.StrLiteral("categories")),
						"subtitles":     ϒsubtitles,
					}), ϒinfo)
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL":     PornHubIE__VALID_URL,
				"_extract_count": PornHubIE__extract_count,
				"_real_extract":  PornHubIE__real_extract,
			})
		}())
		PornHubPlaylistBaseIE = λ.Cal(λ.TypeType, λ.StrLiteral("PornHubPlaylistBaseIE"), λ.NewTuple(PornHubBaseIE), func() λ.Dict {

			return λ.ClassDictLiteral(map[λ.Object]λ.Object{})
		}())
		PornHubUserIE = λ.Cal(λ.TypeType, λ.StrLiteral("PornHubUserIE"), λ.NewTuple(PornHubPlaylistBaseIE), func() λ.Dict {
			var (
				PornHubUserIE__VALID_URL λ.Object
			)
			PornHubUserIE__VALID_URL = λ.StrLiteral("(?P<url>https?://(?:[^/]+\\.)?(?P<host>pornhub(?:premium)?\\.(?:com|net|org))/(?:(?:user|channel)s|model|pornstar)/(?P<id>[^/?#&]+))(?:[?#&]|/(?!videos)|$)")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": PornHubUserIE__VALID_URL,
			})
		}())
		PornHubPagedPlaylistBaseIE = λ.Cal(λ.TypeType, λ.StrLiteral("PornHubPagedPlaylistBaseIE"), λ.NewTuple(PornHubPlaylistBaseIE), func() λ.Dict {

			return λ.ClassDictLiteral(map[λ.Object]λ.Object{})
		}())
		PornHubPagedVideoListIE = λ.Cal(λ.TypeType, λ.StrLiteral("PornHubPagedVideoListIE"), λ.NewTuple(PornHubPagedPlaylistBaseIE), func() λ.Dict {
			var (
				PornHubPagedVideoListIE__VALID_URL λ.Object
				PornHubPagedVideoListIE_suitable   λ.Object
			)
			PornHubPagedVideoListIE__VALID_URL = λ.StrLiteral("https?://(?:[^/]+\\.)?(?P<host>pornhub(?:premium)?\\.(?:com|net|org))/(?P<id>(?:[^/]+/)*[^/?#&]+)")
			PornHubPagedVideoListIE_suitable = λ.NewFunction("suitable",
				[]λ.Param{
					{Name: "cls"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒcls = λargs[0]
						ϒurl = λargs[1]
					)
					return func() λ.Object {
						if λ.IsTrue(func() λ.Object {
							if λv := λ.Calm(PornHubIE, "suitable", ϒurl); λ.IsTrue(λv) {
								return λv
							} else if λv := λ.Calm(PornHubUserIE, "suitable", ϒurl); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Calm(PornHubUserVideosUploadIE, "suitable", ϒurl)
							}
						}()) {
							return λ.False
						} else {
							return λ.Calm(λ.Cal(λ.SuperType, PornHubPagedVideoListIE, ϒcls), "suitable", ϒurl)
						}
					}()
				})
			PornHubPagedVideoListIE_suitable = λ.Cal(λ.ClassMethodType, PornHubPagedVideoListIE_suitable)
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": PornHubPagedVideoListIE__VALID_URL,
				"suitable":   PornHubPagedVideoListIE_suitable,
			})
		}())
		PornHubUserVideosUploadIE = λ.Cal(λ.TypeType, λ.StrLiteral("PornHubUserVideosUploadIE"), λ.NewTuple(PornHubPagedPlaylistBaseIE), func() λ.Dict {
			var (
				PornHubUserVideosUploadIE__VALID_URL λ.Object
			)
			PornHubUserVideosUploadIE__VALID_URL = λ.StrLiteral("(?P<url>https?://(?:[^/]+\\.)?(?P<host>pornhub(?:premium)?\\.(?:com|net|org))/(?:(?:user|channel)s|model|pornstar)/(?P<id>[^/]+)/videos/upload)")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": PornHubUserVideosUploadIE__VALID_URL,
			})
		}())
		λ.GetAttr(PhantomJSwrapper, "__init__", nil)
		λ.GetAttr(PhantomJSwrapper, "get", nil)
	})
}
