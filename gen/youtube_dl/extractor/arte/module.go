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
 * arte/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/arte.py
 */

package arte

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωparse "github.com/tenta-browser/go-video-downloader/gen/urllib/parse"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ArteTVBaseIE     λ.Object
	ArteTVEmbedIE    λ.Object
	ArteTVIE         λ.Object
	ArteTVPlaylistIE λ.Object
	ExtractorError   λ.Object
	InfoExtractor    λ.Object
	ϒcompat_str      λ.Object
	ϒint_or_none     λ.Object
	ϒqualities       λ.Object
	ϒtry_get         λ.Object
	ϒunified_strdate λ.Object
	ϒurl_or_none     λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_str = Ωcompat.ϒcompat_str
		ExtractorError = Ωutils.ExtractorError
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒqualities = Ωutils.ϒqualities
		ϒtry_get = Ωutils.ϒtry_get
		ϒunified_strdate = Ωutils.ϒunified_strdate
		ϒurl_or_none = Ωutils.ϒurl_or_none
		ArteTVBaseIE = λ.Cal(λ.TypeType, λ.StrLiteral("ArteTVBaseIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				ArteTVBaseIE__API_BASE       λ.Object
				ArteTVBaseIE__ARTE_LANGUAGES λ.Object
			)
			ArteTVBaseIE__ARTE_LANGUAGES = λ.StrLiteral("fr|de|en|es|it|pl")
			ArteTVBaseIE__API_BASE = λ.StrLiteral("https://api.arte.tv/api/player/v1")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_API_BASE":       ArteTVBaseIE__API_BASE,
				"_ARTE_LANGUAGES": ArteTVBaseIE__ARTE_LANGUAGES,
			})
		}())
		ArteTVIE = λ.Cal(λ.TypeType, λ.StrLiteral("ArteTVIE"), λ.NewTuple(ArteTVBaseIE), func() λ.Dict {
			var (
				ArteTVIE__VALID_URL    λ.Object
				ArteTVIE__real_extract λ.Object
			)
			ArteTVIE__VALID_URL = λ.Mod(λ.StrLiteral("(?x)\n                    https?://\n                        (?:\n                            (?:www\\.)?arte\\.tv/(?P<lang>%(langs)s)/videos|\n                            api\\.arte\\.tv/api/player/v\\d+/config/(?P<lang_2>%(langs)s)\n                        )\n                        /(?P<id>\\d{6}-\\d{3}-[AF])\n                    "), λ.DictLiteral(map[string]λ.Object{
				"langs": λ.GetAttr(ArteTVBaseIE, "_ARTE_LANGUAGES", nil),
			}))
			ArteTVIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						LANGS            λ.Object
						PREFERENCES      λ.Object
						ϒerror           λ.Object
						ϒf               λ.Object
						ϒformat          λ.Object
						ϒformat_dict     λ.Object
						ϒformat_id       λ.Object
						ϒformat_url      λ.Object
						ϒformats         λ.Object
						ϒinfo            λ.Object
						ϒl               λ.Object
						ϒlang            λ.Object
						ϒlang_pref       λ.Object
						ϒlangcode        λ.Object
						ϒm3u8_format     λ.Object
						ϒm3u8_formats    λ.Object
						ϒmedia_type      λ.Object
						ϒmobj            λ.Object
						ϒp               λ.Object
						ϒplayer_info     λ.Object
						ϒpref            λ.Object
						ϒqfunc           λ.Object
						ϒself            = λargs[0]
						ϒstreamer        λ.Object
						ϒsubtitle        λ.Object
						ϒtitle           λ.Object
						ϒupload_date_str λ.Object
						ϒurl             = λargs[1]
						ϒversionCode     λ.Object
						ϒvideo_id        λ.Object
						ϒvsr             λ.Object
						τmp0             λ.Object
						τmp1             λ.Object
						τmp2             λ.Object
						τmp3             λ.Object
						τmp4             λ.Object
					)
					ϒmobj = λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl)
					ϒvideo_id = λ.Calm(ϒmobj, "group", λ.StrLiteral("id"))
					ϒlang = func() λ.Object {
						if λv := λ.Calm(ϒmobj, "group", λ.StrLiteral("lang")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Calm(ϒmobj, "group", λ.StrLiteral("lang_2"))
						}
					}()
					ϒinfo = λ.Calm(ϒself, "_download_json", λ.Mod(λ.StrLiteral("%s/config/%s/%s"), λ.NewTuple(
						λ.GetAttr(ϒself, "_API_BASE", nil),
						ϒlang,
						ϒvideo_id,
					)), ϒvideo_id)
					ϒplayer_info = λ.GetItem(ϒinfo, λ.StrLiteral("videoJsonPlayer"))
					ϒvsr = λ.Cal(ϒtry_get, ϒplayer_info, λ.NewFunction("<lambda>",
						[]λ.Param{
							{Name: "x"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒx = λargs[0]
							)
							return λ.GetItem(ϒx, λ.StrLiteral("VSR"))
						}), λ.DictType)
					if !λ.IsTrue(ϒvsr) {
						ϒerror = λ.None
						if λ.IsTrue(λ.Eq(λ.Cal(ϒtry_get, ϒplayer_info, λ.NewFunction("<lambda>",
							[]λ.Param{
								{Name: "x"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒx = λargs[0]
								)
								return λ.GetItem(λ.GetItem(ϒx, λ.StrLiteral("custom_msg")), λ.StrLiteral("type"))
							})), λ.StrLiteral("error"))) {
							ϒerror = λ.Cal(ϒtry_get, ϒplayer_info, λ.NewFunction("<lambda>",
								[]λ.Param{
									{Name: "x"},
								},
								0, false, false,
								func(λargs []λ.Object) λ.Object {
									var (
										ϒx = λargs[0]
									)
									return λ.GetItem(λ.GetItem(ϒx, λ.StrLiteral("custom_msg")), λ.StrLiteral("msg"))
								}), ϒcompat_str)
						}
						if !λ.IsTrue(ϒerror) {
							ϒerror = func() λ.Object {
								if λv := λ.Mod(λ.StrLiteral("Video %s is not available"), λ.Calm(ϒplayer_info, "get", λ.StrLiteral("VID"))); λ.IsTrue(λv) {
									return λv
								} else {
									return ϒvideo_id
								}
							}()
						}
						panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(ϒerror), λ.KWArgs{
							{Name: "expected", Value: λ.True},
						})))
					}
					ϒupload_date_str = λ.Calm(ϒplayer_info, "get", λ.StrLiteral("shootingDate"))
					if !λ.IsTrue(ϒupload_date_str) {
						ϒupload_date_str = λ.GetItem(λ.Calm(func() λ.Object {
							if λv := λ.Calm(ϒplayer_info, "get", λ.StrLiteral("VRA")); λ.IsTrue(λv) {
								return λv
							} else if λv := λ.Calm(ϒplayer_info, "get", λ.StrLiteral("VDA")); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.StrLiteral("")
							}
						}(), "split", λ.StrLiteral(" ")), λ.IntLiteral(0))
					}
					ϒtitle = λ.Calm(func() λ.Object {
						if λv := λ.Calm(ϒplayer_info, "get", λ.StrLiteral("VTI")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.GetItem(ϒplayer_info, λ.StrLiteral("VID"))
						}
					}(), "strip")
					ϒsubtitle = λ.Calm(λ.Calm(ϒplayer_info, "get", λ.StrLiteral("VSU"), λ.StrLiteral("")), "strip")
					if λ.IsTrue(ϒsubtitle) {
						τmp0 = λ.IAdd(ϒtitle, λ.Mod(λ.StrLiteral(" - %s"), ϒsubtitle))
						ϒtitle = τmp0
					}
					ϒqfunc = λ.Cal(ϒqualities, λ.NewList(
						λ.StrLiteral("MQ"),
						λ.StrLiteral("HQ"),
						λ.StrLiteral("EQ"),
						λ.StrLiteral("SQ"),
					))
					LANGS = λ.DictLiteral(map[string]string{
						"fr": "F",
						"de": "A",
						"en": "E[ANG]",
						"es": "E[ESP]",
						"it": "E[ITA]",
						"pl": "E[POL]",
					})
					ϒlangcode = λ.Calm(LANGS, "get", ϒlang, ϒlang)
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(ϒvsr, "items"))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						τmp2 = τmp1
						ϒformat_id = λ.GetItem(τmp2, λ.IntLiteral(0))
						ϒformat_dict = λ.GetItem(τmp2, λ.IntLiteral(1))
						ϒf = λ.Cal(λ.DictType, ϒformat_dict)
						ϒformat_url = λ.Cal(ϒurl_or_none, λ.Calm(ϒf, "get", λ.StrLiteral("url")))
						ϒstreamer = λ.Calm(ϒf, "get", λ.StrLiteral("streamer"))
						if λ.IsTrue(func() λ.Object {
							if λv := λ.NewBool(!λ.IsTrue(ϒformat_url)); !λ.IsTrue(λv) {
								return λv
							} else {
								return λ.NewBool(!λ.IsTrue(ϒstreamer))
							}
						}()) {
							continue
						}
						ϒversionCode = λ.Calm(ϒf, "get", λ.StrLiteral("versionCode"))
						ϒl = λ.Cal(Ωre.ϒescape, ϒlangcode)
						PREFERENCES = λ.NewTuple(
							λ.Calm(λ.StrLiteral("VO{0}$"), "format", ϒl),
							λ.Calm(λ.StrLiteral("VO{0}-ST{0}$"), "format", ϒl),
							λ.Calm(λ.StrLiteral("VO{0}-STM{0}$"), "format", ϒl),
							λ.Calm(λ.StrLiteral("V{0}$"), "format", ϒl),
							λ.Calm(λ.StrLiteral("V{0}-ST{0}$"), "format", ϒl),
							λ.Calm(λ.StrLiteral("V{0}-STM{0}$"), "format", ϒl),
							λ.Calm(λ.StrLiteral("VO{0}-ST(?!{0}).+?$"), "format", ϒl),
							λ.Calm(λ.StrLiteral("VO{0}-STM(?!{0}).+?$"), "format", ϒl),
							λ.Calm(λ.StrLiteral("VO(?:(?!{0}).+?)?-ST{0}$"), "format", ϒl),
							λ.Calm(λ.StrLiteral("VO(?:(?!{0}).+?)?-STM{0}$"), "format", ϒl),
							λ.Calm(λ.StrLiteral("VO(?:(?!{0}))?$"), "format", ϒl),
							λ.Calm(λ.StrLiteral("VO(?:(?!{0}).+?)?-ST(?!{0}).+?$"), "format", ϒl),
							λ.Calm(λ.StrLiteral("VO(?:(?!{0}).+?)?-STM(?!{0}).+?$"), "format", ϒl),
						)
						τmp2 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.EnumerateIteratorType, PREFERENCES))
						for {
							if τmp3 = λ.NextDefault(τmp2, λ.AfterLast); τmp3 == λ.AfterLast {
								break
							}
							τmp4 = τmp3
							ϒpref = λ.GetItem(τmp4, λ.IntLiteral(0))
							ϒp = λ.GetItem(τmp4, λ.IntLiteral(1))
							if λ.IsTrue(λ.Cal(Ωre.ϒmatch, ϒp, ϒversionCode)) {
								ϒlang_pref = λ.Sub(λ.Cal(λ.BuiltinLen, PREFERENCES), ϒpref)
								break
							}
						}
						if τmp3 == λ.AfterLast {
							ϒlang_pref = λ.Neg(λ.IntLiteral(1))
						}
						ϒmedia_type = λ.Calm(ϒf, "get", λ.StrLiteral("mediaType"))
						if λ.IsTrue(λ.Eq(ϒmedia_type, λ.StrLiteral("hls"))) {
							ϒm3u8_formats = λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
								ϒformat_url,
								ϒvideo_id,
								λ.StrLiteral("mp4"),
							), λ.KWArgs{
								{Name: "entry_protocol", Value: λ.StrLiteral("m3u8_native")},
								{Name: "m3u8_id", Value: ϒformat_id},
								{Name: "fatal", Value: λ.False},
							})
							τmp2 = λ.Cal(λ.BuiltinIter, ϒm3u8_formats)
							for {
								if τmp3 = λ.NextDefault(τmp2, λ.AfterLast); τmp3 == λ.AfterLast {
									break
								}
								ϒm3u8_format = τmp3
								λ.SetItem(ϒm3u8_format, λ.StrLiteral("language_preference"), ϒlang_pref)
							}
							λ.Calm(ϒformats, "extend", ϒm3u8_formats)
							continue
						}
						ϒformat = λ.DictLiteral(map[string]λ.Object{
							"format_id": ϒformat_id,
							"preference": func() λ.Object {
								if λ.IsTrue(λ.Eq(λ.Calm(ϒf, "get", λ.StrLiteral("videoFormat")), λ.StrLiteral("M3U8"))) {
									return λ.Neg(λ.IntLiteral(10))
								} else {
									return λ.None
								}
							}(),
							"language_preference": ϒlang_pref,
							"format_note": λ.Mod(λ.StrLiteral("%s, %s"), λ.NewTuple(
								λ.Calm(ϒf, "get", λ.StrLiteral("versionCode")),
								λ.Calm(ϒf, "get", λ.StrLiteral("versionLibelle")),
							)),
							"width":   λ.Cal(ϒint_or_none, λ.Calm(ϒf, "get", λ.StrLiteral("width"))),
							"height":  λ.Cal(ϒint_or_none, λ.Calm(ϒf, "get", λ.StrLiteral("height"))),
							"tbr":     λ.Cal(ϒint_or_none, λ.Calm(ϒf, "get", λ.StrLiteral("bitrate"))),
							"quality": λ.Cal(ϒqfunc, λ.Calm(ϒf, "get", λ.StrLiteral("quality"))),
						})
						if λ.IsTrue(λ.Eq(ϒmedia_type, λ.StrLiteral("rtmp"))) {
							λ.SetItem(ϒformat, λ.StrLiteral("url"), λ.GetItem(ϒf, λ.StrLiteral("streamer")))
							λ.SetItem(ϒformat, λ.StrLiteral("play_path"), λ.Add(λ.StrLiteral("mp4:"), λ.GetItem(ϒf, λ.StrLiteral("url"))))
							λ.SetItem(ϒformat, λ.StrLiteral("ext"), λ.StrLiteral("flv"))
						} else {
							λ.SetItem(ϒformat, λ.StrLiteral("url"), λ.GetItem(ϒf, λ.StrLiteral("url")))
						}
						λ.Calm(ϒformats, "append", ϒformat)
					}
					λ.Calm(ϒself, "_sort_formats", ϒformats)
					return λ.DictLiteral(map[string]λ.Object{
						"id": func() λ.Object {
							if λv := λ.Calm(ϒplayer_info, "get", λ.StrLiteral("VID")); λ.IsTrue(λv) {
								return λv
							} else {
								return ϒvideo_id
							}
						}(),
						"title":       ϒtitle,
						"description": λ.Calm(ϒplayer_info, "get", λ.StrLiteral("VDE")),
						"upload_date": λ.Cal(ϒunified_strdate, ϒupload_date_str),
						"thumbnail": func() λ.Object {
							if λv := λ.Calm(ϒplayer_info, "get", λ.StrLiteral("programImage")); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Calm(λ.Calm(ϒplayer_info, "get", λ.StrLiteral("VTU"), λ.DictLiteral(map[λ.Object]λ.Object{})), "get", λ.StrLiteral("IUR"))
							}
						}(),
						"formats": ϒformats,
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL":    ArteTVIE__VALID_URL,
				"_real_extract": ArteTVIE__real_extract,
			})
		}())
		ArteTVEmbedIE = λ.Cal(λ.TypeType, λ.StrLiteral("ArteTVEmbedIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				ArteTVEmbedIE__VALID_URL    λ.Object
				ArteTVEmbedIE__real_extract λ.Object
			)
			ArteTVEmbedIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?arte\\.tv/player/v\\d+/index\\.php\\?.*?\\bjson_url=.+")
			ArteTVEmbedIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒjson_url λ.Object
						ϒqs       λ.Object
						ϒself     = λargs[0]
						ϒurl      = λargs[1]
						ϒvideo_id λ.Object
					)
					ϒqs = λ.Cal(Ωparse.ϒparse_qs, λ.GetAttr(λ.Cal(Ωparse.ϒurlparse, ϒurl), "query", nil))
					ϒjson_url = λ.GetItem(λ.GetItem(ϒqs, λ.StrLiteral("json_url")), λ.IntLiteral(0))
					ϒvideo_id = λ.Calm(ArteTVIE, "_match_id", ϒjson_url)
					return λ.Call(λ.GetAttr(ϒself, "url_result", nil), λ.NewArgs(ϒjson_url), λ.KWArgs{
						{Name: "ie", Value: λ.Calm(ArteTVIE, "ie_key")},
						{Name: "video_id", Value: ϒvideo_id},
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL":    ArteTVEmbedIE__VALID_URL,
				"_real_extract": ArteTVEmbedIE__real_extract,
			})
		}())
		ArteTVPlaylistIE = λ.Cal(λ.TypeType, λ.StrLiteral("ArteTVPlaylistIE"), λ.NewTuple(ArteTVBaseIE), func() λ.Dict {
			var (
				ArteTVPlaylistIE__VALID_URL λ.Object
			)
			ArteTVPlaylistIE__VALID_URL = λ.Mod(λ.StrLiteral("https?://(?:www\\.)?arte\\.tv/(?P<lang>%s)/videos/(?P<id>RC-\\d{6})"), λ.GetAttr(ArteTVBaseIE, "_ARTE_LANGUAGES", nil))
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": ArteTVPlaylistIE__VALID_URL,
			})
		}())
	})
}
