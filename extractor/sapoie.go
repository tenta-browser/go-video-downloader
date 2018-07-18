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
 * sapoie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/sapo.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type SapoIE struct {
	*rnt.CommonIE
	IE_DESC string
	IE_NAME string
}

func NewSapoIE() rnt.InfoExtractor {
	var (
		IE_DESC    string
		IE_NAME    string
		_VALID_URL string
	)
	self := &SapoIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "Sapo"
	IE_DESC = "SAPO Vídeos"
	_VALID_URL = "https?://(?:(?:v2|www)\\.)?videos\\.sapo\\.(?:pt|cv|ao|mz|tl)/(?P<id>[\\da-zA-Z]{20})"
	self.IE_DESC = IE_DESC
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *SapoIE) Key() string {
	return "Sapo"
}

func (self *SapoIE) Name() string {
	return self.IE_NAME
}

func (self *SapoIE) _real_extract(url string) rnt.SDict {
	var (
		age_limit     int
		categories    []interface{}
		comment_count int
		description   string
		duration      rnt.OptString
		formats       []rnt.SDict
		item          rnt.XMLElement
		mobj          rnt.Match
		tags          string
		thumbnail     rnt.OptString
		title         string
		upload_date   rnt.OptString
		uploader      string
		video_id      rnt.OptString
		video_size    []string
		video_url     string
		view_count    int
	)
	mobj = rnt.ReMatch((self).VALIDURL, url, 0)
	video_id = rnt.ReMatchGroupOne(mobj, "id")
	item = rnt.XMLFind((self).DownloadXML(rnt.StrFormat2("http://rd3.videos.sapo.pt/%s/rss2", video_id), video_id.Get(), rnt.AsOptString("Downloading XML"), rnt.AsOptString("Unable to download XML"), nil, true, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{}, nil), "./channel/item")
	title = (rnt.XMLFind(item, "./title")).Text
	description = (rnt.XMLFind(item, "./{http://videos.sapo.pt/mrss/}synopse")).Text
	thumbnail = rnt.XMLGetAttribute(rnt.XMLFind(item, "./{http://search.yahoo.com/mrss/}content"), "url", rnt.OptString{})
	duration = rnt.ParseDuration(rnt.AsOptString((rnt.XMLFind(item, "./{http://videos.sapo.pt/mrss/}time")).Text))
	uploader = (rnt.XMLFind(item, "./{http://videos.sapo.pt/mrss/}author")).Text
	upload_date = rnt.UnifiedStrDate(rnt.AsOptString((rnt.XMLFind(item, "./pubDate")).Text), true)
	view_count = rnt.ConvertToInt((rnt.XMLFind(item, "./{http://videos.sapo.pt/mrss/}views")).Text)
	comment_count = rnt.ConvertToInt((rnt.XMLFind(item, "./{http://videos.sapo.pt/mrss/}comment_count")).Text)
	tags = (rnt.XMLFind(item, "./{http://videos.sapo.pt/mrss/}tags")).Text
	categories = func() []interface{} {
		if (tags) != "" {
			return τ_conv_Ls_to_Lα(rnt.StrSplit(tags, "", -(1)))
		} else {
			return []interface{}{}
		}
	}()
	age_limit = func() int {
		if ((rnt.XMLFind(item, "./{http://videos.sapo.pt/mrss/}m18")).Text) == ("true") {
			return 18
		} else {
			return 0
		}
	}()
	video_url = (rnt.XMLFind(item, "./{http://videos.sapo.pt/mrss/}videoFile")).Text
	video_size = rnt.StrSplit((rnt.XMLFind(item, "./{http://videos.sapo.pt/mrss/}videoSize")).Text, "x", -(1))
	formats = []rnt.SDict{{
		"url":       video_url,
		"ext":       "mp4",
		"format_id": "sd",
		"width":     rnt.ConvertToInt((video_size)[0]),
		"height":    rnt.ConvertToInt((video_size)[1]),
	}}
	if ((rnt.XMLFind(item, "./{http://videos.sapo.pt/mrss/}HD")).Text) == ("true") {
		formats = append(formats, rnt.SDict{
			"url":       rnt.ReSub("/mov/1$", "/mov/39", video_url, 0, 0),
			"ext":       "mp4",
			"format_id": "hd",
			"width":     1280,
			"height":    720,
		})
	}
	formats = (self).SortFormats(formats)
	return rnt.SDict{
		"id":            video_id,
		"title":         title,
		"description":   description,
		"thumbnail":     thumbnail,
		"duration":      duration,
		"uploader":      uploader,
		"upload_date":   upload_date,
		"view_count":    view_count,
		"comment_count": comment_count,
		"categories":    categories,
		"age_limit":     age_limit,
		"formats":       formats,
	}
}

func (self *SapoIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("Sapo", NewSapoIE)
}
