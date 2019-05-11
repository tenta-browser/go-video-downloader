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
 * instagram/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/instagram.py
 */

package instagram

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError            λ.Object
	InfoExtractor             λ.Object
	InstagramIE               λ.Object
	InstagramPlaylistIE       λ.Object
	InstagramTagIE            λ.Object
	InstagramUserIE           λ.Object
	ϒcompat_HTTPError         λ.Object
	ϒcompat_str               λ.Object
	ϒget_element_by_attribute λ.Object
	ϒint_or_none              λ.Object
	ϒstd_headers              λ.Object
	ϒtry_get                  λ.Object
	ϒurl_or_none              λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_str = Ωcompat.ϒcompat_str
		ϒcompat_HTTPError = Ωcompat.ϒcompat_HTTPError
		ExtractorError = Ωutils.ExtractorError
		ϒget_element_by_attribute = Ωutils.ϒget_element_by_attribute
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒstd_headers = Ωutils.ϒstd_headers
		ϒtry_get = Ωutils.ϒtry_get
		ϒurl_or_none = Ωutils.ϒurl_or_none
		InstagramIE = λ.Cal(λ.TypeType, λ.NewStr("InstagramIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				InstagramIE__TESTS        λ.Object
				InstagramIE__VALID_URL    λ.Object
				InstagramIE__real_extract λ.Object
			)
			InstagramIE__VALID_URL = λ.NewStr("(?P<url>https?://(?:www\\.)?instagram\\.com/p/(?P<id>[^/?#&]+))")
			InstagramIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("https://instagram.com/p/aye83DjauH/?foo=bar#abc"),
					λ.NewStr("md5"): λ.NewStr("0d2da106a9d2631273e192b372806516"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):            λ.NewStr("aye83DjauH"),
						λ.NewStr("ext"):           λ.NewStr("mp4"),
						λ.NewStr("title"):         λ.NewStr("Video by naomipq"),
						λ.NewStr("description"):   λ.NewStr("md5:1f17f0ab29bd6fe2bfad705f58de3cb8"),
						λ.NewStr("thumbnail"):     λ.NewStr("re:^https?://.*\\.jpg"),
						λ.NewStr("timestamp"):     λ.NewInt(1371748545),
						λ.NewStr("upload_date"):   λ.NewStr("20130620"),
						λ.NewStr("uploader_id"):   λ.NewStr("naomipq"),
						λ.NewStr("uploader"):      λ.NewStr("Naomi Leonor Phan-Quang"),
						λ.NewStr("like_count"):    λ.IntType,
						λ.NewStr("comment_count"): λ.IntType,
						λ.NewStr("comments"):      λ.ListType,
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("https://www.instagram.com/p/BA-pQFBG8HZ/?taken-by=britneyspears"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):            λ.NewStr("BA-pQFBG8HZ"),
						λ.NewStr("ext"):           λ.NewStr("mp4"),
						λ.NewStr("title"):         λ.NewStr("Video by britneyspears"),
						λ.NewStr("thumbnail"):     λ.NewStr("re:^https?://.*\\.jpg"),
						λ.NewStr("timestamp"):     λ.NewInt(1453760977),
						λ.NewStr("upload_date"):   λ.NewStr("20160125"),
						λ.NewStr("uploader_id"):   λ.NewStr("britneyspears"),
						λ.NewStr("uploader"):      λ.NewStr("Britney Spears"),
						λ.NewStr("like_count"):    λ.IntType,
						λ.NewStr("comment_count"): λ.IntType,
						λ.NewStr("comments"):      λ.ListType,
					}),
					λ.NewStr("params"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("skip_download"): λ.True,
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("https://www.instagram.com/p/BQ0eAlwhDrw/"),
					λ.NewStr("playlist"): λ.NewList(
						λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
								λ.NewStr("id"):    λ.NewStr("BQ0dSaohpPW"),
								λ.NewStr("ext"):   λ.NewStr("mp4"),
								λ.NewStr("title"): λ.NewStr("Video 1"),
							}),
						}),
						λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
								λ.NewStr("id"):    λ.NewStr("BQ0dTpOhuHT"),
								λ.NewStr("ext"):   λ.NewStr("mp4"),
								λ.NewStr("title"): λ.NewStr("Video 2"),
							}),
						}),
						λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
								λ.NewStr("id"):    λ.NewStr("BQ0dT7RBFeF"),
								λ.NewStr("ext"):   λ.NewStr("mp4"),
								λ.NewStr("title"): λ.NewStr("Video 3"),
							}),
						}),
					),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("BQ0eAlwhDrw"),
						λ.NewStr("title"):       λ.NewStr("Post by instagram"),
						λ.NewStr("description"): λ.NewStr("md5:0f9203fc6a2ce4d228da5754bcf54957"),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://instagram.com/p/-Cmh1cukG2/"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://instagram.com/p/9o6LshA7zy/embed/"),
					λ.NewStr("only_matching"): λ.True,
				}),
			)
			InstagramIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒcomment_count  λ.Object
						ϒcomments       λ.Object
						ϒdescription    λ.Object
						ϒedge           λ.Object
						ϒedge_num       λ.Object
						ϒedges          λ.Object
						ϒentries        λ.Object
						ϒformats        λ.Object
						ϒget_count      λ.Object
						ϒheight         λ.Object
						ϒlike_count     λ.Object
						ϒmedia          λ.Object
						ϒmobj           λ.Object
						ϒnode           λ.Object
						ϒnode_video_url λ.Object
						ϒself           = λargs[0]
						ϒshared_data    λ.Object
						ϒthumbnail      λ.Object
						ϒtimestamp      λ.Object
						ϒuploader       λ.Object
						ϒuploader_id    λ.Object
						ϒurl            = λargs[1]
						ϒvideo_id       λ.Object
						ϒvideo_url      λ.Object
						ϒwebpage        λ.Object
						ϒwidth          λ.Object
						τmp0            λ.Object
						τmp1            λ.Object
						τmp2            λ.Object
					)
					ϒmobj = λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("id"))
					ϒurl = λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("url"))
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒvideo_id)
					τmp0 = λ.Mul(λ.NewList(λ.None), λ.NewInt(11))
					ϒvideo_url = λ.GetItem(τmp0, λ.NewInt(0))
					ϒdescription = λ.GetItem(τmp0, λ.NewInt(1))
					ϒthumbnail = λ.GetItem(τmp0, λ.NewInt(2))
					ϒtimestamp = λ.GetItem(τmp0, λ.NewInt(3))
					ϒuploader = λ.GetItem(τmp0, λ.NewInt(4))
					ϒuploader_id = λ.GetItem(τmp0, λ.NewInt(5))
					ϒlike_count = λ.GetItem(τmp0, λ.NewInt(6))
					ϒcomment_count = λ.GetItem(τmp0, λ.NewInt(7))
					ϒcomments = λ.GetItem(τmp0, λ.NewInt(8))
					ϒheight = λ.GetItem(τmp0, λ.NewInt(9))
					ϒwidth = λ.GetItem(τmp0, λ.NewInt(10))
					ϒshared_data = λ.Call(λ.GetAttr(ϒself, "_parse_json", nil), λ.NewArgs(
						λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
							λ.NewStr("window\\._sharedData\\s*=\\s*({.+?});"),
							ϒwebpage,
							λ.NewStr("shared data"),
						), λ.KWArgs{
							{Name: "default", Value: λ.NewStr("{}")},
						}),
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					})
					if λ.IsTrue(ϒshared_data) {
						ϒmedia = λ.Cal(ϒtry_get, ϒshared_data, λ.NewTuple(
							λ.NewFunction("<lambda>",
								[]λ.Param{
									{Name: "x"},
								},
								0, false, false,
								func(λargs []λ.Object) λ.Object {
									var (
										ϒx = λargs[0]
									)
									return λ.GetItem(λ.GetItem(λ.GetItem(λ.GetItem(λ.GetItem(ϒx, λ.NewStr("entry_data")), λ.NewStr("PostPage")), λ.NewInt(0)), λ.NewStr("graphql")), λ.NewStr("shortcode_media"))
								}),
							λ.NewFunction("<lambda>",
								[]λ.Param{
									{Name: "x"},
								},
								0, false, false,
								func(λargs []λ.Object) λ.Object {
									var (
										ϒx = λargs[0]
									)
									return λ.GetItem(λ.GetItem(λ.GetItem(λ.GetItem(ϒx, λ.NewStr("entry_data")), λ.NewStr("PostPage")), λ.NewInt(0)), λ.NewStr("media"))
								}),
						), λ.DictType)
						if λ.IsTrue(ϒmedia) {
							ϒvideo_url = λ.Cal(λ.GetAttr(ϒmedia, "get", nil), λ.NewStr("video_url"))
							ϒheight = λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒmedia, "get", nil), λ.NewStr("dimensions"), λ.NewDictWithTable(map[λ.Object]λ.Object{})), "get", nil), λ.NewStr("height")))
							ϒwidth = λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒmedia, "get", nil), λ.NewStr("dimensions"), λ.NewDictWithTable(map[λ.Object]λ.Object{})), "get", nil), λ.NewStr("width")))
							ϒdescription = func() λ.Object {
								if λv := λ.Cal(ϒtry_get, ϒmedia, λ.NewFunction("<lambda>",
									[]λ.Param{
										{Name: "x"},
									},
									0, false, false,
									func(λargs []λ.Object) λ.Object {
										var (
											ϒx = λargs[0]
										)
										return λ.GetItem(λ.GetItem(λ.GetItem(λ.GetItem(λ.GetItem(ϒx, λ.NewStr("edge_media_to_caption")), λ.NewStr("edges")), λ.NewInt(0)), λ.NewStr("node")), λ.NewStr("text"))
									}), ϒcompat_str); λ.IsTrue(λv) {
									return λv
								} else {
									return λ.Cal(λ.GetAttr(ϒmedia, "get", nil), λ.NewStr("caption"))
								}
							}()
							ϒthumbnail = λ.Cal(λ.GetAttr(ϒmedia, "get", nil), λ.NewStr("display_src"))
							ϒtimestamp = λ.Cal(ϒint_or_none, func() λ.Object {
								if λv := λ.Cal(λ.GetAttr(ϒmedia, "get", nil), λ.NewStr("taken_at_timestamp")); λ.IsTrue(λv) {
									return λv
								} else {
									return λ.Cal(λ.GetAttr(ϒmedia, "get", nil), λ.NewStr("date"))
								}
							}())
							ϒuploader = λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒmedia, "get", nil), λ.NewStr("owner"), λ.NewDictWithTable(map[λ.Object]λ.Object{})), "get", nil), λ.NewStr("full_name"))
							ϒuploader_id = λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒmedia, "get", nil), λ.NewStr("owner"), λ.NewDictWithTable(map[λ.Object]λ.Object{})), "get", nil), λ.NewStr("username"))
							ϒget_count = λ.NewFunction("get_count",
								[]λ.Param{
									{Name: "key"},
									{Name: "kind"},
								},
								0, false, false,
								func(λargs []λ.Object) λ.Object {
									var (
										ϒkey  = λargs[0]
										ϒkind = λargs[1]
									)
									return λ.Cal(ϒint_or_none, λ.Cal(ϒtry_get, ϒmedia, λ.NewTuple(
										λ.NewFunction("<lambda>",
											[]λ.Param{
												{Name: "x"},
											},
											0, false, false,
											func(λargs []λ.Object) λ.Object {
												var (
													ϒx = λargs[0]
												)
												return λ.GetItem(λ.GetItem(ϒx, λ.Mod(λ.NewStr("edge_media_%s"), ϒkey)), λ.NewStr("count"))
											}),
										λ.NewFunction("<lambda>",
											[]λ.Param{
												{Name: "x"},
											},
											0, false, false,
											func(λargs []λ.Object) λ.Object {
												var (
													ϒx = λargs[0]
												)
												return λ.GetItem(λ.GetItem(ϒx, λ.Mod(λ.NewStr("%ss"), ϒkind)), λ.NewStr("count"))
											}),
									)))
								})
							ϒlike_count = λ.Cal(ϒget_count, λ.NewStr("preview_like"), λ.NewStr("like"))
							ϒcomment_count = λ.Cal(ϒget_count, λ.NewStr("to_comment"), λ.NewStr("comment"))
							ϒcomments = λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
								nil,
								0, false, false,
								func(λargs []λ.Object) λ.Object {
									return λ.NewGenerator(func(λgy λ.Yielder) λ.Object {
										var (
											ϒcomment λ.Object
											τmp0     λ.Object
											τmp1     λ.Object
										)
										τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒmedia, "get", nil), λ.NewStr("comments"), λ.NewDictWithTable(map[λ.Object]λ.Object{})), "get", nil), λ.NewStr("nodes"), λ.NewList()))
										for {
											if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
												break
											}
											ϒcomment = τmp1
											if λ.IsTrue(λ.Cal(λ.GetAttr(ϒcomment, "get", nil), λ.NewStr("text"))) {
												λgy.Yield(λ.NewDictWithTable(map[λ.Object]λ.Object{
													λ.NewStr("author"):    λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒcomment, "get", nil), λ.NewStr("user"), λ.NewDictWithTable(map[λ.Object]λ.Object{})), "get", nil), λ.NewStr("username")),
													λ.NewStr("author_id"): λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒcomment, "get", nil), λ.NewStr("user"), λ.NewDictWithTable(map[λ.Object]λ.Object{})), "get", nil), λ.NewStr("id")),
													λ.NewStr("id"):        λ.Cal(λ.GetAttr(ϒcomment, "get", nil), λ.NewStr("id")),
													λ.NewStr("text"):      λ.Cal(λ.GetAttr(ϒcomment, "get", nil), λ.NewStr("text")),
													λ.NewStr("timestamp"): λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒcomment, "get", nil), λ.NewStr("created_at"))),
												}))
											}
										}
										return λ.None
									})
								})))
							if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒvideo_url))) {
								ϒedges = func() λ.Object {
									if λv := λ.Cal(ϒtry_get, ϒmedia, λ.NewFunction("<lambda>",
										[]λ.Param{
											{Name: "x"},
										},
										0, false, false,
										func(λargs []λ.Object) λ.Object {
											var (
												ϒx = λargs[0]
											)
											return λ.GetItem(λ.GetItem(ϒx, λ.NewStr("edge_sidecar_to_children")), λ.NewStr("edges"))
										}), λ.ListType); λ.IsTrue(λv) {
										return λv
									} else {
										return λ.NewList()
									}
								}()
								if λ.IsTrue(ϒedges) {
									ϒentries = λ.NewList()
									τmp0 = λ.Cal(λ.BuiltinIter, λ.Call(λ.EnumerateIteratorType, λ.NewArgs(ϒedges), λ.KWArgs{
										{Name: "start", Value: λ.NewInt(1)},
									}))
									for {
										if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
											break
										}
										τmp2 = τmp1
										ϒedge_num = λ.GetItem(τmp2, λ.NewInt(0))
										ϒedge = λ.GetItem(τmp2, λ.NewInt(1))
										ϒnode = λ.Cal(ϒtry_get, ϒedge, λ.NewFunction("<lambda>",
											[]λ.Param{
												{Name: "x"},
											},
											0, false, false,
											func(λargs []λ.Object) λ.Object {
												var (
													ϒx = λargs[0]
												)
												return λ.GetItem(ϒx, λ.NewStr("node"))
											}), λ.DictType)
										if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒnode))) {
											continue
										}
										ϒnode_video_url = λ.Cal(ϒurl_or_none, λ.Cal(λ.GetAttr(ϒnode, "get", nil), λ.NewStr("video_url")))
										if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒnode_video_url))) {
											continue
										}
										λ.Cal(λ.GetAttr(ϒentries, "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
											λ.NewStr("id"): func() λ.Object {
												if λv := λ.Cal(λ.GetAttr(ϒnode, "get", nil), λ.NewStr("shortcode")); λ.IsTrue(λv) {
													return λv
												} else {
													return λ.GetItem(ϒnode, λ.NewStr("id"))
												}
											}(),
											λ.NewStr("title"):     λ.Mod(λ.NewStr("Video %d"), ϒedge_num),
											λ.NewStr("url"):       ϒnode_video_url,
											λ.NewStr("thumbnail"): λ.Cal(λ.GetAttr(ϒnode, "get", nil), λ.NewStr("display_url")),
											λ.NewStr("width"): λ.Cal(ϒint_or_none, λ.Cal(ϒtry_get, ϒnode, λ.NewFunction("<lambda>",
												[]λ.Param{
													{Name: "x"},
												},
												0, false, false,
												func(λargs []λ.Object) λ.Object {
													var (
														ϒx = λargs[0]
													)
													return λ.GetItem(λ.GetItem(ϒx, λ.NewStr("dimensions")), λ.NewStr("width"))
												}))),
											λ.NewStr("height"): λ.Cal(ϒint_or_none, λ.Cal(ϒtry_get, ϒnode, λ.NewFunction("<lambda>",
												[]λ.Param{
													{Name: "x"},
												},
												0, false, false,
												func(λargs []λ.Object) λ.Object {
													var (
														ϒx = λargs[0]
													)
													return λ.GetItem(λ.GetItem(ϒx, λ.NewStr("dimensions")), λ.NewStr("height"))
												}))),
											λ.NewStr("view_count"): λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒnode, "get", nil), λ.NewStr("video_view_count"))),
										}))
									}
									return λ.Cal(λ.GetAttr(ϒself, "playlist_result", nil), ϒentries, ϒvideo_id, func() λ.Object {
										if λ.IsTrue(ϒuploader_id) {
											return λ.Mod(λ.NewStr("Post by %s"), ϒuploader_id)
										} else {
											return λ.None
										}
									}(), ϒdescription)
								}
							}
						}
					}
					if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒvideo_url))) {
						ϒvideo_url = λ.Call(λ.GetAttr(ϒself, "_og_search_video_url", nil), λ.NewArgs(ϒwebpage), λ.KWArgs{
							{Name: "secure", Value: λ.False},
						})
					}
					ϒformats = λ.NewList(λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("url"):    ϒvideo_url,
						λ.NewStr("width"):  ϒwidth,
						λ.NewStr("height"): ϒheight,
					}))
					if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒuploader_id))) {
						ϒuploader_id = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
							λ.NewStr("\"owner\"\\s*:\\s*{\\s*\"username\"\\s*:\\s*\"(.+?)\""),
							ϒwebpage,
							λ.NewStr("uploader id"),
						), λ.KWArgs{
							{Name: "fatal", Value: λ.False},
						})
					}
					if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒdescription))) {
						ϒdescription = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
							λ.NewStr("\"caption\"\\s*:\\s*\"(.+?)\""),
							ϒwebpage,
							λ.NewStr("description"),
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
						})
						if λ.IsTrue(λ.NewBool(ϒdescription != λ.None)) {
							ϒdescription = λ.Cal(λ.None, ϒdescription)
						}
					}
					if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒthumbnail))) {
						ϒthumbnail = λ.Cal(λ.GetAttr(ϒself, "_og_search_thumbnail", nil), ϒwebpage)
					}
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):            ϒvideo_id,
						λ.NewStr("formats"):       ϒformats,
						λ.NewStr("ext"):           λ.NewStr("mp4"),
						λ.NewStr("title"):         λ.Mod(λ.NewStr("Video by %s"), ϒuploader_id),
						λ.NewStr("description"):   ϒdescription,
						λ.NewStr("thumbnail"):     ϒthumbnail,
						λ.NewStr("timestamp"):     ϒtimestamp,
						λ.NewStr("uploader_id"):   ϒuploader_id,
						λ.NewStr("uploader"):      ϒuploader,
						λ.NewStr("like_count"):    ϒlike_count,
						λ.NewStr("comment_count"): ϒcomment_count,
						λ.NewStr("comments"):      ϒcomments,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TESTS"):        InstagramIE__TESTS,
				λ.NewStr("_VALID_URL"):    InstagramIE__VALID_URL,
				λ.NewStr("_real_extract"): InstagramIE__real_extract,
			})
		}())
		InstagramPlaylistIE = λ.Cal(λ.TypeType, λ.NewStr("InstagramPlaylistIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {

			return λ.NewDictWithTable(map[λ.Object]λ.Object{})
		}())
		InstagramUserIE = λ.Cal(λ.TypeType, λ.NewStr("InstagramUserIE"), λ.NewTuple(InstagramPlaylistIE), func() λ.Dict {
			var (
				InstagramUserIE__VALID_URL λ.Object
			)
			InstagramUserIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?instagram\\.com/(?P<id>[^/]{2,})/?(?:$|[?#])")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): InstagramUserIE__VALID_URL,
			})
		}())
		InstagramTagIE = λ.Cal(λ.TypeType, λ.NewStr("InstagramTagIE"), λ.NewTuple(InstagramPlaylistIE), func() λ.Dict {
			var (
				InstagramTagIE__VALID_URL λ.Object
			)
			InstagramTagIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?instagram\\.com/explore/tags/(?P<id>[^/]+)")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): InstagramTagIE__VALID_URL,
			})
		}())
	})
}
