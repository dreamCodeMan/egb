package egb

import "net/url"

//URLEncode return encode url string.
func URLEncode(str string) string {
	return url.QueryEscape(str)
}

//URLDecode return decode url string.
func URLDecode(str string) (string, error) {
	return url.QueryUnescape(str)
}
