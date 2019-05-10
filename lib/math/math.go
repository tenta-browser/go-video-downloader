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
 * math.go: Various math or integer, number related functions
 */

package math

import (
	"math"

	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

// MaxUInt is the bigest unsigned integer value representable on the current
// platform.
var MaxUInt = rnt.NewInt(int(^uint(0) >> 1))

// Mathematical constants.
var (
	E  = rnt.NewFloat(math.E)
	Pi = rnt.NewFloat(math.Pi)
)

// Ceil wraps math.Ceil.
var Ceil = rnt.NewSimpleFunction("Ceil",
	[]string{"f"},
	func(args []rnt.Object) rnt.Object {
		f := args[0].(rnt.Float).Value()
		return rnt.NewInt(int(math.Ceil(f)))
	})

// Log wraps math.Log.
var Log = rnt.NewSimpleFunction("Log",
	[]string{"f"},
	func(args []rnt.Object) rnt.Object {
		f := args[0].(rnt.Float).Value()
		l := math.Log(f)
		if math.IsNaN(l) || math.IsInf(l, 0) {
			panic(rnt.RaiseType(rnt.ValueErrorType, "math domain error"))
		}
		return rnt.NewFloat(l)
	})
