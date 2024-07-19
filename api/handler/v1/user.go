package v1

import (
	"context"
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type UserReq struct {
	Name     string
	Role     string
	Password string
}

// CreateUser ...
// @Summary CreateUser
// @Description Api for creating a new user
// @Tags user
// @Accept json
// @Produce json
// @Param User body UserReq true "createUserModel"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /v1/users/ [post]
func (h handler) CreatUser(c *gin.Context) {

	var user UserReq
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}

	rdb := redis.NewClient(
		&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		},
	)

	byteUser, err := json.Marshal(user)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
	}

	rdb.Set(context.Background(), user.Name, byteUser, 0)

}
