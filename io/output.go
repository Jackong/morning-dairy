/**
 * User: jackong
 * Date: 11/12/13
 * Time: 2:26 PM
 */
package io

import (
	"net/http"
	"fmt"
	"encoding/json"
)

type Output struct {
	http.ResponseWriter
	ret []interface {}
}

func (this *Output) Append(a []interface {}) {
	this.ret = a
}

func (this *Output) Render(req *http.Request) {
	//accept := req.Header.Get("Accept")
	//if accept == "application/json" {
		js, _ := json.Marshal(this.ret)
		fmt.Fprint(this.ResponseWriter, string(js))
	//}
}

func Puts(writer http.ResponseWriter, a ... interface {}) {
	writer.(*Output).Append(a)
}
