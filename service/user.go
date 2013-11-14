/**
 * User: jackong
 * Date: 11/12/13
 * Time: 9:23 PM
 */
package service
import (
	"morning-dairy/model"
	"github.com/gorilla/securecookie"
	"net/http"
	. "morning-dairy/global"
)

type user struct {
	secureCookie *securecookie.SecureCookie
}

var (
	User user
)

func init() {
    User = user{secureCookie: securecookie.New([]byte(Project.String("cookie", "hash")), []byte(Project.String("cookie", "block")))}
}

func (this user) Exist(name string) bool {
	return model.User.Exist(name)
}


func (this user) Create(name, password string) bool {
	return model.User.Create(name, password)
}


func (this user) CanSignIn(name, password string) bool {
    return model.User.Password(name) == password
}

func (this user) SignIn(writer http.ResponseWriter, name, password string) {
	session := map[string] string {
		"name": name,
		"password": password,
	}
	if encoded, err := this.secureCookie.Encode("gsession", session); err == nil {
		cookie := &http.Cookie{
			Name:  "gsession",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(writer, cookie)
	}
}

func (this user) IsSignIn(req *http.Request) (ok bool, name string) {
	if cookie, err := req.Cookie("gsession"); err == nil {
		value := make(map[string]string)
		if err = this.secureCookie.Decode("gsession", cookie.Value, &value); err == nil {
			name, ok := value["name"]
			_, ko := value["password"]
			if ok && ko {
				return true, name
			}
		}
	}
	return false, ""
}
