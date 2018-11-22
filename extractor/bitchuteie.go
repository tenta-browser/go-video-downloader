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
 * bitchuteie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/bitchute.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type BitChuteIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewBitChuteIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &BitChuteIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "BitChute"
	_VALID_URL = "https?://(?:www\\.)?bitchute\\.com/(?:video|embed|torrent/[^/]+)/(?P<id>[^/?#&]+)"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *BitChuteIE) Key() string {
	return "BitChute"
}

func (self *BitChuteIE) Name() string {
	return self.IE_NAME
}

func (self *BitChuteIE) _real_extract(url string) rnt.SDict {
	var (
		description rnt.OptString
		formats     []rnt.SDict
		thumbnail   rnt.OptString
		title       rnt.OptString
		uploader    rnt.OptString
		video_id    string
		webpage     string
	)
	video_id = (self).MatchID(url)
	webpage = (self).DownloadWebpageURL(rnt.StrFormat2("https://www.bitchute.com/video/%s", video_id), video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, nil, rnt.SDict{
		"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.57 Safari/537.36",
	}, rnt.SDict{}, nil)
	title = func() rnt.OptString {
		if v := ((self).HTMLSearchRegexMulti(τ_conv_Tssω_to_Ls(τ_Tssω{
			Φ0: "<[^>]+\\bid=[\"\\']video-title[^>]+>([^<]+)",
			Φ1: "<title>([^<]+)",
		}), webpage, "title", nil, true, 0, nil)); τ_isTruthy_Os(v) {
			return v
		} else if v := ((self).HTMLSearchMetaOne("description", webpage, rnt.AsOptString("title"), false, nil, 0)); τ_isTruthy_Os(v) {
			return v
		} else {
			return (self).OgSearchDescription(webpage, rnt.NoDefault)
		}
	}()
	formats = func() []rnt.SDict {
		τresult := []rnt.SDict{}
		τmp1 := rnt.ReFindIter("addWebSeed\\s*\\(\\s*([\"\\'])(?P<url>(?:(?!\\1).)+)\\1", webpage, 0)
		for τel, τhv := τmp1.Next(); τhv; τel, τhv = τmp1.Next() {
			var (
				mobj rnt.Match
			)
			mobj = τel
			τresult = append(τresult, rnt.SDict{
				"url": rnt.ReMatchGroupOne(mobj, "url"),
			})
		}
		return τresult
	}()
	formats = (self).SortFormats(formats)
	description = (self).HTMLSearchRegexOne("(?s)<div\\b[^>]+\\bclass=[\"\\']full hidden[^>]+>(.+?)</div>", webpage, "description", rnt.NoDefault, false, 0, nil)
	thumbnail = func() rnt.OptString {
		if v := ((self).OgSearchThumbnail(webpage, nil)); τ_isTruthy_Os(v) {
			return v
		} else {
			return (self).HTMLSearchMetaOne("twitter:image:src", webpage, rnt.AsOptString("thumbnail"), false, rnt.NoDefault, 0)
		}
	}()
	uploader = (self).HTMLSearchRegexOne("(?s)<p\\b[^>]+\\bclass=[\"\\']video-author[^>]+>(.+?)</p>", webpage, "uploader", rnt.NoDefault, false, 0, nil)
	return rnt.SDict{
		"id":          video_id,
		"title":       title,
		"description": description,
		"thumbnail":   thumbnail,
		"uploader":    uploader,
		"formats":     formats,
	}
}

func (self *BitChuteIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("BitChute", NewBitChuteIE)
}
