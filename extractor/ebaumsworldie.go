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
 * ebaumsworldie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/ebaumsworld.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type EbaumsWorldIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewEbaumsWorldIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &EbaumsWorldIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "EbaumsWorld"
	_VALID_URL = "https?://(?:www\\.)?ebaumsworld\\.com/videos/[^/]+/(?P<id>\\d+)"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *EbaumsWorldIE) Key() string {
	return "EbaumsWorld"
}

func (self *EbaumsWorldIE) Name() string {
	return self.IE_NAME
}

func (self *EbaumsWorldIE) _real_extract(url string) rnt.SDict {
	var (
		config    rnt.XMLElement
		video_id  string
		video_url string
	)
	video_id = (self).MatchID(url)
	config = (self).DownloadXML(rnt.StrFormat2("http://www.ebaumsworld.com/video/player/%s", video_id), video_id, rnt.AsOptString("Downloading XML"), rnt.AsOptString("Unable to download XML"), nil, true, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{}, nil)
	video_url = (rnt.XMLFind(config, "file")).Text
	return rnt.SDict{
		"id":          video_id,
		"title":       (rnt.XMLFind(config, "title")).Text,
		"url":         video_url,
		"description": (rnt.XMLFind(config, "description")).Text,
		"thumbnail":   (rnt.XMLFind(config, "image")).Text,
		"uploader":    (rnt.XMLFind(config, "username")).Text,
	}
}

func (self *EbaumsWorldIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("EbaumsWorld", NewEbaumsWorldIE)
}
