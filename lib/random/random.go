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
 * random.go: Re-implementation of Python's random functions (random)
 */

package random

import (
	"math/rand"
	"time"

	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

// RandInt returns a random number in [0,n).
var RandInt = rnt.NewSimpleFunction("RandInt",
	[]string{"n"},
	func(args []rnt.Object) rnt.Object {
		n := args[0].(rnt.Int).Value()
		return rnt.NewInt(rand.Intn(n))
	})

// RandFloat returns a random number in [0,1).
var RandFloat = rnt.NewSimpleFunction("RandFloat",
	nil,
	func(args []rnt.Object) rnt.Object {
		return rnt.NewFloat(rand.Float64())
	})

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}
