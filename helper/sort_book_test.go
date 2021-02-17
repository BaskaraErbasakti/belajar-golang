package helper

import "testing"

func TestSortBook(t *testing.T) {
	input1 := SortBook("3A13 5X19 9Y20 2C18 1N20 3N20 1M21 1F14 9A21 3N21 0E13 5G14 8A23 9E22 3N14")
	if input1 != "0E13 9E22 9A21 9Y20 8A23 1M21 1N20 1F14 2C18 5X19 5G14 3N21 3N20 3N14 3A13" {
		t.Error("Fail output")
	}
}
