/*
 * Copyright (c) 2022 Randy Ardiansyah https://github.com/randyardiansyah25/<repo>
 *
 * Created Date: Wednesday, 16/03/2022, 10:32:08
 * Author: Randy Ardiansyah
 *
 * Filename: /home/Documents/workspace/go/src/router-template/delivery/router/registry.go
 * Project : /home/Documents/workspace/go/src/router-template/delivery/router
 *
 * HISTORY:
 * Date                  	By                 	Comments
 * ----------------------	-------------------	--------------------------------------------------------------------------------------------------------------------
 */

package router

import (
	"itso-task-scheduler/delivery/handler"

	"github.com/gin-gonic/gin"
)

func RegisterHandler(router *gin.Engine) {

	// API Versioning:
	apiv1 := router.Group("api/v1/")

	// API Endpoint:
	apiv1.GET("/version", handler.AppInfo)
	apiv1.GET("/fee-update", handler.FeeUpdateTelkomHalloByAPI)
	apiv1.POST("/repostings/all", handler.RepostingAllByApi)

}
