package downloader

import (
	"errors"
	"net/http"

	"github.com/tenta-browser/go-pcre-matcher"
	"github.com/tenta-browser/go-video-downloader/extractor"
	"github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	masterRegexp     string
	masterRegexpComp matcher.Regexp
)

// CheckResult holds all the info obtained by verifying if an URL contains
// a downloadable video (extractorKey for now)
type CheckResult struct {
	url          string
	extractorKey string
}

// Connector holds settings/callbacks related to establishing connections,
// downloading content, etc.
type Connector struct {
	Client    *http.Client
	UserAgent string
}

// VideoData holds the final result of the extraction
type VideoData struct {
	URL      string
	Title    string
	Filename string
}

// Check verifies if url contains a downloadable video; returns nil if not
func Check(url string) *CheckResult {
	// if the master regexp isn't compiled, compile it
	if masterRegexpComp == nil {
		var err error
		masterRegexpComp, err = matcher.ReEngine.Compile(masterRegexp, 0)
		if err != nil {
			panic(err.Error())
		}
	}
	// check if something matches the url
	matcher := masterRegexpComp.Search(url)
	if matcher == nil {
		return nil
	}
	// something matched the url, now we have to find it :|
	for extractorKey := range extractor.GetFactories() {
		if matcher.GroupPresentByName("extr_" + extractorKey) {
			return &CheckResult{
				url:          url,
				extractorKey: extractorKey,
			}
		}
	}
	// master regexp and the extractors are not in sync
	panic("Url was matched but no extractor found: " + url)
}

// Extract extracts video info based on the results of a successful Check()
func Extract(checkResult *CheckResult, connector *Connector) (*VideoData, error) {
	factory := extractor.GetFactories()[checkResult.extractorKey]
	extractor := factory(&runtime.Context{
		Client:    connector.Client,
		UserAgent: connector.UserAgent,
	})
	res, err := extractor.Extract(checkResult.url)
	if err != nil {
		return nil, err
	}
	if res.Type != "video" {
		return nil, errors.New("Unsupported result type: " + res.Type)
	}
	return &VideoData{
		URL:      res.URL,
		Title:    res.Title,
		Filename: res.Filename,
	}, nil
}
