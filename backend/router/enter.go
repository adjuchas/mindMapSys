package router

import (
	"backend/middleware"
	service "backend/service"
	stuService "backend/service/stu"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTION"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	router.Static("/uploads", "./uploads")

	v1 := router.Group("/api/v1")
	v1.POST("/info", service.GetInfo)
	v1.POST("/recommend", service.GetRecommends)
	v1.GET("/hasTags", service.GetHasTags)

	v1.POST("/drafts", middleware.Auth, service.GetDrafts)
	v1.POST("/classify", middleware.Auth, service.GetClassify)
	v1.POST("/draftSelect", middleware.Auth, service.DraftSelect)
	v1.POST("/draftCommit", middleware.Auth, service.DraftCommit)
	v1.POST("/draftDel", middleware.Auth, service.DraftDel)
	v1.POST("/createDraft", middleware.Auth, service.CreateDraft)
	v1.POST("/dots", middleware.Auth, service.GetDots)
	v1.POST("/setDotState", middleware.Auth, service.SetDotState)
	v1.POST("/dotSelect", middleware.Auth, service.DotSelect)
	v1.POST("/setMd", middleware.Auth, service.SetMd)
	v1.POST("/getMd", middleware.Auth, service.GetMd)
	v1.POST("/dotDel", middleware.Auth, service.DotDel)
	v1.POST("/getDotBody", middleware.Auth, service.GetDotBody)
	//下面这个是用来教师查看分类信息的
	v1.POST("/getClassifyInfo", middleware.Auth, service.GetClassifyInfo)
	v1.POST("/createClassify", middleware.Auth, service.CreateTag)
	v1.POST("/audit", middleware.Auth, service.Audit)
	v1.POST("/passAudit", middleware.Auth, service.PassAudit)
	v1.POST("/unPassAudit", middleware.Auth, service.UnPassAudit)
	v1.POST("/startDot", middleware.Auth, service.StarDot)
	v1.POST("/checkStar", middleware.Auth, service.CheckStar)
	v1.POST("/quitStar", middleware.Auth, service.QuitStar)

	v1.POST("/getStarDots", middleware.Auth, service.GetStarDots)
	v1.POST("/uploadImg", service.UploadImg)
	stu := v1.Group("/stu")
	{
		stu.POST("/getDraftBody", middleware.Auth, stuService.GetDraftBody)
		stu.POST("/setDraftBody", middleware.Auth, stuService.SetDraftBody)
		//未测试：
		//stu.POST("/history", middleware.Auth, stuService.GetHistory)
		stu.POST("/message", middleware.Auth, stuService.GetMsg)

		stu.POST("/setNote", middleware.Auth, stuService.SetNote)
		stu.POST("/getNote", middleware.Auth, stuService.GetNote)
		stu.POST("/notes", middleware.Auth, stuService.GetNoteList)

	}

	return router
}

func Start() {
	server := SetupRouter()
	server.Run(":8080")
}
