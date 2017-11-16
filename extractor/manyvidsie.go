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
 * manyvidsie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/manyvids.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type ManyVidsIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewManyVidsIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &ManyVidsIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "ManyVids"
	_VALID_URL = "(?i)https?://(?:www\\.)?manyvids\\.com/video/(?P<id>\\d+)"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *ManyVidsIE) Key() string {
	return "ManyVids"
}

func (self *ManyVidsIE) Name() string {
	return self.IE_NAME
}

func (self *ManyVidsIE) _real_extract(url string) rnt.SDict {
	var (
		like_count rnt.OptInt
		title      string
		video_id   string
		video_url  rnt.OptString
		view_count rnt.OptInt
		webpage    string
	)
	video_id = (self).MatchID(url)
	webpage = (self).DownloadWebpageURL(url, video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	video_url = (self).SearchRegexOne("data-(?:video-filepath|meta-video)\\s*=s*([\"\\'])(?P<url>(?:(?!\\1).)+)\\1", webpage, "video URL", rnt.NoDefault, true, 0, "url")
	title = rnt.StrFormat2("%s (Preview)", (self).HTMLSearchRegexOne("<h2[^>]+class=\"m-a-0\"[^>]*>([^<]+)", webpage, "title", rnt.NoDefault, true, 0, nil))
	like_count = rnt.IntOrNone((self).SearchRegexOne("data-likes=[\"\\'](\\d+)", webpage, "like count", nil, true, 0, nil), 1, rnt.OptInt{}, 1)
	view_count = rnt.IntOrNone((self).HTMLSearchRegexOne("(?s)<span[^>]+class=\"views-wrapper\"[^>]*>(.+?)</span", webpage, "view count", nil, true, 0, nil), 1, rnt.OptInt{}, 1)
	return rnt.SDict{
		"id":         video_id,
		"title":      title,
		"view_count": view_count,
		"like_count": like_count,
		"formats": []rnt.SDict{{
			"url": video_url,
		}},
	}
}

func (self *ManyVidsIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("ManyVids", NewManyVidsIE)
}
