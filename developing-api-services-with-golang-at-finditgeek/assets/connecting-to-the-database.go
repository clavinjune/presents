const (
	queryMigrate string = `CREATE TABLE IF NOT EXISTS tasks(
		id TEXT PRIMARY KEY,
		name TEXT NOT NULL UNIQUE
	)`

	querySeed string = `INSERT INTO tasks VALUES('seed1', 'task seed 1')`
)

// this function is simplified // HL1
// please don't use this function in production // HL1
func initializeDatabase() *sql.DB {
	db, _ := sql.Open("sqlite3", ":memory:")
	db.Exec(queryMigrate)
	db.Exec(querySeed)
	return db
}
