// Code generated by transpiler. DO NOT EDIT.

/**
 * Go Video Downloader
 *
 *    Copyright 2018 Tenta, LLC
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
 * tvp/module.go: transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/tvp.py
 */

package tvp

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError            λ.Object
	InfoExtractor             λ.Object
	TVPEmbedIE                λ.Object
	TVPIE                     λ.Object
	TVPWebsiteIE              λ.Object
	ϒclean_html               λ.Object
	ϒdetermine_ext            λ.Object
	ϒget_element_by_attribute λ.Object
	ϒorderedSet               λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒclean_html = Ωutils.ϒclean_html
		ϒdetermine_ext = Ωutils.ϒdetermine_ext
		ExtractorError = Ωutils.ExtractorError
		ϒget_element_by_attribute = Ωutils.ϒget_element_by_attribute
		ϒorderedSet = Ωutils.ϒorderedSet
		TVPIE = λ.Cal(λ.TypeType, λ.NewStr("TVPIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				TVPIE_IE_NAME       λ.Object
				TVPIE__TESTS        λ.Object
				TVPIE__VALID_URL    λ.Object
				TVPIE__real_extract λ.Object
			)
			TVPIE_IE_NAME = λ.NewStr("tvp")
			TVPIE__VALID_URL = λ.NewStr("https?://[^/]+\\.tvp\\.(?:pl|info)/(?:video/(?:[^,\\s]*,)*|(?:(?!\\d+/)[^/]+/)*)(?P<id>\\d+)")
			TVPIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("https://vod.tvp.pl/video/czas-honoru,i-seria-odc-13,194536"),
					λ.NewStr("md5"): λ.NewStr("a21eb0aa862f25414430f15fdfb9e76c"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("194536"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Czas honoru, odc. 13 – Władek"),
						λ.NewStr("description"): λ.NewStr("md5:437f48b93558370b031740546b696e24"),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.tvp.pl/there-can-be-anything-so-i-shortened-it/17916176"),
					λ.NewStr("md5"): λ.NewStr("b0005b542e5b4de643a9690326ab1257"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("17916176"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("TVP Gorzów pokaże filmy studentów z podroży dookoła świata"),
						λ.NewStr("description"): λ.NewStr("TVP Gorzów pokaże filmy studentów z podroży dookoła świata"),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("https://wiadomosci.tvp.pl/33908820/28092017-1930"),
					λ.NewStr("md5"): λ.NewStr("84cd3c8aec4840046e5ab712416b73d0"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("33908820"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Wiadomości, 28.09.2017, 19:30"),
						λ.NewStr("description"): λ.NewStr("Wydanie główne codziennego serwisu informacyjnego."),
					}),
					λ.NewStr("skip"): λ.NewStr("HTTP Error 404: Not Found"),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://vod.tvp.pl/seriale/obyczajowe/na-sygnale/sezon-2-27-/odc-39/17834272"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://wiadomosci.tvp.pl/25169746/24052016-1200"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://krakow.tvp.pl/25511623/25lecie-mck-wyjatkowe-miejsce-na-mapie-krakowa"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://teleexpress.tvp.pl/25522307/wierni-wzieli-udzial-w-procesjach"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://sport.tvp.pl/25522165/krychowiak-uspokaja-w-sprawie-kontuzji-dwa-tygodnie-to-maksimum"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://www.tvp.info/25511919/trwa-rewolucja-wladza-zdecydowala-sie-na-pogwalcenie-konstytucji"),
					λ.NewStr("only_matching"): λ.True,
				}),
			)
			TVPIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒpage_id  λ.Object
						ϒself     = λargs[0]
						ϒurl      = λargs[1]
						ϒvideo_id λ.Object
						ϒwebpage  λ.Object
					)
					ϒpage_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒpage_id)
					ϒvideo_id = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewList(
							λ.NewStr("<iframe[^>]+src=\"[^\"]*?object_id=(\\d+)"),
							λ.NewStr("object_id\\s*:\\s*'(\\d+)'"),
							λ.NewStr("data-video-id=\"(\\d+)\""),
						),
						ϒwebpage,
						λ.NewStr("video id"),
					), λ.KWArgs{
						{Name: "default", Value: ϒpage_id},
					})
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("_type"): λ.NewStr("url_transparent"),
						λ.NewStr("url"):   λ.Add(λ.NewStr("tvp:"), ϒvideo_id),
						λ.NewStr("description"): func() λ.Object {
							if λv := λ.Call(λ.GetAttr(ϒself, "_og_search_description", nil), λ.NewArgs(ϒwebpage), λ.KWArgs{
								{Name: "default", Value: λ.None},
							}); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Call(λ.GetAttr(ϒself, "_html_search_meta", nil), λ.NewArgs(
									λ.NewStr("description"),
									ϒwebpage,
								), λ.KWArgs{
									{Name: "default", Value: λ.None},
								})
							}
						}(),
						λ.NewStr("thumbnail"): λ.Call(λ.GetAttr(ϒself, "_og_search_thumbnail", nil), λ.NewArgs(ϒwebpage), λ.KWArgs{
							{Name: "default", Value: λ.None},
						}),
						λ.NewStr("ie_key"): λ.NewStr("TVPEmbed"),
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("IE_NAME"):       TVPIE_IE_NAME,
				λ.NewStr("_TESTS"):        TVPIE__TESTS,
				λ.NewStr("_VALID_URL"):    TVPIE__VALID_URL,
				λ.NewStr("_real_extract"): TVPIE__real_extract,
			})
		}())
		TVPEmbedIE = λ.Cal(λ.TypeType, λ.NewStr("TVPEmbedIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				TVPEmbedIE_IE_NAME       λ.Object
				TVPEmbedIE__TESTS        λ.Object
				TVPEmbedIE__VALID_URL    λ.Object
				TVPEmbedIE__real_extract λ.Object
			)
			TVPEmbedIE_IE_NAME = λ.NewStr("tvp:embed")
			TVPEmbedIE__VALID_URL = λ.NewStr("(?:tvp:|https?://[^/]+\\.tvp\\.(?:pl|info)/sess/tvplayer\\.php\\?.*?object_id=)(?P<id>\\d+)")
			TVPEmbedIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("tvp:194536"),
					λ.NewStr("md5"): λ.NewStr("a21eb0aa862f25414430f15fdfb9e76c"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):    λ.NewStr("194536"),
						λ.NewStr("ext"):   λ.NewStr("mp4"),
						λ.NewStr("title"): λ.NewStr("Czas honoru, odc. 13 – Władek"),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.tvp.pl/sess/tvplayer.php?object_id=22670268"),
					λ.NewStr("md5"): λ.NewStr("8c9cd59d16edabf39331f93bf8a766c7"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):    λ.NewStr("22670268"),
						λ.NewStr("ext"):   λ.NewStr("mp4"),
						λ.NewStr("title"): λ.NewStr("Panorama, 07.12.2015, 15:40"),
					}),
					λ.NewStr("skip"): λ.NewStr("Transmisja została zakończona lub materiał niedostępny"),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("tvp:22670268"),
					λ.NewStr("only_matching"): λ.True,
				}),
			)
			TVPEmbedIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒerror          λ.Object
						ϒf              λ.Object
						ϒformats        λ.Object
						ϒhttp_url       λ.Object
						ϒi              λ.Object
						ϒm3u8_format    λ.Object
						ϒm3u8_formats   λ.Object
						ϒself           = λargs[0]
						ϒseries_title   λ.Object
						ϒthumbnail      λ.Object
						ϒtitle          λ.Object
						ϒurl            = λargs[1]
						ϒvideo_id       λ.Object
						ϒvideo_url      λ.Object
						ϒvideo_url_base λ.Object
						ϒwebpage        λ.Object
						τmp0            λ.Object
						τmp1            λ.Object
						τmp2            λ.Object
					)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), λ.Mod(λ.NewStr("http://www.tvp.pl/sess/tvplayer.php?object_id=%s"), ϒvideo_id), ϒvideo_id)
					ϒerror = func() λ.Object {
						if λv := λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
							λ.NewStr("(?s)<p[^>]+\\bclass=[\"\\']notAvailable__text[\"\\'][^>]*>(.+?)</p>"),
							ϒwebpage,
							λ.NewStr("error"),
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
						}); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Cal(ϒclean_html, λ.Cal(ϒget_element_by_attribute, λ.NewStr("class"), λ.NewStr("msg error"), ϒwebpage))
						}
					}()
					if λ.IsTrue(ϒerror) {
						panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.Mod(λ.NewStr("%s said: %s"), λ.NewTuple(
							λ.GetAttr(ϒself, "IE_NAME", nil),
							λ.Cal(ϒclean_html, ϒerror),
						))), λ.KWArgs{
							{Name: "expected", Value: λ.True},
						})))
					}
					ϒtitle = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr("name\\s*:\\s*([\\'\"])Title\\1\\s*,\\s*value\\s*:\\s*\\1(?P<title>.+?)\\1"),
						ϒwebpage,
						λ.NewStr("title"),
					), λ.KWArgs{
						{Name: "group", Value: λ.NewStr("title")},
					})
					ϒseries_title = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr("name\\s*:\\s*([\\'\"])SeriesTitle\\1\\s*,\\s*value\\s*:\\s*\\1(?P<series>.+?)\\1"),
						ϒwebpage,
						λ.NewStr("series"),
					), λ.KWArgs{
						{Name: "group", Value: λ.NewStr("series")},
						{Name: "default", Value: λ.None},
					})
					if λ.IsTrue(ϒseries_title) {
						ϒtitle = λ.Mod(λ.NewStr("%s, %s"), λ.NewTuple(
							ϒseries_title,
							ϒtitle,
						))
					}
					ϒthumbnail = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr("poster\\s*:\\s*'([^']+)'"),
						ϒwebpage,
						λ.NewStr("thumbnail"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					ϒvideo_url = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr("0:{src:([\\'\"])(?P<url>.*?)\\1"),
						ϒwebpage,
						λ.NewStr("formats"),
					), λ.KWArgs{
						{Name: "group", Value: λ.NewStr("url")},
						{Name: "default", Value: λ.None},
					})
					if λ.IsTrue(func() λ.Object {
						if λv := λ.NewBool(!λ.IsTrue(ϒvideo_url)); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.NewBool(λ.Contains(ϒvideo_url, λ.NewStr("material_niedostepny.mp4")))
						}
					}()) {
						ϒvideo_url = λ.GetItem(λ.Cal(λ.GetAttr(ϒself, "_download_json", nil), λ.Mod(λ.NewStr("http://www.tvp.pl/pub/stat/videofileinfo?video_id=%s"), ϒvideo_id), ϒvideo_id), λ.NewStr("video_url"))
					}
					ϒformats = λ.NewList()
					ϒvideo_url_base = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr("(https?://.+?/video)(?:\\.(?:ism|f4m|m3u8)|-\\d+\\.mp4)"),
						ϒvideo_url,
						λ.NewStr("video base url"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					if λ.IsTrue(ϒvideo_url_base) {
						λ.Cal(λ.GetAttr(ϒformats, "extend", nil), λ.Call(λ.GetAttr(ϒself, "_extract_ism_formats", nil), λ.NewArgs(
							λ.Add(ϒvideo_url_base, λ.NewStr(".ism/Manifest")),
							ϒvideo_id,
							λ.NewStr("mss"),
						), λ.KWArgs{
							{Name: "fatal", Value: λ.False},
						}))
						λ.Cal(λ.GetAttr(ϒformats, "extend", nil), λ.Call(λ.GetAttr(ϒself, "_extract_f4m_formats", nil), λ.NewArgs(
							λ.Add(ϒvideo_url_base, λ.NewStr(".ism/video.f4m")),
							ϒvideo_id,
						), λ.KWArgs{
							{Name: "f4m_id", Value: λ.NewStr("hds")},
							{Name: "fatal", Value: λ.False},
						}))
						ϒm3u8_formats = λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
							λ.Add(ϒvideo_url_base, λ.NewStr(".ism/video.m3u8")),
							ϒvideo_id,
							λ.NewStr("mp4"),
							λ.NewStr("m3u8_native"),
						), λ.KWArgs{
							{Name: "m3u8_id", Value: λ.NewStr("hls")},
							{Name: "fatal", Value: λ.False},
						})
						λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒm3u8_formats)
						ϒm3u8_formats = λ.Cal(λ.ListType, λ.Cal(λ.FilterIteratorType, λ.NewFunction("<lambda>",
							[]λ.Param{
								{Name: "f"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒf = λargs[0]
								)
								return λ.Ne(λ.Cal(λ.GetAttr(ϒf, "get", nil), λ.NewStr("vcodec")), λ.NewStr("none"))
							}), ϒm3u8_formats))
						λ.Cal(λ.GetAttr(ϒformats, "extend", nil), ϒm3u8_formats)
						τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.EnumerateIteratorType, ϒm3u8_formats, λ.NewInt(2)))
						for {
							if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
								break
							}
							τmp2 = τmp1
							ϒi = λ.GetItem(τmp2, λ.NewInt(0))
							ϒm3u8_format = λ.GetItem(τmp2, λ.NewInt(1))
							ϒhttp_url = λ.Mod(λ.NewStr("%s-%d.mp4"), λ.NewTuple(
								ϒvideo_url_base,
								ϒi,
							))
							if λ.IsTrue(λ.Cal(λ.GetAttr(ϒself, "_is_valid_url", nil), ϒhttp_url, ϒvideo_id)) {
								ϒf = λ.Cal(λ.GetAttr(ϒm3u8_format, "copy", nil))
								λ.Cal(λ.GetAttr(ϒf, "update", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
									λ.NewStr("url"):       ϒhttp_url,
									λ.NewStr("format_id"): λ.Cal(λ.GetAttr(λ.GetItem(ϒf, λ.NewStr("format_id")), "replace", nil), λ.NewStr("hls"), λ.NewStr("http")),
									λ.NewStr("protocol"):  λ.NewStr("http"),
								}))
								λ.Cal(λ.GetAttr(ϒformats, "append", nil), ϒf)
							}
						}
					} else {
						ϒformats = λ.NewList(λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("format_id"): λ.NewStr("direct"),
							λ.NewStr("url"):       ϒvideo_url,
							λ.NewStr("ext"):       λ.Cal(ϒdetermine_ext, ϒvideo_url, λ.NewStr("mp4")),
						}))
					}
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):        ϒvideo_id,
						λ.NewStr("title"):     ϒtitle,
						λ.NewStr("thumbnail"): ϒthumbnail,
						λ.NewStr("formats"):   ϒformats,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("IE_NAME"):       TVPEmbedIE_IE_NAME,
				λ.NewStr("_TESTS"):        TVPEmbedIE__TESTS,
				λ.NewStr("_VALID_URL"):    TVPEmbedIE__VALID_URL,
				λ.NewStr("_real_extract"): TVPEmbedIE__real_extract,
			})
		}())
		TVPWebsiteIE = λ.Cal(λ.TypeType, λ.NewStr("TVPWebsiteIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				TVPWebsiteIE__VALID_URL λ.Object
			)
			TVPWebsiteIE__VALID_URL = λ.NewStr("https?://vod\\.tvp\\.pl/website/(?P<display_id>[^,]+),(?P<id>\\d+)")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): TVPWebsiteIE__VALID_URL,
			})
		}())
	})
}
