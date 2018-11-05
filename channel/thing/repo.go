package thing

type Repo struct {
	id         int
	things     Things
	getChan    chan *getReq
	getAllChan chan *getAllReq
	createChan chan *createReq
	updateChan chan *updateReq
	deleteChan chan *deleteReq
}

type getAllReq struct {
	res chan *Things
}

type getReq struct {
	id  int
	res chan *Thing
}

type createReq struct {
	t   *Thing
	res chan *Thing
}

type updateReq struct {
	id  int
	t   *Thing
	res chan *Thing
}

type deleteReq struct {
	id  int
	res chan bool
}

func NewRepo() *Repo {
	r := Repo{
		id:         0,
		things:     Things{},
		getChan:    make(chan *getReq),
		getAllChan: make(chan *getAllReq),
		createChan: make(chan *createReq),
		updateChan: make(chan *updateReq),
		deleteChan: make(chan *deleteReq),
	}

	go func(r *Repo) {
		for {
			select {
			case getReq := <-r.getChan:
				getReq.res <- r.get(getReq.id)
			case getAllReq := <-r.getAllChan:
				getAllReq.res <- r.getAll()
			case createReq := <-r.createChan:
				createReq.res <- r.create(createReq.t)
			case updateReq := <-r.updateChan:
				updateReq.res <- r.update(updateReq.id, updateReq.t)
			case deleteReq := <-r.deleteChan:
				deleteReq.res <- r.delete(deleteReq.id)
			}
		}
	}(&r)

	return &r
}

// -- private versions of repo operations --

func (r *Repo) get(id int) *Thing {
	for i := range r.things {
		if r.things[i].ID == id {
			t := r.things[i]
			return &t
		}
	}
	return nil
}

func (r *Repo) getAll() *Things {
	thingsCopy := make(Things, len(r.things))
	copy(thingsCopy, r.things)
	return &thingsCopy
}

func (r *Repo) create(t *Thing) *Thing {
	r.id++
	t.ID = r.id
	r.things = append(r.things, *t)
	return t
}

func (r *Repo) update(id int, t *Thing) *Thing {
	for i := range r.things {
		if r.things[i].ID == id {
			r.things[i].Val = t.Val
			tt := r.things[i]
			return &tt
		}
	}
	return nil
}

func (r *Repo) delete(id int) bool {
	for i := range r.things {
		if r.things[i].ID == id {
			r.things = append(r.things[:i], r.things[i+1:]...)
			return true
		}
	}
	return false
}

// -- public repo operations --

func (r *Repo) Get(id int) *Thing {
	res := make(chan *Thing)
	r.getChan <- &getReq{id: id, res: res}
	return <-res
}

func (r *Repo) GetAll() *Things {
	res := make(chan *Things)
	r.getAllChan <- &getAllReq{res: res}
	return <-res
}

func (r *Repo) Create(t *Thing) *Thing {
	res := make(chan *Thing)
	r.createChan <- &createReq{t: t, res: res}
	return <-res
}

func (r *Repo) Update(id int, t *Thing) *Thing {
	res := make(chan *Thing)
	r.updateChan <- &updateReq{id: id, t: t, res: res}
	return <-res
}

func (r *Repo) Delete(id int) bool {
	res := make(chan bool)
	r.deleteChan <- &deleteReq{id: id, res: res}
	return <-res
}
