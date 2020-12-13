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
 * vlive/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/vlive.py
 */

package vlive

import (
	Ωjson "github.com/tenta-browser/go-video-downloader/gen/json"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωnaver "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/naver"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError      λ.Object
	NaverBaseIE         λ.Object
	VLiveBaseIE         λ.Object
	VLiveChannelIE      λ.Object
	VLiveIE             λ.Object
	VLivePostIE         λ.Object
	ϒcompat_HTTPError   λ.Object
	ϒcompat_str         λ.Object
	ϒint_or_none        λ.Object
	ϒmerge_dicts        λ.Object
	ϒstr_or_none        λ.Object
	ϒstrip_or_none      λ.Object
	ϒtry_get            λ.Object
	ϒurlencode_postdata λ.Object
)

func init() {
	λ.InitModule(func() {
		NaverBaseIE = Ωnaver.NaverBaseIE
		ϒcompat_HTTPError = Ωcompat.ϒcompat_HTTPError
		ϒcompat_str = Ωcompat.ϒcompat_str
		ExtractorError = Ωutils.ExtractorError
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒmerge_dicts = Ωutils.ϒmerge_dicts
		ϒstr_or_none = Ωutils.ϒstr_or_none
		ϒstrip_or_none = Ωutils.ϒstrip_or_none
		ϒtry_get = Ωutils.ϒtry_get
		ϒurlencode_postdata = Ωutils.ϒurlencode_postdata
		VLiveBaseIE = λ.Cal(λ.TypeType, λ.StrLiteral("VLiveBaseIE"), λ.NewTuple(NaverBaseIE), func() λ.Dict {
			var (
				VLiveBaseIE__APP_ID λ.Object
			)
			VLiveBaseIE__APP_ID = λ.StrLiteral("8c6cc7b45d2568fb668be6e05b6e5a3b")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_APP_ID": VLiveBaseIE__APP_ID,
			})
		}())
		VLiveIE = λ.Cal(λ.TypeType, λ.StrLiteral("VLiveIE"), λ.NewTuple(VLiveBaseIE), func() λ.Dict {
			var (
				VLiveIE_IE_NAME          λ.Object
				VLiveIE__NETRC_MACHINE   λ.Object
				VLiveIE__VALID_URL       λ.Object
				VLiveIE__call_api        λ.Object
				VLiveIE__login           λ.Object
				VLiveIE__real_extract    λ.Object
				VLiveIE__real_initialize λ.Object
			)
			VLiveIE_IE_NAME = λ.StrLiteral("vlive")
			VLiveIE__VALID_URL = λ.StrLiteral("https?://(?:(?:www|m)\\.)?vlive\\.tv/(?:video|embed)/(?P<id>[0-9]+)")
			VLiveIE__NETRC_MACHINE = λ.StrLiteral("vlive")
			VLiveIE__real_initialize = λ.NewFunction("_real_initialize",
				[]λ.Param{
					{Name: "self"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒself = λargs[0]
					)
					λ.Calm(ϒself, "_login")
					return λ.None
				})
			VLiveIE__login = λ.NewFunction("_login",
				[]λ.Param{
					{Name: "self"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						LOGIN_URL     λ.Object
						ϒemail        λ.Object
						ϒis_logged_in λ.Object
						ϒpassword     λ.Object
						ϒself         = λargs[0]
						τmp0          λ.Object
					)
					τmp0 = λ.Calm(ϒself, "_get_login_info")
					ϒemail = λ.GetItem(τmp0, λ.IntLiteral(0))
					ϒpassword = λ.GetItem(τmp0, λ.IntLiteral(1))
					if λ.Contains(λ.NewTuple(
						ϒemail,
						ϒpassword,
					), λ.None) {
						return λ.None
					}
					ϒis_logged_in = λ.NewFunction("is_logged_in",
						nil,
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒlogin_info λ.Object
							)
							ϒlogin_info = λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
								λ.StrLiteral("https://www.vlive.tv/auth/loginInfo"),
								λ.None,
							), λ.KWArgs{
								{Name: "note", Value: λ.StrLiteral("Downloading login info")},
								{Name: "headers", Value: λ.DictLiteral(map[string]string{
									"Referer": "https://www.vlive.tv/home",
								})},
							})
							return func() λ.Object {
								if λv := λ.Cal(ϒtry_get, ϒlogin_info, λ.NewFunction("<lambda>",
									[]λ.Param{
										{Name: "x"},
									},
									0, false, false,
									func(λargs []λ.Object) λ.Object {
										var (
											ϒx = λargs[0]
										)
										return λ.GetItem(λ.GetItem(ϒx, λ.StrLiteral("message")), λ.StrLiteral("login"))
									}), λ.BoolType); λ.IsTrue(λv) {
									return λv
								} else {
									return λ.False
								}
							}()
						})
					LOGIN_URL = λ.StrLiteral("https://www.vlive.tv/auth/email/login")
					λ.Call(λ.GetAttr(ϒself, "_request_webpage", nil), λ.NewArgs(
						LOGIN_URL,
						λ.None,
					), λ.KWArgs{
						{Name: "note", Value: λ.StrLiteral("Downloading login cookies")},
					})
					λ.Call(λ.GetAttr(ϒself, "_download_webpage", nil), λ.NewArgs(
						LOGIN_URL,
						λ.None,
					), λ.KWArgs{
						{Name: "note", Value: λ.StrLiteral("Logging in")},
						{Name: "data", Value: λ.Cal(ϒurlencode_postdata, λ.DictLiteral(map[string]λ.Object{
							"email": ϒemail,
							"pwd":   ϒpassword,
						}))},
						{Name: "headers", Value: λ.DictLiteral(map[string]λ.Object{
							"Referer":      LOGIN_URL,
							"Content-Type": λ.StrLiteral("application/x-www-form-urlencoded"),
						})},
					})
					if !λ.IsTrue(λ.Cal(ϒis_logged_in)) {
						panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.StrLiteral("Unable to log in")), λ.KWArgs{
							{Name: "expected", Value: λ.True},
						})))
					}
					return λ.None
				})
			VLiveIE__call_api = λ.NewFunction("_call_api",
				[]λ.Param{
					{Name: "self"},
					{Name: "path_template"},
					{Name: "video_id"},
					{Name: "fields", Def: λ.None},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒfields        = λargs[3]
						ϒpath_template = λargs[1]
						ϒquery         λ.Object
						ϒself          = λargs[0]
						ϒvideo_id      = λargs[2]
						τmp0           λ.Object
						τmp1           λ.Object
					)
					ϒquery = λ.DictLiteral(map[string]λ.Object{
						"appId": λ.GetAttr(ϒself, "_APP_ID", nil),
						"gcc":   λ.StrLiteral("KR"),
					})
					if λ.IsTrue(ϒfields) {
						λ.SetItem(ϒquery, λ.StrLiteral("fields"), ϒfields)
					}
					τmp0, τmp1 = func() (λexit λ.Object, λret λ.Object) {
						defer λ.CatchMulti(
							nil,
							&λ.Catcher{ExtractorError, func(λex λ.BaseException) {
								var ϒe λ.Object = λex
								if λ.IsTrue(func() λ.Object {
									if λv := λ.Cal(λ.BuiltinIsInstance, λ.GetAttr(ϒe, "cause", nil), ϒcompat_HTTPError); !λ.IsTrue(λv) {
										return λv
									} else {
										return λ.Eq(λ.GetAttr(λ.GetAttr(ϒe, "cause", nil), "code", nil), λ.IntLiteral(403))
									}
								}()) {
									λ.Calm(ϒself, "raise_login_required", λ.GetItem(λ.Cal(Ωjson.ϒloads, λ.Calm(λ.Calm(λ.GetAttr(ϒe, "cause", nil), "read"), "decode")), λ.StrLiteral("message")))
								}
								panic(λ.Raise(λex))
							}},
						)
						λexit, λret = λ.BlockExitReturn, λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
							λ.Add(λ.StrLiteral("https://www.vlive.tv/globalv-web/vam-web/"), λ.Mod(ϒpath_template, ϒvideo_id)),
							ϒvideo_id,
							λ.Mod(λ.StrLiteral("Downloading %s JSON metadata"), λ.GetItem(λ.Calm(λ.GetItem(λ.Calm(ϒpath_template, "split", λ.StrLiteral("/")), λ.Neg(λ.IntLiteral(1))), "split", λ.StrLiteral("-")), λ.IntLiteral(0))),
						), λ.KWArgs{
							{Name: "headers", Value: λ.DictLiteral(map[string]string{
								"Referer": "https://www.vlive.tv/",
							})},
							{Name: "query", Value: ϒquery},
						})
						return
						return λ.BlockExitNormally, nil
					}()
					if τmp0 == λ.BlockExitReturn {
						return τmp1
					}
					return λ.None
				})
			VLiveIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒformats           λ.Object
						ϒget_common_fields λ.Object
						ϒinfo              λ.Object
						ϒinkey             λ.Object
						ϒpost              λ.Object
						ϒself              = λargs[0]
						ϒstatus            λ.Object
						ϒstream_url        λ.Object
						ϒurl               = λargs[1]
						ϒvideo             λ.Object
						ϒvideo_id          λ.Object
						ϒvideo_type        λ.Object
						ϒvod_id            λ.Object
					)
					ϒvideo_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒpost = λ.Calm(ϒself, "_call_api", λ.StrLiteral("post/v1.0/officialVideoPost-%s"), ϒvideo_id, λ.StrLiteral("author{nickname},channel{channelCode,channelName},officialVideo{commentCount,exposeStatus,likeCount,playCount,playTime,status,title,type,vodId}"))
					ϒvideo = λ.GetItem(ϒpost, λ.StrLiteral("officialVideo"))
					ϒget_common_fields = λ.NewFunction("get_common_fields",
						nil,
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒchannel λ.Object
							)
							ϒchannel = func() λ.Object {
								if λv := λ.Calm(ϒpost, "get", λ.StrLiteral("channel")); λ.IsTrue(λv) {
									return λv
								} else {
									return λ.DictLiteral(map[λ.Object]λ.Object{})
								}
							}()
							return λ.DictLiteral(map[string]λ.Object{
								"title":         λ.Calm(ϒvideo, "get", λ.StrLiteral("title")),
								"creator":       λ.Calm(λ.Calm(ϒpost, "get", λ.StrLiteral("author"), λ.DictLiteral(map[λ.Object]λ.Object{})), "get", λ.StrLiteral("nickname")),
								"channel":       λ.Calm(ϒchannel, "get", λ.StrLiteral("channelName")),
								"channel_id":    λ.Calm(ϒchannel, "get", λ.StrLiteral("channelCode")),
								"duration":      λ.Cal(ϒint_or_none, λ.Calm(ϒvideo, "get", λ.StrLiteral("playTime"))),
								"view_count":    λ.Cal(ϒint_or_none, λ.Calm(ϒvideo, "get", λ.StrLiteral("playCount"))),
								"like_count":    λ.Cal(ϒint_or_none, λ.Calm(ϒvideo, "get", λ.StrLiteral("likeCount"))),
								"comment_count": λ.Cal(ϒint_or_none, λ.Calm(ϒvideo, "get", λ.StrLiteral("commentCount"))),
							})
						})
					ϒvideo_type = λ.Calm(ϒvideo, "get", λ.StrLiteral("type"))
					if λ.IsTrue(λ.Eq(ϒvideo_type, λ.StrLiteral("VOD"))) {
						ϒinkey = λ.GetItem(λ.Calm(ϒself, "_call_api", λ.StrLiteral("video/v1.0/vod/%s/inkey"), ϒvideo_id), λ.StrLiteral("inkey"))
						ϒvod_id = λ.GetItem(ϒvideo, λ.StrLiteral("vodId"))
						return λ.Cal(ϒmerge_dicts, λ.Cal(ϒget_common_fields), λ.Calm(ϒself, "_extract_video_info", ϒvideo_id, ϒvod_id, ϒinkey))
					} else {
						if λ.IsTrue(λ.Eq(ϒvideo_type, λ.StrLiteral("LIVE"))) {
							ϒstatus = λ.Calm(ϒvideo, "get", λ.StrLiteral("status"))
							if λ.IsTrue(λ.Eq(ϒstatus, λ.StrLiteral("ON_AIR"))) {
								ϒstream_url = λ.GetItem(λ.GetItem(λ.Calm(ϒself, "_call_api", λ.StrLiteral("old/v3/live/%s/playInfo"), ϒvideo_id), λ.StrLiteral("result")), λ.StrLiteral("adaptiveStreamUrl"))
								ϒformats = λ.Calm(ϒself, "_extract_m3u8_formats", ϒstream_url, ϒvideo_id, λ.StrLiteral("mp4"))
								λ.Calm(ϒself, "_sort_formats", ϒformats)
								ϒinfo = λ.Cal(ϒget_common_fields)
								λ.Calm(ϒinfo, "update", λ.DictLiteral(map[string]λ.Object{
									"title":   λ.Calm(ϒself, "_live_title", λ.GetItem(ϒvideo, λ.StrLiteral("title"))),
									"id":      ϒvideo_id,
									"formats": ϒformats,
									"is_live": λ.True,
								}))
								return ϒinfo
							} else {
								if λ.IsTrue(λ.Eq(ϒstatus, λ.StrLiteral("ENDED"))) {
									panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.StrLiteral("Uploading for replay. Please wait...")), λ.KWArgs{
										{Name: "expected", Value: λ.True},
									})))
								} else {
									if λ.IsTrue(λ.Eq(ϒstatus, λ.StrLiteral("RESERVED"))) {
										panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.StrLiteral("Coming soon!")), λ.KWArgs{
											{Name: "expected", Value: λ.True},
										})))
									} else {
										if λ.IsTrue(λ.Eq(λ.Calm(ϒvideo, "get", λ.StrLiteral("exposeStatus")), λ.StrLiteral("CANCEL"))) {
											panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.StrLiteral("We are sorry, but the live broadcast has been canceled.")), λ.KWArgs{
												{Name: "expected", Value: λ.True},
											})))
										} else {
											panic(λ.Raise(λ.Cal(ExtractorError, λ.Add(λ.StrLiteral("Unknown status "), ϒstatus))))
										}
									}
								}
							}
						}
					}
					return λ.None
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"IE_NAME":          VLiveIE_IE_NAME,
				"_NETRC_MACHINE":   VLiveIE__NETRC_MACHINE,
				"_VALID_URL":       VLiveIE__VALID_URL,
				"_call_api":        VLiveIE__call_api,
				"_login":           VLiveIE__login,
				"_real_extract":    VLiveIE__real_extract,
				"_real_initialize": VLiveIE__real_initialize,
			})
		}())
		VLivePostIE = λ.Cal(λ.TypeType, λ.StrLiteral("VLivePostIE"), λ.NewTuple(VLiveIE), func() λ.Dict {
			var (
				VLivePostIE__FVIDEO_TMPL λ.Object
				VLivePostIE__VALID_URL   λ.Object
			)
			VLivePostIE__VALID_URL = λ.StrLiteral("https?://(?:(?:www|m)\\.)?vlive\\.tv/post/(?P<id>\\d-\\d+)")
			VLivePostIE__FVIDEO_TMPL = λ.StrLiteral("fvideo/v1.0/fvideo-%%s/%s")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_FVIDEO_TMPL": VLivePostIE__FVIDEO_TMPL,
				"_VALID_URL":   VLivePostIE__VALID_URL,
			})
		}())
		VLiveChannelIE = λ.Cal(λ.TypeType, λ.StrLiteral("VLiveChannelIE"), λ.NewTuple(VLiveBaseIE), func() λ.Dict {
			var (
				VLiveChannelIE__VALID_URL λ.Object
			)
			VLiveChannelIE__VALID_URL = λ.StrLiteral("https?://(?:channels\\.vlive\\.tv|(?:(?:www|m)\\.)?vlive\\.tv/channel)/(?P<id>[0-9A-Z]+)")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": VLiveChannelIE__VALID_URL,
			})
		}())
	})
}
