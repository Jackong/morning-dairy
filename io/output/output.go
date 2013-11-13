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
	Render(code int, msg string) (error)
}

func Puts(writer http.ResponseWriter, ret ...interface {}) {
	writer.(Output).Puts(ret)
}
