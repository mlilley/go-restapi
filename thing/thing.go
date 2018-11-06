package thing

type Thing struct {
	ID  int64 `json:"id"`
	Val int   `json:"val"`
}

type Things []Thing

type ThingRepo interface {
	Get(int64) (*Thing, error)
	GetAll() (*Things, error)
	Create(*Thing) (*Thing, error)
	Update(int64, *Thing) (*Thing, error)
	Delete(int64) (bool, error)
}
