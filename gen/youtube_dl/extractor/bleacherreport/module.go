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
 * bleacherreport/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/bleacherreport.py
 */

package bleacherreport

import (
	Ωamp "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/amp"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	AMPIE               λ.Object
	BleacherReportCMSIE λ.Object
	BleacherReportIE    λ.Object
	ExtractorError      λ.Object
	InfoExtractor       λ.Object
	ϒint_or_none        λ.Object
	ϒparse_iso8601      λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		AMPIE = Ωamp.AMPIE
		ExtractorError = Ωutils.ExtractorError
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒparse_iso8601 = Ωutils.ϒparse_iso8601
		BleacherReportIE = λ.Cal(λ.TypeType, λ.NewStr("BleacherReportIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				BleacherReportIE__VALID_URL λ.Object
			)
			BleacherReportIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?bleacherreport\\.com/articles/(?P<id>\\d+)")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): BleacherReportIE__VALID_URL,
			})
		}())
		BleacherReportCMSIE = λ.Cal(λ.TypeType, λ.NewStr("BleacherReportCMSIE"), λ.NewTuple(AMPIE), func() λ.Dict {
			var (
				BleacherReportCMSIE__TESTS        λ.Object
				BleacherReportCMSIE__VALID_URL    λ.Object
				BleacherReportCMSIE__real_extract λ.Object
			)
			BleacherReportCMSIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?bleacherreport\\.com/video_embed\\?id=(?P<id>[0-9a-f-]{36})")
			BleacherReportCMSIE__TESTS = λ.NewList(λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("url"): λ.NewStr("http://bleacherreport.com/video_embed?id=8fd44c2f-3dc5-4821-9118-2c825a98c0e1"),
				λ.NewStr("md5"): λ.NewStr("2e4b0a997f9228ffa31fada5c53d1ed1"),
				λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("id"):          λ.NewStr("8fd44c2f-3dc5-4821-9118-2c825a98c0e1"),
					λ.NewStr("ext"):         λ.NewStr("flv"),
					λ.NewStr("title"):       λ.NewStr("Cena vs. Rollins Would Expose the Heavyweight Division"),
					λ.NewStr("description"): λ.NewStr("md5:984afb4ade2f9c0db35f3267ed88b36e"),
				}),
			}))
			BleacherReportCMSIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒinfo     λ.Object
						ϒself     = λargs[0]
						ϒurl      = λargs[1]
						ϒvideo_id λ.Object
					)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒinfo = λ.Cal(λ.GetAttr(ϒself, "_extract_feed_info", nil), λ.Mod(λ.NewStr("http://cms.bleacherreport.com/media/items/%s/akamai.json"), ϒvideo_id))
					λ.SetItem(ϒinfo, λ.NewStr("id"), ϒvideo_id)
					return ϒinfo
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TESTS"):        BleacherReportCMSIE__TESTS,
				λ.NewStr("_VALID_URL"):    BleacherReportCMSIE__VALID_URL,
				λ.NewStr("_real_extract"): BleacherReportCMSIE__real_extract,
			})
		}())
	})
}
