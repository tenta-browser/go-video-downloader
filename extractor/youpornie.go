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
}

func NewYouPornIE() rnt.InfoExtractor {
	ret := &YouPornIE{}
	ret.CommonIE = rnt.NewCommonIE()
	ret.VALIDURL = "https?://(?:www\\.)?youporn\\.com/watch/(?P<id>\\d+)/(?P<display_id>[^/?#&]+)"
	return ret
}

func (self *YouPornIE) Key() string {
	return "YouPorn"
}

func (self *YouPornIE) Name() string {
	return "YouPorn extractor"
}

func (self *YouPornIE) _real_extract(url string) map[string]interface{} {
	mobj := rnt.ReMatch((self).VALIDURL, url, 0)
	video_id := rnt.ReMatchGroupOne(mobj, "id")
	display_id := rnt.ReMatchGroupOne(mobj, "display_id")
	request := rnt.SanitizedRequest(url)
	rnt.RequestAddHeader(request, "Cookie", "age_verified=1")
	webpage := (self).DownloadWebpageRequest(request, display_id.Get(), rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, rnt.OptString{}, map[string]interface{}{}, map[string]interface{}{})
	title := func() rnt.OptString {
		if ((self).SearchRegexMulti([]string{"(?:video_titles|videoTitle)\\s*[:=]\\s*([\"\\'])(?P<title>(?:(?!\\1).)+)\\1", "<h1[^>]+class=[\"\\']heading\\d?[\"\\'][^>]*>(?P<title>[^<]+)<"}, webpage, "title", nil, true, 0, "title")).IsSet() && ((self).SearchRegexMulti([]string{"(?:video_titles|videoTitle)\\s*[:=]\\s*([\"\\'])(?P<title>(?:(?!\\1).)+)\\1", "<h1[^>]+class=[\"\\']heading\\d?[\"\\'][^>]*>(?P<title>[^<]+)<"}, webpage, "title", nil, true, 0, "title").Get()) != "" {
			return (self).SearchRegexMulti([]string{"(?:video_titles|videoTitle)\\s*[:=]\\s*([\"\\'])(?P<title>(?:(?!\\1).)+)\\1", "<h1[^>]+class=[\"\\']heading\\d?[\"\\'][^>]*>(?P<title>[^<]+)<"}, webpage, "title", nil, true, 0, "title")
		} else if ((self).OgSearchTitle(webpage, nil, true)).IsSet() && ((self).OgSearchTitle(webpage, nil, true).Get()) != "" {
			return (self).OgSearchTitle(webpage, nil, true)
		} else {
			return (self).HTMLSearchMetaOne("title", webpage, rnt.OptString{}, true, rnt.NoDefault, 0)
		}
	}()
	links := []interface{}{}
	definitions := (self).ParseJSONList((self).SearchRegexOne("mediaDefinition\\s*=\\s*(\\[.+?\\]);", webpage, "media definitions", "[]", true, 0, nil).Get(), video_id.Get(), nil, false)
	if len(definitions) > 0 {
		for _, τmp1 := range definitions {
			definition := τmp1
			if !(func() bool {
				_, ok := (definition).(map[string]interface{})
				return ok
			}()) {
				continue
			}
			video_url := rnt.DictGet(func(val interface{}) map[string]interface{} {
				res, ok := val.(map[string]interface{})
				if !ok {
					panic(rnt.NewCastError(val, `map[string]interface{}`))
				}
				return res
			}(definition), "videoUrl", nil)
			if func() bool {
				_, ok := (video_url).(string)
				return ok
			}() && rnt.IsTruthy(video_url) {
				links = append(links, video_url)
			}
		}
	}
	for _, τmp1 := range rnt.ReFindAllMulti("<a[^>]+href=([\"\\'])(http.+?)\\1[^>]+title=[\"\\']Download [Vv]ideo", webpage, 0) {
		τmp2 := τmp1
		_ = (τmp2)[0]
		link := (τmp2)[1]
		links = append(links, link)
	}
	sources := (self).SearchRegexOne("(?s)sources\\s*:\\s*({.+?})", webpage, "sources", nil, true, 0, nil)
	if (sources).IsSet() && (sources.Get()) != "" {
		for _, τmp1 := range rnt.ReFindAllMulti("[^:]+\\s*:\\s*([\"\\'])(http.+?)\\1", sources.Get(), 0) {
			τmp2 := τmp1
			_ = (τmp2)[0]
			link := (τmp2)[1]
			links = append(links, link)
		}
	}
	for _, τmp1 := range rnt.ReFindAllMulti("(?:videoSrc|videoIpadUrl|html5PlayerSrc)\\s*[:=]\\s*([\"\\'])(http.+?)\\1", webpage, 0) {
		τmp2 := τmp1
		_ = (τmp2)[0]
		link := (τmp2)[1]
		links = append(links, link)
	}
	for _, τmp1 := range rnt.ReFindAllMulti("encryptedQuality\\d{3,4}URL\\s*=\\s*([\"\\'])([\\da-zA-Z+/=]+)\\1", webpage, 0) {
		τmp2 := τmp1
		_ = (τmp2)[0]
		encrypted_link := (τmp2)[1]
		links = append(links, rnt.BytesToStr(rnt.AESDecryptText(encrypted_link, title.Get(), 32), "utf-8"))
	}
	formats := []interface{}{}
	for _, τmp1 := range func() []rnt.OptString {
		τmp2 := []rnt.OptString{}
		for τmp1 := range func() map[rnt.OptString]struct{} {
			τmp1 := map[rnt.OptString]struct{}{}
			for _, τmp2 := range func() []rnt.OptString {
				τmp2 := []rnt.OptString{}
				for _, τmp1 := range links {
					link := τmp1

					τmp2 = append(τmp2, rnt.UnescapeHTML(rnt.CastToOptString(link)))
				}
				return τmp2
			}() {
				τmp1[τmp2] = struct{}{}
			}
			return τmp1
		}() {
			τmp2 = append(τmp2, τmp1)
		}
		return τmp2
	}() {
		video_url := τmp1
		f := map[string]interface{}{"url": video_url}
		mobj = rnt.ReSearch("(?P<height>\\d{3,4})[pP]_(?P<bitrate>\\d+)[kK]_\\d+/", video_url.Get(), 0)
		if mobj != nil {
			height := rnt.ConvertToInt(rnt.ReMatchGroupOne(mobj, "height"))
			bitrate := rnt.ConvertToInt(rnt.ReMatchGroupOne(mobj, "bitrate"))
			rnt.DictUpdate(f, map[string]interface{}{"format_id": rnt.StrFormat("%dp-%dk", height, bitrate),
				"height": height,
				"tbr":    bitrate})
		}
		formats = append(formats, f)
	}
	(self).SortFormats(formats)
	description := (self).OgSearchDescription(webpage, nil)
	thumbnail := (self).SearchRegexOne("(?:imageurl\\s*=|poster\\s*:)\\s*([\"\\'])(?P<thumbnail>.+?)\\1", webpage, "thumbnail", rnt.NoDefault, false, 0, "thumbnail")
	uploader := (self).HTMLSearchRegexOne("(?s)<div[^>]+class=[\"\\']submitByLink[\"\\'][^>]*>(.+?)</div>", webpage, "uploader", rnt.NoDefault, false, 0, nil)
	upload_date := rnt.UnifiedStrDate((self).HTMLSearchRegexMulti([]string{"Date\\s+[Aa]dded:\\s*<span>([^<]+)", "(?s)<div[^>]+class=[\"\\']videoInfo(?:Date|Time)[\"\\'][^>]*>(.+?)</div>"}, webpage, "upload date", rnt.NoDefault, false, 0, nil), true)
	age_limit := (self).RTASearch(webpage)
	average_rating := rnt.IntOrNone((self).SearchRegexOne("<div[^>]+class=[\"\\']videoRatingPercentage[\"\\'][^>]*>(\\d+)%</div>", webpage, "average rating", rnt.NoDefault, false, 0, nil), 1, rnt.OptInt{}, 1)
	view_count := rnt.StrToInt((self).SearchRegexOne("(?s)<div[^>]+class=([\"\\']).*?\\bvideoInfoViews\\b.*?\\1[^>]*>.*?(?P<count>[\\d,.]+)<", webpage, "view count", rnt.NoDefault, false, 0, "count"))
	comment_count := rnt.StrToInt((self).SearchRegexOne(">All [Cc]omments? \\(([\\d,.]+)\\)", webpage, "comment count", rnt.NoDefault, false, 0, nil))

	var extract_tag_box func(string, string) []interface{}
	extract_tag_box = func(regex string, title string) []interface{} {
		tag_box := (self).SearchRegexOne(regex, webpage, title, nil, true, 0, nil)
		if !((tag_box).IsSet() && (tag_box.Get()) != "") {
			return []interface{}{}
		}
		return func() []interface{} {
			τmp2 := []interface{}{}
			for _, τmp1 := range rnt.ReFindAllOne("<a[^>]+href=[^>]+>([^<]+)", tag_box.Get(), 0) {
				τmp2 = append(τmp2, τmp1)
			}
			return τmp2
		}()
	}

	categories := extract_tag_box("(?s)Categories:.*?</[^>]+>(.+?)</div>", "categories")
	tags := extract_tag_box("(?s)Tags:.*?</div>\\s*<div[^>]+class=[\"\\']tagBoxContent[\"\\'][^>]*>(.+?)</div>", "tags")
	return map[string]interface{}{"id": video_id,
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
		"formats":        formats}
}

func (self *YouPornIE) Extract(url string) (*rnt.VideoResult, error) {
	return rnt.RunExtractor(url, self._real_extract)
}

func init() {
	registerFactory("YouPorn", NewYouPornIE)
}
