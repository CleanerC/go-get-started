package v1_service

import "context"

type Add struct {
	A int `json:"a"`
	B int `json:"b"`
}

type AddAck struct {
	Res int `json:"res"`
}

type Service interface {
	TestAdd(ctx context.Context, nums Add) AddAck
}

type baseServer struct {}

func NewService() Service {
	return &baseServer{}
}

func (server baseServer) TestAdd(ctx context.Context, in Add) AddAck {
	return AddAck{Res:in.A + in.B}
} 