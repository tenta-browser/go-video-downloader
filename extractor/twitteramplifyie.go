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
 * twitteramplifyie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/twitter.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type TwitterAmplifyIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewTwitterAmplifyIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &TwitterAmplifyIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "TwitterAmplify"
	IE_NAME = "twitter:amplify"
	_VALID_URL = "https?://amp\\.twimg\\.com/v/(?P<id>[0-9a-f\\-]{36})"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *TwitterAmplifyIE) Key() string {
	return "TwitterAmplify"
}

func (self *TwitterAmplifyIE) Name() string {
	return self.IE_NAME
}

func (self *TwitterAmplifyIE) _extract_formats_from_vmap_url(vmap_url rnt.OptString, video_id string) []rnt.SDict {
	var (
		video_url string
		vmap_data rnt.XMLElement
	)
	vmap_data = (self).DownloadXML(vmap_url.Get(), video_id, rnt.AsOptString("Downloading XML"), rnt.AsOptString("Unable to download XML"), nil, true, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{}, nil)
	video_url = rnt.StrStrip(rnt.XPathTextOne(vmap_data, ".//MediaFile", rnt.OptString{}, false, rnt.NoDefault).Get(), "")
	if (rnt.DetermineExt(rnt.AsOptString(video_url), "unknown_video")) == ("m3u8") {
		return (self).ExtractM3U8Formats(video_url, video_id, rnt.AsOptString("mp4"), rnt.AsOptString("m3u8_native"), nil, rnt.AsOptString("hls"), rnt.OptString{}, rnt.OptString{}, true, false)
	}
	return []rnt.SDict{{
		"url": video_url,
	}}
}

func (self *TwitterAmplifyIE) _real_extract(url string) rnt.SDict {
	var (
		_find_dimension func(string) τ_TOiOiω
		formats         []rnt.SDict
		thumbnail       rnt.OptString
		thumbnail_h     rnt.OptInt
		thumbnail_w     rnt.OptInt
		thumbnails      []interface{}
		video_h         rnt.OptInt
		video_id        string
		video_w         rnt.OptInt
		vmap_url        rnt.OptString
		webpage         string
		τmp1            τ_TOiOiω
		τmp2            τ_TOiOiω
	)
	video_id = (self).MatchID(url)
	webpage = (self).DownloadWebpageURL(url, video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{}, nil)
	vmap_url = (self).HTMLSearchMetaOne("twitter:amplify:vmap", webpage, rnt.AsOptString("vmap url"), false, rnt.NoDefault, 0)
	formats = (self)._extract_formats_from_vmap_url(vmap_url, video_id)
	thumbnails = []interface{}{}
	thumbnail = (self).HTMLSearchMetaOne("twitter:image:src", webpage, rnt.AsOptString("thumbnail"), false, rnt.NoDefault, 0)
	_find_dimension = func(target string) τ_TOiOiω {
		var (
			h rnt.OptInt
			w rnt.OptInt
		)
		w = rnt.IntOrNone((self).HTMLSearchMetaOne(rnt.StrFormat2("twitter:%s:width", target), webpage, rnt.OptString{}, false, rnt.NoDefault, 0), 1, rnt.OptInt{}, 1)
		h = rnt.IntOrNone((self).HTMLSearchMetaOne(rnt.StrFormat2("twitter:%s:height", target), webpage, rnt.OptString{}, false, rnt.NoDefault, 0), 1, rnt.OptInt{}, 1)
		return τ_TOiOiω{
			Φ0: w,
			Φ1: h,
		}
	}
	if τ_isTruthy_Os(thumbnail) {
		τmp1 = _find_dimension("image")
		thumbnail_w = (τmp1).Φ0
		thumbnail_h = (τmp1).Φ1
		thumbnails = append(thumbnails, rnt.SDict{
			"url":    thumbnail,
			"width":  thumbnail_w,
			"height": thumbnail_h,
		})
	}
	τmp2 = _find_dimension("player")
	video_w = (τmp2).Φ0
	video_h = (τmp2).Φ1
	rnt.DictUpdate((formats)[0], rnt.SDict{
		"width":  video_w,
		"height": video_h,
	})
	return rnt.SDict{
		"id":         video_id,
		"title":      "Twitter Video",
		"formats":    formats,
		"thumbnails": thumbnails,
	}
}

func (self *TwitterAmplifyIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("TwitterAmplify", NewTwitterAmplifyIE)
}
