package session

import (
	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
)

// Mailer
type Mailer struct {
	config struct {
		SmtpAddr     string `json:"smtp.addr"`
		AuthUsername string `json:"smtp.auth.username"`
		AuthPassword string `json:"smtp.auth.password"`
		AuthHost     string `json:"smtp.auth.host"`
		AuthIdentity string `json:"smtp.auth.identity"`
	}
	deps struct {
		Template services.Template `dependency:"TemplateService"`
	}
}

// Mailer create a mailer instance
func MailerFactory(dp dependency.Provider) (interface{}, error) {
	m := &Mailer{}
	if err := dp.InjectTo(&m.config); err != nil {
		return nil, err
	}
	if err := dp.InjectTo(&m.deps); err != nil {
		return nil, err
	}
	return services.Mailer(m), nil
}

// getSession return session map by session id
func (m *Mailer) Send(to, viewName string, data interface{}, scope app.Scope) {
	htmlTemplate, err := m.deps.Template.View("mail/html", viewName, scope)
	if err != nil {
		scope.Trigger(app.ErrorEvent, err)
		return
	}
	html := htmlTemplate.Execute(wr, data)
	textTemplate, err := m.deps.Template.View("mail/text", viewName, scope)
	if err != nil {
		scope.Trigger(app.ErrorEvent, err)
		return
	}
}
