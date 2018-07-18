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
 * clypie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/clyp.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type ClypIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewClypIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &ClypIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "Clyp"
	_VALID_URL = "https?://(?:www\\.)?clyp\\.it/(?P<id>[a-z0-9]+)"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *ClypIE) Key() string {
	return "Clyp"
}

func (self *ClypIE) Name() string {
	return self.IE_NAME
}

func (self *ClypIE) _real_extract(url string) rnt.SDict {
	var (
		audio_id    string
		description interface{}
		duration    rnt.OptFloat
		ext         string
		format_id   string
		format_url  interface{}
		formats     []interface{}
		metadata    interface{}
		secure      string
		timestamp   rnt.OptInt
		title       interface{}
	)
	audio_id = (self).MatchID(url)
	metadata = (self).DownloadJSON(rnt.StrFormat2("https://api.clyp.it/%s", audio_id), audio_id, rnt.AsOptString("Downloading JSON metadata"), rnt.AsOptString("Unable to download JSON metadata"), nil, true, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{}, nil)
	formats = []interface{}{}
	for _, τel := range τ_conv_Tssω_to_Ls(τ_Tssω{
		Φ0: "",
		Φ1: "Secure",
	}) {
		secure = τel
		for _, τel := range τ_conv_Tssω_to_Ls(τ_Tssω{
			Φ0: "Ogg",
			Φ1: "Mp3",
		}) {
			ext = τel
			format_id = rnt.StrFormat2("%s%s", secure, ext)
			format_url = rnt.DictGet(τ_cast_α_to_d(metadata), rnt.StrFormat2("%sUrl", format_id), nil)
			if rnt.IsTruthy(format_url) {
				formats = append(formats, rnt.SDict{
					"url":       format_url,
					"format_id": format_id,
					"vcodec":    "none",
				})
			}
		}
	}
	formats = τ_conv_Ld_to_Lα((self).SortFormats(τ_conv_Lα_to_Ld(formats)))
	title = rnt.UnsafeSubscript(metadata, "Title")
	description = rnt.DictGet(τ_cast_α_to_d(metadata), "Description", nil)
	duration = rnt.FloatOrNone(rnt.DictGet(τ_cast_α_to_d(metadata), "Duration", nil), 1, 1, rnt.OptFloat{})
	timestamp = rnt.ParseISO8601(rnt.CastToOptString(rnt.DictGet(τ_cast_α_to_d(metadata), "DateCreated", nil)), "T")
	return rnt.SDict{
		"id":          audio_id,
		"title":       title,
		"description": description,
		"duration":    duration,
		"timestamp":   timestamp,
		"formats":     formats,
	}
}

func (self *ClypIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("Clyp", NewClypIE)
}
