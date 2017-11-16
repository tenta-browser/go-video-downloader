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
 * videoszie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/videosz.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type VideosZIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewVideosZIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &VideosZIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "VideosZ"
	_VALID_URL = "https?://(?:www\\.|m\\.|)videosz\\.com/[a-z]+/[a-z]+/(?P<id>[0-9]+)_"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *VideosZIE) Key() string {
	return "VideosZ"
}

func (self *VideosZIE) Name() string {
	return self.IE_NAME
}

func (self *VideosZIE) _real_extract(url string) rnt.SDict {
	var (
		format         string
		format_matches [][]string
		format_url     string
		formats        []interface{}
		formats_source rnt.OptString
		title          rnt.OptString
		video_id       string
		webpage        string
		τmp1           []string
	)
	url = rnt.ReSub("^(https?://(?:.+?\\.)?)m\\.", "\\1", url, 0, 0)
	video_id = (self).MatchID(url)
	webpage = (self).DownloadWebpageURL(url, video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	if rnt.ReSearch("<form[^>]*id=\"login_form\"[^>]*>", webpage, 0) != nil {
		panic(rnt.PyExtractorError("LOGIN_REQUIRED", true, rnt.OptString{}))
	}
	title = (self).HTMLSearchMetaOne("title", webpage, rnt.AsOptString("title"), false, rnt.NoDefault, 0)
	title = rnt.AsOptString(rnt.ReSub("([^:]+):.+", "\\1", title.Get(), 0, 0))
	formats_source = (self).SearchRegexOne("<div[^>]*id=\"download_options_btn_content\"[^>]*>(.+?)</div>", webpage, "formats_source", rnt.NoDefault, true, rnt.ReFlagDotAll, nil)
	formats = []interface{}{}
	format_matches = rnt.ReFindAllMulti("<a[^>]*href=\"(?P<url>[^\"]+)\"[^>]*>(?P<format>[^(]+) \\([^)]+\\)</a>", formats_source.Get(), 0)
	for _, τel := range format_matches {
		τmp1 = τel
		format_url = (τmp1)[0]
		format = (τmp1)[1]
		formats = append(formats, rnt.SDict{
			"format_id": format,
			"url":       rnt.URLJoin(url, rnt.AsOptString(format_url)),
		})
	}
	formats = τ_conv_Ld_to_Lα((self).SortFormats(τ_conv_Lα_to_Ld(formats)))
	return rnt.SDict{
		"id":          video_id,
		"title":       title,
		"description": title,
		"age_limit":   (self).MediaRatingSearch(webpage),
		"formats":     formats,
	}
}

func (self *VideosZIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("VideosZ", NewVideosZIE)
}
