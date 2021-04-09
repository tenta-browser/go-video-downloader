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
 * minds/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/minds.py
 */

package minds

import (
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor   λ.Object
	MindsBaseIE     λ.Object
	MindsChannelIE  λ.Object
	MindsFeedBaseIE λ.Object
	MindsGroupIE    λ.Object
	MindsIE         λ.Object
	ϒclean_html     λ.Object
	ϒcompat_str     λ.Object
	ϒint_or_none    λ.Object
	ϒstr_or_none    λ.Object
	ϒstrip_or_none  λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_str = Ωcompat.ϒcompat_str
		ϒclean_html = Ωutils.ϒclean_html
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒstr_or_none = Ωutils.ϒstr_or_none
		ϒstrip_or_none = Ωutils.ϒstrip_or_none
		MindsBaseIE = λ.Cal(λ.TypeType, λ.StrLiteral("MindsBaseIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				MindsBaseIE__VALID_URL_BASE λ.Object
				MindsBaseIE__call_api       λ.Object
			)
			MindsBaseIE__VALID_URL_BASE = λ.StrLiteral("https?://(?:www\\.)?minds\\.com/")
			MindsBaseIE__call_api = λ.NewFunction("_call_api",
				[]λ.Param{
					{Name: "self"},
					{Name: "path"},
					{Name: "video_id"},
					{Name: "resource"},
					{Name: "query", Def: λ.None},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒapi_url  λ.Object
						ϒpath     = λargs[1]
						ϒquery    = λargs[4]
						ϒresource = λargs[3]
						ϒself     = λargs[0]
						ϒtoken    λ.Object
						ϒvideo_id = λargs[2]
					)
					ϒapi_url = λ.Add(λ.StrLiteral("https://www.minds.com/api/"), ϒpath)
					ϒtoken = λ.Calm(λ.Calm(ϒself, "_get_cookies", ϒapi_url), "get", λ.StrLiteral("XSRF-TOKEN"))
					return λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
						ϒapi_url,
						ϒvideo_id,
						λ.Mod(λ.StrLiteral("Downloading %s JSON metadata"), ϒresource),
					), λ.KWArgs{
						{Name: "headers", Value: λ.DictLiteral(map[string]λ.Object{
							"Referer": λ.StrLiteral("https://www.minds.com/"),
							"X-XSRF-TOKEN": func() λ.Object {
								if λ.IsTrue(ϒtoken) {
									return λ.GetAttr(ϒtoken, "value", nil)
								} else {
									return λ.StrLiteral("")
								}
							}(),
						})},
						{Name: "query", Value: ϒquery},
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL_BASE": MindsBaseIE__VALID_URL_BASE,
				"_call_api":       MindsBaseIE__call_api,
			})
		}())
		MindsIE = λ.Cal(λ.TypeType, λ.StrLiteral("MindsIE"), λ.NewTuple(MindsBaseIE), func() λ.Dict {
			var (
				MindsIE_IE_NAME       λ.Object
				MindsIE__VALID_URL    λ.Object
				MindsIE__real_extract λ.Object
			)
			MindsIE_IE_NAME = λ.StrLiteral("minds")
			MindsIE__VALID_URL = λ.Add(λ.GetAttr(MindsBaseIE, "_VALID_URL_BASE", nil), λ.StrLiteral("(?:media|newsfeed|archive/view)/(?P<id>[0-9]+)"))
			MindsIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒentity      λ.Object
						ϒentity_id   λ.Object
						ϒformats     λ.Object
						ϒowner       λ.Object
						ϒposter      λ.Object
						ϒself        = λargs[0]
						ϒsource      λ.Object
						ϒsrc         λ.Object
						ϒtags        λ.Object
						ϒthumbnail   λ.Object
						ϒuploader_id λ.Object
						ϒurl         = λargs[1]
						ϒurlh        λ.Object
						ϒvideo       λ.Object
						ϒvideo_id    λ.Object
						τmp0         λ.Object
						τmp1         λ.Object
					)
					ϒentity_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒentity = λ.GetItem(λ.Calm(ϒself, "_call_api", λ.Add(λ.StrLiteral("v1/entities/entity/"), ϒentity_id), ϒentity_id, λ.StrLiteral("entity")), λ.StrLiteral("entity"))
					if λ.IsTrue(λ.Eq(λ.Calm(ϒentity, "get", λ.StrLiteral("type")), λ.StrLiteral("activity"))) {
						if λ.IsTrue(λ.Eq(λ.Calm(ϒentity, "get", λ.StrLiteral("custom_type")), λ.StrLiteral("video"))) {
							ϒvideo_id = λ.GetItem(ϒentity, λ.StrLiteral("entity_guid"))
						} else {
							return λ.Calm(ϒself, "url_result", λ.GetItem(ϒentity, λ.StrLiteral("perma_url")))
						}
					} else {
						if !λ.IsTrue(λ.Eq(λ.GetItem(ϒentity, λ.StrLiteral("subtype")), λ.StrLiteral("video"))) {
							panic(λ.Raise(λ.Cal(λ.AssertionErrorType)))
						}
						ϒvideo_id = ϒentity_id
					}
					ϒvideo = λ.Calm(ϒself, "_call_api", λ.Add(λ.StrLiteral("v2/media/video/"), ϒvideo_id), ϒvideo_id, λ.StrLiteral("video"))
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, func() λ.Object {
						if λv := λ.Calm(ϒvideo, "get", λ.StrLiteral("sources")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.NewList()
						}
					}())
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒsource = τmp1
						ϒsrc = λ.Calm(ϒsource, "get", λ.StrLiteral("src"))
						if !λ.IsTrue(ϒsrc) {
							continue
						}
						λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
							"format_id": λ.Calm(ϒsource, "get", λ.StrLiteral("label")),
							"height":    λ.Cal(ϒint_or_none, λ.Calm(ϒsource, "get", λ.StrLiteral("size"))),
							"url":       ϒsrc,
						}))
					}
					λ.Calm(ϒself, "_sort_formats", ϒformats)
					ϒentity = func() λ.Object {
						if λv := λ.Calm(ϒvideo, "get", λ.StrLiteral("entity")); λ.IsTrue(λv) {
							return λv
						} else {
							return ϒentity
						}
					}()
					ϒowner = func() λ.Object {
						if λv := λ.Calm(ϒentity, "get", λ.StrLiteral("ownerObj")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.DictLiteral(map[λ.Object]λ.Object{})
						}
					}()
					ϒuploader_id = λ.Calm(ϒowner, "get", λ.StrLiteral("username"))
					ϒtags = λ.Calm(ϒentity, "get", λ.StrLiteral("tags"))
					if λ.IsTrue(func() λ.Object {
						if λv := ϒtags; !λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Cal(λ.BuiltinIsInstance, ϒtags, ϒcompat_str)
						}
					}()) {
						ϒtags = λ.NewList(ϒtags)
					}
					ϒthumbnail = λ.None
					ϒposter = func() λ.Object {
						if λv := λ.Calm(ϒvideo, "get", λ.StrLiteral("poster")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Calm(ϒentity, "get", λ.StrLiteral("thumbnail_src"))
						}
					}()
					if λ.IsTrue(ϒposter) {
						ϒurlh = λ.Call(λ.GetAttr(ϒself, "_request_webpage", nil), λ.NewArgs(
							ϒposter,
							ϒvideo_id,
						), λ.KWArgs{
							{Name: "fatal", Value: λ.False},
						})
						if λ.IsTrue(ϒurlh) {
							ϒthumbnail = λ.Calm(ϒurlh, "geturl")
						}
					}
					return λ.DictLiteral(map[string]λ.Object{
						"id": ϒvideo_id,
						"title": func() λ.Object {
							if λv := λ.Calm(ϒentity, "get", λ.StrLiteral("title")); λ.IsTrue(λv) {
								return λv
							} else {
								return ϒvideo_id
							}
						}(),
						"formats": ϒformats,
						"description": func() λ.Object {
							if λv := λ.Cal(ϒclean_html, λ.Calm(ϒentity, "get", λ.StrLiteral("description"))); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.None
							}
						}(),
						"license":     λ.Cal(ϒstr_or_none, λ.Calm(ϒentity, "get", λ.StrLiteral("license"))),
						"timestamp":   λ.Cal(ϒint_or_none, λ.Calm(ϒentity, "get", λ.StrLiteral("time_created"))),
						"uploader":    λ.Cal(ϒstrip_or_none, λ.Calm(ϒowner, "get", λ.StrLiteral("name"))),
						"uploader_id": ϒuploader_id,
						"uploader_url": func() λ.Object {
							if λ.IsTrue(ϒuploader_id) {
								return λ.Add(λ.StrLiteral("https://www.minds.com/"), ϒuploader_id)
							} else {
								return λ.None
							}
						}(),
						"view_count":    λ.Cal(ϒint_or_none, λ.Calm(ϒentity, "get", λ.StrLiteral("play:count"))),
						"like_count":    λ.Cal(ϒint_or_none, λ.Calm(ϒentity, "get", λ.StrLiteral("thumbs:up:count"))),
						"dislike_count": λ.Cal(ϒint_or_none, λ.Calm(ϒentity, "get", λ.StrLiteral("thumbs:down:count"))),
						"tags":          ϒtags,
						"comment_count": λ.Cal(ϒint_or_none, λ.Calm(ϒentity, "get", λ.StrLiteral("comments:count"))),
						"thumbnail":     ϒthumbnail,
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"IE_NAME":       MindsIE_IE_NAME,
				"_VALID_URL":    MindsIE__VALID_URL,
				"_real_extract": MindsIE__real_extract,
			})
		}())
		MindsFeedBaseIE = λ.Cal(λ.TypeType, λ.StrLiteral("MindsFeedBaseIE"), λ.NewTuple(MindsBaseIE), func() λ.Dict {

			return λ.ClassDictLiteral(map[λ.Object]λ.Object{})
		}())
		MindsChannelIE = λ.Cal(λ.TypeType, λ.StrLiteral("MindsChannelIE"), λ.NewTuple(MindsFeedBaseIE), func() λ.Dict {
			var (
				MindsChannelIE__FEED_TYPE λ.Object
				MindsChannelIE__VALID_URL λ.Object
			)
			MindsChannelIE__FEED_TYPE = λ.StrLiteral("channel")
			MindsChannelIE__VALID_URL = λ.Add(λ.GetAttr(MindsBaseIE, "_VALID_URL_BASE", nil), λ.StrLiteral("(?!(?:newsfeed|media|api|archive|groups)/)(?P<id>[^/?&#]+)"))
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_FEED_TYPE": MindsChannelIE__FEED_TYPE,
				"_VALID_URL": MindsChannelIE__VALID_URL,
			})
		}())
		MindsGroupIE = λ.Cal(λ.TypeType, λ.StrLiteral("MindsGroupIE"), λ.NewTuple(MindsFeedBaseIE), func() λ.Dict {
			var (
				MindsGroupIE__FEED_TYPE λ.Object
				MindsGroupIE__VALID_URL λ.Object
			)
			MindsGroupIE__FEED_TYPE = λ.StrLiteral("group")
			MindsGroupIE__VALID_URL = λ.Add(λ.GetAttr(MindsBaseIE, "_VALID_URL_BASE", nil), λ.StrLiteral("groups/profile/(?P<id>[0-9]+)"))
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_FEED_TYPE": MindsGroupIE__FEED_TYPE,
				"_VALID_URL": MindsGroupIE__VALID_URL,
			})
		}())
	})
}
