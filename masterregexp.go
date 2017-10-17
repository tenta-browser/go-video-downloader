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
 * masterregexp.go: Automatically generated regexp by combining URL validation expressions of all extractors
 */

package downloader

func init() {
	masterRegexp = `(?<extr_AnySex>https?://(?:www\.)?anysex\.com/(?:\d+))|(?<extr_ATTTechChannel>https?://techchannel\.att\.com/play-video\.cfm/([^/]+/)*(?:.+))|(?<extr_Bang>https?://(?:www\.)?bang\.com/video/(?:[0-9a-zA-Z_-]+))|(?<extr_Bild>https?://(?:www\.)?bild\.de/(?:[^/]+/)+(?:[^/]+)-(?:\d+)(?:,auto=true)?\.bild\.html)|(?<extr_Criterion>https?://(?:www\.)?criterion\.com/films/(?:[0-9]+)-.+)|(?<extr_EchoMsk>https?://(?:www\.)?echo\.msk\.ru/sounds/(?:\d+))|(?<extr_FiveTV>http://(?:www\.)?5-tv\.ru/(?:(?:[^/]+/)+(?:\d+)|(?:[^/?#]+)(?:[/?#])?))|(?<extr_Hark>https?://(?:www\.)?hark\.com/clips/(?:.+?)-.+)|(?<extr_HentaiStigma>^https?://hentai\.animestigma\.com/(?:[^/]+))|(?<extr_HistoricFilms>https?://(?:www\.)?historicfilms\.com/(?:tapes/|play)(?:\d+))|(?<extr_Keek>https?://(?:www\.)?keek\.com/keek/(?:\w+))|(?<extr_Libsyn>(?:https?://html5-player\.libsyn\.com/embed/episode/id/(?:[0-9]+)))|(?<extr_MacGameStore>https?://(?:www\.)?macgamestore\.com/mediaviewer\.php\?trailer=(?:\d+))|(?<extr_ManyVids>(?i)https?://(?:www\.)?manyvids\.com/video/(?:\d+))|(?<extr_Morningstar>https?://(?:(?:www|news)\.)morningstar\.com/[cC]over/video[cC]enter\.aspx\?id=(?:[0-9]+))|(?<extr_PornHub>https?://(?:(?:[a-z]+\.)?pornhub\.com/(?:(?:view_video\.php|video/show)\?viewkey=|embed/)|(?:www\.)?thumbzilla\.com/video/)(?:[\da-z]+))|(?<extr_RadioJavan>https?://(?:www\.)?radiojavan\.com/videos/video/(?:[^/]+)/?)|(?<extr_RedTube>https?://(?:(?:www\.)?redtube\.com/|embed\.redtube\.com/\?.*?\bid=)(?:[0-9]+))|(?<extr_Slutload>^https?://(?:\w+\.)?slutload\.com/video/[^/]+/(?:[^/]+)/?$)|(?<extr_Stitcher>https?://(?:www\.)?stitcher\.com/podcast/(?:[^/]+/)+e/(?:(?:[^/#?&]+?)-)?(?:\d+)(?:[/#?&]|$))|(?<extr_VideosZ>https?://(?:www\.|m\.|)videosz\.com/[a-z]+/[a-z]+/(?:[0-9]+)_)|(?<extr_WDRMobile>https?://mobile-ondemand\.wdr\.de/.*?/fsk(?:[0-9]+)/[0-9]+/[0-9]+/(?:[0-9]+)_(?:[0-9]+))|(?<extr_XHamster>https?://(?:.+?\.)?xhamster\.com/(?:movies/(?:\d+)/(?:[^/]*)\.html|videos/(?:[^/]*)-(?:\d+)))|(?<extr_XNXX>https?://(?:video|www)\.xnxx\.com/video-?(?:[0-9a-z]+)/)|(?<extr_XTube>(?:xtube:|https?://(?:www\.)?xtube\.com/(?:watch\.php\?.*\bv=|video-watch/(?:[^/]+)-))(?:[^/?&#]+))|(?<extr_XXXYMovies>https?://(?:www\.)?xxxymovies\.com/videos/(?:\d+)/(?:[^/]+))|(?<extr_YouPorn>https?://(?:www\.)?youporn\.com/watch/(?:\d+)/(?:[^/?#&]+))`
}
