package users

import (
	"log"
	"net/http"

	"github.com/goatcms/goat-core/dependency"
	"github.com/goatcms/goatcms/forms"
	"github.com/goatcms/goatcms/models"
	"github.com/goatcms/goatcms/models/user"
	"github.com/goatcms/goatcms/services"
	"github.com/gorilla/schema"
)

// UserController is article controller endpoint
type UserController struct {
	tmpl    services.Template
	userDAO models.UserDAO
	crypt   services.Crypt
	auth    services.Auth
}

// NewUserController create instance of a articles controller
func NewUserController(dp dependency.Provider) (*UserController, error) {
	ctrl := &UserController{}
	// load template service from dependency provider
	tmplIns, err := dp.Get(services.TemplateID)
	if err != nil {
		return nil, err
	}
	ctrl.tmpl = tmplIns.(services.Template)
	// load userDAO service from dependency provider
	daoIns, err := dp.Get(models.UserDAOID)
	if err != nil {
		return nil, err
	}
	ctrl.userDAO = daoIns.(models.UserDAO)
	// load crypting service from dependency provider
	cryptIns, err := dp.Get(services.CryptID)
	if err != nil {
		return nil, err
	}
	ctrl.crypt = cryptIns.(services.Crypt)
	// load auth service from dependency provider
	authIns, err := dp.Get(services.AuthID)
	if err != nil {
		return nil, err
	}
	ctrl.auth = authIns.(services.Auth)
	// return
	return ctrl, nil
}

// TemplateSignUp is handler to serve template where one can register new user
func (c *UserController) TemplateSignUp(w http.ResponseWriter, r *http.Request) {
	log.Println("responding to", r.Method, r.URL)
	if err := c.tmpl.ExecuteTemplate(w, "users/register", nil); err != nil {
		log.Fatal("error rendering a template: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// TryToSignUp is handler to save user from form obtained data
func (c *UserController) TryToSignUp(w http.ResponseWriter, r *http.Request) {
	log.Println("responding to", r.Method, r.URL)
	if err := r.ParseForm(); err != nil {
		log.Fatal("error parsing a form: ", err)
		return
	}
	// obtain data from form with gorilla schema decoder
	decoder := schema.NewDecoder()
	registerForm := &forms.RegisterForm{}
	if err := decoder.Decode(registerForm, r.PostForm); err != nil {
		log.Fatal(err)
	}
	// validate form data
	isUser := c.userDAO.FindByEmail(registerForm.Email) // try find user
	if result, errors := registerForm.Validate(isUser); result != true {
		c.tmpl.ExecuteTemplate(w, "users/register", map[string]interface{}{
			"Errors": errors,
			"Email":  registerForm.Email,
		})
		return
	}
	// encrypt password with bcrypt
	passHashed, err := c.crypt.Hash(registerForm.Password)
	if err != nil {
		log.Fatal("error crypting pass: ", err)
		return
	}
	user := usermodel.UserDTO{Email: registerForm.Email, PassHash: passHashed}
	// ...and save to database
	var userToAdd []models.UserDTO
	userToAdd = append(userToAdd, models.UserDTO(&user))
	c.userDAO.PersistAll(userToAdd)
	// redirect
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// TemplateLogin is handler to serve template where one can log in
func (c *UserController) TemplateLogin(w http.ResponseWriter, r *http.Request) {
	log.Println("responding to", r.Method, r.URL)
	if err := c.tmpl.ExecuteTemplate(w, "users/login", nil); err != nil {
		log.Fatal("error rendering a template: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// TryToLogin is handler to try sign in with given data in POST request
func (c *UserController) TryToLogin(w http.ResponseWriter, r *http.Request) {
	log.Println("responding to", r.Method, r.URL)
	err := r.ParseForm()
	if err != nil {
		log.Fatal("error parsing a form: ", err)
	}
	// obtain data from login form...
	decoder := schema.NewDecoder()
	loginForm := &forms.LoginForm{}
	if err2 := decoder.Decode(loginForm, r.PostForm); err != nil {
		log.Fatal(err2)
	}
	// validate form data and check credentials
	user := c.userDAO.FindByEmail(loginForm.Email) // try find user
	if result, errors := loginForm.Validate(user, c.crypt); result != true {
		c.tmpl.ExecuteTemplate(w, "users/login", map[string]interface{}{
			"Errors": errors,
			"Email":  loginForm.Email,
		})
		return
	}
	// if validation ok then set session
	c.auth.SetSession(loginForm.Email, w)
	// log.Println(w.Header()) // DEBUG
	// and redirect
	c.auth.ExecuteTemplateAuth(w, r, "/") // middleware, if no session redir /login
	// http.Redirect(w, r, "/", http.StatusSeeOther)
}

// TryToLogout is handler to try logour from current user
func (c *UserController) TryToLogout(w http.ResponseWriter, r *http.Request) {
	c.auth.ClearSession(w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
