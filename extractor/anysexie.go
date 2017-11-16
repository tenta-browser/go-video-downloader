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
	*rnt.CommonIE
	IE_NAME string
}

func NewAnySexIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &AnySexIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "AnySex"
	_VALID_URL = "https?://(?:www\\.)?anysex\\.com/(?P<id>\\d+)"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *AnySexIE) Key() string {
	return "AnySex"
}

func (self *AnySexIE) Name() string {
	return self.IE_NAME
}

func (self *AnySexIE) _real_extract(url string) rnt.SDict {
	var (
		categories  []string
		description rnt.OptString
		duration    rnt.OptString
		mobj        rnt.Match
		thumbnail   rnt.OptString
		title       rnt.OptString
		video_id    rnt.OptString
		video_url   rnt.OptString
		view_count  rnt.OptInt
		webpage     string
	)
	mobj = rnt.ReMatch((self).VALIDURL, url, 0)
	video_id = rnt.ReMatchGroupOne(mobj, "id")
	webpage = (self).DownloadWebpageURL(url, video_id.Get(), rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	video_url = (self).HTMLSearchRegexOne("video_url\\s*:\\s*'([^']+)'", webpage, "video URL", rnt.NoDefault, true, 0, nil)
	title = (self).HTMLSearchRegexOne("<title>(.*?)</title>", webpage, "title", rnt.NoDefault, true, 0, nil)
	description = (self).HTMLSearchRegexOne("<div class=\"description\"[^>]*>([^<]+)</div>", webpage, "description", rnt.NoDefault, false, 0, nil)
	thumbnail = (self).HTMLSearchRegexOne("preview_url\\s*:\\s*\\'(.*?)\\'", webpage, "thumbnail", rnt.NoDefault, false, 0, nil)
	categories = rnt.ReFindAllOne("<a href=\"http://anysex\\.com/categories/[^\"]+\" title=\"[^\"]*\">([^<]+)</a>", webpage, 0)
	duration = rnt.ParseDuration((self).SearchRegexOne("<b>Duration:</b> (?:<q itemprop=\"duration\">)?(\\d+:\\d+)", webpage, "duration", rnt.NoDefault, false, 0, nil))
	view_count = rnt.IntOrNone((self).HTMLSearchRegexOne("<b>Views:</b> (\\d+)", webpage, "view count", rnt.NoDefault, false, 0, nil), 1, rnt.OptInt{}, 1)
	return rnt.SDict{
		"id":          video_id,
		"url":         video_url,
		"ext":         "mp4",
		"title":       title,
		"description": description,
		"thumbnail":   thumbnail,
		"categories":  categories,
		"duration":    duration,
		"view_count":  view_count,
		"age_limit":   18,
	}
}

func (self *AnySexIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("AnySex", NewAnySexIE)
}
