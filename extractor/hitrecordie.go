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
 * hitrecordie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/hitrecord.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type HitRecordIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewHitRecordIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &HitRecordIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "HitRecord"
	_VALID_URL = "https?://(?:www\\.)?hitrecord\\.org/records/(?P<id>\\d+)"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *HitRecordIE) Key() string {
	return "HitRecord"
}

func (self *HitRecordIE) Name() string {
	return self.IE_NAME
}

func (self *HitRecordIE) _real_extract(url string) rnt.SDict {
	var (
		tags      []interface{}
		tags_list interface{}
		title     interface{}
		video     interface{}
		video_id  string
		video_url interface{}
	)
	video_id = (self).MatchID(url)
	video = (self).DownloadJSON(rnt.StrFormat2("https://hitrecord.org/api/web/records/%s", video_id), video_id, rnt.AsOptString("Downloading JSON metadata"), rnt.AsOptString("Unable to download JSON metadata"), nil, true, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	title = rnt.UnsafeSubscript(video, "title")
	video_url = rnt.UnsafeSubscript(rnt.UnsafeSubscript(video, "source_url"), "mp4_url")
	tags = nil
	tags_list = rnt.TryGetOne(video, func(x interface{}) interface{} {
		return rnt.UnsafeSubscript(x, "tags")
	}, rnt.ConvertToList)
	if rnt.IsTruthy(tags_list) {
		tags = func() []interface{} {
			τresult := []interface{}{}
			for _, τel := range τ_cast_α_to_Lα(tags_list) {
				var (
					t interface{}
				)
				t = τel
				if !(rnt.IsTruthy(func() interface{} {
					if v := (τ_isinstance_d(t)); !(v) {
						return v
					} else if v := (rnt.DictGet(τ_cast_α_to_d(t), "text", nil)); !(rnt.IsTruthy(v)) {
						return v
					} else {
						return τ_isinstance_s(rnt.UnsafeSubscript(t, "text"))
					}
				}())) {
					continue
				}
				τresult = append(τresult, rnt.UnsafeSubscript(t, "text"))
			}
			return τresult
		}()
	}
	return rnt.SDict{
		"id":          video_id,
		"url":         video_url,
		"title":       title,
		"description": rnt.CleanHTML(rnt.CastToOptString(rnt.DictGet(τ_cast_α_to_d(video), "body", nil))),
		"duration":    rnt.FloatOrNone(rnt.DictGet(τ_cast_α_to_d(video), "duration", nil), 1000, 1, rnt.OptFloat{}),
		"timestamp":   rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(video), "created_at_i", nil), 1, rnt.OptInt{}, 1),
		"uploader": rnt.TryGetOne(video, func(x interface{}) interface{} {
			return rnt.UnsafeSubscript(rnt.UnsafeSubscript(x, "user"), "username")
		}, rnt.ConvertToString),
		"uploader_id": rnt.TryGetOne(video, func(x interface{}) interface{} {
			return rnt.ConvertToString(rnt.UnsafeSubscript(rnt.UnsafeSubscript(x, "user"), "id"))
		}, nil),
		"view_count":    rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(video), "total_views_count", nil), 1, rnt.OptInt{}, 1),
		"like_count":    rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(video), "hearts_count", nil), 1, rnt.OptInt{}, 1),
		"comment_count": rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(video), "comments_count", nil), 1, rnt.OptInt{}, 1),
		"tags":          tags,
	}
}

func (self *HitRecordIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("HitRecord", NewHitRecordIE)
}
