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
 * stanfordoc/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/stanfordoc.py
 */

package stanfordoc

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError          λ.Object
	InfoExtractor           λ.Object
	StanfordOpenClassroomIE λ.Object
	ϒorderedSet             λ.Object
	ϒunescapeHTML           λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ExtractorError = Ωutils.ExtractorError
		ϒorderedSet = Ωutils.ϒorderedSet
		ϒunescapeHTML = Ωutils.ϒunescapeHTML
		StanfordOpenClassroomIE = λ.Cal(λ.TypeType, λ.NewStr("StanfordOpenClassroomIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				StanfordOpenClassroomIE_IE_NAME       λ.Object
				StanfordOpenClassroomIE__TEST         λ.Object
				StanfordOpenClassroomIE__VALID_URL    λ.Object
				StanfordOpenClassroomIE__real_extract λ.Object
			)
			StanfordOpenClassroomIE_IE_NAME = λ.NewStr("stanfordoc")
			StanfordOpenClassroomIE__VALID_URL = λ.NewStr("https?://openclassroom\\.stanford\\.edu(?P<path>/?|(/MainFolder/(?:HomePage|CoursePage|VideoPage)\\.php([?]course=(?P<course>[^&]+)(&video=(?P<video>[^&]+))?(&.*)?)?))$")
			StanfordOpenClassroomIE__TEST = λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("url"): λ.NewStr("http://openclassroom.stanford.edu/MainFolder/VideoPage.php?course=PracticalUnix&video=intro-environment&speed=100"),
				λ.NewStr("md5"): λ.NewStr("544a9468546059d4e80d76265b0443b8"),
				λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("id"):    λ.NewStr("PracticalUnix_intro-environment"),
					λ.NewStr("ext"):   λ.NewStr("mp4"),
					λ.NewStr("title"): λ.NewStr("Intro Environment"),
				}),
			})
			StanfordOpenClassroomIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒbaseUrl    λ.Object
						ϒcourse     λ.Object
						ϒcoursepage λ.Object
						ϒinfo       λ.Object
						ϒlinks      λ.Object
						ϒmdoc       λ.Object
						ϒmobj       λ.Object
						ϒrootURL    λ.Object
						ϒrootpage   λ.Object
						ϒself       = λargs[0]
						ϒurl        = λargs[1]
						ϒvideo      λ.Object
						ϒxmlUrl     λ.Object
						τmp0        λ.Object
						τmp1        λ.Object
					)
					_ = τmp0
					_ = τmp1
					ϒmobj = λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl)
					if λ.IsTrue(func() λ.Object {
						if λv := λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("course")); !λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("video"))
						}
					}()) {
						ϒcourse = λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("course"))
						ϒvideo = λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("video"))
						ϒinfo = λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("id"):          λ.Add(λ.Add(ϒcourse, λ.NewStr("_")), ϒvideo),
							λ.NewStr("uploader"):    λ.None,
							λ.NewStr("upload_date"): λ.None,
						})
						ϒbaseUrl = λ.Add(λ.Add(λ.NewStr("http://openclassroom.stanford.edu/MainFolder/courses/"), ϒcourse), λ.NewStr("/videos/"))
						ϒxmlUrl = λ.Add(λ.Add(ϒbaseUrl, ϒvideo), λ.NewStr(".xml"))
						ϒmdoc = λ.Cal(λ.GetAttr(ϒself, "_download_xml", nil), ϒxmlUrl, λ.GetItem(ϒinfo, λ.NewStr("id")))
						τmp0, τmp1 = func() (λexit λ.Object, λret λ.Object) {
							defer λ.CatchMulti(
								nil,
								&λ.Catcher{λ.IndexErrorType, func(λex λ.BaseException) {
									panic(λ.Raise(λ.Cal(ExtractorError, λ.NewStr("Invalid metadata XML file"))))
								}},
							)
							λ.SetItem(ϒinfo, λ.NewStr("title"), λ.GetAttr(λ.GetItem(λ.Cal(λ.GetAttr(ϒmdoc, "findall", nil), λ.NewStr("./title")), λ.NewInt(0)), "text", nil))
							λ.SetItem(ϒinfo, λ.NewStr("url"), λ.Add(ϒbaseUrl, λ.GetAttr(λ.GetItem(λ.Cal(λ.GetAttr(ϒmdoc, "findall", nil), λ.NewStr("./videoFile")), λ.NewInt(0)), "text", nil)))
							return λ.BlockExitNormally, nil
						}()
						return ϒinfo
					} else {
						if λ.IsTrue(λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("course"))) {
							ϒcourse = λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("course"))
							ϒinfo = λ.NewDictWithTable(map[λ.Object]λ.Object{
								λ.NewStr("id"):          ϒcourse,
								λ.NewStr("_type"):       λ.NewStr("playlist"),
								λ.NewStr("uploader"):    λ.None,
								λ.NewStr("upload_date"): λ.None,
							})
							ϒcoursepage = λ.Call(λ.GetAttr(ϒself, "_download_webpage", nil), λ.NewArgs(
								ϒurl,
								λ.GetItem(ϒinfo, λ.NewStr("id")),
							), λ.KWArgs{
								{Name: "note", Value: λ.NewStr("Downloading course info page")},
								{Name: "errnote", Value: λ.NewStr("Unable to download course info page")},
							})
							λ.SetItem(ϒinfo, λ.NewStr("title"), λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
								λ.NewStr("<h1>([^<]+)</h1>"),
								ϒcoursepage,
								λ.NewStr("title"),
							), λ.KWArgs{
								{Name: "default", Value: λ.GetItem(ϒinfo, λ.NewStr("id"))},
							}))
							λ.SetItem(ϒinfo, λ.NewStr("description"), λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
								λ.NewStr("(?s)<description>([^<]+)</description>"),
								ϒcoursepage,
								λ.NewStr("description"),
							), λ.KWArgs{
								{Name: "fatal", Value: λ.False},
							}))
							ϒlinks = λ.Cal(ϒorderedSet, λ.Cal(Ωre.ϒfindall, λ.NewStr("<a href=\"(VideoPage\\.php\\?[^\"]+)\">"), ϒcoursepage))
							λ.SetItem(ϒinfo, λ.NewStr("entries"), λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
								nil,
								0, false, false,
								func(λargs []λ.Object) λ.Object {
									return λ.NewGenerator(func(λgen λ.Generator) λ.Object {
										var (
											ϒl   λ.Object
											τmp0 λ.Object
											τmp1 λ.Object
										)
										τmp0 = λ.Cal(λ.BuiltinIter, ϒlinks)
										for {
											if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
												break
											}
											ϒl = τmp1
											λgen.Yield(λ.Cal(λ.GetAttr(ϒself, "url_result", nil), λ.Mod(λ.NewStr("http://openclassroom.stanford.edu/MainFolder/%s"), λ.Cal(ϒunescapeHTML, ϒl))))
										}
										return λ.None
									})
								}))))
							return ϒinfo
						} else {
							ϒinfo = λ.NewDictWithTable(map[λ.Object]λ.Object{
								λ.NewStr("id"):          λ.NewStr("Stanford OpenClassroom"),
								λ.NewStr("_type"):       λ.NewStr("playlist"),
								λ.NewStr("uploader"):    λ.None,
								λ.NewStr("upload_date"): λ.None,
							})
							λ.SetItem(ϒinfo, λ.NewStr("title"), λ.GetItem(ϒinfo, λ.NewStr("id")))
							ϒrootURL = λ.NewStr("http://openclassroom.stanford.edu/MainFolder/HomePage.php")
							ϒrootpage = λ.Call(λ.GetAttr(ϒself, "_download_webpage", nil), λ.NewArgs(
								ϒrootURL,
								λ.GetItem(ϒinfo, λ.NewStr("id")),
							), λ.KWArgs{
								{Name: "errnote", Value: λ.NewStr("Unable to download course info page")},
							})
							ϒlinks = λ.Cal(ϒorderedSet, λ.Cal(Ωre.ϒfindall, λ.NewStr("<a href=\"(CoursePage\\.php\\?[^\"]+)\">"), ϒrootpage))
							λ.SetItem(ϒinfo, λ.NewStr("entries"), λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
								nil,
								0, false, false,
								func(λargs []λ.Object) λ.Object {
									return λ.NewGenerator(func(λgen λ.Generator) λ.Object {
										var (
											ϒl   λ.Object
											τmp0 λ.Object
											τmp1 λ.Object
										)
										τmp0 = λ.Cal(λ.BuiltinIter, ϒlinks)
										for {
											if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
												break
											}
											ϒl = τmp1
											λgen.Yield(λ.Cal(λ.GetAttr(ϒself, "url_result", nil), λ.Mod(λ.NewStr("http://openclassroom.stanford.edu/MainFolder/%s"), λ.Cal(ϒunescapeHTML, ϒl))))
										}
										return λ.None
									})
								}))))
							return ϒinfo
						}
					}
					return λ.None
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("IE_NAME"):       StanfordOpenClassroomIE_IE_NAME,
				λ.NewStr("_TEST"):         StanfordOpenClassroomIE__TEST,
				λ.NewStr("_VALID_URL"):    StanfordOpenClassroomIE__VALID_URL,
				λ.NewStr("_real_extract"): StanfordOpenClassroomIE__real_extract,
			})
		}())
	})
}
