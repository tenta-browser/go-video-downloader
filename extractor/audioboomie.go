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
 * audioboomie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/audioboom.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type AudioBoomIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewAudioBoomIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &AudioBoomIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "AudioBoom"
	_VALID_URL = "https?://(?:www\\.)?audioboom\\.com/(?:boos|posts)/(?P<id>[0-9]+)"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *AudioBoomIE) Key() string {
	return "AudioBoom"
}

func (self *AudioBoomIE) Name() string {
	return self.IE_NAME
}

func (self *AudioBoomIE) _real_extract(url string) rnt.SDict {
	var (
		audio_url    interface{}
		clip         interface{}
		clip_store   interface{}
		clips        interface{}
		description  interface{}
		duration     rnt.OptFloat
		from_clip    func(string) interface{}
		title        interface{}
		uploader     interface{}
		uploader_url interface{}
		video_id     string
		webpage      string
	)
	video_id = (self).MatchID(url)
	webpage = (self).DownloadWebpageURL(url, video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{}, nil)
	clip = nil
	clip_store = (self).ParseJSON((self).SearchRegexOne(rnt.StrFormat2("data-new-clip-store=([\"\\'])(?P<json>{.*?\"clipId\"\\s*:\\s*%s.*?})\\1", video_id), webpage, "clip store", "{}", true, 0, "json").Get(), video_id, nil, false)
	if rnt.IsTruthy(clip_store) {
		clips = rnt.DictGet(τ_cast_α_to_d(clip_store), "clips", nil)
		if rnt.IsTruthy(func() interface{} {
			if v := (clips); !(rnt.IsTruthy(v)) {
				return v
			} else if v := (rnt.IsList(clips)); !(v) {
				return v
			} else {
				return τ_isinstance_d(rnt.UnsafeSubscript(clips, 0))
			}
		}()) {
			clip = rnt.UnsafeSubscript(clips, 0)
		}
	}
	from_clip = func(field string) interface{} {
		if rnt.IsTruthy(clip) {
			return rnt.DictGet(τ_cast_α_to_d(clip), field, nil)
		}
		return nil
	}
	audio_url = func() interface{} {
		if v := (from_clip("clipURLPriorToLoading")); rnt.IsTruthy(v) {
			return v
		} else {
			return (self).OgSearchPropertyOne("audio", webpage, rnt.AsOptString("audio url"), rnt.NoDefault, true)
		}
	}()
	title = func() interface{} {
		if v := (from_clip("title")); rnt.IsTruthy(v) {
			return v
		} else {
			return (self).OgSearchTitle(webpage, rnt.NoDefault, true)
		}
	}()
	description = func() interface{} {
		if v := (from_clip("description")); rnt.IsTruthy(v) {
			return v
		} else {
			return (self).OgSearchDescription(webpage, rnt.NoDefault)
		}
	}()
	duration = rnt.FloatOrNone(func() interface{} {
		if v := (from_clip("duration")); rnt.IsTruthy(v) {
			return v
		} else {
			return (self).HTMLSearchMetaOne("weibo:audio:duration", webpage, rnt.OptString{}, false, rnt.NoDefault, 0)
		}
	}(), 1, 1, rnt.OptFloat{})
	uploader = func() interface{} {
		if v := (from_clip("author")); rnt.IsTruthy(v) {
			return v
		} else {
			return (self).OgSearchPropertyOne("audio:artist", webpage, rnt.AsOptString("uploader"), rnt.NoDefault, false)
		}
	}()
	uploader_url = func() interface{} {
		if v := (from_clip("author_url")); rnt.IsTruthy(v) {
			return v
		} else {
			return (self).HTMLSearchMetaOne("audioboo:channel", webpage, rnt.AsOptString("uploader url"), false, rnt.NoDefault, 0)
		}
	}()
	return rnt.SDict{
		"id":           video_id,
		"url":          audio_url,
		"title":        title,
		"description":  description,
		"duration":     duration,
		"uploader":     uploader,
		"uploader_url": uploader_url,
	}
}

func (self *AudioBoomIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("AudioBoom", NewAudioBoomIE)
}
