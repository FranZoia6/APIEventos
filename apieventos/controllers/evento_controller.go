package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "strconv"
    
    "apieventos/models"
    "apieventos/services"
)



type SuscripcionRequest struct {
    EventoID int    `json:"eventoID" binding:"required"`
    Email    string `json:"email" binding:"required,email"`
}

func GetEvento(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
        return
    }

    evento := services.GetEventoByID(id)
    if evento == nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Evento no encontrado o no publicado"})
        return
    }

    c.JSON(http.StatusOK, evento)
}

func CreateEvent(c *gin.Context) {
    var newEvent models.Event

    if err := c.ShouldBindJSON(&newEvent); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    services.SetEvent(newEvent)

    c.JSON(http.StatusCreated, newEvent)
}

func UpdateEvent(c *gin.Context) {
    var event models.Event

    if err := c.ShouldBindJSON(&event); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    
    if event.ID == 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID es requerido"})
        return
    }

    err := services.UpdateEvent( event)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Evento actualizado correctamente"})
}


func SuscribirUsuario(c *gin.Context) {
    var req SuscripcionRequest

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
        return
    }


    usuarioID, err := services.GetUsuarioIDByEmail(req.Email)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
        return
    }

    err = services.SuscribirUsuarioAEvento(req.EventoID, usuarioID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Usuario suscrito al evento correctamente"})
}


