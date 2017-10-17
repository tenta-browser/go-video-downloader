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
 * wdrmobileie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/wdr.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type WDRMobileIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewWDRMobileIE() rnt.InfoExtractor {
	ret := &WDRMobileIE{}
	ret.CommonIE = rnt.NewCommonIE()
	ret.VALIDURL = "(?x)\n        https?://mobile-ondemand\\.wdr\\.de/\n        .*?/fsk(?P<age_limit>[0-9]+)\n        /[0-9]+/[0-9]+/\n        (?P<id>[0-9]+)_(?P<title>[0-9]+)"
	ret.IE_NAME = "wdr:mobile"
	return ret
}

func (self *WDRMobileIE) Key() string {
	return "WDRMobile"
}

func (self *WDRMobileIE) Name() string {
	return self.IE_NAME
}

func (self *WDRMobileIE) _real_extract(url string) map[string]interface{} {
	mobj := rnt.ReMatch((self).VALIDURL, url, 0)
	return map[string]interface{}{"id": rnt.ReMatchGroupOne(mobj, "id"),
		"title":        rnt.ReMatchGroupOne(mobj, "title"),
		"age_limit":    rnt.ConvertToInt(rnt.ReMatchGroupOne(mobj, "age_limit")),
		"url":          url,
		"http_headers": map[string]interface{}{"User-Agent": "mobile"}}
}

func (self *WDRMobileIE) Extract(url string) (*rnt.VideoResult, error) {
	return rnt.RunExtractor(url, self._real_extract)
}

func init() {
	registerFactory("WDRMobile", NewWDRMobileIE)
}
