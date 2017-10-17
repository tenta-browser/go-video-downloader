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
}

func NewMorningstarIE() rnt.InfoExtractor {
	ret := &MorningstarIE{}
	ret.CommonIE = rnt.NewCommonIE()
	ret.IE_DESC = "morningstar.com"
	ret.VALIDURL = "https?://(?:(?:www|news)\\.)morningstar\\.com/[cC]over/video[cC]enter\\.aspx\\?id=(?P<id>[0-9]+)"
	return ret
}

func (self *MorningstarIE) Key() string {
	return "Morningstar"
}

func (self *MorningstarIE) Name() string {
	return "Morningstar extractor" + " (" + self.IE_DESC + ")"
}

func (self *MorningstarIE) _real_extract(url string) map[string]interface{} {
	mobj := rnt.ReMatch((self).VALIDURL, url, 0)
	video_id := rnt.ReMatchGroupOne(mobj, "id")
	webpage := (self).DownloadWebpageURL(url, video_id.Get(), rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, rnt.OptString{}, map[string]interface{}{}, map[string]interface{}{})
	title := (self).HTMLSearchRegexOne("<h1 id=\"titleLink\">(.*?)</h1>", webpage, "title", rnt.NoDefault, true, 0, nil)
	video_url := (self).HTMLSearchRegexOne("<input type=\"hidden\" id=\"hidVideoUrl\" value=\"([^\"]+)\"", webpage, "video URL", rnt.NoDefault, true, 0, nil)
	thumbnail := (self).HTMLSearchRegexOne("<input type=\"hidden\" id=\"hidSnapshot\" value=\"([^\"]+)\"", webpage, "thumbnail", rnt.NoDefault, false, 0, nil)
	description := (self).HTMLSearchRegexOne("<div id=\"mstarDeck\".*?>(.*?)</div>", webpage, "description", rnt.NoDefault, false, 0, nil)
	return map[string]interface{}{"id": video_id,
		"title":       title,
		"url":         video_url,
		"thumbnail":   thumbnail,
		"description": description}
}

func (self *MorningstarIE) Extract(url string) (*rnt.VideoResult, error) {
	return rnt.RunExtractor(url, self._real_extract)
}

func init() {
	registerFactory("Morningstar", NewMorningstarIE)
}
