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
 * lrtie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/lrt.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type LRTIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewLRTIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &LRTIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "LRT"
	IE_NAME = "lrt.lt"
	_VALID_URL = "https?://(?:www\\.)?lrt\\.lt/mediateka/irasas/(?P<id>[0-9]+)"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *LRTIE) Key() string {
	return "LRT"
}

func (self *LRTIE) Name() string {
	return self.IE_NAME
}

func (self *LRTIE) _real_extract(url string) rnt.SDict {
	var (
		_           string
		description rnt.OptString
		duration    rnt.OptString
		ext         string
		file_url    string
		formats     []interface{}
		like_count  rnt.OptInt
		thumbnail   rnt.OptString
		title       rnt.OptString
		video_id    string
		view_count  rnt.OptInt
		webpage     string
		τmp1        []string
	)
	video_id = (self).MatchID(url)
	webpage = (self).DownloadWebpageURL(url, video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	title = rnt.RemoveEnd((self).OgSearchTitle(webpage, rnt.NoDefault, true), " - LRT")
	formats = []interface{}{}
	for _, τel := range rnt.ReFindAllMulti("file\\s*:\\s*([\"\\'])(?P<url>(?:(?!\\1).)+)\\1", webpage, 0) {
		τmp1 = τel
		_ = (τmp1)[0]
		file_url = (τmp1)[1]
		ext = rnt.DetermineExt(rnt.AsOptString(file_url), "unknown_video")
		if τ_search_s_notin_Ls(ext, τ_conv_Tssω_to_Ls(τ_Tssω{
			Φ0: "m3u8",
			Φ1: "mp3",
		})) {
			continue
		}
		if (ext) == ("m3u8") && rnt.StrContains(file_url, ".mp3") {
			continue
		}
		if (ext) == ("m3u8") {
			formats = append(formats, τ_conv_Ld_to_Lα((self).ExtractM3U8Formats(file_url, video_id, rnt.AsOptString("mp4"), rnt.AsOptString("m3u8_native"), rnt.OptFloat{}, rnt.OptString{}, rnt.OptString{}, rnt.OptString{}, false, false))...)
		} else {
			if (ext) == ("mp3") {
				formats = append(formats, rnt.SDict{
					"url":    file_url,
					"vcodec": "none",
				})
			}
		}
	}
	formats = τ_conv_Ld_to_Lα((self).SortFormats(τ_conv_Lα_to_Ld(formats)))
	thumbnail = (self).OgSearchThumbnail(webpage, rnt.NoDefault)
	description = (self).OgSearchDescription(webpage, rnt.NoDefault)
	duration = rnt.ParseDuration((self).SearchRegexOne("var\\s+record_len\\s*=\\s*([\"\\'])(?P<duration>[0-9]+:[0-9]+:[0-9]+)\\1", webpage, "duration", nil, true, 0, "duration"))
	view_count = rnt.IntOrNone((self).HTMLSearchRegexOne("<div[^>]+class=([\"\\']).*?record-desc-seen.*?\\1[^>]*>(?P<count>.+?)</div>", webpage, "view count", rnt.NoDefault, false, 0, "count"), 1, rnt.OptInt{}, 1)
	like_count = rnt.IntOrNone((self).SearchRegexOne("<span[^>]+id=([\"\\'])flikesCount.*?\\1>(?P<count>\\d+)<", webpage, "like count", rnt.NoDefault, false, 0, "count"), 1, rnt.OptInt{}, 1)
	return rnt.SDict{
		"id":          video_id,
		"title":       title,
		"formats":     formats,
		"thumbnail":   thumbnail,
		"description": description,
		"duration":    duration,
		"view_count":  view_count,
		"like_count":  like_count,
	}
}

func (self *LRTIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("LRT", NewLRTIE)
}
