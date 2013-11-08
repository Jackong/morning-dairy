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

func LogAppend(code int, req *http.Request, v ... interface {}) (more []interface {}) {
	more = append(more,
		code, DELIMITER,
		req.RemoteAddr, DELIMITER,
		req.Method, DELIMITER,
		req.URL.Path, DELIMITER,
		req.UserAgent(), DELIMITER)
	for _, s := range v {
		more = append(more, s)
	}
	return
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
