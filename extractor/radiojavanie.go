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
 * radiojavanie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/radiojavan.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type RadioJavanIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewRadioJavanIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &RadioJavanIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "RadioJavan"
	_VALID_URL = "https?://(?:www\\.)?radiojavan\\.com/videos/video/(?P<id>[^/]+)/?"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *RadioJavanIE) Key() string {
	return "RadioJavan"
}

func (self *RadioJavanIE) Name() string {
	return self.IE_NAME
}

func (self *RadioJavanIE) _real_extract(url string) rnt.SDict {
	var (
		dislike_count rnt.OptInt
		formats       []rnt.SDict
		like_count    rnt.OptInt
		thumbnail     rnt.OptString
		title         rnt.OptString
		upload_date   rnt.OptString
		video_id      string
		view_count    rnt.OptInt
		webpage       string
	)
	video_id = (self).MatchID(url)
	webpage = (self).DownloadWebpageURL(url, video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	formats = func() []rnt.SDict {
		τresult := []rnt.SDict{}
		for _, τel := range rnt.ReFindAllMulti("RJ\\.video(\\d+)p\\s*=\\s*'/?([^']+)'", webpage, 0) {
			var (
				height     string
				video_path string
				τmp1       []string
			)
			τmp1 = τel
			height = (τmp1)[0]
			video_path = (τmp1)[1]
			τresult = append(τresult, rnt.SDict{
				"url":       rnt.StrFormat2("https://media.rdjavan.com/media/music_video/%s", video_path),
				"format_id": rnt.StrFormat2("%sp", height),
				"height":    rnt.ConvertToInt(height),
			})
		}
		return τresult
	}()
	formats = (self).SortFormats(formats)
	title = (self).OgSearchTitle(webpage, rnt.NoDefault, true)
	thumbnail = (self).OgSearchThumbnail(webpage, rnt.NoDefault)
	upload_date = rnt.UnifiedStrDate((self).SearchRegexOne("class=\"date_added\">Date added: ([^<]+)<", webpage, "upload date", rnt.NoDefault, false, 0, nil), true)
	view_count = rnt.StrToInt((self).SearchRegexOne("class=\"views\">Plays: ([\\d,]+)", webpage, "view count", rnt.NoDefault, false, 0, nil))
	like_count = rnt.StrToInt((self).SearchRegexOne("class=\"rating\">([\\d,]+) likes", webpage, "like count", rnt.NoDefault, false, 0, nil))
	dislike_count = rnt.StrToInt((self).SearchRegexOne("class=\"rating\">([\\d,]+) dislikes", webpage, "dislike count", rnt.NoDefault, false, 0, nil))
	return rnt.SDict{
		"id":            video_id,
		"title":         title,
		"thumbnail":     thumbnail,
		"upload_date":   upload_date,
		"view_count":    view_count,
		"like_count":    like_count,
		"dislike_count": dislike_count,
		"formats":       formats,
	}
}

func (self *RadioJavanIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("RadioJavan", NewRadioJavanIE)
}
