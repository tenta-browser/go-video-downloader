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
 * tvcarticleie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/tvc.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type TVCArticleIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewTVCArticleIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &TVCArticleIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "TVCArticle"
	_VALID_URL = "https?://(?:www\\.)?tvc\\.ru/(?!video/iframe/id/)(?P<id>[^?#]+)"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *TVCArticleIE) Key() string {
	return "TVCArticle"
}

func (self *TVCArticleIE) Name() string {
	return self.IE_NAME
}

func (self *TVCArticleIE) _real_extract(url string) rnt.SDict {
	var (
		webpage string
	)
	webpage = (self).DownloadWebpageURL(url, (self).MatchID(url), rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	return rnt.SDict{
		"_type":       "url_transparent",
		"ie_key":      "TVC",
		"url":         (self).OgSearchVideoURL(webpage, "video url", true, rnt.NoDefault),
		"title":       rnt.CleanHTML((self).OgSearchTitle(webpage, rnt.NoDefault, true)),
		"description": rnt.CleanHTML((self).OgSearchDescription(webpage, rnt.NoDefault)),
		"thumbnail":   (self).OgSearchThumbnail(webpage, rnt.NoDefault),
	}
}

func (self *TVCArticleIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("TVCArticle", NewTVCArticleIE)
}
