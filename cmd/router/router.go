package router

import (
	"test_rudnytskyi/cmd/controllers"
	"test_rudnytskyi/cmd/middleware"

	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(catController *controllers.CatController, missionController *controllers.MissionController) *gin.Engine {
	router := gin.Default()
	router.Use(middleware.LoggerMiddleware())

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})
	baseRouter := router.Group("/api/v1")
	catRouter := baseRouter.Group("/cats")
	catRouter.GET("", catController.FindAll)
	catRouter.GET("/:catId", catController.FindById)
	catRouter.POST("", catController.Create)
	catRouter.PATCH("/:catId", catController.Update)
	catRouter.DELETE("/:catId", catController.Delete)

	missionRouter := baseRouter.Group("/mission")
	missionRouter.POST("", missionController.Create)
	missionRouter.GET("", missionController.FindAll)
	missionRouter.GET("/:missionId", missionController.FindById)
	missionRouter.POST("/find_missions", missionController.FindMissionByCatId)
	missionRouter.POST("/assign", missionController.AssignCatToMission)
	missionRouter.PATCH("/update_name", missionController.UpdateNameRequest)
	missionRouter.DELETE("/:missionId", missionController.Delete)
	missionRouter.PATCH("/complete_mission", missionController.CompleteMission)
	
	missionRouter.PATCH("/add_target", missionController.AddTarget)
	missionRouter.PATCH("/complete_target", missionController.CompleteTarget)
	missionRouter.DELETE("/remove_target", missionController.RemoveTarget)
	missionRouter.PATCH("/update_target", missionController.UpdateTarget)
	missionRouter.PATCH("/update_note", missionController.UpdateNotes)

	return router
}
