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
 * mojvideo/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/mojvideo.py
 */

package mojvideo

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError  λ.Object
	InfoExtractor   λ.Object
	MojvideoIE      λ.Object
	ϒparse_duration λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ExtractorError = Ωutils.ExtractorError
		ϒparse_duration = Ωutils.ϒparse_duration
		MojvideoIE = λ.Cal(λ.TypeType, λ.NewStr("MojvideoIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				MojvideoIE__TEST         λ.Object
				MojvideoIE__VALID_URL    λ.Object
				MojvideoIE__real_extract λ.Object
			)
			MojvideoIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?mojvideo\\.com/video-(?P<display_id>[^/]+)/(?P<id>[a-f0-9]+)")
			MojvideoIE__TEST = λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("url"): λ.NewStr("http://www.mojvideo.com/video-v-avtu-pred-mano-rdecelaska-alfi-nipic/3d1ed4497707730b2906"),
				λ.NewStr("md5"): λ.NewStr("f7fd662cc8ce2be107b0d4f2c0483ae7"),
				λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("id"):         λ.NewStr("3d1ed4497707730b2906"),
					λ.NewStr("display_id"): λ.NewStr("v-avtu-pred-mano-rdecelaska-alfi-nipic"),
					λ.NewStr("ext"):        λ.NewStr("mp4"),
					λ.NewStr("title"):      λ.NewStr("V avtu pred mano rdečelaska - Alfi Nipič"),
					λ.NewStr("thumbnail"):  λ.NewStr("re:^http://.*\\.jpg$"),
					λ.NewStr("duration"):   λ.NewInt(242),
				}),
			})
			MojvideoIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdisplay_id λ.Object
						ϒduration   λ.Object
						ϒerror_desc λ.Object
						ϒmobj       λ.Object
						ϒplayerapi  λ.Object
						ϒself       = λargs[0]
						ϒthumbnail  λ.Object
						ϒtitle      λ.Object
						ϒurl        = λargs[1]
						ϒvideo_id   λ.Object
						ϒvideo_url  λ.Object
					)
					ϒmobj = λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("id"))
					ϒdisplay_id = λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("display_id"))
					ϒplayerapi = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), λ.Mod(λ.NewStr("http://www.mojvideo.com/playerapi.php?v=%s&t=1"), ϒvideo_id), ϒdisplay_id)
					if λ.IsTrue(λ.NewBool(λ.Contains(ϒplayerapi, λ.NewStr("<error>true</error>")))) {
						ϒerror_desc = λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
							λ.NewStr("<errordesc>([^<]*)</errordesc>"),
							ϒplayerapi,
							λ.NewStr("error description"),
						), λ.KWArgs{
							{Name: "fatal", Value: λ.False},
						})
						panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.Mod(λ.NewStr("%s said: %s"), λ.NewTuple(
							λ.GetAttr(ϒself, "IE_NAME", nil),
							ϒerror_desc,
						))), λ.KWArgs{
							{Name: "expected", Value: λ.True},
						})))
					}
					ϒtitle = λ.Cal(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewStr("<title>([^<]+)</title>"), ϒplayerapi, λ.NewStr("title"))
					ϒvideo_url = λ.Cal(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewStr("<file>([^<]+)</file>"), ϒplayerapi, λ.NewStr("video URL"))
					ϒthumbnail = λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
						λ.NewStr("<preview>([^<]+)</preview>"),
						ϒplayerapi,
						λ.NewStr("thumbnail"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					})
					ϒduration = λ.Cal(ϒparse_duration, λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
						λ.NewStr("<duration>([^<]+)</duration>"),
						ϒplayerapi,
						λ.NewStr("duration"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					}))
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):         ϒvideo_id,
						λ.NewStr("display_id"): ϒdisplay_id,
						λ.NewStr("url"):        ϒvideo_url,
						λ.NewStr("title"):      ϒtitle,
						λ.NewStr("thumbnail"):  ϒthumbnail,
						λ.NewStr("duration"):   ϒduration,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TEST"):         MojvideoIE__TEST,
				λ.NewStr("_VALID_URL"):    MojvideoIE__VALID_URL,
				λ.NewStr("_real_extract"): MojvideoIE__real_extract,
			})
		}())
	})
}
