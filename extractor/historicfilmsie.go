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
 * historicfilmsie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/historicfilms.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type HistoricFilmsIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewHistoricFilmsIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &HistoricFilmsIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "HistoricFilms"
	_VALID_URL = "https?://(?:www\\.)?historicfilms\\.com/(?:tapes/|play)(?P<id>\\d+)"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *HistoricFilmsIE) Key() string {
	return "HistoricFilms"
}

func (self *HistoricFilmsIE) Name() string {
	return self.IE_NAME
}

func (self *HistoricFilmsIE) _real_extract(url string) rnt.SDict {
	var (
		description rnt.OptString
		duration    rnt.OptString
		tape_id     rnt.OptString
		thumbnail   rnt.OptString
		title       rnt.OptString
		video_id    string
		video_url   string
		webpage     string
	)
	video_id = (self).MatchID(url)
	webpage = (self).DownloadWebpageURL(url, video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	tape_id = (self).SearchRegexMulti([]string{"class=\"tapeId\"[^>]*>([^<]+)<", "tapeId\\s*:\\s*\"([^\"]+)\""}, webpage, "tape id", rnt.NoDefault, true, 0, nil)
	title = (self).OgSearchTitle(webpage, rnt.NoDefault, true)
	description = (self).OgSearchDescription(webpage, rnt.NoDefault)
	thumbnail = func() rnt.OptString {
		if v := ((self).HTMLSearchMetaOne("thumbnailUrl", webpage, rnt.AsOptString("thumbnails"), false, rnt.NoDefault, 0)); τ_isTruthy_Os(v) {
			return v
		} else {
			return (self).OgSearchThumbnail(webpage, rnt.NoDefault)
		}
	}()
	duration = rnt.ParseDuration((self).HTMLSearchMetaOne("duration", webpage, rnt.AsOptString("duration"), false, rnt.NoDefault, 0))
	video_url = rnt.StrFormat2("http://www.historicfilms.com/video/%s_%s_web.mov", tape_id, video_id)
	return rnt.SDict{
		"id":          video_id,
		"url":         video_url,
		"title":       title,
		"description": description,
		"thumbnail":   thumbnail,
		"duration":    duration,
	}
}

func (self *HistoricFilmsIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("HistoricFilms", NewHistoricFilmsIE)
}
