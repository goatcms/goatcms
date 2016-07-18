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
	sess    services.Session
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
	// load session service from dependency provider
	sessIns, err := dp.Get(services.SessionManagerID)
	if err != nil {
		return nil, err
	}
	ctrl.sess = sessIns.(services.Session)
	return ctrl, nil
}

// TemplateSignUp is handler to serve template where one can register new user
func (c *UserController) TemplateSignUp(w http.ResponseWriter, r *http.Request) {
	if err := c.tmpl.ExecuteTemplate(w, "users/register", nil); err != nil {
		log.Fatal("error rendering a template: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// TryToSignUp is handler to save user from form obtained data
func (c *UserController) TryToSignUp(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatal("error parsing a form: ", err)
		return
	}
	// obtain data from form with gorilla schema decoder and validate
	decoder := schema.NewDecoder()
	registerForm := &forms.RegisterForm{}
	if err := decoder.Decode(registerForm, r.PostForm); err != nil {
		log.Fatal(err)
	}
	isUser := c.userDAO.FindByEmail(registerForm.Email) // try find user
	if result, errors := registerForm.Validate(isUser); result != true {
		c.tmpl.ExecuteTemplate(w, "users/register", map[string]interface{}{
			"Errors": errors,
			"Email":  registerForm.Email,
		})
		return
	}
	// encrypt password with bcrypt and save user
	passHashed, err := c.crypt.Hash(registerForm.Password)
	if err != nil {
		log.Fatal("error crypting pass: ", err)
		return
	}
	user := usermodel.UserDTO{Email: registerForm.Email, PassHash: passHashed}
	var userToAdd []models.UserDTO
	userToAdd = append(userToAdd, models.UserDTO(&user))
	c.userDAO.PersistAll(userToAdd)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// TemplateLogin is handler to serve template where one can log in
func (c *UserController) TemplateLogin(w http.ResponseWriter, r *http.Request) {
	if sessID, err := c.sess.Init(w, r); err == nil {
		userid, _ := c.auth.GetUserID(sessID)
		log.Println("current user is: ", userid)
	}
	if err := c.tmpl.ExecuteTemplate(w, "users/login", nil); err != nil {
		log.Fatal("error rendering a template: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// TryToLogin is handler to try sign in with given data in POST request
func (c *UserController) TryToLogin(w http.ResponseWriter, r *http.Request) {
	// obtain data from login form...
	if err := r.ParseForm(); err != nil {
		log.Fatal("error parsing a form: ", err)
	}
	decoder := schema.NewDecoder()
	loginForm := &forms.LoginForm{}
	if err := decoder.Decode(loginForm, r.PostForm); err != nil {
		log.Fatal(err)
		return
	}
	// validate form data and check credentials
	user := c.userDAO.FindByEmail(loginForm.Email)
	if result, errors := loginForm.Validate(user, c.crypt); result != true {
		c.tmpl.ExecuteTemplate(w, "users/login", map[string]interface{}{
			"Errors": errors,
			"Email":  loginForm.Email,
		})
		return
	}
	// if validation went ok then set session
	// TODO repair below, so successful login will invoke proper session
	sessID, err := c.sess.Init(w, r)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = c.auth.Auth(sessID, loginForm.Email)
	if err != nil {
		log.Fatal(err)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// TryToLogout is handler to try logour from current user
func (c *UserController) TryToLogout(w http.ResponseWriter, r *http.Request) {
	sessID, err := c.sess.Init(w, r)
	if err != nil {
		log.Fatal(err)
		return
	}
	c.auth.Clear(sessID)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
