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
 * builtin-xml.go: Re-implementation of parts of Python's XML parsing module
 */

package runtime

import (
	"fmt"
	"strings"

	"github.com/beevik/etree"
)

type xmlElementStruct struct {
	spaces     map[string]string // maps namespace URL -> namespace name
	ee         *etree.Element
	Text       string
	Attributes SDict
}

// XMLElement is an element in a parsed XML tree
type XMLElement *xmlElementStruct

// XMLParse parses xmlString and returns the root element
func XMLParse(xmlString string) (XMLElement, error) {
	doc := etree.NewDocument()
	err := doc.ReadFromString(xmlString)
	if err != nil {
		return nil, err
	}
	return NewXMLElement(nil, doc.Root()), nil
}

// NewEmptyXMLElement constructs and empty XMLElement
func NewEmptyXMLElement() XMLElement {
	return &xmlElementStruct{}
}

// NewXMLElement wraps a raw element and prepares needed properties
func NewXMLElement(parent XMLElement, ee *etree.Element) XMLElement {
	if ee == nil {
		return nil
	}
	var spaces map[string]string
	if parent == nil {
		// we are at the root; calculate namespace map, used in fixPath()
		spaces = map[string]string{}
		for _, attr := range ee.Attr {
			if attr.Space == "xmlns" {
				spaces[attr.Value] = attr.Key
			}
		}
	} else {
		// inherit namespace map
		spaces = parent.spaces
	}
	// calculate attribute map
	attributes := SDict{}
	for _, attr := range ee.Attr {
		attributes[attr.Key] = attr.Value
	}
	return &xmlElementStruct{
		spaces:     spaces,
		ee:         ee,
		Text:       ee.Text(),
		Attributes: attributes}
}

// XMLGetAttribute implements python/xml.etree.ElementTree.Element.get()
func XMLGetAttribute(el XMLElement, key string, def OptString) OptString {
	val, found := el.Attributes[key]
	if !found {
		return def
	}
	return AsOptString(val.(string))
}

// XMLFind implements python/xml.etree.ElementTree.Element.find()
func XMLFind(el XMLElement, path string) XMLElement {
	path = fixPath(el, path)
	return NewXMLElement(el, el.ee.FindElement(path))
}

// XMLFindAll implements python/xml.etree.ElementTree.Element.findall()
func XMLFindAll(el XMLElement, path string) []XMLElement {
	path = fixPath(el, path)
	result := []XMLElement{}
	for _, ee := range el.ee.FindElements(path) {
		result = append(result, NewXMLElement(el, ee))
	}
	return result
}

// fixPath replaces namespace URLs into namespace names,
// because this lib doesn't support them in paths
func fixPath(el XMLElement, path string) string {
	for spaceURL, spaceName := range el.spaces {
		path = strings.Replace(path,
			fmt.Sprintf("{%s}", spaceURL),
			fmt.Sprintf("%s:", spaceName),
			-1)
	}
	return path
}
