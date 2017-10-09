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
	*rnt.Context
	_VALID_URL string
}

func NewSlutloadIE(ctx *rnt.Context) rnt.InfoExtractor {
	ret := &SlutloadIE{}
	ret.Context = ctx
	ret._VALID_URL = `^https?://(?:\w+\.)?slutload\.com/video/[^/]+/(?P<id>[^/]+)/?$`
	return ret
}

func (self *SlutloadIE) Ctx() *rnt.Context {
	return self.Context
}

func (self *SlutloadIE) Key() string {
	return "Slutload"
}

func (self *SlutloadIE) ValidUrl() string {
	return self._VALID_URL
}

func (self *SlutloadIE) Name() string {
	return `Slutload extractor`
}

func (self *SlutloadIE) Tests() []map[string]interface{} {
	return []map[string]interface{}{}
}

func (self *SlutloadIE) _real_extract(url string) map[string]interface{} {
	url = rnt.ReSub(rnt.Re, `^(https?://)mobile\.`, `\1`, url, 0, 0)
	video_id := rnt.MatchID(self, url)
	webpage := rnt.DownloadWebpage(self, url, video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, rnt.OptString{}, map[string]interface{}{}, map[string]interface{}{})
	video_title := rnt.StrStrip(rnt.HTMLSearchRegex(self, `<h1><strong>([^<]+)</strong>`, webpage, `title`, rnt.NoDefault, true, 0, nil).Get(), ``)
	video_url := rnt.HTMLSearchRegex(self, `(?s)<div id="vidPlayer"\s+data-url="([^"]+)"`, webpage, `video URL`, rnt.NoDefault, true, 0, nil)
	thumbnail := rnt.HTMLSearchRegex(self, `(?s)<div id="vidPlayer"\s+.*?previewer-file="([^"]+)"`, webpage, `thumbnail`, rnt.NoDefault, false, 0, nil)
	return map[string]interface{}{`id`: video_id,
		`url`:       video_url,
		`title`:     video_title,
		`thumbnail`: thumbnail,
		`age_limit`: 18}
}

func (self *SlutloadIE) Extract(url string) (*rnt.VideoResult, error) {
	return rnt.RunExtractor(url, self._real_extract)
}

func init() {
	registerFactory(`Slutload`, NewSlutloadIE)
}
