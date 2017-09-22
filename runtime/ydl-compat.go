package runtime

import "net/url"

// ParseUnquote implements compat.py/compat_urllib_parse_unquote
func ParseUnquote(str string) string {
	ret, err := url.PathUnescape(str)
	if err != nil {
		panic(newExtractorError(err.Error()))
	}
	return ret
}
