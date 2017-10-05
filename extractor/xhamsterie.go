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
	*rnt.Context
	_VALID_URL string
}

func NewXHamsterIE(ctx *rnt.Context) rnt.InfoExtractor {
	ret := &XHamsterIE{}
	ret.Context = ctx
	ret._VALID_URL = `(?x)
                    https?://
                        (?:.+?\.)?xhamster\.com/
                        (?:
                            movies/(?P<id>\d+)/(?P<display_id>[^/]*)\.html|
                            videos/(?P<display_id_2>[^/]*)-(?P<id_2>\d+)
                        )
                    `
	return ret
}

func (self *XHamsterIE) Ctx() *rnt.Context {
	return self.Context
}

func (self *XHamsterIE) Key() string {
	return "XHamster"
}

func (self *XHamsterIE) ValidUrl() string {
	return self._VALID_URL
}

func (self *XHamsterIE) Name() string {
	return `XHamster extractor`
}

func (self *XHamsterIE) Tests() []map[string]interface{} {
	return []map[string]interface{}{}
}

func (self *XHamsterIE) _real_extract(url string) map[string]interface{} {
	url = rnt.ReSub(rnt.Re, `^(https?://(?:.+?\.)?)m\.`, `\1`, url, 0, 0)
	mobj := rnt.ReMatch(rnt.Re, (self)._VALID_URL, url, 0)
	video_id := func() rnt.OptString {
		if (rnt.ReMatchGroupOne(mobj, `id`)).IsSet() && (rnt.ReMatchGroupOne(mobj, `id`).Get()) != "" {
			return rnt.ReMatchGroupOne(mobj, `id`)
		} else {
			return rnt.ReMatchGroupOne(mobj, `id_2`)
		}
	}()
	display_id := func() rnt.OptString {
		if (rnt.ReMatchGroupOne(mobj, `display_id`)).IsSet() && (rnt.ReMatchGroupOne(mobj, `display_id`).Get()) != "" {
			return rnt.ReMatchGroupOne(mobj, `display_id`)
		} else {
			return rnt.ReMatchGroupOne(mobj, `display_id_2`)
		}
	}()
	webpage := rnt.DownloadWebpage(self, url, video_id.Get(), rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, rnt.OptString{}, map[string]interface{}{}, map[string]interface{}{})
	error := rnt.HTMLSearchRegex(self, `<div[^>]+id=["\']videoClosed["\'][^>]*>(.+?)</div>`, webpage, `error`, nil, true, 0, nil)
	if (error).IsSet() && (error.Get()) != "" {
		panic(rnt.PyExtractorError(error.Get(), true, rnt.OptString{}))
	}
	title := rnt.HTMLSearchRegex(self, []string{`<h1[^>]*>([^<]+)</h1>`, `<meta[^>]+itemprop=".*?caption.*?"[^>]+content="(.+?)"`, `<title[^>]*>(.+?)(?:,\s*[^,]*?\s*Porn\s*[^,]*?:\s*xHamster[^<]*| - xHamster\.com)</title>`}, webpage, `title`, rnt.NoDefault, true, 0, nil)
	formats := []interface{}{}
	format_urls := func() map[interface{}]struct{} {
		τmp1 := map[interface{}]struct{}{}
		for _, τmp2 := range []interface{}{} {
			τmp1[τmp2] = struct{}{}
		}
		return τmp1
	}()
	sources := rnt.ParseJSON(self, rnt.SearchRegex(self, `sources\s*:\s*({.+?})\s*,?\s*\n`, webpage, `sources`, `{}`, true, 0, nil).Get(), video_id.Get(), 0, false)
	for τmp1, τmp2 := range sources {
		τmp3 := struct {
			_0 string
			_1 interface{}
		}{_0: τmp1,
			_1: τmp2}
		format_id := (τmp3)._0
		format_url := (τmp3)._1
		if !(func() bool {
			_, ok := (format_url).(string)
			return ok
		}()) {
			continue
		}
		if func() bool {
			τmp4 := format_url
			for τmp5 := range format_urls {
				if τmp5 == τmp4 {
					return true
				}
			}
			return false
		}() {
			continue
		}
		format_urls[format_url] = struct{}{}
		formats = append(formats, map[string]interface{}{`format_id`: format_id,
			`url`:    format_url,
			`height`: rnt.IntOrNone(rnt.SearchRegex(self, `^(\d+)[pP]`, format_id, `height`, nil, true, 0, nil), 1, rnt.OptInt{}, 1)})
	}
	video_url := rnt.SearchRegex(self, []string{`file\s*:\s*(?P<q>["'])(?P<mp4>.+?)(?P=q)`, `<a\s+href=(?P<q>["'])(?P<mp4>.+?)(?P=q)\s+class=["']mp4Thumb`, `<video[^>]+file=(?P<q>["'])(?P<mp4>.+?)(?P=q)[^>]*>`}, webpage, `video url`, nil, true, 0, `mp4`)
	if (video_url).IsSet() && (video_url.Get()) != "" && func() bool {
		τmp1 := video_url
		for τmp2 := range format_urls {
			if τmp2 == τmp1 {
				return false
			}
		}
		return true
	}() {
		formats = append(formats, map[string]interface{}{`url`: video_url})
	}
	rnt.SortFormats(self, formats)
	mobj = rnt.ReSearch(rnt.Re, `<span>Description: </span>([^<]+)`, webpage, 0)
	description := func() rnt.OptString {
		if mobj != nil {
			return rnt.ReMatchGroupOne(mobj, 1)
		} else {
			return rnt.OptString{}
		}
	}()
	upload_date := rnt.UnifiedStrDate(rnt.SearchRegex(self, `hint=["\'](\d{4}-\d{2}-\d{2}) \d{2}:\d{2}:\d{2} [A-Z]{3,4}`, webpage, `upload date`, rnt.NoDefault, false, 0, nil), true)
	uploader := rnt.HTMLSearchRegex(self, `<span[^>]+itemprop=["\']author[^>]+><a[^>]+><span[^>]+>([^<]+)`, webpage, `uploader`, `anonymous`, true, 0, nil)
	thumbnail := rnt.SearchRegex(self, []string{`thumb\s*:\s*(?P<q>["'])(?P<thumbnail>.+?)(?P=q)`, `<video[^>]+poster=(?P<q>["'])(?P<thumbnail>.+?)(?P=q)[^>]*>`}, webpage, `thumbnail`, rnt.NoDefault, false, 0, `thumbnail`)
	duration := rnt.ParseDuration(rnt.SearchRegex(self, []string{`<[^<]+\bitemprop=["\']duration["\'][^<]+\bcontent=["\'](.+?)["\']`, `Runtime:\s*</span>\s*([\d:]+)`}, webpage, `duration`, rnt.NoDefault, false, 0, nil))
	view_count := rnt.IntOrNone(rnt.SearchRegex(self, `content=["\']User(?:View|Play)s:(\d+)`, webpage, `view count`, rnt.NoDefault, false, 0, nil), 1, rnt.OptInt{}, 1)
	mobj = rnt.ReSearch(rnt.Re, `hint=[\'"](?P<likecount>\d+) Likes / (?P<dislikecount>\d+) Dislikes`, webpage, 0)
	τmp1 := func() struct {
		_0 rnt.OptString
		_1 rnt.OptString
	} {
		if mobj != nil {
			return struct {
				_0 rnt.OptString
				_1 rnt.OptString
			}{_0: rnt.ReMatchGroupOne(mobj, `likecount`),
				_1: rnt.ReMatchGroupOne(mobj, `dislikecount`)}
		} else {
			return struct {
				_0 rnt.OptString
				_1 rnt.OptString
			}{_0: rnt.OptString{},
				_1: rnt.OptString{}}
		}
	}()
	like_count := (τmp1)._0
	dislike_count := (τmp1)._1
	mobj = rnt.ReSearch(rnt.Re, `</label>Comments \((?P<commentcount>\d+)\)</div>`, webpage, 0)
	comment_count := func() interface{} {
		if mobj != nil {
			return rnt.ReMatchGroupOne(mobj, `commentcount`)
		} else {
			return 0
		}
	}()
	age_limit := rnt.RTASearch(self, webpage)
	categories_html := rnt.SearchRegex(self, `(?s)<table.+?(<span>Categories:.+?)</table>`, webpage, `categories`, nil, true, 0, nil)
	categories := func() []rnt.OptString {
		if (categories_html).IsSet() && (categories_html.Get()) != "" {
			return func() []rnt.OptString {
				τmp3 := []rnt.OptString{}
				for _, τmp2 := range rnt.ReFindAllOne(rnt.Re, `<a[^>]+>(.+?)</a>`, categories_html.Get(), 0) {
					category := τmp2

					τmp3 = append(τmp3, rnt.CleanHTML(rnt.AsOptString(category)))
				}
				return τmp3
			}()
		} else {
			return nil
		}
	}()
	return map[string]interface{}{`id`: video_id,
		`display_id`:    display_id,
		`title`:         title,
		`description`:   description,
		`upload_date`:   upload_date,
		`uploader`:      uploader,
		`thumbnail`:     thumbnail,
		`duration`:      duration,
		`view_count`:    view_count,
		`like_count`:    rnt.IntOrNone(like_count, 1, rnt.OptInt{}, 1),
		`dislike_count`: rnt.IntOrNone(dislike_count, 1, rnt.OptInt{}, 1),
		`comment_count`: rnt.IntOrNone(comment_count, 1, rnt.OptInt{}, 1),
		`age_limit`:     age_limit,
		`categories`:    categories,
		`formats`:       formats}
}

func (self *XHamsterIE) Extract(url string) (*rnt.VideoResult, error) {
	return rnt.RunExtractor(url, self._real_extract)
}

func init() {
	registerFactory(`XHamster`, NewXHamsterIE)
}
