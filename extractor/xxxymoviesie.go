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
	*rnt.CommonIE
	IE_NAME string
}

func NewXXXYMoviesIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &XXXYMoviesIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "XXXYMovies"
	_VALID_URL = "https?://(?:www\\.)?xxxymovies\\.com/videos/(?P<id>\\d+)/(?P<display_id>[^/]+)"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *XXXYMoviesIE) Key() string {
	return "XXXYMovies"
}

func (self *XXXYMoviesIE) Name() string {
	return self.IE_NAME
}

func (self *XXXYMoviesIE) _real_extract(url string) rnt.SDict {
	var (
		age_limit     int
		categories    []string
		dislike_count rnt.OptInt
		display_id    rnt.OptString
		duration      rnt.OptString
		like_count    rnt.OptInt
		mobj          rnt.Match
		thumbnail     rnt.OptString
		title         rnt.OptString
		video_id      rnt.OptString
		video_url     rnt.OptString
		view_count    rnt.OptInt
		webpage       string
	)
	mobj = rnt.ReMatch((self).VALIDURL, url, 0)
	video_id = rnt.ReMatchGroupOne(mobj, "id")
	display_id = rnt.ReMatchGroupOne(mobj, "display_id")
	webpage = (self).DownloadWebpageURL(url, display_id.Get(), rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	video_url = (self).SearchRegexOne("video_url\\s*:\\s*'([^']+)'", webpage, "video URL", rnt.NoDefault, true, 0, nil)
	title = (self).HTMLSearchRegexMulti([]string{"<div[^>]+\\bclass=\"block_header\"[^>]*>\\s*<h1>([^<]+)<", "<title>(.*?)\\s*-\\s*(?:XXXYMovies\\.com|XXX\\s+Movies)</title>"}, webpage, "title", rnt.NoDefault, true, 0, nil)
	thumbnail = (self).SearchRegexOne("preview_url\\s*:\\s*'([^']+)'", webpage, "thumbnail", rnt.NoDefault, false, 0, nil)
	categories = rnt.StrSplit((self).HTMLSearchMetaOne("keywords", webpage, rnt.AsOptString("categories"), false, "", 0).Get(), ",", -(1))
	duration = rnt.ParseDuration((self).SearchRegexOne("<span>Duration:</span>\\s*(\\d+:\\d+)", webpage, "duration", rnt.NoDefault, false, 0, nil))
	view_count = rnt.IntOrNone((self).HTMLSearchRegexOne("<div class=\"video_views\">\\s*(\\d+)", webpage, "view count", rnt.NoDefault, false, 0, nil), 1, rnt.OptInt{}, 1)
	like_count = rnt.IntOrNone((self).SearchRegexOne(">\\s*Likes? <b>\\((\\d+)\\)", webpage, "like count", rnt.NoDefault, false, 0, nil), 1, rnt.OptInt{}, 1)
	dislike_count = rnt.IntOrNone((self).SearchRegexOne(">\\s*Dislike <b>\\((\\d+)\\)</b>", webpage, "dislike count", rnt.NoDefault, false, 0, nil), 1, rnt.OptInt{}, 1)
	age_limit = (self).RTASearch(webpage)
	return rnt.SDict{
		"id":            video_id,
		"display_id":    display_id,
		"url":           video_url,
		"title":         title,
		"thumbnail":     thumbnail,
		"categories":    categories,
		"duration":      duration,
		"view_count":    view_count,
		"like_count":    like_count,
		"dislike_count": dislike_count,
		"age_limit":     age_limit,
	}
}

func (self *XXXYMoviesIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("XXXYMovies", NewXXXYMoviesIE)
}
