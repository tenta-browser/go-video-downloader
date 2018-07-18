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
 * macgamestoreie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/macgamestore.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type MacGameStoreIE struct {
	*rnt.CommonIE
	IE_DESC string
	IE_NAME string
}

func NewMacGameStoreIE() rnt.InfoExtractor {
	var (
		IE_DESC    string
		IE_NAME    string
		_VALID_URL string
	)
	self := &MacGameStoreIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "MacGameStore"
	IE_NAME = "macgamestore"
	IE_DESC = "MacGameStore trailers"
	_VALID_URL = "https?://(?:www\\.)?macgamestore\\.com/mediaviewer\\.php\\?trailer=(?P<id>\\d+)"
	self.IE_DESC = IE_DESC
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *MacGameStoreIE) Key() string {
	return "MacGameStore"
}

func (self *MacGameStoreIE) Name() string {
	return self.IE_NAME
}

func (self *MacGameStoreIE) _real_extract(url string) rnt.SDict {
	var (
		video_id    string
		video_title rnt.OptString
		video_url   rnt.OptString
		webpage     string
	)
	video_id = (self).MatchID(url)
	webpage = (self).DownloadWebpageURL(url, video_id, rnt.AsOptString("Downloading trailer page"), rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{}, nil)
	if rnt.StrContains(webpage, ">Missing Media<") {
		panic(rnt.PyExtractorError(rnt.StrFormat2("Trailer %s does not exist", video_id), true, rnt.OptString{}))
	}
	video_title = (self).HTMLSearchRegexOne("<title>MacGameStore: (.*?) Trailer</title>", webpage, "title", rnt.NoDefault, true, 0, nil)
	video_url = (self).HTMLSearchRegexOne("(?s)<div\\s+id=\"video-player\".*?href=\"([^\"]+)\"\\s*>", webpage, "video URL", rnt.NoDefault, true, 0, nil)
	return rnt.SDict{
		"id":    video_id,
		"url":   video_url,
		"title": video_title,
	}
}

func (self *MacGameStoreIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("MacGameStore", NewMacGameStoreIE)
}
