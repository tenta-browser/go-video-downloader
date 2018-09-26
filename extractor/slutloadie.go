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
 * slutloadie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/slutload.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type SlutloadIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewSlutloadIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &SlutloadIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "Slutload"
	_VALID_URL = "https?://(?:\\w+\\.)?slutload\\.com/(?:video/[^/]+|embed_player|watch)/(?P<id>[^/]+)"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *SlutloadIE) Key() string {
	return "Slutload"
}

func (self *SlutloadIE) Name() string {
	return self.IE_NAME
}

func (self *SlutloadIE) _real_extract(url string) rnt.SDict {
	var (
		embed_page string
		extract    func(string) rnt.OptString
		title      rnt.OptString
		video_id   string
		video_url  rnt.OptString
	)
	video_id = (self).MatchID(url)
	embed_page = (self).DownloadWebpageURL(rnt.StrFormat2("http://www.slutload.com/embed_player/%s", video_id), video_id, rnt.AsOptString("Downloading embed page"), rnt.OptString{}, false, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{}, nil)
	if (embed_page) != "" {
		extract = func(what string) rnt.OptString {
			return (self).HTMLSearchRegexOne(rnt.StrFormat2("data-video-%s=([\"\\'])(?P<url>(?:(?!\\1).)+)\\1", what), embed_page, rnt.StrFormat2("video %s", what), nil, true, 0, "url")
		}
		video_url = extract("url")
		if τ_isTruthy_Os(video_url) {
			title = (self).HTMLSearchRegexOne("<title>([^<]+)", embed_page, "title", video_id, true, 0, nil)
			return rnt.SDict{
				"id":        video_id,
				"url":       video_url,
				"title":     title,
				"thumbnail": extract("preview"),
				"age_limit": 18,
			}
		}
	}
	return rnt.SDict{}
}

func (self *SlutloadIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("Slutload", NewSlutloadIE)
}
