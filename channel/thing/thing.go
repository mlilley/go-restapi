package thing

type Thing struct {
	ID  int `json:"id"`
	Val int `json:"val"`
}

type Things []Thing
