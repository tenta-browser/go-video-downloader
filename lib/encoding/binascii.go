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
 * binascii.go: Various encoding/decoding methods
 */

package encoding

import (
	"encoding/hex"

	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

// BinasciiErrorType is the error type raised by the binascii functions ..
var BinasciiErrorType = rnt.Cal(
	rnt.TypeType,
	rnt.NewStr("binascii.Error"),
	rnt.NewTuple(rnt.ValueErrorType),
	rnt.NewDict()).(rnt.Type)

// Hexlify implements python/binascii.hexlify
var Hexlify = rnt.NewSimpleFunction("hexlify",
	[]string{"data"}, func(args []rnt.Object) rnt.Object {
		data := args[0].(rnt.Bytes)
		bytes := data.Value()
		result := make([]byte, hex.EncodedLen(len(bytes)))
		hex.Encode(result, bytes)
		return rnt.NewBytes(result...)
	})

// Unhexlify implements python/binascii.unhexlify
var Unhexlify = rnt.NewSimpleFunction("unhexlify",
	[]string{"hexstr"}, func(args []rnt.Object) rnt.Object {
		hexstr := args[0].(rnt.Bytes)
		bytes := hexstr.Value()
		result := make([]byte, hex.DecodedLen(len(bytes)))
		_, err := hex.Decode(result, bytes)
		if err != nil {
			panic(rnt.RaiseType(BinasciiErrorType, "Invalid hex input"))
		}
		return rnt.NewBytes(result...)
	})
