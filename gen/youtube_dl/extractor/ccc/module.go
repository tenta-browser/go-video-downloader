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
 * ccc/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/ccc.py
 */

package ccc

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	CCCIE          λ.Object
	InfoExtractor  λ.Object
	ϒint_or_none   λ.Object
	ϒparse_iso8601 λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒparse_iso8601 = Ωutils.ϒparse_iso8601
		CCCIE = λ.Cal(λ.TypeType, λ.NewStr("CCCIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				CCCIE_IE_NAME       λ.Object
				CCCIE__TESTS        λ.Object
				CCCIE__VALID_URL    λ.Object
				CCCIE__real_extract λ.Object
			)
			CCCIE_IE_NAME = λ.NewStr("media.ccc.de")
			CCCIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?media\\.ccc\\.de/v/(?P<id>[^/?#&]+)")
			CCCIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("https://media.ccc.de/v/30C3_-_5443_-_en_-_saal_g_-_201312281830_-_introduction_to_processor_design_-_byterazor#video"),
					λ.NewStr("md5"): λ.NewStr("3a1eda8f3a29515d27f5adb967d7e740"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("1839"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Introduction to Processor Design"),
						λ.NewStr("description"): λ.NewStr("md5:df55f6d073d4ceae55aae6f2fd98a0ac"),
						λ.NewStr("thumbnail"):   λ.NewStr("re:^https?://.*\\.jpg$"),
						λ.NewStr("upload_date"): λ.NewStr("20131228"),
						λ.NewStr("timestamp"):   λ.NewInt(1388188800),
						λ.NewStr("duration"):    λ.NewInt(3710),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://media.ccc.de/v/32c3-7368-shopshifting#download"),
					λ.NewStr("only_matching"): λ.True,
				}),
			)
			CCCIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdisplay_id    λ.Object
						ϒevent_data    λ.Object
						ϒevent_id      λ.Object
						ϒfolder        λ.Object
						ϒformat_id     λ.Object
						ϒformats       λ.Object
						ϒlanguage      λ.Object
						ϒrecording     λ.Object
						ϒrecording_url λ.Object
						ϒself          = λargs[0]
						ϒurl           = λargs[1]
						ϒvcodec        λ.Object
						ϒwebpage       λ.Object
						τmp0           λ.Object
						τmp1           λ.Object
						τmp2           λ.Object
					)
					ϒdisplay_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒdisplay_id)
					ϒevent_id = λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewStr("data-id='(\\d+)'"), ϒwebpage, λ.NewStr("event id"))
					ϒevent_data = λ.Cal(λ.GetAttr(ϒself, "_download_json", nil), λ.Mod(λ.NewStr("https://media.ccc.de/public/events/%s"), ϒevent_id), ϒevent_id)
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(ϒevent_data, "get", nil), λ.NewStr("recordings"), λ.NewList()))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒrecording = τmp1
						ϒrecording_url = λ.Cal(λ.GetAttr(ϒrecording, "get", nil), λ.NewStr("recording_url"))
						if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒrecording_url))) {
							continue
						}
						ϒlanguage = λ.Cal(λ.GetAttr(ϒrecording, "get", nil), λ.NewStr("language"))
						ϒfolder = λ.Cal(λ.GetAttr(ϒrecording, "get", nil), λ.NewStr("folder"))
						ϒformat_id = λ.None
						if λ.IsTrue(ϒlanguage) {
							ϒformat_id = ϒlanguage
						}
						if λ.IsTrue(ϒfolder) {
							if λ.IsTrue(ϒlanguage) {
								τmp2 = λ.IAdd(ϒformat_id, λ.Add(λ.NewStr("-"), ϒfolder))
								ϒformat_id = τmp2
							} else {
								ϒformat_id = ϒfolder
							}
						}
						ϒvcodec = func() λ.Object {
							if λ.IsTrue(λ.NewBool(λ.Contains(ϒfolder, λ.NewStr("h264")))) {
								return λ.NewStr("h264")
							} else {
								return func() λ.Object {
									if λ.IsTrue(λ.NewBool(λ.Contains(λ.NewTuple(
										λ.NewStr("mp3"),
										λ.NewStr("opus"),
									), ϒfolder))) {
										return λ.NewStr("none")
									} else {
										return λ.None
									}
								}()
							}
						}()
						λ.Cal(λ.GetAttr(ϒformats, "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("format_id"): ϒformat_id,
							λ.NewStr("url"):       ϒrecording_url,
							λ.NewStr("width"):     λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒrecording, "get", nil), λ.NewStr("width"))),
							λ.NewStr("height"):    λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒrecording, "get", nil), λ.NewStr("height"))),
							λ.NewStr("filesize"): λ.Call(ϒint_or_none, λ.NewArgs(λ.Cal(λ.GetAttr(ϒrecording, "get", nil), λ.NewStr("size"))), λ.KWArgs{
								{Name: "invscale", Value: λ.Mul(λ.NewInt(1024), λ.NewInt(1024))},
							}),
							λ.NewStr("language"): ϒlanguage,
							λ.NewStr("vcodec"):   ϒvcodec,
						}))
					}
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          ϒevent_id,
						λ.NewStr("display_id"):  ϒdisplay_id,
						λ.NewStr("title"):       λ.GetItem(ϒevent_data, λ.NewStr("title")),
						λ.NewStr("description"): λ.Cal(λ.GetAttr(ϒevent_data, "get", nil), λ.NewStr("description")),
						λ.NewStr("thumbnail"):   λ.Cal(λ.GetAttr(ϒevent_data, "get", nil), λ.NewStr("thumb_url")),
						λ.NewStr("timestamp"):   λ.Cal(ϒparse_iso8601, λ.Cal(λ.GetAttr(ϒevent_data, "get", nil), λ.NewStr("date"))),
						λ.NewStr("duration"):    λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒevent_data, "get", nil), λ.NewStr("length"))),
						λ.NewStr("tags"):        λ.Cal(λ.GetAttr(ϒevent_data, "get", nil), λ.NewStr("tags")),
						λ.NewStr("formats"):     ϒformats,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("IE_NAME"):       CCCIE_IE_NAME,
				λ.NewStr("_TESTS"):        CCCIE__TESTS,
				λ.NewStr("_VALID_URL"):    CCCIE__VALID_URL,
				λ.NewStr("_real_extract"): CCCIE__real_extract,
			})
		}())
	})
}
