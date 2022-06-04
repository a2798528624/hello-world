package conf

import (
	"giligili/cache"
	"giligili/model"
	"giligili/util"
	"log"
	"os"

	"github.com/joho/godotenv" //从本地读取环境变量
)

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// 设置日志级别
	util.BuildLogger(os.Getenv("LOG_LEVEL"))

	// // 读取翻译文件
	// if err := LoadLocales("conf/locales/zh-cn.yaml"); err != nil {
	// 	util.Log().Panic("翻译文件加载失败", err)
	// }

	// 连接数据库
	model.Database(os.Getenv("MYSQL_DSN"))
	cache.Redis()
}
