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
 * bangie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/bang.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type BangIE struct {
	*rnt.Context
	_VALID_URL string
}

func NewBangIE(ctx *rnt.Context) rnt.InfoExtractor {
	ret := &BangIE{}
	ret.Context = ctx
	ret._VALID_URL = `https?://(?:www\.)?bang\.com/video/(?P<id>[0-9a-zA-Z]+)`
	return ret
}

func (self *BangIE) Ctx() *rnt.Context {
	return self.Context
}

func (self *BangIE) Key() string {
	return "Bang"
}

func (self *BangIE) ValidUrl() string {
	return self._VALID_URL
}

func (self *BangIE) Name() string {
	return `Bang extractor`
}

func (self *BangIE) Tests() []map[string]interface{} {
	return []map[string]interface{}{}
}

func (self *BangIE) _real_extract(url string) map[string]interface{} {
	video_id := rnt.MatchID(self, url)
	webpage := rnt.DownloadWebpage(self, url, video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, rnt.OptString{}, map[string]interface{}{}, map[string]interface{}{})
	api_key := rnt.SearchRegex(self, `apiKey\s*:\s*["\']([^"\']+)["\']`, webpage, `api_key`, nil, true, 0, nil)
	if !((api_key).IsSet() && (api_key.Get()) != "") {
		panic(rnt.PyExtractorError(`LOGIN_REQUIRED`, true, rnt.OptString{}))
	}
	mainjs_url := rnt.SearchRegex(self, `<script[^>]+?src=["\'](?P<mainjs_url>/assets/js/main.[a-z0-9]+.js)`, webpage, `main.js url`, rnt.NoDefault, true, 0, nil)
	mainjs_url = rnt.AsOptString(rnt.URLJoin(rnt.UP, url, mainjs_url))
	mainjs := rnt.DownloadWebpage(self, mainjs_url.Get(), video_id, rnt.AsOptString(`Downloading main.js`), rnt.OptString{}, true, 1, 5, rnt.OptString{}, rnt.OptString{}, map[string]interface{}{}, map[string]interface{}{})
	elasticsearch_url := rnt.SearchRegex(self, `"elasticsearch.url"\s*:\s*"([^"]+)"`, mainjs, `Elasticsearch url`, rnt.NoDefault, true, 0, nil)
	elasticsearch_auth := rnt.SearchRegex(self, `"elasticsearch.auth"\s*:\s*"([^"]+)"`, mainjs, `Elasticsearch auth`, rnt.NoDefault, true, 0, nil)
	video_id_hex := rnt.BytesToStr(rnt.HexEncode(rnt.Codecs, rnt.Base64DecodeString(rnt.Codecs, video_id)), `utf-8`)
	metadata_url := rnt.URLJoin(rnt.UP, elasticsearch_url.Get(), rnt.AsOptString(rnt.StrFormat(`/videos/video/%s`, video_id_hex)))
	metadata := rnt.DownloadJSON(self, metadata_url, video_id, rnt.AsOptString(`Downloading JSON metadata`), rnt.AsOptString(`Unable to download JSON metadata`), 0, true, rnt.OptString{}, rnt.OptString{}, map[string]interface{}{`Authorization`: (`Basic ` + elasticsearch_auth.Get())}, map[string]interface{}{})
	metadata_source := (metadata)[`_source`]
	identifier := rnt.UnsafeSubscript(metadata_source, `identifier`)
	title := rnt.UnsafeSubscript(metadata_source, `name`)
	description := rnt.UnsafeSubscript(metadata_source, `description`)
	links_url := rnt.StrFormat(`https://links.bang.com/video/%s`, identifier)
	links := rnt.DownloadJSONList(self, links_url, video_id, rnt.AsOptString(`Downloading links`), rnt.AsOptString(`Unable to download JSON metadata`), 0, true, rnt.OptString{}, rnt.OptString{}, map[string]interface{}{}, map[string]interface{}{})
	formats := []interface{}{}
	for _, τmp1 := range links {
		link := τmp1
		formats = append(formats, map[string]interface{}{`format_id`: rnt.UnsafeSubscript(link, `FormatName`),
			`filesize`: rnt.IntOrNone(rnt.UnsafeSubscript(link, `Size`), 1, rnt.OptInt{}, 1),
			`url`:      ((rnt.URLJoin(rnt.UP, links_url, rnt.CastToOptString(rnt.UnsafeSubscript(link, `Url`))) + `?apiKey=`) + api_key.Get())})
	}
	rnt.SortFormats(self, formats)
	age_limit := rnt.RTASearch(self, webpage)
	return map[string]interface{}{`id`: video_id,
		`title`:       title,
		`description`: description,
		`formats`:     formats,
		`age_limit`:   age_limit}
}

func (self *BangIE) Extract(url string) (*rnt.VideoResult, error) {
	return rnt.RunExtractor(url, self._real_extract)
}

func init() {
	registerFactory(`Bang`, NewBangIE)
}
