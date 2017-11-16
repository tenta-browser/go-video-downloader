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
 * wistiaie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/wistia.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type WistiaIE struct {
	*rnt.CommonIE
	IE_NAME     string
	_API_URL    string
	_IFRAME_URL string
}

func NewWistiaIE() rnt.InfoExtractor {
	var (
		IE_NAME     string
		_API_URL    string
		_IFRAME_URL string
		_VALID_URL  string
	)
	self := &WistiaIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "Wistia"
	_VALID_URL = "(?:wistia:|https?://(?:fast\\.)?wistia\\.net/embed/iframe/)(?P<id>[a-z0-9]+)"
	_API_URL = "http://fast.wistia.com/embed/medias/%s.json"
	_IFRAME_URL = "http://fast.wistia.net/embed/iframe/%s"
	self.IE_NAME = IE_NAME
	self._API_URL = _API_URL
	self._IFRAME_URL = _IFRAME_URL
	self.VALIDURL = _VALID_URL
	return self
}

func (self *WistiaIE) Key() string {
	return "Wistia"
}

func (self *WistiaIE) Name() string {
	return self.IE_NAME
}

func (self *WistiaIE) _real_extract(url string) rnt.SDict {
	var (
		a          interface{}
		aext       interface{}
		astatus    interface{}
		atype      interface{}
		aurl       interface{}
		data       interface{}
		data_json  interface{}
		formats    []interface{}
		is_m3u8    bool
		thumbnails []interface{}
		title      interface{}
		video_id   string
	)
	video_id = (self).MatchID(url)
	data_json = (self).DownloadJSON(rnt.StrFormat2((self)._API_URL, video_id), video_id, rnt.AsOptString("Downloading JSON metadata"), rnt.AsOptString("Unable to download JSON metadata"), nil, true, rnt.OptString{}, nil, rnt.SDict{
		"Referer": func() string {
			if rnt.StrStartswith(url, "http", rnt.OptInt{}, rnt.OptInt{}) {
				return url
			} else {
				return rnt.StrFormat2((self)._IFRAME_URL, video_id)
			}
		}(),
	}, rnt.SDict{})
	if rnt.IsTruthy(rnt.DictGet(τ_cast_α_to_d(data_json), "error", nil)) {
		panic(rnt.PyExtractorError("Error while getting the playlist", true, rnt.OptString{}))
	}
	data = rnt.UnsafeSubscript(data_json, "media")
	title = rnt.UnsafeSubscript(data, "name")
	formats = []interface{}{}
	thumbnails = []interface{}{}
	for _, τel := range τ_cast_α_to_Lα(rnt.UnsafeSubscript(data, "assets")) {
		a = τel
		aurl = rnt.DictGet(τ_cast_α_to_d(a), "url", nil)
		if !(rnt.IsTruthy(aurl)) {
			continue
		}
		astatus = rnt.DictGet(τ_cast_α_to_d(a), "status", nil)
		atype = rnt.DictGet(τ_cast_α_to_d(a), "type", nil)
		if ((!rnt.IsNone(astatus) && (astatus) != (2)) || τ_search_α_in_Ls(atype, τ_conv_Tssω_to_Ls(τ_Tssω{
			Φ0: "preview",
			Φ1: "storyboard",
		}))) {
			continue
		} else {
			if τ_search_α_in_Ls(atype, τ_conv_Tssω_to_Ls(τ_Tssω{
				Φ0: "still",
				Φ1: "still_image",
			})) {
				thumbnails = append(thumbnails, rnt.SDict{
					"url":    aurl,
					"width":  rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(a), "width", nil), 1, rnt.OptInt{}, 1),
					"height": rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(a), "height", nil), 1, rnt.OptInt{}, 1),
				})
			} else {
				aext = rnt.DictGet(τ_cast_α_to_d(a), "ext", nil)
				is_m3u8 = ((rnt.DictGet(τ_cast_α_to_d(a), "container", nil)) == ("m3u8") || (aext) == ("m3u8"))
				formats = append(formats, rnt.SDict{
					"format_id": atype,
					"url":       aurl,
					"tbr":       rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(a), "bitrate", nil), 1, rnt.OptInt{}, 1),
					"vbr":       rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(a), "opt_vbitrate", nil), 1, rnt.OptInt{}, 1),
					"width":     rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(a), "width", nil), 1, rnt.OptInt{}, 1),
					"height":    rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(a), "height", nil), 1, rnt.OptInt{}, 1),
					"filesize":  rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(a), "size", nil), 1, rnt.OptInt{}, 1),
					"vcodec":    rnt.DictGet(τ_cast_α_to_d(a), "codec", nil),
					"container": rnt.DictGet(τ_cast_α_to_d(a), "container", nil),
					"ext": func() interface{} {
						if is_m3u8 {
							return "mp4"
						} else {
							return aext
						}
					}(),
					"protocol": func() rnt.OptString {
						if is_m3u8 {
							return rnt.AsOptString("m3u8")
						} else {
							return rnt.OptString{}
						}
					}(),
					"preference": func() rnt.OptInt {
						if (atype) == ("original") {
							return rnt.AsOptInt(1)
						} else {
							return rnt.OptInt{}
						}
					}(),
				})
			}
		}
	}
	formats = τ_conv_Ld_to_Lα((self).SortFormats(τ_conv_Lα_to_Ld(formats)))
	return rnt.SDict{
		"id":          video_id,
		"title":       title,
		"description": rnt.DictGet(τ_cast_α_to_d(data), "seoDescription", nil),
		"formats":     formats,
		"thumbnails":  thumbnails,
		"duration":    rnt.FloatOrNone(rnt.DictGet(τ_cast_α_to_d(data), "duration", nil), 1, 1, rnt.OptFloat{}),
		"timestamp":   rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(data), "createdAt", nil), 1, rnt.OptInt{}, 1),
	}
}

func (self *WistiaIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("Wistia", NewWistiaIE)
}
