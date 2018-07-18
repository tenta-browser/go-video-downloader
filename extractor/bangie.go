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
	IE_NAME    string
	_LOGIN_URL string
}

func NewBangIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_LOGIN_URL string
		_VALID_URL string
	)
	self := &BangIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "Bang"
	_VALID_URL = "https?://(?:www\\.)?bang\\.com/video/(?P<id>[0-9a-zA-Z_-]+)"
	_LOGIN_URL = "https://www.bang.com/login_check"
	self.IE_NAME = IE_NAME
	self._LOGIN_URL = _LOGIN_URL
	self.VALIDURL = _VALID_URL
	return self
}

func (self *BangIE) Key() string {
	return "Bang"
}

func (self *BangIE) Name() string {
	return self.IE_NAME
}

func (self *BangIE) _login(video_id string) interface{} {
	var (
		login_data []byte
		password   rnt.OptString
		res        interface{}
		username   rnt.OptString
		τmp1       τ_TOsOsω
	)
	τmp1 = (self).GetLoginInfo()
	username = (τmp1).Φ0
	password = (τmp1).Φ1
	if !(τ_isTruthy_Os(username)) {
		panic(rnt.PyExtractorError("LOGIN_REQUIRED", true, rnt.OptString{}))
	}
	login_data = rnt.StrToBytes(rnt.JSONDumps(rnt.SDict{
		"_username":    username,
		"_password":    password,
		"_remember_me": false,
	}), "utf-8")
	res = (self).DownloadJSON((self)._LOGIN_URL, video_id, rnt.AsOptString("Logging in"), rnt.AsOptString("Unable to download JSON metadata"), nil, true, rnt.OptString{}, login_data, rnt.SDict{}, rnt.SDict{}, nil)
	if !(rnt.IsTruthy(rnt.DictGet(τ_cast_α_to_d(res), "success", nil))) {
		panic(rnt.PyExtractorError("Invalid credentials", true, rnt.OptString{}))
	}
	return nil
}

func (self *BangIE) _real_extract(url string) rnt.SDict {
	var (
		age_limit          int
		api_key            rnt.OptString
		description        interface{}
		elasticsearch_auth rnt.OptString
		elasticsearch_url  rnt.OptString
		formats            []interface{}
		get_api_key        func() rnt.OptString
		identifier         interface{}
		link               interface{}
		links              interface{}
		links_url          string
		mainjs             string
		mainjs_url         rnt.OptString
		metadata           interface{}
		metadata_source    interface{}
		metadata_url       string
		title              interface{}
		video_id           string
		video_id_hex       string
		webpage            string
	)
	video_id = (self).MatchID(url)
	webpage = (self).DownloadWebpageURL(url, video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{}, nil)
	get_api_key = func() rnt.OptString {
		return (self).SearchRegexOne("apiKey\\s*:\\s*[\"\\']([^\"\\']+)[\"\\']", webpage, "api_key", nil, true, 0, nil)
	}
	api_key = get_api_key()
	if !(τ_isTruthy_Os(api_key)) {
		(self)._login(video_id)
		webpage = (self).DownloadWebpageURL(url, video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{}, nil)
		api_key = get_api_key()
	}
	mainjs_url = (self).SearchRegexOne("<script[^>]+?src=[\"\\'](?P<mainjs_url>/assets/js/main.[a-z0-9]+.js)", webpage, "main.js url", rnt.NoDefault, true, 0, nil)
	mainjs_url = rnt.AsOptString(rnt.URLJoin(url, mainjs_url))
	mainjs = (self).DownloadWebpageURL(mainjs_url.Get(), video_id, rnt.AsOptString("Downloading main.js"), rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{}, nil)
	elasticsearch_url = (self).SearchRegexOne("\"elasticsearch.url\"\\s*:\\s*\"([^\"]+)\"", mainjs, "Elasticsearch url", rnt.NoDefault, true, 0, nil)
	elasticsearch_auth = (self).SearchRegexOne("\"elasticsearch.auth\"\\s*:\\s*\"([^\"]+)\"", mainjs, "Elasticsearch auth", rnt.NoDefault, true, 0, nil)
	video_id_hex = rnt.BytesToStr(rnt.HexEncode(rnt.Base64UrlDecodeString(video_id)), "utf-8")
	metadata_url = rnt.URLJoin(elasticsearch_url.Get(), rnt.AsOptString(rnt.StrFormat2("/videos/video/%s", video_id_hex)))
	metadata = (self).DownloadJSON(metadata_url, video_id, rnt.AsOptString("Downloading JSON metadata"), rnt.AsOptString("Unable to download JSON metadata"), nil, true, rnt.OptString{}, nil, rnt.SDict{
		"Authorization": ("Basic " + elasticsearch_auth.Get()),
	}, rnt.SDict{}, nil)
	metadata_source = rnt.UnsafeSubscript(metadata, "_source")
	identifier = rnt.UnsafeSubscript(metadata_source, "identifier")
	title = rnt.UnsafeSubscript(metadata_source, "name")
	description = rnt.UnsafeSubscript(metadata_source, "description")
	links_url = rnt.StrFormat2("https://links.bang.com/video/%s", identifier)
	links = (self).DownloadJSON(links_url, video_id, rnt.AsOptString("Downloading links"), rnt.AsOptString("Unable to download JSON metadata"), nil, true, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{}, nil)
	formats = []interface{}{}
	for _, τel := range τ_cast_α_to_Lα(links) {
		link = τel
		formats = append(formats, rnt.SDict{
			"format_id": rnt.UnsafeSubscript(link, "FormatName"),
			"filesize":  rnt.IntOrNone(rnt.UnsafeSubscript(link, "Size"), 1, rnt.OptInt{}, 1),
			"url":       ((rnt.URLJoin(links_url, rnt.CastToOptString(rnt.UnsafeSubscript(link, "Url"))) + "?apiKey=") + api_key.Get()),
		})
	}
	formats = τ_conv_Ld_to_Lα((self).SortFormats(τ_conv_Lα_to_Ld(formats)))
	age_limit = (self).RTASearch(webpage)
	return rnt.SDict{
		"id":          video_id,
		"title":       title,
		"description": description,
		"formats":     formats,
		"age_limit":   age_limit,
	}
}

func (self *BangIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("Bang", NewBangIE)
}
