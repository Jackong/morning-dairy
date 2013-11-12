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
	ret []interface {}
}

func (this *Json) Append(a []interface {}) {
	this.ret = a
}

func (this *Json) Render() (err error) {
	var data []byte
	if len(this.ret) == 1 {
		data, err = json.Marshal(this.ret[0])
	} else {
		data, err = json.Marshal(this.ret)
	}
	if err != nil {
		return err
	}
	fmt.Fprint(this.ResponseWriter, string(data))
	return
}
