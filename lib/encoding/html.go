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
 * html.go: HTML manipulation methods
 */

package encoding

import (
	"strings"

	rnt "github.com/tenta-browser/go-video-downloader/runtime"
	"golang.org/x/net/html"
)

// HTMLParseErrorType is the error type raised by HTML parsing
var HTMLParseErrorType = rnt.Cal(
	rnt.TypeType,
	rnt.NewStr("HTMLParseError"),
	rnt.NewTuple(rnt.ExceptionType),
	rnt.NewDict()).(rnt.Type)

// HTMLParse parses s as an HTML document, and calls the cb_starttag(tag, attrs) callback
// for each tag encountered (attrs is a list of (key,value) tuples).
var HTMLParse = rnt.NewSimpleFunction("HTMLParse", []string{"s", "cb_starttag"},
	func(args []rnt.Object) rnt.Object {
		s := args[0].(rnt.Str).Value()
		cbStartTag := args[1]
		doc, err := html.Parse(strings.NewReader(s))
		if err != nil {
			panic(rnt.RaiseType(HTMLParseErrorType, err.Error()))
		}
		var visit func(*html.Node)
		visit = func(n *html.Node) {
			if n.Type == html.ElementNode {
				attrs := make([]rnt.Object, len(n.Attr))
				for i, attr := range n.Attr {
					attrs[i] = rnt.NewTuple(rnt.NewStr(attr.Key), rnt.NewStr(attr.Val))
				}
				rnt.Cal(cbStartTag, rnt.NewStr(n.Data), rnt.NewList(attrs...))
			}
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				visit(c)
			}
		}
		visit(doc)
		return rnt.None
	})
