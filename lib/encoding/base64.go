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
 * base64.go: Base64 encoding/decoding methods
 */

package encoding

import (
	"encoding/base64"

	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

// B64Encode implements python/base64.b64encode (mostly)
var B64Encode = rnt.NewSimpleFunction("b64encode",
	[]string{"s", "url"}, func(args []rnt.Object) rnt.Object {
		bytes := args[0].(rnt.Bytes).Value()
		enc := base64.StdEncoding
		if args[1].(rnt.Int).Value() != 0 {
			enc = base64.URLEncoding
		}
		result := make([]byte, enc.EncodedLen(len(bytes)))
		enc.Encode(result, bytes)
		return rnt.NewBytes(result...)
	})

// B64Decode implements python/base64.b64decode (mostly)
var B64Decode = rnt.NewSimpleFunction("b64decode",
	[]string{"s", "url"}, func(args []rnt.Object) rnt.Object {
		var bytes []byte
		if args[0].IsInstance(rnt.BytesType) {
			bytes = args[0].(rnt.Bytes).Value()
		} else {
			bytes = []byte(args[0].(rnt.Str).Value())
		}
		enc := base64.StdEncoding
		if args[1].(rnt.Int).Value() != 0 {
			enc = base64.URLEncoding
		}
		result := make([]byte, enc.DecodedLen(len(bytes)))
		n, err := enc.Decode(result, bytes)
		if err != nil {
			panic(rnt.RaiseType(BinasciiErrorType, "Invalid base64 input"))
		}
		return rnt.NewBytes(result[:n]...)
	})
