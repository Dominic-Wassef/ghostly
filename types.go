package ghostly

import "database/sql"

// initPaths is used when initializing the application. It holds the root
// path for the application, and a slice of strings with the names of
// folders that the application expects to find.
type initPaths struct {
	rootPath    string
	folderNames []string
}

// cookieConfig holds cookie config values
type cookieConfig struct {
	name     string
	lifetime string
	persist  string
	secure   string
	domain   string
}

type databaseConfig struct {
	dsn      string
	database string
}

type Database struct {
	DataType string
	Pool     *sql.DB
}
