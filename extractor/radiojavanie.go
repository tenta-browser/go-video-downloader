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
}

func NewRadioJavanIE() rnt.InfoExtractor {
	ret := &RadioJavanIE{}
	ret.CommonIE = rnt.NewCommonIE()
	ret.VALIDURL = "https?://(?:www\\.)?radiojavan\\.com/videos/video/(?P<id>[^/]+)/?"
	return ret
}

func (self *RadioJavanIE) Key() string {
	return "RadioJavan"
}

func (self *RadioJavanIE) Name() string {
	return "RadioJavan extractor"
}

func (self *RadioJavanIE) _real_extract(url string) map[string]interface{} {
	video_id := (self).MatchID(url)
	webpage := (self).DownloadWebpageURL(url, video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, rnt.OptString{}, map[string]interface{}{}, map[string]interface{}{})
	formats := func() []map[string]interface{} {
		τmp3 := []map[string]interface{}{}
		for _, τmp1 := range rnt.ReFindAllMulti("RJ\\.video(\\d+)p\\s*=\\s*'/?([^']+)'", webpage, 0) {
			τmp2 := τmp1
			height := (τmp2)[0]
			video_path := (τmp2)[1]

			τmp3 = append(τmp3, map[string]interface{}{"url": rnt.StrFormat("https://media.rdjavan.com/media/music_video/%s", video_path),
				"format_id": rnt.StrFormat("%sp", height),
				"height":    rnt.ConvertToInt(height)})
		}
		return τmp3
	}()
	(self).SortFormats(func() []interface{} {
		τmp2 := []interface{}{}
		for _, τmp1 := range formats {
			τmp2 = append(τmp2, τmp1)
		}
		return τmp2
	}())
	title := (self).OgSearchTitle(webpage, rnt.NoDefault, true)
	thumbnail := (self).OgSearchThumbnail(webpage, rnt.NoDefault)
	upload_date := rnt.UnifiedStrDate((self).SearchRegexOne("class=\"date_added\">Date added: ([^<]+)<", webpage, "upload date", rnt.NoDefault, false, 0, nil), true)
	view_count := rnt.StrToInt((self).SearchRegexOne("class=\"views\">Plays: ([\\d,]+)", webpage, "view count", rnt.NoDefault, false, 0, nil))
	like_count := rnt.StrToInt((self).SearchRegexOne("class=\"rating\">([\\d,]+) likes", webpage, "like count", rnt.NoDefault, false, 0, nil))
	dislike_count := rnt.StrToInt((self).SearchRegexOne("class=\"rating\">([\\d,]+) dislikes", webpage, "dislike count", rnt.NoDefault, false, 0, nil))
	return map[string]interface{}{"id": video_id,
		"title":         title,
		"thumbnail":     thumbnail,
		"upload_date":   upload_date,
		"view_count":    view_count,
		"like_count":    like_count,
		"dislike_count": dislike_count,
		"formats":       formats}
}

func (self *RadioJavanIE) Extract(url string) (*rnt.VideoResult, error) {
	return rnt.RunExtractor(url, self._real_extract)
}

func init() {
	registerFactory("RadioJavan", NewRadioJavanIE)
}
