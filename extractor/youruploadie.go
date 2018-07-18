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
 * youruploadie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/yourupload.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type YourUploadIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewYourUploadIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &YourUploadIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "YourUpload"
	_VALID_URL = "https?://(?:www\\.)?(?:yourupload\\.com/(?:watch|embed)|embed\\.yourupload\\.com)/(?P<id>[A-Za-z0-9]+)"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *YourUploadIE) Key() string {
	return "YourUpload"
}

func (self *YourUploadIE) Name() string {
	return self.IE_NAME
}

func (self *YourUploadIE) _real_extract(url string) rnt.SDict {
	var (
		embed_url string
		thumbnail rnt.OptString
		title     rnt.OptString
		video_id  string
		video_url rnt.OptString
		webpage   string
	)
	video_id = (self).MatchID(url)
	embed_url = rnt.StrFormat2("http://www.yourupload.com/embed/%s", video_id)
	webpage = (self).DownloadWebpageURL(embed_url, video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{}, nil)
	title = (self).OgSearchTitle(webpage, rnt.NoDefault, true)
	video_url = rnt.UtilURLJoin(embed_url, (self).OgSearchVideoURL(webpage, "video url", true, rnt.NoDefault))
	thumbnail = (self).OgSearchThumbnail(webpage, nil)
	return rnt.SDict{
		"id":        video_id,
		"title":     title,
		"url":       video_url,
		"thumbnail": thumbnail,
		"http_headers": rnt.SDict{
			"Referer": embed_url,
		},
	}
}

func (self *YourUploadIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("YourUpload", NewYourUploadIE)
}
