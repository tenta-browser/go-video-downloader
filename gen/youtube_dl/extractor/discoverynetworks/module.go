// Code generated by transpiler. DO NOT EDIT.

/**
 * Go Video Downloader
 *
 *    Copyright 2018 Tenta, LLC
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
 * discoverynetworks/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/discoverynetworks.py
 */

package discoverynetworks

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωparse "github.com/tenta-browser/go-video-downloader/gen/urllib/parse"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωbrightcove "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/brightcove"
	Ωdplay "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/dplay"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	BrightcoveLegacyIE    λ.Object
	DPlayIE               λ.Object
	DiscoveryNetworksDeIE λ.Object
	ϒcompat_parse_qs      λ.Object
	ϒsmuggle_url          λ.Object
)

func init() {
	λ.InitModule(func() {
		BrightcoveLegacyIE = Ωbrightcove.BrightcoveLegacyIE
		DPlayIE = Ωdplay.DPlayIE
		ϒcompat_parse_qs = Ωcompat.ϒcompat_parse_qs
		ϒsmuggle_url = Ωutils.ϒsmuggle_url
		DiscoveryNetworksDeIE = λ.Cal(λ.TypeType, λ.NewStr("DiscoveryNetworksDeIE"), λ.NewTuple(DPlayIE), func() λ.Dict {
			var (
				DiscoveryNetworksDeIE_BRIGHTCOVE_URL_TEMPLATE λ.Object
				DiscoveryNetworksDeIE__TESTS                  λ.Object
				DiscoveryNetworksDeIE__VALID_URL              λ.Object
				DiscoveryNetworksDeIE__real_extract           λ.Object
			)
			DiscoveryNetworksDeIE__VALID_URL = λ.NewStr("(?x)https?://(?:www\\.)?(?P<site>discovery|tlc|animalplanet|dmax)\\.de/\n                        (?:\n                           .*\\#(?P<id>\\d+)|\n                           (?:[^/]+/)*videos/(?P<display_id>[^/?#]+)|\n                           programme/(?P<programme>[^/]+)/video/(?P<alternate_id>[^/]+)\n                        )")
			DiscoveryNetworksDeIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.tlc.de/sendungen/breaking-amish/videos/#3235167922001"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("3235167922001"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Breaking Amish: Die Welt da draußen"),
						λ.NewStr("description"): λ.NewStr("Vier Amische und eine Mennonitin wagen in New York  den Sprung in ein komplett anderes Leben. Begleitet sie auf ihrem spannenden Weg."),
						λ.NewStr("timestamp"):   λ.NewInt(1396598084),
						λ.NewStr("upload_date"): λ.NewStr("20140404"),
						λ.NewStr("uploader_id"): λ.NewStr("1659832546"),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://www.dmax.de/programme/storage-hunters-uk/videos/storage-hunters-uk-episode-6/"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://www.discovery.de/#5332316765001"),
					λ.NewStr("only_matching"): λ.True,
				}),
			)
			DiscoveryNetworksDeIE_BRIGHTCOVE_URL_TEMPLATE = λ.NewStr("http://players.brightcove.net/1659832546/default_default/index.html?videoId=%s")
			DiscoveryNetworksDeIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒalternate_id          λ.Object
						ϒbrightcove_id         λ.Object
						ϒbrightcove_legacy_url λ.Object
						ϒmobj                  λ.Object
						ϒself                  = λargs[0]
						ϒtitle                 λ.Object
						ϒurl                   = λargs[1]
						ϒwebpage               λ.Object
					)
					ϒmobj = λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl)
					ϒalternate_id = λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("alternate_id"))
					if λ.IsTrue(ϒalternate_id) {
						λ.Cal(λ.GetAttr(ϒself, "_initialize_geo_bypass", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("countries"): λ.NewList(λ.NewStr("DE")),
						}))
						return λ.Cal(λ.GetAttr(ϒself, "_get_disco_api_info", nil), ϒurl, λ.Mod(λ.NewStr("%s/%s"), λ.NewTuple(
							λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("programme")),
							ϒalternate_id,
						)), λ.NewStr("sonic-eu1-prod.disco-api.com"), λ.Add(λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("site")), λ.NewStr("de")))
					}
					ϒbrightcove_id = λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("id"))
					if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒbrightcove_id))) {
						ϒtitle = λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("title"))
						ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒtitle)
						ϒbrightcove_legacy_url = λ.Cal(λ.GetAttr(BrightcoveLegacyIE, "_extract_brightcove_url", nil), ϒwebpage)
						ϒbrightcove_id = λ.GetItem(λ.GetItem(λ.Cal(ϒcompat_parse_qs, λ.GetAttr(λ.Cal(Ωparse.ϒurlparse, ϒbrightcove_legacy_url), "query", nil)), λ.NewStr("@videoPlayer")), λ.NewInt(0))
					}
					return λ.Cal(λ.GetAttr(ϒself, "url_result", nil), λ.Cal(ϒsmuggle_url, λ.Mod(λ.GetAttr(ϒself, "BRIGHTCOVE_URL_TEMPLATE", nil), ϒbrightcove_id), λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("geo_countries"): λ.NewList(λ.NewStr("DE")),
					})), λ.NewStr("BrightcoveNew"), ϒbrightcove_id)
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("BRIGHTCOVE_URL_TEMPLATE"): DiscoveryNetworksDeIE_BRIGHTCOVE_URL_TEMPLATE,
				λ.NewStr("_TESTS"):                  DiscoveryNetworksDeIE__TESTS,
				λ.NewStr("_VALID_URL"):              DiscoveryNetworksDeIE__VALID_URL,
				λ.NewStr("_real_extract"):           DiscoveryNetworksDeIE__real_extract,
			})
		}())
	})
}
