package interfaces

type DbHandler interface {
	Execute(statement string)
	Query(statement string) (Row, error)
}

type Row interface {
	Scan(dest ...interface{}) error
	Next() bool
}