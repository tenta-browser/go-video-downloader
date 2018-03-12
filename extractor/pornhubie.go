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
	IE_NAME string
}

func NewPornHubIE() rnt.InfoExtractor {
	var (
		IE_DESC    string
		IE_NAME    string
		_VALID_URL string
	)
	self := &PornHubIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "PornHub"
	IE_DESC = "PornHub and Thumbzilla"
	_VALID_URL = "(?x)\n                    https?://\n                        (?:\n                            (?:[a-z]+\\.)?pornhub\\.com/(?:(?:view_video\\.php|video/show)\\?viewkey=|embed/)|\n                            (?:www\\.)?thumbzilla\\.com/video/\n                        )\n                        (?P<id>[\\da-z]+)\n                    "
	self.IE_DESC = IE_DESC
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *PornHubIE) Key() string {
	return "PornHub"
}

func (self *PornHubIE) Name() string {
	return self.IE_NAME
}

func (self *PornHubIE) _extract_count(pattern string, webpage string, name string) rnt.OptInt {
	return rnt.StrToInt((self).SearchRegexOne(pattern, webpage, rnt.StrFormat2("%s count", name), rnt.NoDefault, false, 0, nil))
}

func (self *PornHubIE) _real_extract(url string) rnt.SDict {
	var (
		assignments    []string
		assn           string
		categories     []string
		comment_count  rnt.OptInt
		dislike_count  rnt.OptInt
		dl_webpage     func(string) string
		duration       rnt.OptInt
		error_msg      rnt.OptString
		flashvars      interface{}
		js_vars        rnt.SDict
		like_count     rnt.OptInt
		page_params    interface{}
		parse_js_value func(string) interface{}
		tags           []string
		thumbnail      interface{}
		title          rnt.OptString
		tv_webpage     string
		value          string
		video_id       string
		video_uploader rnt.OptString
		video_url      interface{}
		view_count     rnt.OptInt
		vname          string
		webpage        string
		τmp1           []string
		τmp2           []interface{}
		τmp3           interface{}
	)
	video_id = (self).MatchID(url)
	(self).SetCookie("pornhub.com", "age_verified", "1")
	dl_webpage = func(platform string) string {
		(self).SetCookie("pornhub.com", "platform", platform)
		return (self).DownloadWebpageURL(rnt.StrFormat2("http://www.pornhub.com/view_video.php?viewkey=%s", video_id), video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	}
	webpage = dl_webpage("pc")
	error_msg = (self).HTMLSearchRegexOne("(?s)<div[^>]+class=([\"\\'])(?:(?!\\1).)*\\b(?:removed|userMessageSection)\\b(?:(?!\\1).)*\\1[^>]*>(?P<error>.+?)</div>", webpage, "error message", nil, true, 0, "error")
	if τ_isTruthy_Os(error_msg) {
		error_msg = rnt.AsOptString(rnt.ReSub("\\s+", " ", error_msg.Get(), 0, 0))
		panic(rnt.PyExtractorError(rnt.StrFormat2("PornHub said: %s", error_msg), true, rnt.AsOptString(video_id)))
	}
	tv_webpage = dl_webpage("tv")
	assignments = rnt.StrSplit((self).SearchRegexOne("(var.+?mediastring.+?)</script>", tv_webpage, "encoded url", rnt.NoDefault, true, 0, nil).Get(), ";", -(1))
	js_vars = rnt.SDict{}
	parse_js_value = func(inp string) interface{} {
		var (
			inps []string
		)
		inp = rnt.ReSub("/\\*(?:(?!\\*/).)*?\\*/", "", inp, 0, 0)
		if rnt.StrContains(inp, "+") {
			inps = rnt.StrSplit(inp, "+", -(1))
			return rnt.FuncReduce(rnt.OperatorConcat, τ_map_Fsαω_over_Ls(parse_js_value, inps))
		}
		inp = rnt.StrStrip(inp, "")
		if rnt.DictContains(js_vars, inp) {
			return (js_vars)[inp]
		}
		return rnt.RemoveQuotes(rnt.AsOptString(inp))
	}
	for _, τel := range assignments {
		assn = τel
		assn = rnt.StrStrip(assn, "")
		if !((assn) != "") {
			continue
		}
		assn = rnt.ReSub("var\\s+", "", assn, 0, 0)
		τmp1 = rnt.StrSplit(assn, "=", 1)
		vname = (τmp1)[0]
		value = (τmp1)[1]
		(js_vars)[vname] = parse_js_value(value)
	}
	video_url = (js_vars)["mediastring"]
	title = (self).SearchRegexOne("<h1>([^>]+)</h1>", tv_webpage, "title", nil, true, 0, nil)
	title = func() rnt.OptString {
		if v := (title); τ_isTruthy_Os(v) {
			return v
		} else if v := ((self).HTMLSearchMetaOne("twitter:title", webpage, rnt.OptString{}, false, nil, 0)); τ_isTruthy_Os(v) {
			return v
		} else {
			return (self).SearchRegexMulti(τ_conv_Tsssω_to_Ls(τ_Tsssω{
				Φ0: "<h1[^>]+class=[\"\\']title[\"\\'][^>]*>(?P<title>[^<]+)",
				Φ1: "<div[^>]+data-video-title=([\"\\'])(?P<title>.+?)\\1",
				Φ2: "shareTitle\\s*=\\s*([\"\\'])(?P<title>.+?)\\1",
			}), webpage, "title", rnt.NoDefault, true, 0, "title")
		}
	}()
	flashvars = (self).ParseJSON((self).SearchRegexOne("var\\s+flashvars_\\d+\\s*=\\s*({.+?});", webpage, "flashvars", "{}", true, 0, nil).Get(), video_id, nil, true)
	if rnt.IsTruthy(flashvars) {
		thumbnail = rnt.DictGet(τ_cast_α_to_d(flashvars), "image_url", nil)
		duration = rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(flashvars), "video_duration", nil), 1, rnt.OptInt{}, 1)
	} else {
		τmp2 = τ_multiply_L0_by_i([]interface{}{nil}, 3)
		title = τ_sink_Os((τmp2)[0])
		thumbnail = (τmp2)[1]
		duration = τ_sink_Oi((τmp2)[2])
	}
	video_uploader = (self).HTMLSearchRegexOne("(?s)From:&nbsp;.+?<(?:a\\b[^>]+\\bhref=[\"\\']/(?:user|channel)s/|span\\b[^>]+\\bclass=[\"\\']username)[^>]+>(.+?)<", webpage, "uploader", rnt.NoDefault, false, 0, nil)
	view_count = (self)._extract_count("<span class=\"count\">([\\d,\\.]+)</span> views", webpage, "view")
	like_count = (self)._extract_count("<span class=\"votesUp\">([\\d,\\.]+)</span>", webpage, "like")
	dislike_count = (self)._extract_count("<span class=\"votesDown\">([\\d,\\.]+)</span>", webpage, "dislike")
	comment_count = (self)._extract_count("All Comments\\s*<span>\\(([\\d,.]+)\\)", webpage, "comment")
	page_params = (self).ParseJSON((self).SearchRegexOne("page_params\\.zoneDetails\\[([\\'\"])[^\\'\"]+\\1\\]\\s*=\\s*(?P<data>{[^}]+})", webpage, "page parameters", "{}", true, 0, "data").Get(), video_id, rnt.JsToJSON, false)
	τmp3 = nil
	tags = τ_sink_Ls(τmp3)
	categories = τ_sink_Ls(τmp3)
	if rnt.IsTruthy(page_params) {
		tags = rnt.StrSplit(rnt.CastToString(rnt.DictGet(τ_cast_α_to_d(page_params), "tags", "")), ",", -(1))
		categories = rnt.StrSplit(rnt.CastToString(rnt.DictGet(τ_cast_α_to_d(page_params), "categories", "")), ",", -(1))
	}
	return rnt.SDict{
		"id":            video_id,
		"url":           video_url,
		"uploader":      video_uploader,
		"title":         title,
		"thumbnail":     thumbnail,
		"duration":      duration,
		"view_count":    view_count,
		"like_count":    like_count,
		"dislike_count": dislike_count,
		"comment_count": comment_count,
		"age_limit":     18,
		"tags":          tags,
		"categories":    categories,
	}
}

func (self *PornHubIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("PornHub", NewPornHubIE)
}
