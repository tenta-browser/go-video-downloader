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
 * lemondeie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/lemonde.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type LemondeIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewLemondeIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &LemondeIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "Lemonde"
	_VALID_URL = "https?://(?:.+?\\.)?lemonde\\.fr/(?:[^/]+/)*(?P<id>[^/]+)\\.html"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *LemondeIE) Key() string {
	return "Lemonde"
}

func (self *LemondeIE) Name() string {
	return self.IE_NAME
}

func (self *LemondeIE) _real_extract(url string) rnt.SDict {
	var (
		digiteka_url rnt.OptString
		display_id   string
		webpage      string
	)
	display_id = (self).MatchID(url)
	webpage = (self).DownloadWebpageURL(url, display_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{}, nil)
	digiteka_url = (self).ProtoRelativeURL((self).SearchRegexOne("url\\s*:\\s*([\"\\'])(?P<url>(?:https?://)?//(?:www\\.)?(?:digiteka\\.net|ultimedia\\.com)/deliver/.+?)\\1", webpage, "digiteka url", nil, true, 0, "url"), rnt.OptString{})
	if τ_isTruthy_Os(digiteka_url) {
		return (self).URLResult(digiteka_url.Get(), rnt.AsOptString("Digiteka"), rnt.OptString{}, rnt.OptString{})
	}
	return (self).URLResult(url, rnt.AsOptString("Generic"), rnt.OptString{}, rnt.OptString{})
}

func (self *LemondeIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("Lemonde", NewLemondeIE)
}
