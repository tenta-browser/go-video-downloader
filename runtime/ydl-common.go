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
	"compress/zlib"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"sort"
	"strings"
	"time"

	"github.com/tenta-browser/go-pcre-matcher"
	"github.com/tenta-browser/go-video-downloader/utils"
)

var (
	// SleepBeforeRequest is the duration to sleep before doing an HTTP request
	// The tester sets this to avoid being banned from sites for doing requests too frequently
	SleepBeforeRequest time.Duration
)

// CommonIE represents the common extractor base
type CommonIE struct {
	Context  *Context
	VALIDURL string
	TESTS    []SDict
	TEST     SDict
}

// NewCommonIE constructs a new common extractor base
func NewCommonIE() *CommonIE {
	return &CommonIE{}
}

// SetContext sets the runtime extractor context
func (ie *CommonIE) SetContext(context *Context) {
	ie.Context = context
}

// ValidURL returns the regexp matching URLs supported by this extractor
func (ie *CommonIE) ValidURL() string {
	return ie.VALIDURL
}

// Tests returns the test cases for this extractor
func (ie *CommonIE) Tests() []SDict {
	if len(ie.TEST) > 0 {
		return []SDict{ie.TEST}
	}
	return ie.TESTS
}

func (ie *CommonIE) log(msg string, args ...interface{}) {
	if utils.LogLevel >= utils.StdLogLevel {
		utils.Log("[%s] %s", ie.Context.ExtractorKey, fmt.Sprintf(msg, args...))
	}
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
	fatal bool, tries int, timeout int, encoding OptString, data []byte,
	headers, query SDict, expectedStatus interface{}) string {

	return ie.DownloadWebpageRequest(SanitizedRequest(url, nil, nil), videoID, note, errNote,
		fatal, tries, timeout, encoding, data, headers, query, expectedStatus)
}

// DownloadWebpageRequest implements common.py/_download_webpage (for requests)
// TODO handle all the params: tries, timeout
func (ie *CommonIE) DownloadWebpageRequest(req Request, videoID string, note, errNote OptString,
	fatal bool, tries int, timeout int, encoding OptString, data []byte,
	headers, query SDict, expectedStatus interface{}) string {

	return ie.DownloadWebpageHandleRequest(req, videoID, note, errNote, fatal, encoding, data, headers, query, expectedStatus).Φ0
}

// DownloadWebpageHandleResponse is tuple returned by the DownloadWebpageHandle* functions
// consisting of the resulting content and the HTTP response object
type DownloadWebpageHandleResponse = struct {
	Φ0 string
	Φ1 Response
}

// DownloadWebpageHandleURL implements common.py/_download_webpage_handle (for urls)
func (ie *CommonIE) DownloadWebpageHandleURL(url, videoID string, note, errNote OptString,
	fatal bool, encoding OptString, data []byte, headers, query SDict, expectedStatus interface{}) DownloadWebpageHandleResponse {

	return ie.DownloadWebpageHandleRequest(SanitizedRequest(url, nil, nil), videoID, note, errNote,
		fatal, encoding, data, headers, query, expectedStatus)
}

// DownloadWebpageHandleRequest implements common.py/_download_webpage_handle (for requests)
// TODO handle all the params: fatal
func (ie *CommonIE) DownloadWebpageHandleRequest(req Request, videoID string, note, errNote OptString,
	fatal bool, encoding OptString, data []byte, headers, query SDict, expectedStatus interface{}) DownloadWebpageHandleResponse {

	res, err := ie.requestWebpage(req, false, videoID, note, errNote, data, headers, query, expectedStatus)
	if err != nil {
		panic(newExtractorError(err.Error()))
	}
	defer res.Body.Close()

	// TODO workaround for https://github.com/golang/go/issues/18779
	// remove this in Go 1.10
	reader := res.Body
	if res.Header.Get("Content-Encoding") == "deflate" {
		readerZ, errZ := zlib.NewReader(res.Body)
		defer readerZ.Close()
		if errZ == nil {
			reader = readerZ
		}
	}

	ret, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(newExtractorError(err.Error()))
	}

	retstr := BytesToStr(ret, encoding.GetOrDef("utf-8"))

	return DownloadWebpageHandleResponse{retstr, res}
}

