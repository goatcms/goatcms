package services

import "github.com/goatcms/goat-core/dependency"

// DefaultProvider is external dependency provider implementation
type DefaultProvider struct {
	dependency.Provider
}

// NewProvider create new service dependency provider
func NewProvider() Provider {
	return &DefaultProvider{dependency.NewProvider()}
}

// Database return instance of default database implementation
func (p *DefaultProvider) Database() (Database, error) {
	ins, err := p.Get(DBID)
	if err != nil {
		return nil, err
	}
	return ins.(Database), nil
}

// Mux return instance of default mux implementation
func (p *DefaultProvider) Mux() (Mux, error) {
	ins, err := p.Get(MuxID)
	if err != nil {
		return nil, err
	}
	return ins.(Mux), nil
}

// Template return instance of default template implementation
func (p *DefaultProvider) Template() (Template, error) {
	ins, err := p.Get(TemplateID)
	if err != nil {
		return nil, err
	}
	return ins.(Template), nil
}

// Crypt return instance of default crypt implementation
func (p *DefaultProvider) Crypt() (Crypt, error) {
	ins, err := p.Get(CryptID)
	if err != nil {
		return nil, err
	}
	return ins.(Crypt), nil
}

// Auth return instance of default auth implementation
func (p *DefaultProvider) Auth() (Auth, error) {
	ins, err := p.Get(AuthID)
	if err != nil {
		return nil, err
	}
	return ins.(Auth), nil
}

// SessionManager return instance of default session manager implementation
func (p *DefaultProvider) SessionManager() (SessionManager, error) {
	ins, err := p.Get(SessionManagerID)
	if err != nil {
		return nil, err
	}
	return ins.(SessionManager), nil
}
