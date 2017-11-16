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
 * bambuserie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/bambuser.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type BambuserIE struct {
	*rnt.CommonIE
	IE_NAME        string
	_API_KEY       string
	_LOGIN_URL     string
	_NETRC_MACHINE string
}

func NewBambuserIE() rnt.InfoExtractor {
	var (
		IE_NAME        string
		_API_KEY       string
		_LOGIN_URL     string
		_NETRC_MACHINE string
		_VALID_URL     string
	)
	self := &BambuserIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "Bambuser"
	IE_NAME = "bambuser"
	_VALID_URL = "https?://bambuser\\.com/v/(?P<id>\\d+)"
	_API_KEY = "005f64509e19a868399060af746a00aa"
	_LOGIN_URL = "https://bambuser.com/user"
	_NETRC_MACHINE = "bambuser"
	self.IE_NAME = IE_NAME
	self._API_KEY = _API_KEY
	self._LOGIN_URL = _LOGIN_URL
	self._NETRC_MACHINE = _NETRC_MACHINE
	self.VALIDURL = _VALID_URL
	return self
}

func (self *BambuserIE) Key() string {
	return "Bambuser"
}

func (self *BambuserIE) Name() string {
	return self.IE_NAME
}

func (self *BambuserIE) _real_extract(url string) rnt.SDict {
	var (
		error    interface{}
		info     interface{}
		result   interface{}
		video_id string
	)
	video_id = (self).MatchID(url)
	info = (self).DownloadJSON(rnt.StrFormat2("http://player-c.api.bambuser.com/getVideo.json?api_key=%s&vid=%s", (self)._API_KEY, video_id), video_id, rnt.AsOptString("Downloading JSON metadata"), rnt.AsOptString("Unable to download JSON metadata"), nil, true, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	error = rnt.DictGet(τ_cast_α_to_d(info), "error", nil)
	if rnt.IsTruthy(error) {
		panic(rnt.PyExtractorError(rnt.StrFormat2("%s returned error: %s", (self).IE_NAME, error), true, rnt.OptString{}))
	}
	result = rnt.UnsafeSubscript(info, "result")
	return rnt.SDict{
		"id":            video_id,
		"title":         rnt.UnsafeSubscript(result, "title"),
		"url":           rnt.UnsafeSubscript(result, "url"),
		"thumbnail":     rnt.DictGet(τ_cast_α_to_d(result), "preview", nil),
		"duration":      rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(result), "length", nil), 1, rnt.OptInt{}, 1),
		"uploader":      rnt.DictGet(τ_cast_α_to_d(result), "username", nil),
		"uploader_id":   rnt.ConvertToString(rnt.DictGet(τ_cast_α_to_d(rnt.DictGet(τ_cast_α_to_d(result), "owner", rnt.SDict{})), "uid", nil)),
		"timestamp":     rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(result), "created", nil), 1, rnt.OptInt{}, 1),
		"fps":           rnt.FloatOrNone(rnt.DictGet(τ_cast_α_to_d(result), "framerate", nil), 1, 1, rnt.OptFloat{}),
		"view_count":    rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(result), "views_total", nil), 1, rnt.OptInt{}, 1),
		"comment_count": rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(result), "comment_count", nil), 1, rnt.OptInt{}, 1),
	}
}

func (self *BambuserIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("Bambuser", NewBambuserIE)
}
