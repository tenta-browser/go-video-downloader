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
 * morningstarie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/morningstar.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type MorningstarIE struct {
	*rnt.CommonIE
	IE_DESC string
	IE_NAME string
}

func NewMorningstarIE() rnt.InfoExtractor {
	var (
		IE_DESC    string
		IE_NAME    string
		_VALID_URL string
	)
	self := &MorningstarIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "Morningstar"
	IE_DESC = "morningstar.com"
	_VALID_URL = "https?://(?:(?:www|news)\\.)morningstar\\.com/[cC]over/video[cC]enter\\.aspx\\?id=(?P<id>[0-9]+)"
	self.IE_DESC = IE_DESC
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *MorningstarIE) Key() string {
	return "Morningstar"
}

func (self *MorningstarIE) Name() string {
	return self.IE_NAME
}

func (self *MorningstarIE) _real_extract(url string) rnt.SDict {
	var (
		description rnt.OptString
		mobj        rnt.Match
		thumbnail   rnt.OptString
		title       rnt.OptString
		video_id    rnt.OptString
		video_url   rnt.OptString
		webpage     string
	)
	mobj = rnt.ReMatch((self).VALIDURL, url, 0)
	video_id = rnt.ReMatchGroupOne(mobj, "id")
	webpage = (self).DownloadWebpageURL(url, video_id.Get(), rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	title = (self).HTMLSearchRegexOne("<h1 id=\"titleLink\">(.*?)</h1>", webpage, "title", rnt.NoDefault, true, 0, nil)
	video_url = (self).HTMLSearchRegexOne("<input type=\"hidden\" id=\"hidVideoUrl\" value=\"([^\"]+)\"", webpage, "video URL", rnt.NoDefault, true, 0, nil)
	thumbnail = (self).HTMLSearchRegexOne("<input type=\"hidden\" id=\"hidSnapshot\" value=\"([^\"]+)\"", webpage, "thumbnail", rnt.NoDefault, false, 0, nil)
	description = (self).HTMLSearchRegexOne("<div id=\"mstarDeck\".*?>(.*?)</div>", webpage, "description", rnt.NoDefault, false, 0, nil)
	return rnt.SDict{
		"id":          video_id,
		"title":       title,
		"url":         video_url,
		"thumbnail":   thumbnail,
		"description": description,
	}
}

func (self *MorningstarIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("Morningstar", NewMorningstarIE)
}
