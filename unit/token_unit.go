package unit

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"todoProject/config"
)

func GetUserId(c *gin.Context) (int64, error) {
	var userId int64
	var error error
	temp, isOk := c.Get(config.TOKEN_CURRENT_USER_ID)
	if isOk {
		userId, error = strconv.ParseInt(fmt.Sprint(temp), 10, 64)
		if error == nil {
			return userId, nil
		}
	}
	return 0, nil
}
