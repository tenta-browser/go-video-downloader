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
 * historicfilmsie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/historicfilms.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type HistoricFilmsIE struct {
	*rnt.CommonIE
}

func NewHistoricFilmsIE() rnt.InfoExtractor {
	ret := &HistoricFilmsIE{}
	ret.CommonIE = rnt.NewCommonIE()
	ret.VALIDURL = "https?://(?:www\\.)?historicfilms\\.com/(?:tapes/|play)(?P<id>\\d+)"
	return ret
}

func (self *HistoricFilmsIE) Key() string {
	return "HistoricFilms"
}

func (self *HistoricFilmsIE) Name() string {
	return "HistoricFilms extractor"
}

func (self *HistoricFilmsIE) _real_extract(url string) map[string]interface{} {
	video_id := (self).MatchID(url)
	webpage := (self).DownloadWebpageURL(url, video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, rnt.OptString{}, map[string]interface{}{}, map[string]interface{}{})
	tape_id := (self).SearchRegexMulti([]string{"class=\"tapeId\"[^>]*>([^<]+)<", "tapeId\\s*:\\s*\"([^\"]+)\""}, webpage, "tape id", rnt.NoDefault, true, 0, nil)
	title := (self).OgSearchTitle(webpage, rnt.NoDefault, true)
	description := (self).OgSearchDescription(webpage, rnt.NoDefault)
	thumbnail := func() rnt.OptString {
		if ((self).HTMLSearchMetaOne("thumbnailUrl", webpage, rnt.AsOptString("thumbnails"), false, rnt.NoDefault, 0)).IsSet() && ((self).HTMLSearchMetaOne("thumbnailUrl", webpage, rnt.AsOptString("thumbnails"), false, rnt.NoDefault, 0).Get()) != "" {
			return (self).HTMLSearchMetaOne("thumbnailUrl", webpage, rnt.AsOptString("thumbnails"), false, rnt.NoDefault, 0)
		} else {
			return (self).OgSearchThumbnail(webpage, rnt.NoDefault)
		}
	}()
	duration := rnt.ParseDuration((self).HTMLSearchMetaOne("duration", webpage, rnt.AsOptString("duration"), false, rnt.NoDefault, 0))
	video_url := rnt.StrFormat("http://www.historicfilms.com/video/%s_%s_web.mov", tape_id, video_id)
	return map[string]interface{}{"id": video_id,
		"url":         video_url,
		"title":       title,
		"description": description,
		"thumbnail":   thumbnail,
		"duration":    duration}
}

func (self *HistoricFilmsIE) Extract(url string) (*rnt.VideoResult, error) {
	return rnt.RunExtractor(url, self._real_extract)
}

func init() {
	registerFactory("HistoricFilms", NewHistoricFilmsIE)
}
