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
 * playfmie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/playfm.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type PlayFMIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewPlayFMIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &PlayFMIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "PlayFM"
	IE_NAME = "play.fm"
	_VALID_URL = "https?://(?:www\\.)?play\\.fm/(?P<slug>(?:[^/]+/)+(?P<id>[^/]+))/?(?:$|[?#])"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *PlayFMIE) Key() string {
	return "PlayFM"
}

func (self *PlayFMIE) Name() string {
	return self.IE_NAME
}

func (self *PlayFMIE) _real_extract(url string) rnt.SDict {
	var (
		audio_url     interface{}
		categories    []interface{}
		comment_count rnt.OptInt
		description   interface{}
		duration      rnt.OptInt
		error         interface{}
		mobj          rnt.Match
		recordings    interface{}
		slug          rnt.OptString
		timestamp     rnt.OptInt
		title         interface{}
		uploader      interface{}
		uploader_id   string
		video_id      rnt.OptString
		view_count    rnt.OptInt
	)
	mobj = rnt.ReMatch((self).VALIDURL, url, 0)
	video_id = rnt.ReMatchGroupOne(mobj, "id")
	slug = rnt.ReMatchGroupOne(mobj, "slug")
	recordings = (self).DownloadJSON(rnt.StrFormat2("http://v2api.play.fm/recordings/slug/%s", slug), video_id.Get(), rnt.AsOptString("Downloading JSON metadata"), rnt.AsOptString("Unable to download JSON metadata"), nil, true, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	error = rnt.DictGet(τ_cast_α_to_d(recordings), "error", nil)
	if τ_isinstance_d(error) {
		panic(rnt.PyExtractorError(rnt.StrFormat2("%s returned error: %s", (self).IE_NAME, rnt.DictGet(τ_cast_α_to_d(error), "message", nil)), true, rnt.OptString{}))
	}
	audio_url = rnt.UnsafeSubscript(recordings, "audio")
	video_id = rnt.AsOptString(rnt.ConvertToString(func() interface{} {
		if v := (rnt.DictGet(τ_cast_α_to_d(recordings), "id", nil)); rnt.IsTruthy(v) {
			return v
		} else {
			return video_id
		}
	}()))
	title = rnt.UnsafeSubscript(recordings, "title")
	description = rnt.DictGet(τ_cast_α_to_d(recordings), "description", nil)
	duration = rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(recordings), "recordingDuration", nil), 1, rnt.OptInt{}, 1)
	timestamp = rnt.ParseISO8601(rnt.CastToOptString(rnt.DictGet(τ_cast_α_to_d(recordings), "created_at", nil)), "T")
	uploader = rnt.DictGet(τ_cast_α_to_d(rnt.DictGet(τ_cast_α_to_d(recordings), "page", rnt.SDict{})), "title", nil)
	uploader_id = rnt.ConvertToString(rnt.DictGet(τ_cast_α_to_d(rnt.DictGet(τ_cast_α_to_d(recordings), "page", rnt.SDict{})), "id", nil))
	view_count = rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(recordings), "playCount", nil), 1, rnt.OptInt{}, 1)
	comment_count = rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(recordings), "commentCount", nil), 1, rnt.OptInt{}, 1)
	categories = func() []interface{} {
		τresult := []interface{}{}
		for _, τel := range τ_cast_α_to_Lα(rnt.DictGet(τ_cast_α_to_d(recordings), "tags", []interface{}{})) {
			var (
				tag interface{}
			)
			tag = τel
			if !(rnt.IsTruthy(rnt.DictGet(τ_cast_α_to_d(tag), "name", nil))) {
				continue
			}
			τresult = append(τresult, rnt.UnsafeSubscript(tag, "name"))
		}
		return τresult
	}()
	return rnt.SDict{
		"id":            video_id,
		"url":           audio_url,
		"title":         title,
		"description":   description,
		"duration":      duration,
		"timestamp":     timestamp,
		"uploader":      uploader,
		"uploader_id":   uploader_id,
		"view_count":    view_count,
		"comment_count": comment_count,
		"categories":    categories,
	}
}

func (self *PlayFMIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("PlayFM", NewPlayFMIE)
}
