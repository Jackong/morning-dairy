/**
 * User: jackong
 * Date: 11/12/13
 * Time: 5:01 PM
 */
package output

import (
	"net/http"
	"encoding/json"
	"fmt"
)

type Json struct {
	http.ResponseWriter
	ret map[string]interface {}
}

func NewJson(writer http.ResponseWriter) *Json{
	return &Json{ResponseWriter: writer, ret: make(map[string]interface {})}
}

func (this *Json) Puts(ret []interface {}) {
	for index := 0; index < len(ret); index += 2 {
		this.ret[ret[index].(string)] = ret[index + 1]
	}
}

func (this *Json) Render(code int, msg string) (err error) {
	this.ret["code"] = code
	this.ret["msg"] = msg

	var data []byte
	data, err = json.Marshal(this.ret)
	if err != nil {
		return err
	}
	fmt.Fprint(this.ResponseWriter, string(data))
	return
}
