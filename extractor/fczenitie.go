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
 * fczenitie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/fczenit.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type FczenitIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewFczenitIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &FczenitIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "Fczenit"
	_VALID_URL = "https?://(?:www\\.)?fc-zenit\\.ru/video/(?P<id>[0-9]+)"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *FczenitIE) Key() string {
	return "Fczenit"
}

func (self *FczenitIE) Name() string {
	return self.IE_NAME
}

func (self *FczenitIE) _real_extract(url string) rnt.SDict {
	var (
		formats  []rnt.SDict
		msi_data interface{}
		msi_id   rnt.OptString
		tags     []interface{}
		title    interface{}
		video_id string
		webpage  string
	)
	video_id = (self).MatchID(url)
	webpage = (self).DownloadWebpageURL(url, video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	msi_id = (self).SearchRegexOne("(?s)config\\s*=\\s*{.+?video_id\\s*:\\s*'([^']+)'", webpage, "msi id", rnt.NoDefault, true, 0, nil)
	msi_data = rnt.UnsafeSubscript((self).DownloadJSON("http://player.fc-zenit.ru/msi/video", msi_id.Get(), rnt.AsOptString("Downloading JSON metadata"), rnt.AsOptString("Unable to download JSON metadata"), nil, true, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{
		"video": msi_id,
	}), "data")
	title = rnt.UnsafeSubscript(msi_data, "name")
	formats = func() []rnt.SDict {
		τresult := []rnt.SDict{}
		for _, τel := range τ_cast_α_to_Lα(rnt.UnsafeSubscript(msi_data, "qualities")) {
			var (
				q interface{}
			)
			q = τel
			if !(rnt.IsTruthy(rnt.DictGet(τ_cast_α_to_d(q), "url", nil))) {
				continue
			}
			τresult = append(τresult, rnt.SDict{
				"format_id": rnt.DictGet(τ_cast_α_to_d(q), "label", nil),
				"url":       rnt.UnsafeSubscript(q, "url"),
				"height":    rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(q), "label", nil), 1, rnt.OptInt{}, 1),
			})
		}
		return τresult
	}()
	formats = (self).SortFormats(formats)
	tags = func() []interface{} {
		τresult := []interface{}{}
		for _, τel := range τ_cast_α_to_Lα(rnt.DictGet(τ_cast_α_to_d(msi_data), "tags", []interface{}{})) {
			var (
				tag interface{}
			)
			tag = τel
			if !(rnt.IsTruthy(rnt.DictGet(τ_cast_α_to_d(tag), "label", nil))) {
				continue
			}
			τresult = append(τresult, rnt.UnsafeSubscript(tag, "label"))
		}
		return τresult
	}()
	return rnt.SDict{
		"id":        video_id,
		"title":     title,
		"thumbnail": rnt.DictGet(τ_cast_α_to_d(msi_data), "preview", nil),
		"formats":   formats,
		"duration":  rnt.FloatOrNone(rnt.DictGet(τ_cast_α_to_d(msi_data), "duration", nil), 1, 1, rnt.OptFloat{}),
		"timestamp": rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(msi_data), "date", nil), 1, rnt.OptInt{}, 1),
		"tags":      tags,
	}
}

func (self *FczenitIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("Fczenit", NewFczenitIE)
}
