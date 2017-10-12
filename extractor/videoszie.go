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
	*rnt.CommonIE
}

func NewVideosZIE() rnt.InfoExtractor {
	ret := &VideosZIE{}
	ret.CommonIE = rnt.NewCommonIE()
	ret.VALIDURL = `https?://(?:www\.|m\.|)videosz\.com/[a-z]+/[a-z]+/(?P<id>[0-9]+)_`
	return ret
}

func (self *VideosZIE) Key() string {
	return "VideosZ"
}

func (self *VideosZIE) Name() string {
	return `VideosZ extractor`
}

func (self *VideosZIE) _real_extract(url string) map[string]interface{} {
	url = rnt.ReSub(`^(https?://(?:.+?\.)?)m\.`, `\1`, url, 0, 0)
	video_id := (self).MatchID(url)
	webpage := (self).DownloadWebpage(url, video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, rnt.OptString{}, map[string]interface{}{}, map[string]interface{}{})
	if rnt.ReSearch(`<form[^>]*id="login_form"[^>]*>`, webpage, 0) != nil {
		panic(rnt.PyExtractorError(`LOGIN_REQUIRED`, true, rnt.OptString{}))
	}
	title := (self).HTMLSearchMeta(`title`, webpage, rnt.AsOptString(`title`), false, rnt.NoDefault, 0)
	title = rnt.AsOptString(rnt.ReSub(`([^:]+):.+`, `\1`, title.Get(), 0, 0))
	formats_source := (self).SearchRegex(`<div[^>]*id="download_options_btn_content"[^>]*>(.+?)</div>`, webpage, `formats_source`, rnt.NoDefault, true, rnt.ReFlagDotAll, nil)
	formats := []interface{}{}
	format_matches := rnt.ReFindAllMulti(`<a[^>]*href="(?P<url>[^"]+)"[^>]*>(?P<format>[^(]+) \([^)]+\)</a>`, formats_source.Get(), 0)
	for _, τmp1 := range format_matches {
		τmp2 := τmp1
		format_url := (τmp2)[0]
		format := (τmp2)[1]
		formats = append(formats, map[string]interface{}{`format_id`: format,
			`url`: rnt.URLJoin(url, rnt.AsOptString(format_url))})
	}
	(self).SortFormats(formats)
	return map[string]interface{}{`id`: video_id,
		`title`:       title,
		`description`: title,
		`age_limit`:   (self).MediaRatingSearch(webpage),
		`formats`:     formats}
}

func (self *VideosZIE) Extract(url string) (*rnt.VideoResult, error) {
	return rnt.RunExtractor(url, self._real_extract)
}

func init() {
	registerFactory(`VideosZ`, NewVideosZIE)
}
