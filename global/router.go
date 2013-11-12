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

type BeforeFunc func(http.ResponseWriter, *http.Request) AccessError

var (
	beforeFuncs []BeforeFunc
	afterFuncs []http.HandlerFunc
)

type router struct {
	*mux.Router
}

func (this *router) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	defer func() {
		if e := recover(); e != nil {
			err := e.(AccessError)
			Access.Error(err.Code, req, err.Msg)
			http.Error(writer, err.Msg, err.Code)
			return
		}
	}()

	if !fireBefore(writer, req, beforeFuncs) {
		return
	}
	this.Router.ServeHTTP(writer, req)

	fireAfter(writer, req)
}

func OnBefore(before ...BeforeFunc) {
    beforeFuncs = append(beforeFuncs, before...)
}

func fireBefore(writer http.ResponseWriter, req *http.Request, beforeFuncs []BeforeFunc) bool {
	for _, beforeFunc := range beforeFuncs {
		if accessErr := beforeFunc(writer, req); accessErr.Code != CODE_OK {
			panic(accessErr)
			return false
		}
	}
	return true
}

func OnAfter(after ...http.HandlerFunc) {
    afterFuncs = append(afterFuncs, after...)
}

func fireAfter(writer http.ResponseWriter, req *http.Request) {
    for _, afterFunc := range afterFuncs {
		afterFunc(writer, req)
	}
}

func OnBeforeRoute(handlerFunc http.HandlerFunc, before ...BeforeFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, req *http.Request) {
		if fireBefore(writer, req, before) {
			handlerFunc(writer, req)
		}
	}
}
