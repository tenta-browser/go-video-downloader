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
 * telembie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/telemb.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type TeleMBIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewTeleMBIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &TeleMBIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "TeleMB"
	_VALID_URL = "https?://(?:www\\.)?telemb\\.be/(?P<display_id>.+?)_d_(?P<id>\\d+)\\.html"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *TeleMBIE) Key() string {
	return "TeleMB"
}

func (self *TeleMBIE) Name() string {
	return self.IE_NAME
}

func (self *TeleMBIE) _real_extract(url string) rnt.SDict {
	var (
		description rnt.OptString
		display_id  rnt.OptString
		fmt         rnt.SDict
		formats     []interface{}
		mobj        rnt.Match
		rtmp        rnt.Match
		thumbnail   rnt.OptString
		title       rnt.OptString
		video_id    rnt.OptString
		video_url   string
		webpage     string
	)
	mobj = rnt.ReMatch((self).VALIDURL, url, 0)
	video_id = rnt.ReMatchGroupOne(mobj, "id")
	display_id = rnt.ReMatchGroupOne(mobj, "display_id")
	webpage = (self).DownloadWebpageURL(url, display_id.Get(), rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	formats = []interface{}{}
	for _, τel := range rnt.ReFindAllOne("file\\s*:\\s*\"([^\"]+)\"", webpage, 0) {
		video_url = τel
		fmt = rnt.SDict{
			"url":       video_url,
			"format_id": (rnt.StrSplit(video_url, ":", -(1)))[0],
		}
		rtmp = rnt.ReSearch("^(?P<url>rtmp://[^/]+/(?P<app>.+))/(?P<playpath>mp4:.+)$", video_url, 0)
		if rtmp != nil {
			rnt.DictUpdate(fmt, rnt.SDict{
				"play_path":  rnt.ReMatchGroupOne(rtmp, "playpath"),
				"app":        rnt.ReMatchGroupOne(rtmp, "app"),
				"player_url": "http://p.jwpcdn.com/6/10/jwplayer.flash.swf",
				"page_url":   "http://www.telemb.be",
				"preference": -(1),
			})
		}
		formats = append(formats, fmt)
	}
	formats = τ_conv_Ld_to_Lα((self).SortFormats(τ_conv_Lα_to_Ld(formats)))
	title = rnt.RemoveStart((self).OgSearchTitle(webpage, rnt.NoDefault, true), "TéléMB : ")
	description = (self).HTMLSearchRegexOne("<meta property=\"og:description\" content=\"(.+?)\" />", webpage, "description", rnt.NoDefault, false, 0, nil)
	thumbnail = (self).OgSearchThumbnail(webpage, rnt.NoDefault)
	return rnt.SDict{
		"id":          video_id,
		"display_id":  display_id,
		"title":       title,
		"description": description,
		"thumbnail":   thumbnail,
		"formats":     formats,
	}
}

func (self *TeleMBIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("TeleMB", NewTeleMBIE)
}
