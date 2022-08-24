package app

import "os"

type database interface {
	NewDB() interface{}
}

func GetDatabaseName() string {
	return os.Getenv("DB_NAME")
}
