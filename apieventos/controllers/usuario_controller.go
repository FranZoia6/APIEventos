package controllers

import (
	"apieventos/models"
	"apieventos/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var mySigningKey = []byte("your-secret-key")

// Handler para obtener eventos
func GetEventsPublished(c *gin.Context) {
    eventosPublicados := services.GetEventsPublished()
    if eventosPublicados == nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Evento no encontrado o no publicado"})
        return
    }
    c.JSON(http.StatusOK, eventosPublicados)
}

func GetEvents(c *gin.Context) {
    eventosPublicados := services.GetEvents()
    if eventosPublicados == nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Evento no encontrado o no publicado"})
        return
    }
    c.JSON(http.StatusOK, eventosPublicados)
}


func Login(c *gin.Context) {
    var loginDetails map[string]string
    if err := c.BindJSON(&loginDetails); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
        return
    }

    adminname := loginDetails["adminname"]
    password := loginDetails["password"]

    admin, valid := services.ValidateAdmin(adminname, password)
    if !valid {
        c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
        return
    }

    claims := jwt.MapClaims{
        "role": admin.Role,
        "exp":  time.Now().Add(time.Hour * 1).Unix(), // Expira en 1 hora
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(mySigningKey)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not generate token"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func RegisterUser(c *gin.Context) {
    var newUser models.User


    if err := c.ShouldBindJSON(&newUser); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
        return
    }


    err := services.SetUser(newUser)
    if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
    }

    c.JSON(http.StatusCreated, newUser)
}




