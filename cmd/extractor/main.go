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
 * main.go: Example program
 */

package main

import (
	"fmt"
	"github.com/tenta-browser/go-pcre-matcher"
	"github.com/tenta-browser/go-pcre-matcher/matcherpcre"
	"github.com/tenta-browser/go-video-downloader"
	"net/http"
	"os"
)

func main() {
	// Validate arguments.
	if len(os.Args) != 2 {
		fmt.Printf("Usage: extractor URL\n")
		os.Exit(2)
	}
	url := os.Args[1]

	// Select PCRE matcher engine.
	// You can plug in any regular expression engine, implementing the matcher.Engine interface.
	matcher.ReEngine = matcherpcre.NewEngine()

	// The downloader has to be initialized once, before first usage.
	if err := downloader.Init(); err != nil {
		fmt.Printf("Downloader intialization failed: %v\n", err)
		os.Exit(1)
	}

	// Quick check to see if the specified URL contains a video (for which extraction is supported).
	hasVideo, err := downloader.Check(url)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Video extraction supported for %s: %v\n", url, hasVideo)

	// Pluggable HTTP client and common headers to be used at each page request
	// done by the extractor.
	connector := &downloader.Connector{
		Client:    &http.Client{},
		UserAgent: "Mozilla/5.0 (X11; Linux x86_64) ..",
		Cookie:    "SESSIONID=12345",
	}

	// Extracts various properties of the video contained at the specified URL,
	// most importantly it's direct URL.
	// This is a slow call, the extractor will possibly do multiple HTTP requests.
	videoData, err := downloader.Extract(url, connector)
	if err != nil {
		fmt.Printf("Video extraction failed: %v\n", err)
		os.Exit(1)
	}
	println("Video title:", videoData.Title)
	println("Raw video URL:", videoData.URL)
	println("Suggested filename:", videoData.Filename)
}
