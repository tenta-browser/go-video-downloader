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
 * sunpornoie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/sunporno.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type SunPornoIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewSunPornoIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &SunPornoIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "SunPorno"
	_VALID_URL = "https?://(?:(?:www\\.)?sunporno\\.com/videos|embeds\\.sunporno\\.com/embed)/(?P<id>\\d+)"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *SunPornoIE) Key() string {
	return "SunPorno"
}

func (self *SunPornoIE) Name() string {
	return self.IE_NAME
}

func (self *SunPornoIE) _real_extract(url string) rnt.SDict {
	var (
		comment_count rnt.OptInt
		description   rnt.OptString
		duration      rnt.OptString
		formats       []interface{}
		quality       func(string) int
		thumbnail     rnt.OptString
		title         rnt.OptString
		video_ext     string
		video_id      string
		video_url     string
		view_count    rnt.OptInt
		webpage       string
	)
	video_id = (self).MatchID(url)
	webpage = (self).DownloadWebpageURL(rnt.StrFormat2("http://www.sunporno.com/videos/%s", video_id), video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	title = (self).HTMLSearchRegexOne("<title>([^<]+)</title>", webpage, "title", rnt.NoDefault, true, 0, nil)
	description = (self).HTMLSearchMetaOne("description", webpage, rnt.AsOptString("description"), false, rnt.NoDefault, 0)
	thumbnail = (self).HTMLSearchRegexOne("poster=\"([^\"]+)\"", webpage, "thumbnail", rnt.NoDefault, false, 0, nil)
	duration = rnt.ParseDuration((self).SearchRegexMulti(τ_conv_Tssω_to_Ls(τ_Tssω{
		Φ0: "itemprop=\"duration\"[^>]*>\\s*(\\d+:\\d+)\\s*<",
		Φ1: ">Duration:\\s*<span[^>]+>\\s*(\\d+:\\d+)\\s*<",
	}), webpage, "duration", rnt.NoDefault, false, 0, nil))
	view_count = rnt.IntOrNone((self).HTMLSearchRegexOne("class=\"views\">(?:<noscript>)?\\s*(\\d+)\\s*<", webpage, "view count", rnt.NoDefault, false, 0, nil), 1, rnt.OptInt{}, 1)
	comment_count = rnt.IntOrNone((self).HTMLSearchRegexOne("(\\d+)</b> Comments?", webpage, "comment count", nil, false, 0, nil), 1, rnt.OptInt{}, 1)
	formats = []interface{}{}
	quality = rnt.Qualities([]string{"mp4", "flv"})
	for _, τel := range rnt.ReFindAllOne("<(?:source|video) src=\"([^\"]+)\"", webpage, 0) {
		video_url = τel
		video_ext = rnt.DetermineExt(rnt.AsOptString(video_url), "unknown_video")
		formats = append(formats, rnt.SDict{
			"url":       video_url,
			"format_id": video_ext,
			"quality":   quality(video_ext),
		})
	}
	formats = τ_conv_Ld_to_Lα((self).SortFormats(τ_conv_Lα_to_Ld(formats)))
	return rnt.SDict{
		"id":            video_id,
		"title":         title,
		"description":   description,
		"thumbnail":     thumbnail,
		"duration":      duration,
		"view_count":    view_count,
		"comment_count": comment_count,
		"formats":       formats,
		"age_limit":     18,
	}
}

func (self *SunPornoIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("SunPorno", NewSunPornoIE)
}
