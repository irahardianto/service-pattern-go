package interfaces

type IDbHandler interface {
	Execute(statement string)
	Query(statement string) (IRow, error)
}

type IRow interface {
	Scan(dest ...interface{}) error
	Next() bool
}