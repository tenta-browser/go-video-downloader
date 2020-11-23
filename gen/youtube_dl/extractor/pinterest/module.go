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
 * pinterest/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/pinterest.py
 */

package pinterest

import (
	Ωjson "github.com/tenta-browser/go-video-downloader/gen/json"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor         λ.Object
	PinterestBaseIE       λ.Object
	PinterestCollectionIE λ.Object
	PinterestIE           λ.Object
	ϒcompat_str           λ.Object
	ϒdetermine_ext        λ.Object
	ϒfloat_or_none        λ.Object
	ϒint_or_none          λ.Object
	ϒtry_get              λ.Object
	ϒunified_timestamp    λ.Object
	ϒurl_or_none          λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_str = Ωcompat.ϒcompat_str
		ϒdetermine_ext = Ωutils.ϒdetermine_ext
		ϒfloat_or_none = Ωutils.ϒfloat_or_none
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒtry_get = Ωutils.ϒtry_get
		ϒunified_timestamp = Ωutils.ϒunified_timestamp
		ϒurl_or_none = Ωutils.ϒurl_or_none
		PinterestBaseIE = λ.Cal(λ.TypeType, λ.StrLiteral("PinterestBaseIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				PinterestBaseIE__VALID_URL_BASE λ.Object
				PinterestBaseIE__call_api       λ.Object
				PinterestBaseIE__extract_video  λ.Object
			)
			PinterestBaseIE__VALID_URL_BASE = λ.StrLiteral("https?://(?:[^/]+\\.)?pinterest\\.(?:com|fr|de|ch|jp|cl|ca|it|co\\.uk|nz|ru|com\\.au|at|pt|co\\.kr|es|com\\.mx|dk|ph|th|com\\.uy|co|nl|info|kr|ie|vn|com\\.vn|ec|mx|in|pe|co\\.at|hu|co\\.in|co\\.nz|id|com\\.ec|com\\.py|tw|be|uk|com\\.bo|com\\.pe)")
			PinterestBaseIE__call_api = λ.NewFunction("_call_api",
				[]λ.Param{
					{Name: "self"},
					{Name: "resource"},
					{Name: "video_id"},
					{Name: "options"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒoptions  = λargs[3]
						ϒresource = λargs[1]
						ϒself     = λargs[0]
						ϒvideo_id = λargs[2]
					)
					return λ.GetItem(λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
						λ.Mod(λ.StrLiteral("https://www.pinterest.com/resource/%sResource/get/"), ϒresource),
						ϒvideo_id,
						λ.Mod(λ.StrLiteral("Download %s JSON metadata"), ϒresource),
					), λ.KWArgs{
						{Name: "query", Value: λ.DictLiteral(map[string]λ.Object{
							"data": λ.Cal(Ωjson.ϒdumps, λ.DictLiteral(map[string]λ.Object{
								"options": ϒoptions,
							})),
						})},
					}), λ.StrLiteral("resource_response"))
				})
			PinterestBaseIE__extract_video = λ.NewFunction("_extract_video",
				[]λ.Param{
					{Name: "self"},
					{Name: "data"},
					{Name: "extract_formats", Def: λ.True},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒ_u              λ.Object
						ϒcategories      λ.Object
						ϒcomment_count   λ.Object
						ϒdata            = λargs[1]
						ϒdescription     λ.Object
						ϒduration        λ.Object
						ϒext             λ.Object
						ϒextract_formats = λargs[2]
						ϒformat_dict     λ.Object
						ϒformat_id       λ.Object
						ϒformat_url      λ.Object
						ϒformats         λ.Object
						ϒimages          λ.Object
						ϒrepost_count    λ.Object
						ϒself            = λargs[0]
						ϒtags            λ.Object
						ϒthumbnail       λ.Object
						ϒthumbnail_id    λ.Object
						ϒthumbnail_url   λ.Object
						ϒthumbnails      λ.Object
						ϒtimestamp       λ.Object
						ϒtitle           λ.Object
						ϒuploader        λ.Object
						ϒuploader_id     λ.Object
						ϒvideo_id        λ.Object
						τmp0             λ.Object
						τmp1             λ.Object
						τmp2             λ.Object
					)
					_ = ϒthumbnail_id
					ϒvideo_id = λ.GetItem(ϒdata, λ.StrLiteral("id"))
					ϒtitle = λ.Calm(func() λ.Object {
						if λv := λ.Calm(ϒdata, "get", λ.StrLiteral("title")); λ.IsTrue(λv) {
							return λv
						} else if λv := λ.Calm(ϒdata, "get", λ.StrLiteral("grid_title")); λ.IsTrue(λv) {
							return λv
						} else {
							return ϒvideo_id
						}
					}(), "strip")
					ϒformats = λ.NewList()
					ϒduration = λ.None
					if λ.IsTrue(ϒextract_formats) {
						τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(λ.GetItem(λ.GetItem(ϒdata, λ.StrLiteral("videos")), λ.StrLiteral("video_list")), "items"))
						for {
							if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
								break
							}
							τmp2 = τmp1
							ϒformat_id = λ.GetItem(τmp2, λ.IntLiteral(0))
							ϒformat_dict = λ.GetItem(τmp2, λ.IntLiteral(1))
							if !λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒformat_dict, λ.DictType)) {
								continue
							}
							ϒformat_url = λ.Cal(ϒurl_or_none, λ.Calm(ϒformat_dict, "get", λ.StrLiteral("url")))
							if !λ.IsTrue(ϒformat_url) {
								continue
							}
							ϒduration = λ.Call(ϒfloat_or_none, λ.NewArgs(λ.Calm(ϒformat_dict, "get", λ.StrLiteral("duration"))), λ.KWArgs{
								{Name: "scale", Value: λ.IntLiteral(1000)},
							})
							ϒext = λ.Cal(ϒdetermine_ext, ϒformat_url)
							if λ.IsTrue(func() λ.Object {
								if λv := λ.NewBool(λ.Contains(λ.Calm(ϒformat_id, "lower"), λ.StrLiteral("hls"))); λ.IsTrue(λv) {
									return λv
								} else {
									return λ.Eq(ϒext, λ.StrLiteral("m3u8"))
								}
							}()) {
								λ.Calm(ϒformats, "extend", λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
									ϒformat_url,
									ϒvideo_id,
									λ.StrLiteral("mp4"),
								), λ.KWArgs{
									{Name: "entry_protocol", Value: λ.StrLiteral("m3u8_native")},
									{Name: "m3u8_id", Value: ϒformat_id},
									{Name: "fatal", Value: λ.False},
								}))
							} else {
								λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
									"url":       ϒformat_url,
									"format_id": ϒformat_id,
									"width":     λ.Cal(ϒint_or_none, λ.Calm(ϒformat_dict, "get", λ.StrLiteral("width"))),
									"height":    λ.Cal(ϒint_or_none, λ.Calm(ϒformat_dict, "get", λ.StrLiteral("height"))),
									"duration":  ϒduration,
								}))
							}
						}
						λ.Call(λ.GetAttr(ϒself, "_sort_formats", nil), λ.NewArgs(ϒformats), λ.KWArgs{
							{Name: "field_preference", Value: λ.NewTuple(
								λ.StrLiteral("height"),
								λ.StrLiteral("width"),
								λ.StrLiteral("tbr"),
								λ.StrLiteral("format_id"),
							)},
						})
					}
					ϒdescription = func() λ.Object {
						if λv := λ.Calm(ϒdata, "get", λ.StrLiteral("description")); λ.IsTrue(λv) {
							return λv
						} else if λv := λ.Calm(ϒdata, "get", λ.StrLiteral("description_html")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Calm(ϒdata, "get", λ.StrLiteral("seo_description"))
						}
					}()
					ϒtimestamp = λ.Cal(ϒunified_timestamp, λ.Calm(ϒdata, "get", λ.StrLiteral("created_at")))
					ϒ_u = λ.NewFunction("_u",
						[]λ.Param{
							{Name: "field"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒfield = λargs[0]
							)
							return λ.Cal(ϒtry_get, ϒdata, λ.NewFunction("<lambda>",
								[]λ.Param{
									{Name: "x"},
								},
								0, false, false,
								func(λargs []λ.Object) λ.Object {
									var (
										ϒx = λargs[0]
									)
									return λ.GetItem(λ.GetItem(ϒx, λ.StrLiteral("closeup_attribution")), ϒfield)
								}), ϒcompat_str)
						})
					ϒuploader = λ.Cal(ϒ_u, λ.StrLiteral("full_name"))
					ϒuploader_id = λ.Cal(ϒ_u, λ.StrLiteral("id"))
					ϒrepost_count = λ.Cal(ϒint_or_none, λ.Calm(ϒdata, "get", λ.StrLiteral("repin_count")))
					ϒcomment_count = λ.Cal(ϒint_or_none, λ.Calm(ϒdata, "get", λ.StrLiteral("comment_count")))
					ϒcategories = λ.Cal(ϒtry_get, ϒdata, λ.NewFunction("<lambda>",
						[]λ.Param{
							{Name: "x"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒx = λargs[0]
							)
							return λ.GetItem(λ.GetItem(ϒx, λ.StrLiteral("pin_join")), λ.StrLiteral("visual_annotation"))
						}), λ.ListType)
					ϒtags = λ.Calm(ϒdata, "get", λ.StrLiteral("hashtags"))
					ϒthumbnails = λ.NewList()
					ϒimages = λ.Calm(ϒdata, "get", λ.StrLiteral("images"))
					if λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒimages, λ.DictType)) {
						τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(ϒimages, "items"))
						for {
							if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
								break
							}
							τmp2 = τmp1
							ϒthumbnail_id = λ.GetItem(τmp2, λ.IntLiteral(0))
							ϒthumbnail = λ.GetItem(τmp2, λ.IntLiteral(1))
							if !λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒthumbnail, λ.DictType)) {
								continue
							}
							ϒthumbnail_url = λ.Cal(ϒurl_or_none, λ.Calm(ϒthumbnail, "get", λ.StrLiteral("url")))
							if !λ.IsTrue(ϒthumbnail_url) {
								continue
							}
							λ.Calm(ϒthumbnails, "append", λ.DictLiteral(map[string]λ.Object{
								"url":    ϒthumbnail_url,
								"width":  λ.Cal(ϒint_or_none, λ.Calm(ϒthumbnail, "get", λ.StrLiteral("width"))),
								"height": λ.Cal(ϒint_or_none, λ.Calm(ϒthumbnail, "get", λ.StrLiteral("height"))),
							}))
						}
					}
					return λ.DictLiteral(map[string]λ.Object{
						"id":            ϒvideo_id,
						"title":         ϒtitle,
						"description":   ϒdescription,
						"duration":      ϒduration,
						"timestamp":     ϒtimestamp,
						"thumbnails":    ϒthumbnails,
						"uploader":      ϒuploader,
						"uploader_id":   ϒuploader_id,
						"repost_count":  ϒrepost_count,
						"comment_count": ϒcomment_count,
						"categories":    ϒcategories,
						"tags":          ϒtags,
						"formats":       ϒformats,
						"extractor_key": λ.Calm(PinterestIE, "ie_key"),
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL_BASE": PinterestBaseIE__VALID_URL_BASE,
				"_call_api":       PinterestBaseIE__call_api,
				"_extract_video":  PinterestBaseIE__extract_video,
			})
		}())
		PinterestIE = λ.Cal(λ.TypeType, λ.StrLiteral("PinterestIE"), λ.NewTuple(PinterestBaseIE), func() λ.Dict {
			var (
				PinterestIE__VALID_URL    λ.Object
				PinterestIE__real_extract λ.Object
			)
			PinterestIE__VALID_URL = λ.Mod(λ.StrLiteral("%s/pin/(?P<id>\\d+)"), λ.GetAttr(PinterestBaseIE, "_VALID_URL_BASE", nil))
			PinterestIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdata     λ.Object
						ϒself     = λargs[0]
						ϒurl      = λargs[1]
						ϒvideo_id λ.Object
					)
					ϒvideo_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒdata = λ.GetItem(λ.Calm(ϒself, "_call_api", λ.StrLiteral("Pin"), ϒvideo_id, λ.DictLiteral(map[string]λ.Object{
						"field_set_key": λ.StrLiteral("unauth_react_main_pin"),
						"id":            ϒvideo_id,
					})), λ.StrLiteral("data"))
					return λ.Calm(ϒself, "_extract_video", ϒdata)
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL":    PinterestIE__VALID_URL,
				"_real_extract": PinterestIE__real_extract,
			})
		}())
		PinterestCollectionIE = λ.Cal(λ.TypeType, λ.StrLiteral("PinterestCollectionIE"), λ.NewTuple(PinterestBaseIE), func() λ.Dict {
			var (
				PinterestCollectionIE__VALID_URL λ.Object
				PinterestCollectionIE_suitable   λ.Object
			)
			PinterestCollectionIE__VALID_URL = λ.Mod(λ.StrLiteral("%s/(?P<username>[^/]+)/(?P<id>[^/?#&]+)"), λ.GetAttr(PinterestBaseIE, "_VALID_URL_BASE", nil))
			PinterestCollectionIE_suitable = λ.NewFunction("suitable",
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
						if λ.IsTrue(λ.Calm(PinterestIE, "suitable", ϒurl)) {
							return λ.False
						} else {
							return λ.Calm(λ.Cal(λ.SuperType, PinterestCollectionIE, ϒcls), "suitable", ϒurl)
						}
					}()
				})
			PinterestCollectionIE_suitable = λ.Cal(λ.ClassMethodType, PinterestCollectionIE_suitable)
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": PinterestCollectionIE__VALID_URL,
				"suitable":   PinterestCollectionIE_suitable,
			})
		}())
	})
}
