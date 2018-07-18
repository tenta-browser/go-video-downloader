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
 * tinypicie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/tinypic.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type TinyPicIE struct {
	*rnt.CommonIE
	IE_DESC string
	IE_NAME string
}

func NewTinyPicIE() rnt.InfoExtractor {
	var (
		IE_DESC    string
		IE_NAME    string
		_VALID_URL string
	)
	self := &TinyPicIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "TinyPic"
	IE_NAME = "tinypic"
	IE_DESC = "tinypic.com videos"
	_VALID_URL = "https?://(?:.+?\\.)?tinypic\\.com/player\\.php\\?v=(?P<id>[^&]+)&s=\\d+"
	self.IE_DESC = IE_DESC
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *TinyPicIE) Key() string {
	return "TinyPic"
}

func (self *TinyPicIE) Name() string {
	return self.IE_NAME
}

func (self *TinyPicIE) _real_extract(url string) rnt.SDict {
	var (
		KEYWORDS_SUFFIX string
		file_id         rnt.OptString
		keywords        rnt.OptString
		mobj            rnt.Match
		server_id       rnt.OptString
		thumbnail       string
		title           string
		video_id        rnt.OptString
		video_url       string
		webpage         string
	)
	mobj = rnt.ReMatch((self).VALIDURL, url, 0)
	video_id = rnt.ReMatchGroupOne(mobj, "id")
	webpage = (self).DownloadWebpageURL(url, video_id.Get(), rnt.AsOptString("Downloading page"), rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{}, nil)
	mobj = rnt.ReSearch("(?m)fo\\.addVariable\\(\"file\",\\s\"(?P<fileid>[\\da-z]+)\"\\);\\n\\s+fo\\.addVariable\\(\"s\",\\s\"(?P<serverid>\\d+)\"\\);", webpage, 0)
	if mobj == nil {
		panic(rnt.PyExtractorError(rnt.StrFormat2("Video %s does not exist", video_id), true, rnt.OptString{}))
	}
	file_id = rnt.ReMatchGroupOne(mobj, "fileid")
	server_id = rnt.ReMatchGroupOne(mobj, "serverid")
	KEYWORDS_SUFFIX = ", Video, images, photos, videos, myspace, ebay, video hosting, photo hosting"
	keywords = (self).HTMLSearchMetaOne("keywords", webpage, rnt.AsOptString("title"), false, rnt.NoDefault, 0)
	title = func() string {
		if rnt.StrEndswith(keywords.Get(), KEYWORDS_SUFFIX, rnt.OptInt{}, rnt.OptInt{}) {
			return rnt.StrSlice(keywords.Get(), rnt.OptInt{}, rnt.AsOptInt(-(len(KEYWORDS_SUFFIX))), rnt.OptInt{})
		} else {
			return ""
		}
	}()
	video_url = rnt.StrFormat2("http://v%s.tinypic.com/%s.flv", server_id, file_id)
	thumbnail = rnt.StrFormat2("http://v%s.tinypic.com/%s_th.jpg", server_id, file_id)
	return rnt.SDict{
		"id":        file_id,
		"url":       video_url,
		"thumbnail": thumbnail,
		"title":     title,
	}
}

func (self *TinyPicIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("TinyPic", NewTinyPicIE)
}
