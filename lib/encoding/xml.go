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
 * xml.go: XML parsing and searching methods
 */

package encoding

import (
	"fmt"
	"io"
	"strings"

	"github.com/beevik/etree"
	"github.com/tenta-browser/go-video-downloader/lib"
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type xmlElement struct {
	spaces map[string]string // maps namespace URL -> namespace name
	ee     *etree.Element
}

func newXMLElement(parent *xmlElement, ee *etree.Element) *xmlElement {
	var spaces map[string]string
	if parent == nil {
		// we are at the root; calculate namespace map, used in fixPath()
		spaces = map[string]string{}
		for _, attr := range ee.Attr {
			if attr.Space == "xmlns" {
				spaces[attr.Value] = attr.Key
			} else if attr.Key == "xmlns" {
				spaces[attr.Value] = ""
			}
		}
	} else {
		// inherit namespace map
		spaces = parent.spaces
	}
	return &xmlElement{
		spaces: spaces,
		ee:     ee,
	}
}

// XMLParseErrorType is the error type raised by XML parsing ..
var XMLParseErrorType = rnt.Cal(
	rnt.TypeType,
	rnt.NewStr("XMLParseError"),
	rnt.NewTuple(rnt.ExceptionType),
	rnt.NewDict()).(rnt.Type)

// XMLParse implements python/xml.etree.ElementTree.XML
var XMLParse = rnt.NewSimpleFunction("XMLParse", []string{"data"},
	func(args []rnt.Object) rnt.Object {
		data, _ := lib.StringlikeVal(args[0])
		doc := etree.NewDocument()
		doc.ReadSettings.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
			return input, nil
		}
		err := doc.ReadFromString(data)
		if err != nil {
			panic(rnt.RaiseType(XMLParseErrorType, err.Error()))
		}
		root := doc.Root()
		if root == nil {
			return rnt.None
		}
		return rnt.NewNative(newXMLElement(nil, root))
	})

// XMLElementInfos returns (tag, text, attrib, childno) tuple for the provided element
var XMLElementInfos = rnt.NewSimpleFunction("XMLElementInfos", []string{"element"},
	func(args []rnt.Object) rnt.Object {
		el := rnt.UnwrapNative(args[0]).(*xmlElement)
		tag := rnt.NewStr(el.ee.Tag)
		text := rnt.NewStr(el.ee.Text())
		attrib := rnt.NewDict()
		for _, attr := range el.ee.Attr {
			attrib.SetItem(rnt.NewStr(attr.Key), rnt.NewStr(attr.Value))
		}
		childno := rnt.NewInt(len(el.ee.ChildElements()))
		return rnt.NewTuple(tag, text, attrib, childno)
	})

// XMLGetChildElement return the idx-th child element of element
var XMLGetChildElement = rnt.NewSimpleFunction("XMLGetChildElement", []string{"element", "idx"},
	func(args []rnt.Object) rnt.Object {
		el := rnt.UnwrapNative(args[0]).(*xmlElement)
		idx := args[1].(rnt.Int).Value()
		children := el.ee.ChildElements()
		if idx >= len(children) {
			panic(rnt.Raise(rnt.IndexErrorType))
		}
		child := children[idx]
		return rnt.NewNative(newXMLElement(el, child))
	})

// XMLFind implements python/xml.etree.ElementTree.Element.find()
var XMLFind = rnt.NewSimpleFunction("XMLFind", []string{"element", "path"},
	func(args []rnt.Object) rnt.Object {
		el := rnt.UnwrapNative(args[0]).(*xmlElement)
		path, _ := lib.StringlikeVal(args[1])
		path = fixPath(el, path)
		resEl := el.ee.FindElement(path)
		if resEl == nil {
			return rnt.None
		}
		return rnt.NewNative(newXMLElement(el, resEl))
	})

// XMLFindAll implements python/xml.etree.ElementTree.Element.findall()
var XMLFindAll = rnt.NewSimpleFunction("XMLFindAll", []string{"element", "path"},
	func(args []rnt.Object) rnt.Object {
		el := rnt.UnwrapNative(args[0]).(*xmlElement)
		path, _ := lib.StringlikeVal(args[1])
		path = fixPath(el, path)
		res := []rnt.Object{}
		for _, resEl := range el.ee.FindElements(path) {
			res = append(res, rnt.NewNative(newXMLElement(el, resEl)))
		}
		return rnt.NewList(res...)
	})

// fixPath replaces namespace URLs into namespace names,
// because this lib doesn't support them in paths
func fixPath(el *xmlElement, path string) string {
	for spaceURL, spaceName := range el.spaces {
		repl := ""
		if spaceName != "" {
			repl = fmt.Sprintf("%s:", spaceName)
		}
		path = strings.Replace(path,
			fmt.Sprintf("{%s}", spaceURL),
			repl,
			-1)
	}
	return path
}
