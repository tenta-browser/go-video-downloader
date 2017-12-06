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
 * streamableie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/streamable.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type StreamableIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewStreamableIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &StreamableIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "Streamable"
	_VALID_URL = "https?://streamable\\.com/(?:[es]/)?(?P<id>\\w+)"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *StreamableIE) Key() string {
	return "Streamable"
}

func (self *StreamableIE) Name() string {
	return self.IE_NAME
}

func (self *StreamableIE) _real_extract(url string) rnt.SDict {
	var (
		formats  []interface{}
		info     interface{}
		key      string
		status   interface{}
		title    interface{}
		video    interface{}
		video_id string
		τmp1     τ_Tsαω
	)
	video_id = (self).MatchID(url)
	video = (self).DownloadJSON(rnt.StrFormat2("https://ajax.streamable.com/videos/%s", video_id), video_id, rnt.AsOptString("Downloading JSON metadata"), rnt.AsOptString("Unable to download JSON metadata"), nil, true, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	status = rnt.DictGet(τ_cast_α_to_d(video), "status", nil)
	if (status) != (2) {
		panic(rnt.PyExtractorError("This video is currently unavailable. It may still be uploading or processing.", true, rnt.OptString{}))
	}
	title = func() interface{} {
		if v := (rnt.DictGet(τ_cast_α_to_d(video), "reddit_title", nil)); rnt.IsTruthy(v) {
			return v
		} else {
			return rnt.UnsafeSubscript(video, "title")
		}
	}()
	formats = []interface{}{}
	for _, τel := range rnt.DictItems(τ_cast_α_to_d(rnt.UnsafeSubscript(video, "files"))) {
		τmp1 = τel
		key = (τmp1).Φ0
		info = (τmp1).Φ1
		if !(rnt.IsTruthy(rnt.DictGet(τ_cast_α_to_d(info), "url", nil))) {
			continue
		}
		formats = append(formats, rnt.SDict{
			"format_id": key,
			"url":       (self).ProtoRelativeURL(rnt.CastToOptString(rnt.UnsafeSubscript(info, "url")), rnt.OptString{}),
			"width":     rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(info), "width", nil), 1, rnt.OptInt{}, 1),
			"height":    rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(info), "height", nil), 1, rnt.OptInt{}, 1),
			"filesize":  rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(info), "size", nil), 1, rnt.OptInt{}, 1),
			"fps":       rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(info), "framerate", nil), 1, rnt.OptInt{}, 1),
			"vbr":       rnt.FloatOrNone(rnt.DictGet(τ_cast_α_to_d(info), "bitrate", nil), 1000, 1, rnt.OptFloat{}),
		})
	}
	formats = τ_conv_Ld_to_Lα((self).SortFormats(τ_conv_Lα_to_Ld(formats)))
	return rnt.SDict{
		"id":          video_id,
		"title":       title,
		"description": rnt.DictGet(τ_cast_α_to_d(video), "description", nil),
		"thumbnail":   (self).ProtoRelativeURL(rnt.CastToOptString(rnt.DictGet(τ_cast_α_to_d(video), "thumbnail_url", nil)), rnt.OptString{}),
		"uploader":    rnt.DictGet(τ_cast_α_to_d(rnt.DictGet(τ_cast_α_to_d(video), "owner", rnt.SDict{})), "user_name", nil),
		"timestamp":   rnt.FloatOrNone(rnt.DictGet(τ_cast_α_to_d(video), "date_added", nil), 1, 1, rnt.OptFloat{}),
		"duration":    rnt.FloatOrNone(rnt.DictGet(τ_cast_α_to_d(video), "duration", nil), 1, 1, rnt.OptFloat{}),
		"view_count":  rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(video), "plays", nil), 1, rnt.OptInt{}, 1),
		"formats":     formats,
	}
}

func (self *StreamableIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("Streamable", NewStreamableIE)
}
