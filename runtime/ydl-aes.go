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
 * ydl-aes.go: Re-implementation of various functions from
 *             https://github.com/rg3/youtube-dl/blob/master/youtube_dl/aes.py
 */

package runtime

import (
	"crypto/aes"
	"crypto/cipher"
)

const (
	nonceByteSize  = 8
	blockSizeBytes = aes.BlockSize
)

// AESDecryptText implements aes.py/aes_decrypt_text
func AESDecryptText(data, password string, keySizeBytes int) []byte {
	dataBytes := Base64StdDecodeString(data) // data = nonce(8) + cipher
	nonceBytes := append([]byte{}, dataBytes[:nonceByteSize]...)
	cipherBytes := dataBytes[nonceByteSize:]

	// pad up nonce (IV) to blockSizeBytes
	nonceBytes = pad(nonceBytes, blockSizeBytes)

	// keyBytes should be the first keySizeBytes bytes of password (eventually padded with 0s)
	keyBytes := []byte(password)
	if len(password) > keySizeBytes {
		keyBytes = keyBytes[:keySizeBytes]
	} else {
		keyBytes = pad(keyBytes, keySizeBytes)
	}

	// encrypt the first blockSizeBytes bytes of the password with keyBytes
	cipherKeyBytes := make([]byte, blockSizeBytes)
	keyBlock, err := aes.NewCipher(keyBytes)
	if err != nil {
		panic(err.Error())
	}
	keyBlock.Encrypt(cipherKeyBytes, keyBytes[:blockSizeBytes])

	// multiplicate key up so that it gets keySizeBytes long
	shortCipherKeyBytes := cipherKeyBytes
	for i := 0; i < keySizeBytes/blockSizeBytes-1; i++ {
		cipherKeyBytes = append(cipherKeyBytes, shortCipherKeyBytes...)
	}

	// decrypt cipherBytes using CTR mode
	plainBytes := make([]byte, len(cipherBytes))
	block, err := aes.NewCipher(cipherKeyBytes)
	if err != nil {
		panic(err.Error())
	}
	cipher.NewCTR(block, nonceBytes).XORKeyStream(plainBytes, cipherBytes)

	return plainBytes
}

func pad(bytes []byte, n int) []byte {
	diff := n - len(bytes)
	for i := 0; i < diff; i++ {
		bytes = append(bytes, 0)
	}
	return bytes
}
