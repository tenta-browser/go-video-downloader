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
 * teachertubeie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/teachertube.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type TeacherTubeIE struct {
	*rnt.CommonIE
	IE_DESC string
	IE_NAME string
}

func NewTeacherTubeIE() rnt.InfoExtractor {
	var (
		IE_DESC    string
		IE_NAME    string
		_VALID_URL string
	)
	self := &TeacherTubeIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "TeacherTube"
	IE_NAME = "teachertube"
	IE_DESC = "teachertube.com videos"
	_VALID_URL = "https?://(?:www\\.)?teachertube\\.com/(viewVideo\\.php\\?video_id=|music\\.php\\?music_id=|video/(?:[\\da-z-]+-)?|audio/)(?P<id>\\d+)"
	self.IE_DESC = IE_DESC
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *TeacherTubeIE) Key() string {
	return "TeacherTube"
}

func (self *TeacherTubeIE) Name() string {
	return self.IE_NAME
}

func (self *TeacherTubeIE) _real_extract(url string) rnt.SDict {
	var (
		TITLE_SUFFIX string
		description  rnt.OptString
		formats      []rnt.SDict
		media_urls   []string
		quality      func(string) int
		title        rnt.OptString
		video_id     string
		webpage      string
	)
	video_id = (self).MatchID(url)
	webpage = (self).DownloadWebpageURL(url, video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{}, rnt.SDict{})
	title = (self).HTMLSearchMetaOne("title", webpage, rnt.AsOptString("title"), true, rnt.NoDefault, 0)
	TITLE_SUFFIX = " - TeacherTube"
	if rnt.StrEndswith(title.Get(), TITLE_SUFFIX, rnt.OptInt{}, rnt.OptInt{}) {
		title = rnt.AsOptString(rnt.StrStrip(rnt.StrSlice(title.Get(), rnt.OptInt{}, rnt.AsOptInt(-(len(TITLE_SUFFIX))), rnt.OptInt{}), ""))
	}
	description = (self).HTMLSearchMetaOne("description", webpage, rnt.AsOptString("description"), false, rnt.NoDefault, 0)
	if τ_isTruthy_Os(description) {
		description = rnt.AsOptString(rnt.StrStrip(description.Get(), ""))
	}
	quality = rnt.Qualities([]string{"mp3", "flv", "mp4"})
	media_urls = rnt.ReFindAllOne("data-contenturl=\"([^\"]+)\"", webpage, 0)
	media_urls = append(media_urls, τ_conv_Ls_to_Ls(rnt.ReFindAllOne("var\\s+filePath\\s*=\\s*\"([^\"]+)\"", webpage, 0))...)
	media_urls = append(media_urls, τ_conv_Ls_to_Ls(rnt.ReFindAllOne("\\'file\\'\\s*:\\s*[\"\\']([^\"\\']+)[\"\\'],", webpage, 0))...)
	formats = func() []rnt.SDict {
		τresult := []rnt.SDict{}
		for _, τel := range τ_conv_Ss_to_Ls(τ_constr_Ss_from_Ls(media_urls)) {
			var (
				media_url string
			)
			media_url = τel
			τresult = append(τresult, rnt.SDict{
				"url":     media_url,
				"quality": quality(rnt.DetermineExt(rnt.AsOptString(media_url), "unknown_video")),
			})
		}
		return τresult
	}()
	formats = (self).SortFormats(formats)
	return rnt.SDict{
		"id":          video_id,
		"title":       title,
		"thumbnail":   (self).HTMLSearchRegexOne("\\'image\\'\\s*:\\s*[\"\\']([^\"\\']+)[\"\\']", webpage, "thumbnail", rnt.NoDefault, true, 0, nil),
		"formats":     formats,
		"description": description,
	}
}

func (self *TeacherTubeIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("TeacherTube", NewTeacherTubeIE)
}
