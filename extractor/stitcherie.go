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
 * stitcherie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/stitcher.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type StitcherIE struct {
	*rnt.CommonIE
}

func NewStitcherIE() rnt.InfoExtractor {
	ret := &StitcherIE{}
	ret.CommonIE = rnt.NewCommonIE()
	ret.VALIDURL = "https?://(?:www\\.)?stitcher\\.com/podcast/(?:[^/]+/)+e/(?:(?P<display_id>[^/#?&]+?)-)?(?P<id>\\d+)(?:[/#?&]|$)"
	return ret
}

func (self *StitcherIE) Key() string {
	return "Stitcher"
}

func (self *StitcherIE) Name() string {
	return "Stitcher extractor"
}

func (self *StitcherIE) _real_extract(url string) map[string]interface{} {
	mobj := rnt.ReMatch((self).VALIDURL, url, 0)
	audio_id := rnt.ReMatchGroupOne(mobj, "id")
	display_id := func() rnt.OptString {
		if (rnt.ReMatchGroupOne(mobj, "display_id")).IsSet() && (rnt.ReMatchGroupOne(mobj, "display_id").Get()) != "" {
			return rnt.ReMatchGroupOne(mobj, "display_id")
		} else {
			return audio_id
		}
	}()
	webpage := (self).DownloadWebpageURL(url, display_id.Get(), rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, rnt.OptString{}, map[string]interface{}{}, map[string]interface{}{})
	episode := rnt.UnsafeSubscript(((self).ParseJSON(rnt.JsToJSON((self).SearchRegexOne("(?s)var\\s+stitcher(?:Config)?\\s*=\\s*({.+?});\\n", webpage, "episode config", rnt.NoDefault, true, 0, nil).Get()), display_id.Get(), nil, true))["config"], "episode")
	title := rnt.UnescapeHTML(rnt.CastToOptString(rnt.UnsafeSubscript(episode, "title")))
	formats := func() []map[string]interface{} {
		τmp2 := []map[string]interface{}{}
		for _, τmp1 := range func() []string {
			τmp1 := struct {
				_0 string
			}{_0: "episodeURL"}
			return []string{τmp1._0}
		}() {
			episode_key := τmp1
			if !(rnt.IsTruthy(rnt.DictGet(func(val interface{}) map[string]interface{} {
				res, ok := val.(map[string]interface{})
				if !ok {
					panic(rnt.NewCastError(val, `map[string]interface{}`))
				}
				return res
			}(episode), episode_key, nil))) {
				continue
			}
			τmp2 = append(τmp2, map[string]interface{}{"url": rnt.UnsafeSubscript(episode, episode_key),
				"ext": func() string {
					if (rnt.DetermineExt(rnt.CastToOptString(rnt.UnsafeSubscript(episode, episode_key)), "unknown_video")) != "" {
						return rnt.DetermineExt(rnt.CastToOptString(rnt.UnsafeSubscript(episode, episode_key)), "unknown_video")
					} else {
						return "mp3"
					}
				}(),
				"vcodec": "none"})
		}
		return τmp2
	}()
	description := (self).SearchRegexOne("Episode Info:\\s*</span>([^<]+)<", webpage, "description", rnt.NoDefault, false, 0, nil)
	duration := rnt.IntOrNone(rnt.DictGet(func(val interface{}) map[string]interface{} {
		res, ok := val.(map[string]interface{})
		if !ok {
			panic(rnt.NewCastError(val, `map[string]interface{}`))
		}
		return res
	}(episode), "duration", nil), 1, rnt.OptInt{}, 1)
	thumbnail := rnt.DictGet(func(val interface{}) map[string]interface{} {
		res, ok := val.(map[string]interface{})
		if !ok {
			panic(rnt.NewCastError(val, `map[string]interface{}`))
		}
		return res
	}(episode), "episodeImage", nil)
	return map[string]interface{}{"id": audio_id,
		"display_id":  display_id,
		"title":       title,
		"description": description,
		"duration":    duration,
		"thumbnail":   thumbnail,
		"formats":     formats}
}

func (self *StitcherIE) Extract(url string) (*rnt.VideoResult, error) {
	return rnt.RunExtractor(url, self._real_extract)
}

func init() {
	registerFactory("Stitcher", NewStitcherIE)
}
