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
 * videoszie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/videosz.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type VideosZIE struct {
	*rnt.Context
	_VALID_URL string
}

func NewVideosZIE(ctx *rnt.Context) rnt.InfoExtractor {
	ret := &VideosZIE{}
	ret.Context = ctx
	ret._VALID_URL = `https?://(?:www\.|m\.|)videosz\.com/[a-z]+/[a-z]+/(?P<id>[0-9]+)_`
	return ret
}

func (self *VideosZIE) Ctx() *rnt.Context {
	return self.Context
}

func (self *VideosZIE) Key() string {
	return "VideosZ"
}

func (self *VideosZIE) ValidUrl() string {
	return self._VALID_URL
}

func (self *VideosZIE) Name() string {
	return `VideosZ extractor`
}

func (self *VideosZIE) Tests() []map[string]interface{} {
	return []map[string]interface{}{}
}

func (self *VideosZIE) _real_extract(url string) map[string]interface{} {
	url = rnt.ReSub(rnt.Re, `^(https?://(?:.+?\.)?)m\.`, `\1`, url, 0, 0)
	video_id := rnt.MatchID(self, url)
	webpage := rnt.DownloadWebpage(self, url, video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, rnt.OptString{}, map[string]interface{}{}, map[string]interface{}{})
	if rnt.ReSearch(rnt.Re, `<form[^>]*id="login_form"[^>]*>`, webpage, 0) != nil {
		panic(rnt.PyExtractorError(`LOGIN_REQUIRED`, true, rnt.OptString{}))
	}
	title := rnt.HTMLSearchMeta(self, `title`, webpage, rnt.AsOptString(`title`), false, rnt.NoDefault, 0)
	title = rnt.AsOptString(rnt.ReSub(rnt.Re, `([^:]+):.+`, `\1`, title.Get(), 0, 0))
	formats_source := rnt.SearchRegex(self, `<div[^>]*id="download_options_btn_content"[^>]*>(.+?)</div>`, webpage, `formats_source`, rnt.NoDefault, true, rnt.ReFlagDotAll, nil)
	formats := []interface{}{}
	format_matches := rnt.ReFindAllMulti(rnt.Re, `<a[^>]*href="(?P<url>[^"]+)"[^>]*>(?P<format>[^(]+) \([^)]+\)</a>`, formats_source.Get(), 0)
	for _, τmp1 := range format_matches {
		τmp2 := τmp1
		format_url := (τmp2)[0]
		format := (τmp2)[1]
		formats = append(formats, map[string]interface{}{`format_id`: format,
			`url`: rnt.URLJoin(rnt.UP, url, rnt.AsOptString(format_url))})
	}
	rnt.SortFormats(self, formats)
	return map[string]interface{}{`id`: video_id,
		`title`:       title,
		`description`: title,
		`age_limit`:   rnt.MediaRatingSearch(self, webpage),
		`formats`:     formats}
}

func (self *VideosZIE) Extract(url string) (*rnt.VideoResult, error) {
	return rnt.RunExtractor(url, self._real_extract)
}

func init() {
	registerFactory(`VideosZ`, NewVideosZIE)
}
