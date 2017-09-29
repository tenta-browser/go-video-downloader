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
 * builtin-codecs.go: Re-implementation of Python's string/bytes encoding/decoding functions
 */

package runtime

import (
	"encoding/base64"
	"encoding/hex"
)

// Codecs is dummy
const Codecs = 0

var b64 = base64.StdEncoding

// Base64Encode implements python/base64.b64encode
func Base64Encode(codecs int, bytes []byte) []byte {
	result := make([]byte, b64.EncodedLen(len(bytes)))
	b64.Encode(result, bytes)
	return result
}

// Base64DecodeBytes implements python/base64.b64decode(bytes)
func Base64DecodeBytes(codecs int, bytes []byte) []byte {
	result := make([]byte, b64.DecodedLen(len(bytes)))
	n, err := b64.Decode(result, bytes)
	if err != nil {
		panic(newExtractorError("Invalid base64 input"))
	}
	return result[:n]
}

// Base64DecodeString implements python/base64.b64decode(str)
func Base64DecodeString(codecs int, str string) []byte {
	return Base64DecodeBytes(codecs, []byte(str))
}

// HexEncode implements python/binascii.hexlify
func HexEncode(codecs int, bytes []byte) []byte {
	result := make([]byte, hex.EncodedLen(len(bytes)))
	hex.Encode(result, bytes)
	return result
}

// HexDecode implements python/binascii.unhexlify
func HexDecode(codecs int, bytes []byte) []byte {
	result := make([]byte, hex.DecodedLen(len(bytes)))
	_, err := hex.Decode(result, bytes)
	if err != nil {
		panic(newExtractorError("Invalid hex input"))
	}
	return result
}
