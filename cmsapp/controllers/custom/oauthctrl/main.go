package oauthctrl

const (
	githubCSRFCookie = "github-login-csrf"

	githubAuthorizeURL     = "https://github.com/login/oauth/authorize"
	githubTokenURL         = "https://github.com/login/oauth/access_token"
	githubRedirectURL      = "/user/signin/github/auth"
	githubStartURL         = "/user/signin/github/start"
	redirectAfterSigninURL = "/"
)
