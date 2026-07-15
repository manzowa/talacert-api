package auth

import "github.com/gin-gonic/gin"

const (
	ContextUser = "user"
)

func GetUserID(c *gin.Context) uint {

	userID, _ := c.Get("userID")

	return userID.(uint)
}

func GetUsername(c *gin.Context) string {

	username, _ := c.Get("username")

	return username.(string)
}

func GetEmail(c *gin.Context) string {

	email, _ := c.Get("email")

	return email.(string)
}

func GetRole(c *gin.Context) string {

	role, _ := c.Get("role")

	return role.(string)
}
