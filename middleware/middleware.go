package middleware

import (
	"net/http"
	"sesi-2/helper"
	"sesi-2/model"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(ctx *gin.Context) {
	auth := ctx.GetHeader("Authorization")

	token := strings.Split(auth, " ")[1]

	jwtToken, err := helper.VerifyToken(token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.MyError{
			Err: err.Error(),
		})
		return
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.MyError{
			Err: err.Error(),
		})
		return
	}

	ctx.Set("user_id", claims["user_id"])

	ctx.Next()
}
