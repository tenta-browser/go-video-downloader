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
 * kinja/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/kinja.py
 */

package kinja

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor                λ.Object
	KinjaEmbedIE                 λ.Object
	ϒcompat_str                  λ.Object
	ϒcompat_urllib_parse_unquote λ.Object
	ϒint_or_none                 λ.Object
	ϒparse_iso8601               λ.Object
	ϒstrip_or_none               λ.Object
	ϒtry_get                     λ.Object
	ϒunescapeHTML                λ.Object
	ϒurljoin                     λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_str = Ωcompat.ϒcompat_str
		ϒcompat_urllib_parse_unquote = Ωcompat.ϒcompat_urllib_parse_unquote
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒparse_iso8601 = Ωutils.ϒparse_iso8601
		ϒstrip_or_none = Ωutils.ϒstrip_or_none
		ϒtry_get = Ωutils.ϒtry_get
		ϒunescapeHTML = Ωutils.ϒunescapeHTML
		ϒurljoin = Ωutils.ϒurljoin
		KinjaEmbedIE = λ.Cal(λ.TypeType, λ.StrLiteral("KinjaEmbedIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				KinjaEmbedIE__COMMON_REGEX        λ.Object
				KinjaEmbedIE__DOMAIN_REGEX        λ.Object
				KinjaEmbedIE__JWPLATFORM_PROVIDER λ.Object
				KinjaEmbedIE__PROVIDER_MAP        λ.Object
				KinjaEmbedIE__VALID_URL           λ.Object
				KinjaEmbedIE__real_extract        λ.Object
			)
			KinjaEmbedIE__DOMAIN_REGEX = λ.StrLiteral("(?:[^.]+\\.)?\n        (?:\n            avclub|\n            clickhole|\n            deadspin|\n            gizmodo|\n            jalopnik|\n            jezebel|\n            kinja|\n            kotaku|\n            lifehacker|\n            splinternews|\n            the(?:inventory|onion|root|takeout)\n        )\\.com")
			KinjaEmbedIE__COMMON_REGEX = λ.StrLiteral("/\n        (?:\n            ajax/inset|\n            embed/video\n        )/iframe\\?.*?\\bid=")
			KinjaEmbedIE__VALID_URL = λ.Mod(λ.StrLiteral("(?x)https?://%s%s\n        (?P<type>\n            fb|\n            imgur|\n            instagram|\n            jwp(?:layer)?-video|\n            kinjavideo|\n            mcp|\n            megaphone|\n            ooyala|\n            soundcloud(?:-playlist)?|\n            tumblr-post|\n            twitch-stream|\n            twitter|\n            ustream-channel|\n            vimeo|\n            vine|\n            youtube-(?:list|video)\n        )-(?P<id>[^&]+)"), λ.NewTuple(
				KinjaEmbedIE__DOMAIN_REGEX,
				KinjaEmbedIE__COMMON_REGEX,
			))
			KinjaEmbedIE__JWPLATFORM_PROVIDER = λ.NewTuple(
				λ.StrLiteral("cdn.jwplayer.com/v2/media/"),
				λ.StrLiteral("JWPlatform"),
			)
			KinjaEmbedIE__PROVIDER_MAP = λ.DictLiteral(map[string]λ.Object{
				"fb": λ.NewTuple(
					λ.StrLiteral("facebook.com/video.php?v="),
					λ.StrLiteral("Facebook"),
				),
				"imgur": λ.NewTuple(
					λ.StrLiteral("imgur.com/"),
					λ.StrLiteral("Imgur"),
				),
				"instagram": λ.NewTuple(
					λ.StrLiteral("instagram.com/p/"),
					λ.StrLiteral("Instagram"),
				),
				"jwplayer-video": KinjaEmbedIE__JWPLATFORM_PROVIDER,
				"jwp-video":      KinjaEmbedIE__JWPLATFORM_PROVIDER,
				"megaphone": λ.NewTuple(
					λ.StrLiteral("player.megaphone.fm/"),
					λ.StrLiteral("Generic"),
				),
				"ooyala": λ.NewTuple(
					λ.StrLiteral("player.ooyala.com/player.js?embedCode="),
					λ.StrLiteral("Ooyala"),
				),
				"soundcloud": λ.NewTuple(
					λ.StrLiteral("api.soundcloud.com/tracks/"),
					λ.StrLiteral("Soundcloud"),
				),
				"soundcloud-playlist": λ.NewTuple(
					λ.StrLiteral("api.soundcloud.com/playlists/"),
					λ.StrLiteral("SoundcloudPlaylist"),
				),
				"tumblr-post": λ.NewTuple(
					λ.StrLiteral("%s.tumblr.com/post/%s"),
					λ.StrLiteral("Tumblr"),
				),
				"twitch-stream": λ.NewTuple(
					λ.StrLiteral("twitch.tv/"),
					λ.StrLiteral("TwitchStream"),
				),
				"twitter": λ.NewTuple(
					λ.StrLiteral("twitter.com/i/cards/tfw/v1/"),
					λ.StrLiteral("TwitterCard"),
				),
				"ustream-channel": λ.NewTuple(
					λ.StrLiteral("ustream.tv/embed/"),
					λ.StrLiteral("Ustream"),
				),
				"vimeo": λ.NewTuple(
					λ.StrLiteral("vimeo.com/"),
					λ.StrLiteral("Vimeo"),
				),
				"vine": λ.NewTuple(
					λ.StrLiteral("vine.co/v/"),
					λ.StrLiteral("Vine"),
				),
				"youtube-list": λ.NewTuple(
					λ.StrLiteral("youtube.com/embed/%s?list=%s"),
					λ.StrLiteral("YoutubePlaylist"),
				),
				"youtube-video": λ.NewTuple(
					λ.StrLiteral("youtube.com/embed/"),
					λ.StrLiteral("Youtube"),
				),
			})
			KinjaEmbedIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒblog                   λ.Object
						ϒdata                   λ.Object
						ϒfallback_rendition_url λ.Object
						ϒfmg                    λ.Object
						ϒformats                λ.Object
						ϒiptc                   λ.Object
						ϒk                      λ.Object
						ϒm3u8_url               λ.Object
						ϒplaylist_id            λ.Object
						ϒposter                 λ.Object
						ϒposter_id              λ.Object
						ϒprovider               λ.Object
						ϒrendition_url          λ.Object
						ϒresult_url             λ.Object
						ϒself                   = λargs[0]
						ϒthumbnail              λ.Object
						ϒtitle                  λ.Object
						ϒtvss_domain            λ.Object
						ϒurl                    = λargs[1]
						ϒvideo_data             λ.Object
						ϒvideo_id               λ.Object
						ϒvideo_type             λ.Object
						τmp0                    λ.Object
						τmp1                    λ.Object
					)
					τmp0 = λ.UnpackIterable(λ.Calm(λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl), "groups"), 2)
					ϒvideo_type = λ.GetItem(τmp0, λ.IntLiteral(0))
					ϒvideo_id = λ.GetItem(τmp0, λ.IntLiteral(1))
					ϒprovider = λ.Calm(λ.GetAttr(ϒself, "_PROVIDER_MAP", nil), "get", ϒvideo_type)
					if λ.IsTrue(ϒprovider) {
						ϒvideo_id = λ.Cal(ϒcompat_urllib_parse_unquote, ϒvideo_id)
						if λ.IsTrue(λ.Eq(ϒvideo_type, λ.StrLiteral("tumblr-post"))) {
							τmp0 = λ.UnpackIterable(λ.Calm(ϒvideo_id, "split", λ.StrLiteral("-"), λ.IntLiteral(1)), 2)
							ϒvideo_id = λ.GetItem(τmp0, λ.IntLiteral(0))
							ϒblog = λ.GetItem(τmp0, λ.IntLiteral(1))
							ϒresult_url = λ.Mod(λ.GetItem(ϒprovider, λ.IntLiteral(0)), λ.NewTuple(
								ϒblog,
								ϒvideo_id,
							))
						} else {
							if λ.IsTrue(λ.Eq(ϒvideo_type, λ.StrLiteral("youtube-list"))) {
								τmp0 = λ.UnpackIterable(λ.Calm(ϒvideo_id, "split", λ.StrLiteral("/")), 2)
								ϒvideo_id = λ.GetItem(τmp0, λ.IntLiteral(0))
								ϒplaylist_id = λ.GetItem(τmp0, λ.IntLiteral(1))
								ϒresult_url = λ.Mod(λ.GetItem(ϒprovider, λ.IntLiteral(0)), λ.NewTuple(
									ϒvideo_id,
									ϒplaylist_id,
								))
							} else {
								if λ.IsTrue(λ.Eq(ϒvideo_type, λ.StrLiteral("ooyala"))) {
									ϒvideo_id = λ.GetItem(λ.Calm(ϒvideo_id, "split", λ.StrLiteral("/")), λ.IntLiteral(0))
								}
								ϒresult_url = λ.Add(λ.GetItem(ϒprovider, λ.IntLiteral(0)), ϒvideo_id)
							}
						}
						return λ.Calm(ϒself, "url_result", λ.Add(λ.StrLiteral("http://"), ϒresult_url), λ.GetItem(ϒprovider, λ.IntLiteral(1)))
					}
					if λ.IsTrue(λ.Eq(ϒvideo_type, λ.StrLiteral("kinjavideo"))) {
						ϒdata = λ.GetItem(λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
							λ.StrLiteral("https://kinja.com/api/core/video/views/videoById"),
							ϒvideo_id,
						), λ.KWArgs{
							{Name: "query", Value: λ.DictLiteral(map[string]λ.Object{
								"videoId": ϒvideo_id,
							})},
						}), λ.StrLiteral("data"))
						ϒtitle = λ.GetItem(ϒdata, λ.StrLiteral("title"))
						ϒformats = λ.NewList()
						τmp0 = λ.Cal(λ.BuiltinIter, λ.NewTuple(
							λ.StrLiteral("signedPlaylist"),
							λ.StrLiteral("streaming"),
						))
						for {
							if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
								break
							}
							ϒk = τmp1
							ϒm3u8_url = λ.Calm(ϒdata, "get", λ.Add(ϒk, λ.StrLiteral("Url")))
							if λ.IsTrue(ϒm3u8_url) {
								λ.Calm(ϒformats, "extend", λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
									ϒm3u8_url,
									ϒvideo_id,
									λ.StrLiteral("mp4"),
									λ.StrLiteral("m3u8_native"),
								), λ.KWArgs{
									{Name: "m3u8_id", Value: λ.StrLiteral("hls")},
									{Name: "fatal", Value: λ.False},
								}))
							}
						}
						λ.Calm(ϒself, "_sort_formats", ϒformats)
						ϒthumbnail = λ.None
						ϒposter = func() λ.Object {
							if λv := λ.Calm(ϒdata, "get", λ.StrLiteral("poster")); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.DictLiteral(map[λ.Object]λ.Object{})
							}
						}()
						ϒposter_id = λ.Calm(ϒposter, "get", λ.StrLiteral("id"))
						if λ.IsTrue(ϒposter_id) {
							ϒthumbnail = λ.Mod(λ.StrLiteral("https://i.kinja-img.com/gawker-media/image/upload/%s.%s"), λ.NewTuple(
								ϒposter_id,
								func() λ.Object {
									if λv := λ.Calm(ϒposter, "get", λ.StrLiteral("format")); λ.IsTrue(λv) {
										return λv
									} else {
										return λ.StrLiteral("jpg")
									}
								}(),
							))
						}
						return λ.DictLiteral(map[string]λ.Object{
							"id":          ϒvideo_id,
							"title":       ϒtitle,
							"description": λ.Cal(ϒstrip_or_none, λ.Calm(ϒdata, "get", λ.StrLiteral("description"))),
							"formats":     ϒformats,
							"tags":        λ.Calm(ϒdata, "get", λ.StrLiteral("tags")),
							"timestamp": λ.Cal(ϒint_or_none, λ.Cal(ϒtry_get, ϒdata, λ.NewFunction("<lambda>",
								[]λ.Param{
									{Name: "x"},
								},
								0, false, false,
								func(λargs []λ.Object) λ.Object {
									var (
										ϒx = λargs[0]
									)
									return λ.GetItem(λ.GetItem(ϒx, λ.StrLiteral("postInfo")), λ.StrLiteral("publishTimeMillis"))
								})), λ.IntLiteral(1000)),
							"thumbnail": ϒthumbnail,
							"uploader":  λ.Calm(ϒdata, "get", λ.StrLiteral("network")),
						})
					} else {
						ϒvideo_data = λ.GetItem(λ.Calm(ϒself, "_download_json", λ.Add(λ.StrLiteral("https://api.vmh.univision.com/metadata/v1/content/"), ϒvideo_id), ϒvideo_id), λ.StrLiteral("videoMetadata"))
						ϒiptc = λ.GetItem(ϒvideo_data, λ.StrLiteral("photoVideoMetadataIPTC"))
						ϒtitle = λ.GetItem(λ.GetItem(ϒiptc, λ.StrLiteral("title")), λ.StrLiteral("en"))
						ϒfmg = func() λ.Object {
							if λv := λ.Calm(ϒvideo_data, "get", λ.StrLiteral("photoVideoMetadata_fmg")); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.DictLiteral(map[λ.Object]λ.Object{})
							}
						}()
						ϒtvss_domain = func() λ.Object {
							if λv := λ.Calm(ϒfmg, "get", λ.StrLiteral("tvssDomain")); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.StrLiteral("https://auth.univision.com")
							}
						}()
						ϒdata = λ.GetItem(λ.GetItem(λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
							λ.Add(ϒtvss_domain, λ.StrLiteral("/api/v3/video-auth/url-signature-tokens")),
							ϒvideo_id,
						), λ.KWArgs{
							{Name: "query", Value: λ.DictLiteral(map[string]λ.Object{
								"mcpids": ϒvideo_id,
							})},
						}), λ.StrLiteral("data")), λ.IntLiteral(0))
						ϒformats = λ.NewList()
						ϒrendition_url = λ.Calm(ϒdata, "get", λ.StrLiteral("renditionUrl"))
						if λ.IsTrue(ϒrendition_url) {
							ϒformats = λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
								ϒrendition_url,
								ϒvideo_id,
								λ.StrLiteral("mp4"),
								λ.StrLiteral("m3u8_native"),
							), λ.KWArgs{
								{Name: "m3u8_id", Value: λ.StrLiteral("hls")},
								{Name: "fatal", Value: λ.False},
							})
						}
						ϒfallback_rendition_url = λ.Calm(ϒdata, "get", λ.StrLiteral("fallbackRenditionUrl"))
						if λ.IsTrue(ϒfallback_rendition_url) {
							λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
								"format_id": λ.StrLiteral("fallback"),
								"tbr": λ.Cal(ϒint_or_none, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
									λ.StrLiteral("_(\\d+)\\.mp4"),
									ϒfallback_rendition_url,
									λ.StrLiteral("bitrate"),
								), λ.KWArgs{
									{Name: "default", Value: λ.None},
								})),
								"url": ϒfallback_rendition_url,
							}))
						}
						λ.Calm(ϒself, "_sort_formats", ϒformats)
						return λ.DictLiteral(map[string]λ.Object{
							"id":    ϒvideo_id,
							"title": ϒtitle,
							"thumbnail": λ.Cal(ϒtry_get, ϒiptc, λ.NewFunction("<lambda>",
								[]λ.Param{
									{Name: "x"},
								},
								0, false, false,
								func(λargs []λ.Object) λ.Object {
									var (
										ϒx = λargs[0]
									)
									return λ.GetItem(λ.GetItem(ϒx, λ.StrLiteral("cloudinaryLink")), λ.StrLiteral("link"))
								}), ϒcompat_str),
							"uploader": λ.Calm(ϒfmg, "get", λ.StrLiteral("network")),
							"duration": λ.Cal(ϒint_or_none, λ.Calm(ϒiptc, "get", λ.StrLiteral("fileDuration"))),
							"formats":  ϒformats,
							"description": λ.Cal(ϒtry_get, ϒiptc, λ.NewFunction("<lambda>",
								[]λ.Param{
									{Name: "x"},
								},
								0, false, false,
								func(λargs []λ.Object) λ.Object {
									var (
										ϒx = λargs[0]
									)
									return λ.GetItem(λ.GetItem(ϒx, λ.StrLiteral("description")), λ.StrLiteral("en"))
								}), ϒcompat_str),
							"timestamp": λ.Cal(ϒparse_iso8601, λ.Calm(ϒiptc, "get", λ.StrLiteral("dateReleased"))),
						})
					}
					return λ.None
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_COMMON_REGEX":        KinjaEmbedIE__COMMON_REGEX,
				"_DOMAIN_REGEX":        KinjaEmbedIE__DOMAIN_REGEX,
				"_JWPLATFORM_PROVIDER": KinjaEmbedIE__JWPLATFORM_PROVIDER,
				"_PROVIDER_MAP":        KinjaEmbedIE__PROVIDER_MAP,
				"_VALID_URL":           KinjaEmbedIE__VALID_URL,
				"_real_extract":        KinjaEmbedIE__real_extract,
			})
		}())
	})
}
