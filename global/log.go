/**
 * User: jackong
 * Date: 11/6/13
 * Time: 4:28 PM
 */
package global

import "net/http"

const (
	DELIMITER = "|"
)
func logMore(code int, req *http.Request) (more []interface {}) {
	return append(more,
		code, DELIMITER,
		req.RemoteAddr, DELIMITER,
		req.Method, DELIMITER,
		req.URL.Path, DELIMITER,
		req.UserAgent(), DELIMITER)
}

func Debug(code int, req *http.Request, v...interface {}) {
    Log.Debug(append(logMore(code, req), v)...)
}


func Info(code int, req *http.Request, v...interface {}) {
	Log.Info(append(logMore(code, req), v)...)
}


func Config(code int, req *http.Request, v...interface {}) {
	Log.Config(append(logMore(code, req), v)...)
}


func Warn(code int, req *http.Request, v...interface {}) {
	Log.Warn(append(logMore(code, req), v)...)
}


func Error(code int, req *http.Request, v...interface {}) {
	Log.Error(append(logMore(code, req), v)...)
}

func Alert(code int, req *http.Request, v...interface {}) {
	Log.Alert(append(logMore(code, req), v)...)
}

func Fatal(code int, req *http.Request, v...interface {}) {
	Log.Fatal(append(logMore(code, req), v)...)
}
