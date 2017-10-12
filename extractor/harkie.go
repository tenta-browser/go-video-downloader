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
 * harkie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/hark.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type HarkIE struct {
	*rnt.CommonIE
}

func NewHarkIE() rnt.InfoExtractor {
	ret := &HarkIE{}
	ret.CommonIE = rnt.NewCommonIE()
	ret.VALIDURL = `https?://(?:www\.)?hark\.com/clips/(?P<id>.+?)-.+`
	return ret
}

func (self *HarkIE) Key() string {
	return "Hark"
}

func (self *HarkIE) Name() string {
	return `Hark extractor`
}

func (self *HarkIE) _real_extract(url string) map[string]interface{} {
	video_id := (self).MatchID(url)
	data := (self).DownloadJSON(rnt.StrFormat(`http://www.hark.com/clips/%s.json`, video_id), video_id, rnt.AsOptString(`Downloading JSON metadata`), rnt.AsOptString(`Unable to download JSON metadata`), nil, true, rnt.OptString{}, rnt.OptString{}, map[string]interface{}{}, map[string]interface{}{})
	return map[string]interface{}{`id`: video_id,
		`url`:         (data)[`url`],
		`title`:       (data)[`name`],
		`description`: rnt.DictGet(data, `description`, nil),
		`thumbnail`:   rnt.DictGet(data, `image_original`, nil),
		`duration`:    rnt.DictGet(data, `duration`, nil)}
}

func (self *HarkIE) Extract(url string) (*rnt.VideoResult, error) {
	return rnt.RunExtractor(url, self._real_extract)
}

func init() {
	registerFactory(`Hark`, NewHarkIE)
}
