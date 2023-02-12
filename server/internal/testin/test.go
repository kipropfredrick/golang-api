package testin

import "context"

type TestIn struct {
	ID   int64  `json:"id"`
	NAME string `json:"name"`
}

// test response
type TestResponse struct {
	ID   int64  `json:"id"`
	NAME string `json:"name"`
}

// test request
type TestRequest struct {
	NAME string `json:"name"`
}

// service test interface
type TestService interface {
	CreateTest(ctx context.Context, req *TestRequest) (*TestResponse, error)
}
//repository interface
type Testrepo interface {
	CreatTestRepository(ctx context.Context, test *TestIn) (*TestIn, error)
}