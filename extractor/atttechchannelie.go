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
 * atttechchannelie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/atttechchannel.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type ATTTechChannelIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewATTTechChannelIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &ATTTechChannelIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "ATTTechChannel"
	_VALID_URL = "https?://techchannel\\.att\\.com/play-video\\.cfm/([^/]+/)*(?P<id>.+)"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *ATTTechChannelIE) Key() string {
	return "ATTTechChannel"
}

func (self *ATTTechChannelIE) Name() string {
	return self.IE_NAME
}

func (self *ATTTechChannelIE) _real_extract(url string) rnt.SDict {
	var (
		description rnt.OptString
		display_id  string
		thumbnail   rnt.OptString
		title       rnt.OptString
		upload_date rnt.OptString
		video_id    rnt.OptString
		video_url   rnt.OptString
		webpage     string
	)
	display_id = (self).MatchID(url)
	webpage = (self).DownloadWebpageURL(url, display_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	video_url = (self).SearchRegexOne("url\\s*:\\s*'(rtmp://[^']+)'", webpage, "video URL", rnt.NoDefault, true, 0, nil)
	video_id = (self).SearchRegexOne("mediaid\\s*=\\s*(\\d+)", webpage, "video id", rnt.NoDefault, false, 0, nil)
	title = (self).OgSearchTitle(webpage, rnt.NoDefault, true)
	description = (self).OgSearchDescription(webpage, rnt.NoDefault)
	thumbnail = (self).OgSearchThumbnail(webpage, rnt.NoDefault)
	upload_date = rnt.UnifiedStrDate((self).SearchRegexOne("[Rr]elease\\s+date:\\s*(\\d{1,2}/\\d{1,2}/\\d{4})", webpage, "upload date", rnt.NoDefault, false, 0, nil), false)
	return rnt.SDict{
		"id":          video_id,
		"display_id":  display_id,
		"url":         video_url,
		"ext":         "flv",
		"title":       title,
		"description": description,
		"thumbnail":   thumbnail,
		"upload_date": upload_date,
	}
}

func (self *ATTTechChannelIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("ATTTechChannel", NewATTTechChannelIE)
}
