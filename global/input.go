/**
 * User: jackong
 * Date: 11/12/13
 * Time: 2:09 PM
 */
package global

import (
	"net/http"
	"regexp"
)

type input string

func (this input) Get(req *http.Request, name, pattern, defo string) string {
	value := req.FormValue(name)
	if value == "" {
		req.ParseForm()
		value = req.Form.Get(name)
	}
	if value == "" {
		if defo != "" {
			return defo
		}
		panic(AccessError{Code: http.StatusBadRequest, Msg: "Invalid param: " + name})
	}
	if pattern == "" {
		return value
	}
	if match, _ := regexp.MatchString(pattern, value); !match {
		panic(AccessError{Code: http.StatusBadRequest, Msg: "Invalid param: " + name})
	}
	return value
}

func (this input) Default(req *http.Request, name, defo string) string {
	return this.Get(req, name, "", defo)
}

func (this input) Pattern(req *http.Request, name, pattern string) string {
	return this.Get(req, name, pattern, "")
}

func (this input) Required(req *http.Request, name string) string {
	return this.Pattern(req, name, "")
}
