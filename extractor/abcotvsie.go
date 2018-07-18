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
 * abcotvsie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/abcotvs.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type ABCOTVSIE struct {
	*rnt.CommonIE
	IE_DESC string
	IE_NAME string
}

func NewABCOTVSIE() rnt.InfoExtractor {
	var (
		IE_DESC    string
		IE_NAME    string
		_VALID_URL string
	)
	self := &ABCOTVSIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "ABCOTVS"
	IE_NAME = "abcotvs"
	IE_DESC = "ABC Owned Television Stations"
	_VALID_URL = "https?://(?:abc(?:7(?:news|ny|chicago)?|11|13|30)|6abc)\\.com(?:/[^/]+/(?P<display_id>[^/]+))?/(?P<id>\\d+)"
	self.IE_DESC = IE_DESC
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *ABCOTVSIE) Key() string {
	return "ABCOTVS"
}

func (self *ABCOTVSIE) Name() string {
	return self.IE_NAME
}

func (self *ABCOTVSIE) _real_extract(url string) rnt.SDict {
	var (
		description string
		display_id  rnt.OptString
		formats     []rnt.SDict
		m3u8        string
		mobj        rnt.Match
		thumbnail   rnt.OptString
		timestamp   rnt.OptInt
		title       string
		uploader    rnt.OptString
		video_id    rnt.OptString
		webpage     string
	)
	mobj = rnt.ReMatch((self).VALIDURL, url, 0)
	video_id = rnt.ReMatchGroupOne(mobj, "id")
	display_id = func() rnt.OptString {
		if v := (rnt.ReMatchGroupOne(mobj, "display_id")); τ_isTruthy_Os(v) {
			return v
		} else {
			return video_id
		}
	}()
	webpage = (self).DownloadWebpageURL(url, display_id.Get(), rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{}, nil)
	m3u8 = (rnt.StrSplit((self).HTMLSearchMetaOne("contentURL", webpage, rnt.AsOptString("m3u8 url"), true, rnt.NoDefault, 0).Get(), "?", -(1)))[0]
	formats = (self).ExtractM3U8Formats(m3u8, display_id.Get(), rnt.AsOptString("mp4"), rnt.AsOptString("m3u8"), nil, rnt.OptString{}, rnt.OptString{}, rnt.OptString{}, true, false)
	formats = (self).SortFormats(formats)
	title = rnt.StrStrip((self).OgSearchTitle(webpage, rnt.NoDefault, true).Get(), "")
	description = rnt.StrStrip((self).OgSearchDescription(webpage, rnt.NoDefault).Get(), "")
	thumbnail = (self).OgSearchThumbnail(webpage, rnt.NoDefault)
	timestamp = rnt.ParseISO8601((self).SearchRegexOne("<div class=\"meta\">\\s*<time class=\"timeago\" datetime=\"([^\"]+)\">", webpage, "upload date", rnt.NoDefault, false, 0, nil), "T")
	uploader = (self).SearchRegexOne("rel=\"author\">([^<]+)</a>", webpage, "uploader", nil, true, 0, nil)
	return rnt.SDict{
		"id":          video_id,
		"display_id":  display_id,
		"title":       title,
		"description": description,
		"thumbnail":   thumbnail,
		"timestamp":   timestamp,
		"uploader":    uploader,
		"formats":     formats,
	}
}

func (self *ABCOTVSIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("ABCOTVS", NewABCOTVSIE)
}
