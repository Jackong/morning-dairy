/**
 * User: jackong
 * Date: 11/14/13
 * Time: 7:21 PM
 */
package service

type dairy struct {

}

var (
	Dairy dairy
)

const (
	RE_FROM = "^[0-4]{1}$"
	FROM_GLOBAL = "0"
	FROM_LEFT = "1"
	FROM_RIGHT = "2"
	FROM_UP = "3"
	FROM_DOWN = "4"
)

func init() {
	Dairy = dairy{}
}
