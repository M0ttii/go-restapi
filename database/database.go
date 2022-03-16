package database

import (
	"context"
	"database/sql"
)

type Database struct {
	SqlDB *sql.DB
}

var dbContext = context.Background()
