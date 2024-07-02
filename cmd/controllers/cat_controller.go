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

type CatController struct {
	catService services.CatService
}

func NewCatController(service services.CatService) *CatController {
	return &CatController{catService: service}
}

func (controller *CatController) FindAll(ctx *gin.Context) {
	catsResponse, err := controller.catService.FindAll()
	if err != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Cannot find cats",
		}
		log.Printf("err: %v", err)
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	webResponse := response.Response{
		Code:    http.StatusOK,
		Status:  "Ok",
		Message: "Successfully found!",
		Data:    catsResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *CatController) FindById(ctx *gin.Context) {
	catId := ctx.Param("catId")
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

	catResponse, err := controller.catService.FindById(uint32(id))
	if err != nil || catResponse.Id == 0 {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Cannot find cat by id",
		}
		log.Printf("err: %v", err)
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	webResponse := response.Response{
		Code:    http.StatusOK,
		Status:  "Ok",
		Message: "Successfully found!",
		Data:    catResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *CatController) Create(ctx *gin.Context) {
	ccr := request.CreateCatRequest{}
	err := ctx.ShouldBindJSON(&ccr)
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

	err = controller.catService.Create(ccr)
	if err != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Cannot create cat",
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

func (controller *CatController) Update(ctx *gin.Context) {
	catId := ctx.Param("catId")
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

	ucr := request.UpdateCatRequest{Id: uint32(id)}
	err = ctx.ShouldBindJSON(&ucr)
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

	err = controller.catService.Update(ucr)
	if err != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Cannot update cat by id",
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

func (controller *CatController) Delete(ctx *gin.Context) {
	catId := ctx.Param("catId")
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

	err = controller.catService.Delete(uint32(id))
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
