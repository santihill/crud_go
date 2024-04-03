package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/santihill/crud_go/initializers"
	"github.com/santihill/crud_go/models"
	"gorm.io/gorm"
)

func IndexEmpleados(c *gin.Context) {
	//Get the posts
	var empleado []models.Empleado
	result := initializers.DB.Find(&empleado)

	//Respond with them
	c.JSON(200, gin.H{
		"Rows Affected": result.RowsAffected,
		"posts":         empleado,
	})

}

func SearchByIdEmpleado(c *gin.Context) {
	//Get id from URL
	id := c.Param("id")

	//Get the posts
	var empleado models.Empleado
	result := initializers.DB.First(&empleado, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(400, gin.H{
			"error": "Empresa not found",
		})
		return
	}

	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": "Empresa not found",
		})
		return
	}

	//Respond with them
	c.JSON(200, gin.H{
		"Rows affected": result.RowsAffected,
		"error":         result.Error,
		"valor enviado": id,
		"empresa":       empleado,
	})
}

func SetEmpleado(c *gin.Context) {
	//Get the empleado and pass off req body
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

	result := initializers.DB.Create(&empleado) // pass pointer of data to Create
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create empleado",
		})
		return
	}

	//Respond
	c.JSON(http.StatusOK, gin.H{
		"Success":    "Empleado has being created",
		"ID":         empleado.ID,
		"Nombre:":    body.Nombre,
		"Empresa:":   body.Empresa,
		"CUIT:":      body.NumFiscal,
		"Direccion:": body.Direccion,
		"Telefono:":  body.Telefono,
		"Email:":     body.Email,
		"Trabajo":    body.Trabajo,
	})

}

func UpdateEmpleado(c *gin.Context) {
	//Get the id off the url
	id := c.Param("id")

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

	var empleado2 models.Empleado
	// Get first matched record
	initializers.DB.Where("numfiscal = ?", body.Email).First(&empleado2)
	// SELECT * FROM empleado2 WHERE empresa = 'empresa' ORDER BY id LIMIT 1;
	if body.Email == empleado2.Email {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Empleado already exist",
		})
		return
	}

	//Find the empleado were updating
	var empleado models.Empleado
	initializers.DB.First(&empleado, id)

	//Update it
	initializers.DB.Model(&empleado).Updates(models.Empleado{
		Nombre:    body.Nombre,
		Empresa:   body.Empresa,
		NumFiscal: body.NumFiscal,
		Direccion: body.Direccion,
		Telefono:  body.Telefono,
		Email:     body.Email,
		Trabajo:   body.Trabajo,
	})

	//Respond with it
	c.JSON(200, gin.H{
		"Empleado actualizado": empleado,
	})
}

func EmpleadoDelete(c *gin.Context) {
	//Get the id off url
	id := c.Param("id")

	//Delete the empleado
	initializers.DB.Delete(&models.Empleado{}, id)
	//Respond

	c.Status(200)
}
