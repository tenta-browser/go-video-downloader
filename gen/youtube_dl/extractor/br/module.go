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
 * br/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/br.py
 */

package br

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	BRIE            λ.Object
	BRMediathekIE   λ.Object
	ExtractorError  λ.Object
	InfoExtractor   λ.Object
	ϒdetermine_ext  λ.Object
	ϒint_or_none    λ.Object
	ϒparse_duration λ.Object
	ϒparse_iso8601  λ.Object
	ϒxpath_element  λ.Object
	ϒxpath_text     λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒdetermine_ext = Ωutils.ϒdetermine_ext
		ExtractorError = Ωutils.ExtractorError
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒparse_duration = Ωutils.ϒparse_duration
		ϒparse_iso8601 = Ωutils.ϒparse_iso8601
		ϒxpath_element = Ωutils.ϒxpath_element
		ϒxpath_text = Ωutils.ϒxpath_text
		BRIE = λ.Cal(λ.TypeType, λ.NewStr("BRIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				BRIE__TESTS              λ.Object
				BRIE__VALID_URL          λ.Object
				BRIE__extract_formats    λ.Object
				BRIE__extract_thumbnails λ.Object
				BRIE__real_extract       λ.Object
			)
			BRIE__VALID_URL = λ.NewStr("(?P<base_url>https?://(?:www\\.)?br(?:-klassik)?\\.de)/(?:[a-z0-9\\-_]+/)+(?P<id>[a-z0-9\\-_]+)\\.html")
			BRIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.br.de/mediathek/video/sendungen/abendschau/betriebliche-altersvorsorge-104.html"),
					λ.NewStr("md5"): λ.NewStr("83a0477cf0b8451027eb566d88b51106"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("48f656ef-287e-486f-be86-459122db22cc"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Die böse Überraschung"),
						λ.NewStr("description"): λ.NewStr("md5:ce9ac81b466ce775b8018f6801b48ac9"),
						λ.NewStr("duration"):    λ.NewInt(180),
						λ.NewStr("uploader"):    λ.NewStr("Reinhard Weber"),
						λ.NewStr("upload_date"): λ.NewStr("20150422"),
					}),
					λ.NewStr("skip"): λ.NewStr("404 not found"),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.br.de/nachrichten/oberbayern/inhalt/muenchner-polizeipraesident-schreiber-gestorben-100.html"),
					λ.NewStr("md5"): λ.NewStr("af3a3a4aa43ff0ce6a89504c67f427ef"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("a4b83e34-123d-4b81-9f4e-c0d3121a4e05"),
						λ.NewStr("ext"):         λ.NewStr("flv"),
						λ.NewStr("title"):       λ.NewStr("Manfred Schreiber ist tot"),
						λ.NewStr("description"): λ.NewStr("md5:b454d867f2a9fc524ebe88c3f5092d97"),
						λ.NewStr("duration"):    λ.NewInt(26),
					}),
					λ.NewStr("skip"): λ.NewStr("404 not found"),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("https://www.br-klassik.de/audio/peeping-tom-premierenkritik-dance-festival-muenchen-100.html"),
					λ.NewStr("md5"): λ.NewStr("8b5b27c0b090f3b35eac4ab3f7a73d3d"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("74c603c9-26d3-48bb-b85b-079aeed66e0b"),
						λ.NewStr("ext"):         λ.NewStr("aac"),
						λ.NewStr("title"):       λ.NewStr("Kurzweilig und sehr bewegend"),
						λ.NewStr("description"): λ.NewStr("md5:0351996e3283d64adeb38ede91fac54e"),
						λ.NewStr("duration"):    λ.NewInt(296),
					}),
					λ.NewStr("skip"): λ.NewStr("404 not found"),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.br.de/radio/bayern1/service/team/videos/team-video-erdelt100.html"),
					λ.NewStr("md5"): λ.NewStr("dbab0aef2e047060ea7a21fc1ce1078a"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("6ba73750-d405-45d3-861d-1ce8c524e059"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Umweltbewusster Häuslebauer"),
						λ.NewStr("description"): λ.NewStr("md5:d52dae9792d00226348c1dbb13c9bae2"),
						λ.NewStr("duration"):    λ.NewInt(116),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.br.de/fernsehen/br-alpha/sendungen/kant-fuer-anfaenger/kritik-der-reinen-vernunft/kant-kritik-01-metaphysik100.html"),
					λ.NewStr("md5"): λ.NewStr("23bca295f1650d698f94fc570977dae3"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("d982c9ce-8648-4753-b358-98abb8aec43d"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Folge 1 - Metaphysik"),
						λ.NewStr("description"): λ.NewStr("md5:bb659990e9e59905c3d41e369db1fbe3"),
						λ.NewStr("duration"):    λ.NewInt(893),
						λ.NewStr("uploader"):    λ.NewStr("Eva Maria Steimle"),
						λ.NewStr("upload_date"): λ.NewStr("20170208"),
					}),
				}),
			)
			BRIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒbase_url       λ.Object
						ϒbroadcast_date λ.Object
						ϒdisplay_id     λ.Object
						ϒmedia          λ.Object
						ϒmedia_id       λ.Object
						ϒmedias         λ.Object
						ϒpage           λ.Object
						ϒself           = λargs[0]
						ϒurl            = λargs[1]
						ϒxml            λ.Object
						ϒxml_media      λ.Object
						ϒxml_url        λ.Object
						τmp0            λ.Object
						τmp1            λ.Object
					)
					τmp0 = λ.Cal(λ.GetAttr(λ.Cal(Ωre.ϒsearch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl), "groups", nil))
					ϒbase_url = λ.GetItem(τmp0, λ.NewInt(0))
					ϒdisplay_id = λ.GetItem(τmp0, λ.NewInt(1))
					ϒpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒdisplay_id)
					ϒxml_url = λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewStr("return BRavFramework\\.register\\(BRavFramework\\('avPlayer_(?:[a-f0-9-]{36})'\\)\\.setup\\({dataURL:'(/(?:[a-z0-9\\-]+/)+[a-z0-9/~_.-]+)'}\\)\\);"), ϒpage, λ.NewStr("XMLURL"))
					ϒxml = λ.Cal(λ.GetAttr(ϒself, "_download_xml", nil), λ.Add(ϒbase_url, ϒxml_url), ϒdisplay_id)
					ϒmedias = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Add(λ.Cal(λ.GetAttr(ϒxml, "findall", nil), λ.NewStr("video")), λ.Cal(λ.GetAttr(ϒxml, "findall", nil), λ.NewStr("audio"))))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒxml_media = τmp1
						ϒmedia_id = λ.Cal(λ.GetAttr(ϒxml_media, "get", nil), λ.NewStr("externalId"))
						ϒmedia = λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("id"):          ϒmedia_id,
							λ.NewStr("title"):       λ.Cal(ϒxpath_text, ϒxml_media, λ.NewStr("title"), λ.NewStr("title"), λ.True),
							λ.NewStr("duration"):    λ.Cal(ϒparse_duration, λ.Cal(ϒxpath_text, ϒxml_media, λ.NewStr("duration"))),
							λ.NewStr("formats"):     λ.Cal(λ.GetAttr(ϒself, "_extract_formats", nil), λ.Cal(ϒxpath_element, ϒxml_media, λ.NewStr("assets")), ϒmedia_id),
							λ.NewStr("thumbnails"):  λ.Cal(λ.GetAttr(ϒself, "_extract_thumbnails", nil), λ.Cal(ϒxpath_element, ϒxml_media, λ.NewStr("teaserImage/variants")), ϒbase_url),
							λ.NewStr("description"): λ.Cal(ϒxpath_text, ϒxml_media, λ.NewStr("desc")),
							λ.NewStr("webpage_url"): λ.Cal(ϒxpath_text, ϒxml_media, λ.NewStr("permalink")),
							λ.NewStr("uploader"):    λ.Cal(ϒxpath_text, ϒxml_media, λ.NewStr("author")),
						})
						ϒbroadcast_date = λ.Cal(ϒxpath_text, ϒxml_media, λ.NewStr("broadcastDate"))
						if λ.IsTrue(ϒbroadcast_date) {
							λ.SetItem(ϒmedia, λ.NewStr("upload_date"), λ.Cal(λ.GetAttr(λ.NewStr(""), "join", nil), λ.Cal(λ.ReversedIteratorType, λ.Cal(λ.GetAttr(ϒbroadcast_date, "split", nil), λ.NewStr(".")))))
						}
						λ.Cal(λ.GetAttr(ϒmedias, "append", nil), ϒmedia)
					}
					if λ.IsTrue(λ.Gt(λ.Cal(λ.BuiltinLen, ϒmedias), λ.NewInt(1))) {
						λ.Cal(λ.GetAttr(λ.GetAttr(ϒself, "_downloader", nil), "report_warning", nil), λ.NewStr("found multiple medias; please report this with the video URL to http://yt-dl.org/bug"))
					}
					if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒmedias))) {
						panic(λ.Raise(λ.Cal(ExtractorError, λ.NewStr("No media entries found"))))
					}
					return λ.GetItem(ϒmedias, λ.NewInt(0))
				})
			BRIE__extract_formats = λ.NewFunction("_extract_formats",
				[]λ.Param{
					{Name: "self"},
					{Name: "assets"},
					{Name: "media_id"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒasset            λ.Object
						ϒasset_type       λ.Object
						ϒassets           = λargs[1]
						ϒformat_info      λ.Object
						ϒformat_url       λ.Object
						ϒformats          λ.Object
						ϒhttp_format_info λ.Object
						ϒmedia_id         = λargs[2]
						ϒrtmp_format_info λ.Object
						ϒself             = λargs[0]
						ϒserver_prefix    λ.Object
						τmp0              λ.Object
						τmp1              λ.Object
					)
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(ϒassets, "findall", nil), λ.NewStr("asset")))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒasset = τmp1
						ϒformat_url = λ.Cal(ϒxpath_text, ϒasset, λ.NewList(
							λ.NewStr("downloadUrl"),
							λ.NewStr("url"),
						))
						ϒasset_type = λ.Cal(λ.GetAttr(ϒasset, "get", nil), λ.NewStr("type"))
						if λ.IsTrue(λ.Cal(λ.GetAttr(ϒasset_type, "startswith", nil), λ.NewStr("HDS"))) {
							λ.Cal(λ.GetAttr(ϒformats, "extend", nil), λ.Call(λ.GetAttr(ϒself, "_extract_f4m_formats", nil), λ.NewArgs(
								λ.Add(ϒformat_url, λ.NewStr("?hdcore=3.2.0")),
								ϒmedia_id,
							), λ.KWArgs{
								{Name: "f4m_id", Value: λ.NewStr("hds")},
								{Name: "fatal", Value: λ.False},
							}))
						} else {
							if λ.IsTrue(λ.Cal(λ.GetAttr(ϒasset_type, "startswith", nil), λ.NewStr("HLS"))) {
								λ.Cal(λ.GetAttr(ϒformats, "extend", nil), λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
									ϒformat_url,
									ϒmedia_id,
									λ.NewStr("mp4"),
									λ.NewStr("m3u8_native"),
								), λ.KWArgs{
									{Name: "m3u8_id", Value: λ.NewStr("hds")},
									{Name: "fatal", Value: λ.False},
								}))
							} else {
								ϒformat_info = λ.NewDictWithTable(map[λ.Object]λ.Object{
									λ.NewStr("ext"):       λ.Cal(ϒxpath_text, ϒasset, λ.NewStr("mediaType")),
									λ.NewStr("width"):     λ.Cal(ϒint_or_none, λ.Cal(ϒxpath_text, ϒasset, λ.NewStr("frameWidth"))),
									λ.NewStr("height"):    λ.Cal(ϒint_or_none, λ.Cal(ϒxpath_text, ϒasset, λ.NewStr("frameHeight"))),
									λ.NewStr("tbr"):       λ.Cal(ϒint_or_none, λ.Cal(ϒxpath_text, ϒasset, λ.NewStr("bitrateVideo"))),
									λ.NewStr("abr"):       λ.Cal(ϒint_or_none, λ.Cal(ϒxpath_text, ϒasset, λ.NewStr("bitrateAudio"))),
									λ.NewStr("vcodec"):    λ.Cal(ϒxpath_text, ϒasset, λ.NewStr("codecVideo")),
									λ.NewStr("acodec"):    λ.Cal(ϒxpath_text, ϒasset, λ.NewStr("codecAudio")),
									λ.NewStr("container"): λ.Cal(ϒxpath_text, ϒasset, λ.NewStr("mediaType")),
									λ.NewStr("filesize"):  λ.Cal(ϒint_or_none, λ.Cal(ϒxpath_text, ϒasset, λ.NewStr("size"))),
								})
								ϒformat_url = λ.Cal(λ.GetAttr(ϒself, "_proto_relative_url", nil), ϒformat_url)
								if λ.IsTrue(ϒformat_url) {
									ϒhttp_format_info = λ.Cal(λ.GetAttr(ϒformat_info, "copy", nil))
									λ.Cal(λ.GetAttr(ϒhttp_format_info, "update", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
										λ.NewStr("url"):       ϒformat_url,
										λ.NewStr("format_id"): λ.Mod(λ.NewStr("http-%s"), ϒasset_type),
									}))
									λ.Cal(λ.GetAttr(ϒformats, "append", nil), ϒhttp_format_info)
								}
								ϒserver_prefix = λ.Cal(ϒxpath_text, ϒasset, λ.NewStr("serverPrefix"))
								if λ.IsTrue(ϒserver_prefix) {
									ϒrtmp_format_info = λ.Cal(λ.GetAttr(ϒformat_info, "copy", nil))
									λ.Cal(λ.GetAttr(ϒrtmp_format_info, "update", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
										λ.NewStr("url"):       ϒserver_prefix,
										λ.NewStr("play_path"): λ.Cal(ϒxpath_text, ϒasset, λ.NewStr("fileName")),
										λ.NewStr("format_id"): λ.Mod(λ.NewStr("rtmp-%s"), ϒasset_type),
									}))
									λ.Cal(λ.GetAttr(ϒformats, "append", nil), ϒrtmp_format_info)
								}
							}
						}
					}
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					return ϒformats
				})
			BRIE__extract_thumbnails = λ.NewFunction("_extract_thumbnails",
				[]λ.Param{
					{Name: "self"},
					{Name: "variants"},
					{Name: "base_url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒbase_url   = λargs[2]
						ϒself       = λargs[0]
						ϒthumbnails λ.Object
						ϒvariants   = λargs[1]
					)
					_ = ϒself
					ϒthumbnails = λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
						nil,
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							return λ.NewGenerator(func(λgen λ.Generator) λ.Object {
								var (
									ϒvariant λ.Object
									τmp0     λ.Object
									τmp1     λ.Object
								)
								τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(ϒvariants, "findall", nil), λ.NewStr("variant")))
								for {
									if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
										break
									}
									ϒvariant = τmp1
									if λ.IsTrue(λ.Cal(ϒxpath_text, ϒvariant, λ.NewStr("url"))) {
										λgen.Yield(λ.NewDictWithTable(map[λ.Object]λ.Object{
											λ.NewStr("url"):    λ.Add(ϒbase_url, λ.Cal(ϒxpath_text, ϒvariant, λ.NewStr("url"))),
											λ.NewStr("width"):  λ.Cal(ϒint_or_none, λ.Cal(ϒxpath_text, ϒvariant, λ.NewStr("width"))),
											λ.NewStr("height"): λ.Cal(ϒint_or_none, λ.Cal(ϒxpath_text, ϒvariant, λ.NewStr("height"))),
										}))
									}
								}
								return λ.None
							})
						})))
					λ.Call(λ.GetAttr(ϒthumbnails, "sort", nil), nil, λ.KWArgs{
						{Name: "key", Value: λ.NewFunction("<lambda>",
							[]λ.Param{
								{Name: "x"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒx = λargs[0]
								)
								return λ.Mul(λ.GetItem(ϒx, λ.NewStr("width")), λ.GetItem(ϒx, λ.NewStr("height")))
							})},
						{Name: "reverse", Value: λ.True},
					})
					return ϒthumbnails
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TESTS"):              BRIE__TESTS,
				λ.NewStr("_VALID_URL"):          BRIE__VALID_URL,
				λ.NewStr("_extract_formats"):    BRIE__extract_formats,
				λ.NewStr("_extract_thumbnails"): BRIE__extract_thumbnails,
				λ.NewStr("_real_extract"):       BRIE__real_extract,
			})
		}())
		BRMediathekIE = λ.Cal(λ.TypeType, λ.NewStr("BRMediathekIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				BRMediathekIE__VALID_URL λ.Object
			)
			BRMediathekIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?br\\.de/mediathek/video/[^/?&#]*?-(?P<id>av:[0-9a-f]{24})")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): BRMediathekIE__VALID_URL,
			})
		}())
	})
}
