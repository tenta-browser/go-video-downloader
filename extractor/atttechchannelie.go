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
 * atttechchannelie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/atttechchannel.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type ATTTechChannelIE struct {
	*rnt.Context
	_VALID_URL string
}

func NewATTTechChannelIE(ctx *rnt.Context) rnt.InfoExtractor {
	ret := &ATTTechChannelIE{}
	ret.Context = ctx
	ret._VALID_URL = `https?://techchannel\.att\.com/play-video\.cfm/([^/]+/)*(?P<id>.+)`
	return ret
}

func (self *ATTTechChannelIE) Ctx() *rnt.Context {
	return self.Context
}

func (self *ATTTechChannelIE) Key() string {
	return "ATTTechChannel"
}

func (self *ATTTechChannelIE) ValidUrl() string {
	return self._VALID_URL
}

func (self *ATTTechChannelIE) Name() string {
	return `ATTTechChannel extractor`
}

func (self *ATTTechChannelIE) Tests() []map[string]interface{} {
	return []map[string]interface{}{}
}

func (self *ATTTechChannelIE) _real_extract(url string) map[string]interface{} {
	display_id := rnt.MatchID(self, url)
	webpage := rnt.DownloadWebpage(self, url, display_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, rnt.OptString{}, map[string]interface{}{}, map[string]interface{}{})
	video_url := rnt.SearchRegex(self, `url\s*:\s*'(rtmp://[^']+)'`, webpage, `video URL`, rnt.NoDefault, true, 0, nil)
	video_id := rnt.SearchRegex(self, `mediaid\s*=\s*(\d+)`, webpage, `video id`, rnt.NoDefault, false, 0, nil)
	title := rnt.OgSearchTitle(self, webpage, rnt.NoDefault, true)
	description := rnt.OgSearchDescription(self, webpage, rnt.NoDefault)
	thumbnail := rnt.OgSearchThumbnail(self, webpage, rnt.NoDefault)
	upload_date := rnt.UnifiedStrDate(rnt.SearchRegex(self, `[Rr]elease\s+date:\s*(\d{1,2}/\d{1,2}/\d{4})`, webpage, `upload date`, rnt.NoDefault, false, 0, nil), false)
	return map[string]interface{}{`id`: video_id,
		`display_id`:  display_id,
		`url`:         video_url,
		`ext`:         `flv`,
		`title`:       title,
		`description`: description,
		`thumbnail`:   thumbnail,
		`upload_date`: upload_date}
}

func (self *ATTTechChannelIE) Extract(url string) (*rnt.VideoResult, error) {
	return rnt.RunExtractor(url, self._real_extract)
}

func init() {
	registerFactory(`ATTTechChannel`, NewATTTechChannelIE)
}
