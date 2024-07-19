package api

import (
	v1 "lessons/casbin/api/handler/v1"
	middlewareCasbin "lessons/casbin/api/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

type Option struct {
	Enforcer *casbin.Enforcer
}

func New(option Option) *gin.Engine {
	router := gin.New()

	router.Use(middlewareCasbin.CheckPermissionMiddleware(option.Enforcer))

	handlerV1 := v1.New(&v1.HandlerConfig{
		Enforcer: option.Enforcer,
	})

	router.POST("/users", handlerV1.CreatUser)

	return router
}
