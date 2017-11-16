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
 * movingimageie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/movingimage.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type MovingImageIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewMovingImageIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &MovingImageIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "MovingImage"
	_VALID_URL = "https?://movingimage\\.nls\\.uk/film/(?P<id>\\d+)"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *MovingImageIE) Key() string {
	return "MovingImage"
}

func (self *MovingImageIE) Name() string {
	return self.IE_NAME
}

func (self *MovingImageIE) _real_extract(url string) rnt.SDict {
	var (
		description  rnt.OptString
		duration     rnt.OptString
		formats      []rnt.SDict
		search_field func(string, bool) rnt.OptString
		thumbnail    rnt.OptString
		title        string
		video_id     string
		webpage      string
	)
	video_id = (self).MatchID(url)
	webpage = (self).DownloadWebpageURL(url, video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	formats = (self).ExtractM3U8Formats((self).HTMLSearchRegexOne("file\\s*:\\s*\"([^\"]+)\"", webpage, "m3u8 manifest URL", rnt.NoDefault, true, 0, nil).Get(), video_id, rnt.AsOptString("mp4"), rnt.AsOptString("m3u8_native"), rnt.OptFloat{}, rnt.OptString{}, rnt.OptString{}, rnt.OptString{}, true, false)
	search_field = func(field_name string, fatal bool) rnt.OptString {
		return (self).SearchRegexOne(rnt.StrFormat2("<span\\s+class=\"field_title\">%s:</span>\\s*<span\\s+class=\"field_content\">([^<]+)</span>", field_name), webpage, "title", rnt.NoDefault, fatal, 0, nil)
	}
	title = rnt.StrStrip(rnt.UnescapeHTML(search_field("Title", true)).Get(), "()[]")
	description = rnt.UnescapeHTML(search_field("Description", false))
	duration = rnt.ParseDuration(search_field("Running time", false))
	thumbnail = (self).SearchRegexOne("image\\s*:\\s*'([^']+)'", webpage, "thumbnail", rnt.NoDefault, false, 0, nil)
	return rnt.SDict{
		"id":          video_id,
		"formats":     formats,
		"title":       title,
		"description": description,
		"duration":    duration,
		"thumbnail":   thumbnail,
	}
}

func (self *MovingImageIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("MovingImage", NewMovingImageIE)
}
