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
	atr := request.AddTargetRequest{}
	err := ctx.ShouldBindJSON(&atr)
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

	err = controller.missionService.AddTarget(atr)
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

	missionResponse, err := controller.missionService.FindById(uint32(id))
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
	cmr := request.CreateMissionRequest{}
	err := ctx.ShouldBindJSON(&cmr)
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

	err = controller.missionService.Create(cmr)
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
	umr := request.UpdateNameMissionRequest{}
	err := ctx.ShouldBindJSON(&umr)
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

	err = controller.missionService.UpdateNameRequest(umr)
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
	umr := request.UpdateNotesRequest{}
	err := ctx.ShouldBindJSON(&umr)
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

	err = controller.missionService.UpdateNotes(umr)
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
	utr := request.UpdateTargetRequest{}
	err := ctx.ShouldBindJSON(&utr)
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

	err = controller.missionService.UpdateTarget(utr)
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

	err = controller.missionService.Delete(uint32(id))
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
	ctr := request.CompleteTargetRequest{}
	err := ctx.ShouldBindJSON(&ctr)
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

	err = controller.missionService.CompleteTarget(uint32(ctr.MissionId), ctr.Id)
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
	cmr := request.CompleteMissionRequest{}
	err := ctx.ShouldBindJSON(&cmr)
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

	err = controller.missionService.CompleteMission(uint32(cmr.Id))
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
	acmr := request.AssignCatToMissionRequest{}
	err := ctx.ShouldBindJSON(&acmr)
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

	err = controller.missionService.AssignCatToMission(acmr.CatId, acmr.MissionId)
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
	rtr := request.RemoveTarget{}
	err := ctx.ShouldBindJSON(&rtr)
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

	err = controller.missionService.RemoveTarget(rtr.TargetId, uint32(rtr.MissionId))
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
	fmcr := request.FindMissionByCatId{}
	err := ctx.ShouldBindJSON(&fmcr)
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

	missions, err := controller.missionService.FindMissionByCatId(fmcr.CatId)
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
