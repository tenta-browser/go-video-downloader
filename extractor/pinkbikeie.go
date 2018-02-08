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
 * pinkbikeie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/pinkbike.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type PinkbikeIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewPinkbikeIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &PinkbikeIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "Pinkbike"
	_VALID_URL = "https?://(?:(?:www\\.)?pinkbike\\.com/video/|es\\.pinkbike\\.org/i/kvid/kvid-y5\\.swf\\?id=)(?P<id>[0-9]+)"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *PinkbikeIE) Key() string {
	return "Pinkbike"
}

func (self *PinkbikeIE) Name() string {
	return self.IE_NAME
}

func (self *PinkbikeIE) _real_extract(url string) rnt.SDict {
	var (
		_             string
		comment_count rnt.OptInt
		description   rnt.OptString
		duration      rnt.OptInt
		extract_count func(string, string) rnt.OptInt
		format_id     string
		formats       []interface{}
		height        rnt.OptInt
		location      rnt.OptString
		src           string
		thumbnail     rnt.OptString
		title         rnt.OptString
		upload_date   rnt.OptString
		uploader      rnt.OptString
		video_id      string
		view_count    rnt.OptInt
		webpage       string
		τmp1          []string
	)
	video_id = (self).MatchID(url)
	webpage = (self).DownloadWebpageURL(rnt.StrFormat2("http://www.pinkbike.com/video/%s", video_id), video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	formats = []interface{}{}
	for _, τel := range rnt.ReFindAllMulti("data-quality=((?:\\\\)?[\"\\'])(.+?)\\1[^>]+src=\\1(.+?)\\1", webpage, 0) {
		τmp1 = τel
		_ = (τmp1)[0]
		format_id = (τmp1)[1]
		src = (τmp1)[2]
		height = rnt.IntOrNone((self).SearchRegexOne("^(\\d+)[pP]$", format_id, "height", nil, true, 0, nil), 1, rnt.OptInt{}, 1)
		formats = append(formats, rnt.SDict{
			"url":       src,
			"format_id": format_id,
			"height":    height,
		})
	}
	formats = τ_conv_Ld_to_Lα((self).SortFormats(τ_conv_Lα_to_Ld(formats)))
	title = rnt.RemoveEnd((self).OgSearchTitle(webpage, rnt.NoDefault, true), " Video - Pinkbike")
	description = func() rnt.OptString {
		if v := ((self).HTMLSearchRegexOne("(?s)id=\"media-description\"[^>]*>(.+?)<", webpage, "description", nil, true, 0, nil)); τ_isTruthy_Os(v) {
			return v
		} else {
			return rnt.RemoveStart((self).OgSearchDescription(webpage, rnt.NoDefault), (title.Get() + ". "))
		}
	}()
	thumbnail = (self).OgSearchThumbnail(webpage, rnt.NoDefault)
	duration = rnt.IntOrNone((self).HTMLSearchMetaOne("video:duration", webpage, rnt.AsOptString("duration"), false, rnt.NoDefault, 0), 1, rnt.OptInt{}, 1)
	uploader = (self).SearchRegexOne("<a[^>]+\\brel=[\"\\']author[^>]+>([^<]+)", webpage, "uploader", rnt.NoDefault, false, 0, nil)
	upload_date = rnt.UnifiedStrDate((self).SearchRegexOne("class=\"fullTime\"[^>]+title=\"([^\"]+)\"", webpage, "upload date", rnt.NoDefault, false, 0, nil), true)
	location = (self).HTMLSearchRegexOne("(?s)<dt>Location</dt>\\s*<dd>(.+?)<", webpage, "location", rnt.NoDefault, false, 0, nil)
	extract_count = func(webpage string, label string) rnt.OptInt {
		return rnt.StrToInt((self).SearchRegexOne(rnt.StrFormat2("<span[^>]+class=\"stat-num\"[^>]*>([\\d,.]+)</span>\\s*<span[^>]+class=\"stat-label\"[^>]*>%s", label), webpage, label, rnt.NoDefault, false, 0, nil))
	}
	view_count = extract_count(webpage, "Views")
	comment_count = extract_count(webpage, "Comments")
	return rnt.SDict{
		"id":            video_id,
		"title":         title,
		"description":   description,
		"thumbnail":     thumbnail,
		"duration":      duration,
		"upload_date":   upload_date,
		"uploader":      uploader,
		"location":      location,
		"view_count":    view_count,
		"comment_count": comment_count,
		"formats":       formats,
	}
}

func (self *PinkbikeIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("Pinkbike", NewPinkbikeIE)
}
