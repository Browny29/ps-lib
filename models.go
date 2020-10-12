package main

import "time"

type process struct {
	PID  string
	TTY  string
	Time time.Time
	CMD  string
}
