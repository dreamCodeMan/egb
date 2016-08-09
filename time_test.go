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

func Test_TimeDayToUnix(t *testing.T) {
	input := "2016-08-05"
	correct := "1470326400"
	result := TimeDayToUnix(input)
	if result != correct {
		t.Fail()
	}
}

func Test_TimeSecondToUnix(t *testing.T) {
	input := "2016-08-05 15:04:02"
	correct := "1470380642"
	result := TimeSecondToUnix(input)
	if result != correct {
		t.Fatal(result)
		t.Fail()
	}
}