// RequestWebpageURL implements common.py/_request_webpage (for urls)
func (ie *CommonIE) RequestWebpageURL(url, videoID string, note, errNote OptString,
	fatal bool, data []byte, headers, query SDict, expectedStatus interface{}) Response {

	return ie.RequestWebpageRequest(SanitizedRequest(url, nil, nil), videoID, note, errNote,
		fatal, data, headers, query, expectedStatus)
}

// RequestWebpageRequest implements common.py/_request_webpage (for requests)
func (ie *CommonIE) RequestWebpageRequest(req Request, videoID string, note, errNote OptString,
	fatal bool, data []byte, headers, query SDict, expectedStatus interface{}) Response {

	res, err := ie.requestWebpage(req, true, videoID, note, errNote, data, headers, query, expectedStatus)
	if err != nil {
		panic(newExtractorError(err.Error()))
	}

	return res
}

func (ie *CommonIE) requestWebpage(req Request, closeBody bool, videoID string, note, errNote OptString,
	data []byte, headers, query SDict, expectedStatus interface{}) (Response, error) {

	ie.log("%s: %s (%s)", videoID, note.GetOrDef("Downloading webpage"), req.URL)

	if SleepBeforeRequest > 0 {
		time.Sleep(SleepBeforeRequest)
	}

	// apply headers, priority: StdHeaders < context < request < args
	baseHeaders := make(map[string]string)
	for headerName, headerVal := range StdHeaders {
		baseHeaders[headerName] = headerVal.(string)
	}
	for headerName, headerVal := range ie.Context.Headers {
		baseHeaders[headerName] = headerVal
	}
	for headerName, headerVal := range baseHeaders {
		_, found := req.Header[headerName]
		if !found {
			RequestAddHeader(req, headerName, headerVal)
		}
	}

	// apply data, headers, query args
	updateRequest(req, data, headers, query)

	// remove 'Accept-Encoding' header, let the HTTP client set its default
	req.Header.Del("Accept-Encoding")

	res, err := ie.Context.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 && !canAcceptStatusCode(res.StatusCode, expectedStatus) {
		res.Body.Close()
		return nil, fmt.Errorf("Unable to download webpage: HTTP Error %d: %s", res.StatusCode, res.Status)
	}

	if closeBody {
		res.Body.Close()
	}

	return res, nil
}

// canAcceptStatusCode implements common.py/__can_accept_status_code
func canAcceptStatusCode(actual int, expected interface{}) bool {
	if expected == nil {
		return false
	}
	switch rexpected := reflect.ValueOf(expected); rexpected.Kind() {
	case reflect.Slice:
		// find actual in expected list
		for idx := 0; idx < rexpected.Len(); idx++ {
			if CastToInt(rexpected.Index(idx).Interface()) == actual {
				return true
			}
		}
		return false
	default:
		// find actual in expected tuple
		if IsTuple(rexpected) {
			for idx := 0; idx < rexpected.NumField(); idx++ {
				if CastToInt(rexpected.Field(idx).Interface()) == actual {
					return true
				}
			}
			return false
		}
		panic(newExtractorError(fmt.Sprintf("Invalid type passed to CanAcceptStatusCode: %T", expected)))
	}
}

// GetLoginInfoResult is tuple returned by the GetLoginInfo function
// consisting of a username and a password
type GetLoginInfoResult = struct {
	Φ0 OptString
	Φ1 OptString
}

