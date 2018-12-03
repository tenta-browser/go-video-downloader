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
 * extractor-interface.go: Interface methods called by extractors on Python side,
 *                         to make themselves available to the downloader on Go side
 */

package lib

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

// Extractors ..
var Extractors = make(map[string]rnt.Object)

// ExtractorRunner ..
var ExtractorRunner rnt.Object

// RegisterExtractor registers a new extractor class
var RegisterExtractor = rnt.NewSimpleFunction("RegisterExtractor",
	[]string{"key", "class"},
	func(args []rnt.Object) rnt.Object {
		key := args[0].(rnt.Str).Value()
		class := args[1]
		Extractors[key] = class
		return rnt.None
	})

// SetExtractorRunner is called by the generated code to provide an entry
// point into it, usable by the fixe code.
var SetExtractorRunner = rnt.NewSimpleFunction("SetExtractorRunner",
	[]string{"runner"},
	func(args []rnt.Object) rnt.Object {
		ExtractorRunner = args[0]
		return rnt.None
	})
