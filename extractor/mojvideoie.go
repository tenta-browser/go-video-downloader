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
 * mojvideoie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/mojvideo.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type MojvideoIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewMojvideoIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &MojvideoIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "Mojvideo"
	_VALID_URL = "https?://(?:www\\.)?mojvideo\\.com/video-(?P<display_id>[^/]+)/(?P<id>[a-f0-9]+)"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *MojvideoIE) Key() string {
	return "Mojvideo"
}

func (self *MojvideoIE) Name() string {
	return self.IE_NAME
}

func (self *MojvideoIE) _real_extract(url string) rnt.SDict {
	var (
		display_id rnt.OptString
		duration   rnt.OptString
		error_desc rnt.OptString
		mobj       rnt.Match
		playerapi  string
		thumbnail  rnt.OptString
		title      rnt.OptString
		video_id   rnt.OptString
		video_url  rnt.OptString
	)
	mobj = rnt.ReMatch((self).VALIDURL, url, 0)
	video_id = rnt.ReMatchGroupOne(mobj, "id")
	display_id = rnt.ReMatchGroupOne(mobj, "display_id")
	playerapi = (self).DownloadWebpageURL(rnt.StrFormat2("http://www.mojvideo.com/playerapi.php?v=%s&t=1", video_id), display_id.Get(), rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{}, nil)
	if rnt.StrContains(playerapi, "<error>true</error>") {
		error_desc = (self).HTMLSearchRegexOne("<errordesc>([^<]*)</errordesc>", playerapi, "error description", rnt.NoDefault, false, 0, nil)
		panic(rnt.PyExtractorError(rnt.StrFormat2("%s said: %s", (self).IE_NAME, error_desc), true, rnt.OptString{}))
	}
	title = (self).HTMLSearchRegexOne("<title>([^<]+)</title>", playerapi, "title", rnt.NoDefault, true, 0, nil)
	video_url = (self).HTMLSearchRegexOne("<file>([^<]+)</file>", playerapi, "video URL", rnt.NoDefault, true, 0, nil)
	thumbnail = (self).HTMLSearchRegexOne("<preview>([^<]+)</preview>", playerapi, "thumbnail", rnt.NoDefault, false, 0, nil)
	duration = rnt.ParseDuration((self).HTMLSearchRegexOne("<duration>([^<]+)</duration>", playerapi, "duration", rnt.NoDefault, false, 0, nil))
	return rnt.SDict{
		"id":         video_id,
		"display_id": display_id,
		"url":        video_url,
		"title":      title,
		"thumbnail":  thumbnail,
		"duration":   duration,
	}
}

func (self *MojvideoIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("Mojvideo", NewMojvideoIE)
}
