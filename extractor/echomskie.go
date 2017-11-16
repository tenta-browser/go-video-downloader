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
 * echomskie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/echomsk.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type EchoMskIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewEchoMskIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &EchoMskIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "EchoMsk"
	_VALID_URL = "https?://(?:www\\.)?echo\\.msk\\.ru/sounds/(?P<id>\\d+)"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *EchoMskIE) Key() string {
	return "EchoMsk"
}

func (self *EchoMskIE) Name() string {
	return self.IE_NAME
}

func (self *EchoMskIE) _real_extract(url string) rnt.SDict {
	var (
		air_date  rnt.OptString
		audio_url rnt.OptString
		title     rnt.OptString
		video_id  string
		webpage   string
	)
	video_id = (self).MatchID(url)
	webpage = (self).DownloadWebpageURL(url, video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	audio_url = (self).SearchRegexOne("<a rel=\"mp3\" href=\"([^\"]+)\">", webpage, "audio URL", rnt.NoDefault, true, 0, nil)
	title = (self).HTMLSearchRegexOne("<a href=\"/programs/[^\"]+\" target=\"_blank\">([^<]+)</a>", webpage, "title", rnt.NoDefault, true, 0, nil)
	air_date = (self).HTMLSearchRegexOne("(?s)<div class=\"date\">(.+?)</div>", webpage, "date", nil, false, 0, nil)
	if τ_isTruthy_Os(air_date) {
		air_date = rnt.AsOptString(rnt.ReSub("(\\s)\\1+", "\\1", air_date.Get(), 0, 0))
		if τ_isTruthy_Os(air_date) {
			title = rnt.AsOptString(rnt.StrFormat2("%s - %s", title, air_date))
		}
	}
	return rnt.SDict{
		"id":    video_id,
		"url":   audio_url,
		"title": title,
	}
}

func (self *EchoMskIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("EchoMsk", NewEchoMskIE)
}
