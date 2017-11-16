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
 * bildie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/bild.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type BildIE struct {
	*rnt.CommonIE
	IE_DESC string
	IE_NAME string
}

func NewBildIE() rnt.InfoExtractor {
	var (
		IE_DESC    string
		IE_NAME    string
		_VALID_URL string
	)
	self := &BildIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "Bild"
	_VALID_URL = "https?://(?:www\\.)?bild\\.de/(?:[^/]+/)+(?P<display_id>[^/]+)-(?P<id>\\d+)(?:,auto=true)?\\.bild\\.html"
	IE_DESC = "Bild.de"
	self.IE_DESC = IE_DESC
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *BildIE) Key() string {
	return "Bild"
}

func (self *BildIE) Name() string {
	return self.IE_NAME
}

func (self *BildIE) _real_extract(url string) rnt.SDict {
	var (
		video_data interface{}
		video_id   string
	)
	video_id = (self).MatchID(url)
	video_data = (self).DownloadJSON(((rnt.StrSplit(url, ".bild.html", -(1)))[0] + ",view=json.bild.html"), video_id, rnt.AsOptString("Downloading JSON metadata"), rnt.AsOptString("Unable to download JSON metadata"), nil, true, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	return rnt.SDict{
		"id":          video_id,
		"title":       rnt.StrStrip(rnt.UnescapeHTML(rnt.CastToOptString(rnt.UnsafeSubscript(video_data, "title"))).Get(), ""),
		"description": rnt.UnescapeHTML(rnt.CastToOptString(rnt.DictGet(τ_cast_α_to_d(video_data), "description", nil))),
		"url":         rnt.UnsafeSubscript(rnt.UnsafeSubscript(rnt.UnsafeSubscript(rnt.UnsafeSubscript(rnt.UnsafeSubscript(video_data, "clipList"), 0), "srces"), 0), "src"),
		"thumbnail":   rnt.DictGet(τ_cast_α_to_d(video_data), "poster", nil),
		"duration":    rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(video_data), "durationSec", nil), 1, rnt.OptInt{}, 1),
	}
}

func (self *BildIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("Bild", NewBildIE)
}
