package session

import (
	"github.com/goatcms/goat-core/dependency"
	"github.com/goatcms/goat-core/http/post"
	"github.com/goatcms/goatcms/services"
)

// HttpPostDecoderFactory build new post decoder
func HTTPPostDecoderFactory(dp dependency.Provider) (dependency.Instance, error) {
	return post.NewDecoder()
}

// InitDep initialize a new session manager dependency
func InitDep(prov dependency.Provider) error {
	if err := prov.AddService(services.SessionManagerID, HTTPPostDecoderFactory); err != nil {
		return err
	}
	return nil
}
