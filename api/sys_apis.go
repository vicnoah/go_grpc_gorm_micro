package api

import (
	"context"
	"go_grpc_gorm_micro/proto/proto"
	"go_grpc_gorm_micro/lib/response"
	"go_grpc_gorm_micro/service"
)

type SysApis struct{}

// 生成curd代码
func (s *SysApis) Create(ctx context.Context, req *proto.SysApis) (*proto.Response, error) {
	data, err := service.CreateSysApis(req)
	return response.SuccessAny(data), err
}


func (s *SysApis) Delete(ctx context.Context, req *proto.SysApis) (*proto.Response, error) {
	data, err := service.DeleteSysApis(req)
	return response.SuccessAny(data), err
}

func (s *SysApis) DeleteById(ctx context.Context, req *proto.SysApis) (*proto.Response, error) {
	data, err := service.DeleteByIdSysApis(req)
	return response.SuccessAny(data), err
}

func (s *SysApis) Update(ctx context.Context, req *proto.SysApis) (*proto.Response, error) {
	data, err := service.UpdateSysApis(req)
	return response.SuccessAny(data), err
}

func (s *SysApis) Find(ctx context.Context, req *proto.SysApis) (*proto.Response, error) {
	data, err := service.FindSysApis(req)
	return response.SuccessAny(data), err
}

func (s *SysApis) Lists(ctx context.Context, req *proto.Request) (*proto.Responses, error) {
	data, err := service.GetListSysApis(req)
	return data, err
	//return response.SuccessAny(data), err
}


