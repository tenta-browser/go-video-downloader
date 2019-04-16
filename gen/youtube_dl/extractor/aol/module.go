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
 * aol/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/aol.py
 */

package aol

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	AolIE                         λ.Object
	ExtractorError                λ.Object
	InfoExtractor                 λ.Object
	ϒcompat_parse_qs              λ.Object
	ϒcompat_urllib_parse_urlparse λ.Object
	ϒint_or_none                  λ.Object
	ϒurl_or_none                  λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_parse_qs = Ωcompat.ϒcompat_parse_qs
		ϒcompat_urllib_parse_urlparse = Ωcompat.ϒcompat_urllib_parse_urlparse
		ExtractorError = Ωutils.ExtractorError
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒurl_or_none = Ωutils.ϒurl_or_none
		AolIE = λ.Cal(λ.TypeType, λ.NewStr("AolIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				AolIE_IE_NAME       λ.Object
				AolIE__TESTS        λ.Object
				AolIE__VALID_URL    λ.Object
				AolIE__real_extract λ.Object
			)
			AolIE_IE_NAME = λ.NewStr("aol.com")
			AolIE__VALID_URL = λ.NewStr("(?:aol-video:|https?://(?:www\\.)?aol\\.(?:com|ca|co\\.uk|de|jp)/video/(?:[^/]+/)*)(?P<id>[0-9a-f]+)")
			AolIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("https://www.aol.com/video/view/u-s--official-warns-of-largest-ever-irs-phone-scam/518167793/"),
					λ.NewStr("md5"): λ.NewStr("18ef68f48740e86ae94b98da815eec42"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("518167793"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("U.S. Official Warns Of 'Largest Ever' IRS Phone Scam"),
						λ.NewStr("description"): λ.NewStr("A major phone scam has cost thousands of taxpayers more than $1 million, with less than a month until income tax returns are due to the IRS."),
						λ.NewStr("timestamp"):   λ.NewInt(1395405060),
						λ.NewStr("upload_date"): λ.NewStr("20140321"),
						λ.NewStr("uploader"):    λ.NewStr("Newsy Studio"),
					}),
					λ.NewStr("params"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("skip_download"): λ.True,
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("https://www.aol.com/video/view/netflix-is-raising-rates/5707d6b8e4b090497b04f706/"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("5707d6b8e4b090497b04f706"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Netflix is Raising Rates"),
						λ.NewStr("description"): λ.NewStr("Netflix is rewarding millions of it’s long-standing members with an increase in cost. Veuer’s Carly Figueroa has more."),
						λ.NewStr("upload_date"): λ.NewStr("20160408"),
						λ.NewStr("timestamp"):   λ.NewInt(1460123280),
						λ.NewStr("uploader"):    λ.NewStr("Veuer"),
					}),
					λ.NewStr("params"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("skip_download"): λ.True,
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://www.aol.com/video/view/park-bench-season-2-trailer/559a1b9be4b0c3bfad3357a7/"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://www.aol.com/video/view/donald-trump-spokeswoman-tones-down-megyn-kelly-attacks/519442220/"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("aol-video:5707d6b8e4b090497b04f706"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://www.aol.com/video/playlist/PL8245/5ca79d19d21f1a04035db606/"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://www.aol.ca/video/view/u-s-woman-s-family-arrested-for-murder-first-pinned-on-panhandler-police/5c7ccf45bc03931fa04b2fe1/"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://www.aol.co.uk/video/view/-one-dead-and-22-hurt-in-bus-crash-/5cb3a6f3d21f1a072b457347/"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://www.aol.de/video/view/eva-braun-privataufnahmen-von-hitlers-geliebter-werden-digitalisiert/5cb2d49de98ab54c113d3d5d/"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://www.aol.jp/video/playlist/5a28e936a1334d000137da0c/5a28f3151e642219fde19831/"),
					λ.NewStr("only_matching"): λ.True,
				}),
			)
			AolIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒext        λ.Object
						ϒf          λ.Object
						ϒformats    λ.Object
						ϒm3u8_url   λ.Object
						ϒmobj       λ.Object
						ϒqs         λ.Object
						ϒrendition  λ.Object
						ϒresponse   λ.Object
						ϒself       = λargs[0]
						ϒurl        = λargs[1]
						ϒvideo_data λ.Object
						ϒvideo_id   λ.Object
						ϒvideo_url  λ.Object
						τmp0        λ.Object
						τmp1        λ.Object
					)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒresponse = λ.GetItem(λ.Cal(λ.GetAttr(ϒself, "_download_json", nil), λ.Mod(λ.NewStr("https://feedapi.b2c.on.aol.com/v1.0/app/videos/aolon/%s/details"), ϒvideo_id), ϒvideo_id), λ.NewStr("response"))
					if λ.IsTrue(λ.Ne(λ.GetItem(ϒresponse, λ.NewStr("statusText")), λ.NewStr("Ok"))) {
						panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.Mod(λ.NewStr("%s said: %s"), λ.NewTuple(
							λ.GetAttr(ϒself, "IE_NAME", nil),
							λ.GetItem(ϒresponse, λ.NewStr("statusText")),
						))), λ.KWArgs{
							{Name: "expected", Value: λ.True},
						})))
					}
					ϒvideo_data = λ.GetItem(ϒresponse, λ.NewStr("data"))
					ϒformats = λ.NewList()
					ϒm3u8_url = λ.Cal(ϒurl_or_none, λ.Cal(λ.GetAttr(ϒvideo_data, "get", nil), λ.NewStr("videoMasterPlaylist")))
					if λ.IsTrue(ϒm3u8_url) {
						λ.Cal(λ.GetAttr(ϒformats, "extend", nil), λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
							ϒm3u8_url,
							ϒvideo_id,
							λ.NewStr("mp4"),
						), λ.KWArgs{
							{Name: "m3u8_id", Value: λ.NewStr("hls")},
							{Name: "fatal", Value: λ.False},
						}))
					}
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(ϒvideo_data, "get", nil), λ.NewStr("renditions"), λ.NewList()))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒrendition = τmp1
						ϒvideo_url = λ.Cal(ϒurl_or_none, λ.Cal(λ.GetAttr(ϒrendition, "get", nil), λ.NewStr("url")))
						if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒvideo_url))) {
							continue
						}
						ϒext = λ.Cal(λ.GetAttr(ϒrendition, "get", nil), λ.NewStr("format"))
						if λ.IsTrue(λ.Eq(ϒext, λ.NewStr("m3u8"))) {
							λ.Cal(λ.GetAttr(ϒformats, "extend", nil), λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
								ϒvideo_url,
								ϒvideo_id,
								λ.NewStr("mp4"),
							), λ.KWArgs{
								{Name: "m3u8_id", Value: λ.NewStr("hls")},
								{Name: "fatal", Value: λ.False},
							}))
						} else {
							ϒf = λ.NewDictWithTable(map[λ.Object]λ.Object{
								λ.NewStr("url"):       ϒvideo_url,
								λ.NewStr("format_id"): λ.Cal(λ.GetAttr(ϒrendition, "get", nil), λ.NewStr("quality")),
							})
							ϒmobj = λ.Cal(Ωre.ϒsearch, λ.NewStr("(\\d+)x(\\d+)"), ϒvideo_url)
							if λ.IsTrue(ϒmobj) {
								λ.Cal(λ.GetAttr(ϒf, "update", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
									λ.NewStr("width"):  λ.Cal(λ.IntType, λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewInt(1))),
									λ.NewStr("height"): λ.Cal(λ.IntType, λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewInt(2))),
								}))
							} else {
								ϒqs = λ.Cal(ϒcompat_parse_qs, λ.GetAttr(λ.Cal(ϒcompat_urllib_parse_urlparse, ϒvideo_url), "query", nil))
								λ.Cal(λ.GetAttr(ϒf, "update", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
									λ.NewStr("width"):  λ.Cal(ϒint_or_none, λ.GetItem(λ.Cal(λ.GetAttr(ϒqs, "get", nil), λ.NewStr("w"), λ.NewList(λ.None)), λ.NewInt(0))),
									λ.NewStr("height"): λ.Cal(ϒint_or_none, λ.GetItem(λ.Cal(λ.GetAttr(ϒqs, "get", nil), λ.NewStr("h"), λ.NewList(λ.None)), λ.NewInt(0))),
								}))
							}
							λ.Cal(λ.GetAttr(ϒformats, "append", nil), ϒf)
						}
					}
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats, λ.NewTuple(
						λ.NewStr("width"),
						λ.NewStr("height"),
						λ.NewStr("tbr"),
						λ.NewStr("format_id"),
					))
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          ϒvideo_id,
						λ.NewStr("title"):       λ.GetItem(ϒvideo_data, λ.NewStr("title")),
						λ.NewStr("duration"):    λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒvideo_data, "get", nil), λ.NewStr("duration"))),
						λ.NewStr("timestamp"):   λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒvideo_data, "get", nil), λ.NewStr("publishDate"))),
						λ.NewStr("view_count"):  λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒvideo_data, "get", nil), λ.NewStr("views"))),
						λ.NewStr("description"): λ.Cal(λ.GetAttr(ϒvideo_data, "get", nil), λ.NewStr("description")),
						λ.NewStr("uploader"):    λ.Cal(λ.GetAttr(ϒvideo_data, "get", nil), λ.NewStr("videoOwner")),
						λ.NewStr("formats"):     ϒformats,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("IE_NAME"):       AolIE_IE_NAME,
				λ.NewStr("_TESTS"):        AolIE__TESTS,
				λ.NewStr("_VALID_URL"):    AolIE__VALID_URL,
				λ.NewStr("_real_extract"): AolIE__real_extract,
			})
		}())
	})
}
