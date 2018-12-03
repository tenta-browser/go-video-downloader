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
 * util.go: Miscellaneous utilities
 */

package runtime

import (
	"hash"
	"hash/fnv"
	"strconv"
)

const arch32 = strconv.IntSize == 32

func plural(amount int, singular string, plural string) string {
	if amount == 1 {
		return singular
	}
	return plural
}

func makeCompareFn(fn func(compareOp, Object, Object) Object, op compareOp) binaryOpFn {
	return func(a, b Object) Object {
		return fn(op, a, b)
	}
}

func makeBinaryOpFn(fn func(binaryOpType, Object, Object, bool) Object, opt binaryOpType, refl bool) binaryOpFn {
	return func(a, b Object) Object {
		return fn(opt, a, b, refl)
	}
}

func newHash() hash.Hash {
	if arch32 {
		return fnv.New32a()
	}
	return fnv.New64a()
}

func sumHash(h hash.Hash) int {
	if arch32 {
		return int(h.(hash.Hash32).Sum32())
	}
	return int(h.(hash.Hash64).Sum64())
}
