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
 * all.go: Registry of all available extractors
 */

package extractor

import (
	"github.com/tenta-browser/go-video-downloader/runtime"
)

var factories = make(map[string]runtime.InfoExtractorFactory)

func registerFactory(key string, factory runtime.InfoExtractorFactory) {
	factories[key] = factory
}

// GetFactories returns factories for every InfoExtractor
func GetFactories() map[string]runtime.InfoExtractorFactory {
	return factories
}
