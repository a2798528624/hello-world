package api

import (
	"giligili/service"

	"github.com/gin-gonic/gin"
)

//视频投稿接口
func CreateVideo(c *gin.Context) {
	var service service.CreatVideoService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, ErrorResponse(err))
		return
	}

	res := service.Create()
	c.JSON(200, res)
}

//获取单个视频接口
func ShowVideo(c *gin.Context) {
	var service service.ShowVideoService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, ErrorResponse(err))
		return
	}

	res := service.Show(c.Param("id"))
	c.JSON(200, res)
}

//获取全部视频接口
func ListVideo(c *gin.Context) {
	var service service.ListVideoService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, ErrorResponse(err))
		return
	}

	res := service.List()
	c.JSON(200, res)
}

//更新单个视频接口
func UpdateVideo(c *gin.Context) {
	var service service.UpdateVideoService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, ErrorResponse(err))
		return
	}

	res := service.Update(c.Param("id"))
	c.JSON(200, res)
}

//删除单个视频接口
func DeleteVideo(c *gin.Context) {
	var service service.DeleteVideoService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, ErrorResponse(err))
		return
	}

	res := service.Delete(c.Param("id"))
	c.JSON(200, res)
}
