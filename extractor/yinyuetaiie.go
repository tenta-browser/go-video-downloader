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
 * yinyuetaiie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/yinyuetai.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type YinYueTaiIE struct {
	*rnt.CommonIE
	IE_DESC string
	IE_NAME string
}

func NewYinYueTaiIE() rnt.InfoExtractor {
	var (
		IE_DESC    string
		IE_NAME    string
		_VALID_URL string
	)
	self := &YinYueTaiIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "YinYueTai"
	IE_NAME = "yinyuetai:video"
	IE_DESC = "音悦Tai"
	_VALID_URL = "https?://v\\.yinyuetai\\.com/video(?:/h5)?/(?P<id>[0-9]+)"
	self.IE_DESC = IE_DESC
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *YinYueTaiIE) Key() string {
	return "YinYueTai"
}

func (self *YinYueTaiIE) Name() string {
	return self.IE_NAME
}

func (self *YinYueTaiIE) _real_extract(url string) rnt.SDict {
	var (
		formats  []rnt.SDict
		info     interface{}
		video_id string
	)
	video_id = (self).MatchID(url)
	info = rnt.UnsafeSubscript(rnt.UnsafeSubscript((self).DownloadJSON(rnt.StrFormat2("http://ext.yinyuetai.com/main/get-h-mv-info?json=true&videoId=%s", video_id), video_id, rnt.AsOptString("Downloading mv info"), rnt.AsOptString("Unable to download JSON metadata"), nil, true, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{}), "videoInfo"), "coreVideoInfo")
	if rnt.IsTruthy(rnt.UnsafeSubscript(info, "error")) {
		panic(rnt.PyExtractorError(rnt.CastToString(rnt.UnsafeSubscript(info, "errorMsg")), true, rnt.OptString{}))
	}
	formats = func() []rnt.SDict {
		τresult := []rnt.SDict{}
		for _, τel := range τ_cast_α_to_Lα(rnt.UnsafeSubscript(info, "videoUrlModels")) {
			var (
				format_info interface{}
			)
			format_info = τel
			τresult = append(τresult, rnt.SDict{
				"url":       rnt.UnsafeSubscript(format_info, "videoUrl"),
				"format_id": rnt.UnsafeSubscript(format_info, "qualityLevel"),
				"format":    rnt.DictGet(τ_cast_α_to_d(format_info), "qualityLevelName", nil),
				"filesize":  rnt.DictGet(τ_cast_α_to_d(format_info), "fileSize", nil),
				"ext":       "mp4",
				"tbr":       rnt.DictGet(τ_cast_α_to_d(format_info), "bitrate", nil),
			})
		}
		return τresult
	}()
	formats = (self).SortFormats(formats)
	return rnt.SDict{
		"id":        video_id,
		"title":     rnt.UnsafeSubscript(info, "videoName"),
		"thumbnail": rnt.DictGet(τ_cast_α_to_d(info), "bigHeadImage", nil),
		"creator":   rnt.DictGet(τ_cast_α_to_d(info), "artistNames", nil),
		"duration":  rnt.DictGet(τ_cast_α_to_d(info), "duration", nil),
		"formats":   formats,
	}
}

func (self *YinYueTaiIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("YinYueTai", NewYinYueTaiIE)
}
