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
 * xxxymoviesie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/xxxymovies.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type XXXYMoviesIE struct {
	*rnt.Context
	_VALID_URL string
}

func NewXXXYMoviesIE(ctx *rnt.Context) rnt.InfoExtractor {
	ret := &XXXYMoviesIE{}
	ret.Context = ctx
	ret._VALID_URL = `https?://(?:www\.)?xxxymovies\.com/videos/(?P<id>\d+)/(?P<display_id>[^/]+)`
	return ret
}

func (self *XXXYMoviesIE) Ctx() *rnt.Context {
	return self.Context
}

func (self *XXXYMoviesIE) Key() string {
	return "XXXYMovies"
}

func (self *XXXYMoviesIE) ValidUrl() string {
	return self._VALID_URL
}

func (self *XXXYMoviesIE) Name() string {
	return `XXXYMovies extractor`
}

func (self *XXXYMoviesIE) Tests() []map[string]interface{} {
	return []map[string]interface{}{}
}

func (self *XXXYMoviesIE) _real_extract(url string) map[string]interface{} {
	mobj := rnt.ReMatch(rnt.Re, (self)._VALID_URL, url, 0)
	video_id := rnt.ReMatchGroupOne(mobj, `id`)
	display_id := rnt.ReMatchGroupOne(mobj, `display_id`)
	webpage := rnt.DownloadWebpage(self, url, display_id.Get(), rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, rnt.OptString{}, map[string]interface{}{}, map[string]interface{}{})
	video_url := rnt.SearchRegex(self, `video_url\s*:\s*'([^']+)'`, webpage, `video URL`, rnt.NoDefault, true, 0, nil)
	title := rnt.HTMLSearchRegex(self, []string{`<div[^>]+\bclass="block_header"[^>]*>\s*<h1>([^<]+)<`, `<title>(.*?)\s*-\s*(?:XXXYMovies\.com|XXX\s+Movies)</title>`}, webpage, `title`, rnt.NoDefault, true, 0, nil)
	thumbnail := rnt.SearchRegex(self, `preview_url\s*:\s*'([^']+)'`, webpage, `thumbnail`, rnt.NoDefault, false, 0, nil)
	categories := rnt.StrSplit(rnt.HTMLSearchMeta(self, `keywords`, webpage, rnt.AsOptString(`categories`), false, ``, 0).Get(), `,`, -(1))
	duration := rnt.ParseDuration(rnt.SearchRegex(self, `<span>Duration:</span>\s*(\d+:\d+)`, webpage, `duration`, rnt.NoDefault, false, 0, nil))
	view_count := rnt.IntOrNone(rnt.HTMLSearchRegex(self, `<div class="video_views">\s*(\d+)`, webpage, `view count`, rnt.NoDefault, false, 0, nil), 1, rnt.OptInt{}, 1)
	like_count := rnt.IntOrNone(rnt.SearchRegex(self, `>\s*Likes? <b>\((\d+)\)`, webpage, `like count`, rnt.NoDefault, false, 0, nil), 1, rnt.OptInt{}, 1)
	dislike_count := rnt.IntOrNone(rnt.SearchRegex(self, `>\s*Dislike <b>\((\d+)\)</b>`, webpage, `dislike count`, rnt.NoDefault, false, 0, nil), 1, rnt.OptInt{}, 1)
	age_limit := rnt.RTASearch(self, webpage)
	return map[string]interface{}{`id`: video_id,
		`display_id`:    display_id,
		`url`:           video_url,
		`title`:         title,
		`thumbnail`:     thumbnail,
		`categories`:    categories,
		`duration`:      duration,
		`view_count`:    view_count,
		`like_count`:    like_count,
		`dislike_count`: dislike_count,
		`age_limit`:     age_limit}
}

func (self *XXXYMoviesIE) Extract(url string) (*rnt.VideoResult, error) {
	return rnt.RunExtractor(url, self._real_extract)
}

func init() {
	registerFactory(`XXXYMovies`, NewXXXYMoviesIE)
}
