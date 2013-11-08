/**
 * User: jackong
 * Date: 11/6/13
 * Time: 4:28 PM
 */
package global

import (
	"net/http"
	"sync"
	"github.com/Jackong/log/writer"
	"github.com/Jackong/log"
)

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


type asyncMail struct {
	*writer.Email
}


func (this *asyncMail) Write(p []byte) (n int, err error) {
	go this.Email.Write(p)
	return
}


type dateLog struct {
	mu     sync.Mutex
	dir string
	date string
	log.Logger
}

func (this *dateLog) Output(level, depth int, s string) {
	this.ensureDate()
	this.Logger.Output(level, depth + 1, s)
}

func (this *dateLog) ensureDate() {
	this.mu.Lock()
	defer this.mu.Unlock()
	today := Today()
	if this.date != today {
		this.date = today
		this.Logger = fileLog(this.dir, today)
	}
}
