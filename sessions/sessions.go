package sessions

import (
	"net/http"

	"github.com/gorilla/sessions"
)

// cookie store to store session data in cookie
var Store = sessions.NewCookieStore([]byte("secret-password"))

// check if user is logged in
func IsLoggedIn(r *http.Request) bool {
	session, _ := Store.Get(r, "session")
	return session.Values["loggedin"] == "true"

}
