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
 * camdemyie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/camdemy.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type CamdemyIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewCamdemyIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &CamdemyIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "Camdemy"
	_VALID_URL = "https?://(?:www\\.)?camdemy\\.com/media/(?P<id>\\d+)"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *CamdemyIE) Key() string {
	return "Camdemy"
}

func (self *CamdemyIE) Name() string {
	return self.IE_NAME
}

func (self *CamdemyIE) _real_extract(url string) rnt.SDict {
	var (
		description   rnt.OptString
		file_list_doc rnt.XMLElement
		file_name     string
		oembed_obj    interface{}
		src_from      rnt.OptString
		thumb_url     interface{}
		title         interface{}
		upload_date   rnt.OptString
		video_folder  string
		video_id      string
		video_url     string
		view_count    rnt.OptInt
		webpage       string
	)
	video_id = (self).MatchID(url)
	webpage = (self).DownloadWebpageURL(url, video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	src_from = (self).HTMLSearchRegexOne("class=['\\\"]srcFrom['\\\"][^>]*>Sources?(?:\\s+from)?\\s*:\\s*<a[^>]+(?:href|title)=(['\\\"])(?P<url>(?:(?!\\1).)+)\\1", webpage, "external source", nil, true, 0, "url")
	if τ_isTruthy_Os(src_from) {
		return (self).URLResult(src_from.Get(), rnt.OptString{}, rnt.OptString{}, rnt.OptString{})
	}
	oembed_obj = (self).DownloadJSON(("http://www.camdemy.com/oembed/?format=json&url=" + url), video_id, rnt.AsOptString("Downloading JSON metadata"), rnt.AsOptString("Unable to download JSON metadata"), nil, true, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	title = rnt.UnsafeSubscript(oembed_obj, "title")
	thumb_url = rnt.UnsafeSubscript(oembed_obj, "thumbnail_url")
	video_folder = rnt.URLJoin(rnt.CastToString(thumb_url), rnt.AsOptString("video/"))
	file_list_doc = (self).DownloadXML(rnt.URLJoin(video_folder, rnt.AsOptString("fileList.xml")), video_id, rnt.AsOptString("Downloading filelist XML"), rnt.AsOptString("Unable to download XML"), nil, true, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	file_name = (rnt.XMLFind(file_list_doc, "./video/item/fileName")).Text
	video_url = rnt.URLJoin(video_folder, rnt.AsOptString(file_name))
	upload_date = rnt.UnifiedStrDate((self).SearchRegexOne(">published on ([^<]+)<", webpage, "upload date", nil, true, 0, nil), true)
	view_count = rnt.StrToInt((self).SearchRegexOne("role=[\"\\']viewCnt[\"\\'][^>]*>([\\d,.]+) views", webpage, "view count", nil, true, 0, nil))
	description = func() rnt.OptString {
		if τ_isTruthy_Os((self).HTMLSearchMetaOne("description", webpage, rnt.OptString{}, false, nil, 0)) {
			return (self).HTMLSearchMetaOne("description", webpage, rnt.OptString{}, false, nil, 0)
		} else {
			return rnt.CleanHTML(rnt.CastToOptString(rnt.DictGet(τ_cast_α_to_d(oembed_obj), "description", nil)))
		}
	}()
	return rnt.SDict{
		"id":          video_id,
		"url":         video_url,
		"title":       title,
		"thumbnail":   thumb_url,
		"description": description,
		"creator":     rnt.DictGet(τ_cast_α_to_d(oembed_obj), "author_name", nil),
		"duration":    rnt.ParseDuration(rnt.CastToOptString(rnt.DictGet(τ_cast_α_to_d(oembed_obj), "duration", nil))),
		"upload_date": upload_date,
		"view_count":  view_count,
	}
}

func (self *CamdemyIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("Camdemy", NewCamdemyIE)
}
