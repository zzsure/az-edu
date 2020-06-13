package v1

import (
	"az-edu/controller/request"
	"az-edu/controller/response"
	"az-edu/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func LabelAdd(c *gin.Context) {
	var req request.LabelAdd
	if err := c.ShouldBindWith(&req, binding.JSON); err != nil {
		response.ClientLogErr(c, logger, err.Error())
		return
	}
	err := models.AddLabel(req.Name)
	if err != nil {
		response.ServerLogErr(c, logger, err.Error())
		return
	}
	response.ServerSucc(c, "success", nil)
}