package fragments

import (
	"html/template"
	"strconv"
	"strings"
	"time"

	"github.com/goatcms/goatcms/cmsapp/dao"
	"github.com/goatcms/goatcms/cmsapp/entities"
	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/scope"
	"github.com/goatcms/goatcore/dependency"
	"github.com/microcosm-cc/bluemonday"
	blackfriday "gopkg.in/russross/blackfriday.v2"
)

// Cache is fragment cache service
type Cache struct {
	deps struct {
		Logger          services.Logger     `dependency:"LoggerService"`
		AppScope        app.Scope           `dependency:"AppScope"`
		FragmentFindAll dao.FragmentFindAll `dependency:"FragmentFindAll"`
	}
	inited bool
	data   map[string]Row
}

// CacheFactory create new Cache instance
func CacheFactory(dp dependency.Provider) (in interface{}, err error) {
	cache := &Cache{
		data:   map[string]Row{},
		inited: false,
	}
	if err = dp.InjectTo(&cache.deps); err != nil {
		return nil, err
	}
	return services.FragmentCache(cache), nil
}

func (cache *Cache) init() {
	if !cache.inited {
		cache.inited = true
		go cache.startRefreshLoop()
	}
}

// startRefreshLoop start refresh cached data. Run it in background.
func (cache *Cache) startRefreshLoop() {
	var err error
	for {
		if err = cache.Refresh(); err != nil {
			cache.deps.Logger.ErrorLog("%v", err.Error())
		}
		time.Sleep(10 * time.Minute)
	}
}

// Refresh get new cached data
func (cache *Cache) Refresh() (err error) {
	var (
		data     = map[string]Row{}
		rows     dao.FragmentRows
		fragment *entities.Fragment
	)
	refreshScope := scope.NewScope("")
	if rows, err = cache.deps.FragmentFindAll.Find(refreshScope, &entities.FragmentFields{
		ID:      true,
		Lang:    true,
		Name:    true,
		Content: true,
	}); err != nil {
		return err
	}
	for rows.Next() {
		if fragment, err = rows.Get(); err != nil {
			return err
		}
		key := *fragment.Lang + "." + *fragment.Name
		unsafe := blackfriday.Run([]byte(*fragment.Content))
		html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
		data[key] = Row{
			ID:   *fragment.ID,
			HTML: string(html),
		}
	}
	cache.data = data
	return nil
}

// Get return fragment for key
func (cache *Cache) Get(key string) *Row {
	var (
		row Row
		ok  bool
	)
	cache.init()
	if row, ok = cache.data[key]; !ok {
		return nil
	}
	return &row
}

// RenderFragment return a HTML content for fragment. It is uset for small block with inline editor
func (cache *Cache) RenderFragment(key, defaultValue string) (result template.HTML) {
	var (
		row Row
		ok  bool
	)
	cache.init()
	if row, ok = cache.data[key]; !ok {
		row.ID = 0
		row.HTML = defaultValue
	}
	return template.HTML(strings.Join([]string{
		`<div class="fragment" g-small-fragment g-fragment-key="`,
		key,
		`" g-fragment-id="`,
		strconv.FormatInt(row.ID, 10),
		`">`,
		row.HTML,
		`</div>`,
	}, ""))
}
