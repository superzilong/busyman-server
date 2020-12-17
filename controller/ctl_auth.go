package controller

import (
	"gg/models"
	"gg/pkg/jwt"
	"gg/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login is the callback of /login route.
func Login(c *gin.Context) {
	var user models.UserInfo
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 2001,
			"msg":  "无效的参数",
		})
		return
	}

	if service.GetUserInfoByName(user.Name).Password != user.Password {
		c.JSON(http.StatusOK, gin.H{
			"code": 2002,
			"msg":  "用户名或密码不对",
		})
		return
	}

	acessTokenString, err := jwt.GenAccessToken(user.Name)
	refreshTokenString, err2 := jwt.GenRefreshToken()
	if err != nil || err2 != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 2003,
			"msg":  "生成token出错",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 2000,
		"msg":  "success",
		"data": gin.H{
			"accessToken":  acessTokenString,
			"refreshToken": refreshTokenString},
	})
}

// RefreshToken will refresh the access token.
func RefreshToken(c *gin.Context) {
	m := make(map[string]string)
	if err := c.BindJSON(&m); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 2001,
			"msg":  "无效的参数",
		})
	}
	strRefreshToken, ok := m["refreshToken"]
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"code": 2001,
			"msg":  "参数中没有refreshToken",
		})
	}
	err := jwt.ParseRefreshToken(strRefreshToken)
	if err != nil {
		if jwt.IsTokenExpiredErr(err) {
			c.JSON(http.StatusOK, gin.H{
				"code": 2004,
				"msg":  "refreshToken expires.",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 2005,
				"msg":  "invalid refreshToken.",
			})
		}
	}
	username, _ := c.Get("username")
	strAccessToken, err := jwt.GenAccessToken(username.(string))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 2003,
			"msg":  "生成token出错",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 2000,
		"msg":  "success",
		"data": gin.H{
			"accessToken": strAccessToken,
		},
	})
}
