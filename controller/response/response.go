package response

import (
	"az-edu/consts"
	"az-edu/library/log"
	"github.com/gin-gonic/gin"
	"github.com/op/go-logging"
)

func Response(c *gin.Context, code int, msg string, data interface{}) {
	requestID := c.MustGet(consts.REQUEST_ID_KEY)
	c.JSON(200, map[string]interface{}{
		"data":                data,
		"error_no":            code,
		"error_msg":           msg,
		consts.REQUEST_ID_KEY: requestID,
	})
}

func ClientErr(c *gin.Context, msg string) {
	Response(c, 400, msg, nil)
}

func ServerErr(c *gin.Context, msg string) {
	Response(c, 500, msg, nil)
}

func Success(c *gin.Context) {
	Response(c, 0, "success", nil)
}

func ClientLogErr(c *gin.Context, logger *logging.Logger, msg string) {
	requestId := c.MustGet(consts.REQUEST_ID_KEY).(log.RequestID)
	logger.Error(requestId, msg)
	ClientErr(c, msg)
}

func ServerLogErr(c *gin.Context, logger *logging.Logger, msg string) {
	requestId := c.MustGet(consts.REQUEST_ID_KEY).(log.RequestID)
	logger.Error(requestId, msg)
	ServerErr(c, msg)
}

func ServerSucc(c *gin.Context, msg string, data interface{}) {
	requestId := c.MustGet(consts.REQUEST_ID_KEY).(log.RequestID)
	c.JSON(consts.HTTP_SUCC_CODE, map[string]interface{}{
		"data":       data,
		"error_no":   consts.ERR_NO_SUCC,
		"error_msg":  msg,
		"request_id": requestId,
	})
}
