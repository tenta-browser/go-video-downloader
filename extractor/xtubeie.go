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
 * xtubeie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/xtube.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type XTubeIE struct {
	*rnt.CommonIE
}

func NewXTubeIE() rnt.InfoExtractor {
	ret := &XTubeIE{}
	ret.CommonIE = rnt.NewCommonIE()
	ret.VALIDURL = "(?x)\n                        (?:\n                            xtube:|\n                            https?://(?:www\\.)?xtube\\.com/(?:watch\\.php\\?.*\\bv=|video-watch/(?P<display_id>[^/]+)-)\n                        )\n                        (?P<id>[^/?&#]+)\n                    "
	return ret
}

func (self *XTubeIE) Key() string {
	return "XTube"
}

func (self *XTubeIE) Name() string {
	return "XTube extractor"
}

func (self *XTubeIE) _real_extract(url string) map[string]interface{} {
	mobj := rnt.ReMatch((self).VALIDURL, url, 0)
	video_id := rnt.ReMatchGroupOne(mobj, "id")
	display_id := rnt.ReMatchGroupOne(mobj, "display_id")
	if !((display_id).IsSet() && (display_id.Get()) != "") {
		display_id = video_id
	}
	url_pattern := ""
	if rnt.StrIsDigit(video_id.Get()) && (len(video_id.Get())) < (11) {
		url_pattern = "http://www.xtube.com/video-watch/-%s"
	} else {
		url_pattern = "http://www.xtube.com/watch.php?v=%s"
	}
	webpage := (self).DownloadWebpageURL(rnt.StrFormat(url_pattern, video_id), display_id.Get(), rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, rnt.OptString{}, map[string]interface{}{"Cookie": "age_verified=1; cookiesAccepted=1"}, map[string]interface{}{})
	sources := (self).ParseJSON((self).SearchRegexOne("([\"\\'])?sources\\1?\\s*:\\s*(?P<sources>{.+?}),", webpage, "sources", rnt.NoDefault, true, 0, "sources").Get(), video_id.Get(), rnt.JsToJSON, true)
	formats := []interface{}{}
	for τmp1, τmp2 := range sources {
		τmp3 := struct {
			_0 string
			_1 interface{}
		}{_0: τmp1,
			_1: τmp2}
		format_id := (τmp3)._0
		format_url := (τmp3)._1
		formats = append(formats, map[string]interface{}{"url": format_url,
			"format_id": format_id,
			"height":    rnt.IntOrNone(format_id, 1, rnt.OptInt{}, 1)})
	}
	formats = (self).RemoveDuplicateFormats(formats)
	(self).SortFormats(formats)
	title := (self).SearchRegexMulti(func() []string {
		τmp1 := struct {
			_0 string
			_1 string
		}{_0: "<h1>\\s*(?P<title>[^<]+?)\\s*</h1>",
			_1: "videoTitle\\s*:\\s*([\"\\'])(?P<title>.+?)\\1"}
		return []string{τmp1._0, τmp1._1}
	}(), webpage, "title", rnt.NoDefault, true, 0, "title")
	description := (self).SearchRegexOne("</h1>\\s*<p>([^<]+)", webpage, "description", rnt.NoDefault, false, 0, nil)
	uploader := (self).SearchRegexMulti(func() []string {
		τmp1 := struct {
			_0 string
			_1 string
		}{_0: "<input[^>]+name=\"contentOwnerId\"[^>]+value=\"([^\"]+)\"",
			_1: "<span[^>]+class=\"nickname\"[^>]*>([^<]+)"}
		return []string{τmp1._0, τmp1._1}
	}(), webpage, "uploader", rnt.NoDefault, false, 0, nil)
	duration := rnt.ParseDuration((self).SearchRegexOne("<dt>Runtime:?</dt>\\s*<dd>([^<]+)</dd>", webpage, "duration", rnt.NoDefault, false, 0, nil))
	view_count := rnt.StrToInt((self).SearchRegexOne("<dt>Views:?</dt>\\s*<dd>([\\d,\\.]+)</dd>", webpage, "view count", rnt.NoDefault, false, 0, nil))
	comment_count := rnt.StrToInt((self).HTMLSearchRegexOne(">Comments? \\(([\\d,\\.]+)\\)<", webpage, "comment count", rnt.NoDefault, false, 0, nil))
	return map[string]interface{}{"id": video_id,
		"display_id":    display_id,
		"title":         title,
		"description":   description,
		"uploader":      uploader,
		"duration":      duration,
		"view_count":    view_count,
		"comment_count": comment_count,
		"age_limit":     18,
		"formats":       formats}
}

func (self *XTubeIE) Extract(url string) (*rnt.VideoResult, error) {
	return rnt.RunExtractor(url, self._real_extract)
}

func init() {
	registerFactory("XTube", NewXTubeIE)
}
