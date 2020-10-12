package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	processes, _ := getProc()

	fmt.Println("PID TTY Time CMD")
	for _, p := range processes {
		fmt.Println(fmt.Sprintf("%s %s %s %s", p.PID, p.TTY, p.Time.Format("15:04:05"), p.CMD))
	}

}

type process struct {
	PID  string
	TTY  string
	Time time.Time
	CMD  string
}

func getProc() ([]process, error) {
	procDir, err := os.Open("/proc")
	if err != nil {
		return nil, err
	}
	defer procDir.Close()

	dirNames, err := procDir.Readdirnames(0)
	processes := make([]process, len(dirNames))
	if err != nil {
		return nil, err
	}

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

	statPath := fmt.Sprintf("/proc/%s/stat", pid)
	bytes, err := ioutil.ReadFile(statPath)
	if err != nil {
		return nil, err
	}

	// Get the name of the executable
	data := string(bytes)
	nameStart := strings.IndexRune(data, '(') + 1
	nameEnd := strings.IndexRune(data[nameStart:], ')')
	proc.CMD = data[nameStart : nameEnd+nameStart]

	proc.TTY = findProcValue(data, 8) // TTY is at index 8 in the stat file

	// These two lines could be fit into a one liner, but for readability I split them
	stringTime := findProcValue(data, 23)
	seconds, err := strconv.ParseInt(stringTime, 10, 64) // TTY is at index 23 in the stat file
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