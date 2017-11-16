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
 * bigflixie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/bigflix.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type BigflixIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewBigflixIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &BigflixIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "Bigflix"
	_VALID_URL = "https?://(?:www\\.)?bigflix\\.com/.+/(?P<id>[0-9]+)"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *BigflixIE) Key() string {
	return "Bigflix"
}

func (self *BigflixIE) Name() string {
	return self.IE_NAME
}

func (self *BigflixIE) _real_extract(url string) rnt.SDict {
	var (
		decode_url  func(string) string
		description rnt.OptString
		encoded_url string
		f           rnt.SDict
		file_url    rnt.OptString
		formats     []interface{}
		height      string
		title       rnt.OptString
		video_id    string
		video_url   string
		webpage     string
		τmp1        []string
	)
	video_id = (self).MatchID(url)
	webpage = (self).DownloadWebpageURL(url, video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	title = (self).HTMLSearchRegexOne("<div[^>]+class=[\"\\']pagetitle[\"\\'][^>]*>(.+?)</div>", webpage, "title", rnt.NoDefault, true, 0, nil)
	decode_url = func(quoted_b64_url string) string {
		return rnt.BytesToStr(rnt.Base64StdDecodeBytes(rnt.StrToBytes(rnt.ParseUnquote(quoted_b64_url), "ascii")), "utf-8")
	}
	formats = []interface{}{}
	for _, τel := range rnt.ReFindAllMulti("ContentURL_(\\d{3,4})[pP][^=]+=([^&]+)", webpage, 0) {
		τmp1 = τel
		height = (τmp1)[0]
		encoded_url = (τmp1)[1]
		video_url = decode_url(encoded_url)
		f = rnt.SDict{
			"url":       video_url,
			"format_id": rnt.StrFormat2("%sp", height),
			"height":    rnt.ConvertToInt(height),
		}
		if rnt.StrStartswith(video_url, "rtmp", rnt.OptInt{}, rnt.OptInt{}) {
			(f)["ext"] = "flv"
		}
		formats = append(formats, f)
	}
	file_url = (self).SearchRegexOne("file=([^&]+)", webpage, "video url", nil, true, 0, nil)
	if τ_isTruthy_Os(file_url) {
		video_url = decode_url(file_url.Get())
		if τ_all_Lb(func() []bool {
			τresult := []bool{}
			for _, τel := range formats {
				var (
					f interface{}
				)
				f = τel
				τresult = append(τresult, (rnt.UnsafeSubscript(f, "url")) != (video_url))
			}
			return τresult
		}()) {
			formats = append(formats, rnt.SDict{
				"url": decode_url(file_url.Get()),
			})
		}
	}
	formats = τ_conv_Ld_to_Lα((self).SortFormats(τ_conv_Lα_to_Ld(formats)))
	description = (self).HTMLSearchMetaOne("description", webpage, rnt.OptString{}, false, rnt.NoDefault, 0)
	return rnt.SDict{
		"id":          video_id,
		"title":       title,
		"description": description,
		"formats":     formats,
	}
}

func (self *BigflixIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("Bigflix", NewBigflixIE)
}
