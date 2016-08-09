package egb

// SliceAppendStr appends string to slice with no duplicates.
func SliceAppendStr(strs []string, str string) []string {
	for _, s := range strs {
		if s == str {
			return strs
		}
	}
	return append(strs, str)
}

// SliceCompareStr compares two 'string' type slices.
// It returns true if elements and order are both the same.
func SliceCompareStr(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

// SliceCompareStrIgnoreOrder compares two 'string' type slices.
// It returns true if elements are the same, and ignores the order.
func SliceCompareStrIgnoreOrder(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		for j := len(s2) - 1; j >= 0; j-- {
			if s1[i] == s2[j] {
				s2 = append(s2[:j], s2[j + 1:]...)
				break
			}
		}
	}
	if len(s2) > 0 {
		return false
	}
	return true
}

// SliceCompareStr compares two 'int' type slices.
// It returns true if elements and order are both the same.
func SliceCompareInt(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

// SliceCompareStrIgnoreOrder compares two 'int' type slices.
// It returns true if elements are the same, and ignores the order.
func SliceCompareIntIgnoreOrder(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		for j := len(s2) - 1; j >= 0; j-- {
			if s1[i] == s2[j] {
				s2 = append(s2[:j], s2[j + 1:]...)
				break
			}
		}
	}
	if len(s2) > 0 {
		return false
	}
	return true
}







