package utilz

import (
	"database/sql"
	"fmt"
	spb "github.com/explodes/serving/proto"
)

const (
	sqliteDriver = "sqlite3"
	pgDriver     = "postgres"
)

func OpenDatabaseAddress(database *spb.DatabaseAddress) (*sql.DB, error) {
	switch t := database.DatabaseAddress.(type) {
	case *spb.DatabaseAddress_Sqlite3:
		return OpenSqlite3Address(t.Sqlite3)
	case *spb.DatabaseAddress_Postgres:
		return OpenPostgresAddress(t.Postgres)
	default:
		return nil, fmt.Errorf("missing/unknown database connection type")
	}
}

func OpenSqlite3Address(addr *spb.Sqlite3Address) (*sql.DB, error) {
	db, err := sql.Open(sqliteDriver, addr.Url)
	if err != nil {
		return nil, fmt.Errorf("error opening sqlite3: %v", err)
	}
	return db, nil
}

func OpenPostgresAddress(addr *spb.PostgresAddress) (*sql.DB, error) {
	db, err := sql.Open(pgDriver, addr.Url)
	if err != nil {
		return nil, fmt.Errorf("error opening postgres: %v", err)
	}
	return db, nil
}
