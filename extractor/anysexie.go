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
 * anysexie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/anysex.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type AnySexIE struct {
	*rnt.Context
	_VALID_URL string
}

func NewAnySexIE(ctx *rnt.Context) rnt.InfoExtractor {
	ret := &AnySexIE{}
	ret.Context = ctx
	ret._VALID_URL = `https?://(?:www\.)?anysex\.com/(?P<id>\d+)`
	return ret
}

func (self *AnySexIE) Ctx() *rnt.Context {
	return self.Context
}

func (self *AnySexIE) Key() string {
	return "AnySex"
}

func (self *AnySexIE) ValidUrl() string {
	return self._VALID_URL
}

func (self *AnySexIE) Name() string {
	return `AnySex extractor`
}

func (self *AnySexIE) Tests() []map[string]interface{} {
	return []map[string]interface{}{}
}

func (self *AnySexIE) _real_extract(url string) map[string]interface{} {
	mobj := rnt.ReMatch(rnt.Re, (self)._VALID_URL, url, 0)
	video_id := rnt.ReMatchGroupOne(mobj, `id`)
	webpage := rnt.DownloadWebpage(self, url, video_id.Get(), rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, rnt.OptString{}, map[string]interface{}{}, map[string]interface{}{})
	video_url := rnt.HTMLSearchRegex(self, `video_url\s*:\s*'([^']+)'`, webpage, `video URL`, rnt.NoDefault, true, 0, nil)
	title := rnt.HTMLSearchRegex(self, `<title>(.*?)</title>`, webpage, `title`, rnt.NoDefault, true, 0, nil)
	description := rnt.HTMLSearchRegex(self, `<div class="description"[^>]*>([^<]+)</div>`, webpage, `description`, rnt.NoDefault, false, 0, nil)
	thumbnail := rnt.HTMLSearchRegex(self, `preview_url\s*:\s*\'(.*?)\'`, webpage, `thumbnail`, rnt.NoDefault, false, 0, nil)
	categories := rnt.ReFindAllOne(rnt.Re, `<a href="http://anysex\.com/categories/[^"]+" title="[^"]*">([^<]+)</a>`, webpage, 0)
	duration := rnt.ParseDuration(rnt.SearchRegex(self, `<b>Duration:</b> (?:<q itemprop="duration">)?(\d+:\d+)`, webpage, `duration`, rnt.NoDefault, false, 0, nil))
	view_count := rnt.IntOrNone(rnt.HTMLSearchRegex(self, `<b>Views:</b> (\d+)`, webpage, `view count`, rnt.NoDefault, false, 0, nil), 1, rnt.OptInt{}, 1)
	return map[string]interface{}{`id`: video_id,
		`url`:         video_url,
		`ext`:         `mp4`,
		`title`:       title,
		`description`: description,
		`thumbnail`:   thumbnail,
		`categories`:  categories,
		`duration`:    duration,
		`view_count`:  view_count,
		`age_limit`:   18}
}

func (self *AnySexIE) Extract(url string) (*rnt.VideoResult, error) {
	return rnt.RunExtractor(url, self._real_extract)
}

func init() {
	registerFactory(`AnySex`, NewAnySexIE)
}
