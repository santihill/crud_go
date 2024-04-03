package main

import (
	"github.com/gin-gonic/gin"
	"github.com/santihill/crud_go/controllers"
	"github.com/santihill/crud_go/initializers"
	"github.com/santihill/crud_go/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

// albun Estructura de albun
//type album struct {
//	ID     string `json:"id"`
//	Title  string `json:"title"`
//	Artist string `json:"artist"`
//	Year   int    `json:"year"`
//}

// Lista de albuns
//var albums = []album{
//	{ID: "1", Title: "Familia", Artist: "Camila Cabello", Year: 2022},
//	{ID: "2", Title: "21", Artist: "Adele", Year: 2011},
//	{ID: "3", Title: "The Eminem Show", Artist: "Eminem", Year: 2002},
//	{ID: "4", Title: "Meteora", Artist: "Linkin Park", Year: 2003},
//	{ID: "5", Title: "25", Artist: "Adele", Year: 2015},
//}

// Obtener lista de albums
//func getAlbums(c *gin.Context) {
//	c.IndentedJSON(http.StatusOK, albums)
//}

// Agregar un albums
//func postAlbums(c *gin.Context) {
//	var newAlbum album
//
//	if err := c.BindJSON(&newAlbum); err != nil {
//		return
//	}
//
//	albums = append(albums, newAlbum)
//	c.IndentedJSON(http.StatusCreated, newAlbum)
//}

// Obtener po id un albuns
//func getAlbumByID(c *gin.Context) {
//	id := c.Param("id")
//
//	for _, a := range albums {
//		if a.ID == id {
//			c.IndentedJSON(http.StatusOK, a)
//			return
//		}
//	}
//
//	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Albun no encontrado"})
//}

func main() {
	//	router := gin.Default()
	//	router.GET("/albums", getAlbums)
	//	router.POST("/albums", postAlbums)
	//	router.GET("/albums/:id", getAlbumByID)

	//router.Run("localhost:8080")
	r := gin.Default()
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	r.POST("/proveedores", controllers.SetProveedores)
	r.GET("/proveedores", controllers.IndexProveedores)
	r.GET("/proveedores/:id", controllers.SearchByIdProveedor)
	r.PUT("/proveedores/:id", controllers.UpdateProveedor)
	r.DELETE("/proveedores/:id", controllers.ProveedorDelete)

	r.POST("/empleado", controllers.SetEmpleado)
	r.GET("/empleado", controllers.IndexEmpleados)
	r.GET("/empleado/:id", controllers.SearchByIdEmpleado)
	r.PUT("/empleado/:id", controllers.UpdateEmpleado)
	r.DELETE("/empleado/:id", controllers.EmpleadoDelete)

	r.POST("/propietario", controllers.SetPropietario)
	r.GET("/propietario", controllers.IndexPropietarios)
	r.GET("/propietario/:id", controllers.SearchByIdPropietario)
	r.PUT("/propietario/:id", controllers.UpdatePropietario)
	r.DELETE("/propietario/:id", controllers.PropietarioDelete)

	//r.GET("/empleados", getEmpleados)
	//r.GET("/propietarios", getPropietarios)
	r.POST("/posts", controllers.PostsCreate)
	r.PUT("/posts/:id", controllers.PostsUpdate)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostsShow)
	r.DELETE("/posts/:id", controllers.PostsDelete)

	r.Run()
}
