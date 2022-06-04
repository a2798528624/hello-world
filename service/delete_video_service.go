package service

import (
	"giligili/model"
	"giligili/serializer"
)

type DeleteVideoService struct {
	ID string `form:"id" json:"id" binding:"required"`
}

func (service *DeleteVideoService) Delete(id string) serializer.Response {
	video := model.Video{}
	if err := model.DB.First(&video, id).Error; err != nil {
		return serializer.ParamErr("视频不存在", err)
	}
	if err := model.DB.Delete(&video).Error; err != nil {
		return serializer.DBErr("删除视频失败", err)
	}
	return serializer.Response{
		Code: 200,
		Msg:  "删除视频成功",
	}
}
