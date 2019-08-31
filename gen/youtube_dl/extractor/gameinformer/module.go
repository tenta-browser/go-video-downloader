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
 * gameinformer/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/gameinformer.py
 */

package gameinformer

import (
	Ωbrightcove "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/brightcove"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	BrightcoveNewIE       λ.Object
	GameInformerIE        λ.Object
	InfoExtractor         λ.Object
	ϒclean_html           λ.Object
	ϒget_element_by_class λ.Object
	ϒget_element_by_id    λ.Object
)

func init() {
	λ.InitModule(func() {
		BrightcoveNewIE = Ωbrightcove.BrightcoveNewIE
		InfoExtractor = Ωcommon.InfoExtractor
		ϒclean_html = Ωutils.ϒclean_html
		ϒget_element_by_class = Ωutils.ϒget_element_by_class
		ϒget_element_by_id = Ωutils.ϒget_element_by_id
		GameInformerIE = λ.Cal(λ.TypeType, λ.NewStr("GameInformerIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				GameInformerIE_BRIGHTCOVE_URL_TEMPLATE λ.Object
				GameInformerIE__TESTS                  λ.Object
				GameInformerIE__VALID_URL              λ.Object
				GameInformerIE__real_extract           λ.Object
			)
			GameInformerIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?gameinformer\\.com/(?:[^/]+/)*(?P<id>[^.?&#]+)")
			GameInformerIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.gameinformer.com/b/features/archive/2015/09/26/replay-animal-crossing.aspx"),
					λ.NewStr("md5"): λ.NewStr("292f26da1ab4beb4c9099f1304d2b071"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("4515472681001"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Replay - Animal Crossing"),
						λ.NewStr("description"): λ.NewStr("md5:2e211891b215c85d061adc7a4dd2d930"),
						λ.NewStr("timestamp"):   λ.NewInt(1443457610),
						λ.NewStr("upload_date"): λ.NewStr("20150928"),
						λ.NewStr("uploader_id"): λ.NewStr("694940074001"),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("https://www.gameinformer.com/video-feature/new-gameplay-today/2019/07/09/new-gameplay-today-streets-of-rogue"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("6057111913001"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("New Gameplay Today – Streets Of Rogue"),
						λ.NewStr("timestamp"):   λ.NewInt(1562699001),
						λ.NewStr("upload_date"): λ.NewStr("20190709"),
						λ.NewStr("uploader_id"): λ.NewStr("694940074001"),
					}),
				}),
			)
			GameInformerIE_BRIGHTCOVE_URL_TEMPLATE = λ.NewStr("http://players.brightcove.net/694940074001/default_default/index.html?videoId=%s")
			GameInformerIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒbrightcove_id  λ.Object
						ϒbrightcove_url λ.Object
						ϒdisplay_id     λ.Object
						ϒself           = λargs[0]
						ϒurl            = λargs[1]
						ϒwebpage        λ.Object
					)
					ϒdisplay_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒwebpage = λ.Call(λ.GetAttr(ϒself, "_download_webpage", nil), λ.NewArgs(
						ϒurl,
						ϒdisplay_id,
					), λ.KWArgs{
						{Name: "headers", Value: λ.Cal(λ.GetAttr(ϒself, "geo_verification_headers", nil))},
					})
					ϒbrightcove_id = λ.Cal(ϒclean_html, func() λ.Object {
						if λv := λ.Cal(ϒget_element_by_class, λ.NewStr("field--name-field-brightcove-video-id"), ϒwebpage); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Cal(ϒget_element_by_id, λ.NewStr("video-source-content"), ϒwebpage)
						}
					}())
					ϒbrightcove_url = func() λ.Object {
						if λ.IsTrue(ϒbrightcove_id) {
							return λ.Mod(λ.GetAttr(ϒself, "BRIGHTCOVE_URL_TEMPLATE", nil), ϒbrightcove_id)
						} else {
							return λ.Cal(λ.GetAttr(BrightcoveNewIE, "_extract_url", nil), ϒself, ϒwebpage)
						}
					}()
					return λ.Cal(λ.GetAttr(ϒself, "url_result", nil), ϒbrightcove_url, λ.NewStr("BrightcoveNew"), ϒbrightcove_id)
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("BRIGHTCOVE_URL_TEMPLATE"): GameInformerIE_BRIGHTCOVE_URL_TEMPLATE,
				λ.NewStr("_TESTS"):                  GameInformerIE__TESTS,
				λ.NewStr("_VALID_URL"):              GameInformerIE__VALID_URL,
				λ.NewStr("_real_extract"):           GameInformerIE__real_extract,
			})
		}())
	})
}