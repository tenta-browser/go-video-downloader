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
 * thisamericanlifeie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/thisamericanlife.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type ThisAmericanLifeIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewThisAmericanLifeIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &ThisAmericanLifeIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "ThisAmericanLife"
	_VALID_URL = "https?://(?:www\\.)?thisamericanlife\\.org/(?:radio-archives/episode/|play_full\\.php\\?play=)(?P<id>\\d+)"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *ThisAmericanLifeIE) Key() string {
	return "ThisAmericanLife"
}

func (self *ThisAmericanLifeIE) Name() string {
	return self.IE_NAME
}

func (self *ThisAmericanLifeIE) _real_extract(url string) rnt.SDict {
	var (
		video_id string
		webpage  string
	)
	video_id = (self).MatchID(url)
	webpage = (self).DownloadWebpageURL(rnt.StrFormat2("http://www.thisamericanlife.org/radio-archives/episode/%s", video_id), video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{}, nil)
	return rnt.SDict{
		"id":          video_id,
		"url":         rnt.StrFormat3("http://stream.thisamericanlife.org/{0}/stream/{0}_64k.m3u8", video_id),
		"protocol":    "m3u8_native",
		"ext":         "m4a",
		"acodec":      "aac",
		"vcodec":      "none",
		"abr":         64,
		"title":       (self).HTMLSearchMetaOne("twitter:title", webpage, rnt.AsOptString("title"), true, rnt.NoDefault, 0),
		"description": (self).HTMLSearchMetaOne("description", webpage, rnt.AsOptString("description"), false, rnt.NoDefault, 0),
		"thumbnail":   (self).OgSearchThumbnail(webpage, rnt.NoDefault),
	}
}

func (self *ThisAmericanLifeIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("ThisAmericanLife", NewThisAmericanLifeIE)
}
