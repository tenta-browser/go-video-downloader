/**
 * Go Video Downloader
 *
 *    Copyright 2019 Tenta, LLC
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

// Code generated by transpiler. DO NOT EDIT.

package downloader

func init() {
	masterRegexps = []string{
		`(?:https?://(?:abcnews\.go\.com/(?:[^/]+/video/(?:[0-9a-z-]+)-|video/embed\?.*?\` +
			`bid=)|fivethirtyeight\.abcnews\.go\.com/video/embed/\d+/)(?:\d+))|(?:https?://(?` +
			`:(?:(?:embed|www)\.)?acast\.com/|play\.acast\.com/s/)(?:[^/]+)/(?:[^/#?]+))|(?:h` +
			`ttps?://tv\.adobe\.com/(?:(?:fr|de|es|jp)/)?watch/(?:[^/]+)/(?:[^/]+))|(?:https?` +
			`://video\.tv\.adobe\.com/v/(?:\d+))|(?:(?:aol-video:|https?://(?:www\.)?aol\.(?:` +
			`com|ca|co\.uk|de|jp)/video/(?:[^/]+/)*)(?:[0-9a-f]+))|(?:https?://(?:www\.)?allo` +
			`cine\.fr/(?:article|video|film)/(?:fichearticle_gen_carticle=|player_gen_cmedia=` +
			`|fichefilm_gen_cfilm=|video-)(?:[0-9]+)(?:\.html)?)|(?:https?://(?:www\.)?aparat` +
			`\.com/(?:v/|video/video/embed/videohash/)(?:[a-zA-Z0-9]+))|(?:^https?://(?:(?:(?` +
			`:www|classic)\.)?ardmediathek\.de|mediathek\.(?:daserste|rbb-online)\.de|one\.ar` +
			`d\.de)/(?:.*/)(?:[0-9]+|[^0-9][^/\?]+)[^/\?]*(?:\?.*)?)|(?:https?://(?:www\.)?ar` +
			`te\.tv/(?:fr|de|en|es|it|pl)/videos/(?:\d{6}-\d{3}-[AF]))|(?:https?://(?:www\.)?` +
			`(?:(?:(?:asiancrush|yuyutv|midnightpulp)\.com|cocoro\.tv))/video/(?:[^/]+/)?0+(?` +
			`:\d+)v\b)|(?:https?://(?:www\.)?audiomack\.com/song/(?:[\w/-]+))|(?:https?://(?:` +
			`www\.)?(?:telezueri\.ch|telebaern\.tv|telem1\.ch)/[^/]+/(?:[^/]+-(?:\d+))(?:\#vi` +
			`deo=(?:[_0-9a-z]+))?)|(?:https?://[^/]+\.bandcamp\.com/track/(?:[^/?#&]+))|(?:ht` +
			`tps?://(?:www\.)?bandcamp\.com/?\?(?:.*?&)?show=(?:\d+))|(?:https?://(?:www\.|pr` +
			`o\.)?beatport\.com/track/(?:[^/]+)/(?:[0-9]+))|(?:https?://(?:www\.)?bigflix\.co` +
			`m/.+/(?:[0-9]+))|(?:https?://(?:www\.)?bild\.de/(?:[^/]+/)+(?:[^/]+)-(?:\d+)(?:,` +
			`auto=true)?\.bild\.html)|(?:https?://(?:www\.)?bilibili\.com/audio/au(?:\d+))|(?` +
			`:https?://(?:www\.)?bitchute\.com/(?:video|embed|torrent/[^/]+)/(?:[^/?#&]+))|(?` +
			`:https?://(?:www\.)?biqle\.(?:com|org|ru)/watch/(?:-?\d+_\d+))|(?:https?://(?:ww` +
			`w\.)?bleacherreport\.com/articles/(?:\d+))|(?:https?://(?:www\.)?bleacherreport\` +
			`.com/video_embed\?id=(?:[0-9a-f-]{36}|\d{5}))|(?:https?://(?:www\.)?bpb\.de/medi` +
			`athek/(?:[0-9]+)/)|(?:(?:https?://(?:www\.)?br(?:-klassik)?\.de)/(?:[a-z0-9\-_]+` +
			`/)+(?:[a-z0-9\-_]+)\.html)|(?:https?://(?:www\.)?bravotv\.com/(?:[^/]+/)+(?:[^/?` +
			`#]+))|(?:https?://(?:www\.)?camdemy\.com/media/(?:\d+))|(?:https?://(?:www\.)?(?` +
			`:mycanal|piwiplus)\.fr/(?:[^/]+/)*(?:[^?/]+)(?:\.html\?.*\bvid=|/p/)(?:\d+))|(?:` +
			`https?://(?:(?:www\.)?canalc2\.tv/video/|archives-canalc2\.u-strasbg\.fr/video\.` +
			`asp\?.*\bidVideo=)(?:\d+))|(?:(?:cbcplayer:|https?://(?:www\.)?cbc\.ca/(?:player` +
			`/play/|i/caffeine/syndicate/\?mediaId=))(?:\d+))|(?:https?://(?:www\.)?media\.cc` +
			`c\.de/v/(?:[^/?#&]+))|(?:https?://(?:www\.)?ccma\.cat/(?:[^/]+/)*?(?:video|audio` +
			`)/(?:\d+))|(?:https?://(?:(?:[^/]+)\.(?:cntv|cctv)\.(?:com|cn)|(?:www\.)?ncpa-cl` +
			`assic\.com)/(?:[^/]+/)*?(?:[^/?#&]+?)(?:/index)?(?:\.s?html|[?#&]|$))|(?:https?:` +
			`//(?:www\.)?charlierose\.com/(?:video|episode)(?:s|/player)/(?:\d+))|(?:https?:/` +
			`/(?:www\.)?chilloutzone\.net/video/(?:[\w|-]+)\.html)|(?:https?://(?:www\.)?chir` +
			`b\.it/(?:(?:wp|pl)/|fb_chirbit_player\.swf\?key=)?(?:[\da-zA-Z]+))|(?:https?://(` +
			`?:www\.)?cinemax\.com/(?:[^/]+/video/[0-9a-z-]+-(?:\d+)))|(?:https?://(?:www\.)?` +
			`clippituser\.tv/c/(?:[a-z]+))|(?:https?://(?:www\.)?clip\.rs/(?:[^/]+)/\d+)|(?:h` +
			`ttps?://(?:chic|www)\.clipsyndicate\.com/video/play(list/\d+)?/(?:\d+))|(?:https` +
			`?://(?:www\.)?closertotruth\.com/(?:[^/]+/)*(?:[^/?#&]+))|(?:https?://(?:www\.)?` +
			`clyp\.it/(?:[a-z0-9]+))|(?:https?://(?:(?:edition|www|money)\.)?cnn\.com/(?:vide` +
			`o/(?:data/.+?|\?)/)?videos?/(?:.+?/(?:[^/]+?)(?:\.(?:[a-z\-]+)|(?=&))))|(?:https` +
			`?://[^\.]+\.blogs\.cnn\.com/.+)|(?:https?://(?:(?:edition|www)\.)?cnn\.com/(?!vi` +
			`deos?/))|(?:(?:coub:|https?://(?:coub\.com/(?:view|embed|coubs)/|c-cdn\.coub\.co` +
			`m/fb-player\.swf\?.*\bcoub(?:ID|id)=))(?:[\da-z]+))|(?:https?://(?:video|www|pla` +
			`yer(?:-backend)?)\.(?:allure|architecturaldigest|arstechnica|bonappetit|brides|c` +
			`nevids|cntraveler|details|epicurious|glamour|golfdigest|gq|newyorker|self|teenvo` +
			`gue|vanityfair|vogue|wired|wmagazine)\.com/(?:(?:embed(?:js)?|(?:script|inline)/` +
			`video)/(?:[0-9a-f]{24})(?:/(?:[0-9a-f]{24}))?(?:.+?\btarget=(?:[^&]+))?|(?:watch` +
			`|series|video)/(?:[^/?#]+)))|(?:https?://(?:www\.)?cracked\.com/video_(?:\d+)_[\` +
			`da-z-]+\.html)|(?:https?://news\.cts\.com\.tw/[a-z]+/[a-z]+/\d+/(?:\d+)\.html)|(` +
			`?:https?://(?:app\.)?curiositystream\.com/video/(?:\d+))|(?:https?://(?:www\.)?d` +
			`ailymail\.co\.uk/(?:video/[^/]+/video-|embed/video/)(?:[0-9]+))|(?:(?i)https?://` +
			`(?:(?:(?:www|touch)\.)?dailymotion\.[a-z]{2,3}/(?:(?:(?:embed|swf|\#)/)?video|sw` +
			`f)|(?:www\.)?lequipe\.fr/video)/(?:[^/?_]+))|(?:https?://(?:www\.)?dagbladet\.no` +
			`/video/(?:(?:embed|(?:[^/]+))/)?(?:[0-9A-Za-z_-]{11}|[a-zA-Z0-9]{8}))|(?:https?:` +
			`//(?:www\.)?dotsub\.com/view/(?:[^/]+))|(?:https?://(?:(?:www|m)\.)?drtuber\.com` +
			`/(?:video|embed)/(?:\d+)(?:/(?:[\w-]+))?)|(?:https?://video\.aktualne\.cz/(?:[^/` +
			`]+/)+r~(?:[0-9a-f]{32}))|(?:https?://(?:www\.)?(?:discovery|tlc|animalplanet|dma` +
			`x)\.de/(?:.*\#(?:\d+)|(?:[^/]+/)*videos/(?:[^/?#]+)|programme/(?:[^/]+)/video/(?` +
			`:[^/]+)))|(?:https?://(?:(?:[^/]+\.)?(?:disney\.[a-z]{2,3}(?:\.[a-z]{2})?|disney` +
			`(?:(?:me|latino)\.com|turkiye\.com\.tr|channel\.de)|(?:starwars|marvelkids)\.com` +
			`))/(?:(?:embed/|(?:[^/]+/)+[\w-]+-)(?:[a-z0-9]{24})|(?:[^/]+/)?(?:[^/?#]+)))|(?:` +
			`https?://(?:s?evt\.dispeak|events\.digitallyspeaking)\.com/(?:[^/]+/)+xml/(?:[^.` +
			`]+)\.xml)|(?:https?://(?:www\.)?dw\.com/(?:[^/]+/)+(?:av|e)-(?:\d+))|(?:(?:eagle` +
			`platform:(?:[^/]+):|https?://(?:.+?\.media\.eagleplatform\.com)/index/player\?.*` +
			`\brecord_id=)(?:\d+))|(?:https?://(?:www\.)?ebaumsworld\.com/videos/[^/]+/(?:\d+` +
			`))|(?:https?://(?:www\.)?echo\.msk\.ru/sounds/(?:\d+))|(?:https?://(?:[^.]+\.)?e` +
			`lpais\.com/.*/(?:[^/#?]+)\.html(?:$|[?#]))|(?:https?://(?:www\.)?engadget\.com/v` +
			`ideo/(?:[^/?#]+))|(?:https?://(?:www\.)?eporner\.com/(?:hd-porn|embed)/(?:\w+)(?` +
			`:/(?:[\w-]+))?)|(?:https?://?(?:(?:www|v1)\.)?escapistmagazine\.com/videos/view/` +
			`[^/]+/(?:[0-9]+))|(?:https?://(?:www\.)?extremetube\.com/(?:[^/]+/)?video/(?:[^/` +
			`#?&]+))|(?:(?:https?://(?:[\w-]+\.)?(?:facebook\.com|facebookcorewwwi\.onion)/(?` +
			`:[^#]*?\#!/)?(?:(?:video/video\.php|photo\.php|video\.php|video/embed|story\.php` +
			`)\?(?:.*?)(?:v|video_id|story_fbid)=|[^/]+/videos/(?:[^/]+/)?|[^/]+/posts/|group` +
			`s/[^/]+/permalink/)|facebook:)(?:[0-9]+))|(?:https?://(?:[\w-]+\.)?facebook\.com` +
			`/plugins/video\.php\?.*?\bhref=(?:https.+))|(?:https?://(?:www\.)?faz\.net/(?:[^` +
			`/]+/)*.*?-(?:\d+)\.html)|(?:https?://(?:www\.)?fc-zenit\.ru/video/(?:[0-9]+))|(?` +
			`:https?://(?:www\.)?filmweb\.no/(?:trailere|filmnytt)/article(?:\d+)\.ece)|(?:(?` +
			`:5min:|https?://(?:[^/]*?5min\.com/|delivery\.vidible\.tv/aol)(?:(?:Scripts/Play` +
			`erSeed\.js|playerseed/?)?\?.*?playList=)?)(?:\d+))|(?:https?://(?:www\.)?5-tv\.r` +
			`u/(?:(?:[^/]+/)+(?:\d+)|(?:[^/?#]+)(?:[/?#])?))|(?:https?://(?:www\.|secure\.)?f` +
			`lickr\.com/photos/[\w\-_@]+/(?:\d+))|(?:https?://(?:www\.)?formula1\.com/(?:cont` +
			`ent/fom-website/)?en/video/\d{4}/\d{1,2}/(?:.+?)\.html)|(?:https?://(?:video\.(?` +
			`:insider\.)?fox(?:news|business)\.com)/v/(?:video-embed\.html\?video_id=)?(?:\d+` +
			`))|(?:https?://(?:www\.)?(?:insider\.)?foxnews\.com/(?!v)([^/]+/)+(?:[a-z-]+))|(` +
			`?:https?://(?:www\.)?franceculture\.fr/emissions/(?:[^/]+/)*(?:[^/?#&]+))|(?:htt` +
			`ps?://(?:www\.)?franceinter\.fr/emissions/(?:[^?#]+))|(?:https?://generation-wha` +
			`t\.francetv\.fr/[^/]+/video/(?:[^/?#&]+))|(?:https?://(?:www\.)?freesound\.org/p` +
			`eople/[^/]+/sounds/(?:[^/]+))|(?:https?://(?:www\.)?funk\.net/(?:channel|playlis` +
			`t)/[^/]+/(?:[0-9a-z-]+)-(?:\d+))|(?:https?://(?:www\.)?(?:fxnetworks|simpsonswor` +
			`ld)\.com/video/(?:\d+))|(?:https?://(?:www\.)?gameinformer\.com/(?:[^/]+/)*(?:[^` +
			`.?&#]+))|(?:https?://(?:www\.)?gaskrank\.tv/tv/(?:[^/]+)/(?:[^/]+)\.htm)|(?:(?:h` +
			`ttps?://(?:www\.)?gazeta\.ru/(?:[^/]+/)?video/(?:main/)*(?:\d{4}/\d{2}/\d{2}/)?(` +
			`?:[A-Za-z0-9-_.]+)\.s?html))|(?:https?://(?:www\.)?gdcvault\.com/play/(?:\d+)(?:` +
			`/(?:[\w-]+))?)|(?:https?://(?:www\.)?gfycat\.com/(?:ru/|ifr/|gifs/detail/)?(?:[^` +
			`-/?#]+))|(?:https?://(?:www\.)?godtube\.com/watch/\?v=(?:[\da-zA-Z]+))|(?:^https` +
			`?://video\.golem\.de/.+?/(?:.+?)/)|(?:https?://(?:(?:docs|drive)\.google\.com/(?` +
			`:(?:uc|open)\?.*?id=|file/d/)|video\.google\.com/get_player\?.*?docid=)(?:[a-zA-` +
			`Z0-9_-]{28,}))|(?:https?://on-demand\.gputechconf\.com/gtc/2015/video/S(?:\d+)\.` +
			`html)|(?:https?://(?:www\.)?hbo\.com/(?:video|embed)(?:/[^/]+)*/(?:[^/?#]+))|(?:` +
			`https?://(?:www\.)?heise\.de/(?:[^/]+/)+[^/]+-(?:[0-9]+)\.html)|(?:https?://(?:w` +
			`ww\.)?historicfilms\.com/(?:tapes/|play)(?:\d+))|(?:https?://(?:www\.)?hitrecord` +
			`\.org/records/(?:\d+))|(?:http?://(?:www\.)?hornbunny\.com/videos/(?:[a-z-]+)-(?` +
			`:\d+)\.html)|(?:https?://[\da-z-]+\.(?:howstuffworks|stuff(?:(?:youshould|theydo` +
			`ntwantyouto)know|toblowyourmind|momnevertoldyou)|(?:brain|car)stuffshow|fwthinki` +
			`ng|geniusstuff)\.com/(?:[^/]+/)*(?:\d+-)?(?:.+?)-video\.htm)|(?:https?://(?:www\` +
			`.)?hypem\.com/track/(?:[0-9a-z]{5}))|(?:https?://.+?\.ign\.com/(?:[^/]+/)?(?:vid` +
			`eos|show_videos|articles|feature|(?:[^/]+/\d+/video))(/.+)?/(?:.+))|(?:https?://` +
			`(?:www|m)\.imdb\.com/(?:video|title|list).+?[/-]vi(?:\d+))|(?:https?://(?:i\.)?i` +
			`mgur\.com/(?!(?:a|gallery|(?:t(?:opic)?|r)/[^/]+)/)(?:[a-zA-Z0-9]+))|(?:https?:/` +
			`/(?:i\.)?imgur\.com/(?:gallery|(?:t(?:opic)?|r)/[^/]+)/(?:[a-zA-Z0-9]+))|(?:http` +
			`s?://(?:www\.)?ina\.fr/(?:video|audio)/(?:[A-Z0-9_]+))|(?:https?://(?:www\.)?inf` +
			`oq\.com/(?:[^/]+/)+(?:[^/]+))|(?:(?:https?://(?:www\.)?instagram\.com/(?:p|tv)/(` +
			`?:[^/?#&]+)))|(?:https?://(?:www\.)?90tv\.ir/video/(?:[0-9]+)/.*)|(?:https?://(?` +
			`:www\.)?ivideon\.com/tv/(?:[^/]+/)*camera/(?:\d+-[\da-f]+)/(?:\d+))|(?:https?://` +
			`(?:(?:www|m)\.)?izlesene\.com/(?:video|embedplayer)/(?:[^/]+/)?(?:[0-9]+))|(?:ht` +
			`tps?://(?:licensing\.jamendo\.com/[^/]+|(?:www\.)?jamendo\.com)/track/(?:[0-9]+)` +
			`/(?:[^/?#&]+))|(?:https?://.*?\.jeuxvideo\.com/.*/(.*?)\.htm)|(?:(?:joj:|https?:` +
			`//media\.joj\.sk/embed/)(?:[^/?#^]+))|(?:(?:https?://(?:content\.jwplatform|cdn\` +
			`.jwplayer)\.com/(?:(?:feed|player|thumb|preview)s|jw6|v2/media)/|jwplatform:)(?:` +
			`[a-zA-Z0-9]{8}))|(?:https?://tv\.kakao\.com/channel/(?:\d+)/cliplink/(?:\d+))|(?` +
			`:https?://(?:www\.)?keezmovies\.com/video/(?:(?:[^/]+)-)?(?:\d+))|(?:^https?://(` +
			`?:(?:www|api)\.)?khanacademy\.org/(?:[^/]+)/(?:[^/]+/){,2}(?:[^?#/]+)(?:$|[?#]))` +
			`|(?:https?://(?:www\.)?kickstarter\.com/projects/(?:[^/]*)/.*)|(?:https?://(?:ww` +
			`w\.)?loc\.gov/(?:item/|today/cyberlc/feature_wdesc\.php\?.*\brec=)(?:[0-9a-z_.]+` +
			`))|(?:(?:https?://html5-player\.libsyn\.com/embed/episode/id/(?:[0-9]+)))|(?:(?:` +
			`limelight:media:|https?://(?:link\.videoplatform\.limelight\.com/media/|assets\.` +
			`delvenetworks\.com/player/loader\.swf)\?.*?\bmediaId=)(?:[a-z0-9]{32}))|(?:https` +
			`?://(?:new\.)?livestream\.com/(?:accounts/(?:\d+)|(?:[^/]+))/(?:events/(?:\d+)|(` +
			`?:[^/]+))(?:/videos/(?:\d+))?)|(?:https?://(?:www\.)?lovehomeporn\.com/video/(?:` +
			`\d+)(?:/(?:[^/?#&]+))?)|(?:https?://(?:www\.)?(?:lynda\.com|educourse\.ga)/(?:(?` +
			`:[^/]+/){2,3}(?:\d+)|player/embed)/(?:\d+))|(?:https?://my\.mail\.ru/music/songs` +
			`/[^/?#&]+-(?:[\da-f]+))|(?:(?i)https?://(?:www\.)?manyvids\.com/video/(?:\d+))|(` +
			`?:https?://(?:www\.)?massengeschmack\.tv/play/(?:[^?&#]+))|(?:(?:mediaset:|https` +
			`?://(?:(?:www|static3)\.)?mediasetplay\.mediaset\.it/(?:(?:video|on-demand)/(?:[` +
			`^/]+/)+[^/]+_|player/index\.html\?.*?\bprogramGuid=))(?:[0-9A-Z]{16}))|(?:https:` +
			`//player\.megaphone\.fm/(?:[A-Z0-9]+))|(?:https?://(?:www\.)?metacritic\.com/.+?` +
			`/trailers/(?:\d+))|(?:(?:mva:|https?://(?:mva\.microsoft|(?:www\.)?microsoftvirt` +
			`ualacademy)\.com/[^/]+/training-courses/[^/?#&]+-)(?:\d+)(?::|\?l=)(?:[\da-zA-Z]` +
			`+_\d+))|(?:https?://(?:[\da-z_-]+\.)*(?:mlb)\.com/(?:(?:(?:[^/]+/)*c-|(?:shared/` +
			`video/embed/(?:embed|m-internal-embed)\.html|(?:[^/]+/)+(?:play|index)\.jsp|)\?.` +
			`*?\bcontent_id=)(?:\d+)))|(?:https?://(?:www\.)?mofosex\.com/videos/(?:\d+)/(?:[` +
			`^/?#&.]+)\.html)|(?:https?://(?:www\.)?mojvideo\.com/video-(?:[^/]+)/(?:[a-f0-9]` +
			`+))|(?:https?://myspace\.com/[^/]+/(?:video/[^/]+/(?:\d+)|music/song/[^/?#&]+-(?` +
			`:\d+)-\d+(?:[/?#&]|$)))|(?:https?://video\.nationalgeographic\.com/.*?)|(?:https` +
			`?://(?:m\.)?tv(?:cast)?\.naver\.com/v/(?:\d+))|(?:https?://(?:watch\.|www\.)?nba` +
			`\.com/(?:(?:[^/]+/)+(?:[^?]*?))/?(?:/index\.html)?(?:\?.*)?$)|(?:https?://(?:www` +
			`\.)?csnne\.com/video/(?:[0-9a-z-]+))|(?:https?://(?:www\.)?(?:nbcnews|today|msnb` +
			`c)\.com/([^/]+/)*(?:.*-)?(?:[^/?]+))|(?:https?://(?:www\.)?nbcsports\.com//?(?:[` +
			`^/]+/)+(?:[0-9a-z-]+))|(?:https?://vplayer\.nbcsports\.com/(?:[^/]+/)+(?:[0-9a-z` +
			`A-Z_]+))|(?:https?://(?:www\.)?ndr\.de/(?:[^/]+/)*(?:[^/?#]+),[\da-z]+\.html)|(?` +
			`:https?://(?:www\.)?ndr\.de/(?:[^/]+/)*(?:[\da-z]+)-(?:player|externalPlayer)\.h` +
			`tml)|(?:https?://(?:www\.)?n-joy\.de/(?:[^/]+/)*(?:[\da-z]+)-(?:player|externalP` +
			`layer)_[^/]+\.html)|(?:https?://(?:[^/]+\.)?ndtv\.com/(?:[^/]+/)*videos?/?(?:[^/` +
			`]+/)*[^/?^&]+-(?:\d+))|(?:https?://(?:www\.)?netzkino\.de/\#!/(?:[^/]+)/(?:[^/]+` +
			`))|(?:https?://(?:www\.)?newgrounds\.com/(?:audio/listen|portal/view)/(?:[0-9]+)` +
			`)|(?:(?:https?://api\.nexx(?:\.cloud|cdn\.com)/v3/(?:\d+)/videos/byid/|nexx:(?:(` +
			`?:\d+):)?|https?://arc\.nexx\.cloud/api/video/)(?:\d+))|(?:https?://(?:www\.)?(?` +
			`:nhl|wch2016)\.com/(?:[^/]+/)*c-(?:\d+))|(?:https?://(?:www\.)?nonktube\.com/(?:` +
			`(?:video|embed)/|media/nuevo/embed\.php\?.*?\bid=)(?:\d+))|(?:https?://(?:[^/]+\` +
			`.)?noovo\.ca/videos/(?:[^/]+/[^/?#&]+))|(?:https?://media\.cms\.nova\.cz/embed/(` +
			`?:[^/?#&]+))|(?:https?://(?:[^.]+\.)?(?:tv(?:noviny)?|tn|novaplus|vymena|fanda|k` +
			`rasna|doma|prask)\.nova\.cz/(?:[^/]+/)+(?:[^/]+?)(?:\.html|/|$))|(?:https?://(?:` +
			`(?:www|cn)\.)?nowness\.com/(?:story|(?:series|category)/[^/]+)/(?:[^/]+?)(?:$|[?` +
			`#]))|(?:https?://(?:www\.)?ntv\.ru/(?:[^/]+/)*(?:[^/?#&]+))|(?:https?://(?:(?:ww` +
			`w|m|mobile)\.)?(?:odnoklassniki|ok)\.ru/(?:video(?:embed)?/|web-api/video/movieP` +
			`layer/|live/|dk\?.*?st\.mvId=)(?:[\d-]+))|(?:https?://(?:www\.)?onet\.tv/[a-z]/[` +
			`a-z]+/(?:[0-9a-z-]+)/(?:[0-9a-z]+))|(?:https?://(?:www\.)?onet\.tv/[a-z]/(?:[a-z` +
			`]+)(?:[?#]|$))|(?:(?:ooyala:|https?://.+?\.ooyala\.com/.*?(?:embedCode|ec)=)(?:.` +
			`+?)(&|$))|(?:https?://(?:(?:www\.)?(?:verystream\.com|woof\.tube))/(?:stream|e)/` +
			`(?:[a-zA-Z0-9-_]+))|(?:https?://(?:www\.)?outsidetv\.com/(?:[^/]+/)*?play/[a-zA-` +
			`Z0-9]{8}/\d+/\d+/(?:[a-zA-Z0-9]{8}))|(?:https?://(?:(?:www\.)?packtpub\.com/mapt` +
			`|subscription\.packtpub\.com)/video/[^/]+/(?:\d+)/(?:[^/]+)/(?:[^/]+)(?:/(?:[^/?` +
			`&#]+))?)|(?:https?://(?:(?:www\.)?pandora\.tv/view/(?:[^/]+)/(?:\d+)|(?:.+?\.)?c` +
			`hannel\.pandora\.tv/channel/video\.ptv\?|m\.pandora\.tv/?\?))|(?:(?i)https?://(?` +
			`:www\.)?parliamentlive\.tv/Event/Index/(?:[\da-f]{8}-[\da-f]{4}-[\da-f]{4}-[\da-` +
			`f]{4}-[\da-f]{12}))|(?:https?://(?:www\.)?patreon\.com/(?:creation\?hid=|posts/(` +
			`?:[\w-]+-)?)(?:\d+))|(?:https?://(?:(?:(?:video|www|player)\.pbs\.org|video\.apt` +
			`v\.org|video\.gpb\.org|video\.mpbonline\.org|video\.wnpt\.org|video\.wfsu\.org|v` +
			`ideo\.wsre\.org|video\.wtcitv\.org|video\.pba\.org|video\.alaskapublic\.org|vide` +
			`o\.azpbs\.org|portal\.knme\.org|video\.vegaspbs\.org|watch\.aetn\.org|video\.ket` +
			`\.org|video\.wkno\.org|video\.lpb\.org|videos\.oeta\.tv|video\.optv\.org|watch\.` +
			`wsiu\.org|video\.keet\.org|pbs\.kixe\.org|video\.kpbs\.org|video\.kqed\.org|vids` +
			`\.kvie\.org|video\.pbssocal\.org|video\.valleypbs\.org|video\.cptv\.org|watch\.k` +
			`npb\.org|video\.soptv\.org|video\.rmpbs\.org|video\.kenw\.org|video\.kued\.org|v` +
			`ideo\.wyomingpbs\.org|video\.cpt12\.org|video\.kbyueleven\.org|video\.thirteen\.` +
			`org|video\.wgbh\.org|video\.wgby\.org|watch\.njtvonline\.org|watch\.wliw\.org|vi` +
			`deo\.mpt\.tv|watch\.weta\.org|video\.whyy\.org|video\.wlvt\.org|video\.wvpt\.net` +
			`|video\.whut\.org|video\.wedu\.org|video\.wgcu\.org|video\.wpbt2\.org|video\.wuc` +
			`ftv\.org|video\.wuft\.org|watch\.wxel\.org|video\.wlrn\.org|video\.wusf\.usf\.ed` +
			`u|video\.scetv\.org|video\.unctv\.org|video\.pbshawaii\.org|video\.idahoptv\.org` +
			`|video\.ksps\.org|watch\.opb\.org|watch\.nwptv\.org|video\.will\.illinois\.edu|v` +
			`ideo\.networkknowledge\.tv|video\.wttw\.com|video\.iptv\.org|video\.ninenet\.org` +
			`|video\.wfwa\.org|video\.wfyi\.org|video\.mptv\.org|video\.wnin\.org|video\.wnit` +
			`\.org|video\.wpt\.org|video\.wvut\.org|video\.weiu\.net|video\.wqpt\.org|video\.` +
			`wycc\.org|video\.wipb\.org|video\.indianapublicmedia\.org|watch\.cetconnect\.org` +
			`|video\.thinktv\.org|video\.wbgu\.org|video\.wgvu\.org|video\.netnebraska\.org|v` +
			`ideo\.pioneer\.org|watch\.sdpb\.org|video\.tpt\.org|watch\.ksmq\.org|watch\.kpts` +
			`\.org|watch\.ktwu\.org|watch\.easttennesseepbs\.org|video\.wcte\.tv|video\.wljt\` +
			`.org|video\.wosu\.org|video\.woub\.org|video\.wvpublic\.org|video\.wkyupbs\.org|` +
			`video\.kera\.org|video\.mpbn\.net|video\.mountainlake\.org|video\.nhptv\.org|vid` +
			`eo\.vpt\.org|video\.witf\.org|watch\.wqed\.org|video\.wmht\.org|video\.deltabroa` +
			`dcasting\.org|video\.dptv\.org|video\.wcmu\.org|video\.wkar\.org|wnmuvideo\.nmu\` +
			`.edu|video\.wdse\.org|video\.wgte\.org|video\.lptv\.org|video\.kmos\.org|watch\.` +
			`montanapbs\.org|video\.krwg\.org|video\.kacvtv\.org|video\.kcostv\.org|video\.wc` +
			`ny\.org|video\.wned\.org|watch\.wpbstv\.org|video\.wskg\.org|video\.wxxi\.org|vi` +
			`deo\.wpsu\.org|on-demand\.wvia\.org|video\.wtvi\.org|video\.westernreservepublic` +
			`media\.org|video\.ideastream\.org|video\.kcts9\.org|video\.basinpbs\.org|video\.` +
			`houstonpbs\.org|video\.klrn\.org|video\.klru\.tv|video\.wtjx\.org|video\.ideasta` +
			`tions\.org|video\.kbtc\.org)/(?:(?:vir|port)alplayer|video)/(?:[0-9]+)(?:[?/]|$)` +
			`|(?:www\.)?pbs\.org/(?:[^/]+/){1,5}(?:[^/]+?)(?:\.html)?/?(?:$|[?\#])|(?:video|p` +
			`layer)\.pbs\.org/(?:widget/)?partnerplayer/(?:[^/]+)/))|(?:https?://(?:www\.)?pe` +
			`arvideo\.com/video_(?:\d+))|(?:https?://(?:www\.)?people\.com/people/videos/0,,(` +
			`?:\d+),00\.html)|(?:https?://(?:[a-z0-9]+\.)?photobucket\.com/.*(([\?\&]current=` +
			`)|_)(?:.*)\.(?:(flv)|(mp4)))|(?:https?://player\.piksel\.com/v/(?:refid/[^/]+/pr` +
			`efid/)?(?:[a-z0-9_]+))|(?:https?://(?:www\.)?play\.fm/(?:(?:[^/]+/)+(?:[^/]+))/?` +
			`(?:$|[?#]))|(?:https?://(?:www\.)?plays\.tv/(?:video|embeds)/(?:[0-9a-f]{18}))|(` +
			`?:https?://(?:www\.)?playvid\.com/watch(\?v=|/)(?:.+?)(?:#|$))|(?:(?:https?)://(` +
			`?:(?:[^.]+)\.podomatic\.com/entry|(?:www\.)?podomatic\.com/podcasts/(?:[^/]+)/ep` +
			`isodes)/(?:[^/?#&]+))|(?:https?://(?:www\.)?pokemon\.com/[a-z]{2}(?:.*?play=(?:[` +
			`a-z0-9]{32})|/(?:[^/]+/)+(?:[^/?#&]+)))|(?:(?:https?://)(?:www\.|)91porn\.com/.+` +
			`?\?viewkey=(?:[\w\d]+))|(?:https?://(?:[a-zA-Z]+\.)?porn\.com/videos/(?:(?:[^/]+` +
			`)-)?(?:\d+))|(?:https?://(?:www\.)?pornhd\.com/(?:[a-z]{2,4}/)?videos/(?:\d+)(?:` +
			`/(?:.+))?)|(?:https?://(?:(?:[^/]+\.)?(?:pornhub\.(?:com|net))/(?:(?:view_video\` +
			`.php|video/show)\?viewkey=|embed/)|(?:www\.)?thumbzilla\.com/video/)(?:[\da-z]+)` +
			`)|(?:https?://(?:\w+\.)?pornotube\.com/(?:[^?#]*?)/video/(?:[0-9]+))|(?:https?:/` +
			`/(?:www\.)?presstv\.ir/[^/]+/(?:\d+)/(?:\d+)/(?:\d+)/(?:\d+)/(?:[^/]+)?)|(?:http` +
			`s?://(?:.+?)\.(?:radio\.(?:de|at|fr|pt|es|pl|it)|rad\.io))|(?:https?://(?:www\.)` +
			`?radiojavan\.com/videos/video/(?:[^/]+)/?)|(?:https?://[^/]+\.(?:rai\.(?:it|tv)|` +
			`rainews\.it)/.+?-(?:[\da-f]{8}-[\da-f]{4}-[\da-f]{4}-[\da-f]{4}-[\da-f]{12})(?:-` +
			`.+?)?\.html)|(?:https?://(?:videos\.raywenderlich\.com/courses|(?:www\.)?raywend` +
			`erlich\.com)/(?:[^/]+)/lessons/(?:\d+))|(?:https?://(?:(?:www\.)?redtube\.com/|e` +
			`mbed\.redtube\.com/\?.*?\bid=)(?:[0-9]+))|(?:(?:rentv:|https?://(?:www\.)?ren\.t` +
			`v/(?:player|video/epizod)/)(?:\d+))|(?:^https?://(?:www\.)?reverbnation\.com/.*?` +
			`/song/(?:\d+).*?$)|(?:https?://(?:www\.)?rockstargames\.com/videos(?:/video/|#?/` +
			`?\?.*\bvideo=)(?:\d+))|(?:https?://(?:www\.)?prehravac\.rozhlas\.cz/audio/(?:[0-` +
			`9]+))|(?:https?://(?:www\.)?rtp\.pt/play/p(?:[0-9]+)/(?:[^/?#]+)/?)|(?:rts:(?:\d` +
			`+)|https?://(?:.+?\.)?rts\.ch/(?:[^/]+/){2,}(?:[0-9]+)-(?:.+?)\.html)|(?:https?:` +
			`//(?:test)?player\.(?:rutv\.ru|vgtrk\.com)/(?:flash\d+v/container\.swf\?id=|ifra` +
			`me/(?:swf|video|live)/id/|index/iframe/cast_id/)(?:\d+))|(?:https?://(?:www\.)?(` +
			`?:ruutu|supla)\.fi/(?:video|supla)/(?:\d+))|(?:https?://(?:www\.)?(?:safaribooks` +
			`online|(?:learning\.)?oreilly)\.com/(?:library/view/[^/]+/(?:[^/]+)/(?:[^/?\#&]+` +
			`)\.html|videos/[^/]+/[^/]+/(?:[^-]+-[^/?\#&]+)))|(?:https?://(?:(?:v2|www)\.)?vi` +
			`deos\.sapo\.(?:pt|cv|ao|mz|tl)/(?:[\da-zA-Z]{20}))|(?:https?://(?:www\.)?sbs\.co` +
			`m\.au/(?:ondemand|news)/video/(?:single/)?(?:[0-9]+))|(?:https?://(?:www\.)?scre` +
			`encast\.com/t/(?:[a-zA-Z0-9]+))|(?:https?://screencast-o-matic\.com/watch/(?:[0-` +
			`9a-zA-Z]+))|(?:https?://(?:www\.)?senate\.gov/isvp/?\?(?:.+))|(?:https?://(?:www` +
			`\.)?seznamzpravy\.cz/iframe/player\?.*\bsrc=)|(?:https?://(?:.*?\.)?video\.sina\` +
			`.com\.cn/(?:(?:view/|.*\#)(?:\d+)|.+?/(?:[^/?#]+)(?:\.s?html)|api/sinawebApi/out` +
			`play.php/(?:.+?)\.swf))|(?:https?://news\.sky\.com/video/[0-9a-z-]+-(?:[0-9]+))|` +
			`(?:https?://(?:www\.)?slideshare\.net/[^/]+?/(?:.+?)($|\?))|(?:https?://slidesli` +
			`ve\.com/(?:[0-9]+))|(?:https?://(?:www\.)?sonyliv\.com/details/[^/]+/(?:\d+))|(?` +
			`:^(?:https?://)?(?:(?:(?:www\.|m\.)?soundcloud\.com/(?!stations/track)(?:[\w\d-]` +
			`+)/(?!(?:tracks|albums|sets(?:/.+?)?|reposts|likes|spotlight)/?(?:$|[?#]))(?:[\w` +
			`\d-]+)/?(?:[^?]+?)?(?:[?].*)?$)|(?:api\.soundcloud\.com/tracks/(?:\d+)(?:/?\?sec` +
			`ret_token=(?:[^&]+))?)|(?:(?:w|player|p.)\.soundcloud\.com/player/?.*?url=.*)))|` +
			`(?:https?://(?:www\.)?soundgasm\.net/u/(?:[0-9a-zA-Z_-]+)/(?:[0-9a-zA-Z_-]+))|(?` +
			`:https?://(?:[^/]+\.)?spankbang\.com/(?:[\da-z]+)/(?:video|play|embed)\b)|(?:htt` +
			`ps?://(?:www\.)?spiegel\.de/video/[^/]*-(?:[0-9]+)(?:-embed|-iframe)?(?:\.html)?` +
			`(?:#.*)?$)|(?:https?://(?:www\.)?stitcher\.com/podcast/(?:[^/]+/)+e/(?:(?:[^/#?&` +
			`]+?)-)?(?:\d+)(?:[/#?&]|$))|(?:https?://(?:(?:www|play)\.)?(?:srf|rts|rsi|rtr|sw` +
			`issinfo)\.ch/play/(?:tv|radio)/(?:[^/]+/(?:video|audio)/[^?]+|popup(?:video|audi` +
			`o)player)\?id=(?:[0-9a-f\-]{36}|\d+))|(?:https?://openclassroom\.stanford\.edu(?` +
			`:/?|(/MainFolder/(?:HomePage|CoursePage|VideoPage)\.php([?]course=(?:[^&]+)(&vid` +
			`eo=(?:[^&]+))?(&.*)?)?))$)|(?:https?://streamable\.com/(?:[es]/)?(?:\w+))|(?:htt` +
			`ps?://(?:(?:www\.)?sunporno\.com/videos|embeds\.sunporno\.com/embed)/(?:\d+))|(?` +
			`:https?://(?:www\.)?sverigesradio\.se/(?:sida/)?avsnitt/(?:[0-9]+))|(?:https?://` +
			`(?:www\.)?sverigesradio\.se/sida/(?:artikel|gruppsida)\.aspx\?.*?\bartikel=(?:[0` +
			`-9]+))|(?:https?://(?:www\.)?tagesschau\.de/(?:[^/]+/(?:[^/]+/)*?(?:[^/#?]+?(?:-` +
			`?[0-9]+)?))(?:~_?[^/#?]+?)?\.html)|(?:https?://(?:www\.)?tastytrade\.com/tt/show` +
			`s/[^/]+/episodes/(?:[^/?#&]+))|(?:https?://tds\.lifeway\.com/v1/trainingdelivery` +
			`system/courses/(?:\d+)/index\.html)|(?:https?://(?:\w+\.)?teamcoco\.com/(?:([^/]` +
			`+/)*[^/?#]+))|(?:https?://(?:www\.)?teamtreehouse\.com/library/(?:[^/]+))|(?:(?:` +
			`https?://)(?:www|embed(?:-ssl)?)(?:\.ted\.com/((?:playlists(?:/(?:\d+))?)|((?:ta` +
			`lks))|(?:watch)/[^/]+/[^/]+)(/lang/(.*?))?/(?:[\w-]+).*)$)|(?:https?://(?:www\.)` +
			`?tenta\.com/how-to-download-videos)|(?:https?://(?:www\.)?tfo\.org/(?:en|fr)/(?:` +
			`[^/]+/){2}(?:\d+))|(?:(?:https?://(?:link|player)\.theplatform\.com/[sp]/(?:[^/]` +
			`+)/(?:(?:(?:[^/]+/)+select/)?(?:media/(?:guid/\d+/)?)?|(?:(?:[^/\?]+/(?:swf|conf` +
			`ig)|onsite)/select/))?|theplatform:)(?:[^/\?&]+))|(?:https?://thescene\.com/watc` +
			`h/[^/]+/(?:[^/#?]+))|(?:https?://(?:www\.)?thisoldhouse\.com/(?:watch|how-to|tv-` +
			`episode)/(?:[^/?#]+))|(?:https?://(?:www\.)?tmz\.com/videos/(?:[^/?#]+))|(?:http` +
			`s?://player\.(?:tna|emp)flix\.com/video/(?:\d+))|(?:https?://(?:www\.)?tnaflix\.` +
			`com/[^/]+/(?:[^/]+)/video(?:\d+))|(?:https?://(?:www\.)?empflix\.com/(?:videos/(` +
			`?:.+?)-|[^/]+/(?:[^/]+)/video)(?:[0-9]+))|(?:https?://(?:www\.)?moviefap\.com/vi` +
			`deos/(?:[0-9a-f]+)/(?:[^/]+)\.html)|(?:https?://(?:www\.)?toongoggles\.com/shows` +
			`/(?:\d+)(?:/[^/]+/episodes/(?:\d+))?)|(?:https?://videos\.toypics\.net/view/(?:[` +
			`0-9]+))|(?:https?://(?:(?:www|m)\.)?trilulilu\.ro/(?:[^/]+/)?(?:[^/#\?]+))|(?:ht` +
			`tps?://(?:www\.)?tube8\.com/(?:[^/]+/)+(?:[^/]+)/(?:\d+))|(?:https?://(?:[^/?#&]` +
			`+)\.tumblr\.com/(?:post|video)/(?:[0-9]+)(?:$|[/?#]))|(?:https?://(?:www\.)?tvan` +
			`ouvelles\.ca/videos/(?:\d+))|(?:https?://(?:www\.)?tvc\.ru/video/iframe/id/(?:\d` +
			`+))|(?:https?://(?:www\.)?tvc\.ru/(?!video/iframe/id/)(?:[^?#]+))|(?:https?://(?` +
			`:(?:[^/]+)\.)?tvn24(?:bis)?\.pl/(?:[^/]+/)*(?:[^/]+))|(?:(?:tvp:|https?://[^/]+\` +
			`.tvp\.(?:pl|info)/sess/tvplayer\.php\?.*?object_id=)(?:\d+))|(?:https?://[^/]+\.` +
			`tvp\.(?:pl|info)/(?:video/(?:[^,\s]*,)*|(?:(?!\d+/)[^/]+/)*)(?:\d+))|(?:https?:/` +
			`/(?:www\.)?20min\.ch/(?:videotv/*\?.*?\bvid=|videoplayer/videoplayer\.html\?.*?\` +
			`bvideoId@)(?:\d+))|(?:https?://video\.(?:twentythree\.net|23video\.com|filmweb\.` +
			`no)/v\.ihtml/player\.html\?(?:.*?\bphoto(?:_|%5f)id=(?:\d+).*))|(?:https?://(?:c` +
			`lips\.twitch\.tv/(?:[^/]+/)*|(?:www\.)?twitch\.tv/[^/]+/clip/)(?:[^/?#&]+))|(?:h` +
			`ttps?://(?:www\.)?twitter\.com/i/(?:cards/tfw/v1|videos(?:/tweet)?)/(?:\d+))|(?:` +
			`https?://amp\.twimg\.com/v/(?:[0-9a-f\-]{36}))|(?:https?://video\.udn\.com/(?:em` +
			`bed|play)/news/(?:\d+))|(?:https?://(?:www\.)?(?:digiteka\.net|ultimedia\.com)/(` +
			`?:deliver/(?:generic|musique)(?:/[^/]+)*/(?:src|article)|default/index/video(?:g` +
			`eneric|music)/id)/(?:[\d+a-z]+))|(?:https?://utv\.unistra\.fr/(?:index|video)\.p` +
			`hp\?id_video\=(?:\d+))|(?:https?://(?:www\.)?usatoday\.com/(?:[^/]+/)*(?:[^?/#]+` +
			`))|(?:https?://(?:(?:app|embed)\.)?ustudio\.com/embed/(?:[^/]+)/(?:[^/]+))|(?:ht` +
			`tps?://(?:www\.)?video\.varzesh3\.com/(?:[^/]+/)+(?:[^/]+)/?)|(?:https?://(?:[^/` +
			`]+\.)?vbox7\.com/(?:play:|(?:emb/external\.php|player/ext\.swf)\?.*?\bvid=)(?:[\` +
			`da-fA-F]+))|(?:https?://veehd\.com/video/(?:\d+))|(?:https?://(?:www\.)?veoh\.co` +
			`m/(?:watch|embed|iphone/#_Watch)/(?:(?:v|e|yapi-)[\da-zA-Z]+))|(?:https?://(?:.+` +
			`?\.)?vesti\.ru/(?:.+))|(?:(?:https?://(?:www\.)?vevo\.com/watch/(?!playlist|genr` +
			`e)(?:[^/]+/(?:[^/]+/)?)?|https?://cache\.vevo\.com/m/html/embed\.html\?video=|ht` +
			`tps?://videoplayer\.vevo\.com/embed/embedded\?videoId=|https?://embed\.vevo\.com` +
			`/.*?[?&]isrc=|vevo:)(?:[^&?#]+))|(?:(?:https?://(?:www\.)?(?:vgtv.no|bt.no/tv|af` +
			`tenbladet.no/tv|fvn.no/fvntv|aftenposten.no/webtv|ap.vgtv.no/webtv|tv.aftonblade` +
			`t.se/abtv|www.aftonbladet.se/tv)/?(?:(?:\#!/)?(?:video|live)/|embed?.*id=|a(?:rt` +
			`icles)?/)|(?:vgtv|bttv|satv|fvntv|aptv|abtv):)(?:\d+))|(?:https?://(?:www\.)?vid` +
			`dler\.com/(?:v|embed|player)/(?:[a-z0-9]+)(?:.+?\bsecret=(\d+))?)|(?:https?://vi` +
			`dea(?:kid)?\.hu/(?:videok/(?:[^/]+/)*[^?#&]+-|player\?.*?\bv=|player/v/)(?:[^?#&` +
			`]+))|(?:https?://videopress\.com/embed/(?:[\da-zA-Z]+))|(?:https?://(?:www\.|m\.` +
			`|)videosz\.com/[a-z]+/[a-z]+/(?:[0-9]+)_)|(?:https?://(?:www\.)?vidlii\.com/(?:w` +
			`atch|embed)\?.*?\bv=(?:[0-9A-Za-z_-]{11}))|(?:https?://(?:(?:www|(?:player))\.)?` +
			`vimeo(?:pro)?\.com/(?!(?:channels|album|showcase)/[^/?#]+/?(?:$|[?#])|[^/]+/revi` +
			`ew/|ondemand/)(?:.*?/)?(?:(?:play_redirect_hls|moogaloop\.swf)\?clip_id=)?(?:vid` +
			`eos?/)?(?:[0-9]+)(?:/[\da-f]+)?/?(?:[?&].*)?(?:[#].*)?$)|(?:https?://(?:www\.)?v` +
			`imeo\.com/ondemand/(?:[^/?#&]+))|(?:https?://(?:player\.vimple\.(?:ru|co)/iframe` +
			`|vimple\.(?:ru|co))/(?:[\da-f-]{32,36}))|(?:https?://(?:www\.)?vine\.co/(?:v|oem` +
			`bed)/(?:\w+))|(?:(?:viqeo:|https?://cdn\.viqeo\.tv/embed/*\?.*?\bvid=|https?://a` +
			`pi\.viqeo\.tv/v\d+/data/startup?.*?\bvideo(?:%5B%5D|\[\])=)(?:[\da-f]+))|(?:http` +
			`s?://(?:(?:(?:(?:m|new)\.)?vk\.com/video_|(?:www\.)?daxab.com/)ext\.php\?(?:.*?\` +
			`boid=(?:-?\d+).*?\bid=(?:\d+).*)|(?:(?:(?:m|new)\.)?vk\.com/(?:.+?\?.*?z=)?video` +
			`|(?:www\.)?daxab.com/embed/)(?:-?\d+_\d+)(?:.*\blist=(?:[\da-f]+))?))|(?:https?:` +
			`//(?:(?:www|m)\.)?vlive\.tv/video/(?:[0-9]+))|(?:https?://(?:(?:www|m)\.)?vlive\` +
			`.tv/video/(?:[0-9]+)/playlist/(?:[0-9]+))|(?:https?://(?:www\.)?(?:(?:theverge|v` +
			`ox|sbnation|eater|polygon|curbed|racked|funnyordie)\.com|recode\.net)/(?:[^/]+/)` +
			`*(?:[^/?]+))|(?:https?://(?:m\.)?vuclip\.com/w\?.*?cid=(?:[0-9]+))|(?:https?://(` +
			`?:(?:www|view)\.)?vzaar\.com/(?:videos/)?(?:\d+))|(?:(?:washingtonpost:|https?:/` +
			`/(?:www\.)?washingtonpost\.com/video/(?:[^/]+/)*)(?:[\da-f]{8}-[\da-f]{4}-[\da-f` +
			`]{4}-[\da-f]{4}-[\da-f]{12}))|(?:https?://m\.weibo\.cn/status/(?:[0-9]+)(\?.+)?)` +
			`|(?:https?://(?:www|m)\.worldstar(?:candy|hiphop)\.com/(?:videos|android)/video\` +
			`.php\?.*?\bv=(?:[^&]+))|(?:(?:https?://video-api\.wsj\.com/api-video/player/ifra` +
			`me\.html\?.*?\bguid=|https?://(?:www\.)?(?:wsj|barrons)\.com/video/(?:[^/]+/)+|w` +
			`sj:)(?:[a-fA-F0-9-]{36}))|(?:(?i)https?://(?:www\.)?wsj\.com/articles/(?:[^/?#&]` +
			`+))|(?:https?://(?:.+?\.)?(?:xhamster\.(?:com|one|desi)|xhms\.pro|xhamster[27]\.` +
			`com)/(?:movies/(?:\d+)/(?:[^/]*)\.html|videos/(?:[^/]*)-(?:\d+)))|(?:https?://(?` +
			`:.+?\.)?(?:xhamster\.(?:com|one|desi)|xhms\.pro|xhamster[27]\.com)/xembed\.php\?` +
			`video=(?:\d+))|(?:https?://(?:video|www)\.xnxx\.com/video-?(?:[0-9a-z]+)/)|(?:(?` +
			`:xtube:|https?://(?:www\.)?xtube\.com/(?:watch\.php\?.*\bv=|video-watch/(?:embed` +
			`ded/)?(?:[^/]+)-))(?:[^/?&#]+))|(?:https?://(?:(?:www\.)?xvideos\.com/video|flas` +
			`hservice\.xvideos\.com/embedframe/|static-hw\.xvideos\.com/swf/xv-player\.swf\?.` +
			`*?\bid_video=)(?:[0-9]+))|(?:https?://(?:www\.)?xxxymovies\.com/videos/(?:\d+)/(` +
			`?:[^/]+))|(?:(?:https?://(?:(?:[a-zA-Z]{2})\.)?[\da-zA-Z_-]+\.yahoo\.com)/(?:[^/` +
			`]+/)*(?:(?:.+)?-)?(?:[0-9]+)(?:-[a-z]+)?(?:\.html)?)|(?:https?://v\.yinyuetai\.c` +
			`om/video(?:/h5)?/(?:[0-9]+))|(?:https?://(?:\w+\.)?youjizz\.com/videos/(?:[^/#?]` +
			`*-(?:\d+)\.html|embed/(?:\d+)))|(?:https?://(?:www\.)?youporn\.com/watch/(?:\d+)` +
			`/(?:[^/?#&]+))|(?:https?://(?:www\.)?(?:yourupload\.com/(?:watch|embed)|embed\.y` +
			`ourupload\.com)/(?:[A-Za-z0-9]+))|(?:^((?:https?://|//)(?:(?:(?:(?:\w+\.)?[yY][o` +
			`O][uU][tT][uU][bB][eE](?:-nocookie)?\.com/|(?:www\.)?deturl\.com/www\.youtube\.c` +
			`om/|(?:www\.)?pwnyoutube\.com/|(?:www\.)?hooktube\.com/|(?:www\.)?yourepeat\.com` +
			`/|tube\.majestyc\.net/|(?:(?:www|dev)\.)?invidio\.us/|(?:(?:www|no)\.)?invidiou\` +
			`.sh/|(?:(?:www|fi|de)\.)?invidious\.snopyta\.org/|(?:www\.)?invidious\.kabi\.tk/` +
			`|(?:www\.)?invidious\.enkirton\.net/|(?:www\.)?invidious\.13ad\.de/|(?:www\.)?in` +
			`vidious\.mastodon\.host/|(?:www\.)?invidious\.nixnet\.xyz/|(?:www\.)?invidious\.` +
			`drycat\.fr/|(?:www\.)?tube\.poal\.co/|(?:www\.)?vid\.wxzm\.sx/|(?:www\.)?yt\.elu` +
			`kerio\.org/|(?:www\.)?kgg2m7yk5aybusll\.onion/|(?:www\.)?qklhadlycap4cnod\.onion` +
			`/|(?:www\.)?axqzx4s6s54s32yentfqojs3x5i7faxza6xo3ehd4bzzsg2ii4fv2iid\.onion/|(?:` +
			`www\.)?c7hqkpkpemu6e7emz5b4vyz7idjgdvgaaa3dyimmeojqbgpea3xqjoid\.onion/|(?:www\.` +
			`)?fz253lmuao3strwbfbmx46yu7acac2jz27iwtorgmbqlkurlclmancad\.onion/|(?:www\.)?inv` +
			`idious\.l4qlywnpwqsluw65ts7md3khrivpirse744un3x7mlskqauz5pyuzgqd\.onion/|(?:www\` +
			`.)?owxfohz4kjyv25fvlqilyxast7inivgiktls3th44jhk3ej3i7ya\.b32\.i2p/|youtube\.goog` +
			`leapis\.com/)(?:.*?\#/)?(?:(?:(?:v|embed|e)/(?!videoseries))|(?:(?:(?:watch|movi` +
			`e)(?:_popup)?(?:\.php)?/?)?(?:\?|\#!?)(?:.*?[&;])??v=)))|(?:youtu\.be|vid\.plus|` +
			`zwearz\.com/watch|)/|(?:www\.)?cleanvideosearch\.com/media/action/yt/watch\?vide` +
			`oId=))?([0-9A-Za-z_-]{11})(?!.*?\blist=(?:(?:PL|LL|EC|UU|FL|RD|UL|TL|OLAK5uy_)[0` +
			`-9A-Za-z-_]{10,}|WL)))|(?:(?:(?:https?://)?(?:\w+\.)?(?:(?:youtube\.com|invidio\` +
			`.us)/(?:(?:course|view_play_list|my_playlists|artist|playlist|watch|embed/(?:vid` +
			`eoseries|[0-9A-Za-z_-]{11}))\?(?:.*?[&;])*?(?:p|a|list)=|p/)|youtu\.be/[0-9A-Za-` +
			`z_-]{11}\?.*?\blist=)((?:PL|LL|EC|UU|FL|RD|UL|TL|OLAK5uy_)?[0-9A-Za-z-_]{10,}|(?` +
			`:MC)[\w\.]*).*|((?:PL|LL|EC|UU|FL|RD|UL|TL|OLAK5uy_)[0-9A-Za-z-_]{10,})))`,
	}
}
