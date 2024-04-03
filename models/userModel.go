package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
}

type Proveedor struct {
	gorm.Model
	Nombre    string
	Empresa   string
	NumFiscal string
	Direccion string
	Telefono  string
	Email     string
}

type Empleado struct {
	gorm.Model
	Nombre    string
	Empresa   string
	NumFiscal string
	Direccion string
	Telefono  string
	Email     string
	Trabajo   string
}

type Propietario struct {
	gorm.Model
	Nombre       string
	Departamento string
	NumFiscal    string
	Direccion    string
	Telefono     string
	Email        string
}

type Post struct {
	gorm.Model
	Tittle string
	Body   string
}
