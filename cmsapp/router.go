package cmsapp

import (
	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/varutil"
)

// RouterStaticMappingRow is single static mapping configuration record
type RouterStaticMappingRow struct {
	Prefix string `json:"prefix"`
	Path   string `json:"path"`
}

// RouterMappingConfig contains mapping configuration
type RouterMappingConfig struct {
	Static []RouterStaticMappingRow `json:"static"`
}

// InitRouter load routing configuration
func InitRouter(a app.App) (err error) {
	var (
		fs     = a.RootFilespace()
		config RouterMappingConfig
		bytes  []byte
		deps   struct {
			Router services.Router `dependency:"RouterService"`
			Logger services.Logger `dependency:"LoggerService"`
		}
	)
	if err = a.DependencyProvider().InjectTo(&deps); err != nil {
		return err
	}
	if !fs.IsExist(RouterMappingFile) {
		deps.Logger.DevLog("Routing loading is skipped. File %s is not exist.")
		return nil
	}
	if bytes, err = fs.ReadFile(RouterMappingFile); err != nil {
		return err
	}
	if err = varutil.ObjectFromJSON(&config, string(bytes)); err != nil {
		return err
	}
	for _, row := range config.Static {
		deps.Logger.DevLog("Static routing: %s -> %s", row.Prefix, row.Path)
		if deps.Router.ServeStatic(row.Prefix, row.Path); err != nil {
			return err
		}
	}
	return nil
}
