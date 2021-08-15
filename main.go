package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-zelus/api-layout/docs"
	"github.com/go-zelus/api-layout/router"
	"github.com/go-zelus/zelus/core/config"
	"github.com/go-zelus/zelus/core/logx"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title API文档
// @version 1.0
// @description common project.
// @termsOfService https://archer.plus

// @contact.name Archer
// @contact.url https://archer.plus
// @contact.email archer.plus@qq.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8080
// @BasePath /api/v1
func main() {
	// 初始化配置文件
	config.Load("./config/config.yml")
	// 初始化日志
	logx.Init()

	r := router.New()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if !config.GetBool("app.debug") {
		gin.SetMode(gin.ReleaseMode)
	}

	addr := ":" + config.GetString("app.port")
	server := &http.Server{
		Addr:           addr,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go server.ListenAndServe()
	fmt.Println("server listen port", addr)
	listenSignal(server)
}

func listenSignal(server *http.Server) {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGQUIT)
	<-ch
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	server.Shutdown(ctx)
}
