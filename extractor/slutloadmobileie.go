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
 * slutloadmobileie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/slutload.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type SlutloadMobileIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewSlutloadMobileIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &SlutloadMobileIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "SlutloadMobile"
	_VALID_URL = "^https?://mobile\\.slutload\\.com/video/[^/]+/(?P<id>[^/]+)/?$"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *SlutloadMobileIE) Key() string {
	return "SlutloadMobile"
}

func (self *SlutloadMobileIE) Name() string {
	return self.IE_NAME
}

func (self *SlutloadMobileIE) _real_extract(url string) rnt.SDict {
	var (
		thumbnail   rnt.OptString
		video_id    string
		video_title string
		video_url   rnt.OptString
		webpage     string
	)
	video_id = (self).MatchID(url)
	webpage = (self).DownloadWebpageURL(url, video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{}, nil)
	video_title = rnt.StrStrip((self).HTMLSearchRegexOne("<div class=[\"\\']videoHd[^\"\\']+[\"\\']>([^<]+)</div>", webpage, "title", rnt.NoDefault, true, 0, nil).Get(), "")
	video_url = (self).HTMLSearchRegexOne("(?s)<video id=[\"\\']html5video[\"\\'].+?src=[\"\\']([^\"\\']+)[\"\\']", webpage, "video URL", rnt.NoDefault, true, 0, nil)
	thumbnail = (self).HTMLSearchRegexOne("(?s)<video id=[\"\\']html5video[\"\\'].+?poster=[\"\\']([^\"\\']+)[\"\\']", webpage, "thumbnail, fatal=False", rnt.NoDefault, true, 0, nil)
	return rnt.SDict{
		"id":        video_id,
		"url":       video_url,
		"title":     video_title,
		"thumbnail": thumbnail,
		"age_limit": 18,
	}
}

func (self *SlutloadMobileIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("SlutloadMobile", NewSlutloadMobileIE)
}
