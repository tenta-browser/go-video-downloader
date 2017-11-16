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
 * nextmediaie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/nextmedia.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type NextMediaIE struct {
	*rnt.CommonIE
	IE_DESC      string
	IE_NAME      string
	_URL_PATTERN string
}

func NewNextMediaIE() rnt.InfoExtractor {
	var (
		IE_DESC      string
		IE_NAME      string
		_URL_PATTERN string
		_VALID_URL   string
	)
	self := &NextMediaIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "NextMedia"
	IE_DESC = "蘋果日報"
	_VALID_URL = "https?://hk\\.apple\\.nextmedia\\.com/[^/]+/[^/]+/(?P<date>\\d+)/(?P<id>\\d+)"
	_URL_PATTERN = "\\{ url: \\'(.+)\\' \\}"
	self.IE_DESC = IE_DESC
	self.IE_NAME = IE_NAME
	self._URL_PATTERN = _URL_PATTERN
	self.VALIDURL = _VALID_URL
	return self
}

func (self *NextMediaIE) Key() string {
	return "NextMedia"
}

func (self *NextMediaIE) Name() string {
	return self.IE_NAME
}

func (self *NextMediaIE) _fetch_title(page string) rnt.OptString {
	return (self).OgSearchTitle(page, rnt.NoDefault, true)
}

func (self *NextMediaIE) _fetch_thumbnail(page string) rnt.OptString {
	return (self).OgSearchThumbnail(page, rnt.NoDefault)
}

func (self *NextMediaIE) _fetch_description(page string) rnt.OptString {
	return (self).OgSearchPropertyOne("description", page, rnt.OptString{}, rnt.NoDefault, true)
}

func (self *NextMediaIE) _fetch_timestamp(page string) rnt.OptInt {
	var (
		dateCreated rnt.OptString
	)
	dateCreated = (self).SearchRegexOne("\"dateCreated\":\"([^\"]+)\"", page, "created time", rnt.NoDefault, true, 0, nil)
	return rnt.ParseISO8601(dateCreated, "T")
}

func (self *NextMediaIE) _fetch_upload_date(url string) rnt.OptString {
	return (self).SearchRegexOne((self).VALIDURL, url, "upload date", rnt.NoDefault, true, 0, "date")
}

func (self *NextMediaIE) _extract_from_nextmedia_page(news_id string, url string, page string) rnt.SDict {
	var (
		attrs           rnt.SDict
		redirection_url rnt.OptString
		timestamp       rnt.OptInt
		title           rnt.OptString
		video_url       rnt.OptString
	)
	redirection_url = (self).SearchRegexOne("window\\.location\\.href\\s*=\\s*([\\'\"])(?P<url>(?!\\1).+)\\1", page, "redirection URL", nil, true, 0, "url")
	if τ_isTruthy_Os(redirection_url) {
		return (self).URLResult(rnt.URLJoin(url, redirection_url), rnt.OptString{}, rnt.OptString{}, rnt.OptString{})
	}
	title = (self)._fetch_title(page)
	video_url = (self).SearchRegexOne((self)._URL_PATTERN, page, "video url", rnt.NoDefault, true, 0, nil)
	attrs = rnt.SDict{
		"id":          news_id,
		"title":       title,
		"url":         video_url,
		"thumbnail":   (self)._fetch_thumbnail(page),
		"description": (self)._fetch_description(page),
	}
	timestamp = (self)._fetch_timestamp(page)
	if τ_isTruthy_Oi(timestamp) {
		(attrs)["timestamp"] = timestamp
	} else {
		(attrs)["upload_date"] = (self)._fetch_upload_date(url)
	}
	return attrs
}

func (self *NextMediaIE) _real_extract(url string) rnt.SDict {
	var (
		news_id string
		page    string
	)
	news_id = (self).MatchID(url)
	page = (self).DownloadWebpageURL(url, news_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	return (self)._extract_from_nextmedia_page(news_id, url, page)
}

func (self *NextMediaIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("NextMedia", NewNextMediaIE)
}
