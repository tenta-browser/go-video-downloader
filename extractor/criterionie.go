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
	*rnt.CommonIE
	IE_NAME string
}

func NewCriterionIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &CriterionIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "Criterion"
	_VALID_URL = "https?://(?:www\\.)?criterion\\.com/films/(?P<id>[0-9]+)-.+"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *CriterionIE) Key() string {
	return "Criterion"
}

func (self *CriterionIE) Name() string {
	return self.IE_NAME
}

func (self *CriterionIE) _real_extract(url string) rnt.SDict {
	var (
		description rnt.OptString
		final_url   rnt.OptString
		thumbnail   rnt.OptString
		title       rnt.OptString
		video_id    string
		webpage     string
	)
	video_id = (self).MatchID(url)
	webpage = (self).DownloadWebpageURL(url, video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	final_url = (self).SearchRegexOne("so\\.addVariable\\(\"videoURL\", \"(.+?)\"\\)\\;", webpage, "video url", rnt.NoDefault, true, 0, nil)
	title = (self).OgSearchTitle(webpage, rnt.NoDefault, true)
	description = (self).HTMLSearchMetaOne("description", webpage, rnt.OptString{}, false, rnt.NoDefault, 0)
	thumbnail = (self).SearchRegexOne("so\\.addVariable\\(\"thumbnailURL\", \"(.+?)\"\\)\\;", webpage, "thumbnail url", rnt.NoDefault, true, 0, nil)
	return rnt.SDict{
		"id":          video_id,
		"url":         final_url,
		"title":       title,
		"description": description,
		"thumbnail":   thumbnail,
	}
}

func (self *CriterionIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("Criterion", NewCriterionIE)
}
