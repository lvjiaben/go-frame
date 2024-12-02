package bootstrap

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	zapLog "go.uber.org/zap"

	"github.com/gin-gonic/gin"
	"github.com/lvjiaben/go-wheel/init/mysql"
	"github.com/lvjiaben/go-wheel/init/redis"
	"github.com/lvjiaben/go-wheel/init/validate"
	"github.com/lvjiaben/go-wheel/init/viper"
	"github.com/lvjiaben/go-wheel/init/zap"
	"github.com/lvjiaben/go-wheel/routes"
)

func init() {
	viper.Load()
	logger := zap.Load()
	defer logger.Sync()
	logger.Debug("Logger init success")
	mysql.Load()
	defer mysql.Close()
	redis.Load()
	defer redis.Close()
	validate.Load()
	// GIN启动
	gin.SetMode(viper.Conf.App.Mode)
	r := gin.New()
	r.Use(zap.GinLogger(), zap.GinRecovery(true))
	// GIN注册
	routes.RegisterRoutes(r)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.Conf.App.Port),
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	// 热退出
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("Server Shutdown", zapLog.Error(err))
	}
	logger.Info("Server exiting")
}
