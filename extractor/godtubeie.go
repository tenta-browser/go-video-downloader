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
 * godtubeie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/godtube.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type GodTubeIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewGodTubeIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &GodTubeIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "GodTube"
	_VALID_URL = "https?://(?:www\\.)?godtube\\.com/watch/\\?v=(?P<id>[\\da-zA-Z]+)"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *GodTubeIE) Key() string {
	return "GodTube"
}

func (self *GodTubeIE) Name() string {
	return self.IE_NAME
}

func (self *GodTubeIE) _real_extract(url string) rnt.SDict {
	var (
		config    rnt.XMLElement
		duration  rnt.OptString
		media     rnt.XMLElement
		mobj      rnt.Match
		thumbnail string
		timestamp rnt.OptInt
		title     string
		uploader  string
		video_id  rnt.OptString
		video_url string
	)
	mobj = rnt.ReMatch((self).VALIDURL, url, 0)
	video_id = rnt.ReMatchGroupOne(mobj, "id")
	config = (self).DownloadXML(rnt.StrFormat2("http://www.godtube.com/resource/mediaplayer/%s.xml", rnt.StrLower(video_id.Get())), video_id.Get(), rnt.AsOptString("Downloading player config XML"), rnt.AsOptString("Unable to download XML"), nil, true, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{}, nil)
	video_url = (rnt.XMLFind(config, "file")).Text
	uploader = (rnt.XMLFind(config, "author")).Text
	timestamp = rnt.ParseISO8601(rnt.AsOptString((rnt.XMLFind(config, "date")).Text), "T")
	duration = rnt.ParseDuration(rnt.AsOptString((rnt.XMLFind(config, "duration")).Text))
	thumbnail = (rnt.XMLFind(config, "image")).Text
	media = (self).DownloadXML(rnt.StrFormat2("http://www.godtube.com/media/xml/?v=%s", video_id), video_id.Get(), rnt.AsOptString("Downloading media XML"), rnt.AsOptString("Unable to download XML"), nil, true, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{}, nil)
	title = (rnt.XMLFind(media, "title")).Text
	return rnt.SDict{
		"id":        video_id,
		"url":       video_url,
		"title":     title,
		"thumbnail": thumbnail,
		"timestamp": timestamp,
		"uploader":  uploader,
		"duration":  duration,
	}
}

func (self *GodTubeIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("GodTube", NewGodTubeIE)
}
