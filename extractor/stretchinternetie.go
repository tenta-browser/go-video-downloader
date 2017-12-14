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
 * stretchinternetie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/stretchinternet.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type StretchInternetIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewStretchInternetIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &StretchInternetIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "StretchInternet"
	_VALID_URL = "https?://portal\\.stretchinternet\\.com/[^/]+/portal\\.htm\\?.*?\\beventId=(?P<id>\\d+)"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *StretchInternetIE) Key() string {
	return "StretchInternet"
}

func (self *StretchInternetIE) Name() string {
	return self.IE_NAME
}

func (self *StretchInternetIE) _real_extract(url string) rnt.SDict {
	var (
		description interface{}
		event       interface{}
		stream      interface{}
		timestamp   rnt.OptInt
		title       interface{}
		video_id    string
		video_url   string
	)
	video_id = (self).MatchID(url)
	stream = (self).DownloadJSON(rnt.StrFormat2("https://neo-client.stretchinternet.com/streamservice/v1/media/stream/v%s", video_id), video_id, rnt.AsOptString("Downloading JSON metadata"), rnt.AsOptString("Unable to download JSON metadata"), nil, true, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	video_url = rnt.StrFormat2("https://%s", rnt.UnsafeSubscript(stream, "source"))
	event = rnt.UnsafeSubscript((self).DownloadJSON("https://neo-client.stretchinternet.com/portal-ws/getEvent.json", video_id, rnt.AsOptString("Downloading JSON metadata"), rnt.AsOptString("Unable to download JSON metadata"), nil, true, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{
		"clientID": 99997,
		"eventID":  video_id,
		"token":    "asdf",
	}), "event")
	title = func() interface{} {
		if v := (rnt.DictGet(τ_cast_α_to_d(event), "title", nil)); rnt.IsTruthy(v) {
			return v
		} else {
			return rnt.UnsafeSubscript(event, "mobileTitle")
		}
	}()
	description = rnt.DictGet(τ_cast_α_to_d(event), "customText", nil)
	timestamp = rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(event), "longtime", nil), 1, rnt.OptInt{}, 1)
	return rnt.SDict{
		"id":          video_id,
		"title":       title,
		"description": description,
		"timestamp":   timestamp,
		"url":         video_url,
	}
}

func (self *StretchInternetIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("StretchInternet", NewStretchInternetIE)
}
