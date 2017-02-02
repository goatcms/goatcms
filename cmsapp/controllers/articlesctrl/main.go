package articles

import (
	"github.com/goatcms/goat-core/http/post"
	"github.com/goatcms/goat-core/types"
	"github.com/goatcms/goatcms/models"
	"github.com/goatcms/goatcms/services"
)

// Dependency is default set of dependency
type Dependency struct {
	DP       services.Provider `inject:"provider"`
	Template services.Template `inject:"template"`
	Mux      services.Mux      `inject:"mux"`
	Database services.Database `inject:"database"`

	ArticleType    types.CustomType  `inject:"type.article"`
	ArticleDAO     models.ArticleDAO `inject:"dao.article"`
	ArticleDecoder *post.Decoder     `inject:"decoder.article"`
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
	d := &Dependency{}
	if err := dp.InjectTo(d); err != nil {
		return nil, err
	}
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
