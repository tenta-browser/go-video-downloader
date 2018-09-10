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
 * chirbitie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/chirbit.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type ChirbitIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewChirbitIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &ChirbitIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "Chirbit"
	IE_NAME = "chirbit"
	_VALID_URL = "https?://(?:www\\.)?chirb\\.it/(?:(?:wp|pl)/|fb_chirbit_player\\.swf\\?key=)?(?P<id>[\\da-zA-Z]+)"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *ChirbitIE) Key() string {
	return "Chirbit"
}

func (self *ChirbitIE) Name() string {
	return self.IE_NAME
}

func (self *ChirbitIE) _real_extract(url string) rnt.SDict {
	var (
		audio_id    string
		audio_url   string
		data_fd     rnt.OptString
		description rnt.OptString
		duration    rnt.OptString
		title       rnt.OptString
		uploader    rnt.OptString
		webpage     string
	)
	audio_id = (self).MatchID(url)
	webpage = (self).DownloadWebpageURL(rnt.StrFormat2("http://chirb.it/%s", audio_id), audio_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{}, nil)
	data_fd = (self).SearchRegexOne("data-fd=([\"\\'])(?P<url>(?:(?!\\1).)+)\\1", webpage, "data fd", rnt.NoDefault, true, 0, "url")
	audio_url = rnt.BytesToStr(rnt.Base64StdDecodeString(rnt.StrSlice(data_fd.Get(), rnt.OptInt{}, rnt.OptInt{}, rnt.AsOptInt(-(1)))), "utf-8")
	title = (self).SearchRegexOne("class=[\"\\']chirbit-title[\"\\'][^>]*>([^<]+)", webpage, "title", rnt.NoDefault, true, 0, nil)
	description = (self).SearchRegexOne("<h3>Description</h3>\\s*<pre[^>]*>([^<]+)</pre>", webpage, "description", nil, true, 0, nil)
	duration = rnt.ParseDuration((self).SearchRegexOne("class=[\"\\']c-length[\"\\'][^>]*>([^<]+)", webpage, "duration", rnt.NoDefault, false, 0, nil))
	uploader = (self).SearchRegexOne("id=[\"\\']chirbit-username[\"\\'][^>]*>([^<]+)", webpage, "uploader", rnt.NoDefault, false, 0, nil)
	return rnt.SDict{
		"id":          audio_id,
		"url":         audio_url,
		"title":       title,
		"description": description,
		"duration":    duration,
		"uploader":    uploader,
	}
}

func (self *ChirbitIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("Chirbit", NewChirbitIE)
}
