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
 * jojie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/joj.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type JojIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewJojIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &JojIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "Joj"
	_VALID_URL = "(?x)\n                    (?:\n                        joj:|\n                        https?://media\\.joj\\.sk/embed/\n                    )\n                    (?P<id>[^/?#^]+)\n                "
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *JojIE) Key() string {
	return "Joj"
}

func (self *JojIE) Name() string {
	return self.IE_NAME
}

func (self *JojIE) _real_extract(url string) rnt.SDict {
	var (
		bitrates   interface{}
		duration   rnt.OptInt
		file_el    rnt.XMLElement
		format_id  rnt.OptString
		format_url interface{}
		formats    []interface{}
		height     rnt.OptString
		path       rnt.OptString
		playlist   rnt.XMLElement
		thumbnail  rnt.OptString
		title      rnt.OptString
		video_id   string
		webpage    string
	)
	video_id = (self).MatchID(url)
	webpage = (self).DownloadWebpageURL(rnt.StrFormat2("https://media.joj.sk/embed/%s", video_id), video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{}, nil)
	title = func() rnt.OptString {
		if v := ((self).SearchRegexMulti(τ_conv_Tssω_to_Ls(τ_Tssω{
			Φ0: "videoTitle\\s*:\\s*([\"\\'])(?P<title>(?:(?!\\1).)+)\\1",
			Φ1: "<title>(?P<title>[^<]+)",
		}), webpage, "title", nil, true, 0, "title")); τ_isTruthy_Os(v) {
			return v
		} else {
			return (self).OgSearchTitle(webpage, rnt.NoDefault, true)
		}
	}()
	bitrates = (self).ParseJSON((self).SearchRegexOne("(?s)(?:src|bitrates)\\s*=\\s*({.+?});", webpage, "bitrates", "{}", true, 0, nil).Get(), video_id, rnt.JsToJSON, false)
	formats = []interface{}{}
	for _, τel := range τ_cast_α_to_Lα(func() interface{} {
		if v := (rnt.TryGetOne(bitrates, func(x interface{}) interface{} {
			return rnt.UnsafeSubscript(x, "mp4")
		}, rnt.ConvertToList)); rnt.IsTruthy(v) {
			return v
		} else {
			return []interface{}{}
		}
	}()) {
		format_url = τel
		if τ_isinstance_s(format_url) {
			height = (self).SearchRegexOne("(\\d+)[pP]\\.", rnt.CastToString(format_url), "height", nil, true, 0, nil)
			formats = append(formats, rnt.SDict{
				"url": format_url,
				"format_id": func() rnt.OptString {
					if τ_isTruthy_Os(height) {
						return rnt.AsOptString(rnt.StrFormat2("%sp", height))
					} else {
						return rnt.OptString{}
					}
				}(),
				"height": rnt.ConvertToInt(height),
			})
		}
	}
	if !(len(formats) > 0) {
		playlist = (self).DownloadXML(rnt.StrFormat2("https://media.joj.sk/services/Video.php?clip=%s", video_id), video_id, rnt.AsOptString("Downloading XML"), rnt.AsOptString("Unable to download XML"), nil, true, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{}, nil)
		for _, τel := range rnt.XMLFindAll(playlist, "./files/file") {
			file_el = τel
			path = rnt.XMLGetAttribute(file_el, "path", rnt.OptString{})
			if !(τ_isTruthy_Os(path)) {
				continue
			}
			format_id = func() rnt.OptString {
				if v := (rnt.XMLGetAttribute(file_el, "id", rnt.OptString{})); τ_isTruthy_Os(v) {
					return v
				} else {
					return rnt.XMLGetAttribute(file_el, "label", rnt.OptString{})
				}
			}()
			formats = append(formats, rnt.SDict{
				"url":       rnt.StrFormat2("http://n16.joj.sk/storage/%s", rnt.StrReplace(path.Get(), "dat/", "", 1)),
				"format_id": format_id,
				"height": rnt.IntOrNone((self).SearchRegexOne("(\\d+)[pP]", func() rnt.OptString {
					if v := (format_id); τ_isTruthy_Os(v) {
						return v
					} else {
						return path
					}
				}().Get(), "height", nil, true, 0, nil), 1, rnt.OptInt{}, 1),
			})
		}
	}
	formats = τ_conv_Ld_to_Lα((self).SortFormats(τ_conv_Lα_to_Ld(formats)))
	thumbnail = (self).OgSearchThumbnail(webpage, rnt.NoDefault)
	duration = rnt.IntOrNone((self).SearchRegexOne("videoDuration\\s*:\\s*(\\d+)", webpage, "duration", rnt.NoDefault, false, 0, nil), 1, rnt.OptInt{}, 1)
	return rnt.SDict{
		"id":        video_id,
		"title":     title,
		"thumbnail": thumbnail,
		"duration":  duration,
		"formats":   formats,
	}
}

func (self *JojIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("Joj", NewJojIE)
}
