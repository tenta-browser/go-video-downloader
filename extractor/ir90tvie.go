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
 * ir90tvie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/ir90tv.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type Ir90TvIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewIr90TvIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &Ir90TvIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "Ir90Tv"
	_VALID_URL = "https?://(?:www\\.)?90tv\\.ir/video/(?P<id>[0-9]+)/.*"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *Ir90TvIE) Key() string {
	return "Ir90Tv"
}

func (self *Ir90TvIE) Name() string {
	return self.IE_NAME
}

func (self *Ir90TvIE) _real_extract(url string) rnt.SDict {
	var (
		thumbnail rnt.OptString
		title     rnt.OptString
		video_id  string
		video_url rnt.OptString
		webpage   string
	)
	video_id = (self).MatchID(url)
	webpage = (self).DownloadWebpageURL(url, video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	title = rnt.RemoveStart((self).HTMLSearchRegexOne("<title>([^<]+)</title>", webpage, "title", rnt.NoDefault, true, 0, nil), "90tv.ir :: ")
	video_url = (self).SearchRegexOne("<source[^>]+src=\"([^\"]+)\"", webpage, "video url", rnt.NoDefault, true, 0, nil)
	thumbnail = (self).SearchRegexOne("poster=\"([^\"]+)\"", webpage, "thumbnail url", rnt.NoDefault, false, 0, nil)
	return rnt.SDict{
		"url":       video_url,
		"id":        video_id,
		"title":     title,
		"video_url": video_url,
		"thumbnail": thumbnail,
	}
}

func (self *Ir90TvIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("Ir90Tv", NewIr90TvIE)
}
