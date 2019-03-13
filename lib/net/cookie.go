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
 * cookie.go: Cookie parsing and cookie jar handling methods
 */

package http

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"sort"
	"strings"

	"github.com/tenta-browser/go-video-downloader/lib"
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
	"golang.org/x/net/publicsuffix"
)

// NewJar creates a new cookie jar
func NewJar() (http.CookieJar, error) {
	jar, err := cookiejar.New(&cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	})
	if err != nil {
		return nil, err
	}
	return jar, nil
}

// SetCookies parses the cookiesstr string and adds the resulting cookies
// (corresponding to the domain of rawurl) to the jar
func SetCookies(jar http.CookieJar, rawurl string, cookiesstr string) error {
	u, err := url.Parse(rawurl)
	if err != nil {
		return err
	}
	// we set the cookies on the empty path and http scheme
	u.Path = ""
	u.Scheme = "http"
	// we set the cookies on tld+1 domain
	dom, err := publicsuffix.EffectiveTLDPlusOne(u.Hostname())
	cookies := parseCookieString(cookiesstr)
	if err != nil {
		return err
	}
	if len(cookies) > 0 {
		for _, cookie := range cookies {
			cookie.Domain = dom
		}
		// actually set the cookies
		jar.SetCookies(u, cookies)
	}
	return nil
}

func parseCookieString(cookies string) []*http.Cookie {
	header := http.Header{}
	header.Add("Cookie", cookies)
	request := http.Request{Header: header}
	return request.Cookies()
}

// SetCookie adds a new cookie to the jar
var SetCookie = rnt.NewSimpleFunction("SetCookie",
	[]string{"jar", "domain", "name", "value"},
	func(args []rnt.Object) rnt.Object {
		jar := rnt.UnwrapNative(args[0]).(http.CookieJar)
		domain := args[1].(rnt.Str).Value()
		name := args[2].(rnt.Str).Value()
		value := args[3].(rnt.Str).Value()
		// create artificial source URL from domain
		u := &url.URL{
			Scheme: "http",
			Host:   domain,
		}
		// create cookie
		c := &http.Cookie{
			Name:   name,
			Value:  value,
			Domain: domain,
		}
		jar.SetCookies(u, []*http.Cookie{c})
		return rnt.None
	})

// GetCookieHeader returns the Cookie request header value for url.
var GetCookieHeader = rnt.NewSimpleFunction("GetCookieHeader",
	[]string{"jar", "url"},
	func(args []rnt.Object) rnt.Object {
		jar := rnt.UnwrapNative(args[0]).(http.CookieJar)
		urls := args[1].(rnt.Str).Value()
		u, err := url.Parse(urls)
		if err != nil {
			panic(rnt.RaiseType(lib.ExtractorErrorType, err.Error()))
		}
		cookies := jar.Cookies(u)
		var cs []string
		for _, c := range cookies {
			cs = append(cs, c.Name+"="+c.Value)
		}
		sort.Strings(cs)
		return rnt.NewStr(strings.Join(cs, "; "))
	})

// ParseCookieString parses cookies and returns a name-value map of them.
var ParseCookieString = rnt.NewSimpleFunction("ParseCookieString",
	[]string{"cs"},
	func(args []rnt.Object) rnt.Object {
		cs := args[0].(rnt.Str).Value()
		cookies := parseCookieString(cs)
		res := make(map[rnt.Object]rnt.Object)
		for _, cookie := range cookies {
			res[rnt.NewStr(cookie.Name)] = rnt.NewStr(cookie.Value)
		}
		return rnt.NewDictWithTable(res)
	})
