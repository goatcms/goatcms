package services

import (
	"github.com/goatcms/goat-core/dependency"
	"github.com/goatcms/goatcms/models"
)

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

// Files return instance of default files implementation
func (p *DefaultProvider) Files() (Files, error) {
	ins, err := p.Get(FilesID)
	if err != nil {
		return nil, err
	}
	return ins.(Files), nil
}

// UserDAO return instance of default user dao
func (p *DefaultProvider) UserDAO() (models.UserDAO, error) {
	ins, err := p.Get(models.UserDAOID)
	if err != nil {
		return nil, err
	}
	return ins.(models.UserDAO), nil
}

// ArticleDAO return instance of default article dao
func (p *DefaultProvider) ArticleDAO() (models.ArticleDAO, error) {
	ins, err := p.Get(models.ArticleDAOID)
	if err != nil {
		return nil, err
	}
	return ins.(models.ArticleDAO), nil
}

// ImageDAO return instance of default image dao
func (p *DefaultProvider) ImageDAO() (models.ImageDAO, error) {
	ins, err := p.Get(models.ImageDAOID)
	if err != nil {
		return nil, err
	}
	return ins.(models.ImageDAO), nil
}
