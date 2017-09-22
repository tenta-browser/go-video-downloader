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
 * extractors.go: Structures and interfaces defining extractors,
 *                capable of extracting video data from a specific domain
 */

package runtime

import (
	"fmt"
	"net/http"

	"github.com/tenta-browser/go-video-downloader/utils"
)

// VideoResult is the most common extractor result,
// it represents a single downloadable video with direct url
type VideoResult struct {
	Type     string
	ID       string
	URL      string
	Title    string
	Filename string
}

// Context represents extraction context data
type Context struct {
	Client    *http.Client
	UserAgent string
}

// InfoExtractor represent the common interface of all extractors
type InfoExtractor interface {
	Ctx() *Context
	Key() string
	ValidUrl() string
	Name() string
	Extract(url string) (*VideoResult, error)
	Tests() []map[string]interface{}
}

// InfoExtractorFactory produces instances of InfoExtractors
type InfoExtractorFactory = func(ctx *Context) InfoExtractor

type extractorError struct {
	msg string
}

func (ee *extractorError) Error() string {
	return ee.msg
}

func newExtractorError(errorMsg string) error {
	return &extractorError{errorMsg}
}

// RunExtractor runs an extractor, handles errors, gathers/normalizes results
func RunExtractor(url string, extrFunc func(string) map[string]interface{}) (res *VideoResult, err error) {
	defer func() {
		// TODO this is a quick-n-dirty exception-like error handling
		// Maybe in the far future transpile functions to return error codes
		if r := recover(); r != nil {
			if e, ok := r.(*extractorError); ok {
				// if it's a extractor error 'catch' it and return it
				res, err = nil, e
			} else {
				panic(r) // otherwise re-panic the recovered panic
			}
		}
	}()

	resDict := extrFunc(url)
	utils.Log("Result dict: %+v", resDict)

	getStr := func(field string, required bool, def string) string {
		if val, ok := resDict[field]; ok {
			switch sval := val.(type) {
			case string:
				return sval
			case OptString:
				return sval.Get()
			default:
				panic(newExtractorError(fmt.Sprintf("Non-string %v found in the result dict", field)))
			}
		} else {
			if required {
				panic(newExtractorError(fmt.Sprintf("No %v found in the result dict", field)))
			} else {
				return def
			}
		}
	}

	resType := getStr("_type", false, "video")
	resTypeVideo := resType == "video"

	res = &VideoResult{
		Type:  resType,
		ID:    getStr("id", resTypeVideo, ""),
		URL:   getStr("url", resTypeVideo, ""),
		Title: getStr("title", resTypeVideo, ""),
	}

	var ext string
	ext, ok := resDict["ext"].(string)
	if !ok {
		ext = DetermineExt(AsOptString(res.URL), "unknown_video")
	}

	res.Filename = fmt.Sprintf("%s-%s.%s",
		utils.Ellipsize(SanitizeFilename(res.Title, true, false), 100, ".."),
		SanitizeFilename(res.ID, true, true),
		SanitizeFilename(ext, true, false))

	return res, nil
}
