package thing

// TODO - make threadsafe

type Repo struct {
	id     int
	things Things
}

func NewRepo() *Repo {
	return &Repo{id: 0, things: Things{}}
}

func (r *Repo) GetAll() *Things {
	return &r.things
}

func (r *Repo) Create(t Thing) *Thing {
	r.id++
	t.ID = r.id
	r.things = append(r.things, t)
	return &t
}

func (r *Repo) Get(id int) *Thing {
	for i := range r.things {
		if r.things[i].ID == id {
			return &r.things[i]
		}
	}
	return nil
}

func (r *Repo) Update(id int, t Thing) *Thing {
	for i := range r.things {
		if r.things[i].ID == id {
			r.things[i].Val = t.Val
			return &r.things[i]
		}
	}
	return nil
}

func (r *Repo) Delete(id int) bool {
	for i := range r.things {
		if r.things[i].ID == id {
			r.things = append(r.things[:i], r.things[i+1:]...)
			return true
		}
	}
	return false
}
