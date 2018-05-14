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
 * tentaie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/tenta.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type TentaIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewTentaIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &TentaIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "Tenta"
	_VALID_URL = "https?://(?:www\\.)?tenta\\.com/how-to-download-videos"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *TentaIE) Key() string {
	return "Tenta"
}

func (self *TentaIE) Name() string {
	return self.IE_NAME
}

func (self *TentaIE) _real_extract(url string) rnt.SDict {
	var (
		title     rnt.OptString
		video_id  string
		video_url rnt.OptString
		webpage   string
	)
	video_id = "howto"
	webpage = (self).DownloadWebpageURL(url, video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	title = (self).OgSearchTitle(webpage, nil, true)
	video_url = (self).SearchRegexOne("<source[^>]+src=[\"\\'](.+?)[\"\\'][^>]+type=[\"\\']video/mp4[\"\\'][^>]*/>", webpage, "video_url", rnt.NoDefault, true, 0, nil)
	video_url = rnt.AsOptString(rnt.URLJoin(url, video_url))
	return rnt.SDict{
		"id":    video_id,
		"title": title,
		"url":   video_url,
	}
}

func (self *TentaIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("Tenta", NewTentaIE)
}
