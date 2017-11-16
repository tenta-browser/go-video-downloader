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
 * xvideosie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/xvideos.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type XVideosIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewXVideosIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &XVideosIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "XVideos"
	_VALID_URL = "https?://(?:www\\.)?xvideos\\.com/video(?P<id>[0-9]+)(?:.*)"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *XVideosIE) Key() string {
	return "XVideos"
}

func (self *XVideosIE) Name() string {
	return self.IE_NAME
}

func (self *XVideosIE) _real_extract(url string) rnt.SDict {
	var (
		_               string
		format_id       string
		format_url      string
		formats         []interface{}
		kind            string
		mobj            rnt.Match
		video_duration  interface{}
		video_id        string
		video_thumbnail rnt.OptString
		video_title     rnt.OptString
		video_url       string
		webpage         string
		τmp1            []string
	)
	video_id = (self).MatchID(url)
	webpage = (self).DownloadWebpageURL(url, video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	mobj = rnt.ReSearch("<h1 class=\"inlineError\">(.+?)</h1>", webpage, 0)
	if mobj != nil {
		panic(rnt.PyExtractorError(rnt.StrFormat2("%s said: %s", (self).IE_NAME, rnt.CleanHTML(rnt.ReMatchGroupOne(mobj, 1))), true, rnt.OptString{}))
	}
	video_title = (self).HTMLSearchRegexOne("<title>(.*?)\\s+-\\s+XVID", webpage, "title", rnt.NoDefault, true, 0, nil)
	video_thumbnail = (self).SearchRegexOne("url_bigthumb=(.+?)&amp", webpage, "thumbnail", rnt.NoDefault, false, 0, nil)
	video_duration = func() interface{} {
		if τ_isTruthy_Oi(rnt.IntOrNone((self).OgSearchPropertyOne("duration", webpage, rnt.OptString{}, nil, true), 1, rnt.OptInt{}, 1)) {
			return rnt.IntOrNone((self).OgSearchPropertyOne("duration", webpage, rnt.OptString{}, nil, true), 1, rnt.OptInt{}, 1)
		} else {
			return rnt.ParseDuration((self).SearchRegexOne("<span[^>]+class=[\"\\']duration[\"\\'][^>]*>.*?(\\d[^<]+)", webpage, "duration", rnt.NoDefault, false, 0, nil))
		}
	}()
	formats = []interface{}{}
	video_url = rnt.ParseUnquote((self).SearchRegexOne("flv_url=(.+?)&", webpage, "video URL", "", true, 0, nil).Get())
	if (video_url) != "" {
		formats = append(formats, rnt.SDict{
			"url":       video_url,
			"format_id": "flv",
		})
	}
	for _, τel := range rnt.ReFindAllMulti("setVideo([^(]+)\\(([\"\\'])(http.+?)\\2\\)", webpage, 0) {
		τmp1 = τel
		kind = (τmp1)[0]
		_ = (τmp1)[1]
		format_url = (τmp1)[2]
		format_id = rnt.StrLower(kind)
		if (format_id) == ("hls") {
			formats = append(formats, τ_conv_Ld_to_Lα((self).ExtractM3U8Formats(format_url, video_id, rnt.AsOptString("mp4"), rnt.AsOptString("m3u8_native"), rnt.OptFloat{}, rnt.AsOptString("hls"), rnt.OptString{}, rnt.OptString{}, false, false))...)
		} else {
			if τ_search_s_in_Ls(format_id, τ_conv_Tssω_to_Ls(τ_Tssω{
				Φ0: "urllow",
				Φ1: "urlhigh",
			})) {
				formats = append(formats, rnt.SDict{
					"url":       format_url,
					"format_id": rnt.StrFormat2("%s-%s", rnt.DetermineExt(rnt.AsOptString(format_url), "mp4"), rnt.StrSlice(format_id, rnt.AsOptInt(3), rnt.OptInt{}, rnt.OptInt{})),
					"quality": func() rnt.OptInt {
						if rnt.StrEndswith(format_id, "low", rnt.OptInt{}, rnt.OptInt{}) {
							return rnt.AsOptInt(-(2))
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
		"id":        video_id,
		"formats":   formats,
		"title":     video_title,
		"duration":  video_duration,
		"thumbnail": video_thumbnail,
		"age_limit": 18,
	}
}

func (self *XVideosIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("XVideos", NewXVideosIE)
}
