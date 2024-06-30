package database

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

// Database instance => DB Gorm connector
var DBConn *pgxpool.Pool
