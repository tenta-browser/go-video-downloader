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
 * redtubeie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/redtube.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type RedTubeIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewRedTubeIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &RedTubeIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "RedTube"
	_VALID_URL = "https?://(?:(?:www\\.)?redtube\\.com/|embed\\.redtube\\.com/\\?.*?\\bid=)(?P<id>[0-9]+)"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *RedTubeIE) Key() string {
	return "RedTube"
}

func (self *RedTubeIE) Name() string {
	return self.IE_NAME
}

func (self *RedTubeIE) _real_extract(url string) rnt.SDict {
	var (
		age_limit   int
		duration    rnt.OptInt
		format_id   interface{}
		format_url  interface{}
		formats     []interface{}
		media       interface{}
		medias      interface{}
		sources     interface{}
		thumbnail   rnt.OptString
		title       rnt.OptString
		upload_date rnt.OptString
		video_id    string
		video_url   rnt.OptString
		view_count  rnt.OptInt
		webpage     string
		τmp1        τ_Tsαω
	)
	video_id = (self).MatchID(url)
	webpage = (self).DownloadWebpageURL(rnt.StrFormat2("http://www.redtube.com/%s", video_id), video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{}, nil)
	if τ_any_Lb(func() []bool {
		τresult := []bool{}
		for _, τel := range []string{"video-deleted-info", ">This video has been removed"} {
			var (
				s string
			)
			s = τel
			τresult = append(τresult, rnt.StrContains(webpage, s))
		}
		return τresult
	}()) {
		panic(rnt.PyExtractorError(rnt.StrFormat2("Video %s has been removed", video_id), true, rnt.OptString{}))
	}
	title = func() rnt.OptString {
		if v := ((self).HTMLSearchRegexMulti(τ_conv_Tssω_to_Ls(τ_Tssω{
			Φ0: "<h(\\d)[^>]+class=\"(?:video_title_text|videoTitle)[^\"]*\">(?P<title>(?:(?!\\1).)+)</h\\1>",
			Φ1: "(?:videoTitle|title)\\s*:\\s*([\"\\'])(?P<title>(?:(?!\\1).)+)\\1",
		}), webpage, "title", nil, true, 0, "title")); τ_isTruthy_Os(v) {
			return v
		} else {
			return (self).OgSearchTitle(webpage, rnt.NoDefault, true)
		}
	}()
	formats = []interface{}{}
	sources = (self).ParseJSON((self).SearchRegexOne("sources\\s*:\\s*({.+?})", webpage, "source", "{}", true, 0, nil).Get(), video_id, nil, false)
	if rnt.IsTruthy(func() interface{} {
		if v := (sources); !(rnt.IsTruthy(v)) {
			return v
		} else {
			return τ_isinstance_d(sources)
		}
	}()) {
		for _, τel := range rnt.DictItems(τ_cast_α_to_d(sources)) {
			τmp1 = τel
			format_id = (τmp1).Φ0
			format_url = (τmp1).Φ1
			if rnt.IsTruthy(format_url) {
				formats = append(formats, rnt.SDict{
					"url":       format_url,
					"format_id": format_id,
					"height":    rnt.IntOrNone(format_id, 1, rnt.OptInt{}, 1),
				})
			}
		}
	}
	medias = (self).ParseJSON((self).SearchRegexOne("mediaDefinition\\s*:\\s*(\\[.+?\\])", webpage, "media definitions", "{}", true, 0, nil).Get(), video_id, nil, false)
	if rnt.IsTruthy(func() interface{} {
		if v := (medias); !(rnt.IsTruthy(v)) {
			return v
		} else {
			return rnt.IsList(medias)
		}
	}()) {
		for _, τel := range τ_cast_α_to_Lα(medias) {
			media = τel
			format_url = rnt.URLOrNone(rnt.CastToOptString(rnt.DictGet(τ_cast_α_to_d(media), "videoUrl", nil)))
			if !(rnt.IsTruthy(format_url)) {
				continue
			}
			format_id = rnt.DictGet(τ_cast_α_to_d(media), "quality", nil)
			formats = append(formats, rnt.SDict{
				"url":       format_url,
				"format_id": format_id,
				"height":    rnt.IntOrNone(format_id, 1, rnt.OptInt{}, 1),
			})
		}
	}
	if !(len(formats) > 0) {
		video_url = (self).HTMLSearchRegexOne("<source src=\"(.+?)\" type=\"video/mp4\">", webpage, "video URL", rnt.NoDefault, true, 0, nil)
		formats = append(formats, rnt.SDict{
			"url": video_url,
		})
	}
	formats = τ_conv_Ld_to_Lα((self).SortFormats(τ_conv_Lα_to_Ld(formats)))
	thumbnail = (self).OgSearchThumbnail(webpage, rnt.NoDefault)
	upload_date = rnt.UnifiedStrDate((self).SearchRegexOne("<span[^>]+>ADDED ([^<]+)<", webpage, "upload date", rnt.NoDefault, false, 0, nil), true)
	duration = rnt.IntOrNone(func() rnt.OptString {
		if v := ((self).OgSearchPropertyOne("video:duration", webpage, rnt.OptString{}, nil, true)); τ_isTruthy_Os(v) {
			return v
		} else {
			return (self).SearchRegexOne("videoDuration\\s*:\\s*(\\d+)", webpage, "duration", nil, true, 0, nil)
		}
	}(), 1, rnt.OptInt{}, 1)
	view_count = rnt.StrToInt((self).SearchRegexMulti(τ_conv_Tssω_to_Ls(τ_Tssω{
		Φ0: "<div[^>]*>Views</div>\\s*<div[^>]*>\\s*([\\d,.]+)",
		Φ1: "<span[^>]*>VIEWS</span>\\s*</td>\\s*<td>\\s*([\\d,.]+)",
	}), webpage, "view count", rnt.NoDefault, false, 0, nil))
	age_limit = 18
	return rnt.SDict{
		"id":          video_id,
		"ext":         "mp4",
		"title":       title,
		"thumbnail":   thumbnail,
		"upload_date": upload_date,
		"duration":    duration,
		"view_count":  view_count,
		"age_limit":   age_limit,
		"formats":     formats,
	}
}

func (self *RedTubeIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("RedTube", NewRedTubeIE)
}
