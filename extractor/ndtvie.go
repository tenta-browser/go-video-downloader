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
 * ndtvie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/ndtv.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type NDTVIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewNDTVIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &NDTVIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "NDTV"
	_VALID_URL = "https?://(?:[^/]+\\.)?ndtv\\.com/(?:[^/]+/)*videos?/?(?:[^/]+/)*[^/?^&]+-(?P<id>\\d+)"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *NDTVIE) Key() string {
	return "NDTV"
}

func (self *NDTVIE) Name() string {
	return self.IE_NAME
}

func (self *NDTVIE) _real_extract(url string) rnt.SDict {
	var (
		description rnt.OptString
		duration    rnt.OptString
		filename    rnt.OptString
		title       string
		upload_date rnt.OptString
		video_id    string
		video_url   rnt.OptString
		webpage     string
	)
	video_id = (self).MatchID(url)
	webpage = (self).DownloadWebpageURL(url, video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{}, nil)
	title = rnt.ParseUnquotePlus(func() rnt.OptString {
		if v := ((self).SearchRegexOne("__title\\s*=\\s*'([^']+)'", webpage, "title", nil, true, 0, nil)); τ_isTruthy_Os(v) {
			return v
		} else {
			return (self).OgSearchTitle(webpage, rnt.NoDefault, true)
		}
	}().Get())
	filename = (self).SearchRegexOne("(?:__)?filename\\s*[:=]\\s*'([^']+)'", webpage, "video filename", rnt.NoDefault, true, 0, nil)
	video_url = rnt.UtilURLJoin("https://ndtvod.bc-ssl.cdn.bitgravity.com/23372/ndtv/", rnt.StrLStrip(filename.Get(), "/"))
	duration = rnt.ParseDuration((self).SearchRegexOne("(?:__)?duration\\s*[:=]\\s*'([^']+)'", webpage, "duration", rnt.NoDefault, false, 0, nil))
	upload_date = rnt.UnifiedStrDate(func() rnt.OptString {
		if v := ((self).HTMLSearchMetaOne("publish-date", webpage, rnt.AsOptString("upload date"), false, nil, 0)); τ_isTruthy_Os(v) {
			return v
		} else if v := ((self).HTMLSearchMetaOne("uploadDate", webpage, rnt.AsOptString("upload date"), false, nil, 0)); τ_isTruthy_Os(v) {
			return v
		} else {
			return (self).SearchRegexOne("datePublished\"\\s*:\\s*\"([^\"]+)\"", webpage, "upload date", rnt.NoDefault, false, 0, nil)
		}
	}(), true)
	description = rnt.RemoveEnd((self).OgSearchDescription(webpage, rnt.NoDefault), " (Read more)")
	return rnt.SDict{
		"id":          video_id,
		"url":         video_url,
		"title":       title,
		"description": description,
		"thumbnail":   (self).OgSearchThumbnail(webpage, rnt.NoDefault),
		"duration":    duration,
		"upload_date": upload_date,
	}
}

func (self *NDTVIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("NDTV", NewNDTVIE)
}
