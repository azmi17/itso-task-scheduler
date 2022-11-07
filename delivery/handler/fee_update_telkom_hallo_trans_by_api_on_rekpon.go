package handler

import (
	"itso-task-scheduler/delivery/handler/httpio"
	"itso-task-scheduler/entities"
	"itso-task-scheduler/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FeeUpdateTelkomHalloByAPI(ctx *gin.Context) {

	httpio := httpio.NewRequestIO(ctx)
	httpio.Recv()

	usecase := usecase.NewRekponSchedulerUsecase()
	data := usecase.RekponUpdateFeeOnTelkomHalloTrans()

	resp := entities.SchedulerResponse{}

	if data != nil {
		entities.PrintLog(data.Error())
		resp.ResponseCode = "0000"
		resp.ResponseMessage = "Tidak ada data fee transaksi (telkom, hallo) untuk di update"

	} else {
		resp.ResponseCode = "0000"
		resp.ResponseMessage = "Update data fee transaksi (telkom, hallo) sukses"
	}
	httpio.Response(http.StatusOK, resp)

}
