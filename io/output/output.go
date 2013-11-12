/**
 * User: jackong
 * Date: 11/12/13
 * Time: 2:26 PM
 */
package output

import (
	"net/http"
)

type Output interface {
	Puts(ret []interface {})
	IsReturn() bool
	Render() (error)
}

func Puts(writer http.ResponseWriter, ret ...interface {}) {
	writer.(Output).Puts(ret)
}

func Return(writer http.ResponseWriter, code int, msg string) {
	Puts(writer, "code", code, "msg", msg)
}

func IsReturn(writer http.ResponseWriter) bool {
	return writer.(Output).IsReturn()
}
