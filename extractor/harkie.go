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
 * harkie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/hark.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type HarkIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewHarkIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &HarkIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "Hark"
	_VALID_URL = "https?://(?:www\\.)?hark\\.com/clips/(?P<id>.+?)-.+"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *HarkIE) Key() string {
	return "Hark"
}

func (self *HarkIE) Name() string {
	return self.IE_NAME
}

func (self *HarkIE) _real_extract(url string) rnt.SDict {
	var (
		data     interface{}
		video_id string
	)
	video_id = (self).MatchID(url)
	data = (self).DownloadJSON(rnt.StrFormat2("http://www.hark.com/clips/%s.json", video_id), video_id, rnt.AsOptString("Downloading JSON metadata"), rnt.AsOptString("Unable to download JSON metadata"), nil, true, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{}, nil)
	return rnt.SDict{
		"id":          video_id,
		"url":         rnt.UnsafeSubscript(data, "url"),
		"title":       rnt.UnsafeSubscript(data, "name"),
		"description": rnt.DictGet(τ_cast_α_to_d(data), "description", nil),
		"thumbnail":   rnt.DictGet(τ_cast_α_to_d(data), "image_original", nil),
		"duration":    rnt.DictGet(τ_cast_α_to_d(data), "duration", nil),
	}
}

func (self *HarkIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("Hark", NewHarkIE)
}
