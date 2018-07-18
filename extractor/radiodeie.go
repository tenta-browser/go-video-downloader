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
 * radiodeie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/radiode.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type RadioDeIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewRadioDeIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &RadioDeIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "RadioDe"
	IE_NAME = "radio.de"
	_VALID_URL = "https?://(?P<id>.+?)\\.(?:radio\\.(?:de|at|fr|pt|es|pl|it)|rad\\.io)"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *RadioDeIE) Key() string {
	return "RadioDe"
}

func (self *RadioDeIE) Name() string {
	return self.IE_NAME
}

func (self *RadioDeIE) _real_extract(url string) rnt.SDict {
	var (
		broadcast   interface{}
		description interface{}
		formats     []rnt.SDict
		jscode      rnt.OptString
		radio_id    string
		thumbnail   interface{}
		title       string
		webpage     string
	)
	radio_id = (self).MatchID(url)
	webpage = (self).DownloadWebpageURL(url, radio_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{}, nil)
	jscode = (self).SearchRegexOne("'components/station/stationService':\\s*\\{\\s*'?station'?:\\s*(\\{.*?\\s*\\}),\\n", webpage, "broadcast", rnt.NoDefault, true, 0, nil)
	broadcast = (self).ParseJSON(jscode.Get(), radio_id, nil, true)
	title = (self).LiveTitle(rnt.CastToString(rnt.UnsafeSubscript(broadcast, "name")))
	description = func() interface{} {
		if v := (rnt.DictGet(τ_cast_α_to_d(broadcast), "description", nil)); rnt.IsTruthy(v) {
			return v
		} else {
			return rnt.DictGet(τ_cast_α_to_d(broadcast), "shortDescription", nil)
		}
	}()
	thumbnail = func() interface{} {
		if v := (rnt.DictGet(τ_cast_α_to_d(broadcast), "picture4Url", nil)); rnt.IsTruthy(v) {
			return v
		} else if v := (rnt.DictGet(τ_cast_α_to_d(broadcast), "picture4TransUrl", nil)); rnt.IsTruthy(v) {
			return v
		} else {
			return rnt.DictGet(τ_cast_α_to_d(broadcast), "logo100x100", nil)
		}
	}()
	formats = func() []rnt.SDict {
		τresult := []rnt.SDict{}
		for _, τel := range τ_cast_α_to_Lα(rnt.UnsafeSubscript(broadcast, "streamUrls")) {
			var (
				stream interface{}
			)
			stream = τel
			τresult = append(τresult, rnt.SDict{
				"url":    rnt.UnsafeSubscript(stream, "streamUrl"),
				"ext":    rnt.StrLower(rnt.CastToString(rnt.UnsafeSubscript(stream, "streamContentFormat"))),
				"acodec": rnt.UnsafeSubscript(stream, "streamContentFormat"),
				"abr":    rnt.UnsafeSubscript(stream, "bitRate"),
				"asr":    rnt.UnsafeSubscript(stream, "sampleRate"),
			})
		}
		return τresult
	}()
	formats = (self).SortFormats(formats)
	return rnt.SDict{
		"id":          radio_id,
		"title":       title,
		"description": description,
		"thumbnail":   thumbnail,
		"is_live":     true,
		"formats":     formats,
	}
}

func (self *RadioDeIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("RadioDe", NewRadioDeIE)
}
