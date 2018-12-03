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
 * http.go: Simple http client
 */

package http

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

// Execute creates a request from the parameters, executes it,
// and returns a tuple of response info on success or an error
// message on failure.
var Execute = rnt.NewSimpleFunction("Execute",
	[]string{"client", "method", "url", "data", "headers"},
	func(args []rnt.Object) rnt.Object {
		// get client
		client := rnt.UnwrapNative(args[0]).(*http.Client)

		// initialize request
		method := args[1].(rnt.Str).Value()
		url := args[2].(rnt.Str).Value()
		req, err := http.NewRequest(method, url, nil)
		if err != nil {
			return rnt.NewTuple(rnt.None, rnt.NewStr(err.Error()))
		}

		// set headers
		if args[4] != rnt.None {
			it := args[4].(rnt.Dict).Iterator()
			for e := it(); e != nil; e = it() {
				key := e.Key.(rnt.Str).Value()
				value := e.Value.(rnt.Str).Value()
				req.Header.Set(key, value)
			}
			// remove 'Accept-Encoding' header, let the HTTP client set its default
			req.Header.Del("Accept-Encoding")
		}

		// set data
		if args[3] != rnt.None {
			data := args[3].(rnt.Bytes).Value()
			if len(data) > 0 {
				req.Body = ioutil.NopCloser(bytes.NewBuffer(data))
				req.ContentLength = int64(len(data))
			}
		}

		// execute!
		res, err := client.Do(req)
		if err != nil {
			return rnt.NewTuple(rnt.None, rnt.NewStr(err.Error()))
		}

		// extract response info
		nativeRes := rnt.NewNative(res)
		resURL := rnt.NewStr(res.Request.URL.String())

		// extract response status
		statusCode := rnt.NewInt(res.StatusCode)
		_status := res.Status
		if pref := fmt.Sprintf("%d ", res.StatusCode); strings.HasPrefix(_status, pref) {
			_status = _status[len(pref):]
		}
		status := rnt.NewStr(_status)

		// extract response headers
		_resHeaders := make(map[rnt.Object]rnt.Object)
		for key := range res.Header {
			value := res.Header.Get(key)
			_resHeaders[rnt.NewStr(key)] = rnt.NewStr(value)
		}
		resHeaders := rnt.NewDictWithTable(_resHeaders)

		return rnt.NewTuple(
			rnt.NewTuple(nativeRes, resURL, statusCode, status, resHeaders),
			rnt.None)
	})

// ReadResponseBody reads the entire body of the response and closes it.
var ReadResponseBody = rnt.NewSimpleFunction("ReadResponseBody", []string{"res"},
	func(args []rnt.Object) rnt.Object {
		res := rnt.UnwrapNative(args[0]).(*http.Response)
		defer res.Body.Close()
		bytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return rnt.NewTuple(rnt.None, rnt.NewStr(err.Error()))
		}
		return rnt.NewTuple(rnt.NewBytes(bytes...), rnt.None)
	})

// CloseResponse closes the response.
var CloseResponse = rnt.NewSimpleFunction("CloseResponse", []string{"res"},
	func(args []rnt.Object) rnt.Object {
		res := rnt.UnwrapNative(args[0]).(*http.Response)
		res.Body.Close()
		return rnt.None
	})
