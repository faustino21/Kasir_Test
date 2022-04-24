package api

import (
	"Kasir_Test/Delivery/commonResp"
	"Kasir_Test/Delivery/httpReq"
	"Kasir_Test/Delivery/middleware"
	"Kasir_Test/usecase"
	"Kasir_Test/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type LoginApi struct {
	commonResp.BaseApi
	login usecase.LoginUseCase
}

func (l *LoginApi) GetPassCode() gin.HandlerFunc {
	funcName := "LoginApi.GetPasscode"

	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("cashierId"))
		if err != nil {
			util.Log.Error().Msgf(funcName+".cashierID : %v", err)
			return
		}

		passcode, err := l.login.GetPasscode(id)
		if err != nil {
			util.Log.Error().Msgf(funcName+" : %v", err)
			commonResp.NewAppHttpResponse(c).FailedResp(http.StatusBadRequest, commonResp.NewFailedMessage(err.Error()))
			return
		}
		commonResp.NewAppHttpResponse(c).SuccessResp(http.StatusOK, commonResp.NewSuccessMessage(passcode))
	}
}

func (l *LoginApi) VerifyLoginPasscode() gin.HandlerFunc {
	funcName := "LoginApi.VerifyLogin"

	return func(c *gin.Context) {
		var cashier httpReq.CashierReq
		id, err := strconv.Atoi(c.Param("cashierId"))
		if err != nil {
			util.Log.Error().Msgf(funcName+".cashierID : %v", err)
			return
		}
		err = l.ParseRequestBody(c, &cashier)
		if err != nil {
			l.ParsingError(c, err)
			return
		}
		if len(cashier.Passcode) != 6 {
			commonResp.NewAppHttpResponse(c).FailedResp(http.StatusBadRequest, commonResp.NewFailedMessage("Password must be 6 characters"))
			return
		}
		data, err := l.login.Verify(id, cashier.Passcode)
		if err != nil {
			util.Log.Error().Msgf(funcName+" : %v", err)
			commonResp.NewAppHttpResponse(c).FailedResp(http.StatusBadRequest, commonResp.NewFailedMessage(err.Error()))
			return
		}
		token, err := middleware.GenerateToken(data.CashierName, data.CreatedAt)
		if err != nil {
			util.Log.Error().Msgf(funcName+" : %v", err)
			commonResp.NewAppHttpResponse(c).FailedResp(http.StatusBadRequest, commonResp.NewFailedMessage(err.Error()))
		}
		err = l.login.InsertingToken(id, token)
		if err != nil {
			util.Log.Error().Msgf(funcName+" : %v", err)
			commonResp.NewAppHttpResponse(c).FailedResp(http.StatusBadRequest, commonResp.NewFailedMessage(err.Error()))
		}
		commonResp.NewAppHttpResponse(c).SuccessResp(http.StatusOK, commonResp.NewSuccessMessage(gin.H{
			"token": token,
		}))
	}
}

func (l *LoginApi) VerifyLogout() gin.HandlerFunc {
	funcName := "LoginApi.VerifyLogout"

	return func(c *gin.Context) {
		var cashier httpReq.CashierReq
		id, err := strconv.Atoi(c.Param("cashierId"))
		if err != nil {
			util.Log.Error().Msgf(funcName+".cashierID : %v", err)
			return
		}
		err = l.ParseRequestBody(c, &cashier)
		if err != nil {
			l.ParsingError(c, err)
			return
		}
		if len(cashier.Passcode) != 6 {
			commonResp.NewAppHttpResponse(c).FailedResp(http.StatusBadRequest, commonResp.NewFailedMessage("Password must be 6 characters"))
			return
		}
		_, err = l.login.Verify(id, cashier.Passcode)
		if err != nil {
			util.Log.Error().Msgf(funcName+" : %v", err)
			commonResp.NewAppHttpResponse(c).FailedResp(http.StatusBadRequest, commonResp.NewFailedMessage(err.Error()))
			return
		}
		err = l.login.InsertingToken(id, "")
		if err != nil {
			util.Log.Error().Msgf(funcName+" : %v", err)
			commonResp.NewAppHttpResponse(c).FailedResp(http.StatusBadRequest, commonResp.NewFailedMessage(err.Error()))
		}
		commonResp.NewAppHttpResponse(c).SuccessResp2(http.StatusOK, commonResp.NewSuccessMessage2())
	}
}

func LoginApiRoute(route *gin.RouterGroup, login usecase.LoginUseCase) *LoginApi {
	loginApi := LoginApi{
		login: login,
	}

	route.GET("/:cashierId/passcode", loginApi.GetPassCode())
	route.POST("/:cashierId/login", loginApi.VerifyLoginPasscode())
	route.POST("/:cashierId/logout", loginApi.VerifyLogout())
	return &loginApi
}
