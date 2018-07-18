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
 * pornhdie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/pornhd.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type PornHdIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewPornHdIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &PornHdIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "PornHd"
	_VALID_URL = "https?://(?:www\\.)?pornhd\\.com/(?:[a-z]{2,4}/)?videos/(?P<id>\\d+)(?:/(?P<display_id>.+))?"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *PornHdIE) Key() string {
	return "PornHd"
}

func (self *PornHdIE) Name() string {
	return self.IE_NAME
}

func (self *PornHdIE) _real_extract(url string) rnt.SDict {
	var (
		description rnt.OptString
		display_id  rnt.OptString
		format_id   string
		formats     []interface{}
		height      rnt.OptInt
		message     rnt.OptString
		mobj        rnt.Match
		sources     interface{}
		thumbnail   rnt.OptString
		title       rnt.OptString
		video_id    rnt.OptString
		video_url   interface{}
		view_count  rnt.OptInt
		webpage     string
		τmp1        τ_Tsαω
	)
	mobj = rnt.ReMatch((self).VALIDURL, url, 0)
	video_id = rnt.ReMatchGroupOne(mobj, "id")
	display_id = rnt.ReMatchGroupOne(mobj, "display_id")
	webpage = (self).DownloadWebpageURL(url, func() rnt.OptString {
		if v := (display_id); τ_isTruthy_Os(v) {
			return v
		} else {
			return video_id
		}
	}().Get(), rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{}, nil)
	title = (self).HTMLSearchRegexMulti([]string{"<span[^>]+class=[\"\\']video-name[\"\\'][^>]*>([^<]+)", "<title>(.+?) - .*?[Pp]ornHD.*?</title>"}, webpage, "title", rnt.NoDefault, true, 0, nil)
	sources = (self).ParseJSON(rnt.JsToJSON((self).SearchRegexOne("(?s)sources'?\\s*[:=]\\s*(\\{.+?\\})", webpage, "sources", "{}", true, 0, nil).Get()), video_id.Get(), nil, true)
	if !(rnt.IsTruthy(sources)) {
		message = (self).HTMLSearchRegexOne("(?s)<(div|p)[^>]+class=\"no-video\"[^>]*>(?P<value>.+?)</\\1", webpage, "error message", rnt.NoDefault, true, 0, "value")
		panic(rnt.PyExtractorError(rnt.StrFormat2("%s said: %s", (self).IE_NAME, message), true, rnt.OptString{}))
	}
	formats = []interface{}{}
	for _, τel := range rnt.DictItems(τ_cast_α_to_d(sources)) {
		τmp1 = τel
		format_id = (τmp1).Φ0
		video_url = (τmp1).Φ1
		if !(rnt.IsTruthy(video_url)) {
			continue
		}
		height = rnt.IntOrNone((self).SearchRegexOne("^(\\d+)[pP]", format_id, "height", nil, true, 0, nil), 1, rnt.OptInt{}, 1)
		formats = append(formats, rnt.SDict{
			"url":       video_url,
			"format_id": format_id,
			"height":    height,
		})
	}
	formats = τ_conv_Ld_to_Lα((self).SortFormats(τ_conv_Lα_to_Ld(formats)))
	description = (self).HTMLSearchRegexOne("<(div|p)[^>]+class=\"description\"[^>]*>(?P<value>[^<]+)</\\1", webpage, "description", rnt.NoDefault, false, 0, "value")
	view_count = rnt.IntOrNone((self).HTMLSearchRegexOne("(\\d+) views\\s*<", webpage, "view count", rnt.NoDefault, false, 0, nil), 1, rnt.OptInt{}, 1)
	thumbnail = (self).SearchRegexOne("poster'?\\s*:\\s*([\\\"'])(?P<url>(?:(?!\\1).)+)\\1", webpage, "thumbnail", rnt.NoDefault, false, 0, "url")
	return rnt.SDict{
		"id":          video_id,
		"display_id":  display_id,
		"title":       title,
		"description": description,
		"thumbnail":   thumbnail,
		"view_count":  view_count,
		"formats":     formats,
		"age_limit":   18,
	}
}

func (self *PornHdIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("PornHd", NewPornHdIE)
}
