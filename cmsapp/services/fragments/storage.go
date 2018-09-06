package fragments

import (
	"sync"
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

// Storage is fragment storage service
type Storage struct {
	deps struct {
		Logger          services.Logger     `dependency:"LoggerService"`
		EngineScope     app.Scope           `dependency:"EngineScope"`
		FragmentFindAll dao.FragmentFindAll `dependency:"FragmentFindAll"`
	}
	expirationDuration  time.Duration
	expirationDeadline  time.Time
	initedMU            sync.Mutex
	inited              bool
	dataMU              sync.RWMutex
	data                map[string]*services.Fragment
	refreshCounterMU    sync.Mutex
	lastExecutedCounter int64
	refreshCounter      int64
}

// StorageFactory create new Storage instance
func StorageFactory(dp dependency.Provider) (in interface{}, err error) {
	storage := &Storage{
		data:                map[string]*services.Fragment{},
		inited:              false,
		expirationDuration:  2 * time.Hour,
		lastExecutedCounter: 0,
		refreshCounter:      0,
	}
	if err = dp.InjectTo(&storage.deps); err != nil {
		return nil, err
	}
	return services.FragmentStorage(storage), nil
}

func (storage *Storage) init() {
	// check without lock rutine
	if storage.inited {
		return
	}
	// lock rutine and init code
	storage.initedMU.Lock()
	defer storage.initedMU.Unlock()
	if !storage.inited {
		storage.deps.EngineScope.On(dao.FragmentDeleteEvent, storage.RefreshHandler)
		storage.deps.EngineScope.On(dao.FragmentUpdateEvent, storage.RefreshHandler)
		storage.deps.EngineScope.On(dao.FragmentInsertEvent, storage.RefreshHandler)
		storage.Refresh()
		go storage.runRefreshLoop()
		storage.inited = true
	}
}

// runRefreshLoop start refresh storaged data endless loop. Run it in background.
// It is used to sync changes from other app instances
func (storage *Storage) runRefreshLoop() {
	for {
		time.Sleep(storage.expirationDuration)
		if time.Now().After(storage.expirationDeadline) {
			storage.Refresh()
		}
	}
}

// RefreshHandler is event callback
func (storage *Storage) RefreshHandler(in interface{}) (err error) {
	return storage.Refresh()
}

// Refresh get new storaged data
func (storage *Storage) Refresh() (err error) {
	storage.newRefreshCounter()
	storage.dataMU.Lock()
	defer storage.dataMU.Unlock()
	// do it for all waiting task
	//for storage.lastExecutedCounter < storage.refreshCounter {
	storage.lastExecutedCounter = storage.refreshCounter
	if err = storage.refresh(); err != nil {
		storage.deps.Logger.ErrorLog("%v", err.Error())
		return err
	}
	//}
	return nil
}

func (storage *Storage) refresh() (err error) {
	var (
		data     = map[string]*services.Fragment{}
		rows     dao.FragmentRows
		fragment *entities.Fragment
	)
	// execute only last policy - break if new event come
	if storage.lastExecutedCounter != storage.refreshCounter {
		return nil
	}
	// get fragments from database
	refreshScope := scope.NewScope("")
	if rows, err = storage.deps.FragmentFindAll.Find(refreshScope, &entities.FragmentFields{
		ID:      true,
		Lang:    true,
		Name:    true,
		Content: true,
	}); err != nil {
		return err
	}
	defer rows.Close()
	// process pragments
	for rows.Next() {
		// execute only last policy - break if new event come
		if storage.lastExecutedCounter != storage.refreshCounter {
			return nil
		}
		// process single fragment dao
		if fragment, err = rows.Get(); err != nil {
			return err
		}
		key := *fragment.Lang + "." + *fragment.Name
		unsafe := blackfriday.Run([]byte(*fragment.Content))
		html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
		data[key] = &services.Fragment{
			ID:   *fragment.ID,
			HTML: string(html),
		}
	}
	storage.data = data
	storage.expirationDeadline = time.Now().Add(storage.expirationDuration)
	return nil
}

func (storage *Storage) newRefreshCounter() int64 {
	storage.refreshCounterMU.Lock()
	defer storage.refreshCounterMU.Unlock()
	storage.refreshCounter++
	return storage.refreshCounter
}

// Get return fragment for key
func (storage *Storage) Get(key string) (result *services.Fragment) {
	var ok bool
	storage.init()
	storage.dataMU.RLock()
	defer storage.dataMU.RUnlock()
	if result, ok = storage.data[key]; !ok {
		return nil
	}
	return result
}
