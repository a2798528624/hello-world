package service

import (
	"giligili/model"
	"giligili/serializer"
)

type UpdateVideoService struct {
	Title string `json:"title" form:"title" binding:"required,min=2,max=30"`
	Url   string `json:"url" form:"url"`
	Info  string `json:"info" form:"info" binding:"required,min=0,max=300"`
}

func (service *UpdateVideoService) Update(id string) serializer.Response {
	video := model.Video{}
	err := model.DB.First(&video, id).Error
	if err != nil {
		return serializer.DBErr("视频不存在", err)
	}
	video.Title = service.Title
	video.Url = service.Url
	video.Info = service.Info
	err = model.DB.Save(&video).Error
	if err != nil {
		return serializer.DBErr("视频更新失败", err)
	}
	return serializer.Response{
		Code: 200,
		Data: serializer.BuildVideo(video),
	}
}
