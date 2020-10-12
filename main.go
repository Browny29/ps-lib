package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

var procGetter iProcGetter

type iProcGetter interface {
	GetProcByPID(pid string) ([]byte, error)
	GetProcDirectories() ([]string, error)
}

func main() {
	procGetter = NewProcService()

	processes, err := getProc()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("PID TTY Time CMD")
	for _, p := range processes {
		fmt.Println(fmt.Sprintf("%s %s %s %s", p.PID, p.TTY, p.Time.Format("15:04:05"), p.CMD))
	}

}

func getProc() ([]process, error) {
	dirNames, err := procGetter.GetProcDirectories()
	if err != nil {
		return nil, err
	}
	processes := make([]process, len(dirNames))

	i := 0
	for _, name := range dirNames {
		// If the name is not a number it is not a process, so we continue
		_, err := strconv.Atoi(name)
		if err != nil {
			continue
		}

		result, err := getProcessInformation(name)
		if err != nil {
			return processes, err
		}

		processes[i] = *result
		i++ // We want to increment the i only when a process has been added. So we do it at the end
	}

	return processes[:i], nil
}

func getProcessInformation(pid string) (*process, error) {
	proc := &process{
		PID: pid,
	}

	bytes, err := procGetter.GetProcByPID(pid)
	if err != nil {
		return nil, err
	}

	// Get the name of the executable
	data := string(bytes)
	nameStart := strings.IndexRune(data, '(') + 1
	nameEnd := strings.IndexRune(data[nameStart:], ')')
	proc.CMD = data[nameStart : nameEnd+nameStart]

	proc.TTY = findProcValue(data, 7) // TTY is at index 7 in the stat file

	// These two lines could be fit into a one liner, but for readability I split them
	stringTime := findProcValue(data, 24) // TTY is at index 24 in the stat file
	seconds, err := strconv.ParseInt(stringTime[:11], 10, 64)
	if err != nil {
		return nil, err
	}

	proc.Time = time.Unix(seconds, 0)
	return proc, nil
}

func findProcValue(procData string, index int) string {
	var procValue string
	for i, char := range procData {
		if index <= 0 {
			procValue = procData[i:]
			valueEnd := strings.IndexRune(procValue, ' ')
			procValue = procValue[:valueEnd]

			return procValue
		}
		if char == ' ' {
			index--
		}
	}

	return ""
}
