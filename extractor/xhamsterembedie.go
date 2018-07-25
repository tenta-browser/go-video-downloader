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
 * xhamsterembedie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/xhamster.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type XHamsterEmbedIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewXHamsterEmbedIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &XHamsterEmbedIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "XHamsterEmbed"
	_VALID_URL = "https?://(?:.+?\\.)?xhamster\\.com/xembed\\.php\\?video=(?P<id>\\d+)"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *XHamsterEmbedIE) Key() string {
	return "XHamsterEmbed"
}

func (self *XHamsterEmbedIE) Name() string {
	return self.IE_NAME
}

func (self *XHamsterEmbedIE) _real_extract(url string) rnt.SDict {
	var (
		vars      interface{}
		video_id  string
		video_url interface{}
		webpage   string
	)
	video_id = (self).MatchID(url)
	webpage = (self).DownloadWebpageURL(url, video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{}, nil)
	video_url = (self).SearchRegexOne(rnt.StrFormat3("href=\"(https?://xhamster\\.com/(?:movies/{0}/[^\"]*\\.html|videos/[^/]*-{0})[^\"]*)\"", video_id), webpage, "xhamster url", nil, true, 0, nil)
	if !(rnt.IsTruthy(video_url)) {
		vars = (self).ParseJSON((self).SearchRegexOne("vars\\s*:\\s*({.+?})\\s*,\\s*\\n", webpage, "vars", rnt.NoDefault, true, 0, nil).Get(), video_id, nil, true)
		video_url = rnt.UtilDictGet(τ_cast_α_to_d(vars), τ_Tssssω{
			Φ0: "downloadLink",
			Φ1: "homepageLink",
			Φ2: "commentsLink",
			Φ3: "shareUrl",
		}, nil, true)
	}
	return (self).URLResult(rnt.CastToString(video_url), rnt.AsOptString("XHamster"), rnt.OptString{}, rnt.OptString{})
}

func (self *XHamsterEmbedIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("XHamsterEmbed", NewXHamsterEmbedIE)
}
