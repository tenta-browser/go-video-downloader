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
 * rockstargamesie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/rockstargames.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type RockstarGamesIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewRockstarGamesIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &RockstarGamesIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "RockstarGames"
	_VALID_URL = "https?://(?:www\\.)?rockstargames\\.com/videos(?:/video/|#?/?\\?.*\\bvideo=)(?P<id>\\d+)"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *RockstarGamesIE) Key() string {
	return "RockstarGames"
}

func (self *RockstarGamesIE) Name() string {
	return self.IE_NAME
}

func (self *RockstarGamesIE) _real_extract(url string) rnt.SDict {
	var (
		formats    []interface{}
		height     rnt.OptInt
		resolution interface{}
		title      interface{}
		video      interface{}
		video_id   string
		youtube_id interface{}
	)
	video_id = (self).MatchID(url)
	video = rnt.UnsafeSubscript((self).DownloadJSON("https://www.rockstargames.com/videoplayer/videos/get-video.json", video_id, rnt.AsOptString("Downloading JSON metadata"), rnt.AsOptString("Unable to download JSON metadata"), nil, true, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{
		"id":     video_id,
		"locale": "en_us",
	}), "video")
	title = rnt.UnsafeSubscript(video, "title")
	formats = []interface{}{}
	for _, τel := range τ_cast_α_to_Lα(rnt.UnsafeSubscript(rnt.UnsafeSubscript(video, "files_processed"), "video/mp4")) {
		video = τel
		if !(rnt.IsTruthy(rnt.DictGet(τ_cast_α_to_d(video), "src", nil))) {
			continue
		}
		resolution = rnt.DictGet(τ_cast_α_to_d(video), "resolution", nil)
		height = rnt.IntOrNone((self).SearchRegexOne("^(\\d+)[pP]$", rnt.CastToString(func() interface{} {
			if v := (resolution); rnt.IsTruthy(v) {
				return v
			} else {
				return ""
			}
		}()), "height", nil, true, 0, nil), 1, rnt.OptInt{}, 1)
		formats = append(formats, rnt.SDict{
			"url":       (self).ProtoRelativeURL(rnt.CastToOptString(rnt.UnsafeSubscript(video, "src")), rnt.OptString{}),
			"format_id": resolution,
			"height":    height,
		})
	}
	if !(len(formats) > 0) {
		youtube_id = rnt.DictGet(τ_cast_α_to_d(video), "youtube_id", nil)
		if rnt.IsTruthy(youtube_id) {
			return (self).URLResult(rnt.CastToString(youtube_id), rnt.AsOptString("Youtube"), rnt.OptString{}, rnt.OptString{})
		}
	}
	formats = τ_conv_Ld_to_Lα((self).SortFormats(τ_conv_Lα_to_Ld(formats)))
	return rnt.SDict{
		"id":          video_id,
		"title":       title,
		"description": rnt.DictGet(τ_cast_α_to_d(video), "description", nil),
		"thumbnail":   (self).ProtoRelativeURL(rnt.CastToOptString(rnt.DictGet(τ_cast_α_to_d(video), "screencap", nil)), rnt.OptString{}),
		"timestamp":   rnt.ParseISO8601(rnt.CastToOptString(rnt.DictGet(τ_cast_α_to_d(video), "created", nil)), "T"),
		"formats":     formats,
	}
}

func (self *RockstarGamesIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("RockstarGames", NewRockstarGamesIE)
}
