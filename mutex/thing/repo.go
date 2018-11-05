package thing

import (
	"sync"
)

type Repo struct {
	id     int
	things Things
	lock   *sync.RWMutex
}

func NewRepo() *Repo {
	return &Repo{id: 0, things: Things{}, lock: &sync.RWMutex{}}
}

func (r *Repo) GetAll() *Things {
	r.lock.RLock()
	thingsCopy := make(Things, len(r.things))
	copy(thingsCopy, r.things)
	r.lock.RUnlock()
	return &thingsCopy
}

func (r *Repo) Create(t Thing) *Thing {
	r.lock.Lock()
	r.id++
	t.ID = r.id
	r.things = append(r.things, t) // copy of t stored
	r.lock.Unlock()
	return &t
}

func (r *Repo) Get(id int) *Thing {
	r.lock.RLock()
	for i := range r.things {
		if r.things[i].ID == id {
			t := r.things[i]
			r.lock.RUnlock()
			return &t
		}
	}
	r.lock.RUnlock()
	return nil
}

func (r *Repo) Update(id int, t Thing) *Thing {
	r.lock.Lock()
	for i := range r.things {
		if r.things[i].ID == id {
			r.things[i].Val = t.Val
			tt := r.things[i]
			r.lock.Unlock()
			return &tt
		}
	}
	r.lock.Unlock()
	return nil
}

func (r *Repo) Delete(id int) bool {
	r.lock.Lock()
	for i := range r.things {
		if r.things[i].ID == id {
			r.things = append(r.things[:i], r.things[i+1:]...)
			r.lock.Unlock()
			return true
		}
	}
	r.lock.Unlock()
	return false
}
