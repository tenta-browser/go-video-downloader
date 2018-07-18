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
 * soundgasmie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/soundgasm.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type SoundgasmIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewSoundgasmIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &SoundgasmIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "Soundgasm"
	IE_NAME = "soundgasm"
	_VALID_URL = "https?://(?:www\\.)?soundgasm\\.net/u/(?P<user>[0-9a-zA-Z_-]+)/(?P<display_id>[0-9a-zA-Z_-]+)"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *SoundgasmIE) Key() string {
	return "Soundgasm"
}

func (self *SoundgasmIE) Name() string {
	return self.IE_NAME
}

func (self *SoundgasmIE) _real_extract(url string) rnt.SDict {
	var (
		audio_id    rnt.OptString
		audio_url   rnt.OptString
		description rnt.OptString
		display_id  rnt.OptString
		mobj        rnt.Match
		title       rnt.OptString
		webpage     string
	)
	mobj = rnt.ReMatch((self).VALIDURL, url, 0)
	display_id = rnt.ReMatchGroupOne(mobj, "display_id")
	webpage = (self).DownloadWebpageURL(url, display_id.Get(), rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{}, nil)
	audio_url = (self).HTMLSearchRegexOne("(?s)m4a\\s*:\\s*([\"\\'])(?P<url>(?:(?!\\1).)+)\\1", webpage, "audio URL", rnt.NoDefault, true, 0, "url")
	title = (self).SearchRegexOne("<div[^>]+\\bclass=[\"\\']jp-title[^>]+>([^<]+)", webpage, "title", display_id, true, 0, nil)
	description = (self).HTMLSearchRegexMulti(τ_conv_Tssω_to_Ls(τ_Tssω{
		Φ0: "(?s)<div[^>]+\\bclass=[\"\\']jp-description[^>]+>(.+?)</div>",
		Φ1: "(?s)<li>Description:\\s(.*?)<\\/li>",
	}), webpage, "description", rnt.NoDefault, false, 0, nil)
	audio_id = (self).SearchRegexOne("/([^/]+)\\.m4a", audio_url.Get(), "audio id", display_id, true, 0, nil)
	return rnt.SDict{
		"id":          audio_id,
		"display_id":  display_id,
		"url":         audio_url,
		"vcodec":      "none",
		"title":       title,
		"description": description,
		"uploader":    rnt.ReMatchGroupOne(mobj, "user"),
	}
}

func (self *SoundgasmIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("Soundgasm", NewSoundgasmIE)
}
