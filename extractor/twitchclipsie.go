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
 * twitchclipsie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/twitch.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type TwitchClipsIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewTwitchClipsIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &TwitchClipsIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "TwitchClips"
	IE_NAME = "twitch:clips"
	_VALID_URL = "https?://clips\\.twitch\\.tv/(?:[^/]+/)*(?P<id>[^/?#&]+)"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *TwitchClipsIE) Key() string {
	return "TwitchClips"
}

func (self *TwitchClipsIE) Name() string {
	return self.IE_NAME
}

func (self *TwitchClipsIE) _real_extract(url string) rnt.SDict {
	var (
		clip     interface{}
		formats  []rnt.SDict
		title    interface{}
		video_id string
		webpage  string
	)
	video_id = (self).MatchID(url)
	webpage = (self).DownloadWebpageURL(url, video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	clip = (self).ParseJSON((self).SearchRegexOne("(?s)clipInfo\\s*=\\s*({.+?});", webpage, "clip info", rnt.NoDefault, true, 0, nil).Get(), video_id, rnt.JsToJSON, true)
	title = func() interface{} {
		if rnt.IsTruthy(rnt.DictGet(τ_cast_α_to_d(clip), "channel_title", nil)) {
			return rnt.DictGet(τ_cast_α_to_d(clip), "channel_title", nil)
		} else {
			return (self).OgSearchTitle(webpage, rnt.NoDefault, true)
		}
	}()
	formats = func() []rnt.SDict {
		τresult := []rnt.SDict{}
		for _, τel := range τ_cast_α_to_Lα(rnt.DictGet(τ_cast_α_to_d(clip), "quality_options", []interface{}{})) {
			var (
				option interface{}
			)
			option = τel
			if !(rnt.IsTruthy(rnt.DictGet(τ_cast_α_to_d(option), "source", nil))) {
				continue
			}
			τresult = append(τresult, rnt.SDict{
				"url":       rnt.UnsafeSubscript(option, "source"),
				"format_id": rnt.DictGet(τ_cast_α_to_d(option), "quality", nil),
				"height":    rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(option), "quality", nil), 1, rnt.OptInt{}, 1),
			})
		}
		return τresult
	}()
	if !(len(formats) > 0) {
		formats = []rnt.SDict{{
			"url": rnt.UnsafeSubscript(clip, "clip_video_url"),
		}}
	}
	formats = (self).SortFormats(formats)
	return rnt.SDict{
		"id":        video_id,
		"title":     title,
		"thumbnail": (self).OgSearchThumbnail(webpage, rnt.NoDefault),
		"creator": func() interface{} {
			if rnt.IsTruthy(rnt.DictGet(τ_cast_α_to_d(clip), "broadcaster_display_name", nil)) {
				return rnt.DictGet(τ_cast_α_to_d(clip), "broadcaster_display_name", nil)
			} else {
				return rnt.DictGet(τ_cast_α_to_d(clip), "broadcaster_login", nil)
			}
		}(),
		"uploader":    rnt.DictGet(τ_cast_α_to_d(clip), "curator_login", nil),
		"uploader_id": rnt.DictGet(τ_cast_α_to_d(clip), "curator_display_name", nil),
		"formats":     formats,
	}
}

func (self *TwitchClipsIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("TwitchClips", NewTwitchClipsIE)
}
