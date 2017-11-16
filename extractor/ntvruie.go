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
 * ntvruie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/ntvru.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type NTVRuIE struct {
	*rnt.CommonIE
	IE_NAME           string
	_VIDEO_ID_REGEXES []string
}

func NewNTVRuIE() rnt.InfoExtractor {
	var (
		IE_NAME           string
		_VALID_URL        string
		_VIDEO_ID_REGEXES []string
	)
	self := &NTVRuIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "NTVRu"
	IE_NAME = "ntv.ru"
	_VALID_URL = "https?://(?:www\\.)?ntv\\.ru/(?:[^/]+/)*(?P<id>[^/?#&]+)"
	_VIDEO_ID_REGEXES = []string{"<meta property=\"og:url\" content=\"http://www\\.ntv\\.ru/video/(\\d+)", "<video embed=[^>]+><id>(\\d+)</id>", "<video restriction[^>]+><key>(\\d+)</key>"}
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	self._VIDEO_ID_REGEXES = _VIDEO_ID_REGEXES
	return self
}

func (self *NTVRuIE) Key() string {
	return "NTVRu"
}

func (self *NTVRuIE) Name() string {
	return self.IE_NAME
}

func (self *NTVRuIE) _real_extract(url string) rnt.SDict {
	var (
		description rnt.OptString
		duration    rnt.OptInt
		file_       rnt.XMLElement
		format_id   string
		formats     []interface{}
		player      rnt.XMLElement
		size        rnt.XMLElement
		thumbnail   rnt.OptString
		title       rnt.OptString
		token       string
		video       rnt.XMLElement
		video_id    rnt.OptString
		video_url   rnt.OptString
		view_count  rnt.OptInt
		webpage     string
	)
	video_id = rnt.AsOptString((self).MatchID(url))
	webpage = (self).DownloadWebpageURL(url, video_id.Get(), rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	video_url = (self).OgSearchPropertyMulti(τ_conv_Tssω_to_Ls(τ_Tssω{
		Φ0: "video",
		Φ1: "video:iframe",
	}), webpage, rnt.OptString{}, nil, true)
	if τ_isTruthy_Os(video_url) {
		video_id = (self).SearchRegexOne("https?://(?:www\\.)?ntv\\.ru/video/(?:embed/)?(\\d+)", video_url.Get(), "video id", nil, true, 0, nil)
	}
	if !(τ_isTruthy_Os(video_id)) {
		video_id = (self).HTMLSearchRegexMulti((self)._VIDEO_ID_REGEXES, webpage, "video id", rnt.NoDefault, true, 0, nil)
	}
	player = (self).DownloadXML(rnt.StrFormat2("http://www.ntv.ru/vi%s/", video_id), video_id.Get(), rnt.AsOptString("Downloading video XML"), rnt.AsOptString("Unable to download XML"), nil, true, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	title = rnt.CleanHTML(rnt.XPathTextOne(player, "./data/title", rnt.AsOptString("title"), true, rnt.NoDefault))
	description = rnt.CleanHTML(rnt.XPathTextOne(player, "./data/description", rnt.AsOptString("description"), false, rnt.NoDefault))
	video = rnt.XMLFind(player, "./data/video")
	video_id = rnt.XPathTextOne(video, "./id", rnt.AsOptString("video id"), false, rnt.NoDefault)
	thumbnail = rnt.XPathTextOne(video, "./splash", rnt.AsOptString("thumbnail"), false, rnt.NoDefault)
	duration = rnt.IntOrNone(rnt.XPathTextOne(video, "./totaltime", rnt.AsOptString("duration"), false, rnt.NoDefault), 1, rnt.OptInt{}, 1)
	view_count = rnt.IntOrNone(rnt.XPathTextOne(video, "./views", rnt.AsOptString("view count"), false, rnt.NoDefault), 1, rnt.OptInt{}, 1)
	token = (self).DownloadWebpageURL("http://stat.ntv.ru/services/access/token", video_id.Get(), rnt.AsOptString("Downloading access token"), rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	formats = []interface{}{}
	for _, τel := range []string{"", "hi", "webm"} {
		format_id = τel
		file_ = rnt.XMLFind(video, rnt.StrFormat2("./%sfile", format_id))
		if file_ == nil {
			continue
		}
		size = rnt.XMLFind(video, rnt.StrFormat2("./%ssize", format_id))
		formats = append(formats, rnt.SDict{
			"url": rnt.StrFormat2("http://media2.ntv.ru/vod/%s&tok=%s", (file_).Text, token),
			"filesize": rnt.IntOrNone(func() rnt.OptString {
				if size != nil {
					return rnt.AsOptString((size).Text)
				} else {
					return rnt.OptString{}
				}
			}(), 1, rnt.OptInt{}, 1),
		})
	}
	formats = τ_conv_Ld_to_Lα((self).SortFormats(τ_conv_Lα_to_Ld(formats)))
	return rnt.SDict{
		"id":          video_id,
		"title":       title,
		"description": description,
		"thumbnail":   thumbnail,
		"duration":    duration,
		"view_count":  view_count,
		"formats":     formats,
	}
}

func (self *NTVRuIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("NTVRu", NewNTVRuIE)
}
