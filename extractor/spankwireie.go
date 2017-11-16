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
 * spankwireie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/spankwire.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type SpankwireIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewSpankwireIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &SpankwireIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "Spankwire"
	_VALID_URL = "https?://(?:www\\.)?(?P<url>spankwire\\.com/[^/]*/video(?P<id>[0-9]+)/?)"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *SpankwireIE) Key() string {
	return "Spankwire"
}

func (self *SpankwireIE) Name() string {
	return self.IE_NAME
}

func (self *SpankwireIE) _real_extract(url string) rnt.SDict {
	var (
		age_limit     int
		comment_count rnt.OptInt
		description   rnt.OptString
		formats       []interface{}
		height        int
		heights       []int
		m             rnt.Match
		mobj          rnt.Match
		password      string
		path          string
		req           rnt.Request
		tbr           rnt.OptInt
		thumbnail     rnt.OptString
		title         rnt.OptString
		upload_date   rnt.OptString
		uploader      rnt.OptString
		uploader_id   rnt.OptString
		video_id      rnt.OptString
		video_url     string
		video_urls    []string
		videos        [][]string
		view_count    rnt.OptInt
		webpage       string
		τmp1          τ_Tisω
	)
	mobj = rnt.ReMatch((self).VALIDURL, url, 0)
	video_id = rnt.ReMatchGroupOne(mobj, "id")
	req = rnt.SanitizedRequest(("http://www." + rnt.ReMatchGroupOne(mobj, "url").Get()), nil, rnt.SDict{})
	rnt.RequestAddHeader(req, "Cookie", "age_verified=1")
	webpage = (self).DownloadWebpageRequest(req, video_id.Get(), rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	title = (self).HTMLSearchRegexOne("<h1>([^<]+)", webpage, "title", rnt.NoDefault, true, 0, nil)
	description = (self).HTMLSearchRegexOne("(?s)<div\\s+id=\"descriptionContent\">(.+?)</div>", webpage, "description", rnt.NoDefault, false, 0, nil)
	thumbnail = (self).HTMLSearchRegexOne("playerData\\.screenShot\\s*=\\s*[\"\\']([^\"\\']+)[\"\\']", webpage, "thumbnail", rnt.NoDefault, false, 0, nil)
	uploader = (self).HTMLSearchRegexOne("by:\\s*<a [^>]*>(.+?)</a>", webpage, "uploader", rnt.NoDefault, false, 0, nil)
	uploader_id = (self).HTMLSearchRegexOne("by:\\s*<a href=\"/(?:user/viewProfile|Profile\\.aspx)\\?.*?UserId=(\\d+).*?\"", webpage, "uploader id", rnt.NoDefault, false, 0, nil)
	upload_date = rnt.UnifiedStrDate((self).HTMLSearchRegexOne("</a> on (.+?) at \\d+:\\d+", webpage, "upload date", rnt.NoDefault, false, 0, nil), true)
	view_count = rnt.StrToInt((self).HTMLSearchRegexOne("<div id=\"viewsCounter\"><span>([\\d,\\.]+)</span> views</div>", webpage, "view count", rnt.NoDefault, false, 0, nil))
	comment_count = rnt.StrToInt((self).HTMLSearchRegexOne("<span\\s+id=\"spCommentCount\"[^>]*>([\\d,\\.]+)</span>", webpage, "comment count", rnt.NoDefault, false, 0, nil))
	videos = rnt.ReFindAllMulti("playerData\\.cdnPath([0-9]{3,})\\s*=\\s*(?:encodeURIComponent\\()?[\"\\']([^\"\\']+)[\"\\']", webpage, 0)
	heights = func() []int {
		τresult := []int{}
		for _, τel := range videos {
			var (
				video []string
			)
			video = τel
			τresult = append(τresult, rnt.ConvertToInt((video)[0]))
		}
		return τresult
	}()
	video_urls = τ_map_Fssω_over_Ls(rnt.ParseUnquote, func() []string {
		τresult := []string{}
		for _, τel := range videos {
			var (
				video []string
			)
			video = τel
			τresult = append(τresult, (video)[1])
		}
		return τresult
	}())
	if (rnt.StrFind(webpage, "flashvars\\.encrypted = \"true\"")) != (-(1)) {
		password = rnt.StrReplace((self).SearchRegexOne("flashvars\\.video_title = \"([^\"]+)", webpage, "password", rnt.NoDefault, true, 0, nil).Get(), "+", " ", -(1))
		video_urls = τ_map_Fssω_over_Ls(func(s string) string {
			return rnt.BytesToStr(rnt.AESDecryptText(s, password, 32), "utf-8")
		}, video_urls)
	}
	formats = []interface{}{}
	for _, τel := range τ_zip_i_s(heights, video_urls) {
		τmp1 = τel
		height = (τmp1).Φ0
		video_url = (τmp1).Φ1
		path = (rnt.URLParse(rnt.AsOptString(video_url))).Path
		m = rnt.ReSearch("/(?P<height>\\d+)[pP]_(?P<tbr>\\d+)[kK]", path, 0)
		if m != nil {
			tbr = rnt.AsOptInt(rnt.ConvertToInt(rnt.ReMatchGroupOne(m, "tbr")))
			height = rnt.ConvertToInt(rnt.ReMatchGroupOne(m, "height"))
		} else {
			tbr = rnt.OptInt{}
		}
		formats = append(formats, rnt.SDict{
			"url":       video_url,
			"format_id": rnt.StrFormat2("%dp", height),
			"height":    height,
			"tbr":       tbr,
		})
	}
	formats = τ_conv_Ld_to_Lα((self).SortFormats(τ_conv_Lα_to_Ld(formats)))
	age_limit = (self).RTASearch(webpage)
	return rnt.SDict{
		"id":            video_id,
		"title":         title,
		"description":   description,
		"thumbnail":     thumbnail,
		"uploader":      uploader,
		"uploader_id":   uploader_id,
		"upload_date":   upload_date,
		"view_count":    view_count,
		"comment_count": comment_count,
		"formats":       formats,
		"age_limit":     age_limit,
	}
}

func (self *SpankwireIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("Spankwire", NewSpankwireIE)
}
