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
 * dvtv/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/dvtv.py
 */

package dvtv

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	DVTVIE         λ.Object
	ExtractorError λ.Object
	InfoExtractor  λ.Object
	ϒdetermine_ext λ.Object
	ϒint_or_none   λ.Object
	ϒjs_to_json    λ.Object
	ϒmimetype2ext  λ.Object
	ϒparse_iso8601 λ.Object
	ϒtry_get       λ.Object
	ϒunescapeHTML  λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒdetermine_ext = Ωutils.ϒdetermine_ext
		ExtractorError = Ωutils.ExtractorError
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒjs_to_json = Ωutils.ϒjs_to_json
		ϒmimetype2ext = Ωutils.ϒmimetype2ext
		ϒtry_get = Ωutils.ϒtry_get
		ϒunescapeHTML = Ωutils.ϒunescapeHTML
		ϒparse_iso8601 = Ωutils.ϒparse_iso8601
		DVTVIE = λ.Cal(λ.TypeType, λ.NewStr("DVTVIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				DVTVIE_IE_NAME               λ.Object
				DVTVIE__TESTS                λ.Object
				DVTVIE__VALID_URL            λ.Object
				DVTVIE__parse_video_metadata λ.Object
				DVTVIE__real_extract         λ.Object
			)
			DVTVIE_IE_NAME = λ.NewStr("dvtv")
			DVTVIE__VALID_URL = λ.NewStr("https?://video\\.aktualne\\.cz/(?:[^/]+/)+r~(?P<id>[0-9a-f]{32})")
			DVTVIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://video.aktualne.cz/dvtv/vondra-o-ceskem-stoleti-pri-pohledu-na-havla-mi-bylo-trapne/r~e5efe9ca855511e4833a0025900fea04/"),
					λ.NewStr("md5"): λ.NewStr("67cb83e4a955d36e1b5d31993134a0c2"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("dc0768de855511e49e4b0025900fea04"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Vondra o Českém století: Při pohledu na Havla mi bylo trapně"),
						λ.NewStr("duration"):    λ.NewInt(1484),
						λ.NewStr("upload_date"): λ.NewStr("20141217"),
						λ.NewStr("timestamp"):   λ.NewInt(1418792400),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://video.aktualne.cz/dvtv/dvtv-16-12-2014-utok-talibanu-boj-o-kliniku-uprchlici/r~973eb3bc854e11e498be002590604f2e/"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("title"): λ.NewStr("DVTV 16. 12. 2014: útok Talibanu, boj o kliniku, uprchlíci"),
						λ.NewStr("id"):    λ.NewStr("973eb3bc854e11e498be002590604f2e"),
					}),
					λ.NewStr("playlist"): λ.NewList(
						λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("md5"): λ.NewStr("da7ca6be4935532241fa9520b3ad91e4"),
							λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
								λ.NewStr("id"):          λ.NewStr("b0b40906854d11e4bdad0025900fea04"),
								λ.NewStr("ext"):         λ.NewStr("mp4"),
								λ.NewStr("title"):       λ.NewStr("Drtinová Veselovský TV 16. 12. 2014: Témata dne"),
								λ.NewStr("description"): λ.NewStr("md5:0916925dea8e30fe84222582280b47a0"),
								λ.NewStr("timestamp"):   λ.NewInt(1418760010),
								λ.NewStr("upload_date"): λ.NewStr("20141216"),
							}),
						}),
						λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("md5"): λ.NewStr("5f7652a08b05009c1292317b449ffea2"),
							λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
								λ.NewStr("id"):          λ.NewStr("420ad9ec854a11e4bdad0025900fea04"),
								λ.NewStr("ext"):         λ.NewStr("mp4"),
								λ.NewStr("title"):       λ.NewStr("Školní masakr možná změní boj s Talibanem, říká novinářka"),
								λ.NewStr("description"): λ.NewStr("md5:ff2f9f6de73c73d7cef4f756c1c1af42"),
								λ.NewStr("timestamp"):   λ.NewInt(1418760010),
								λ.NewStr("upload_date"): λ.NewStr("20141216"),
							}),
						}),
						λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("md5"): λ.NewStr("498eb9dfa97169f409126c617e2a3d64"),
							λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
								λ.NewStr("id"):          λ.NewStr("95d35580846a11e4b6d20025900fea04"),
								λ.NewStr("ext"):         λ.NewStr("mp4"),
								λ.NewStr("title"):       λ.NewStr("Boj o kliniku: Veřejný zájem, nebo právo na majetek?"),
								λ.NewStr("description"): λ.NewStr("md5:889fe610a70fee5511dc3326a089188e"),
								λ.NewStr("timestamp"):   λ.NewInt(1418760010),
								λ.NewStr("upload_date"): λ.NewStr("20141216"),
							}),
						}),
						λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("md5"): λ.NewStr("b8dc6b744844032dab6ba3781a7274b9"),
							λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
								λ.NewStr("id"):          λ.NewStr("6fe14d66853511e4833a0025900fea04"),
								λ.NewStr("ext"):         λ.NewStr("mp4"),
								λ.NewStr("title"):       λ.NewStr("Pánek: Odmítání syrských uprchlíků je ostudou české vlády"),
								λ.NewStr("description"): λ.NewStr("md5:544f86de6d20c4815bea11bf2ac3004f"),
								λ.NewStr("timestamp"):   λ.NewInt(1418760010),
								λ.NewStr("upload_date"): λ.NewStr("20141216"),
							}),
						}),
					),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("https://video.aktualne.cz/dvtv/zeman-si-jen-leci-mindraky-sobotku-nenavidi-a-babis-se-mu-te/r~960cdb3a365a11e7a83b0025900fea04/"),
					λ.NewStr("md5"): λ.NewStr("f8efe9656017da948369aa099788c8ea"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("3c496fec365911e7a6500025900fea04"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Zeman si jen léčí mindráky, Sobotku nenávidí a Babiš se mu teď hodí, tvrdí Kmenta"),
						λ.NewStr("duration"):    λ.NewInt(1103),
						λ.NewStr("upload_date"): λ.NewStr("20170511"),
						λ.NewStr("timestamp"):   λ.NewInt(1494514200),
					}),
					λ.NewStr("params"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("skip_download"): λ.True,
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://video.aktualne.cz/v-cechach-poprve-zazni-zelenkova-zrestaurovana-mse/r~45b4b00483ec11e4883b002590604f2e/"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("https://video.aktualne.cz/dvtv/zive-mistryne-sveta-eva-samkova-po-navratu-ze-sampionatu/r~182654c2288811e990fd0cc47ab5f122/"),
					λ.NewStr("md5"): λ.NewStr("2e552e483f2414851ca50467054f9d5d"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("8d116360288011e98c840cc47ab5f122"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Živě: Mistryně světa Eva Samková po návratu ze šampionátu"),
						λ.NewStr("upload_date"): λ.NewStr("20190204"),
						λ.NewStr("timestamp"):   λ.NewInt(1549289591),
					}),
					λ.NewStr("params"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("skip_download"): λ.True,
					}),
				}),
			)
			DVTVIE__parse_video_metadata = λ.NewFunction("_parse_video_metadata",
				[]λ.Param{
					{Name: "self"},
					{Name: "js"},
					{Name: "video_id"},
					{Name: "timestamp"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdata         λ.Object
						ϒext          λ.Object
						ϒf            λ.Object
						ϒformat_id    λ.Object
						ϒformats      λ.Object
						ϒheight       λ.Object
						ϒjs           = λargs[1]
						ϒlabel        λ.Object
						ϒlive_starter λ.Object
						ϒself         = λargs[0]
						ϒtimestamp    = λargs[3]
						ϒtitle        λ.Object
						ϒtracks       λ.Object
						ϒvideo        λ.Object
						ϒvideo_id     = λargs[2]
						ϒvideo_type   λ.Object
						ϒvideo_url    λ.Object
						τmp0          λ.Object
						τmp1          λ.Object
						τmp2          λ.Object
						τmp3          λ.Object
						τmp4          λ.Object
						τmp5          λ.Object
					)
					ϒdata = λ.Call(λ.GetAttr(ϒself, "_parse_json", nil), λ.NewArgs(
						ϒjs,
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "transform_source", Value: ϒjs_to_json},
					})
					ϒtitle = λ.Cal(ϒunescapeHTML, λ.GetItem(ϒdata, λ.NewStr("title")))
					ϒlive_starter = λ.Cal(ϒtry_get, ϒdata, λ.NewFunction("<lambda>",
						[]λ.Param{
							{Name: "x"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒx = λargs[0]
							)
							return λ.GetItem(λ.GetItem(ϒx, λ.NewStr("plugins")), λ.NewStr("liveStarter"))
						}), λ.DictType)
					if λ.IsTrue(ϒlive_starter) {
						λ.Cal(λ.GetAttr(ϒdata, "update", nil), ϒlive_starter)
					}
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒdata, "get", nil), λ.NewStr("tracks"), λ.NewDictWithTable(map[λ.Object]λ.Object{})), "values", nil)))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒtracks = τmp1
						τmp2 = λ.Cal(λ.BuiltinIter, ϒtracks)
						for {
							if τmp3 = λ.NextDefault(τmp2, λ.AfterLast); τmp3 == λ.AfterLast {
								break
							}
							ϒvideo = τmp3
							ϒvideo_url = λ.Cal(λ.GetAttr(ϒvideo, "get", nil), λ.NewStr("src"))
							if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒvideo_url))) {
								continue
							}
							ϒvideo_type = λ.Cal(λ.GetAttr(ϒvideo, "get", nil), λ.NewStr("type"))
							ϒext = λ.Cal(ϒdetermine_ext, ϒvideo_url, λ.Cal(ϒmimetype2ext, ϒvideo_type))
							if λ.IsTrue(func() λ.Object {
								if λv := λ.Eq(ϒvideo_type, λ.NewStr("application/vnd.apple.mpegurl")); λ.IsTrue(λv) {
									return λv
								} else {
									return λ.Eq(ϒext, λ.NewStr("m3u8"))
								}
							}()) {
								λ.Cal(λ.GetAttr(ϒformats, "extend", nil), λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
									ϒvideo_url,
									ϒvideo_id,
									λ.NewStr("mp4"),
								), λ.KWArgs{
									{Name: "entry_protocol", Value: λ.NewStr("m3u8_native")},
									{Name: "m3u8_id", Value: λ.NewStr("hls")},
									{Name: "fatal", Value: λ.False},
								}))
							} else {
								if λ.IsTrue(func() λ.Object {
									if λv := λ.Eq(ϒvideo_type, λ.NewStr("application/dash+xml")); λ.IsTrue(λv) {
										return λv
									} else {
										return λ.Eq(ϒext, λ.NewStr("mpd"))
									}
								}()) {
									λ.Cal(λ.GetAttr(ϒformats, "extend", nil), λ.Call(λ.GetAttr(ϒself, "_extract_mpd_formats", nil), λ.NewArgs(
										ϒvideo_url,
										ϒvideo_id,
									), λ.KWArgs{
										{Name: "mpd_id", Value: λ.NewStr("dash")},
										{Name: "fatal", Value: λ.False},
									}))
								} else {
									ϒlabel = λ.Cal(λ.GetAttr(ϒvideo, "get", nil), λ.NewStr("label"))
									ϒheight = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
										λ.NewStr("^(\\d+)[pP]"),
										func() λ.Object {
											if λv := ϒlabel; λ.IsTrue(λv) {
												return λv
											} else {
												return λ.NewStr("")
											}
										}(),
										λ.NewStr("height"),
									), λ.KWArgs{
										{Name: "default", Value: λ.None},
									})
									ϒformat_id = λ.NewList(λ.NewStr("http"))
									τmp4 = λ.Cal(λ.BuiltinIter, λ.NewTuple(
										ϒext,
										ϒlabel,
									))
									for {
										if τmp5 = λ.NextDefault(τmp4, λ.AfterLast); τmp5 == λ.AfterLast {
											break
										}
										ϒf = τmp5
										if λ.IsTrue(ϒf) {
											λ.Cal(λ.GetAttr(ϒformat_id, "append", nil), ϒf)
										}
									}
									λ.Cal(λ.GetAttr(ϒformats, "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
										λ.NewStr("url"):       ϒvideo_url,
										λ.NewStr("format_id"): λ.Cal(λ.GetAttr(λ.NewStr("-"), "join", nil), ϒformat_id),
										λ.NewStr("height"):    λ.Cal(ϒint_or_none, ϒheight),
									}))
								}
							}
						}
					}
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"): func() λ.Object {
							if λv := λ.Cal(λ.GetAttr(ϒdata, "get", nil), λ.NewStr("mediaid")); λ.IsTrue(λv) {
								return λv
							} else {
								return ϒvideo_id
							}
						}(),
						λ.NewStr("title"):       ϒtitle,
						λ.NewStr("description"): λ.Cal(λ.GetAttr(ϒdata, "get", nil), λ.NewStr("description")),
						λ.NewStr("thumbnail"):   λ.Cal(λ.GetAttr(ϒdata, "get", nil), λ.NewStr("image")),
						λ.NewStr("duration"):    λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒdata, "get", nil), λ.NewStr("duration"))),
						λ.NewStr("timestamp"):   λ.Cal(ϒint_or_none, ϒtimestamp),
						λ.NewStr("formats"):     ϒformats,
					})
				})
			DVTVIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒitem      λ.Object
						ϒitems     λ.Object
						ϒself      = λargs[0]
						ϒtimestamp λ.Object
						ϒurl       = λargs[1]
						ϒvideo_id  λ.Object
						ϒwebpage   λ.Object
					)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒvideo_id)
					ϒtimestamp = λ.Cal(ϒparse_iso8601, λ.Call(λ.GetAttr(ϒself, "_html_search_meta", nil), λ.NewArgs(
						λ.NewStr("article:published_time"),
						ϒwebpage,
						λ.NewStr("published time"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					}))
					ϒitems = λ.Cal(Ωre.ϒfindall, λ.NewStr("(?s)playlist\\.push\\(({.+?})\\);"), ϒwebpage)
					if λ.IsTrue(ϒitems) {
						return λ.Cal(λ.GetAttr(ϒself, "playlist_result", nil), λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
							nil,
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								return λ.NewGenerator(func(λgen λ.Generator) λ.Object {
									var (
										ϒi   λ.Object
										τmp0 λ.Object
										τmp1 λ.Object
									)
									τmp0 = λ.Cal(λ.BuiltinIter, ϒitems)
									for {
										if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
											break
										}
										ϒi = τmp1
										λgen.Yield(λ.Cal(λ.GetAttr(ϒself, "_parse_video_metadata", nil), ϒi, ϒvideo_id, ϒtimestamp))
									}
									return λ.None
								})
							}))), ϒvideo_id, λ.Cal(λ.GetAttr(ϒself, "_html_search_meta", nil), λ.NewStr("twitter:title"), ϒwebpage))
					}
					ϒitem = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr("(?s)BBXPlayer\\.setup\\((.+?)\\);"),
						ϒwebpage,
						λ.NewStr("video"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					if λ.IsTrue(ϒitem) {
						ϒitem = λ.Cal(Ωre.ϒsub, λ.NewStr("\\w+?\\((.+)\\)"), λ.NewStr("\\1"), ϒitem)
						return λ.Cal(λ.GetAttr(ϒself, "_parse_video_metadata", nil), ϒitem, ϒvideo_id, ϒtimestamp)
					}
					panic(λ.Raise(λ.Cal(ExtractorError, λ.NewStr("Could not find neither video nor playlist"))))
					return λ.None
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("IE_NAME"):               DVTVIE_IE_NAME,
				λ.NewStr("_TESTS"):                DVTVIE__TESTS,
				λ.NewStr("_VALID_URL"):            DVTVIE__VALID_URL,
				λ.NewStr("_parse_video_metadata"): DVTVIE__parse_video_metadata,
				λ.NewStr("_real_extract"):         DVTVIE__real_extract,
			})
		}())
	})
}