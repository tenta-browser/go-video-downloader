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
 * criterionie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/criterion.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type CriterionIE struct {
	*rnt.Context
	_VALID_URL string
}

func NewCriterionIE(ctx *rnt.Context) rnt.InfoExtractor {
	ret := &CriterionIE{}
	ret.Context = ctx
	ret._VALID_URL = `https?://(?:www\.)?criterion\.com/films/(?P<id>[0-9]+)-.+`
	return ret
}

func (self *CriterionIE) Ctx() *rnt.Context {
	return self.Context
}

func (self *CriterionIE) Key() string {
	return "Criterion"
}

func (self *CriterionIE) ValidUrl() string {
	return self._VALID_URL
}

func (self *CriterionIE) Name() string {
	return `Criterion extractor`
}

func (self *CriterionIE) Tests() []map[string]interface{} {
	return []map[string]interface{}{}
}

func (self *CriterionIE) _real_extract(url string) map[string]interface{} {
	video_id := rnt.MatchID(self, url)
	webpage := rnt.DownloadWebpage(self, url, video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, rnt.OptString{}, map[string]interface{}{}, map[string]interface{}{})
	final_url := rnt.SearchRegex(self, `so\.addVariable\("videoURL", "(.+?)"\)\;`, webpage, `video url`, rnt.NoDefault, true, 0, nil)
	title := rnt.OgSearchTitle(self, webpage, rnt.NoDefault, true)
	description := rnt.HTMLSearchMeta(self, `description`, webpage, rnt.OptString{}, false, rnt.NoDefault, 0)
	thumbnail := rnt.SearchRegex(self, `so\.addVariable\("thumbnailURL", "(.+?)"\)\;`, webpage, `thumbnail url`, rnt.NoDefault, true, 0, nil)
	return map[string]interface{}{`id`: video_id,
		`url`:         final_url,
		`title`:       title,
		`description`: description,
		`thumbnail`:   thumbnail}
}

func (self *CriterionIE) Extract(url string) (*rnt.VideoResult, error) {
	return rnt.RunExtractor(url, self._real_extract)
}

func init() {
	registerFactory(`Criterion`, NewCriterionIE)
}
