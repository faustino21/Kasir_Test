package api

import (
	"Kasir_Test/Delivery/commonResp"
	"Kasir_Test/Delivery/httpResp"
	"Kasir_Test/usecase"
	"Kasir_Test/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ProductApi struct {
	commonResp.BaseApi
	product usecase.ProductUseCase
}

func (p *ProductApi) GetAllProduct() gin.HandlerFunc {
	funcName := "ProductApi.GetAllProduct"

	return func(c *gin.Context) {
		limit, err := strconv.Atoi(c.Query("limit"))
		if err != nil {
			util.Log.Error().Msgf(funcName+".limit : %v", err)
			return
		}
		skip, err := strconv.Atoi(c.Query("skip"))
		if err != nil {
			util.Log.Error().Msgf(funcName+".skip : %v", err)
			return
		}
		categoryId, err := strconv.Atoi(c.Query("categoryId"))
		if err != nil {
			util.Log.Error().Msgf(funcName+".categoryId : %v", err)
			return
		}
		q := c.Query("q")

		total, res, err := p.product.GetAllProduct(limit, skip, categoryId, q)
		if err != nil {
			util.Log.Error().Msgf(funcName+".getProduct : %v", err)
			commonResp.NewAppHttpResponse(c).FailedResp(http.StatusBadRequest, commonResp.NewFailedMessage(err.Error()))
			return
		}

		metaResp := httpResp.NewMetaResp(total, limit, skip)
		commonResp.NewAppHttpResponse(c).SuccessResp(http.StatusOK, commonResp.NewSuccessMessage(gin.H{
			"product": res,
			"meta":    metaResp,
		}))
	}
}

func ProductApiRoute(route *gin.RouterGroup, product usecase.ProductUseCase) *ProductApi {
	productApi := ProductApi{
		product: product,
	}

	route.GET("", productApi.GetAllProduct())

	return &productApi
}
