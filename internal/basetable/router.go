package basetable

import (
	"basetable.com/internal/basetable/api/v1"
	"basetable.com/internal/pkg/core"
	"basetable.com/internal/pkg/errno"
	"basetable.com/internal/pkg/log"
	mw "basetable.com/internal/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func initRouter(g *gin.Engine) error {
	// 注册 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		core.WriteResponse(c, errno.ErrPageNotFound, nil)
	})

	// 注册 /healthz handler.
	g.GET("/healthz", func(c *gin.Context) {
		log.C(c).Infow("Healthz function calle")
		core.WriteResponse(c, nil, map[string]string{"status": "ok"})
	})

	userController := v1.ApiGroupApp.UserApi
	g.POST("/auth/login", userController.Login)

	v1 := g.Group("/v1")
	v1.Use(mw.Auth())
	{
		userV1 := v1.Group("/users")
		{
			userV1.POST("", userController.Create)
			userV1.GET("/:id", userController.GetById)
			userV1.PUT("", userController.Update)
			userV1.DELETE("", userController.Deleted)
			userV1.GET("/list", userController.List)
		}
	}

	return nil
}
