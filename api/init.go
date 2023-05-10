package api

import (
	"log"

	"github.com/gin-gonic/gin"
)

// Init 初始化服务
func Init() {

	// conf.Init() // 配置的初始化
	// sql.Init()  // 数据库初始化

	// initLogger()

	// 其他需要初始化的sdk和internal包在这个位置完成

	// 路由初始化
	r := gin.Default()
	initRouter(r)
	log.Fatal(r.Run(":8081"))
}
