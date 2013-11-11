/**
 * User: jackong
 * Date: 11/11/13
 * Time: 7:29 PM
 */
package global

import (
	"net/http"
	"github.com/gorilla/mux"
)

type BeforeFunc func(http.ResponseWriter, *http.Request) bool

var (
	beforeFuncs []BeforeFunc
)

type Handler struct {
	*mux.Router
}

func (this *Handler) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	for _, beforeFunc := range beforeFuncs {
		if !beforeFunc(writer, req) {
			http.Error(writer, "400 Bad Request", http.StatusBadRequest)
			return
		}
	}
	this.Router.ServeHTTP(writer, req)
}

func OnBefore(before BeforeFunc) {
    beforeFuncs = append(beforeFuncs, before)
}
