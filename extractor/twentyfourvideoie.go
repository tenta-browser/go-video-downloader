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
 * twentyfourvideoie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/twentyfourvideo.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type TwentyFourVideoIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewTwentyFourVideoIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &TwentyFourVideoIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "TwentyFourVideo"
	IE_NAME = "24video"
	_VALID_URL = "https?://(?P<host>(?:www\\.)?24video\\.(?:net|me|xxx|sex|tube|adult))/(?:video/(?:view|xml)/|player/new24_play\\.swf\\?id=)(?P<id>\\d+)"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *TwentyFourVideoIE) Key() string {
	return "TwentyFourVideo"
}

func (self *TwentyFourVideoIE) Name() string {
	return self.IE_NAME
}

func (self *TwentyFourVideoIE) _real_extract(url string) rnt.SDict {
	var (
		age_limit     int
		comment_count rnt.OptInt
		description   rnt.OptString
		dislike_count rnt.OptInt
		duration      rnt.OptInt
		formats       []rnt.SDict
		host          rnt.OptString
		like_count    rnt.OptInt
		mobj          rnt.Match
		thumbnail     rnt.OptString
		timestamp     rnt.OptInt
		title         rnt.OptString
		uploader      rnt.OptString
		video         rnt.XMLElement
		video_id      rnt.OptString
		video_xml     rnt.XMLElement
		view_count    rnt.OptInt
		webpage       string
	)
	mobj = rnt.ReMatch((self).VALIDURL, url, 0)
	video_id = rnt.ReMatchGroupOne(mobj, "id")
	host = rnt.ReMatchGroupOne(mobj, "host")
	webpage = (self).DownloadWebpageURL(rnt.StrFormat2("http://%s/video/view/%s", host, video_id), video_id.Get(), rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	title = (self).OgSearchTitle(webpage, rnt.NoDefault, true)
	description = (self).HTMLSearchRegexOne("<(p|span)[^>]+itemprop=\"description\"[^>]*>(?P<description>[^<]+)</\\1>", webpage, "description", rnt.NoDefault, false, 0, "description")
	thumbnail = (self).OgSearchThumbnail(webpage, rnt.NoDefault)
	duration = rnt.IntOrNone((self).OgSearchPropertyOne("duration", webpage, rnt.AsOptString("duration"), rnt.NoDefault, false), 1, rnt.OptInt{}, 1)
	timestamp = rnt.ParseISO8601((self).SearchRegexOne("<time[^>]+\\bdatetime=\"([^\"]+)\"[^>]+itemprop=\"uploadDate\"", webpage, "upload date", rnt.NoDefault, false, 0, nil), "T")
	uploader = (self).HTMLSearchRegexOne("class=\"video-uploaded\"[^>]*>\\s*<a href=\"/jsecUser/movies/[^\"]+\"[^>]*>([^<]+)</a>", webpage, "uploader", rnt.NoDefault, false, 0, nil)
	view_count = rnt.IntOrNone((self).HTMLSearchRegexOne("<span class=\"video-views\">(\\d+) просмотр", webpage, "view count", rnt.NoDefault, false, 0, nil), 1, rnt.OptInt{}, 1)
	comment_count = rnt.IntOrNone((self).HTMLSearchRegexOne("<a[^>]+href=\"#tab-comments\"[^>]*>(\\d+) комментари", webpage, "comment count", nil, true, 0, nil), 1, rnt.OptInt{}, 1)
	(self).DownloadXML(rnt.StrFormat2("http://%s/video/xml/%s?mode=init", host, video_id), video_id.Get(), rnt.AsOptString("Downloading init XML"), rnt.AsOptString("Unable to download XML"), nil, true, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	video_xml = (self).DownloadXML(rnt.StrFormat2("http://%s/video/xml/%s?mode=play", host, video_id), video_id.Get(), rnt.AsOptString("Downloading video XML"), rnt.AsOptString("Unable to download XML"), nil, true, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	video = rnt.XPathElementOne(video_xml, ".//video", rnt.AsOptString("video"), true, rnt.NoDefault)
	formats = []rnt.SDict{{
		"url": rnt.XPathAttr(video, "", "url", rnt.AsOptString("video URL"), true, rnt.NoDefault),
	}}
	like_count = rnt.IntOrNone(rnt.XMLGetAttribute(video, "ratingPlus", rnt.OptString{}), 1, rnt.OptInt{}, 1)
	dislike_count = rnt.IntOrNone(rnt.XMLGetAttribute(video, "ratingMinus", rnt.OptString{}), 1, rnt.OptInt{}, 1)
	age_limit = func() int {
		if (rnt.XMLGetAttribute(video, "adult", rnt.OptString{}).Get()) == ("true") {
			return 18
		} else {
			return 0
		}
	}()
	return rnt.SDict{
		"id":            video_id,
		"title":         title,
		"description":   description,
		"thumbnail":     thumbnail,
		"uploader":      uploader,
		"duration":      duration,
		"timestamp":     timestamp,
		"view_count":    view_count,
		"comment_count": comment_count,
		"like_count":    like_count,
		"dislike_count": dislike_count,
		"age_limit":     age_limit,
		"formats":       formats,
	}
}

func (self *TwentyFourVideoIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("TwentyFourVideo", NewTwentyFourVideoIE)
}
