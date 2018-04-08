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
 * drtuberie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/drtuber.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type DrTuberIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewDrTuberIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &DrTuberIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "DrTuber"
	_VALID_URL = "https?://(?:(?:www|m)\\.)?drtuber\\.com/(?:video|embed)/(?P<id>\\d+)(?:/(?P<display_id>[\\w-]+))?"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *DrTuberIE) Key() string {
	return "DrTuber"
}

func (self *DrTuberIE) Name() string {
	return self.IE_NAME
}

func (self *DrTuberIE) _real_extract(url string) rnt.SDict {
	var (
		categories    []interface{}
		cats_str      rnt.OptString
		comment_count rnt.OptInt
		dislike_count rnt.OptInt
		display_id    rnt.OptString
		extract_count func(string, string, interface{}) rnt.OptInt
		format_id     string
		formats       []interface{}
		like_count    rnt.OptInt
		mobj          rnt.Match
		thumbnail     rnt.OptString
		title         rnt.OptString
		video_data    interface{}
		video_id      rnt.OptString
		video_url     interface{}
		webpage       string
		τmp1          τ_Tsαω
	)
	mobj = rnt.ReMatch((self).VALIDURL, url, 0)
	video_id = rnt.ReMatchGroupOne(mobj, "id")
	display_id = func() rnt.OptString {
		if v := (rnt.ReMatchGroupOne(mobj, "display_id")); τ_isTruthy_Os(v) {
			return v
		} else {
			return video_id
		}
	}()
	webpage = (self).DownloadWebpageURL(rnt.StrFormat2("http://www.drtuber.com/video/%s", video_id), display_id.Get(), rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	video_data = (self).DownloadJSON("http://www.drtuber.com/player_config_json/", video_id.Get(), rnt.AsOptString("Downloading JSON metadata"), rnt.AsOptString("Unable to download JSON metadata"), nil, true, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{
		"vid":       video_id,
		"embed":     0,
		"aid":       0,
		"domain_id": 0,
	})
	formats = []interface{}{}
	for _, τel := range rnt.DictItems(τ_cast_α_to_d(rnt.UnsafeSubscript(video_data, "files"))) {
		τmp1 = τel
		format_id = (τmp1).Φ0
		video_url = (τmp1).Φ1
		if rnt.IsTruthy(video_url) {
			formats = append(formats, rnt.SDict{
				"format_id": format_id,
				"quality": func() int {
					if (format_id) == ("hq") {
						return 2
					} else {
						return 1
					}
				}(),
				"url": video_url,
			})
		}
	}
	formats = τ_conv_Ld_to_Lα((self).SortFormats(τ_conv_Lα_to_Ld(formats)))
	title = (self).HTMLSearchRegexMulti(τ_conv_Tsssssω_to_Ls(τ_Tsssssω{
		Φ0: "<h1[^>]+class=[\"\\']title[^>]+>([^<]+)",
		Φ1: "<title>([^<]+)\\s*@\\s+DrTuber",
		Φ2: "class=\"title_watch\"[^>]*><(?:p|h\\d+)[^>]*>([^<]+)<",
		Φ3: "<p[^>]+class=\"title_substrate\">([^<]+)</p>",
		Φ4: "<title>([^<]+) - \\d+",
	}), webpage, "title", rnt.NoDefault, true, 0, nil)
	thumbnail = (self).HTMLSearchRegexOne("poster=\"([^\"]+)\"", webpage, "thumbnail", rnt.NoDefault, false, 0, nil)
	extract_count = func(id_ string, name string, κdefault interface{}) rnt.OptInt {
		return rnt.StrToInt((self).HTMLSearchRegexOne(rnt.StrFormat2("<span[^>]+(?:class|id)=\"%s\"[^>]*>([\\d,\\.]+)</span>", id_), webpage, rnt.StrFormat2("%s count", name), κdefault, false, 0, nil))
	}
	like_count = extract_count("rate_likes", "like", rnt.NoDefault)
	dislike_count = extract_count("rate_dislikes", "dislike", nil)
	comment_count = extract_count("comments_count", "comment", rnt.NoDefault)
	cats_str = (self).SearchRegexOne("<div[^>]+class=\"categories_list\">(.+?)</div>", webpage, "categories", rnt.NoDefault, false, 0, nil)
	categories = func() []interface{} {
		if !(τ_isTruthy_Os(cats_str)) {
			return []interface{}{}
		} else {
			return τ_conv_Ls_to_Lα(rnt.ReFindAllOne("<a title=\"([^\"]+)\"", cats_str.Get(), 0))
		}
	}()
	return rnt.SDict{
		"id":            video_id,
		"display_id":    display_id,
		"formats":       formats,
		"title":         title,
		"thumbnail":     thumbnail,
		"like_count":    like_count,
		"dislike_count": dislike_count,
		"comment_count": comment_count,
		"categories":    categories,
		"age_limit":     (self).RTASearch(webpage),
	}
}

func (self *DrTuberIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("DrTuber", NewDrTuberIE)
}
