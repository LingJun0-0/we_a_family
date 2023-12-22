package service

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
	"we_a_family/we_a_family/global"
	Models "we_a_family/we_a_family/models"
	"we_a_family/we_a_family/utils"
)

func Auth(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	payload, err := ParseToken(token)
	if err != nil {
		utils.FailwithMessage(err.Error(), ctx)
		return
	}
	memberId := payload.UserID
	ctx.Set("memberId", memberId)
}

// GenToken 生成 JWT 令牌
func GenToken(userID int) (string, error) {
	// 设置 JWT 负载
	payload := Models.AuthPayload{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 有效期为 24 小时
			IssuedAt:  time.Now().Unix(),
			Subject:   "auth",
		},
	}

	// 生成令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString([]byte(global.Config.JwtSecret.Auth))
}

func ParseToken(tokenString string) (*Models.AuthPayload, error) {
	var mc = new(Models.AuthPayload)
	token, err := jwt.ParseWithClaims(tokenString, mc, keyFunc)
	if err != nil {
		global.Log.Println(err.Error())
		return nil, err
	}
	if token.Valid {
		return mc, nil
	}
	return nil, errors.New("invalid token")
}

func keyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(global.Config.JwtSecret.Auth), nil
}
