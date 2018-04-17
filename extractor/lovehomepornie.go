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
 * lovehomepornie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/lovehomeporn.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type LoveHomePornIE struct {
	*rnt.CommonIE
	IE_NAME string
}

func NewLoveHomePornIE() rnt.InfoExtractor {
	var (
		IE_NAME    string
		_VALID_URL string
	)
	self := &LoveHomePornIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "LoveHomePorn"
	_VALID_URL = "https?://(?:www\\.)?lovehomeporn\\.com/video/(?P<id>\\d+)(?:/(?P<display_id>[^/?#&]+))?"
	self.IE_NAME = IE_NAME
	self.VALIDURL = _VALID_URL
	return self
}

func (self *LoveHomePornIE) Key() string {
	return "LoveHomePorn"
}

func (self *LoveHomePornIE) Name() string {
	return self.IE_NAME
}

func (self *LoveHomePornIE) _extract_nuevo(config_url string, video_id rnt.OptString, headers rnt.SDict) rnt.SDict {
	var (
		config       rnt.XMLElement
		duration     rnt.OptFloat
		element_name string
		format_id    string
		formats      []interface{}
		thumbnail    rnt.OptString
		title        string
		video_url    rnt.OptString
		τmp1         τ_Tssω
	)
	config = (self).DownloadXML(config_url, video_id.Get(), rnt.AsOptString("Downloading XML"), rnt.AsOptString("Unable to download XML"), func(s string) string {
		return rnt.StrStrip(s, "")
	}, true, rnt.OptString{}, nil, headers, rnt.SDict{})
	title = rnt.StrStrip(rnt.XPathTextOne(config, "./title", rnt.AsOptString("title"), true, rnt.NoDefault).Get(), "")
	video_id = rnt.XPathTextOne(config, "./mediaid", rnt.OptString{}, false, video_id)
	thumbnail = rnt.XPathTextMulti(config, []string{"./image", "./thumb"}, rnt.OptString{}, false, rnt.NoDefault)
	duration = rnt.FloatOrNone(rnt.XPathTextOne(config, "./duration", rnt.OptString{}, false, rnt.NoDefault), 1, 1, rnt.OptFloat{})
	formats = []interface{}{}
	for _, τel := range τ_conv_TTssωTssωω_to_LTssω(τ_TTssωTssωω{
		Φ0: τ_Tssω{
			Φ0: "file",
			Φ1: "sd",
		},
		Φ1: τ_Tssω{
			Φ0: "filehd",
			Φ1: "hd",
		},
	}) {
		τmp1 = τel
		element_name = (τmp1).Φ0
		format_id = (τmp1).Φ1
		video_url = rnt.XPathTextOne(config, element_name, rnt.OptString{}, false, rnt.NoDefault)
		if τ_isTruthy_Os(video_url) {
			formats = append(formats, rnt.SDict{
				"url":       video_url,
				"format_id": format_id,
			})
		}
	}
	formats = τ_conv_Ld_to_Lα((self).CheckFormats(τ_conv_Lα_to_Ld(formats), video_id.Get()))
	return rnt.SDict{
		"id":        video_id,
		"title":     title,
		"thumbnail": thumbnail,
		"duration":  duration,
		"formats":   formats,
	}
}

func (self *LoveHomePornIE) _real_extract(url string) rnt.SDict {
	var (
		display_id rnt.OptString
		info       rnt.SDict
		mobj       rnt.Match
		video_id   rnt.OptString
	)
	mobj = rnt.ReMatch((self).VALIDURL, url, 0)
	video_id = rnt.ReMatchGroupOne(mobj, "id")
	display_id = rnt.ReMatchGroupOne(mobj, "display_id")
	info = (self)._extract_nuevo(rnt.StrFormat2("http://lovehomeporn.com/media/nuevo/config.php?key=%s", video_id), video_id, rnt.SDict{})
	rnt.DictUpdate(info, rnt.SDict{
		"display_id": display_id,
		"age_limit":  18,
	})
	return info
}

func (self *LoveHomePornIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("LoveHomePorn", NewLoveHomePornIE)
}
