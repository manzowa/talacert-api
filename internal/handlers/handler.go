package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func parseID(cxt *gin.Context) (uint, error) {

	id, err := strconv.ParseUint(
		cxt.Param("id"),
		10,
		64,
	)

	return uint(id), err
}
