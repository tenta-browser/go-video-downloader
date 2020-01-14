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
 * safari/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/safari.py
 */

package safari

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
	ExtractorError    λ.Object
	InfoExtractor     λ.Object
	SafariApiIE       λ.Object
	SafariBaseIE      λ.Object
	SafariCourseIE    λ.Object
	SafariIE          λ.Object
	ϒcompat_parse_qs  λ.Object
	ϒcompat_str       λ.Object
	ϒupdate_url_query λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_parse_qs = Ωcompat.ϒcompat_parse_qs
		ϒcompat_str = Ωcompat.ϒcompat_str
		ExtractorError = Ωutils.ExtractorError
		ϒupdate_url_query = Ωutils.ϒupdate_url_query
		SafariBaseIE = λ.Cal(λ.TypeType, λ.NewStr("SafariBaseIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				SafariBaseIE_LOGGED_IN        λ.Object
				SafariBaseIE__NETRC_MACHINE   λ.Object
				SafariBaseIE__login           λ.Object
				SafariBaseIE__real_initialize λ.Object
			)
			SafariBaseIE__NETRC_MACHINE = λ.NewStr("safari")
			SafariBaseIE_LOGGED_IN = λ.False
			SafariBaseIE__real_initialize = λ.NewFunction("_real_initialize",
				[]λ.Param{
					{Name: "self"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒself = λargs[0]
					)
					λ.Cal(λ.GetAttr(ϒself, "_login", nil))
					return λ.None
				})
			SafariBaseIE__login = λ.NewFunction("_login",
				[]λ.Param{
					{Name: "self"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒauth         λ.Object
						ϒcookie       λ.Object
						ϒcredentials  λ.Object
						ϒis_logged    λ.Object
						ϒnext_uri     λ.Object
						ϒparsed_url   λ.Object
						ϒpassword     λ.Object
						ϒqs           λ.Object
						ϒredirect_url λ.Object
						ϒself         = λargs[0]
						ϒurlh         λ.Object
						ϒusername     λ.Object
						τmp0          λ.Object
						τmp1          λ.Object
					)
					τmp0 = λ.Cal(λ.GetAttr(ϒself, "_get_login_info", nil))
					ϒusername = λ.GetItem(τmp0, λ.NewInt(0))
					ϒpassword = λ.GetItem(τmp0, λ.NewInt(1))
					if λ.IsTrue(λ.NewBool(ϒusername == λ.None)) {
						return λ.None
					}
					τmp0 = λ.Cal(λ.GetAttr(ϒself, "_download_webpage_handle", nil), λ.NewStr("https://learning.oreilly.com/accounts/login-check/"), λ.None, λ.NewStr("Downloading login page"))
					_ = λ.GetItem(τmp0, λ.NewInt(0))
					ϒurlh = λ.GetItem(τmp0, λ.NewInt(1))
					ϒis_logged = λ.NewFunction("is_logged",
						[]λ.Param{
							{Name: "urlh"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒurlh = λargs[0]
							)
							return λ.NewBool(λ.Contains(λ.Cal(ϒcompat_str, λ.Cal(λ.GetAttr(ϒurlh, "geturl", nil))), λ.NewStr("learning.oreilly.com/home/")))
						})
					if λ.IsTrue(λ.Cal(ϒis_logged, ϒurlh)) {
						λ.SetAttr(ϒself, "LOGGED_IN", λ.True)
						return λ.None
					}
					ϒredirect_url = λ.Cal(ϒcompat_str, λ.Cal(λ.GetAttr(ϒurlh, "geturl", nil)))
					ϒparsed_url = λ.Cal(Ωparse.ϒurlparse, ϒredirect_url)
					ϒqs = λ.Cal(ϒcompat_parse_qs, λ.GetAttr(ϒparsed_url, "query", nil))
					ϒnext_uri = λ.Cal(Ωparse.ϒurljoin, λ.NewStr("https://api.oreilly.com"), λ.GetItem(λ.GetItem(ϒqs, λ.NewStr("next")), λ.NewInt(0)))
					τmp0 = λ.Call(λ.GetAttr(ϒself, "_download_json_handle", nil), λ.NewArgs(
						λ.NewStr("https://www.oreilly.com/member/auth/login/"),
						λ.None,
						λ.NewStr("Logging in"),
					), λ.KWArgs{
						{Name: "data", Value: λ.Cal(λ.GetAttr(λ.Cal(Ωjson.ϒdumps, λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("email"):        ϒusername,
							λ.NewStr("password"):     ϒpassword,
							λ.NewStr("redirect_uri"): ϒnext_uri,
						})), "encode", nil))},
						{Name: "headers", Value: λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("Content-Type"): λ.NewStr("application/json"),
							λ.NewStr("Referer"):      ϒredirect_url,
						})},
						{Name: "expected_status", Value: λ.NewInt(400)},
					})
					ϒauth = λ.GetItem(τmp0, λ.NewInt(0))
					ϒurlh = λ.GetItem(τmp0, λ.NewInt(1))
					ϒcredentials = λ.Cal(λ.GetAttr(ϒauth, "get", nil), λ.NewStr("credentials"))
					if λ.IsTrue(func() λ.Object {
						if λv := λ.NewBool(!λ.IsTrue(λ.Cal(λ.GetAttr(ϒauth, "get", nil), λ.NewStr("logged_in")))); !λ.IsTrue(λv) {
							return λv
						} else if λv := λ.NewBool(!λ.IsTrue(λ.Cal(λ.GetAttr(ϒauth, "get", nil), λ.NewStr("redirect_uri")))); !λ.IsTrue(λv) {
							return λv
						} else {
							return ϒcredentials
						}
					}()) {
						panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.Mod(λ.NewStr("Unable to login: %s"), ϒcredentials)), λ.KWArgs{
							{Name: "expected", Value: λ.True},
						})))
					}
					τmp0 = λ.Cal(λ.BuiltinIter, λ.NewTuple(
						λ.NewStr("groot_sessionid"),
						λ.NewStr("orm-jwt"),
						λ.NewStr("orm-rt"),
					))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒcookie = τmp1
						λ.Cal(λ.GetAttr(ϒself, "_apply_first_set_cookie_header", nil), ϒurlh, ϒcookie)
					}
					τmp0 = λ.Cal(λ.GetAttr(ϒself, "_download_webpage_handle", nil), func() λ.Object {
						if λv := λ.Cal(λ.GetAttr(ϒauth, "get", nil), λ.NewStr("redirect_uri")); λ.IsTrue(λv) {
							return λv
						} else {
							return ϒnext_uri
						}
					}(), λ.None, λ.NewStr("Completing login"))
					_ = λ.GetItem(τmp0, λ.NewInt(0))
					ϒurlh = λ.GetItem(τmp0, λ.NewInt(1))
					if λ.IsTrue(λ.Cal(ϒis_logged, ϒurlh)) {
						λ.SetAttr(ϒself, "LOGGED_IN", λ.True)
						return λ.None
					}
					panic(λ.Raise(λ.Cal(ExtractorError, λ.NewStr("Unable to log in"))))
					return λ.None
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("LOGGED_IN"):        SafariBaseIE_LOGGED_IN,
				λ.NewStr("_NETRC_MACHINE"):   SafariBaseIE__NETRC_MACHINE,
				λ.NewStr("_login"):           SafariBaseIE__login,
				λ.NewStr("_real_initialize"): SafariBaseIE__real_initialize,
			})
		}())
		SafariIE = λ.Cal(λ.TypeType, λ.NewStr("SafariIE"), λ.NewTuple(SafariBaseIE), func() λ.Dict {
			var (
				SafariIE_IE_NAME       λ.Object
				SafariIE__PARTNER_ID   λ.Object
				SafariIE__UICONF_ID    λ.Object
				SafariIE__VALID_URL    λ.Object
				SafariIE__real_extract λ.Object
			)
			SafariIE_IE_NAME = λ.NewStr("safari")
			SafariIE__VALID_URL = λ.NewStr("(?x)\n                        https?://\n                            (?:www\\.)?(?:safaribooksonline|(?:learning\\.)?oreilly)\\.com/\n                            (?:\n                                library/view/[^/]+/(?P<course_id>[^/]+)/(?P<part>[^/?\\#&]+)\\.html|\n                                videos/[^/]+/[^/]+/(?P<reference_id>[^-]+-[^/?\\#&]+)\n                            )\n                    ")
			SafariIE__PARTNER_ID = λ.NewStr("1926081")
			SafariIE__UICONF_ID = λ.NewStr("29375172")
			SafariIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒkaltura_session λ.Object
						ϒmobj            λ.Object
						ϒpartner_id      λ.Object
						ϒquery           λ.Object
						ϒreference_id    λ.Object
						ϒself            = λargs[0]
						ϒsession         λ.Object
						ϒui_id           λ.Object
						ϒurl             = λargs[1]
						ϒurlh            λ.Object
						ϒvideo_id        λ.Object
						ϒwebpage         λ.Object
						τmp0             λ.Object
					)
					ϒmobj = λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl)
					ϒreference_id = λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("reference_id"))
					if λ.IsTrue(ϒreference_id) {
						ϒvideo_id = ϒreference_id
						ϒpartner_id = λ.GetAttr(ϒself, "_PARTNER_ID", nil)
						ϒui_id = λ.GetAttr(ϒself, "_UICONF_ID", nil)
					} else {
						ϒvideo_id = λ.Mod(λ.NewStr("%s-%s"), λ.NewTuple(
							λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("course_id")),
							λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("part")),
						))
						τmp0 = λ.Cal(λ.GetAttr(ϒself, "_download_webpage_handle", nil), ϒurl, ϒvideo_id)
						ϒwebpage = λ.GetItem(τmp0, λ.NewInt(0))
						ϒurlh = λ.GetItem(τmp0, λ.NewInt(1))
						ϒmobj = λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), λ.Cal(λ.GetAttr(ϒurlh, "geturl", nil)))
						ϒreference_id = λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("reference_id"))
						if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒreference_id))) {
							ϒreference_id = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
								λ.NewStr("data-reference-id=([\"\\'])(?P<id>(?:(?!\\1).)+)\\1"),
								ϒwebpage,
								λ.NewStr("kaltura reference id"),
							), λ.KWArgs{
								{Name: "group", Value: λ.NewStr("id")},
							})
						}
						ϒpartner_id = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
							λ.NewStr("data-partner-id=([\"\\'])(?P<id>(?:(?!\\1).)+)\\1"),
							ϒwebpage,
							λ.NewStr("kaltura widget id"),
						), λ.KWArgs{
							{Name: "default", Value: λ.GetAttr(ϒself, "_PARTNER_ID", nil)},
							{Name: "group", Value: λ.NewStr("id")},
						})
						ϒui_id = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
							λ.NewStr("data-ui-id=([\"\\'])(?P<id>(?:(?!\\1).)+)\\1"),
							ϒwebpage,
							λ.NewStr("kaltura uiconf id"),
						), λ.KWArgs{
							{Name: "default", Value: λ.GetAttr(ϒself, "_UICONF_ID", nil)},
							{Name: "group", Value: λ.NewStr("id")},
						})
					}
					ϒquery = λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("wid"):                    λ.Mod(λ.NewStr("_%s"), ϒpartner_id),
						λ.NewStr("uiconf_id"):              ϒui_id,
						λ.NewStr("flashvars[referenceId]"): ϒreference_id,
					})
					if λ.IsTrue(λ.GetAttr(ϒself, "LOGGED_IN", nil)) {
						ϒkaltura_session = λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
							λ.Mod(λ.NewStr("%s/player/kaltura_session/?reference_id=%s"), λ.NewTuple(
								λ.GetAttr(ϒself, "_API_BASE", nil),
								ϒreference_id,
							)),
							ϒvideo_id,
							λ.NewStr("Downloading kaltura session JSON"),
							λ.NewStr("Unable to download kaltura session JSON"),
						), λ.KWArgs{
							{Name: "fatal", Value: λ.False},
							{Name: "headers", Value: λ.NewDictWithTable(map[λ.Object]λ.Object{
								λ.NewStr("Accept"): λ.NewStr("application/json"),
							})},
						})
						if λ.IsTrue(ϒkaltura_session) {
							ϒsession = λ.Cal(λ.GetAttr(ϒkaltura_session, "get", nil), λ.NewStr("session"))
							if λ.IsTrue(ϒsession) {
								λ.SetItem(ϒquery, λ.NewStr("flashvars[ks]"), ϒsession)
							}
						}
					}
					return λ.Cal(λ.GetAttr(ϒself, "url_result", nil), λ.Cal(ϒupdate_url_query, λ.NewStr("https://cdnapisec.kaltura.com/html5/html5lib/v2.37.1/mwEmbedFrame.php"), ϒquery), λ.NewStr("Kaltura"))
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("IE_NAME"):       SafariIE_IE_NAME,
				λ.NewStr("_PARTNER_ID"):   SafariIE__PARTNER_ID,
				λ.NewStr("_UICONF_ID"):    SafariIE__UICONF_ID,
				λ.NewStr("_VALID_URL"):    SafariIE__VALID_URL,
				λ.NewStr("_real_extract"): SafariIE__real_extract,
			})
		}())
		SafariApiIE = λ.Cal(λ.TypeType, λ.NewStr("SafariApiIE"), λ.NewTuple(SafariBaseIE), func() λ.Dict {
			var (
				SafariApiIE__VALID_URL λ.Object
			)
			SafariApiIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?(?:safaribooksonline|(?:learning\\.)?oreilly)\\.com/api/v1/book/(?P<course_id>[^/]+)/chapter(?:-content)?/(?P<part>[^/?#&]+)\\.html")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): SafariApiIE__VALID_URL,
			})
		}())
		SafariCourseIE = λ.Cal(λ.TypeType, λ.NewStr("SafariCourseIE"), λ.NewTuple(SafariBaseIE), func() λ.Dict {
			var (
				SafariCourseIE__VALID_URL λ.Object
				SafariCourseIE_suitable   λ.Object
			)
			SafariCourseIE__VALID_URL = λ.NewStr("(?x)\n                    https?://\n                        (?:\n                            (?:www\\.)?(?:safaribooksonline|(?:learning\\.)?oreilly)\\.com/\n                            (?:\n                                library/view/[^/]+|\n                                api/v1/book|\n                                videos/[^/]+\n                            )|\n                            techbus\\.safaribooksonline\\.com\n                        )\n                        /(?P<id>[^/]+)\n                    ")
			SafariCourseIE_suitable = λ.NewFunction("suitable",
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
							if λv := λ.Cal(λ.GetAttr(SafariIE, "suitable", nil), ϒurl); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Cal(λ.GetAttr(SafariApiIE, "suitable", nil), ϒurl)
							}
						}()) {
							return λ.False
						} else {
							return λ.Cal(λ.GetAttr(λ.Cal(λ.SuperType, SafariCourseIE, ϒcls), "suitable", nil), ϒurl)
						}
					}()
				})
			SafariCourseIE_suitable = λ.Cal(λ.ClassMethodType, SafariCourseIE_suitable)
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): SafariCourseIE__VALID_URL,
				λ.NewStr("suitable"):   SafariCourseIE_suitable,
			})
		}())
	})
}
