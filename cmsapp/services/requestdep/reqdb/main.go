package reqdb

import (
	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcms/cmsapp/services/requestdep"
	"github.com/goatcms/goatcore/db"
	"github.com/goatcms/goatcore/dependency"
)

// RequestDB provide request database transaction
type RequestDB struct {
	deps struct {
		DatabaseService services.Database `dependency:"DatabaseService"`
	}
	tx db.TX
}

// RequestDBFactory create an database service instance
func RequestDBFactory(dp dependency.Provider) (interface{}, error) {
	instance := &RequestDB{}
	if err := dp.InjectTo(&instance.deps); err != nil {
		return nil, err
	}
	return requestdep.DB(instance), nil
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
