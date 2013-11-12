/**
 * User: jackong
 * Date: 11/11/13
 * Time: 7:29 PM
 */
package global

import (
	"net/http"
	"github.com/gorilla/mux"
	"morning-dairy/err"
	"morning-dairy/io/output"
)

type BeforeFunc func(http.ResponseWriter, *http.Request) error

var (
	beforeFuncs []BeforeFunc
	afterFuncs []http.HandlerFunc
)

type router struct {
	*mux.Router
}

func (this *router) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	op := output.NewJson(writer)

	defer func() {
		if e := recover(); e != nil {
			accessErr := e.(err.AccessError)
			Access.Error(accessErr.Code, req, accessErr.Msg)
			output.Return(op, accessErr.Code, accessErr.Msg)
		}
		if re := op.Render(); re != nil {
			Access.Alert(http.StatusBadRequest, req, re)
		}
	}()

	if !fireBefore(op, req, beforeFuncs) {
		return
	}
	this.Router.ServeHTTP(op, req)

	fireAfter(op, req)
}

func OnBefore(before ...BeforeFunc) {
    beforeFuncs = append(beforeFuncs, before...)
}

func fireBefore(writer http.ResponseWriter, req *http.Request, beforeFuncs []BeforeFunc) bool {
	for _, beforeFunc := range beforeFuncs {
		if accessErr := beforeFunc(writer, req); accessErr != nil {
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
