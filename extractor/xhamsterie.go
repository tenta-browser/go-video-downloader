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
 * xhamsterie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/xhamster.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type XHamsterIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewXHamsterIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &XHamsterIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "XHamster"
	_VALID_URL = "(?x)\n                    https?://\n                        (?:.+?\\.)?xhamster\\.com/\n                        (?:\n                            movies/(?P<id>\\d+)/(?P<display_id>[^/]*)\\.html|\n                            videos/(?P<display_id_2>[^/]*)-(?P<id_2>\\d+)\n                        )\n                    "
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *XHamsterIE) Key() string {
	return "XHamster"
}

func (self *XHamsterIE) Name() string {
	return self.IE_NAME
}

func (self *XHamsterIE) _real_extract(url string) rnt.SDict {
	var (
		age_limit       int
		c               interface{}
		c_name          interface{}
		categories      []interface{}
		categories_html rnt.OptString
		categories_list interface{}
		comment_count   interface{}
		description     rnt.OptString
		desktop_url     string
		dislike_count   rnt.OptString
		display_id      rnt.OptString
		duration        rnt.OptString
		error           rnt.OptString
		filesize        rnt.OptInt
		format_id       string
		format_item     interface{}
		format_url      interface{}
		format_urls     map[interface{}]struct{}
		formats         []interface{}
		formats_dict    interface{}
		get_height      func(string) rnt.OptInt
		initials        interface{}
		like_count      rnt.OptString
		mobj            rnt.Match
		quality         string
		sources         interface{}
		thumbnail       rnt.OptString
		title           interface{}
		upload_date     rnt.OptString
		uploader        rnt.OptString
		video           interface{}
		video_id        rnt.OptString
		video_url       rnt.OptString
		view_count      rnt.OptInt
		webpage         string
		τmp1            τ_Tsαω
		τmp2            τ_Tsαω
		τmp3            τ_Tsαω
		τmp4            τ_TOsOsω
	)
	url = rnt.ReSub("^(https?://(?:.+?\\.)?)m\\.", "\\1", url, 0, 0)
	mobj = rnt.ReMatch((self).VALIDURL, url, 0)
	video_id = func() rnt.OptString {
		if v := (rnt.ReMatchGroupOne(mobj, "id")); τ_isTruthy_Os(v) {
			return v
		} else {
			return rnt.ReMatchGroupOne(mobj, "id_2")
		}
	}()
	display_id = func() rnt.OptString {
		if v := (rnt.ReMatchGroupOne(mobj, "display_id")); τ_isTruthy_Os(v) {
			return v
		} else {
			return rnt.ReMatchGroupOne(mobj, "display_id_2")
		}
	}()
	desktop_url = rnt.ReSub("^(https?://(?:.+?\\.)?)m\\.", "\\1", url, 0, 0)
	webpage = (self).DownloadWebpageURL(desktop_url, video_id.Get(), rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{}, nil)
	error = (self).HTMLSearchRegexOne("<div[^>]+id=[\"\\']videoClosed[\"\\'][^>]*>(.+?)</div>", webpage, "error", nil, true, 0, nil)
	if τ_isTruthy_Os(error) {
		panic(rnt.PyExtractorError(error.Get(), true, rnt.OptString{}))
	}
	age_limit = (self).RTASearch(webpage)
	get_height = func(s string) rnt.OptInt {
		return rnt.IntOrNone((self).SearchRegexOne("^(\\d+)[pP]", s, "height", nil, true, 0, nil), 1, rnt.OptInt{}, 1)
	}
	initials = (self).ParseJSON((self).SearchRegexOne("window\\.initials\\s*=\\s*({.+?})\\s*;\\s*\\n", webpage, "initials", "{}", true, 0, nil).Get(), video_id.Get(), nil, false)
	if rnt.IsTruthy(initials) {
		video = rnt.UnsafeSubscript(initials, "videoModel")
		title = rnt.UnsafeSubscript(video, "title")
		formats = []interface{}{}
		for _, τel := range rnt.DictItems(τ_cast_α_to_d(rnt.UnsafeSubscript(video, "sources"))) {
			τmp1 = τel
			format_id = (τmp1).Φ0
			formats_dict = (τmp1).Φ1
			if !(τ_isinstance_d(formats_dict)) {
				continue
			}
			for _, τel := range rnt.DictItems(τ_cast_α_to_d(formats_dict)) {
				τmp2 = τel
				quality = (τmp2).Φ0
				format_item = (τmp2).Φ1
				if (format_id) == ("download") {
					continue
					if !(τ_isinstance_d(format_item)) {
						continue
					}
					format_url = rnt.DictGet(τ_cast_α_to_d(format_item), "link", nil)
					filesize = rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(format_item), "size", nil), 1, rnt.OptInt{}, 1000000)
				} else {
					format_url = format_item
					filesize = rnt.OptInt{}
				}
				format_url = rnt.URLOrNone(rnt.CastToOptString(format_url))
				if !(rnt.IsTruthy(format_url)) {
					continue
				}
				formats = append(formats, rnt.SDict{
					"format_id": rnt.StrFormat2("%s-%s", format_id, quality),
					"url":       format_url,
					"ext":       rnt.DetermineExt(rnt.CastToOptString(format_url), "mp4"),
					"height":    get_height(quality),
					"filesize":  filesize,
				})
			}
		}
		formats = τ_conv_Ld_to_Lα((self).SortFormats(τ_conv_Lα_to_Ld(formats)))
		categories_list = rnt.DictGet(τ_cast_α_to_d(video), "categories", nil)
		if rnt.IsList(categories_list) {
			categories = []interface{}{}
			for _, τel := range τ_cast_α_to_Lα(categories_list) {
				c = τel
				if !(τ_isinstance_d(c)) {
					continue
				}
				c_name = rnt.DictGet(τ_cast_α_to_d(c), "name", nil)
				if τ_isinstance_s(c_name) {
					categories = append(categories, c_name)
				}
			}
		} else {
			categories = nil
		}
		return rnt.SDict{
			"id":          video_id,
			"display_id":  display_id,
			"title":       title,
			"description": rnt.DictGet(τ_cast_α_to_d(video), "description", nil),
			"timestamp":   rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(video), "created", nil), 1, rnt.OptInt{}, 1),
			"uploader": rnt.TryGetOne(video, func(x interface{}) interface{} {
				return rnt.UnsafeSubscript(rnt.UnsafeSubscript(x, "author"), "name")
			}, rnt.ConvertToString),
			"thumbnail":  rnt.DictGet(τ_cast_α_to_d(video), "thumbURL", nil),
			"duration":   rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(video), "duration", nil), 1, rnt.OptInt{}, 1),
			"view_count": rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(video), "views", nil), 1, rnt.OptInt{}, 1),
			"like_count": rnt.IntOrNone(rnt.TryGetOne(video, func(x interface{}) interface{} {
				return rnt.UnsafeSubscript(rnt.UnsafeSubscript(x, "rating"), "likes")
			}, rnt.ConvertToInt), 1, rnt.OptInt{}, 1),
			"dislike_count": rnt.IntOrNone(rnt.TryGetOne(video, func(x interface{}) interface{} {
				return rnt.UnsafeSubscript(rnt.UnsafeSubscript(x, "rating"), "dislikes")
			}, rnt.ConvertToInt), 1, rnt.OptInt{}, 1),
			"comment_count": rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(video), "views", nil), 1, rnt.OptInt{}, 1),
			"age_limit":     age_limit,
			"categories":    categories,
			"formats":       formats,
		}
	}
	title = (self).HTMLSearchRegexMulti([]string{"<h1[^>]*>([^<]+)</h1>", "<meta[^>]+itemprop=\".*?caption.*?\"[^>]+content=\"(.+?)\"", "<title[^>]*>(.+?)(?:,\\s*[^,]*?\\s*Porn\\s*[^,]*?:\\s*xHamster[^<]*| - xHamster\\.com)</title>"}, webpage, "title", rnt.NoDefault, true, 0, nil)
	formats = []interface{}{}
	format_urls = τ_constr_Sα_from_Lα([]interface{}{})
	sources = (self).ParseJSON((self).SearchRegexOne("sources\\s*:\\s*({.+?})\\s*,?\\s*\\n", webpage, "sources", "{}", true, 0, nil).Get(), video_id.Get(), nil, false)
	for _, τel := range rnt.DictItems(τ_cast_α_to_d(sources)) {
		τmp3 = τel
		format_id = (τmp3).Φ0
		format_url = (τmp3).Φ1
		format_url = rnt.URLOrNone(rnt.CastToOptString(format_url))
		if !(rnt.IsTruthy(format_url)) {
			continue
		}
		if τ_search_α_in_Sα(format_url, format_urls) {
			continue
		}
		format_urls[format_url] = struct{}{}
		formats = append(formats, rnt.SDict{
			"format_id": format_id,
			"url":       format_url,
			"height":    get_height(format_id),
		})
	}
	video_url = (self).SearchRegexMulti([]string{"file\\s*:\\s*(?P<q>[\"'])(?P<mp4>.+?)(?P=q)", "<a\\s+href=(?P<q>[\"'])(?P<mp4>.+?)(?P=q)\\s+class=[\"']mp4Thumb", "<video[^>]+file=(?P<q>[\"'])(?P<mp4>.+?)(?P=q)[^>]*>"}, webpage, "video url", nil, true, 0, "mp4")
	if rnt.IsTruthy(func() interface{} {
		if v := (video_url); !(τ_isTruthy_Os(v)) {
			return v
		} else {
			return τ_search_s_notin_Sα(video_url.Get(), format_urls)
		}
	}()) {
		formats = append(formats, rnt.SDict{
			"url": video_url,
		})
	}
	formats = τ_conv_Ld_to_Lα((self).SortFormats(τ_conv_Lα_to_Ld(formats)))
	mobj = rnt.ReSearch("<span>Description: </span>([^<]+)", webpage, 0)
	description = func() rnt.OptString {
		if mobj != nil {
			return rnt.ReMatchGroupOne(mobj, 1)
		} else {
			return rnt.OptString{}
		}
	}()
	upload_date = rnt.UnifiedStrDate((self).SearchRegexOne("hint=[\"\\'](\\d{4}-\\d{2}-\\d{2}) \\d{2}:\\d{2}:\\d{2} [A-Z]{3,4}", webpage, "upload date", rnt.NoDefault, false, 0, nil), true)
	uploader = (self).HTMLSearchRegexOne("<span[^>]+itemprop=[\"\\']author[^>]+><a[^>]+><span[^>]+>([^<]+)", webpage, "uploader", "anonymous", true, 0, nil)
	thumbnail = (self).SearchRegexMulti([]string{"[\"']thumbUrl[\"']\\s*:\\s*(?P<q>[\"'])(?P<thumbnail>.+?)(?P=q)", "<video[^>]+\"poster\"=(?P<q>[\"'])(?P<thumbnail>.+?)(?P=q)[^>]*>"}, webpage, "thumbnail", rnt.NoDefault, false, 0, "thumbnail")
	duration = rnt.ParseDuration((self).SearchRegexMulti([]string{"<[^<]+\\bitemprop=[\"\\']duration[\"\\'][^<]+\\bcontent=[\"\\'](.+?)[\"\\']", "Runtime:\\s*</span>\\s*([\\d:]+)"}, webpage, "duration", rnt.NoDefault, false, 0, nil))
	view_count = rnt.IntOrNone((self).SearchRegexOne("content=[\"\\']User(?:View|Play)s:(\\d+)", webpage, "view count", rnt.NoDefault, false, 0, nil), 1, rnt.OptInt{}, 1)
	mobj = rnt.ReSearch("hint=[\\'\"](?P<likecount>\\d+) Likes / (?P<dislikecount>\\d+) Dislikes", webpage, 0)
	τmp4 = func() τ_TOsOsω {
		if mobj != nil {
			return τ_TOsOsω{
				Φ0: rnt.ReMatchGroupOne(mobj, "likecount"),
				Φ1: rnt.ReMatchGroupOne(mobj, "dislikecount"),
			}
		} else {
			return τ_conv_T00ω_to_TOsOsω(τ_T00ω{
				Φ0: nil,
				Φ1: nil,
			})
		}
	}()
	like_count = (τmp4).Φ0
	dislike_count = (τmp4).Φ1
	mobj = rnt.ReSearch("</label>Comments \\((?P<commentcount>\\d+)\\)</div>", webpage, 0)
	comment_count = func() interface{} {
		if mobj != nil {
			return rnt.ReMatchGroupOne(mobj, "commentcount")
		} else {
			return 0
		}
	}()
	categories_html = (self).SearchRegexOne("(?s)<table.+?(<span>Categories:.+?)</table>", webpage, "categories", nil, true, 0, nil)
	categories = τ_conv_LOs_to_Lα(func() []rnt.OptString {
		if τ_isTruthy_Os(categories_html) {
			return func() []rnt.OptString {
				τresult := []rnt.OptString{}
				for _, τel := range rnt.ReFindAllOne("<a[^>]+>(.+?)</a>", categories_html.Get(), 0) {
					var (
						category string
					)
					category = τel
					τresult = append(τresult, rnt.CleanHTML(rnt.AsOptString(category)))
				}
				return τresult
			}()
		} else {
			return nil
		}
	}())
	return rnt.SDict{
		"id":            video_id,
		"display_id":    display_id,
		"title":         title,
		"description":   description,
		"upload_date":   upload_date,
		"uploader":      uploader,
		"thumbnail":     thumbnail,
		"duration":      duration,
		"view_count":    view_count,
		"like_count":    rnt.IntOrNone(like_count, 1, rnt.OptInt{}, 1),
		"dislike_count": rnt.IntOrNone(dislike_count, 1, rnt.OptInt{}, 1),
		"comment_count": rnt.IntOrNone(comment_count, 1, rnt.OptInt{}, 1),
		"age_limit":     age_limit,
		"categories":    categories,
		"formats":       formats,
	}
}

func (self *XHamsterIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("XHamster", NewXHamsterIE)
}
