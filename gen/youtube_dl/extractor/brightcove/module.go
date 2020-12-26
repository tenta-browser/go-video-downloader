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
 * brightcove/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/brightcove.py
 */

package brightcove

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωadobepass "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/adobepass"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	AdobePassIE                   λ.Object
	BrightcoveLegacyIE            λ.Object
	BrightcoveNewIE               λ.Object
	ExtractorError                λ.Object
	InfoExtractor                 λ.Object
	ϒclean_html                   λ.Object
	ϒcompat_HTTPError             λ.Object
	ϒcompat_etree_fromstring      λ.Object
	ϒcompat_parse_qs              λ.Object
	ϒcompat_urllib_parse_urlparse λ.Object
	ϒcompat_xml_parse_error       λ.Object
	ϒextract_attributes           λ.Object
	ϒfind_xpath_attr              λ.Object
	ϒfix_xml_ampersands           λ.Object
	ϒfloat_or_none                λ.Object
	ϒint_or_none                  λ.Object
	ϒjs_to_json                   λ.Object
	ϒmimetype2ext                 λ.Object
	ϒparse_iso8601                λ.Object
	ϒsmuggle_url                  λ.Object
	ϒstr_or_none                  λ.Object
	ϒtry_get                      λ.Object
	ϒunescapeHTML                 λ.Object
	ϒunsmuggle_url                λ.Object
	ϒupdate_url_query             λ.Object
	ϒurl_or_none                  λ.Object
)

