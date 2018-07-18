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
 * pikselie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/piksel.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type PikselIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewPikselIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &PikselIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "Piksel"
	_VALID_URL = "https?://player\\.piksel\\.com/v/(?P<id>[a-z0-9]+)"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *PikselIE) Key() string {
	return "Piksel"
}

func (self *PikselIE) Name() string {
	return self.IE_NAME
}

func (self *PikselIE) _real_extract(url string) rnt.SDict {
	var (
		abr        rnt.OptInt
		app_token  rnt.OptString
		asset_file interface{}
		asset_type interface{}
		failure    interface{}
		format_id  []string
		formats    []interface{}
		http_url   interface{}
		m3u8_url   interface{}
		response   interface{}
		tbr        rnt.OptInt
		title      interface{}
		vbr        rnt.OptInt
		video_data interface{}
		video_id   string
		webpage    string
	)
	video_id = (self).MatchID(url)
	webpage = (self).DownloadWebpageURL(url, video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{}, nil)
	app_token = (self).SearchRegexMulti([]string{"clientAPI\\s*:\\s*\"([^\"]+)\"", "data-de-api-key\\s*=\\s*\"([^\"]+)\""}, webpage, "app token", rnt.NoDefault, true, 0, nil)
	response = rnt.UnsafeSubscript((self).DownloadJSON(rnt.StrFormat2("http://player.piksel.com/ws/ws_program/api/%s/mode/json/apiv/5", app_token), video_id, rnt.AsOptString("Downloading JSON metadata"), rnt.AsOptString("Unable to download JSON metadata"), nil, true, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{
		"v": video_id,
	}, nil), "response")
	failure = rnt.DictGet(τ_cast_α_to_d(response), "failure", nil)
	if rnt.IsTruthy(failure) {
		panic(rnt.PyExtractorError(rnt.CastToString(rnt.UnsafeSubscript(rnt.UnsafeSubscript(response, "failure"), "reason")), true, rnt.OptString{}))
	}
	video_data = rnt.UnsafeSubscript(rnt.UnsafeSubscript(rnt.UnsafeSubscript(response, "WsProgramResponse"), "program"), "asset")
	title = rnt.UnsafeSubscript(video_data, "title")
	formats = []interface{}{}
	m3u8_url = rnt.UtilDictGet(τ_cast_α_to_d(video_data), []string{"m3u8iPadURL", "ipadM3u8Url", "m3u8AndroidURL", "m3u8iPhoneURL", "iphoneM3u8Url"}, nil, true)
	if rnt.IsTruthy(m3u8_url) {
		formats = append(formats, τ_conv_Ld_to_Lα((self).ExtractM3U8Formats(rnt.CastToString(m3u8_url), video_id, rnt.AsOptString("mp4"), rnt.AsOptString("m3u8_native"), nil, rnt.AsOptString("hls"), rnt.OptString{}, rnt.OptString{}, false, false))...)
	}
	asset_type = rnt.UtilDictGet(τ_cast_α_to_d(video_data), []string{"assetType", "asset_type"}, nil, true)
	for _, τel := range τ_cast_α_to_Lα(rnt.DictGet(τ_cast_α_to_d(video_data), "assetFiles", []interface{}{})) {
		asset_file = τel
		http_url = rnt.DictGet(τ_cast_α_to_d(asset_file), "http_url", nil)
		if !(rnt.IsTruthy(http_url)) {
			continue
		}
		tbr = rnt.OptInt{}
		vbr = rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(asset_file), "videoBitrate", nil), 1024, rnt.OptInt{}, 1)
		abr = rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(asset_file), "audioBitrate", nil), 1024, rnt.OptInt{}, 1)
		if (asset_type) == ("video") {
			tbr = rnt.AsOptInt((vbr.Get() + abr.Get()))
		} else {
			if (asset_type) == ("audio") {
				tbr = abr
			}
		}
		format_id = []string{"http"}
		if τ_isTruthy_Oi(tbr) {
			format_id = append(format_id, rnt.ConvertToString(tbr))
		}
		formats = append(formats, rnt.SDict{
			"format_id": rnt.StrJoin("-", format_id),
			"url":       rnt.UnescapeHTML(rnt.CastToOptString(http_url)),
			"vbr":       vbr,
			"abr":       abr,
			"width":     rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(asset_file), "videoWidth", nil), 1, rnt.OptInt{}, 1),
			"height":    rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(asset_file), "videoHeight", nil), 1, rnt.OptInt{}, 1),
			"filesize":  rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(asset_file), "filesize", nil), 1, rnt.OptInt{}, 1),
			"tbr":       tbr,
		})
	}
	formats = τ_conv_Ld_to_Lα((self).SortFormats(τ_conv_Lα_to_Ld(formats)))
	return rnt.SDict{
		"id":          video_id,
		"title":       title,
		"description": rnt.DictGet(τ_cast_α_to_d(video_data), "description", nil),
		"thumbnail":   rnt.DictGet(τ_cast_α_to_d(video_data), "thumbnailUrl", nil),
		"timestamp":   rnt.ParseISO8601(rnt.CastToOptString(rnt.DictGet(τ_cast_α_to_d(video_data), "dateadd", nil)), "T"),
		"formats":     formats,
	}
}

func (self *PikselIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("Piksel", NewPikselIE)
}
