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
 * senateisvp/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/senateisvp.py
 */

package senateisvp

import (
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
	SenateISVPIE     λ.Object
	ϒcompat_parse_qs λ.Object
	ϒunsmuggle_url   λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ExtractorError = Ωutils.ExtractorError
		ϒunsmuggle_url = Ωutils.ϒunsmuggle_url
		ϒcompat_parse_qs = Ωcompat.ϒcompat_parse_qs
		SenateISVPIE = λ.Cal(λ.TypeType, λ.NewStr("SenateISVPIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				SenateISVPIE__COMM_MAP          λ.Object
				SenateISVPIE__TESTS             λ.Object
				SenateISVPIE__VALID_URL         λ.Object
				SenateISVPIE__get_info_for_comm λ.Object
				SenateISVPIE__real_extract      λ.Object
				SenateISVPIE__search_iframe_url λ.Object
			)
			SenateISVPIE__COMM_MAP = λ.NewList(
				λ.NewList(
					λ.NewStr("ag"),
					λ.NewStr("76440"),
					λ.NewStr("http://ag-f.akamaihd.net"),
				),
				λ.NewList(
					λ.NewStr("aging"),
					λ.NewStr("76442"),
					λ.NewStr("http://aging-f.akamaihd.net"),
				),
				λ.NewList(
					λ.NewStr("approps"),
					λ.NewStr("76441"),
					λ.NewStr("http://approps-f.akamaihd.net"),
				),
				λ.NewList(
					λ.NewStr("armed"),
					λ.NewStr("76445"),
					λ.NewStr("http://armed-f.akamaihd.net"),
				),
				λ.NewList(
					λ.NewStr("banking"),
					λ.NewStr("76446"),
					λ.NewStr("http://banking-f.akamaihd.net"),
				),
				λ.NewList(
					λ.NewStr("budget"),
					λ.NewStr("76447"),
					λ.NewStr("http://budget-f.akamaihd.net"),
				),
				λ.NewList(
					λ.NewStr("cecc"),
					λ.NewStr("76486"),
					λ.NewStr("http://srs-f.akamaihd.net"),
				),
				λ.NewList(
					λ.NewStr("commerce"),
					λ.NewStr("80177"),
					λ.NewStr("http://commerce1-f.akamaihd.net"),
				),
				λ.NewList(
					λ.NewStr("csce"),
					λ.NewStr("75229"),
					λ.NewStr("http://srs-f.akamaihd.net"),
				),
				λ.NewList(
					λ.NewStr("dpc"),
					λ.NewStr("76590"),
					λ.NewStr("http://dpc-f.akamaihd.net"),
				),
				λ.NewList(
					λ.NewStr("energy"),
					λ.NewStr("76448"),
					λ.NewStr("http://energy-f.akamaihd.net"),
				),
				λ.NewList(
					λ.NewStr("epw"),
					λ.NewStr("76478"),
					λ.NewStr("http://epw-f.akamaihd.net"),
				),
				λ.NewList(
					λ.NewStr("ethics"),
					λ.NewStr("76449"),
					λ.NewStr("http://ethics-f.akamaihd.net"),
				),
				λ.NewList(
					λ.NewStr("finance"),
					λ.NewStr("76450"),
					λ.NewStr("http://finance-f.akamaihd.net"),
				),
				λ.NewList(
					λ.NewStr("foreign"),
					λ.NewStr("76451"),
					λ.NewStr("http://foreign-f.akamaihd.net"),
				),
				λ.NewList(
					λ.NewStr("govtaff"),
					λ.NewStr("76453"),
					λ.NewStr("http://govtaff-f.akamaihd.net"),
				),
				λ.NewList(
					λ.NewStr("help"),
					λ.NewStr("76452"),
					λ.NewStr("http://help-f.akamaihd.net"),
				),
				λ.NewList(
					λ.NewStr("indian"),
					λ.NewStr("76455"),
					λ.NewStr("http://indian-f.akamaihd.net"),
				),
				λ.NewList(
					λ.NewStr("intel"),
					λ.NewStr("76456"),
					λ.NewStr("http://intel-f.akamaihd.net"),
				),
				λ.NewList(
					λ.NewStr("intlnarc"),
					λ.NewStr("76457"),
					λ.NewStr("http://intlnarc-f.akamaihd.net"),
				),
				λ.NewList(
					λ.NewStr("jccic"),
					λ.NewStr("85180"),
					λ.NewStr("http://jccic-f.akamaihd.net"),
				),
				λ.NewList(
					λ.NewStr("jec"),
					λ.NewStr("76458"),
					λ.NewStr("http://jec-f.akamaihd.net"),
				),
				λ.NewList(
					λ.NewStr("judiciary"),
					λ.NewStr("76459"),
					λ.NewStr("http://judiciary-f.akamaihd.net"),
				),
				λ.NewList(
					λ.NewStr("rpc"),
					λ.NewStr("76591"),
					λ.NewStr("http://rpc-f.akamaihd.net"),
				),
				λ.NewList(
					λ.NewStr("rules"),
					λ.NewStr("76460"),
					λ.NewStr("http://rules-f.akamaihd.net"),
				),
				λ.NewList(
					λ.NewStr("saa"),
					λ.NewStr("76489"),
					λ.NewStr("http://srs-f.akamaihd.net"),
				),
				λ.NewList(
					λ.NewStr("smbiz"),
					λ.NewStr("76461"),
					λ.NewStr("http://smbiz-f.akamaihd.net"),
				),
				λ.NewList(
					λ.NewStr("srs"),
					λ.NewStr("75229"),
					λ.NewStr("http://srs-f.akamaihd.net"),
				),
				λ.NewList(
					λ.NewStr("uscc"),
					λ.NewStr("76487"),
					λ.NewStr("http://srs-f.akamaihd.net"),
				),
				λ.NewList(
					λ.NewStr("vetaff"),
					λ.NewStr("76462"),
					λ.NewStr("http://vetaff-f.akamaihd.net"),
				),
				λ.NewList(
					λ.NewStr("arch"),
					λ.NewStr(""),
					λ.NewStr("http://ussenate-f.akamaihd.net/"),
				),
			)
			SenateISVPIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?senate\\.gov/isvp/?\\?(?P<qs>.+)")
			SenateISVPIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.senate.gov/isvp/?comm=judiciary&type=live&stt=&filename=judiciary031715&auto_play=false&wmode=transparent&poster=http%3A%2F%2Fwww.judiciary.senate.gov%2Fthemes%2Fjudiciary%2Fimages%2Fvideo-poster-flash-fit.png"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):        λ.NewStr("judiciary031715"),
						λ.NewStr("ext"):       λ.NewStr("mp4"),
						λ.NewStr("title"):     λ.NewStr("Integrated Senate Video Player"),
						λ.NewStr("thumbnail"): λ.NewStr("re:^https?://.*\\.(?:jpg|png)$"),
					}),
					λ.NewStr("params"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("skip_download"): λ.True,
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.senate.gov/isvp/?type=live&comm=commerce&filename=commerce011514.mp4&auto_play=false"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):    λ.NewStr("commerce011514"),
						λ.NewStr("ext"):   λ.NewStr("mp4"),
						λ.NewStr("title"): λ.NewStr("Integrated Senate Video Player"),
					}),
					λ.NewStr("params"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("skip_download"): λ.True,
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.senate.gov/isvp/?type=arch&comm=intel&filename=intel090613&hc_location=ufi"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):    λ.NewStr("intel090613"),
						λ.NewStr("ext"):   λ.NewStr("mp4"),
						λ.NewStr("title"): λ.NewStr("Integrated Senate Video Player"),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://www.senate.gov/isvp?type=live&comm=banking&filename=banking012715"),
					λ.NewStr("only_matching"): λ.True,
				}),
			)
			SenateISVPIE__search_iframe_url = λ.NewFunction("_search_iframe_url",
				[]λ.Param{
					{Name: "webpage"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒmobj    λ.Object
						ϒwebpage = λargs[0]
					)
					ϒmobj = λ.Cal(Ωre.ϒsearch, λ.NewStr("<iframe[^>]+src=['\\\"](?P<url>https?://www\\.senate\\.gov/isvp/?\\?[^'\\\"]+)['\\\"]"), ϒwebpage)
					if λ.IsTrue(ϒmobj) {
						return λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("url"))
					}
					return λ.None
				})
			SenateISVPIE__search_iframe_url = λ.Cal(λ.StaticMethodType, SenateISVPIE__search_iframe_url)
			SenateISVPIE__get_info_for_comm = λ.NewFunction("_get_info_for_comm",
				[]λ.Param{
					{Name: "self"},
					{Name: "committee"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒcommittee = λargs[1]
						ϒentry     λ.Object
						ϒself      = λargs[0]
						τmp0       λ.Object
						τmp1       λ.Object
					)
					τmp0 = λ.Cal(λ.BuiltinIter, λ.GetAttr(ϒself, "_COMM_MAP", nil))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒentry = τmp1
						if λ.IsTrue(λ.Eq(λ.GetItem(ϒentry, λ.NewInt(0)), ϒcommittee)) {
							return λ.GetItem(ϒentry, λ.NewSlice(λ.NewInt(1), λ.None, λ.None))
						}
					}
					return λ.None
				})
			SenateISVPIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒcommittee     λ.Object
						ϒdomain        λ.Object
						ϒentry         λ.Object
						ϒf4m_url       λ.Object
						ϒfilename      λ.Object
						ϒformats       λ.Object
						ϒhdcore_sign   λ.Object
						ϒm3u8_url      λ.Object
						ϒmobj          λ.Object
						ϒposter        λ.Object
						ϒqs            λ.Object
						ϒself          = λargs[0]
						ϒsmuggled_data λ.Object
						ϒstream_num    λ.Object
						ϒthumbnail     λ.Object
						ϒtitle         λ.Object
						ϒurl           = λargs[1]
						ϒurl_params    λ.Object
						ϒvideo_id      λ.Object
						ϒvideo_type    λ.Object
						ϒwebpage       λ.Object
						τmp0           λ.Object
						τmp1           λ.Object
						τmp2           λ.Object
					)
					τmp0 = λ.Cal(ϒunsmuggle_url, ϒurl, λ.NewDictWithTable(map[λ.Object]λ.Object{}))
					ϒurl = λ.GetItem(τmp0, λ.NewInt(0))
					ϒsmuggled_data = λ.GetItem(τmp0, λ.NewInt(1))
					ϒqs = λ.Cal(ϒcompat_parse_qs, λ.Cal(λ.GetAttr(λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl), "group", nil), λ.NewStr("qs")))
					if λ.IsTrue(func() λ.Object {
						if λv := λ.NewBool(!λ.IsTrue(λ.Cal(λ.GetAttr(ϒqs, "get", nil), λ.NewStr("filename")))); λ.IsTrue(λv) {
							return λv
						} else if λv := λ.NewBool(!λ.IsTrue(λ.Cal(λ.GetAttr(ϒqs, "get", nil), λ.NewStr("type")))); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.NewBool(!λ.IsTrue(λ.Cal(λ.GetAttr(ϒqs, "get", nil), λ.NewStr("comm"))))
						}
					}()) {
						panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.NewStr("Invalid URL")), λ.KWArgs{
							{Name: "expected", Value: λ.True},
						})))
					}
					ϒvideo_id = λ.Cal(Ωre.ϒsub, λ.NewStr(".mp4$"), λ.NewStr(""), λ.GetItem(λ.GetItem(ϒqs, λ.NewStr("filename")), λ.NewInt(0)))
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒvideo_id)
					if λ.IsTrue(λ.Cal(λ.GetAttr(ϒsmuggled_data, "get", nil), λ.NewStr("force_title"))) {
						ϒtitle = λ.GetItem(ϒsmuggled_data, λ.NewStr("force_title"))
					} else {
						ϒtitle = λ.Cal(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewStr("<title>([^<]+)</title>"), ϒwebpage, ϒvideo_id)
					}
					ϒposter = λ.Cal(λ.GetAttr(ϒqs, "get", nil), λ.NewStr("poster"))
					ϒthumbnail = func() λ.Object {
						if λ.IsTrue(ϒposter) {
							return λ.GetItem(ϒposter, λ.NewInt(0))
						} else {
							return λ.None
						}
					}()
					ϒvideo_type = λ.GetItem(λ.GetItem(ϒqs, λ.NewStr("type")), λ.NewInt(0))
					ϒcommittee = func() λ.Object {
						if λ.IsTrue(λ.Eq(ϒvideo_type, λ.NewStr("arch"))) {
							return ϒvideo_type
						} else {
							return λ.GetItem(λ.GetItem(ϒqs, λ.NewStr("comm")), λ.NewInt(0))
						}
					}()
					τmp0 = λ.Cal(λ.GetAttr(ϒself, "_get_info_for_comm", nil), ϒcommittee)
					ϒstream_num = λ.GetItem(τmp0, λ.NewInt(0))
					ϒdomain = λ.GetItem(τmp0, λ.NewInt(1))
					ϒformats = λ.NewList()
					if λ.IsTrue(λ.Eq(ϒvideo_type, λ.NewStr("arch"))) {
						ϒfilename = func() λ.Object {
							if λ.IsTrue(λ.NewBool(λ.Contains(ϒvideo_id, λ.NewStr(".")))) {
								return ϒvideo_id
							} else {
								return λ.Add(ϒvideo_id, λ.NewStr(".mp4"))
							}
						}()
						ϒformats = λ.NewList(λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("url"): λ.Add(λ.Cal(Ωparse.ϒurljoin, ϒdomain, ϒfilename), λ.NewStr("?v=3.1.0&fp=&r=&g=")),
						}))
					} else {
						ϒhdcore_sign = λ.NewStr("hdcore=3.1.0")
						ϒurl_params = λ.NewTuple(
							ϒdomain,
							ϒvideo_id,
							ϒstream_num,
						)
						ϒf4m_url = λ.Add(λ.Mod(λ.NewStr("%s/z/%s_1@%s/manifest.f4m?"), ϒurl_params), ϒhdcore_sign)
						ϒm3u8_url = λ.Mod(λ.NewStr("%s/i/%s_1@%s/master.m3u8"), ϒurl_params)
						τmp0 = λ.Cal(λ.BuiltinIter, λ.Call(λ.GetAttr(ϒself, "_extract_f4m_formats", nil), λ.NewArgs(
							ϒf4m_url,
							ϒvideo_id,
						), λ.KWArgs{
							{Name: "f4m_id", Value: λ.NewStr("f4m")},
						}))
						for {
							if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
								break
							}
							ϒentry = τmp1
							λ.Cal(λ.GetAttr(ϒentry, "update", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
								λ.NewStr("extra_param_to_segment_url"): ϒhdcore_sign,
							}))
							λ.Cal(λ.GetAttr(ϒformats, "append", nil), ϒentry)
						}
						τmp0 = λ.Cal(λ.BuiltinIter, λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
							ϒm3u8_url,
							ϒvideo_id,
						), λ.KWArgs{
							{Name: "ext", Value: λ.NewStr("mp4")},
							{Name: "m3u8_id", Value: λ.NewStr("m3u8")},
						}))
						for {
							if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
								break
							}
							ϒentry = τmp1
							ϒmobj = λ.Cal(Ωre.ϒsearch, λ.NewStr("(?P<tag>(?:-p|-b)).m3u8"), λ.GetItem(ϒentry, λ.NewStr("url")))
							if λ.IsTrue(ϒmobj) {
								τmp2 = λ.IAdd(λ.GetItem(ϒentry, λ.NewStr("format_id")), λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("tag")))
								λ.SetItem(ϒentry, λ.NewStr("format_id"), τmp2)
							}
							λ.Cal(λ.GetAttr(ϒformats, "append", nil), ϒentry)
						}
						λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					}
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):        ϒvideo_id,
						λ.NewStr("title"):     ϒtitle,
						λ.NewStr("formats"):   ϒformats,
						λ.NewStr("thumbnail"): ϒthumbnail,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_COMM_MAP"):          SenateISVPIE__COMM_MAP,
				λ.NewStr("_TESTS"):             SenateISVPIE__TESTS,
				λ.NewStr("_VALID_URL"):         SenateISVPIE__VALID_URL,
				λ.NewStr("_get_info_for_comm"): SenateISVPIE__get_info_for_comm,
				λ.NewStr("_real_extract"):      SenateISVPIE__real_extract,
				λ.NewStr("_search_iframe_url"): SenateISVPIE__search_iframe_url,
			})
		}())
	})
}
