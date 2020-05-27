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
