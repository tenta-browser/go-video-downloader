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
 * hitboxie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/hitbox.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type HitboxIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewHitboxIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &HitboxIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "Hitbox"
	IE_NAME = "hitbox"
	_VALID_URL = "https?://(?:www\\.)?(?:hitbox|smashcast)\\.tv/(?:[^/]+/)*videos?/(?P<id>[0-9]+)"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *HitboxIE) Key() string {
	return "Hitbox"
}

func (self *HitboxIE) Name() string {
	return self.IE_NAME
}

func (self *HitboxIE) _extract_metadata(url string, video_id string) rnt.SDict {
	var (
		alt_title   interface{}
		categories  []interface{}
		date        string
		description rnt.OptString
		duration    rnt.OptFloat
		media_type  string
		metadata    interface{}
		thumb_base  string
		thumbs      []rnt.SDict
		timestamp   rnt.OptInt
		title       interface{}
		uploader    interface{}
		video_meta  interface{}
		views       rnt.OptInt
	)
	thumb_base = "https://edge.sf.hitbox.tv"
	metadata = (self).DownloadJSON(rnt.StrFormat2("%s/%s", url, video_id), video_id, rnt.AsOptString("Downloading metadata JSON"), rnt.AsOptString("Unable to download JSON metadata"), nil, true, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	date = "media_live_since"
	media_type = "livestream"
	if (rnt.DictGet(τ_cast_α_to_d(metadata), "media_type", nil)) == ("video") {
		media_type = "video"
		date = "media_date_added"
	}
	video_meta = rnt.UnsafeSubscript(rnt.DictGet(τ_cast_α_to_d(metadata), media_type, []interface{}{}), 0)
	title = rnt.DictGet(τ_cast_α_to_d(video_meta), "media_status", nil)
	alt_title = rnt.DictGet(τ_cast_α_to_d(video_meta), "media_title", nil)
	description = rnt.CleanHTML(rnt.CastToOptString(func() interface{} {
		if v := (rnt.DictGet(τ_cast_α_to_d(video_meta), "media_description", nil)); rnt.IsTruthy(v) {
			return v
		} else {
			return rnt.DictGet(τ_cast_α_to_d(video_meta), "media_description_md", nil)
		}
	}()))
	duration = rnt.FloatOrNone(rnt.DictGet(τ_cast_α_to_d(video_meta), "media_duration", nil), 1, 1, rnt.OptFloat{})
	uploader = rnt.DictGet(τ_cast_α_to_d(video_meta), "media_user_name", nil)
	views = rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(video_meta), "media_views", nil), 1, rnt.OptInt{}, 1)
	timestamp = rnt.ParseISO8601(rnt.CastToOptString(rnt.DictGet(τ_cast_α_to_d(video_meta), date, nil)), " ")
	categories = []interface{}{rnt.DictGet(τ_cast_α_to_d(video_meta), "category_name", nil)}
	thumbs = []rnt.SDict{{
		"url":    (thumb_base + rnt.CastToString(rnt.DictGet(τ_cast_α_to_d(video_meta), "media_thumbnail", nil))),
		"width":  320,
		"height": 180,
	}, {
		"url":    (thumb_base + rnt.CastToString(rnt.DictGet(τ_cast_α_to_d(video_meta), "media_thumbnail_large", nil))),
		"width":  768,
		"height": 432,
	}}
	return rnt.SDict{
		"id":          video_id,
		"title":       title,
		"alt_title":   alt_title,
		"description": description,
		"ext":         "mp4",
		"thumbnails":  thumbs,
		"duration":    duration,
		"uploader":    uploader,
		"view_count":  views,
		"timestamp":   timestamp,
		"categories":  categories,
	}
}

func (self *HitboxIE) _real_extract(url string) rnt.SDict {
	var (
		bitrate       rnt.OptInt
		formats       []interface{}
		label         interface{}
		metadata      rnt.SDict
		player_config interface{}
		video         interface{}
		video_id      string
		video_url     interface{}
	)
	video_id = (self).MatchID(url)
	player_config = (self).DownloadJSON(rnt.StrFormat2("https://www.smashcast.tv/api/player/config/video/%s", video_id), video_id, rnt.AsOptString("Downloading video JSON"), rnt.AsOptString("Unable to download JSON metadata"), nil, true, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	formats = []interface{}{}
	for _, τel := range τ_cast_α_to_Lα(rnt.UnsafeSubscript(rnt.UnsafeSubscript(player_config, "clip"), "bitrates")) {
		video = τel
		label = rnt.DictGet(τ_cast_α_to_d(video), "label", nil)
		if (label) == ("Auto") {
			continue
		}
		video_url = rnt.DictGet(τ_cast_α_to_d(video), "url", nil)
		if !(rnt.IsTruthy(video_url)) {
			continue
		}
		bitrate = rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(video), "bitrate", nil), 1, rnt.OptInt{}, 1)
		if (rnt.DetermineExt(rnt.CastToOptString(video_url), "unknown_video")) == ("m3u8") {
			if !(rnt.StrStartswith(rnt.CastToString(video_url), "http", rnt.OptInt{}, rnt.OptInt{})) {
				continue
			}
			formats = append(formats, rnt.SDict{
				"url":         video_url,
				"ext":         "mp4",
				"tbr":         bitrate,
				"format_note": label,
				"protocol":    "m3u8_native",
			})
		} else {
			formats = append(formats, rnt.SDict{
				"url":         video_url,
				"tbr":         bitrate,
				"format_note": label,
			})
		}
	}
	formats = τ_conv_Ld_to_Lα((self).SortFormats(τ_conv_Lα_to_Ld(formats)))
	metadata = (self)._extract_metadata("https://www.smashcast.tv/api/media/video", video_id)
	(metadata)["formats"] = formats
	return metadata
}

func (self *HitboxIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("Hitbox", NewHitboxIE)
}
