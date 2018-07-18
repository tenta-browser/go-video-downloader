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
 * tvcie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/tvc.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type TVCIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewTVCIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &TVCIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "TVC"
	_VALID_URL = "https?://(?:www\\.)?tvc\\.ru/video/iframe/id/(?P<id>\\d+)"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *TVCIE) Key() string {
	return "TVC"
}

func (self *TVCIE) Name() string {
	return self.IE_NAME
}

func (self *TVCIE) _real_extract(url string) rnt.SDict {
	var (
		format_id rnt.OptString
		formats   []interface{}
		info      interface{}
		video     interface{}
		video_id  string
		video_url interface{}
	)
	video_id = (self).MatchID(url)
	video = (self).DownloadJSON(rnt.StrFormat2("http://www.tvc.ru/video/json/id/%s", video_id), video_id, rnt.AsOptString("Downloading JSON metadata"), rnt.AsOptString("Unable to download JSON metadata"), nil, true, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{}, nil)
	formats = []interface{}{}
	for _, τel := range τ_cast_α_to_Lα(rnt.DictGet(τ_cast_α_to_d(rnt.DictGet(τ_cast_α_to_d(video), "path", rnt.SDict{})), "quality", []interface{}{})) {
		info = τel
		video_url = rnt.DictGet(τ_cast_α_to_d(info), "url", nil)
		if !(rnt.IsTruthy(video_url)) {
			continue
		}
		format_id = (self).SearchRegexOne("cdnvideo/([^/]+?)(?:-[^/]+?)?/", rnt.CastToString(video_url), "format id", nil, true, 0, nil)
		formats = append(formats, rnt.SDict{
			"url":       video_url,
			"format_id": format_id,
			"width":     rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(info), "width", nil), 1, rnt.OptInt{}, 1),
			"height":    rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(info), "height", nil), 1, rnt.OptInt{}, 1),
			"tbr":       rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(info), "bitrate", nil), 1, rnt.OptInt{}, 1),
		})
	}
	formats = τ_conv_Ld_to_Lα((self).SortFormats(τ_conv_Lα_to_Ld(formats)))
	return rnt.SDict{
		"id":        video_id,
		"title":     rnt.UnsafeSubscript(video, "title"),
		"thumbnail": rnt.DictGet(τ_cast_α_to_d(video), "picture", nil),
		"duration":  rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(video), "duration", nil), 1, rnt.OptInt{}, 1),
		"formats":   formats,
	}
}

func (self *TVCIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("TVC", NewTVCIE)
}
