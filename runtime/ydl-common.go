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

	"github.com/tenta-browser/go-pcre-matcher"
	"github.com/tenta-browser/go-video-downloader/utils"
)

// MatchID implements common.py/_match_id
func MatchID(ie InfoExtractor, url string) string {
	match := ReMatch(Re, ie.ValidUrl(), url, 0)
	if match == nil || !match.GroupPresentByName("id") {
		panic(newExtractorError("Failed to match id"))
	}
	return match.GroupByName("id")
}

// DownloadWebpage implements common.py/_download_webpage
// TODO handle all the params!
func DownloadWebpage(ie InfoExtractor, url, videoID string, note, errNote OptString,
	fatal bool, tries int, timeout int, encoding OptString, data OptString,
	headers, query map[string]interface{}) string {

	utils.Log("[%s] %s (%s)", videoID, note.GetOrDef("Downloading webpage"), url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(newExtractorError(err.Error()))
	}

	for headerName, headerVal := range ie.Ctx().Headers {
		if len(headerVal) > 0 {
			req.Header.Set(headerName, headerVal)
		}
	}
	for headerName := range headers {
		req.Header.Set(headerName, GetStringField(headers, headerName, true, ""))
	}

	res, err := ie.Ctx().Client.Do(req)
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

// SearchRegex implements common.py/_search_regex
func SearchRegex(ie InfoExtractor, pattern interface{}, str, name string,
	def interface{}, fatal bool, flags int, group interface{}) OptString {

	patterns, ok := pattern.([]string)
	if !ok {
		patterns = []string{pattern.(string)}
	}

	var match matcher.Match
	for _, pattern := range patterns {
		if match = ReSearch(Re, pattern, str, flags); match != nil {
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

// ParseJSONList implements common.py/_parse_json (returning lists)
func ParseJSONList(ie InfoExtractor, jsonString string, videoID string, transformSource int, fatal bool) []interface{} {
	// TODO handle transformSource constants
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
func DownloadJSON(ie InfoExtractor, url, videoID string, note, errNote OptString,
	transformSource int, fatal bool, encoding OptString, data OptString,
	headers, query map[string]interface{}) map[string]interface{} {

	jsonString := DownloadWebpage(ie, url, videoID, note, errNote, fatal, 1, 5, encoding, data, headers, query)

	// TODO some fatal handling

	return ParseJSON(ie, jsonString, videoID, transformSource, fatal)
}

// DownloadJSONList implements common.py/_download_json (returning lists)
func DownloadJSONList(ie InfoExtractor, url, videoID string, note, errNote OptString,
	transformSource int, fatal bool, encoding OptString, data OptString,
	headers, query map[string]interface{}) []interface{} {

	jsonString := DownloadWebpage(ie, url, videoID, note, errNote, fatal, 1, 5, encoding, data, headers, query)

	// TODO some fatal handling

	return ParseJSONList(ie, jsonString, videoID, transformSource, fatal)
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

// SortFormats implements common.py/_sort_formats
func SortFormats(ie InfoExtractor, formats []interface{}) {
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

// RTASearch implements common.py/_rta_search
func RTASearch(ie InfoExtractor, html string) int {
	match := ReSearch(Re, `(?ix)<meta\s+name="rating"\s+content="RTA-5042-1996-1400-1577-RTA"`, html, 0)
	if match != nil {
		return 18
	}
	return 0
}
