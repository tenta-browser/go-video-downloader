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
 * skylinewebcamsie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/skylinewebcams.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type SkylineWebcamsIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewSkylineWebcamsIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &SkylineWebcamsIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "SkylineWebcams"
	_VALID_URL = "https?://(?:www\\.)?skylinewebcams\\.com/[^/]+/webcam/(?:[^/]+/)+(?P<id>[^/]+)\\.html"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *SkylineWebcamsIE) Key() string {
	return "SkylineWebcams"
}

func (self *SkylineWebcamsIE) Name() string {
	return self.IE_NAME
}

func (self *SkylineWebcamsIE) _real_extract(url string) rnt.SDict {
	var (
		description rnt.OptString
		stream_url  rnt.OptString
		title       rnt.OptString
		video_id    string
		webpage     string
	)
	video_id = (self).MatchID(url)
	webpage = (self).DownloadWebpageURL(url, video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{}, nil)
	stream_url = (self).SearchRegexOne("url\\s*:\\s*([\"\\'])(?P<url>(?:https?:)?//.+?\\.m3u8.*?)\\1", webpage, "stream url", rnt.NoDefault, true, 0, "url")
	title = (self).OgSearchTitle(webpage, rnt.NoDefault, true)
	description = (self).OgSearchDescription(webpage, rnt.NoDefault)
	return rnt.SDict{
		"id":          video_id,
		"url":         stream_url,
		"ext":         "mp4",
		"title":       (self).LiveTitle(title.Get()),
		"description": description,
		"is_live":     true,
	}
}

func (self *SkylineWebcamsIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("SkylineWebcams", NewSkylineWebcamsIE)
}
