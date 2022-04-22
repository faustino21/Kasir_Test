package commonResp

import (
	"Kasir_Test/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseApi struct {
}

func (b *BaseApi) ParseRequestBody(c *gin.Context, body interface{}) error {
	err := c.ShouldBindJSON(body)
	if err != nil {
		return err
	}
	return nil
}

func (b *BaseApi) ParsingError(c *gin.Context, err error) {
	util.Log.Error().Msg("Parsing Error")
	NewAppHttpResponse(c).FailedResp(http.StatusBadRequest, NewFailedMessage(err.Error()))
}
