/**
 * User: jackong
 * Date: 11/12/13
 * Time: 8:17 PM
 */
package sign

import (
	"net/http"
	"morning-dairy/io/input"
	."morning-dairy/global"
	"morning-dairy/service"
	"morning-dairy/io/output"
)

var (
	user service.User
)

type beforeHnd struct {

}

func init() {
	user = service.User{}
	Router.Handle("/sign/up", &beforeHnd{})
}

func (this *beforeHnd) ServeHTTP(writer http.ResponseWriter, req * http.Request) {
	userName := input.Required(req, "userName")
	input.Required(req, "password")
	if user.IsExist(userName) {
		output.Puts(writer, "code", CODE_FAIL)
		return
	}
	output.Puts(writer, "code", CODE_OK)
}
