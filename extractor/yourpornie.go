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
 * yourpornie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/yourporn.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type YourPornIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewYourPornIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &YourPornIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "YourPorn"
	_VALID_URL = "https?://(?:www\\.)?yourporn\\.sexy/post/(?P<id>[^/?#&.]+)"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *YourPornIE) Key() string {
	return "YourPorn"
}

func (self *YourPornIE) Name() string {
	return self.IE_NAME
}

func (self *YourPornIE) _real_extract(url string) rnt.SDict {
	var (
		thumbnail rnt.OptString
		title     string
		video_id  string
		video_url rnt.OptString
		webpage   string
	)
	video_id = (self).MatchID(url)
	webpage = (self).DownloadWebpageURL(url, video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{}, nil)
	video_url = rnt.UtilURLJoin(url, rnt.UnsafeSubscript((self).ParseJSON((self).SearchRegexOne("data-vnfo=([\"\\'])(?P<data>{.+?})\\1", webpage, "data info", rnt.NoDefault, true, 0, "data").Get(), video_id, nil, true), video_id))
	title = rnt.StrStrip(func() rnt.OptString {
		if v := ((self).SearchRegexOne("<[^>]+\\bclass=[\"\\']PostEditTA[^>]+>([^<]+)", webpage, "title", nil, true, 0, nil)); τ_isTruthy_Os(v) {
			return v
		} else {
			return (self).OgSearchDescription(webpage, rnt.NoDefault)
		}
	}().Get(), "")
	thumbnail = (self).OgSearchThumbnail(webpage, rnt.NoDefault)
	return rnt.SDict{
		"id":        video_id,
		"url":       video_url,
		"title":     title,
		"thumbnail": thumbnail,
	}
}

func (self *YourPornIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("YourPorn", NewYourPornIE)
}
