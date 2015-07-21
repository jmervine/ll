# ll

short for "linelogger"

```
PACKAGE DOCUMENTATION

package ll
    import "."

    ll (short for linelogger) is a vary simple Scrolls'y
    (https://github.com/asenchi/scrolls) * style logger. * * Example: * *
    package main * import "time" * import "gopkg.in/jmervine/ll.v1" * * func
    main() { * begin := time.Now() * * // ... do stuff ... * *
    ll.Info(&begin, map[string]interface{} { * "at": "main", * "data":
    "foo", * } * } * * // Output: * // * // YYYY-MM-DD HH:MM:SS at=main
    data=foo durration=108.258us *

VARIABLES

var DebugLogger *log.Logger
    DebugLogger is exported as a pass through to log.Logger under the hood,
    so default functions can still be called.

    Example:

    DebugLogger.Println("ack")

    Warning: this will exit with no output when not in debug mode.
    DebugLogger.Fatal("ack")

    TODO: create noop logger for when not in debug mode.

var InfoLogger *log.Logger
    InfoLogger is exported as a pass through to log.Logger under the hood,
    so default functions can still be called.

    Example:

    InfoLogger.Fatal("ack")

FUNCTIONS

func Debug(begin *time.Time, meta map[string]interface{})
    Debug is the debug (DEBUG=true) logger, logging to os.Stdout by default
    when os.Getenv("DEBUG") is true.

func Info(begin *time.Time, meta map[string]interface{})
    Info is the standard logger, always logging to os.Stdout by default.

    Example usage:

	begin := time.Now()
	// do stuff
	Info(&begin, map[string]interface{} {
	    "at": "request",
	    "method": "GET",
	    "url: "/path/to/file.html",
	    "error": errors.New("something bad happened"),
	})

	// Outputs:
	YYYY-MM-DD HH:MM:SS at=request method=GET url=/path/to/file.html error="something bad happened" durration=#ns

	// without time
	Info(nil, map[string]interface{}{
	    "at": "request",
	    "request": fmt.Sprintf("%+v", *req),
	})

	// Outputs:
	YYYY-MM-DD HH:MM:SS at=request request={ ... request args ... }

func SetOutput(out io.Writer)
    SetOutput allows you to change the output destination of both InfoLogger
    and DebugLogger in one shot.

    Example:

	SetOutput(os.Stderr)
```
