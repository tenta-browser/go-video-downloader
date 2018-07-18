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
 * dumpertie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/dumpert.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type DumpertIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewDumpertIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &DumpertIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "Dumpert"
	_VALID_URL = "(?P<protocol>https?)://(?:www\\.)?dumpert\\.nl/(?:mediabase|embed)/(?P<id>[0-9]+/[0-9a-zA-Z]+)"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *DumpertIE) Key() string {
	return "Dumpert"
}

func (self *DumpertIE) Name() string {
	return self.IE_NAME
}

func (self *DumpertIE) _real_extract(url string) rnt.SDict {
	var (
		description  rnt.OptString
		files        interface{}
		files_base64 rnt.OptString
		formats      []rnt.SDict
		mobj         rnt.Match
		protocol     rnt.OptString
		quality      func(string) int
		req          rnt.Request
		thumbnail    interface{}
		title        rnt.OptString
		video_id     rnt.OptString
		webpage      string
	)
	mobj = rnt.ReMatch((self).VALIDURL, url, 0)
	video_id = rnt.ReMatchGroupOne(mobj, "id")
	protocol = rnt.ReMatchGroupOne(mobj, "protocol")
	url = rnt.StrFormat2("%s://www.dumpert.nl/mediabase/%s", protocol, video_id)
	req = rnt.SanitizedRequest(url, nil, rnt.SDict{})
	rnt.RequestAddHeader(req, "Cookie", "nsfw=1; cpc=10")
	webpage = (self).DownloadWebpageRequest(req, video_id.Get(), rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{}, nil)
	files_base64 = (self).SearchRegexOne("data-files=\"([^\"]+)\"", webpage, "data files", rnt.NoDefault, true, 0, nil)
	files = (self).ParseJSON(rnt.BytesToStr(rnt.Base64StdDecodeString(files_base64.Get()), "utf-8"), video_id.Get(), nil, true)
	quality = rnt.Qualities([]string{"flv", "mobile", "tablet", "720p"})
	formats = func() []rnt.SDict {
		τresult := []rnt.SDict{}
		for _, τel := range rnt.DictItems(τ_cast_α_to_d(files)) {
			var (
				format_id string
				video_url interface{}
				τmp1      τ_Tsαω
			)
			τmp1 = τel
			format_id = (τmp1).Φ0
			video_url = (τmp1).Φ1
			if !((format_id) != ("still")) {
				continue
			}
			τresult = append(τresult, rnt.SDict{
				"url":       video_url,
				"format_id": format_id,
				"quality":   quality(format_id),
			})
		}
		return τresult
	}()
	formats = (self).SortFormats(formats)
	title = func() rnt.OptString {
		if v := ((self).HTMLSearchMetaOne("title", webpage, rnt.OptString{}, false, rnt.NoDefault, 0)); τ_isTruthy_Os(v) {
			return v
		} else {
			return (self).OgSearchTitle(webpage, rnt.NoDefault, true)
		}
	}()
	description = func() rnt.OptString {
		if v := ((self).HTMLSearchMetaOne("description", webpage, rnt.OptString{}, false, rnt.NoDefault, 0)); τ_isTruthy_Os(v) {
			return v
		} else {
			return (self).OgSearchDescription(webpage, rnt.NoDefault)
		}
	}()
	thumbnail = func() interface{} {
		if v := (rnt.DictGet(τ_cast_α_to_d(files), "still", nil)); rnt.IsTruthy(v) {
			return v
		} else {
			return (self).OgSearchThumbnail(webpage, rnt.NoDefault)
		}
	}()
	return rnt.SDict{
		"id":          video_id,
		"title":       title,
		"description": description,
		"thumbnail":   thumbnail,
		"formats":     formats,
	}
}

func (self *DumpertIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("Dumpert", NewDumpertIE)
}
