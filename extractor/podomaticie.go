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
 * podomaticie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/podomatic.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type PodomaticIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewPodomaticIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &PodomaticIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "Podomatic"
	IE_NAME = "podomatic"
	_VALID_URL = "(?x)\n                    (?P<proto>https?)://\n                        (?:\n                            (?P<channel>[^.]+)\\.podomatic\\.com/entry|\n                            (?:www\\.)?podomatic\\.com/podcasts/(?P<channel_2>[^/]+)/episodes\n                        )/\n                        (?P<id>[^/?#&]+)\n                "
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *PodomaticIE) Key() string {
	return "Podomatic"
}

func (self *PodomaticIE) Name() string {
	return self.IE_NAME
}

func (self *PodomaticIE) _real_extract(url string) rnt.SDict {
	var (
		channel   rnt.OptString
		data      interface{}
		data_json string
		duration  rnt.OptInt
		json_url  string
		mobj      rnt.Match
		thumbnail interface{}
		title     interface{}
		uploader  interface{}
		video_id  rnt.OptString
		video_url interface{}
	)
	mobj = rnt.ReMatch((self).VALIDURL, url, 0)
	video_id = rnt.ReMatchGroupOne(mobj, "id")
	channel = func() rnt.OptString {
		if v := (rnt.ReMatchGroupOne(mobj, "channel")); τ_isTruthy_Os(v) {
			return v
		} else {
			return rnt.ReMatchGroupOne(mobj, "channel_2")
		}
	}()
	json_url = rnt.StrFormat2(("%s://%s.podomatic.com/entry/embed_params/%s" + "?permalink=true&rtmp=0"), rnt.ReMatchGroupOne(mobj, "proto"), channel, video_id)
	data_json = (self).DownloadWebpageURL(json_url, video_id.Get(), rnt.AsOptString("Downloading video info"), rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	data = rnt.JSONLoads(data_json)
	video_url = rnt.UnsafeSubscript(data, "downloadLink")
	if !(rnt.IsTruthy(video_url)) {
		video_url = rnt.StrFormat2("%s/%s", rnt.StrReplace(rnt.CastToString(rnt.UnsafeSubscript(data, "streamer")), "rtmp", "http", -(1)), rnt.UnsafeSubscript(data, "mediaLocation"))
	}
	uploader = rnt.UnsafeSubscript(data, "podcast")
	title = rnt.UnsafeSubscript(data, "title")
	thumbnail = rnt.UnsafeSubscript(data, "imageLocation")
	duration = rnt.IntOrNone(rnt.DictGet(τ_cast_α_to_d(data), "length", nil), 1000, rnt.OptInt{}, 1)
	return rnt.SDict{
		"id":          video_id,
		"url":         video_url,
		"title":       title,
		"uploader":    uploader,
		"uploader_id": channel,
		"thumbnail":   thumbnail,
		"duration":    duration,
	}
}

func (self *PodomaticIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("Podomatic", NewPodomaticIE)
}
