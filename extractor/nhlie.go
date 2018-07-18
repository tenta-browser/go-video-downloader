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
 * nhlie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/nhl.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type NHLIE struct {
	*rnt.CommonIE
	IE_NAME         string
	_CONTENT_DOMAIN string
}

func NewNHLIE() rnt.InfoExtractor {
	var (
		IE_NAME         string
		_CONTENT_DOMAIN string
		_VALID_URL      string
	)
	self := &NHLIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "NHL"
	IE_NAME = "nhl.com"
	_VALID_URL = "https?://(?:www\\.)?(?P<site>nhl|wch2016)\\.com/(?:[^/]+/)*c-(?P<id>\\d+)"
	_CONTENT_DOMAIN = "nhl.bamcontent.com"
	self.IE_NAME = IE_NAME
	self._CONTENT_DOMAIN = _CONTENT_DOMAIN
	self.VALIDURL = _VALID_URL
	return self
}

func (self *NHLIE) Key() string {
	return "NHL"
}

func (self *NHLIE) Name() string {
	return self.IE_NAME
}

func (self *NHLIE) _real_extract(url string) rnt.SDict {
	var (
		cuts           interface{}
		ext            string
		formats        []interface{}
		height         rnt.OptInt
		m3u8_formats   []rnt.SDict
		playback       interface{}
		playback_url   interface{}
		site           rnt.OptString
		thumbnail_data interface{}
		thumbnail_url  interface{}
		thumbnails     []interface{}
		title          interface{}
		tmp_id         rnt.OptString
		video          interface{}
		video_data     interface{}
		video_id       string
		videos         interface{}
		τmp1           []rnt.OptString
	)
	τmp1 = rnt.ReMatchGroups(rnt.ReMatch((self).VALIDURL, url, 0), rnt.OptString{})
	site = (τmp1)[0]
	tmp_id = (τmp1)[1]
	video_data = (self).DownloadJSON(rnt.StrFormat2("https://%s/%s/%sid/v1/%s/details/web-v1.json", (self)._CONTENT_DOMAIN, rnt.StrSlice(site.Get(), rnt.OptInt{}, rnt.AsOptInt(3), rnt.OptInt{}), func() string {
		if (site.Get()) == ("mlb") {
			return "item/"
		} else {
			return ""
		}
	}(), tmp_id), tmp_id.Get(), rnt.AsOptString("Downloading JSON metadata"), rnt.AsOptString("Unable to download JSON metadata"), nil, true, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{}, nil)
	if (rnt.DictGet(τ_cast_α_to_d(video_data), "type", nil)) != ("video") {
		video_data = rnt.UnsafeSubscript(video_data, "media")
		video = rnt.DictGet(τ_cast_α_to_d(video_data), "video", nil)
		if rnt.IsTruthy(video) {
			video_data = video
		} else {
			videos = rnt.DictGet(τ_cast_α_to_d(video_data), "videos", nil)
			if rnt.IsTruthy(videos) {
				video_data = rnt.UnsafeSubscript(videos, 0)
			}
		}
	}
	video_id = rnt.ConvertToString(rnt.UnsafeSubscript(video_data, "id"))
	title = rnt.UnsafeSubscript(video_data, "title")
	formats = []interface{}{}
	for _, τel := range τ_cast_α_to_Lα(rnt.DictGet(τ_cast_α_to_d(video_data), "playbacks", []interface{}{})) {
		playback = τel
		playback_url = rnt.DictGet(τ_cast_α_to_d(playback), "url", nil)
		if !(rnt.IsTruthy(playback_url)) {
			continue
		}
		ext = rnt.DetermineExt(rnt.CastToOptString(playback_url), "unknown_video")
		if (ext) == ("m3u8") {
			m3u8_formats = (self).ExtractM3U8Formats(rnt.CastToString(playback_url), video_id, rnt.AsOptString("mp4"), rnt.AsOptString("m3u8_native"), nil, rnt.CastToOptString(rnt.DictGet(τ_cast_α_to_d(playback), "name", "hls")), rnt.OptString{}, rnt.OptString{}, false, false)
			m3u8_formats = (self).CheckFormats(m3u8_formats, video_id)
			formats = append(formats, τ_conv_Ld_to_Lα(m3u8_formats)...)
		} else {
			height = rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(playback), "height", nil), 1, rnt.OptInt{}, 1)
			formats = append(formats, rnt.SDict{
				"format_id": rnt.DictGet(τ_cast_α_to_d(playback), "name", ("http" + func() string {
					if τ_isTruthy_Oi(height) {
						return rnt.StrFormat2("-%dp", height)
					} else {
						return ""
					}
				}())),
				"url":    playback_url,
				"width":  rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(playback), "width", nil), 1, rnt.OptInt{}, 1),
				"height": height,
				"tbr":    rnt.IntOrNone((self).SearchRegexOne("_(\\d+)[kK]", rnt.CastToString(playback_url), "bitrate", nil, true, 0, nil), 1, rnt.OptInt{}, 1),
			})
		}
	}
	formats = τ_conv_Ld_to_Lα((self).SortFormats(τ_conv_Lα_to_Ld(formats)))
	thumbnails = []interface{}{}
	cuts = func() interface{} {
		if v := (rnt.DictGet(τ_cast_α_to_d(rnt.DictGet(τ_cast_α_to_d(video_data), "image", rnt.SDict{})), "cuts", nil)); rnt.IsTruthy(v) {
			return v
		} else {
			return []interface{}{}
		}
	}()
	if τ_isinstance_d(cuts) {
		cuts = rnt.DictValues(τ_cast_α_to_d(cuts))
	}
	for _, τel := range τ_cast_α_to_Lα(cuts) {
		thumbnail_data = τel
		thumbnail_url = rnt.DictGet(τ_cast_α_to_d(thumbnail_data), "src", nil)
		if !(rnt.IsTruthy(thumbnail_url)) {
			continue
		}
		thumbnails = append(thumbnails, rnt.SDict{
			"url":    thumbnail_url,
			"width":  rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(thumbnail_data), "width", nil), 1, rnt.OptInt{}, 1),
			"height": rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(thumbnail_data), "height", nil), 1, rnt.OptInt{}, 1),
		})
	}
	return rnt.SDict{
		"id":          video_id,
		"title":       title,
		"description": rnt.DictGet(τ_cast_α_to_d(video_data), "description", nil),
		"timestamp":   rnt.ParseISO8601(rnt.CastToOptString(rnt.DictGet(τ_cast_α_to_d(video_data), "date", nil)), "T"),
		"duration":    rnt.ParseDuration(rnt.CastToOptString(rnt.DictGet(τ_cast_α_to_d(video_data), "duration", nil))),
		"thumbnails":  thumbnails,
		"formats":     formats,
	}
}

func (self *NHLIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("NHL", NewNHLIE)
}
