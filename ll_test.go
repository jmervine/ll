package ll

import (
	"bytes"
	"errors"
	"os"
	"strings"
	"testing"
	"time"

	. "gopkg.in/jmervine/GoT.v1"
)

var INFO_MAP = map[string]interface{}{
	"at":     "at",
	"method": "meth",
	"status": 200,
	"source": "test",
	"url":    "url",
}

var ERROR_MAP = map[string]interface{}{
	"at":     "at",
	"method": "meth",
	"status": 500,
	"error":  errors.New("error message"),
	"url":    "url",
}

// fetch buffer, remove whitespace and date for assertion
func cleanLine(s string) string {
	defer func() {
		recover() // ignore errors
	}()
	return strings.Join(strings.Split(strings.TrimSpace(s), " ")[2:], " ")
}

func TestStandardLogger(T *testing.T) {
	recorder := bytes.NewBuffer(nil)
	SetOutput(recorder)

	n := time.Now()
	Info(&n, INFO_MAP)

	out := cleanLine(recorder.String())

	Go(T).AssertContains(out, "durration=")
	Go(T).AssertContains(out, "at=at")
	Go(T).AssertContains(out, "status=200")
}

func TestErrorLogger(T *testing.T) {
	recorder := bytes.NewBuffer(nil)
	SetOutput(recorder)

	Info(nil, ERROR_MAP)

	out := cleanLine(recorder.String())

	Go(T).AssertContains(out, "status=500")
	Go(T).AssertContains(out, "error=\"error message\"")
}

func TestNoDebugLogger(T *testing.T) {
	recorder := bytes.NewBuffer(nil)
	SetOutput(recorder)

	Debug(nil, INFO_MAP)
	Go(T).AssertEqual("", recorder.String())
}

func TestDebugLogger(T *testing.T) {
	os.Setenv("DEBUG", "true")
	recorder := bytes.NewBuffer(nil)
	SetOutput(recorder)

	Debug(nil, ERROR_MAP)

	out := cleanLine(recorder.String())

	Go(T).AssertContains(out, "status=500")
	Go(T).AssertContains(out, "error=\"error message\"")
}
