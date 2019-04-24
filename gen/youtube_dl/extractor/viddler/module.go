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
 * viddler/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/viddler.py
 */

package viddler

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor  λ.Object
	ViddlerIE      λ.Object
	ϒfloat_or_none λ.Object
	ϒint_or_none   λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒfloat_or_none = Ωutils.ϒfloat_or_none
		ϒint_or_none = Ωutils.ϒint_or_none
		ViddlerIE = λ.Cal(λ.TypeType, λ.NewStr("ViddlerIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				ViddlerIE__TESTS        λ.Object
				ViddlerIE__VALID_URL    λ.Object
				ViddlerIE__real_extract λ.Object
			)
			ViddlerIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?viddler\\.com/(?:v|embed|player)/(?P<id>[a-z0-9]+)(?:.+?\\bsecret=(\\d+))?")
			ViddlerIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.viddler.com/v/43903784"),
					λ.NewStr("md5"): λ.NewStr("9eee21161d2c7f5b39690c3e325fab2f"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):            λ.NewStr("43903784"),
						λ.NewStr("ext"):           λ.NewStr("mov"),
						λ.NewStr("title"):         λ.NewStr("Video Made Easy"),
						λ.NewStr("description"):   λ.NewStr("md5:6a697ebd844ff3093bd2e82c37b409cd"),
						λ.NewStr("uploader"):      λ.NewStr("viddler"),
						λ.NewStr("timestamp"):     λ.NewInt(1335371429),
						λ.NewStr("upload_date"):   λ.NewStr("20120425"),
						λ.NewStr("duration"):      λ.NewFloat(100.89),
						λ.NewStr("thumbnail"):     λ.NewStr("re:^https?://.*\\.jpg$"),
						λ.NewStr("view_count"):    λ.IntType,
						λ.NewStr("comment_count"): λ.IntType,
						λ.NewStr("categories"): λ.NewList(
							λ.NewStr("video content"),
							λ.NewStr("high quality video"),
							λ.NewStr("video made easy"),
							λ.NewStr("how to produce video with limited resources"),
							λ.NewStr("viddler"),
						),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.viddler.com/v/4d03aad9/"),
					λ.NewStr("md5"): λ.NewStr("f12c5a7fa839c47a79363bfdf69404fb"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):            λ.NewStr("4d03aad9"),
						λ.NewStr("ext"):           λ.NewStr("ts"),
						λ.NewStr("title"):         λ.NewStr("WALL-TO-GORTAT"),
						λ.NewStr("upload_date"):   λ.NewStr("20150126"),
						λ.NewStr("uploader"):      λ.NewStr("deadspin"),
						λ.NewStr("timestamp"):     λ.NewInt(1422285291),
						λ.NewStr("view_count"):    λ.IntType,
						λ.NewStr("comment_count"): λ.IntType,
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.viddler.com/player/221ebbbd/0/"),
					λ.NewStr("md5"): λ.NewStr("740511f61d3d1bb71dc14a0fe01a1c10"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):            λ.NewStr("221ebbbd"),
						λ.NewStr("ext"):           λ.NewStr("mov"),
						λ.NewStr("title"):         λ.NewStr("LETeens-Grammar-snack-third-conditional"),
						λ.NewStr("description"):   λ.NewStr(" "),
						λ.NewStr("upload_date"):   λ.NewStr("20140929"),
						λ.NewStr("uploader"):      λ.NewStr("BCLETeens"),
						λ.NewStr("timestamp"):     λ.NewInt(1411997190),
						λ.NewStr("view_count"):    λ.IntType,
						λ.NewStr("comment_count"): λ.IntType,
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.viddler.com/v/890c0985?secret=34051570"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):            λ.NewStr("890c0985"),
						λ.NewStr("ext"):           λ.NewStr("mp4"),
						λ.NewStr("title"):         λ.NewStr("Complete Property Training - Traineeships"),
						λ.NewStr("description"):   λ.NewStr(" "),
						λ.NewStr("upload_date"):   λ.NewStr("20130606"),
						λ.NewStr("uploader"):      λ.NewStr("TiffanyBowtell"),
						λ.NewStr("timestamp"):     λ.NewInt(1370496993),
						λ.NewStr("view_count"):    λ.IntType,
						λ.NewStr("comment_count"): λ.IntType,
					}),
					λ.NewStr("params"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("skip_download"): λ.True,
					}),
				}),
			)
			ViddlerIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒcategories λ.Object
						ϒdata       λ.Object
						ϒf          λ.Object
						ϒfiled      λ.Object
						ϒformat_id  λ.Object
						ϒformats    λ.Object
						ϒquery      λ.Object
						ϒsecret     λ.Object
						ϒself       = λargs[0]
						ϒurl        = λargs[1]
						ϒvideo_id   λ.Object
						τmp0        λ.Object
						τmp1        λ.Object
					)
					τmp0 = λ.Cal(λ.GetAttr(λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl), "groups", nil))
					ϒvideo_id = λ.GetItem(τmp0, λ.NewInt(0))
					ϒsecret = λ.GetItem(τmp0, λ.NewInt(1))
					ϒquery = λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("video_id"): ϒvideo_id,
						λ.NewStr("key"):      λ.NewStr("v0vhrt7bg2xq1vyxhkct"),
					})
					if λ.IsTrue(ϒsecret) {
						λ.SetItem(ϒquery, λ.NewStr("secret"), ϒsecret)
					}
					ϒdata = λ.GetItem(λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
						λ.NewStr("http://api.viddler.com/api/v2/viddler.videos.getPlaybackDetails.json"),
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "headers", Value: λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("Referer"): ϒurl,
						})},
						{Name: "query", Value: ϒquery},
					}), λ.NewStr("video"))
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.GetItem(ϒdata, λ.NewStr("files")))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒfiled = τmp1
						if λ.IsTrue(λ.Ne(λ.Cal(λ.GetAttr(ϒfiled, "get", nil), λ.NewStr("status"), λ.NewStr("ready")), λ.NewStr("ready"))) {
							continue
						}
						ϒformat_id = func() λ.Object {
							if λv := λ.Cal(λ.GetAttr(ϒfiled, "get", nil), λ.NewStr("profile_id")); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.GetItem(ϒfiled, λ.NewStr("profile_name"))
							}
						}()
						ϒf = λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("format_id"):         ϒformat_id,
							λ.NewStr("format_note"):       λ.GetItem(ϒfiled, λ.NewStr("profile_name")),
							λ.NewStr("url"):               λ.Cal(λ.GetAttr(ϒself, "_proto_relative_url", nil), λ.GetItem(ϒfiled, λ.NewStr("url"))),
							λ.NewStr("width"):             λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒfiled, "get", nil), λ.NewStr("width"))),
							λ.NewStr("height"):            λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒfiled, "get", nil), λ.NewStr("height"))),
							λ.NewStr("filesize"):          λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒfiled, "get", nil), λ.NewStr("size"))),
							λ.NewStr("ext"):               λ.Cal(λ.GetAttr(ϒfiled, "get", nil), λ.NewStr("ext")),
							λ.NewStr("source_preference"): λ.Neg(λ.NewInt(1)),
						})
						λ.Cal(λ.GetAttr(ϒformats, "append", nil), ϒf)
						if λ.IsTrue(λ.Cal(λ.GetAttr(ϒfiled, "get", nil), λ.NewStr("cdn_url"))) {
							ϒf = λ.Cal(λ.GetAttr(ϒf, "copy", nil))
							λ.SetItem(ϒf, λ.NewStr("url"), λ.Cal(λ.GetAttr(ϒself, "_proto_relative_url", nil), λ.GetItem(ϒfiled, λ.NewStr("cdn_url")), λ.NewStr("http:")))
							λ.SetItem(ϒf, λ.NewStr("format_id"), λ.Add(ϒformat_id, λ.NewStr("-cdn")))
							λ.SetItem(ϒf, λ.NewStr("source_preference"), λ.NewInt(1))
							λ.Cal(λ.GetAttr(ϒformats, "append", nil), ϒf)
						}
						if λ.IsTrue(λ.Cal(λ.GetAttr(ϒfiled, "get", nil), λ.NewStr("html5_video_source"))) {
							ϒf = λ.Cal(λ.GetAttr(ϒf, "copy", nil))
							λ.SetItem(ϒf, λ.NewStr("url"), λ.Cal(λ.GetAttr(ϒself, "_proto_relative_url", nil), λ.GetItem(ϒfiled, λ.NewStr("html5_video_source"))))
							λ.SetItem(ϒf, λ.NewStr("format_id"), λ.Add(ϒformat_id, λ.NewStr("-html5")))
							λ.SetItem(ϒf, λ.NewStr("source_preference"), λ.NewInt(0))
							λ.Cal(λ.GetAttr(ϒformats, "append", nil), ϒf)
						}
					}
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					ϒcategories = λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
						nil,
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							return λ.NewGenerator(func(λgen λ.Generator) λ.Object {
								var (
									ϒt   λ.Object
									τmp0 λ.Object
									τmp1 λ.Object
								)
								τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(ϒdata, "get", nil), λ.NewStr("tags"), λ.NewList()))
								for {
									if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
										break
									}
									ϒt = τmp1
									if λ.IsTrue(λ.NewBool(λ.Contains(ϒt, λ.NewStr("text")))) {
										λgen.Yield(λ.Cal(λ.GetAttr(ϒt, "get", nil), λ.NewStr("text")))
									}
								}
								return λ.None
							})
						})))
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):            ϒvideo_id,
						λ.NewStr("title"):         λ.GetItem(ϒdata, λ.NewStr("title")),
						λ.NewStr("formats"):       ϒformats,
						λ.NewStr("description"):   λ.Cal(λ.GetAttr(ϒdata, "get", nil), λ.NewStr("description")),
						λ.NewStr("timestamp"):     λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒdata, "get", nil), λ.NewStr("upload_time"))),
						λ.NewStr("thumbnail"):     λ.Cal(λ.GetAttr(ϒself, "_proto_relative_url", nil), λ.Cal(λ.GetAttr(ϒdata, "get", nil), λ.NewStr("thumbnail_url"))),
						λ.NewStr("uploader"):      λ.Cal(λ.GetAttr(ϒdata, "get", nil), λ.NewStr("author")),
						λ.NewStr("duration"):      λ.Cal(ϒfloat_or_none, λ.Cal(λ.GetAttr(ϒdata, "get", nil), λ.NewStr("length"))),
						λ.NewStr("view_count"):    λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒdata, "get", nil), λ.NewStr("view_count"))),
						λ.NewStr("comment_count"): λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒdata, "get", nil), λ.NewStr("comment_count"))),
						λ.NewStr("categories"):    ϒcategories,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TESTS"):        ViddlerIE__TESTS,
				λ.NewStr("_VALID_URL"):    ViddlerIE__VALID_URL,
				λ.NewStr("_real_extract"): ViddlerIE__real_extract,
			})
		}())
	})
}