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
 * abcotvsclipsie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/abcotvs.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type ABCOTVSClipsIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewABCOTVSClipsIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &ABCOTVSClipsIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "ABCOTVSClips"
	IE_NAME = "abcotvs:clips"
	_VALID_URL = "https?://clips\\.abcotvs\\.com/(?:[^/]+/)*video/(?P<id>\\d+)"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *ABCOTVSClipsIE) Key() string {
	return "ABCOTVSClips"
}

func (self *ABCOTVSClipsIE) Name() string {
	return self.IE_NAME
}

func (self *ABCOTVSClipsIE) _real_extract(url string) rnt.SDict {
	var (
		formats    []rnt.SDict
		title      interface{}
		video_data interface{}
		video_id   string
	)
	video_id = (self).MatchID(url)
	video_data = rnt.UnsafeSubscript(rnt.UnsafeSubscript((self).DownloadJSON(("https://clips.abcotvs.com/vogo/video/getByIds?ids="+video_id), video_id, rnt.AsOptString("Downloading JSON metadata"), rnt.AsOptString("Unable to download JSON metadata"), nil, true, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{}), "results"), 0)
	title = rnt.UnsafeSubscript(video_data, "title")
	formats = (self).ExtractM3U8Formats((rnt.StrSplit(rnt.CastToString(rnt.UnsafeSubscript(video_data, "videoURL")), "?", -(1)))[0], video_id, rnt.AsOptString("mp4"), rnt.AsOptString("m3u8"), nil, rnt.OptString{}, rnt.OptString{}, rnt.OptString{}, true, false)
	formats = (self).SortFormats(formats)
	return rnt.SDict{
		"id":          video_id,
		"title":       title,
		"description": rnt.DictGet(τ_cast_α_to_d(video_data), "description", nil),
		"thumbnail":   rnt.DictGet(τ_cast_α_to_d(video_data), "thumbnailURL", nil),
		"duration":    rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(video_data), "duration", nil), 1, rnt.OptInt{}, 1),
		"timestamp":   rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(video_data), "pubDate", nil), 1, rnt.OptInt{}, 1),
		"formats":     formats,
	}
}

func (self *ABCOTVSClipsIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("ABCOTVSClips", NewABCOTVSClipsIE)
}
