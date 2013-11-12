/**
 * User: jackong
 * Date: 11/12/13
 * Time: 2:26 PM
 */
package io

import (
	"net/http"
	"fmt"
)

type Output interface {
	Append(a []interface {})
	Render() (error)
}

type Normal struct {
	http.ResponseWriter
	ret []interface {}
}

func (this *Normal) Append(a []interface {}) {
	this.ret = a
}

func (this *Normal) Render() (err error) {
	fmt.Fprint(this.ResponseWriter, this.ret)
	return err
}


func Puts(writer http.ResponseWriter, a ... interface {}) {
	writer.(Output).Append(a)
}
