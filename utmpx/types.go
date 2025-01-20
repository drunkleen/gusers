package utmpx

import "time"

const (
	HostSize = 16
	IDSize   = 4
	LineSize = 32
	NameSize = 32
)

const (
	Accounting   = 9
	BootTime     = 2
	DeadProcess  = 8
	Empty        = 0
	InitProcess  = 5
	LoginProcess = 6
	NewTime      = 3
	OldTime      = 4
	RunLvl       = 1
	UserProcess  = 7
)

var Types = map[int]string{
	Accounting:   "accounting",
	BootTime:     "boot time",
	Empty:        "empty",
	DeadProcess:  "dead process",
	InitProcess:  "init process",
	LoginProcess: "login process",
	NewTime:      "new time",
	OldTime:      "old time",
	RunLvl:       "run level",
	UserProcess:  "user process",
}

type Entry struct {
	ID        string    `json:"id"`
	Host      string    `json:"host"`
	Line      string    `json:"line"`
	PID       int       `json:"pid"`
	SessionID int       `json:"session_id"`
	Type      int       `json:"type"`
	TypeStr   string    `json:"typestr"`
	User      string    `json:"user"`
	Timestamp time.Time `json:"timestamp"`
}
