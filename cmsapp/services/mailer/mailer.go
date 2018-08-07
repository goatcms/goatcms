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

// Mailer is email sender service
type Mailer struct {
	smtpConfig struct {
		SMTPAddr     string `config:"mailer.smtp.addr"`
		AuthUsername string `config:"mailer.smtp.auth.username"`
		AuthPassword string `config:"mailer.smtp.auth.password"`
		AuthIdentity string `config:"?mailer.smtp.auth.identity"`
	}
	deps struct {
		Logger     services.Logger   `dependency:"LoggerService"`
		Template   services.Template `dependency:"TemplateService"`
		HTMLLayout string            `config:"?mailer.layout.html"`
		TextLayout string            `config:"?mailer.layout.text"`
	}
	sender *smtpmail.MailSender
}

// Factory create a Mailer instance
func Factory(dp dependency.Provider) (interface{}, error) {
	m := &Mailer{}
	if err := dp.InjectTo(&m.smtpConfig); err != nil {
		return nil, err
	}
	if err := dp.InjectTo(&m.deps); err != nil {
		return nil, err
	}
	if m.deps.HTMLLayout == "" {
		m.deps.HTMLLayout = htmlLayout
	}
	if m.deps.TextLayout == "" {
		m.deps.HTMLLayout = textLayout
	}
	config := smtpmail.Config(m.smtpConfig)
	m.sender = smtpmail.NewMailSender(config)
	return services.Mailer(m), nil
}

// Send render text and html email by named template.
// Next send it to users by smtp server.
func (m *Mailer) Send(to, name string, data interface{}, attachments []goatmail.Attachment, scope app.Scope) (err error) {
	m.deps.Logger.TestLog("Mailer.Send: send to: %v; name: %v; data: %v", to, name, data)
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
	m.sender.Send(&goatmail.Mail{
		Date:    time.Now(),
		Subject: name,
		Body: map[string]io.Reader{
			"text/plain": textReader,
			"text/html":  htmlReader,
		},
		Attachments: attachments,
	}, lifecycle)
	return nil
}
