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
	if !FireBefore(writer, req) {
		return
	}
	this.Router.ServeHTTP(writer, req)

	FireAfter(writer, req)
}

func OnBefore(before ...BeforeFunc) {
    beforeFuncs = append(beforeFuncs, before...)
}

func FireBefore(writer http.ResponseWriter, req *http.Request) bool {
	for _, beforeFunc := range beforeFuncs {
		if accessErr := beforeFunc(writer, req); accessErr.Code != CODE_OK {
			Access.Error(accessErr.Code, req, accessErr.Msg)
			http.Error(writer, accessErr.Msg, accessErr.Code)
			return false
		}
	}
	return true
}

func OnAfter(after ...http.HandlerFunc) {
    afterFuncs = append(afterFuncs, after...)
}

func FireAfter(writer http.ResponseWriter, req *http.Request) {
    for _, afterFunc := range afterFuncs {
		afterFunc(writer, req)
	}
}
