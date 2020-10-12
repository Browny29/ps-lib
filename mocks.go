package main

type MockedService struct {}

func NewMockedService() *MockedService {
	return new(MockedService)
}

func (s *MockedService) GetProcDirectories() ([]string, error) {
	return []string{"13936"}, nil
}

func (s *MockedService) GetProcByPID(pid string) ([]byte, error) {
	return []byte("13936 (bash) S 1 13936 13936 196608 -1 0 3436 3436 0 0 78 359 78 359 8 0 0 0 0 0 0 00854085461 7196672 2628 345"), nil
}
