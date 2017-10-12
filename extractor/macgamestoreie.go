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
 * macgamestoreie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/macgamestore.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type MacGameStoreIE struct {
	*rnt.CommonIE
	IE_NAME string
	IE_DESC string
}

func NewMacGameStoreIE() rnt.InfoExtractor {
	ret := &MacGameStoreIE{}
	ret.CommonIE = rnt.NewCommonIE()
	ret.IE_NAME = `macgamestore`
	ret.IE_DESC = `MacGameStore trailers`
	ret.VALIDURL = `https?://(?:www\.)?macgamestore\.com/mediaviewer\.php\?trailer=(?P<id>\d+)`
	return ret
}

func (self *MacGameStoreIE) Key() string {
	return "MacGameStore"
}

func (self *MacGameStoreIE) Name() string {
	return self.IE_NAME + " (" + self.IE_DESC + ")"
}

func (self *MacGameStoreIE) _real_extract(url string) map[string]interface{} {
	video_id := (self).MatchID(url)
	webpage := (self).DownloadWebpage(url, video_id, rnt.AsOptString(`Downloading trailer page`), rnt.OptString{}, true, 1, 5, rnt.OptString{}, rnt.OptString{}, map[string]interface{}{}, map[string]interface{}{})
	if rnt.StrContains(webpage, `>Missing Media<`) {
		panic(rnt.PyExtractorError(rnt.StrFormat(`Trailer %s does not exist`, video_id), true, rnt.OptString{}))
	}
	video_title := (self).HTMLSearchRegex(`<title>MacGameStore: (.*?) Trailer</title>`, webpage, `title`, rnt.NoDefault, true, 0, nil)
	video_url := (self).HTMLSearchRegex(`(?s)<div\s+id="video-player".*?href="([^"]+)"\s*>`, webpage, `video URL`, rnt.NoDefault, true, 0, nil)
	return map[string]interface{}{`id`: video_id,
		`url`:   video_url,
		`title`: video_title}
}

func (self *MacGameStoreIE) Extract(url string) (*rnt.VideoResult, error) {
	return rnt.RunExtractor(url, self._real_extract)
}

func init() {
	registerFactory(`MacGameStore`, NewMacGameStoreIE)
}
