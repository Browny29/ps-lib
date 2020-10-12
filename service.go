package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

type ProcService struct {}

func NewProcService() *ProcService{
	return new(ProcService)
}

// GetProcDirectories returns the names of all the directories inside the /proc directory
func (s *ProcService) GetProcDirectories() ([]string, error) {
	procDir, err := os.Open("/proc")
	if err != nil {
		return nil, err
	}
	defer procDir.Close()
	return procDir.Readdirnames(0)
}

// GetProcByPID returns the contents of the file /proc/:pid/stat where pid is the id of the process
func (s *ProcService) GetProcByPID(pid string) ([]byte, error) {
	statPath := fmt.Sprintf("/proc/%s/stat", pid)
	return ioutil.ReadFile(statPath)
}