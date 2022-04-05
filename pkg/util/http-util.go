package util

import (
	"context"
	"github.com/alexliesenfeld/health"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/qkgo/scaff/pkg/cfg"
	"github.com/qkgo/scaff/pkg/log"
	"github.com/qkgo/scaff/pkg/serialize"
	"github.com/qkgo/scaff/pkg/util/crypt"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func DefaultView(ctx *gin.Context) {
	ctx.JSON(200, map[string]interface{}{
		"success": true,
		"code":    200,
		"uri":     ctx.Request.RequestURI,
		"addr":    ctx.Request.RemoteAddr,
		"host":    ctx.Request.Host,
	})
}

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

// HealthCheckDatabase
// @Summary HealthCheckDatabase
// @Description HealthCheckDatabase
// @Produce  application/json
// @Router /health [GET]
func HealthCheckDatabase(ctx *gin.Context) {
	var checkerResult health.CheckerResult
	var checkerResultSecond health.CheckerResult
	if serialize.DB != nil {
		checker := health.NewChecker(

			// Set the time-to-live for our cache to 1 second (default).
			health.WithCacheDuration(1*time.Second),

			// Configure a global timeout that will be applied to all checks.
			health.WithTimeout(10*time.Second),

			// A check configuration to see if our database connection is up.
			// The check function will be executed for each HTTP request.
			health.WithCheck(health.Check{
				Name:    "database",      // A unique check name.
				Timeout: 2 * time.Second, // A check specific timeout.
				Check:   serialize.DB.DB().PingContext,
			}),
		)
		checkerResult = checker.Check(ctx)
	}
	if serialize.SecondDB != nil {
		checkerSecond := health.NewChecker(

			// Set the time-to-live for our cache to 1 second (default).
			health.WithCacheDuration(1*time.Second),

			// Configure a global timeout that will be applied to all checks.
			health.WithTimeout(10*time.Second),

			// A check configuration to see if our database connection is up.
			// The check function will be executed for each HTTP request.
			health.WithCheck(health.Check{
				Name:    "database",      // A unique check name.
				Timeout: 2 * time.Second, // A check specific timeout.
				Check:   serialize.SecondDB.DB().PingContext,
			}),
		)
		checkerResultSecond = checkerSecond.Check(ctx)
	}

	if checkerResult.Status == health.StatusDown || checkerResultSecond.Status == health.StatusDown {
		ctx.JSON(400, map[string]interface{}{
			"db-main":   checkerResult,
			"db-second": checkerResultSecond,
		})
		return
	}
	ctx.JSON(200, map[string]interface{}{
		"db-main":   checkerResult,
		"db-second": checkerResultSecond,
	})
	return
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
	router.GET("/", DefaultView)
	router.GET("/health", HealthCheckDatabase)
	router.GET("/health/database", HealthCheckDatabase)
	router.GET("/hc", HealthCheckDatabase)
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

	log.I("start listen:", addr)
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
	log.I("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second/10)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	select {
	case <-ctx.Done():
		log.I("timeout of 0.1 seconds.")
	}
	log.I("Server exiting")
	return srv
}
