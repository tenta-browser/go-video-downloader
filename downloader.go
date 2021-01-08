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
 * downloader.go: Downloader interface
 */

package downloader

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/tenta-browser/go-pcre-matcher"
	"github.com/tenta-browser/go-video-downloader/lib"
	"github.com/tenta-browser/go-video-downloader/lib/browser"
	"github.com/tenta-browser/go-video-downloader/lib/re"

	netlib "github.com/tenta-browser/go-video-downloader/lib/net"
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	masterRegexps     []string
	masterRegexpComps []matcher.Regexp
)

// BrowserRequest respresents a request towards a Browser.
type BrowserRequest struct {
	UserAgent string
	Cookies   string
	URL       string
}

// BrowserResponse respresents a response from a Browser.
type BrowserResponse struct {
	HTML    string
	Cookies string
}

// Browser represents an interface towards a real browser,
// which can load pages evaluating JavaScript.
type Browser interface {
	Load(*BrowserRequest) (*BrowserResponse, error)
}

// Connector holds settings/callbacks related to establishing connections,
// downloading content, etc.
type Connector struct {
	Browser   Browser
	Client    *http.Client
	UserAgent string
	Cookie    string
	Username  string
	Password  string
}

// VideoData holds the final result of the extraction
type VideoData struct {
	URL       string
	Title     string
	Filename  string
	AgeLimit  int
	Thumbnail string
	Cookie    string
}

// Init initializes the downloader:
// runs transpiled module initialization and compiles the master regexps
func Init() (err error) {
	defer func() {
		err = handlePanic(recover(), err)
	}()
	rnt.InitModules()
	if err := compileMasterRegexps(); err != nil {
		return err
	}
	return nil
}

// Check verifies if url contains a downloadable video
func Check(url string) (found bool, err error) {
	defer func() {
		err = handlePanic(recover(), err)
	}()
	// check if something matches the url
	for _, masterRegexpComp := range masterRegexpComps {
		if masterRegexpComp.Search(url) != nil {
			return true, nil
		}
	}
	return false, nil
}

// Extract extracts video info from the page at url
func Extract(url string, connector *Connector) (resData *VideoData, err error) {
	defer func() {
		err = handlePanic(recover(), err)
	}()
	defer rnt.FinalizeGenerators()
	connectorDict := rnt.NewDict()
	var client *http.Client
	if connector != nil {
		client = connector.Client
	}
	if client == nil {
		client = new(http.Client)
	}
	// we can't be sure the Python code cleans up after itself
	defer client.CloseIdleConnections()
	jar := client.Jar
	if jar == nil {
		jar, err = netlib.NewJar()
		if err != nil {
			return nil, err
		}
		client.Jar = jar
	}
	if connector != nil && connector.Cookie != "" {
		netlib.SetCookies(jar, url, connector.Cookie)
	}
	connectorDict.SetItem(rnt.NewStr("client"), rnt.NewNative(client))
	connectorDict.SetItem(rnt.NewStr("jar"), rnt.NewNative(jar))
	if connector != nil {
		if connector.Browser != nil {
			var b browser.Browser = &browserImpl{connector.Browser}
			connectorDict.SetItem(rnt.NewStr("browser"), rnt.NewNative(b))
		}
		if connector.UserAgent != "" {
			headersDict := rnt.NewDict()
			headersDict.SetItem(rnt.NewStr("User-Agent"), rnt.NewStr(connector.UserAgent))
			connectorDict.SetItem(rnt.NewStr("headers"), headersDict)
		}
		if connector.Username != "" {
			connectorDict.SetItem(rnt.NewStr("username"), rnt.NewStr(connector.Username))
		}
		if connector.Password != "" {
			connectorDict.SetItem(rnt.NewStr("password"), rnt.NewStr(connector.Password))
		}
	}
	res := rnt.Cal(lib.ExtractorRunner,
		rnt.NewStr(url),
		connectorDict,
	)
	if !res.IsInstance(rnt.DictType) {
		panic(rnt.RaiseType(lib.ExtractorErrorType, fmt.Sprintf(
			"result is not a dict, but %s", res.Type().Name())))
	}
	resDict := res.(rnt.Dict)
	resData = &VideoData{
		URL:       resStringField(resDict, "url", true, ""),
		Title:     resStringField(resDict, "title", true, ""),
		Filename:  resStringField(resDict, "_filename", false, "unknown.video"),
		AgeLimit:  resIntField(resDict, "age_limit", false, 0),
		Thumbnail: resStringField(resDict, "thumbnail", false, ""),
	}
	// TODO it would be great to return all the cookies of the jar
	if cookie, err := netlib.GetCookieString(jar, resData.URL); err == nil {
		resData.Cookie = cookie
	}
	return resData, nil
}

func handlePanic(r interface{}, originalErr error) error {
	if r == nil {
		return originalErr
	}
	var msg string
	if e, ok := r.(rnt.BaseException); ok {
		msg = e.String()
	} else {
		msg = fmt.Sprintf("Crash: %v", r)
	}
	return errors.New(msg)
}

func compileMasterRegexps() error {
	if masterRegexpComps != nil {
		// already compiled
		return nil
	}
	masterRegexpComps = make([]matcher.Regexp, len(masterRegexps))
	for idx, masterRegexp := range masterRegexps {
		var masterRegexpComp matcher.Regexp
		var err error
		if masterRegexpComp, err = re.CompileInternal(masterRegexp, 0); err != nil {
			return fmt.Errorf("compiling regexp %d: %s", idx, err.Error())
		}
		masterRegexpComps[idx] = masterRegexpComp
	}
	return nil
}

func resStringField(resDict rnt.Dict, name string, required bool, def string) string {
	val := resField(resDict, name, required)
	if val == nil {
		return def
	}
	if !val.IsInstance(rnt.StrType) {
		panic(rnt.RaiseType(lib.ExtractorErrorType, fmt.Sprintf(
			"result field '%s' should be a string, but it is a %s", name, val.Type().Name())))
	}
	return val.(rnt.Str).Value()
}

func resIntField(resDict rnt.Dict, name string, required bool, def int) int {
	val := resField(resDict, name, required)
	if val == nil {
		return def
	}
	if !val.IsInstance(rnt.IntType) {
		panic(rnt.RaiseType(lib.ExtractorErrorType, fmt.Sprintf(
			"result field '%s' should be an int, but it is a %s", name, val.Type().Name())))
	}
	return val.(rnt.Int).Value()
}

func resField(resDict rnt.Dict, name string, required bool) rnt.Object {
	val := resDict.GetItem(rnt.NewStr(name))
	if val == rnt.None {
		val = nil
	}
	if val == nil && required {
		panic(rnt.RaiseType(lib.ExtractorErrorType, fmt.Sprintf(
			"result dict is missing '%s'", name)))
	}
	return val
}

type browserImpl struct {
	Browser
}

func (b *browserImpl) Load(url, ua, inCookies string) (html string, outCookies string, err error) {
	req := &BrowserRequest{
		UserAgent: ua,
		Cookies:   inCookies,
		URL:       url,
	}
	res, err := b.Browser.Load(req)
	if err != nil {
		return "", "", err
	}
	return res.HTML, res.Cookies, nil
}
