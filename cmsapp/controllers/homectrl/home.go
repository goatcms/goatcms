package homectrl

import (
	"fmt"
	"net/http"

	"github.com/goatcms/goat-core/app"
	"github.com/goatcms/goat-core/dependency"
	"github.com/goatcms/goatcms/cmsapp/services"
)

type Home struct {
	deps struct {
		Template services.Template `dependency:"TemplateService"`
	}
}

func NewHome(dp dependency.Provider) (*Home, error) {
	ctrl := &Home{}
	if err := dp.InjectTo(&ctrl.deps); err != nil {
		return nil, err
	}
	return ctrl, nil
}

func (c *Home) Get(requestScope app.Scope) {
	var requestDeps struct {
		RequestError services.RequestError `request:"RequestErrorService"`
		Response     http.ResponseWriter   `request:"Response"`
	}
	if err := requestScope.InjectTo(&requestDeps); err != nil {
		fmt.Println(err)
		return
	}
	if err := c.deps.Template.ExecuteTemplate(requestDeps.Response, "home/index", nil); err != nil {
		requestDeps.RequestError.Error(312, err)
		return
	}
}
