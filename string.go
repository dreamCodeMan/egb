package egb

import (
	"encoding/json"
	"strings"
	"bytes"
	"crypto/md5"
	"fmt"
	"encoding/hex"
	"crypto/sha1"
	"crypto/sha256"
	"strconv"
	"crypto/aes"
	"encoding/base64"
	"io"
	"crypto/cipher"
	"crypto/rand"
	"errors"
)

//StringMapFunc execute a function on each element of the slice of string.
func StringMapFunc(f func(string) string, data []string) []string {
	size := len(data)
	result := make([]string, size, size)
	for i := 0; i < size; i++ {
		result[i] = f(data[i])
	}
	return result
}

//StringMarshalJSON marshals data to an json string.
func StringMarshalJSON(data interface{}) string {
	buffer, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(buffer)
}

//StringListContains judge if the slice of string contains the given string element.
func StringListContains(list []string, str string) bool {
	for i := range list {
		if list[i] == str {
			return true
		}
	}
	return false
}

//StringListContainsCaseInsensitive judge if the slice of string contains the given string element case-insensitive.
func StringListContainsCaseInsensitive(list []string, str string) bool {
	str = strings.ToLower(str)
	for i := range list {
		if strings.ToLower(list[i]) == str {
			return true
		}
	}
	return false
}

// StringStripHTMLTags strips HTML/XML tags from text.
func StringStripHTMLTags(text string) (plainText string) {
	var buf *bytes.Buffer
	tagClose := -1
	tagStart := -1
	for i, char := range text {
		//html标签的开始标志
		if char == '<' {
			if buf == nil {
				buf = bytes.NewBufferString(text)
				buf.Reset()
			}
			buf.WriteString(text[tagClose + 1 : i])
			tagStart = i
			//html标签的结束标志并且start不为-1,说明已经存在开始标签
		} else if char == '>' && tagStart != -1 {
			tagClose = i
			tagStart = -1
		}
	}
	if buf == nil {
		return text
	}
	buf.WriteString(text[tagClose + 1:])
	return buf.String()
}

// StringReplaceHTMLTags replaces HTML/XML tags from text with replacement.
func StringReplaceHTMLTags(text, replacement string) (plainText string) {
	var buf *bytes.Buffer
	tagClose := -1
	tagStart := -1
	for i, char := range text {
		if char == '<' {
			if buf == nil {
				buf = bytes.NewBufferString(text)
				buf.Reset()
			}
			buf.WriteString(text[tagClose + 1 : i])
			tagStart = i
		} else if char == '>' && tagStart != -1 {
			buf.WriteString(replacement)
			tagClose = i
			tagStart = -1
		}
	}
	if buf == nil {
		return text
	}
	buf.WriteString(text[tagClose + 1:])
	return buf.String()
}

//StringMD5Hex returns the hex encoded MD5 hash of data.
func StringMD5Hex(data string) string {
	hash := md5.New()
	hash.Write([]byte(data))
	return fmt.Sprintf("%s", hex.EncodeToString(hash.Sum(nil)))
}

//StringMD5Hex returns the hex encoded SHA1 hash of data.
func StringSHA1Hex(data string) string {
	hash := sha1.New()
	hash.Write([]byte(data))
	return fmt.Sprintf("%s", hex.EncodeToString(hash.Sum(nil)))
}

//StringMD5Hex returns the hex encoded SHA256 hash of data.
func StringSHA256Hex(data string) string {
	hash := sha256.New()
	hash.Write([]byte(data))
	return fmt.Sprintf("%s", hex.EncodeToString(hash.Sum(nil)))
}

//StringBase64Encode returns the base64 encoding of data.
//TODO test func
func StringBase64Encode(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

//StringBase64Decode returns the string represented by the base64 string data.
//TODO test func
func StringBase64Decode(data string) (string, error) {
	b, err := base64.StdEncoding.DecodeString(data)
	return string(b), err
}

//StringAddURLParam add url param for given url.
func StringAddURLParam(url, name, value string) string {
	var separator string
	//如果不存在?,则表示是第一个参数
	if strings.IndexRune(url, '?') == -1 {
		separator = "?"
		//如果存在?,则表示不是第一个参数
	} else {
		separator = "&"
	}
	return url + separator + name + "=" + value
}

//StringToInt convert string to int.
func StringToInt(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

//StringFromInt convert int to string.
func StringFromInt(i int) string {
	return strconv.Itoa(i)
}

//StringToInt convert string to int64.
func StringToInt64(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}

//StringFromInt64 convert int64 to string.
func StringFromInt64(i int64) string {
	return strconv.FormatInt(i, 10)
}

//StringToInt convert string to float64.
func StringToFloat64(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

//StringFromFloat64 convert float64 to string.
func StringFromFloat64(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

//StringToInt convert string to bool.
//It accepts 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False.
func StringToBool(s string) bool {
	b, _ := strconv.ParseBool(s)
	return b
}

//StringFilter filter out all strings where the function does not return true.
func StringFilter(f func(string) bool, data []string) []string {
	result := make([]string, 0, 0)
	for _, element := range data {
		if f(element) {
			result = append(result, element)
		}
	}
	return result
}

//StringSubStr returns the substr from start to length.
func StringSubStr(s string, start, length int) string {
	bt := []rune(s)
	if start < 0 {
		start = 0
	}
	if start > len(bt) {
		start = start % len(bt)
	}
	var end int
	if (start + length) > (len(bt) - 1) {
		end = len(bt)
	} else {
		end = start + length
	}
	return string(bt[start:end])
}

//StringToIntArray convert string to a int array
//eg: "1,2,3" => [1,2,3]
func StringToIntArray(str string) ([]int, error) {
	strarr := strings.Split(str, ",")
	intarr := make([]int, 0)
	for _, v := range strarr {
		tempint := StringToInt(string(v))
		intarr = append(intarr, tempint)
	}
	return intarr, nil
}

// AESEncrypt encrypts text and given key with AES.
func AESEncrypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	b := base64.StdEncoding.EncodeToString(text)
	ciphertext := make([]byte, aes.BlockSize + len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
	return ciphertext, nil
}

// AESDecrypt decrypts text and given key with AES.
func AESDecrypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(text) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	data, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		return nil, err
	}
	return data, nil
}










