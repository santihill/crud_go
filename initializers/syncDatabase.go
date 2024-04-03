package initializers

import "github.com/santihill/crud_go/models"

func SyncDatabase() {
	// Migrate the schema
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Proveedor{})
	DB.AutoMigrate(&models.Empleado{})
	DB.AutoMigrate(&models.Propietario{})
}

// TODO: agregar un método para insertar proveedores a la BD
