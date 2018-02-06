package itbase

import (
	"github.com/goatcms/goatcms/cmsapp/commands/servec"
	"github.com/goatcms/goatcore/app"
)

func Run(mapp app.App) {
	go servec.Run(mapp)
}
