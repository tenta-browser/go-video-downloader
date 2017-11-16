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
 * tumblrie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/tumblr.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type TumblrIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewTumblrIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &TumblrIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "Tumblr"
	_VALID_URL = "https?://(?P<blog_name>[^/?#&]+)\\.tumblr\\.com/(?:post|video)/(?P<id>[0-9]+)(?:$|[/?#])"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *TumblrIE) Key() string {
	return "Tumblr"
}

func (self *TumblrIE) Name() string {
	return self.IE_NAME
}

func (self *TumblrIE) _real_extract(url string) rnt.SDict {
	var (
		blog        rnt.OptString
		duration    rnt.OptInt
		formats     []rnt.SDict
		hd_url      interface{}
		iframe      string
		iframe_url  rnt.OptString
		m_url       rnt.Match
		options     interface{}
		sd_url      rnt.OptString
		sources     []interface{}
		urlh        rnt.Response
		video_id    rnt.OptString
		video_title rnt.OptString
		webpage     string
		τmp1        τ_TsζResponseωω
	)
	m_url = rnt.ReMatch((self).VALIDURL, url, 0)
	video_id = rnt.ReMatchGroupOne(m_url, "id")
	blog = rnt.ReMatchGroupOne(m_url, "blog_name")
	url = rnt.StrFormat2("http://%s.tumblr.com/post/%s/", blog, video_id)
	τmp1 = (self).DownloadWebpageHandleURL(url, video_id.Get(), rnt.OptString{}, rnt.OptString{}, true, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	webpage = (τmp1).Φ0
	urlh = (τmp1).Φ1
	iframe_url = (self).SearchRegexOne("src=\\'(https?://www\\.tumblr\\.com/video/[^\\']+)\\'", webpage, "iframe url", nil, true, 0, nil)
	if !iframe_url.IsSet() {
		return (self).URLResult(rnt.ResponseGetURL(urlh), rnt.AsOptString("Generic"), rnt.OptString{}, rnt.OptString{})
	}
	iframe = (self).DownloadWebpageURL(iframe_url.Get(), video_id.Get(), rnt.AsOptString("Downloading iframe page"), rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	duration = rnt.OptInt{}
	sources = []interface{}{}
	sd_url = (self).SearchRegexOne("<source[^>]+src=([\"\\'])(?P<url>.+?)\\1", iframe, "sd video url", nil, true, 0, "url")
	if τ_isTruthy_Os(sd_url) {
		sources = append(sources, τ_TOssω{
			Φ0: sd_url,
			Φ1: "sd",
		})
	}
	options = (self).ParseJSON((self).SearchRegexOne("data-crt-options=([\"\\'])(?P<options>.+?)\\1", iframe, "hd video url", "", true, 0, "options").Get(), video_id.Get(), nil, false)
	if rnt.IsTruthy(options) {
		duration = rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(options), "duration", nil), 1, rnt.OptInt{}, 1)
		hd_url = rnt.DictGet(τ_cast_α_to_d(options), "hdUrl", nil)
		if rnt.IsTruthy(hd_url) {
			sources = append(sources, τ_Tαsω{
				Φ0: hd_url,
				Φ1: "hd",
			})
		}
	}
	formats = func() []rnt.SDict {
		τresult := []rnt.SDict{}
		for _, τel := range τ_enumerate_Lα(sources, 0) {
			var (
				format_id interface{}
				quality   int
				video_url interface{}
				τmp2      τ_Tiαω
				τmp3      interface{}
			)
			τmp2 = τel
			quality = (τmp2).Φ0
			τmp3 = (τmp2).Φ1
			video_url = rnt.UnsafeSubscript(τmp3, 0)
			format_id = rnt.UnsafeSubscript(τmp3, 1)
			τresult = append(τresult, rnt.SDict{
				"url":       video_url,
				"ext":       "mp4",
				"format_id": format_id,
				"height":    rnt.IntOrNone((self).SearchRegexOne("/(\\d{3,4})$", rnt.CastToString(video_url), "height", nil, true, 0, nil), 1, rnt.OptInt{}, 1),
				"quality":   quality,
			})
		}
		return τresult
	}()
	formats = (self).SortFormats(formats)
	video_title = (self).HTMLSearchRegexOne("(?s)<title>(?P<title>.*?)(?: \\| Tumblr)?</title>", webpage, "title", rnt.NoDefault, true, 0, nil)
	return rnt.SDict{
		"id":          video_id,
		"title":       video_title,
		"description": (self).OgSearchDescription(webpage, nil),
		"thumbnail":   (self).OgSearchThumbnail(webpage, nil),
		"duration":    duration,
		"formats":     formats,
	}
}

func (self *TumblrIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("Tumblr", NewTumblrIE)
}
