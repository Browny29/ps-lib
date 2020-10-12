package main

import "testing"

func Test_FindProcValue(t *testing.T) {
	testString := "1 2 3 4 5 6 7 8 9"

	result := findProcValue(testString, 0)
	if result != "1" {
		t.Fail()
	}

	result = findProcValue(testString, 4)
	if result != "5" {
		t.Fail()
	}

	result = findProcValue(testString, 100)
	if result != "" {
		t.Fail()
	}
}


