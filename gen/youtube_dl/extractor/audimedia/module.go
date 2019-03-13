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
 * audimedia/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/audimedia.py
 */

package audimedia

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	AudiMediaIE    λ.Object
	InfoExtractor  λ.Object
	ϒint_or_none   λ.Object
	ϒparse_iso8601 λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒparse_iso8601 = Ωutils.ϒparse_iso8601
		AudiMediaIE = λ.Cal(λ.TypeType, λ.NewStr("AudiMediaIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				AudiMediaIE__TESTS        λ.Object
				AudiMediaIE__VALID_URL    λ.Object
				AudiMediaIE__real_extract λ.Object
			)
			AudiMediaIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?audi-mediacenter\\.com/(?:en|de)/audimediatv/(?:video/)?(?P<id>[^/?#]+)")
			AudiMediaIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("https://www.audi-mediacenter.com/en/audimediatv/60-seconds-of-audi-sport-104-2015-wec-bahrain-rookie-test-1467"),
					λ.NewStr("md5"): λ.NewStr("79a8b71c46d49042609795ab59779b66"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("1565"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("60 Seconds of Audi Sport 104/2015 - WEC Bahrain, Rookie Test"),
						λ.NewStr("description"): λ.NewStr("md5:60e5d30a78ced725f7b8d34370762941"),
						λ.NewStr("upload_date"): λ.NewStr("20151124"),
						λ.NewStr("timestamp"):   λ.NewInt(1448354940),
						λ.NewStr("duration"):    λ.NewInt(74022),
						λ.NewStr("view_count"):  λ.IntType,
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://www.audi-mediacenter.com/en/audimediatv/video/60-seconds-of-audi-sport-104-2015-wec-bahrain-rookie-test-2991"),
					λ.NewStr("only_matching"): λ.True,
				}),
			)
			AudiMediaIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒbitrate           λ.Object
						ϒdisplay_id        λ.Object
						ϒf                 λ.Object
						ϒformats           λ.Object
						ϒraw_payload       λ.Object
						ϒself              = λargs[0]
						ϒstage_mode        λ.Object
						ϒstream_url_hds    λ.Object
						ϒstream_url_hls    λ.Object
						ϒurl               = λargs[1]
						ϒvideo_data        λ.Object
						ϒvideo_id          λ.Object
						ϒvideo_version     λ.Object
						ϒvideo_version_url λ.Object
						ϒwebpage           λ.Object
						τmp0               λ.Object
						τmp1               λ.Object
					)
					ϒdisplay_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒdisplay_id)
					ϒraw_payload = λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewList(
						λ.NewStr("class=\"amtv-embed\"[^>]+id=\"([0-9a-z-]+)\""),
						λ.NewStr("id=\"([0-9a-z-]+)\"[^>]+class=\"amtv-embed\""),
						λ.NewStr("class=\\\\\"amtv-embed\\\\\"[^>]+id=\\\\\"([0-9a-z-]+)\\\\\""),
						λ.NewStr("id=\\\\\"([0-9a-z-]+)\\\\\"[^>]+class=\\\\\"amtv-embed\\\\\""),
						λ.NewStr("id=(?:\\\\)?\"(amtve-[a-z]-\\d+-[a-z]{2})"),
					), ϒwebpage, λ.NewStr("raw payload"))
					τmp0 = λ.Cal(λ.GetAttr(ϒraw_payload, "split", nil), λ.NewStr("-"))
					_ = λ.GetItem(τmp0, λ.NewInt(0))
					ϒstage_mode = λ.GetItem(τmp0, λ.NewInt(1))
					ϒvideo_id = λ.GetItem(τmp0, λ.NewInt(2))
					_ = λ.GetItem(τmp0, λ.NewInt(3))
					if λ.IsTrue(λ.NewBool(!λ.Contains(λ.NewTuple(
						λ.NewStr("s"),
						λ.NewStr("e"),
					), ϒstage_mode))) {
						ϒvideo_data = λ.GetItem(λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
							λ.Add(λ.NewStr("https://www.audimedia.tv/api/video/v1/videos/"), ϒvideo_id),
							ϒvideo_id,
						), λ.KWArgs{
							{Name: "query", Value: λ.NewDictWithTable(map[λ.Object]λ.Object{
								λ.NewStr("embed[]"): λ.NewList(
									λ.NewStr("video_versions"),
									λ.NewStr("thumbnail_image"),
								),
							})},
						}), λ.NewStr("results"))
						ϒformats = λ.NewList()
						ϒstream_url_hls = λ.Cal(λ.GetAttr(ϒvideo_data, "get", nil), λ.NewStr("stream_url_hls"))
						if λ.IsTrue(ϒstream_url_hls) {
							λ.Cal(λ.GetAttr(ϒformats, "extend", nil), λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
								ϒstream_url_hls,
								ϒvideo_id,
								λ.NewStr("mp4"),
							), λ.KWArgs{
								{Name: "entry_protocol", Value: λ.NewStr("m3u8_native")},
								{Name: "m3u8_id", Value: λ.NewStr("hls")},
								{Name: "fatal", Value: λ.False},
							}))
						}
						ϒstream_url_hds = λ.Cal(λ.GetAttr(ϒvideo_data, "get", nil), λ.NewStr("stream_url_hds"))
						if λ.IsTrue(ϒstream_url_hds) {
							λ.Cal(λ.GetAttr(ϒformats, "extend", nil), λ.Call(λ.GetAttr(ϒself, "_extract_f4m_formats", nil), λ.NewArgs(
								λ.Add(ϒstream_url_hds, λ.NewStr("?hdcore=3.4.0")),
								ϒvideo_id,
							), λ.KWArgs{
								{Name: "f4m_id", Value: λ.NewStr("hds")},
								{Name: "fatal", Value: λ.False},
							}))
						}
						τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(ϒvideo_data, "get", nil), λ.NewStr("video_versions"), λ.NewList()))
						for {
							if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
								break
							}
							ϒvideo_version = τmp1
							ϒvideo_version_url = func() λ.Object {
								if λv := λ.Cal(λ.GetAttr(ϒvideo_version, "get", nil), λ.NewStr("download_url")); λ.IsTrue(λv) {
									return λv
								} else {
									return λ.Cal(λ.GetAttr(ϒvideo_version, "get", nil), λ.NewStr("stream_url"))
								}
							}()
							if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒvideo_version_url))) {
								continue
							}
							ϒf = λ.NewDictWithTable(map[λ.Object]λ.Object{
								λ.NewStr("url"):    ϒvideo_version_url,
								λ.NewStr("width"):  λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒvideo_version, "get", nil), λ.NewStr("width"))),
								λ.NewStr("height"): λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒvideo_version, "get", nil), λ.NewStr("height"))),
								λ.NewStr("abr"):    λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒvideo_version, "get", nil), λ.NewStr("audio_bitrate"))),
								λ.NewStr("vbr"):    λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒvideo_version, "get", nil), λ.NewStr("video_bitrate"))),
							})
							ϒbitrate = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
								λ.NewStr("(\\d+)k"),
								ϒvideo_version_url,
								λ.NewStr("bitrate"),
							), λ.KWArgs{
								{Name: "default", Value: λ.None},
							})
							if λ.IsTrue(ϒbitrate) {
								λ.Cal(λ.GetAttr(ϒf, "update", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
									λ.NewStr("format_id"): λ.Mod(λ.NewStr("http-%s"), ϒbitrate),
								}))
							}
							λ.Cal(λ.GetAttr(ϒformats, "append", nil), ϒf)
						}
						λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
						return λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("id"):          ϒvideo_id,
							λ.NewStr("title"):       λ.GetItem(ϒvideo_data, λ.NewStr("title")),
							λ.NewStr("description"): λ.Cal(λ.GetAttr(ϒvideo_data, "get", nil), λ.NewStr("subtitle")),
							λ.NewStr("thumbnail"):   λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒvideo_data, "get", nil), λ.NewStr("thumbnail_image"), λ.NewDictWithTable(map[λ.Object]λ.Object{})), "get", nil), λ.NewStr("file")),
							λ.NewStr("timestamp"):   λ.Cal(ϒparse_iso8601, λ.Cal(λ.GetAttr(ϒvideo_data, "get", nil), λ.NewStr("publication_date"))),
							λ.NewStr("duration"):    λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒvideo_data, "get", nil), λ.NewStr("duration"))),
							λ.NewStr("view_count"):  λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒvideo_data, "get", nil), λ.NewStr("view_count"))),
							λ.NewStr("formats"):     ϒformats,
						})
					}
					return λ.None
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TESTS"):        AudiMediaIE__TESTS,
				λ.NewStr("_VALID_URL"):    AudiMediaIE__VALID_URL,
				λ.NewStr("_real_extract"): AudiMediaIE__real_extract,
			})
		}())
	})
}
