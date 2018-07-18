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
 * vidioie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/vidio.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type VidioIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewVidioIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &VidioIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "Vidio"
	_VALID_URL = "https?://(?:www\\.)?vidio\\.com/watch/(?P<id>\\d+)-(?P<display_id>[^/?#&]+)"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *VidioIE) Key() string {
	return "Vidio"
}

func (self *VidioIE) Name() string {
	return self.IE_NAME
}

func (self *VidioIE) _real_extract(url string) rnt.SDict {
	var (
		clip       interface{}
		clips      interface{}
		display_id rnt.OptString
		duration   interface{}
		formats    []rnt.SDict
		like_count rnt.OptInt
		m3u8_url   interface{}
		mobj       rnt.Match
		thumbnail  interface{}
		title      rnt.OptString
		video_id   rnt.OptString
		webpage    string
		τmp1       τ_TOsOsω
		τmp2       []interface{}
	)
	mobj = rnt.ReMatch((self).VALIDURL, url, 0)
	τmp1 = rnt.ReMatchGroupTwo(mobj, "id", "display_id")
	video_id = (τmp1).Φ0
	display_id = (τmp1).Φ1
	webpage = (self).DownloadWebpageURL(url, display_id.Get(), rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{}, nil)
	title = (self).OgSearchTitle(webpage, rnt.NoDefault, true)
	τmp2 = τ_multiply_L0_by_i([]interface{}{nil}, 3)
	m3u8_url = (τmp2)[0]
	duration = (τmp2)[1]
	thumbnail = (τmp2)[2]
	clips = (self).ParseJSON((self).HTMLSearchRegexOne("data-json-clips\\s*=\\s*([\"\\'])(?P<data>\\[.+?\\])\\1", webpage, "video data", "[]", true, 0, "data").Get(), display_id.Get(), nil, false)
	if rnt.IsTruthy(clips) {
		clip = rnt.UnsafeSubscript(clips, 0)
		m3u8_url = rnt.DictGet(τ_cast_α_to_d(rnt.UnsafeSubscript(rnt.DictGet(τ_cast_α_to_d(clip), "sources", []rnt.SDict{{}}), 0)), "file", nil)
		duration = rnt.DictGet(τ_cast_α_to_d(clip), "clip_duration", nil)
		thumbnail = rnt.DictGet(τ_cast_α_to_d(clip), "image", nil)
	}
	m3u8_url = func() interface{} {
		if v := (m3u8_url); rnt.IsTruthy(v) {
			return v
		} else {
			return (self).SearchRegexOne("data(?:-vjs)?-clip-hls-url=([\"\\'])(?P<url>(?:(?!\\1).)+)\\1", webpage, "hls url", rnt.NoDefault, true, 0, "url")
		}
	}()
	formats = (self).ExtractM3U8Formats(rnt.CastToString(m3u8_url), display_id.Get(), rnt.AsOptString("mp4"), rnt.AsOptString("m3u8_native"), nil, rnt.OptString{}, rnt.OptString{}, rnt.OptString{}, true, false)
	formats = (self).SortFormats(formats)
	duration = rnt.IntOrNone(func() interface{} {
		if v := (duration); rnt.IsTruthy(v) {
			return v
		} else {
			return (self).SearchRegexOne("data-video-duration=([\"\\'])(?P<duration>\\d+)\\1", webpage, "duration", rnt.NoDefault, false, 0, "duration")
		}
	}(), 1, rnt.OptInt{}, 1)
	thumbnail = func() interface{} {
		if v := (thumbnail); rnt.IsTruthy(v) {
			return v
		} else {
			return (self).OgSearchThumbnail(webpage, rnt.NoDefault)
		}
	}()
	like_count = rnt.IntOrNone((self).SearchRegexMulti(τ_conv_Tssω_to_Ls(τ_Tssω{
		Φ0: "<span[^>]+data-comment-vote-count=[\"\\'](\\d+)",
		Φ1: "<span[^>]+class=[\"\\'].*?\\blike(?:__|-)count\\b.*?[\"\\'][^>]*>\\s*(\\d+)",
	}), webpage, "like count", rnt.NoDefault, false, 0, nil), 1, rnt.OptInt{}, 1)
	return rnt.SDict{
		"id":          video_id,
		"display_id":  display_id,
		"title":       title,
		"description": (self).OgSearchDescription(webpage, rnt.NoDefault),
		"thumbnail":   thumbnail,
		"duration":    duration,
		"like_count":  like_count,
		"formats":     formats,
	}
}

func (self *VidioIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("Vidio", NewVidioIE)
}
