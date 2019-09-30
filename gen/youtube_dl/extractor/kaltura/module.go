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
 * kaltura/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/kaltura.py
 */

package kaltura

import (
	Ωbase64 "github.com/tenta-browser/go-video-downloader/gen/base64"
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωparse "github.com/tenta-browser/go-video-downloader/gen/urllib/parse"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError   λ.Object
	InfoExtractor    λ.Object
	KalturaIE        λ.Object
	ϒclean_html      λ.Object
	ϒcompat_parse_qs λ.Object
	ϒint_or_none     λ.Object
	ϒsmuggle_url     λ.Object
	ϒunsmuggle_url   λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_parse_qs = Ωcompat.ϒcompat_parse_qs
		ϒclean_html = Ωutils.ϒclean_html
		ExtractorError = Ωutils.ExtractorError
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒunsmuggle_url = Ωutils.ϒunsmuggle_url
		ϒsmuggle_url = Ωutils.ϒsmuggle_url
		KalturaIE = λ.Cal(λ.TypeType, λ.NewStr("KalturaIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				KalturaIE__SERVICE_BASE     λ.Object
				KalturaIE__SERVICE_URL      λ.Object
				KalturaIE__VALID_URL        λ.Object
				KalturaIE__extract_url      λ.Object
				KalturaIE__get_video_info   λ.Object
				KalturaIE__kaltura_api_call λ.Object
				KalturaIE__real_extract     λ.Object
			)
			KalturaIE__VALID_URL = λ.NewStr("(?x)\n                (?:\n                    kaltura:(?P<partner_id>\\d+):(?P<id>[0-9a-z_]+)|\n                    https?://\n                        (:?(?:www|cdnapi(?:sec)?)\\.)?kaltura\\.com(?::\\d+)?/\n                        (?:\n                            (?:\n                                # flash player\n                                index\\.php/(?:kwidget|extwidget/preview)|\n                                # html5 player\n                                html5/html5lib/[^/]+/mwEmbedFrame\\.php\n                            )\n                        )(?:/(?P<path>[^?]+))?(?:\\?(?P<query>.*))?\n                )\n                ")
			KalturaIE__SERVICE_URL = λ.NewStr("http://cdnapi.kaltura.com")
			KalturaIE__SERVICE_BASE = λ.NewStr("/api_v3/index.php")
			KalturaIE__extract_url = λ.NewFunction("_extract_url",
				[]λ.Param{
					{Name: "webpage"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒembed_info  λ.Object
						ϒescaped_pid λ.Object
						ϒk           λ.Object
						ϒmobj        λ.Object
						ϒservice_url λ.Object
						ϒurl         λ.Object
						ϒv           λ.Object
						ϒwebpage     = λargs[0]
						τmp0         λ.Object
						τmp1         λ.Object
						τmp2         λ.Object
					)
					ϒmobj = func() λ.Object {
						if λv := λ.Cal(Ωre.ϒsearch, λ.NewStr("(?xs)\n                    kWidget\\.(?:thumb)?[Ee]mbed\\(\n                    \\{.*?\n                        (?P<q1>['\"])wid(?P=q1)\\s*:\\s*\n                        (?P<q2>['\"])_?(?P<partner_id>(?:(?!(?P=q2)).)+)(?P=q2),.*?\n                        (?P<q3>['\"])entry_?[Ii]d(?P=q3)\\s*:\\s*\n                        (?P<q4>['\"])(?P<id>(?:(?!(?P=q4)).)+)(?P=q4)(?:,|\\s*\\})\n                "), ϒwebpage); λ.IsTrue(λv) {
							return λv
						} else if λv := λ.Cal(Ωre.ϒsearch, λ.NewStr("(?xs)\n                    (?P<q1>[\"'])\n                        (?:https?:)?//cdnapi(?:sec)?\\.kaltura\\.com(?::\\d+)?/(?:(?!(?P=q1)).)*\\b(?:p|partner_id)/(?P<partner_id>\\d+)(?:(?!(?P=q1)).)*\n                    (?P=q1).*?\n                    (?:\n                        (?:\n                            entry_?[Ii]d|\n                            (?P<q2>[\"'])entry_?[Ii]d(?P=q2)\n                        )\\s*:\\s*|\n                        \\[\\s*(?P<q2_1>[\"'])entry_?[Ii]d(?P=q2_1)\\s*\\]\\s*=\\s*\n                    )\n                    (?P<q3>[\"'])(?P<id>(?:(?!(?P=q3)).)+)(?P=q3)\n                "), ϒwebpage); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Cal(Ωre.ϒsearch, λ.NewStr("(?xs)\n                    <(?:iframe[^>]+src|meta[^>]+\\bcontent)=(?P<q1>[\"'])\n                      (?:https?:)?//(?:(?:www|cdnapi(?:sec)?)\\.)?kaltura\\.com/(?:(?!(?P=q1)).)*\\b(?:p|partner_id)/(?P<partner_id>\\d+)\n                      (?:(?!(?P=q1)).)*\n                      [?&;]entry_id=(?P<id>(?:(?!(?P=q1))[^&])+)\n                      (?:(?!(?P=q1)).)*\n                    (?P=q1)\n                "), ϒwebpage)
						}
					}()
					if λ.IsTrue(ϒmobj) {
						ϒembed_info = λ.Cal(λ.GetAttr(ϒmobj, "groupdict", nil))
						τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(ϒembed_info, "items", nil)))
						for {
							if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
								break
							}
							τmp2 = τmp1
							ϒk = λ.GetItem(τmp2, λ.NewInt(0))
							ϒv = λ.GetItem(τmp2, λ.NewInt(1))
							λ.SetItem(ϒembed_info, ϒk, λ.Cal(λ.GetAttr(ϒv, "strip", nil)))
						}
						ϒurl = λ.Mod(λ.NewStr("kaltura:%(partner_id)s:%(id)s"), ϒembed_info)
						ϒescaped_pid = λ.Cal(Ωre.ϒescape, λ.GetItem(ϒembed_info, λ.NewStr("partner_id")))
						ϒservice_url = λ.Cal(Ωre.ϒsearch, λ.Mod(λ.NewStr("<script[^>]+src=[\"\\']((?:https?:)?//.+?)/p/%s/sp/%s00/embedIframeJs"), λ.NewTuple(
							ϒescaped_pid,
							ϒescaped_pid,
						)), ϒwebpage)
						if λ.IsTrue(ϒservice_url) {
							ϒurl = λ.Cal(ϒsmuggle_url, ϒurl, λ.NewDictWithTable(map[λ.Object]λ.Object{
								λ.NewStr("service_url"): λ.Cal(λ.GetAttr(ϒservice_url, "group", nil), λ.NewInt(1)),
							}))
						}
						return ϒurl
					}
					return λ.None
				})
			KalturaIE__extract_url = λ.Cal(λ.StaticMethodType, KalturaIE__extract_url)
			KalturaIE__kaltura_api_call = λ.NewFunction("_kaltura_api_call",
				[]λ.Param{
					{Name: "self"},
					{Name: "video_id"},
					{Name: "actions"},
					{Name: "service_url", Def: λ.None},
				},
				0, true, true,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒa           λ.Object
						ϒactions     = λargs[2]
						ϒargs        = λargs[4]
						ϒdata        λ.Object
						ϒi           λ.Object
						ϒk           λ.Object
						ϒkwargs      = λargs[5]
						ϒparams      λ.Object
						ϒself        = λargs[0]
						ϒservice_url = λargs[3]
						ϒstatus      λ.Object
						ϒv           λ.Object
						ϒvideo_id    = λargs[1]
						τmp0         λ.Object
						τmp1         λ.Object
						τmp2         λ.Object
						τmp3         λ.Object
						τmp4         λ.Object
					)
					ϒparams = λ.GetItem(ϒactions, λ.NewInt(0))
					if λ.IsTrue(λ.Gt(λ.Cal(λ.BuiltinLen, ϒactions), λ.NewInt(1))) {
						τmp0 = λ.Cal(λ.BuiltinIter, λ.Call(λ.EnumerateIteratorType, λ.NewArgs(λ.GetItem(ϒactions, λ.NewSlice(λ.NewInt(1), λ.None, λ.None))), λ.KWArgs{
							{Name: "start", Value: λ.NewInt(1)},
						}))
						for {
							if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
								break
							}
							τmp2 = τmp1
							ϒi = λ.GetItem(τmp2, λ.NewInt(0))
							ϒa = λ.GetItem(τmp2, λ.NewInt(1))
							τmp2 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(ϒa, "items", nil)))
							for {
								if τmp3 = λ.NextDefault(τmp2, λ.AfterLast); τmp3 == λ.AfterLast {
									break
								}
								τmp4 = τmp3
								ϒk = λ.GetItem(τmp4, λ.NewInt(0))
								ϒv = λ.GetItem(τmp4, λ.NewInt(1))
								λ.SetItem(ϒparams, λ.Mod(λ.NewStr("%d:%s"), λ.NewTuple(
									ϒi,
									ϒk,
								)), ϒv)
							}
						}
					}
					ϒdata = λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(λ.Unpack(
						λ.Add(func() λ.Object {
							if λv := ϒservice_url; λ.IsTrue(λv) {
								return λv
							} else {
								return λ.GetAttr(ϒself, "_SERVICE_URL", nil)
							}
						}(), λ.GetAttr(ϒself, "_SERVICE_BASE", nil)),
						ϒvideo_id,
						λ.AsStarred(ϒargs),
					)...), λ.KWArgs{
						{Name: "query", Value: ϒparams},
						{Name: "", Value: ϒkwargs},
					})
					ϒstatus = func() λ.Object {
						if λ.IsTrue(λ.Eq(λ.Cal(λ.BuiltinLen, ϒactions), λ.NewInt(1))) {
							return ϒdata
						} else {
							return λ.GetItem(ϒdata, λ.NewInt(0))
						}
					}()
					if λ.IsTrue(λ.Eq(λ.Cal(λ.GetAttr(ϒstatus, "get", nil), λ.NewStr("objectType")), λ.NewStr("KalturaAPIException"))) {
						panic(λ.Raise(λ.Cal(ExtractorError, λ.Mod(λ.NewStr("%s said: %s"), λ.NewTuple(
							λ.GetAttr(ϒself, "IE_NAME", nil),
							λ.GetItem(ϒstatus, λ.NewStr("message")),
						)))))
					}
					return ϒdata
				})
			KalturaIE__get_video_info = λ.NewFunction("_get_video_info",
				[]λ.Param{
					{Name: "self"},
					{Name: "video_id"},
					{Name: "partner_id"},
					{Name: "service_url", Def: λ.None},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒactions     λ.Object
						ϒpartner_id  = λargs[2]
						ϒself        = λargs[0]
						ϒservice_url = λargs[3]
						ϒvideo_id    = λargs[1]
					)
					ϒactions = λ.NewList(
						λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("action"):     λ.NewStr("null"),
							λ.NewStr("apiVersion"): λ.NewStr("3.1.5"),
							λ.NewStr("clientTag"):  λ.NewStr("kdp:v3.8.5"),
							λ.NewStr("format"):     λ.NewInt(1),
							λ.NewStr("service"):    λ.NewStr("multirequest"),
						}),
						λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("expiry"):   λ.NewInt(86400),
							λ.NewStr("service"):  λ.NewStr("session"),
							λ.NewStr("action"):   λ.NewStr("startWidgetSession"),
							λ.NewStr("widgetId"): λ.Mod(λ.NewStr("_%s"), ϒpartner_id),
						}),
						λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("action"):                 λ.NewStr("get"),
							λ.NewStr("entryId"):                ϒvideo_id,
							λ.NewStr("service"):                λ.NewStr("baseentry"),
							λ.NewStr("ks"):                     λ.NewStr("{1:result:ks}"),
							λ.NewStr("responseProfile:fields"): λ.NewStr("createdAt,dataUrl,duration,name,plays,thumbnailUrl,userId"),
							λ.NewStr("responseProfile:type"):   λ.NewInt(1),
						}),
						λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("action"):  λ.NewStr("getbyentryid"),
							λ.NewStr("entryId"): ϒvideo_id,
							λ.NewStr("service"): λ.NewStr("flavorAsset"),
							λ.NewStr("ks"):      λ.NewStr("{1:result:ks}"),
						}),
						λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("action"):              λ.NewStr("list"),
							λ.NewStr("filter:entryIdEqual"): ϒvideo_id,
							λ.NewStr("service"):             λ.NewStr("caption_captionasset"),
							λ.NewStr("ks"):                  λ.NewStr("{1:result:ks}"),
						}),
					)
					return λ.Call(λ.GetAttr(ϒself, "_kaltura_api_call", nil), λ.NewArgs(
						ϒvideo_id,
						ϒactions,
						ϒservice_url,
					), λ.KWArgs{
						{Name: "note", Value: λ.NewStr("Downloading video info JSON")},
					})
				})
			KalturaIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒcaption        λ.Object
						ϒcaption_format λ.Object
						ϒcaptions       λ.Object
						ϒdata_url       λ.Object
						ϒentry_data     λ.Object
						ϒentry_id       λ.Object
						ϒf              λ.Object
						ϒflavor_assets  λ.Object
						ϒformat_id      λ.Object
						ϒformats        λ.Object
						ϒinfo           λ.Object
						ϒks             λ.Object
						ϒm3u8_url       λ.Object
						ϒmobj           λ.Object
						ϒparams         λ.Object
						ϒpartner_id     λ.Object
						ϒpath           λ.Object
						ϒquery          λ.Object
						ϒreference_id   λ.Object
						ϒreferrer       λ.Object
						ϒself           = λargs[0]
						ϒsign_url       λ.Object
						ϒsmuggled_data  λ.Object
						ϒsource_url     λ.Object
						ϒsplitted_path  λ.Object
						ϒsubtitles      λ.Object
						ϒurl            = λargs[1]
						ϒvcodec         λ.Object
						ϒvideo_url      λ.Object
						ϒwebpage        λ.Object
						τmp0            λ.Object
						τmp1            λ.Object
						τmp2            λ.Object
					)
					τmp0 = λ.Cal(ϒunsmuggle_url, ϒurl, λ.NewDictWithTable(map[λ.Object]λ.Object{}))
					ϒurl = λ.GetItem(τmp0, λ.NewInt(0))
					ϒsmuggled_data = λ.GetItem(τmp0, λ.NewInt(1))
					ϒmobj = λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl)
					τmp0 = λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("partner_id"), λ.NewStr("id"))
					ϒpartner_id = λ.GetItem(τmp0, λ.NewInt(0))
					ϒentry_id = λ.GetItem(τmp0, λ.NewInt(1))
					ϒks = λ.None
					ϒcaptions = λ.None
					if λ.IsTrue(func() λ.Object {
						if λv := ϒpartner_id; !λ.IsTrue(λv) {
							return λv
						} else {
							return ϒentry_id
						}
					}()) {
						τmp0 = λ.Cal(λ.GetAttr(ϒself, "_get_video_info", nil), ϒentry_id, ϒpartner_id, λ.Cal(λ.GetAttr(ϒsmuggled_data, "get", nil), λ.NewStr("service_url")))
						_ = λ.GetItem(τmp0, λ.NewInt(0))
						ϒinfo = λ.GetItem(τmp0, λ.NewInt(1))
						ϒflavor_assets = λ.GetItem(τmp0, λ.NewInt(2))
						ϒcaptions = λ.GetItem(τmp0, λ.NewInt(3))
					} else {
						τmp0 = λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("path"), λ.NewStr("query"))
						ϒpath = λ.GetItem(τmp0, λ.NewInt(0))
						ϒquery = λ.GetItem(τmp0, λ.NewInt(1))
						if λ.IsTrue(func() λ.Object {
							if λv := λ.NewBool(!λ.IsTrue(ϒpath)); !λ.IsTrue(λv) {
								return λv
							} else {
								return λ.NewBool(!λ.IsTrue(ϒquery))
							}
						}()) {
							panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.NewStr("Invalid URL")), λ.KWArgs{
								{Name: "expected", Value: λ.True},
							})))
						}
						ϒparams = λ.NewDictWithTable(map[λ.Object]λ.Object{})
						if λ.IsTrue(ϒquery) {
							ϒparams = λ.Cal(ϒcompat_parse_qs, ϒquery)
						}
						if λ.IsTrue(ϒpath) {
							ϒsplitted_path = λ.Cal(λ.GetAttr(ϒpath, "split", nil), λ.NewStr("/"))
							λ.Cal(λ.GetAttr(ϒparams, "update", nil), λ.Cal(λ.DictType, λ.Cal(λ.ZipIteratorType, λ.GetItem(ϒsplitted_path, λ.NewSlice(λ.None, λ.None, λ.NewInt(2))), λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
								nil,
								0, false, false,
								func(λargs []λ.Object) λ.Object {
									return λ.NewGenerator(func(λgy λ.Yielder) λ.Object {
										var (
											ϒv   λ.Object
											τmp0 λ.Object
											τmp1 λ.Object
										)
										τmp0 = λ.Cal(λ.BuiltinIter, λ.GetItem(ϒsplitted_path, λ.NewSlice(λ.NewInt(1), λ.None, λ.NewInt(2))))
										for {
											if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
												break
											}
											ϒv = τmp1
											λgy.Yield(λ.NewList(ϒv))
										}
										return λ.None
									})
								}))))))
						}
						if λ.IsTrue(λ.NewBool(λ.Contains(ϒparams, λ.NewStr("wid")))) {
							ϒpartner_id = λ.GetItem(λ.GetItem(λ.GetItem(ϒparams, λ.NewStr("wid")), λ.NewInt(0)), λ.NewSlice(λ.NewInt(1), λ.None, λ.None))
						} else {
							if λ.IsTrue(λ.NewBool(λ.Contains(ϒparams, λ.NewStr("p")))) {
								ϒpartner_id = λ.GetItem(λ.GetItem(ϒparams, λ.NewStr("p")), λ.NewInt(0))
							} else {
								if λ.IsTrue(λ.NewBool(λ.Contains(ϒparams, λ.NewStr("partner_id")))) {
									ϒpartner_id = λ.GetItem(λ.GetItem(ϒparams, λ.NewStr("partner_id")), λ.NewInt(0))
								} else {
									panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.NewStr("Invalid URL")), λ.KWArgs{
										{Name: "expected", Value: λ.True},
									})))
								}
							}
						}
						if λ.IsTrue(λ.NewBool(λ.Contains(ϒparams, λ.NewStr("entry_id")))) {
							ϒentry_id = λ.GetItem(λ.GetItem(ϒparams, λ.NewStr("entry_id")), λ.NewInt(0))
							τmp0 = λ.Cal(λ.GetAttr(ϒself, "_get_video_info", nil), ϒentry_id, ϒpartner_id)
							_ = λ.GetItem(τmp0, λ.NewInt(0))
							ϒinfo = λ.GetItem(τmp0, λ.NewInt(1))
							ϒflavor_assets = λ.GetItem(τmp0, λ.NewInt(2))
							ϒcaptions = λ.GetItem(τmp0, λ.NewInt(3))
						} else {
							if λ.IsTrue(func() λ.Object {
								if λv := λ.NewBool(λ.Contains(ϒparams, λ.NewStr("uiconf_id"))); !λ.IsTrue(λv) {
									return λv
								} else {
									return λ.NewBool(λ.Contains(ϒparams, λ.NewStr("flashvars[referenceId]")))
								}
							}()) {
								ϒreference_id = λ.GetItem(λ.GetItem(ϒparams, λ.NewStr("flashvars[referenceId]")), λ.NewInt(0))
								ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒreference_id)
								ϒentry_data = λ.GetItem(λ.Cal(λ.GetAttr(ϒself, "_parse_json", nil), λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewStr("window\\.kalturaIframePackageData\\s*=\\s*({.*});"), ϒwebpage, λ.NewStr("kalturaIframePackageData")), ϒreference_id), λ.NewStr("entryResult"))
								τmp0 = λ.NewTuple(
									λ.GetItem(ϒentry_data, λ.NewStr("meta")),
									λ.GetItem(λ.GetItem(ϒentry_data, λ.NewStr("contextData")), λ.NewStr("flavorAssets")),
								)
								ϒinfo = λ.GetItem(τmp0, λ.NewInt(0))
								ϒflavor_assets = λ.GetItem(τmp0, λ.NewInt(1))
								ϒentry_id = λ.GetItem(ϒinfo, λ.NewStr("id"))
								τmp0, τmp1 = func() (λexit λ.Object, λret λ.Object) {
									defer λ.CatchMulti(
										nil,
										&λ.Catcher{ExtractorError, func(λex λ.BaseException) {
											// pass
										}},
									)
									τmp2 = λ.Cal(λ.GetAttr(ϒself, "_get_video_info", nil), ϒentry_id, ϒpartner_id)
									_ = λ.GetItem(τmp2, λ.NewInt(0))
									ϒinfo = λ.GetItem(τmp2, λ.NewInt(1))
									ϒflavor_assets = λ.GetItem(τmp2, λ.NewInt(2))
									ϒcaptions = λ.GetItem(τmp2, λ.NewInt(3))
									return λ.BlockExitNormally, nil
								}()
							} else {
								panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.NewStr("Invalid URL")), λ.KWArgs{
									{Name: "expected", Value: λ.True},
								})))
							}
						}
						ϒks = λ.GetItem(λ.Cal(λ.GetAttr(ϒparams, "get", nil), λ.NewStr("flashvars[ks]"), λ.NewList(λ.None)), λ.NewInt(0))
					}
					ϒsource_url = λ.Cal(λ.GetAttr(ϒsmuggled_data, "get", nil), λ.NewStr("source_url"))
					if λ.IsTrue(ϒsource_url) {
						ϒreferrer = λ.Cal(λ.GetAttr(λ.Cal(Ωbase64.ϒb64encode, λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(λ.NewStr("://"), "join", nil), λ.GetItem(λ.Cal(Ωparse.ϒurlparse, ϒsource_url), λ.NewSlice(λ.None, λ.NewInt(2), λ.None))), "encode", nil), λ.NewStr("utf-8"))), "decode", nil), λ.NewStr("utf-8"))
					} else {
						ϒreferrer = λ.None
					}
					ϒsign_url = λ.NewFunction("sign_url",
						[]λ.Param{
							{Name: "unsigned_url"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒunsigned_url = λargs[0]
								τmp0          λ.Object
							)
							if λ.IsTrue(ϒks) {
								τmp0 = λ.IAdd(ϒunsigned_url, λ.Mod(λ.NewStr("/ks/%s"), ϒks))
								ϒunsigned_url = τmp0
							}
							if λ.IsTrue(ϒreferrer) {
								τmp0 = λ.IAdd(ϒunsigned_url, λ.Mod(λ.NewStr("?referrer=%s"), ϒreferrer))
								ϒunsigned_url = τmp0
							}
							return ϒunsigned_url
						})
					ϒdata_url = λ.GetItem(ϒinfo, λ.NewStr("dataUrl"))
					if λ.IsTrue(λ.NewBool(λ.Contains(ϒdata_url, λ.NewStr("/flvclipper/")))) {
						ϒdata_url = λ.Cal(Ωre.ϒsub, λ.NewStr("/flvclipper/.*"), λ.NewStr("/serveFlavor"), ϒdata_url)
					}
					ϒformats = λ.NewList()
					τmp1 = λ.Cal(λ.BuiltinIter, ϒflavor_assets)
					for {
						if τmp0 = λ.NextDefault(τmp1, λ.AfterLast); τmp0 == λ.AfterLast {
							break
						}
						ϒf = τmp0
						if λ.IsTrue(λ.Ne(λ.Cal(λ.GetAttr(ϒf, "get", nil), λ.NewStr("status")), λ.NewInt(2))) {
							continue
						}
						if λ.IsTrue(λ.Eq(λ.Cal(λ.GetAttr(ϒf, "get", nil), λ.NewStr("fileExt")), λ.NewStr("chun"))) {
							continue
						}
						if λ.IsTrue(λ.Eq(λ.Cal(λ.GetAttr(ϒf, "get", nil), λ.NewStr("fileExt")), λ.NewStr("wvm"))) {
							continue
						}
						if λ.IsTrue(λ.NewBool(!λ.IsTrue(λ.Cal(λ.GetAttr(ϒf, "get", nil), λ.NewStr("fileExt"))))) {
							if λ.IsTrue(λ.Eq(λ.Cal(λ.GetAttr(ϒf, "get", nil), λ.NewStr("containerFormat")), λ.NewStr("qt"))) {
								λ.SetItem(ϒf, λ.NewStr("fileExt"), λ.NewStr("mov"))
							} else {
								λ.SetItem(ϒf, λ.NewStr("fileExt"), λ.NewStr("mp4"))
							}
						}
						ϒvideo_url = λ.Cal(ϒsign_url, λ.Mod(λ.NewStr("%s/flavorId/%s"), λ.NewTuple(
							ϒdata_url,
							λ.GetItem(ϒf, λ.NewStr("id")),
						)))
						ϒformat_id = λ.Mod(λ.NewStr("%(fileExt)s-%(bitrate)s"), ϒf)
						if λ.IsTrue(func() λ.Object {
							if λv := λ.NewBool(λ.Cal(λ.GetAttr(ϒf, "get", nil), λ.NewStr("isOriginal")) == λ.True); !λ.IsTrue(λv) {
								return λv
							} else {
								return λ.NewBool(!λ.IsTrue(λ.Cal(λ.GetAttr(ϒself, "_is_valid_url", nil), ϒvideo_url, ϒentry_id, ϒformat_id)))
							}
						}()) {
							continue
						}
						ϒvcodec = func() λ.Object {
							if λ.IsTrue(func() λ.Object {
								if λv := λ.NewBool(!λ.Contains(ϒf, λ.NewStr("videoCodecId"))); !λ.IsTrue(λv) {
									return λv
								} else {
									return λ.Eq(λ.Cal(λ.GetAttr(ϒf, "get", nil), λ.NewStr("frameRate")), λ.NewInt(0))
								}
							}()) {
								return λ.NewStr("none")
							} else {
								return λ.Cal(λ.GetAttr(ϒf, "get", nil), λ.NewStr("videoCodecId"))
							}
						}()
						λ.Cal(λ.GetAttr(ϒformats, "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("format_id"): ϒformat_id,
							λ.NewStr("ext"):       λ.Cal(λ.GetAttr(ϒf, "get", nil), λ.NewStr("fileExt")),
							λ.NewStr("tbr"):       λ.Cal(ϒint_or_none, λ.GetItem(ϒf, λ.NewStr("bitrate"))),
							λ.NewStr("fps"):       λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒf, "get", nil), λ.NewStr("frameRate"))),
							λ.NewStr("filesize_approx"): λ.Call(ϒint_or_none, λ.NewArgs(λ.Cal(λ.GetAttr(ϒf, "get", nil), λ.NewStr("size"))), λ.KWArgs{
								{Name: "invscale", Value: λ.NewInt(1024)},
							}),
							λ.NewStr("container"): λ.Cal(λ.GetAttr(ϒf, "get", nil), λ.NewStr("containerFormat")),
							λ.NewStr("vcodec"):    ϒvcodec,
							λ.NewStr("height"):    λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒf, "get", nil), λ.NewStr("height"))),
							λ.NewStr("width"):     λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒf, "get", nil), λ.NewStr("width"))),
							λ.NewStr("url"):       ϒvideo_url,
						}))
					}
					if λ.IsTrue(λ.NewBool(λ.Contains(ϒdata_url, λ.NewStr("/playManifest/")))) {
						ϒm3u8_url = λ.Cal(ϒsign_url, λ.Cal(λ.GetAttr(ϒdata_url, "replace", nil), λ.NewStr("format/url"), λ.NewStr("format/applehttp")))
						λ.Cal(λ.GetAttr(ϒformats, "extend", nil), λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
							ϒm3u8_url,
							ϒentry_id,
							λ.NewStr("mp4"),
							λ.NewStr("m3u8_native"),
						), λ.KWArgs{
							{Name: "m3u8_id", Value: λ.NewStr("hls")},
							{Name: "fatal", Value: λ.False},
						}))
					}
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					ϒsubtitles = λ.NewDictWithTable(map[λ.Object]λ.Object{})
					if λ.IsTrue(ϒcaptions) {
						τmp1 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(ϒcaptions, "get", nil), λ.NewStr("objects"), λ.NewList()))
						for {
							if τmp0 = λ.NextDefault(τmp1, λ.AfterLast); τmp0 == λ.AfterLast {
								break
							}
							ϒcaption = τmp0
							if λ.IsTrue(λ.Ne(λ.Cal(λ.GetAttr(ϒcaption, "get", nil), λ.NewStr("status")), λ.NewInt(2))) {
								continue
							}
							if λ.IsTrue(λ.NewBool(!λ.IsTrue(λ.Cal(λ.GetAttr(ϒcaption, "get", nil), λ.NewStr("id"))))) {
								continue
							}
							ϒcaption_format = λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒcaption, "get", nil), λ.NewStr("format")))
							λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒsubtitles, "setdefault", nil), func() λ.Object {
								if λv := λ.Cal(λ.GetAttr(ϒcaption, "get", nil), λ.NewStr("languageCode")); λ.IsTrue(λv) {
									return λv
								} else {
									return λ.Cal(λ.GetAttr(ϒcaption, "get", nil), λ.NewStr("language"))
								}
							}(), λ.NewList()), "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
								λ.NewStr("url"): λ.Mod(λ.NewStr("%s/api_v3/service/caption_captionasset/action/serve/captionAssetId/%s"), λ.NewTuple(
									λ.GetAttr(ϒself, "_SERVICE_URL", nil),
									λ.GetItem(ϒcaption, λ.NewStr("id")),
								)),
								λ.NewStr("ext"): func() λ.Object {
									if λv := λ.Cal(λ.GetAttr(ϒcaption, "get", nil), λ.NewStr("fileExt")); λ.IsTrue(λv) {
										return λv
									} else if λv := λ.Cal(λ.GetAttr(λ.GetAttr(ϒself, "_CAPTION_TYPES", nil), "get", nil), ϒcaption_format); λ.IsTrue(λv) {
										return λv
									} else {
										return λ.NewStr("ttml")
									}
								}(),
							}))
						}
					}
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          ϒentry_id,
						λ.NewStr("title"):       λ.GetItem(ϒinfo, λ.NewStr("name")),
						λ.NewStr("formats"):     ϒformats,
						λ.NewStr("subtitles"):   ϒsubtitles,
						λ.NewStr("description"): λ.Cal(ϒclean_html, λ.Cal(λ.GetAttr(ϒinfo, "get", nil), λ.NewStr("description"))),
						λ.NewStr("thumbnail"):   λ.Cal(λ.GetAttr(ϒinfo, "get", nil), λ.NewStr("thumbnailUrl")),
						λ.NewStr("duration"):    λ.Cal(λ.GetAttr(ϒinfo, "get", nil), λ.NewStr("duration")),
						λ.NewStr("timestamp"):   λ.Cal(λ.GetAttr(ϒinfo, "get", nil), λ.NewStr("createdAt")),
						λ.NewStr("uploader_id"): func() λ.Object {
							if λ.IsTrue(λ.Ne(λ.Cal(λ.GetAttr(ϒinfo, "get", nil), λ.NewStr("userId")), λ.NewStr("None"))) {
								return λ.Cal(λ.GetAttr(ϒinfo, "get", nil), λ.NewStr("userId"))
							} else {
								return λ.None
							}
						}(),
						λ.NewStr("view_count"): λ.Cal(λ.GetAttr(ϒinfo, "get", nil), λ.NewStr("plays")),
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_SERVICE_BASE"):     KalturaIE__SERVICE_BASE,
				λ.NewStr("_SERVICE_URL"):      KalturaIE__SERVICE_URL,
				λ.NewStr("_VALID_URL"):        KalturaIE__VALID_URL,
				λ.NewStr("_extract_url"):      KalturaIE__extract_url,
				λ.NewStr("_get_video_info"):   KalturaIE__get_video_info,
				λ.NewStr("_kaltura_api_call"): KalturaIE__kaltura_api_call,
				λ.NewStr("_real_extract"):     KalturaIE__real_extract,
			})
		}())
	})
}
