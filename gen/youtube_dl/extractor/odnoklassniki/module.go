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
 * odnoklassniki/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/odnoklassniki.py
 */

package odnoklassniki

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError                λ.Object
	InfoExtractor                 λ.Object
	OdnoklassnikiIE               λ.Object
	ϒcompat_etree_fromstring      λ.Object
	ϒcompat_parse_qs              λ.Object
	ϒcompat_urllib_parse_unquote  λ.Object
	ϒcompat_urllib_parse_urlparse λ.Object
	ϒint_or_none                  λ.Object
	ϒqualities                    λ.Object
	ϒunescapeHTML                 λ.Object
	ϒunified_strdate              λ.Object
	ϒurlencode_postdata           λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_etree_fromstring = Ωcompat.ϒcompat_etree_fromstring
		ϒcompat_parse_qs = Ωcompat.ϒcompat_parse_qs
		ϒcompat_urllib_parse_unquote = Ωcompat.ϒcompat_urllib_parse_unquote
		ϒcompat_urllib_parse_urlparse = Ωcompat.ϒcompat_urllib_parse_urlparse
		ExtractorError = Ωutils.ExtractorError
		ϒunified_strdate = Ωutils.ϒunified_strdate
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒqualities = Ωutils.ϒqualities
		ϒunescapeHTML = Ωutils.ϒunescapeHTML
		ϒurlencode_postdata = Ωutils.ϒurlencode_postdata
		OdnoklassnikiIE = λ.Cal(λ.TypeType, λ.StrLiteral("OdnoklassnikiIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				OdnoklassnikiIE__VALID_URL    λ.Object
				OdnoklassnikiIE__extract_url  λ.Object
				OdnoklassnikiIE__real_extract λ.Object
			)
			OdnoklassnikiIE__VALID_URL = λ.StrLiteral("(?x)\n                https?://\n                    (?:(?:www|m|mobile)\\.)?\n                    (?:odnoklassniki|ok)\\.ru/\n                    (?:\n                        video(?:embed)?/|\n                        web-api/video/moviePlayer/|\n                        live/|\n                        dk\\?.*?st\\.mvId=\n                    )\n                    (?P<id>[\\d-]+)\n                ")
			OdnoklassnikiIE__extract_url = λ.NewFunction("_extract_url",
				[]λ.Param{
					{Name: "webpage"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒmobj    λ.Object
						ϒwebpage = λargs[0]
					)
					ϒmobj = λ.Cal(Ωre.ϒsearch, λ.StrLiteral("<iframe[^>]+src=([\"\\'])(?P<url>(?:https?:)?//(?:odnoklassniki|ok)\\.ru/videoembed/.+?)\\1"), ϒwebpage)
					if λ.IsTrue(ϒmobj) {
						return λ.Calm(ϒmobj, "group", λ.StrLiteral("url"))
					}
					return λ.None
				})
			OdnoklassnikiIE__extract_url = λ.Cal(λ.StaticMethodType, OdnoklassnikiIE__extract_url)
			OdnoklassnikiIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒadult         λ.Object
						ϒage_limit     λ.Object
						ϒauthor        λ.Object
						ϒdash_manifest λ.Object
						ϒdata          λ.Object
						ϒduration      λ.Object
						ϒerror         λ.Object
						ϒflashvars     λ.Object
						ϒfmt           λ.Object
						ϒfmt_type      λ.Object
						ϒformats       λ.Object
						ϒinfo          λ.Object
						ϒlike_count    λ.Object
						ϒm3u8_url      λ.Object
						ϒmetadata      λ.Object
						ϒmovie         λ.Object
						ϒpayment_info  λ.Object
						ϒplayer        λ.Object
						ϒprovider      λ.Object
						ϒquality       λ.Object
						ϒrtmp_url      λ.Object
						ϒself          = λargs[0]
						ϒst_location   λ.Object
						ϒstart_time    λ.Object
						ϒthumbnail     λ.Object
						ϒtitle         λ.Object
						ϒupload_date   λ.Object
						ϒuploader      λ.Object
						ϒuploader_id   λ.Object
						ϒurl           = λargs[1]
						ϒvideo_id      λ.Object
						ϒwebpage       λ.Object
						τmp0           λ.Object
						τmp1           λ.Object
					)
					ϒstart_time = λ.Cal(ϒint_or_none, λ.GetItem(λ.Calm(λ.Cal(ϒcompat_parse_qs, λ.GetAttr(λ.Cal(ϒcompat_urllib_parse_urlparse, ϒurl), "query", nil)), "get", λ.StrLiteral("fromTime"), λ.NewList(λ.None)), λ.IntLiteral(0)))
					ϒvideo_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒwebpage = λ.Calm(ϒself, "_download_webpage", λ.Mod(λ.StrLiteral("http://ok.ru/video/%s"), ϒvideo_id), ϒvideo_id)
					ϒerror = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("[^>]+class=\"vp_video_stub_txt\"[^>]*>([^<]+)<"),
						ϒwebpage,
						λ.StrLiteral("error"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					if λ.IsTrue(ϒerror) {
						panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(ϒerror), λ.KWArgs{
							{Name: "expected", Value: λ.True},
						})))
					}
					ϒplayer = λ.Calm(ϒself, "_parse_json", λ.Cal(ϒunescapeHTML, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.Mod(λ.StrLiteral("data-options=(?P<quote>[\"\\'])(?P<player>{.+?%s.+?})(?P=quote)"), ϒvideo_id),
						ϒwebpage,
						λ.StrLiteral("player"),
					), λ.KWArgs{
						{Name: "group", Value: λ.StrLiteral("player")},
					})), ϒvideo_id)
					ϒflashvars = λ.GetItem(ϒplayer, λ.StrLiteral("flashvars"))
					ϒmetadata = λ.Calm(ϒflashvars, "get", λ.StrLiteral("metadata"))
					if λ.IsTrue(ϒmetadata) {
						ϒmetadata = λ.Calm(ϒself, "_parse_json", ϒmetadata, ϒvideo_id)
					} else {
						ϒdata = λ.DictLiteral(map[λ.Object]λ.Object{})
						ϒst_location = λ.Calm(ϒflashvars, "get", λ.StrLiteral("location"))
						if λ.IsTrue(ϒst_location) {
							λ.SetItem(ϒdata, λ.StrLiteral("st.location"), ϒst_location)
						}
						ϒmetadata = λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
							λ.Cal(ϒcompat_urllib_parse_unquote, λ.GetItem(ϒflashvars, λ.StrLiteral("metadataUrl"))),
							ϒvideo_id,
							λ.StrLiteral("Downloading metadata JSON"),
						), λ.KWArgs{
							{Name: "data", Value: λ.Cal(ϒurlencode_postdata, ϒdata)},
						})
					}
					ϒmovie = λ.GetItem(ϒmetadata, λ.StrLiteral("movie"))
					ϒprovider = λ.Calm(ϒmetadata, "get", λ.StrLiteral("provider"))
					ϒtitle = func() λ.Object {
						if λ.IsTrue(λ.Eq(ϒprovider, λ.StrLiteral("UPLOADED_ODKL"))) {
							return λ.GetItem(ϒmovie, λ.StrLiteral("title"))
						} else {
							return λ.Calm(ϒmovie, "get", λ.StrLiteral("title"))
						}
					}()
					ϒthumbnail = λ.Calm(ϒmovie, "get", λ.StrLiteral("poster"))
					ϒduration = λ.Cal(ϒint_or_none, λ.Calm(ϒmovie, "get", λ.StrLiteral("duration")))
					ϒauthor = λ.Calm(ϒmetadata, "get", λ.StrLiteral("author"), λ.DictLiteral(map[λ.Object]λ.Object{}))
					ϒuploader_id = λ.Calm(ϒauthor, "get", λ.StrLiteral("id"))
					ϒuploader = λ.Calm(ϒauthor, "get", λ.StrLiteral("name"))
					ϒupload_date = λ.Cal(ϒunified_strdate, λ.Call(λ.GetAttr(ϒself, "_html_search_meta", nil), λ.NewArgs(
						λ.StrLiteral("ya:ovs:upload_date"),
						ϒwebpage,
						λ.StrLiteral("upload date"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					}))
					ϒage_limit = λ.None
					ϒadult = λ.Call(λ.GetAttr(ϒself, "_html_search_meta", nil), λ.NewArgs(
						λ.StrLiteral("ya:ovs:adult"),
						ϒwebpage,
						λ.StrLiteral("age limit"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					if λ.IsTrue(ϒadult) {
						ϒage_limit = func() λ.Object {
							if λ.IsTrue(λ.Eq(ϒadult, λ.StrLiteral("true"))) {
								return λ.IntLiteral(18)
							} else {
								return λ.IntLiteral(0)
							}
						}()
					}
					ϒlike_count = λ.Cal(ϒint_or_none, λ.Calm(ϒmetadata, "get", λ.StrLiteral("likeCount")))
					ϒinfo = λ.DictLiteral(map[string]λ.Object{
						"id":          ϒvideo_id,
						"title":       ϒtitle,
						"thumbnail":   ϒthumbnail,
						"duration":    ϒduration,
						"upload_date": ϒupload_date,
						"uploader":    ϒuploader,
						"uploader_id": ϒuploader_id,
						"like_count":  ϒlike_count,
						"age_limit":   ϒage_limit,
						"start_time":  ϒstart_time,
					})
					if λ.IsTrue(λ.Eq(ϒprovider, λ.StrLiteral("USER_YOUTUBE"))) {
						λ.Calm(ϒinfo, "update", λ.DictLiteral(map[string]λ.Object{
							"_type": λ.StrLiteral("url_transparent"),
							"url":   λ.GetItem(ϒmovie, λ.StrLiteral("contentId")),
						}))
						return ϒinfo
					}
					if !λ.IsTrue(ϒtitle) {
						panic(λ.Raise(λ.Cal(λ.AssertionErrorType)))
					}
					if λ.IsTrue(λ.Eq(ϒprovider, λ.StrLiteral("LIVE_TV_APP"))) {
						λ.SetItem(ϒinfo, λ.StrLiteral("title"), λ.Calm(ϒself, "_live_title", ϒtitle))
					}
					ϒquality = λ.Cal(ϒqualities, λ.NewTuple(
						λ.StrLiteral("4"),
						λ.StrLiteral("0"),
						λ.StrLiteral("1"),
						λ.StrLiteral("2"),
						λ.StrLiteral("3"),
						λ.StrLiteral("5"),
					))
					ϒformats = λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
						nil,
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							return λ.NewGenerator(func(λgy λ.Yielder) λ.Object {
								var (
									ϒf   λ.Object
									τmp0 λ.Object
									τmp1 λ.Object
								)
								τmp0 = λ.Cal(λ.BuiltinIter, λ.GetItem(ϒmetadata, λ.StrLiteral("videos")))
								for {
									if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
										break
									}
									ϒf = τmp1
									λgy.Yield(λ.DictLiteral(map[string]λ.Object{
										"url":       λ.GetItem(ϒf, λ.StrLiteral("url")),
										"ext":       λ.StrLiteral("mp4"),
										"format_id": λ.GetItem(ϒf, λ.StrLiteral("name")),
									}))
								}
								return λ.None
							})
						})))
					ϒm3u8_url = λ.Calm(ϒmetadata, "get", λ.StrLiteral("hlsManifestUrl"))
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
					ϒdash_manifest = λ.Calm(ϒmetadata, "get", λ.StrLiteral("metadataEmbedded"))
					if λ.IsTrue(ϒdash_manifest) {
						λ.Calm(ϒformats, "extend", λ.Calm(ϒself, "_parse_mpd_formats", λ.Cal(ϒcompat_etree_fromstring, ϒdash_manifest), λ.StrLiteral("mpd")))
					}
					τmp0 = λ.Cal(λ.BuiltinIter, ϒformats)
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒfmt = τmp1
						ϒfmt_type = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
							λ.StrLiteral("\\btype[/=](\\d)"),
							λ.GetItem(ϒfmt, λ.StrLiteral("url")),
							λ.StrLiteral("format type"),
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
						})
						if λ.IsTrue(ϒfmt_type) {
							λ.SetItem(ϒfmt, λ.StrLiteral("quality"), λ.Cal(ϒquality, ϒfmt_type))
						}
					}
					ϒm3u8_url = λ.Calm(ϒmetadata, "get", λ.StrLiteral("hlsMasterPlaylistUrl"))
					if λ.IsTrue(ϒm3u8_url) {
						λ.Calm(ϒformats, "extend", λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
							ϒm3u8_url,
							ϒvideo_id,
							λ.StrLiteral("mp4"),
						), λ.KWArgs{
							{Name: "entry_protocol", Value: λ.StrLiteral("m3u8")},
							{Name: "m3u8_id", Value: λ.StrLiteral("hls")},
							{Name: "fatal", Value: λ.False},
						}))
					}
					ϒrtmp_url = λ.Calm(ϒmetadata, "get", λ.StrLiteral("rtmpUrl"))
					if λ.IsTrue(ϒrtmp_url) {
						λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
							"url":       ϒrtmp_url,
							"format_id": λ.StrLiteral("rtmp"),
							"ext":       λ.StrLiteral("flv"),
						}))
					}
					if !λ.IsTrue(ϒformats) {
						ϒpayment_info = λ.Calm(ϒmetadata, "get", λ.StrLiteral("paymentInfo"))
						if λ.IsTrue(ϒpayment_info) {
							panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.StrLiteral("This video is paid, subscribe to download it")), λ.KWArgs{
								{Name: "expected", Value: λ.True},
							})))
						}
					}
					λ.Calm(ϒself, "_sort_formats", ϒformats)
					λ.SetItem(ϒinfo, λ.StrLiteral("formats"), ϒformats)
					return ϒinfo
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL":    OdnoklassnikiIE__VALID_URL,
				"_extract_url":  OdnoklassnikiIE__extract_url,
				"_real_extract": OdnoklassnikiIE__real_extract,
			})
		}())
	})
}
