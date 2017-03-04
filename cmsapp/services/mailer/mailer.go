package mailer

import (
	"io"
	"time"

	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/scope/scopesync"
	"github.com/goatcms/goatcore/dependency"
	"github.com/goatcms/goatcore/goatmail"
	"github.com/goatcms/goatcore/goatmail/smtpmail"
)

const (
	htmlLayout = "mail/html"
	textLayout = "mail/text"
)

// Mailer
type Mailer struct {
	smtpConfig struct {
		SmtpAddr     string `config:"mailer.smtp.addr"`
		AuthUsername string `config:"mailer.smtp.auth.username"`
		AuthPassword string `config:"mailer.smtp.auth.password"`
		AuthIdentity string `config:"?mailer.smtp.auth.identity"`
	}
	deps struct {
		Template   services.Template `dependency:"TemplateService"`
		HtmlLayout string            `config:"?mailer.layout.html"`
		TextLayout string            `config:"?mailer.layout.text"`
	}
	sender *smtpmail.MailSender
}

// Mailer create a mailer instance
func MailerFactory(dp dependency.Provider) (interface{}, error) {
	m := &Mailer{}
	if err := dp.InjectTo(&m.smtpConfig); err != nil {
		return nil, err
	}
	if err := dp.InjectTo(&m.deps); err != nil {
		return nil, err
	}
	if m.deps.HtmlLayout == "" {
		m.deps.HtmlLayout = htmlLayout
	}
	if m.deps.TextLayout == "" {
		m.deps.HtmlLayout = textLayout
	}
	config := smtpmail.Config(m.smtpConfig)
	m.sender = smtpmail.NewMailSender(config)
	return services.Mailer(m), nil
}

// getSession return session map by session id
func (m *Mailer) Send(to, name string, data interface{}, attachments []goatmail.Attachment, scope app.Scope) {
	lifecycle := scopesync.Lifecycle(scope)
	htmlTemplate, err := m.deps.Template.View("mail/html", name, scope)
	if err != nil {
		scope.Trigger(app.ErrorEvent, err)
		return
	}
	htmlReader, htmlWriter := io.Pipe()
	go func() {
		if err = htmlTemplate.Execute(htmlWriter, data); err != nil {
			lifecycle.Error(err)
		}
		htmlWriter.Close()
	}()
	textTemplate, err := m.deps.Template.View("mail/text", name, scope)
	if err != nil {
		scope.Trigger(app.ErrorEvent, err)
		return
	}
	textReader, textWriter := io.Pipe()
	go func() {
		if err = textTemplate.Execute(textWriter, data); err != nil {
			lifecycle.Error(err)
		}
		textWriter.Close()
	}()
	mail := &goatmail.Mail{
		Date:    time.Now(),
		Subject: name,
		Body: map[string]io.Reader{
			"text/plain": textReader,
			"text/html":  htmlReader,
		},
		Attachments: attachments,
	}
	m.sender.Send(mail, lifecycle)
}
