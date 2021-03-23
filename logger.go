package elogger

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"

	jsoniter "github.com/json-iterator/go"
)

// Structured Logging functions

// ServiceName name of service
var ServiceName = ""

// StructuredLog struct for structured loggin
type StructuredLog struct {
	Timestamp string `json:"@timestamp,omitempty"`
	Service   string `json:"service,omitempty"`
	Thread    string `json:"thread,omitempty"`
	IP        string `json:"ip,omitempty"`
	Env       string `json:"env,omitempty"`
	Server    string `json:"server,omitempty"`
	Path      string `json:"path,omitempty"`

	Level        string      `json:"level,omitempty"`
	Event        string      `json:"event,omitempty"`
	Message      string      `json:"message,omitempty"`
	Account      string      `json:"account,omitempty"`
	ID           string      `json:"id,omitempty"`
	Raw          string      `json:"raw,omitempty"`
	RawInterface interface{} `json:"rawInterface,omitempty"`
}

// Event in json format
func Event(thelog StructuredLog, level, event, msg string) {
	thelog.Timestamp = time.Now().Format(time.RFC3339)
	hostname, _ := os.Hostname()
	thelog.Server = hostname
	thelog.Level = level
	thelog.Event = event
	thelog.Message = msg
	thelog.Service = ServiceName
	_, fn, line, ok := runtime.Caller(1)
	if ok {
		thelog.Path = fmt.Sprintf("%s:%v", fn, line)
	}
	logJSON, err := jsoniter.MarshalToString(thelog)
	if err != nil {
		log.Println("Structured Logger: Logger JSON Marshal failed !", err.Error())
	}
	log.Println(logJSON)
}

// LogNew is to log with a new StructuredLog struct
func LogNew(level, event, msg string) {
	var thelog StructuredLog
	thelog.Timestamp = time.Now().Format(time.RFC3339)
	hostname, _ := os.Hostname()
	thelog.Server = hostname
	thelog.Level = level
	thelog.Event = event
	thelog.Message = msg
	thelog.Service = ServiceName
	logJSON, err := jsoniter.MarshalToString(thelog)
	if err != nil {
		log.Println("Structured logger: Logger JSON Marshal failed !", err.Error())
	}
	log.Println(logJSON)
}

func (thelog *StructuredLog) Fatal(event string, err error) {
	if err != nil {
		thelog.Timestamp = time.Now().Format(time.RFC3339)
		hostname, _ := os.Hostname()
		thelog.Server = hostname
		thelog.Level = "error"
		thelog.Event = event
		thelog.Message = err.Error()
		thelog.Service = ServiceName
		_, fn, line, ok := runtime.Caller(1)
		if ok {
			thelog.Path = fmt.Sprintf("%s:%v", fn, line)
		}
		logJSON, err := jsoniter.MarshalToString(thelog)
		if err != nil {
			log.Println("Structured Logger: elogger JSON Marshal failed !", err.Error())
		}
		log.Fatal(logJSON)
	}
}

func (thelog *StructuredLog) Error(event, msg string) {
	thelog.Timestamp = time.Now().Format(time.RFC3339)
	hostname, _ := os.Hostname()
	thelog.Server = hostname
	thelog.Level = "error"
	thelog.Event = event
	thelog.Message = msg
	thelog.Service = ServiceName
	_, fn, line, ok := runtime.Caller(1)
	if ok {
		thelog.Path = fmt.Sprintf("%s:%v", fn, line)
	}
	logJSON, err := jsoniter.MarshalToString(thelog)
	if err != nil {
		log.Println("Structured Logger: elogger JSON Marshal failed !", err.Error())
	}
	log.Print(logJSON)
}
