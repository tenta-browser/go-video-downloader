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
 * dotsubie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/dotsub.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type DotsubIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewDotsubIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &DotsubIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "Dotsub"
	_VALID_URL = "https?://(?:www\\.)?dotsub\\.com/view/(?P<id>[^/]+)"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *DotsubIE) Key() string {
	return "Dotsub"
}

func (self *DotsubIE) Name() string {
	return self.IE_NAME
}

func (self *DotsubIE) _real_extract(url string) rnt.SDict {
	var (
		info       interface{}
		info_dict  rnt.SDict
		setup_data interface{}
		video_id   string
		video_url  interface{}
		webpage    string
	)
	video_id = (self).MatchID(url)
	info = (self).DownloadJSON(rnt.StrFormat2("https://dotsub.com/api/media/%s/metadata", video_id), video_id, rnt.AsOptString("Downloading JSON metadata"), rnt.AsOptString("Unable to download JSON metadata"), nil, true, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	video_url = rnt.DictGet(τ_cast_α_to_d(info), "mediaURI", nil)
	if !(rnt.IsTruthy(video_url)) {
		webpage = (self).DownloadWebpageURL(url, video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
		video_url = (self).SearchRegexMulti([]string{"<source[^>]+src=\"([^\"]+)\"", "\"file\"\\s*:\\s*\\'([^\\']+)"}, webpage, "video url", nil, true, 0, nil)
		info_dict = rnt.SDict{
			"id":  video_id,
			"url": video_url,
			"ext": "flv",
		}
	}
	if !(rnt.IsTruthy(video_url)) {
		setup_data = (self).ParseJSON((self).HTMLSearchRegexOne("(?s)data-setup=([\\'\"])(?P<content>(?!\\1).+?)\\1", webpage, "setup data", rnt.NoDefault, true, 0, "content").Get(), video_id, nil, true)
		info_dict = rnt.SDict{
			"_type": "url_transparent",
			"url":   rnt.UnsafeSubscript(setup_data, "src"),
		}
	}
	rnt.DictUpdate(info_dict, rnt.SDict{
		"title":       rnt.UnsafeSubscript(info, "title"),
		"description": rnt.DictGet(τ_cast_α_to_d(info), "description", nil),
		"thumbnail":   rnt.DictGet(τ_cast_α_to_d(info), "screenshotURI", nil),
		"duration":    rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(info), "duration", nil), 1000, rnt.OptInt{}, 1),
		"uploader":    rnt.DictGet(τ_cast_α_to_d(info), "user", nil),
		"timestamp":   rnt.FloatOrNone(rnt.DictGet(τ_cast_α_to_d(info), "dateCreated", nil), 1000, 1, rnt.OptFloat{}),
		"view_count":  rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(info), "numberOfViews", nil), 1, rnt.OptInt{}, 1),
	})
	return info_dict
}

func (self *DotsubIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("Dotsub", NewDotsubIE)
}
