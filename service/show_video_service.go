package service

import (
	"giligili/model"
	"giligili/serializer"
)

type ShowVideoService struct {
}

func (service *ShowVideoService) Show(id string) serializer.Response {
	video := model.Video{}
	err := model.DB.First(&video, id).Error
	if err != nil {
		return serializer.DBErr("获取视频失败", err)
	}
	return serializer.Response{
		Code: 200,
		Data: serializer.BuildVideo(video),
	}

}
