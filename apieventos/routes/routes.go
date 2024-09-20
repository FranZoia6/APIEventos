package routes

import (
    "apieventos/controllers"
    "github.com/gin-gonic/gin"
    "apieventos/middleware"
    "github.com/gin-contrib/cors"
    "time"
)



func SetupRoutes(router *gin.Engine) {


   router.Use(cors.New(cors.Config{
    AllowOrigins:     []string{"*"}, 
    AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
    AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
    ExposeHeaders:    []string{"Content-Length"},
    AllowCredentials: true,
    MaxAge: 12 * time.Hour,
    }))


    router.POST("/login", controllers.Login)

    
    authorized := router.Group("/")
    authorized.Use(middleware.AuthMiddleware())
    {
        authorized.POST("/addEvent", controllers.CreateEvent) 
        authorized.GET("/eventosAdmin", controllers.GetEvents)
        authorized.PUT("/updateEvent", controllers.UpdateEvent)
    }
    
    
    router.GET("/eventos", controllers.GetEventsPublished)
    router.GET("/evento/:id", controllers.GetEvento) 
    router.POST("/eventos/suscribir", controllers.SuscribirUsuario)
    router.POST("/register", controllers.RegisterUser)
}



