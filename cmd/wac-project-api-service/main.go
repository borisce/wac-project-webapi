package main

import (
    "log"
    "os"
    "strings"
    "github.com/gin-gonic/gin"
    "github.com/borisce/wac-project-webapi/api"
    "github.com/borisce/wac-project-webapi/internal/wac_project"
)

func main() {
    log.Printf("Server started")
    port := os.Getenv("WAC_PROJECT_API_PORT")
    if port == "" {
        port = "8080"
    }
    environment := os.Getenv("WAC_PROJECT_API_ENVIRONMENT")
    if !strings.EqualFold(environment, "production") { // case insensitive comparison
        gin.SetMode(gin.DebugMode)
    }
    engine := gin.New()
    engine.Use(gin.Recovery())
    // request routings
    wac_project.AddRoutes(engine)
    engine.GET("/openapi", api.HandleOpenApi)
    engine.Run(":" + port)
}