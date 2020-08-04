package work

import (
	"fmt"
	"github.com/Carrotxyy/syncSense/common/api"
	"github.com/Carrotxyy/syncSense/common/setting"
	"github.com/Carrotxyy/syncSense/repository"
)

type Work struct {
	Repository *repository.BaseRepository `inject:""`
	HttpApi    *api.Api                   `inject:""`
	Config     *setting.Config            `inject:""`
}

/**
添加路由参数

@path 访问路由
*/
func (w *Work) SplicUrl(path string) string {

	//// 获取 key
	//key := api.GetKey(w.Config.WxAddr +"/interactive/key").(string)
	//// 加密 key
	//enkey := api.Encryption(key)
	//
	// 构建url
	url := fmt.Sprintf("%s%s", w.Config.WxAddr, path)
	return url
}
