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
 * pornhubie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/pornhub.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type PornHubIE struct {
	*rnt.CommonIE
	IE_DESC string
}

func NewPornHubIE() rnt.InfoExtractor {
	ret := &PornHubIE{}
	ret.CommonIE = rnt.NewCommonIE()
	ret.IE_DESC = `PornHub and Thumbzilla`
	ret.VALIDURL = `(?x)
                    https?://
                        (?:
                            (?:[a-z]+\.)?pornhub\.com/(?:(?:view_video\.php|video/show)\?viewkey=|embed/)|
                            (?:www\.)?thumbzilla\.com/video/
                        )
                        (?P<id>[\da-z]+)
                    `
	return ret
}

func (self *PornHubIE) Key() string {
	return "PornHub"
}

func (self *PornHubIE) Name() string {
	return `PornHub extractor` + " (" + self.IE_DESC + ")"
}

func (self *PornHubIE) _extract_urls(webpage string) []string {
	return rnt.ReFindAllOne(`<iframe[^>]+?src=["\'](?P<url>(?:https?:)?//(?:www\.)?pornhub\.com/embed/[\da-z]+)`, webpage, 0)
}

func (self *PornHubIE) _extract_count(pattern string, webpage string, name string) int {
	return rnt.StrToInt((self).SearchRegex(pattern, webpage, rnt.StrFormat(`%s count`, name), rnt.NoDefault, false, 0, nil)).Get()
}

