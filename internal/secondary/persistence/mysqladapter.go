package persistence

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type MySqlAdapter struct {
	driver, user, password, addr, port, dbname string
}

func NewMySqlAdapter(driver, user, password, addr, port, dbname string) *MySqlAdapter {
	return &MySqlAdapter{driver, user, password, addr, port, dbname}
}

func (m *MySqlAdapter) TestConnection() {
	db := m.Open()
	if err := db.Close(); err != nil {
		log.Fatal("Couldn't close database connection")
	}
}

func (m *MySqlAdapter) Open() *sql.DB {
	db, err := sql.Open(
		m.driver,
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", m.user, m.password, m.addr, m.port, m.dbname),
	)
	if err != nil {
		log.Fatal("Couldn't connect to database")
	}
	return db
}

func (m *MySqlAdapter) Exec(query string, args ...string) (int, error) {
	db := m.Open()
	defer closeDB(db)
	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}
	result, err := tx.Exec(query, args)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), tx.Commit()
}

func (m *MySqlAdapter) Query(query string, args ...string) (*sql.Row, error) {
	db := m.Open()
	defer closeDB(db)
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	result := tx.QueryRow(query, args)
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func closeDB(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Println(err.Error())
	}
}
