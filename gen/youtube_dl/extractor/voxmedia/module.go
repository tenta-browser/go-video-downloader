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
 * voxmedia/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/voxmedia.py
 */

package voxmedia

import (
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωonce "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/once"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError               λ.Object
	InfoExtractor                λ.Object
	OnceIE                       λ.Object
	VoxMediaIE                   λ.Object
	VoxMediaVolumeIE             λ.Object
	ϒcompat_urllib_parse_unquote λ.Object
	ϒint_or_none                 λ.Object
	ϒtry_get                     λ.Object
	ϒunified_timestamp           λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		OnceIE = Ωonce.OnceIE
		ϒcompat_urllib_parse_unquote = Ωcompat.ϒcompat_urllib_parse_unquote
		ExtractorError = Ωutils.ExtractorError
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒtry_get = Ωutils.ϒtry_get
		ϒunified_timestamp = Ωutils.ϒunified_timestamp
		VoxMediaVolumeIE = λ.Cal(λ.TypeType, λ.StrLiteral("VoxMediaVolumeIE"), λ.NewTuple(OnceIE), func() λ.Dict {
			var (
				VoxMediaVolumeIE__VALID_URL    λ.Object
				VoxMediaVolumeIE__real_extract λ.Object
			)
			VoxMediaVolumeIE__VALID_URL = λ.StrLiteral("https?://volume\\.vox-cdn\\.com/embed/(?P<id>[0-9a-f]{9})")
			VoxMediaVolumeIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒasset               λ.Object
						ϒformat_id           λ.Object
						ϒformats             λ.Object
						ϒformatted_metadata  λ.Object
						ϒhls_url             λ.Object
						ϒinfo                λ.Object
						ϒmp4_url             λ.Object
						ϒplayer_setup        λ.Object
						ϒprovider_video_id   λ.Object
						ϒprovider_video_type λ.Object
						ϒself                = λargs[0]
						ϒsetup               λ.Object
						ϒtbr                 λ.Object
						ϒurl                 = λargs[1]
						ϒvideo_data          λ.Object
						ϒvideo_id            λ.Object
						ϒwebpage             λ.Object
						τmp0                 λ.Object
						τmp1                 λ.Object
					)
					ϒvideo_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒwebpage = λ.Calm(ϒself, "_download_webpage", ϒurl, ϒvideo_id)
					ϒsetup = λ.Calm(ϒself, "_parse_json", λ.Calm(ϒself, "_search_regex", λ.StrLiteral("setup\\s*=\\s*({.+});"), ϒwebpage, λ.StrLiteral("setup")), ϒvideo_id)
					ϒplayer_setup = func() λ.Object {
						if λv := λ.Calm(ϒsetup, "get", λ.StrLiteral("player_setup")); λ.IsTrue(λv) {
							return λv
						} else {
							return ϒsetup
						}
					}()
					ϒvideo_data = func() λ.Object {
						if λv := λ.Calm(ϒplayer_setup, "get", λ.StrLiteral("video")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.DictLiteral(map[λ.Object]λ.Object{})
						}
					}()
					ϒformatted_metadata = func() λ.Object {
						if λv := λ.Calm(ϒvideo_data, "get", λ.StrLiteral("formatted_metadata")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.DictLiteral(map[λ.Object]λ.Object{})
						}
					}()
					ϒinfo = λ.DictLiteral(map[string]λ.Object{
						"id": ϒvideo_id,
						"title": func() λ.Object {
							if λv := λ.Calm(ϒplayer_setup, "get", λ.StrLiteral("title")); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Calm(ϒvideo_data, "get", λ.StrLiteral("title_short"))
							}
						}(),
						"description": func() λ.Object {
							if λv := λ.Calm(ϒvideo_data, "get", λ.StrLiteral("description_long")); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Calm(ϒvideo_data, "get", λ.StrLiteral("description_short"))
							}
						}(),
						"thumbnail": func() λ.Object {
							if λv := λ.Calm(ϒformatted_metadata, "get", λ.StrLiteral("thumbnail")); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Calm(ϒvideo_data, "get", λ.StrLiteral("brightcove_thumbnail"))
							}
						}(),
						"timestamp": λ.Cal(ϒunified_timestamp, λ.Calm(ϒformatted_metadata, "get", λ.StrLiteral("video_publish_date"))),
					})
					ϒasset = func() λ.Object {
						if λv := λ.Cal(ϒtry_get, ϒsetup, λ.NewFunction("<lambda>",
							[]λ.Param{
								{Name: "x"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒx = λargs[0]
								)
								return λ.GetItem(λ.GetItem(ϒx, λ.StrLiteral("embed_assets")), λ.StrLiteral("chorus"))
							}), λ.DictType); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.DictLiteral(map[λ.Object]λ.Object{})
						}
					}()
					ϒformats = λ.NewList()
					ϒhls_url = λ.Calm(ϒasset, "get", λ.StrLiteral("hls_url"))
					if λ.IsTrue(ϒhls_url) {
						λ.Calm(ϒformats, "extend", λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
							ϒhls_url,
							ϒvideo_id,
							λ.StrLiteral("mp4"),
							λ.StrLiteral("m3u8_native"),
						), λ.KWArgs{
							{Name: "m3u8_id", Value: λ.StrLiteral("hls")},
							{Name: "fatal", Value: λ.False},
						}))
					}
					ϒmp4_url = λ.Calm(ϒasset, "get", λ.StrLiteral("mp4_url"))
					if λ.IsTrue(ϒmp4_url) {
						ϒtbr = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
							λ.StrLiteral("-(\\d+)k\\."),
							ϒmp4_url,
							λ.StrLiteral("bitrate"),
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
						})
						ϒformat_id = λ.StrLiteral("http")
						if λ.IsTrue(ϒtbr) {
							τmp0 = λ.IAdd(ϒformat_id, λ.Add(λ.StrLiteral("-"), ϒtbr))
							ϒformat_id = τmp0
						}
						λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
							"format_id": ϒformat_id,
							"url":       ϒmp4_url,
							"tbr":       λ.Cal(ϒint_or_none, ϒtbr),
						}))
					}
					if λ.IsTrue(ϒformats) {
						λ.Calm(ϒself, "_sort_formats", ϒformats)
						λ.SetItem(ϒinfo, λ.StrLiteral("formats"), ϒformats)
						λ.SetItem(ϒinfo, λ.StrLiteral("duration"), λ.Cal(ϒint_or_none, λ.Calm(ϒasset, "get", λ.StrLiteral("duration"))))
						return ϒinfo
					}
					τmp0 = λ.Cal(λ.BuiltinIter, λ.NewTuple(
						λ.StrLiteral("ooyala"),
						λ.StrLiteral("youtube"),
						λ.StrLiteral("brightcove"),
					))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒprovider_video_type = τmp1
						ϒprovider_video_id = λ.Calm(ϒvideo_data, "get", λ.Mod(λ.StrLiteral("%s_id"), ϒprovider_video_type))
						if !λ.IsTrue(ϒprovider_video_id) {
							continue
						}
						if λ.IsTrue(λ.Eq(ϒprovider_video_type, λ.StrLiteral("brightcove"))) {
							λ.SetItem(ϒinfo, λ.StrLiteral("formats"), λ.Calm(ϒself, "_extract_once_formats", ϒprovider_video_id))
							λ.Calm(ϒself, "_sort_formats", λ.GetItem(ϒinfo, λ.StrLiteral("formats")))
						} else {
							λ.Calm(ϒinfo, "update", λ.DictLiteral(map[string]λ.Object{
								"_type": λ.StrLiteral("url_transparent"),
								"url": func() λ.Object {
									if λ.IsTrue(λ.Eq(ϒprovider_video_type, λ.StrLiteral("youtube"))) {
										return ϒprovider_video_id
									} else {
										return λ.Mod(λ.StrLiteral("%s:%s"), λ.NewTuple(
											ϒprovider_video_type,
											ϒprovider_video_id,
										))
									}
								}(),
								"ie_key": λ.Calm(ϒprovider_video_type, "capitalize"),
							}))
						}
						return ϒinfo
					}
					panic(λ.Raise(λ.Cal(ExtractorError, λ.StrLiteral("Unable to find provider video id"))))
					return λ.None
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL":    VoxMediaVolumeIE__VALID_URL,
				"_real_extract": VoxMediaVolumeIE__real_extract,
			})
		}())
		VoxMediaIE = λ.Cal(λ.TypeType, λ.StrLiteral("VoxMediaIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				VoxMediaIE__VALID_URL    λ.Object
				VoxMediaIE__real_extract λ.Object
			)
			VoxMediaIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?(?:(?:theverge|vox|sbnation|eater|polygon|curbed|racked|funnyordie)\\.com|recode\\.net)/(?:[^/]+/)*(?P<id>[^/?]+)")
			VoxMediaIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒcreate_entry        λ.Object
						ϒdisplay_id          λ.Object
						ϒentries             λ.Object
						ϒentries_data        λ.Object
						ϒprovider_video_id   λ.Object
						ϒprovider_video_type λ.Object
						ϒself                = λargs[0]
						ϒurl                 = λargs[1]
						ϒvideo_data          λ.Object
						ϒvolume_uuid         λ.Object
						ϒwebpage             λ.Object
						τmp0                 λ.Object
						τmp1                 λ.Object
					)
					ϒdisplay_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒwebpage = λ.Cal(ϒcompat_urllib_parse_unquote, λ.Calm(ϒself, "_download_webpage", ϒurl, ϒdisplay_id))
					ϒcreate_entry = λ.NewFunction("create_entry",
						[]λ.Param{
							{Name: "provider_video_id"},
							{Name: "provider_video_type"},
							{Name: "title", Def: λ.None},
							{Name: "description", Def: λ.None},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒdescription         = λargs[3]
								ϒprovider_video_id   = λargs[0]
								ϒprovider_video_type = λargs[1]
								ϒtitle               = λargs[2]
								ϒvideo_url           λ.Object
							)
							ϒvideo_url = λ.Mod(λ.GetItem(λ.DictLiteral(map[string]string{
								"youtube": "%s",
								"ooyala":  "ooyala:%s",
								"volume":  "http://volume.vox-cdn.com/embed/%s",
							}), ϒprovider_video_type), ϒprovider_video_id)
							return λ.DictLiteral(map[string]λ.Object{
								"_type": λ.StrLiteral("url_transparent"),
								"url":   ϒvideo_url,
								"title": func() λ.Object {
									if λv := ϒtitle; λ.IsTrue(λv) {
										return λv
									} else {
										return λ.Calm(ϒself, "_og_search_title", ϒwebpage)
									}
								}(),
								"description": func() λ.Object {
									if λv := ϒdescription; λ.IsTrue(λv) {
										return λv
									} else {
										return λ.Calm(ϒself, "_og_search_description", ϒwebpage)
									}
								}(),
							})
						})
					ϒentries = λ.NewList()
					ϒentries_data = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewList(
							λ.StrLiteral("Chorus\\.VideoContext\\.addVideo\\((\\[{.+}\\])\\);"),
							λ.StrLiteral("var\\s+entry\\s*=\\s*({.+});"),
							λ.StrLiteral("SBN\\.VideoLinkset\\.entryGroup\\(\\s*(\\[.+\\])"),
						),
						ϒwebpage,
						λ.StrLiteral("video data"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					if λ.IsTrue(ϒentries_data) {
						ϒentries_data = λ.Calm(ϒself, "_parse_json", ϒentries_data, ϒdisplay_id)
						if λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒentries_data, λ.DictType)) {
							ϒentries_data = λ.NewList(ϒentries_data)
						}
						τmp0 = λ.Cal(λ.BuiltinIter, ϒentries_data)
						for {
							if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
								break
							}
							ϒvideo_data = τmp1
							ϒprovider_video_id = λ.Calm(ϒvideo_data, "get", λ.StrLiteral("provider_video_id"))
							ϒprovider_video_type = λ.Calm(ϒvideo_data, "get", λ.StrLiteral("provider_video_type"))
							if λ.IsTrue(func() λ.Object {
								if λv := ϒprovider_video_id; !λ.IsTrue(λv) {
									return λv
								} else {
									return ϒprovider_video_type
								}
							}()) {
								λ.Calm(ϒentries, "append", λ.Cal(ϒcreate_entry, ϒprovider_video_id, ϒprovider_video_type, λ.Calm(ϒvideo_data, "get", λ.StrLiteral("title")), λ.Calm(ϒvideo_data, "get", λ.StrLiteral("description"))))
							}
						}
					}
					ϒprovider_video_id = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("data-ooyala-id=\"([^\"]+)\""),
						ϒwebpage,
						λ.StrLiteral("ooyala id"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					if λ.IsTrue(ϒprovider_video_id) {
						λ.Calm(ϒentries, "append", λ.Cal(ϒcreate_entry, ϒprovider_video_id, λ.StrLiteral("ooyala")))
					}
					ϒvolume_uuid = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("data-volume-uuid=\"([^\"]+)\""),
						ϒwebpage,
						λ.StrLiteral("volume uuid"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					if λ.IsTrue(ϒvolume_uuid) {
						λ.Calm(ϒentries, "append", λ.Cal(ϒcreate_entry, ϒvolume_uuid, λ.StrLiteral("volume")))
					}
					if λ.IsTrue(λ.Eq(λ.Cal(λ.BuiltinLen, ϒentries), λ.IntLiteral(1))) {
						return λ.GetItem(ϒentries, λ.IntLiteral(0))
					} else {
						return λ.Calm(ϒself, "playlist_result", ϒentries, ϒdisplay_id, λ.Calm(ϒself, "_og_search_title", ϒwebpage), λ.Calm(ϒself, "_og_search_description", ϒwebpage))
					}
					return λ.None
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL":    VoxMediaIE__VALID_URL,
				"_real_extract": VoxMediaIE__real_extract,
			})
		}())
	})
}
