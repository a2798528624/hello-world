package service

import (
	"giligili/model"
	"giligili/serializer"
)

type ListVideoService struct {
}

func (service *ListVideoService) List() serializer.Response {
	videos := []model.Video{}
	err := model.DB.Find(&videos).Error
	if err != nil {
		return serializer.DBErr("获取视频列表失败", err)
	}
	return serializer.Response{
		Code: 200,
		Data: serializer.BuildVideos(videos),
	}
}
