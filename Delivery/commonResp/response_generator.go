package commonResp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type AppHttpResponse interface {
	SuccessResp(httCode int, message *SuccessMessage)
	FailedResp(httpCode int, message *FailedMessage)
	SuccessResp2(httCode int, message *SuccessMessage2)
}

type JsonResponse struct {
	ctx *gin.Context
}

func (j *JsonResponse) SuccessResp(httCode int, message *SuccessMessage) {
	j.ctx.JSON(http.StatusOK, message)
	j.ctx.Abort()
}

func (j *JsonResponse) SuccessResp2(httCode int, message *SuccessMessage2) {
	j.ctx.JSON(http.StatusOK, message)
	j.ctx.Abort()
}

func (j *JsonResponse) FailedResp(httpCode int, message *FailedMessage) {
	j.ctx.JSON(http.StatusBadRequest, message)
	j.ctx.Abort()
}

func NewAppHttpResponse(ctx *gin.Context) AppHttpResponse {
	return &JsonResponse{
		ctx,
	}
}
