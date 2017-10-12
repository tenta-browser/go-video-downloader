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
	AgeLimit int
}

// Context represents extraction context data
type Context struct {
	Client  *http.Client
	Headers map[string]string
}

// InfoExtractor represent the common interface of all extractors
type InfoExtractor interface {
	SetContext(context *Context)
	Key() string
	ValidURL() string
	Name() string
	Extract(url string) (*VideoResult, error)
	Tests() []map[string]interface{}
}

// InfoExtractorFactory produces instances of InfoExtractors
type InfoExtractorFactory = func() InfoExtractor

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

	resType := GetStringField(resDict, "_type", false, "video")
	resTypeVideo := resType == "video"

	res = &VideoResult{
		Type:     resType,
		ID:       GetStringField(resDict, "id", resTypeVideo, ""),
		Title:    GetStringField(resDict, "title", resTypeVideo, ""),
		AgeLimit: GetIntField(resDict, "age_limit", false, 0),
	}

	if resTypeVideo {
		if formats, ok := resDict["formats"]; ok {
			if _formats, ok := formats.([]interface{}); ok {
				if format, ok := _formats[len(_formats)-1].(map[string]interface{}); ok {
					res.URL = GetStringField(format, "url", true, "")
				}
			}
			if res.URL == "" {
				panic(newExtractorError("Failed to extract URL from formats"))
			}
		} else {
			res.URL = GetStringField(resDict, "url", true, "")
		}
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
