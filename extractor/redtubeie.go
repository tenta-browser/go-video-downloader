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
 * redtubeie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/redtube.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type RedTubeIE struct {
	*rnt.CommonIE
}

func NewRedTubeIE() rnt.InfoExtractor {
	ret := &RedTubeIE{}
	ret.CommonIE = rnt.NewCommonIE()
	ret.VALIDURL = "https?://(?:(?:www\\.)?redtube\\.com/|embed\\.redtube\\.com/\\?.*?\\bid=)(?P<id>[0-9]+)"
	return ret
}

func (self *RedTubeIE) Key() string {
	return "RedTube"
}

func (self *RedTubeIE) Name() string {
	return "RedTube extractor"
}

func (self *RedTubeIE) _extract_urls(webpage string) []string {
	return rnt.ReFindAllOne("<iframe[^>]+?src=[\"\\'](?P<url>(?:https?:)?//embed\\.redtube\\.com/\\?.*?\\bid=\\d+)", webpage, 0)
}

func (self *RedTubeIE) _real_extract(url string) map[string]interface{} {
	video_id := (self).MatchID(url)
	webpage := (self).DownloadWebpageURL(rnt.StrFormat("http://www.redtube.com/%s", video_id), video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, rnt.OptString{}, map[string]interface{}{}, map[string]interface{}{})
	if func() bool {
		for _, τmp2 := range func() []bool {
			τmp3 := []bool{}
			for _, τmp2 := range []string{"video-deleted-info", ">This video has been removed"} {
				s := τmp2

				τmp3 = append(τmp3, rnt.StrContains(webpage, s))
			}
			return τmp3
		}() {
			τmp1 := τmp2
			if τmp1 {
				return true
			}
		}
		return false
	}() {
		panic(rnt.PyExtractorError(rnt.StrFormat("Video %s has been removed", video_id), true, rnt.OptString{}))
	}
	title := (self).HTMLSearchRegexMulti(func() []string {
		τmp1 := struct {
			_0 string
			_1 string
		}{_0: "<h1 class=\"videoTitle[^\"]*\">(?P<title>.+?)</h1>",
			_1: "videoTitle\\s*:\\s*([\"\\'])(?P<title>)\\1"}
		return []string{τmp1._0, τmp1._1}
	}(), webpage, "title", rnt.NoDefault, true, 0, "title")
	formats := []interface{}{}
	sources := (self).ParseJSON((self).SearchRegexOne("sources\\s*:\\s*({.+?})", webpage, "source", "{}", true, 0, nil).Get(), video_id, nil, false)
	if len(sources) > 0 && true {
		for τmp1, τmp2 := range sources {
			τmp3 := struct {
				_0 string
				_1 interface{}
			}{_0: τmp1,
				_1: τmp2}
			format_id := (τmp3)._0
			format_url := (τmp3)._1
			if rnt.IsTruthy(format_url) {
				formats = append(formats, map[string]interface{}{"url": format_url,
					"format_id": format_id,
					"height":    rnt.IntOrNone(format_id, 1, rnt.OptInt{}, 1)})
			}
		}
	}
	medias := (self).ParseJSONList((self).SearchRegexOne("mediaDefinition\\s*:\\s*(\\[.+?\\])", webpage, "media definitions", "{}", true, 0, nil).Get(), video_id, nil, false)
	if len(medias) > 0 && true {
		for _, τmp1 := range medias {
			media := τmp1
			format_url := rnt.DictGet(func(val interface{}) map[string]interface{} {
				res, ok := val.(map[string]interface{})
				if !ok {
					panic(rnt.NewCastError(val, `map[string]interface{}`))
				}
				return res
			}(media), "videoUrl", nil)
			if !(rnt.IsTruthy(format_url)) || !(func() bool {
				_, ok := (format_url).(string)
				return ok
			}()) {
				continue
			}
			format_id := rnt.DictGet(func(val interface{}) map[string]interface{} {
				res, ok := val.(map[string]interface{})
				if !ok {
					panic(rnt.NewCastError(val, `map[string]interface{}`))
				}
				return res
			}(media), "quality", nil)
			formats = append(formats, map[string]interface{}{"url": format_url,
				"format_id": format_id,
				"height":    rnt.IntOrNone(format_id, 1, rnt.OptInt{}, 1)})
		}
	}
	if !(len(formats) > 0) {
		video_url := (self).HTMLSearchRegexOne("<source src=\"(.+?)\" type=\"video/mp4\">", webpage, "video URL", rnt.NoDefault, true, 0, nil)
		formats = append(formats, map[string]interface{}{"url": video_url})
	}
	(self).SortFormats(formats)
	thumbnail := (self).OgSearchThumbnail(webpage, rnt.NoDefault)
	upload_date := rnt.UnifiedStrDate((self).SearchRegexOne("<span[^>]+class=\"added-time\"[^>]*>ADDED ([^<]+)<", webpage, "upload date", rnt.NoDefault, false, 0, nil), true)
	duration := rnt.IntOrNone((self).SearchRegexOne("videoDuration\\s*:\\s*(\\d+)", webpage, "duration", nil, true, 0, nil), 1, rnt.OptInt{}, 1)
	view_count := rnt.StrToInt((self).SearchRegexOne("<span[^>]*>VIEWS</span></td>\\s*<td>([\\d,.]+)", webpage, "view count", rnt.NoDefault, false, 0, nil))
	age_limit := 18
	return map[string]interface{}{"id": video_id,
		"ext":         "mp4",
		"title":       title,
		"thumbnail":   thumbnail,
		"upload_date": upload_date,
		"duration":    duration,
		"view_count":  view_count,
		"age_limit":   age_limit,
		"formats":     formats}
}

func (self *RedTubeIE) Extract(url string) (*rnt.VideoResult, error) {
	return rnt.RunExtractor(url, self._real_extract)
}

func init() {
	registerFactory("RedTube", NewRedTubeIE)
}
