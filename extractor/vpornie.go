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
 * vpornie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/vporn.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type VpornIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewVpornIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &VpornIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "Vporn"
	_VALID_URL = "https?://(?:www\\.)?vporn\\.com/[^/]+/(?P<display_id>[^/]+)/(?P<id>\\d+)"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *VpornIE) Key() string {
	return "Vporn"
}

func (self *VpornIE) Name() string {
	return self.IE_NAME
}

func (self *VpornIE) _real_extract(url string) rnt.SDict {
	var (
		categories    []string
		comment_count rnt.OptInt
		description   rnt.OptString
		display_id    rnt.OptString
		duration      rnt.OptString
		errmsg        string
		fmt           rnt.SDict
		formats       []interface{}
		m             rnt.Match
		mobj          rnt.Match
		thumbnail     rnt.OptString
		title         string
		uploader      rnt.OptString
		video         []string
		video_id      rnt.OptString
		video_url     string
		view_count    rnt.OptInt
		webpage       string
	)
	mobj = rnt.ReMatch((self).VALIDURL, url, 0)
	video_id = rnt.ReMatchGroupOne(mobj, "id")
	display_id = rnt.ReMatchGroupOne(mobj, "display_id")
	webpage = (self).DownloadWebpageURL(url, display_id.Get(), rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	errmsg = "This video has been deleted due to Copyright Infringement or by the account owner!"
	if rnt.StrContains(webpage, errmsg) {
		panic(rnt.PyExtractorError(rnt.StrFormat2("%s said: %s", (self).IE_NAME, errmsg), true, rnt.OptString{}))
	}
	title = rnt.StrStrip((self).HTMLSearchRegexOne("videoname\\s*=\\s*\\'([^\\']+)\\'", webpage, "title", rnt.NoDefault, true, 0, nil).Get(), "")
	description = (self).HTMLSearchRegexOne("class=\"(?:descr|description_txt)\">(.*?)</div>", webpage, "description", rnt.NoDefault, false, 0, nil)
	thumbnail = rnt.UtilURLJoin("http://www.vporn.com", (self).HTMLSearchRegexOne("flashvars\\.imageUrl\\s*=\\s*\"([^\"]+)\"", webpage, "description", nil, true, 0, nil))
	uploader = (self).HTMLSearchRegexOne("(?s)Uploaded by:.*?<a href=\"/user/[^\"]+\"[^>]*>(.+?)</a>", webpage, "uploader", rnt.NoDefault, false, 0, nil)
	categories = rnt.ReFindAllOne("<a href=\"/cat/[^\"]+\"[^>]*>([^<]+)</a>", webpage, 0)
	duration = rnt.ParseDuration((self).SearchRegexOne("Runtime:\\s*</span>\\s*(\\d+ min \\d+ sec)", webpage, "duration", rnt.NoDefault, false, 0, nil))
	view_count = rnt.StrToInt((self).SearchRegexOne("class=\"views\">([\\d,\\.]+) [Vv]iews<", webpage, "view count", rnt.NoDefault, false, 0, nil))
	comment_count = rnt.StrToInt((self).HTMLSearchRegexOne("'Comments \\(([\\d,\\.]+)\\)'", webpage, "comment count", nil, true, 0, nil))
	formats = []interface{}{}
	for _, τel := range rnt.ReFindAllMulti("flashvars\\.videoUrl([^=]+?)\\s*=\\s*\"(https?://[^\"]+)\"", webpage, 0) {
		video = τel
		video_url = (video)[1]
		fmt = rnt.SDict{
			"url":       video_url,
			"format_id": (video)[0],
		}
		m = rnt.ReSearch("_(?P<width>\\d+)x(?P<height>\\d+)_(?P<vbr>\\d+)k\\.mp4$", video_url, 0)
		if m != nil {
			rnt.DictUpdate(fmt, rnt.SDict{
				"width":  rnt.ConvertToInt(rnt.ReMatchGroupOne(m, "width")),
				"height": rnt.ConvertToInt(rnt.ReMatchGroupOne(m, "height")),
				"vbr":    rnt.ConvertToInt(rnt.ReMatchGroupOne(m, "vbr")),
			})
		}
		formats = append(formats, fmt)
	}
	formats = τ_conv_Ld_to_Lα((self).SortFormats(τ_conv_Lα_to_Ld(formats)))
	return rnt.SDict{
		"id":            video_id,
		"display_id":    display_id,
		"title":         title,
		"description":   description,
		"thumbnail":     thumbnail,
		"uploader":      uploader,
		"categories":    categories,
		"duration":      duration,
		"view_count":    view_count,
		"comment_count": comment_count,
		"age_limit":     18,
		"formats":       formats,
	}
}

func (self *VpornIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("Vporn", NewVpornIE)
}
