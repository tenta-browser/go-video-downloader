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
 * adobetvie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/adobetv.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type AdobeTVIE struct {
	*rnt.CommonIE
	IE_NAME       string
	_API_BASE_URL string
}

func NewAdobeTVIE() rnt.InfoExtractor {
	var (
		IE_NAME       string
		_API_BASE_URL string
		_VALID_URL    string
	)
	self := &AdobeTVIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "AdobeTV"
	_API_BASE_URL = "http://tv.adobe.com/api/v4/"
	_VALID_URL = "https?://tv\\.adobe\\.com/(?:(?P<language>fr|de|es|jp)/)?watch/(?P<show_urlname>[^/]+)/(?P<id>[^/]+)"
	self.IE_NAME = IE_NAME
	self._API_BASE_URL = _API_BASE_URL
	self.VALIDURL = _VALID_URL
	return self
}

func (self *AdobeTVIE) Key() string {
	return "AdobeTV"
}

func (self *AdobeTVIE) Name() string {
	return self.IE_NAME
}

func (self *AdobeTVIE) _real_extract(url string) rnt.SDict {
	var (
		formats      []rnt.SDict
		language     rnt.OptString
		show_urlname rnt.OptString
		urlname      rnt.OptString
		video_data   interface{}
		τmp1         []rnt.OptString
	)
	τmp1 = rnt.ReMatchGroups(rnt.ReMatch((self).VALIDURL, url, 0), rnt.OptString{})
	language = (τmp1)[0]
	show_urlname = (τmp1)[1]
	urlname = (τmp1)[2]
	if !(τ_isTruthy_Os(language)) {
		language = rnt.AsOptString("en")
	}
	video_data = rnt.UnsafeSubscript(rnt.UnsafeSubscript((self).DownloadJSON(((self)._API_BASE_URL+rnt.StrFormat2("episode/get/?language=%s&show_urlname=%s&urlname=%s&disclosure=standard", language, show_urlname, urlname)), urlname.Get(), rnt.AsOptString("Downloading JSON metadata"), rnt.AsOptString("Unable to download JSON metadata"), nil, true, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{}, nil), "data"), 0)
	formats = func() []rnt.SDict {
		τresult := []rnt.SDict{}
		for _, τel := range τ_cast_α_to_Lα(rnt.UnsafeSubscript(video_data, "videos")) {
			var (
				source interface{}
			)
			source = τel
			τresult = append(τresult, rnt.SDict{
				"url": rnt.UnsafeSubscript(source, "url"),
				"format_id": func() interface{} {
					if v := (rnt.DictGet(τ_cast_α_to_d(source), "quality_level", nil)); rnt.IsTruthy(v) {
						return v
					} else if v := ((rnt.StrSplit(τ_index_Ls(rnt.StrSplit(rnt.CastToString(rnt.UnsafeSubscript(source, "url")), "-", -(1)), -(1)), ".", -(1)))[0]); (v) != "" {
						return v
					} else {
						return nil
					}
				}(),
				"width":  rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(source), "width", nil), 1, rnt.OptInt{}, 1),
				"height": rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(source), "height", nil), 1, rnt.OptInt{}, 1),
				"tbr":    rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(source), "video_data_rate", nil), 1, rnt.OptInt{}, 1),
			})
		}
		return τresult
	}()
	formats = (self).SortFormats(formats)
	return rnt.SDict{
		"id":          rnt.ConvertToString(rnt.UnsafeSubscript(video_data, "id")),
		"title":       rnt.UnsafeSubscript(video_data, "title"),
		"description": rnt.DictGet(τ_cast_α_to_d(video_data), "description", nil),
		"thumbnail":   rnt.DictGet(τ_cast_α_to_d(video_data), "thumbnail", nil),
		"upload_date": rnt.UnifiedStrDate(rnt.CastToOptString(rnt.DictGet(τ_cast_α_to_d(video_data), "start_date", nil)), true),
		"duration":    rnt.ParseDuration(rnt.CastToOptString(rnt.DictGet(τ_cast_α_to_d(video_data), "duration", nil))),
		"view_count":  rnt.StrToInt(rnt.CastToOptString(rnt.DictGet(τ_cast_α_to_d(video_data), "playcount", nil))),
		"formats":     formats,
	}
}

func (self *AdobeTVIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("AdobeTV", NewAdobeTVIE)
}
