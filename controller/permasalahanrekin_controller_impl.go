package controller

import (
	"ekak_kab_sleman/helper"
	"ekak_kab_sleman/model/web"
	"ekak_kab_sleman/model/web/permasalahan"
	"ekak_kab_sleman/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type PermasalahanRekinControllerImpl struct {
	PermasalahanRekinService service.PermasalahanRekinService
}

func NewPermasalahanRekinControllerImpl(permasalahanRekinService service.PermasalahanRekinService) *PermasalahanRekinControllerImpl {
	return &PermasalahanRekinControllerImpl{
		PermasalahanRekinService: permasalahanRekinService,
	}
}

func (controller *PermasalahanRekinControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	permasalahanRekinCreateRequest := permasalahan.PermasalahanRekinCreateRequest{}
	helper.ReadFromRequestBody(request, &permasalahanRekinCreateRequest)

	permasalahanRekinResponse, err := controller.PermasalahanRekinService.Create(request.Context(), permasalahanRekinCreateRequest)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   400,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   permasalahanRekinResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *PermasalahanRekinControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	permasalahanRekinUpdateRequest := permasalahan.PermasalahanRekinUpdateRequest{}
	helper.ReadFromRequestBody(request, &permasalahanRekinUpdateRequest)

	id := params.ByName("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   400,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}
	permasalahanRekinUpdateRequest.Id = idInt

	permasalahanRekinResponse, err := controller.PermasalahanRekinService.Update(request.Context(), permasalahanRekinUpdateRequest)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   400,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   permasalahanRekinResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)

}

func (controller *PermasalahanRekinControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   400,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	err = controller.PermasalahanRekinService.Delete(request.Context(), idInt)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   400,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "success delete permasalahan rekin",
		Data:   nil,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *PermasalahanRekinControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   400,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	permasalahanRekinResponse, err := controller.PermasalahanRekinService.FindById(request.Context(), idInt)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   400,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "success get data permasalahan rekin by id",
		Data:   permasalahanRekinResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *PermasalahanRekinControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	rekinId := params.ByName("rekinId")

	var rekinIdPtr *string
	if rekinId != "" {
		rekinIdPtr = &rekinId
	}

	permasalahanRekinResponse, err := controller.PermasalahanRekinService.FindAll(request.Context(), rekinIdPtr)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   400,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "success get all data permasalahan rekin",
		Data:   permasalahanRekinResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}
