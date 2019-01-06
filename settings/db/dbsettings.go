package db

type DbSettings interface {
	isEnableMongoDb() bool
	isEnablePostgres() bool
}

func IsEnableMongoDb() bool {
	return true
}

func IsEnablePostgres() bool {
	return false
}
