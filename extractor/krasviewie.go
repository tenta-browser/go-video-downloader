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
 * krasviewie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/krasview.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type KrasViewIE struct {
	*rnt.CommonIE
	IE_DESC string
	IE_NAME string
}

func NewKrasViewIE() rnt.InfoExtractor {
	var (
		IE_DESC    string
		IE_NAME    string
		_VALID_URL string
	)
	self := &KrasViewIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "KrasView"
	IE_DESC = "Красвью"
	_VALID_URL = "https?://krasview\\.ru/(?:video|embed)/(?P<id>\\d+)"
	self.IE_DESC = IE_DESC
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *KrasViewIE) Key() string {
	return "KrasView"
}

func (self *KrasViewIE) Name() string {
	return self.IE_NAME
}

func (self *KrasViewIE) _real_extract(url string) rnt.SDict {
	var (
		description rnt.OptString
		duration    rnt.OptInt
		flashvars   interface{}
		height      rnt.OptInt
		thumbnail   interface{}
		title       rnt.OptString
		video_id    string
		video_url   interface{}
		webpage     string
		width       rnt.OptInt
	)
	video_id = (self).MatchID(url)
	webpage = (self).DownloadWebpageURL(url, video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{}, nil)
	flashvars = rnt.JSONLoads(rnt.JsToJSON((self).SearchRegexOne("video_Init\\(({.+?})", webpage, "flashvars", rnt.NoDefault, true, 0, nil).Get()))
	video_url = rnt.UnsafeSubscript(flashvars, "url")
	title = (self).OgSearchTitle(webpage, rnt.NoDefault, true)
	description = (self).OgSearchDescription(webpage, nil)
	thumbnail = func() interface{} {
		if v := (rnt.DictGet(τ_cast_α_to_d(flashvars), "image", nil)); rnt.IsTruthy(v) {
			return v
		} else {
			return (self).OgSearchThumbnail(webpage, rnt.NoDefault)
		}
	}()
	duration = rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(flashvars), "duration", nil), 1, rnt.OptInt{}, 1)
	width = rnt.IntOrNone((self).OgSearchPropertyOne("video:width", webpage, rnt.AsOptString("video width"), nil, true), 1, rnt.OptInt{}, 1)
	height = rnt.IntOrNone((self).OgSearchPropertyOne("video:height", webpage, rnt.AsOptString("video height"), nil, true), 1, rnt.OptInt{}, 1)
	return rnt.SDict{
		"id":          video_id,
		"url":         video_url,
		"title":       title,
		"description": description,
		"thumbnail":   thumbnail,
		"duration":    duration,
		"width":       width,
		"height":      height,
	}
}

func (self *KrasViewIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("KrasView", NewKrasViewIE)
}
