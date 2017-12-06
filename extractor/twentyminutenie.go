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
 * twentyminutenie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/twentymin.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type TwentyMinutenIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewTwentyMinutenIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &TwentyMinutenIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "TwentyMinuten"
	IE_NAME = "20min"
	_VALID_URL = "(?x)\n                    https?://\n                        (?:www\\.)?20min\\.ch/\n                        (?:\n                            videotv/*\\?.*?\\bvid=|\n                            videoplayer/videoplayer\\.html\\?.*?\\bvideoId@\n                        )\n                        (?P<id>\\d+)\n                    "
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *TwentyMinutenIE) Key() string {
	return "TwentyMinuten"
}

func (self *TwentyMinutenIE) Name() string {
	return self.IE_NAME
}

func (self *TwentyMinutenIE) _real_extract(url string) rnt.SDict {
	var (
		description   interface{}
		dislike_count interface{}
		extract_count func(string) interface{}
		formats       []rnt.SDict
		like_count    interface{}
		thumbnail     interface{}
		title         interface{}
		video         interface{}
		video_id      string
	)
	video_id = (self).MatchID(url)
	video = rnt.UnsafeSubscript((self).DownloadJSON(rnt.StrFormat2("http://api.20min.ch/video/%s/show", video_id), video_id, rnt.AsOptString("Downloading JSON metadata"), rnt.AsOptString("Unable to download JSON metadata"), nil, true, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{}), "content")
	title = rnt.UnsafeSubscript(video, "title")
	formats = func() []rnt.SDict {
		τresult := []rnt.SDict{}
		for _, τel := range τ_enumerate_LTssω([]τ_Tssω{{
			Φ0: "sd",
			Φ1: "",
		}, {
			Φ0: "hd",
			Φ1: "h",
		}}, 0) {
			var (
				format_id string
				p         string
				quality   int
				τmp1      τ_TiTssωω
				τmp2      τ_Tssω
			)
			τmp1 = τel
			quality = (τmp1).Φ0
			τmp2 = (τmp1).Φ1
			format_id = (τmp2).Φ0
			p = (τmp2).Φ1
			τresult = append(τresult, rnt.SDict{
				"format_id": format_id,
				"url":       rnt.StrFormat2("http://podcast.20min-tv.ch/podcast/20min/%s%s.mp4", video_id, p),
				"quality":   quality,
			})
		}
		return τresult
	}()
	formats = (self).SortFormats(formats)
	description = rnt.DictGet(τ_cast_α_to_d(video), "lead", nil)
	thumbnail = rnt.DictGet(τ_cast_α_to_d(video), "thumbnail", nil)
	extract_count = func(kind string) interface{} {
		return rnt.TryGetOne(video, func(x interface{}) interface{} {
			return rnt.IntOrNone(rnt.UnsafeSubscript(rnt.UnsafeSubscript(x, "communityobject"), rnt.StrFormat2("thumbs_%s", kind)), 1, rnt.OptInt{}, 1)
		}, nil)
	}
	like_count = extract_count("up")
	dislike_count = extract_count("down")
	return rnt.SDict{
		"id":            video_id,
		"title":         title,
		"description":   description,
		"thumbnail":     thumbnail,
		"like_count":    like_count,
		"dislike_count": dislike_count,
		"formats":       formats,
	}
}

func (self *TwentyMinutenIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("TwentyMinuten", NewTwentyMinutenIE)
}
