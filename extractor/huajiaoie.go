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
 * huajiaoie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/huajiao.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type HuajiaoIE struct {
	*rnt.CommonIE
	IE_DESC string
	IE_NAME string
}

func NewHuajiaoIE() rnt.InfoExtractor {
	var (
		IE_DESC    string
		IE_NAME    string
		_VALID_URL string
	)
	self := &HuajiaoIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "Huajiao"
	IE_DESC = "花椒直播"
	_VALID_URL = "https?://(?:www\\.)?huajiao\\.com/l/(?P<id>[0-9]+)"
	self.IE_DESC = IE_DESC
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *HuajiaoIE) Key() string {
	return "Huajiao"
}

func (self *HuajiaoIE) Name() string {
	return self.IE_NAME
}

func (self *HuajiaoIE) _real_extract(url string) rnt.SDict {
	var (
		description rnt.OptString
		feed        interface{}
		feed_json   rnt.OptString
		get         func(string, string) interface{}
		video_id    string
		webpage     string
	)
	video_id = (self).MatchID(url)
	webpage = (self).DownloadWebpageURL(url, video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	feed_json = (self).SearchRegexOne("var\\s+feed\\s*=\\s*({.+})", webpage, "feed json", rnt.NoDefault, true, 0, nil)
	feed = (self).ParseJSON(feed_json.Get(), video_id, nil, true)
	description = (self).HTMLSearchMetaOne("description", webpage, rnt.AsOptString("description"), false, rnt.NoDefault, 0)
	get = func(section string, field string) interface{} {
		return rnt.DictGet(τ_cast_α_to_d(rnt.DictGet(τ_cast_α_to_d(feed), section, rnt.SDict{})), field, nil)
	}
	return rnt.SDict{
		"id":          video_id,
		"title":       rnt.UnsafeSubscript(rnt.UnsafeSubscript(feed, "feed"), "formated_title"),
		"description": description,
		"duration":    rnt.ParseDuration(rnt.CastToOptString(get("feed", "duration"))),
		"thumbnail":   get("feed", "image"),
		"timestamp":   rnt.ParseISO8601(rnt.CastToOptString(rnt.DictGet(τ_cast_α_to_d(feed), "creatime", nil)), " "),
		"uploader":    get("author", "nickname"),
		"uploader_id": get("author", "uid"),
		"formats":     (self).ExtractM3U8Formats(rnt.CastToString(rnt.UnsafeSubscript(rnt.UnsafeSubscript(feed, "feed"), "m3u8")), video_id, rnt.AsOptString("mp4"), rnt.AsOptString("m3u8_native"), rnt.OptFloat{}, rnt.OptString{}, rnt.OptString{}, rnt.OptString{}, true, false),
	}
}

func (self *HuajiaoIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("Huajiao", NewHuajiaoIE)
}
