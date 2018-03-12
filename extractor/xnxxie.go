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
 * xnxxie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/xnxx.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type XNXXIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewXNXXIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &XNXXIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "XNXX"
	_VALID_URL = "https?://(?:video|www)\\.xnxx\\.com/video-?(?P<id>[0-9a-z]+)/"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *XNXXIE) Key() string {
	return "XNXX"
}

func (self *XNXXIE) Name() string {
	return self.IE_NAME
}

func (self *XNXXIE) _real_extract(url string) rnt.SDict {
	var (
		duration   rnt.OptInt
		format_id  rnt.OptString
		format_url rnt.OptString
		formats    []interface{}
		get        func(string, interface{}, bool) rnt.OptString
		mobj       rnt.Match
		thumbnail  rnt.OptString
		title      rnt.OptString
		video_id   string
		view_count rnt.OptInt
		webpage    string
	)
	video_id = (self).MatchID(url)
	webpage = (self).DownloadWebpageURL(url, video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	get = func(meta string, κdefault interface{}, fatal bool) rnt.OptString {
		return (self).SearchRegexOne(rnt.StrFormat2("set%s\\s*\\(\\s*([\"\\'])(?P<value>(?:(?!\\1).)+)\\1", meta), webpage, meta, κdefault, fatal, 0, "value")
	}
	title = func() rnt.OptString {
		if v := ((self).OgSearchTitle(webpage, nil, true)); τ_isTruthy_Os(v) {
			return v
		} else {
			return get("VideoTitle", rnt.NoDefault, true)
		}
	}()
	formats = []interface{}{}
	τmp1 := rnt.ReFindIter("setVideo(?:Url(?P<id>Low|High)|HLS)\\s*\\(\\s*(?P<q>[\"\\'])(?P<url>(?:https?:)?//.+?)(?P=q)", webpage, 0)
	for τel, τhv := τmp1.Next(); τhv; τel, τhv = τmp1.Next() {
		mobj = τel
		format_url = rnt.ReMatchGroupOne(mobj, "url")
		if (rnt.DetermineExt(format_url, "unknown_video")) == ("m3u8") {
			formats = append(formats, τ_conv_Ld_to_Lα((self).ExtractM3U8Formats(format_url.Get(), video_id, rnt.AsOptString("mp4"), rnt.AsOptString("m3u8_native"), 1, rnt.AsOptString("hls"), rnt.OptString{}, rnt.OptString{}, false, false))...)
		} else {
			format_id = rnt.ReMatchGroupOne(mobj, "id")
			if τ_isTruthy_Os(format_id) {
				format_id = rnt.AsOptString(rnt.StrLower(format_id.Get()))
			}
			formats = append(formats, rnt.SDict{
				"url":       format_url,
				"format_id": format_id,
				"quality": func() int {
					if (format_id.Get()) == ("low") {
						return -(1)
					} else {
						return 0
					}
				}(),
			})
		}
	}
	formats = τ_conv_Ld_to_Lα((self).SortFormats(τ_conv_Lα_to_Ld(formats)))
	thumbnail = func() rnt.OptString {
		if v := ((self).OgSearchThumbnail(webpage, nil)); τ_isTruthy_Os(v) {
			return v
		} else if v := (get("ThumbUrl", rnt.NoDefault, false)); τ_isTruthy_Os(v) {
			return v
		} else {
			return get("ThumbUrl169", rnt.NoDefault, false)
		}
	}()
	duration = rnt.IntOrNone((self).OgSearchPropertyOne("duration", webpage, rnt.OptString{}, rnt.NoDefault, true), 1, rnt.OptInt{}, 1)
	view_count = rnt.StrToInt((self).SearchRegexOne("id=[\"\\']nb-views-number[^>]+>([\\d,.]+)", webpage, "view count", nil, true, 0, nil))
	return rnt.SDict{
		"id":         video_id,
		"title":      title,
		"thumbnail":  thumbnail,
		"duration":   duration,
		"view_count": view_count,
		"age_limit":  18,
		"formats":    formats,
	}
}

func (self *XNXXIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("XNXX", NewXNXXIE)
}
