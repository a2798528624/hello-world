package service

import (
	"giligili/model"
	"giligili/serializer"
)

type CreatVideoService struct {
	Title string `json:"title" form:"title" binding:"required,min=2,max=30"`
	Url   string `json:"url" form:"url"`
	Info  string `json:"info" form:"info" binding:"required,min=0,max=300"`
}

func (service *CreatVideoService) Create() serializer.Response {
	video := model.Video{
		Title: service.Title,
		Url:   service.Url,
		Info:  service.Info,
	}
	err := model.DB.Create(&video).Error //这里必须传递地址，否则会报错
	if err != nil {
		return serializer.DBErr("视频存储失败", err)
	}
	return serializer.Response{
		Code: 200,
		Data: serializer.BuildVideo(video),
	}
}
