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
 * builtin-url.go: Re-implementation of Python's url manipulation functions (urllib.parse)
 */

package runtime

import "net/url"
import "fmt"

// URLParse impements python/urllib.parse.urlparse
func URLParse(urlString OptString) *url.URL {
	_urlString := urlString.GetOrDef("")
	url, err := url.Parse(urlString.GetOrDef(""))
	if err != nil {
		panic(newExtractorError("Failed to parse: " + _urlString))
	}
	return url
}

// URLJoin implements python/urllib.parse.urljoin
func URLJoin(base string, urlString OptString) string {
	baseURL, err := url.Parse(base)
	if err != nil {
		panic(newExtractorError("Failed to parse: " + base))
	}
	_urlString := urlString.GetOrDef("")
	resultURL, err := baseURL.Parse(_urlString)
	if err != nil {
		panic(newExtractorError(fmt.Sprintf("Failed to join %s with %s", base, _urlString)))
	}
	return resultURL.String()
}

// ParseUnquote implements python/urllib.parse.unquote
func ParseUnquote(str string) string {
	ret, err := url.PathUnescape(str)
	if err != nil {
		panic(newExtractorError(err.Error()))
	}
	return ret
}

// ParseUnquotePlus implements python/urllib.parse.unquote_plus
func ParseUnquotePlus(str string) string {
	ret, err := url.QueryUnescape(str)
	if err != nil {
		panic(newExtractorError(err.Error()))
	}
	return ret
}
