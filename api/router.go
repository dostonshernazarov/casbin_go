package api

import (
	v1 "lessons/casbin/api/handler/v1"
	middlewareCasbin "lessons/casbin/api/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "lessons/casbin/api/docs"
)

type Option struct {
	Enforcer *casbin.Enforcer
}

// NewRouter
// @title Welcome To CV Maker API
// @Description API for CV Maker
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func New(option Option) *gin.Engine {
	router := gin.New()

	router.Use(middlewareCasbin.CheckPermissionMiddleware(option.Enforcer))

	handlerV1 := v1.New(&v1.HandlerConfig{
		Enforcer: option.Enforcer,
	})

	api := router.Group("/v1")

	api.POST("/users", handlerV1.CreatUser)
	api.POST("/media", handlerV1.UploadMedia)

	url := ginSwagger.URL("swagger/doc.json")
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}
