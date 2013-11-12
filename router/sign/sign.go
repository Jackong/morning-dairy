/**
 * User: jackong
 * Date: 11/12/13
 * Time: 8:17 PM
 */
package sign

import (
	"net/http"
	"morning-dairy/io/input"
	. "morning-dairy/global"
)

func init() {
	Router.HandleFunc("/sign/up", signUp)
}

func signUp(writer http.ResponseWriter, req * http.Request) {
	input.Required(req, "userName")
	input.Required(req, "password")
}
