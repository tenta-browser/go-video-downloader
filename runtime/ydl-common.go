package runtime

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/tenta-browser/go-pcre-matcher"
	"github.com/tenta-browser/go-video-downloader/utils"
)

// MatchID implements common.py/_match_id
func MatchID(ie InfoExtractor, url string) string {
	matcher := ReMatch(Re, ie.ValidUrl(), url, 0)
	if matcher == nil || !matcher.GroupPresentByName("id") {
		panic(newExtractorError("Failed to match id"))
	}
	return matcher.GroupByName("id")
}

// DownloadWebpage implements common.py/_download_webpage
// TODO handle all the params!
func DownloadWebpage(ie InfoExtractor, url, videoID string, note, errNote OptString,
	fatal bool, tries int, timeout int, encoding OptString, data OptString,
	headers, query map[string]interface{}) string {

	utils.Log("Downloading webpage for %s [%s]", videoID, url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(newExtractorError(err.Error()))
	}
	req.Header.Set("User-Agent", ie.Ctx().UserAgent)
	res, err := ie.Ctx().Client.Do(req)
	if err != nil {
		panic(newExtractorError(err.Error()))
	}
	ret, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		panic(newExtractorError(err.Error()))
	}
	return string(ret)
}

// SearchRegex implements common.py/_search_regex
func SearchRegex(ie InfoExtractor, pattern interface{}, str, name string,
	def interface{}, fatal bool, flags int, group interface{}) OptString {

	patterns, ok := pattern.([]string)
	if !ok {
		patterns = []string{pattern.(string)}
	}

	var matcher matcher.Matcher
	for _, pattern := range patterns {
		if matcher = ReSearch(Re, pattern, str, flags); matcher != nil {
			break
		}
	}

	if matcher != nil {
		if group == nil {
			// return first group that is present in the match
			for i := 1; i <= matcher.Groups(); i++ {
				if matcher.GroupPresentByIdx(i) {
					return AsOptString(matcher.GroupByIdx(i))
				}
			}
			// this should not be
			return AsOptString("")
		}
		// return the group requested by index/name or None if it isn't present
		return ReMatchGroupOne(matcher, group)
	} else if !isNoDefault(def) {
		if def == nil {
			return OptString{}
		}
		return AsOptString(def.(string))
	} else if fatal {
		panic(newExtractorError(fmt.Sprintf("Unable to extract %v", name)))
	} else {
		utils.Log("Unable to extract %v", name)
		return OptString{}
	}
}

// HTMLSearchRegex implements common.py/_html_search_regex
func HTMLSearchRegex(ie InfoExtractor, pattern interface{}, body, name string,
	def interface{}, fatal bool, flags int, group interface{}) OptString {

	res := SearchRegex(ie, pattern, body, name, def, fatal, flags, group)

	return CleanHTML(res)
}

// HTMLSearchMeta implements common.py/_html_search_meta
func HTMLSearchMeta(ie InfoExtractor, name interface{}, html string, displayName OptString,
	fatal bool, def interface{}, flags int) OptString {

	names, ok := name.([]string)
	if !ok {
		names = []string{name.(string)}
	}

	_displayName := displayName.GetOrDef(names[0])

	patterns := make([]string, len(names))
	for i, name := range names {
		patterns[i] = metaRegex(name)
	}

	return HTMLSearchRegex(ie, patterns, html, _displayName, def, fatal, flags, "content")
}

func metaRegex(prop string) string {
	return fmt.Sprintf(
		`(?is)<meta(?=[^>]+(?:itemprop|name|property|id|http-equiv)=(["\']?)%s\1)[^>]+?content=(["\'])(?P<content>.*?)\2`,
		matcher.ReEngine.Quote(prop))
}

