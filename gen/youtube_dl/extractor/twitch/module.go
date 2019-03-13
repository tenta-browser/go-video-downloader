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
 * twitch/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/twitch.py
 */

package twitch

import (
	Ωjson "github.com/tenta-browser/go-video-downloader/gen/json"
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError                 λ.Object
	InfoExtractor                  λ.Object
	TwitchAllVideosIE              λ.Object
	TwitchBaseIE                   λ.Object
	TwitchChapterIE                λ.Object
	TwitchClipsIE                  λ.Object
	TwitchHighlightsIE             λ.Object
	TwitchItemBaseIE               λ.Object
	TwitchPastBroadcastsIE         λ.Object
	TwitchPlaylistBaseIE           λ.Object
	TwitchProfileIE                λ.Object
	TwitchStreamIE                 λ.Object
	TwitchUploadsIE                λ.Object
	TwitchVideoIE                  λ.Object
	TwitchVideosBaseIE             λ.Object
	TwitchVodIE                    λ.Object
	ϒclean_html                    λ.Object
	ϒcompat_kwargs                 λ.Object
	ϒcompat_parse_qs               λ.Object
	ϒcompat_str                    λ.Object
	ϒcompat_urllib_parse_urlencode λ.Object
	ϒcompat_urllib_parse_urlparse  λ.Object
	ϒfloat_or_none                 λ.Object
	ϒint_or_none                   λ.Object
	ϒorderedSet                    λ.Object
	ϒparse_duration                λ.Object
	ϒparse_iso8601                 λ.Object
	ϒqualities                     λ.Object
	ϒtry_get                       λ.Object
	ϒunified_timestamp             λ.Object
	ϒupdate_url_query              λ.Object
	ϒurl_or_none                   λ.Object
	ϒurljoin                       λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_kwargs = Ωcompat.ϒcompat_kwargs
		ϒcompat_parse_qs = Ωcompat.ϒcompat_parse_qs
		ϒcompat_str = Ωcompat.ϒcompat_str
		ϒcompat_urllib_parse_urlencode = Ωcompat.ϒcompat_urllib_parse_urlencode
		ϒcompat_urllib_parse_urlparse = Ωcompat.ϒcompat_urllib_parse_urlparse
		ϒclean_html = Ωutils.ϒclean_html
		ExtractorError = Ωutils.ExtractorError
		ϒfloat_or_none = Ωutils.ϒfloat_or_none
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒorderedSet = Ωutils.ϒorderedSet
		ϒparse_duration = Ωutils.ϒparse_duration
		ϒparse_iso8601 = Ωutils.ϒparse_iso8601
		ϒqualities = Ωutils.ϒqualities
		ϒtry_get = Ωutils.ϒtry_get
		ϒunified_timestamp = Ωutils.ϒunified_timestamp
		ϒupdate_url_query = Ωutils.ϒupdate_url_query
		ϒurl_or_none = Ωutils.ϒurl_or_none
		ϒurljoin = Ωutils.ϒurljoin
		TwitchBaseIE = λ.Cal(λ.TypeType, λ.NewStr("TwitchBaseIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				TwitchBaseIE__API_BASE        λ.Object
				TwitchBaseIE__CLIENT_ID       λ.Object
				TwitchBaseIE__NETRC_MACHINE   λ.Object
				TwitchBaseIE__VALID_URL_BASE  λ.Object
				TwitchBaseIE__call_api        λ.Object
				TwitchBaseIE__handle_error    λ.Object
				TwitchBaseIE__login           λ.Object
				TwitchBaseIE__real_initialize λ.Object
			)
			TwitchBaseIE__VALID_URL_BASE = λ.NewStr("https?://(?:(?:www|go|m)\\.)?twitch\\.tv")
			TwitchBaseIE__API_BASE = λ.NewStr("https://api.twitch.tv")
			TwitchBaseIE__CLIENT_ID = λ.NewStr("kimne78kx3ncx6brgo4mv6wki5h1ko")
			TwitchBaseIE__NETRC_MACHINE = λ.NewStr("twitch")
			TwitchBaseIE__handle_error = λ.NewFunction("_handle_error",
				[]λ.Param{
					{Name: "self"},
					{Name: "response"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒerror    λ.Object
						ϒresponse = λargs[1]
						ϒself     = λargs[0]
					)
					if λ.IsTrue(λ.NewBool(!λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒresponse, λ.DictType)))) {
						return λ.None
					}
					ϒerror = λ.Cal(λ.GetAttr(ϒresponse, "get", nil), λ.NewStr("error"))
					if λ.IsTrue(ϒerror) {
						panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.Mod(λ.NewStr("%s returned error: %s - %s"), λ.NewTuple(
							λ.GetAttr(ϒself, "IE_NAME", nil),
							ϒerror,
							λ.Cal(λ.GetAttr(ϒresponse, "get", nil), λ.NewStr("message")),
						))), λ.KWArgs{
							{Name: "expected", Value: λ.True},
						})))
					}
					return λ.None
				})
			TwitchBaseIE__call_api = λ.NewFunction("_call_api",
				[]λ.Param{
					{Name: "self"},
					{Name: "path"},
					{Name: "item_id"},
				},
				0, true, true,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒargs     = λargs[3]
						ϒheaders  λ.Object
						ϒitem_id  = λargs[2]
						ϒkwargs   = λargs[4]
						ϒpath     = λargs[1]
						ϒresponse λ.Object
						ϒself     = λargs[0]
					)
					ϒheaders = λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒkwargs, "get", nil), λ.NewStr("headers"), λ.NewDictWithTable(map[λ.Object]λ.Object{})), "copy", nil))
					λ.SetItem(ϒheaders, λ.NewStr("Client-ID"), λ.GetAttr(ϒself, "_CLIENT_ID", nil))
					λ.SetItem(ϒkwargs, λ.NewStr("headers"), ϒheaders)
					ϒresponse = λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(λ.Unpack(
						λ.Mod(λ.NewStr("%s/%s"), λ.NewTuple(
							λ.GetAttr(ϒself, "_API_BASE", nil),
							ϒpath,
						)),
						ϒitem_id,
						λ.AsStarred(ϒargs),
					)...), λ.KWArgs{
						{Name: "", Value: λ.Cal(ϒcompat_kwargs, ϒkwargs)},
					})
					λ.Cal(λ.GetAttr(ϒself, "_handle_error", nil), ϒresponse)
					return ϒresponse
				})
			TwitchBaseIE__real_initialize = λ.NewFunction("_real_initialize",
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
			TwitchBaseIE__login = λ.NewFunction("_login",
				[]λ.Param{
					{Name: "self"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒfail          λ.Object
						ϒhandle        λ.Object
						ϒlogin_page    λ.Object
						ϒlogin_step    λ.Object
						ϒpassword      λ.Object
						ϒredirect_page λ.Object
						ϒself          = λargs[0]
						ϒtfa_token     λ.Object
						ϒusername      λ.Object
						τmp0           λ.Object
					)
					τmp0 = λ.Cal(λ.GetAttr(ϒself, "_get_login_info", nil))
					ϒusername = λ.GetItem(τmp0, λ.NewInt(0))
					ϒpassword = λ.GetItem(τmp0, λ.NewInt(1))
					if λ.IsTrue(λ.NewBool(ϒusername == λ.None)) {
						return λ.None
					}
					ϒfail = λ.NewFunction("fail",
						[]λ.Param{
							{Name: "message"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒmessage = λargs[0]
							)
							panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.Mod(λ.NewStr("Unable to login. Twitch said: %s"), ϒmessage)), λ.KWArgs{
								{Name: "expected", Value: λ.True},
							})))
							return λ.None
						})
					ϒlogin_step = λ.NewFunction("login_step",
						[]λ.Param{
							{Name: "page"},
							{Name: "urlh"},
							{Name: "note"},
							{Name: "data"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒdata         = λargs[3]
								ϒerror        λ.Object
								ϒform         λ.Object
								ϒheaders      λ.Object
								ϒnote         = λargs[2]
								ϒpage         = λargs[0]
								ϒpage_url     λ.Object
								ϒpost_url     λ.Object
								ϒredirect_url λ.Object
								ϒresponse     λ.Object
								ϒurlh         = λargs[1]
							)
							ϒform = λ.Cal(λ.GetAttr(ϒself, "_hidden_inputs", nil), ϒpage)
							λ.Cal(λ.GetAttr(ϒform, "update", nil), ϒdata)
							ϒpage_url = λ.Cal(λ.GetAttr(ϒurlh, "geturl", nil))
							ϒpost_url = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
								λ.NewStr("<form[^>]+action=([\"\\'])(?P<url>.+?)\\1"),
								ϒpage,
								λ.NewStr("post url"),
							), λ.KWArgs{
								{Name: "default", Value: λ.GetAttr(ϒself, "_LOGIN_POST_URL", nil)},
								{Name: "group", Value: λ.NewStr("url")},
							})
							ϒpost_url = λ.Cal(ϒurljoin, ϒpage_url, ϒpost_url)
							ϒheaders = λ.NewDictWithTable(map[λ.Object]λ.Object{
								λ.NewStr("Referer"):      ϒpage_url,
								λ.NewStr("Origin"):       ϒpage_url,
								λ.NewStr("Content-Type"): λ.NewStr("text/plain;charset=UTF-8"),
							})
							ϒresponse = λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
								ϒpost_url,
								λ.None,
								ϒnote,
							), λ.KWArgs{
								{Name: "data", Value: λ.Cal(λ.GetAttr(λ.Cal(Ωjson.ϒdumps, ϒform), "encode", nil))},
								{Name: "headers", Value: ϒheaders},
								{Name: "expected_status", Value: λ.NewInt(400)},
							})
							ϒerror = func() λ.Object {
								if λv := λ.Cal(λ.GetAttr(ϒresponse, "get", nil), λ.NewStr("error_description")); λ.IsTrue(λv) {
									return λv
								} else {
									return λ.Cal(λ.GetAttr(ϒresponse, "get", nil), λ.NewStr("error_code"))
								}
							}()
							if λ.IsTrue(ϒerror) {
								λ.Cal(ϒfail, ϒerror)
							}
							if λ.IsTrue(λ.NewBool(λ.Contains(λ.Cal(λ.GetAttr(ϒresponse, "get", nil), λ.NewStr("message"), λ.NewStr("")), λ.NewStr("Authenticated successfully")))) {
								return λ.NewTuple(
									λ.None,
									λ.None,
								)
							}
							ϒredirect_url = λ.Cal(ϒurljoin, ϒpost_url, func() λ.Object {
								if λv := λ.Cal(λ.GetAttr(ϒresponse, "get", nil), λ.NewStr("redirect")); λ.IsTrue(λv) {
									return λv
								} else {
									return λ.GetItem(ϒresponse, λ.NewStr("redirect_path"))
								}
							}())
							return λ.Call(λ.GetAttr(ϒself, "_download_webpage_handle", nil), λ.NewArgs(
								ϒredirect_url,
								λ.None,
								λ.NewStr("Downloading login redirect page"),
							), λ.KWArgs{
								{Name: "headers", Value: ϒheaders},
							})
						})
					τmp0 = λ.Cal(λ.GetAttr(ϒself, "_download_webpage_handle", nil), λ.GetAttr(ϒself, "_LOGIN_FORM_URL", nil), λ.None, λ.NewStr("Downloading login page"))
					ϒlogin_page = λ.GetItem(τmp0, λ.NewInt(0))
					ϒhandle = λ.GetItem(τmp0, λ.NewInt(1))
					if λ.IsTrue(λ.NewBool(λ.Contains(ϒlogin_page, λ.NewStr("blacklist_message")))) {
						λ.Cal(ϒfail, λ.Cal(ϒclean_html, ϒlogin_page))
					}
					τmp0 = λ.Cal(ϒlogin_step, ϒlogin_page, ϒhandle, λ.NewStr("Logging in"), λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("username"):  ϒusername,
						λ.NewStr("password"):  ϒpassword,
						λ.NewStr("client_id"): λ.GetAttr(ϒself, "_CLIENT_ID", nil),
					}))
					ϒredirect_page = λ.GetItem(τmp0, λ.NewInt(0))
					ϒhandle = λ.GetItem(τmp0, λ.NewInt(1))
					if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒredirect_page))) {
						return λ.None
					}
					if λ.IsTrue(λ.NewBool(λ.Cal(Ωre.ϒsearch, λ.NewStr("(?i)<form[^>]+id=\"two-factor-submit\""), ϒredirect_page) != λ.None)) {
						ϒtfa_token = λ.Cal(λ.GetAttr(ϒself, "_get_tfa_info", nil), λ.NewStr("two-factor authentication token"))
						λ.Cal(ϒlogin_step, ϒredirect_page, ϒhandle, λ.NewStr("Submitting TFA token"), λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("authy_token"):  ϒtfa_token,
							λ.NewStr("remember_2fa"): λ.NewStr("true"),
						}))
					}
					return λ.None
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_API_BASE"):        TwitchBaseIE__API_BASE,
				λ.NewStr("_CLIENT_ID"):       TwitchBaseIE__CLIENT_ID,
				λ.NewStr("_NETRC_MACHINE"):   TwitchBaseIE__NETRC_MACHINE,
				λ.NewStr("_VALID_URL_BASE"):  TwitchBaseIE__VALID_URL_BASE,
				λ.NewStr("_call_api"):        TwitchBaseIE__call_api,
				λ.NewStr("_handle_error"):    TwitchBaseIE__handle_error,
				λ.NewStr("_login"):           TwitchBaseIE__login,
				λ.NewStr("_real_initialize"): TwitchBaseIE__real_initialize,
			})
		}())
		TwitchItemBaseIE = λ.Cal(λ.TypeType, λ.NewStr("TwitchItemBaseIE"), λ.NewTuple(TwitchBaseIE), func() λ.Dict {

			return λ.NewDictWithTable(map[λ.Object]λ.Object{})
		}())
		TwitchVideoIE = λ.Cal(λ.TypeType, λ.NewStr("TwitchVideoIE"), λ.NewTuple(TwitchItemBaseIE), func() λ.Dict {
			var (
				TwitchVideoIE__VALID_URL λ.Object
			)
			TwitchVideoIE__VALID_URL = λ.Mod(λ.NewStr("%s/[^/]+/b/(?P<id>\\d+)"), λ.GetAttr(TwitchBaseIE, "_VALID_URL_BASE", nil))
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): TwitchVideoIE__VALID_URL,
			})
		}())
		TwitchChapterIE = λ.Cal(λ.TypeType, λ.NewStr("TwitchChapterIE"), λ.NewTuple(TwitchItemBaseIE), func() λ.Dict {
			var (
				TwitchChapterIE__VALID_URL λ.Object
			)
			TwitchChapterIE__VALID_URL = λ.Mod(λ.NewStr("%s/[^/]+/c/(?P<id>\\d+)"), λ.GetAttr(TwitchBaseIE, "_VALID_URL_BASE", nil))
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): TwitchChapterIE__VALID_URL,
			})
		}())
		TwitchVodIE = λ.Cal(λ.TypeType, λ.NewStr("TwitchVodIE"), λ.NewTuple(TwitchItemBaseIE), func() λ.Dict {
			var (
				TwitchVodIE__VALID_URL λ.Object
			)
			TwitchVodIE__VALID_URL = λ.NewStr("(?x)\n                    https?://\n                        (?:\n                            (?:(?:www|go|m)\\.)?twitch\\.tv/(?:[^/]+/v(?:ideo)?|videos)/|\n                            player\\.twitch\\.tv/\\?.*?\\bvideo=v\n                        )\n                        (?P<id>\\d+)\n                    ")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): TwitchVodIE__VALID_URL,
			})
		}())
		TwitchPlaylistBaseIE = λ.Cal(λ.TypeType, λ.NewStr("TwitchPlaylistBaseIE"), λ.NewTuple(TwitchBaseIE), func() λ.Dict {
			var (
				TwitchPlaylistBaseIE__PLAYLIST_PATH λ.Object
			)
			TwitchPlaylistBaseIE__PLAYLIST_PATH = λ.NewStr("kraken/channels/%s/videos/?offset=%d&limit=%d")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_PLAYLIST_PATH"): TwitchPlaylistBaseIE__PLAYLIST_PATH,
			})
		}())
		TwitchProfileIE = λ.Cal(λ.TypeType, λ.NewStr("TwitchProfileIE"), λ.NewTuple(TwitchPlaylistBaseIE), func() λ.Dict {
			var (
				TwitchProfileIE__VALID_URL λ.Object
			)
			TwitchProfileIE__VALID_URL = λ.Mod(λ.NewStr("%s/(?P<id>[^/]+)/profile/?(?:\\#.*)?$"), λ.GetAttr(TwitchBaseIE, "_VALID_URL_BASE", nil))
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): TwitchProfileIE__VALID_URL,
			})
		}())
		TwitchVideosBaseIE = λ.Cal(λ.TypeType, λ.NewStr("TwitchVideosBaseIE"), λ.NewTuple(TwitchPlaylistBaseIE), func() λ.Dict {
			var (
				TwitchVideosBaseIE__PLAYLIST_PATH         λ.Object
				TwitchVideosBaseIE__VALID_URL_VIDEOS_BASE λ.Object
			)
			TwitchVideosBaseIE__VALID_URL_VIDEOS_BASE = λ.Mod(λ.NewStr("%s/(?P<id>[^/]+)/videos"), λ.GetAttr(TwitchBaseIE, "_VALID_URL_BASE", nil))
			TwitchVideosBaseIE__PLAYLIST_PATH = λ.Add(λ.GetAttr(TwitchPlaylistBaseIE, "_PLAYLIST_PATH", nil), λ.NewStr("&broadcast_type="))
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_PLAYLIST_PATH"):         TwitchVideosBaseIE__PLAYLIST_PATH,
				λ.NewStr("_VALID_URL_VIDEOS_BASE"): TwitchVideosBaseIE__VALID_URL_VIDEOS_BASE,
			})
		}())
		TwitchAllVideosIE = λ.Cal(λ.TypeType, λ.NewStr("TwitchAllVideosIE"), λ.NewTuple(TwitchVideosBaseIE), func() λ.Dict {
			var (
				TwitchAllVideosIE__VALID_URL λ.Object
			)
			TwitchAllVideosIE__VALID_URL = λ.Mod(λ.NewStr("%s/all"), λ.GetAttr(TwitchVideosBaseIE, "_VALID_URL_VIDEOS_BASE", nil))
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): TwitchAllVideosIE__VALID_URL,
			})
		}())
		TwitchUploadsIE = λ.Cal(λ.TypeType, λ.NewStr("TwitchUploadsIE"), λ.NewTuple(TwitchVideosBaseIE), func() λ.Dict {
			var (
				TwitchUploadsIE__VALID_URL λ.Object
			)
			TwitchUploadsIE__VALID_URL = λ.Mod(λ.NewStr("%s/uploads"), λ.GetAttr(TwitchVideosBaseIE, "_VALID_URL_VIDEOS_BASE", nil))
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): TwitchUploadsIE__VALID_URL,
			})
		}())
		TwitchPastBroadcastsIE = λ.Cal(λ.TypeType, λ.NewStr("TwitchPastBroadcastsIE"), λ.NewTuple(TwitchVideosBaseIE), func() λ.Dict {
			var (
				TwitchPastBroadcastsIE__VALID_URL λ.Object
			)
			TwitchPastBroadcastsIE__VALID_URL = λ.Mod(λ.NewStr("%s/past-broadcasts"), λ.GetAttr(TwitchVideosBaseIE, "_VALID_URL_VIDEOS_BASE", nil))
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): TwitchPastBroadcastsIE__VALID_URL,
			})
		}())
		TwitchHighlightsIE = λ.Cal(λ.TypeType, λ.NewStr("TwitchHighlightsIE"), λ.NewTuple(TwitchVideosBaseIE), func() λ.Dict {
			var (
				TwitchHighlightsIE__VALID_URL λ.Object
			)
			TwitchHighlightsIE__VALID_URL = λ.Mod(λ.NewStr("%s/highlights"), λ.GetAttr(TwitchVideosBaseIE, "_VALID_URL_VIDEOS_BASE", nil))
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): TwitchHighlightsIE__VALID_URL,
			})
		}())
		TwitchStreamIE = λ.Cal(λ.TypeType, λ.NewStr("TwitchStreamIE"), λ.NewTuple(TwitchBaseIE), func() λ.Dict {
			var (
				TwitchStreamIE__VALID_URL λ.Object
				TwitchStreamIE_suitable   λ.Object
			)
			TwitchStreamIE__VALID_URL = λ.NewStr("(?x)\n                    https?://\n                        (?:\n                            (?:(?:www|go|m)\\.)?twitch\\.tv/|\n                            player\\.twitch\\.tv/\\?.*?\\bchannel=\n                        )\n                        (?P<id>[^/#?]+)\n                    ")
			TwitchStreamIE_suitable = λ.NewFunction("suitable",
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
						if λ.IsTrue(λ.Cal(λ.BuiltinAny, λ.Cal(λ.NewFunction("<generator>",
							nil,
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								return λ.NewGenerator(func(λgen λ.Generator) λ.Object {
									var (
										ϒie  λ.Object
										τmp0 λ.Object
										τmp1 λ.Object
									)
									τmp0 = λ.Cal(λ.BuiltinIter, λ.NewTuple(
										TwitchVideoIE,
										TwitchChapterIE,
										TwitchVodIE,
										TwitchProfileIE,
										TwitchAllVideosIE,
										TwitchUploadsIE,
										TwitchPastBroadcastsIE,
										TwitchHighlightsIE,
										TwitchClipsIE,
									))
									for {
										if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
											break
										}
										ϒie = τmp1
										λgen.Yield(λ.Cal(λ.GetAttr(ϒie, "suitable", nil), ϒurl))
									}
									return λ.None
								})
							})))) {
							return λ.False
						} else {
							return λ.Cal(λ.GetAttr(λ.Cal(λ.SuperType, TwitchStreamIE, ϒcls), "suitable", nil), ϒurl)
						}
					}()
				})
			TwitchStreamIE_suitable = λ.Cal(λ.ClassMethodType, TwitchStreamIE_suitable)
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): TwitchStreamIE__VALID_URL,
				λ.NewStr("suitable"):   TwitchStreamIE_suitable,
			})
		}())
		TwitchClipsIE = λ.Cal(λ.TypeType, λ.NewStr("TwitchClipsIE"), λ.NewTuple(TwitchBaseIE), func() λ.Dict {
			var (
				TwitchClipsIE_IE_NAME       λ.Object
				TwitchClipsIE__TESTS        λ.Object
				TwitchClipsIE__VALID_URL    λ.Object
				TwitchClipsIE__real_extract λ.Object
			)
			TwitchClipsIE_IE_NAME = λ.NewStr("twitch:clips")
			TwitchClipsIE__VALID_URL = λ.NewStr("https?://(?:clips\\.twitch\\.tv/(?:[^/]+/)*|(?:www\\.)?twitch\\.tv/[^/]+/clip/)(?P<id>[^/?#&]+)")
			TwitchClipsIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("https://clips.twitch.tv/FaintLightGullWholeWheat"),
					λ.NewStr("md5"): λ.NewStr("761769e1eafce0ffebfb4089cb3847cd"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("42850523"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("EA Play 2016 Live from the Novo Theatre"),
						λ.NewStr("thumbnail"):   λ.NewStr("re:^https?://.*\\.jpg"),
						λ.NewStr("timestamp"):   λ.NewInt(1465767393),
						λ.NewStr("upload_date"): λ.NewStr("20160612"),
						λ.NewStr("creator"):     λ.NewStr("EA"),
						λ.NewStr("uploader"):    λ.NewStr("stereotype_"),
						λ.NewStr("uploader_id"): λ.NewStr("43566419"),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://clips.twitch.tv/rflegendary/UninterestedBeeDAESuppy"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://www.twitch.tv/sergeynixon/clip/StormyThankfulSproutFutureMan"),
					λ.NewStr("only_matching"): λ.True,
				}),
			)
			TwitchClipsIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒclip            λ.Object
						ϒformats         λ.Object
						ϒinfo            λ.Object
						ϒoption          λ.Object
						ϒquality_key     λ.Object
						ϒself            = λargs[0]
						ϒsource          λ.Object
						ϒstatus          λ.Object
						ϒthumbnail_id    λ.Object
						ϒthumbnail_url   λ.Object
						ϒthumbnails      λ.Object
						ϒthumbnails_dict λ.Object
						ϒurl             = λargs[1]
						ϒvideo_id        λ.Object
						τmp0             λ.Object
						τmp1             λ.Object
						τmp2             λ.Object
					)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒstatus = λ.Cal(λ.GetAttr(ϒself, "_download_json", nil), λ.Mod(λ.NewStr("https://clips.twitch.tv/api/v2/clips/%s/status"), ϒvideo_id), ϒvideo_id)
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.GetItem(ϒstatus, λ.NewStr("quality_options")))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒoption = τmp1
						if λ.IsTrue(λ.NewBool(!λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒoption, λ.DictType)))) {
							continue
						}
						ϒsource = λ.Cal(ϒurl_or_none, λ.Cal(λ.GetAttr(ϒoption, "get", nil), λ.NewStr("source")))
						if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒsource))) {
							continue
						}
						λ.Cal(λ.GetAttr(ϒformats, "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("url"):       ϒsource,
							λ.NewStr("format_id"): λ.Cal(λ.GetAttr(ϒoption, "get", nil), λ.NewStr("quality")),
							λ.NewStr("height"):    λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒoption, "get", nil), λ.NewStr("quality"))),
							λ.NewStr("fps"):       λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒoption, "get", nil), λ.NewStr("frame_rate"))),
						}))
					}
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					ϒinfo = λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("formats"): ϒformats,
					})
					ϒclip = λ.Call(λ.GetAttr(ϒself, "_call_api", nil), λ.NewArgs(
						λ.Mod(λ.NewStr("kraken/clips/%s"), ϒvideo_id),
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
						{Name: "headers", Value: λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("Accept"): λ.NewStr("application/vnd.twitchtv.v5+json"),
						})},
					})
					if λ.IsTrue(ϒclip) {
						ϒquality_key = λ.Cal(ϒqualities, λ.NewTuple(
							λ.NewStr("tiny"),
							λ.NewStr("small"),
							λ.NewStr("medium"),
						))
						ϒthumbnails = λ.NewList()
						ϒthumbnails_dict = λ.Cal(λ.GetAttr(ϒclip, "get", nil), λ.NewStr("thumbnails"))
						if λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒthumbnails_dict, λ.DictType)) {
							τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(ϒthumbnails_dict, "items", nil)))
							for {
								if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
									break
								}
								τmp2 = τmp1
								ϒthumbnail_id = λ.GetItem(τmp2, λ.NewInt(0))
								ϒthumbnail_url = λ.GetItem(τmp2, λ.NewInt(1))
								λ.Cal(λ.GetAttr(ϒthumbnails, "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
									λ.NewStr("id"):         ϒthumbnail_id,
									λ.NewStr("url"):        ϒthumbnail_url,
									λ.NewStr("preference"): λ.Cal(ϒquality_key, ϒthumbnail_id),
								}))
							}
						}
						λ.Cal(λ.GetAttr(ϒinfo, "update", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("id"): func() λ.Object {
								if λv := λ.Cal(λ.GetAttr(ϒclip, "get", nil), λ.NewStr("tracking_id")); λ.IsTrue(λv) {
									return λv
								} else {
									return ϒvideo_id
								}
							}(),
							λ.NewStr("title"): func() λ.Object {
								if λv := λ.Cal(λ.GetAttr(ϒclip, "get", nil), λ.NewStr("title")); λ.IsTrue(λv) {
									return λv
								} else {
									return ϒvideo_id
								}
							}(),
							λ.NewStr("duration"):   λ.Cal(ϒfloat_or_none, λ.Cal(λ.GetAttr(ϒclip, "get", nil), λ.NewStr("duration"))),
							λ.NewStr("views"):      λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒclip, "get", nil), λ.NewStr("views"))),
							λ.NewStr("timestamp"):  λ.Cal(ϒunified_timestamp, λ.Cal(λ.GetAttr(ϒclip, "get", nil), λ.NewStr("created_at"))),
							λ.NewStr("thumbnails"): ϒthumbnails,
							λ.NewStr("creator"): λ.Cal(ϒtry_get, ϒclip, λ.NewFunction("<lambda>",
								[]λ.Param{
									{Name: "x"},
								},
								0, false, false,
								func(λargs []λ.Object) λ.Object {
									var (
										ϒx = λargs[0]
									)
									return λ.GetItem(λ.GetItem(ϒx, λ.NewStr("broadcaster")), λ.NewStr("display_name"))
								}), ϒcompat_str),
							λ.NewStr("uploader"): λ.Cal(ϒtry_get, ϒclip, λ.NewFunction("<lambda>",
								[]λ.Param{
									{Name: "x"},
								},
								0, false, false,
								func(λargs []λ.Object) λ.Object {
									var (
										ϒx = λargs[0]
									)
									return λ.GetItem(λ.GetItem(ϒx, λ.NewStr("curator")), λ.NewStr("display_name"))
								}), ϒcompat_str),
							λ.NewStr("uploader_id"): λ.Cal(ϒtry_get, ϒclip, λ.NewFunction("<lambda>",
								[]λ.Param{
									{Name: "x"},
								},
								0, false, false,
								func(λargs []λ.Object) λ.Object {
									var (
										ϒx = λargs[0]
									)
									return λ.GetItem(λ.GetItem(ϒx, λ.NewStr("curator")), λ.NewStr("id"))
								}), ϒcompat_str),
						}))
					} else {
						λ.Cal(λ.GetAttr(ϒinfo, "update", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("title"): ϒvideo_id,
							λ.NewStr("id"):    ϒvideo_id,
						}))
					}
					return ϒinfo
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("IE_NAME"):       TwitchClipsIE_IE_NAME,
				λ.NewStr("_TESTS"):        TwitchClipsIE__TESTS,
				λ.NewStr("_VALID_URL"):    TwitchClipsIE__VALID_URL,
				λ.NewStr("_real_extract"): TwitchClipsIE__real_extract,
			})
		}())
	})
}
