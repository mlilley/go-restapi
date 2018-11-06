package thing

import (
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"
)

type ThingSqlite3Repo struct {
	db *sql.DB
}

func NewThingSqlite3Repo() (*ThingSqlite3Repo, error) {
	db, err := sql.Open("sqlite3", "things.sqlite3")
	if err != nil {
		return nil, err
	}
	if db == nil {
		return nil, errors.New("nil db")
	}

	_, err = db.Exec(`DROP TABLE IF EXISTS things;`)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE things(id INTEGER PRIMARY KEY AUTOINCREMENT, val INTEGER);`)
	if err != nil {
		return nil, err
	}

	return &ThingSqlite3Repo{db: db}, nil
}

func (r *ThingSqlite3Repo) Get(id int64) (*Thing, error) {
	rows, err := r.db.Query(`SELECT id, val FROM things WHERE id = ?;`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		t := Thing{}
		err := rows.Scan(&t.ID, &t.Val)
		if err != nil {
			return nil, err
		}
		return &t, nil
	}

	return nil, nil
}

func (r *ThingSqlite3Repo) GetAll() (*Things, error) {
	rows, err := r.db.Query(`SELECT ID, Val FROM things ORDER BY ID;`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ts := Things{}
	for rows.Next() {
		t := Thing{}
		err := rows.Scan(&t.ID, &t.Val)
		if err != nil {
			return nil, err
		}
		ts = append(ts, t)
	}

	return &ts, nil
}

func (r *ThingSqlite3Repo) Create(t *Thing) (*Thing, error) {
	res, err := r.db.Exec(`INSERT INTO things (val) VALUES (?);`, t.Val)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &Thing{ID: id, Val: t.Val}, nil
}

func (r *ThingSqlite3Repo) Update(id int64, t *Thing) (*Thing, error) {
	res, err := r.db.Exec(`UPDATE things SET val = ? WHERE id = ?;`, t.Val, id)
	if err != nil {
		return nil, err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if n == 0 {
		return nil, nil
	}

	return &Thing{ID: id, Val: t.Val}, nil
}

func (r *ThingSqlite3Repo) Delete(id int64) (bool, error) {
	res, err := r.db.Exec(`DELETE FROM things WHERE id = ?;`, id)
	if err != nil {
		return false, err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	return n != 0, nil
}
