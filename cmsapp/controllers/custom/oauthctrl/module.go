package oauthctrl

import (
	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcore/app"
)

// InitDependencies init all dependency modules
func InitDependencies(a app.App) (err error) {
	var (
		deps struct {
			Router services.Router `dependency:"RouterService"`
		}
		githubSignin *GithubSignin
	)
	if err = a.DependencyProvider().InjectTo(&deps); err != nil {
		return err
	}
	// signup
	if githubSignin, err = NewGithubSignin(a.DependencyProvider()); err != nil {
		return err
	}
	deps.Router.OnGet(githubStartURL, githubSignin.Get)
	deps.Router.OnGet(githubRedirectURL, githubSignin.Post)
	deps.Router.OnPost(githubRedirectURL, githubSignin.Post)
	return nil
}
