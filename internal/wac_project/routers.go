/*
 * Waiting List Api
 *
 * Ambulance Waiting List management for Web-In-Cloud system
 *
 * API version: 1.0.0
 * Contact: xcernak@stuba.sk
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package wac_project

import (
    "github.com/gin-gonic/gin"
)

func AddRoutes(engine *gin.Engine) {
  group := engine.Group("/api")
  
  {
    api := newAppointmentsListAPI()
    api.addRoutes(group)
  }
  
}