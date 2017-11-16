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
 * acastie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/acast.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type ACastIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewACastIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &ACastIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "ACast"
	IE_NAME = "acast"
	_VALID_URL = "https?://(?:www\\.)?acast\\.com/(?P<channel>[^/]+)/(?P<id>[^/#?]+)"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *ACastIE) Key() string {
	return "ACast"
}

func (self *ACastIE) Name() string {
	return self.IE_NAME
}

func (self *ACastIE) _real_extract(url string) rnt.SDict {
	var (
		cast_data  interface{}
		channel    rnt.OptString
		display_id rnt.OptString
		τmp1       []rnt.OptString
	)
	τmp1 = rnt.ReMatchGroups(rnt.ReMatch((self).VALIDURL, url, 0), rnt.OptString{})
	channel = (τmp1)[0]
	display_id = (τmp1)[1]
	cast_data = (self).DownloadJSON(rnt.StrFormat2("https://embed.acast.com/api/acasts/%s/%s", channel, display_id), display_id.Get(), rnt.AsOptString("Downloading JSON metadata"), rnt.AsOptString("Unable to download JSON metadata"), nil, true, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	return rnt.SDict{
		"id":         rnt.ConvertToString(rnt.UnsafeSubscript(cast_data, "id")),
		"display_id": display_id,
		"url": (func() []interface{} {
			τresult := []interface{}{}
			for _, τel := range τ_cast_α_to_Lα(rnt.UnsafeSubscript(cast_data, "blings")) {
				var (
					b interface{}
				)
				b = τel
				if !((rnt.UnsafeSubscript(b, "type")) == ("BlingAudio")) {
					continue
				}
				τresult = append(τresult, rnt.UnsafeSubscript(b, "audio"))
			}
			return τresult
		}())[0],
		"title":       rnt.UnsafeSubscript(cast_data, "name"),
		"description": rnt.DictGet(τ_cast_α_to_d(cast_data), "description", nil),
		"thumbnail":   rnt.DictGet(τ_cast_α_to_d(cast_data), "image", nil),
		"timestamp":   rnt.ParseISO8601(rnt.CastToOptString(rnt.DictGet(τ_cast_α_to_d(cast_data), "publishingDate", nil)), "T"),
		"duration":    rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(cast_data), "duration", nil), 1, rnt.OptInt{}, 1),
	}
}

func (self *ACastIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("ACast", NewACastIE)
}
