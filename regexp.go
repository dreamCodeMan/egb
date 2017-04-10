package egb

import "regexp"

//RegexpIsPhoneNumber return true if given string is a phone number.
func RegexpIsPhoneNumber(mobileNum string) bool {
	regular := "^(13[0-9]|14[0-9]|15[0-9]|17[0-9]|18[0-9])\\d{8}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}

//RegexpIsStrongPassword return true if the given string is a strong pwd(6位以上数字和字母的组合).
func RegexpIsStrongPassword(pwd string) bool {
	if len(pwd) < 6 {
		return false
	}
	regex1 := `[a-zA-Z]`
	reg1 := regexp.MustCompile(regex1)
	r1 := reg1.MatchString(pwd)
	regex2 := `[0-9]`
	reg2 := regexp.MustCompile(regex2)
	r2 := reg2.MatchString(pwd)
	return r1 && r2
}

//RegexpIsUrl return true if the given string is a true url.
func RegexpIsUrl(url string) bool {
	regular := `^(http://|https://)?((?:[A-Za-z0-9]+-[A-Za-z0-9]+|[A-Za-z0-9]+)\.)+([A-Za-z]+)[/\?\:]?.*$`
	reg := regexp.MustCompile(regular)
	return reg.MatchString(url)
}

//RegexpIsEmailAddress return true if the given string is a email address.
func RegexpIsEmailAddress(email string) bool {
	regular := `^[^@]+@[^@]+\.[^@]+$`
	emailFormatPregex := regexp.MustCompile(regular)
	return emailFormatPregex.MatchString(email)
}
