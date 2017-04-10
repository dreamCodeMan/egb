package egb

import "testing"

func Test_RandString(t *testing.T) {
	length := 6
	result := RandString(length)
	if len(result) != length {
		t.Fail()
	}
}

func Test_RandNum(t *testing.T) {
	length := 6
	result := RandNum(length)
	if len(result) != length {
		t.Fail()
	}
}

func Test_RandIntBetween(t *testing.T) {
	min := 5
	max := 10
	result := RandIntBetween(min, max)
	if result < min || result > max {
		t.Fail()
	}
}

func Test_RandInt64Between(t *testing.T) {
	min := int64(5)
	max := int64(10)
	result := RandInt64Between(min, max)
	if result < min || result > max {
		t.Fail()
	}
}

func Test_RandNumber(t *testing.T) {
	min := 0
	max := 10
	count := 3
	result := RandNumber(min, max, count)
	if len(result) != count {
		t.Fail()
	}
	for _, v := range result {
		if v < min || v >= max {
			t.Fail()
		}
	}
}
