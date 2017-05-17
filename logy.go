package logy

import (
	"log"
	"os"
)

// Logfile : struct for logfile
type Logfile struct {
	filename string
	umask    os.FileMode
}

// NewLog : init new logger
func NewLog(logname string) (*Logfile, error) {
	l := &Logfile{
		filename: logname,
		umask:    0644,
	}
	f, err := os.OpenFile(l.filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, l.umask)
	if err != nil {
		return l, err
	}
	defer f.Close()
	return l, nil
}

// Log : write message to file
func (e *Logfile) Log(msg string) (string, error) {
	l, err := os.OpenFile(e.filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, e.umask)
	if err != nil {
		return e.filename, err
	}
	defer l.Close()
	log.SetOutput(l)
	log.Println(msg)
	return msg, nil
}

// Logfile : returns the Path of the Logfile
func (e *Logfile) Logfile() string {
	return e.filename
}
