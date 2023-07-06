package basetable

import (
	"basetable.com/internal/basetable/controller/v1/user"
	"basetable.com/internal/basetable/store"
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

	userController := user.New(store.S)
	g.POST("/auth/login", userController.Login)

	v1 := g.Group("/v1")
	v1.Use(mw.Auth())
	{
		userV1 := v1.Group("/users")
		{
			userV1.POST("", userController.Create)
			userV1.GET("/:id", userController.GetOne)
			//userV1.GET("", userController.List)
			userV1.GET("/list", userController.List)
			userV1.DELETE("/:id", userController.Deleted)
		}
	}

	return nil
}
