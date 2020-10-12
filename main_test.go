package main

import (
	"testing"
)

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

func Test_GetProcessInformation(t *testing.T) {
	procGetter = NewMockedService()
	resultProc, err := getProcessInformation("13936")
	if err != nil {
		t.Fatal(err)
	}

	if resultProc.PID != "13936" {t.Fail()}
	if resultProc.TTY != "-1" {t.Fail()}
	if resultProc.CMD != "bash" {t.Fail()}
}
