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
 * ydl-common.go: Re-implementation of various functions from
 *                https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/common.py
 */

package runtime

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"

	"github.com/tenta-browser/go-pcre-matcher"
	"github.com/tenta-browser/go-video-downloader/utils"
)

// CommonIE represents the common extractor base
type CommonIE struct {
	context  *Context
	VALIDURL string
	TESTS    []map[string]interface{}
	TEST     map[string]interface{}
}

// NewCommonIE constructs a new common extractor base
func NewCommonIE() *CommonIE {
	return &CommonIE{}
}

// SetContext sets the runtime extractor context
func (ie *CommonIE) SetContext(context *Context) {
	ie.context = context
}

// ValidURL returns the regexp matching URLs supported by this extractor
func (ie *CommonIE) ValidURL() string {
	return ie.VALIDURL
}

// Tests returns the test cases for this extractor
func (ie *CommonIE) Tests() []map[string]interface{} {
	if len(ie.TEST) > 0 {
		return []map[string]interface{}{ie.TEST}
	}
	return ie.TESTS
}

// MatchID implements common.py/_match_id
func (ie *CommonIE) MatchID(url string) string {
	match := ReMatch(ie.VALIDURL, url, 0)
	if match == nil || !match.GroupPresentByName("id") {
		panic(newExtractorError("Failed to match id"))
	}
	return match.GroupByName("id")
}

// DownloadWebpageURL implements common.py/_download_webpage (for urls)
func (ie *CommonIE) DownloadWebpageURL(url, videoID string, note, errNote OptString,
	fatal bool, tries int, timeout int, encoding OptString, data OptString,
	headers, query map[string]interface{}) string {

	return ie.DownloadWebpageRequest(SanitizedRequest(url), videoID, note, errNote,
		fatal, tries, timeout, encoding, data, headers, query)
}

// DownloadWebpageRequest implements common.py/_download_webpage (for requests)
// TODO handle all the params!
func (ie *CommonIE) DownloadWebpageRequest(req *http.Request, videoID string, note, errNote OptString,
	fatal bool, tries int, timeout int, encoding OptString, data OptString,
	headers, query map[string]interface{}) string {

	utils.Log("[%s] %s (%s)", videoID, note.GetOrDef("Downloading webpage"), req.URL)

	// priority of headers: context < request < args
	for headerName, headerVal := range ie.context.Headers {
		if len(headerVal) > 0 && len(req.Header.Get(headerVal)) == 0 {
			req.Header.Set(headerName, headerVal)
		}
	}
	for headerName := range headers {
		req.Header.Set(headerName, GetStringField(headers, headerName, true, ""))
	}

	res, err := ie.context.Client.Do(req)
	if err != nil {
		panic(newExtractorError(err.Error()))
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		panic(newExtractorError(fmt.Sprintf("Unable to download webpage: HTTP Error %d: %s",
			res.StatusCode, res.Status)))
	}

	ret, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(newExtractorError(err.Error()))
	}

	return string(ret)
}

// SearchRegexOne implements common.py/_search_regex (for single pattern)
func (ie *CommonIE) SearchRegexOne(pattern string, str, name string,
	def interface{}, fatal bool, flags int, group interface{}) OptString {

	return ie.SearchRegexMulti([]string{pattern}, str, name, def, fatal, flags, group)
}

