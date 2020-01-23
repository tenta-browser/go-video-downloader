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
		YinYueTaiIE = λ.Cal(λ.TypeType, λ.StrLiteral("YinYueTaiIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				YinYueTaiIE_IE_NAME       λ.Object
				YinYueTaiIE__VALID_URL    λ.Object
				YinYueTaiIE__real_extract λ.Object
			)
			YinYueTaiIE_IE_NAME = λ.StrLiteral("yinyuetai:video")
			YinYueTaiIE__VALID_URL = λ.StrLiteral("https?://v\\.yinyuetai\\.com/video(?:/h5)?/(?P<id>[0-9]+)")
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
					ϒvideo_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒinfo = λ.GetItem(λ.GetItem(λ.Calm(ϒself, "_download_json", λ.Mod(λ.StrLiteral("http://ext.yinyuetai.com/main/get-h-mv-info?json=true&videoId=%s"), ϒvideo_id), ϒvideo_id, λ.StrLiteral("Downloading mv info")), λ.StrLiteral("videoInfo")), λ.StrLiteral("coreVideoInfo"))
					if λ.IsTrue(λ.GetItem(ϒinfo, λ.StrLiteral("error"))) {
						panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.GetItem(ϒinfo, λ.StrLiteral("errorMsg"))), λ.KWArgs{
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
								τmp0 = λ.Cal(λ.BuiltinIter, λ.GetItem(ϒinfo, λ.StrLiteral("videoUrlModels")))
								for {
									if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
										break
									}
									ϒformat_info = τmp1
									λgy.Yield(λ.DictLiteral(map[string]λ.Object{
										"url":       λ.GetItem(ϒformat_info, λ.StrLiteral("videoUrl")),
										"format_id": λ.GetItem(ϒformat_info, λ.StrLiteral("qualityLevel")),
										"format":    λ.Calm(ϒformat_info, "get", λ.StrLiteral("qualityLevelName")),
										"filesize":  λ.Calm(ϒformat_info, "get", λ.StrLiteral("fileSize")),
										"ext":       λ.StrLiteral("mp4"),
										"tbr":       λ.Calm(ϒformat_info, "get", λ.StrLiteral("bitrate")),
									}))
								}
								return λ.None
							})
						})))
					λ.Calm(ϒself, "_sort_formats", ϒformats)
					return λ.DictLiteral(map[string]λ.Object{
						"id":        ϒvideo_id,
						"title":     λ.GetItem(ϒinfo, λ.StrLiteral("videoName")),
						"thumbnail": λ.Calm(ϒinfo, "get", λ.StrLiteral("bigHeadImage")),
						"creator":   λ.Calm(ϒinfo, "get", λ.StrLiteral("artistNames")),
						"duration":  λ.Calm(ϒinfo, "get", λ.StrLiteral("duration")),
						"formats":   ϒformats,
					})
				})
			return λ.DictLiteral(map[string]λ.Object{
				"IE_NAME":       YinYueTaiIE_IE_NAME,
				"_VALID_URL":    YinYueTaiIE__VALID_URL,
				"_real_extract": YinYueTaiIE__real_extract,
			})
		}())
	})
}
