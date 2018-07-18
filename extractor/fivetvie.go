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
 * fivetvie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/fivetv.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type FiveTVIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewFiveTVIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &FiveTVIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "FiveTV"
	_VALID_URL = "(?x)\n                    http://\n                        (?:www\\.)?5-tv\\.ru/\n                        (?:\n                            (?:[^/]+/)+(?P<id>\\d+)|\n                            (?P<path>[^/?#]+)(?:[/?#])?\n                        )\n                    "
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *FiveTVIE) Key() string {
	return "FiveTV"
}

func (self *FiveTVIE) Name() string {
	return self.IE_NAME
}

func (self *FiveTVIE) _real_extract(url string) rnt.SDict {
	var (
		duration  rnt.OptInt
		mobj      rnt.Match
		title     rnt.OptString
		video_id  rnt.OptString
		video_url rnt.OptString
		webpage   string
	)
	mobj = rnt.ReMatch((self).VALIDURL, url, 0)
	video_id = func() rnt.OptString {
		if v := (rnt.ReMatchGroupOne(mobj, "id")); τ_isTruthy_Os(v) {
			return v
		} else {
			return rnt.ReMatchGroupOne(mobj, "path")
		}
	}()
	webpage = (self).DownloadWebpageURL(url, video_id.Get(), rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{}, nil)
	video_url = (self).SearchRegexMulti([]string{"<div[^>]+?class=\"flowplayer[^>]+?data-href=\"([^\"]+)\"", "<a[^>]+?href=\"([^\"]+)\"[^>]+?class=\"videoplayer\""}, webpage, "video url", rnt.NoDefault, true, 0, nil)
	title = func() rnt.OptString {
		if v := ((self).OgSearchTitle(webpage, nil, true)); τ_isTruthy_Os(v) {
			return v
		} else {
			return (self).SearchRegexOne("<title>([^<]+)</title>", webpage, "title", rnt.NoDefault, true, 0, nil)
		}
	}()
	duration = rnt.IntOrNone((self).OgSearchPropertyOne("video:duration", webpage, rnt.AsOptString("duration"), nil, true), 1, rnt.OptInt{}, 1)
	return rnt.SDict{
		"id":          video_id,
		"url":         video_url,
		"title":       title,
		"description": (self).OgSearchDescription(webpage, nil),
		"thumbnail":   (self).OgSearchThumbnail(webpage, nil),
		"duration":    duration,
	}
}

func (self *FiveTVIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("FiveTV", NewFiveTVIE)
}
