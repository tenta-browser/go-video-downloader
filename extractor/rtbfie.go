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
 * rtbfie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/rtbf.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type RTBFIE struct {
	*rnt.CommonIE
	IE_NAME     string
	_IMAGE_HOST string
	_PROVIDERS  rnt.SDict
	_QUALITIES  []τ_Tssω
}

func NewRTBFIE() rnt.InfoExtractor {
	var (
		IE_NAME     string
		_IMAGE_HOST string
		_PROVIDERS  rnt.SDict
		_QUALITIES  []τ_Tssω
		_VALID_URL  string
	)
	self := &RTBFIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "RTBF"
	_VALID_URL = "(?x)\n        https?://(?:www\\.)?rtbf\\.be/\n        (?:\n            video/[^?]+\\?.*\\bid=|\n            ouftivi/(?:[^/]+/)*[^?]+\\?.*\\bvideoId=|\n            auvio/[^/]+\\?.*id=\n        )(?P<id>\\d+)"
	_IMAGE_HOST = "http://ds1.ds.static.rtbf.be"
	_PROVIDERS = rnt.SDict{
		"YOUTUBE":     "Youtube",
		"DAILYMOTION": "Dailymotion",
		"VIMEO":       "Vimeo",
	}
	_QUALITIES = []τ_Tssω{{
		Φ0: "mobile",
		Φ1: "SD",
	}, {
		Φ0: "web",
		Φ1: "MD",
	}, {
		Φ0: "high",
		Φ1: "HD",
	}}
	self.IE_NAME = IE_NAME
	self._IMAGE_HOST = _IMAGE_HOST
	self._PROVIDERS = _PROVIDERS
	self._QUALITIES = _QUALITIES
	self.VALIDURL = _VALID_URL
	return self
}

func (self *RTBFIE) Key() string {
	return "RTBF"
}

func (self *RTBFIE) Name() string {
	return self.IE_NAME
}

func (self *RTBFIE) _real_extract(url string) rnt.SDict {
	var (
		data          interface{}
		error         interface{}
		format_id     string
		format_url    interface{}
		formats       []interface{}
		key           string
		provider      interface{}
		thumbnail_id  string
		thumbnail_url interface{}
		thumbnails    []interface{}
		video_id      string
		τmp1          τ_Tssω
		τmp2          τ_Tsαω
	)
	video_id = (self).MatchID(url)
	data = (self).DownloadJSON(rnt.StrFormat2("http://www.rtbf.be/api/media/video?method=getVideoDetail&args[]=%s", video_id), video_id, rnt.AsOptString("Downloading JSON metadata"), rnt.AsOptString("Unable to download JSON metadata"), nil, true, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	error = rnt.DictGet(τ_cast_α_to_d(data), "error", nil)
	if rnt.IsTruthy(error) {
		panic(rnt.PyExtractorError(rnt.StrFormat2("%s said: %s", (self).IE_NAME, error), true, rnt.OptString{}))
	}
	data = rnt.UnsafeSubscript(data, "data")
	provider = rnt.DictGet(τ_cast_α_to_d(data), "provider", nil)
	if rnt.DictContains((self)._PROVIDERS, rnt.CastToString(provider)) {
		return (self).URLResult(rnt.CastToString(rnt.UnsafeSubscript(data, "url")), rnt.CastToOptString(((self)._PROVIDERS)[rnt.CastToString(provider)]), rnt.OptString{}, rnt.OptString{})
	}
	formats = []interface{}{}
	for _, τel := range (self)._QUALITIES {
		τmp1 = τel
		key = (τmp1).Φ0
		format_id = (τmp1).Φ1
		format_url = rnt.DictGet(τ_cast_α_to_d(data), (key + "Url"), nil)
		if rnt.IsTruthy(format_url) {
			formats = append(formats, rnt.SDict{
				"format_id": format_id,
				"url":       format_url,
			})
		}
	}
	thumbnails = []interface{}{}
	for _, τel := range rnt.DictItems(τ_cast_α_to_d(rnt.DictGet(τ_cast_α_to_d(data), "thumbnail", rnt.SDict{}))) {
		τmp2 = τel
		thumbnail_id = (τmp2).Φ0
		thumbnail_url = (τmp2).Φ1
		if (thumbnail_id) != ("default") {
			thumbnails = append(thumbnails, rnt.SDict{
				"url": ((self)._IMAGE_HOST + rnt.CastToString(thumbnail_url)),
				"id":  thumbnail_id,
			})
		}
	}
	return rnt.SDict{
		"id":      video_id,
		"formats": formats,
		"title":   rnt.UnsafeSubscript(data, "title"),
		"description": func() interface{} {
			if v := (rnt.DictGet(τ_cast_α_to_d(data), "description", nil)); rnt.IsTruthy(v) {
				return v
			} else {
				return rnt.DictGet(τ_cast_α_to_d(data), "subtitle", nil)
			}
		}(),
		"thumbnails": thumbnails,
		"duration": func() interface{} {
			if v := (rnt.DictGet(τ_cast_α_to_d(data), "duration", nil)); rnt.IsTruthy(v) {
				return v
			} else {
				return rnt.DictGet(τ_cast_α_to_d(data), "realDuration", nil)
			}
		}(),
		"timestamp":  rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(data), "created", nil), 1, rnt.OptInt{}, 1),
		"view_count": rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(data), "viewCount", nil), 1, rnt.OptInt{}, 1),
		"uploader":   rnt.DictGet(τ_cast_α_to_d(data), "channel", nil),
		"tags":       rnt.DictGet(τ_cast_α_to_d(data), "tags", nil),
	}
}

func (self *RTBFIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("RTBF", NewRTBFIE)
}
