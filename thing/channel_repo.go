package thing

type ThingChannelRepo struct {
	id         int64
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
	id  int64
	res chan *Thing
}

type createReq struct {
	t   *Thing
	res chan *Thing
}

type updateReq struct {
	id  int64
	t   *Thing
	res chan *Thing
}

type deleteReq struct {
	id  int64
	res chan bool
}

// public

func NewThingChannelRepo() *ThingChannelRepo {
	r := ThingChannelRepo{
		id:         0,
		things:     Things{},
		getChan:    make(chan *getReq),
		getAllChan: make(chan *getAllReq),
		createChan: make(chan *createReq),
		updateChan: make(chan *updateReq),
		deleteChan: make(chan *deleteReq),
	}

	go func(r *ThingChannelRepo) {
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

func (r *ThingChannelRepo) Get(id int64) (*Thing, error) {
	res := make(chan *Thing)
	r.getChan <- &getReq{id: id, res: res}
	return <-res, nil
}

func (r *ThingChannelRepo) GetAll() (*Things, error) {
	res := make(chan *Things)
	r.getAllChan <- &getAllReq{res: res}
	return <-res, nil
}

func (r *ThingChannelRepo) Create(t *Thing) (*Thing, error) {
	res := make(chan *Thing)
	r.createChan <- &createReq{t: t, res: res}
	return <-res, nil
}

func (r *ThingChannelRepo) Update(id int64, t *Thing) (*Thing, error) {
	res := make(chan *Thing)
	r.updateChan <- &updateReq{id: id, t: t, res: res}
	return <-res, nil
}

func (r *ThingChannelRepo) Delete(id int64) (bool, error) {
	res := make(chan bool)
	r.deleteChan <- &deleteReq{id: id, res: res}
	return <-res, nil
}

// private

func (r *ThingChannelRepo) get(id int64) *Thing {
	for i := range r.things {
		if r.things[i].ID == id {
			t := r.things[i]
			return &t
		}
	}
	return nil
}

func (r *ThingChannelRepo) getAll() *Things {
	thingsCopy := make(Things, len(r.things))
	copy(thingsCopy, r.things)
	return &thingsCopy
}

func (r *ThingChannelRepo) create(t *Thing) *Thing {
	r.id++
	tt := Thing{ID: r.id, Val: t.Val}
	r.things = append(r.things, tt)
	return &tt
}

func (r *ThingChannelRepo) update(id int64, t *Thing) *Thing {
	for i := range r.things {
		if r.things[i].ID == id {
			r.things[i].Val = t.Val
			tt := r.things[i]
			return &tt
		}
	}
	return nil
}

func (r *ThingChannelRepo) delete(id int64) bool {
	for i := range r.things {
		if r.things[i].ID == id {
			r.things = append(r.things[:i], r.things[i+1:]...)
			return true
		}
	}
	return false
}
