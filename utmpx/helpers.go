package utmpx

// #include <utmpx.h>
import "C"
import (
	"time"
)

func (e *Entry) String() string {
	return "Type: '" + string(e.Type) + " " + e.TypeStr + "'" +
		"PID: " + string(e.PID) +
		"Line: " + e.Line +
		"ID: " + e.ID +
		"User: " + e.User +
		"Host: " + e.Host +
		" Session: " + string(e.SessionID) +
		" Timestamp: " + e.Timestamp.String()
}

func next() (entry *Entry) {
	centry := C.getutxent()
	if centry == nil {
		return nil
	}
	entry = &Entry{
		Type:      int(centry.ut_type),
		TypeStr:   Types[int(centry.ut_type)],
		PID:       int(centry.ut_pid),
		Line:      toStr(&centry.ut_line[0], LineSize),
		ID:        toStr(&centry.ut_id[0], IDSize),
		User:      toStr(&centry.ut_user[0], NameSize),
		Host:      toStr(&centry.ut_host[0], HostSize),
		SessionID: int(centry.ut_session),
		Timestamp: time.Unix(int64(centry.ut_tv.tv_sec), int64(centry.ut_tv.tv_usec)*1000),
	}
	return entry
}
