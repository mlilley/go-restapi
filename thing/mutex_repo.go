package thing

import (
	"sync"
)

type ThingMutexRepo struct {
	id     int64
	things Things
	lock   *sync.RWMutex
}

func NewThingMutexRepo() (*ThingMutexRepo, error) {
	return &ThingMutexRepo{id: 0, things: Things{}, lock: &sync.RWMutex{}}, nil
}

func (r *ThingMutexRepo) Get(id int64) (*Thing, error) {
	r.lock.RLock()
	for i := range r.things {
		if r.things[i].ID == id {
			t := r.things[i]
			r.lock.RUnlock()
			return &t, nil
		}
	}
	r.lock.RUnlock()
	return nil, nil
}

func (r *ThingMutexRepo) GetAll() (*Things, error) {
	r.lock.RLock()
	thingsCopy := make(Things, len(r.things))
	copy(thingsCopy, r.things)
	r.lock.RUnlock()
	return &thingsCopy, nil
}

func (r *ThingMutexRepo) Create(t *Thing) (*Thing, error) {
	r.lock.Lock()
	r.id++
	tCopy := Thing{ID: r.id, Val: t.Val}
	r.things = append(r.things, tCopy) // (copy of tCopy stored)
	r.lock.Unlock()
	return &tCopy, nil
}

func (r *ThingMutexRepo) Update(id int64, t *Thing) (*Thing, error) {
	r.lock.Lock()
	for i := range r.things {
		if r.things[i].ID == id {
			r.things[i].Val = t.Val
			tt := r.things[i]
			r.lock.Unlock()
			return &tt, nil
		}
	}
	r.lock.Unlock()
	return nil, nil
}

func (r *ThingMutexRepo) Delete(id int64) (bool, error) {
	r.lock.Lock()
	for i := range r.things {
		if r.things[i].ID == id {
			r.things = append(r.things[:i], r.things[i+1:]...)
			r.lock.Unlock()
			return true, nil
		}
	}
	r.lock.Unlock()
	return false, nil
}
