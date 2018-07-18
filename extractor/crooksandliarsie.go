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
 * crooksandliarsie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/crooksandliars.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type CrooksAndLiarsIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewCrooksAndLiarsIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &CrooksAndLiarsIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "CrooksAndLiars"
	_VALID_URL = "https?://embed\\.crooksandliars\\.com/(?:embed|v)/(?P<id>[A-Za-z0-9]+)"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *CrooksAndLiarsIE) Key() string {
	return "CrooksAndLiars"
}

func (self *CrooksAndLiarsIE) Name() string {
	return self.IE_NAME
}

func (self *CrooksAndLiarsIE) _real_extract(url string) rnt.SDict {
	var (
		formats  []rnt.SDict
		manifest interface{}
		quality  func(string) int
		video_id string
		webpage  string
	)
	video_id = (self).MatchID(url)
	webpage = (self).DownloadWebpageURL(rnt.StrFormat2("http://embed.crooksandliars.com/embed/%s", video_id), video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{}, nil)
	manifest = (self).ParseJSON((self).SearchRegexOne("var\\s+manifest\\s*=\\s*({.+?})\\n", webpage, "manifest JSON", rnt.NoDefault, true, 0, nil).Get(), video_id, nil, true)
	quality = rnt.Qualities(τ_conv_Tssssω_to_Ls(τ_Tssssω{
		Φ0: "webm_low",
		Φ1: "mp4_low",
		Φ2: "webm_high",
		Φ3: "mp4_high",
	}))
	formats = func() []rnt.SDict {
		τresult := []rnt.SDict{}
		for _, τel := range τ_cast_α_to_Lα(rnt.UnsafeSubscript(manifest, "flavors")) {
			var (
				item interface{}
			)
			item = τel
			if !(rnt.StrStartswith(rnt.CastToString(rnt.UnsafeSubscript(item, "mime")), "video/", rnt.OptInt{}, rnt.OptInt{})) {
				continue
			}
			τresult = append(τresult, rnt.SDict{
				"url":       rnt.UnsafeSubscript(item, "url"),
				"format_id": rnt.UnsafeSubscript(item, "type"),
				"quality":   quality(rnt.CastToString(rnt.UnsafeSubscript(item, "type"))),
			})
		}
		return τresult
	}()
	formats = (self).SortFormats(formats)
	return rnt.SDict{
		"url":         url,
		"id":          video_id,
		"title":       rnt.UnsafeSubscript(manifest, "title"),
		"description": rnt.DictGet(τ_cast_α_to_d(manifest), "description", nil),
		"thumbnail":   (self).ProtoRelativeURL(rnt.CastToOptString(rnt.DictGet(τ_cast_α_to_d(manifest), "poster", nil)), rnt.OptString{}),
		"timestamp":   rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(manifest), "created", nil), 1, rnt.OptInt{}, 1),
		"uploader":    rnt.DictGet(τ_cast_α_to_d(manifest), "author", nil),
		"duration":    rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(manifest), "duration", nil), 1, rnt.OptInt{}, 1),
		"formats":     formats,
	}
}

func (self *CrooksAndLiarsIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("CrooksAndLiars", NewCrooksAndLiarsIE)
}
