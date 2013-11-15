/**
 * User: jackong
 * Date: 10/17/13
 * Time: 6:24 PM
 */
package main

import (
	"net/http"
	"fmt"
	. "morning-dairy/global"
	_ "morning-dairy/router/sign"
	_ "morning-dairy/router/dairy"
)


func main() {
	err := http.ListenAndServe(Project.String("server", "addr"), Router)
	if	err != nil {
		fmt.Println(err)
	}
	ShutDown()
}
