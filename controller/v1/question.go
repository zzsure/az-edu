package v1

import (
	"az-edu/controller/request"
	"az-edu/controller/response"
	"az-edu/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func QuestionAdd(c *gin.Context) {
	var req request.QuestionAdd
	if err := c.ShouldBindWith(&req, binding.JSON); err != nil {
		response.ClientLogErr(c, logger, err.Error())
		return
	}
	// TODO: transaction
	question, err := models.AddQuestion(req.Title, req.Content)
	if err != nil {
		response.ServerLogErr(c, logger, err.Error())
		return
	}
	for _, labelID := range req.LabelIDS {
		// TODO: check label id
		err := models.AddQuestionLabel(question.ID, labelID)
		if err != nil {
			response.ClientLogErr(c, logger, err.Error())
			return
		}
	}
	response.ServerSucc(c, "success", nil)
}
