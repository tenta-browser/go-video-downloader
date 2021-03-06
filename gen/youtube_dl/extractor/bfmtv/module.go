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
 * bfmtv/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/bfmtv.py
 */

package bfmtv

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	BFMTVArticleIE      λ.Object
	BFMTVBaseIE         λ.Object
	BFMTVIE             λ.Object
	BFMTVLiveIE         λ.Object
	InfoExtractor       λ.Object
	ϒextract_attributes λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒextract_attributes = Ωutils.ϒextract_attributes
		BFMTVBaseIE = λ.Cal(λ.TypeType, λ.StrLiteral("BFMTVBaseIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				BFMTVBaseIE_BRIGHTCOVE_URL_TEMPLATE λ.Object
				BFMTVBaseIE__VALID_URL_BASE         λ.Object
				BFMTVBaseIE__VALID_URL_TMPL         λ.Object
				BFMTVBaseIE__VIDEO_BLOCK_REGEX      λ.Object
				BFMTVBaseIE__brightcove_url_result  λ.Object
			)
			BFMTVBaseIE__VALID_URL_BASE = λ.StrLiteral("https?://(?:www\\.)?bfmtv\\.com/")
			BFMTVBaseIE__VALID_URL_TMPL = λ.Add(BFMTVBaseIE__VALID_URL_BASE, λ.StrLiteral("(?:[^/]+/)*[^/?&#]+_%s[A-Z]-(?P<id>\\d{12})\\.html"))
			BFMTVBaseIE__VIDEO_BLOCK_REGEX = λ.StrLiteral("(<div[^>]+class=\"video_block\"[^>]*>)")
			BFMTVBaseIE_BRIGHTCOVE_URL_TEMPLATE = λ.StrLiteral("http://players.brightcove.net/%s/%s_default/index.html?videoId=%s")
			BFMTVBaseIE__brightcove_url_result = λ.NewFunction("_brightcove_url_result",
				[]λ.Param{
					{Name: "self"},
					{Name: "video_id"},
					{Name: "video_block"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒaccount_id  λ.Object
						ϒplayer_id   λ.Object
						ϒself        = λargs[0]
						ϒvideo_block = λargs[2]
						ϒvideo_id    = λargs[1]
					)
					ϒaccount_id = func() λ.Object {
						if λv := λ.Calm(ϒvideo_block, "get", λ.StrLiteral("accountid")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.StrLiteral("876450612001")
						}
					}()
					ϒplayer_id = func() λ.Object {
						if λv := λ.Calm(ϒvideo_block, "get", λ.StrLiteral("playerid")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.StrLiteral("I2qBTln4u")
						}
					}()
					return λ.Calm(ϒself, "url_result", λ.Mod(λ.GetAttr(ϒself, "BRIGHTCOVE_URL_TEMPLATE", nil), λ.NewTuple(
						ϒaccount_id,
						ϒplayer_id,
						ϒvideo_id,
					)), λ.StrLiteral("BrightcoveNew"), ϒvideo_id)
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"BRIGHTCOVE_URL_TEMPLATE": BFMTVBaseIE_BRIGHTCOVE_URL_TEMPLATE,
				"_VALID_URL_BASE":         BFMTVBaseIE__VALID_URL_BASE,
				"_VALID_URL_TMPL":         BFMTVBaseIE__VALID_URL_TMPL,
				"_VIDEO_BLOCK_REGEX":      BFMTVBaseIE__VIDEO_BLOCK_REGEX,
				"_brightcove_url_result":  BFMTVBaseIE__brightcove_url_result,
			})
		}())
		BFMTVIE = λ.Cal(λ.TypeType, λ.StrLiteral("BFMTVIE"), λ.NewTuple(BFMTVBaseIE), func() λ.Dict {
			var (
				BFMTVIE_IE_NAME       λ.Object
				BFMTVIE__VALID_URL    λ.Object
				BFMTVIE__real_extract λ.Object
			)
			BFMTVIE_IE_NAME = λ.StrLiteral("bfmtv")
			BFMTVIE__VALID_URL = λ.Mod(λ.GetAttr(BFMTVBaseIE, "_VALID_URL_TMPL", nil), λ.StrLiteral("V"))
			BFMTVIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒbfmtv_id    λ.Object
						ϒself        = λargs[0]
						ϒurl         = λargs[1]
						ϒvideo_block λ.Object
						ϒwebpage     λ.Object
					)
					ϒbfmtv_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒwebpage = λ.Calm(ϒself, "_download_webpage", ϒurl, ϒbfmtv_id)
					ϒvideo_block = λ.Cal(ϒextract_attributes, λ.Calm(ϒself, "_search_regex", λ.GetAttr(ϒself, "_VIDEO_BLOCK_REGEX", nil), ϒwebpage, λ.StrLiteral("video block")))
					return λ.Calm(ϒself, "_brightcove_url_result", λ.GetItem(ϒvideo_block, λ.StrLiteral("videoid")), ϒvideo_block)
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"IE_NAME":       BFMTVIE_IE_NAME,
				"_VALID_URL":    BFMTVIE__VALID_URL,
				"_real_extract": BFMTVIE__real_extract,
			})
		}())
		BFMTVLiveIE = λ.Cal(λ.TypeType, λ.StrLiteral("BFMTVLiveIE"), λ.NewTuple(BFMTVIE), func() λ.Dict {
			var (
				BFMTVLiveIE__VALID_URL λ.Object
			)
			BFMTVLiveIE__VALID_URL = λ.Add(λ.GetAttr(BFMTVBaseIE, "_VALID_URL_BASE", nil), λ.StrLiteral("(?P<id>(?:[^/]+/)?en-direct)"))
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": BFMTVLiveIE__VALID_URL,
			})
		}())
		BFMTVArticleIE = λ.Cal(λ.TypeType, λ.StrLiteral("BFMTVArticleIE"), λ.NewTuple(BFMTVBaseIE), func() λ.Dict {
			var (
				BFMTVArticleIE__VALID_URL λ.Object
			)
			BFMTVArticleIE__VALID_URL = λ.Mod(λ.GetAttr(BFMTVBaseIE, "_VALID_URL_TMPL", nil), λ.StrLiteral("A"))
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": BFMTVArticleIE__VALID_URL,
			})
		}())
	})
}
