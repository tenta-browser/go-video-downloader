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
 * youpornie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/youporn.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type YouPornIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewYouPornIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &YouPornIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "YouPorn"
	_VALID_URL = "https?://(?:www\\.)?youporn\\.com/watch/(?P<id>\\d+)/(?P<display_id>[^/?#&]+)"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *YouPornIE) Key() string {
	return "YouPorn"
}

func (self *YouPornIE) Name() string {
	return self.IE_NAME
}

func (self *YouPornIE) _real_extract(url string) rnt.SDict {
	var (
		_               string
		age_limit       int
		average_rating  rnt.OptInt
		bitrate         int
		categories      []interface{}
		comment_count   rnt.OptInt
		definition      interface{}
		definitions     interface{}
		description     rnt.OptString
		display_id      rnt.OptString
		encrypted_link  string
		extract_tag_box func(string, string) []interface{}
		f               rnt.SDict
		formats         []interface{}
		height          int
		link            string
		links           []interface{}
		mobj            rnt.Match
		request         rnt.Request
		sources         rnt.OptString
		tags            []interface{}
		thumbnail       rnt.OptString
		title           rnt.OptString
		upload_date     rnt.OptString
		uploader        rnt.OptString
		video_id        rnt.OptString
		video_url       interface{}
		view_count      rnt.OptInt
		webpage         string
		τmp1            []string
		τmp2            []string
		τmp3            []string
		τmp4            []string
	)
	mobj = rnt.ReMatch((self).VALIDURL, url, 0)
	video_id = rnt.ReMatchGroupOne(mobj, "id")
	display_id = rnt.ReMatchGroupOne(mobj, "display_id")
	request = rnt.SanitizedRequest(url, nil, rnt.SDict{})
	rnt.RequestAddHeader(request, "Cookie", "age_verified=1")
	webpage = (self).DownloadWebpageRequest(request, display_id.Get(), rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	title = func() rnt.OptString {
		if τ_isTruthy_Os((self).SearchRegexMulti([]string{"(?:video_titles|videoTitle)\\s*[:=]\\s*([\"\\'])(?P<title>(?:(?!\\1).)+)\\1", "<h1[^>]+class=[\"\\']heading\\d?[\"\\'][^>]*>(?P<title>[^<]+)<"}, webpage, "title", nil, true, 0, "title")) {
			return (self).SearchRegexMulti([]string{"(?:video_titles|videoTitle)\\s*[:=]\\s*([\"\\'])(?P<title>(?:(?!\\1).)+)\\1", "<h1[^>]+class=[\"\\']heading\\d?[\"\\'][^>]*>(?P<title>[^<]+)<"}, webpage, "title", nil, true, 0, "title")
		} else if τ_isTruthy_Os((self).OgSearchTitle(webpage, nil, true)) {
			return (self).OgSearchTitle(webpage, nil, true)
		} else {
			return (self).HTMLSearchMetaOne("title", webpage, rnt.OptString{}, true, rnt.NoDefault, 0)
		}
	}()
	links = []interface{}{}
	definitions = (self).ParseJSON((self).SearchRegexOne("mediaDefinition\\s*=\\s*(\\[.+?\\]);", webpage, "media definitions", "[]", true, 0, nil).Get(), video_id.Get(), nil, false)
	if rnt.IsTruthy(definitions) {
		for _, τel := range τ_cast_α_to_Lα(definitions) {
			definition = τel
			if !(τ_isinstance_d(definition)) {
				continue
			}
			video_url = rnt.DictGet(τ_cast_α_to_d(definition), "videoUrl", nil)
			if rnt.IsTruthy(func() interface{} {
				if !(τ_isinstance_s(video_url)) {
					return τ_isinstance_s(video_url)
				} else {
					return video_url
				}
			}()) {
				links = append(links, video_url)
			}
		}
	}
	for _, τel := range rnt.ReFindAllMulti("<a[^>]+href=([\"\\'])(http.+?)\\1[^>]+title=[\"\\']Download [Vv]ideo", webpage, 0) {
		τmp1 = τel
		_ = (τmp1)[0]
		link = (τmp1)[1]
		links = append(links, link)
	}
	sources = (self).SearchRegexOne("(?s)sources\\s*:\\s*({.+?})", webpage, "sources", nil, true, 0, nil)
	if τ_isTruthy_Os(sources) {
		for _, τel := range rnt.ReFindAllMulti("[^:]+\\s*:\\s*([\"\\'])(http.+?)\\1", sources.Get(), 0) {
			τmp2 = τel
			_ = (τmp2)[0]
			link = (τmp2)[1]
			links = append(links, link)
		}
	}
	for _, τel := range rnt.ReFindAllMulti("(?:videoSrc|videoIpadUrl|html5PlayerSrc)\\s*[:=]\\s*([\"\\'])(http.+?)\\1", webpage, 0) {
		τmp3 = τel
		_ = (τmp3)[0]
		link = (τmp3)[1]
		links = append(links, link)
	}
	for _, τel := range rnt.ReFindAllMulti("encryptedQuality\\d{3,4}URL\\s*=\\s*([\"\\'])([\\da-zA-Z+/=]+)\\1", webpage, 0) {
		τmp4 = τel
		_ = (τmp4)[0]
		encrypted_link = (τmp4)[1]
		links = append(links, rnt.BytesToStr(rnt.AESDecryptText(encrypted_link, title.Get(), 32), "utf-8"))
	}
	formats = []interface{}{}
	for _, τel := range τ_conv_SOs_to_LOs(τ_constr_SOs_from_LOs(func() []rnt.OptString {
		τresult := []rnt.OptString{}
		for _, τel := range links {
			var (
				link interface{}
			)
			link = τel
			τresult = append(τresult, rnt.UnescapeHTML(rnt.CastToOptString(link)))
		}
		return τresult
	}())) {
		video_url = τel
		f = rnt.SDict{
			"url": video_url,
		}
		mobj = rnt.ReSearch("(?P<height>\\d{3,4})[pP]_(?P<bitrate>\\d+)[kK]_\\d+/", rnt.CastToString(video_url), 0)
		if mobj != nil {
			height = rnt.ConvertToInt(rnt.ReMatchGroupOne(mobj, "height"))
			bitrate = rnt.ConvertToInt(rnt.ReMatchGroupOne(mobj, "bitrate"))
			rnt.DictUpdate(f, rnt.SDict{
				"format_id": rnt.StrFormat2("%dp-%dk", height, bitrate),
				"height":    height,
				"tbr":       bitrate,
			})
		}
		formats = append(formats, f)
	}
	formats = τ_conv_Ld_to_Lα((self).SortFormats(τ_conv_Lα_to_Ld(formats)))
	description = (self).OgSearchDescription(webpage, nil)
	thumbnail = (self).SearchRegexOne("(?:imageurl\\s*=|poster\\s*:)\\s*([\"\\'])(?P<thumbnail>.+?)\\1", webpage, "thumbnail", rnt.NoDefault, false, 0, "thumbnail")
	uploader = (self).HTMLSearchRegexOne("(?s)<div[^>]+class=[\"\\']submitByLink[\"\\'][^>]*>(.+?)</div>", webpage, "uploader", rnt.NoDefault, false, 0, nil)
	upload_date = rnt.UnifiedStrDate((self).HTMLSearchRegexMulti([]string{"Date\\s+[Aa]dded:\\s*<span>([^<]+)", "(?s)<div[^>]+class=[\"\\']videoInfo(?:Date|Time)[\"\\'][^>]*>(.+?)</div>"}, webpage, "upload date", rnt.NoDefault, false, 0, nil), true)
	age_limit = (self).RTASearch(webpage)
	average_rating = rnt.IntOrNone((self).SearchRegexOne("<div[^>]+class=[\"\\']videoRatingPercentage[\"\\'][^>]*>(\\d+)%</div>", webpage, "average rating", rnt.NoDefault, false, 0, nil), 1, rnt.OptInt{}, 1)
	view_count = rnt.StrToInt((self).SearchRegexOne("(?s)<div[^>]+class=([\"\\']).*?\\bvideoInfoViews\\b.*?\\1[^>]*>.*?(?P<count>[\\d,.]+)<", webpage, "view count", rnt.NoDefault, false, 0, "count"))
	comment_count = rnt.StrToInt((self).SearchRegexOne(">All [Cc]omments? \\(([\\d,.]+)\\)", webpage, "comment count", rnt.NoDefault, false, 0, nil))
	extract_tag_box = func(regex string, title string) []interface{} {
		var (
			tag_box rnt.OptString
		)
		tag_box = (self).SearchRegexOne(regex, webpage, title, nil, true, 0, nil)
		if !(τ_isTruthy_Os(tag_box)) {
			return []interface{}{}
		}
		return τ_conv_Ls_to_Lα(rnt.ReFindAllOne("<a[^>]+href=[^>]+>([^<]+)", tag_box.Get(), 0))
	}
	categories = extract_tag_box("(?s)Categories:.*?</[^>]+>(.+?)</div>", "categories")
	tags = extract_tag_box("(?s)Tags:.*?</div>\\s*<div[^>]+class=[\"\\']tagBoxContent[\"\\'][^>]*>(.+?)</div>", "tags")
	return rnt.SDict{
		"id":             video_id,
		"display_id":     display_id,
		"title":          title,
		"description":    description,
		"thumbnail":      thumbnail,
		"uploader":       uploader,
		"upload_date":    upload_date,
		"average_rating": average_rating,
		"view_count":     view_count,
		"comment_count":  comment_count,
		"categories":     categories,
		"tags":           tags,
		"age_limit":      age_limit,
		"formats":        formats,
	}
}

func (self *YouPornIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("YouPorn", NewYouPornIE)
}
