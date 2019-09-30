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
 * yinyuetai/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/yinyuetai.py
 */

package yinyuetai

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError λ.Object
	InfoExtractor  λ.Object
	YinYueTaiIE    λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ExtractorError = Ωutils.ExtractorError
		YinYueTaiIE = λ.Cal(λ.TypeType, λ.NewStr("YinYueTaiIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				YinYueTaiIE_IE_NAME       λ.Object
				YinYueTaiIE__TESTS        λ.Object
				YinYueTaiIE__VALID_URL    λ.Object
				YinYueTaiIE__real_extract λ.Object
			)
			YinYueTaiIE_IE_NAME = λ.NewStr("yinyuetai:video")
			YinYueTaiIE__VALID_URL = λ.NewStr("https?://v\\.yinyuetai\\.com/video(?:/h5)?/(?P<id>[0-9]+)")
			YinYueTaiIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://v.yinyuetai.com/video/2322376"),
					λ.NewStr("md5"): λ.NewStr("6e3abe28d38e3a54b591f9f040595ce0"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):        λ.NewStr("2322376"),
						λ.NewStr("ext"):       λ.NewStr("mp4"),
						λ.NewStr("title"):     λ.NewStr("少女时代_PARTY_Music Video Teaser"),
						λ.NewStr("creator"):   λ.NewStr("少女时代"),
						λ.NewStr("duration"):  λ.NewInt(25),
						λ.NewStr("thumbnail"): λ.NewStr("re:^https?://.*\\.jpg$"),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://v.yinyuetai.com/video/h5/2322376"),
					λ.NewStr("only_matching"): λ.True,
				}),
			)
			YinYueTaiIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒformats  λ.Object
						ϒinfo     λ.Object
						ϒself     = λargs[0]
						ϒurl      = λargs[1]
						ϒvideo_id λ.Object
					)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒinfo = λ.GetItem(λ.GetItem(λ.Cal(λ.GetAttr(ϒself, "_download_json", nil), λ.Mod(λ.NewStr("http://ext.yinyuetai.com/main/get-h-mv-info?json=true&videoId=%s"), ϒvideo_id), ϒvideo_id, λ.NewStr("Downloading mv info")), λ.NewStr("videoInfo")), λ.NewStr("coreVideoInfo"))
					if λ.IsTrue(λ.GetItem(ϒinfo, λ.NewStr("error"))) {
						panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.GetItem(ϒinfo, λ.NewStr("errorMsg"))), λ.KWArgs{
							{Name: "expected", Value: λ.True},
						})))
					}
					ϒformats = λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
						nil,
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							return λ.NewGenerator(func(λgy λ.Yielder) λ.Object {
								var (
									ϒformat_info λ.Object
									τmp0         λ.Object
									τmp1         λ.Object
								)
								τmp0 = λ.Cal(λ.BuiltinIter, λ.GetItem(ϒinfo, λ.NewStr("videoUrlModels")))
								for {
									if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
										break
									}
									ϒformat_info = τmp1
									λgy.Yield(λ.NewDictWithTable(map[λ.Object]λ.Object{
										λ.NewStr("url"):       λ.GetItem(ϒformat_info, λ.NewStr("videoUrl")),
										λ.NewStr("format_id"): λ.GetItem(ϒformat_info, λ.NewStr("qualityLevel")),
										λ.NewStr("format"):    λ.Cal(λ.GetAttr(ϒformat_info, "get", nil), λ.NewStr("qualityLevelName")),
										λ.NewStr("filesize"):  λ.Cal(λ.GetAttr(ϒformat_info, "get", nil), λ.NewStr("fileSize")),
										λ.NewStr("ext"):       λ.NewStr("mp4"),
										λ.NewStr("tbr"):       λ.Cal(λ.GetAttr(ϒformat_info, "get", nil), λ.NewStr("bitrate")),
									}))
								}
								return λ.None
							})
						})))
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):        ϒvideo_id,
						λ.NewStr("title"):     λ.GetItem(ϒinfo, λ.NewStr("videoName")),
						λ.NewStr("thumbnail"): λ.Cal(λ.GetAttr(ϒinfo, "get", nil), λ.NewStr("bigHeadImage")),
						λ.NewStr("creator"):   λ.Cal(λ.GetAttr(ϒinfo, "get", nil), λ.NewStr("artistNames")),
						λ.NewStr("duration"):  λ.Cal(λ.GetAttr(ϒinfo, "get", nil), λ.NewStr("duration")),
						λ.NewStr("formats"):   ϒformats,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("IE_NAME"):       YinYueTaiIE_IE_NAME,
				λ.NewStr("_TESTS"):        YinYueTaiIE__TESTS,
				λ.NewStr("_VALID_URL"):    YinYueTaiIE__VALID_URL,
				λ.NewStr("_real_extract"): YinYueTaiIE__real_extract,
			})
		}())
	})
}
