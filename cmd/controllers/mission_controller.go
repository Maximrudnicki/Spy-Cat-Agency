package controllers

import (
	"log"
	"net/http"
	"strconv"
	"test_rudnytskyi/cmd/data/request"
	"test_rudnytskyi/cmd/data/response"
	"test_rudnytskyi/cmd/services"

	"github.com/gin-gonic/gin"
)

type MissionController struct {
	missionService services.MissionService
}

func NewMissionController(service services.MissionService) *MissionController {
	return &MissionController{missionService: service}
}

func (controller *MissionController) AddTarget(ctx *gin.Context) {
	req := request.AddTargetRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "StatusBadRequest",
			Message: "Cannot parse data",
		}
		log.Printf("err: %v", err)
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	err = controller.missionService.AddTarget(req)
	if err != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Cannot add target",
		}
		log.Printf("err: %v", err)
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	webResponse := response.Response{
		Code:    http.StatusOK,
		Status:  "Ok",
		Message: "Successfully added target!",
		Data:    nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *MissionController) FindAll(ctx *gin.Context) {
	missionsResponse, err := controller.missionService.FindAll()
	if err != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Cannot find missions",
		}
		log.Printf("err: %v", err)
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	webResponse := response.Response{
		Code:    http.StatusOK,
		Status:  "Ok",
		Message: "Successfully found!",
		Data:    missionsResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *MissionController) FindById(ctx *gin.Context) {
	missionId := ctx.Param("missionId")
	id, err := strconv.Atoi(missionId)
	if err != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "StatusBadRequest",
			Message: "Cannot parse id",
		}
		log.Printf("err: %v", err)
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	missionResponse, err := controller.missionService.FindById(id)
	if err != nil || missionResponse.Id == 0 {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Cannot find mission by id",
		}
		log.Printf("err: %v", err)
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	webResponse := response.Response{
		Code:    http.StatusOK,
		Status:  "Ok",
		Message: "Successfully found!",
		Data:    missionResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *MissionController) Create(ctx *gin.Context) {
	req := request.CreateMissionRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "StatusBadRequest",
			Message: "Cannot parse data",
		}
		log.Printf("err: %v", err)
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	err = controller.missionService.Create(req)
	if err != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Cannot create mission",
		}
		log.Printf("err: %v", err)
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	webResponse := response.Response{
		Code:    http.StatusOK,
		Status:  "Ok",
		Message: "Successfully created!",
		Data:    nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *MissionController) UpdateNameRequest(ctx *gin.Context) {
	req := request.UpdateNameMissionRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "StatusBadRequest",
			Message: "Cannot parse data",
		}
		log.Printf("err: %v", err)
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	err = controller.missionService.UpdateNameRequest(req)
	if err != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Cannot update mission's name",
		}
		log.Printf("err: %v", err)
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	webResponse := response.Response{
		Code:    http.StatusOK,
		Status:  "Ok",
		Message: "Successfully updated!",
		Data:    nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *MissionController) UpdateNotes(ctx *gin.Context) {
	req := request.UpdateNotesRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "StatusBadRequest",
			Message: "Cannot parse data",
		}
		log.Printf("err: %v", err)
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	err = controller.missionService.UpdateNotes(req)
	if err != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Cannot update mission",
		}
		log.Printf("err: %v", err)
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	webResponse := response.Response{
		Code:    http.StatusOK,
		Status:  "Ok",
		Message: "Successfully updated!",
		Data:    nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *MissionController) UpdateTarget(ctx *gin.Context) {
	req := request.UpdateTargetRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "StatusBadRequest",
			Message: "Cannot parse data",
		}
		log.Printf("err: %v", err)
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	err = controller.missionService.UpdateTarget(req)
	if err != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Cannot update mission",
		}
		log.Printf("err: %v", err)
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	webResponse := response.Response{
		Code:    http.StatusOK,
		Status:  "Ok",
		Message: "Successfully updated!",
		Data:    nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *MissionController) Delete(ctx *gin.Context) {
	catId := ctx.Param("missionId")
	id, err := strconv.Atoi(catId)
	if err != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "StatusBadRequest",
			Message: "Cannot parse id",
		}
		log.Printf("err: %v", err)
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	err = controller.missionService.Delete(id)
	if err != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Cannot delete cat by id",
		}
		log.Printf("err: %v", err)
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	webResponse := response.Response{
		Code:    http.StatusOK,
		Status:  "Ok",
		Message: "Successfully deleted!",
		Data:    nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *MissionController) CompleteTarget(ctx *gin.Context) {
	req := request.CompleteTargetRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "StatusBadRequest",
			Message: "Cannot parse data",
		}
		log.Printf("err: %v", err)
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	err = controller.missionService.CompleteTarget(req.MissionId, req.Id)
	if err != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Cannot completed target",
		}
		log.Printf("err: %v", err)
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	webResponse := response.Response{
		Code:    http.StatusOK,
		Status:  "Ok",
		Message: "Successfully completed!",
		Data:    nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *MissionController) CompleteMission(ctx *gin.Context) {
	req := request.CompleteMissionRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "StatusBadRequest",
			Message: "Cannot parse data",
		}
		log.Printf("err: %v", err)
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	err = controller.missionService.CompleteMission(req.Id)
	if err != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Cannot completed mission",
		}
		log.Printf("err: %v", err)
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	webResponse := response.Response{
		Code:    http.StatusOK,
		Status:  "Ok",
		Message: "Successfully completed!",
		Data:    nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *MissionController) AssignCatToMission(ctx *gin.Context) {
	req := request.AssignCatToMissionRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "StatusBadRequest",
			Message: "Cannot parse data",
		}
		log.Printf("err: %v", err)
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	err = controller.missionService.AssignCatToMission(req.CatId, req.MissionId)
	if err != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Cannot completed mission",
		}
		log.Printf("err: %v", err)
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	webResponse := response.Response{
		Code:    http.StatusOK,
		Status:  "Ok",
		Message: "Successfully completed!",
		Data:    nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *MissionController) RemoveTarget(ctx *gin.Context) {
	req := request.RemoveTarget{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "StatusBadRequest",
			Message: "Cannot parse data",
		}
		log.Printf("err: %v", err)
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	err = controller.missionService.RemoveTarget(req.TargetId, req.MissionId)
	if err != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Cannot completed mission",
		}
		log.Printf("err: %v", err)
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	webResponse := response.Response{
		Code:    http.StatusOK,
		Status:  "Ok",
		Message: "Successfully completed!",
		Data:    nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *MissionController) FindMissionByCatId(ctx *gin.Context) {
	req := request.FindMissionByCatId{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "StatusBadRequest",
			Message: "Cannot parse data",
		}
		log.Printf("err: %v", err)
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	missions, err := controller.missionService.FindMissionByCatId(req.CatId)
	if err != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Cannot find missions",
		}
		log.Printf("err: %v", err)
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	webResponse := response.Response{
		Code:    http.StatusOK,
		Status:  "Ok",
		Message: "Successfully found!",
		Data:    missions,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
