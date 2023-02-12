package testin

import (
	"context"
	"database/sql"
)

// create interface of the database instance that
type DTBX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

// create a struct to subscribe to dtbx
type testrepo struct {
	db DTBX
}

// CreatTestRepository implements Testrepo
// func (s *testrepo) CreatTestRepository(ctx context.Context, test *TestIn) (*TestIn, error) {
// 	panic("unimplemented")
// }

// Create a constructor to return repository interface and database
func NewTestrepo(db DTBX) Testrepo {
	return &testrepo{
		db: db,
	}
}

func (t *testrepo) CreatTestRepository(ctx context.Context,test *TestIn)(*TestIn,error){
	//perform query insertion here
	var lastinsertid int
	query := "INSERT INTO tests (name) VALUES ($1) returning id"
	err:= t.db.QueryRowContext(ctx, query,test.NAME).Scan(&lastinsertid)
	if err != nil {
		return &TestIn{},err
	}
	//typecast id to interface
	test.ID=int64(lastinsertid)
	return test, nil
}
