package serializer

import "giligili/model"

type Video struct {
	Url       string `json:"url"`
	Title     string `json:"title"`
	Info      string `json:"info"`
	CreatedAt int64  `json:"create_at"`
}

//序列化单个视屏
func BuildVideo(videos model.Video) Video {
	return Video{
		Url:       videos.Url,
		Title:     videos.Title,
		Info:      videos.Info,
		CreatedAt: videos.CreatedAt.Unix(),
	}
}

//序列化视屏列表
func BuildVideos(videos []model.Video) []Video {
	var videoList []Video
	for _, v := range videos {
		videoList = append(videoList, BuildVideo(v))
	}
	return videoList
}
