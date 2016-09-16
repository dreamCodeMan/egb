package egb

import (
	"strings"
	"testing"
)

func Test_StringMapFunc(t *testing.T) {
	result := StringMapFunc(strings.TrimSpace, []string{"  a  ", " b ", "c", "  d", "e  "})
	correct := []string{"a", "b", "c", "d", "e"}
	if len(result) != len(correct) {
		t.Fail()
	}
	for i, _ := range result {
		if result[i] != correct[i] {
			t.Fail()
		}
	}
}

func Test_StringMarshalJSON(t *testing.T) {
	data := []string{"angelina", "test", "json"}
	result := StringMarshalJSON(data)
	correct := `["angelina","test","json"]`
	if result != correct {
		t.Fail()
	}
}

func Test_StringListContains(t *testing.T) {
	list := []string{"angelina", "test", "Json"}
	result := StringListContains(list, "angelina")
	if !result {
		t.Fail()
	}
	result = StringListContains(list, "json")
	if result {
		t.Fail()
	}
}

func Test_StringListContainsCaseInsensitive(t *testing.T) {
	list := []string{"angelina", "test", "Json"}
	result := StringListContainsCaseInsensitive(list, "Json")
	if !result {
		t.Fail()
	}
	result = StringListContainsCaseInsensitive(list, "json")
	if !result {
		t.Fail()
	}
}

func Test_StringStripHTMLTags(t *testing.T) {
	withHTML := "<div>Hello > World <br/> <im src='xxx'/>"
	skippedHTML := "Hello > World  "
	if StringStripHTMLTags(withHTML) != skippedHTML {
		t.Fail()
	}
}

func Test_StringReplaceHTMLTags(t *testing.T) {
	withHTML := "<div>Hello > World <br/> <im src='xxx'/>"
	replacedHTML := "xxHello > World xx xx"
	if StringReplaceHTMLTags(withHTML, "xx") != replacedHTML {
		t.Fail()
	}
}

func Test_StringMD5Hex(t *testing.T) {
	data := "123456"
	correct := "e10adc3949ba59abbe56e057f20f883e"
	result := StringMD5Hex(data)
	if result != correct {
		t.Fail()
	}
}

func Test_StringSHA1Hex(t *testing.T) {
	data := "123456"
	correct := "7c4a8d09ca3762af61e59520943dc26494f8941b"
	result := StringSHA1Hex(data)
	if result != correct {
		t.Fail()
	}
}

func Test_StringSHA256Hex(t *testing.T) {
	data := "123456"
	correct := "8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92"
	result := StringSHA256Hex(data)
	if result != correct {
		t.Fail()
	}
}

func Test_StringAddURLParam(t *testing.T) {
	url := "http://www.baidu.com"
	name := "name"
	value := "angelina"
	correct := "http://www.baidu.com?name=angelina"
	result := StringAddURLParam(url, name, value)
	if result != correct {
		t.Fail()
	}
	correct = "http://www.baidu.com?name=angelina&name=angelina"
	result = StringAddURLParam(result, name, value)
	if result != correct {
		t.Fail()
	}
}

func Test_StringToInt(t *testing.T) {
	str := "1"
	correct := int(1)
	result := StringToInt(str)
	if result != correct {
		t.Fail()
	}
}

func Test_StringToInt64(t *testing.T) {
	str := "1"
	correct := int64(1)
	result := StringToInt64(str)
	if result != correct {
		t.Fail()
	}
}

func Test_StringToFloat64(t *testing.T) {
	str := "1.11"
	correct := float64(1.11)
	result := StringToFloat64(str)
	if result != correct {
		t.Fail()
	}
}

func Test_StringToBool(t *testing.T) {
	str := "true"
	correct := true
	result := StringToBool(str)
	if result != correct {
		t.Fail()
	}
}

func Test_StringFilter(t *testing.T) {
	hFunc := func(s string) bool {
		return strings.HasPrefix(s, "h")
	}
	result := StringFilter(hFunc, []string{"cheese", "mouse", "hi", "there", "horse"})
	correct := []string{"hi", "horse"}
	if len(result) != len(correct) {
		t.Fail()
	}
	for i, _ := range result {
		if result[i] != correct[i] {
			t.Fail()
		}
	}
}

func Test_StringSubStr(t *testing.T) {
	str := "angelina123"
	correct := "123"
	result := StringSubStr(str, 8, 3)
	if result != correct {
		t.Fail()
	}
}

func Test_StringToIntArray(t *testing.T) {
	str := "1,2,3"
	correct := []int{1, 2, 3}
	result, err := StringToIntArray(str)
	if err != nil {
		t.Errorf("StringToIntArray error: %s", err.Error())
	}
	if !SliceCompareInt(result, correct) {
		t.Errorf("StringToIntArray:\n Expect => %#v\n Got => %#v\n", correct, result)
	}
}
