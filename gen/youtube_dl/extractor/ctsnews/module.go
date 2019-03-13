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
 * ctsnews/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/ctsnews.py
 */

package ctsnews

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	CtsNewsIE          λ.Object
	InfoExtractor      λ.Object
	ϒunified_timestamp λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒunified_timestamp = Ωutils.ϒunified_timestamp
		CtsNewsIE = λ.Cal(λ.TypeType, λ.NewStr("CtsNewsIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				CtsNewsIE__TESTS        λ.Object
				CtsNewsIE__VALID_URL    λ.Object
				CtsNewsIE__real_extract λ.Object
			)
			CtsNewsIE__VALID_URL = λ.NewStr("https?://news\\.cts\\.com\\.tw/[a-z]+/[a-z]+/\\d+/(?P<id>\\d+)\\.html")
			CtsNewsIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://news.cts.com.tw/cts/international/201501/201501291578109.html"),
					λ.NewStr("md5"): λ.NewStr("a9875cb790252b08431186d741beaabe"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("201501291578109"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("以色列.真主黨交火 3人死亡"),
						λ.NewStr("description"): λ.NewStr("以色列和黎巴嫩真主黨，爆發五年最嚴重衝突，雙方砲轟交火，兩名以軍死亡，還有一名西班牙籍的聯合國維和人..."),
						λ.NewStr("timestamp"):   λ.NewInt(1422528540),
						λ.NewStr("upload_date"): λ.NewStr("20150129"),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://news.cts.com.tw/cts/international/201309/201309031304098.html"),
					λ.NewStr("md5"): λ.NewStr("3aee7e0df7cdff94e43581f54c22619e"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("201309031304098"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("韓國31歲童顏男 貌如十多歲小孩"),
						λ.NewStr("description"): λ.NewStr("越有年紀的人，越希望看起來年輕一點，而南韓卻有一位31歲的男子，看起來像是11、12歲的小孩，身..."),
						λ.NewStr("thumbnail"):   λ.NewStr("re:^https?://.*\\.jpg$"),
						λ.NewStr("timestamp"):   λ.NewInt(1378205880),
						λ.NewStr("upload_date"): λ.NewStr("20130903"),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://news.cts.com.tw/cts/money/201501/201501291578003.html"),
					λ.NewStr("md5"): λ.NewStr("e4726b2ccd70ba2c319865e28f0a91d1"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("OVbfO7d0_hQ"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("iPhone6熱銷 蘋果財報亮眼"),
						λ.NewStr("description"): λ.NewStr("md5:f395d4f485487bb0f992ed2c4b07aa7d"),
						λ.NewStr("thumbnail"):   λ.NewStr("re:^https?://.*\\.jpg$"),
						λ.NewStr("upload_date"): λ.NewStr("20150128"),
						λ.NewStr("uploader_id"): λ.NewStr("TBSCTS"),
						λ.NewStr("uploader"):    λ.NewStr("中華電視公司"),
					}),
					λ.NewStr("add_ie"): λ.NewList(λ.NewStr("Youtube")),
				}),
			)
			CtsNewsIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdatetime_str λ.Object
						ϒdescription  λ.Object
						ϒmp4_feed     λ.Object
						ϒnews_id      λ.Object
						ϒpage         λ.Object
						ϒself         = λargs[0]
						ϒthumbnail    λ.Object
						ϒtimestamp    λ.Object
						ϒtitle        λ.Object
						ϒurl          = λargs[1]
						ϒvideo_url    λ.Object
						ϒyoutube_url  λ.Object
					)
					ϒnews_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒnews_id)
					ϒnews_id = λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒself, "_hidden_inputs", nil), ϒpage), "get", nil), λ.NewStr("get_id"))
					if λ.IsTrue(ϒnews_id) {
						ϒmp4_feed = λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
							λ.NewStr("http://news.cts.com.tw/action/test_mp4feed.php"),
							ϒnews_id,
						), λ.KWArgs{
							{Name: "note", Value: λ.NewStr("Fetching feed")},
							{Name: "query", Value: λ.NewDictWithTable(map[λ.Object]λ.Object{
								λ.NewStr("news_id"): ϒnews_id,
							})},
						})
						ϒvideo_url = λ.GetItem(ϒmp4_feed, λ.NewStr("source_url"))
					} else {
						λ.Cal(λ.GetAttr(ϒself, "to_screen", nil), λ.NewStr("Not CTSPlayer video, trying Youtube..."))
						ϒyoutube_url = λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewStr("src=\"(//www\\.youtube\\.com/embed/[^\"]+)\""), ϒpage, λ.NewStr("youtube url"))
						return λ.Call(λ.GetAttr(ϒself, "url_result", nil), λ.NewArgs(ϒyoutube_url), λ.KWArgs{
							{Name: "ie", Value: λ.NewStr("Youtube")},
						})
					}
					ϒdescription = λ.Cal(λ.GetAttr(ϒself, "_html_search_meta", nil), λ.NewStr("description"), ϒpage)
					ϒtitle = λ.Call(λ.GetAttr(ϒself, "_html_search_meta", nil), λ.NewArgs(
						λ.NewStr("title"),
						ϒpage,
					), λ.KWArgs{
						{Name: "fatal", Value: λ.True},
					})
					ϒthumbnail = λ.Cal(λ.GetAttr(ϒself, "_html_search_meta", nil), λ.NewStr("image"), ϒpage)
					ϒdatetime_str = λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
						λ.NewStr("(\\d{4}/\\d{2}/\\d{2} \\d{2}:\\d{2})"),
						ϒpage,
						λ.NewStr("date and time"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					})
					ϒtimestamp = λ.None
					if λ.IsTrue(ϒdatetime_str) {
						ϒtimestamp = λ.Sub(λ.Cal(ϒunified_timestamp, ϒdatetime_str), λ.Mul(λ.NewInt(8), λ.NewInt(3600)))
					}
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          ϒnews_id,
						λ.NewStr("url"):         ϒvideo_url,
						λ.NewStr("title"):       ϒtitle,
						λ.NewStr("description"): ϒdescription,
						λ.NewStr("thumbnail"):   ϒthumbnail,
						λ.NewStr("timestamp"):   ϒtimestamp,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TESTS"):        CtsNewsIE__TESTS,
				λ.NewStr("_VALID_URL"):    CtsNewsIE__VALID_URL,
				λ.NewStr("_real_extract"): CtsNewsIE__real_extract,
			})
		}())
	})
}
