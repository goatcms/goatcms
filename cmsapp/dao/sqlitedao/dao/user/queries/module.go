package queries

import (
	"github.com/goatcms/goatcore/dependency"
)

func RegisterDependencies(dp dependency.Provider) error {
	if err := dp.AddDefaultFactory("UserLoginQuery", UserLoginQueryFactory); err != nil {
		return err
	}
	return nil
}
