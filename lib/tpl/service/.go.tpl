package service

import (
	"errors"
	"github.com/golang/protobuf/ptypes"
	"github.com/jinzhu/gorm"
	"go_grpc_gorm_micro/lib/constant"
	"go_grpc_gorm_micro/lib/global"
	"go_grpc_gorm_micro/proto/proto"
)

func Create{{.ModelName}}(req *proto.{{.ModelName}}) (*proto.{{.ModelName}}, error) {
	// validator校验
	data := req

	//
	if !errors.Is(global.CURD_DB.Where("path = ? AND method = ?", req.Path, req.Method).First(&proto.{{.ModelName}}{}).Error, gorm.ErrRecordNotFound) {
		return data, errors.New("重复创建～")
	}

	//
	err := global.CURD_DB.Create(&req).Error
	return data, err
}

func Delete{{.ModelName}}(req *proto.{{.ModelName}}) (*proto.{{.ModelName}}, error) {
	err := global.CURD_DB.Where(&req).First(&req).Delete(&req).Error
	return req, err
}

func DeleteById{{.ModelName}}(req *proto.{{.ModelName}}) (*proto.{{.ModelName}}, error) {
	err := global.CURD_DB.First(&req).Delete(&req).Error
	return req, err
}

func Update{{.ModelName}}(req *proto.{{.ModelName}}) (*proto.{{.ModelName}}, error) {
	err := global.CURD_DB.Update(&req).Error
	return req, err
}

func Find{{.ModelName}}(req *proto.{{.ModelName}}) (*proto.{{.ModelName}}, error) {
	err := global.CURD_DB.Where(&req).First(&req).Error
	return req, err
}

func GetList{{.ModelName}}(req *proto.Request) (*proto.Responses, error) {
	meta := &proto.Meta{Total:0} // 需要进行初始化

	rsp := &proto.Responses{}
	rsp.Meta = meta

	if req.PageSize == 0 {
		req.PageSize = constant.PAGESIZE
	}
	if req.Page == 0 {
		req.Page = constant.PAGE
	}
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)

	unmarshal := &proto.{{.ModelName}}{}
	err := ptypes.UnmarshalAny(req.Query, unmarshal)

	db := global.CURD_DB.Model(&rsp.Data)

	if unmarshal.Path != "" {
		db = db.Where("path LIKE ?", "%"+unmarshal.Path+"%")
	}

	if unmarshal.Description != "" {
		db = db.Where("description LIKE ?", "%"+unmarshal.Description+"%")
	}

	if unmarshal.Method != "" {
		db = db.Where("method = ?", unmarshal.Method)
	}

	if unmarshal.ApiGroup != "" {
		db = db.Where("api_group = ?", unmarshal.ApiGroup)
	}

	err = db.Count(&rsp.Meta.Total).Error

	if err != nil {
		return rsp, err
	} else {
		db = db.Limit(limit).Offset(offset)
		if req.OrderKey != "" {
			var OrderStr string
			if req.OrderDesc != "" {
				OrderStr = req.OrderKey + " desc"
			} else {
				OrderStr = req.OrderKey
			}
			err = db.Order(OrderStr).Find(&rsp.Data).Error
		} else {
			err = db.Order("id").Find(&rsp.Data).Error
		}
	}

	return rsp, err
}


