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
	*rnt.CommonIE
}

func NewBangIE() rnt.InfoExtractor {
	ret := &BangIE{}
	ret.CommonIE = rnt.NewCommonIE()
	ret.VALIDURL = "https?://(?:www\\.)?bang\\.com/video/(?P<id>[0-9a-zA-Z_-]+)"
	return ret
}

func (self *BangIE) Key() string {
	return "Bang"
}

func (self *BangIE) Name() string {
	return "Bang extractor"
}

func (self *BangIE) _real_extract(url string) map[string]interface{} {
	video_id := (self).MatchID(url)
	webpage := (self).DownloadWebpageURL(url, video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, rnt.OptString{}, map[string]interface{}{}, map[string]interface{}{})
	api_key := (self).SearchRegexOne("apiKey\\s*:\\s*[\"\\']([^\"\\']+)[\"\\']", webpage, "api_key", nil, true, 0, nil)
	if !((api_key).IsSet() && (api_key.Get()) != "") {
		panic(rnt.PyExtractorError("LOGIN_REQUIRED", true, rnt.OptString{}))
	}
	mainjs_url := (self).SearchRegexOne("<script[^>]+?src=[\"\\'](?P<mainjs_url>/assets/js/main.[a-z0-9]+.js)", webpage, "main.js url", rnt.NoDefault, true, 0, nil)
	mainjs_url = rnt.AsOptString(rnt.URLJoin(url, mainjs_url))
	mainjs := (self).DownloadWebpageURL(mainjs_url.Get(), video_id, rnt.AsOptString("Downloading main.js"), rnt.OptString{}, true, 1, 5, rnt.OptString{}, rnt.OptString{}, map[string]interface{}{}, map[string]interface{}{})
	elasticsearch_url := (self).SearchRegexOne("\"elasticsearch.url\"\\s*:\\s*\"([^\"]+)\"", mainjs, "Elasticsearch url", rnt.NoDefault, true, 0, nil)
	elasticsearch_auth := (self).SearchRegexOne("\"elasticsearch.auth\"\\s*:\\s*\"([^\"]+)\"", mainjs, "Elasticsearch auth", rnt.NoDefault, true, 0, nil)
	video_id_hex := rnt.BytesToStr(rnt.HexEncode(rnt.Base64UrlDecodeString(video_id)), "utf-8")
	metadata_url := rnt.URLJoin(elasticsearch_url.Get(), rnt.AsOptString(rnt.StrFormat("/videos/video/%s", video_id_hex)))
	metadata := (self).DownloadJSON(metadata_url, video_id, rnt.AsOptString("Downloading JSON metadata"), rnt.AsOptString("Unable to download JSON metadata"), nil, true, rnt.OptString{}, rnt.OptString{}, map[string]interface{}{"Authorization": ("Basic " + elasticsearch_auth.Get())}, map[string]interface{}{})
	metadata_source := (metadata)["_source"]
	identifier := rnt.UnsafeSubscript(metadata_source, "identifier")
	title := rnt.UnsafeSubscript(metadata_source, "name")
	description := rnt.UnsafeSubscript(metadata_source, "description")
	links_url := rnt.StrFormat("https://links.bang.com/video/%s", identifier)
	links := (self).DownloadJSONList(links_url, video_id, rnt.AsOptString("Downloading links"), rnt.AsOptString("Unable to download JSON metadata"), nil, true, rnt.OptString{}, rnt.OptString{}, map[string]interface{}{}, map[string]interface{}{})
	formats := []interface{}{}
	for _, τmp1 := range links {
		link := τmp1
		formats = append(formats, map[string]interface{}{"format_id": rnt.UnsafeSubscript(link, "FormatName"),
			"filesize": rnt.IntOrNone(rnt.UnsafeSubscript(link, "Size"), 1, rnt.OptInt{}, 1),
			"url":      ((rnt.URLJoin(links_url, rnt.CastToOptString(rnt.UnsafeSubscript(link, "Url"))) + "?apiKey=") + api_key.Get())})
	}
	(self).SortFormats(formats)
	age_limit := (self).RTASearch(webpage)
	return map[string]interface{}{"id": video_id,
		"title":       title,
		"description": description,
		"formats":     formats,
		"age_limit":   age_limit}
}

func (self *BangIE) Extract(url string) (*rnt.VideoResult, error) {
	return rnt.RunExtractor(url, self._real_extract)
}

func init() {
	registerFactory("Bang", NewBangIE)
}
