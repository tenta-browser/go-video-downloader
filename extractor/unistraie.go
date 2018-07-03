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
 * unistraie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/unistra.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type UnistraIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewUnistraIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &UnistraIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "Unistra"
	_VALID_URL = "https?://utv\\.unistra\\.fr/(?:index|video)\\.php\\?id_video\\=(?P<id>\\d+)"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *UnistraIE) Key() string {
	return "Unistra"
}

func (self *UnistraIE) Name() string {
	return self.IE_NAME
}

func (self *UnistraIE) _real_extract(url string) rnt.SDict {
	var (
		description rnt.OptString
		file_path   string
		files       map[string]struct{}
		format_id   string
		formats     []interface{}
		mobj        rnt.Match
		quality     func(string) int
		thumbnail   rnt.OptString
		title       rnt.OptString
		video_id    rnt.OptString
		webpage     string
	)
	mobj = rnt.ReMatch((self).VALIDURL, url, 0)
	video_id = rnt.ReMatchGroupOne(mobj, "id")
	webpage = (self).DownloadWebpageURL(url, video_id.Get(), rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	files = τ_constr_Ss_from_Ls(rnt.ReFindAllOne("file\\s*:\\s*\"(/[^\"]+)\"", webpage, 0))
	quality = rnt.Qualities([]string{"SD", "HD"})
	formats = []interface{}{}
	for _, τel := range τ_conv_Ss_to_Ls(files) {
		file_path = τel
		format_id = func() string {
			if rnt.StrEndswith(file_path, "-HD.mp4", rnt.OptInt{}, rnt.OptInt{}) {
				return "HD"
			} else {
				return "SD"
			}
		}()
		formats = append(formats, rnt.SDict{
			"url":       rnt.StrFormat2("http://vod-flash.u-strasbg.fr:8080%s", file_path),
			"format_id": format_id,
			"quality":   quality(format_id),
		})
	}
	formats = τ_conv_Ld_to_Lα((self).SortFormats(τ_conv_Lα_to_Ld(formats)))
	title = (self).HTMLSearchRegexOne("<title>UTV - (.*?)</", webpage, "title", rnt.NoDefault, true, 0, nil)
	description = (self).HTMLSearchRegexOne("<meta name=\"Description\" content=\"(.*?)\"", webpage, "description", rnt.NoDefault, true, rnt.ReFlagDotAll, nil)
	thumbnail = (self).SearchRegexOne("image: \"(.*?)\"", webpage, "thumbnail", rnt.NoDefault, true, 0, nil)
	return rnt.SDict{
		"id":          video_id,
		"title":       title,
		"description": description,
		"thumbnail":   thumbnail,
		"formats":     formats,
	}
}

func (self *UnistraIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("Unistra", NewUnistraIE)
}
