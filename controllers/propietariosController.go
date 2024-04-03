package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/santihill/crud_go/initializers"
	"github.com/santihill/crud_go/models"
	"gorm.io/gorm"
)

func IndexPropietarios(c *gin.Context) {
	//Get the posts
	var propietario []models.Propietario
	result := initializers.DB.Find(&propietario)

	//Respond with them
	c.JSON(200, gin.H{
		"Rows Affected": result.RowsAffected,
		"propietario":   propietario,
	})

}

func SearchByIdPropietario(c *gin.Context) {
	//Get id from URL
	id := c.Param("id")

	//Get the posts
	var propietario models.Propietario
	result := initializers.DB.First(&propietario, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(400, gin.H{
			"error": "Departamento not found",
		})
		return
	}

	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": "Departamento not found",
		})
		return
	}

	//Respond with them
	c.JSON(200, gin.H{
		"Rows affected": result.RowsAffected,
		"error":         result.Error,
		"valor enviado": id,
		"propietario":   propietario,
	})
}

func SetPropietario(c *gin.Context) {
	//Get the propietario and pass off req body
	var body struct {
		Nombre       string
		Departamento string
		NumFiscal    string
		Direccion    string
		Telefono     string
		Email        string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	//Create the user
	propietario := models.Propietario{Nombre: body.Nombre, Departamento: body.Departamento, NumFiscal: body.NumFiscal, Direccion: body.Direccion, Telefono: body.Telefono, Email: body.Email}
	var propietario2 models.Propietario
	// Get first matched record
	initializers.DB.Where("departamento = ?", body.Departamento).First(&propietario2)
	// SELECT * FROM proovedor2 WHERE departamento = 'departamento' ORDER BY id LIMIT 1;
	if propietario.Departamento == propietario2.Departamento {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Propietario already exist",
		})
		return
	}

	result := initializers.DB.Create(&propietario) // pass pointer of data to Create
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create propietario",
		})
		return
	}

	//Respond
	c.JSON(http.StatusOK, gin.H{
		"Success":       "Propietario has being created",
		"ID":            propietario.ID,
		"Nombre:":       body.Nombre,
		"Departamento:": body.Departamento,
		"CUIT:":         body.NumFiscal,
		"Direccion:":    body.Direccion,
		"Telefono:":     body.Telefono,
		"Email:":        body.Email,
	})

}

func UpdatePropietario(c *gin.Context) {
	//Get the id off the url
	id := c.Param("id")

	var body struct {
		Nombre       string
		Departamento string
		NumFiscal    string
		Direccion    string
		Telefono     string
		Email        string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	var propietario2 models.Propietario
	// Get first matched record
	initializers.DB.Where("departamento = ?", body.Departamento).First(&propietario2)
	// SELECT * FROM proovedor2 WHERE departamento = 'departamento' ORDER BY id LIMIT 1;
	if body.Departamento == propietario2.Departamento {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Propietario already exist",
		})
		return
	}

	//Find the propietario were updating
	var propietario models.Propietario
	initializers.DB.First(&propietario, id)

	//Update it
	initializers.DB.Model(&propietario).Updates(models.Propietario{
		Nombre:       body.Nombre,
		Departamento: body.Departamento,
		NumFiscal:    body.NumFiscal,
		Direccion:    body.Direccion,
		Telefono:     body.Telefono,
		Email:        body.Email,
	})

	//Respond with it
	c.JSON(200, gin.H{
		"Propietario actualizado": propietario,
	})
}

func PropietarioDelete(c *gin.Context) {
	//Get the id off url
	id := c.Param("id")

	//Delete the propietario
	initializers.DB.Delete(&models.Propietario{}, id)
	//Respond

	c.Status(200)
}
