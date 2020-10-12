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

func Test_getProcessInformation(t *testing.T) {
	testString := "13936 (bash) S 1 13936 13936 196608 -1 0 3436 3436 0 0 78 359 78 359 8 0 0 0 854085461 7196672 2628 345"
	resultProc := getProcessInformation()
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
