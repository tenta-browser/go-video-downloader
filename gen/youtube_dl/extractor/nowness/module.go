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
 * nowness/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/nowness.py
 */

package nowness

import (
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωbrightcove "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/brightcove"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	BrightcoveLegacyIE λ.Object
	BrightcoveNewIE    λ.Object
	ExtractorError     λ.Object
	InfoExtractor      λ.Object
	NownessBaseIE      λ.Object
	NownessIE          λ.Object
	NownessPlaylistIE  λ.Object
	NownessSeriesIE    λ.Object
	ϒcompat_str        λ.Object
	ϒsanitized_Request λ.Object
)

func init() {
	λ.InitModule(func() {
		BrightcoveLegacyIE = Ωbrightcove.BrightcoveLegacyIE
		BrightcoveNewIE = Ωbrightcove.BrightcoveNewIE
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_str = Ωcompat.ϒcompat_str
		ExtractorError = Ωutils.ExtractorError
		ϒsanitized_Request = Ωutils.ϒsanitized_Request
		NownessBaseIE = λ.Cal(λ.TypeType, λ.NewStr("NownessBaseIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				NownessBaseIE__api_request        λ.Object
				NownessBaseIE__extract_url_result λ.Object
			)
			NownessBaseIE__extract_url_result = λ.NewFunction("_extract_url_result",
				[]λ.Param{
					{Name: "self"},
					{Name: "post"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒbc_url      λ.Object
						ϒmedia       λ.Object
						ϒplayer_code λ.Object
						ϒpost        = λargs[1]
						ϒself        = λargs[0]
						ϒsource      λ.Object
						ϒvideo_id    λ.Object
						τmp0         λ.Object
						τmp1         λ.Object
					)
					if λ.IsTrue(λ.Eq(λ.GetItem(ϒpost, λ.NewStr("type")), λ.NewStr("video"))) {
						τmp0 = λ.Cal(λ.BuiltinIter, λ.GetItem(ϒpost, λ.NewStr("media")))
						for {
							if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
								break
							}
							ϒmedia = τmp1
							if λ.IsTrue(λ.Eq(λ.GetItem(ϒmedia, λ.NewStr("type")), λ.NewStr("video"))) {
								ϒvideo_id = λ.GetItem(ϒmedia, λ.NewStr("content"))
								ϒsource = λ.GetItem(ϒmedia, λ.NewStr("source"))
								if λ.IsTrue(λ.Eq(ϒsource, λ.NewStr("brightcove"))) {
									ϒplayer_code = λ.Call(λ.GetAttr(ϒself, "_download_webpage", nil), λ.NewArgs(
										λ.Mod(λ.NewStr("http://www.nowness.com/iframe?id=%s"), ϒvideo_id),
										ϒvideo_id,
									), λ.KWArgs{
										{Name: "note", Value: λ.NewStr("Downloading player JavaScript")},
										{Name: "errnote", Value: λ.NewStr("Unable to download player JavaScript")},
									})
									ϒbc_url = λ.Cal(λ.GetAttr(BrightcoveLegacyIE, "_extract_brightcove_url", nil), ϒplayer_code)
									if λ.IsTrue(ϒbc_url) {
										return λ.Cal(λ.GetAttr(ϒself, "url_result", nil), ϒbc_url, λ.Cal(λ.GetAttr(BrightcoveLegacyIE, "ie_key", nil)))
									}
									ϒbc_url = λ.Cal(λ.GetAttr(BrightcoveNewIE, "_extract_url", nil), ϒself, ϒplayer_code)
									if λ.IsTrue(ϒbc_url) {
										return λ.Cal(λ.GetAttr(ϒself, "url_result", nil), ϒbc_url, λ.Cal(λ.GetAttr(BrightcoveNewIE, "ie_key", nil)))
									}
									panic(λ.Raise(λ.Cal(ExtractorError, λ.NewStr("Could not find player definition"))))
								} else {
									if λ.IsTrue(λ.Eq(ϒsource, λ.NewStr("vimeo"))) {
										return λ.Cal(λ.GetAttr(ϒself, "url_result", nil), λ.Mod(λ.NewStr("http://vimeo.com/%s"), ϒvideo_id), λ.NewStr("Vimeo"))
									} else {
										if λ.IsTrue(λ.Eq(ϒsource, λ.NewStr("youtube"))) {
											return λ.Cal(λ.GetAttr(ϒself, "url_result", nil), ϒvideo_id, λ.NewStr("Youtube"))
										} else {
											if λ.IsTrue(λ.Eq(ϒsource, λ.NewStr("cinematique"))) {
												// pass
											}
										}
									}
								}
							}
						}
					}
					return λ.None
				})
			NownessBaseIE__api_request = λ.NewFunction("_api_request",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
					{Name: "request_path"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdisplay_id   λ.Object
						ϒrequest      λ.Object
						ϒrequest_path = λargs[2]
						ϒself         = λargs[0]
						ϒurl          = λargs[1]
					)
					ϒdisplay_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒrequest = λ.Call(ϒsanitized_Request, λ.NewArgs(λ.Add(λ.NewStr("http://api.nowness.com/api/"), λ.Mod(ϒrequest_path, ϒdisplay_id))), λ.KWArgs{
						{Name: "headers", Value: λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("X-Nowness-Language"): func() λ.Object {
								if λ.IsTrue(λ.NewBool(λ.Contains(ϒurl, λ.NewStr("cn.nowness.com")))) {
									return λ.NewStr("zh-cn")
								} else {
									return λ.NewStr("en-us")
								}
							}(),
						})},
					})
					return λ.NewTuple(
						ϒdisplay_id,
						λ.Cal(λ.GetAttr(ϒself, "_download_json", nil), ϒrequest, ϒdisplay_id),
					)
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_api_request"):        NownessBaseIE__api_request,
				λ.NewStr("_extract_url_result"): NownessBaseIE__extract_url_result,
			})
		}())
		NownessIE = λ.Cal(λ.TypeType, λ.NewStr("NownessIE"), λ.NewTuple(NownessBaseIE), func() λ.Dict {
			var (
				NownessIE_IE_NAME       λ.Object
				NownessIE__TESTS        λ.Object
				NownessIE__VALID_URL    λ.Object
				NownessIE__real_extract λ.Object
			)
			NownessIE_IE_NAME = λ.NewStr("nowness")
			NownessIE__VALID_URL = λ.NewStr("https?://(?:(?:www|cn)\\.)?nowness\\.com/(?:story|(?:series|category)/[^/]+)/(?P<id>[^/]+?)(?:$|[?#])")
			NownessIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("https://www.nowness.com/story/candor-the-art-of-gesticulation"),
					λ.NewStr("md5"): λ.NewStr("068bc0202558c2e391924cb8cc470676"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("2520295746001"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Candor: The Art of Gesticulation"),
						λ.NewStr("description"): λ.NewStr("Candor: The Art of Gesticulation"),
						λ.NewStr("thumbnail"):   λ.NewStr("re:^https?://.*\\.jpg"),
						λ.NewStr("timestamp"):   λ.NewInt(1446745676),
						λ.NewStr("upload_date"): λ.NewStr("20151105"),
						λ.NewStr("uploader_id"): λ.NewStr("2385340575001"),
					}),
					λ.NewStr("add_ie"): λ.NewList(λ.NewStr("BrightcoveNew")),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("https://cn.nowness.com/story/kasper-bjorke-ft-jaakko-eino-kalevi-tnr"),
					λ.NewStr("md5"): λ.NewStr("e79cf125e387216f86b2e0a5b5c63aa3"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("3716354522001"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Kasper Bjørke ft. Jaakko Eino Kalevi: TNR"),
						λ.NewStr("description"): λ.NewStr("Kasper Bjørke ft. Jaakko Eino Kalevi: TNR"),
						λ.NewStr("thumbnail"):   λ.NewStr("re:^https?://.*\\.jpg"),
						λ.NewStr("timestamp"):   λ.NewInt(1407315371),
						λ.NewStr("upload_date"): λ.NewStr("20140806"),
						λ.NewStr("uploader_id"): λ.NewStr("2385340575001"),
					}),
					λ.NewStr("add_ie"): λ.NewList(λ.NewStr("BrightcoveNew")),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("https://www.nowness.com/series/nowness-picks/jean-luc-godard-supercut"),
					λ.NewStr("md5"): λ.NewStr("9a5a6a8edf806407e411296ab6bc2a49"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("130020913"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Bleu, Blanc, Rouge - A Godard Supercut"),
						λ.NewStr("description"): λ.NewStr("md5:f0ea5f1857dffca02dbd37875d742cec"),
						λ.NewStr("thumbnail"):   λ.NewStr("re:^https?://.*\\.jpg"),
						λ.NewStr("upload_date"): λ.NewStr("20150607"),
						λ.NewStr("uploader"):    λ.NewStr("Cinema Sem Lei"),
						λ.NewStr("uploader_id"): λ.NewStr("cinemasemlei"),
					}),
					λ.NewStr("add_ie"): λ.NewList(λ.NewStr("Vimeo")),
				}),
			)
			NownessIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒpost λ.Object
						ϒself = λargs[0]
						ϒurl  = λargs[1]
						τmp0  λ.Object
					)
					τmp0 = λ.Cal(λ.GetAttr(ϒself, "_api_request", nil), ϒurl, λ.NewStr("post/getBySlug/%s"))
					_ = λ.GetItem(τmp0, λ.NewInt(0))
					ϒpost = λ.GetItem(τmp0, λ.NewInt(1))
					return λ.Cal(λ.GetAttr(ϒself, "_extract_url_result", nil), ϒpost)
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("IE_NAME"):       NownessIE_IE_NAME,
				λ.NewStr("_TESTS"):        NownessIE__TESTS,
				λ.NewStr("_VALID_URL"):    NownessIE__VALID_URL,
				λ.NewStr("_real_extract"): NownessIE__real_extract,
			})
		}())
		NownessPlaylistIE = λ.Cal(λ.TypeType, λ.NewStr("NownessPlaylistIE"), λ.NewTuple(NownessBaseIE), func() λ.Dict {
			var (
				NownessPlaylistIE__VALID_URL λ.Object
			)
			NownessPlaylistIE__VALID_URL = λ.NewStr("https?://(?:(?:www|cn)\\.)?nowness\\.com/playlist/(?P<id>\\d+)")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): NownessPlaylistIE__VALID_URL,
			})
		}())
		NownessSeriesIE = λ.Cal(λ.TypeType, λ.NewStr("NownessSeriesIE"), λ.NewTuple(NownessBaseIE), func() λ.Dict {
			var (
				NownessSeriesIE__VALID_URL λ.Object
			)
			NownessSeriesIE__VALID_URL = λ.NewStr("https?://(?:(?:www|cn)\\.)?nowness\\.com/series/(?P<id>[^/]+?)(?:$|[?#])")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): NownessSeriesIE__VALID_URL,
			})
		}())
	})
}
