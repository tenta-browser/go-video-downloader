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
 * spankbangie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/spankbang.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type SpankBangIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewSpankBangIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &SpankBangIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "SpankBang"
	_VALID_URL = "https?://(?:(?:www|m|[a-z]{2})\\.)?spankbang\\.com/(?P<id>[\\da-z]+)/video"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *SpankBangIE) Key() string {
	return "SpankBang"
}

func (self *SpankBangIE) Name() string {
	return self.IE_NAME
}

func (self *SpankBangIE) _real_extract(url string) rnt.SDict {
	var (
		age_limit   int
		description rnt.OptString
		duration    rnt.OptString
		f           rnt.SDict
		format_id   rnt.OptString
		format_url  rnt.OptString
		formats     []interface{}
		mobj        rnt.Match
		thumbnail   rnt.OptString
		title       rnt.OptString
		uploader    rnt.OptString
		video_id    string
		view_count  rnt.OptInt
		webpage     string
		τmp2        τ_TOsOsω
	)
	video_id = (self).MatchID(url)
	webpage = (self).DownloadWebpageURL(url, video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{
		"Cookie": "country=US",
	}, rnt.SDict{})
	if rnt.ReSearch("<[^>]+\\bid=[\"\\']video_removed", webpage, 0) != nil {
		panic(rnt.PyExtractorError(rnt.StrFormat2("Video %s is not available", video_id), true, rnt.OptString{}))
	}
	formats = []interface{}{}
	τmp1 := rnt.ReFindIter("stream_url_(?P<id>[^\\s=]+)\\s*=\\s*([\"\\'])(?P<url>(?:(?!\\2).)+)\\2", webpage, 0)
	for τel, τhv := τmp1.Next(); τhv; τel, τhv = τmp1.Next() {
		mobj = τel
		τmp2 = rnt.ReMatchGroupTwo(mobj, "id", "url")
		format_id = (τmp2).Φ0
		format_url = (τmp2).Φ1
		f = rnt.ParseResolution(format_id)
		rnt.DictUpdate(f, rnt.SDict{
			"url":       format_url,
			"format_id": format_id,
		})
		formats = append(formats, f)
	}
	formats = τ_conv_Ld_to_Lα((self).SortFormats(τ_conv_Lα_to_Ld(formats)))
	title = (self).HTMLSearchRegexOne("(?s)<h1[^>]*>(.+?)</h1>", webpage, "title", rnt.NoDefault, true, 0, nil)
	description = (self).SearchRegexOne("<div[^>]+\\bclass=[\"\\']bottom[^>]+>\\s*<p>[^<]*</p>\\s*<p>([^<]+)", webpage, "description", rnt.NoDefault, false, 0, nil)
	thumbnail = (self).OgSearchThumbnail(webpage, rnt.NoDefault)
	uploader = (self).SearchRegexOne("class=\"user\"[^>]*><img[^>]+>([^<]+)", webpage, "uploader", nil, true, 0, nil)
	duration = rnt.ParseDuration((self).SearchRegexOne("<div[^>]+\\bclass=[\"\\']right_side[^>]+>\\s*<span>([^<]+)", webpage, "duration", rnt.NoDefault, false, 0, nil))
	view_count = rnt.StrToInt((self).SearchRegexOne("([\\d,.]+)\\s+plays", webpage, "view count", rnt.NoDefault, false, 0, nil))
	age_limit = (self).RTASearch(webpage)
	return rnt.SDict{
		"id":          video_id,
		"title":       title,
		"description": description,
		"thumbnail":   thumbnail,
		"uploader":    uploader,
		"duration":    duration,
		"view_count":  view_count,
		"formats":     formats,
		"age_limit":   age_limit,
	}
}

func (self *SpankBangIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("SpankBang", NewSpankBangIE)
}
