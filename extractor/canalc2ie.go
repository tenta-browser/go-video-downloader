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
 * canalc2ie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/canalc2.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type Canalc2IE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewCanalc2IE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &Canalc2IE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "Canalc2"
	IE_NAME = "canalc2.tv"
	_VALID_URL = "https?://(?:(?:www\\.)?canalc2\\.tv/video/|archives-canalc2\\.u-strasbg\\.fr/video\\.asp\\?.*\\bidVideo=)(?P<id>\\d+)"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *Canalc2IE) Key() string {
	return "Canalc2"
}

func (self *Canalc2IE) Name() string {
	return self.IE_NAME
}

func (self *Canalc2IE) _real_extract(url string) rnt.SDict {
	var (
		_         string
		duration  rnt.OptString
		formats   []interface{}
		rtmp      rnt.Match
		title     rnt.OptString
		video_id  string
		video_url string
		webpage   string
		τmp1      []string
	)
	video_id = (self).MatchID(url)
	webpage = (self).DownloadWebpageURL(rnt.StrFormat2("http://www.canalc2.tv/video/%s", video_id), video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	formats = []interface{}{}
	for _, τel := range rnt.ReFindAllMulti("file\\s*=\\s*([\"\\'])(.+?)\\1", webpage, 0) {
		τmp1 = τel
		_ = (τmp1)[0]
		video_url = (τmp1)[1]
		if rnt.StrStartswith(video_url, "rtmp://", rnt.OptInt{}, rnt.OptInt{}) {
			rtmp = rnt.ReSearch("^(?P<url>rtmp://[^/]+/(?P<app>.+/))(?P<play_path>mp4:.+)$", video_url, 0)
			formats = append(formats, rnt.SDict{
				"url":       rnt.ReMatchGroupOne(rtmp, "url"),
				"format_id": "rtmp",
				"ext":       "flv",
				"app":       rnt.ReMatchGroupOne(rtmp, "app"),
				"play_path": rnt.ReMatchGroupOne(rtmp, "play_path"),
				"page_url":  url,
			})
		} else {
			formats = append(formats, rnt.SDict{
				"url":       video_url,
				"format_id": "http",
			})
		}
	}
	formats = τ_conv_Ld_to_Lα((self).SortFormats(τ_conv_Lα_to_Ld(formats)))
	title = (self).HTMLSearchRegexOne("(?s)class=\"[^\"]*col_description[^\"]*\">.*?<h3>(.*?)</h3>", webpage, "title", rnt.NoDefault, true, 0, nil)
	duration = rnt.ParseDuration((self).SearchRegexOne("id=[\"\\']video_duree[\"\\'][^>]*>([^<]+)", webpage, "duration", rnt.NoDefault, false, 0, nil))
	return rnt.SDict{
		"id":       video_id,
		"title":    title,
		"duration": duration,
		"formats":  formats,
	}
}

func (self *Canalc2IE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("Canalc2", NewCanalc2IE)
}