// GetLoginInfo implements common.py/_get_login_info
func (ie *CommonIE) GetLoginInfo() GetLoginInfoResult {
	credentials := ie.Context.Credentials
	if credentials == nil {
		return GetLoginInfoResult{
			OptString{},
			OptString{},
		}
	}
	return GetLoginInfoResult{
		AsOptString(credentials.Username),
		AsOptString(credentials.Password),
	}
}

// SearchRegexOne implements common.py/_search_regex (for single pattern)
func (ie *CommonIE) SearchRegexOne(pattern string, str, name string,
	def interface{}, fatal bool, flags int, group interface{}) OptString {

	return ie.SearchRegexMulti([]string{pattern}, str, name, def, fatal, flags, group)
}

// SearchRegexMulti implements common.py/_search_regex (for multiple patterns)
func (ie *CommonIE) SearchRegexMulti(patterns []string, str, name string,
	def interface{}, fatal bool, flags int, group interface{}) OptString {

	var match Match
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
		ie.log("WARNING: Unable to extract %v", name)
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
func (ie *CommonIE) ParseJSON(jsonString string, videoID string, transformSource func(string) string, fatal bool) interface{} {
	if transformSource != nil {
		jsonString = transformSource(jsonString)
	}
	result, err := jsonLoads(jsonString)
	if err != nil {
		errMsg := fmt.Sprintf("[%v] Failed to parse JSON", videoID)
		if fatal {
			panic(newExtractorError(errMsg))
		} else {
			ie.log(errMsg)
			return nil
		}
	}
	return utils.FixJSONFloats(result)
}

// DownloadJSON implements common.py/_download_json
func (ie *CommonIE) DownloadJSON(url, videoID string, note, errNote OptString,
	transformSource func(string) string, fatal bool, encoding OptString, data []byte,
	headers, query SDict, expectedStatus interface{}) interface{} {

	jsonString := ie.DownloadWebpageURL(url, videoID, note, errNote, fatal, 1, 5, encoding, data, headers, query, expectedStatus)

	// TODO some fatal handling

	return ie.ParseJSON(jsonString, videoID, transformSource, fatal)
}

// ParseXML implements common.py/_parse_xml
func (ie *CommonIE) ParseXML(xmlString string, videoID string, transformSource func(string) string, fatal bool) XMLElement {
	if transformSource != nil {
		xmlString = transformSource(xmlString)
	}
	root, err := XMLParse(xmlString)
	if err != nil {
		errMsg := fmt.Sprintf("[%v] Failed to parse XML", videoID)
		if fatal {
			panic(newExtractorError(errMsg))
		} else {
			ie.log(errMsg)
			return nil
		}
	}
	return root
}

// DownloadXML implements common.py/_download_xml
func (ie *CommonIE) DownloadXML(url, videoID string, note, errNote OptString,
	transformSource func(string) string, fatal bool, encoding OptString, data []byte,
	headers, query SDict, expectedStatus interface{}) XMLElement {

	// TODO handle return False
	res := ie.DownloadXMLHandle(url, videoID, note, errNote, transformSource, fatal, encoding, data, headers, query, expectedStatus)
	return res.Φ0
}

// DownloadXMLHandleResponse is tuple returned by the DownloadXMLHandle* functions
// consisting of the resulting XML tree and the HTTP response object
type DownloadXMLHandleResponse = struct {
	Φ0 XMLElement
	Φ1 Response
}

// DownloadXMLHandle implements common.py/_download_xml_handle
func (ie *CommonIE) DownloadXMLHandle(url, videoID string, note, errNote OptString,
	transformSource func(string) string, fatal bool, encoding OptString, data []byte,
	headers, query SDict, expectedStatus interface{}) DownloadXMLHandleResponse {

	// TODO handle return False
	res := ie.DownloadWebpageHandleURL(url, videoID, note, errNote, fatal, encoding, data, headers, query, expectedStatus)

	xmlString := res.Φ0
	urlh := res.Φ1

	return DownloadXMLHandleResponse{ie.ParseXML(xmlString, videoID, transformSource, fatal), urlh}
}

// URLResult implements common.py/url_result
func (ie *CommonIE) URLResult(url string, ieKey, videoID, videoTitle OptString) SDict {
	result := SDict{
		"_type":  "url",
		"url":    url,
		"ie_key": ieKey,
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
func (ie *CommonIE) SortFormats(formats []SDict) []SDict {
	if len(formats) == 0 {
		panic(newExtractorError("No video formats found"))
	}

	fs := make([]SDict, len(formats))
	for idx, f := range formats {
		fs[idx] = f

		if _, ok := f["tbr"]; !ok {
			if abr, ok := f["abr"]; ok {
				if vbr, ok := f["vbr"]; ok {
					f["tbr"] = ConvertToFloat(abr) + ConvertToFloat(vbr)
				}
			}
		} else {
			f["tbr"] = ConvertToFloat(f["tbr"])
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
			preference = ConvertToFloat(_preference)
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
		{"tbr", -1.0},
		{"filesize", -1},
		{"vbr", -1},
		{"height", -1},
		{"width", -1},
		{"protoPreference", -1.0},
		{"extPreference", -1},
		{"abr", -1},
		{"audioExtPreference", -1},
		{"fps", -1.0},
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

	return formats
}

// CheckFormats implements common.py/_check_formats
func (ie *CommonIE) CheckFormats(formats []SDict, videoID string) []SDict {
	newFormats := make([]SDict, 0, len(formats))
	for _, format := range formats {
		isValid := ie.IsValidURL(
			GetStringField(format, "url", true, ""),
			videoID,
			fmt.Sprintf("%s video format", GetStringField(format, "format_id", false, "video")),
			nil)
		if isValid {
			newFormats = append(newFormats, format)
		}
	}
	return newFormats
}

// RemoveDuplicateFormats implements common.py/_remove_duplicate_formats
func (ie *CommonIE) RemoveDuplicateFormats(formats []SDict) []SDict {
	formatURLs := map[string]struct{}{}
	uniqueFormats := make([]SDict, 0, len(formats))
	for _, format := range formats {
		url := GetStringField(format, "url", true, "")
		if _, found := formatURLs[url]; !found {
			formatURLs[url] = struct{}{}
			uniqueFormats = append(uniqueFormats, format)
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

// ExtractM3U8Formats implements common.py/_extract_m3u8_formats
func (ie *CommonIE) ExtractM3U8Formats(m3u8URL, videoID string, ext, entryProtocol OptString, preference interface{},
	m3u8ID, note, errnote OptString, fatal bool, live bool) []SDict {

	res := ie.DownloadWebpageHandleURL(m3u8URL, videoID,
		AsOptString(note.GetOrDef("Downloading m3u8 information")),
		AsOptString(errnote.GetOrDef("Failed to download m3u8 information")),
		fatal, OptString{}, nil, nil, nil, nil)

	m3u8Doc := res.Φ0
	m3u8URL = ResponseGetURL(res.Φ1)

	return ie.ParseM3U8Formats(m3u8Doc, m3u8URL, ext, entryProtocol, preference, m3u8ID, live)
}

// ParseM3U8Formats implements common.py/_parse_m3u8_formats
func (ie *CommonIE) ParseM3U8Formats(m3u8Doc, m3u8URL string, ext, entryProtocol OptString, preference interface{},
	m3u8ID OptString, live bool) []SDict {

	// TODO handle int/float casting properly
	switch pref := preference.(type) {
	case nil:
		preference = OptFloat{}
	case int:
		preference = AsOptFloat(float64(pref))
	case OptInt:
		if pref.IsSet() {
			preference = AsOptFloat(float64(pref.Get()))
		} else {
			preference = OptFloat{}
		}
	case float64:
		preference = AsOptFloat(pref)
	case OptFloat:
		// great!
	default:
		panic("Preference should be int or float")
	}

	if strings.Contains(m3u8Doc, "#EXT-X-FAXS-CM:") { // Adobe Flash Access
		return []SDict{}
	}

	if ReSearch(`#EXT-X-SESSION-KEY:.*?URI="skd://`, m3u8Doc, 0) != nil { // Apple FairPlay
		return []SDict{}
	}

	formats := []SDict{}

	formatURL := func(u string) string {
		if ReMatch("^https?://", u, 0) != nil {
			return u
		}
		return URLJoin(m3u8URL, AsOptString(u))
	}

	if strings.Contains(m3u8Doc, "#EXT-X-TARGETDURATION") {
		return []SDict{{
			"url":        m3u8URL,
			"format_id":  m3u8ID,
			"ext":        ext,
			"protocol":   entryProtocol,
			"preference": preference,
		}}
	}

	groups := make(map[string][]map[string]string)
	lastStreamInf := make(map[string]string)

	extractMedia := func(xMediaLine string) {
		media := parseM3U8Attributes(xMediaLine)
		mediaType := media["TYPE"]
		groupID := media["GROUP-ID"]
		name := media["NAME"]
		if mediaType == "" || groupID == "" || name == "" {
			return
		}
		groups[groupID] = append(groups[groupID], media)
		if !utils.StrIn(mediaType, "VIDEO", "AUDIO") {
			return
		}
		mediaURL := media["URI"]
		if mediaURL != "" {
			formatID := []string{}
			if m3u8ID.GetOrDef("") != "" {
				formatID = append(formatID, m3u8ID.Get())
			}
			if groupID != "" {
				formatID = append(formatID, groupID)
			}
			if name != "" {
				formatID = append(formatID, name)
			}
			f := SDict{
				"format_id":    strings.Join(formatID, "-"),
				"url":          formatURL(mediaURL),
				"manifest_url": m3u8URL,
				"language":     media["LANGUAGE"],
				"ext":          ext,
				"protocol":     entryProtocol,
				"preference":   preference,
			}
			if mediaType == "AUDIO" {
				f["vcodec"] = "none"
			}
			formats = append(formats, f)
		}
	}

	buildStreamName := func() string {
		streamName := lastStreamInf["NAME"]
		if streamName != "" {
			return streamName
		}
		streamGroupID := lastStreamInf["VIDEO"]
		if streamGroupID == "" {
			return ""
		}
		streamGroup := groups[streamGroupID]
		if len(streamGroup) == 0 {
			return streamGroupID
		}
		rendition := streamGroup[0]
		if renditionName := rendition["NAME"]; renditionName != "" {
			return renditionName
		}
		return streamGroupID
	}

	for _, line := range strings.Split(m3u8Doc, "\n") {
		if strings.HasPrefix(line, "#EXT-X-STREAM-INF:") {
			lastStreamInf = parseM3U8Attributes(line)
		} else if strings.HasPrefix(line, "#EXT-X-MEDIA:") {
			extractMedia(line)
		} else if strings.HasPrefix(line, "#") || strings.TrimSpace(line) == "" {
			continue
		} else {
			var tbr OptFloat
			{
				averageBandwidth := lastStreamInf["AVERAGE-BANDWIDTH"]
				bandwidth := lastStreamInf["BANDWIDTH"]
				if averageBandwidth != "" {
					tbr = FloatOrNone(averageBandwidth, 1000, 1, OptFloat{})
				} else {
					tbr = FloatOrNone(bandwidth, 1000, 1, OptFloat{})
				}
			}
			formatID := []string{}
			if m3u8ID.GetOrDef("") != "" {
				formatID = append(formatID, m3u8ID.Get())
			}
			streamName := buildStreamName()
			if !live {
				if streamName != "" {
					formatID = append(formatID, streamName)
				} else if tbr.GetOrDef(0) != 0 {
					formatID = append(formatID, fmt.Sprintf("%v", int(tbr.Get())))
				} else {
					formatID = append(formatID, fmt.Sprintf("%v", len(formats)))
				}
			}
			manifestURL := formatURL(strings.TrimSpace(line))
			f := SDict{
				"format_id":    strings.Join(formatID, "-"),
				"url":          manifestURL,
				"manifest_url": m3u8URL,
				"tbr":          tbr,
				"ext":          ext,
				"fps":          FloatOrNone(lastStreamInf["FRAME-RATE"], 1, 1, OptFloat{}),
				"protocol":     entryProtocol,
				"preference":   preference,
			}
			resolution := lastStreamInf["RESOLUTION"]
			if resolution != "" {
				mobj := ReSearch(`(?P<width>\d+)[xX](?P<height>\d+)`, resolution, 0)
				if mobj != nil {
					f["width"] = ConvertToInt(mobj.GroupByName("width"))
					f["height"] = ConvertToInt(mobj.GroupByName("height"))
				}
			}
			mobj := ReSearch(`audio.*?(?:%3D|=)(\d+)(?:-video.*?(?:%3D|=)(\d+))?`, GetStringField(f, "url", true, ""), 0)
			if mobj != nil {
				abr := FloatOrNone(mobj.GroupByIdx(1), 1000, 1, OptFloat{})
				vbr := FloatOrNone(mobj.GroupByIdx(2), 1000, 1, OptFloat{})
				f["abr"] = abr
				f["vbr"] = vbr
			}
			codecs := parseCodecs(lastStreamInf["CODECS"])
			for key, val := range codecs {
				f[key] = val
			}
			audioGroupID := lastStreamInf["AUDIO"]
			if audioGroupID != "" && len(codecs) > 0 && f["vcodec"] != "none" {
				audioGroup := groups[audioGroupID]
				if len(audioGroup) > 0 && audioGroup[0]["URI"] != "" {
					f["acodec"] = "none"
				}
			}
			formats = append(formats, f)
			lastStreamInf = make(map[string]string)
		}
	}

	return formats
}

// HTTPScheme implements common.py/http_scheme
func (ie *CommonIE) HTTPScheme() string {
	return "https:"
}

// ProtoRelativeURL implements common.py/_proto_relative_url
func (ie *CommonIE) ProtoRelativeURL(url, scheme OptString) OptString {
	if !url.IsSet() {
		return url
	}
	_url := url.Get()
	if strings.HasPrefix(_url, "//") {
		_scheme := scheme.GetOrDef(ie.HTTPScheme())
		return AsOptString(_scheme + _url)
	}
	return url
}

// IsValidURL implements common.py/_is_valid_url
func (ie *CommonIE) IsValidURL(url, videoID, item string, headers SDict) bool {
	url = ie.ProtoRelativeURL(AsOptString(url), AsOptString("http:")).Get()
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		return true
	}
	_, err := ie.requestWebpage(SanitizedRequest(url, nil, nil), true, videoID,
		AsOptString(fmt.Sprintf("Checking %s URL", item)), OptString{}, nil, headers, nil, nil)
	if err != nil {
		ie.log("%s: %s URL is invalid, skipping", videoID, item)
		return false
	}
	return true
}

// SetCookie implements common.py/_set_cookie
func (ie *CommonIE) SetCookie(domain, name, value string) {
	// create artificial source URL from domain
	u := &url.URL{
		Scheme: "http",
		Host:   domain,
	}

	// create cookie
	c := &http.Cookie{
		Name:   name,
		Value:  value,
		Domain: domain,
	}

	ie.Context.Client.Jar.SetCookies(u, []*http.Cookie{c})
}

// LiveTitle implements common.py/_live_title
func (ie *CommonIE) LiveTitle(name string) string {
	nowStr := time.Now().Format("2006-01-02 15:04")
	return name + " " + nowStr
}
