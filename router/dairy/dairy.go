/**
 * User: jackong
 * Date: 11/14/13
 * Time: 7:09 PM
 */
package dairy

import (
	. "morning-dairy/global"
	"net/http"
	"github.com/gorilla/mux"
	"morning-dairy/io/output"
	"morning-dairy/io/input"
)

func init() {
	Router.HandleFunc("/dairy/{date:[0-9]{8}}", dairy)
}

func dairy(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	date := vars["date"]
	from := input.Get(req, "from", "^[0-4]{1}$", "0")

	output.Puts(writer, "date", date)
	output.Puts(writer, "from", from)
	output.Puts(writer, "code", CODE_OK)
}
