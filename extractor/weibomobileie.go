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
 * weibomobileie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/weibo.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type WeiboMobileIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewWeiboMobileIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &WeiboMobileIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "WeiboMobile"
	_VALID_URL = "https?://m\\.weibo\\.cn/status/(?P<id>[0-9]+)(\\?.+)?"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *WeiboMobileIE) Key() string {
	return "WeiboMobile"
}

func (self *WeiboMobileIE) Name() string {
	return self.IE_NAME
}

func (self *WeiboMobileIE) _real_extract(url string) rnt.SDict {
	var (
		page_info   interface{}
		status_data interface{}
		title       interface{}
		uploader    interface{}
		video_id    string
		webpage     string
		weibo_info  interface{}
	)
	video_id = (self).MatchID(url)
	webpage = (self).DownloadWebpageURL(url, video_id, rnt.AsOptString("visit the page"), rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	weibo_info = (self).ParseJSON((self).SearchRegexOne("var\\s+\\$render_data\\s*=\\s*\\[({.*})\\]\\[0\\]\\s*\\|\\|\\s*{};", webpage, "js_code", rnt.NoDefault, true, rnt.ReFlagDotAll, nil).Get(), video_id, rnt.JsToJSON, true)
	status_data = rnt.DictGet(τ_cast_α_to_d(weibo_info), "status", rnt.SDict{})
	page_info = rnt.DictGet(τ_cast_α_to_d(status_data), "page_info", nil)
	title = rnt.UnsafeSubscript(status_data, "status_title")
	uploader = rnt.DictGet(τ_cast_α_to_d(rnt.DictGet(τ_cast_α_to_d(status_data), "user", rnt.SDict{})), "screen_name", nil)
	return rnt.SDict{
		"id":       video_id,
		"title":    title,
		"uploader": uploader,
		"url":      rnt.UnsafeSubscript(rnt.UnsafeSubscript(page_info, "media_info"), "stream_url"),
	}
}

func (self *WeiboMobileIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("WeiboMobile", NewWeiboMobileIE)
}
