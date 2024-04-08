package service

import (
	"backend/model/models"
	mysqlConn "backend/model/mysql"
	redisConn "backend/model/redis"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
)

type AllMe struct {
	Status string `json:"status"`
	Info   struct {
		YbId          string `json:"yb_userid"`
		YbUsername    string `json:"yb_username"`
		YbUsernick    string `json:"yb_usernick"`
		YbSex         string `json:"yb_sex"`
		YbBirthday    string `json:"yb_birthday"`
		YbMoney       string `json:"yb_money"`
		YbExp         string `json:"yb_exp"`
		YbUserhead    string `json:"yb_userhead"`
		YbRegtime     string `json:"yb_regtime"`
		YbSchoolid    string `json:"yb_schoolid"`
		YbSchoolname  string `json:"yb_schoolname"`
		InstituteId   string `json:"yb_collegeid"`
		YbCollegename string `json:"yb_collegename"`
		ClassId       string `json:"yb_classid"`
		YbClassname   string `json:"yb_classname"`
		Name          string `json:"yb_realname"`
		YbEnteryear   string `json:"yb_enteryear"`
		YbVerifykey   string `json:"yb_verifykey"`
		StuId         string `json:"yb_studentid"`
		YbExamid      string `json:"yb_examid"`
		YbAdmissionid string `json:"yb_admissionid"`
		YbEmployid    string `json:"yb_employid"`
		YbIdentity    string `json:"yb_identity"`
	} `json:"info"`
}

func GetInfo(c *gin.Context) {
	var resp models.RequestInfo
	var data map[string]interface{}
	c.ShouldBindJSON(&resp)
	info, err := redisConn.GetRedisUserInfo(resp.AccessToken)

	if err != nil {
		allme := GetYbInfo(resp.AccessToken)
		ClassifyIdentity(allme)
		userInfo := ToRedisUserInfo(allme)
		redisConn.RedisSetUserInfo(resp.AccessToken, userInfo)
		info, _ := redisConn.GetRedisUserInfo(resp.AccessToken)
		data = map[string]interface{}{
			"info": info,
		}
	} else {
		data = map[string]interface{}{
			"info": info,
		}
	}
	c.AsciiJSON(http.StatusOK, data)
}

func GetYbInfo(accessToken string) AllMe {
	url := fmt.Sprintf("https://openapi.yiban.cn/user/verify_me?access_token=%s", accessToken)
	method := "GET"
	payload := strings.NewReader(`<body data here>`)
	client := &http.Client{}
	req, _ := http.NewRequest(method, url, payload)
	res, _ := client.Do(req)

	body, _ := ioutil.ReadAll(res.Body)
	bodyStr := string(body)
	bodyStr = strings.ReplaceAll(bodyStr, "\n", "")

	var allme AllMe
	json.Unmarshal([]byte(bodyStr), &allme)
	return allme
}

func ClassifyIdentity(allme AllMe) {
	if allme.Info.YbIdentity != "学生" {
		tea := models.TeacherInfo{
			Teacher_ID: allme.Info.YbEmployid,
			Yb_ID:      allme.Info.YbId,
			Name:       allme.Info.Name,
		}
		if !mysqlConn.FindTea(tea.Teacher_ID) {
			mysqlConn.SetTea(tea)
		}
	} else {
		stu := models.Stu{
			YbId:        allme.Info.YbId,
			StuId:       allme.Info.StuId,
			Name:        allme.Info.Name,
			InstituteId: allme.Info.InstituteId,
			ClassId:     allme.Info.ClassId,
		}
		if !mysqlConn.FindStu(stu.StuId) {
			mysqlConn.SetStu(stu)
		}
	}
}

func ToRedisUserInfo(me AllMe) models.RedisUserInfo {
	var userInfo models.RedisUserInfo
	userInfo = models.RedisUserInfo{
		Yb_userid:      me.Info.YbId,
		Yb_collegeid:   me.Info.InstituteId,
		Yb_collegename: me.Info.YbCollegename,
		Yb_classid:     me.Info.ClassId,
		Yb_classname:   me.Info.YbClassname,
		Yb_realname:    me.Info.Name,
		Yb_studentid:   me.Info.StuId,
		Yb_employid:    me.Info.YbEmployid,
		Yb_identity:    me.Info.YbIdentity,
	}
	return userInfo
}