func (self *PornHubIE) _real_extract(url string) map[string]interface{} {
	video_id := (self).MatchID(url)

	var dl_webpage func(string) string
	dl_webpage = func(platform string) string {
		return (self).DownloadWebpage(rnt.StrFormat(`http://www.pornhub.com/view_video.php?viewkey=%s`, video_id), video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, rnt.OptString{}, map[string]interface{}{`Cookie`: rnt.StrFormat(`age_verified=1; platform=%s`, platform)}, map[string]interface{}{})
	}

	webpage := dl_webpage(`pc`)
	error_msg := (self).HTMLSearchRegex(`(?s)<div[^>]+class=(["\'])(?:(?!\1).)*\b(?:removed|userMessageSection)\b(?:(?!\1).)*\1[^>]*>(?P<error>.+?)</div>`, webpage, `error message`, nil, true, 0, `error`)
	if (error_msg).IsSet() && (error_msg.Get()) != "" {
		error_msg = rnt.AsOptString(rnt.ReSub(`\s+`, ` `, error_msg.Get(), 0, 0))
		panic(rnt.PyExtractorError(rnt.StrFormat(`PornHub said: %s`, error_msg), true, rnt.AsOptString(video_id)))
	}
	tv_webpage := dl_webpage(`tv`)
	assignments := rnt.StrSplit((self).SearchRegex(`(var.+?mediastring.+?)</script>`, tv_webpage, `encoded url`, rnt.NoDefault, true, 0, nil).Get(), `;`, -(1))
	js_vars := map[string]interface{}{}

	var parse_js_value func(string) interface{}
	parse_js_value = func(inp string) interface{} {
		inp = rnt.ReSub(`/\*(?:(?!\*/).)*?\*/`, ``, inp, 0, 0)
		if rnt.StrContains(inp, `+`) {
			inps := rnt.StrSplit(inp, `+`, -(1))
			return rnt.FuncReduce(rnt.OperatorConcat, func() []interface{} {
				τmp3 := []interface{}{}
				for _, τmp2 := range inps {
					τmp1 := τmp2

					τmp3 = append(τmp3, parse_js_value(τmp1))
				}
				return τmp3
			}())
		}
		inp = rnt.StrStrip(inp, ``)
		if rnt.DictContains(js_vars, inp) {
			return (js_vars)[inp]
		}
		return rnt.RemoveQuotes(rnt.AsOptString(inp))
	}

	for _, τmp1 := range assignments {
		assn := τmp1
		assn = rnt.StrStrip(assn, ``)
		if !((assn) != "") {
			continue
		}
		assn = rnt.ReSub(`var\s+`, ``, assn, 0, 0)
		τmp2 := rnt.StrSplit(assn, `=`, 1)
		vname := (τmp2)[0]
		value := (τmp2)[1]
		(js_vars)[vname] = parse_js_value(value)
	}
	video_url := (js_vars)[`mediastring`]
	title := (self).SearchRegex(`<h1>([^>]+)</h1>`, tv_webpage, `title`, nil, true, 0, nil)
	title = func() rnt.OptString {
		if (title).IsSet() && (title.Get()) != "" {
			return title
		} else if ((self).HTMLSearchMeta(`twitter:title`, webpage, rnt.OptString{}, false, nil, 0)).IsSet() && ((self).HTMLSearchMeta(`twitter:title`, webpage, rnt.OptString{}, false, nil, 0).Get()) != "" {
			return (self).HTMLSearchMeta(`twitter:title`, webpage, rnt.OptString{}, false, nil, 0)
		} else {
			return (self).SearchRegex(struct {
				_0 string
				_1 string
				_2 string
			}{_0: `<h1[^>]+class=["\']title["\'][^>]*>(?P<title>[^<]+)`,
				_1: `<div[^>]+data-video-title=(["\'])(?P<title>.+?)\1`,
				_2: `shareTitle\s*=\s*(["\'])(?P<title>.+?)\1`}, webpage, `title`, rnt.NoDefault, true, 0, `title`)
		}
	}()
	flashvars := (self).ParseJSON((self).SearchRegex(`var\s+flashvars_\d+\s*=\s*({.+?});`, webpage, `flashvars`, `{}`, true, 0, nil).Get(), video_id, nil, true)
	thumbnail := ``
	duration := 0
	if len(flashvars) > 0 {
		thumbnail = rnt.CastToString(rnt.DictGet(flashvars, `image_url`, nil))
		duration = rnt.IntOrNone(rnt.DictGet(flashvars, `video_duration`, nil), 1, rnt.OptInt{}, 1).Get()
	} else {
		τmp1 := func(τmp2 []interface{}, τmp3 int) []interface{} {
			τmp4 := []interface{}{}
			for τmp5 := 0; τmp5 < τmp3; τmp5++ {
				τmp4 = append(τmp4, τmp2...)
			}
			return τmp4
		}([]interface{}{nil}, 3)
		title = rnt.CastToOptString((τmp1)[0])
		thumbnail = rnt.CastToString((τmp1)[1])
		duration = rnt.CastToInt((τmp1)[2])
	}
	video_uploader := (self).HTMLSearchRegex(`(?s)From:&nbsp;.+?<(?:a\b[^>]+\bhref=["\']/(?:user|channel)s/|span\b[^>]+\bclass=["\']username)[^>]+>(.+?)<`, webpage, `uploader`, rnt.NoDefault, false, 0, nil)
	view_count := (self)._extract_count(`<span class="count">([\d,\.]+)</span> views`, webpage, `view`)
	like_count := (self)._extract_count(`<span class="votesUp">([\d,\.]+)</span>`, webpage, `like`)
	dislike_count := (self)._extract_count(`<span class="votesDown">([\d,\.]+)</span>`, webpage, `dislike`)
	comment_count := (self)._extract_count(`All Comments\s*<span>\(([\d,.]+)\)`, webpage, `comment`)
	page_params := (self).ParseJSON((self).SearchRegex(`page_params\.zoneDetails\[([\'"])[^\'"]+\1\]\s*=\s*(?P<data>{[^}]+})`, webpage, `page parameters`, `{}`, true, 0, `data`).Get(), video_id, rnt.JsToJSON, false)
	τmp1 := rnt.StrSplit(``, `,`, -(1))
	tags := τmp1
	categories := τmp1
	if len(page_params) > 0 {
		tags = rnt.StrSplit(rnt.CastToString(rnt.DictGet(page_params, `tags`, ``)), `,`, -(1))
		categories = rnt.StrSplit(rnt.CastToString(rnt.DictGet(page_params, `categories`, ``)), `,`, -(1))
	}
	return map[string]interface{}{`id`: video_id,
		`url`:           video_url,
		`uploader`:      video_uploader,
		`title`:         title,
		`thumbnail`:     thumbnail,
		`duration`:      duration,
		`view_count`:    view_count,
		`like_count`:    like_count,
		`dislike_count`: dislike_count,
		`comment_count`: comment_count,
		`age_limit`:     18,
		`tags`:          tags,
		`categories`:    categories}
}

func (self *PornHubIE) Extract(url string) (*rnt.VideoResult, error) {
	return rnt.RunExtractor(url, self._real_extract)
}

func init() {
	registerFactory(`PornHub`, NewPornHubIE)
}
