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
	"morning-dairy/service"
)

const (
	RE_FROM = "^[0-4]{1}$"
)

func init() {
	Router.HandleFunc("/dairy/{date:[0-9]{4}-[0-9]{2}-[0-9]{2}}", OnBeforeRouteFunc(dairy, service.Before2Sign))
	Router.HandleFunc("/dairy", OnBeforeRouteFunc(dairy, service.Before2Sign))
}

func dairy(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	date := vars["date"]
	from := input.Get(req, "from", RE_FROM, service.FROM_GLOBAL)

	dairies := service.Dairy.Dairies(date, from)

	output.Puts(writer, "dairys", dairies)
	output.Puts(writer, "code", CODE_OK)
}
