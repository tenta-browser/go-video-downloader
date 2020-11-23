// Code generated by transpiler. DO NOT EDIT.
// +build restricted

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
 * youtube/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/youtube.py
 */

package youtube

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError           λ.Object
	InfoExtractor            λ.Object
	YoutubeBaseInfoExtractor λ.Object
	YoutubeIE                λ.Object
	YoutubePlaylistIE        λ.Object
	YoutubeTabIE             λ.Object
	YoutubeYtUserIE          λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ExtractorError = Ωutils.ExtractorError
		YoutubeBaseInfoExtractor = λ.Cal(λ.TypeType, λ.StrLiteral("YoutubeBaseInfoExtractor"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				YoutubeBaseInfoExtractor__PLAYLIST_ID_RE λ.Object
				YoutubeBaseInfoExtractor__real_extract   λ.Object
			)
			YoutubeBaseInfoExtractor__PLAYLIST_ID_RE = λ.StrLiteral("(?:(?:PL|LL|EC|UU|FL|RD|UL|TL|PU|OLAK5uy_)[0-9A-Za-z-_]{10,}|RDMM)")
			YoutubeBaseInfoExtractor__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "_"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒ_    = λargs[1]
						ϒself = λargs[0]
					)
					_ = ϒ_
					_ = ϒself
					panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.StrLiteral("YOUTUBE_PLACEHOLDER")), λ.KWArgs{
						{Name: "expected", Value: λ.True},
					})))
					return λ.None
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_PLAYLIST_ID_RE": YoutubeBaseInfoExtractor__PLAYLIST_ID_RE,
				"_real_extract":   YoutubeBaseInfoExtractor__real_extract,
			})
		}())
		YoutubeIE = λ.Cal(λ.TypeType, λ.StrLiteral("YoutubeIE"), λ.NewTuple(YoutubeBaseInfoExtractor), func() λ.Dict {
			var (
				YoutubeIE__GEO_BYPASS λ.Object
				YoutubeIE__VALID_URL  λ.Object
			)
			YoutubeIE__VALID_URL = λ.Mod(λ.StrLiteral("(?x)^\n                     (\n                         (?:https?://|//)                                    # http(s):// or protocol-independent URL\n                         (?:(?:(?:(?:\\w+\\.)?[yY][oO][uU][tT][uU][bB][eE](?:-nocookie|kids)?\\.com/|\n                            (?:www\\.)?deturl\\.com/www\\.youtube\\.com/|\n                            (?:www\\.)?pwnyoutube\\.com/|\n                            (?:www\\.)?hooktube\\.com/|\n                            (?:www\\.)?yourepeat\\.com/|\n                            tube\\.majestyc\\.net/|\n                            # Invidious instances taken from https://github.com/omarroth/invidious/wiki/Invidious-Instances\n                            (?:(?:www|dev)\\.)?invidio\\.us/|\n                            (?:(?:www|no)\\.)?invidiou\\.sh/|\n                            (?:(?:www|fi|de)\\.)?invidious\\.snopyta\\.org/|\n                            (?:www\\.)?invidious\\.kabi\\.tk/|\n                            (?:www\\.)?invidious\\.13ad\\.de/|\n                            (?:www\\.)?invidious\\.mastodon\\.host/|\n                            (?:www\\.)?invidious\\.nixnet\\.xyz/|\n                            (?:www\\.)?invidious\\.drycat\\.fr/|\n                            (?:www\\.)?tube\\.poal\\.co/|\n                            (?:www\\.)?vid\\.wxzm\\.sx/|\n                            (?:www\\.)?yewtu\\.be/|\n                            (?:www\\.)?yt\\.elukerio\\.org/|\n                            (?:www\\.)?yt\\.lelux\\.fi/|\n                            (?:www\\.)?invidious\\.ggc-project\\.de/|\n                            (?:www\\.)?yt\\.maisputain\\.ovh/|\n                            (?:www\\.)?invidious\\.13ad\\.de/|\n                            (?:www\\.)?invidious\\.toot\\.koeln/|\n                            (?:www\\.)?invidious\\.fdn\\.fr/|\n                            (?:www\\.)?watch\\.nettohikari\\.com/|\n                            (?:www\\.)?kgg2m7yk5aybusll\\.onion/|\n                            (?:www\\.)?qklhadlycap4cnod\\.onion/|\n                            (?:www\\.)?axqzx4s6s54s32yentfqojs3x5i7faxza6xo3ehd4bzzsg2ii4fv2iid\\.onion/|\n                            (?:www\\.)?c7hqkpkpemu6e7emz5b4vyz7idjgdvgaaa3dyimmeojqbgpea3xqjoid\\.onion/|\n                            (?:www\\.)?fz253lmuao3strwbfbmx46yu7acac2jz27iwtorgmbqlkurlclmancad\\.onion/|\n                            (?:www\\.)?invidious\\.l4qlywnpwqsluw65ts7md3khrivpirse744un3x7mlskqauz5pyuzgqd\\.onion/|\n                            (?:www\\.)?owxfohz4kjyv25fvlqilyxast7inivgiktls3th44jhk3ej3i7ya\\.b32\\.i2p/|\n                            (?:www\\.)?4l2dgddgsrkf2ous66i6seeyi6etzfgrue332grh2n7madpwopotugyd\\.onion/|\n                            youtube\\.googleapis\\.com/)                        # the various hostnames, with wildcard subdomains\n                         (?:.*?\\#/)?                                          # handle anchor (#/) redirect urls\n                         (?:                                                  # the various things that can precede the ID:\n                             (?:(?:v|embed|e)/(?!videoseries))                # v/ or embed/ or e/\n                             |(?:                                             # or the v= param in all its forms\n                                 (?:(?:watch|movie)(?:_popup)?(?:\\.php)?/?)?  # preceding watch(_popup|.php) or nothing (like /?v=xxxx)\n                                 (?:\\?|\\#!?)                                  # the params delimiter ? or # or #!\n                                 # any other preceding param (like /?s=tuff&v=xxxx or ?s=tuff&amp;v=V36LpHqtcDY)\n                                 (?:.*?[&;])??\n                                 v=\n                             )\n                         ))\n                         |(?:\n                            youtu\\.be|                                        # just youtu.be/xxxx\n                            vid\\.plus|                                        # or vid.plus/xxxx\n                            zwearz\\.com/watch|                                # or zwearz.com/watch/xxxx\n                         )/\n                         |(?:www\\.)?cleanvideosearch\\.com/media/action/yt/watch\\?videoId=\n                         )\n                     )?                                                       # all until now is optional -> you can pass the naked ID\n                     ([0-9A-Za-z_-]{11})                                      # here is it! the YouTube video ID\n                     (?!.*?\\blist=\n                        (?:\n                            %(playlist_id)s|                                  # combined list/video URLs are handled by the playlist IE\n                            WL                                                # WL are handled by the watch later IE\n                        )\n                     )\n                     #(?(1).+)?                                                # if we found the ID, everything can follow\n                     "), λ.DictLiteral(map[string]λ.Object{
				"playlist_id": λ.GetAttr(YoutubeBaseInfoExtractor, "_PLAYLIST_ID_RE", nil),
			}))
			YoutubeIE__GEO_BYPASS = λ.False
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_GEO_BYPASS": YoutubeIE__GEO_BYPASS,
				"_VALID_URL":  YoutubeIE__VALID_URL,
			})
		}())
		YoutubeTabIE = λ.Cal(λ.TypeType, λ.StrLiteral("YoutubeTabIE"), λ.NewTuple(YoutubeBaseInfoExtractor), func() λ.Dict {
			var (
				YoutubeTabIE__VALID_URL λ.Object
			)
			YoutubeTabIE__VALID_URL = λ.StrLiteral("https?://(?:\\w+\\.)?(?:youtube(?:kids)?\\.com|invidio\\.us)/(?:(?:channel|c|user)/|(?:playlist|watch)\\?.*?\\blist=)(?P<id>[^/?#&]+)")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": YoutubeTabIE__VALID_URL,
			})
		}())
		YoutubePlaylistIE = λ.Cal(λ.TypeType, λ.StrLiteral("YoutubePlaylistIE"), λ.NewTuple(YoutubeBaseInfoExtractor), func() λ.Dict {
			var (
				YoutubePlaylistIE__VALID_URL    λ.Object
				YoutubePlaylistIE__real_extract λ.Object
			)
			YoutubePlaylistIE__VALID_URL = λ.Mod(λ.StrLiteral("(?x)(?:\n                        (?:https?://)?\n                        (?:\\w+\\.)?\n                        (?:\n                            (?:\n                                youtube(?:kids)?\\.com|\n                                invidio\\.us|\n                                youtu\\.be\n                            )\n                            /.*?\\?.*?\\blist=\n                        )?\n                        (?P<id>%(playlist_id)s)\n                     )"), λ.DictLiteral(map[string]λ.Object{
				"playlist_id": λ.GetAttr(YoutubeBaseInfoExtractor, "_PLAYLIST_ID_RE", nil),
			}))
			YoutubePlaylistIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "_"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒ_    = λargs[1]
						ϒself = λargs[0]
					)
					_ = ϒ_
					_ = ϒself
					panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.StrLiteral("YOUTUBE_PLACEHOLDER")), λ.KWArgs{
						{Name: "expected", Value: λ.True},
					})))
					return λ.None
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL":    YoutubePlaylistIE__VALID_URL,
				"_real_extract": YoutubePlaylistIE__real_extract,
			})
		}())
		YoutubeYtUserIE = λ.Cal(λ.TypeType, λ.StrLiteral("YoutubeYtUserIE"), λ.NewTuple(YoutubeBaseInfoExtractor), func() λ.Dict {
			var (
				YoutubeYtUserIE__VALID_URL λ.Object
			)
			YoutubeYtUserIE__VALID_URL = λ.StrLiteral("ytuser:(?P<id>.+)")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": YoutubeYtUserIE__VALID_URL,
			})
		}())
	})
}
