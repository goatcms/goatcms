package queries

import (
	"github.com/goatcms/goatcore/dependency"
)

func RegisterDependencies(dp dependency.Provider) error {
	if err := dp.AddDefaultFactory("UserSigninQuery", UserSigninQueryFactory); err != nil {
		return err
	}
	return nil
}
