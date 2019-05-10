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
 * browser.go: Browser interface through downloader
 */

package browser

import (
	"net/http"

	"github.com/tenta-browser/go-video-downloader/lib"

	netlib "github.com/tenta-browser/go-video-downloader/lib/net"
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

// Browser represents an interface towards a real browser,
// which can load pages evaluating JavaScript.
type Browser interface {
	Load(url, ua, inCookies string) (html string, outCookies string, err error)
}

// BrowserLoad loads the url in the browser, keeping the cookie jar in sync.
var BrowserLoad = rnt.NewSimpleFunction("BrowserLoad",
	[]string{"jar", "browser", "url", "ua"},
	func(args []rnt.Object) rnt.Object {
		jar := rnt.UnwrapNative(args[0]).(http.CookieJar)
		browser := rnt.UnwrapNative(args[1]).(Browser)
		urls := args[2].(rnt.Str).Value()
		ua := args[3].(rnt.Str).Value()

		// collect cookies from jar
		inCookies, err := netlib.GetCookieString(jar, urls)
		if err != nil {
			panic(rnt.RaiseType(lib.ExtractorErrorType, err.Error()))
		}

		// run url through browser
		html, outCookies, err := browser.Load(urls, ua, inCookies)
		if err != nil {
			panic(rnt.RaiseType(lib.ExtractorErrorType, err.Error()))
		}

		// put back resulting cookies into the jar
		err = netlib.SetCookies(jar, urls, outCookies)
		if err != nil {
			panic(rnt.RaiseType(lib.ExtractorErrorType, err.Error()))
		}

		return rnt.NewStr(html)
	})
