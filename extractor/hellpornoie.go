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
 * hellpornoie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/hellporno.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type HellPornoIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewHellPornoIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &HellPornoIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "HellPorno"
	_VALID_URL = "https?://(?:www\\.)?hellporno\\.(?:com/videos|net/v)/(?P<id>[^/]+)"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *HellPornoIE) Key() string {
	return "HellPorno"
}

func (self *HellPornoIE) Name() string {
	return self.IE_NAME
}

func (self *HellPornoIE) _real_extract(url string) rnt.SDict {
	var (
		categories    []string
		display_id    string
		ext           string
		flashvars     interface{}
		fmt           rnt.SDict
		formats       []interface{}
		m             rnt.Match
		thumbnail     interface{}
		title         rnt.OptString
		video_id      interface{}
		video_text    interface{}
		video_url     interface{}
		video_url_key string
		webpage       string
	)
	display_id = (self).MatchID(url)
	webpage = (self).DownloadWebpageURL(url, display_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	title = rnt.RemoveEnd((self).HTMLSearchRegexOne("<title>([^<]+)</title>", webpage, "title", rnt.NoDefault, true, 0, nil), " - Hell Porno")
	flashvars = (self).ParseJSON((self).SearchRegexOne("var\\s+flashvars\\s*=\\s*({.+?});", webpage, "flashvars", rnt.NoDefault, true, 0, nil).Get(), display_id, rnt.JsToJSON, true)
	video_id = rnt.DictGet(τ_cast_α_to_d(flashvars), "video_id", nil)
	thumbnail = rnt.DictGet(τ_cast_α_to_d(flashvars), "preview_url", nil)
	ext = rnt.DetermineExt(rnt.CastToOptString(rnt.DictGet(τ_cast_α_to_d(flashvars), "postfix", nil)), "mp4")
	formats = []interface{}{}
	for _, τel := range []string{"video_url", "video_alt_url"} {
		video_url_key = τel
		video_url = rnt.DictGet(τ_cast_α_to_d(flashvars), video_url_key, nil)
		if !(rnt.IsTruthy(video_url)) {
			continue
		}
		video_text = rnt.DictGet(τ_cast_α_to_d(flashvars), rnt.StrFormat2("%s_text", video_url_key), nil)
		fmt = rnt.SDict{
			"url":       video_url,
			"ext":       ext,
			"format_id": video_text,
		}
		m = rnt.ReSearch("^(?P<height>\\d+)[pP]", rnt.CastToString(video_text), 0)
		if m != nil {
			(fmt)["height"] = rnt.ConvertToInt(rnt.ReMatchGroupOne(m, "height"))
		}
		formats = append(formats, fmt)
	}
	formats = τ_conv_Ld_to_Lα((self).SortFormats(τ_conv_Lα_to_Ld(formats)))
	categories = rnt.StrSplit((self).HTMLSearchMetaOne("keywords", webpage, rnt.AsOptString("categories"), false, "", 0).Get(), ",", -(1))
	return rnt.SDict{
		"id":         video_id,
		"display_id": display_id,
		"title":      title,
		"thumbnail":  thumbnail,
		"categories": categories,
		"age_limit":  18,
		"formats":    formats,
	}
}

func (self *HellPornoIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("HellPorno", NewHellPornoIE)
}
