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
		categories      []rnt.OptString
		categories_html rnt.OptString
		comment_count   interface{}
		description     rnt.OptString
		dislike_count   rnt.OptString
		display_id      rnt.OptString
		duration        rnt.OptString
		error           rnt.OptString
		format_id       string
		format_url      interface{}
		format_urls     map[interface{}]struct{}
		formats         []interface{}
		like_count      rnt.OptString
		mobj            rnt.Match
		sources         interface{}
		thumbnail       rnt.OptString
		title           rnt.OptString
		upload_date     rnt.OptString
		uploader        rnt.OptString
		video_id        rnt.OptString
		video_url       rnt.OptString
		view_count      rnt.OptInt
		webpage         string
		τmp1            τ_Tsαω
		τmp2            τ_TOsOsω
	)
	url = rnt.ReSub("^(https?://(?:.+?\\.)?)m\\.", "\\1", url, 0, 0)
	mobj = rnt.ReMatch((self).VALIDURL, url, 0)
	video_id = func() rnt.OptString {
		if τ_isTruthy_Os(rnt.ReMatchGroupOne(mobj, "id")) {
			return rnt.ReMatchGroupOne(mobj, "id")
		} else {
			return rnt.ReMatchGroupOne(mobj, "id_2")
		}
	}()
	display_id = func() rnt.OptString {
		if τ_isTruthy_Os(rnt.ReMatchGroupOne(mobj, "display_id")) {
			return rnt.ReMatchGroupOne(mobj, "display_id")
		} else {
			return rnt.ReMatchGroupOne(mobj, "display_id_2")
		}
	}()
	webpage = (self).DownloadWebpageURL(url, video_id.Get(), rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	error = (self).HTMLSearchRegexOne("<div[^>]+id=[\"\\']videoClosed[\"\\'][^>]*>(.+?)</div>", webpage, "error", nil, true, 0, nil)
	if τ_isTruthy_Os(error) {
		panic(rnt.PyExtractorError(error.Get(), true, rnt.OptString{}))
	}
	title = (self).HTMLSearchRegexMulti([]string{"<h1[^>]*>([^<]+)</h1>", "<meta[^>]+itemprop=\".*?caption.*?\"[^>]+content=\"(.+?)\"", "<title[^>]*>(.+?)(?:,\\s*[^,]*?\\s*Porn\\s*[^,]*?:\\s*xHamster[^<]*| - xHamster\\.com)</title>"}, webpage, "title", rnt.NoDefault, true, 0, nil)
	formats = []interface{}{}
	format_urls = τ_constr_Sα_from_Lα([]interface{}{})
	sources = (self).ParseJSON((self).SearchRegexOne("sources\\s*:\\s*({.+?})\\s*,?\\s*\\n", webpage, "sources", "{}", true, 0, nil).Get(), video_id.Get(), nil, false)
	for _, τel := range rnt.DictItems(τ_cast_α_to_d(sources)) {
		τmp1 = τel
		format_id = (τmp1).Φ0
		format_url = (τmp1).Φ1
		if !(τ_isinstance_s(format_url)) {
			continue
		}
		if τ_search_α_in_Sα(format_url, format_urls) {
			continue
		}
		format_urls[format_url] = struct{}{}
		formats = append(formats, rnt.SDict{
			"format_id": format_id,
			"url":       format_url,
			"height":    rnt.IntOrNone((self).SearchRegexOne("^(\\d+)[pP]", format_id, "height", nil, true, 0, nil), 1, rnt.OptInt{}, 1),
		})
	}
	video_url = (self).SearchRegexMulti([]string{"file\\s*:\\s*(?P<q>[\"'])(?P<mp4>.+?)(?P=q)", "<a\\s+href=(?P<q>[\"'])(?P<mp4>.+?)(?P=q)\\s+class=[\"']mp4Thumb", "<video[^>]+file=(?P<q>[\"'])(?P<mp4>.+?)(?P=q)[^>]*>"}, webpage, "video url", nil, true, 0, "mp4")
	if rnt.IsTruthy(func() interface{} {
		if !(τ_isTruthy_Os(video_url)) {
			return video_url
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
	thumbnail = (self).SearchRegexMulti([]string{"thumb\\s*:\\s*(?P<q>[\"'])(?P<thumbnail>.+?)(?P=q)", "<video[^>]+poster=(?P<q>[\"'])(?P<thumbnail>.+?)(?P=q)[^>]*>"}, webpage, "thumbnail", rnt.NoDefault, false, 0, "thumbnail")
	duration = rnt.ParseDuration((self).SearchRegexMulti([]string{"<[^<]+\\bitemprop=[\"\\']duration[\"\\'][^<]+\\bcontent=[\"\\'](.+?)[\"\\']", "Runtime:\\s*</span>\\s*([\\d:]+)"}, webpage, "duration", rnt.NoDefault, false, 0, nil))
	view_count = rnt.IntOrNone((self).SearchRegexOne("content=[\"\\']User(?:View|Play)s:(\\d+)", webpage, "view count", rnt.NoDefault, false, 0, nil), 1, rnt.OptInt{}, 1)
	mobj = rnt.ReSearch("hint=[\\'\"](?P<likecount>\\d+) Likes / (?P<dislikecount>\\d+) Dislikes", webpage, 0)
	τmp2 = func() τ_TOsOsω {
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
	like_count = (τmp2).Φ0
	dislike_count = (τmp2).Φ1
	mobj = rnt.ReSearch("</label>Comments \\((?P<commentcount>\\d+)\\)</div>", webpage, 0)
	comment_count = func() interface{} {
		if mobj != nil {
			return rnt.ReMatchGroupOne(mobj, "commentcount")
		} else {
			return 0
		}
	}()
	age_limit = (self).RTASearch(webpage)
	categories_html = (self).SearchRegexOne("(?s)<table.+?(<span>Categories:.+?)</table>", webpage, "categories", nil, true, 0, nil)
	categories = func() []rnt.OptString {
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
	}()
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
