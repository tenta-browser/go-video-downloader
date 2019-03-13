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
 * rutv/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/rutv.py
 */

package rutv

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError λ.Object
	InfoExtractor  λ.Object
	RUTVIE         λ.Object
	ϒint_or_none   λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ExtractorError = Ωutils.ExtractorError
		ϒint_or_none = Ωutils.ϒint_or_none
		RUTVIE = λ.Cal(λ.TypeType, λ.NewStr("RUTVIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				RUTVIE__TESTS        λ.Object
				RUTVIE__VALID_URL    λ.Object
				RUTVIE__extract_url  λ.Object
				RUTVIE__real_extract λ.Object
			)
			RUTVIE__VALID_URL = λ.NewStr("(?x)\n                    https?://\n                        (?:test)?player\\.(?:rutv\\.ru|vgtrk\\.com)/\n                        (?P<path>\n                            flash\\d+v/container\\.swf\\?id=|\n                            iframe/(?P<type>swf|video|live)/id/|\n                            index/iframe/cast_id/\n                        )\n                        (?P<id>\\d+)\n                    ")
			RUTVIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://player.rutv.ru/flash2v/container.swf?id=774471&sid=kultura&fbv=true&isPlay=true&ssl=false&i=560&acc_video_id=episode_id/972347/video_id/978186/brand_id/31724"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("774471"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Монологи на все времена"),
						λ.NewStr("description"): λ.NewStr("md5:18d8b5e6a41fb1faa53819471852d5d5"),
						λ.NewStr("duration"):    λ.NewInt(2906),
					}),
					λ.NewStr("params"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("skip_download"): λ.True,
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("https://player.vgtrk.com/flash2v/container.swf?id=774016&sid=russiatv&fbv=true&isPlay=true&ssl=false&i=560&acc_video_id=episode_id/972098/video_id/977760/brand_id/57638"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("774016"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Чужой в семье Сталина"),
						λ.NewStr("description"): λ.NewStr(""),
						λ.NewStr("duration"):    λ.NewInt(2539),
					}),
					λ.NewStr("params"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("skip_download"): λ.True,
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://player.rutv.ru/iframe/swf/id/766888/sid/hitech/?acc_video_id=4000"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("766888"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Вести.net: интернет-гиганты начали перетягивание программных \"одеял\""),
						λ.NewStr("description"): λ.NewStr("md5:65ddd47f9830c4f42ed6475f8730c995"),
						λ.NewStr("duration"):    λ.NewInt(279),
					}),
					λ.NewStr("params"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("skip_download"): λ.True,
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://player.rutv.ru/iframe/video/id/771852/start_zoom/true/showZoomBtn/false/sid/russiatv/?acc_video_id=episode_id/970443/video_id/975648/brand_id/5169"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("771852"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Прямой эфир. Жертвы загадочной болезни: смерть от старости в 17 лет"),
						λ.NewStr("description"): λ.NewStr("md5:b81c8c55247a4bd996b43ce17395b2d8"),
						λ.NewStr("duration"):    λ.NewInt(3096),
					}),
					λ.NewStr("params"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("skip_download"): λ.True,
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://player.rutv.ru/iframe/live/id/51499/showZoomBtn/false/isPlay/true/sid/sochi2014"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("51499"),
						λ.NewStr("ext"):         λ.NewStr("flv"),
						λ.NewStr("title"):       λ.NewStr("Сочи-2014. Биатлон. Индивидуальная гонка. Мужчины "),
						λ.NewStr("description"): λ.NewStr("md5:9e0ed5c9d2fa1efbfdfed90c9a6d179c"),
					}),
					λ.NewStr("skip"): λ.NewStr("Translation has finished"),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://player.rutv.ru/iframe/live/id/21/showZoomBtn/false/isPlay/true/"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):      λ.NewStr("21"),
						λ.NewStr("ext"):     λ.NewStr("mp4"),
						λ.NewStr("title"):   λ.NewStr("re:^Россия 24. Прямой эфир [0-9]{4}-[0-9]{2}-[0-9]{2} [0-9]{2}:[0-9]{2}$"),
						λ.NewStr("is_live"): λ.True,
					}),
					λ.NewStr("params"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("skip_download"): λ.True,
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://testplayer.vgtrk.com/iframe/live/id/19201/showZoomBtn/false/isPlay/true/"),
					λ.NewStr("only_matching"): λ.True,
				}),
			)
			RUTVIE__extract_url = λ.NewFunction("_extract_url",
				[]λ.Param{
					{Name: "cls"},
					{Name: "webpage"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒcls     = λargs[0]
						ϒmobj    λ.Object
						ϒwebpage = λargs[1]
					)
					_ = ϒcls
					ϒmobj = λ.Cal(Ωre.ϒsearch, λ.NewStr("<iframe[^>]+?src=([\"\\'])(?P<url>https?://(?:test)?player\\.(?:rutv\\.ru|vgtrk\\.com)/(?:iframe/(?:swf|video|live)/id|index/iframe/cast_id)/.+?)\\1"), ϒwebpage)
					if λ.IsTrue(ϒmobj) {
						return λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("url"))
					}
					ϒmobj = λ.Cal(Ωre.ϒsearch, λ.NewStr("<meta[^>]+?property=([\"\\'])og:video\\1[^>]+?content=([\"\\'])(?P<url>https?://(?:test)?player\\.(?:rutv\\.ru|vgtrk\\.com)/flash\\d+v/container\\.swf\\?id=.+?\\2)"), ϒwebpage)
					if λ.IsTrue(ϒmobj) {
						return λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("url"))
					}
					return λ.None
				})
			RUTVIE__extract_url = λ.Cal(λ.ClassMethodType, RUTVIE__extract_url)
			RUTVIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdescription        λ.Object
						ϒduration           λ.Object
						ϒfmt                λ.Object
						ϒformats            λ.Object
						ϒheight             λ.Object
						ϒis_live            λ.Object
						ϒjson_data          λ.Object
						ϒlinks              λ.Object
						ϒmedia              λ.Object
						ϒmedialist          λ.Object
						ϒmobj               λ.Object
						ϒplaylist           λ.Object
						ϒpreference         λ.Object
						ϒpriority_transport λ.Object
						ϒquality            λ.Object
						ϒself               = λargs[0]
						ϒthumbnail          λ.Object
						ϒtitle              λ.Object
						ϒtransport          λ.Object
						ϒurl                = λargs[1]
						ϒvideo_id           λ.Object
						ϒvideo_path         λ.Object
						ϒvideo_type         λ.Object
						ϒview_count         λ.Object
						ϒwidth              λ.Object
						τmp0                λ.Object
						τmp1                λ.Object
						τmp2                λ.Object
						τmp3                λ.Object
						τmp4                λ.Object
					)
					ϒmobj = λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("id"))
					ϒvideo_path = λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("path"))
					if λ.IsTrue(λ.Cal(Ωre.ϒmatch, λ.NewStr("flash\\d+v"), ϒvideo_path)) {
						ϒvideo_type = λ.NewStr("video")
					} else {
						if λ.IsTrue(λ.Cal(λ.GetAttr(ϒvideo_path, "startswith", nil), λ.NewStr("iframe"))) {
							ϒvideo_type = λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("type"))
							if λ.IsTrue(λ.Eq(ϒvideo_type, λ.NewStr("swf"))) {
								ϒvideo_type = λ.NewStr("video")
							}
						} else {
							if λ.IsTrue(λ.Cal(λ.GetAttr(ϒvideo_path, "startswith", nil), λ.NewStr("index/iframe/cast_id"))) {
								ϒvideo_type = λ.NewStr("live")
							}
						}
					}
					ϒis_live = λ.Eq(ϒvideo_type, λ.NewStr("live"))
					ϒjson_data = λ.Cal(λ.GetAttr(ϒself, "_download_json", nil), λ.Mod(λ.NewStr("http://player.rutv.ru/iframe/data%s/id/%s"), λ.NewTuple(
						func() λ.Object {
							if λ.IsTrue(ϒis_live) {
								return λ.NewStr("live")
							} else {
								return λ.NewStr("video")
							}
						}(),
						ϒvideo_id,
					)), ϒvideo_id, λ.NewStr("Downloading JSON"))
					if λ.IsTrue(λ.GetItem(ϒjson_data, λ.NewStr("errors"))) {
						panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.Mod(λ.NewStr("%s said: %s"), λ.NewTuple(
							λ.GetAttr(ϒself, "IE_NAME", nil),
							λ.GetItem(ϒjson_data, λ.NewStr("errors")),
						))), λ.KWArgs{
							{Name: "expected", Value: λ.True},
						})))
					}
					ϒplaylist = λ.GetItem(λ.GetItem(ϒjson_data, λ.NewStr("data")), λ.NewStr("playlist"))
					ϒmedialist = λ.GetItem(ϒplaylist, λ.NewStr("medialist"))
					ϒmedia = λ.GetItem(ϒmedialist, λ.NewInt(0))
					if λ.IsTrue(λ.GetItem(ϒmedia, λ.NewStr("errors"))) {
						panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.Mod(λ.NewStr("%s said: %s"), λ.NewTuple(
							λ.GetAttr(ϒself, "IE_NAME", nil),
							λ.GetItem(ϒmedia, λ.NewStr("errors")),
						))), λ.KWArgs{
							{Name: "expected", Value: λ.True},
						})))
					}
					ϒview_count = λ.Cal(λ.GetAttr(ϒplaylist, "get", nil), λ.NewStr("count_views"))
					ϒpriority_transport = λ.GetItem(ϒplaylist, λ.NewStr("priority_transport"))
					ϒthumbnail = λ.GetItem(ϒmedia, λ.NewStr("picture"))
					ϒwidth = λ.Cal(ϒint_or_none, λ.GetItem(ϒmedia, λ.NewStr("width")))
					ϒheight = λ.Cal(ϒint_or_none, λ.GetItem(ϒmedia, λ.NewStr("height")))
					ϒdescription = λ.GetItem(ϒmedia, λ.NewStr("anons"))
					ϒtitle = λ.GetItem(ϒmedia, λ.NewStr("title"))
					ϒduration = λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒmedia, "get", nil), λ.NewStr("duration")))
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(λ.GetItem(ϒmedia, λ.NewStr("sources")), "items", nil)))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						τmp2 = τmp1
						ϒtransport = λ.GetItem(τmp2, λ.NewInt(0))
						ϒlinks = λ.GetItem(τmp2, λ.NewInt(1))
						τmp2 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(ϒlinks, "items", nil)))
						for {
							if τmp3 = λ.NextDefault(τmp2, λ.AfterLast); τmp3 == λ.AfterLast {
								break
							}
							τmp4 = τmp3
							ϒquality = λ.GetItem(τmp4, λ.NewInt(0))
							ϒurl = λ.GetItem(τmp4, λ.NewInt(1))
							ϒpreference = func() λ.Object {
								if λ.IsTrue(λ.Eq(ϒpriority_transport, ϒtransport)) {
									return λ.Neg(λ.NewInt(1))
								} else {
									return λ.Neg(λ.NewInt(2))
								}
							}()
							if λ.IsTrue(λ.Eq(ϒtransport, λ.NewStr("rtmp"))) {
								ϒmobj = λ.Cal(Ωre.ϒsearch, λ.NewStr("^(?P<url>rtmp://[^/]+/(?P<app>.+))/(?P<playpath>.+)$"), ϒurl)
								if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒmobj))) {
									continue
								}
								ϒfmt = λ.NewDictWithTable(map[λ.Object]λ.Object{
									λ.NewStr("url"):        λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("url")),
									λ.NewStr("play_path"):  λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("playpath")),
									λ.NewStr("app"):        λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("app")),
									λ.NewStr("page_url"):   λ.NewStr("http://player.rutv.ru"),
									λ.NewStr("player_url"): λ.NewStr("http://player.rutv.ru/flash3v/osmf.swf?i=22"),
									λ.NewStr("rtmp_live"):  λ.True,
									λ.NewStr("ext"):        λ.NewStr("flv"),
									λ.NewStr("vbr"):        λ.Cal(λ.IntType, ϒquality),
									λ.NewStr("preference"): ϒpreference,
								})
							} else {
								if λ.IsTrue(λ.Eq(ϒtransport, λ.NewStr("m3u8"))) {
									λ.Cal(λ.GetAttr(ϒformats, "extend", nil), λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
										ϒurl,
										ϒvideo_id,
										λ.NewStr("mp4"),
									), λ.KWArgs{
										{Name: "preference", Value: ϒpreference},
										{Name: "m3u8_id", Value: λ.NewStr("hls")},
									}))
									continue
								} else {
									ϒfmt = λ.NewDictWithTable(map[λ.Object]λ.Object{
										λ.NewStr("url"): ϒurl,
									})
								}
							}
							λ.Cal(λ.GetAttr(ϒfmt, "update", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
								λ.NewStr("width"):  ϒwidth,
								λ.NewStr("height"): ϒheight,
								λ.NewStr("format_id"): λ.Mod(λ.NewStr("%s-%s"), λ.NewTuple(
									ϒtransport,
									ϒquality,
								)),
							}))
							λ.Cal(λ.GetAttr(ϒformats, "append", nil), ϒfmt)
						}
					}
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"): ϒvideo_id,
						λ.NewStr("title"): func() λ.Object {
							if λ.IsTrue(ϒis_live) {
								return λ.Cal(λ.GetAttr(ϒself, "_live_title", nil), ϒtitle)
							} else {
								return ϒtitle
							}
						}(),
						λ.NewStr("description"): ϒdescription,
						λ.NewStr("thumbnail"):   ϒthumbnail,
						λ.NewStr("view_count"):  ϒview_count,
						λ.NewStr("duration"):    ϒduration,
						λ.NewStr("formats"):     ϒformats,
						λ.NewStr("is_live"):     ϒis_live,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TESTS"):        RUTVIE__TESTS,
				λ.NewStr("_VALID_URL"):    RUTVIE__VALID_URL,
				λ.NewStr("_extract_url"):  RUTVIE__extract_url,
				λ.NewStr("_real_extract"): RUTVIE__real_extract,
			})
		}())
	})
}
