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
 * glideie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/glide.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type GlideIE struct {
	*rnt.CommonIE
	IE_DESC string
	IE_NAME string
}

func NewGlideIE() rnt.InfoExtractor {
	var (
		IE_DESC    string
		IE_NAME    string
		_VALID_URL string
	)
	self := &GlideIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "Glide"
	IE_DESC = "Glide mobile video messages (glide.me)"
	_VALID_URL = "https?://share\\.glide\\.me/(?P<id>[A-Za-z0-9\\-=_+]+)"
	self.IE_DESC = IE_DESC
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *GlideIE) Key() string {
	return "Glide"
}

func (self *GlideIE) Name() string {
	return self.IE_NAME
}

func (self *GlideIE) _real_extract(url string) rnt.SDict {
	var (
		thumbnail rnt.OptString
		title     rnt.OptString
		video_id  string
		video_url rnt.OptString
		webpage   string
	)
	video_id = (self).MatchID(url)
	webpage = (self).DownloadWebpageURL(url, video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	title = func() rnt.OptString {
		if τ_isTruthy_Os((self).HTMLSearchRegexOne("<title>(.+?)</title>", webpage, "title", nil, true, 0, nil)) {
			return (self).HTMLSearchRegexOne("<title>(.+?)</title>", webpage, "title", nil, true, 0, nil)
		} else {
			return (self).OgSearchTitle(webpage, rnt.NoDefault, true)
		}
	}()
	video_url = func() rnt.OptString {
		if τ_isTruthy_Os((self).ProtoRelativeURL((self).SearchRegexOne("<source[^>]+src=([\"\\'])(?P<url>.+?)\\1", webpage, "video URL", nil, true, 0, "url"), rnt.OptString{})) {
			return (self).ProtoRelativeURL((self).SearchRegexOne("<source[^>]+src=([\"\\'])(?P<url>.+?)\\1", webpage, "video URL", nil, true, 0, "url"), rnt.OptString{})
		} else {
			return (self).OgSearchVideoURL(webpage, "video url", true, rnt.NoDefault)
		}
	}()
	thumbnail = func() rnt.OptString {
		if τ_isTruthy_Os((self).ProtoRelativeURL((self).SearchRegexOne("<img[^>]+id=[\"\\']video-thumbnail[\"\\'][^>]+src=([\"\\'])(?P<url>.+?)\\1", webpage, "thumbnail url", nil, true, 0, "url"), rnt.OptString{})) {
			return (self).ProtoRelativeURL((self).SearchRegexOne("<img[^>]+id=[\"\\']video-thumbnail[\"\\'][^>]+src=([\"\\'])(?P<url>.+?)\\1", webpage, "thumbnail url", nil, true, 0, "url"), rnt.OptString{})
		} else {
			return (self).OgSearchThumbnail(webpage, rnt.NoDefault)
		}
	}()
	return rnt.SDict{
		"id":        video_id,
		"title":     title,
		"url":       video_url,
		"thumbnail": thumbnail,
	}
}

func (self *GlideIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("Glide", NewGlideIE)
}