// SearchRegexMulti implements common.py/_search_regex (for multiple patterns)
func (ie *CommonIE) SearchRegexMulti(patterns []string, str, name string,
	def interface{}, fatal bool, flags int, group interface{}) OptString {

	var match matcher.Match
	for _, pattern := range patterns {
		if match = ReSearch(pattern, str, flags); match != nil {
			break
		}
	}

	if match != nil {
		if group == nil {
			// return first group that is present in the match
			for i := 1; i <= match.Groups(); i++ {
				if match.GroupPresentByIdx(i) {
					return AsOptString(match.GroupByIdx(i))
				}
			}
			// this should not be
			return AsOptString("")
		}
		// return the group requested by index/name or None if it isn't present
		return ReMatchGroupOne(match, group)
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

// HTMLSearchRegexOne implements common.py/_html_search_regex (for single pattern)
func (ie *CommonIE) HTMLSearchRegexOne(pattern string, body, name string,
	def interface{}, fatal bool, flags int, group interface{}) OptString {

	return ie.HTMLSearchRegexMulti([]string{pattern}, body, name, def, fatal, flags, group)
}

// HTMLSearchRegexMulti implements common.py/_html_search_regex (for multiple patterns)
func (ie *CommonIE) HTMLSearchRegexMulti(patterns []string, body, name string,
	def interface{}, fatal bool, flags int, group interface{}) OptString {

	res := ie.SearchRegexMulti(patterns, body, name, def, fatal, flags, group)

	return CleanHTML(res)
}

// HTMLSearchMetaOne implements common.py/_html_search_meta (for single name)
func (ie *CommonIE) HTMLSearchMetaOne(name string, html string, displayName OptString,
	fatal bool, def interface{}, flags int) OptString {

	return ie.HTMLSearchMetaMulti([]string{name}, html, displayName, fatal, def, flags)
}

// HTMLSearchMetaMulti implements common.py/_html_search_meta (for multiple names)
func (ie *CommonIE) HTMLSearchMetaMulti(names []string, html string, displayName OptString,
	fatal bool, def interface{}, flags int) OptString {

	_displayName := displayName.GetOrDef(names[0])

	patterns := make([]string, len(names))
	for i, name := range names {
		patterns[i] = metaRegex(name)
	}

	return ie.HTMLSearchRegexMulti(patterns, html, _displayName, def, fatal, flags, "content")
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

// OgSearchPropertyOne implements common.py/_og_search_property (for single property)
func (ie *CommonIE) OgSearchPropertyOne(prop string, html string, name OptString, def interface{}, fatal bool) OptString {
	return ie.OgSearchPropertyMulti([]string{prop}, html, name, def, fatal)
}

// OgSearchPropertyMulti implements common.py/_og_search_property (for multiple properties)
func (ie *CommonIE) OgSearchPropertyMulti(props []string, html string, name OptString, def interface{}, fatal bool) OptString {
	_name := name.GetOrDef(fmt.Sprintf("OpenGraph %s", props[0]))

	regexes := []string{}
	for _, prop := range props {
		regexes = append(regexes, ogRegexes(prop)...)
	}

	escaped := ie.SearchRegexMulti(regexes, html, _name, def, fatal, ReFlagDotAll, nil)

	return UnescapeHTML(escaped)
}

// OgSearchThumbnail implements common.py/_og_search_thumbnail
func (ie *CommonIE) OgSearchThumbnail(html string, def interface{}) OptString {
	return ie.OgSearchPropertyOne("image", html, AsOptString("thumbnail URL"), def, false)
}

// OgSearchDescription implements common.py/_og_search_description
func (ie *CommonIE) OgSearchDescription(html string, def interface{}) OptString {
	return ie.OgSearchPropertyOne("description", html, OptString{}, def, false)
}

// OgSearchTitle implements common.py/_og_search_title
func (ie *CommonIE) OgSearchTitle(html string, def interface{}, fatal bool) OptString {
	return ie.OgSearchPropertyOne("title", html, OptString{}, def, fatal)
}

// OgSearchVideoURL implements common.py/_og_search_video_url
func (ie *CommonIE) OgSearchVideoURL(html string, name string, secure bool, def interface{}) OptString {
	regexes := append(ogRegexes("video"), ogRegexes("video:url")...)
	if secure {
		regexes = append(ogRegexes("video:secure_url"), regexes...)
	}
	return ie.HTMLSearchRegexMulti(regexes, html, name, def, true, 0, nil)
}

// OgSearchURL implements common.py/_og_search_url
func (ie *CommonIE) OgSearchURL(html string) OptString {
	return ie.OgSearchPropertyOne("url", html, OptString{}, NoDefault, true)
}

// ParseJSON implements common.py/_parse_json
func (ie *CommonIE) ParseJSON(jsonString string, videoID string, transformSource func(string) string, fatal bool) map[string]interface{} {
	if transformSource != nil {
		jsonString = transformSource(jsonString)
	}
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

// ParseJSONList implements common.py/_parse_json (returning lists)
func (ie *CommonIE) ParseJSONList(jsonString string, videoID string, transformSource func(string) string, fatal bool) []interface{} {
	if transformSource != nil {
		jsonString = transformSource(jsonString)
	}
	var result []interface{}
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
func (ie *CommonIE) DownloadJSON(url, videoID string, note, errNote OptString,
	transformSource func(string) string, fatal bool, encoding OptString, data OptString,
	headers, query map[string]interface{}) map[string]interface{} {

	jsonString := ie.DownloadWebpageURL(url, videoID, note, errNote, fatal, 1, 5, encoding, data, headers, query)

	// TODO some fatal handling

	return ie.ParseJSON(jsonString, videoID, transformSource, fatal)
}

// DownloadJSONList implements common.py/_download_json (returning lists)
func (ie *CommonIE) DownloadJSONList(url, videoID string, note, errNote OptString,
	transformSource func(string) string, fatal bool, encoding OptString, data OptString,
	headers, query map[string]interface{}) []interface{} {

	jsonString := ie.DownloadWebpageURL(url, videoID, note, errNote, fatal, 1, 5, encoding, data, headers, query)

	// TODO some fatal handling

	return ie.ParseJSONList(jsonString, videoID, transformSource, fatal)
}

// URLResult implements common.py/url_result
func (ie *CommonIE) URLResult(url string, ieKey, videoID, videoTitle OptString) map[string]interface{} {
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

// SortFormats implements common.py/_sort_formats
func (ie *CommonIE) SortFormats(formats []interface{}) {
	if len(formats) == 0 {
		panic(newExtractorError("No video formats found"))
	}

	fs := make([]map[string]interface{}, len(formats))
	for idx, format := range formats {
		f, ok := format.(map[string]interface{})
		if !ok {
			panic(newExtractorError("Format is not a dict"))
		}
		fs[idx] = f

		if _, ok := f["tbr"]; !ok {
			if abr, ok := f["abr"]; ok {
				if vbr, ok := f["vbr"]; ok {
					f["tbr"] = CastToInt(abr) + CastToInt(vbr)
				}
			}
		}

		if ext, ok := f["ext"]; !ok || ext == "" {
			if url, ok := f["url"]; ok {
				f["ext"] = DetermineExt(CastToOptString(url), "unknown_video")
			}
		}

		var preference float64
		if _preference, ok := f["preference"]; !ok {
			preference = 0
			if f["ext"] == "f4f" || f["ext"] == "f4m" {
				preference -= 0.5
			}
		} else {
			preference = CastToFloat(_preference)
		}

		protocol := DetermineProtocol(f)
		if protocol == "http" || protocol == "https" {
			f["proto_preference"] = 0.0
		} else if protocol == "rtsp" {
			f["proto_preference"] = -0.5
		} else {
			f["proto_preference"] = -0.1
		}

		if f["vcodec"] == "none" {
			preference -= 50
			order := []string{"webm", "opus", "ogg", "mp3", "aac", "m4a"}
			f["ext_preference"] = 0
			f["audio_ext_preference"] = -1
			for idx, ext := range order {
				if f["ext"] == ext {
					f["audio_ext_preference"] = idx
				}
			}
		} else {
			if f["acodec"] == "none" {
				preference -= 40
			}
			order := []string{"webm", "mp4", "flv"}
			f["audio_ext_preference"] = 0
			f["ext_preference"] = -1
			for idx, ext := range order {
				if f["ext"] == ext {
					f["ext_preference"] = idx
				}
			}
		}

		f["preference"] = preference
	}

	sortFields := []struct {
		key string
		def interface{}
	}{
		{"preference", -1.0},
		{"languagePreference", -1},
		{"quality", -1},
		{"tbr", -1},
		{"filesize", -1},
		{"vbr", -1},
		{"height", -1},
		{"width", -1},
		{"protoPreference", -1.0},
		{"extPreference", -1},
		{"abr", -1},
		{"audioExtPreference", -1},
		{"fps", -1},
		{"filesizeApprox", -1},
		{"sourcePreference", -1},
		{"formatID", ""},
	}

	sort.Slice(formats, func(i, j int) bool {
		for _, sortField := range sortFields {
			key := sortField.key
			switch def := sortField.def.(type) {
			case int:
				ival := GetIntField(fs[i], key, false, def)
				jval := GetIntField(fs[j], key, false, def)
				if ival < jval {
					return true
				} else if ival > jval {
					return false
				}
			case float64:
				ival := GetFloatField(fs[i], key, false, def)
				jval := GetFloatField(fs[j], key, false, def)
				if ival < jval {
					return true
				} else if ival > jval {
					return false
				}
			case string:
				ival := GetStringField(fs[i], key, false, def)
				jval := GetStringField(fs[j], key, false, def)
				if ival < jval {
					return true
				} else if ival > jval {
					return false
				}
			}
		}
		return false
	})
}

// RemoveDuplicateFormats implements common.py/_remove_duplicate_formats
func (ie *CommonIE) RemoveDuplicateFormats(formats []interface{}) []interface{} {
	formatURLs := map[string]struct{}{}
	uniqueFormats := make([]interface{}, 0, len(formats))
	for _, _format := range formats {
		format, ok := _format.(map[string]interface{})
		if !ok {
			panic(newExtractorError("Format is not a dict"))
		}
		url := GetStringField(format, "url", true, "")
		if _, found := formatURLs[url]; !found {
			formatURLs[url] = struct{}{}
			uniqueFormats = append(uniqueFormats, _format)
		}
	}
	return uniqueFormats
}

// RTASearch implements common.py/_rta_search
func (ie *CommonIE) RTASearch(html string) int {
	match := ReSearch(`(?ix)<meta\s+name="rating"\s+content="RTA-5042-1996-1400-1577-RTA"`, html, 0)
	if match != nil {
		return 18
	}
	return 0
}

var ratingTable = map[string]int{
	"safe for kids": 0,
	"general":       8,
	"14 years":      14,
	"mature":        17,
	"restricted":    19,
}

// MediaRatingSearch implements common.py/_media_rating_search
func (ie *CommonIE) MediaRatingSearch(html string) OptInt {
	rating := ie.HTMLSearchMetaOne("rating", html, OptString{}, false, NoDefault, 0)
	if !rating.IsSet() {
		return OptInt{}
	}
	age, ok := ratingTable[strings.ToLower(rating.Get())]
	if !ok {
		return OptInt{}
	}
	return AsOptInt(age)
}
