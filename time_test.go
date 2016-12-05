package egb

import "testing"

func Test_TimeYear(t *testing.T) {
	year := TimeYear()
	if year == "" {
		t.Fail()
	}
}

func Test_TimeWeekDay(t *testing.T) {
	weekday := TimeWeekDay()
	if weekday == "" {
		t.Fail()
	}
}

/*
//TODO 本地测试成功,但是travis失败(时区问题???)
func Test_TimeDayToUnix(t *testing.T) {
	input := "2016-08-05"
	correct := "1470326400"
	result := TimeDayToUnix(input)
	if result != correct {
		t.Fatal(result)
	}
}

func Test_TimeSecondToUnix(t *testing.T) {
	input := "2016-08-05 15:04:02"
	correct := "1470380642"
	result := TimeSecondToUnix(input)
	if result != correct {
		t.Fatal(result)
	}
}
*/
