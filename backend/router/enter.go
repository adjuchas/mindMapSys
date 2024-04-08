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

	// 没测试

	//v1.POST("/startDot", middleware.Auth, service.StarDot)

	stu := v1.Group("/stu")
	{
		stu.POST("/getDraftBody", middleware.Auth, stuService.GetDraftBody)
		stu.POST("/setDraftBody", middleware.Auth, stuService.SetDraftBody)
		//未测试：
		stu.POST("/history", stuService.GetHistory)
		stu.POST("/message")

		stu.POST("/notes", middleware.Auth, stuService.GetNotes)
		stu.POST("/setNoteState", middleware.Auth, stuService.SetNoteState)
		stu.POST("/starNote", middleware.Auth, stuService.StarNote)
		stu.POST("/getStartNote", middleware.Auth, stuService.GetStartNotes)
		stu.POST("/getStarDot", middleware.Auth, stuService.GetStarDots)

		//stu.POST("/getDotBody", stuService.GetDotBody)
		//stu.POST("/getNoteBody", stuService.GetNoteBody)
	}

	return router
}

func Start() {
	server := SetupRouter()
	server.Run(":8080")
}
