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
 * xtubeie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/xtube.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type XTubeIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewXTubeIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &XTubeIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "XTube"
	_VALID_URL = "(?x)\n                        (?:\n                            xtube:|\n                            https?://(?:www\\.)?xtube\\.com/(?:watch\\.php\\?.*\\bv=|video-watch/(?P<display_id>[^/]+)-)\n                        )\n                        (?P<id>[^/?&#]+)\n                    "
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *XTubeIE) Key() string {
	return "XTube"
}

func (self *XTubeIE) Name() string {
	return self.IE_NAME
}

func (self *XTubeIE) _real_extract(url string) rnt.SDict {
	var (
		comment_count rnt.OptInt
		description   rnt.OptString
		display_id    rnt.OptString
		duration      rnt.OptString
		format_id     string
		format_url    interface{}
		formats       []interface{}
		mobj          rnt.Match
		sources       interface{}
		title         rnt.OptString
		uploader      rnt.OptString
		url_pattern   string
		video_id      rnt.OptString
		view_count    rnt.OptInt
		webpage       string
		τmp1          τ_Tsαω
	)
	mobj = rnt.ReMatch((self).VALIDURL, url, 0)
	video_id = rnt.ReMatchGroupOne(mobj, "id")
	display_id = rnt.ReMatchGroupOne(mobj, "display_id")
	if !(τ_isTruthy_Os(display_id)) {
		display_id = video_id
	}
	if rnt.StrIsDigit(video_id.Get()) && (len(video_id.Get())) < (11) {
		url_pattern = "http://www.xtube.com/video-watch/-%s"
	} else {
		url_pattern = "http://www.xtube.com/watch.php?v=%s"
	}
	webpage = (self).DownloadWebpageURL(rnt.StrFormat2(url_pattern, video_id), display_id.Get(), rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{
		"Cookie": "age_verified=1; cookiesAccepted=1",
	}, rnt.SDict{})
	sources = (self).ParseJSON((self).SearchRegexOne("([\"\\'])?sources\\1?\\s*:\\s*(?P<sources>{.+?}),", webpage, "sources", rnt.NoDefault, true, 0, "sources").Get(), video_id.Get(), rnt.JsToJSON, true)
	formats = []interface{}{}
	for _, τel := range rnt.DictItems(τ_cast_α_to_d(sources)) {
		τmp1 = τel
		format_id = (τmp1).Φ0
		format_url = (τmp1).Φ1
		formats = append(formats, rnt.SDict{
			"url":       format_url,
			"format_id": format_id,
			"height":    rnt.IntOrNone(format_id, 1, rnt.OptInt{}, 1),
		})
	}
	formats = τ_conv_Ld_to_Lα((self).RemoveDuplicateFormats(τ_conv_Lα_to_Ld(formats)))
	formats = τ_conv_Ld_to_Lα((self).SortFormats(τ_conv_Lα_to_Ld(formats)))
	title = (self).SearchRegexMulti(τ_conv_Tssω_to_Ls(τ_Tssω{
		Φ0: "<h1>\\s*(?P<title>[^<]+?)\\s*</h1>",
		Φ1: "videoTitle\\s*:\\s*([\"\\'])(?P<title>.+?)\\1",
	}), webpage, "title", rnt.NoDefault, true, 0, "title")
	description = (self).SearchRegexOne("</h1>\\s*<p>([^<]+)", webpage, "description", rnt.NoDefault, false, 0, nil)
	uploader = (self).SearchRegexMulti(τ_conv_Tssω_to_Ls(τ_Tssω{
		Φ0: "<input[^>]+name=\"contentOwnerId\"[^>]+value=\"([^\"]+)\"",
		Φ1: "<span[^>]+class=\"nickname\"[^>]*>([^<]+)",
	}), webpage, "uploader", rnt.NoDefault, false, 0, nil)
	duration = rnt.ParseDuration((self).SearchRegexOne("<dt>Runtime:?</dt>\\s*<dd>([^<]+)</dd>", webpage, "duration", rnt.NoDefault, false, 0, nil))
	view_count = rnt.StrToInt((self).SearchRegexOne("<dt>Views:?</dt>\\s*<dd>([\\d,\\.]+)</dd>", webpage, "view count", rnt.NoDefault, false, 0, nil))
	comment_count = rnt.StrToInt((self).HTMLSearchRegexOne(">Comments? \\(([\\d,\\.]+)\\)<", webpage, "comment count", rnt.NoDefault, false, 0, nil))
	return rnt.SDict{
		"id":            video_id,
		"display_id":    display_id,
		"title":         title,
		"description":   description,
		"uploader":      uploader,
		"duration":      duration,
		"view_count":    view_count,
		"comment_count": comment_count,
		"age_limit":     18,
		"formats":       formats,
	}
}

func (self *XTubeIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("XTube", NewXTubeIE)
}
