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
 * fivetvie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/fivetv.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type FiveTVIE struct {
	*rnt.CommonIE
}

func NewFiveTVIE() rnt.InfoExtractor {
	ret := &FiveTVIE{}
	ret.CommonIE = rnt.NewCommonIE()
	ret.VALIDURL = `(?x)
                    http://
                        (?:www\.)?5-tv\.ru/
                        (?:
                            (?:[^/]+/)+(?P<id>\d+)|
                            (?P<path>[^/?#]+)(?:[/?#])?
                        )
                    `
	return ret
}

func (self *FiveTVIE) Key() string {
	return "FiveTV"
}

func (self *FiveTVIE) Name() string {
	return `FiveTV extractor`
}

func (self *FiveTVIE) _real_extract(url string) map[string]interface{} {
	mobj := rnt.ReMatch((self).VALIDURL, url, 0)
	video_id := func() rnt.OptString {
		if (rnt.ReMatchGroupOne(mobj, `id`)).IsSet() && (rnt.ReMatchGroupOne(mobj, `id`).Get()) != "" {
			return rnt.ReMatchGroupOne(mobj, `id`)
		} else {
			return rnt.ReMatchGroupOne(mobj, `path`)
		}
	}()
	webpage := (self).DownloadWebpage(url, video_id.Get(), rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, rnt.OptString{}, map[string]interface{}{}, map[string]interface{}{})
	video_url := (self).SearchRegex([]string{`<div[^>]+?class="flowplayer[^>]+?data-href="([^"]+)"`, `<a[^>]+?href="([^"]+)"[^>]+?class="videoplayer"`}, webpage, `video url`, rnt.NoDefault, true, 0, nil)
	title := func() rnt.OptString {
		if ((self).OgSearchTitle(webpage, nil, true)).IsSet() && ((self).OgSearchTitle(webpage, nil, true).Get()) != "" {
			return (self).OgSearchTitle(webpage, nil, true)
		} else {
			return (self).SearchRegex(`<title>([^<]+)</title>`, webpage, `title`, rnt.NoDefault, true, 0, nil)
		}
	}()
	duration := rnt.IntOrNone((self).OgSearchProperty(`video:duration`, webpage, rnt.AsOptString(`duration`), nil, true), 1, rnt.OptInt{}, 1)
	return map[string]interface{}{`id`: video_id,
		`url`:         video_url,
		`title`:       title,
		`description`: (self).OgSearchDescription(webpage, nil),
		`thumbnail`:   (self).OgSearchThumbnail(webpage, nil),
		`duration`:    duration}
}

func (self *FiveTVIE) Extract(url string) (*rnt.VideoResult, error) {
	return rnt.RunExtractor(url, self._real_extract)
}

func init() {
	registerFactory(`FiveTV`, NewFiveTVIE)
}
