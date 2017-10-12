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
	*rnt.CommonIE
}

func NewATTTechChannelIE() rnt.InfoExtractor {
	ret := &ATTTechChannelIE{}
	ret.CommonIE = rnt.NewCommonIE()
	ret.VALIDURL = `https?://techchannel\.att\.com/play-video\.cfm/([^/]+/)*(?P<id>.+)`
	return ret
}

func (self *ATTTechChannelIE) Key() string {
	return "ATTTechChannel"
}

func (self *ATTTechChannelIE) Name() string {
	return `ATTTechChannel extractor`
}

func (self *ATTTechChannelIE) _real_extract(url string) map[string]interface{} {
	display_id := (self).MatchID(url)
	webpage := (self).DownloadWebpage(url, display_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, rnt.OptString{}, map[string]interface{}{}, map[string]interface{}{})
	video_url := (self).SearchRegex(`url\s*:\s*'(rtmp://[^']+)'`, webpage, `video URL`, rnt.NoDefault, true, 0, nil)
	video_id := (self).SearchRegex(`mediaid\s*=\s*(\d+)`, webpage, `video id`, rnt.NoDefault, false, 0, nil)
	title := (self).OgSearchTitle(webpage, rnt.NoDefault, true)
	description := (self).OgSearchDescription(webpage, rnt.NoDefault)
	thumbnail := (self).OgSearchThumbnail(webpage, rnt.NoDefault)
	upload_date := rnt.UnifiedStrDate((self).SearchRegex(`[Rr]elease\s+date:\s*(\d{1,2}/\d{1,2}/\d{4})`, webpage, `upload date`, rnt.NoDefault, false, 0, nil), false)
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
