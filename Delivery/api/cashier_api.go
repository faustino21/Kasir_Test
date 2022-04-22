package api

import (
	"Kasir_Test/Delivery/commonResp"
	"Kasir_Test/Delivery/httpReq"
	"Kasir_Test/Delivery/httpResp"
	"Kasir_Test/usecase"
	"Kasir_Test/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CashierApi struct {
	commonResp.BaseApi
	cashier usecase.CashierUseCase
}

func (cs *CashierApi) GetCashierList() gin.HandlerFunc {
	funcName := "CashierApi.GetCashier"

	return func(c *gin.Context) {
		limit, err := strconv.Atoi(c.Query("limit"))
		if err != nil {
			util.Log.Error().Msgf(funcName+".limit : %w", err)
			return
		}
		skip, err := strconv.Atoi(c.Query("skip"))
		if err != nil {
			util.Log.Error().Msgf(funcName+".skip : %w", err)
			return
		}

		total, res, err := cs.cashier.GetCashier(skip, limit)
		if err != nil {
			util.Log.Error().Msgf(funcName+".getCashier : %w", err)
			commonResp.NewAppHttpResponse(c).FailedResp(http.StatusBadRequest, commonResp.NewFailedMessage(err.Error()))
			return
		}

		metaResp := httpResp.NewMetaResp(total, limit, skip)
		commonResp.NewAppHttpResponse(c).SuccessResp(http.StatusOK, commonResp.NewSuccessMessage(gin.H{
			"cashiers": res,
			"meta":     metaResp,
		}))
	}
}

func (cs *CashierApi) GetCashierDetail() gin.HandlerFunc {
	funcName := "CashierApi.GetCashier"

	return func(c *gin.Context) {
		cashierId, err := strconv.Atoi(c.Param("cashierId"))
		if err != nil {
			util.Log.Error().Msgf(funcName+" : %w", err)
			return
		}
		cashier, err := cs.cashier.GetCashierDetail(cashierId)
		if err != nil {
			util.Log.Error().Msgf(funcName+" : %w", err)
			commonResp.NewAppHttpResponse(c).FailedResp(http.StatusBadRequest, commonResp.NewFailedMessage(err.Error()))
			return
		}
		commonResp.NewAppHttpResponse(c).SuccessResp(http.StatusOK, commonResp.NewSuccessMessage(cashier))
	}
}

func (cs *CashierApi) RegisterCashier() gin.HandlerFunc {
	funcName := "CashierApi.RegisterCashier"

	return func(c *gin.Context) {
		var cashierReq httpReq.CashierReq
		err := cs.ParseRequestBody(c, &cashierReq)
		if err != nil {
			cs.ParsingError(c, err)
			return
		}

		data, err := cs.cashier.RegisterCashier(cashierReq.Name, cashierReq.Passcode)
		if err != nil {
			util.Log.Error().Msgf(funcName+" : %w", err)
			commonResp.NewAppHttpResponse(c).FailedResp(http.StatusBadRequest, commonResp.NewFailedMessage(err.Error()))
			return
		}
		commonResp.NewAppHttpResponse(c).SuccessResp(http.StatusOK, commonResp.NewSuccessMessage(data))
	}
}

func (cs *CashierApi) UpdateCashier() gin.HandlerFunc {
	funcName := "CashierApi.UpdateCashier"

	return func(c *gin.Context) {
		var updateCashier httpReq.CashierReq
		id, err := strconv.Atoi(c.Param("cashierId"))
		if err != nil {
			util.Log.Error().Msgf(funcName+".cashierID : %w", err)
			return
		}
		err = cs.ParseRequestBody(c, &updateCashier)
		if err != nil {
			cs.ParsingError(c, err)
			return
		}

		err = cs.cashier.UpdateCashier(id, updateCashier.Name, updateCashier.Passcode)
		if err != nil {
			util.Log.Error().Msgf(funcName+".usCaseUpdate : %w", err)
			commonResp.NewAppHttpResponse(c).FailedResp(http.StatusBadRequest, commonResp.NewFailedMessage(err.Error()))
			return
		}
		commonResp.NewAppHttpResponse(c).SuccessResp2(http.StatusOK, commonResp.NewSuccessMessage2())
	}
}

func (cs *CashierApi) DeleteCashier() gin.HandlerFunc {
	funcName := "CashierApi.DeleteCashier"

	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("cashierId"))
		if err != nil {
			util.Log.Error().Msgf(funcName+".cashierID : %w", err)
			return
		}

		err = cs.cashier.DeleteCashier(id)
		if err != nil {
			util.Log.Error().Msgf(funcName+".usCaseDelete : %w", err)
			commonResp.NewAppHttpResponse(c).FailedResp(http.StatusBadRequest, commonResp.NewFailedMessage(err.Error()))
			return
		}
		commonResp.NewAppHttpResponse(c).SuccessResp2(http.StatusOK, commonResp.NewSuccessMessage2())
	}
}

func CashierApiRoute(route *gin.RouterGroup, cashier usecase.CashierUseCase) *CashierApi {
	cashierApi := CashierApi{
		cashier: cashier,
	}

	route.GET("", cashierApi.GetCashierList())
	route.GET("/:cashierId", cashierApi.GetCashierDetail())
	route.POST("", cashierApi.RegisterCashier())
	route.PUT("/:cashierId", cashierApi.UpdateCashier())
	route.DELETE(":cashierId", cashierApi.DeleteCashier())
	return &cashierApi
}
