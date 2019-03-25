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
 * codecs.go: Provides interface to the unicode codecs of the runtime
 */

package encoding

import (
	"strings"

	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

// GetEncoder returns the encoder function of the codec corresponding to encoding.
var GetEncoder = rnt.NewSimpleFunction("getencoder",
	[]string{"encoding"},
	func(args []rnt.Object) rnt.Object {
		encoding := args[0].(rnt.Str).Value()
		encoder := rnt.GetCodec(encoding).Encode
		return rnt.NewFunction(namifyEncoding(encoding)+"_encode", []rnt.Param{
			{Name: "obj"},
			{Name: "errors", Def: rnt.NewStr("strict")},
		}, 0, false, false, func(args []rnt.Object) rnt.Object {
			res, cnt := encoder(args[0], args[1].(rnt.Str))
			return rnt.NewTuple(res, cnt)
		})
	})

// GetDecoder returns the decoder function of the codec corresponding to encoding.
var GetDecoder = rnt.NewSimpleFunction("getdecoder",
	[]string{"encoding"},
	func(args []rnt.Object) rnt.Object {
		encoding := args[0].(rnt.Str).Value()
		decoder := rnt.GetCodec(encoding).Decode
		return rnt.NewFunction(namifyEncoding(encoding)+"_decode", []rnt.Param{
			{Name: "s"},
			{Name: "errors", Def: rnt.NewStr("strict")},
		}, 0, false, false, func(args []rnt.Object) rnt.Object {
			res, cnt := decoder(args[0], args[1].(rnt.Str))
			return rnt.NewTuple(res, cnt)
		})
	})

func namifyEncoding(encoding string) string {
	return strings.Replace(encoding, "-", "_", -1)
}
