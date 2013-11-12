/**
 * User: jackong
 * Date: 11/12/13
 * Time: 2:09 PM
 */
package io

import (
	"net/http"
	"regexp"
	"morning-dairy/err"
)

func Get(req *http.Request, name, pattern, defo string) string {
	value := req.FormValue(name)
	if value == "" {
		req.ParseForm()
		value = req.Form.Get(name)
	}
	if value == "" {
		if defo != "" {
			return defo
		}
		panic(err.AccessError{Code: http.StatusBadRequest, Msg: "Invalid param: " + name})
	}
	if pattern == "" {
		return value
	}
	if match, _ := regexp.MatchString(pattern, value); !match {
		panic(err.AccessError{Code: http.StatusBadRequest, Msg: "Invalid param: " + name})
	}
	return value
}

func Default(req *http.Request, name, defo string) string {
	return Get(req, name, "", defo)
}

func Pattern(req *http.Request, name, pattern string) string {
	return Get(req, name, pattern, "")
}

func Required(req *http.Request, name string) string {
	return Pattern(req, name, "")
}
