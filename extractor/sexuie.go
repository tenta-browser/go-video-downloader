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
 * sexuie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/sexu.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type SexuIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewSexuIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &SexuIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "Sexu"
	_VALID_URL = "https?://(?:www\\.)?sexu\\.com/(?P<id>\\d+)"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *SexuIE) Key() string {
	return "Sexu"
}

func (self *SexuIE) Name() string {
	return self.IE_NAME
}

func (self *SexuIE) _real_extract(url string) rnt.SDict {
	var (
		categories     []string
		categories_str rnt.OptString
		description    rnt.OptString
		formats        []rnt.SDict
		jwvideo        interface{}
		sources        interface{}
		thumbnail      interface{}
		title          rnt.OptString
		video_id       string
		webpage        string
	)
	video_id = (self).MatchID(url)
	webpage = (self).DownloadWebpageURL(url, video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	jwvideo = (self).ParseJSON((self).SearchRegexOne("\\.setup\\(\\s*({.+?})\\s*\\);", webpage, "jwvideo", rnt.NoDefault, true, 0, nil).Get(), video_id, nil, true)
	sources = rnt.UnsafeSubscript(jwvideo, "sources")
	formats = func() []rnt.SDict {
		τresult := []rnt.SDict{}
		for _, τel := range τ_cast_α_to_Lα(sources) {
			var (
				source interface{}
			)
			source = τel
			if !(rnt.IsTruthy(rnt.DictGet(τ_cast_α_to_d(source), "file", nil))) {
				continue
			}
			τresult = append(τresult, rnt.SDict{
				"url":       rnt.StrReplace(rnt.CastToString(rnt.UnsafeSubscript(source, "file")), "\\", "", -(1)),
				"format_id": rnt.DictGet(τ_cast_α_to_d(source), "label", nil),
				"height":    rnt.ConvertToInt((self).SearchRegexOne("^(\\d+)[pP]", rnt.CastToString(rnt.DictGet(τ_cast_α_to_d(source), "label", "")), "height", nil, true, 0, nil)),
			})
		}
		return τresult
	}()
	formats = (self).SortFormats(formats)
	title = (self).HTMLSearchRegexOne("<title>([^<]+)\\s*-\\s*Sexu\\.Com</title>", webpage, "title", rnt.NoDefault, true, 0, nil)
	description = (self).HTMLSearchMetaOne("description", webpage, rnt.AsOptString("description"), false, rnt.NoDefault, 0)
	thumbnail = rnt.DictGet(τ_cast_α_to_d(jwvideo), "image", nil)
	categories_str = (self).HTMLSearchMetaOne("keywords", webpage, rnt.AsOptString("categories"), false, rnt.NoDefault, 0)
	categories = func() []string {
		if !categories_str.IsSet() {
			return nil
		} else {
			return rnt.StrSplit(categories_str.Get(), ",", -(1))
		}
	}()
	return rnt.SDict{
		"id":          video_id,
		"title":       title,
		"description": description,
		"thumbnail":   thumbnail,
		"categories":  categories,
		"formats":     formats,
		"age_limit":   18,
	}
}

func (self *SexuIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("Sexu", NewSexuIE)
}
