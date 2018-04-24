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
 * rentvie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/rentv.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type RENTVIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewRENTVIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &RENTVIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "RENTV"
	_VALID_URL = "(?:rentv:|https?://(?:www\\.)?ren\\.tv/(?:player|video/epizod)/)(?P<id>\\d+)"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *RENTVIE) Key() string {
	return "RENTV"
}

func (self *RENTVIE) Name() string {
	return self.IE_NAME
}

func (self *RENTVIE) _real_extract(url string) rnt.SDict {
	var (
		config   interface{}
		ext      string
		formats  []interface{}
		src      interface{}
		title    interface{}
		video    interface{}
		video_id string
		webpage  string
	)
	video_id = (self).MatchID(url)
	webpage = (self).DownloadWebpageURL(("http://ren.tv/player/" + video_id), video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	config = (self).ParseJSON((self).SearchRegexOne("config\\s*=\\s*({.+})\\s*;", webpage, "config", rnt.NoDefault, true, 0, nil).Get(), video_id, nil, true)
	title = rnt.UnsafeSubscript(config, "title")
	formats = []interface{}{}
	for _, τel := range τ_cast_α_to_Lα(rnt.UnsafeSubscript(config, "src")) {
		video = τel
		src = rnt.DictGet(τ_cast_α_to_d(video), "src", nil)
		if !(rnt.IsTruthy(src)) || !(τ_isinstance_s(src)) {
			continue
		}
		ext = rnt.DetermineExt(rnt.CastToOptString(src), "unknown_video")
		if (ext) == ("m3u8") {
			formats = append(formats, τ_conv_Ld_to_Lα((self).ExtractM3U8Formats(rnt.CastToString(src), video_id, rnt.AsOptString("mp4"), rnt.AsOptString("m3u8_native"), nil, rnt.AsOptString("hls"), rnt.OptString{}, rnt.OptString{}, false, false))...)
		} else {
			formats = append(formats, rnt.SDict{
				"url": src,
			})
		}
	}
	formats = τ_conv_Ld_to_Lα((self).SortFormats(τ_conv_Lα_to_Ld(formats)))
	return rnt.SDict{
		"id":          video_id,
		"title":       title,
		"description": rnt.DictGet(τ_cast_α_to_d(config), "description", nil),
		"thumbnail":   rnt.DictGet(τ_cast_α_to_d(config), "image", nil),
		"duration":    rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(config), "duration", nil), 1, rnt.OptInt{}, 1),
		"timestamp":   rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(config), "date", nil), 1, rnt.OptInt{}, 1),
		"formats":     formats,
	}
}

func (self *RENTVIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("RENTV", NewRENTVIE)
}
