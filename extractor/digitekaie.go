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
 * digitekaie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/digiteka.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type DigitekaIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewDigitekaIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &DigitekaIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "Digiteka"
	_VALID_URL = "(?x)\n        https?://(?:www\\.)?(?:digiteka\\.net|ultimedia\\.com)/\n        (?:\n            deliver/\n            (?P<embed_type>\n                generic|\n                musique\n            )\n            (?:/[^/]+)*/\n            (?:\n                src|\n                article\n            )|\n            default/index/video\n            (?P<site_type>\n                generic|\n                music\n            )\n            /id\n        )/(?P<id>[\\d+a-z]+)"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *DigitekaIE) Key() string {
	return "Digiteka"
}

func (self *DigitekaIE) Name() string {
	return self.IE_NAME
}

func (self *DigitekaIE) _real_extract(url string) rnt.SDict {
	var (
		deliver_info interface{}
		duration     rnt.OptInt
		formats      []interface{}
		jwconf       interface{}
		mobj         rnt.Match
		source       interface{}
		thumbnail    interface{}
		timestamp    rnt.OptInt
		title        interface{}
		uploader_id  interface{}
		video_id     rnt.OptString
		video_type   rnt.OptString
		yt_id        interface{}
	)
	mobj = rnt.ReMatch((self).VALIDURL, url, 0)
	video_id = rnt.ReMatchGroupOne(mobj, "id")
	video_type = func() rnt.OptString {
		if τ_isTruthy_Os(rnt.ReMatchGroupOne(mobj, "embed_type")) {
			return rnt.ReMatchGroupOne(mobj, "embed_type")
		} else {
			return rnt.ReMatchGroupOne(mobj, "site_type")
		}
	}()
	if (video_type.Get()) == ("music") {
		video_type = rnt.AsOptString("musique")
	}
	deliver_info = (self).DownloadJSON(rnt.StrFormat2("http://www.ultimedia.com/deliver/video?video=%s&topic=%s", video_id, video_type), video_id.Get(), rnt.AsOptString("Downloading JSON metadata"), rnt.AsOptString("Unable to download JSON metadata"), nil, true, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	yt_id = rnt.DictGet(τ_cast_α_to_d(deliver_info), "yt_id", nil)
	if rnt.IsTruthy(yt_id) {
		return (self).URLResult(rnt.CastToString(yt_id), rnt.AsOptString("Youtube"), rnt.OptString{}, rnt.OptString{})
	}
	jwconf = rnt.UnsafeSubscript(deliver_info, "jwconf")
	formats = []interface{}{}
	for _, τel := range τ_cast_α_to_Lα(rnt.UnsafeSubscript(rnt.UnsafeSubscript(rnt.UnsafeSubscript(jwconf, "playlist"), 0), "sources")) {
		source = τel
		formats = append(formats, rnt.SDict{
			"url":       rnt.UnsafeSubscript(source, "file"),
			"format_id": rnt.DictGet(τ_cast_α_to_d(source), "label", nil),
		})
	}
	formats = τ_conv_Ld_to_Lα((self).SortFormats(τ_conv_Lα_to_Ld(formats)))
	title = rnt.UnsafeSubscript(deliver_info, "title")
	thumbnail = rnt.DictGet(τ_cast_α_to_d(jwconf), "image", nil)
	duration = rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(deliver_info), "duration", nil), 1, rnt.OptInt{}, 1)
	timestamp = rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(deliver_info), "release_time", nil), 1, rnt.OptInt{}, 1)
	uploader_id = rnt.DictGet(τ_cast_α_to_d(deliver_info), "owner_id", nil)
	return rnt.SDict{
		"id":          video_id,
		"title":       title,
		"thumbnail":   thumbnail,
		"duration":    duration,
		"timestamp":   timestamp,
		"uploader_id": uploader_id,
		"formats":     formats,
	}
}

func (self *DigitekaIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("Digiteka", NewDigitekaIE)
}
