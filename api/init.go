package api

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var started time.Time

// Init 初始化服务
func Init() {

	// conf.Init() // 配置的初始化
	// sql.Init()  // 数据库初始化

	// initLogger()

	// 其他需要初始化的sdk和internal包在这个位置完成

	// 路由初始化
	started = time.Now()
	r := gin.Default()

	r.GET("/health", health)

	initRouter(r)
	log.Fatal(r.Run(":9093"))
}

func health(ctx *gin.Context) {
	w := ctx.Writer
	duration := time.Now().Sub(started)
	if duration.Seconds() > 20 {
		log.Println("healthz 健康检查失败")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error: #{duration.Seconds()}")))
		return
	} else {
		log.Println("Healthz 健康检查成功")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	}

}
