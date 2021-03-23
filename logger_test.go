package elogger

import (
	"bytes"
	"io"
	"log"
	"os"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
)

// captureOutput capture os Stdout logs to test
func captureOutput(f func()) string {
	reader, writer, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	stdout := os.Stdout
	stderr := os.Stderr
	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
		log.SetOutput(os.Stderr)
	}()
	os.Stdout = writer
	os.Stderr = writer
	log.SetOutput(writer)
	out := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		var buf bytes.Buffer
		wg.Done()
		io.Copy(&buf, reader)
		out <- buf.String()
	}()
	wg.Wait()
	f()
	writer.Close()
	return <-out
}
func TestLogNew(t *testing.T) {
	t.Run("LogNew should log in json format", func(t *testing.T) {
		re := captureOutput(func() {
			LogNew("info", "testNewLog", "Test new log")
		})
		assert.Equal(t, gjson.Get(re, "level").String(), "info")
		assert.Equal(t, gjson.Get(re, "event").String(), "testNewLog")
	})
}

func TestEvent(t *testing.T) {
	t.Run("LogEvent should log in json format", func(t *testing.T) {
		var logData StructuredLog
		logData.Account = "account_id"

		re := captureOutput(func() {
			Event(logData, "info", "testEventLog", "test event logger")
		})
		assert.Equal(t, gjson.Get(re, "account").String(), "account_id")
		assert.Equal(t, gjson.Get(re, "level").String(), "info")
		assert.Equal(t, gjson.Get(re, "event").String(), "testEventLog")
	})
}

func TestError(t *testing.T) {
	t.Run("LogEvent should log in json format", func(t *testing.T) {
		var slog StructuredLog
		slog.Account = "account_id"

		re := captureOutput(func() {
			slog.Error("testEventLog", "test event logger")
		})
		assert.Equal(t, gjson.Get(re, "account").String(), "account_id")
		assert.Equal(t, gjson.Get(re, "level").String(), "error")
		assert.Equal(t, gjson.Get(re, "event").String(), "testEventLog")
	})
}
