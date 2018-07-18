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
 * tassie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/tass.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type TassIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewTassIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &TassIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "Tass"
	_VALID_URL = "https?://(?:tass\\.ru|itar-tass\\.com)/[^/]+/(?P<id>\\d+)"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *TassIE) Key() string {
	return "Tass"
}

func (self *TassIE) Name() string {
	return self.IE_NAME
}

func (self *TassIE) _real_extract(url string) rnt.SDict {
	var (
		formats   []interface{}
		label     interface{}
		quality   func(string) int
		source    interface{}
		sources   interface{}
		video_id  string
		video_url interface{}
		webpage   string
	)
	video_id = (self).MatchID(url)
	webpage = (self).DownloadWebpageURL(url, video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{}, nil)
	sources = rnt.JSONLoads(rnt.JsToJSON((self).SearchRegexOne("(?s)sources\\s*:\\s*(\\[.+?\\])", webpage, "sources", rnt.NoDefault, true, 0, nil).Get()))
	quality = rnt.Qualities([]string{"sd", "hd"})
	formats = []interface{}{}
	for _, τel := range τ_cast_α_to_Lα(sources) {
		source = τel
		video_url = rnt.DictGet(τ_cast_α_to_d(source), "file", nil)
		if !(rnt.IsTruthy(video_url)) || !(rnt.StrStartswith(rnt.CastToString(video_url), "http", rnt.OptInt{}, rnt.OptInt{})) || !(rnt.StrEndswith(rnt.CastToString(video_url), ".mp4", rnt.OptInt{}, rnt.OptInt{})) {
			continue
		}
		label = rnt.DictGet(τ_cast_α_to_d(source), "label", nil)
		formats = append(formats, rnt.SDict{
			"url":       video_url,
			"format_id": label,
			"quality":   quality(rnt.CastToString(label)),
		})
	}
	formats = τ_conv_Ld_to_Lα((self).SortFormats(τ_conv_Lα_to_Ld(formats)))
	return rnt.SDict{
		"id":          video_id,
		"title":       (self).OgSearchTitle(webpage, rnt.NoDefault, true),
		"description": (self).OgSearchDescription(webpage, rnt.NoDefault),
		"thumbnail":   (self).OgSearchThumbnail(webpage, rnt.NoDefault),
		"formats":     formats,
	}
}

func (self *TassIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("Tass", NewTassIE)
}
