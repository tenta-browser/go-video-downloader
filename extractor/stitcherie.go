/**
 * Go Video Downloader
 *
 *    Copyright 2017 Tenta, LLC
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
 * stitcherie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/stitcher.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type StitcherIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewStitcherIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &StitcherIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "Stitcher"
	_VALID_URL = "https?://(?:www\\.)?stitcher\\.com/podcast/(?:[^/]+/)+e/(?:(?P<display_id>[^/#?&]+?)-)?(?P<id>\\d+)(?:[/#?&]|$)"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *StitcherIE) Key() string {
	return "Stitcher"
}

func (self *StitcherIE) Name() string {
	return self.IE_NAME
}

func (self *StitcherIE) _real_extract(url string) rnt.SDict {
	var (
		audio_id    rnt.OptString
		description rnt.OptString
		display_id  rnt.OptString
		duration    rnt.OptInt
		episode     interface{}
		formats     []rnt.SDict
		mobj        rnt.Match
		thumbnail   interface{}
		title       rnt.OptString
		webpage     string
	)
	mobj = rnt.ReMatch((self).VALIDURL, url, 0)
	audio_id = rnt.ReMatchGroupOne(mobj, "id")
	display_id = func() rnt.OptString {
		if v := (rnt.ReMatchGroupOne(mobj, "display_id")); τ_isTruthy_Os(v) {
			return v
		} else {
			return audio_id
		}
	}()
	webpage = (self).DownloadWebpageURL(url, display_id.Get(), rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{}, nil)
	episode = rnt.UnsafeSubscript(rnt.UnsafeSubscript((self).ParseJSON(rnt.JsToJSON((self).SearchRegexOne("(?s)var\\s+stitcher(?:Config)?\\s*=\\s*({.+?});\\n", webpage, "episode config", rnt.NoDefault, true, 0, nil).Get()), display_id.Get(), nil, true), "config"), "episode")
	title = rnt.UnescapeHTML(rnt.CastToOptString(rnt.UnsafeSubscript(episode, "title")))
	formats = func() []rnt.SDict {
		τresult := []rnt.SDict{}
		for _, τel := range τ_conv_Tsω_to_Ls(τ_Tsω{
			Φ0: "episodeURL",
		}) {
			var (
				episode_key string
			)
			episode_key = τel
			if !(rnt.IsTruthy(rnt.DictGet(τ_cast_α_to_d(episode), episode_key, nil))) {
				continue
			}
			τresult = append(τresult, rnt.SDict{
				"url": rnt.UnsafeSubscript(episode, episode_key),
				"ext": func() string {
					if v := (rnt.DetermineExt(rnt.CastToOptString(rnt.UnsafeSubscript(episode, episode_key)), "unknown_video")); (v) != "" {
						return v
					} else {
						return "mp3"
					}
				}(),
				"vcodec": "none",
			})
		}
		return τresult
	}()
	description = (self).SearchRegexOne("Episode Info:\\s*</span>([^<]+)<", webpage, "description", rnt.NoDefault, false, 0, nil)
	duration = rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(episode), "duration", nil), 1, rnt.OptInt{}, 1)
	thumbnail = rnt.DictGet(τ_cast_α_to_d(episode), "episodeImage", nil)
	return rnt.SDict{
		"id":          audio_id,
		"display_id":  display_id,
		"title":       title,
		"description": description,
		"duration":    duration,
		"thumbnail":   thumbnail,
		"formats":     formats,
	}
}

func (self *StitcherIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("Stitcher", NewStitcherIE)
}
