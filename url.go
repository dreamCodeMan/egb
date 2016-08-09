package egb

import "net/url"

//UrlEncode return encode url string.
func UrlEncode(str string) string {
	return url.QueryEscape(str)
}

//UrlDecode return decode url string.
func UrlDecode(str string) (string, error) {
	return url.QueryUnescape(str)
}