package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/santihill/crud_go/initializers"
	"github.com/santihill/crud_go/models"
)

func SearchRepeat(c *gin.Context) {
	var body struct {
		Nombre    string
		Empresa   string
		NumFiscal string
		Direccion string
		Telefono  string
		Email     string
		Trabajo   string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	//Create the user
	empleado := models.Empleado{Nombre: body.Nombre, Empresa: body.Empresa, NumFiscal: body.NumFiscal, Direccion: body.Direccion, Telefono: body.Telefono, Email: body.Email, Trabajo: body.Trabajo}
	var empleado2 models.Empleado
	// Get first matched record
	initializers.DB.Where("email = ?", empleado.Email).First(&empleado2)
	// SELECT * FROM empleado2 WHERE CUIT = 'CUIT' ORDER BY id LIMIT 1;
	if empleado.Email == empleado2.Email {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Empleado already exist",
		})
		return
	}

	switch {
	case empleado.Email == empleado2.Email:
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Empleado already exist",
		})
	case empleado.Nombre == empleado2.Nombre:
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Empleado already exist",
		})
	case empleado.NumFiscal == empleado2.NumFiscal:
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Empleado already exist",
		})
	}
}
