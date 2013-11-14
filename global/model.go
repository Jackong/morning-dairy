/**
 * User: jackong
 * Date: 11/14/13
 * Time: 2:03 PM
 */
package global

import (
	"fmt"
	"os"
	"github.com/Jackong/db"
)

func Collection(name string) db.Collection {
	collection, err := Conn.Collection(name)
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
	return collection
}
