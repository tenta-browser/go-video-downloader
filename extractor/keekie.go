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
 * keekie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/keek.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type KeekIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewKeekIE() rnt.InfoExtractor {
	ret := &KeekIE{}
	ret.CommonIE = rnt.NewCommonIE()
	ret.VALIDURL = "https?://(?:www\\.)?keek\\.com/keek/(?P<id>\\w+)"
	ret.IE_NAME = "keek"
	return ret
}

func (self *KeekIE) Key() string {
	return "Keek"
}

func (self *KeekIE) Name() string {
	return self.IE_NAME
}

func (self *KeekIE) _real_extract(url string) map[string]interface{} {
	video_id := (self).MatchID(url)
	webpage := (self).DownloadWebpageURL(url, video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, rnt.OptString{}, map[string]interface{}{}, map[string]interface{}{})
	return map[string]interface{}{"id": video_id,
		"url":         (self).OgSearchVideoURL(webpage, "video url", true, rnt.NoDefault),
		"ext":         "mp4",
		"title":       rnt.StrStrip((self).OgSearchDescription(webpage, rnt.NoDefault).Get(), ""),
		"thumbnail":   (self).OgSearchThumbnail(webpage, rnt.NoDefault),
		"uploader":    (self).SearchRegexOne("data-username=([\"\\'])(?P<uploader>.+?)\\1", webpage, "uploader", rnt.NoDefault, false, 0, "uploader"),
		"uploader_id": (self).SearchRegexOne("data-user-id=([\"\\'])(?P<uploader_id>.+?)\\1", webpage, "uploader id", rnt.NoDefault, false, 0, "uploader_id")}
}

func (self *KeekIE) Extract(url string) (*rnt.VideoResult, error) {
	return rnt.RunExtractor(url, self._real_extract)
}

func init() {
	registerFactory("Keek", NewKeekIE)
}
