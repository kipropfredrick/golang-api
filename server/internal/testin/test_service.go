package testin

import "context"

type testSerStruct struct {
	Testrepo
}

// constructor to create a new test
func NewTestservice(test Testrepo) TestService {
	return &testSerStruct{
		test,
	}
}

// functuion to create test service
func (s *testSerStruct) CreateTest(ctx context.Context, req *TestRequest) (*TestResponse, error) {
	tes :=&TestIn{
		NAME: req.NAME,
	}
	resp,err:=s.Testrepo.CreatTestRepository(ctx,tes)
	if err != nil {
		return nil, err
	}
	//response
	r:=&TestResponse{
		resp.ID,
		resp.NAME,
	}
	return r, nil
}