package middleware

import (
	"errors"
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

type casbinPermission struct {
	enforcer *casbin.Enforcer
}

func (casb *casbinPermission) GetRole(c *gin.Context) (string, int) {
	role := c.GetHeader("Authorization")

	if role == "" {
		return "unauthorized", http.StatusUnauthorized
	}

	return role, 0
}

func (casb *casbinPermission) CheckPermission(c *gin.Context) (bool, error) {

	act := c.Request.Method
	sub, status := casb.GetRole(c)
	if status != 0 {
		return false, errors.New("error in get role")
	}
	obj := c.Request.URL

	allow, err := casb.enforcer.Enforce(sub, obj.String(), act)
	if err != nil {
		return false, err
	}

	return allow, nil
}

func CheckPermissionMiddleware(enf *casbin.Enforcer) gin.HandlerFunc {
	casbHandler := &casbinPermission{
		enforcer: enf,
	}

	return func(c *gin.Context) {
		result, err := casbHandler.CheckPermission(c)

		if err != nil {
			c.AbortWithError(500, err)
		}
		if !result {
			c.AbortWithStatusJSON(401, gin.H{
				"message": "unauthorized",
			})
		}

		c.Next()
	}
}
