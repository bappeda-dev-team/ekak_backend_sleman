package controller

import (
	"ekak_kab_sleman/helper"
	"ekak_kab_sleman/model/web"
	visimisipemda "ekak_kab_sleman/model/web/visimisi"
	"ekak_kab_sleman/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type MisiPemdaControllerImpl struct {
	MisiPemdaService service.MisiPemdaService
}

func NewMisiPemdaControllerImpl(misiPemdaService service.MisiPemdaService) *MisiPemdaControllerImpl {
	return &MisiPemdaControllerImpl{MisiPemdaService: misiPemdaService}
}

func (controller *MisiPemdaControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	misiPemdaCreateRequest := visimisipemda.MisiPemdaCreateRequest{}
	helper.ReadFromRequestBody(request, &misiPemdaCreateRequest)

	misiPemdaResponse, err := controller.MisiPemdaService.Create(request.Context(), misiPemdaCreateRequest)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebResponse{
		Code:   http.StatusCreated,
		Status: "success create misi pemda",
		Data:   misiPemdaResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *MisiPemdaControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	misiPemdaUpdateRequest := visimisipemda.MisiPemdaUpdateRequest{}
	helper.ReadFromRequestBody(request, &misiPemdaUpdateRequest)

	id := params.ByName("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	misiPemdaUpdateRequest.Id = idInt

	misiPemdaResponse, err := controller.MisiPemdaService.Update(request.Context(), misiPemdaUpdateRequest)
	if err != nil {
		helper.WriteToResponseBody(writer, err.Error())
		return
	}

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "success update misi pemda",
		Data:   misiPemdaResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)

}

func (controller *MisiPemdaControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	err = controller.MisiPemdaService.Delete(request.Context(), idInt)
	if err != nil {
		helper.WriteToResponseBody(writer, err.Error())
		return
	}

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "success delete misi pemda",
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *MisiPemdaControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	misiPemdaResponse, err := controller.MisiPemdaService.FindById(request.Context(), idInt)
	if err != nil {
		helper.WriteToResponseBody(writer, err.Error())
		return
	}

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "success find misi pemda",
		Data:   misiPemdaResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *MisiPemdaControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	tahunAwal := params.ByName("tahun_awal")
	tahunAkhir := params.ByName("tahun_akhir")
	jenisPeriode := params.ByName("jenis_periode")

	misiPemdaResponses, err := controller.MisiPemdaService.FindAll(request.Context(), tahunAwal, tahunAkhir, jenisPeriode)
	if err != nil {
		helper.WriteToResponseBody(writer, err.Error())
		return
	}

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "success find all misi pemda",
		Data:   misiPemdaResponses,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *MisiPemdaControllerImpl) FindByIdVisi(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	idVisistr := params.ByName("id_visi")

	idVisi, err := strconv.Atoi(idVisistr)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}
	misiPemdaResponses, err := controller.MisiPemdaService.FindByIdVisi(request.Context(), idVisi)
	if err != nil {
		helper.WriteToResponseBody(writer, err.Error())
		return
	}

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "success find all misi pemda by id visi",
		Data:   misiPemdaResponses,
	}
	helper.WriteToResponseBody(writer, webResponse)
}
