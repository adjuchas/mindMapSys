package middleware

import (
	redisConn "backend/model/redis"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthParams struct {
	Id          string                 `json:"Id"`
	Identity    string                 `json:"Identity"`
	AccessToken string                 `json:"accessToken"`
	Data        map[string]string      `json:"Data"`
	NodeJson    map[string]interface{} `json:"NodeJson"`
}

func Auth(c *gin.Context) {
	var authParams AuthParams
	var auth bool
	c.ShouldBindJSON(&authParams)
	c.Set("Identity", authParams.Identity)
	if authParams.Identity == "学生" {
		auth = redisConn.CheckUserByStuid(authParams.AccessToken, authParams.Id)
	} else {
		auth = redisConn.CheckTeaUserByStuid(authParams.AccessToken, authParams.Id)
	}
	if !auth {
		c.Abort()
		data := map[string]interface{}{
			"result": false,
		}
		c.AsciiJSON(http.StatusOK, data)
	}
	c.Set("data", authParams.Data)
	c.Set("nodeJson", authParams.NodeJson)
	c.Next()
}
