package users

import (
	"log"
	"net/http"

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
	sess    services.SessionManager
}

// NewUserController create instance of a articles controller
func NewUserController(dp services.Provider) (*UserController, error) {
	var err error
	ctrl := &UserController{}
	ctrl.tmpl, err = dp.Template()
	if err != nil {
		return nil, err
	}
	ctrl.crypt, err = dp.Crypt()
	if err != nil {
		return nil, err
	}
	ctrl.auth, err = dp.Auth()
	if err != nil {
		return nil, err
	}
	ctrl.sess, err = dp.SessionManager()
	if err != nil {
		return nil, err
	}
	ctrl.userDAO, err = dp.UserDAO()
	if err != nil {
		return nil, err
	}
	return ctrl, nil
}

// TemplateSignUp is handler to serve template where one can register new user
func (c *UserController) TemplateSignUp(scope services.RequestScope) {
	if err := c.tmpl.ExecuteTemplate(scope.Response(), "users/register", nil); err != nil {
		scope.Error(err)
		return
	}
}

// TryToSignUp is handler to save user from form obtained data
func (c *UserController) TryToSignUp(scope services.RequestScope) {
	if err := scope.Request().ParseForm(); err != nil {
		scope.Error(err)
		return
	}
	// obtain data from form with gorilla schema decoder and validate
	decoder := schema.NewDecoder()
	registerForm := &forms.RegisterForm{}
	if err := decoder.Decode(registerForm, scope.Request().PostForm); err != nil {
		scope.Error(err)
		return
	}
	isUser := c.userDAO.FindByEmail(registerForm.Email) // try find user
	if result, errors := registerForm.Validate(isUser); result != true {
		c.tmpl.ExecuteTemplate(scope.Response(), "users/register", map[string]interface{}{
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
	http.Redirect(scope.Response(), scope.Request(), "/", http.StatusSeeOther)
}

// TemplateLogin is handler to serve template where one can log in
func (c *UserController) TemplateLogin(scope services.RequestScope) {
	if sessID, err := c.sess.Init(scope.Response(), scope.Request()); err == nil {
		userid, _ := c.auth.GetUserID(sessID)
		log.Println("current user is: ", userid)
	}
	if err := c.tmpl.ExecuteTemplate(scope.Response(), "users/login", nil); err != nil {
		log.Fatal("error rendering a template: ", err)
		http.Error(scope.Response(), err.Error(), http.StatusInternalServerError)
		return
	}
}

// TryToLogin is handler to try sign in with given data in POST request
func (c *UserController) TryToLogin(scope services.RequestScope) {
	// obtain data from login form...
	if err := scope.Request().ParseForm(); err != nil {
		scope.Error(err)
		return
	}
	decoder := schema.NewDecoder()
	loginForm := &forms.LoginForm{}
	if err := decoder.Decode(loginForm, scope.Request().PostForm); err != nil {
		scope.Error(err)
		return
	}
	// validate form data and check credentials
	user := c.userDAO.FindByEmail(loginForm.Email)
	if result, errors := loginForm.Validate(user, c.crypt); result != true {
		c.tmpl.ExecuteTemplate(scope.Response(), "users/login", map[string]interface{}{
			"Errors": errors,
			"Email":  loginForm.Email,
		})
		return
	}
	// if validation went ok then set session
	// TODO repair below, so successful login will invoke proper session
	sessID, err := c.sess.Init(scope.Response(), scope.Request())
	if err != nil {
		scope.Error(err)
		return
	}
	err = c.auth.Auth(sessID, loginForm.Email)
	if err != nil {
		scope.Error(err)
		return
	}
	http.Redirect(scope.Response(), scope.Request(), "/", http.StatusSeeOther)
}

// TryToLogout is handler to try logour from current user
func (c *UserController) TryToLogout(scope services.RequestScope) {
	sessID, err := c.sess.Init(scope.Response(), scope.Request())
	if err != nil {
		scope.Error(err)
		return
	}
	c.auth.Clear(sessID)
	http.Redirect(scope.Response(), scope.Request(), "/", http.StatusSeeOther)
}
