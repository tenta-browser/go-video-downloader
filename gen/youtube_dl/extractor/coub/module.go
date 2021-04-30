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
 * coub/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/coub.py
 */

package coub

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	CoubIE         λ.Object
	ExtractorError λ.Object
	InfoExtractor  λ.Object
	ϒfloat_or_none λ.Object
	ϒint_or_none   λ.Object
	ϒparse_iso8601 λ.Object
	ϒqualities     λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ExtractorError = Ωutils.ExtractorError
		ϒfloat_or_none = Ωutils.ϒfloat_or_none
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒparse_iso8601 = Ωutils.ϒparse_iso8601
		ϒqualities = Ωutils.ϒqualities
		CoubIE = λ.Cal(λ.TypeType, λ.StrLiteral("CoubIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				CoubIE__VALID_URL    λ.Object
				CoubIE__real_extract λ.Object
			)
			CoubIE__VALID_URL = λ.StrLiteral("(?:coub:|https?://(?:coub\\.com/(?:view|embed|coubs)/|c-cdn\\.coub\\.com/fb-player\\.swf\\?.*\\bcoub(?:ID|id)=))(?P<id>[\\da-z]+)")
			CoubIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						HTML5             λ.Object
						IPHONE            λ.Object
						MOBILE            λ.Object
						QUALITIES         λ.Object
						SOURCE_PREFERENCE λ.Object
						ϒage_limit        λ.Object
						ϒage_restricted   λ.Object
						ϒcoub             λ.Object
						ϒduration         λ.Object
						ϒfile_versions    λ.Object
						ϒformats          λ.Object
						ϒiphone_url       λ.Object
						ϒitem             λ.Object
						ϒitem_url         λ.Object
						ϒitems            λ.Object
						ϒkind             λ.Object
						ϒlike_count       λ.Object
						ϒmobile_url       λ.Object
						ϒpreference_key   λ.Object
						ϒquality          λ.Object
						ϒquality_key      λ.Object
						ϒrepost_count     λ.Object
						ϒself             = λargs[0]
						ϒthumbnail        λ.Object
						ϒtimestamp        λ.Object
						ϒtitle            λ.Object
						ϒuploader         λ.Object
						ϒuploader_id      λ.Object
						ϒurl              = λargs[1]
						ϒvideo_id         λ.Object
						ϒview_count       λ.Object
						τmp0              λ.Object
						τmp1              λ.Object
						τmp2              λ.Object
						τmp3              λ.Object
						τmp4              λ.Object
					)
					ϒvideo_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒcoub = λ.Calm(ϒself, "_download_json", λ.Mod(λ.StrLiteral("http://coub.com/api/v2/coubs/%s.json"), ϒvideo_id), ϒvideo_id)
					if λ.IsTrue(λ.Calm(ϒcoub, "get", λ.StrLiteral("error"))) {
						panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.Mod(λ.StrLiteral("%s said: %s"), λ.NewTuple(
							λ.GetAttr(ϒself, "IE_NAME", nil),
							λ.GetItem(ϒcoub, λ.StrLiteral("error")),
						))), λ.KWArgs{
							{Name: "expected", Value: λ.True},
						})))
					}
					ϒtitle = λ.GetItem(ϒcoub, λ.StrLiteral("title"))
					ϒfile_versions = λ.GetItem(ϒcoub, λ.StrLiteral("file_versions"))
					QUALITIES = λ.NewTuple(
						λ.StrLiteral("low"),
						λ.StrLiteral("med"),
						λ.StrLiteral("high"),
					)
					MOBILE = λ.StrLiteral("mobile")
					IPHONE = λ.StrLiteral("iphone")
					HTML5 = λ.StrLiteral("html5")
					SOURCE_PREFERENCE = λ.NewTuple(
						MOBILE,
						IPHONE,
						HTML5,
					)
					ϒquality_key = λ.Cal(ϒqualities, QUALITIES)
					ϒpreference_key = λ.Cal(ϒqualities, SOURCE_PREFERENCE)
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(λ.Calm(ϒfile_versions, "get", HTML5, λ.DictLiteral(map[λ.Object]λ.Object{})), "items"))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						τmp2 = λ.UnpackIterable(τmp1, 2)
						ϒkind = λ.GetItem(τmp2, λ.IntLiteral(0))
						ϒitems = λ.GetItem(τmp2, λ.IntLiteral(1))
						if !λ.Contains(λ.NewTuple(
							λ.StrLiteral("video"),
							λ.StrLiteral("audio"),
						), ϒkind) {
							continue
						}
						if !λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒitems, λ.DictType)) {
							continue
						}
						τmp2 = λ.Cal(λ.BuiltinIter, λ.Calm(ϒitems, "items"))
						for {
							if τmp3 = λ.NextDefault(τmp2, λ.AfterLast); τmp3 == λ.AfterLast {
								break
							}
							τmp4 = λ.UnpackIterable(τmp3, 2)
							ϒquality = λ.GetItem(τmp4, λ.IntLiteral(0))
							ϒitem = λ.GetItem(τmp4, λ.IntLiteral(1))
							if !λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒitem, λ.DictType)) {
								continue
							}
							ϒitem_url = λ.Calm(ϒitem, "get", λ.StrLiteral("url"))
							if !λ.IsTrue(ϒitem_url) {
								continue
							}
							λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
								"url": ϒitem_url,
								"format_id": λ.Mod(λ.StrLiteral("%s-%s-%s"), λ.NewTuple(
									HTML5,
									ϒkind,
									ϒquality,
								)),
								"filesize": λ.Cal(ϒint_or_none, λ.Calm(ϒitem, "get", λ.StrLiteral("size"))),
								"vcodec": func() λ.Object {
									if λ.IsTrue(λ.Eq(ϒkind, λ.StrLiteral("audio"))) {
										return λ.StrLiteral("none")
									} else {
										return λ.None
									}
								}(),
								"quality":    λ.Cal(ϒquality_key, ϒquality),
								"preference": λ.Cal(ϒpreference_key, HTML5),
							}))
						}
					}
					ϒiphone_url = λ.Calm(λ.Calm(ϒfile_versions, "get", IPHONE, λ.DictLiteral(map[λ.Object]λ.Object{})), "get", λ.StrLiteral("url"))
					if λ.IsTrue(ϒiphone_url) {
						λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
							"url":        ϒiphone_url,
							"format_id":  IPHONE,
							"preference": λ.Cal(ϒpreference_key, IPHONE),
						}))
					}
					ϒmobile_url = λ.Calm(λ.Calm(ϒfile_versions, "get", MOBILE, λ.DictLiteral(map[λ.Object]λ.Object{})), "get", λ.StrLiteral("audio_url"))
					if λ.IsTrue(ϒmobile_url) {
						λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
							"url":        ϒmobile_url,
							"format_id":  λ.Mod(λ.StrLiteral("%s-audio"), MOBILE),
							"preference": λ.Cal(ϒpreference_key, MOBILE),
						}))
					}
					λ.Calm(ϒself, "_sort_formats", ϒformats)
					ϒthumbnail = λ.Calm(ϒcoub, "get", λ.StrLiteral("picture"))
					ϒduration = λ.Cal(ϒfloat_or_none, λ.Calm(ϒcoub, "get", λ.StrLiteral("duration")))
					ϒtimestamp = λ.Cal(ϒparse_iso8601, func() λ.Object {
						if λv := λ.Calm(ϒcoub, "get", λ.StrLiteral("published_at")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Calm(ϒcoub, "get", λ.StrLiteral("created_at"))
						}
					}())
					ϒuploader = λ.Calm(λ.Calm(ϒcoub, "get", λ.StrLiteral("channel"), λ.DictLiteral(map[λ.Object]λ.Object{})), "get", λ.StrLiteral("title"))
					ϒuploader_id = λ.Calm(λ.Calm(ϒcoub, "get", λ.StrLiteral("channel"), λ.DictLiteral(map[λ.Object]λ.Object{})), "get", λ.StrLiteral("permalink"))
					ϒview_count = λ.Cal(ϒint_or_none, func() λ.Object {
						if λv := λ.Calm(ϒcoub, "get", λ.StrLiteral("views_count")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Calm(ϒcoub, "get", λ.StrLiteral("views_increase_count"))
						}
					}())
					ϒlike_count = λ.Cal(ϒint_or_none, λ.Calm(ϒcoub, "get", λ.StrLiteral("likes_count")))
					ϒrepost_count = λ.Cal(ϒint_or_none, λ.Calm(ϒcoub, "get", λ.StrLiteral("recoubs_count")))
					ϒage_restricted = λ.Calm(ϒcoub, "get", λ.StrLiteral("age_restricted"), λ.Calm(ϒcoub, "get", λ.StrLiteral("age_restricted_by_admin")))
					if ϒage_restricted != λ.None {
						ϒage_limit = func() λ.Object {
							if ϒage_restricted == λ.True {
								return λ.IntLiteral(18)
							} else {
								return λ.IntLiteral(0)
							}
						}()
					} else {
						ϒage_limit = λ.None
					}
					return λ.DictLiteral(map[string]λ.Object{
						"id":           ϒvideo_id,
						"title":        ϒtitle,
						"thumbnail":    ϒthumbnail,
						"duration":     ϒduration,
						"timestamp":    ϒtimestamp,
						"uploader":     ϒuploader,
						"uploader_id":  ϒuploader_id,
						"view_count":   ϒview_count,
						"like_count":   ϒlike_count,
						"repost_count": ϒrepost_count,
						"age_limit":    ϒage_limit,
						"formats":      ϒformats,
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL":    CoubIE__VALID_URL,
				"_real_extract": CoubIE__real_extract,
			})
		}())
	})
}
