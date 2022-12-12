package handler

import (
	"itso-task-scheduler/delivery/handler/httpio"
	"itso-task-scheduler/entities"
	"itso-task-scheduler/entities/err"
	"itso-task-scheduler/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RepostingAllByApi(ctx *gin.Context) {

	httpio := httpio.NewRequestIO(ctx)
	httpio.Recv()

	usecase := usecase.NewApexSchedulerUsecase()
	er := usecase.RepostingSaldoApexByScheduler()

	// go RepostingSchedulerRepoObserver()

	resp := entities.SchedulerResponse{}
	if er != nil {
		if er == err.NoRecord {
			resp.ResponseCode = "1111"
			resp.ResponseMessage = er.Error()
		} else {
			entities.PrintError(er.Error())
			httpio.ResponseString(http.StatusInternalServerError, "internal service error", nil)
			return
		}
	} else {
		resp.ResponseCode = "0000"
		resp.ResponseMessage = "Reposting saldo succeeded"
	}

	httpio.Response(http.StatusOK, resp)
}
