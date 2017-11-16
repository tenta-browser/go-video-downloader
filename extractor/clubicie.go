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
 * clubicie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/clubic.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type ClubicIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewClubicIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &ClubicIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "Clubic"
	_VALID_URL = "https?://(?:www\\.)?clubic\\.com/video/(?:[^/]+/)*video.*-(?P<id>[0-9]+)\\.html"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *ClubicIE) Key() string {
	return "Clubic"
}

func (self *ClubicIE) Name() string {
	return self.IE_NAME
}

func (self *ClubicIE) _real_extract(url string) rnt.SDict {
	var (
		config        interface{}
		formats       []rnt.SDict
		player_page   string
		player_url    string
		quality_order func(string) int
		sources       interface{}
		video_id      string
		video_info    interface{}
	)
	video_id = (self).MatchID(url)
	player_url = rnt.StrFormat2("http://player.m6web.fr/v1/player/clubic/%s.html", video_id)
	player_page = (self).DownloadWebpageURL(player_url, video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	config = (self).ParseJSON((self).SearchRegexOne("(?m)M6\\.Player\\.config\\s*=\\s*(\\{.+?\\});$", player_page, "configuration", rnt.NoDefault, true, 0, nil).Get(), video_id, nil, true)
	video_info = rnt.UnsafeSubscript(config, "videoInfo")
	sources = rnt.UnsafeSubscript(config, "sources")
	quality_order = rnt.Qualities([]string{"sd", "hq"})
	formats = func() []rnt.SDict {
		τresult := []rnt.SDict{}
		for _, τel := range τ_cast_α_to_Lα(sources) {
			var (
				src interface{}
			)
			src = τel
			τresult = append(τresult, rnt.SDict{
				"format_id": rnt.UnsafeSubscript(src, "streamQuality"),
				"url":       rnt.UnsafeSubscript(src, "src"),
				"quality":   quality_order(rnt.CastToString(rnt.UnsafeSubscript(src, "streamQuality"))),
			})
		}
		return τresult
	}()
	formats = (self).SortFormats(formats)
	return rnt.SDict{
		"id":          video_id,
		"title":       rnt.UnsafeSubscript(video_info, "title"),
		"formats":     formats,
		"description": rnt.CleanHTML(rnt.CastToOptString(rnt.DictGet(τ_cast_α_to_d(video_info), "description", nil))),
		"thumbnail":   rnt.DictGet(τ_cast_α_to_d(config), "poster", nil),
	}
}

func (self *ClubicIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("Clubic", NewClubicIE)
}
