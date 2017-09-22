package extractor

import (
	"github.com/tenta-browser/go-video-downloader/runtime"
)

var factories = make(map[string]runtime.InfoExtractorFactory)

func registerFactory(key string, factory runtime.InfoExtractorFactory) {
	factories[key] = factory
}

// GetFactories returns factories for every InfoExtractor
func GetFactories() map[string]runtime.InfoExtractorFactory {
	return factories
}
