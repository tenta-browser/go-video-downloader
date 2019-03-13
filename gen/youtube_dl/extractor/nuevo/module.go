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
 * nuevo/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/nuevo.py
 */

package nuevo

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor  λ.Object
	NuevoBaseIE    λ.Object
	ϒfloat_or_none λ.Object
	ϒxpath_text    λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒfloat_or_none = Ωutils.ϒfloat_or_none
		ϒxpath_text = Ωutils.ϒxpath_text
		NuevoBaseIE = λ.Cal(λ.TypeType, λ.NewStr("NuevoBaseIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				NuevoBaseIE__extract_nuevo λ.Object
			)
			NuevoBaseIE__extract_nuevo = λ.NewFunction("_extract_nuevo",
				[]λ.Param{
					{Name: "self"},
					{Name: "config_url"},
					{Name: "video_id"},
					{Name: "headers", Def: λ.NewDictWithTable(map[λ.Object]λ.Object{})},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒconfig       λ.Object
						ϒconfig_url   = λargs[1]
						ϒduration     λ.Object
						ϒelement_name λ.Object
						ϒformat_id    λ.Object
						ϒformats      λ.Object
						ϒheaders      = λargs[3]
						ϒself         = λargs[0]
						ϒthumbnail    λ.Object
						ϒtitle        λ.Object
						ϒvideo_id     = λargs[2]
						ϒvideo_url    λ.Object
						τmp0          λ.Object
						τmp1          λ.Object
						τmp2          λ.Object
					)
					ϒconfig = λ.Call(λ.GetAttr(ϒself, "_download_xml", nil), λ.NewArgs(
						ϒconfig_url,
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "transform_source", Value: λ.NewFunction("<lambda>",
							[]λ.Param{
								{Name: "s"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒs = λargs[0]
								)
								return λ.Cal(λ.GetAttr(ϒs, "strip", nil))
							})},
						{Name: "headers", Value: ϒheaders},
					})
					ϒtitle = λ.Cal(λ.GetAttr(λ.Call(ϒxpath_text, λ.NewArgs(
						ϒconfig,
						λ.NewStr("./title"),
						λ.NewStr("title"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.True},
					}), "strip", nil))
					ϒvideo_id = λ.Call(ϒxpath_text, λ.NewArgs(
						ϒconfig,
						λ.NewStr("./mediaid"),
					), λ.KWArgs{
						{Name: "default", Value: ϒvideo_id},
					})
					ϒthumbnail = λ.Cal(ϒxpath_text, ϒconfig, λ.NewList(
						λ.NewStr("./image"),
						λ.NewStr("./thumb"),
					))
					ϒduration = λ.Cal(ϒfloat_or_none, λ.Cal(ϒxpath_text, ϒconfig, λ.NewStr("./duration")))
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.NewTuple(
						λ.NewTuple(
							λ.NewStr("file"),
							λ.NewStr("sd"),
						),
						λ.NewTuple(
							λ.NewStr("filehd"),
							λ.NewStr("hd"),
						),
					))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						τmp2 = τmp1
						ϒelement_name = λ.GetItem(τmp2, λ.NewInt(0))
						ϒformat_id = λ.GetItem(τmp2, λ.NewInt(1))
						ϒvideo_url = λ.Cal(ϒxpath_text, ϒconfig, ϒelement_name)
						if λ.IsTrue(ϒvideo_url) {
							λ.Cal(λ.GetAttr(ϒformats, "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
								λ.NewStr("url"):       ϒvideo_url,
								λ.NewStr("format_id"): ϒformat_id,
							}))
						}
					}
					λ.Cal(λ.GetAttr(ϒself, "_check_formats", nil), ϒformats, ϒvideo_id)
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):        ϒvideo_id,
						λ.NewStr("title"):     ϒtitle,
						λ.NewStr("thumbnail"): ϒthumbnail,
						λ.NewStr("duration"):  ϒduration,
						λ.NewStr("formats"):   ϒformats,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_extract_nuevo"): NuevoBaseIE__extract_nuevo,
			})
		}())
	})
}
