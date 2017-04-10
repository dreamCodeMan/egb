package egb

import "testing"

func Test_RegexpIsPhoneNumber(t *testing.T) {
	correct := "15008477777"
	wrong := "12345678909"
	result := RegexpIsPhoneNumber(correct)
	if !result {
		t.Fail()
	}
	result = RegexpIsPhoneNumber(wrong)
	if result {
		t.Fail()
	}
}

func Test_RegexpIsStrongPassword(t *testing.T) {
	strongpwd := "zaq1xsw2"
	weakpwd := "123567890"
	result := RegexpIsStrongPassword(strongpwd)
	if !result {
		t.Fail()
	}
	result = RegexpIsStrongPassword(weakpwd)
	if result {
		t.Fail()
	}
}

func Test_RegexpIsUrl(t *testing.T) {
	righturl := "http://www.baidu.com"
	wrongurl := "xxxxxx"
	result := RegexpIsUrl(righturl)
	if !result {
		t.Fail()
	}
	result = RegexpIsUrl(wrongurl)
	if result {
		t.Fail()
	}
}

func Test_RegexpIsEmailAddress(t *testing.T) {
	rightEmail := "80928@qq.com"
	wrongEmail := "xxx.xxx"
	result := RegexpIsEmailAddress(rightEmail)
	if !result {
		t.Fail()
	}
	result = RegexpIsEmailAddress(wrongEmail)
	if result {
		t.Fail()
	}
}
