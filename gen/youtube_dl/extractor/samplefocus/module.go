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
 * samplefocus/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/samplefocus.py
 */

package samplefocus

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor             λ.Object
	SampleFocusIE             λ.Object
	ϒextract_attributes       λ.Object
	ϒget_element_by_attribute λ.Object
	ϒint_or_none              λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒextract_attributes = Ωutils.ϒextract_attributes
		ϒget_element_by_attribute = Ωutils.ϒget_element_by_attribute
		ϒint_or_none = Ωutils.ϒint_or_none
		SampleFocusIE = λ.Cal(λ.TypeType, λ.StrLiteral("SampleFocusIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				SampleFocusIE__VALID_URL    λ.Object
				SampleFocusIE__real_extract λ.Object
			)
			SampleFocusIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?samplefocus\\.com/samples/(?P<id>[^/?&#]+)")
			SampleFocusIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒauthor        λ.Object
						ϒauthor_id     λ.Object
						ϒbody          λ.Object
						ϒbreadcrumb    λ.Object
						ϒcategories    λ.Object
						ϒcomments      λ.Object
						ϒdisplay_id    λ.Object
						ϒextract_count λ.Object
						ϒmobj          λ.Object
						ϒmp3_url       λ.Object
						ϒname          λ.Object
						ϒsample_id     λ.Object
						ϒself          = λargs[0]
						ϒthumbnail     λ.Object
						ϒtitle         λ.Object
						ϒuploader      λ.Object
						ϒuploader_id   λ.Object
						ϒurl           = λargs[1]
						ϒwebpage       λ.Object
						τmp0           λ.Object
						τmp1           λ.Object
						τmp2           λ.Object
					)
					ϒdisplay_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒwebpage = λ.Calm(ϒself, "_download_webpage", ϒurl, ϒdisplay_id)
					ϒsample_id = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("<input[^>]+id=([\"\\'])sample_id\\1[^>]+value=(?:[\"\\'])(?P<id>\\d+)"),
						ϒwebpage,
						λ.StrLiteral("sample id"),
					), λ.KWArgs{
						{Name: "group", Value: λ.StrLiteral("id")},
					})
					ϒtitle = func() λ.Object {
						if λv := λ.Call(λ.GetAttr(ϒself, "_og_search_title", nil), λ.NewArgs(ϒwebpage), λ.KWArgs{
							{Name: "fatal", Value: λ.False},
						}); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Calm(ϒself, "_html_search_regex", λ.StrLiteral("<h1>(.+?)</h1>"), ϒwebpage, λ.StrLiteral("title"))
						}
					}()
					ϒmp3_url = func() λ.Object {
						if λv := λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
							λ.StrLiteral("<input[^>]+id=([\"\\'])sample_mp3\\1[^>]+value=([\"\\'])(?P<url>(?:(?!\\2).)+)"),
							ϒwebpage,
							λ.StrLiteral("mp3"),
						), λ.KWArgs{
							{Name: "fatal", Value: λ.False},
							{Name: "group", Value: λ.StrLiteral("url")},
						}); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.GetItem(λ.Cal(ϒextract_attributes, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
								λ.StrLiteral("<meta[^>]+itemprop=([\"\\'])contentUrl\\1[^>]*>"),
								ϒwebpage,
								λ.StrLiteral("mp3 url"),
							), λ.KWArgs{
								{Name: "group", Value: λ.IntLiteral(0)},
							})), λ.StrLiteral("content"))
						}
					}()
					ϒthumbnail = func() λ.Object {
						if λv := λ.Calm(ϒself, "_og_search_thumbnail", ϒwebpage); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
								λ.StrLiteral("<img[^>]+class=(?:[\"\\'])waveform responsive-img[^>]+src=([\"\\'])(?P<url>(?:(?!\\1).)+)"),
								ϒwebpage,
								λ.StrLiteral("mp3"),
							), λ.KWArgs{
								{Name: "fatal", Value: λ.False},
								{Name: "group", Value: λ.StrLiteral("url")},
							})
						}
					}()
					ϒcomments = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(Ωre.ϒfindall, λ.StrLiteral("(?s)<p[^>]+class=\"comment-author\"><a[^>]+href=\"/users/([^\"]+)\">([^\"]+)</a>.+?<p[^>]+class=\"comment-body\">([^>]+)</p>"), ϒwebpage))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						τmp2 = λ.UnpackIterable(τmp1, 3)
						ϒauthor_id = λ.GetItem(τmp2, λ.IntLiteral(0))
						ϒauthor = λ.GetItem(τmp2, λ.IntLiteral(1))
						ϒbody = λ.GetItem(τmp2, λ.IntLiteral(2))
						λ.Calm(ϒcomments, "append", λ.DictLiteral(map[string]λ.Object{
							"author":    ϒauthor,
							"author_id": ϒauthor_id,
							"text":      ϒbody,
						}))
					}
					τmp0 = λ.None
					ϒuploader_id = τmp0
					ϒuploader = τmp0
					ϒmobj = λ.Cal(Ωre.ϒsearch, λ.StrLiteral(">By <a[^>]+href=\"/users/([^\"]+)\"[^>]*>([^<]+)"), ϒwebpage)
					if λ.IsTrue(ϒmobj) {
						τmp0 = λ.UnpackIterable(λ.Calm(ϒmobj, "groups"), 2)
						ϒuploader_id = λ.GetItem(τmp0, λ.IntLiteral(0))
						ϒuploader = λ.GetItem(τmp0, λ.IntLiteral(1))
					}
					ϒbreadcrumb = λ.Cal(ϒget_element_by_attribute, λ.StrLiteral("typeof"), λ.StrLiteral("BreadcrumbList"), ϒwebpage)
					ϒcategories = λ.NewList()
					if λ.IsTrue(ϒbreadcrumb) {
						τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(Ωre.ϒfindall, λ.StrLiteral("<span[^>]+property=([\"\\'])name\\1[^>]*>([^<]+)"), ϒbreadcrumb))
						for {
							if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
								break
							}
							τmp2 = λ.UnpackIterable(τmp1, 2)
							_ = λ.GetItem(τmp2, λ.IntLiteral(0))
							ϒname = λ.GetItem(τmp2, λ.IntLiteral(1))
							λ.Calm(ϒcategories, "append", ϒname)
						}
					}
					ϒextract_count = λ.NewFunction("extract_count",
						[]λ.Param{
							{Name: "klass"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒklass = λargs[0]
							)
							return λ.Cal(ϒint_or_none, λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
								λ.Mod(λ.StrLiteral("<span[^>]+class=(?:[\"\\'])?%s-count[^>]*>(\\d+)"), ϒklass),
								ϒwebpage,
								ϒklass,
							), λ.KWArgs{
								{Name: "fatal", Value: λ.False},
							}))
						})
					return λ.DictLiteral(map[string]λ.Object{
						"id":         ϒsample_id,
						"title":      ϒtitle,
						"url":        ϒmp3_url,
						"display_id": ϒdisplay_id,
						"thumbnail":  ϒthumbnail,
						"uploader":   ϒuploader,
						"license": λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
							λ.StrLiteral("<a[^>]+href=([\"\\'])/license\\1[^>]*>(?P<license>[^<]+)<"),
							ϒwebpage,
							λ.StrLiteral("license"),
						), λ.KWArgs{
							{Name: "fatal", Value: λ.False},
							{Name: "group", Value: λ.StrLiteral("license")},
						}),
						"uploader_id":   ϒuploader_id,
						"like_count":    λ.Cal(ϒextract_count, λ.Mod(λ.StrLiteral("sample-%s-favorites"), ϒsample_id)),
						"comment_count": λ.Cal(ϒextract_count, λ.StrLiteral("comments")),
						"comments":      ϒcomments,
						"categories":    ϒcategories,
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL":    SampleFocusIE__VALID_URL,
				"_real_extract": SampleFocusIE__real_extract,
			})
		}())
	})
}
