package v1

import (
	"context"
	"encoding/json"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type UserReq struct {
	Name     string
	Role     string
	Password string
}

type UserRes struct {
	ID       string
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

	id := uuid.NewString()

	rdb.Set(context.Background(), id, byteUser, 0)

	c.JSON(201, &UserRes{
		ID:       id,
		Name:     user.Name,
		Role:     user.Role,
		Password: user.Password,
	})

}

// @Summary Upload Photo
// @Description Api for upload a new photo
// @Tags media
// @Accept multipart/form-data
// @Param file formData file true "createUserModel"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /v1/media/ [post]
func (h handler) UploadMedia(c *gin.Context) {

	header, _ := c.FormFile("file")

	url := filepath.Join("media", header.Filename)

	err := c.SaveUploadedFile(header, url)
	if err != nil {
		return
	}

	c.JSON(200, gin.H{
		"message": url,
	})

}
