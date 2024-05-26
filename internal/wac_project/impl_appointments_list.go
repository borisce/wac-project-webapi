package wac_project

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

func (this *implAppointmentsListAPI) CreateAppointment(ctx *gin.Context) {
	ctx.AbortWithStatus(http.StatusNotImplemented)
}
	
//DeleteAppointment - Deletes an appointment
func (this *implAppointmentsListAPI) DeleteAppointment(ctx *gin.Context) {
  	ctx.AbortWithStatus(http.StatusNotImplemented)
}
	
//GetAppointmentsList - Provides the appointments list
func (this *implAppointmentsListAPI) GetAppointmentsList(ctx *gin.Context) {
 	ctx.AbortWithStatus(http.StatusNotImplemented)
}
	
//UpdateAppointment - Updates an appointment
func (this *implAppointmentsListAPI) UpdateAppointment(ctx *gin.Context) {
 	ctx.AbortWithStatus(http.StatusNotImplemented)
}