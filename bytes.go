package egb

import (
	"io"
	"bytes"
	"strings"
	"fmt"
)

//BytesReader convert interface{} to io.Reader
func BytesReader(data interface{}) io.Reader {
	switch s := data.(type) {
	case io.Reader:
		return s
	case []byte:
		return bytes.NewReader(s)
	case string:
		return strings.NewReader(s)
	case fmt.Stringer:
		return strings.NewReader(s.String())
	case error:
		return strings.NewReader(s.Error())
	}
	return nil
}
