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
 * patreonie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/patreon.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type PatreonIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewPatreonIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &PatreonIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "Patreon"
	_VALID_URL = "https?://(?:www\\.)?patreon\\.com/(?:creation\\?hid=|posts/(?:[\\w-]+-)?)(?P<id>\\d+)"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *PatreonIE) Key() string {
	return "Patreon"
}

func (self *PatreonIE) Name() string {
	return self.IE_NAME
}

func (self *PatreonIE) _real_extract(url string) rnt.SDict {
	var (
		add_file        func(interface{}) interface{}
		attributes      interface{}
		i               interface{}
		i_type          interface{}
		image           interface{}
		info            rnt.SDict
		post            interface{}
		title           string
		user_attributes interface{}
		video_id        string
	)
	video_id = (self).MatchID(url)
	post = (self).DownloadJSON(("https://www.patreon.com/api/posts/" + video_id), video_id, rnt.AsOptString("Downloading JSON metadata"), rnt.AsOptString("Unable to download JSON metadata"), nil, true, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{}, nil)
	attributes = rnt.UnsafeSubscript(rnt.UnsafeSubscript(post, "data"), "attributes")
	title = rnt.StrStrip(rnt.CastToString(rnt.UnsafeSubscript(attributes, "title")), "")
	image = func() interface{} {
		if v := (rnt.DictGet(τ_cast_α_to_d(attributes), "image", nil)); rnt.IsTruthy(v) {
			return v
		} else {
			return rnt.SDict{}
		}
	}()
	info = rnt.SDict{
		"id":          video_id,
		"title":       title,
		"description": rnt.CleanHTML(rnt.CastToOptString(rnt.DictGet(τ_cast_α_to_d(attributes), "content", nil))),
		"thumbnail": func() interface{} {
			if v := (rnt.DictGet(τ_cast_α_to_d(image), "large_url", nil)); rnt.IsTruthy(v) {
				return v
			} else {
				return rnt.DictGet(τ_cast_α_to_d(image), "url", nil)
			}
		}(),
		"timestamp":     rnt.ParseISO8601(rnt.CastToOptString(rnt.DictGet(τ_cast_α_to_d(attributes), "published_at", nil)), "T"),
		"like_count":    rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(attributes), "like_count", nil), 1, rnt.OptInt{}, 1),
		"comment_count": rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(attributes), "comment_count", nil), 1, rnt.OptInt{}, 1),
	}
	add_file = func(file_data interface{}) interface{} {
		var (
			file_url interface{}
		)
		file_url = rnt.DictGet(τ_cast_α_to_d(file_data), "url", nil)
		if rnt.IsTruthy(file_url) {
			rnt.DictUpdate(info, rnt.SDict{
				"url": file_url,
				"ext": rnt.DetermineExt(rnt.CastToOptString(rnt.DictGet(τ_cast_α_to_d(file_data), "name", nil)), "mp3"),
			})
		}
		return nil
	}
	for _, τel := range τ_cast_α_to_Lα(rnt.DictGet(τ_cast_α_to_d(post), "included", []interface{}{})) {
		i = τel
		i_type = rnt.DictGet(τ_cast_α_to_d(i), "type", nil)
		if (i_type) == ("attachment") {
			add_file(func() interface{} {
				if v := (rnt.DictGet(τ_cast_α_to_d(i), "attributes", nil)); rnt.IsTruthy(v) {
					return v
				} else {
					return rnt.SDict{}
				}
			}())
		} else {
			if (i_type) == ("user") {
				user_attributes = rnt.DictGet(τ_cast_α_to_d(i), "attributes", nil)
				if rnt.IsTruthy(user_attributes) {
					rnt.DictUpdate(info, rnt.SDict{
						"uploader":     rnt.DictGet(τ_cast_α_to_d(user_attributes), "full_name", nil),
						"uploader_url": rnt.DictGet(τ_cast_α_to_d(user_attributes), "url", nil),
					})
				}
			}
		}
	}
	if !(rnt.IsTruthy(rnt.DictGet(info, "url", nil))) {
		add_file(func() interface{} {
			if v := (rnt.DictGet(τ_cast_α_to_d(attributes), "post_file", nil)); rnt.IsTruthy(v) {
				return v
			} else {
				return rnt.SDict{}
			}
		}())
	}
	if !(rnt.IsTruthy(rnt.DictGet(info, "url", nil))) {
		rnt.DictUpdate(info, rnt.SDict{
			"_type": "url",
			"url":   rnt.UnsafeSubscript(rnt.UnsafeSubscript(attributes, "embed"), "url"),
		})
	}
	return info
}

func (self *PatreonIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("Patreon", NewPatreonIE)
}
