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
 * mailruie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/mailru.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type MailRuIE struct {
	*rnt.CommonIE
	IE_DESC string
	IE_NAME string
}

func NewMailRuIE() rnt.InfoExtractor {
	var (
		IE_DESC    string
		IE_NAME    string
		_VALID_URL string
	)
	self := &MailRuIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "MailRu"
	IE_NAME = "mailru"
	IE_DESC = "Видео@Mail.Ru"
	_VALID_URL = "(?x)\n                    https?://\n                        (?:(?:www|m)\\.)?my\\.mail\\.ru/\n                        (?:\n                            video/.*\\#video=/?(?P<idv1>(?:[^/]+/){3}\\d+)|\n                            (?:(?P<idv2prefix>(?:[^/]+/){2})video/(?P<idv2suffix>[^/]+/\\d+))\\.html|\n                            (?:video/embed|\\+/video/meta)/(?P<metaid>\\d+)\n                        )\n                    "
	self.IE_DESC = IE_DESC
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *MailRuIE) Key() string {
	return "MailRu"
}

func (self *MailRuIE) Name() string {
	return self.IE_NAME
}

func (self *MailRuIE) _real_extract(url string) rnt.SDict {
	var (
		acc_id      interface{}
		author      interface{}
		content_id  rnt.OptString
		duration    rnt.OptInt
		f           interface{}
		format_id   interface{}
		formats     []interface{}
		height      rnt.OptInt
		item_id     interface{}
		meta_data   interface{}
		meta_id     rnt.OptString
		meta_url    interface{}
		mobj        rnt.Match
		page_config interface{}
		thumbnail   interface{}
		timestamp   rnt.OptInt
		title       rnt.OptString
		uploader    interface{}
		uploader_id interface{}
		video_data  interface{}
		video_id    rnt.OptString
		video_url   interface{}
		view_count  rnt.OptInt
		webpage     string
	)
	mobj = rnt.ReMatch((self).VALIDURL, url, 0)
	meta_id = rnt.ReMatchGroupOne(mobj, "metaid")
	video_id = rnt.OptString{}
	if τ_isTruthy_Os(meta_id) {
		meta_url = rnt.StrFormat2("https://my.mail.ru/+/video/meta/%s", meta_id)
	} else {
		video_id = rnt.ReMatchGroupOne(mobj, "idv1")
		if !(τ_isTruthy_Os(video_id)) {
			video_id = rnt.AsOptString((rnt.ReMatchGroupOne(mobj, "idv2prefix").Get() + rnt.ReMatchGroupOne(mobj, "idv2suffix").Get()))
		}
		webpage = (self).DownloadWebpageURL(url, video_id.Get(), rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{}, nil)
		page_config = (self).ParseJSON((self).SearchRegexOne("(?s)<script[^>]+class=\"sp-video__page-config\"[^>]*>(.+?)</script>", webpage, "page config", "{}", true, 0, nil).Get(), video_id.Get(), nil, false)
		if rnt.IsTruthy(page_config) {
			meta_url = func() interface{} {
				if v := (rnt.DictGet(τ_cast_α_to_d(page_config), "metaUrl", nil)); rnt.IsTruthy(v) {
					return v
				} else {
					return rnt.DictGet(τ_cast_α_to_d(rnt.DictGet(τ_cast_α_to_d(page_config), "video", rnt.SDict{})), "metaUrl", nil)
				}
			}()
		} else {
			meta_url = nil
		}
	}
	video_data = nil
	if rnt.IsTruthy(meta_url) {
		video_data = (self).DownloadJSON(rnt.CastToString(meta_url), func() rnt.OptString {
			if v := (video_id); τ_isTruthy_Os(v) {
				return v
			} else {
				return meta_id
			}
		}().Get(), rnt.AsOptString("Downloading video meta JSON"), rnt.AsOptString("Unable to download JSON metadata"), nil, !(τ_isTruthy_Os(video_id)), rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{}, nil)
	}
	if !(rnt.IsTruthy(video_data)) {
		video_data = (self).DownloadJSON(rnt.StrFormat2("http://api.video.mail.ru/videos/%s.json?new=1", video_id), video_id.Get(), rnt.AsOptString("Downloading video JSON"), rnt.AsOptString("Unable to download JSON metadata"), nil, true, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{}, nil)
	}
	formats = []interface{}{}
	for _, τel := range τ_cast_α_to_Lα(rnt.UnsafeSubscript(video_data, "videos")) {
		f = τel
		video_url = rnt.DictGet(τ_cast_α_to_d(f), "url", nil)
		if !(rnt.IsTruthy(video_url)) {
			continue
		}
		format_id = rnt.DictGet(τ_cast_α_to_d(f), "key", nil)
		height = func() rnt.OptInt {
			if rnt.IsTruthy(format_id) {
				return rnt.IntOrNone((self).SearchRegexOne("^(\\d+)[pP]$", rnt.CastToString(format_id), "height", nil, true, 0, nil), 1, rnt.OptInt{}, 1)
			} else {
				return rnt.OptInt{}
			}
		}()
		formats = append(formats, rnt.SDict{
			"url":       video_url,
			"format_id": format_id,
			"height":    height,
		})
	}
	formats = τ_conv_Ld_to_Lα((self).SortFormats(τ_conv_Lα_to_Ld(formats)))
	meta_data = rnt.UnsafeSubscript(video_data, "meta")
	title = rnt.RemoveEnd(rnt.CastToOptString(rnt.UnsafeSubscript(meta_data, "title")), ".mp4")
	author = rnt.DictGet(τ_cast_α_to_d(video_data), "author", nil)
	uploader = rnt.DictGet(τ_cast_α_to_d(author), "name", nil)
	uploader_id = func() interface{} {
		if v := (rnt.DictGet(τ_cast_α_to_d(author), "id", nil)); rnt.IsTruthy(v) {
			return v
		} else {
			return rnt.DictGet(τ_cast_α_to_d(author), "email", nil)
		}
	}()
	view_count = rnt.IntOrNone(func() interface{} {
		if v := (rnt.DictGet(τ_cast_α_to_d(video_data), "viewsCount", nil)); rnt.IsTruthy(v) {
			return v
		} else {
			return rnt.DictGet(τ_cast_α_to_d(video_data), "views_count", nil)
		}
	}(), 1, rnt.OptInt{}, 1)
	acc_id = rnt.DictGet(τ_cast_α_to_d(meta_data), "accId", nil)
	item_id = rnt.DictGet(τ_cast_α_to_d(meta_data), "itemId", nil)
	content_id = func() rnt.OptString {
		if rnt.IsTruthy(func() interface{} {
			if v := (acc_id); !(rnt.IsTruthy(v)) {
				return v
			} else {
				return item_id
			}
		}()) {
			return rnt.AsOptString(rnt.StrFormat2("%s_%s", acc_id, item_id))
		} else {
			return video_id
		}
	}()
	thumbnail = rnt.DictGet(τ_cast_α_to_d(meta_data), "poster", nil)
	duration = rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(meta_data), "duration", nil), 1, rnt.OptInt{}, 1)
	timestamp = rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(meta_data), "timestamp", nil), 1, rnt.OptInt{}, 1)
	return rnt.SDict{
		"id":          content_id,
		"title":       title,
		"thumbnail":   thumbnail,
		"timestamp":   timestamp,
		"uploader":    uploader,
		"uploader_id": uploader_id,
		"duration":    duration,
		"view_count":  view_count,
		"formats":     formats,
	}
}

func (self *MailRuIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("MailRu", NewMailRuIE)
}
