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
 * rozhlasie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/rozhlas.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type RozhlasIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewRozhlasIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &RozhlasIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "Rozhlas"
	_VALID_URL = "https?://(?:www\\.)?prehravac\\.rozhlas\\.cz/audio/(?P<id>[0-9]+)"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *RozhlasIE) Key() string {
	return "Rozhlas"
}

func (self *RozhlasIE) Name() string {
	return self.IE_NAME
}

func (self *RozhlasIE) _real_extract(url string) rnt.SDict {
	var (
		audio_id    string
		description rnt.OptString
		duration    rnt.OptInt
		title       rnt.OptString
		webpage     string
	)
	audio_id = (self).MatchID(url)
	webpage = (self).DownloadWebpageURL(rnt.StrFormat2("http://prehravac.rozhlas.cz/audio/%s", audio_id), audio_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	title = func() rnt.OptString {
		if v := ((self).HTMLSearchRegexOne("<h3>(.+?)</h3>\\s*<p[^>]*>.*?</p>\\s*<div[^>]+id=[\"\\']player-track", webpage, "title", nil, true, 0, nil)); τ_isTruthy_Os(v) {
			return v
		} else {
			return rnt.RemoveStart((self).OgSearchTitle(webpage, rnt.NoDefault, true), "Radio Wave - ")
		}
	}()
	description = (self).HTMLSearchRegexOne("<p[^>]+title=([\"\\'])(?P<url>(?:(?!\\1).)+)\\1[^>]*>.*?</p>\\s*<div[^>]+id=[\"\\']player-track", webpage, "description", rnt.NoDefault, false, 0, "url")
	duration = rnt.IntOrNone((self).SearchRegexOne("data-duration=[\"\\'](\\d+)", webpage, "duration", nil, true, 0, nil), 1, rnt.OptInt{}, 1)
	return rnt.SDict{
		"id":          audio_id,
		"url":         rnt.StrFormat2("http://media.rozhlas.cz/_audio/%s.mp3", audio_id),
		"title":       title,
		"description": description,
		"duration":    duration,
		"vcodec":      "none",
	}
}

func (self *RozhlasIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("Rozhlas", NewRozhlasIE)
}
