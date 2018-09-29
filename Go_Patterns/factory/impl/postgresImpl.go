package impl

import (
	"database/sql"
)

type PostgreSQLDataStore struct {
	DSN string
	DB  sql.DB
}

func (pds *PostgreSQLDataStore) Name() string {
	return "PostgreSQLDataStore"
}

func (pds *PostgreSQLDataStore) FindUserById(id int64) (string, error) {
	//query pds.DB and get name
	return "", nil
}
