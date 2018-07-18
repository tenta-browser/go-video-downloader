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
 * hentaistigmaie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/hentaistigma.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type HentaiStigmaIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewHentaiStigmaIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &HentaiStigmaIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "HentaiStigma"
	_VALID_URL = "^https?://hentai\\.animestigma\\.com/(?P<id>[^/]+)"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *HentaiStigmaIE) Key() string {
	return "HentaiStigma"
}

func (self *HentaiStigmaIE) Name() string {
	return self.IE_NAME
}

func (self *HentaiStigmaIE) _real_extract(url string) rnt.SDict {
	var (
		title        rnt.OptString
		video_id     string
		video_url    rnt.OptString
		webpage      string
		wrap_url     rnt.OptString
		wrap_webpage string
	)
	video_id = (self).MatchID(url)
	webpage = (self).DownloadWebpageURL(url, video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{}, nil)
	title = (self).HTMLSearchRegexOne("<h2[^>]+class=\"posttitle\"[^>]*><a[^>]*>([^<]+)</a>", webpage, "title", rnt.NoDefault, true, 0, nil)
	wrap_url = (self).HTMLSearchRegexOne("<iframe[^>]+src=\"([^\"]+mp4)\"", webpage, "wrapper url", rnt.NoDefault, true, 0, nil)
	wrap_webpage = (self).DownloadWebpageURL(wrap_url.Get(), video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{}, nil)
	video_url = (self).HTMLSearchRegexOne("file\\s*:\\s*\"([^\"]+)\"", wrap_webpage, "video url", rnt.NoDefault, true, 0, nil)
	return rnt.SDict{
		"id":        video_id,
		"url":       video_url,
		"title":     title,
		"age_limit": 18,
	}
}

func (self *HentaiStigmaIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("HentaiStigma", NewHentaiStigmaIE)
}
