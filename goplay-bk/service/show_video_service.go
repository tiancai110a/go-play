package service

import (
	"github.com/tiancai110a/go-play/goplay-bk/model"
	"github.com/tiancai110a/go-play/goplay-bk/serializer"
)

// ShowVideoService 投稿详情的服务
type ShowVideoService struct {
}

// Show 创建视频
func (service *ShowVideoService) Show(id string) serializer.Response {
	var video model.Video
	err := model.DB.First(&video, id).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "视频不存在",
			Error:  err.Error(),
		}
	}
	//处理视频被观看的一系问题
	video.AddView()
	return serializer.Response{
		Data: serializer.BuildVideo(video),
	}
}
