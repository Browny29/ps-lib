package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	fmt.Println("hello word")
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
	processes := make([]process, 0, len(dirNames))
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

		i++

		result, err := getProcessInformation(name)
		if err != nil {
			return processes, err
		}

		processes[i] = *result
	}

	return nil, nil
}

func getProcessInformation(pid string) (*process, error) {
	return nil, nil
}
