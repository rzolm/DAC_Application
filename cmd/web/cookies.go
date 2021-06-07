package cookies

import (
	"fmt"
	"net/http"
)

type Cookie struct {
	Name  string
	Value string
}

func CreateCookie(w http.ResponseWriter, r http.Request) {
	fmt.Println("creating cookie")

	if login == true {
		//create cookie
	}
	http.SetCookie(w, &http.Cookie{
		Name:  "advisor_login",
		Value: "yes",
	})
}

func VerifyCookie(w http.ResponseWriter, req *http.Request) {
	if cookieExists == true {
		c, err := req.Cookie("login-cookie")
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		fmt.Fprintln(w, "cookie:", c)
	}
}
