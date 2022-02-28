package util

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/qkgo/scaff/pkg/cfg"
	"github.com/qkgo/scaff/pkg/util/crypt"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func NoRouterHandle() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(404, map[string]interface{}{
			"error":   "Not router",
			"success": false,
			"code":    -404.1,
		})
	}
}

func NoMethodHandle() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(404, map[string]interface{}{
			"error":   "Not method",
			"success": false,
			"code":    -404.2,
		})
	}
}

func NoRouterHandler(context *gin.Context) {
	context.JSON(404, map[string]interface{}{
		"error":   "Not router",
		"success": false,
		"code":    -404.1,
	})
	return
}

func NoMethodHandler(context *gin.Context) {
	context.JSON(404, map[string]interface{}{
		"error":   "Not method",
		"success": false,
		"code":    -404.2,
	})
	return
}

func VersionHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Header("server-version", "2.1.6")
		return
	}
}

func GetRouter(
	needCrypto bool,
	ConfigCustomRouter func(*gin.Engine) *gin.Engine,
	apiRouterConfig func(*gin.Engine) *gin.Engine,
) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard
	router := gin.Default()
	corsDefault := cors.DefaultConfig()
	corsDefault.AllowCredentials = true
	corsDefault.AllowHeaders = []string{"Origin", "token", "Content-Length", "Content-Type", "session", "DNT", "content-type", "s", "timezone", "tz", "specify", "order"}
	corsDefault.AllowAllOrigins = true
	router.Use(GinToLogrus())
	router.Use(cors.New(corsDefault))
	c := gin.LoggerConfig{
		Output:    ioutil.Discard,
		SkipPaths: []string{"*"},
	}
	router.Use(gin.LoggerWithConfig(c))
	router.Use(gin.Recovery())
	router.GET("/doc/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	if ConfigCustomRouter != nil {
		ConfigCustomRouter(router)
	}
	if needCrypto {
		if cfg.OzConfig.IsProduction() {
			router.Use(crypt.QuietCrypto())
		} else {
			router.Use(crypt.Crypto())
		}
	}
	router.Use(crypt.TokenRole())
	router.Use(VersionHandler())
	router.NoRoute(NoRouterHandle())
	router.NoMethod(NoMethodHandle())
	if apiRouterConfig != nil {
		apiRouterConfig(router)
	}
	return router
}

func BootstrapHttp(
	needCrypto bool,
	projectName string,
	customRouterFunc func(*gin.Engine) *gin.Engine,
	apiRouterFunc func(*gin.Engine) *gin.Engine,
) *http.Server {
	router := GetRouter(needCrypto, customRouterFunc, apiRouterFunc)
	var addr string
	if os.Getenv("PORT") != "" {
		addr = ":" + os.Getenv("PORT")
	} else if cfg.OzConfig.GetServerBindPort(projectName) != "" {
		addr = cfg.OzConfig.GetServerBindPort(projectName)
	} else {
		addr = ":8896"
	}

	cfg.LogInfo.Info("start listen:", addr)
	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	cfg.LogInfo.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second/10)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	select {
	case <-ctx.Done():
		cfg.LogInfo.Info("timeout of 0.1 seconds.")
	}
	cfg.LogInfo.Info("Server exiting")
	return srv
}
