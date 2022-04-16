package storage

import "server1/internal/grpcapi/protobuff/user"

type StubDB struct {
	data []User
}

type User struct {
	Id   int32
	Name string
}

func (s *StubDB) GetUserByID(id int32) *user.UserResponse {
	for i, val := range s.data {
		if int32(i+1) == id {
			return &user.UserResponse{
				Id:   val.Id,
				Name: val.Name,
			}
		}
	}

	return &user.UserResponse{
		Id:   0,
		Name: "Test",
	}
}
