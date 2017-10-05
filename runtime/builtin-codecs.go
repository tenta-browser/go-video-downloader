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

// Base64StdEncode implements python/base64.b64encode
func Base64StdEncode(codecs int, bytes []byte) []byte {
	return base64Encode(base64.StdEncoding, bytes)
}

// Base64StdDecodeBytes implements python/base64.b64decode(bytes)
func Base64StdDecodeBytes(codecs int, bytes []byte) []byte {
	return base64DecodeBytes(base64.StdEncoding, bytes)
}

// Base64StdDecodeString implements python/base64.b64decode(str)
func Base64StdDecodeString(codecs int, str string) []byte {
	return base64DecodeBytes(base64.StdEncoding, []byte(str))
}

// Base64UrlEncode implements python/base64.urlsafe_b64encode
func Base64UrlEncode(codecs int, bytes []byte) []byte {
	return base64Encode(base64.URLEncoding, bytes)
}

// Base64UrlDecodeBytes implements python/base64.urlsafe_b64decode(bytes)
func Base64UrlDecodeBytes(codecs int, bytes []byte) []byte {
	return base64DecodeBytes(base64.URLEncoding, bytes)
}

// Base64UrlDecodeString implements python/base64.urlsafe_b64decode(str)
func Base64UrlDecodeString(codecs int, str string) []byte {
	return base64DecodeBytes(base64.URLEncoding, []byte(str))
}

func base64Encode(enc *base64.Encoding, bytes []byte) []byte {
	result := make([]byte, enc.EncodedLen(len(bytes)))
	enc.Encode(result, bytes)
	return result
}

func base64DecodeBytes(enc *base64.Encoding, bytes []byte) []byte {
	result := make([]byte, enc.DecodedLen(len(bytes)))
	n, err := enc.Decode(result, bytes)
	if err != nil {
		panic(newExtractorError("Invalid base64 input"))
	}
	return result[:n]
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
