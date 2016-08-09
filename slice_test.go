package egb

import "testing"

func Test_SliceAppendStr(t *testing.T) {
	strs := []string{"a"}
	str := "b"
	result := SliceAppendStr(strs, str)
	if len(result) != 2 {
		t.Errorf("SliceAppendStr:\n Expect length => %d\n Got =>%d", 2, len(result))
	}
	result = SliceAppendStr(strs, str)
	if len(result) != 2 {
		t.Errorf("SliceAppendStr:\n Expect length => %d\n Got =>%d", 2, len(result))
	}
}

func Test_SliceCompareStr(t *testing.T) {
	strs1 := []string{"angelina", "golang", "language"}
	strs2 := []string{"angelina", "golang", "language"}
	strs3 := []string{"angelina", "go", "language"}
	if !SliceCompareStr(strs1, strs2) {
		t.Errorf("SliceCompareStr:\n Expect => %v\n Got => %v", true, SliceCompareStr(strs1, strs2))
	}
	if SliceCompareStr(strs1, strs3) {
		t.Errorf("SliceCompareStr:\n Expect => %v\n Got => %v", false, SliceCompareStr(strs1, strs3))
	}
}

func Test_SliceCompareStrIgnoreOrder(t *testing.T) {
	strs1 := []string{"a", "b", "c"}
	strs2 := []string{"b", "c", "a"}
	strs3 := []string{"a", "c", "d"}
	if !SliceCompareStrIgnoreOrder(strs1, strs2) {
		t.Errorf("SliceCompareStr:\n Expect => %v\n Got => %v", true, SliceCompareStrIgnoreOrder(strs1, strs2))
	}
	if SliceCompareStrIgnoreOrder(strs1, strs3) {
		t.Errorf("SliceCompareStr:\n Expect => %v\n Got => %v", false, SliceCompareStr(strs1, strs3))
	}
}

