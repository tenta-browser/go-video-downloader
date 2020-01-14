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
 * vice/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/vice.py
 */

package vice

import (
	Ωjson "github.com/tenta-browser/go-video-downloader/gen/json"
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωtime "github.com/tenta-browser/go-video-downloader/gen/time"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωadobepass "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/adobepass"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωyoutube "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/youtube"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	AdobePassIE      λ.Object
	ExtractorError   λ.Object
	InfoExtractor    λ.Object
	ViceArticleIE    λ.Object
	ViceBaseIE       λ.Object
	ViceIE           λ.Object
	ViceShowIE       λ.Object
	YoutubeIE        λ.Object
	ϒclean_html      λ.Object
	ϒcompat_str      λ.Object
	ϒint_or_none     λ.Object
	ϒparse_age_limit λ.Object
	ϒstr_or_none     λ.Object
	ϒtry_get         λ.Object
)

func init() {
	λ.InitModule(func() {
		AdobePassIE = Ωadobepass.AdobePassIE
		InfoExtractor = Ωcommon.InfoExtractor
		YoutubeIE = Ωyoutube.YoutubeIE
		ϒcompat_str = Ωcompat.ϒcompat_str
		ϒclean_html = Ωutils.ϒclean_html
		ExtractorError = Ωutils.ExtractorError
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒparse_age_limit = Ωutils.ϒparse_age_limit
		ϒstr_or_none = Ωutils.ϒstr_or_none
		ϒtry_get = Ωutils.ϒtry_get
		ViceBaseIE = λ.Cal(λ.TypeType, λ.NewStr("ViceBaseIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				ViceBaseIE__call_api λ.Object
			)
			ViceBaseIE__call_api = λ.NewFunction("_call_api",
				[]λ.Param{
					{Name: "self"},
					{Name: "resource"},
					{Name: "resource_key"},
					{Name: "resource_id"},
					{Name: "locale"},
					{Name: "fields"},
					{Name: "args", Def: λ.NewStr("")},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒargs         = λargs[6]
						ϒfields       = λargs[5]
						ϒlocale       = λargs[4]
						ϒresource     = λargs[1]
						ϒresource_id  = λargs[3]
						ϒresource_key = λargs[2]
						ϒself         = λargs[0]
					)
					return λ.GetItem(λ.GetItem(λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
						λ.NewStr("https://video.vice.com/api/v1/graphql"),
						ϒresource_id,
					), λ.KWArgs{
						{Name: "query", Value: λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("query"): λ.Mod(λ.NewStr("{\n  %s(locale: \"%s\", %s: \"%s\"%s) {\n    %s\n  }\n}"), λ.NewTuple(
								ϒresource,
								ϒlocale,
								ϒresource_key,
								ϒresource_id,
								ϒargs,
								ϒfields,
							)),
						})},
					}), λ.NewStr("data")), ϒresource)
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_call_api"): ViceBaseIE__call_api,
			})
		}())
		ViceIE = λ.Cal(λ.TypeType, λ.NewStr("ViceIE"), λ.NewTuple(
			ViceBaseIE,
			AdobePassIE,
		), func() λ.Dict {
			var (
				ViceIE_IE_NAME       λ.Object
				ViceIE__VALID_URL    λ.Object
				ViceIE__extract_url  λ.Object
				ViceIE__extract_urls λ.Object
				ViceIE__real_extract λ.Object
			)
			ViceIE_IE_NAME = λ.NewStr("vice")
			ViceIE__VALID_URL = λ.NewStr("https?://(?:(?:video|vms)\\.vice|(?:www\\.)?vice(?:land|tv))\\.com/(?P<locale>[^/]+)/(?:video/[^/]+|embed)/(?P<id>[\\da-f]{24})")
			ViceIE__extract_urls = λ.NewFunction("_extract_urls",
				[]λ.Param{
					{Name: "webpage"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒwebpage = λargs[0]
					)
					return λ.Cal(Ωre.ϒfindall, λ.NewStr("<iframe\\b[^>]+\\bsrc=[\"\\']((?:https?:)?//video\\.vice\\.com/[^/]+/embed/[\\da-f]{24})"), ϒwebpage)
				})
			ViceIE__extract_urls = λ.Cal(λ.StaticMethodType, ViceIE__extract_urls)
			ViceIE__extract_url = λ.NewFunction("_extract_url",
				[]λ.Param{
					{Name: "webpage"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒurls    λ.Object
						ϒwebpage = λargs[0]
					)
					ϒurls = λ.Cal(λ.GetAttr(ViceIE, "_extract_urls", nil), ϒwebpage)
					return func() λ.Object {
						if λ.IsTrue(ϒurls) {
							return λ.GetItem(ϒurls, λ.NewInt(0))
						} else {
							return λ.None
						}
					}()
				})
			ViceIE__extract_url = λ.Cal(λ.StaticMethodType, ViceIE__extract_url)
			ViceIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒcc_url        λ.Object
						ϒchannel       λ.Object
						ϒepisode       λ.Object
						ϒerror         λ.Object
						ϒerror_message λ.Object
						ϒexp           λ.Object
						ϒformats       λ.Object
						ϒlanguage_code λ.Object
						ϒlocale        λ.Object
						ϒpreplay       λ.Object
						ϒquery         λ.Object
						ϒrating        λ.Object
						ϒresource      λ.Object
						ϒseason        λ.Object
						ϒself          = λargs[0]
						ϒsubtitle      λ.Object
						ϒsubtitles     λ.Object
						ϒtitle         λ.Object
						ϒurl           = λargs[1]
						ϒvideo         λ.Object
						ϒvideo_data    λ.Object
						ϒvideo_id      λ.Object
						τmp0           λ.Object
						τmp1           λ.Object
					)
					τmp0 = λ.Cal(λ.GetAttr(λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl), "groups", nil))
					ϒlocale = λ.GetItem(τmp0, λ.NewInt(0))
					ϒvideo_id = λ.GetItem(τmp0, λ.NewInt(1))
					ϒvideo = λ.GetItem(λ.Cal(λ.GetAttr(ϒself, "_call_api", nil), λ.NewStr("videos"), λ.NewStr("id"), ϒvideo_id, ϒlocale, λ.NewStr("body\n    locked\n    rating\n    thumbnail_url\n    title")), λ.NewInt(0))
					ϒtitle = λ.Cal(λ.GetAttr(λ.GetItem(ϒvideo, λ.NewStr("title")), "strip", nil))
					ϒrating = λ.Cal(λ.GetAttr(ϒvideo, "get", nil), λ.NewStr("rating"))
					ϒquery = λ.NewDictWithTable(map[λ.Object]λ.Object{})
					if λ.IsTrue(λ.Cal(λ.GetAttr(ϒvideo, "get", nil), λ.NewStr("locked"))) {
						ϒresource = λ.Cal(λ.GetAttr(ϒself, "_get_mvpd_resource", nil), λ.NewStr("VICELAND"), ϒtitle, ϒvideo_id, ϒrating)
						λ.SetItem(ϒquery, λ.NewStr("tvetoken"), λ.Cal(λ.GetAttr(ϒself, "_extract_mvpd_auth", nil), ϒurl, ϒvideo_id, λ.NewStr("VICELAND"), ϒresource))
					}
					ϒexp = λ.Add(λ.Cal(λ.IntType, λ.Cal(Ωtime.ϒtime)), λ.NewInt(1440))
					λ.Cal(λ.GetAttr(ϒquery, "update", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("exp"): ϒexp,
						λ.NewStr("sign"): λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(λ.None, "sha512", nil), λ.Cal(λ.GetAttr(λ.Mod(λ.NewStr("%s:GET:%d"), λ.NewTuple(
							ϒvideo_id,
							ϒexp,
						)), "encode", nil))), "hexdigest", nil)),
						λ.NewStr("skipadstitching"): λ.NewInt(1),
						λ.NewStr("platform"):        λ.NewStr("desktop"),
						λ.NewStr("rn"):              λ.Cal(λ.None, λ.NewInt(10000), λ.NewInt(100000)),
					}))
					τmp0, τmp1 = func() (λexit λ.Object, λret λ.Object) {
						defer λ.CatchMulti(
							nil,
							&λ.Catcher{ExtractorError, func(λex λ.BaseException) {
								var ϒe λ.Object = λex
								if λ.IsTrue(func() λ.Object {
									if λv := λ.Cal(λ.BuiltinIsInstance, λ.GetAttr(ϒe, "cause", nil), λ.None); !λ.IsTrue(λv) {
										return λv
									} else {
										return λ.NewBool(λ.Contains(λ.NewTuple(
											λ.NewInt(400),
											λ.NewInt(401),
										), λ.GetAttr(λ.GetAttr(ϒe, "cause", nil), "code", nil)))
									}
								}()) {
									ϒerror = λ.Cal(Ωjson.ϒloads, λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(λ.GetAttr(ϒe, "cause", nil), "read", nil)), "decode", nil)))
									ϒerror_message = func() λ.Object {
										if λv := λ.Cal(λ.GetAttr(ϒerror, "get", nil), λ.NewStr("error_description")); λ.IsTrue(λv) {
											return λv
										} else {
											return λ.GetItem(ϒerror, λ.NewStr("details"))
										}
									}()
									panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.Mod(λ.NewStr("%s said: %s"), λ.NewTuple(
										λ.GetAttr(ϒself, "IE_NAME", nil),
										ϒerror_message,
									))), λ.KWArgs{
										{Name: "expected", Value: λ.True},
									})))
								}
								panic(λ.Raise(λex))
							}},
						)
						ϒpreplay = λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
							λ.Mod(λ.NewStr("https://vms.vice.com/%s/video/preplay/%s"), λ.NewTuple(
								ϒlocale,
								ϒvideo_id,
							)),
							ϒvideo_id,
						), λ.KWArgs{
							{Name: "query", Value: ϒquery},
						})
						return λ.BlockExitNormally, nil
					}()
					ϒvideo_data = λ.GetItem(ϒpreplay, λ.NewStr("video"))
					ϒformats = λ.Cal(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.GetItem(ϒpreplay, λ.NewStr("playURL")), ϒvideo_id, λ.NewStr("mp4"), λ.NewStr("m3u8_native"))
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					ϒepisode = func() λ.Object {
						if λv := λ.Cal(λ.GetAttr(ϒvideo_data, "get", nil), λ.NewStr("episode")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.NewDictWithTable(map[λ.Object]λ.Object{})
						}
					}()
					ϒchannel = func() λ.Object {
						if λv := λ.Cal(λ.GetAttr(ϒvideo_data, "get", nil), λ.NewStr("channel")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.NewDictWithTable(map[λ.Object]λ.Object{})
						}
					}()
					ϒseason = func() λ.Object {
						if λv := λ.Cal(λ.GetAttr(ϒvideo_data, "get", nil), λ.NewStr("season")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.NewDictWithTable(map[λ.Object]λ.Object{})
						}
					}()
					ϒsubtitles = λ.NewDictWithTable(map[λ.Object]λ.Object{})
					τmp1 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(ϒpreplay, "get", nil), λ.NewStr("subtitleURLs"), λ.NewList()))
					for {
						if τmp0 = λ.NextDefault(τmp1, λ.AfterLast); τmp0 == λ.AfterLast {
							break
						}
						ϒsubtitle = τmp0
						ϒcc_url = λ.Cal(λ.GetAttr(ϒsubtitle, "get", nil), λ.NewStr("url"))
						if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒcc_url))) {
							continue
						}
						ϒlanguage_code = func() λ.Object {
							if λv := λ.Cal(ϒtry_get, ϒsubtitle, λ.NewFunction("<lambda>",
								[]λ.Param{
									{Name: "x"},
								},
								0, false, false,
								func(λargs []λ.Object) λ.Object {
									var (
										ϒx = λargs[0]
									)
									return λ.GetItem(λ.GetItem(λ.GetItem(ϒx, λ.NewStr("languages")), λ.NewInt(0)), λ.NewStr("language_code"))
								}), ϒcompat_str); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.NewStr("en")
							}
						}()
						λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒsubtitles, "setdefault", nil), ϒlanguage_code, λ.NewList()), "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("url"): ϒcc_url,
						}))
					}
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("formats"):     ϒformats,
						λ.NewStr("id"):          ϒvideo_id,
						λ.NewStr("title"):       ϒtitle,
						λ.NewStr("description"): λ.Cal(ϒclean_html, λ.Cal(λ.GetAttr(ϒvideo, "get", nil), λ.NewStr("body"))),
						λ.NewStr("thumbnail"):   λ.Cal(λ.GetAttr(ϒvideo, "get", nil), λ.NewStr("thumbnail_url")),
						λ.NewStr("duration"):    λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒvideo_data, "get", nil), λ.NewStr("video_duration"))),
						λ.NewStr("timestamp"):   λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒvideo_data, "get", nil), λ.NewStr("created_at")), λ.NewInt(1000)),
						λ.NewStr("age_limit"): λ.Cal(ϒparse_age_limit, func() λ.Object {
							if λv := λ.Cal(λ.GetAttr(ϒvideo_data, "get", nil), λ.NewStr("video_rating")); λ.IsTrue(λv) {
								return λv
							} else {
								return ϒrating
							}
						}()),
						λ.NewStr("series"): λ.Cal(ϒtry_get, ϒvideo_data, λ.NewFunction("<lambda>",
							[]λ.Param{
								{Name: "x"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒx = λargs[0]
								)
								return λ.GetItem(λ.GetItem(λ.GetItem(ϒx, λ.NewStr("show")), λ.NewStr("base")), λ.NewStr("display_title"))
							}), ϒcompat_str),
						λ.NewStr("episode_number"): λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒepisode, "get", nil), λ.NewStr("episode_number"))),
						λ.NewStr("episode_id"): λ.Cal(ϒstr_or_none, func() λ.Object {
							if λv := λ.Cal(λ.GetAttr(ϒepisode, "get", nil), λ.NewStr("id")); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Cal(λ.GetAttr(ϒvideo_data, "get", nil), λ.NewStr("episode_id"))
							}
						}()),
						λ.NewStr("season_number"): λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒseason, "get", nil), λ.NewStr("season_number"))),
						λ.NewStr("season_id"): λ.Cal(ϒstr_or_none, func() λ.Object {
							if λv := λ.Cal(λ.GetAttr(ϒseason, "get", nil), λ.NewStr("id")); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Cal(λ.GetAttr(ϒvideo_data, "get", nil), λ.NewStr("season_id"))
							}
						}()),
						λ.NewStr("uploader"):    λ.Cal(λ.GetAttr(ϒchannel, "get", nil), λ.NewStr("name")),
						λ.NewStr("uploader_id"): λ.Cal(ϒstr_or_none, λ.Cal(λ.GetAttr(ϒchannel, "get", nil), λ.NewStr("id"))),
						λ.NewStr("subtitles"):   ϒsubtitles,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("IE_NAME"):       ViceIE_IE_NAME,
				λ.NewStr("_VALID_URL"):    ViceIE__VALID_URL,
				λ.NewStr("_extract_url"):  ViceIE__extract_url,
				λ.NewStr("_extract_urls"): ViceIE__extract_urls,
				λ.NewStr("_real_extract"): ViceIE__real_extract,
			})
		}())
		ViceShowIE = λ.Cal(λ.TypeType, λ.NewStr("ViceShowIE"), λ.NewTuple(ViceBaseIE), func() λ.Dict {
			var (
				ViceShowIE__VALID_URL λ.Object
			)
			ViceShowIE__VALID_URL = λ.NewStr("https?://(?:video\\.vice|(?:www\\.)?vice(?:land|tv))\\.com/(?P<locale>[^/]+)/show/(?P<id>[^/?#&]+)")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): ViceShowIE__VALID_URL,
			})
		}())
		ViceArticleIE = λ.Cal(λ.TypeType, λ.NewStr("ViceArticleIE"), λ.NewTuple(ViceBaseIE), func() λ.Dict {
			var (
				ViceArticleIE_IE_NAME       λ.Object
				ViceArticleIE__VALID_URL    λ.Object
				ViceArticleIE__real_extract λ.Object
			)
			ViceArticleIE_IE_NAME = λ.NewStr("vice:article")
			ViceArticleIE__VALID_URL = λ.NewStr("https://(?:www\\.)?vice\\.com/(?P<locale>[^/]+)/article/(?:[0-9a-z]{6}/)?(?P<id>[^?#]+)")
			ViceArticleIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒ_url_res    λ.Object
						ϒarticle     λ.Object
						ϒbody        λ.Object
						ϒdisplay_id  λ.Object
						ϒembed_code  λ.Object
						ϒlocale      λ.Object
						ϒself        = λargs[0]
						ϒurl         = λargs[1]
						ϒvice_url    λ.Object
						ϒvideo_url   λ.Object
						ϒyoutube_url λ.Object
						τmp0         λ.Object
					)
					τmp0 = λ.Cal(λ.GetAttr(λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl), "groups", nil))
					ϒlocale = λ.GetItem(τmp0, λ.NewInt(0))
					ϒdisplay_id = λ.GetItem(τmp0, λ.NewInt(1))
					ϒarticle = λ.GetItem(λ.Cal(λ.GetAttr(ϒself, "_call_api", nil), λ.NewStr("articles"), λ.NewStr("slug"), ϒdisplay_id, ϒlocale, λ.NewStr("body\n    embed_code")), λ.NewInt(0))
					ϒbody = λ.GetItem(ϒarticle, λ.NewStr("body"))
					ϒ_url_res = λ.NewFunction("_url_res",
						[]λ.Param{
							{Name: "video_url"},
							{Name: "ie_key"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒie_key    = λargs[1]
								ϒvideo_url = λargs[0]
							)
							return λ.NewDictWithTable(map[λ.Object]λ.Object{
								λ.NewStr("_type"):      λ.NewStr("url_transparent"),
								λ.NewStr("url"):        ϒvideo_url,
								λ.NewStr("display_id"): ϒdisplay_id,
								λ.NewStr("ie_key"):     ϒie_key,
							})
						})
					ϒvice_url = λ.Cal(λ.GetAttr(ViceIE, "_extract_url", nil), ϒbody)
					if λ.IsTrue(ϒvice_url) {
						return λ.Cal(ϒ_url_res, ϒvice_url, λ.Cal(λ.GetAttr(ViceIE, "ie_key", nil)))
					}
					ϒembed_code = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr("embedCode=([^&\\'\"]+)"),
						ϒbody,
						λ.NewStr("ooyala embed code"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					if λ.IsTrue(ϒembed_code) {
						return λ.Cal(ϒ_url_res, λ.Mod(λ.NewStr("ooyala:%s"), ϒembed_code), λ.NewStr("Ooyala"))
					}
					ϒyoutube_url = λ.Cal(λ.GetAttr(YoutubeIE, "_extract_url", nil), ϒbody)
					if λ.IsTrue(ϒyoutube_url) {
						return λ.Cal(ϒ_url_res, ϒyoutube_url, λ.Cal(λ.GetAttr(YoutubeIE, "ie_key", nil)))
					}
					ϒvideo_url = λ.Cal(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewStr("data-video-url=\"([^\"]+)\""), λ.GetItem(ϒarticle, λ.NewStr("embed_code")), λ.NewStr("video URL"))
					return λ.Cal(ϒ_url_res, ϒvideo_url, λ.Cal(λ.GetAttr(ViceIE, "ie_key", nil)))
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("IE_NAME"):       ViceArticleIE_IE_NAME,
				λ.NewStr("_VALID_URL"):    ViceArticleIE__VALID_URL,
				λ.NewStr("_real_extract"): ViceArticleIE__real_extract,
			})
		}())
	})
}
