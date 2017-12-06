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
		formats     []rnt.SDict
		stream_key  rnt.OptString
		thumbnail   rnt.OptString
		title       rnt.OptString
		uploader    rnt.OptString
		video_id    string
		webpage     string
	)
	video_id = (self).MatchID(url)
	webpage = (self).DownloadWebpageURL(url, video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	if rnt.ReSearch("<[^>]+\\bid=[\"\\']video_removed", webpage, 0) != nil {
		panic(rnt.PyExtractorError(rnt.StrFormat2("Video %s is not available", video_id), true, rnt.OptString{}))
	}
	stream_key = (self).HTMLSearchRegexOne("var\\s+stream_key\\s*=\\s*['\"](.+?)['\"]", webpage, "stream key", rnt.NoDefault, true, 0, nil)
	formats = func() []rnt.SDict {
		τresult := []rnt.SDict{}
		for _, τel := range rnt.ReFindAllOne("<(?:span|li|p)[^>]+[qb]_(\\d+)p", webpage, 0) {
			var (
				height string
			)
			height = τel
			τresult = append(τresult, rnt.SDict{
				"url":       rnt.StrFormat2("http://spankbang.com/_%s/%s/title/%sp__mp4", video_id, stream_key, height),
				"ext":       "mp4",
				"format_id": rnt.StrFormat2("%sp", height),
				"height":    rnt.ConvertToInt(height),
			})
		}
		return τresult
	}()
	formats = (self).CheckFormats(formats, video_id)
	formats = (self).SortFormats(formats)
	title = (self).HTMLSearchRegexOne("(?s)<h1[^>]*>(.+?)</h1>", webpage, "title", rnt.NoDefault, true, 0, nil)
	description = (self).OgSearchDescription(webpage, rnt.NoDefault)
	thumbnail = (self).OgSearchThumbnail(webpage, rnt.NoDefault)
	uploader = (self).SearchRegexOne("class=\"user\"[^>]*><img[^>]+>([^<]+)", webpage, "uploader", nil, true, 0, nil)
	age_limit = (self).RTASearch(webpage)
	return rnt.SDict{
		"id":          video_id,
		"title":       title,
		"description": description,
		"thumbnail":   thumbnail,
		"uploader":    uploader,
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
