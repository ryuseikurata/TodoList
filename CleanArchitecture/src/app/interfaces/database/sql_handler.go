package database

type SqlHandler interface {
	Execute(string, ...interface{}) (Result, error)
	Query(string, ...interface{}) (Row, error)
}

type Result interface {
	LastInsertId() (int64, error)
	Rowsaffected() (int64, error)
}

type Row interface {
	Scan(...interface{}) error
	Next() bool
	Close() bool
}