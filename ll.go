/*
ll (short for linelogger) is a vary simple Scrolls'y (https://github.com/asenchi/scrolls)
style logger.

Example:

    package main
    import "time"
    import "gopkg.in/jmervine/ll.v1"

    func main() {
        begin := time.Now()

        // ... do stuff ...

        ll.Info(&begin, map[string]interface{} {
            "at": "main",
            "data": "foo",
        }
    }

    // Output:
    //
    // YYYY-MM-DD HH:MM:SS at=main data=foo durration=108.258us
*/
package ll

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

// InfoLogger is exported as a pass through to log.Logger under the
// hood, so default functions can still be called.
//
// Example:
//
// InfoLogger.Fatal("ack")
var InfoLogger *log.Logger

// DebugLogger is exported as a pass through to log.Logger under the
// hood, so default functions can still be called.
//
// Example:
//
// DebugLogger.Println("ack")
//
// Warning: this will exit with no output when not in debug mode.
// DebugLogger.Fatal("ack")
//
// TODO: create noop logger for when not in debug mode.
var DebugLogger *log.Logger

func init() {
	InfoLogger = log.New(os.Stdout, "", log.Flags())

	// if anything but '/true/i', '/t/i' or '1' hide debug output
	if !hasDebug() {
		DebugLogger = log.New(ioutil.Discard, "", log.Flags())
	} else {
		DebugLogger = InfoLogger
	}
}

func hasDebug() bool {
	if ok, _ := strconv.ParseBool(os.Getenv("DEBUG")); !ok {
		return false
	}
	return true
}

// Logger provides a good loggging mechanism.
func logger(target *log.Logger, level string, begin *time.Time, meta map[string]interface{}) {
	toS := func(k string, i interface{}) string {
		return fmt.Sprintf("%s=%v", k, i)
	}
	toQ := func(k string, i interface{}) string {
		return fmt.Sprintf("%s=\"%s\"", k, i)
	}

	var line = []string{toS("level", level)}
	for key, val := range meta {
		var pair string
		switch v := val.(type) {
		case error:
			pair = toQ(key, v.Error())
		case []string:
			pair = toQ(key, strings.Join(v, ","))
		}

		if pair == "" {
			pair = toS(key, val)
		}

		line = append(line, pair)
	}

	if begin != nil {
		line = append(line, toS("durration", time.Since(*begin)))
	}

	target.Println(strings.Join(line, " "))
}

// SetOutput allows you to change the output destination of both InfoLogger
// and DebugLogger in one shot.
//
// Example:
//
//     SetOutput(os.Stderr)
func SetOutput(out io.Writer) {
	InfoLogger = log.New(out, "", log.Flags())
	if hasDebug() {
		DebugLogger = log.New(out, "", log.Flags())
	}
}

// Info is the standard logger, always logging to os.Stdout by default.
//
// Example usage:
//
//     begin := time.Now()
//     // do stuff
//     Info(&begin, map[string]interface{} {
//         "at": "request",
//         "method": "GET",
//         "url: "/path/to/file.html",
//         "error": errors.New("something bad happened"),
//     })
//
//     // Outputs:
//     YYYY-MM-DD HH:MM:SS at=request method=GET url=/path/to/file.html error="something bad happened" durration=#ns
//
//     // without time
//     Info(nil, map[string]interface{}{
//         "at": "request",
//         "request": fmt.Sprintf("%+v", *req),
//     })
//
//     // Outputs:
//     YYYY-MM-DD HH:MM:SS at=request request={ ... request args ... }
//
func Info(begin *time.Time, meta map[string]interface{}) {
	logger(InfoLogger, "info", begin, meta)
}

// Debug is the debug (DEBUG=true) logger, logging to os.Stdout by default
// when os.Getenv("DEBUG") is true.
func Debug(begin *time.Time, meta map[string]interface{}) {
	logger(DebugLogger, "debug", begin, meta)
}
