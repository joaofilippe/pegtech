package database

// NewDatabase creates a new database connection
func NewDatabase() (*PostgresDB, error) {
	config := NewDBConfig()
	return NewPostgresDB(
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.DBName,
	)
}
