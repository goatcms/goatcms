package articles

import (
	"github.com/goatcms/goat-core/http/post"
	"github.com/goatcms/goatcms/models"
	"github.com/goatcms/goatcms/models/article"
	"github.com/goatcms/goatcms/services"
)

// Dependency is default set of dependency
type Dependency struct {
	DP       services.Provider
	Template services.Template
	Mux      services.Mux
	Database services.Database

	ArticleDAO     models.ArticleDAO
	ArticleDecoder *post.Decoder
}

const (
	// InsertURL is url to insert page
	InsertURL = "/article/add"
	// ListURL is url to articles list page
	ListURL = "/article"
	// ViewURL is url to single article page
	ViewURL = "/article/{id:[0-9]+}"
)

// NewDependency is default set of dependency
func NewDependency(dp services.Provider) (*Dependency, error) {
	var err error
	d := &Dependency{
		DP: dp,
	}
	if d.Template, err = dp.Template(); err != nil {
		return nil, err
	}
	if d.Mux, err = dp.Mux(); err != nil {
		return nil, err
	}
	if d.Database, err = dp.Database(); err != nil {
		return nil, err
	}
	if d.ArticleDAO, err = dp.ArticleDAO(); err != nil {
		return nil, err
	}
	articleType := articlemodel.NewArticleType()
	d.ArticleDecoder = post.NewDecoder(articleType)
	return d, nil
}

// Init initialize the article controllers package
func Init(dp services.Provider) error {
	d, err := NewDependency(dp)
	if err != nil {
		return err
	}
	insertCtrl := NewInsertCtrl(d)
	listCtrl := NewListCtrl(d)
	viewCtrl := NewViewCtrl(d)
	d.Mux.Get(InsertURL, insertCtrl.Get)
	d.Mux.Post(InsertURL, insertCtrl.Post)
	d.Mux.Get(ListURL, listCtrl.Get)
	d.Mux.Get(ViewURL, viewCtrl.Get)
	return nil
}
