package oauthctrl

import (
	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcore/app"
)

// InitDependencies init all dependency modules
func InitDependencies(a app.App) error {
	var deps struct {
		Router services.Router `dependency:"RouterService"`
	}
	if err := a.DependencyProvider().InjectTo(&deps); err != nil {
		return err
	}
	// signup
	github, err := NewGithubSignin(a.DependencyProvider())
	if err != nil {
		return err
	}
	deps.Router.OnGet("/user/signin/github/start", github.Get)
	deps.Router.OnGet(githubRedirectURL, github.Post)
	deps.Router.OnPost(githubRedirectURL, github.Post)
	return nil
}
