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
 * ydl-utils.go: Re-implementation of various functions from
 *               https://github.com/rg3/youtube-dl/blob/master/youtube_dl/utils.py
 */

package runtime

import (
	"bytes"
	"fmt"
	htmlLib "html"
	"io/ioutil"
	"net/http"
	"net/url"
	urlLib "net/url"
	"strconv"
	"strings"
	"unicode"

	"github.com/tenta-browser/go-pcre-matcher/replacer"

	"github.com/tenta-browser/go-video-downloader/utils"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

type (
	// Request is the request type exposed to extractors
	Request = *http.Request
	// Response is the response type exposed to extractors
	Response = *http.Response
)

var cleanHTMLRe1, cleanHTMLRe2, cleanHTMLRe3 Regexp

// CleanHTML implements utils.py/clean_html
func CleanHTML(html OptString) OptString {
	// these expressions must be compiled lazily, since the regexp engine isn't set at init time
	if cleanHTMLRe1 == nil {
		cleanHTMLRe1 = ReMustCompile(`\s*<\s*br\s*/?\s*>\s*`, 0)
	}
	if cleanHTMLRe2 == nil {
		cleanHTMLRe2 = ReMustCompile(`<\s*/\s*p\s*>\s*<\s*p[^>]*>`, 0)
	}
	if cleanHTMLRe3 == nil {
		cleanHTMLRe3 = ReMustCompile(`<.*?>`, 0)
	}
	if !html.IsSet() {
		return html
	}
	_html := html.Get()
	_html = strings.Replace(_html, "\n", " ", -1)
	_html = cleanHTMLRe1.Replace(_html, "\n")
	_html = cleanHTMLRe2.Replace(_html, "\n")
	_html = cleanHTMLRe3.Replace(_html, "")
	_html = htmlLib.UnescapeString(_html)
	_html = strings.TrimSpace(_html)
	return AsOptString(_html)
}

// UnescapeHTML implements utils.py/unescapeHTML
func UnescapeHTML(html OptString) OptString {
	if !html.IsSet() {
		return html
	}
	return AsOptString(htmlLib.UnescapeString(html.Get()))
}

var knownExtensions = utils.MakeSet([]string{
	"mp4", "m4a", "m4p", "m4b", "m4r", "m4v", "aac",
	"flv", "f4v", "f4a", "f4b",
	"webm", "ogg", "ogv", "oga", "ogx", "spx", "opus",
	"mkv", "mka", "mk3d",
	"avi", "divx",
	"mov",
	"asf", "wmv", "wma",
	"3gp", "3g2",
	"mp3",
	"flac",
	"ape",
	"wav",
	"f4f", "f4m", "m3u8", "smil",
})

// DetermineExt implements utils.py/determine_ext
func DetermineExt(url OptString, defaultExt string) string {
	if !url.IsSet() {
		return defaultExt
	}

	// get it from before '?' and after the last '.'
	guess := url.Get()
	pos := utils.IndexRune(guess, '?')
	if pos < 0 {
		pos = len(guess)
	}
	guess = guess[:pos]
	guess = guess[utils.LastIndexRune(guess, '.')+1:]

	if match := ReMatch("^[A-Za-z0-9]+$", guess, 0); match != nil {
		return guess
	}

	guess = strings.TrimRight(guess, "/")
	if _, ok := knownExtensions[guess]; ok {
		return guess
	}

	return defaultExt
}

// taken from https://stackoverflow.com/a/26722698
func removeDiacritics(s string) string {
	isMn := func(r rune) bool {
		return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
	}
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	rs, _, _ := transform.String(t, s)
	return rs
}

// SanitizeFilename implements utils.py/sanitize_filename
func SanitizeFilename(s string, restricted bool, isID bool) string {
	// TODO handle timestamps (needs replace with func callback)

	if restricted {
		s = removeDiacritics(s)
	}

	replaceInsaneChar := func(char rune) string {
		if char == '?' || int(char) < 32 || int(char) == 127 {
			return ""
		} else if char == '"' {
			if restricted {
				return ""
			}
			return "'"
		} else if char == ':' {
			if restricted {
				return "_-"
			}
			return " -"
		} else if utils.IndexRune("\\/|*<>", char) >= 0 {
			return "_"
		} else if restricted && (utils.IndexRune("!&'()[]{}$;`^,#", char) >= 0 || unicode.IsSpace(char)) {
			return "_"
		} else if restricted && int(char) > 127 {
			return "_"
		} else {
			return string(char)
		}
	}

	var sparts []string
	for _, char := range s {
		sparts = append(sparts, replaceInsaneChar(char))
	}
	s = strings.Join(sparts, "")

	if !isID {
		for strings.Index(s, "__") >= 0 {
			s = strings.Replace(s, "__", "_", -1)
		}
		s = strings.Trim(s, "_")
		if restricted && strings.HasPrefix(s, "-_") {
			s = s[2:]
		}
		if strings.HasPrefix(s, "-") {
			s = "_" + s[1:]
		}
		s = strings.TrimLeft(s, ".")
		if s == "" {
			s = "_"
		}
	}

	return s
}

// StripOrNone implements utils.py/strip_or_none
func StripOrNone(s OptString) OptString {
	if !s.IsSet() {
		return s
	}
	return AsOptString(StrStrip(s.Get(), ""))
}

// PyExtractorError implements utils.py/ExtractorError
func PyExtractorError(msg string, expected bool, videoID OptString) error {
	if videoID.IsSet() {
		msg = videoID.Get() + ": " + msg
	}
	return newExtractorError(msg)
}

// IntOrNone implements utils.py/int_or_none
func IntOrNone(val interface{}, scale int, def OptInt, invScale int) OptInt {
	if val == nil || val == "" || IsNone(val) {
		return def
	}
	i, err := convertToInt(val)
	if err != nil {
		return def
	}
	return AsOptInt(i * invScale / scale)
}

// StrToInt implements utils.py/str_to_int
func StrToInt(s OptString) OptInt {
	if !s.IsSet() {
		return OptInt{}
	}
	ss := ReSub(`[,\.\+]`, "", s.Get(), 0, 0)
	i, err := strconv.Atoi(ss)
	if err != nil {
		panic(newExtractorError(fmt.Sprintf("Failed to convert to int: '%s'", s.Get())))
	}
	return AsOptInt(i)
}

// FloatOrNone implements utils.py/float_or_none
func FloatOrNone(val interface{}, scale, invScale int, def OptFloat) OptFloat {
	if IsNone(val) {
		return def
	}
	f, err := convertToFloat(val)
	if err != nil {
		return def
	}
	return AsOptFloat(f * float64(invScale) / float64(scale))
}

// UtilDictGet implements utils.py/dict_get
func UtilDictGet(dict SDict, key interface{}, def interface{}, skipFalseValues bool) interface{} {
	keys, ok := key.([]string)
	if !ok {
		keys = []string{key.(string)}
	}
	for _, key := range keys {
		val, ok := dict[key]
		if !ok || val == nil {
			continue
		}
		if skipFalseValues {
			// TODO possibly handle empty slices and maps too
			if val == 0 || val == "" || val == nil {
				continue
			}
		}
		return val
	}
	return def
}

// TryGetOne implements utils.py/try_get_one (for one getter)
func TryGetOne(src interface{}, getter func(idx interface{}) interface{}, expectedType interface{}) interface{} {
	return TryGetMulti(src, []func(idx interface{}) interface{}{getter}, expectedType)
}

// TryGetMulti implements utils.py/try_get_one (for multiple getters)
func TryGetMulti(src interface{}, getters []func(idx interface{}) interface{}, expectedType interface{}) interface{} {
	for _, getter := range getters {
		getterRunner := func() (val interface{}, ok bool) {
			defer func() {
				// poor mans imitation of "except (AttributeError, KeyError, TypeError, IndexError)"
				if r := recover(); r != nil {
					if e, ok := r.(*extractorError); ok && utils.StrIn(e.kind, "cast", "convert", "subscript") {
						val, ok = nil, false
					} else {
						panic(r)
					}
				}
			}()
			return getter(src), true
		}
		v, ok := getterRunner()
		if ok && (expectedType == nil || IsInstance(v, expectedType)) {
			return v
		}
	}
	return nil
}

// UnifiedStrDate implements utils.py/unified_strdate
func UnifiedStrDate(dateStr OptString, dayFirst bool) OptString {
	// TODO implement me
	return OptString{}
}

// ParseDuration implements utils.py/parse_duration
func ParseDuration(s OptString) OptString {
	// TODO implement me
	return OptString{}
}

// ParseISO8601 implements utils.py/parse_iso8601
func ParseISO8601(dateStr OptString, delimiter string) OptInt {
	// TODO implement me
	return OptInt{}
}

// DetermineProtocol implements utils.py/determine_protocol
func DetermineProtocol(infoDict SDict) string {
	if protocol := GetStringField(infoDict, "protocol", false, ""); protocol != "" {
		return protocol
	}

	url := GetStringField(infoDict, "url", true, "")

	if strings.HasPrefix(url, "rtmp") {
		return "rtmp"
	} else if strings.HasPrefix(url, "mms") {
		return "mms"
	} else if strings.HasPrefix(url, "rtsp") {
		return "rtsp"
	}

	ext := DetermineExt(AsOptString(url), "unknown_video")

	if ext == "m3u8" {
		return "m3u8"
	} else if ext == "f4m" {
		return "f4m"
	}

	purl, err := urlLib.Parse(url)
	if err != nil {
		panic(newExtractorError("Invalid URL: " + url))
	}

	return purl.Scheme
}

// RemoveStart implements utils.py/remove_start
func RemoveStart(s OptString, start string) OptString {
	if !s.IsSet() {
		return s
	}
	ss := s.Get()
	if !strings.HasPrefix(ss, start) {
		return s
	}
	return AsOptString(ss[len(start):])
}

// RemoveEnd implements utils.py/remove_end
func RemoveEnd(s OptString, end string) OptString {
	if !s.IsSet() {
		return s
	}
	ss := s.Get()
	if !strings.HasSuffix(ss, end) {
		return s
	}
	return AsOptString(ss[:len(ss)-len(end)])
}

// RemoveQuotes implements utils.py/remove_quotes
func RemoveQuotes(s OptString) OptString {
	if !s.IsSet() {
		return s
	}
	ss := s.Get()
	if len(ss) < 2 {
		return s
	}
	for _, quote := range []byte{'\'', '"'} {
		if ss[0] == quote && ss[len(ss)-1] == quote {
			return AsOptString(ss[1 : len(ss)-1])
		}
	}
	return s
}

// JsToJSON implements utils.py/js_to_json
func JsToJSON(code string) string {
	commentRe := `/\*(?:(?!\*/).)*?\*/|//[^\n]*`
	skipRe := fmt.Sprintf(`\s*(?:%s)?\s*`, commentRe)
	integerTable := []struct {
		regex string
		base  int
	}{
		{fmt.Sprintf(`(?s)^(0[xX][0-9a-fA-F]+)%s:?$`, skipRe), 16},
		{fmt.Sprintf(`(?s)^(0+[0-7]+)%s:?$`, skipRe), 8},
	}

	fixKv := func(m Match) string {
		v := m.GroupByIdx(0)
		if v == "true" || v == "false" || v == "null" {
			return v
		} else if strings.HasPrefix(v, "/*") || strings.HasPrefix(v, "//") || v == "," {
			return ""
		}

		if v[0] == '\'' || v[0] == '"' {
			v = ReMustCompile(`(?s)\\.|"`, 0).ReplaceFunc(v[1:len(v)-1], replacer.NewReplacer(func(m Match) string {
				m0 := m.GroupByIdx(0)
				switch m0 {
				case "\"":
					return "\\\""
				case "\\'":
					return "'"
				case "\\\n":
					return ""
				case "\\x":
					return "\\\\x" // no need to convert to \u, Go likes \x
				default:
					return m0
				}
			}))
		}

		for _, rb := range integerTable {
			im := ReMatch(rb.regex, v, 0)
			if im != nil {
				i64, err := strconv.ParseInt(im.GroupByIdx(1), rb.base, 0)
				if err != nil {
					panic(newExtractorError("Int conversion failed: " + err.Error()))
				}
				if strings.HasSuffix(v, ":") {
					return fmt.Sprintf(`"%d"`, i64)
				}
				return fmt.Sprintf(`%d`, i64)
			}
		}

		return fmt.Sprintf(`"%s"`, v)
	}

	return ReMustCompile(fmt.Sprintf(`(?sx)
		"(?:[^"\\]*(?:\\\\|\\['"nurtbfx/\n]))*[^"\\]*"|
		'(?:[^'\\]*(?:\\\\|\\['"nurtbfx/\n]))*[^'\\]*'|
		%[1]s|,(?=%[2]s[\]}}])|
		[a-zA-Z_][.a-zA-Z_0-9]*|
		\b(?:0[xX][0-9a-fA-F]+|0+[0-7]+)(?:%[2]s:)?|
		[0-9]+(?=%[2]s:)`,
		commentRe, skipRe), 0).ReplaceFunc(code, replacer.NewReplacer(fixKv))
}

// StdHeaders implements utils.py/std_headers
var StdHeaders = SDict{
	"User-Agent":      "Mozilla/5.0 (X11; Linux x86_64; rv:10.0) Gecko/20150101 Firefox/47.0 (Chrome)",
	"Accept-Charset":  "ISO-8859-1,utf-8;q=0.7,*;q=0.7",
	"Accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
	"Accept-Encoding": "gzip, deflate",
	"Accept-Language": "en-us,en;q=0.5",
}

// SanitizedRequest implements utils.py/sanitized_Request
func SanitizedRequest(url string, data []byte, headers SDict) Request {
	req, err := http.NewRequest("GET", utils.SanitizeURL(url), nil)
	if err != nil {
		panic(newExtractorError(err.Error()))
	}
	updateRequest(req, data, headers, nil)
	return req
}

func updateRequest(req Request, data []byte, headers, query SDict) {
	// apply headers
	for headerName := range headers {
		RequestAddHeader(req, headerName, ConvertToString(DictGet(headers, headerName, nil)))
	}

	// apply body data
	if data != nil {
		req.Body = ioutil.NopCloser(bytes.NewBuffer(data))
		req.ContentLength = int64(len(data))
	}

	// apply query parameters
	// Note: even if there are no new parameters to add, check if the existing
	// parameters are properly encoded, if not, rebuild the query;
	// Todo: We could just simply rebuild it every time, but rebuilding changes the order
	// of keys and some sites break. The ideal solution would escape existing keys and
	// add new ones without changing the order.
	{
		q := req.URL.Query()
		needsReencoding := false
	loop:
		for queryKey, queryVals := range q {
			if queryKey != url.QueryEscape(queryKey) {
				needsReencoding = true
				break
			}
			for _, queryVal := range queryVals {
				if queryVal != url.QueryEscape(queryVal) {
					needsReencoding = true
					break loop
				}
			}
		}
		for queryKey := range query {
			q.Set(queryKey, ConvertToString(DictGet(query, queryKey, nil)))
			needsReencoding = true
		}
		if needsReencoding {
			req.URL.RawQuery = q.Encode()
		}
	}
}

// RequestAddHeader implements python/Request.add_header
func RequestAddHeader(req Request, key, val string) {
	req.Header.Set(key, val)
}

// ResponseGetURL implements python/Response.geturl
func ResponseGetURL(res Response) string {
	return res.Request.URL.String()
}

// parseM3U8Attributes implements utils.py/parse_m3u8_attributes
func parseM3U8Attributes(attrib string) map[string]string {
	info := make(map[string]string)
	for _, kv := range ReFindAllMulti(`(?P<key>[A-Z0-9-]+)=(?P<val>"[^"]+"|[^",]+)(?:,|$)`, attrib, 0) {
		key := kv[0]
		val := kv[1]
		if val[0] == '"' {
			val = val[1 : len(val)-1]
		}
		info[key] = val
	}
	return info
}

// parseCodecs implements utils.py/parse_codecs
func parseCodecs(codecsStr string) map[string]OptString {
	if codecsStr == "" {
		return make(map[string]OptString)
	}
	splitedCodecs := []string{}
	for _, str := range strings.Split(strings.Trim(strings.TrimSpace(codecsStr), ","), ",") {
		str = strings.TrimSpace(str)
		if str != "" {
			splitedCodecs = append(splitedCodecs, str)
		}
	}
	vcodec := OptString{}
	acodec := OptString{}
	for _, fullCodec := range splitedCodecs {
		codec := strings.Split(fullCodec, ".")[0]
		if utils.StrIn(codec, "avc1", "avc2", "avc3", "avc4", "vp9", "vp8", "hev1", "hev2", "h263", "h264", "mp4v", "hvc1") {
			if vcodec.GetOrDef("") == "" {
				vcodec = AsOptString(fullCodec)
			}
		} else if utils.StrIn(codec, "mp4a", "opus", "vorbis", "mp3", "aac", "ac-3", "ec-3", "eac3", "dtsc", "dtse", "dtsh", "dtsl") {
			if acodec.GetOrDef("") == "" {
				acodec = AsOptString(fullCodec)
			}
		} else {
			if utils.Debug {
				utils.Log("WARNING: Unknown codec %s", fullCodec)
			}
		}
	}
	if vcodec.GetOrDef("") == "" && acodec.GetOrDef("") == "" {
		if len(splitedCodecs) == 2 {
			return map[string]OptString{
				"vcodec": vcodec,
				"acodec": acodec,
			}
		} else if len(splitedCodecs) == 1 {
			return map[string]OptString{
				"vcodec": AsOptString("none"),
				"acodec": vcodec,
			}
		}
	} else {
		_or := func(ostr OptString, def string) OptString {
			if ostr.GetOrDef("") == "" {
				return AsOptString(def)

			}
			return ostr
		}
		return map[string]OptString{
			"vcodec": _or(vcodec, "none"),
			"acodec": _or(acodec, "none"),
		}
	}
	return make(map[string]OptString)
}

// Qualities implements utils.py/qualities
func Qualities(qualityIDs []string) func(string) int {
	return func(qID string) int {
		for qIdx, _qID := range qualityIDs {
			if _qID == qID {
				return qIdx
			}
		}
		return -1
	}
}

// UtilURLJoin implements utils.py/urljoin
func UtilURLJoin(base interface{}, path interface{}) OptString {
	var paths, bases string
	var ok bool
	if pathb, ok := path.([]byte); ok {
		path = BytesToStr(pathb, "utf-8")
	}
	if paths, ok = path.(string); !ok || paths == "" {
		return OptString{}
	}
	if ReMatch("^(?:https?:)?//", paths, 0) != nil {
		return AsOptString(paths)
	}
	if baseb, ok := base.([]byte); ok {
		base = BytesToStr(baseb, "utf-8")
	}
	if bases, ok = base.(string); !ok || ReMatch("^(?:https?:)?//", bases, 0) == nil {
		return OptString{}
	}
	return AsOptString(URLJoin(bases, AsOptString(paths)))
}

// FindXPathAttr implements utils.py/find_xpath_attr
func FindXPathAttr(node XMLElement, xpath string, key string, val OptString) XMLElement {
	if utils.Debug {
		if ReMatch("^[a-zA-Z_-]+$", key, 0) == nil {
			panic(newExtractorError("Invalid attribute: " + key))
		}
	}
	if val.IsSet() {
		xpath = xpath + fmt.Sprintf("[@%s='%s']", key, val.Get())
	} else {
		xpath = xpath + fmt.Sprintf("[@%s]", key)
	}
	return XMLFind(node, xpath)
}

// XPathElementOne implements utils.py/xpath_element (for a single xpath)
func XPathElementOne(node XMLElement, xpath string, name OptString, fatal bool, def interface{}) XMLElement {
	return XPathElementMulti(node, []string{xpath}, name, fatal, def)
}

// XPathElementMulti implements utils.py/xpath_element (for multiple xpaths)
func XPathElementMulti(node XMLElement, xpaths []string, name OptString, fatal bool, def interface{}) XMLElement {
	var n XMLElement
	for _, xpath := range xpaths {
		n = XMLFind(node, xpath)
		if n != nil {
			break
		}
	}
	if n == nil {
		if !isNoDefault(def) {
			return def.(XMLElement)
		} else if fatal {
			_name := name.GetOrDef(fmt.Sprintf("%v", xpaths))
			panic(newExtractorError(fmt.Sprintf("Could not find XML element %s", _name)))
		} else {
			return nil
		}
	}
	return n
}

// XPathTextOne implements utils.py/xpath_text (for a single xpath)
func XPathTextOne(node XMLElement, xpath string, name OptString, fatal bool, def interface{}) OptString {
	return XPathTextMulti(node, []string{xpath}, name, fatal, def)
}

// XPathTextMulti implements utils.py/xpath_element (for multiple xpaths)
func XPathTextMulti(node XMLElement, xpaths []string, name OptString, fatal bool, def interface{}) OptString {
	var _def interface{}
	if isNoDefault(def) {
		_def = def
	} else {
		_def = NewEmptyXMLElement()
	}
	n := XPathElementMulti(node, xpaths, name, fatal, _def)
	if n == nil || n == _def {
		if n == nil {
			return OptString{}
		}
		return AsOptString(def.(string))
	}
	if n.Text == "" {
		if !isNoDefault(def) {
			if def == nil {
				return OptString{}
			}
			return AsOptString(def.(string))
		} else if fatal {
			_name := name.GetOrDef(fmt.Sprintf("%v", xpaths))
			panic(newExtractorError(fmt.Sprintf("Could not find XML element %s", _name)))
		} else {
			return OptString{}
		}
	}
	return AsOptString(n.Text)
}

// XPathAttr implements utils.py/xpath_attr
func XPathAttr(node XMLElement, xpath, key string, name OptString, fatal bool, def interface{}) OptString {
	n := FindXPathAttr(node, xpath, key, OptString{})
	if n == nil {
		if !isNoDefault(def) {
			if def == nil {
				return OptString{}
			}
			return AsOptString(def.(string))
		} else if fatal {
			_name := name.GetOrDef(fmt.Sprintf("%s[@%s]", xpath, key))
			panic(newExtractorError(fmt.Sprintf("Could not find XML attribute %s", _name)))
		} else {
			return OptString{}
		}
	}
	return AsOptString(n.Attributes[key].(string))
}
