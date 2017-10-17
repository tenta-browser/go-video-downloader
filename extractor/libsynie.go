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
 * libsynie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/libsyn.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type LibsynIE struct {
	*rnt.CommonIE
}

func NewLibsynIE() rnt.InfoExtractor {
	ret := &LibsynIE{}
	ret.CommonIE = rnt.NewCommonIE()
	ret.VALIDURL = "(?P<mainurl>https?://html5-player\\.libsyn\\.com/embed/episode/id/(?P<id>[0-9]+))"
	return ret
}

func (self *LibsynIE) Key() string {
	return "Libsyn"
}

func (self *LibsynIE) Name() string {
	return "Libsyn extractor"
}

func (self *LibsynIE) _real_extract(url string) map[string]interface{} {
	m := rnt.ReMatch((self).VALIDURL, url, 0)
	video_id := rnt.ReMatchGroupOne(m, "id")
	url = rnt.ReMatchGroupOne(m, "mainurl").Get()
	webpage := (self).DownloadWebpageURL(url, video_id.Get(), rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, rnt.OptString{}, map[string]interface{}{}, map[string]interface{}{})
	formats := func() []map[string]interface{} {
		τmp2 := []map[string]interface{}{}
		for _, τmp1 := range func() []string {
			τmp2 := []string{}
			for τmp1 := range func() map[string]struct{} {
				τmp1 := map[string]struct{}{}
				for _, τmp2 := range rnt.ReFindAllOne("var\\s+mediaURL(?:Libsyn)?\\s*=\\s*\"([^\"]+)\"", webpage, 0) {
					τmp1[τmp2] = struct{}{}
				}
				return τmp1
			}() {
				τmp2 = append(τmp2, τmp1)
			}
			return τmp2
		}() {
			media_url := τmp1

			τmp2 = append(τmp2, map[string]interface{}{"url": media_url})
		}
		return τmp2
	}()
	podcast_title := (self).SearchRegexOne("<h2>([^<]+)</h2>", webpage, "podcast title", nil, true, 0, nil)
	episode_title := (self).SearchRegexOne("(?:<div class=\"episode-title\">|<h3>)([^<]+)</", webpage, "episode title", rnt.NoDefault, true, 0, nil)
	title := func() string {
		if (podcast_title).IsSet() && (podcast_title.Get()) != "" {
			return rnt.StrFormat("%s - %s", podcast_title, episode_title)
		} else {
			return episode_title.Get()
		}
	}()
	description := (self).HTMLSearchRegexOne("<div id=\"info_text_body\">(.+?)</div>", webpage, "description", nil, true, 0, nil)
	thumbnail := (self).SearchRegexOne("<img[^>]+class=\"info-show-icon\"[^>]+src=\"([^\"]+)\"", webpage, "thumbnail", rnt.NoDefault, false, 0, nil)
	release_date := rnt.UnifiedStrDate((self).SearchRegexOne("<div class=\"release_date\">Released: ([^<]+)<", webpage, "release date", rnt.NoDefault, false, 0, nil), true)
	return map[string]interface{}{"id": video_id,
		"title":       title,
		"description": description,
		"thumbnail":   thumbnail,
		"upload_date": release_date,
		"formats":     formats}
}

func (self *LibsynIE) Extract(url string) (*rnt.VideoResult, error) {
	return rnt.RunExtractor(url, self._real_extract)
}

func init() {
	registerFactory("Libsyn", NewLibsynIE)
}
