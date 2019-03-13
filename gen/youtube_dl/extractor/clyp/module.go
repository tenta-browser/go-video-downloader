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
 * clyp/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/clyp.py
 */

package clyp

import (
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ClypIE                        λ.Object
	InfoExtractor                 λ.Object
	ϒcompat_parse_qs              λ.Object
	ϒcompat_urllib_parse_urlparse λ.Object
	ϒfloat_or_none                λ.Object
	ϒunified_timestamp            λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_parse_qs = Ωcompat.ϒcompat_parse_qs
		ϒcompat_urllib_parse_urlparse = Ωcompat.ϒcompat_urllib_parse_urlparse
		ϒfloat_or_none = Ωutils.ϒfloat_or_none
		ϒunified_timestamp = Ωutils.ϒunified_timestamp
		ClypIE = λ.Cal(λ.TypeType, λ.NewStr("ClypIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				ClypIE__TESTS        λ.Object
				ClypIE__VALID_URL    λ.Object
				ClypIE__real_extract λ.Object
			)
			ClypIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?clyp\\.it/(?P<id>[a-z0-9]+)")
			ClypIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("https://clyp.it/ojz2wfah"),
					λ.NewStr("md5"): λ.NewStr("1d4961036c41247ecfdcc439c0cddcbb"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("ojz2wfah"),
						λ.NewStr("ext"):         λ.NewStr("mp3"),
						λ.NewStr("title"):       λ.NewStr("Krisson80 - bits wip wip"),
						λ.NewStr("description"): λ.NewStr("#Krisson80BitsWipWip #chiptune\n#wip"),
						λ.NewStr("duration"):    λ.NewFloat(263.21),
						λ.NewStr("timestamp"):   λ.NewInt(1443515251),
						λ.NewStr("upload_date"): λ.NewStr("20150929"),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("https://clyp.it/b04p1odi?token=b0078e077e15835845c528a44417719d"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("b04p1odi"),
						λ.NewStr("ext"):         λ.NewStr("mp3"),
						λ.NewStr("title"):       λ.NewStr("GJ! (Reward Edit)"),
						λ.NewStr("description"): λ.NewStr("Metal Resistance (THE ONE edition)"),
						λ.NewStr("duration"):    λ.NewFloat(177.789),
						λ.NewStr("timestamp"):   λ.NewInt(1528241278),
						λ.NewStr("upload_date"): λ.NewStr("20180605"),
					}),
					λ.NewStr("params"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("skip_download"): λ.True,
					}),
				}),
			)
			ClypIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒaudio_id    λ.Object
						ϒdescription λ.Object
						ϒduration    λ.Object
						ϒext         λ.Object
						ϒformat_id   λ.Object
						ϒformat_url  λ.Object
						ϒformats     λ.Object
						ϒmetadata    λ.Object
						ϒqs          λ.Object
						ϒquery       λ.Object
						ϒsecure      λ.Object
						ϒself        = λargs[0]
						ϒtimestamp   λ.Object
						ϒtitle       λ.Object
						ϒtoken       λ.Object
						ϒurl         = λargs[1]
						τmp0         λ.Object
						τmp1         λ.Object
						τmp2         λ.Object
						τmp3         λ.Object
					)
					ϒaudio_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒqs = λ.Cal(ϒcompat_parse_qs, λ.GetAttr(λ.Cal(ϒcompat_urllib_parse_urlparse, ϒurl), "query", nil))
					ϒtoken = λ.GetItem(λ.Cal(λ.GetAttr(ϒqs, "get", nil), λ.NewStr("token"), λ.NewList(λ.None)), λ.NewInt(0))
					ϒquery = λ.NewDictWithTable(map[λ.Object]λ.Object{})
					if λ.IsTrue(ϒtoken) {
						λ.SetItem(ϒquery, λ.NewStr("token"), ϒtoken)
					}
					ϒmetadata = λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
						λ.Mod(λ.NewStr("https://api.clyp.it/%s"), ϒaudio_id),
						ϒaudio_id,
					), λ.KWArgs{
						{Name: "query", Value: ϒquery},
					})
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.NewTuple(
						λ.NewStr(""),
						λ.NewStr("Secure"),
					))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒsecure = τmp1
						τmp2 = λ.Cal(λ.BuiltinIter, λ.NewTuple(
							λ.NewStr("Ogg"),
							λ.NewStr("Mp3"),
						))
						for {
							if τmp3 = λ.NextDefault(τmp2, λ.AfterLast); τmp3 == λ.AfterLast {
								break
							}
							ϒext = τmp3
							ϒformat_id = λ.Mod(λ.NewStr("%s%s"), λ.NewTuple(
								ϒsecure,
								ϒext,
							))
							ϒformat_url = λ.Cal(λ.GetAttr(ϒmetadata, "get", nil), λ.Mod(λ.NewStr("%sUrl"), ϒformat_id))
							if λ.IsTrue(ϒformat_url) {
								λ.Cal(λ.GetAttr(ϒformats, "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
									λ.NewStr("url"):       ϒformat_url,
									λ.NewStr("format_id"): ϒformat_id,
									λ.NewStr("vcodec"):    λ.NewStr("none"),
								}))
							}
						}
					}
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					ϒtitle = λ.GetItem(ϒmetadata, λ.NewStr("Title"))
					ϒdescription = λ.Cal(λ.GetAttr(ϒmetadata, "get", nil), λ.NewStr("Description"))
					ϒduration = λ.Cal(ϒfloat_or_none, λ.Cal(λ.GetAttr(ϒmetadata, "get", nil), λ.NewStr("Duration")))
					ϒtimestamp = λ.Cal(ϒunified_timestamp, λ.Cal(λ.GetAttr(ϒmetadata, "get", nil), λ.NewStr("DateCreated")))
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          ϒaudio_id,
						λ.NewStr("title"):       ϒtitle,
						λ.NewStr("description"): ϒdescription,
						λ.NewStr("duration"):    ϒduration,
						λ.NewStr("timestamp"):   ϒtimestamp,
						λ.NewStr("formats"):     ϒformats,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TESTS"):        ClypIE__TESTS,
				λ.NewStr("_VALID_URL"):    ClypIE__VALID_URL,
				λ.NewStr("_real_extract"): ClypIE__real_extract,
			})
		}())
	})
}