func init() {
	λ.InitModule(func() {
		AdobePassIE = Ωadobepass.AdobePassIE
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_etree_fromstring = Ωcompat.ϒcompat_etree_fromstring
		ϒcompat_HTTPError = Ωcompat.ϒcompat_HTTPError
		ϒcompat_parse_qs = Ωcompat.ϒcompat_parse_qs
		ϒcompat_urllib_parse_urlparse = Ωcompat.ϒcompat_urllib_parse_urlparse
		ϒcompat_xml_parse_error = Ωcompat.ϒcompat_xml_parse_error
		ϒclean_html = Ωutils.ϒclean_html
		ϒextract_attributes = Ωutils.ϒextract_attributes
		ExtractorError = Ωutils.ExtractorError
		ϒfind_xpath_attr = Ωutils.ϒfind_xpath_attr
		ϒfix_xml_ampersands = Ωutils.ϒfix_xml_ampersands
		ϒfloat_or_none = Ωutils.ϒfloat_or_none
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒjs_to_json = Ωutils.ϒjs_to_json
		ϒmimetype2ext = Ωutils.ϒmimetype2ext
		ϒparse_iso8601 = Ωutils.ϒparse_iso8601
		ϒsmuggle_url = Ωutils.ϒsmuggle_url
		ϒstr_or_none = Ωutils.ϒstr_or_none
		ϒtry_get = Ωutils.ϒtry_get
		ϒunescapeHTML = Ωutils.ϒunescapeHTML
		ϒunsmuggle_url = Ωutils.ϒunsmuggle_url
		ϒupdate_url_query = Ωutils.ϒupdate_url_query
		ϒurl_or_none = Ωutils.ϒurl_or_none
		BrightcoveLegacyIE = λ.Cal(λ.TypeType, λ.StrLiteral("BrightcoveLegacyIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				BrightcoveLegacyIE__VALID_URL               λ.Object
				BrightcoveLegacyIE__extract_brightcove_url  λ.Object
				BrightcoveLegacyIE__extract_brightcove_urls λ.Object
			)
			BrightcoveLegacyIE__VALID_URL = λ.StrLiteral("(?:https?://.*brightcove\\.com/(services|viewer).*?\\?|brightcove:)(?P<query>.*)")
			BrightcoveLegacyIE__extract_brightcove_url = λ.NewFunction("_extract_brightcove_url",
				[]λ.Param{
					{Name: "cls"},
					{Name: "webpage"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒcls     = λargs[0]
						ϒurls    λ.Object
						ϒwebpage = λargs[1]
					)
					ϒurls = λ.Calm(ϒcls, "_extract_brightcove_urls", ϒwebpage)
					return func() λ.Object {
						if λ.IsTrue(ϒurls) {
							return λ.GetItem(ϒurls, λ.IntLiteral(0))
						} else {
							return λ.None
						}
					}()
				})
			BrightcoveLegacyIE__extract_brightcove_url = λ.Cal(λ.ClassMethodType, BrightcoveLegacyIE__extract_brightcove_url)
			BrightcoveLegacyIE__extract_brightcove_urls = λ.NewFunction("_extract_brightcove_urls",
				[]λ.Param{
					{Name: "cls"},
					{Name: "webpage"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒcls     = λargs[0]
						ϒmatches λ.Object
						ϒurl     λ.Object
						ϒurl_m   λ.Object
						ϒwebpage = λargs[1]
					)
					ϒurl_m = λ.Cal(Ωre.ϒsearch, λ.StrLiteral("(?x)\n                <meta\\s+\n                    (?:property|itemprop)=([\\'\"])(?:og:video|embedURL)\\1[^>]+\n                    content=([\\'\"])(?P<url>https?://(?:secure|c)\\.brightcove.com/(?:(?!\\2).)+)\\2\n            "), ϒwebpage)
					if λ.IsTrue(ϒurl_m) {
						ϒurl = λ.Cal(ϒunescapeHTML, λ.Calm(ϒurl_m, "group", λ.StrLiteral("url")))
						if λ.IsTrue(func() λ.Object {
							if λv := λ.NewBool(λ.Contains(ϒurl, λ.StrLiteral("playerKey"))); λ.IsTrue(λv) {
								return λv
							} else if λv := λ.NewBool(λ.Contains(ϒurl, λ.StrLiteral("videoId"))); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.NewBool(λ.Contains(ϒurl, λ.StrLiteral("idVideo")))
							}
						}()) {
							return λ.NewList(ϒurl)
						}
					}
					ϒmatches = λ.Cal(Ωre.ϒfindall, λ.StrLiteral("(?sx)<object\n            (?:\n                [^>]+?class=[\\'\"][^>]*?BrightcoveExperience.*?[\\'\"] |\n                [^>]*?>\\s*<param\\s+name=\"movie\"\\s+value=\"https?://[^/]*brightcove\\.com/\n            ).+?>\\s*</object>"), ϒwebpage)
					if λ.IsTrue(ϒmatches) {
						return λ.Cal(λ.ListType, λ.Cal(λ.FilterIteratorType, λ.None, λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
							nil,
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								return λ.NewGenerator(func(λgy λ.Yielder) λ.Object {
									var (
										ϒm   λ.Object
										τmp0 λ.Object
										τmp1 λ.Object
									)
									τmp0 = λ.Cal(λ.BuiltinIter, ϒmatches)
									for {
										if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
											break
										}
										ϒm = τmp1
										λgy.Yield(λ.Calm(ϒcls, "_build_brightcove_url", ϒm))
									}
									return λ.None
								})
							})))))
					}
					ϒmatches = λ.Cal(Ωre.ϒfindall, λ.StrLiteral("(customBC\\.createVideo\\(.+?\\);)"), ϒwebpage)
					if λ.IsTrue(ϒmatches) {
						return λ.Cal(λ.ListType, λ.Cal(λ.FilterIteratorType, λ.None, λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
							nil,
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								return λ.NewGenerator(func(λgy λ.Yielder) λ.Object {
									var (
										ϒcustom_bc λ.Object
										τmp0       λ.Object
										τmp1       λ.Object
									)
									τmp0 = λ.Cal(λ.BuiltinIter, ϒmatches)
									for {
										if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
											break
										}
										ϒcustom_bc = τmp1
										λgy.Yield(λ.Calm(ϒcls, "_build_brightcove_url_from_js", ϒcustom_bc))
									}
									return λ.None
								})
							})))))
					}
					return λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
						nil,
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							return λ.NewGenerator(func(λgy λ.Yielder) λ.Object {
								var (
									ϒsrc λ.Object
									τmp0 λ.Object
									τmp1 λ.Object
									τmp2 λ.Object
								)
								τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(Ωre.ϒfindall, λ.StrLiteral("<iframe[^>]+src=([\\'\"])((?:https?:)?//link\\.brightcove\\.com/services/player/(?!\\1).+)\\1"), ϒwebpage))
								for {
									if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
										break
									}
									τmp2 = τmp1
									_ = λ.GetItem(τmp2, λ.IntLiteral(0))
									ϒsrc = λ.GetItem(τmp2, λ.IntLiteral(1))
									λgy.Yield(ϒsrc)
								}
								return λ.None
							})
						})))
				})
			BrightcoveLegacyIE__extract_brightcove_urls = λ.Cal(λ.ClassMethodType, BrightcoveLegacyIE__extract_brightcove_urls)
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL":               BrightcoveLegacyIE__VALID_URL,
				"_extract_brightcove_url":  BrightcoveLegacyIE__extract_brightcove_url,
				"_extract_brightcove_urls": BrightcoveLegacyIE__extract_brightcove_urls,
			})
		}())
		BrightcoveNewIE = λ.Cal(λ.TypeType, λ.StrLiteral("BrightcoveNewIE"), λ.NewTuple(AdobePassIE), func() λ.Dict {
			var (
				BrightcoveNewIE_IE_NAME                    λ.Object
				BrightcoveNewIE__VALID_URL                 λ.Object
				BrightcoveNewIE__extract_url               λ.Object
				BrightcoveNewIE__extract_urls              λ.Object
				BrightcoveNewIE__parse_brightcove_metadata λ.Object
				BrightcoveNewIE__real_extract              λ.Object
			)
			BrightcoveNewIE_IE_NAME = λ.StrLiteral("brightcove:new")
			BrightcoveNewIE__VALID_URL = λ.StrLiteral("https?://players\\.brightcove\\.net/(?P<account_id>\\d+)/(?P<player_id>[^/]+)_(?P<embed>[^/]+)/index\\.html\\?.*(?P<content_type>video|playlist)Id=(?P<video_id>\\d+|ref:[^&]+)")
			BrightcoveNewIE__extract_url = λ.NewFunction("_extract_url",
				[]λ.Param{
					{Name: "ie"},
					{Name: "webpage"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒie      = λargs[0]
						ϒurls    λ.Object
						ϒwebpage = λargs[1]
					)
					ϒurls = λ.Calm(BrightcoveNewIE, "_extract_urls", ϒie, ϒwebpage)
					return func() λ.Object {
						if λ.IsTrue(ϒurls) {
							return λ.GetItem(ϒurls, λ.IntLiteral(0))
						} else {
							return λ.None
						}
					}()
				})
			BrightcoveNewIE__extract_url = λ.Cal(λ.StaticMethodType, BrightcoveNewIE__extract_url)
			BrightcoveNewIE__extract_urls = λ.NewFunction("_extract_urls",
				[]λ.Param{
					{Name: "ie"},
					{Name: "webpage"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒaccount_id λ.Object
						ϒattrs      λ.Object
						ϒbc_url     λ.Object
						ϒembed      λ.Object
						ϒentries    λ.Object
						ϒie         = λargs[0]
						ϒplayer_id  λ.Object
						ϒscript_tag λ.Object
						ϒurl        λ.Object
						ϒvideo      λ.Object
						ϒvideo_id   λ.Object
						ϒwebpage    = λargs[1]
						τmp0        λ.Object
						τmp1        λ.Object
						τmp2        λ.Object
					)
					ϒentries = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(Ωre.ϒfindall, λ.StrLiteral("<iframe[^>]+src=([\"\\'])((?:https?:)?//players\\.brightcove\\.net/\\d+/[^/]+/index\\.html.+?)\\1"), ϒwebpage))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						τmp2 = τmp1
						_ = λ.GetItem(τmp2, λ.IntLiteral(0))
						ϒurl = λ.GetItem(τmp2, λ.IntLiteral(1))
						λ.Calm(ϒentries, "append", func() λ.Object {
							if λ.IsTrue(λ.Calm(ϒurl, "startswith", λ.StrLiteral("http"))) {
								return ϒurl
							} else {
								return λ.Add(λ.StrLiteral("http:"), ϒurl)
							}
						}())
					}
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(Ωre.ϒfindall, λ.StrLiteral("(?isx)\n                    (<video(?:-js)?\\s+[^>]*\\bdata-video-id\\s*=\\s*['\"]?[^>]+>)\n                    (?:.*?\n                        (<script[^>]+\n                            src=[\"\\'](?:https?:)?//players\\.brightcove\\.net/\n                            (\\d+)/([^/]+)_([^/]+)/index(?:\\.min)?\\.js\n                        )\n                    )?\n                "), ϒwebpage))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						τmp2 = τmp1
						ϒvideo = λ.GetItem(τmp2, λ.IntLiteral(0))
						ϒscript_tag = λ.GetItem(τmp2, λ.IntLiteral(1))
						ϒaccount_id = λ.GetItem(τmp2, λ.IntLiteral(2))
						ϒplayer_id = λ.GetItem(τmp2, λ.IntLiteral(3))
						ϒembed = λ.GetItem(τmp2, λ.IntLiteral(4))
						ϒattrs = λ.Cal(ϒextract_attributes, ϒvideo)
						ϒvideo_id = λ.Calm(ϒattrs, "get", λ.StrLiteral("data-video-id"))
						if !λ.IsTrue(ϒvideo_id) {
							continue
						}
						ϒaccount_id = func() λ.Object {
							if λv := ϒaccount_id; λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Calm(ϒattrs, "get", λ.StrLiteral("data-account"))
							}
						}()
						if !λ.IsTrue(ϒaccount_id) {
							continue
						}
						ϒplayer_id = func() λ.Object {
							if λv := ϒplayer_id; λ.IsTrue(λv) {
								return λv
							} else if λv := λ.Calm(ϒattrs, "get", λ.StrLiteral("data-player")); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.StrLiteral("default")
							}
						}()
						ϒembed = func() λ.Object {
							if λv := ϒembed; λ.IsTrue(λv) {
								return λv
							} else if λv := λ.Calm(ϒattrs, "get", λ.StrLiteral("data-embed")); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.StrLiteral("default")
							}
						}()
						ϒbc_url = λ.Mod(λ.StrLiteral("http://players.brightcove.net/%s/%s_%s/index.html?videoId=%s"), λ.NewTuple(
							ϒaccount_id,
							ϒplayer_id,
							ϒembed,
							ϒvideo_id,
						))
						if λ.IsTrue(func() λ.Object {
							if λv := λ.NewBool(!λ.IsTrue(ϒscript_tag)); !λ.IsTrue(λv) {
								return λv
							} else {
								return λ.NewBool(!λ.IsTrue(λ.Calm(ϒie, "_is_valid_url", ϒbc_url, ϒvideo_id, λ.StrLiteral("possible brightcove video"))))
							}
						}()) {
							continue
						}
						λ.Calm(ϒentries, "append", ϒbc_url)
					}
					return ϒentries
				})
			BrightcoveNewIE__extract_urls = λ.Cal(λ.StaticMethodType, BrightcoveNewIE__extract_urls)
			BrightcoveNewIE__parse_brightcove_metadata = λ.NewFunction("_parse_brightcove_metadata",
				[]λ.Param{
					{Name: "self"},
					{Name: "json_data"},
					{Name: "video_id"},
					{Name: "headers", Def: λ.DictLiteral(map[λ.Object]λ.Object{})},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒapp_name        λ.Object
						ϒbuild_format_id λ.Object
						ϒcontainer       λ.Object
						ϒduration        λ.Object
						ϒerror           λ.Object
						ϒerrors          λ.Object
						ϒext             λ.Object
						ϒf               λ.Object
						ϒformats         λ.Object
						ϒheaders         = λargs[3]
						ϒheight          λ.Object
						ϒis_live         λ.Object
						ϒjson_data       = λargs[1]
						ϒlang            λ.Object
						ϒs3_source_url   λ.Object
						ϒself            = λargs[0]
						ϒsource          λ.Object
						ϒsrc             λ.Object
						ϒstream_name     λ.Object
						ϒstreaming_src   λ.Object
						ϒsubtitles       λ.Object
						ϒtbr             λ.Object
						ϒtext_track      λ.Object
						ϒtext_track_url  λ.Object
						ϒtitle           λ.Object
						ϒvideo_id        = λargs[2]
						ϒwidth           λ.Object
						τmp0             λ.Object
						τmp1             λ.Object
						τmp2             λ.Object
					)
					ϒtitle = λ.Calm(λ.GetItem(ϒjson_data, λ.StrLiteral("name")), "strip")
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(ϒjson_data, "get", λ.StrLiteral("sources"), λ.NewList()))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒsource = τmp1
						ϒcontainer = λ.Calm(ϒsource, "get", λ.StrLiteral("container"))
						ϒext = λ.Cal(ϒmimetype2ext, λ.Calm(ϒsource, "get", λ.StrLiteral("type")))
						ϒsrc = λ.Calm(ϒsource, "get", λ.StrLiteral("src"))
						if λ.IsTrue(func() λ.Object {
							if λv := λ.Eq(ϒext, λ.StrLiteral("ism")); λ.IsTrue(λv) {
								return λv
							} else if λv := λ.Eq(ϒcontainer, λ.StrLiteral("WVM")); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Calm(ϒsource, "get", λ.StrLiteral("key_systems"))
							}
						}()) {
							continue
						} else {
							if λ.IsTrue(func() λ.Object {
								if λv := λ.Eq(ϒext, λ.StrLiteral("m3u8")); λ.IsTrue(λv) {
									return λv
								} else {
									return λ.Eq(ϒcontainer, λ.StrLiteral("M2TS"))
								}
							}()) {
								if !λ.IsTrue(ϒsrc) {
									continue
								}
								λ.Calm(ϒformats, "extend", λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
									ϒsrc,
									ϒvideo_id,
									λ.StrLiteral("mp4"),
									λ.StrLiteral("m3u8_native"),
								), λ.KWArgs{
									{Name: "m3u8_id", Value: λ.StrLiteral("hls")},
									{Name: "fatal", Value: λ.False},
								}))
							} else {
								if λ.IsTrue(λ.Eq(ϒext, λ.StrLiteral("mpd"))) {
									if !λ.IsTrue(ϒsrc) {
										continue
									}
									λ.Calm(ϒformats, "extend", λ.Call(λ.GetAttr(ϒself, "_extract_mpd_formats", nil), λ.NewArgs(
										ϒsrc,
										ϒvideo_id,
										λ.StrLiteral("dash"),
									), λ.KWArgs{
										{Name: "fatal", Value: λ.False},
									}))
								} else {
									ϒstreaming_src = λ.Calm(ϒsource, "get", λ.StrLiteral("streaming_src"))
									τmp2 = λ.NewTuple(
										λ.Calm(ϒsource, "get", λ.StrLiteral("stream_name")),
										λ.Calm(ϒsource, "get", λ.StrLiteral("app_name")),
									)
									ϒstream_name = λ.GetItem(τmp2, λ.IntLiteral(0))
									ϒapp_name = λ.GetItem(τmp2, λ.IntLiteral(1))
									if λ.IsTrue(func() λ.Object {
										if λv := λ.NewBool(!λ.IsTrue(ϒsrc)); !λ.IsTrue(λv) {
											return λv
										} else if λv := λ.NewBool(!λ.IsTrue(ϒstreaming_src)); !λ.IsTrue(λv) {
											return λv
										} else {
											return func() λ.Object {
												if λv := λ.NewBool(!λ.IsTrue(ϒstream_name)); λ.IsTrue(λv) {
													return λv
												} else {
													return λ.NewBool(!λ.IsTrue(ϒapp_name))
												}
											}()
										}
									}()) {
										continue
									}
									ϒtbr = λ.Cal(ϒfloat_or_none, λ.Calm(ϒsource, "get", λ.StrLiteral("avg_bitrate")), λ.IntLiteral(1000))
									ϒheight = λ.Cal(ϒint_or_none, λ.Calm(ϒsource, "get", λ.StrLiteral("height")))
									ϒwidth = λ.Cal(ϒint_or_none, λ.Calm(ϒsource, "get", λ.StrLiteral("width")))
									ϒf = λ.DictLiteral(map[string]λ.Object{
										"tbr":       ϒtbr,
										"filesize":  λ.Cal(ϒint_or_none, λ.Calm(ϒsource, "get", λ.StrLiteral("size"))),
										"container": ϒcontainer,
										"ext": func() λ.Object {
											if λv := ϒext; λ.IsTrue(λv) {
												return λv
											} else {
												return λ.Calm(ϒcontainer, "lower")
											}
										}(),
									})
									if λ.IsTrue(func() λ.Object {
										if λv := λ.Eq(ϒwidth, λ.IntLiteral(0)); !λ.IsTrue(λv) {
											return λv
										} else {
											return λ.Eq(ϒheight, λ.IntLiteral(0))
										}
									}()) {
										λ.Calm(ϒf, "update", λ.DictLiteral(map[string]string{
											"vcodec": "none",
										}))
									} else {
										λ.Calm(ϒf, "update", λ.DictLiteral(map[string]λ.Object{
											"width":  ϒwidth,
											"height": ϒheight,
											"vcodec": λ.Calm(ϒsource, "get", λ.StrLiteral("codec")),
										}))
									}
									ϒbuild_format_id = λ.NewFunction("build_format_id",
										[]λ.Param{
											{Name: "kind"},
										},
										0, false, false,
										func(λargs []λ.Object) λ.Object {
											var (
												ϒformat_id λ.Object
												ϒkind      = λargs[0]
												τmp0       λ.Object
											)
											ϒformat_id = ϒkind
											if λ.IsTrue(ϒtbr) {
												τmp0 = λ.IAdd(ϒformat_id, λ.Mod(λ.StrLiteral("-%dk"), λ.Cal(λ.IntType, ϒtbr)))
												ϒformat_id = τmp0
											}
											if λ.IsTrue(ϒheight) {
												τmp0 = λ.IAdd(ϒformat_id, λ.Mod(λ.StrLiteral("-%dp"), ϒheight))
												ϒformat_id = τmp0
											}
											return ϒformat_id
										})
									if λ.IsTrue(func() λ.Object {
										if λv := ϒsrc; λ.IsTrue(λv) {
											return λv
										} else {
											return ϒstreaming_src
										}
									}()) {
										λ.Calm(ϒf, "update", λ.DictLiteral(map[string]λ.Object{
											"url": func() λ.Object {
												if λv := ϒsrc; λ.IsTrue(λv) {
													return λv
												} else {
													return ϒstreaming_src
												}
											}(),
											"format_id": λ.Cal(ϒbuild_format_id, func() λ.Object {
												if λ.IsTrue(ϒsrc) {
													return λ.StrLiteral("http")
												} else {
													return λ.StrLiteral("http-streaming")
												}
											}()),
											"source_preference": func() λ.Object {
												if λ.IsTrue(ϒsrc) {
													return λ.IntLiteral(0)
												} else {
													return λ.Neg(λ.IntLiteral(1))
												}
											}(),
										}))
									} else {
										λ.Calm(ϒf, "update", λ.DictLiteral(map[string]λ.Object{
											"url":       ϒapp_name,
											"play_path": ϒstream_name,
											"format_id": λ.Cal(ϒbuild_format_id, λ.StrLiteral("rtmp")),
										}))
									}
									λ.Calm(ϒformats, "append", ϒf)
								}
							}
						}
					}
					if !λ.IsTrue(ϒformats) {
						ϒs3_source_url = λ.Calm(λ.Calm(ϒjson_data, "get", λ.StrLiteral("custom_fields"), λ.DictLiteral(map[λ.Object]λ.Object{})), "get", λ.StrLiteral("s3sourceurl"))
						if λ.IsTrue(ϒs3_source_url) {
							λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
								"url":       ϒs3_source_url,
								"format_id": λ.StrLiteral("source"),
							}))
						}
					}
					ϒerrors = λ.Calm(ϒjson_data, "get", λ.StrLiteral("errors"))
					if λ.IsTrue(func() λ.Object {
						if λv := λ.NewBool(!λ.IsTrue(ϒformats)); !λ.IsTrue(λv) {
							return λv
						} else {
							return ϒerrors
						}
					}()) {
						ϒerror = λ.GetItem(ϒerrors, λ.IntLiteral(0))
						panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(func() λ.Object {
							if λv := λ.Calm(ϒerror, "get", λ.StrLiteral("message")); λ.IsTrue(λv) {
								return λv
							} else if λv := λ.Calm(ϒerror, "get", λ.StrLiteral("error_subcode")); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.GetItem(ϒerror, λ.StrLiteral("error_code"))
							}
						}()), λ.KWArgs{
							{Name: "expected", Value: λ.True},
						})))
					}
					λ.Calm(ϒself, "_sort_formats", ϒformats)
					τmp0 = λ.Cal(λ.BuiltinIter, ϒformats)
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒf = τmp1
						λ.Calm(λ.Calm(ϒf, "setdefault", λ.StrLiteral("http_headers"), λ.DictLiteral(map[λ.Object]λ.Object{})), "update", ϒheaders)
					}
					ϒsubtitles = λ.DictLiteral(map[λ.Object]λ.Object{})
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(ϒjson_data, "get", λ.StrLiteral("text_tracks"), λ.NewList()))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒtext_track = τmp1
						if λ.IsTrue(λ.Ne(λ.Calm(ϒtext_track, "get", λ.StrLiteral("kind")), λ.StrLiteral("captions"))) {
							continue
						}
						ϒtext_track_url = λ.Cal(ϒurl_or_none, λ.Calm(ϒtext_track, "get", λ.StrLiteral("src")))
						if !λ.IsTrue(ϒtext_track_url) {
							continue
						}
						ϒlang = λ.Calm(func() λ.Object {
							if λv := λ.Cal(ϒstr_or_none, λ.Calm(ϒtext_track, "get", λ.StrLiteral("srclang"))); λ.IsTrue(λv) {
								return λv
							} else if λv := λ.Cal(ϒstr_or_none, λ.Calm(ϒtext_track, "get", λ.StrLiteral("label"))); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.StrLiteral("en")
							}
						}(), "lower")
						λ.Calm(λ.Calm(ϒsubtitles, "setdefault", ϒlang, λ.NewList()), "append", λ.DictLiteral(map[string]λ.Object{
							"url": ϒtext_track_url,
						}))
					}
					ϒis_live = λ.False
					ϒduration = λ.Cal(ϒfloat_or_none, λ.Calm(ϒjson_data, "get", λ.StrLiteral("duration")), λ.IntLiteral(1000))
					if λ.IsTrue(func() λ.Object {
						if λv := λ.NewBool(ϒduration != λ.None); !λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Le(ϒduration, λ.IntLiteral(0))
						}
					}()) {
						ϒis_live = λ.True
					}
					return λ.DictLiteral(map[string]λ.Object{
						"id": ϒvideo_id,
						"title": func() λ.Object {
							if λ.IsTrue(ϒis_live) {
								return λ.Calm(ϒself, "_live_title", ϒtitle)
							} else {
								return ϒtitle
							}
						}(),
						"description": λ.Cal(ϒclean_html, λ.Calm(ϒjson_data, "get", λ.StrLiteral("description"))),
						"thumbnail": func() λ.Object {
							if λv := λ.Calm(ϒjson_data, "get", λ.StrLiteral("thumbnail")); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Calm(ϒjson_data, "get", λ.StrLiteral("poster"))
							}
						}(),
						"duration":    ϒduration,
						"timestamp":   λ.Cal(ϒparse_iso8601, λ.Calm(ϒjson_data, "get", λ.StrLiteral("published_at"))),
						"uploader_id": λ.Calm(ϒjson_data, "get", λ.StrLiteral("account_id")),
						"formats":     ϒformats,
						"subtitles":   ϒsubtitles,
						"tags":        λ.Calm(ϒjson_data, "get", λ.StrLiteral("tags"), λ.NewList()),
						"is_live":     ϒis_live,
					})
				})
			BrightcoveNewIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒaccount_id           λ.Object
						ϒapi_url              λ.Object
						ϒcontent_type         λ.Object
						ϒcustom_fields        λ.Object
						ϒembed                λ.Object
						ϒerrors               λ.Object
						ϒextract_policy_key   λ.Object
						ϒheaders              λ.Object
						ϒjson_data            λ.Object
						ϒmessage              λ.Object
						ϒplayer_id            λ.Object
						ϒpolicy_key           λ.Object
						ϒpolicy_key_extracted λ.Object
						ϒpolicy_key_id        λ.Object
						ϒreferrer             λ.Object
						ϒself                 = λargs[0]
						ϒsmuggled_data        λ.Object
						ϒstore_pk             λ.Object
						ϒtve_token            λ.Object
						ϒurl                  = λargs[1]
						ϒvideo_id             λ.Object
						τmp0                  λ.Object
						τmp1                  λ.Object
						τmp2                  λ.Object
						τmp3                  λ.Object
					)
					_ = τmp3
					τmp0 = λ.Cal(ϒunsmuggle_url, ϒurl, λ.DictLiteral(map[λ.Object]λ.Object{}))
					ϒurl = λ.GetItem(τmp0, λ.IntLiteral(0))
					ϒsmuggled_data = λ.GetItem(τmp0, λ.IntLiteral(1))
					λ.Calm(ϒself, "_initialize_geo_bypass", λ.DictLiteral(map[string]λ.Object{
						"countries": λ.Calm(ϒsmuggled_data, "get", λ.StrLiteral("geo_countries")),
						"ip_blocks": λ.Calm(ϒsmuggled_data, "get", λ.StrLiteral("geo_ip_blocks")),
					}))
					τmp0 = λ.Calm(λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl), "groups")
					ϒaccount_id = λ.GetItem(τmp0, λ.IntLiteral(0))
					ϒplayer_id = λ.GetItem(τmp0, λ.IntLiteral(1))
					ϒembed = λ.GetItem(τmp0, λ.IntLiteral(2))
					ϒcontent_type = λ.GetItem(τmp0, λ.IntLiteral(3))
					ϒvideo_id = λ.GetItem(τmp0, λ.IntLiteral(4))
					ϒpolicy_key_id = λ.Mod(λ.StrLiteral("%s_%s"), λ.NewTuple(
						ϒaccount_id,
						ϒplayer_id,
					))
					ϒpolicy_key = λ.Calm(λ.GetAttr(λ.GetAttr(ϒself, "_downloader", nil), "cache", nil), "load", λ.StrLiteral("brightcove"), ϒpolicy_key_id)
					ϒpolicy_key_extracted = λ.False
					ϒstore_pk = λ.NewFunction("<lambda>",
						[]λ.Param{
							{Name: "x"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒx = λargs[0]
							)
							return λ.Calm(λ.GetAttr(λ.GetAttr(ϒself, "_downloader", nil), "cache", nil), "store", λ.StrLiteral("brightcove"), ϒpolicy_key_id, ϒx)
						})
					ϒextract_policy_key = λ.NewFunction("extract_policy_key",
						nil,
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒbase_url   λ.Object
								ϒcatalog    λ.Object
								ϒconfig     λ.Object
								ϒpolicy_key λ.Object
								ϒwebpage    λ.Object
							)
							ϒbase_url = λ.Mod(λ.StrLiteral("http://players.brightcove.net/%s/%s_%s/"), λ.NewTuple(
								ϒaccount_id,
								ϒplayer_id,
								ϒembed,
							))
							ϒconfig = func() λ.Object {
								if λv := λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
									λ.Add(ϒbase_url, λ.StrLiteral("config.json")),
									ϒvideo_id,
								), λ.KWArgs{
									{Name: "fatal", Value: λ.False},
								}); λ.IsTrue(λv) {
									return λv
								} else {
									return λ.DictLiteral(map[λ.Object]λ.Object{})
								}
							}()
							ϒpolicy_key = λ.Cal(ϒtry_get, ϒconfig, λ.NewFunction("<lambda>",
								[]λ.Param{
									{Name: "x"},
								},
								0, false, false,
								func(λargs []λ.Object) λ.Object {
									var (
										ϒx = λargs[0]
									)
									return λ.GetItem(λ.GetItem(ϒx, λ.StrLiteral("video_cloud")), λ.StrLiteral("policy_key"))
								}))
							if !λ.IsTrue(ϒpolicy_key) {
								ϒwebpage = λ.Calm(ϒself, "_download_webpage", λ.Add(ϒbase_url, λ.StrLiteral("index.min.js")), ϒvideo_id)
								ϒcatalog = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
									λ.StrLiteral("catalog\\(({.+?})\\);"),
									ϒwebpage,
									λ.StrLiteral("catalog"),
								), λ.KWArgs{
									{Name: "default", Value: λ.None},
								})
								if λ.IsTrue(ϒcatalog) {
									ϒcatalog = λ.Call(λ.GetAttr(ϒself, "_parse_json", nil), λ.NewArgs(
										λ.Cal(ϒjs_to_json, ϒcatalog),
										ϒvideo_id,
									), λ.KWArgs{
										{Name: "fatal", Value: λ.False},
									})
									if λ.IsTrue(ϒcatalog) {
										ϒpolicy_key = λ.Calm(ϒcatalog, "get", λ.StrLiteral("policyKey"))
									}
								}
								if !λ.IsTrue(ϒpolicy_key) {
									ϒpolicy_key = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
										λ.StrLiteral("policyKey\\s*:\\s*([\"\\'])(?P<pk>.+?)\\1"),
										ϒwebpage,
										λ.StrLiteral("policy key"),
									), λ.KWArgs{
										{Name: "group", Value: λ.StrLiteral("pk")},
									})
								}
							}
							λ.Cal(ϒstore_pk, ϒpolicy_key)
							return ϒpolicy_key
						})
					ϒapi_url = λ.Mod(λ.StrLiteral("https://edge.api.brightcove.com/playback/v1/accounts/%s/%ss/%s"), λ.NewTuple(
						ϒaccount_id,
						ϒcontent_type,
						ϒvideo_id,
					))
					ϒheaders = λ.DictLiteral(map[λ.Object]λ.Object{})
					ϒreferrer = λ.Calm(ϒsmuggled_data, "get", λ.StrLiteral("referrer"))
					if λ.IsTrue(ϒreferrer) {
						λ.Calm(ϒheaders, "update", λ.DictLiteral(map[string]λ.Object{
							"Referer": ϒreferrer,
							"Origin":  λ.Calm(λ.Cal(Ωre.ϒsearch, λ.StrLiteral("https?://[^/]+"), ϒreferrer), "group", λ.IntLiteral(0)),
						}))
					}
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.RangeType, λ.IntLiteral(2)))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						_ = τmp1
						if !λ.IsTrue(ϒpolicy_key) {
							ϒpolicy_key = λ.Cal(ϒextract_policy_key)
							ϒpolicy_key_extracted = λ.True
						}
						λ.SetItem(ϒheaders, λ.StrLiteral("Accept"), λ.Mod(λ.StrLiteral("application/json;pk=%s"), ϒpolicy_key))
						τmp2, τmp3 = func() (λexit λ.Object, λret λ.Object) {
							defer λ.CatchMulti(
								nil,
								&λ.Catcher{ExtractorError, func(λex λ.BaseException) {
									var ϒe λ.Object = λex
									if λ.IsTrue(func() λ.Object {
										if λv := λ.Cal(λ.BuiltinIsInstance, λ.GetAttr(ϒe, "cause", nil), ϒcompat_HTTPError); !λ.IsTrue(λv) {
											return λv
										} else {
											return λ.NewBool(λ.Contains(λ.NewTuple(
												λ.IntLiteral(401),
												λ.IntLiteral(403),
											), λ.GetAttr(λ.GetAttr(ϒe, "cause", nil), "code", nil)))
										}
									}()) {
										ϒjson_data = λ.GetItem(λ.Calm(ϒself, "_parse_json", λ.Calm(λ.Calm(λ.GetAttr(ϒe, "cause", nil), "read"), "decode"), ϒvideo_id), λ.IntLiteral(0))
										ϒmessage = func() λ.Object {
											if λv := λ.Calm(ϒjson_data, "get", λ.StrLiteral("message")); λ.IsTrue(λv) {
												return λv
											} else {
												return λ.GetItem(ϒjson_data, λ.StrLiteral("error_code"))
											}
										}()
										if λ.IsTrue(λ.Eq(λ.Calm(ϒjson_data, "get", λ.StrLiteral("error_subcode")), λ.StrLiteral("CLIENT_GEO"))) {
											λ.Call(λ.GetAttr(ϒself, "raise_geo_restricted", nil), nil, λ.KWArgs{
												{Name: "msg", Value: ϒmessage},
											})
										} else {
											if λ.IsTrue(func() λ.Object {
												if λv := λ.Eq(λ.Calm(ϒjson_data, "get", λ.StrLiteral("error_code")), λ.StrLiteral("INVALID_POLICY_KEY")); !λ.IsTrue(λv) {
													return λv
												} else {
													return λ.NewBool(!λ.IsTrue(ϒpolicy_key_extracted))
												}
											}()) {
												ϒpolicy_key = λ.None
												λ.Cal(ϒstore_pk, λ.None)
												λexit = λ.BlockExitContinue
												return
											}
										}
										panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(ϒmessage), λ.KWArgs{
											{Name: "expected", Value: λ.True},
										})))
									}
									panic(λ.Raise(λex))
								}},
							)
							ϒjson_data = λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
								ϒapi_url,
								ϒvideo_id,
							), λ.KWArgs{
								{Name: "headers", Value: ϒheaders},
							})
							λexit = λ.BlockExitBreak
							return
							return λ.BlockExitNormally, nil
						}()
						if τmp2 == λ.BlockExitBreak {
							break
						}
						if τmp2 == λ.BlockExitContinue {
							continue
						}
					}
					ϒerrors = λ.Calm(ϒjson_data, "get", λ.StrLiteral("errors"))
					if λ.IsTrue(func() λ.Object {
						if λv := ϒerrors; !λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Eq(λ.Calm(λ.GetItem(ϒerrors, λ.IntLiteral(0)), "get", λ.StrLiteral("error_subcode")), λ.StrLiteral("TVE_AUTH"))
						}
					}()) {
						ϒcustom_fields = λ.GetItem(ϒjson_data, λ.StrLiteral("custom_fields"))
						ϒtve_token = λ.Calm(ϒself, "_extract_mvpd_auth", λ.GetItem(ϒsmuggled_data, λ.StrLiteral("source_url")), ϒvideo_id, λ.GetItem(ϒcustom_fields, λ.StrLiteral("bcadobepassrequestorid")), λ.GetItem(ϒcustom_fields, λ.StrLiteral("bcadobepassresourceid")))
						ϒjson_data = λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
							ϒapi_url,
							ϒvideo_id,
						), λ.KWArgs{
							{Name: "headers", Value: λ.DictLiteral(map[string]λ.Object{
								"Accept": λ.Mod(λ.StrLiteral("application/json;pk=%s"), ϒpolicy_key),
							})},
							{Name: "query", Value: λ.DictLiteral(map[string]λ.Object{
								"tveToken": ϒtve_token,
							})},
						})
					}
					if λ.IsTrue(λ.Eq(ϒcontent_type, λ.StrLiteral("playlist"))) {
						return λ.Calm(ϒself, "playlist_result", λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
							nil,
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								return λ.NewGenerator(func(λgy λ.Yielder) λ.Object {
									var (
										ϒvid λ.Object
										τmp0 λ.Object
										τmp1 λ.Object
									)
									τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(ϒjson_data, "get", λ.StrLiteral("videos"), λ.NewList()))
									for {
										if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
											break
										}
										ϒvid = τmp1
										if λ.IsTrue(λ.Calm(ϒvid, "get", λ.StrLiteral("id"))) {
											λgy.Yield(λ.Calm(ϒself, "_parse_brightcove_metadata", ϒvid, λ.Calm(ϒvid, "get", λ.StrLiteral("id")), ϒheaders))
										}
									}
									return λ.None
								})
							}))), λ.Calm(ϒjson_data, "get", λ.StrLiteral("id")), λ.Calm(ϒjson_data, "get", λ.StrLiteral("name")), λ.Calm(ϒjson_data, "get", λ.StrLiteral("description")))
					}
					return λ.Call(λ.GetAttr(ϒself, "_parse_brightcove_metadata", nil), λ.NewArgs(
						ϒjson_data,
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "headers", Value: ϒheaders},
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"IE_NAME":                    BrightcoveNewIE_IE_NAME,
				"_VALID_URL":                 BrightcoveNewIE__VALID_URL,
				"_extract_url":               BrightcoveNewIE__extract_url,
				"_extract_urls":              BrightcoveNewIE__extract_urls,
				"_parse_brightcove_metadata": BrightcoveNewIE__parse_brightcove_metadata,
				"_real_extract":              BrightcoveNewIE__real_extract,
			})
		}())
	})
}
