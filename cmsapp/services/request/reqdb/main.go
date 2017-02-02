package reqdb

import (
	"github.com/goatcms/goat-core/db"
	"github.com/goatcms/goat-core/dependency"
	"github.com/goatcms/goatcms/cmsapp/services"
)

// RequestDB provide request database transaction
type RequestDB struct {
	deps struct {
		DatabaseService services.Database `dependency:"DatabaseService"`
	}
	tx db.TX
}

// AuthFactory create an authentification service instance
func RequestDBFactory(dp dependency.Provider) (interface{}, error) {
	instance := &RequestDB{}
	if err := dp.InjectTo(&instance.deps); err != nil {
		return nil, err
	}
	return services.RequestDB(instance), nil
}

// TX return request transaction request scope singleton
func (rdb *RequestDB) TX() (db.TX, error) {
	var err error
	if rdb.tx == nil {
		rdb.tx, err = rdb.deps.DatabaseService.TX()
		if err != nil {
			return nil, err
		}
	}
	return rdb.tx, nil
}
