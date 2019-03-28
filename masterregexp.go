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
		`(?:https?://abcnews\.go\.com/(?:[^/]+/video/(?:[0-9a-z-]+)-|video/embed\?.*?\bid` +
			`=)(?:\d+))|(?:https?://(?:(?:(?:embed|www)\.)?acast\.com/|play\.acast\.com/s/)(?` +
			`:[^/]+)/(?:[^/#?]+))|(?:https?://tv\.adobe\.com/(?:(?:fr|de|es|jp)/)?watch/(?:[^` +
			`/]+)/(?:[^/]+))|(?:https?://video\.tv\.adobe\.com/v/(?:\d+))|(?:(?:aol-video:|ht` +
			`tps?://(?:(?:www|on)\.)?aol\.com/(?:[^/]+/)*(?:[^/?#&]+-)?)(?:[^/?#&]+))|(?:http` +
			`s?://(?:www\.)?allocine\.fr/(?:article|video|film)/(?:fichearticle_gen_carticle=` +
			`|player_gen_cmedia=|fichefilm_gen_cfilm=|video-)(?:[0-9]+)(?:\.html)?)|(?:https?` +
			`://(?:www\.)?aparat\.com/(?:v/|video/video/embed/videohash/)(?:[a-zA-Z0-9]+))|(?` +
			`:https?://(?:www\.)?archive\.org/(?:details|embed)/(?:[^/?#]+)(?:[?].*)?$)|(?:^h` +
			`ttps?://(?:(?:(?:www|classic)\.)?ardmediathek\.de|mediathek\.(?:daserste|rbb-onl` +
			`ine)\.de|one\.ard\.de)/(?:.*/)(?:[0-9]+|[^0-9][^/\?]+)[^/\?]*(?:\?.*)?)|(?:https` +
			`?://(?:www\.)?asiancrush\.com/video/(?:[^/]+/)?0+(?:\d+)v\b)|(?:https?://(?:www\` +
			`.)?audiomack\.com/song/(?:[\w/-]+))|(?:https?://(?:www\.)?(?:telezueri\.ch|teleb` +
			`aern\.tv|telem1\.ch)/[^/]+/(?:[^/]+-(?:\d+))(?:\#video=(?:[_0-9a-z]+))?)|(?:http` +
			`s?://[^/]+\.bandcamp\.com/track/(?:[^/?#&]+))|(?:https?://(?:www\.)?bandcamp\.co` +
			`m/?\?(?:.*?&)?show=(?:\d+))|(?:https?://(?:www\.)?beeg\.(?:com|porn(?:/video)?)/` +
			`(?:\d+))|(?:https?://(?:www\.|pro\.)?beatport\.com/track/(?:[^/]+)/(?:[0-9]+))|(` +
			`?:https?://(?:www\.)?bigflix\.com/.+/(?:[0-9]+))|(?:https?://(?:www\.)?bild\.de/` +
			`(?:[^/]+/)+(?:[^/]+)-(?:\d+)(?:,auto=true)?\.bild\.html)|(?:https?://(?:www\.)?b` +
			`itchute\.com/(?:video|embed|torrent/[^/]+)/(?:[^/?#&]+))|(?:https?://(?:www\.)?b` +
			`iqle\.(?:com|org|ru)/watch/(?:-?\d+_\d+))|(?:https?://(?:www\.)?bleacherreport\.` +
			`com/articles/(?:\d+))|(?:https?://(?:www\.)?bleacherreport\.com/video_embed\?id=` +
			`(?:[0-9a-f-]{36}))|(?:https?://(?:www\.)?bpb\.de/mediathek/(?:[0-9]+)/)|(?:(?:ht` +
			`tps?://(?:www\.)?br(?:-klassik)?\.de)/(?:[a-z0-9\-_]+/)+(?:[a-z0-9\-_]+)\.html)|` +
			`(?:(?:https?://.*brightcove\.com/(services|viewer).*?\?|brightcove:)(?:.*))|(?:h` +
			`ttps?://players\.brightcove\.net/(?:\d+)/(?:[^/]+)_(?:[^/]+)/index\.html\?.*vide` +
			`oId=(?:\d+|ref:[^&]+))|(?:https?://(?:[^/]+\.)?businessinsider\.(?:com|nl)/(?:[^` +
			`/]+/)*(?:[^/?#&]+))|(?:https?://(?:www\.)?byutv\.org/(?:watch|player)/(?!event/)` +
			`(?:[0-9a-f-]+)(?:/(?:[^/?#&]+))?)|(?:https?://(?:www\.)?camdemy\.com/media/(?:\d` +
			`+))|(?:https?://(?:www\.)?(?:mycanal|piwiplus)\.fr/(?:[^/]+/)*(?:[^?/]+)(?:\.htm` +
			`l\?.*\bvid=|/p/)(?:\d+))|(?:https?://(?:(?:www\.)?canalc2\.tv/video/|archives-ca` +
			`nalc2\.u-strasbg\.fr/video\.asp\?.*\bidVideo=)(?:\d+))|(?:(?:cbcplayer:|https?:/` +
			`/(?:www\.)?cbc\.ca/(?:player/play/|i/caffeine/syndicate/\?mediaId=))(?:\d+))|(?:` +
			`https?://(?:www\.)?cbssports\.com/[^/]+/(?:video|news)/(?:[^/?#&]+))|(?:https?:/` +
			`/(?:www\.)?media\.ccc\.de/v/(?:[^/?#&]+))|(?:https?://(?:www\.)?ccma\.cat/(?:[^/` +
			`]+/)*?(?:video|audio)/(?:\d+))|(?:https?://(?:(?:[^/]+)\.(?:cntv|cctv)\.(?:com|c` +
			`n)|(?:www\.)?ncpa-classic\.com)/(?:[^/]+/)*?(?:[^/?#&]+?)(?:/index)?(?:\.s?html|` +
			`[?#&]|$))|(?:https?://(?:www\.)?charlierose\.com/(?:video|episode)(?:s|/player)/` +
			`(?:\d+))|(?:https?://(?:www\.)?chirb\.it/(?:(?:wp|pl)/|fb_chirbit_player\.swf\?k` +
			`ey=)?(?:[\da-zA-Z]+))|(?:https?://player\.cinchcast\.com/.*?(?:assetId|show_id)=` +
			`(?:[0-9]+))|(?:https?://(?:www\.)?clippituser\.tv/c/(?:[a-z]+))|(?:https?://(?:w` +
			`ww\.)?clip\.rs/(?:[^/]+)/\d+)|(?:https?://(?:chic|www)\.clipsyndicate\.com/video` +
			`/play(list/\d+)?/(?:\d+))|(?:https?://(?:www\.)?closertotruth\.com/(?:[^/]+/)*(?` +
			`:[^/?#&]+))|(?:https?://(?:www\.)?clyp\.it/(?:[a-z0-9]+))|(?:https?://(?:(?:edit` +
			`ion|www|money)\.)?cnn\.com/(?:video/(?:data/.+?|\?)/)?videos?/(?:.+?/(?:[^/]+?)(` +
			`?:\.(?:[a-z\-]+)|(?=&))))|(?:https?://[^\.]+\.blogs\.cnn\.com/.+)|(?:https?://(?` +
			`:(?:edition|www)\.)?cnn\.com/(?!videos?/))|(?:(?:coub:|https?://(?:coub\.com/(?:` +
			`view|embed|coubs)/|c-cdn\.coub\.com/fb-player\.swf\?.*\bcoub(?:ID|id)=))(?:[\da-` +
			`z]+))|(?:https?://(?:video|www|player(?:-backend)?)\.(?:allure|architecturaldige` +
			`st|arstechnica|bonappetit|brides|cnevids|cntraveler|details|epicurious|glamour|g` +
			`olfdigest|gq|newyorker|self|teenvogue|vanityfair|vogue|wired|wmagazine)\.com/(?:` +
			`(?:embed(?:js)?|(?:script|inline)/video)/(?:[0-9a-f]{24})(?:/(?:[0-9a-f]{24}))?(` +
			`?:.+?\btarget=(?:[^&]+))?|(?:watch|series|video)/(?:[^/?#]+)))|(?:https?://(?:ww` +
			`w\.)?cracked\.com/video_(?:\d+)_[\da-z-]+\.html)|(?:https?://(?:www\.)?c-span\.o` +
			`rg/video/\?(?:[0-9a-f]+))|(?:https?://news\.cts\.com\.tw/[a-z]+/[a-z]+/\d+/(?:\d` +
			`+)\.html)|(?:https?://(?:app\.)?curiositystream\.com/video/(?:\d+))|(?:https?://` +
			`(?:www\.)?dailymail\.co\.uk/(?:video/[^/]+/video-|embed/video/)(?:[0-9]+))|(?:(?` +
			`i)https?://(?:(www|touch)\.)?dailymotion\.[a-z]{2,3}/(?:(?:(?:embed|swf|#)/)?vid` +
			`eo|swf)/(?:[^/?_]+))|(?:https?://(?:www\.)?digg\.com/video/(?:[^/?#&]+))|(?:http` +
			`s?://(?:www\.)?dotsub\.com/view/(?:[^/]+))|(?:https?://(?:(?:www|m)\.)?drtuber\.` +
			`com/(?:video|embed)/(?:\d+)(?:/(?:[\w-]+))?)|(?:https?://(?:www\.)?(?:discovery|` +
			`tlc|animalplanet|dmax)\.de/(?:.*\#(?:\d+)|(?:[^/]+/)*videos/(?:[^/?#]+)|programm` +
			`e/(?:[^/]+)/video/(?:[^/]+)))|(?:https?://(?:(?:[^/]+\.)?(?:disney\.[a-z]{2,3}(?` +
			`:\.[a-z]{2})?|disney(?:(?:me|latino)\.com|turkiye\.com\.tr|channel\.de)|(?:starw` +
			`ars|marvelkids)\.com))/(?:(?:embed/|(?:[^/]+/)+[\w-]+-)(?:[a-z0-9]{24})|(?:[^/]+` +
			`/)?(?:[^/?#]+)))|(?:https?://(?:s?evt\.dispeak|events\.digitallyspeaking)\.com/(` +
			`?:[^/]+/)+xml/(?:[^.]+)\.xml)|(?:https?://(?:www\.)?dw\.com/(?:[^/]+/)+(?:av|e)-` +
			`(?:\d+))|(?:(?:eagleplatform:(?:[^/]+):|https?://(?:.+?\.media\.eagleplatform\.c` +
			`om)/index/player\?.*\brecord_id=)(?:\d+))|(?:https?://(?:www\.)?ebaumsworld\.com` +
			`/videos/[^/]+/(?:\d+))|(?:https?://(?:www\.)?echo\.msk\.ru/sounds/(?:\d+))|(?:ht` +
			`tps?://(?:[^.]+\.)?elpais\.com/.*/(?:[^/#?]+)\.html(?:$|[?#]))|(?:https?://(?:ww` +
			`w\.)?engadget\.com/video/(?:[^/?#]+))|(?:https?://?(?:(?:www|v1)\.)?escapistmaga` +
			`zine\.com/videos/view/[^/]+/(?:[0-9]+))|(?:https?://(?:www\.)?extremetube\.com/(` +
			`?:[^/]+/)?video/(?:[^/#?&]+))|(?:(?:https?://(?:[\w-]+\.)?(?:facebook\.com|faceb` +
			`ookcorewwwi\.onion)/(?:[^#]*?\#!/)?(?:(?:video/video\.php|photo\.php|video\.php|` +
			`video/embed|story\.php)\?(?:.*?)(?:v|video_id|story_fbid)=|[^/]+/videos/(?:[^/]+` +
			`/)?|[^/]+/posts/|groups/[^/]+/permalink/)|facebook:)(?:[0-9]+))|(?:https?://(?:[` +
			`\w-]+\.)?facebook\.com/plugins/video\.php\?.*?\bhref=(?:https.+))|(?:https?://(?` +
			`:www\.)?faz\.net/(?:[^/]+/)*.*?-(?:\d+)\.html)|(?:https?://(?:www\.)?fc-zenit\.r` +
			`u/video/(?:[0-9]+))|(?:https?://(?:www\.)?filmweb\.no/(?:trailere|filmnytt)/arti` +
			`cle(?:\d+)\.ece)|(?:(?:5min:|https?://(?:[^/]*?5min\.com/|delivery\.vidible\.tv/` +
			`aol)(?:(?:Scripts/PlayerSeed\.js|playerseed/?)?\?.*?playList=)?)(?:\d+))|(?:http` +
			`s?://(?:www\.|secure\.)?flickr\.com/photos/[\w\-_@]+/(?:\d+))|(?:https?://(?:www` +
			`\.)?formula1\.com/(?:content/fom-website/)?en/video/\d{4}/\d{1,2}/(?:.+?)\.html)` +
			`|(?:https?://(?:video\.(?:insider\.)?fox(?:news|business)\.com)/v/(?:video-embed` +
			`\.html\?video_id=)?(?:\d+))|(?:https?://(?:www\.)?(?:insider\.)?foxnews\.com/(?!` +
			`v)([^/]+/)+(?:[a-z-]+))|(?:https?://(?:www\.)?franceculture\.fr/emissions/(?:[^/` +
			`]+/)*(?:[^/?#&]+))|(?:https?://(?:www\.)?franceinter\.fr/emissions/(?:[^?#]+))|(` +
			`?:https?://generation-what\.francetv\.fr/[^/]+/video/(?:[^/?#&]+))|(?:https?://(` +
			`?:www\.)?freesound\.org/people/[^/]+/sounds/(?:[^/]+))|(?:https?://(?:www\.)?(?:` +
			`fxnetworks|simpsonsworld)\.com/video/(?:\d+))|(?:https?://(?:www\.)?gamespot\.co` +
			`m/(?:video|article|review)s/(?:[^/]+/\d+-|embed/)(?:\d+))|(?:https?://(?:www\.)?` +
			`gaskrank\.tv/tv/(?:[^/]+)/(?:[^/]+)\.htm)|(?:(?:https?://(?:www\.)?gazeta\.ru/(?` +
			`:[^/]+/)?video/(?:main/)*(?:\d{4}/\d{2}/\d{2}/)?(?:[A-Za-z0-9-_.]+)\.s?html))|(?` +
			`:https?://(?:www\.)?gdcvault\.com/play/(?:\d+)/(?:(\w|-)+)?)|(?:https?://(?:www\` +
			`.)?gfycat\.com/(?:ifr/|gifs/detail/)?(?:[^/?#]+))|(?:https?://(?:www\.)?godtube\` +
			`.com/watch/\?v=(?:[\da-zA-Z]+))|(?:^https?://video\.golem\.de/.+?/(?:.+?)/)|(?:h` +
			`ttps?://(?:(?:docs|drive)\.google\.com/(?:(?:uc|open)\?.*?id=|file/d/)|video\.go` +
			`ogle\.com/get_player\?.*?docid=)(?:[a-zA-Z0-9_-]{28,}))|(?:https?://on-demand\.g` +
			`putechconf\.com/gtc/2015/video/S(?:\d+)\.html)|(?:https?://(?:www\.)?heise\.de/(` +
			`?:[^/]+/)+[^/]+-(?:[0-9]+)\.html)|(?:https?://(?:www\.)?historicfilms\.com/(?:ta` +
			`pes/|play)(?:\d+))|(?:https?://(?:www\.)?hitrecord\.org/records/(?:\d+))|(?:http` +
			`?://(?:www\.)?hornbunny\.com/videos/(?:[a-z-]+)-(?:\d+)\.html)|(?:https?://[\da-` +
			`z-]+\.(?:howstuffworks|stuff(?:(?:youshould|theydontwantyouto)know|toblowyourmin` +
			`d|momnevertoldyou)|(?:brain|car)stuffshow|fwthinking|geniusstuff)\.com/(?:[^/]+/` +
			`)*(?:\d+-)?(?:.+?)-video\.htm)|(?:https?://(?:www\.)?hypem\.com/track/(?:[0-9a-z` +
			`]{5}))|(?:https?://.+?\.ign\.com/(?:[^/]+/)?(?:videos|show_videos|articles|featu` +
			`re|(?:[^/]+/\d+/video))(/.+)?/(?:.+))|(?:https?://(?:www|m)\.imdb\.com/(?:video|` +
			`title|list).+?[/-]vi(?:\d+))|(?:https?://(?:i\.)?imgur\.com/(?!(?:a|gallery|(?:t` +
			`(?:opic)?|r)/[^/]+)/)(?:[a-zA-Z0-9]+))|(?:https?://(?:i\.)?imgur\.com/(?:gallery` +
			`|(?:t(?:opic)?|r)/[^/]+)/(?:[a-zA-Z0-9]+))|(?:https?://(?:www\.)?ina\.fr/(?:vide` +
			`o|audio)/(?:[A-Z0-9_]+))|(?:https?://(?:www\.)?infoq\.com/(?:[^/]+/)+(?:[^/]+))|` +
			`(?:(?:https?://(?:www\.)?instagram\.com/p/(?:[^/?#&]+)))|(?:https?://(?:www\.)?9` +
			`0tv\.ir/video/(?:[0-9]+)/.*)|(?:https?://(?:www\.)?ivideon\.com/tv/(?:[^/]+/)*ca` +
			`mera/(?:\d+-[\da-f]+)/(?:\d+))|(?:https?://(?:(?:www|m)\.)?izlesene\.com/(?:vide` +
			`o|embedplayer)/(?:[^/]+/)?(?:[0-9]+))|(?:https?://(?:licensing\.jamendo\.com/[^/` +
			`]+|(?:www\.)?jamendo\.com)/track/(?:[0-9]+)/(?:[^/?#&]+))|(?:https?://.*?\.jeuxv` +
			`ideo\.com/.*/(.*?)\.htm)|(?:(?:joj:|https?://media\.joj\.sk/embed/)(?:[^/?#^]+))` +
			`|(?:(?:https?://(?:content\.jwplatform|cdn\.jwplayer)\.com/(?:(?:feed|player|thu` +
			`mb|preview|video|manifest)s|jw6|v2/media)/|jwplatform:)(?:[a-zA-Z0-9]{8}))|(?:ht` +
			`tps?://tv\.kakao\.com/channel/(?:\d+)/cliplink/(?:\d+))|(?:(?:kaltura:(?:\d+):(?` +
			`:[0-9a-z_]+)|https?://(:?(?:www|cdnapi(?:sec)?)\.)?kaltura\.com(?::\d+)?/(?:(?:i` +
			`ndex\.php/(?:kwidget|extwidget/preview)|html5/html5lib/[^/]+/mwEmbedFrame\.php))` +
			`(?:/(?:[^?]+))?(?:\?(?:.*))?))|(?:https?://(?:www\.)?keezmovies\.com/video/(?:(?` +
			`:[^/]+)-)?(?:\d+))|(?:^https?://(?:(?:www|api)\.)?khanacademy\.org/(?:[^/]+)/(?:` +
			`[^/]+/){,2}(?:[^?#/]+)(?:$|[?#]))|(?:https?://(?:www\.)?kickstarter\.com/project` +
			`s/(?:[^/]*)/.*)|(?:https?://krasview\.ru/(?:video|embed)/(?:\d+))|(?:https?://(?` +
			`:www\.)?kuwo\.cn/mv/(?:\d+?)/)|(?:(https?://)?(?:(?:www\.)?la7\.it/([^/]+)/(?:ri` +
			`vedila7|video)/|tg\.la7\.it/repliche-tgla7\?id=)(?:.+))|(?:https?://(?:www\.)?lc` +
			`i\.fr/[^/]+/[\w-]+-(?:\d+)\.html)|(?:https?://(?:www\.)?loc\.gov/(?:item/|today/` +
			`cyberlc/feature_wdesc\.php\?.*\brec=)(?:[0-9a-z_.]+))|(?:(?:https?://html5-playe` +
			`r\.libsyn\.com/embed/episode/id/(?:[0-9]+)))|(?:(?:limelight:media:|https?://(?:` +
			`link\.videoplatform\.limelight\.com/media/|assets\.delvenetworks\.com/player/loa` +
			`der\.swf)\?.*?\bmediaId=)(?:[a-z0-9]{32}))|(?:https?://(?:new\.)?livestream\.com` +
			`/(?:accounts/(?:\d+)|(?:[^/]+))/(?:events/(?:\d+)|(?:[^/]+))(?:/videos/(?:\d+))?` +
			`)|(?:https?://(?:www\.)?lovehomeporn\.com/video/(?:\d+)(?:/(?:[^/?#&]+))?)|(?:ht` +
			`tps?://(?:www\.)?(?:lynda\.com|educourse\.ga)/(?:(?:[^/]+/){2,3}(?:\d+)|player/e` +
			`mbed)/(?:\d+))|(?:https?://(?:www\.)?macgamestore\.com/mediaviewer\.php\?trailer` +
			`=(?:\d+))|(?:https?://my\.mail\.ru/music/songs/[^/?#&]+-(?:[\da-f]+))|(?:(?i)htt` +
			`ps?://(?:www\.)?manyvids\.com/video/(?:\d+))|(?:https?://(?:www\.)?massengeschma` +
			`ck\.tv/play/(?:[^?&#]+))|(?:(?:mediaset:|https?://(?:(?:www|static3)\.)?mediaset` +
			`play\.mediaset\.it/(?:(?:video|on-demand)/(?:[^/]+/)+[^/]+_|player/index\.html\?` +
			`.*?\bprogramGuid=))(?:[0-9A-Z]{16}))|(?:https://player\.megaphone\.fm/(?:[A-Z0-9` +
			`]+))|(?:https?://(?:www\.)?metacritic\.com/.+?/trailers/(?:\d+))|(?:(?:mva:|http` +
			`s?://(?:mva\.microsoft|(?:www\.)?microsoftvirtualacademy)\.com/[^/]+/training-co` +
			`urses/[^/?#&]+-)(?:\d+)(?::|\?l=)(?:[\da-zA-Z]+_\d+))|(?:https?://(?:[\da-z_-]+\` +
			`.)*(?:mlb)\.com/(?:(?:(?:[^/]+/)*c-|(?:shared/video/embed/(?:embed|m-internal-em` +
			`bed)\.html|(?:[^/]+/)+(?:play|index)\.jsp|)\?.*?\bcontent_id=)(?:\d+)))|(?:https` +
			`?://(?:www\.)?mofosex\.com/videos/(?:\d+)/(?:[^/?#&.]+)\.html)|(?:https?://(?:ww` +
			`w\.)?mojvideo\.com/video-(?:[^/]+)/(?:[a-f0-9]+))|(?:https?://myspace\.com/[^/]+` +
			`/(?:video/[^/]+/(?:\d+)|music/song/[^/?#&]+-(?:\d+)-\d+(?:[/?#&]|$)))|(?:(?:http` +
			`s?://(?:www\.)?myvi\.(?:(?:ru/player|tv)/(?:(?:embed/html|flash|api/Video/Get)/|` +
			`content/preloader\.swf\?.*\bid=)|ru/watch/)|myvi:)(?:[\da-zA-Z_-]+))|(?:https?:/` +
			`/video\.nationalgeographic\.com/.*?)|(?:https?://(?:m\.)?tv(?:cast)?\.naver\.com` +
			`/v/(?:\d+))|(?:https?://(?:watch\.|www\.)?nba\.com/(?:(?:[^/]+/)+(?:[^?]*?))/?(?` +
			`:/index\.html)?(?:\?.*)?$)|(?:https?://(?:www\.)?csnne\.com/video/(?:[0-9a-z-]+)` +
			`)|(?:https?://(?:www\.)?(?:nbcnews|today|msnbc)\.com/([^/]+/)*(?:.*-)?(?:[^/?]+)` +
			`)|(?:https?://(?:www\.)?nbcsports\.com//?(?:[^/]+/)+(?:[0-9a-z-]+))|(?:https?://` +
			`vplayer\.nbcsports\.com/(?:[^/]+/)+(?:[0-9a-zA-Z_]+))|(?:https?://(?:www\.)?ndr\` +
			`.de/(?:[^/]+/)*(?:[^/?#]+),[\da-z]+\.html)|(?:https?://(?:www\.)?ndr\.de/(?:[^/]` +
			`+/)*(?:[\da-z]+)-(?:player|externalPlayer)\.html)|(?:https?://(?:www\.)?n-joy\.d` +
			`e/(?:[^/]+/)*(?:[\da-z]+)-(?:player|externalPlayer)_[^/]+\.html)|(?:https?://(?:` +
			`[^/]+\.)?ndtv\.com/(?:[^/]+/)*videos?/?(?:[^/]+/)*[^/?^&]+-(?:\d+))|(?:https?://` +
			`(?:www\.)?netzkino\.de/\#!/(?:[^/]+)/(?:[^/]+))|(?:https?://(?:www\.)?newgrounds` +
			`\.com/(?:audio/listen|portal/view)/(?:[0-9]+))|(?:https?://(?:www\.)?(?:nhl|wch2` +
			`016)\.com/(?:[^/]+/)*c-(?:\d+))|(?:https?://(?:(?:www|cn)\.)?nowness\.com/(?:sto` +
			`ry|(?:series|category)/[^/]+)/(?:[^/]+?)(?:$|[?#]))|(?:(?:npo:|https?://(?:www\.` +
			`)?(?:npo\.nl/(?:[^/]+/)*|(?:ntr|npostart)\.nl/(?:[^/]+/){2,}|omroepwnl\.nl/video` +
			`/fragment/[^/]+__|(?:zapp|npo3)\.nl/(?:[^/]+/){2,}))(?:[^/?#]+))|(?:https?://(?:` +
			`www\.)?ntv\.ru/(?:[^/]+/)*(?:[^/?#&]+))|(?:https?://(?:(?:www|m|mobile)\.)?(?:od` +
			`noklassniki|ok)\.ru/(?:video(?:embed)?/|web-api/video/moviePlayer/|live/|dk\?.*?` +
			`st\.mvId=)(?:[\d-]+))|(?:https?://(?:www\.)?onet\.tv/[a-z]/[a-z]+/(?:[0-9a-z-]+)` +
			`/(?:[0-9a-z]+))|(?:https?://(?:www\.)?onet\.tv/[a-z]/(?:[a-z]+)(?:[?#]|$))|(?:(?` +
			`:ooyala:|https?://.+?\.ooyala\.com/.*?(?:embedCode|ec)=)(?:.+?)(&|$))|(?:https?:` +
			`//(?:www\.)?outsidetv\.com/(?:[^/]+/)*?play/[a-zA-Z0-9]{8}/\d+/\d+/(?:[a-zA-Z0-9` +
			`]{8}))|(?:https?://(?:(?:www\.)?packtpub\.com/mapt|subscription\.packtpub\.com)/` +
			`video/[^/]+/(?:\d+)/(?:\d+)/(?:\d+))|(?:https?://(?:(?:www\.)?pandora\.tv/view/(` +
			`?:[^/]+)/(?:\d+)|(?:.+?\.)?channel\.pandora\.tv/channel/video\.ptv\?|m\.pandora\` +
			`.tv/?\?))|(?:(?i)https?://(?:www\.)?parliamentlive\.tv/Event/Index/(?:[\da-f]{8}` +
			`-[\da-f]{4}-[\da-f]{4}-[\da-f]{4}-[\da-f]{12}))|(?:https?://(?:www\.)?patreon\.c` +
			`om/(?:creation\?hid=|posts/(?:[\w-]+-)?)(?:\d+))|(?:https?://(?:www\.)?pearvideo` +
			`\.com/video_(?:\d+))|(?:https?://(?:www\.)?people\.com/people/videos/0,,(?:\d+),` +
			`00\.html)|(?:https?://player\.performgroup\.com/eplayer(?:/eplayer\.html|\.js)#/` +
			`?(?:[0-9a-f]{26})\.(?:[0-9a-z]{26}))|(?:https?://(?:[a-z0-9]+\.)?photobucket\.co` +
			`m/.*(([\?\&]current=)|_)(?:.*)\.(?:(flv)|(mp4)))|(?:https?://player\.piksel\.com` +
			`/v/(?:[a-z0-9]+))|(?:https?://(?:www\.)?play\.fm/(?:(?:[^/]+/)+(?:[^/]+))/?(?:$|` +
			`[?#]))|(?:https?://(?:www\.)?plays\.tv/(?:video|embeds)/(?:[0-9a-f]{18}))|(?:htt` +
			`ps?://(?:www\.)?playvid\.com/watch(\?v=|/)(?:.+?)(?:#|$))|(?:(?:https?)://(?:(?:` +
			`[^.]+)\.podomatic\.com/entry|(?:www\.)?podomatic\.com/podcasts/(?:[^/]+)/episode` +
			`s)/(?:[^/?#&]+))|(?:https?://(?:www\.)?pokemon\.com/[a-z]{2}(?:.*?play=(?:[a-z0-` +
			`9]{32})|/(?:[^/]+/)+(?:[^/?#&]+)))|(?:https?://(?:[a-zA-Z]+\.)?porn\.com/videos/` +
			`(?:(?:[^/]+)-)?(?:\d+))|(?:https?://(?:www\.)?pornhd\.com/(?:[a-z]{2,4}/)?videos` +
			`/(?:\d+)(?:/(?:.+))?)|(?:https?://(?:(?:[^/]+\.)?(?:pornhub\.(?:com|net))/(?:(?:` +
			`view_video\.php|video/show)\?viewkey=|embed/)|(?:www\.)?thumbzilla\.com/video/)(` +
			`?:[\da-z]+))|(?:https?://(?:\w+\.)?pornotube\.com/(?:[^?#]*?)/video/(?:[0-9]+))|` +
			`(?:https?://(?:www\.)?puhutv\.com/(?:[^/?#&]+)-izle)|(?:https?://(?:www\.)?press` +
			`tv\.ir/[^/]+/(?:\d+)/(?:\d+)/(?:\d+)/(?:\d+)/(?:[^/]+)?)|(?:https?://(?:.+?)\.(?` +
			`:radio\.(?:de|at|fr|pt|es|pl|it)|rad\.io))|(?:https?://(?:www\.)?radiojavan\.com` +
			`/videos/video/(?:[^/]+)/?)|(?:https?://[^/]+\.(?:rai\.(?:it|tv)|rainews\.it)/.+?` +
			`-(?:[\da-f]{8}-[\da-f]{4}-[\da-f]{4}-[\da-f]{4}-[\da-f]{12})(?:-.+?)?\.html)|(?:` +
			`https?://(?:videos\.raywenderlich\.com/courses|(?:www\.)?raywenderlich\.com)/(?:` +
			`[^/]+)/lessons/(?:\d+))|(?:https?://(?:www\.)?(?:rbmaradio|redbullradio)\.com/sh` +
			`ows/(?:[^/]+)/episodes/(?:[^/?#&]+))|(?:https?://(?:(?:www\.)?redtube\.com/|embe` +
			`d\.redtube\.com/\?.*?\bid=)(?:[0-9]+))|(?:(?:rentv:|https?://(?:www\.)?ren\.tv/(` +
			`?:player|video/epizod)/)(?:\d+))|(?:^https?://(?:www\.)?reverbnation\.com/.*?/so` +
			`ng/(?:\d+).*?$)|(?:https?://(?:www\.)?rockstargames\.com/videos(?:/video/|#?/?\?` +
			`.*\bvideo=)(?:\d+))|(?:https?://(?:www\.)?prehravac\.rozhlas\.cz/audio/(?:[0-9]+` +
			`))|(?:rts:(?:\d+)|https?://(?:.+?\.)?rts\.ch/(?:[^/]+/){2,}(?:[0-9]+)-(?:.+?)\.h` +
			`tml)|(?:https?://(?:www\.)?rtvs\.sk/(?:radio|televizia)/archiv/\d+/(?:\d+))|(?:h` +
			`ttps?://(?:test)?player\.(?:rutv\.ru|vgtrk\.com)/(?:flash\d+v/container\.swf\?id` +
			`=|iframe/(?:swf|video|live)/id/|index/iframe/cast_id/)(?:\d+))|(?:https?://(?:ww` +
			`w\.)?(?:ruutu|supla)\.fi/(?:video|supla)/(?:\d+))|(?:https?://(?:(?:v2|www)\.)?v` +
			`ideos\.sapo\.(?:pt|cv|ao|mz|tl)/(?:[\da-zA-Z]{20}))|(?:https?://(?:www\.)?sbs\.c` +
			`om\.au/(?:ondemand|news)/video/(?:single/)?(?:[0-9]+))|(?:https?://(?:www\.)?scr` +
			`eencast\.com/t/(?:[a-zA-Z0-9]+))|(?:https?://(?:www\.)?senate\.gov/isvp/?\?(?:.+` +
			`))|(?:https?://(?:www\.)?seznamzpravy\.cz/iframe/player\?.*\bsrc=)|(?:https?://(` +
			`?:.*?\.)?video\.sina\.com\.cn/(?:(?:view/|.*\#)(?:\d+)|.+?/(?:[^/?#]+)(?:\.s?htm` +
			`l)|api/sinawebApi/outplay.php/(?:.+?)\.swf))|(?:https?://(?:www\.)?skysports\.co` +
			`m/watch/video/(?:[0-9]+))|(?:https?://(?:www\.)?slideshare\.net/[^/]+?/(?:.+?)($` +
			`|\?))|(?:https?://slideslive\.com/(?:[0-9]+))|(?:https?://(?:\w+\.)?slutload\.co` +
			`m/(?:video/[^/]+|embed_player|watch)/(?:[^/]+))|(?:^(?:https?://)?(?:(?:(?:www\.` +
			`|m\.)?soundcloud\.com/(?!stations/track)(?:[\w\d-]+)/(?!(?:tracks|albums|sets(?:` +
			`/.+?)?|reposts|likes|spotlight)/?(?:$|[?#]))(?:[\w\d-]+)/?(?:[^?]+?)?(?:[?].*)?$` +
			`)|(?:api\.soundcloud\.com/tracks/(?:\d+)(?:/?\?secret_token=(?:[^&]+))?)|(?:(?:w` +
			`|player|p.)\.soundcloud\.com/player/?.*?url=.*)))|(?:https?://(?:www\.)?soundgas` +
			`m\.net/u/(?:[0-9a-zA-Z_-]+)/(?:[0-9a-zA-Z_-]+))|(?:https?://(?:[^/]+\.)?spankban` +
			`g\.com/(?:[\da-z]+)/(?:video|play|embed)\b)|(?:https?://(?:www\.)?stitcher\.com/` +
			`podcast/(?:[^/]+/)+e/(?:(?:[^/#?&]+?)-)?(?:\d+)(?:[/#?&]|$))|(?:https?://(?:(?:w` +
			`ww|play)\.)?(?:srf|rts|rsi|rtr|swissinfo)\.ch/play/(?:tv|radio)/[^/]+/(?:video|a` +
			`udio)/[^?]+\?id=(?:[0-9a-f\-]{36}|\d+))|(?:https?://openclassroom\.stanford\.edu` +
			`(?:/?|(/MainFolder/(?:HomePage|CoursePage|VideoPage)\.php([?]course=(?:[^&]+)(&v` +
			`ideo=(?:[^&]+))?(&.*)?)?))$)|(?:https?://streamable\.com/(?:[es]/)?(?:\w+))|(?:h` +
			`ttps?://(?:www\.)?(?:streamango\.com|fruithosts\.net)/(?:f|embed)/(?:[^/?#&]+))|` +
			`(?:https?://(?:(?:www\.)?sunporno\.com/videos|embeds\.sunporno\.com/embed)/(?:\d` +
			`+))|(?:https?://(?:www\.)?swrmediathek\.de/(?:content/)?player\.htm\?show=(?:[\d` +
			`a-f]{8}-[\da-f]{4}-[\da-f]{4}-[\da-f]{4}-[\da-f]{12}))|(?:https?://(?:www\.)?tag` +
			`esschau\.de/(?:[^/]+/(?:[^/]+/)*?(?:[^/#?]+?(?:-?[0-9]+)?))(?:~_?[^/#?]+?)?\.htm` +
			`l)|(?:https?://(?:www\.)?tastytrade\.com/tt/shows/[^/]+/episodes/(?:[^/?#&]+))|(` +
			`?:https?://tds\.lifeway\.com/v1/trainingdeliverysystem/courses/(?:\d+)/index\.ht` +
			`ml)|(?:https?://(?:www\.)?teachertube\.com/(viewVideo\.php\?video_id=|music\.php` +
			`\?music_id=|video/(?:[\da-z-]+-)?|audio/)(?:\d+))|(?:https?://(?:www\.)?teamtree` +
			`house\.com/library/(?:[^/]+))|(?:(?:https?://)(?:www|embed(?:-ssl)?)(?:\.ted\.co` +
			`m/((?:playlists(?:/\d+)?)|((?:talks))|(?:watch)/[^/]+/[^/]+)(/lang/(.*?))?/(?:[\` +
			`w-]+).*)$)|(?:^https?://(?:www\.)?t13\.cl/videos(?:/[^/]+)+/(?:[\w-]+))|(?:https` +
			`?://(?:www\.)?tenta\.com/how-to-download-videos)|(?:https?://(?:www\.)?tfo\.org/` +
			`(?:en|fr)/(?:[^/]+/){2}(?:\d+))|(?:(?:https?://(?:link|player)\.theplatform\.com` +
			`/[sp]/(?:[^/]+)/(?:(?:(?:[^/]+/)+select/)?(?:media/(?:guid/\d+/)?)?|(?:(?:[^/\?]` +
			`+/(?:swf|config)|onsite)/select/))?|theplatform:)(?:[^/\?&]+))|(?:https?://feed\` +
			`.theplatform\.com/f/(?:[^/]+)/(?:[^?/]+)\?(?:[^&]+&)*(?:by(?:Gui|I)d=(?:[^&]+)))` +
			`|(?:https?://thescene\.com/watch/[^/]+/(?:[^/#?]+))|(?:https?://(?:www\.)?thisol` +
			`dhouse\.com/(?:watch|how-to|tv-episode)/(?:[^/?#]+))|(?:https?://(?:m\.)?tiktok\` +
			`.com/v/(?:\d+))|(?:https?://(?:.+?\.)?tinypic\.com/player\.php\?v=(?:[^&]+)&s=\d` +
			`+)|(?:https?://(?:www\.)?tmz\.com/videos/(?:[^/?#]+))|(?:https?://(?:www\.)?tmz\` +
			`.com/\d{4}/\d{2}/\d{2}/(?:[^/]+)/?)|(?:https?://player\.(?:tna|emp)flix\.com/vid` +
			`eo/(?:\d+))|(?:https?://(?:www\.)?tnaflix\.com/[^/]+/(?:[^/]+)/video(?:\d+))|(?:` +
			`https?://(?:www\.)?empflix\.com/(?:videos/(?:.+?)-|[^/]+/(?:[^/]+)/video)(?:[0-9` +
			`]+))|(?:https?://(?:www\.)?moviefap\.com/videos/(?:[0-9a-f]+)/(?:[^/]+)\.html)|(` +
			`?:https?://(?:www\.)?toongoggles\.com/shows/(?:\d+)(?:/[^/]+/episodes/(?:\d+))?)` +
			`|(?:https?://videos\.toypics\.net/view/(?:[0-9]+))|(?:https?://(?:(?:www|m)\.)?t` +
			`rilulilu\.ro/(?:[^/]+/)?(?:[^/#\?]+))|(?:https?://(?:www\.)?tube8\.com/(?:[^/]+/` +
			`)+(?:[^/]+)/(?:\d+))|(?:https?://(?:[^/?#&]+)\.tumblr\.com/(?:post|video)/(?:[0-` +
			`9]+)(?:$|[/?#]))|(?:https?://(?:(?:www\.)?tune\.pk/(?:video/|player/embed_player` +
			`.php?.*?\bvid=)|embed\.tune\.pk/play/)(?:\d+))|(?:https?://(?:www\.)?tvanouvelle` +
			`s\.ca/videos/(?:\d+))|(?:https?://(?:www\.)?tvc\.ru/video/iframe/id/(?:\d+))|(?:` +
			`https?://(?:www\.)?tvc\.ru/(?!video/iframe/id/)(?:[^?#]+))|(?:https?://(?:(?:[^/` +
			`]+)\.)?tvn24(?:bis)?\.pl/(?:[^/]+/)*(?:[^/]+))|(?:(?:tvp:|https?://[^/]+\.tvp\.(` +
			`?:pl|info)/sess/tvplayer\.php\?.*?object_id=)(?:\d+))|(?:https?://[^/]+\.tvp\.(?` +
			`:pl|info)/(?:video/(?:[^,\s]*,)*|(?:(?!\d+/)[^/]+/)*)(?:\d+))|(?:https?://tvplay` +
			`\.(?:tv3\.lt|skaties\.lv|tv3\.ee)/[^/]+/[^/?#&]+-(?:\d+))|(?:https?://(?:www\.)?` +
			`20min\.ch/(?:videotv/*\?.*?\bvid=|videoplayer/videoplayer\.html\?.*?\bvideoId@)(` +
			`?:\d+))|(?:https?://video\.(?:twentythree\.net|23video\.com|filmweb\.no)/v\.ihtm` +
			`l/player\.html\?(?:.*?\bphoto(?:_|%5f)id=(?:\d+).*))|(?:https?://(?:clips\.twitc` +
			`h\.tv/(?:[^/]+/)*|(?:www\.)?twitch\.tv/[^/]+/clip/)(?:[^/?#&]+))|(?:https?://(?:` +
			`www\.)?twitter\.com/i/(?:cards/tfw/v1|videos(?:/tweet)?)/(?:\d+))|(?:https?://am` +
			`p\.twimg\.com/v/(?:[0-9a-f\-]{36}))|(?:https?://video\.udn\.com/(?:embed|play)/n` +
			`ews/(?:\d+))|(?:https?://(?:www\.)?(?:digiteka\.net|ultimedia\.com)/(?:deliver/(` +
			`?:generic|musique)(?:/[^/]+)*/(?:src|article)|default/index/video(?:generic|musi` +
			`c)/id)/(?:[\d+a-z]+))|(?:https?://utv\.unistra\.fr/(?:index|video)\.php\?id_vide` +
			`o\=(?:\d+))|(?:https?://(?:www\.)?unity3d\.com/learn/tutorials/(?:[^/]+/)*(?:[^/` +
			`?#&]+))|(?:https?://(?:.+?\.)?uol\.com\.br/.*?(?:(?:mediaId|v)=|view/(?:[a-z0-9]` +
			`+/)?|video(?:=|/(?:\d{4}/\d{2}/\d{2}/)?))(?:\d+|[\w-]+-[A-Z0-9]+))|(?:https?://(` +
			`?:www\.)?usatoday\.com/(?:[^/]+/)*(?:[^?/#]+))|(?:https?://(?:(?:app|embed)\.)?u` +
			`studio\.com/embed/(?:[^/]+)/(?:[^/]+))|(?:https?://(?:www\.)?video\.varzesh3\.co` +
			`m/(?:[^/]+/)+(?:[^/]+)/?)|(?:https?://(?:[^/]+\.)?vbox7\.com/(?:play:|(?:emb/ext` +
			`ernal\.php|player/ext\.swf)\?.*?\bvid=)(?:[\da-fA-F]+))|(?:https?://veehd\.com/v` +
			`ideo/(?:\d+))|(?:https?://(?:www\.)?veoh\.com/(?:watch|embed|iphone/#_Watch)/(?:` +
			`(?:v|e|yapi-)[\da-zA-Z]+))|(?:https?://(?:.+?\.)?vesti\.ru/(?:.+))|(?:(?:https?:` +
			`//(?:www\.)?vevo\.com/watch/(?!playlist|genre)(?:[^/]+/(?:[^/]+/)?)?|https?://ca` +
			`che\.vevo\.com/m/html/embed\.html\?video=|https?://videoplayer\.vevo\.com/embed/` +
			`embedded\?videoId=|vevo:)(?:[^&?#]+))|(?:(?:https?://(?:www\.)?(?:vgtv.no|bt.no/` +
			`tv|aftenbladet.no/tv|fvn.no/fvntv|aftenposten.no/webtv|ap.vgtv.no/webtv|tv.afton` +
			`bladet.se/abtv|www.aftonbladet.se/tv)/?(?:(?:\#!/)?(?:video|live)/|embed?.*id=|a` +
			`(?:rticles)?/)|(?:vgtv|bttv|satv|fvntv|aptv|abtv):)(?:\d+))|(?:https://www\.vice` +
			`\.com/[^/]+/article/(?:[^?#]+))|(?:https?://(?:www\.)?viddler\.com/(?:v|embed|pl` +
			`ayer)/(?:[a-z0-9]+)(?:.+?\bsecret=(\d+))?)|(?:https?://videa(?:kid)?\.hu/(?:vide` +
			`ok/(?:[^/]+/)*[^?#&]+-|player\?.*?\bv=|player/v/)(?:[^?#&]+))|(?:https?://videop` +
			`ress\.com/embed/(?:[\da-zA-Z]+))|(?:https?://(?:www\.|m\.|)videosz\.com/[a-z]+/[` +
			`a-z]+/(?:[0-9]+)_)|(?:https?://(?:www\.)?vidlii\.com/(?:watch|embed)\?.*?\bv=(?:` +
			`[0-9A-Za-z_-]{11}))|(?:https?://(?:(?:www|(?:player))\.)?vimeo(?:pro)?\.com/(?!(` +
			`?:channels|album)/[^/?#]+/?(?:$|[?#])|[^/]+/review/|ondemand/)(?:.*?/)?(?:(?:pla` +
			`y_redirect_hls|moogaloop\.swf)\?clip_id=)?(?:videos?/)?(?:[0-9]+)(?:/[\da-f]+)?/` +
			`?(?:[?&].*)?(?:[#].*)?$)|(?:https?://(?:www\.)?vimeo\.com/ondemand/(?:[^/?#&]+))` +
			`|(?:(?:https://vimeo\.com/[^/]+/review/(?:[^/]+)/[0-9a-f]{10}))|(?:https?://(?:p` +
			`layer\.vimple\.(?:ru|co)/iframe|vimple\.(?:ru|co))/(?:[\da-f-]{32,36}))|(?:https` +
			`?://(?:www\.)?vine\.co/(?:v|oembed)/(?:\w+))|(?:(?:viqeo:|https?://cdn\.viqeo\.t` +
			`v/embed/*\?.*?\bvid=|https?://api\.viqeo\.tv/v\d+/data/startup?.*?\bvideo(?:%5B%` +
			`5D|\[\])=)(?:[\da-f]+))|(?:https?://(?:(?:(?:(?:m|new)\.)?vk\.com/video_|(?:www\` +
			`.)?daxab.com/)ext\.php\?(?:.*?\boid=(?:-?\d+).*?\bid=(?:\d+).*)|(?:(?:(?:m|new)\` +
			`.)?vk\.com/(?:.+?\?.*?z=)?video|(?:www\.)?daxab.com/embed/)(?:-?\d+_\d+)(?:.*\bl` +
			`ist=(?:[\da-f]+))?))|(?:https?://(?:(?:www|m)\.)?vlive\.tv/video/(?:[0-9]+))|(?:` +
			`https?://(?:(?:www|m)\.)?vlive\.tv/video/(?:[0-9]+)/playlist/(?:[0-9]+))|(?:http` +
			`s?://vod\.pl/(?:[^/]+/)+(?:[0-9a-zA-Z]+))|(?:(?:washingtonpost:|https?://(?:www\` +
			`.)?washingtonpost\.com/video/(?:[^/]+/)*)(?:[\da-f]{8}-[\da-f]{4}-[\da-f]{4}-[\d` +
			`a-f]{4}-[\da-f]{12}))|(?:(?:wat:|https?://(?:www\.)?wat\.tv/video/.*-)(?:[0-9a-z` +
			`]+))|(?:https?://m\.weibo\.cn/status/(?:[0-9]+)(\?.+)?)|(?:https?://(?:www|m)\.w` +
			`orldstar(?:candy|hiphop)\.com/(?:videos|android)/video\.php\?.*?\bv=(?:[^&]+))|(` +
			`?:(?:https?://video-api\.wsj\.com/api-video/player/iframe\.html\?.*?\bguid=|http` +
			`s?://(?:www\.)?(?:wsj|barrons)\.com/video/(?:[^/]+/)+|wsj:)(?:[a-fA-F0-9-]{36}))` +
			`|(?:(?i)https?://(?:www\.)?wsj\.com/articles/(?:[^/?#&]+))|(?:https?://(?:.+?\.)` +
			`?xhamster\.com/(?:movies/(?:\d+)/(?:[^/]*)\.html|videos/(?:[^/]*)-(?:\d+)))|(?:h` +
			`ttps?://(?:.+?\.)?xhamster\.com/xembed\.php\?video=(?:\d+))|(?:https?://(?:video` +
			`|www)\.xnxx\.com/video-?(?:[0-9a-z]+)/)|(?:(?:xtube:|https?://(?:www\.)?xtube\.c` +
			`om/(?:watch\.php\?.*\bv=|video-watch/(?:embedded/)?(?:[^/]+)-))(?:[^/?&#]+))|(?:` +
			`https?://(?:(?:www\.)?xvideos\.com/video|flashservice\.xvideos\.com/embedframe/|` +
			`static-hw\.xvideos\.com/swf/xv-player\.swf\?.*?\bid_video=)(?:[0-9]+))|(?:https?` +
			`://(?:www\.)?xxxymovies\.com/videos/(?:\d+)/(?:[^/]+))|(?:(?:https?://(?:(?:[a-z` +
			`A-Z]{2})\.)?[\da-zA-Z_-]+\.yahoo\.com)/(?:[^/]+/)*(?:(?:.+)?-)?(?:[0-9]+)(?:-[a-` +
			`z]+)?(?:\.html)?)|(?:https?://v\.yinyuetai\.com/video(?:/h5)?/(?:[0-9]+))|(?:htt` +
			`ps?://(?:\w+\.)?youjizz\.com/videos/(?:[^/#?]*-(?:\d+)\.html|embed/(?:\d+)))|(?:` +
			`https?://(?:www\.)?youporn\.com/watch/(?:\d+)/(?:[^/?#&]+))|(?:https?://(?:www\.` +
			`)?yourporn\.sexy/post/(?:[^/?#&.]+))|(?:https?://(?:www\.)?(?:yourupload\.com/(?` +
			`:watch|embed)|embed\.yourupload\.com)/(?:[A-Za-z0-9]+))|(?:^((?:https?://|//)(?:` +
			`(?:(?:(?:\w+\.)?[yY][oO][uU][tT][uU][bB][eE](?:-nocookie)?\.com/|(?:www\.)?detur` +
			`l\.com/www\.youtube\.com/|(?:www\.)?pwnyoutube\.com/|(?:www\.)?hooktube\.com/|(?` +
			`:www\.)?yourepeat\.com/|tube\.majestyc\.net/|(?:(?:www|dev)\.)?invidio\.us/|(?:w` +
			`ww\.)?invidiou\.sh/|(?:www\.)?invidious\.snopyta\.org/|(?:www\.)?invidious\.kabi` +
			`\.tk/|(?:www\.)?vid\.wxzm\.sx/|youtube\.googleapis\.com/)(?:.*?\#/)?(?:(?:(?:v|e` +
			`mbed|e)/(?!videoseries))|(?:(?:(?:watch|movie)(?:_popup)?(?:\.php)?/?)?(?:\?|\#!` +
			`?)(?:.*?[&;])??v=)))|(?:youtu\.be|vid\.plus|zwearz\.com/watch|)/|(?:www\.)?clean` +
			`videosearch\.com/media/action/yt/watch\?videoId=))?([0-9A-Za-z_-]{11})(?!.*?\bli` +
			`st=(?:(?:PL|LL|EC|UU|FL|RD|UL|TL|OLAK5uy_)[0-9A-Za-z-_]{10,}|WL)))|(?:(?:(?:http` +
			`s?://)?(?:\w+\.)?(?:(?:youtube\.com|invidio\.us)/(?:(?:course|view_play_list|my_` +
			`playlists|artist|playlist|watch|embed/(?:videoseries|[0-9A-Za-z_-]{11}))\?(?:.*?` +
			`[&;])*?(?:p|a|list)=|p/)|youtu\.be/[0-9A-Za-z_-]{11}\?.*?\blist=)((?:PL|LL|EC|UU` +
			`|FL|RD|UL|TL|OLAK5uy_)?[0-9A-Za-z-_]{10,}|(?:MC)[\w\.]*).*|((?:PL|LL|EC|UU|FL|RD` +
			`|UL|TL|OLAK5uy_)[0-9A-Za-z-_]{10,})))`,
	}
}
