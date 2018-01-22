package reqresponser

import (
	"fmt"
	"html/template"
	"net/http"
	"sync"

	"github.com/goatcms/goatcms/cmsapp/entities"
	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcms/cmsapp/services/requestdep"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
)

// TranslateService provide translate system
type ResponserService struct {
	deps struct {
		Template        services.Template         `dependency:"TemplateService"`
		Logger          services.Logger           `dependency:"LoggerService"`
		DependencyScope app.Scope                 `dependency:"DependencyScope"`
		RequestScope    app.Scope                 `request:"RequestScope"`
		Translate       requestdep.Translate      `request:"TranslateService"`
		SessionManager  requestdep.SessionManager `request:"SessionService"`
		ACL             requestdep.ACL            `request:"ACLService"`
		Request         *http.Request             `request:"Request"`
		Response        http.ResponseWriter       `request:"Response"`
	}
	muSended sync.RWMutex
	sended   bool
}

func (rs *ResponserService) IsSended() bool {
	rs.muSended.RLock()
	defer rs.muSended.RUnlock()
	return rs.sended
}

func (rs *ResponserService) Execute(view *template.Template, data interface{}) error {
	var (
		loggedInUser *entities.User
		session      *entities.Session
		err          error
	)
	rs.muSended.Lock()
	if rs.sended {
		rs.muSended.Unlock()
		return fmt.Errorf("Response sended")
	}
	rs.sended = true
	rs.muSended.Unlock()
	if session, err = rs.deps.SessionManager.Get(); err == nil {
		// only for initied sessions
		loggedInUser = session.User
	}
	if err = view.Execute(rs.deps.Response, map[string]interface{}{
		"Data":         data,
		"Lang":         rs.deps.Translate.Lang(),
		"LoggedInUser": loggedInUser,
		"ACL":          rs.deps.ACL,
	}); err != nil {
		return err
	}
	return nil
}

func (rs *ResponserService) JSON(code int, json string) (err error) {
	rs.muSended.Lock()
	if rs.sended {
		rs.muSended.Unlock()
		return fmt.Errorf("Response sended")
	}
	rs.sended = true
	rs.muSended.Unlock()
	rs.deps.Response.Header().Set("Content-Type", "application/json")
	rs.deps.Response.WriteHeader(code)
	rs.deps.Response.Write([]byte(json))
	return nil
}

func (rs *ResponserService) Redirect(url string) {
	http.Redirect(rs.deps.Response, rs.deps.Request, url, http.StatusSeeOther)
}

// ResponseFactory create new Response service
func ResponserFactory(dp dependency.Provider) (interface{}, error) {
	responser := &ResponserService{
		sended: false,
	}
	if err := dp.InjectTo(&responser.deps); err != nil {
		return nil, err
	}
	return requestdep.Responser(responser), nil
}
