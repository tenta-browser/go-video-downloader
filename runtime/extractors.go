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
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"runtime/debug"
	"strings"

	"github.com/tenta-browser/go-video-downloader/utils"
)

// ExtractorResult is the base interface for all types of results
type ExtractorResult interface {
	Type() string
	Dict() SDict
}

type extractorResult struct {
	resType string
	resDict SDict
}

func (res *extractorResult) Type() string {
	return res.resType
}

func (res *extractorResult) Dict() SDict {
	return res.resDict
}

// VideoResult is the most common extractor result,
// it represents a single downloadable video with direct url
// (has type "video")
type VideoResult struct {
	*extractorResult
	ID       string
	URL      string
	Title    string
	Filename string
	AgeLimit int
}

// URLResult contains a resulting URL for which extraction should be rerun
// (has type "url" or "url_transparent")
type URLResult struct {
	*extractorResult
	URL          string
	ExtractorKey string
}

// Context represents extraction context data
type Context struct {
	ExtractorKey string          // Key of the current extractor
	Referrer     ExtractorResult // Result of the extraction that initiated this extraction
	Client       *http.Client
	Headers      map[string]string
}

// InfoExtractor represent the common interface of all extractors
type InfoExtractor interface {
	SetContext(context *Context)
	Key() string
	ValidURL() string
	Name() string
	Extract(url string) (ExtractorResult, error)
	Tests() []SDict
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
func RunExtractor(url string, ctx *Context, extrFunc func(string) SDict) (res ExtractorResult, err error) {
	defer func() {
		// TODO this is a quick-n-dirty exception-like error handling
		// Maybe in the far future transpile functions to return error codes
		if r := recover(); r != nil {
			if e, ok := r.(*extractorError); ok {
				// if it's a extractor error 'catch' it and return it
				res, err = nil, e
				if utils.Debug {
					utils.Log("Panic: %v", string(debug.Stack()))
				}
			} else {
				panic(r) // otherwise re-panic the recovered panic
			}
		}
	}()

	var (
		resDict = extrFunc(url)
		resType = GetStringField(resDict, "_type", false, "video")
		eres    = &extractorResult{resType, resDict}
	)

	if ctx.Referrer != nil {
		refDict := ctx.Referrer.Dict()
		delete(refDict, "url")
		delete(refDict, "ie_key")
		switch t := ctx.Referrer.Type(); t {
		case "url":
			// extend resDict with properties from refDict
			for k, v := range refDict {
				if _, found := resDict[k]; !found {
					resDict[k] = v
				}
			}
		case "url_transparent":
			// overwrite resDict properties from refDict
			for k, v := range refDict {
				resDict[k] = v
			}
		default:
			panic(newExtractorError("Unhandled referrer type: " + t))
		}
	}

	if resType == "video" {
		url := ""
		ext := GetStringField(resDict, "ext", false, "")
		id := GetStringField(resDict, "id", true, "")
		title := GetStringField(resDict, "title", true, "")

		if formats, ok := resDict["formats"]; ok {
			format, err := selectFormat(formats)
			if err != nil {
				panic(newExtractorError(err.Error()))
			}
			url = GetStringField(format, "url", true, "")
			if ext == "" {
				ext = GetStringField(format, "ext", false, "")
			}
		} else {
			url = GetStringField(resDict, "url", true, "")
		}

		url = utils.SanitizeURL(url)

		filename := fmt.Sprintf("%s-%s.%s",
			utils.Ellipsize(SanitizeFilename(title, true, false), 100, ".."),
			SanitizeFilename(id, true, true),
			SanitizeFilename(ext, true, false))

		return &VideoResult{
			extractorResult: eres,
			ID:              id,
			URL:             url,
			Title:           title,
			Filename:        filename,
			AgeLimit:        GetIntField(resDict, "age_limit", false, 0),
		}, nil
	} else if resType == "url" || resType == "url_transparent" {
		return &URLResult{
			extractorResult: eres,
			URL:             utils.SanitizeURL(GetStringField(resDict, "url", true, "")),
			ExtractorKey:    GetStringField(resDict, "ie_key", false, ""),
		}, nil
	} else {
		panic(newExtractorError("Unsupported result type: " + resType))
	}
}

func selectFormat(formats interface{}) (SDict, error) {
	formatsVal := reflect.ValueOf(formats)
	if formatsVal.Kind() != reflect.Slice {
		return nil, errors.New("Formats is not a slice")
	}

	formatCount := formatsVal.Len()
	for i := formatCount - 1; i >= 0; i-- {
		formatVal := formatsVal.Index(i)
		format, ok := formatVal.Interface().(SDict)
		if !ok {
			// not an SDict, skip
			continue
		}
		if GetStringField(format, "url", false, "") == "" {
			// no url found, skip
			continue
		}
		formatID := strings.ToUpper(GetStringField(format, "format_id", false, ""))
		if strings.Contains(formatID, "HLS") {
			// TODO we don't support HLS downloading yet
			if utils.Debug {
				utils.Log("WARNING: Skipping format: %v", formatID)
			}
			continue
		}
		// we found a good format
		return format, nil
	}

	return nil, errors.New("No suitable format found")
}
