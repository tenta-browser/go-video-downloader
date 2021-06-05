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
 * ard/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/ard.py
 */

package ard

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ARDBetaMediathekIE       λ.Object
	ARDIE                    λ.Object
	ARDMediathekBaseIE       λ.Object
	ARDMediathekIE           λ.Object
	ExtractorError           λ.Object
	InfoExtractor            λ.Object
	ϒcompat_etree_fromstring λ.Object
	ϒdetermine_ext           λ.Object
	ϒint_or_none             λ.Object
	ϒparse_duration          λ.Object
	ϒqualities               λ.Object
	ϒstr_or_none             λ.Object
	ϒtry_get                 λ.Object
	ϒunified_strdate         λ.Object
	ϒunified_timestamp       λ.Object
	ϒupdate_url_query        λ.Object
	ϒurl_or_none             λ.Object
	ϒxpath_text              λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒdetermine_ext = Ωutils.ϒdetermine_ext
		ExtractorError = Ωutils.ExtractorError
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒparse_duration = Ωutils.ϒparse_duration
		ϒqualities = Ωutils.ϒqualities
		ϒstr_or_none = Ωutils.ϒstr_or_none
		ϒtry_get = Ωutils.ϒtry_get
		ϒunified_strdate = Ωutils.ϒunified_strdate
		ϒunified_timestamp = Ωutils.ϒunified_timestamp
		ϒupdate_url_query = Ωutils.ϒupdate_url_query
		ϒurl_or_none = Ωutils.ϒurl_or_none
		ϒxpath_text = Ωutils.ϒxpath_text
		ϒcompat_etree_fromstring = Ωcompat.ϒcompat_etree_fromstring
		ARDMediathekBaseIE = λ.Cal(λ.TypeType, λ.StrLiteral("ARDMediathekBaseIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {

			return λ.ClassDictLiteral(map[λ.Object]λ.Object{})
		}())
		ARDMediathekIE = λ.Cal(λ.TypeType, λ.StrLiteral("ARDMediathekIE"), λ.NewTuple(ARDMediathekBaseIE), func() λ.Dict {
			var (
				ARDMediathekIE__VALID_URL λ.Object
				ARDMediathekIE_suitable   λ.Object
			)
			ARDMediathekIE__VALID_URL = λ.StrLiteral("^https?://(?:(?:(?:www|classic)\\.)?ardmediathek\\.de|mediathek\\.(?:daserste|rbb-online)\\.de|one\\.ard\\.de)/(?:.*/)(?P<video_id>[0-9]+|[^0-9][^/\\?]+)[^/\\?]*(?:\\?.*)?")
			ARDMediathekIE_suitable = λ.NewFunction("suitable",
				[]λ.Param{
					{Name: "cls"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒcls = λargs[0]
						ϒurl = λargs[1]
					)
					return func() λ.Object {
						if λ.IsTrue(λ.Calm(ARDBetaMediathekIE, "suitable", ϒurl)) {
							return λ.False
						} else {
							return λ.Calm(λ.Cal(λ.SuperType, ARDMediathekIE, ϒcls), "suitable", ϒurl)
						}
					}()
				})
			ARDMediathekIE_suitable = λ.Cal(λ.ClassMethodType, ARDMediathekIE_suitable)
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": ARDMediathekIE__VALID_URL,
				"suitable":   ARDMediathekIE_suitable,
			})
		}())
		ARDIE = λ.Cal(λ.TypeType, λ.StrLiteral("ARDIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				ARDIE__VALID_URL    λ.Object
				ARDIE__real_extract λ.Object
			)
			ARDIE__VALID_URL = λ.StrLiteral("(?P<mainurl>https?://(?:www\\.)?daserste\\.de/(?:[^/?#&]+/)+(?P<id>[^/?#&]+))\\.html")
			ARDIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒa             λ.Object
						ϒdisplay_id    λ.Object
						ϒdoc           λ.Object
						ϒext           λ.Object
						ϒf             λ.Object
						ϒfile_name     λ.Object
						ϒformat_type   λ.Object
						ϒformat_url    λ.Object
						ϒformats       λ.Object
						ϒmobj          λ.Object
						ϒplayer_url    λ.Object
						ϒself          = λargs[0]
						ϒserver_prefix λ.Object
						ϒthumbnail     λ.Object
						ϒupload_date   λ.Object
						ϒurl           = λargs[1]
						ϒvideo_node    λ.Object
						τmp0           λ.Object
						τmp1           λ.Object
					)
					ϒmobj = λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl)
					ϒdisplay_id = λ.Calm(ϒmobj, "group", λ.StrLiteral("id"))
					ϒplayer_url = λ.Add(λ.Calm(ϒmobj, "group", λ.StrLiteral("mainurl")), λ.StrLiteral("~playerXml.xml"))
					ϒdoc = λ.Calm(ϒself, "_download_xml", ϒplayer_url, ϒdisplay_id)
					ϒvideo_node = λ.Calm(ϒdoc, "find", λ.StrLiteral("./video"))
					ϒupload_date = λ.Cal(ϒunified_strdate, λ.Cal(ϒxpath_text, ϒvideo_node, λ.StrLiteral("./broadcastDate")))
					ϒthumbnail = λ.Cal(ϒxpath_text, ϒvideo_node, λ.StrLiteral(".//teaserImage//variant/url"))
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(ϒvideo_node, "findall", λ.StrLiteral(".//asset")))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒa = τmp1
						ϒfile_name = λ.Call(ϒxpath_text, λ.NewArgs(
							ϒa,
							λ.StrLiteral("./fileName"),
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
						})
						if !λ.IsTrue(ϒfile_name) {
							continue
						}
						ϒformat_type = λ.Calm(λ.GetAttr(ϒa, "attrib", nil), "get", λ.StrLiteral("type"))
						ϒformat_url = λ.Cal(ϒurl_or_none, ϒfile_name)
						if λ.IsTrue(ϒformat_url) {
							ϒext = λ.Cal(ϒdetermine_ext, ϒfile_name)
							if λ.IsTrue(λ.Eq(ϒext, λ.StrLiteral("m3u8"))) {
								λ.Calm(ϒformats, "extend", λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
									ϒformat_url,
									ϒdisplay_id,
									λ.StrLiteral("mp4"),
								), λ.KWArgs{
									{Name: "entry_protocol", Value: λ.StrLiteral("m3u8_native")},
									{Name: "m3u8_id", Value: func() λ.Object {
										if λv := ϒformat_type; λ.IsTrue(λv) {
											return λv
										} else {
											return λ.StrLiteral("hls")
										}
									}()},
									{Name: "fatal", Value: λ.False},
								}))
								continue
							} else {
								if λ.IsTrue(λ.Eq(ϒext, λ.StrLiteral("f4m"))) {
									λ.Calm(ϒformats, "extend", λ.Call(λ.GetAttr(ϒself, "_extract_f4m_formats", nil), λ.NewArgs(
										λ.Cal(ϒupdate_url_query, ϒformat_url, λ.DictLiteral(map[string]string{
											"hdcore": "3.7.0",
										})),
										ϒdisplay_id,
									), λ.KWArgs{
										{Name: "f4m_id", Value: func() λ.Object {
											if λv := ϒformat_type; λ.IsTrue(λv) {
												return λv
											} else {
												return λ.StrLiteral("hds")
											}
										}()},
										{Name: "fatal", Value: λ.False},
									}))
									continue
								}
							}
						}
						ϒf = λ.DictLiteral(map[string]λ.Object{
							"format_id": ϒformat_type,
							"width":     λ.Cal(ϒint_or_none, λ.Cal(ϒxpath_text, ϒa, λ.StrLiteral("./frameWidth"))),
							"height":    λ.Cal(ϒint_or_none, λ.Cal(ϒxpath_text, ϒa, λ.StrLiteral("./frameHeight"))),
							"vbr":       λ.Cal(ϒint_or_none, λ.Cal(ϒxpath_text, ϒa, λ.StrLiteral("./bitrateVideo"))),
							"abr":       λ.Cal(ϒint_or_none, λ.Cal(ϒxpath_text, ϒa, λ.StrLiteral("./bitrateAudio"))),
							"vcodec":    λ.Cal(ϒxpath_text, ϒa, λ.StrLiteral("./codecVideo")),
							"tbr":       λ.Cal(ϒint_or_none, λ.Cal(ϒxpath_text, ϒa, λ.StrLiteral("./totalBitrate"))),
						})
						ϒserver_prefix = λ.Call(ϒxpath_text, λ.NewArgs(
							ϒa,
							λ.StrLiteral("./serverPrefix"),
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
						})
						if λ.IsTrue(ϒserver_prefix) {
							λ.Calm(ϒf, "update", λ.DictLiteral(map[string]λ.Object{
								"url":      ϒserver_prefix,
								"playpath": ϒfile_name,
							}))
						} else {
							if !λ.IsTrue(ϒformat_url) {
								continue
							}
							λ.SetItem(ϒf, λ.StrLiteral("url"), ϒformat_url)
						}
						λ.Calm(ϒformats, "append", ϒf)
					}
					λ.Calm(ϒself, "_sort_formats", ϒformats)
					return λ.DictLiteral(map[string]λ.Object{
						"id": λ.Call(ϒxpath_text, λ.NewArgs(
							ϒvideo_node,
							λ.StrLiteral("./videoId"),
						), λ.KWArgs{
							{Name: "default", Value: ϒdisplay_id},
						}),
						"formats":     ϒformats,
						"display_id":  ϒdisplay_id,
						"title":       λ.GetAttr(λ.Calm(ϒvideo_node, "find", λ.StrLiteral("./title")), "text", nil),
						"duration":    λ.Cal(ϒparse_duration, λ.GetAttr(λ.Calm(ϒvideo_node, "find", λ.StrLiteral("./duration")), "text", nil)),
						"upload_date": ϒupload_date,
						"thumbnail":   ϒthumbnail,
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL":    ARDIE__VALID_URL,
				"_real_extract": ARDIE__real_extract,
			})
		}())
		ARDBetaMediathekIE = λ.Cal(λ.TypeType, λ.StrLiteral("ARDBetaMediathekIE"), λ.NewTuple(ARDMediathekBaseIE), func() λ.Dict {
			var (
				ARDBetaMediathekIE__VALID_URL λ.Object
			)
			ARDBetaMediathekIE__VALID_URL = λ.StrLiteral("https://(?:(?:beta|www)\\.)?ardmediathek\\.de/(?:[^/]+/)?(?:player|live|video)/(?:[^/]+/)*(?P<id>Y3JpZDovL[a-zA-Z0-9]+)")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": ARDBetaMediathekIE__VALID_URL,
			})
		}())
	})
}
