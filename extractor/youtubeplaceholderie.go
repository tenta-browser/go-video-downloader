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
 * youtubeplaceholderie.go: Automatically transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/youtubeplaceholder.py
 */

package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type YoutubePlaceholderIE struct {
	*rnt.CommonIE
	IE_NAME         string
	_PLAYLIST_ID_RE string
}

func NewYoutubePlaceholderIE() rnt.InfoExtractor {
	var (
		IE_NAME         string
		_PLAYLIST_ID_RE string
		_VALID_URL      string
	)
	self := &YoutubePlaceholderIE{}
	self.CommonIE = rnt.NewCommonIE()
	IE_NAME = "YoutubePlaceholder"
	_PLAYLIST_ID_RE = "(?:PL|LL|EC|UU|FL|RD|UL|TL)[0-9A-Za-z-_]{10,}"
	_VALID_URL = rnt.StrFormat2Named("(?x)^\n                     (\n                         (?:https?://|//)                                    # http(s):// or protocol-independent URL\n                         (?:(?:(?:(?:\\w+\\.)?[yY][oO][uU][tT][uU][bB][eE](?:-nocookie)?\\.com/|\n                            (?:www\\.)?deturl\\.com/www\\.youtube\\.com/|\n                            (?:www\\.)?pwnyoutube\\.com/|\n                            (?:www\\.)?hooktube\\.com/|\n                            (?:www\\.)?yourepeat\\.com/|\n                            tube\\.majestyc\\.net/|\n                            (?:www\\.)?invidio\\.us/|\n                            youtube\\.googleapis\\.com/)                        # the various hostnames, with wildcard subdomains\n                         (?:.*?\\#/)?                                          # handle anchor (#/) redirect urls\n                         (?:                                                  # the various things that can precede the ID:\n                             (?:(?:v|embed|e)/(?!videoseries))                # v/ or embed/ or e/\n                             |(?:                                             # or the v= param in all its forms\n                                 (?:(?:watch|movie)(?:_popup)?(?:\\.php)?/?)?  # preceding watch(_popup|.php) or nothing (like /?v=xxxx)\n                                 (?:\\?|\\#!?)                                  # the params delimiter ? or # or #!\n                                 (?:.*?[&;])??                                # any other preceding param (like /?s=tuff&v=xxxx or ?s=tuff&amp;v=V36LpHqtcDY)\n                                 v=\n                             )\n                         ))\n                         |(?:\n                            youtu\\.be|                                        # just youtu.be/xxxx\n                            vid\\.plus|                                        # or vid.plus/xxxx\n                            zwearz\\.com/watch|                                # or zwearz.com/watch/xxxx\n                         )/\n                         |(?:www\\.)?cleanvideosearch\\.com/media/action/yt/watch\\?videoId=\n                         )\n                     )?                                                       # all until now is optional -> you can pass the naked ID\n                     ([0-9A-Za-z_-]{11})                                      # here is it! the YouTube video ID\n                     (?!.*?\\blist=\n                        (?:\n                            %(playlist_id)s|                                  # combined list/video URLs are handled by the playlist IE\n                            WL                                                # WL are handled by the watch later IE\n                        )\n                     )\n                     #(?(1).+)?                                                # if we found the ID, everything can follow\n                     $", rnt.SDict{
		"playlist_id": _PLAYLIST_ID_RE,
	})
	self.IE_NAME = IE_NAME
	self._PLAYLIST_ID_RE = _PLAYLIST_ID_RE
	self.VALIDURL = _VALID_URL
	return self
}

func (self *YoutubePlaceholderIE) Key() string {
	return "YoutubePlaceholder"
}

func (self *YoutubePlaceholderIE) Name() string {
	return self.IE_NAME
}

func (self *YoutubePlaceholderIE) _real_extract(url string) rnt.SDict {
	panic(rnt.PyExtractorError("YOUTUBE_PLACEHOLDER", true, rnt.OptString{}))
	return nil
}

func (self *YoutubePlaceholderIE) Extract(url string) (rnt.ExtractorResult, error) {
	return rnt.RunExtractor(url, self.Context, self._real_extract)
}

func init() {
	registerFactory("YoutubePlaceholder", NewYoutubePlaceholderIE)
}
