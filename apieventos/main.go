package main

import (
    "github.com/gin-gonic/gin"
    "apieventos/routes"
    // "apieventos/services"

)

func main() {

    // services.ConnectDB()
    router := gin.Default()

    // Configurar rutas
    routes.SetupRoutes(router)

    // Iniciar servidor
    router.Run(":8080")
}
