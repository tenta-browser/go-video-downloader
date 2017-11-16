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
 * megaphoneie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/megaphone.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type MegaphoneIE struct {
	*rnt.CommonIE
	IE_DESC string
	IE_NAME string
}

func NewMegaphoneIE() rnt.InfoExtractor {
	var (
		IE_DESC    string
		IE_NAME    string
		_VALID_URL string
	)
	self := &MegaphoneIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "Megaphone"
	IE_NAME = "megaphone.fm"
	IE_DESC = "megaphone.fm embedded players"
	_VALID_URL = "https://player\\.megaphone\\.fm/(?P<id>[A-Z0-9]+)"
	self.IE_DESC = IE_DESC
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *MegaphoneIE) Key() string {
	return "Megaphone"
}

func (self *MegaphoneIE) Name() string {
	return self.IE_NAME
}

func (self *MegaphoneIE) _real_extract(url string) rnt.SDict {
	var (
		author       rnt.OptString
		episode_data interface{}
		episode_json rnt.OptString
		formats      []rnt.SDict
		thumbnail    rnt.OptString
		title        rnt.OptString
		video_id     string
		video_url    rnt.OptString
		webpage      string
	)
	video_id = (self).MatchID(url)
	webpage = (self).DownloadWebpageURL(url, video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	title = (self).OgSearchPropertyOne("audio:title", webpage, rnt.OptString{}, rnt.NoDefault, true)
	author = (self).OgSearchPropertyOne("audio:artist", webpage, rnt.OptString{}, rnt.NoDefault, true)
	thumbnail = (self).OgSearchThumbnail(webpage, rnt.NoDefault)
	episode_json = (self).SearchRegexOne("(?s)var\\s+episode\\s*=\\s*(\\{.+?\\});", webpage, "episode JSON", rnt.NoDefault, true, 0, nil)
	episode_data = (self).ParseJSON(episode_json.Get(), video_id, rnt.JsToJSON, true)
	video_url = (self).ProtoRelativeURL(rnt.CastToOptString(rnt.UnsafeSubscript(episode_data, "mediaUrl")), rnt.AsOptString("https:"))
	formats = []rnt.SDict{{
		"url": video_url,
	}}
	return rnt.SDict{
		"id":        video_id,
		"thumbnail": thumbnail,
		"title":     title,
		"author":    author,
		"duration":  rnt.UnsafeSubscript(episode_data, "duration"),
		"formats":   formats,
	}
}

func (self *MegaphoneIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("Megaphone", NewMegaphoneIE)
}
