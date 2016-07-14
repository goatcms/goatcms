package auth

import (
	"fmt"
	"net/http"

	"github.com/goatcms/goat-core/dependency"
	"github.com/gorilla/securecookie"
)

// Auth is global auth provider
type Auth struct {
	cookie *securecookie.SecureCookie
}

// NewAuth create a authentification service instance
func NewAuth(dp dependency.Provider) (*Auth, error) {
	var hashKey = []byte("very-secret")   // securecookie.GenerateRandomKey(64)
	var blockKey = []byte("a-lot-secret") // securecookie.GenerateRandomKey(32)
	// http://www.gorillatoolkit.org/pkg/securecookie
	return &Auth{
		cookie: securecookie.New(hashKey, blockKey),
	}, nil
}

// GetCode create HMAC for given string // obsolete as we have securecookie
// func (a *Auth) GetCode(data string) string {
// 	salt := "hereshouldbesomekey"
// 	h := hmac.New(sha256.New, []byte(salt))
// 	io.WriteString(h, data)
// 	return fmt.Sprintf("%x", h.Sum(nil))
// }

// GetUsername retrieve username from cookie decoded by securecookie
func (a *Auth) GetUsername(r *http.Request) (username string) {
	if cookie, err := r.Cookie("session"); err == nil {
		cookieValue := make(map[string]string)
		if err = a.cookie.Decode("session", cookie.Value, &cookieValue); err == nil {
			username = cookieValue["name"]
		}
	}
	return username
}

// SetSession put username in cookie encoded by securecookie
func (a *Auth) SetSession(username string, w http.ResponseWriter) {
	value := map[string]string{
		"name": username,
	}
	if encoded, err := a.cookie.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(w, cookie)
	}
}

// ClearSession set cookie age to -1, so client delete session info cookie
func (a *Auth) ClearSession(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
}

// ExecuteTemplateAuth execute template with auth (redirect to login if no auth)
func (a *Auth) ExecuteTemplateAuth(w http.ResponseWriter, r *http.Request, name string) {
	username := a.GetUsername(r)
	if username != "" {
		fmt.Fprintf(w, name, username)
		// t.tmpl.ExecuteTemplate(wr, name, data) // from template service
	} else {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}
