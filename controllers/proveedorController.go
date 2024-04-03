package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/santihill/crud_go/initializers"
	"github.com/santihill/crud_go/models"
	"gorm.io/gorm"
)

func IndexProveedores(c *gin.Context) {
	//Get the posts
	var posts []models.Proveedor
	result := initializers.DB.Find(&posts)

	//Respond with them
	c.JSON(200, gin.H{
		"Rows Affected": result.RowsAffected,
		"posts":         posts,
	})

}

/*
func SearchByEmpresaProveedor(c *gin.Context) {
	//Get id from URL
	empresa := c.Param("empresa")

	//Get the posts
	var proveedor models.Proveedor
	result := initializers.DB.First(&proveedor, empresa)

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
		"valor enviado": empresa,
		"empresa":       proveedor,
	})
}
*/

func SearchByIdProveedor(c *gin.Context) {
	//Get id from URL
	id := c.Param("id")

	//Get the posts
	var proveedor models.Proveedor
	result := initializers.DB.First(&proveedor, id)

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
		"empresa":       proveedor,
	})
}

func SetProveedores(c *gin.Context) {
	//Get the proveedor and pass off req body
	var body struct {
		Nombre    string
		Empresa   string
		NumFiscal string
		Direccion string
		Telefono  string
		Email     string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	//Create the user
	proveedor := models.Proveedor{Nombre: body.Nombre, Empresa: body.Empresa, NumFiscal: body.NumFiscal, Direccion: body.Direccion, Telefono: body.Telefono, Email: body.Email}
	var proveedor2 models.Proveedor
	// Get first matched record
	initializers.DB.Where("empresa = ?", proveedor.Empresa).First(&proveedor2)
	// SELECT * FROM proovedor2 WHERE empresa = 'empresa' ORDER BY id LIMIT 1;
	if proveedor.Empresa == proveedor2.Empresa {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Proveedor already exist",
		})
		return
	}

	result := initializers.DB.Create(&proveedor) // pass pointer of data to Create
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create proveedor",
		})
		return
	}

	//Respond
	c.JSON(http.StatusOK, gin.H{
		"Success":    "Proveedor has being created",
		"ID":         proveedor.ID,
		"Nombre:":    body.Nombre,
		"Empresa:":   body.Empresa,
		"CUIT:":      body.NumFiscal,
		"Direccion:": body.Direccion,
		"Telefono:":  body.Telefono,
		"Email:":     body.Email,
	})

}

func UpdateProveedor(c *gin.Context) {
	//Get the id off the url
	id := c.Param("id")

	var body struct {
		Nombre    string
		Empresa   string
		NumFiscal string
		Direccion string
		Telefono  string
		Email     string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	var proveedor2 models.Proveedor
	// Get first matched record
	initializers.DB.Where("empresa = ?", body.Empresa).First(&proveedor2)
	// SELECT * FROM proovedor2 WHERE empresa = 'empresa' ORDER BY id LIMIT 1;
	if body.Empresa == proveedor2.Empresa {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Proveedor already exist",
		})
		return
	}

	//Find the proveedor were updating
	var proveedor models.Proveedor
	initializers.DB.First(&proveedor, id)

	//Update it
	initializers.DB.Model(&proveedor).Updates(models.Proveedor{
		Nombre:    body.Nombre,
		Empresa:   body.Empresa,
		NumFiscal: body.NumFiscal,
		Direccion: body.Direccion,
		Telefono:  body.Telefono,
		Email:     body.Email,
	})

	//Respond with it
	c.JSON(200, gin.H{
		"Proveedor actualizado": proveedor,
	})
}

func ProveedorDelete(c *gin.Context) {
	//Get the id off url
	id := c.Param("id")

	//Delete the proveedor
	initializers.DB.Delete(&models.Proveedor{}, id)
	//Respond

	c.Status(200)
}
