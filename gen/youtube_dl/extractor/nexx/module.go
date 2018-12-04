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
 * nexx/module.go: transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/nexx.py
 */

package nexx

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωtime "github.com/tenta-browser/go-video-downloader/gen/time"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError      λ.Object
	InfoExtractor       λ.Object
	NexxEmbedIE         λ.Object
	NexxIE              λ.Object
	ϒcompat_str         λ.Object
	ϒint_or_none        λ.Object
	ϒparse_duration     λ.Object
	ϒtry_get            λ.Object
	ϒurlencode_postdata λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_str = Ωcompat.ϒcompat_str
		ExtractorError = Ωutils.ExtractorError
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒparse_duration = Ωutils.ϒparse_duration
		ϒtry_get = Ωutils.ϒtry_get
		ϒurlencode_postdata = Ωutils.ϒurlencode_postdata
		NexxIE = λ.Cal(λ.TypeType, λ.NewStr("NexxIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				NexxIE__TESTS                 λ.Object
				NexxIE__VALID_URL             λ.Object
				NexxIE__extract_azure_formats λ.Object
				NexxIE__extract_free_formats  λ.Object
				NexxIE__real_extract          λ.Object
			)
			NexxIE__VALID_URL = λ.NewStr("(?x)\n                        (?:\n                            https?://api\\.nexx(?:\\.cloud|cdn\\.com)/v3/(?P<domain_id>\\d+)/videos/byid/|\n                            nexx:(?:(?P<domain_id_s>\\d+):)?|\n                            https?://arc\\.nexx\\.cloud/api/video/\n                        )\n                        (?P<id>\\d+)\n                    ")
			NexxIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("https://api.nexx.cloud/v3/748/videos/byid/128907"),
					λ.NewStr("md5"): λ.NewStr("31899fd683de49ad46f4ee67e53e83fe"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("128907"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Stiftung Warentest"),
						λ.NewStr("alt_title"):   λ.NewStr("Wie ein Test abläuft"),
						λ.NewStr("description"): λ.NewStr("md5:d1ddb1ef63de721132abd38639cc2fd2"),
						λ.NewStr("creator"):     λ.NewStr("SPIEGEL TV"),
						λ.NewStr("thumbnail"):   λ.NewStr("re:^https?://.*\\.jpg$"),
						λ.NewStr("duration"):    λ.NewInt(2509),
						λ.NewStr("timestamp"):   λ.NewInt(1384264416),
						λ.NewStr("upload_date"): λ.NewStr("20131112"),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("https://api.nexx.cloud/v3/741/videos/byid/247858"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):             λ.NewStr("247858"),
						λ.NewStr("ext"):            λ.NewStr("mp4"),
						λ.NewStr("title"):          λ.NewStr("Return of the Golden Child (OV)"),
						λ.NewStr("description"):    λ.NewStr("md5:5d969537509a92b733de21bae249dc63"),
						λ.NewStr("release_year"):   λ.NewInt(2017),
						λ.NewStr("thumbnail"):      λ.NewStr("re:^https?://.*\\.jpg$"),
						λ.NewStr("duration"):       λ.NewInt(1397),
						λ.NewStr("timestamp"):      λ.NewInt(1495033267),
						λ.NewStr("upload_date"):    λ.NewStr("20170517"),
						λ.NewStr("episode_number"): λ.NewInt(2),
						λ.NewStr("season_number"):  λ.NewInt(2),
					}),
					λ.NewStr("params"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("skip_download"): λ.True,
					}),
					λ.NewStr("skip"): λ.NewStr("HTTP Error 404: Not Found"),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("nexx:741:1269984"),
					λ.NewStr("md5"): λ.NewStr("c714b5b238b2958dc8d5642addba6886"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("1269984"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("1 TAG ohne KLO... wortwörtlich! 😑"),
						λ.NewStr("alt_title"):   λ.NewStr("1 TAG ohne KLO... wortwörtlich! 😑"),
						λ.NewStr("thumbnail"):   λ.NewStr("re:^https?://.*\\.jpg$"),
						λ.NewStr("duration"):    λ.NewInt(607),
						λ.NewStr("timestamp"):   λ.NewInt(1518614955),
						λ.NewStr("upload_date"): λ.NewStr("20180214"),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("nexx:747:1533779"),
					λ.NewStr("md5"): λ.NewStr("6bf6883912b82b7069fb86c2297e9893"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("1533779"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Aufregung um ausgebrochene Raubtiere"),
						λ.NewStr("alt_title"):   λ.NewStr("Eifel-Zoo"),
						λ.NewStr("description"): λ.NewStr("md5:f21375c91c74ad741dcb164c427999d2"),
						λ.NewStr("thumbnail"):   λ.NewStr("re:^https?://.*\\.jpg$"),
						λ.NewStr("duration"):    λ.NewInt(111),
						λ.NewStr("timestamp"):   λ.NewInt(1527874460),
						λ.NewStr("upload_date"): λ.NewStr("20180601"),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://api.nexxcdn.com/v3/748/videos/byid/128907"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("nexx:748:128907"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("nexx:128907"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://arc.nexx.cloud/api/video/128907.json"),
					λ.NewStr("only_matching"): λ.True,
				}),
			)
			NexxIE__extract_free_formats = λ.NewFunction("_extract_free_formats",
				[]λ.Param{
					{Name: "self"},
					{Name: "video"},
					{Name: "video_id"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒa            λ.Object
						ϒcdn          λ.Object
						ϒcdn_provider λ.Object
						ϒf            λ.Object
						ϒfd           λ.Object
						ϒfilename     λ.Object
						ϒformats      λ.Object
						ϒh            λ.Object
						ϒhash         λ.Object
						ϒhttp_base    λ.Object
						ϒi            λ.Object
						ϒk            λ.Object
						ϒp            λ.Object
						ϒp0           λ.Object
						ϒps           λ.Object
						ϒs            λ.Object
						ϒself         = λargs[0]
						ϒstream_data  λ.Object
						ϒt            λ.Object
						ϒtbr          λ.Object
						ϒvideo        = λargs[1]
						ϒvideo_id     = λargs[2]
						ϒwidth_height λ.Object
						τmp0          λ.Object
						τmp1          λ.Object
						τmp2          λ.Object
					)
					ϒstream_data = λ.GetItem(ϒvideo, λ.NewStr("streamdata"))
					ϒcdn = λ.GetItem(ϒstream_data, λ.NewStr("cdnType"))
					if !λ.IsTrue(λ.Eq(ϒcdn, λ.NewStr("free"))) {
						panic(λ.Raise(λ.Cal(λ.AssertionErrorType)))
					}
					ϒhash = λ.GetItem(λ.GetItem(ϒvideo, λ.NewStr("general")), λ.NewStr("hash"))
					ϒps = λ.Cal(ϒcompat_str, λ.GetItem(ϒstream_data, λ.NewStr("originalDomain")))
					if λ.IsTrue(λ.Eq(λ.GetItem(ϒstream_data, λ.NewStr("applyFolderHierarchy")), λ.NewInt(1))) {
						ϒs = λ.GetItem(λ.Mod(λ.NewStr("%04d"), λ.Cal(λ.IntType, ϒvideo_id)), λ.NewSlice(λ.None, λ.None, λ.Neg(λ.NewInt(1))))
						τmp0 = λ.IAdd(ϒps, λ.Mod(λ.NewStr("/%s/%s"), λ.NewTuple(
							λ.GetItem(ϒs, λ.NewSlice(λ.NewInt(0), λ.NewInt(2), λ.None)),
							λ.GetItem(ϒs, λ.NewSlice(λ.NewInt(2), λ.NewInt(4), λ.None)),
						)))
						ϒps = τmp0
					}
					τmp0 = λ.IAdd(ϒps, λ.Mod(λ.NewStr("/%s/%s_"), λ.NewTuple(
						ϒvideo_id,
						ϒhash,
					)))
					ϒps = τmp0
					ϒt = λ.Add(λ.NewStr("http://%s"), ϒps)
					ϒfd = λ.Cal(λ.GetAttr(λ.GetItem(ϒstream_data, λ.NewStr("azureFileDistribution")), "split", nil), λ.NewStr(","))
					ϒcdn_provider = λ.GetItem(ϒstream_data, λ.NewStr("cdnProvider"))
					ϒp0 = λ.NewFunction("p0",
						[]λ.Param{
							{Name: "p"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒp = λargs[0]
							)
							return func() λ.Object {
								if λ.IsTrue(λ.Eq(λ.GetItem(ϒstream_data, λ.NewStr("applyAzureStructure")), λ.NewInt(1))) {
									return λ.Mod(λ.NewStr("_%s"), ϒp)
								} else {
									return λ.NewStr("")
								}
							}()
						})
					ϒformats = λ.NewList()
					if λ.IsTrue(λ.Eq(ϒcdn_provider, λ.NewStr("ak"))) {
						τmp0 = λ.IAdd(ϒt, λ.NewStr(","))
						ϒt = τmp0
						τmp0 = λ.Cal(λ.BuiltinIter, ϒfd)
						for {
							if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
								break
							}
							ϒi = τmp1
							ϒp = λ.Cal(λ.GetAttr(ϒi, "split", nil), λ.NewStr(":"))
							τmp2 = λ.IAdd(ϒt, λ.Add(λ.Add(λ.GetItem(ϒp, λ.NewInt(1)), λ.Cal(ϒp0, λ.Cal(λ.IntType, λ.GetItem(ϒp, λ.NewInt(0))))), λ.NewStr(",")))
							ϒt = τmp2
						}
						τmp0 = λ.IAdd(ϒt, λ.NewStr(".mp4.csmil/master.%s"))
						ϒt = τmp0
					} else {
						if λ.IsTrue(λ.Eq(ϒcdn_provider, λ.NewStr("ce"))) {
							ϒk = λ.Cal(λ.GetAttr(ϒt, "split", nil), λ.NewStr("/"))
							ϒh = λ.Cal(λ.GetAttr(ϒk, "pop", nil))
							τmp0 = λ.Cal(λ.GetAttr(λ.NewStr("/"), "join", nil), ϒk)
							ϒhttp_base = τmp0
							ϒt = τmp0
							ϒhttp_base = λ.Mod(ϒhttp_base, λ.GetItem(ϒstream_data, λ.NewStr("cdnPathHTTP")))
							τmp0 = λ.IAdd(ϒt, λ.NewStr("/asset.ism/manifest.%s?dcp_ver=aos4&videostream="))
							ϒt = τmp0
							τmp0 = λ.Cal(λ.BuiltinIter, ϒfd)
							for {
								if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
									break
								}
								ϒi = τmp1
								ϒp = λ.Cal(λ.GetAttr(ϒi, "split", nil), λ.NewStr(":"))
								ϒtbr = λ.Cal(λ.IntType, λ.GetItem(ϒp, λ.NewInt(0)))
								ϒfilename = λ.Mod(λ.NewStr("%s%s%s.mp4"), λ.NewTuple(
									ϒh,
									λ.GetItem(ϒp, λ.NewInt(1)),
									λ.Cal(ϒp0, ϒtbr),
								))
								ϒf = λ.NewDictWithTable(map[λ.Object]λ.Object{
									λ.NewStr("url"): λ.Add(λ.Add(ϒhttp_base, λ.NewStr("/")), ϒfilename),
									λ.NewStr("format_id"): λ.Mod(λ.NewStr("%s-http-%d"), λ.NewTuple(
										ϒcdn,
										ϒtbr,
									)),
									λ.NewStr("tbr"): ϒtbr,
								})
								ϒwidth_height = λ.Cal(λ.GetAttr(λ.GetItem(ϒp, λ.NewInt(1)), "split", nil), λ.NewStr("x"))
								if λ.IsTrue(λ.Eq(λ.Cal(λ.BuiltinLen, ϒwidth_height), λ.NewInt(2))) {
									λ.Cal(λ.GetAttr(ϒf, "update", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
										λ.NewStr("width"):  λ.Cal(ϒint_or_none, λ.GetItem(ϒwidth_height, λ.NewInt(0))),
										λ.NewStr("height"): λ.Cal(ϒint_or_none, λ.GetItem(ϒwidth_height, λ.NewInt(1))),
									}))
								}
								λ.Cal(λ.GetAttr(ϒformats, "append", nil), ϒf)
								ϒa = λ.Add(ϒfilename, λ.Mod(λ.NewStr(":%s"), λ.Mul(ϒtbr, λ.NewInt(1000))))
								τmp2 = λ.IAdd(ϒt, λ.Add(ϒa, λ.NewStr(",")))
								ϒt = τmp2
							}
							ϒt = λ.Add(λ.Add(λ.GetItem(ϒt, λ.NewSlice(λ.None, λ.Neg(λ.NewInt(1)), λ.None)), λ.NewStr("&audiostream=")), λ.GetItem(λ.Cal(λ.GetAttr(ϒa, "split", nil), λ.NewStr(":")), λ.NewInt(0)))
						} else {
							if !λ.IsTrue(λ.False) {
								panic(λ.Raise(λ.Cal(λ.AssertionErrorType)))
							}
						}
					}
					if λ.IsTrue(λ.Eq(ϒcdn_provider, λ.NewStr("ce"))) {
						λ.Cal(λ.GetAttr(ϒformats, "extend", nil), λ.Call(λ.GetAttr(ϒself, "_extract_mpd_formats", nil), λ.NewArgs(
							λ.Mod(ϒt, λ.NewTuple(
								λ.GetItem(ϒstream_data, λ.NewStr("cdnPathDASH")),
								λ.NewStr("mpd"),
							)),
							ϒvideo_id,
						), λ.KWArgs{
							{Name: "mpd_id", Value: λ.Mod(λ.NewStr("%s-dash"), ϒcdn)},
							{Name: "fatal", Value: λ.False},
						}))
					}
					λ.Cal(λ.GetAttr(ϒformats, "extend", nil), λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
						λ.Mod(ϒt, λ.NewTuple(
							λ.GetItem(ϒstream_data, λ.NewStr("cdnPathHLS")),
							λ.NewStr("m3u8"),
						)),
						ϒvideo_id,
						λ.NewStr("mp4"),
					), λ.KWArgs{
						{Name: "entry_protocol", Value: λ.NewStr("m3u8_native")},
						{Name: "m3u8_id", Value: λ.Mod(λ.NewStr("%s-hls"), ϒcdn)},
						{Name: "fatal", Value: λ.False},
					}))
					return ϒformats
				})
			NexxIE__extract_azure_formats = λ.NewFunction("_extract_azure_formats",
				[]λ.Param{
					{Name: "self"},
					{Name: "video"},
					{Name: "video_id"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒazure_file_distribution λ.Object
						ϒazure_locator           λ.Object
						ϒazure_manifest_url      λ.Object
						ϒazure_progressive_base  λ.Object
						ϒazure_stream_base       λ.Object
						ϒcdn                     λ.Object
						ϒf                       λ.Object
						ϒfd                      λ.Object
						ϒfds                     λ.Object
						ϒformats                 λ.Object
						ϒget_cdn_shield_base     λ.Object
						ϒis_ml                   λ.Object
						ϒlanguage                λ.Object
						ϒprotection_token        λ.Object
						ϒself                    = λargs[0]
						ϒss                      λ.Object
						ϒstream_data             λ.Object
						ϒtbr                     λ.Object
						ϒvideo                   = λargs[1]
						ϒvideo_id                = λargs[2]
						ϒwidth_height            λ.Object
						τmp0                     λ.Object
						τmp1                     λ.Object
					)
					ϒstream_data = λ.GetItem(ϒvideo, λ.NewStr("streamdata"))
					ϒcdn = λ.GetItem(ϒstream_data, λ.NewStr("cdnType"))
					if !λ.IsTrue(λ.Eq(ϒcdn, λ.NewStr("azure"))) {
						panic(λ.Raise(λ.Cal(λ.AssertionErrorType)))
					}
					ϒazure_locator = λ.GetItem(ϒstream_data, λ.NewStr("azureLocator"))
					ϒget_cdn_shield_base = λ.NewFunction("get_cdn_shield_base",
						[]λ.Param{
							{Name: "shield_type", Def: λ.NewStr("")},
							{Name: "static", Def: λ.False},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒaccount     λ.Object
								ϒcdn_shield  λ.Object
								ϒprefix      λ.Object
								ϒsecure      λ.Object
								ϒshield_type = λargs[0]
								ϒstatic      = λargs[1]
								τmp0         λ.Object
								τmp1         λ.Object
							)
							τmp0 = λ.Cal(λ.BuiltinIter, λ.NewTuple(
								λ.NewStr(""),
								λ.NewStr("s"),
							))
							for {
								if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
									break
								}
								ϒsecure = τmp1
								ϒcdn_shield = λ.Cal(λ.GetAttr(ϒstream_data, "get", nil), λ.Mod(λ.NewStr("cdnShield%sHTTP%s"), λ.NewTuple(
									ϒshield_type,
									λ.Cal(λ.GetAttr(ϒsecure, "upper", nil)),
								)))
								if λ.IsTrue(ϒcdn_shield) {
									return λ.Mod(λ.NewStr("http%s://%s"), λ.NewTuple(
										ϒsecure,
										ϒcdn_shield,
									))
								}
							}
							if τmp1 == λ.AfterLast {
								if λ.IsTrue(λ.NewBool(λ.Contains(λ.GetItem(ϒstream_data, λ.NewStr("azureAccount")), λ.NewStr("fb")))) {
									ϒprefix = func() λ.Object {
										if λ.IsTrue(ϒstatic) {
											return λ.NewStr("df")
										} else {
											return λ.NewStr("f")
										}
									}()
								} else {
									ϒprefix = func() λ.Object {
										if λ.IsTrue(ϒstatic) {
											return λ.NewStr("d")
										} else {
											return λ.NewStr("p")
										}
									}()
								}
								ϒaccount = λ.Cal(λ.IntType, λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(λ.GetItem(ϒstream_data, λ.NewStr("azureAccount")), "replace", nil), λ.NewStr("nexxplayplus"), λ.NewStr("")), "replace", nil), λ.NewStr("nexxplayfb"), λ.NewStr("")))
								return λ.Mod(λ.NewStr("http://nx-%s%02d.akamaized.net/"), λ.NewTuple(
									ϒprefix,
									ϒaccount,
								))
							}
							return λ.None
						})
					ϒlanguage = func() λ.Object {
						if λv := λ.Cal(λ.GetAttr(λ.GetItem(ϒvideo, λ.NewStr("general")), "get", nil), λ.NewStr("language_raw")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.NewStr("")
						}
					}()
					ϒazure_stream_base = λ.Cal(ϒget_cdn_shield_base)
					ϒis_ml = λ.NewBool(λ.Contains(ϒlanguage, λ.NewStr(",")))
					ϒazure_manifest_url = λ.Add(λ.Mod(λ.NewStr("%s%s/%s_src%s.ism/Manifest"), λ.NewTuple(
						ϒazure_stream_base,
						ϒazure_locator,
						ϒvideo_id,
						func() λ.Object {
							if λ.IsTrue(ϒis_ml) {
								return λ.NewStr("_manifest")
							} else {
								return λ.NewStr("")
							}
						}(),
					)), λ.NewStr("%s"))
					ϒprotection_token = λ.Cal(ϒtry_get, ϒvideo, λ.NewFunction("<lambda>",
						[]λ.Param{
							{Name: "x"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒx = λargs[0]
							)
							return λ.GetItem(λ.GetItem(ϒx, λ.NewStr("protectiondata")), λ.NewStr("token"))
						}), ϒcompat_str)
					if λ.IsTrue(ϒprotection_token) {
						τmp0 = λ.IAdd(ϒazure_manifest_url, λ.Mod(λ.NewStr("?hdnts=%s"), ϒprotection_token))
						ϒazure_manifest_url = τmp0
					}
					ϒformats = λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
						λ.Mod(ϒazure_manifest_url, λ.NewStr("(format=m3u8-aapl)")),
						ϒvideo_id,
						λ.NewStr("mp4"),
						λ.NewStr("m3u8_native"),
					), λ.KWArgs{
						{Name: "m3u8_id", Value: λ.Mod(λ.NewStr("%s-hls"), ϒcdn)},
						{Name: "fatal", Value: λ.False},
					})
					λ.Cal(λ.GetAttr(ϒformats, "extend", nil), λ.Call(λ.GetAttr(ϒself, "_extract_mpd_formats", nil), λ.NewArgs(
						λ.Mod(ϒazure_manifest_url, λ.NewStr("(format=mpd-time-csf)")),
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "mpd_id", Value: λ.Mod(λ.NewStr("%s-dash"), ϒcdn)},
						{Name: "fatal", Value: λ.False},
					}))
					λ.Cal(λ.GetAttr(ϒformats, "extend", nil), λ.Call(λ.GetAttr(ϒself, "_extract_ism_formats", nil), λ.NewArgs(
						λ.Mod(ϒazure_manifest_url, λ.NewStr("")),
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "ism_id", Value: λ.Mod(λ.NewStr("%s-mss"), ϒcdn)},
						{Name: "fatal", Value: λ.False},
					}))
					ϒazure_progressive_base = λ.Cal(ϒget_cdn_shield_base, λ.NewStr("Prog"), λ.True)
					ϒazure_file_distribution = λ.Cal(λ.GetAttr(ϒstream_data, "get", nil), λ.NewStr("azureFileDistribution"))
					if λ.IsTrue(ϒazure_file_distribution) {
						ϒfds = λ.Cal(λ.GetAttr(ϒazure_file_distribution, "split", nil), λ.NewStr(","))
						if λ.IsTrue(ϒfds) {
							τmp0 = λ.Cal(λ.BuiltinIter, ϒfds)
							for {
								if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
									break
								}
								ϒfd = τmp1
								ϒss = λ.Cal(λ.GetAttr(ϒfd, "split", nil), λ.NewStr(":"))
								if λ.IsTrue(λ.Eq(λ.Cal(λ.BuiltinLen, ϒss), λ.NewInt(2))) {
									ϒtbr = λ.Cal(ϒint_or_none, λ.GetItem(ϒss, λ.NewInt(0)))
									if λ.IsTrue(ϒtbr) {
										ϒf = λ.NewDictWithTable(map[λ.Object]λ.Object{
											λ.NewStr("url"): λ.Mod(λ.NewStr("%s%s/%s_src_%s_%d.mp4"), λ.NewTuple(
												ϒazure_progressive_base,
												ϒazure_locator,
												ϒvideo_id,
												λ.GetItem(ϒss, λ.NewInt(1)),
												ϒtbr,
											)),
											λ.NewStr("format_id"): λ.Mod(λ.NewStr("%s-http-%d"), λ.NewTuple(
												ϒcdn,
												ϒtbr,
											)),
											λ.NewStr("tbr"): ϒtbr,
										})
										ϒwidth_height = λ.Cal(λ.GetAttr(λ.GetItem(ϒss, λ.NewInt(1)), "split", nil), λ.NewStr("x"))
										if λ.IsTrue(λ.Eq(λ.Cal(λ.BuiltinLen, ϒwidth_height), λ.NewInt(2))) {
											λ.Cal(λ.GetAttr(ϒf, "update", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
												λ.NewStr("width"):  λ.Cal(ϒint_or_none, λ.GetItem(ϒwidth_height, λ.NewInt(0))),
												λ.NewStr("height"): λ.Cal(ϒint_or_none, λ.GetItem(ϒwidth_height, λ.NewInt(1))),
											}))
										}
										λ.Cal(λ.GetAttr(ϒformats, "append", nil), ϒf)
									}
								}
							}
						}
					}
					return ϒformats
				})
			NexxIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒcdn           λ.Object
						ϒcid           λ.Object
						ϒdevice_id     λ.Object
						ϒdomain_id     λ.Object
						ϒformats       λ.Object
						ϒgeneral       λ.Object
						ϒmobj          λ.Object
						ϒop            λ.Object
						ϒrequest_token λ.Object
						ϒresponse      λ.Object
						ϒresult        λ.Object
						ϒsecret        λ.Object
						ϒself          = λargs[0]
						ϒtitle         λ.Object
						ϒurl           = λargs[1]
						ϒvideo         λ.Object
						ϒvideo_id      λ.Object
					)
					ϒmobj = λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl)
					ϒdomain_id = func() λ.Object {
						if λv := λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("domain_id")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("domain_id_s"))
						}
					}()
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("id"))
					ϒvideo = λ.None
					ϒresponse = λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
						λ.Mod(λ.NewStr("https://arc.nexx.cloud/api/video/%s.json"), ϒvideo_id),
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					})
					if λ.IsTrue(func() λ.Object {
						if λv := ϒresponse; !λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Cal(λ.BuiltinIsInstance, ϒresponse, λ.DictType)
						}
					}()) {
						ϒresult = λ.Cal(λ.GetAttr(ϒresponse, "get", nil), λ.NewStr("result"))
						if λ.IsTrue(func() λ.Object {
							if λv := ϒresult; !λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Cal(λ.BuiltinIsInstance, ϒresult, λ.DictType)
							}
						}()) {
							ϒvideo = ϒresult
						}
					}
					if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒvideo))) {
						ϒdevice_id = λ.Mod(λ.NewStr("%d:%d:%d%d"), λ.NewTuple(
							λ.Cal(λ.None, λ.NewInt(1), λ.NewInt(4)),
							λ.Cal(λ.IntType, λ.Cal(Ωtime.ϒtime)),
							λ.Cal(λ.None, λ.NewFloat(10000.0), λ.NewInt(99999)),
							λ.Cal(λ.None, λ.NewInt(1), λ.NewInt(9)),
						))
						ϒresult = λ.Call(λ.GetAttr(ϒself, "_call_api", nil), λ.NewArgs(
							ϒdomain_id,
							λ.NewStr("session/init"),
							ϒvideo_id,
						), λ.KWArgs{
							{Name: "data", Value: λ.NewDictWithTable(map[λ.Object]λ.Object{
								λ.NewStr("nxp_devh"):         ϒdevice_id,
								λ.NewStr("nxp_userh"):        λ.NewStr(""),
								λ.NewStr("precid"):           λ.NewStr("0"),
								λ.NewStr("playlicense"):      λ.NewStr("0"),
								λ.NewStr("screenx"):          λ.NewStr("1920"),
								λ.NewStr("screeny"):          λ.NewStr("1080"),
								λ.NewStr("playerversion"):    λ.NewStr("6.0.00"),
								λ.NewStr("gateway"):          λ.NewStr("html5"),
								λ.NewStr("adGateway"):        λ.NewStr(""),
								λ.NewStr("explicitlanguage"): λ.NewStr("en-US"),
								λ.NewStr("addTextTemplates"): λ.NewStr("1"),
								λ.NewStr("addDomainData"):    λ.NewStr("1"),
								λ.NewStr("addAdModel"):       λ.NewStr("1"),
							})},
							{Name: "headers", Value: λ.NewDictWithTable(map[λ.Object]λ.Object{
								λ.NewStr("X-Request-Enable-Auth-Fallback"): λ.NewStr("1"),
							})},
						})
						ϒcid = λ.GetItem(λ.GetItem(ϒresult, λ.NewStr("general")), λ.NewStr("cid"))
						ϒsecret = λ.GetItem(λ.GetItem(λ.GetItem(ϒresult, λ.NewStr("device")), λ.NewStr("clienttoken")), λ.NewSlice(λ.Cal(λ.IntType, λ.GetItem(ϒdevice_id, λ.NewInt(0))), λ.None, λ.None))
						ϒsecret = λ.GetItem(ϒsecret, λ.NewSlice(λ.NewInt(0), λ.Sub(λ.Cal(λ.BuiltinLen, ϒsecret), λ.Cal(λ.IntType, λ.GetItem(ϒdevice_id, λ.Neg(λ.NewInt(1))))), λ.None))
						ϒop = λ.NewStr("byid")
						ϒrequest_token = λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(λ.None, "md5", nil), λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(λ.NewStr(""), "join", nil), λ.NewTuple(
							ϒop,
							ϒdomain_id,
							ϒsecret,
						)), "encode", nil), λ.NewStr("utf-8"))), "hexdigest", nil))
						ϒvideo = λ.Call(λ.GetAttr(ϒself, "_call_api", nil), λ.NewArgs(
							ϒdomain_id,
							λ.Mod(λ.NewStr("videos/%s/%s"), λ.NewTuple(
								ϒop,
								ϒvideo_id,
							)),
							ϒvideo_id,
						), λ.KWArgs{
							{Name: "data", Value: λ.NewDictWithTable(map[λ.Object]λ.Object{
								λ.NewStr("additionalfields"):      λ.NewStr("language,channel,actors,studio,licenseby,slug,subtitle,teaser,description"),
								λ.NewStr("addInteractionOptions"): λ.NewStr("1"),
								λ.NewStr("addStatusDetails"):      λ.NewStr("1"),
								λ.NewStr("addStreamDetails"):      λ.NewStr("1"),
								λ.NewStr("addCaptions"):           λ.NewStr("1"),
								λ.NewStr("addScenes"):             λ.NewStr("1"),
								λ.NewStr("addHotSpots"):           λ.NewStr("1"),
								λ.NewStr("addBumpers"):            λ.NewStr("1"),
								λ.NewStr("captionFormat"):         λ.NewStr("data"),
							})},
							{Name: "headers", Value: λ.NewDictWithTable(map[λ.Object]λ.Object{
								λ.NewStr("X-Request-CID"):   ϒcid,
								λ.NewStr("X-Request-Token"): ϒrequest_token,
							})},
						})
					}
					ϒgeneral = λ.GetItem(ϒvideo, λ.NewStr("general"))
					ϒtitle = λ.GetItem(ϒgeneral, λ.NewStr("title"))
					ϒcdn = λ.GetItem(λ.GetItem(ϒvideo, λ.NewStr("streamdata")), λ.NewStr("cdnType"))
					if λ.IsTrue(λ.Eq(ϒcdn, λ.NewStr("azure"))) {
						ϒformats = λ.Cal(λ.GetAttr(ϒself, "_extract_azure_formats", nil), ϒvideo, ϒvideo_id)
					} else {
						if λ.IsTrue(λ.Eq(ϒcdn, λ.NewStr("free"))) {
							ϒformats = λ.Cal(λ.GetAttr(ϒself, "_extract_free_formats", nil), ϒvideo, ϒvideo_id)
						} else {
							if !λ.IsTrue(λ.False) {
								panic(λ.Raise(λ.Cal(λ.AssertionErrorType)))
							}
						}
					}
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):           ϒvideo_id,
						λ.NewStr("title"):        ϒtitle,
						λ.NewStr("alt_title"):    λ.Cal(λ.GetAttr(ϒgeneral, "get", nil), λ.NewStr("subtitle")),
						λ.NewStr("description"):  λ.Cal(λ.GetAttr(ϒgeneral, "get", nil), λ.NewStr("description")),
						λ.NewStr("release_year"): λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒgeneral, "get", nil), λ.NewStr("year"))),
						λ.NewStr("creator"): func() λ.Object {
							if λv := λ.Cal(λ.GetAttr(ϒgeneral, "get", nil), λ.NewStr("studio")); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Cal(λ.GetAttr(ϒgeneral, "get", nil), λ.NewStr("studio_adref"))
							}
						}(),
						λ.NewStr("thumbnail"): λ.Cal(ϒtry_get, ϒvideo, λ.NewFunction("<lambda>",
							[]λ.Param{
								{Name: "x"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒx = λargs[0]
								)
								return λ.GetItem(λ.GetItem(ϒx, λ.NewStr("imagedata")), λ.NewStr("thumb"))
							}), ϒcompat_str),
						λ.NewStr("duration"):  λ.Cal(ϒparse_duration, λ.Cal(λ.GetAttr(ϒgeneral, "get", nil), λ.NewStr("runtime"))),
						λ.NewStr("timestamp"): λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒgeneral, "get", nil), λ.NewStr("uploaded"))),
						λ.NewStr("episode_number"): λ.Cal(ϒint_or_none, λ.Cal(ϒtry_get, ϒvideo, λ.NewFunction("<lambda>",
							[]λ.Param{
								{Name: "x"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒx = λargs[0]
								)
								return λ.GetItem(λ.GetItem(ϒx, λ.NewStr("episodedata")), λ.NewStr("episode"))
							}))),
						λ.NewStr("season_number"): λ.Cal(ϒint_or_none, λ.Cal(ϒtry_get, ϒvideo, λ.NewFunction("<lambda>",
							[]λ.Param{
								{Name: "x"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒx = λargs[0]
								)
								return λ.GetItem(λ.GetItem(ϒx, λ.NewStr("episodedata")), λ.NewStr("season"))
							}))),
						λ.NewStr("formats"): ϒformats,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TESTS"):                 NexxIE__TESTS,
				λ.NewStr("_VALID_URL"):             NexxIE__VALID_URL,
				λ.NewStr("_extract_azure_formats"): NexxIE__extract_azure_formats,
				λ.NewStr("_extract_free_formats"):  NexxIE__extract_free_formats,
				λ.NewStr("_real_extract"):          NexxIE__real_extract,
			})
		}())
		NexxEmbedIE = λ.Cal(λ.TypeType, λ.NewStr("NexxEmbedIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				NexxEmbedIE__VALID_URL λ.Object
			)
			NexxEmbedIE__VALID_URL = λ.NewStr("https?://embed\\.nexx(?:\\.cloud|cdn\\.com)/\\d+/(?P<id>[^/?#&]+)")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): NexxEmbedIE__VALID_URL,
			})
		}())
	})
}