func ogRegexes(prop string) []string {
	contentRe := "content=(?:\"([^\"]+?)\"|'([^']+?)'|\\s*([^\\s\"'=<>`]+?))"
	propertyRe := fmt.Sprintf(`(?:name|property)=(?:\'og:%[1]s\'|"og:%[1]s"|\s*og:%[1]s\b)`, matcher.ReEngine.Quote(prop))
	template := `<meta[^>]+?%s[^>]+?%s`
	return []string{
		fmt.Sprintf(template, propertyRe, contentRe),
		fmt.Sprintf(template, contentRe, propertyRe),
	}
}

// OgSearchProperty implements common.py/_og_search_property
func OgSearchProperty(ie InfoExtractor, prop interface{}, html string, name OptString, def interface{}, fatal bool) OptString {
	props, ok := prop.([]string)
	if !ok {
		props = []string{prop.(string)}
	}

	_name := name.GetOrDef(fmt.Sprintf("OpenGraph %s", props[0]))

	regexes := []string{}
	for _, prop := range props {
		regexes = append(regexes, ogRegexes(prop)...)
	}

	escaped := SearchRegex(ie, regexes, html, _name, def, fatal, ReFlagDotAll, nil)

	return UnescapeHTML(escaped)
}

// OgSearchThumbnail implements common.py/_og_search_thumbnail
func OgSearchThumbnail(ie InfoExtractor, html string, def interface{}) OptString {
	return OgSearchProperty(ie, "image", html, AsOptString("thumbnail URL"), def, false)
}

// OgSearchDescription implements common.py/_og_search_description
func OgSearchDescription(ie InfoExtractor, html string, def interface{}) OptString {
	return OgSearchProperty(ie, "description", html, OptString{}, def, false)
}

// OgSearchTitle implements common.py/_og_search_title
func OgSearchTitle(ie InfoExtractor, html string, def interface{}, fatal bool) OptString {
	return OgSearchProperty(ie, "title", html, OptString{}, def, fatal)
}

// OgSearchVideoURL implements common.py/_og_search_video_url
func OgSearchVideoURL(ie InfoExtractor, html string, name string, secure bool, def interface{}) OptString {
	regexes := append(ogRegexes("video"), ogRegexes("video:url")...)
	if secure {
		regexes = append(ogRegexes("video:secure_url"), regexes...)
	}
	return HTMLSearchRegex(ie, regexes, html, name, def, true, 0, nil)
}

// OgSearchURL implements common.py/_og_search_url
func OgSearchURL(ie InfoExtractor, html string) OptString {
	return OgSearchProperty(ie, "url", html, OptString{}, NoDefault, true)
}

// ParseJSON implements common.py/_parse_json
func ParseJSON(ie InfoExtractor, jsonString string, videoID string, transformSource int, fatal bool) map[string]interface{} {
	// TODO handle transformSource constants
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(jsonString), &result); err != nil {
		errMsg := fmt.Sprintf("%v: Failed to parse JSON", videoID)
		if fatal {
			panic(newExtractorError(errMsg))
		} else {
			utils.Log(errMsg)
			return nil
		}
	}
	return result
}

// DownloadJSON implements common.py/_download_json
func DownloadJSON(ie InfoExtractor, url, videoID string, note, errNote OptString,
	transformSource int, fatal bool, encoding OptString, data OptString,
	headers, query map[string]interface{}) map[string]interface{} {

	jsonString := DownloadWebpage(ie, url, videoID, note, errNote, fatal, 1, 5, encoding, data, headers, query)

	// TODO some fatal handling

	return ParseJSON(ie, jsonString, videoID, transformSource, fatal)
}

// URLResult implements common.py/url_result
func URLResult(ie InfoExtractor, url string, ieKey, videoID, videoTitle OptString) map[string]interface{} {
	result := map[string]interface{}{
		"_type": "url",
		"url":   url,
		"ieKey": ieKey,
	}
	if videoID.IsSet() {
		result["id"] = videoID.Get()
	}
	if videoTitle.IsSet() {
		result["title"] = videoTitle.Get()
	}
	return result
}
