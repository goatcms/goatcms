package oauthctrl

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/goatcms/goatcms/cmsapp/dao"
	"github.com/goatcms/goatcms/cmsapp/entities"
	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcms/cmsapp/services/requestdep"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
	"github.com/goatcms/goatcore/varutil"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// GithubSignin is register controller
type GithubSignin struct {
	deps struct {
		Template                  services.Template             `dependency:"TemplateService"`
		Logger                    services.Logger               `dependency:"LoggerService"`
		UserConnectCriteriaSearch dao.UserConnectCriteriaSearch `dependency:"UserConnectCriteriaSearch"`

		BaseURL                string `config:"app.baseURL"`
		GithubAppID            string `config:"oauth.github.app"`
		GithubAppSecret        string `config:"oauth.github.secret"`
		GithubAuthorizeURL     string `?config:"oauth.github.authorizeURL"`
		GithubTokenURL         string `?config:"oauth.github.tokenURL"`
		GithubRedirectURL      string `?config:"oauth.github.redirectURL"`
		RedirectAfterSigninURL string `?config:"oauth.github.redirectAfterSigninURL"`
	}
	//view *template.Template
	oauthCfg *oauth2.Config
}

// NewGithubSignin create instance of a register form controller
func NewGithubSignin(dp dependency.Provider) (*GithubSignin, error) {
	var err error
	ctrl := &GithubSignin{}
	if err = dp.InjectTo(&ctrl.deps); err != nil {
		return nil, err
	}
	if ctrl.deps.GithubAuthorizeURL == "" {
		ctrl.deps.GithubAuthorizeURL = githubAuthorizeURL
	}
	if ctrl.deps.GithubTokenURL == "" {
		ctrl.deps.GithubTokenURL = githubTokenURL
	}
	if ctrl.deps.GithubRedirectURL == "" {
		ctrl.deps.GithubRedirectURL = githubRedirectURL
	}
	if ctrl.deps.RedirectAfterSigninURL == "" {
		ctrl.deps.RedirectAfterSigninURL = redirectAfterSigninURL
	}
	ctrl.oauthCfg = &oauth2.Config{
		ClientID:     ctrl.deps.GithubAppID,
		ClientSecret: ctrl.deps.GithubAppSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  ctrl.deps.GithubAuthorizeURL,
			TokenURL: ctrl.deps.GithubTokenURL,
		},
		RedirectURL: ctrl.deps.BaseURL + ctrl.deps.GithubRedirectURL,
		Scopes:      []string{"read:user"},
	}
	return ctrl, nil
}

// Get is a endpoint to start signin process
func (c *GithubSignin) Get(scope app.Scope) (err error) {
	var (
		deps struct {
			SessionManager requestdep.SessionManager `request:"SessionService"`
			Responser      requestdep.Responser      `request:"ResponserService"`
			Response       http.ResponseWriter       `request:"Response"`
			RequestError   requestdep.Error          `request:"ErrorService"`
		}
		secret string
	)
	if err = scope.InjectTo(&deps); err != nil {
		return err
	}
	secret = varutil.RandString(20, varutil.StrongBytes)
	http.SetCookie(deps.Response, &http.Cookie{
		Name:     githubCSRFCookie,
		Value:    secret,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Path:     "/",
	})
	url := c.oauthCfg.AuthCodeURL(secret)
	deps.Responser.Redirect(url)
	return nil
}

// Post is a endpoint fo user auth callback
func (c *GithubSignin) Post(scope app.Scope) (err error) {
	var (
		deps struct {
			SessionManager requestdep.SessionManager `request:"SessionService"`
			Request        *http.Request             `request:"Request"`
			Response       http.ResponseWriter       `request:"Response"`
			RequestError   requestdep.Error          `request:"ErrorService"`
			RequestAuth    requestdep.Auth           `request:"AuthService"`
			Responser      requestdep.Responser      `request:"ResponserService"`
		}
		token       *oauth2.Token
		githubUser  *github.User
		rows        dao.UserConnectRows
		userConnect *entities.UserConnect
		cookie      *http.Cookie
	)
	if err = scope.InjectTo(&deps); err != nil {
		return err
	}
	if cookie, err = deps.Request.Cookie(githubCSRFCookie); err != nil {
		return err
	}
	http.SetCookie(deps.Response, &http.Cookie{
		Name:    githubCSRFCookie,
		Value:   "",
		Expires: time.Unix(0, 0),
		Path:    "/",
	})
	if deps.Request.URL.Query().Get("state") != cookie.Value {
		return fmt.Errorf("no state match; possible csrf OR cookies not enabled")
	}
	if token, err = c.oauthCfg.Exchange(oauth2.NoContext, deps.Request.URL.Query().Get("code")); err != nil {
		return fmt.Errorf("there was an issue getting your token")
	}
	if !token.Valid() {
		return fmt.Errorf("retreived invalid token")
	}
	client := github.NewClient(c.oauthCfg.Client(oauth2.NoContext, token))
	deadline := time.Now().Add(5000 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()
	if githubUser, _, err = client.Users.Get(ctx, ""); err != nil {
		return err
	}
	githubUseLogin := strings.ToLower(githubUser.GetLogin())
	if rows, err = c.deps.UserConnectCriteriaSearch.Find(scope, &dao.UserConnectCriteria{
		Fields: &entities.UserConnectFields{},
		Where: dao.UserConnectCriteriaWhere{
			RemoteID: &dao.StringFieldCriteria{
				Value: []string{githubUseLogin},
				Type:  dao.EQ,
			},
			Service: &dao.StringFieldCriteria{
				Value: []string{"github"},
				Type:  dao.EQ,
			},
		},
		Related: dao.UserConnectCriteriaRelated{
			User: &dao.UserCriteria{
				Fields: entities.UserAllFieldsAndID,
			},
		},
		Order: dao.UserConnectCriteriaOrder{},
	}, &dao.Pager{
		Limit:  1,
		Offset: 0,
	}); err != nil {
		return err
	}
	defer rows.Close()
	if !rows.Next() {
		return fmt.Errorf("no results for %v", githubUseLogin)
	}
	if userConnect, err = rows.Get(); err != nil {
		return err
	}
	if _, err = deps.RequestAuth.ForceSignin(userConnect.User); err != nil {
		return err
	}
	if err = scope.Trigger(app.CommitEvent, nil); err != nil {
		return err
	}
	deps.Responser.Redirect(c.deps.BaseURL + c.deps.RedirectAfterSigninURL)
	return err
}
