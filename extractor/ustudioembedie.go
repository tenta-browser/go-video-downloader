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
 * ustudioembedie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/ustudio.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type UstudioEmbedIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewUstudioEmbedIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &UstudioEmbedIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "UstudioEmbed"
	IE_NAME = "ustudio:embed"
	_VALID_URL = "https?://(?:(?:app|embed)\\.)?ustudio\\.com/embed/(?P<uid>[^/]+)/(?P<id>[^/]+)"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *UstudioEmbedIE) Key() string {
	return "UstudioEmbed"
}

func (self *UstudioEmbedIE) Name() string {
	return self.IE_NAME
}

func (self *UstudioEmbedIE) _real_extract(url string) rnt.SDict {
	var (
		ext         string
		formats     []interface{}
		height      rnt.OptInt
		image       interface{}
		image_url   interface{}
		qualities   interface{}
		quality     interface{}
		quality_url interface{}
		thumbnails  []interface{}
		title       interface{}
		uploader_id rnt.OptString
		video_data  interface{}
		video_id    rnt.OptString
		τmp1        []rnt.OptString
		τmp2        τ_Tsαω
	)
	τmp1 = rnt.ReMatchGroups(rnt.ReMatch((self).VALIDURL, url, 0), rnt.OptString{})
	uploader_id = (τmp1)[0]
	video_id = (τmp1)[1]
	video_data = rnt.UnsafeSubscript(rnt.UnsafeSubscript((self).DownloadJSON(rnt.StrFormat2("http://app.ustudio.com/embed/%s/%s/config.json", uploader_id, video_id), video_id.Get(), rnt.AsOptString("Downloading JSON metadata"), rnt.AsOptString("Unable to download JSON metadata"), nil, true, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{}), "videos"), 0)
	title = rnt.UnsafeSubscript(video_data, "name")
	formats = []interface{}{}
	for _, τel := range rnt.DictItems(τ_cast_α_to_d(rnt.DictGet(τ_cast_α_to_d(video_data), "transcodes", rnt.SDict{}))) {
		τmp2 = τel
		ext = (τmp2).Φ0
		qualities = (τmp2).Φ1
		for _, τel := range τ_cast_α_to_Lα(qualities) {
			quality = τel
			quality_url = rnt.DictGet(τ_cast_α_to_d(quality), "url", nil)
			if !(rnt.IsTruthy(quality_url)) {
				continue
			}
			height = rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(quality), "height", nil), 1, rnt.OptInt{}, 1)
			formats = append(formats, rnt.SDict{
				"format_id": func() string {
					if τ_isTruthy_Oi(height) {
						return rnt.StrFormat2("%s-%dp", ext, height)
					} else {
						return ext
					}
				}(),
				"url":    quality_url,
				"width":  rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(quality), "width", nil), 1, rnt.OptInt{}, 1),
				"height": height,
			})
		}
	}
	formats = τ_conv_Ld_to_Lα((self).SortFormats(τ_conv_Lα_to_Ld(formats)))
	thumbnails = []interface{}{}
	for _, τel := range τ_cast_α_to_Lα(rnt.DictGet(τ_cast_α_to_d(video_data), "images", []interface{}{})) {
		image = τel
		image_url = rnt.DictGet(τ_cast_α_to_d(image), "url", nil)
		if !(rnt.IsTruthy(image_url)) {
			continue
		}
		thumbnails = append(thumbnails, rnt.SDict{
			"url": image_url,
		})
	}
	return rnt.SDict{
		"id":          video_id,
		"title":       title,
		"description": rnt.DictGet(τ_cast_α_to_d(video_data), "description", nil),
		"duration":    rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(video_data), "duration", nil), 1, rnt.OptInt{}, 1),
		"uploader_id": uploader_id,
		"tags":        rnt.DictGet(τ_cast_α_to_d(video_data), "keywords", nil),
		"thumbnails":  thumbnails,
		"formats":     formats,
	}
}

func (self *UstudioEmbedIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("UstudioEmbed", NewUstudioEmbedIE)
}
