/**
 * User: jackong
 * Date: 10/17/13
 * Time: 6:24 PM
 */
package main

import (
	"net/http"
	"fmt"
	"os"
	. "morning-dairy/global"
	_ "morning-dairy/router/sign"
)


func main() {
	err := http.ListenAndServe(Project.String("server", "addr"), Router)
	if	err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}
