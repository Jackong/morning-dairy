/**
 * User: jackong
 * Date: 11/12/13
 * Time: 9:23 PM
 */
package service
import (
	"morning-dairy/model"
)

type user struct {

}

var (
	User user
)

func init() {
    User = user{}
}

func (this user) Exist(name string) bool {
	return model.User.Exist(name)
}


func (this user) Create(name, password string) bool {
	return model.User.Create(name, password)
}


func (this user) SignIn(name, password string) bool {
    return model.User.Password(name) == password
}
