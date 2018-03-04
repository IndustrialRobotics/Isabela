package calc

import (
	"testing"
)

func AddNum(x, y int) int {
	return x + y
}

func SubstractNum(x, y int) int {
	return x - y
}

func testAddNum(t *testing.T) {
	var result int
	result = AddNum(15, 10)
	if result != 25 {
		t.Error("Expected 25, got", result)
	}
}

func testSubstractNum(t *testing.T) {
	var result int
	result = SubstractNum(15, 10)
	if result != 5 {
		t.Error("Expected 25, got", result)
	}
}
