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
 * threeqsdn/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/threeqsdn.py
 */

package threeqsdn

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError λ.Object
	InfoExtractor  λ.Object
	ThreeQSDNIE    λ.Object
	ϒdetermine_ext λ.Object
	ϒfloat_or_none λ.Object
	ϒint_or_none   λ.Object
	ϒparse_iso8601 λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒdetermine_ext = Ωutils.ϒdetermine_ext
		ExtractorError = Ωutils.ExtractorError
		ϒfloat_or_none = Ωutils.ϒfloat_or_none
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒparse_iso8601 = Ωutils.ϒparse_iso8601
		ThreeQSDNIE = λ.Cal(λ.TypeType, λ.StrLiteral("ThreeQSDNIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				ThreeQSDNIE_IE_NAME       λ.Object
				ThreeQSDNIE__VALID_URL    λ.Object
				ThreeQSDNIE__real_extract λ.Object
			)
			ThreeQSDNIE_IE_NAME = λ.StrLiteral("3qsdn")
			ThreeQSDNIE__VALID_URL = λ.StrLiteral("https?://playout\\.3qsdn\\.com/(?P<id>[\\da-f]{8}-[\\da-f]{4}-[\\da-f]{4}-[\\da-f]{4}-[\\da-f]{12})")
			ThreeQSDNIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒaspect      λ.Object
						ϒconfig      λ.Object
						ϒext         λ.Object
						ϒf           λ.Object
						ϒformat_id   λ.Object
						ϒformats     λ.Object
						ϒheight      λ.Object
						ϒlive        λ.Object
						ϒs           λ.Object
						ϒself        = λargs[0]
						ϒsource      λ.Object
						ϒsource_type λ.Object
						ϒsrc         λ.Object
						ϒsubtitle    λ.Object
						ϒsubtitles   λ.Object
						ϒtitle       λ.Object
						ϒurl         = λargs[1]
						ϒvideo_id    λ.Object
						ϒwidth       λ.Object
						τmp0         λ.Object
						τmp1         λ.Object
						τmp2         λ.Object
						τmp3         λ.Object
					)
					ϒvideo_id = λ.Calm(ϒself, "_match_id", ϒurl)
					τmp0, τmp1 = func() (λexit λ.Object, λret λ.Object) {
						defer λ.CatchMulti(
							nil,
							&λ.Catcher{ExtractorError, func(λex λ.BaseException) {
								var ϒe λ.Object = λex
								if λ.IsTrue(func() λ.Object {
									if λv := λ.Cal(λ.BuiltinIsInstance, λ.GetAttr(ϒe, "cause", nil), λ.None); !λ.IsTrue(λv) {
										return λv
									} else {
										return λ.Eq(λ.GetAttr(λ.GetAttr(ϒe, "cause", nil), "code", nil), λ.IntLiteral(401))
									}
								}()) {
									λ.Calm(ϒself, "raise_geo_restricted")
								}
								panic(λ.Raise(λex))
							}},
						)
						ϒconfig = λ.Calm(ϒself, "_download_json", λ.Calm(ϒurl, "replace", λ.StrLiteral("://playout.3qsdn.com/"), λ.StrLiteral("://playout.3qsdn.com/config/")), ϒvideo_id)
						return λ.BlockExitNormally, nil
					}()
					ϒlive = λ.Eq(λ.Calm(ϒconfig, "get", λ.StrLiteral("streamContent")), λ.StrLiteral("live"))
					ϒaspect = λ.Cal(ϒfloat_or_none, λ.Calm(ϒconfig, "get", λ.StrLiteral("aspect")))
					ϒformats = λ.NewList()
					τmp1 = λ.Cal(λ.BuiltinIter, λ.Calm(func() λ.Object {
						if λv := λ.Calm(ϒconfig, "get", λ.StrLiteral("sources")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.DictLiteral(map[λ.Object]λ.Object{})
						}
					}(), "items"))
					for {
						if τmp0 = λ.NextDefault(τmp1, λ.AfterLast); τmp0 == λ.AfterLast {
							break
						}
						τmp2 = τmp0
						ϒsource_type = λ.GetItem(τmp2, λ.IntLiteral(0))
						ϒsource = λ.GetItem(τmp2, λ.IntLiteral(1))
						if !λ.IsTrue(ϒsource) {
							continue
						}
						if λ.IsTrue(λ.Eq(ϒsource_type, λ.StrLiteral("dash"))) {
							λ.Calm(ϒformats, "extend", λ.Call(λ.GetAttr(ϒself, "_extract_mpd_formats", nil), λ.NewArgs(
								ϒsource,
								ϒvideo_id,
							), λ.KWArgs{
								{Name: "mpd_id", Value: λ.StrLiteral("mpd")},
								{Name: "fatal", Value: λ.False},
							}))
						} else {
							if λ.IsTrue(λ.Eq(ϒsource_type, λ.StrLiteral("hls"))) {
								λ.Calm(ϒformats, "extend", λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
									ϒsource,
									ϒvideo_id,
									λ.StrLiteral("mp4"),
									func() λ.Object {
										if λ.IsTrue(ϒlive) {
											return λ.StrLiteral("m3u8")
										} else {
											return λ.StrLiteral("m3u8_native")
										}
									}(),
								), λ.KWArgs{
									{Name: "m3u8_id", Value: λ.StrLiteral("hls")},
									{Name: "fatal", Value: λ.False},
								}))
							} else {
								if λ.IsTrue(λ.Eq(ϒsource_type, λ.StrLiteral("progressive"))) {
									τmp2 = λ.Cal(λ.BuiltinIter, ϒsource)
									for {
										if τmp3 = λ.NextDefault(τmp2, λ.AfterLast); τmp3 == λ.AfterLast {
											break
										}
										ϒs = τmp3
										ϒsrc = λ.Calm(ϒs, "get", λ.StrLiteral("src"))
										if !λ.IsTrue(func() λ.Object {
											if λv := ϒsrc; !λ.IsTrue(λv) {
												return λv
											} else {
												return λ.Calm(ϒself, "_is_valid_url", ϒsrc, ϒvideo_id)
											}
										}()) {
											continue
										}
										ϒwidth = λ.None
										ϒformat_id = λ.NewList(λ.StrLiteral("http"))
										ϒext = λ.Cal(ϒdetermine_ext, ϒsrc)
										if λ.IsTrue(ϒext) {
											λ.Calm(ϒformat_id, "append", ϒext)
										}
										ϒheight = λ.Cal(ϒint_or_none, λ.Calm(ϒs, "get", λ.StrLiteral("height")))
										if λ.IsTrue(ϒheight) {
											λ.Calm(ϒformat_id, "append", λ.Mod(λ.StrLiteral("%dp"), ϒheight))
											if λ.IsTrue(ϒaspect) {
												ϒwidth = λ.Cal(λ.IntType, λ.Mul(ϒheight, ϒaspect))
											}
										}
										λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
											"ext":               ϒext,
											"format_id":         λ.Calm(λ.StrLiteral("-"), "join", ϒformat_id),
											"height":            ϒheight,
											"source_preference": λ.IntLiteral(0),
											"url":               ϒsrc,
											"vcodec": func() λ.Object {
												if λ.IsTrue(λ.Eq(ϒheight, λ.IntLiteral(0))) {
													return λ.StrLiteral("none")
												} else {
													return λ.None
												}
											}(),
											"width": ϒwidth,
										}))
									}
								}
							}
						}
					}
					τmp1 = λ.Cal(λ.BuiltinIter, ϒformats)
					for {
						if τmp0 = λ.NextDefault(τmp1, λ.AfterLast); τmp0 == λ.AfterLast {
							break
						}
						ϒf = τmp0
						if λ.IsTrue(λ.Eq(λ.Calm(ϒf, "get", λ.StrLiteral("acodec")), λ.StrLiteral("none"))) {
							λ.SetItem(ϒf, λ.StrLiteral("preference"), λ.Neg(λ.IntLiteral(40)))
						} else {
							if λ.IsTrue(λ.Eq(λ.Calm(ϒf, "get", λ.StrLiteral("vcodec")), λ.StrLiteral("none"))) {
								λ.SetItem(ϒf, λ.StrLiteral("preference"), λ.Neg(λ.IntLiteral(50)))
							}
						}
					}
					λ.Calm(ϒself, "_sort_formats", ϒformats, λ.NewTuple(
						λ.StrLiteral("preference"),
						λ.StrLiteral("width"),
						λ.StrLiteral("height"),
						λ.StrLiteral("source_preference"),
						λ.StrLiteral("tbr"),
						λ.StrLiteral("vbr"),
						λ.StrLiteral("abr"),
						λ.StrLiteral("ext"),
						λ.StrLiteral("format_id"),
					))
					ϒsubtitles = λ.DictLiteral(map[λ.Object]λ.Object{})
					τmp1 = λ.Cal(λ.BuiltinIter, func() λ.Object {
						if λv := λ.Calm(ϒconfig, "get", λ.StrLiteral("subtitles")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.NewList()
						}
					}())
					for {
						if τmp0 = λ.NextDefault(τmp1, λ.AfterLast); τmp0 == λ.AfterLast {
							break
						}
						ϒsubtitle = τmp0
						ϒsrc = λ.Calm(ϒsubtitle, "get", λ.StrLiteral("src"))
						if !λ.IsTrue(ϒsrc) {
							continue
						}
						λ.Calm(λ.Calm(ϒsubtitles, "setdefault", func() λ.Object {
							if λv := λ.Calm(ϒsubtitle, "get", λ.StrLiteral("label")); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.StrLiteral("eng")
							}
						}(), λ.NewList()), "append", λ.DictLiteral(map[string]λ.Object{
							"url": ϒsrc,
						}))
					}
					ϒtitle = func() λ.Object {
						if λv := λ.Calm(ϒconfig, "get", λ.StrLiteral("title")); λ.IsTrue(λv) {
							return λv
						} else {
							return ϒvideo_id
						}
					}()
					return λ.DictLiteral(map[string]λ.Object{
						"id": ϒvideo_id,
						"title": func() λ.Object {
							if λ.IsTrue(ϒlive) {
								return λ.Calm(ϒself, "_live_title", ϒtitle)
							} else {
								return ϒtitle
							}
						}(),
						"thumbnail": func() λ.Object {
							if λv := λ.Calm(ϒconfig, "get", λ.StrLiteral("poster")); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.None
							}
						}(),
						"description": func() λ.Object {
							if λv := λ.Calm(ϒconfig, "get", λ.StrLiteral("description")); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.None
							}
						}(),
						"timestamp": λ.Cal(ϒparse_iso8601, λ.Calm(ϒconfig, "get", λ.StrLiteral("upload_date"))),
						"duration": func() λ.Object {
							if λv := λ.Cal(ϒfloat_or_none, λ.Calm(ϒconfig, "get", λ.StrLiteral("vlength"))); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.None
							}
						}(),
						"is_live":   ϒlive,
						"formats":   ϒformats,
						"subtitles": ϒsubtitles,
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"IE_NAME":       ThreeQSDNIE_IE_NAME,
				"_VALID_URL":    ThreeQSDNIE__VALID_URL,
				"_real_extract": ThreeQSDNIE__real_extract,
			})
		}())
	})
}
